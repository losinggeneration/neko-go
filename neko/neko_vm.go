package neko

// #include "neko/neko_vm.h"
import "C"

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
func ThreadCreate() {}

/// TODO param  neko "thread_main_func" f
/// TODO param "void *" p
func ThreadBlocking() {}

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

/// TODO return neko "value"
func (vm *VM) ExecStack() {}

/// TODO return neko "value"
func (vm *VM) CallStack() {}

/// TODO param neko "vkind" k
/// TODO return "void *"
func (vm *VM) Custom() {}

/// TODO param neko "vkind" k
/// TODO param "void *" v
func (vm *VM) SetCustom() {}

/// TODO param "void *" module
/// TODO return neko "value"
func (vm *VM) Execute() {}

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
func (vm *VM) Redirect() {}

/// TODO param "neko_stat_func" fstats
/// TODO param "neko_stat_func" pstats
func (vm *VM) SetStats() {}

func (vm *VM) DebugStack() {
	C.neko_vm_dump_stack(vm.vm)
}

/// TODO return neko "value"
func DefaultLoader(argv []string) {
}

func IsBigEndian() bool {
	return C.neko_is_big_endian() == 1
}
