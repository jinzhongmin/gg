//go:build ccrelease

package cc

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/jinzhongmin/gg/cc/dl"
)

// var voidZero Void
// var voidType = reflect.TypeOf(voidZero)

type DlFunc[F any, Ret any] struct {
	Name string
	Va   bool

	fn   F
	once sync.Once

	proc uptr
	call func(...any) Ret
}

func (sym *DlFunc[F, Ret]) init() {
	sym.proc = Symbol(sym.Name)
	if sym.proc == nil {
		panic(fmt.Errorf("symbol: %s no found", sym.Name))
	}
	if sym.Va {
		RFuncVa[Ret](&sym.fn, sym.proc)
	} else {
		typ := reflect.TypeOf(sym.fn)
		cif := dl.CifPrepByGoFn(dl.AbiDefault, typ)
		sym.call = func(a ...interface{}) Ret { return dl.CallCif[Ret](cif, sym.proc, a) }
	}
}
func (sym *DlFunc[F, Ret]) Fn() func(...any) Ret {
	sym.once.Do(sym.init)
	return sym.call
}
func (sym *DlFunc[F, Ret]) FnVa() F {
	sym.once.Do(sym.init)
	return sym.fn
}

func BindFunc[Fn any, Ret any](name string) func(...any) Ret {
	r := DlFunc[Fn, Ret]{Name: name, Va: false}
	return r.Fn()
}
func BindFuncVa[Fn any, Ret any](name string) Fn {
	r := DlFunc[Fn, Ret]{Name: name, Va: true}
	return r.FnVa()
}
