package neko

import "C"

var redirectFunc Printer

//export goRedirect
func goRedirect(data *C.char, size C.int, param interface{}) {
	redirectFunc(C.GoStringN(data, size), param)
}
