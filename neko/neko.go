package neko

// #include <malloc.h>
// #include "neko/neko.h"
// #cgo LDFLAGS: -lneko
import "C"

import "unsafe"

func inlineValueToCval(v Value) C.value {
	return C.value(unsafe.Pointer(v))
}

func inlineCvalToValue(v C.value) Value {
	return Value(unsafe.Pointer(v))
}

func AllocFloat(t TFloat) Value {
	return inlineCvalToValue(C.neko_alloc_float(C.tfloat(t)))
}

func AllocInt32(i int) Value {
	return inlineCvalToValue(C.neko_alloc_int32(C.int(i)))
}

func AllocString(str string) Value {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return inlineCvalToValue(C.neko_alloc_string(cstr))
}

func AllocEmptyString(size uint) Value {
	return inlineCvalToValue(C.neko_alloc_empty_string(C.uint(size)))
}

func CopyString(str string) Value {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return inlineCvalToValue(C.neko_copy_string(cstr, C.int_val(len(str))))
}

func ValThis() Value {
	return inlineCvalToValue(C.neko_val_this())
}

func ValId(str string) Field {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return Field(C.neko_val_id(cstr))
}

func ValField(o Value, f Field) Value {
	return inlineCvalToValue(C.neko_val_field(inlineValueToCval(o), C.field(f)))
}

func AllocObject(o Value) Value {
	return inlineCvalToValue(C.neko_alloc_object(inlineValueToCval(o)))
}

func AllocField(obj Value, f Field, v Value) {
}

/*
func ValIterFields( obj Value, func ( v Value, f Field, vp interface{}), vp interface{}) {
}
*/

func ValFieldName(f Field) Value {
	return inlineCvalToValue(C.neko_val_field_name(C.field(f)))
}

func AllocArray(n uint) Value {
	return inlineCvalToValue(C.neko_alloc_array(C.uint(n)))
}

func AllocAbstract(k Kind, data interface{}) Value {
	return inlineCvalToValue(C.neko_alloc_abstract(*(*C.vkind)(unsafe.Pointer(&k)), unsafe.Pointer(&data)))
}

func ValCall0(f Value) Value {
	return inlineCvalToValue(C.neko_val_call0(inlineValueToCval(f)))
}

func ValCall1(f, arg Value) Value {
	return inlineCvalToValue(C.neko_val_call1(inlineValueToCval(f), inlineValueToCval(arg)))
}

func ValCall2(f, arg1, arg2 Value) Value {
	return inlineCvalToValue(C.neko_val_call2(inlineValueToCval(f), inlineValueToCval(arg1), inlineValueToCval(arg2)))
}

func ValCall3(f, arg1, arg2, arg3 Value) Value {
	return inlineCvalToValue(C.neko_val_call3(inlineValueToCval(f), inlineValueToCval(arg1), inlineValueToCval(arg2), inlineValueToCval(arg3)))
}

func ValCallN(f Value, args []Value) Value {
	return inlineCvalToValue(C.neko_val_callN(inlineValueToCval(f), (*C.value)(unsafe.Pointer(&args[0])), C.int(len(args))))
}

func ValOCall0(o Value, f Field) Value {
	return inlineCvalToValue(C.neko_val_ocall0(inlineValueToCval(o), C.field(f)))
}

func ValOCall1(o Value, f Field, arg Value) Value {
	return inlineCvalToValue(C.neko_val_ocall1(inlineValueToCval(o), C.field(f), inlineValueToCval(arg)))
}

func ValOCall2(o Value, f Field, arg1 Value, arg2 Value) Value {
	return inlineCvalToValue(C.neko_val_ocall2(inlineValueToCval(o), C.field(f), inlineValueToCval(arg1), inlineValueToCval(arg2)))
}

func ValOCallN(o Value, f Field, args []Value) Value {
	return inlineCvalToValue(C.neko_val_ocallN(inlineValueToCval(o), C.field(f), (*C.value)(unsafe.Pointer(&args[0])), C.int(len(args))))
}

func ValCallEx(this, f Value, args []Value, exc *Value) Value {
	return inlineCvalToValue(C.neko_val_callEx(inlineValueToCval(this),
		inlineValueToCval(f), (*C.value)(unsafe.Pointer(&args[0])),
		C.int(len(args)), (*C.value)(unsafe.Pointer(exc))))
}

func AllocRoot(nvals uint) *Value {
	return (*Value)(unsafe.Pointer(C.neko_alloc_root(C.uint(nvals))))
}

func FreeRoot(r *Value) {
}

func Alloc(nbytes uint) *uint8 {
	return (*uint8)(unsafe.Pointer(C.neko_alloc(C.uint(nbytes))))
}

func AllocPrivate(nbytes uint) *uint8 {
	return (*uint8)(unsafe.Pointer(C.neko_alloc_private(C.uint(nbytes))))
}

func AllocFunction(c_prim interface{}, nargs uint, name string) {
}

func AllocBuffer(init string) Buffer {
	cstr := C.CString(init)
	defer C.free(unsafe.Pointer(cstr))
	return Buffer(C.neko_alloc_buffer(cstr))
}

func BufferAppend(b Buffer, s string) {
}

func BufferAppendSub(b Buffer, s string) {
}

func BufferAppendChar(b Buffer, c uint8) {
}

func BufferToString(b Buffer) Value {
	return inlineCvalToValue(C.neko_buffer_to_string(C.buffer(b)))
}

func ValBuffer(b Buffer, v Value) {
}

func ValCompare(a Value, b Value) int {
	return int(C.neko_val_compare(inlineValueToCval(a), inlineValueToCval(b)))
}

func ValPrint(s Value) {
}

/*
func ValGC( v Value, f Finalizer ) {
}
*/

func ValThrow(v Value) {
}

func ValRethrow(v Value) {
}

func ValHash(v Value) int {
	return int(C.neko_val_hash(inlineValueToCval(v)))
}

func KindShare(k *Kind, name string) {
}

func Failure(msg Value, file string, line int) {
}

func AllocLocal() *MTLocal {
	return (*MTLocal)(C.neko_alloc_local())
}

func LocalGet(l *MTLocal) interface{} {
	return C.neko_local_get((*C.mt_local)(l))
}

func LocalSet(l *MTLocal, v interface{}) {
}

func FreeLocal(l *MTLocal) {
}

func AllocLock() *MTLock {
	return (*MTLock)(C.neko_alloc_lock())
}

func LockAcquire(l *MTLock) {
}

func LockTry(l *MTLock) bool {
	return C.neko_lock_try((*C.mt_lock)(l)) == 1
}

func LockRelease(l *MTLock) {
}

func FreeLock(l *MTLock) {
}
