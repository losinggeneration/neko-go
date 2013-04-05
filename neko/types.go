package neko

// #include "neko/neko.h"
import "C"

type Val C.val_type

const (
	VAL_INT       Val = C.VAL_INT
	VAL_NULL      Val = C.VAL_NULL
	VAL_FLOAT     Val = C.VAL_FLOAT
	VAL_BOOL      Val = C.VAL_BOOL
	VAL_STRING    Val = C.VAL_STRING
	VAL_OBJECT    Val = C.VAL_OBJECT
	VAL_ARRAY     Val = C.VAL_ARRAY
	VAL_FUNCTION  Val = C.VAL_FUNCTION
	VAL_ABSTRACT  Val = C.VAL_ABSTRACT
	VAL_INT32     Val = C.VAL_INT32
	VAL_PRIMITIVE Val = C.VAL_PRIMITIVE
	VAL_JITFUN    Val = C.VAL_JITFUN
	VAL_32_BITS   Val = C.VAL_32_BITS
)

type Value struct {
	Type Val
}

type Field C.field

type ObjCell struct {
	Id  Field
	Val *Value
}

type ObjTable struct {
	Count int
	Cells []ObjCell
}

type Buffer C.buffer
type TFloat C.tfloat

type Float struct {
	Type  Val
	Float TFloat
}

type Int32 struct {
	Type Val
	Int  int
}

