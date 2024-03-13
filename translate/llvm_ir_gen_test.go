package translate

import (
	"os"
	"testing"
)

func TestLLVM(t *testing.T) {
	file, symTree:=getTestFileCtx()
	v1:=NewTypeCollector(symTree)
	v1.Visit(file)
	v2:=NewFuncCollector(symTree)
	v2.Visit(file)
	v3 := NewLlvmIrGen(symTree)
	v3.Visit(file)
	f,err:=os.Create("o.ll")
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(v3.TargetIL())
	if err != nil {
		panic(err)
	}
	println(v3.TargetIL())
}
