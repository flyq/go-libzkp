package main

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
import "C"
import "fmt"

func foo() {
	handle := C.dlopen(C.CString("libfoo.so"), C.RTLD_LAZY)
	bar := C.dlsym(handle, C.CString("bar"))
	fmt.Printf("bar is at %p\n", bar)
}
