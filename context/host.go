package context

// next is the host function that is called by the runtime to execute the next
// function in the chain.
//
//export next
//go:linkname next
func next(offset uint32, length uint32) (packed uint64)

////export debug
////go:linkname debug
//func debug(offset uint32, byteCount uint32)
