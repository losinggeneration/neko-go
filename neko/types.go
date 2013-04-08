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

type Buffer C.buffer
type Field C.field

type Kind struct {
	Zero int
}

type value struct {
	Type Val
}

type Value *value

type ObjCell struct {
	Id  Field
	Val *Value
}

type ObjTable struct {
	Count int
	Cells []ObjCell
}

type TFloat C.tfloat

type Float struct {
	Type  Val
	Float TFloat
}

type Int32 struct {
	Type Val
	Int  int
}

type Object struct {
	Type  Val
	Table ObjTable
	Proto *Object
}

type Function struct {
	Type   Val
	NArgs  int
	Addr   interface{}
	Env    *Value
	Module interface{}
}

type String struct {
	Type Val
	C    int8
}

type Array struct {
	Type Val
	Ptr  *Value
}

type Abstract struct {
	Type Val
	Kind *Kind
	Data interface{}
}

type HCell struct {
	HKey     int
	Key, Val *Value
	Next     *HCell
}

type Hash struct {
	Cells          *[]HCell
	NCells, NItems int
}

type MTLocal C.mt_local
type MTLock C.mt_lock

type Printer func(string, interface{})
