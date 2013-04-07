package neko

// #include <malloc.h>
// #include "neko/neko_vm.h"
import "C"

import "unsafe"

type VM struct {
	vm *C.neko_vm
}

func bool_to_int(b bool) int {
	if b {
		return 1
	}

	return 0
}

func GlobalInit() {
	C.neko_global_init()
}

func GlobalFree() {
	C.neko_vm_select(nil)
	C.neko_global_free()
}

func GcMajor() {
	C.neko_gc_major()
}

func GcLoop() {
	C.neko_gc_loop()
}

func GcStats() (heap, free int) {
	var h, f C.int
	C.neko_gc_stats(&h, &f)
	return int(h), int(f)
}

/// TODO param  neko "thread_main_func" init
/// TODO param  neko "thread_main_func" main
/// TODO param "void *" pparam
/// TODO param "void **" handle
/// TODO return "int"
func ThreadCreate() {
}

/// TODO param  neko "thread_main_func" f
/// TODO param "void *" p
func ThreadBlocking() {
}

func ThreadRegister(t bool) bool {
	return C.neko_thread_register(C.bool(bool_to_int(t))) == 1
}

func NewVM() *VM {
	return &VM{vmAlloc()}
}

func vmAlloc() *C.neko_vm {
	return C.neko_vm_alloc(nil)
}

func CurrentVM() *VM {
	return &VM{C.neko_vm_current()}
}

func (vm *VM) ExecStack() Value {
	return inlineCvalToValue(C.neko_exc_stack(vm.vm))
}

func (vm *VM) CallStack() Value {
	return inlineCvalToValue(C.neko_call_stack(vm.vm))
}

/// FIXME
func (vm *VM) Custom(k Kind) interface{} {
	return nil
	//return C.neko_vm_custom(vm.vm, *(*C.vkind)(unsafe.Pointer(&k)))
}

/// FIXME
func (vm *VM) SetCustom(k Kind, v interface{}) {
	//C.neko_vm_set_custom(vm.vm, *(*C.vkind)(unsafe.Pointer(&k)), unsafe.Pointer(&v))
}

func (vm *VM) Execute(module interface{}) Value {
	return inlineCvalToValue(C.neko_vm_execute(vm.vm, unsafe.Pointer(&module)))
}

func (vm *VM) Select() {
	C.neko_vm_select(vm.vm)
}

func (vm *VM) Jit(enable bool) int {
	return int(C.neko_vm_jit(vm.vm, C.int(bool_to_int(enable))))
}

func (vm *VM) Trusted(trusted bool) int {
	return int(C.neko_vm_trusted(vm.vm, C.int(bool_to_int(trusted))))
}

/// TODO param "neko_printer" print
/// TODO param "void *" param
func (vm *VM) Redirect() {
}

/// TODO param "neko_stat_func" fstats
/// TODO param "neko_stat_func" pstats
func (vm *VM) SetStats() {
}

func (vm *VM) DebugStack() {
	C.neko_vm_dump_stack(vm.vm)
}

func DefaultLoader(argv []string) Value {
	strs := make([]*C.char, len(argv))
	for i, s := range argv {
		strs[i] = C.CString(s)
		defer C.free(unsafe.Pointer(strs[i]))
	}

	return inlineCvalToValue(C.neko_default_loader((**C.char)(unsafe.Pointer(&strs[0])), C.int(len(argv))))
}

func IsBigEndian() bool {
	return C.neko_is_big_endian() == 1
}
