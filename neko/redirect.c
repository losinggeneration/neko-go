#include "neko/neko_vm.h"

extern void goRedirect(const char *data, int size, void *param);

void setRedirectFunc(neko_vm *vm, void *param) {
	neko_vm_redirect(vm, goRedirect, param);
}
