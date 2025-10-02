//go:build !ccrelease
// +build !ccrelease

package cc

import (
	"fmt"
	"sync"
)

type DlFunc[Fn any, Ret any] struct {
	Name string
	Va   bool

	fn   Fn
	once sync.Once
}

func (sym *DlFunc[Fn, Ret]) init() {
	p := Symbol(sym.Name)
	if p == nil {
		panic(fmt.Errorf("symbol: %s no found", sym.Name))
	}
	if sym.Va {
		RFuncVa[Ret](&sym.fn, p)
	} else {
		RFunc[Ret](&sym.fn, p)
	}
}
func (sym *DlFunc[Fn, Ret]) Fn() Fn {
	sym.once.Do(sym.init)
	return sym.fn
}
func (sym *DlFunc[Fn, Ret]) FnVa() Fn {
	sym.once.Do(sym.init)
	return sym.fn
}
func BindFunc[Fn any, Ret any](name string) Fn {
	r := DlFunc[Fn, Ret]{Name: name, Va: false}
	return r.Fn()
}
func BindFuncVa[Fn any, Ret any](name string) Fn {
	r := DlFunc[Fn, Ret]{Name: name, Va: true}
	return r.FnVa()
}
