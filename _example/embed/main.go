package main

import (
	"log"

	"github.com/mattn/anko/vm"
	"github.com/mattn/anko/core"
)

func main() {
	env := vm.NewEnv()
	core.Import(env)
	core.ImportToX(env)
	v, err := env.Execute(`var fmt = import("fmt")
a = make([]int64, 10, 9)
println(a)`)
	if err != nil {
		utils.Printf("err %v\n", err)
		return
	}

	utils.Printf("v %v\n", v)
}
