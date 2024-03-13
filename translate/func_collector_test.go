package translate

import (
	"testing"
)

func TestFuncCollector(t *testing.T) {
	file, symTree:=getTestFileCtx()
	v1:=NewTypeCollector(symTree)
	v1.Visit(file)
	v2:=NewFuncCollector(symTree)
	v2.Visit(file)
	println(v2.SM)
}
