package cc

import (
	"reflect"
	"sync"
	"unsafe"

	"github.com/jinzhongmin/gg/cc/dl"
)

type uptr = unsafe.Pointer
type iptr = uintptr

type Long dl.Long
type Longlong dl.Longlong
type ULong dl.ULong
type ULonglong dl.ULonglong
type Void dl.Void_
type Bool uint8
type SizeT dl.SizeT
type Enum dl.Enum

func MakeBool(b bool) Bool {
	if b {
		return True
	} else {
		return False
	}
}

const (
	True  Bool = 1
	False Bool = 0
)

type Integer interface {
	~int8 | ~int16 | ~int32 | ~int | ~int64 |
		~uint8 | ~uint16 | ~uint32 | ~uint | ~uint64
}

func anyptr(a interface{}) uptr               { return (*(*[2]uptr)((uptr(&a))))[1] }
func Slice[T any, N Integer](ptr *T, l N) []T { return unsafe.Slice(ptr, l) }

func Malloc(size uint64) uptr            { return dl.CMalloc(size) }
func Free(p uptr)                        { dl.CFree(p) }
func Memset(p uptr, v int32, num uint64) { dl.CMemset(p, v, num) }

type Ptr iptr

func (p Ptr) Len() uint64 {
	if p == 0 {
		return 0
	}

	i := (*[1 << 30]unsafe.Pointer)(uptr(p))
	idx := uint64(0)
	for ; i[idx] != nil; idx += 1 {
	}
	return idx
}
func (p Ptr) Free() { Free(uptr(p)) }

// func (p Ptr) Slice(n ...uint64) uptr {
// 	var count uint64
// 	if len(n) != 0 {
// 		count = n[0]
// 	} else {
// 		count = p.Len()
// 	}
// 	return Slice((*uptr)(uptr(p)), count)
// }

type stringStruct struct {
	p unsafe.Pointer
	l int
}

var stringNull = (*string)((unsafe.Pointer)(&stringStruct{nil, -1}))

func StringNil() string { return *stringNull }

type String iptr

// func StringRef(str string) String { return *(*String)(uptr(&str)) }

func NewString(str string) String {
	s, _ := NewStringLen(str)
	return s
}
func NewStringN(num uint64) String {
	sp := dl.CMalloc(num)
	dl.CMemset(sp, 0, num)
	return String(sp)
}
func NewStringLen(str string) (String, uint64) {
	if str == *stringNull {
		return 0, 0
	}

	strLen := len(str)
	size := uint64(strLen + 1)

	malloc := dl.CMalloc
	p := malloc(size)
	if p == nil {
		return 0, 0
	}

	r := String(p)
	r.Copy(str, (size - 1))
	return r, size
}
func NewStringRef(str string) (ptr String, len int64) {
	pp := *(*[2]iptr)(uptr(&str))
	return String(pp[0]), int64(pp[1])
}
func (c String) Len() uint64 {
	if c == 0 {
		return 0
	}

	i := (*[1 << 30]byte)(uptr(c))
	idx := uint64(0)
	for ; i[idx] != 0; idx += 1 {
	}
	return idx
}
func (c String) Free() {
	if c != 0 {
		dl.CFree(uptr(c))
	}
}
func (c String) TakeString(n ...uint64) string {
	defer c.Free()
	return c.String(n...)
}
func (c String) String(n ...uint64) string {
	return string(c.RefBytes(n...))
}
func (c String) TakeBytes(n ...uint64) []byte {
	defer c.Free()
	return c.Bytes(n...)
}
func (c String) Bytes(n ...uint64) []byte {
	l := c.count(n...)
	dst := make([]byte, l)
	copy(dst, c.RefBytes(l))
	return dst
}
func (c String) RefString(n ...uint64) string {
	s := struct {
		p   String
		len int
	}{c, int(c.count(n...))}
	return *(*string)(uptr(&s))
}
func (c String) RefBytes(n ...uint64) []byte {
	count := c.count(n...)
	return unsafe.Slice((*byte)(uptr(c)), count)
}
func (c String) Copy(src string, n uint64) {
	_dst := c.RefBytes(n + 1)
	_src := Slice((*[2]*byte)(uptr(&src))[0], len(src))
	copy(_dst, _src)
	(_dst)[n] = 0
}
func (c String) count(n ...uint64) uint64 {
	if len(n) != 0 {
		return n[0]
	}
	return c.Len()
}

// func (c String) Bytes(n int) []byte { return unsafe.Slice((*byte)(uptr(c)), n) }

type Strings iptr

func NewStrings(strs []string) Strings {
	ss, _ := NewStringsLen(strs)
	return ss
}
func NewStringsLen(strs []string) (Strings, int32) {
	if strs == nil {
		return 0, 0
	}

	ls := len(strs)
	ss := Malloc(uint64(ls*8) + 1)

	list := unsafe.Slice((*String)(ss), ls+1)
	for idx, str := range strs {
		list[idx] = NewString(str)
	}
	list[ls] = 0
	return Strings(ss), int32(ls)
}
func (css Strings) Len() uint64 { return Ptr(css).Len() }
func (css *Strings) Free(n ...uint64) {
	if (*css) == 0 {
		return
	}
	var l uint64
	if len(n) > 0 {
		l = n[0]
	} else {
		l = css.Len()
	}

	cp := (*[1 << 30]String)(uptr(*css))
	for i := 0; i < int(l); i++ {
		cp[i].Free()
	}

	dl.CFree(uptr(*css))
	(*css) = 0
}
func (css Strings) TakeStrings(n ...uint64) []string {
	defer css.Free(n...)
	return css.Strings(n...)
}
func (css Strings) Strings(n ...uint64) []string {
	if css == 0 {
		return nil
	}

	var count uint64
	if len(n) != 0 {
		count = n[0]
	} else {
		count = css.Len()
	}

	ss := make([]string, count)
	cp := (*[1 << 30]String)(uptr(css))
	for i := 0; i < int(count); i++ {
		ss[i] = cp[i].String()
	}
	return ss
}
func (css Strings) Ref(n ...uint64) []string {
	if css == 0 {
		return nil
	}

	var count uint64
	if len(n) != 0 {
		count = n[0]
	} else {
		count = css.Len()
	}

	ss := make([]string, count)
	cp := (*[1 << 30]String)(uptr(css))
	for i := 0; i < int(count); i++ {
		ss[i] = cp[i].RefString()
	}
	return ss
}

type Func Ptr

func (fnp *Func) Bind(fn interface{}) { *(*uintptr)(uptr(fnp)) = iptr(Cbk(fn)) }
func (fnp *Func) Free()               { CbkClose(uptr(*fnp)); *fnp = 0 }

func FuncBindRaw[Fn any](fn *Func, raw func(out, ins uptr)) {
	(*fn) = Func(CbkRaw[Fn](raw))
}

var vmenLock = sync.Mutex{}
var vmen = struct {
	ptrs    []any
	unseted []uint64
}{ptrs: []any{nil}, unseted: []uint64{}}

type VPtr[T any] uint64

func NewVPtr[T any](val T) VPtr[T] {
	vmenLock.Lock()
	defer vmenLock.Unlock()

	l := len(vmen.unseted)
	if l > 0 {
		idx := vmen.unseted[l-1]
		vmen.unseted = vmen.unseted[:idx]
		vmen.ptrs[idx] = &val
		return VPtr[T](idx)
	}

	vmen.ptrs = append(vmen.ptrs, &val)
	return VPtr[T](len(vmen.ptrs) - 1)
}
func (vp VPtr[T]) Get() T   { return *vmen.ptrs[vp].(*T) }
func (vp VPtr[T]) Addr() *T { return vmen.ptrs[vp].(*T) }
func (vp VPtr[T]) UnSet() {
	vmen.ptrs[vp] = nil
	vmenLock.Lock()
	vmen.unseted = append(vmen.unseted, uint64(vp))
	vmenLock.Unlock()
	vp = 0
}

// #region ffi

func Open(lib string) uintptr { return dl.Open(lib) }
func Symbol(sym string) uptr  { return dl.Symbol(sym) }

func RFunc[Ret any](fptr interface{}, cptr uptr) { dl.BindFunc[Ret](fptr, cptr) }
func RFuncVa[Ret any](fptr interface{}, cptr uptr) {

	fn := reflect.ValueOf(fptr).Elem()
	typ := fn.Type()

	out := dl.Void()
	outc := typ.NumOut()
	retfn := func(r Ret) []reflect.Value { return []reflect.Value{} }

	if outc > 0 {
		t := typ.Out(0)
		out = dl.CCifType(t)
		retfn = func(r Ret) []reflect.Value { return []reflect.Value{reflect.ValueOf(r)} }
	}

	v := reflect.MakeFunc(typ, func(args []reflect.Value) (results []reflect.Value) {
		argc := len(args) - 1

		ins := make([]*dl.Type, argc)
		argv := make([]interface{}, argc)

		for i := 0; i < argc; i++ {
			ins[i] = dl.CCifType(args[i].Type())
			argv[i] = args[i].Interface()
		}

		last := args[argc]
		for i := 0; i < last.Len(); i++ {
			val := last.Index(i).Interface()

			argv = append(argv, val)
			ins = append(ins, dl.CCifType(reflect.TypeOf(val)))
		}

		cif := dl.CifPrep(dl.AbiDefault, out, ins)
		defer cif.Close()

		r := dl.CallCif[Ret](cif, cptr, argv)
		_ = args
		return retfn(r)
	})
	reflect.ValueOf(fptr).Elem().Set(v)
}

func Cbk(fn interface{}) uptr {
	if anyptr(fn) == nil {
		return nil
	}
	cls := dl.ClosureAlloc()
	cls.ClosureLoc(fn)
	return cls.Code()
}
func CbkClose(u uptr)           { dl.ClosureCloseByCode(u) }
func CbkCloseLate(ptrs ...uptr) { dl.ClosureCloseByCodeLate(ptrs...) }

func CbkRaw[T any](fn func(out, ins uptr)) uptr {
	if fn == nil {
		return nil
	}
	var f T
	cls := dl.ClosureAlloc()
	cls.ClosureLoc(f)
	cls.RestLocRaw(fn)
	return cls.Code()
}

func RawInAddr[T any](ins uptr, idx int) *T {
	return unsafe.Slice((**T)(ins), idx+1)[idx]
}
func RawInValue[T any](ins uptr, idx int) T {
	return *unsafe.Slice((**T)(ins), idx+1)[idx]
}

// #endregion
