package main

import (
	"fmt"
	"github.com/losinggeneration/neko-go/neko"
)

func load_file(mload neko.Value) {
	args := []neko.Value{neko.AllocString("test.n"), mload}
	var exc *neko.Value
	neko.ValCallEx(mload, neko.ValField(mload, neko.ValId("loadmodule")), args, exc)
}

func main() {
	args := make([]string, 5)

	for i, _ := range args {
		args[i] = fmt.Sprintf("test%v", i+1)
	}

	neko.GlobalInit()
	vm := neko.NewVM()
	vm.Select()
	vm.Jit(true)
	mload := neko.DefaultLoader(args)
	load_file(mload)
	neko.GlobalFree()
}
