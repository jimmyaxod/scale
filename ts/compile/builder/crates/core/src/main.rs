#![allow(dead_code)]
#![allow(unused_imports)]

mod engine;

#[path = "utils/utils.rs"]
mod utils;

use quickjs_wasm_sys::{
  ext_js_exception, ext_js_null, ext_js_undefined, size_t as JS_size_t, JSCFunctionData,
  JSContext, JSValue, JS_Eval, JS_FreeCString, JS_GetGlobalObject, JS_NewArray, JS_NewBigInt64,
  JS_NewBool_Ext, JS_NewCFunctionData, JS_NewContext, JS_NewFloat64_Ext, JS_NewInt32_Ext,
  JS_NewInt64_Ext, JS_NewObject, JS_NewRuntime, JS_NewStringLen, JS_NewUint32_Ext,
  JS_ToCStringLen2, JS_EVAL_TYPE_GLOBAL, JS_GetPropertyStr, JS_GetPropertyUint32, 
  JS_DefinePropertyValueUint32, JS_PROP_C_W_E,
};
use std::os::raw::{c_char, c_int, c_void};

use std::ffi::CString;


use lazy_static::lazy_static;
use std::sync::Mutex;

use quickjs_wasm_rs::{Context, Value};

use once_cell::sync::OnceCell;
use std::io::{self, Read};

use utils::{pack_uint32, unpack_uint32};

use std::io::{Cursor, Write};

use std::mem;
use std::mem::{MaybeUninit};
extern crate wee_alloc;

lazy_static! {
  pub static ref PTR: Mutex<u32> = Mutex::new(0);
  pub static ref LEN: Mutex<u32> = Mutex::new(0);
  pub static ref READ_BUFFER: Mutex<Vec<u8>> = Mutex::new(Vec::with_capacity(0));

  pub static ref NEXT_PTR: Mutex<u32> = Mutex::new(0);
  pub static ref NEXT_LEN: Mutex<u32> = Mutex::new(0);
  pub static ref NEXT_READ_BUFFER: Mutex<Vec<u8>> = Mutex::new(Vec::with_capacity(0));
}

#[cfg(not(test))]
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;

static mut JS_CONTEXT: OnceCell<Context> = OnceCell::new();
static mut ENTRYPOINT: (OnceCell<Value>, OnceCell<Value>) = (OnceCell::new(), OnceCell::new());
static mut ENTRYPOINT2: (OnceCell<Value>, OnceCell<Value>) = (OnceCell::new(), OnceCell::new());
static mut ENTRYPOINT3: (OnceCell<Value>, OnceCell<Value>) = (OnceCell::new(), OnceCell::new());
static SCRIPT_NAME: &str = "script.js";

// TODO
//
// AOT validations:
//  1. Ensure that the required exports are present
//  2. If not present just evaluate the top level statement (?)


// Convert a vec to a JSValue array of bytes
fn vec_to_js(context: *mut JSContext, v: Vec<u8>) -> JSValue {
  unsafe {
    let array = JS_NewArray(context);
    let mut index: u32 = 0;
    for val in v.iter() {
      let jval = JS_NewInt32_Ext(context, i32::from(*val));
      JS_DefinePropertyValueUint32(
          context,
          array,
          index,
          jval,
          JS_PROP_C_W_E as i32,
      );
      index += 1;
    }
    return array;
  }
}

fn js_to_vec(context: *mut JSContext, v: JSValue) -> Vec<u8> {
  let mut cursor = Cursor::new(Vec::new());

  // The input (jsval2) is expected to be an array of bytes.
  let cstring_key = CString::new("length").unwrap();
  let len = unsafe { JS_GetPropertyStr(context, v, cstring_key.as_ptr()) } as u32;

  for i in 0..len {
    let v = unsafe { JS_GetPropertyUint32(context, v, i) } as u8;
    let nval:&[u8] = &[v];
    cursor.write(nval);
  }

  let mut vec = cursor.into_inner();
  vec.shrink_to_fit();
  return vec;
}

// The function env.next exported by the host
#[link(wasm_import_module = "env")]
extern "C" {
    #[link_name = "next"]
    fn _next(ptr: u32, size: u32) -> u64;
}

// Wrap the exported next function so it can be called from js
fn nextwrap(context: *mut JSContext, jsval1: JSValue, int1: c_int, jsval2: *mut JSValue, int2: c_int) -> JSValue {
  unsafe {
    let vec = js_to_vec(context, *jsval2);
    let (ptr, len) = setNextBuffer(vec);

    let packed = _next(ptr, len);
    let (ptr, len) = unpack_uint32(packed);    
    let rvec = Vec::from_raw_parts(ptr as *mut u8, len as usize, len as usize);
    return vec_to_js(context, rvec);
  }
}

fn setBuffer(v: Vec<u8>) -> (u32, u32) {
  let ptr = v.as_ptr() as u32;
  let len = v.len() as u32;

  *READ_BUFFER.lock().unwrap() = v;
  *PTR.lock().unwrap() = ptr;
  *LEN.lock().unwrap() = len;
  return (ptr, len);
}

fn setNextBuffer(v: Vec<u8>) -> (u32, u32) {
  let ptr = v.as_ptr() as u32;
  let len = v.len() as u32;

  *NEXT_READ_BUFFER.lock().unwrap() = v;
  *NEXT_PTR.lock().unwrap() = ptr;
  *NEXT_LEN.lock().unwrap() = len;
  return (ptr, len);
}

#[export_name = "wizer.initialize"]
pub extern "C" fn init() {
    unsafe {
        let mut context = Context::default();
        context
            .register_globals(io::stderr(), io::stderr())
            .unwrap();

        let mut contents = String::new();
        io::stdin().read_to_string(&mut contents).unwrap();

        let _ = context.eval_global(SCRIPT_NAME, &contents).unwrap();
        let global = context.global_object().unwrap();
        let exports = global.get_property("Exports").unwrap();
        let main = exports.get_property("main").unwrap();

        // Setup a function called next() in the global_object
        let cb = context.new_callback(nextwrap).unwrap();
        global.set_property("scale_fn_next", cb);    // Callback to the 'next' host export.

        JS_CONTEXT.set(context).unwrap();
        ENTRYPOINT.0.set(exports).unwrap();
        ENTRYPOINT.1.set(main).unwrap();

        let exports2 = global.get_property("Exports").unwrap();
        let runfn = exports2.get_property("run").unwrap();
        ENTRYPOINT2.0.set(exports2).unwrap();
        ENTRYPOINT2.1.set(runfn).unwrap();

        let exports3 = global.get_property("Exports").unwrap();
        let resizefn = exports3.get_property("resize").unwrap();
        ENTRYPOINT3.0.set(exports3).unwrap();
        ENTRYPOINT3.1.set(resizefn).unwrap();
    }
}

fn main() {
    unsafe {
        let context = JS_CONTEXT.get().unwrap();
        let shopify = ENTRYPOINT.0.get().unwrap();
        let main = ENTRYPOINT.1.get().unwrap();

        let array = context.array_value().unwrap();       // Just send in empty array for now...
        let output_value = main.call(shopify, &[array]);

        if output_value.is_err() {
            panic!("{}", output_value.unwrap_err().to_string());
        }
    }
}

#[export_name = "run"]
fn run((ptr, len): (i32, i32)) -> u64 {
  unsafe {
    let context = JS_CONTEXT.get().unwrap();
    let exports = ENTRYPOINT2.0.get().unwrap();
    let main = ENTRYPOINT2.1.get().unwrap();

    // Look at what has been written to the input buffer...
    // NB Possibly we should use the args ptr/len
    let ptr = PTR.lock().unwrap().clone();
    let len = LEN.lock().unwrap().clone();

    let vec = Vec::from_raw_parts(ptr as *mut u8, len as usize, len as usize);

    let array = context.array_value().unwrap();
    for val in vec.iter() {
      let jval = context.value_from_i32(i32::from(*val)).unwrap();
      array.append_property(jval);
    }
  
    // Now we need to add the bytes in
    let output_value = main.call(exports, &[array]);

    // output_value should be a byte array again...

    if output_value.is_err() {
      panic!("{}", output_value.unwrap_err().to_string());
    }

    // Convert output_value...
    let mut cursor = Cursor::new(Vec::new());
    // TODO: Write to the vector...

    let jsval = output_value.unwrap();
    let len = jsval.get_property("length").unwrap().as_u32_unchecked();

    for i in 0..len {
      let v = jsval.get_indexed_property(i).unwrap().as_u32_unchecked();
      let nval:&[u8] = &[v as u8];
      cursor.write(nval);
    }

    let mut vec = cursor.into_inner();
    vec.shrink_to_fit();

    let (ptr, len) = setBuffer(vec);

    return pack_uint32(ptr, len);
  }
}

#[cfg_attr(all(target_arch = "wasm32"), export_name = "resize")]
#[no_mangle]
pub unsafe extern "C" fn resize(size: u32) -> *const u8 {
   let existing_cap = READ_BUFFER.lock().unwrap().capacity() as u32;
   READ_BUFFER.lock().unwrap().reserve_exact((size - existing_cap) as usize);
   let ptr = READ_BUFFER.lock().unwrap().as_ptr();

   *PTR.lock().unwrap() = ptr as u32;
   *LEN.lock().unwrap() = size;
   return ptr
}
