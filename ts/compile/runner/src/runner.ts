/*
	Copyright 2022 Loophole Labs
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at
		   http://www.apache.org/licenses/LICENSE-2.0
	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

declare var scale: ScaleFunction;    // This should be defined in the global scope.

import { TextEncoder, TextDecoder } from "text-encoding";

global.TextEncoder = TextEncoder;
global.TextDecoder = TextDecoder as typeof global["TextDecoder"];

import { Context, StringList } from "@loopholelabs/scale-signature-http";
import { encodeError } from "@loopholelabs/polyglot-ts";
import { Context as GuestContext } from "./guest";

const SCALE_NEXT: string = "scale_fn_next";

export type ScaleFunction = (a: GuestContext) => GuestContext;

function mainFunction() {
  console.log("Main function called");
}

// Run the scale next function
export function runNext(context: Context): Context {
  // context -> bytes
  let buf = context.encode(new Uint8Array());
  let data = Array.from(buf);

  // Call next()
  let nextfn = (global as any)[SCALE_NEXT];
  let newdata = nextfn(data);

  // bytes -> context
  const oContext = Context.decode(Uint8Array.from(newdata)).value;
  return oContext;
}

// Our own run function wrapper
function runFunction(data: number[]): number[] {
  // Decode data to a context
  const orgContext = Context.decode(Uint8Array.from(data)).value;

  // Use the global 'scale' function.
  // const scale = (global as any).scale as ScaleFunction;

  const GContext = new GuestContext(orgContext);

  try {
    const iContext = scale(GContext);

    // Encode the context back into an array
    let buf = iContext.Context().encode(new Uint8Array());
    let retdata = Array.from(buf);
    return retdata;
  } catch(e) {
    // We need to encode this as an error and return that to the host.
    let buf = encodeError(new Uint8Array(), e as Error);
    let retdata = Array.from(buf);
    return retdata;
  }
}

(global as any).Exports = {
  main: mainFunction,
  run: runFunction,
}
