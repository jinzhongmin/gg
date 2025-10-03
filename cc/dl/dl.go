//go:build !nocgo
// +build !nocgo

package dl

/*
#cgo LDFLAGS: -ldl
#cgo pkg-config: libffi

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <stdbool.h>
#include <dlfcn.h>
#include <ffi.h>

typedef unsigned long ulong;
typedef unsigned long long ulonglong;
typedef signed long slong;
typedef signed long long slonglong;
typedef void _void_t_;
typedef enum _enum_t_{_enum_t_none}_enum_t_;

extern void closure_agent(void* ccif,void* out,void* ins, void* ccls);

typedef struct _go_slice{
	void* ptr;
	int64_t len;
	int64_t cap;
}go_slice;

typedef struct _go_any{
	void* t;
	void* ptr;
}go_any;

typedef struct _go_string{
	char* ptr;
	int64_t len;
}go_string;

typedef struct _go_cif{
	ffi_cif* cif;
	go_slice cins_type;
	ffi_type* cout_type;
	go_slice* go_ins_geter;
	int64_t go_id;
	go_slice c_ins_geter;
} go_cif;

typedef struct _free_fn{
	void (*fn)(void*);
	void *data;
}free_fn;

static void free_ptr(void* p){ free(p);}
static void* getter_ptr(go_any* a, int arg_size, int idx, free_fn *frees){
	frees[idx].fn = NULL;
	return (void*)(&a->ptr);
}
static void* getter_str(go_any* a, int arg_size, int idx, free_fn *frees){
	go_string* sp = (go_string*)a->ptr;

	char* dest = (char*)malloc(sp->len+1);
	memcpy(dest, sp->ptr, sp->len);
	dest[sp->len] = 0;

	frees[idx].fn = free_ptr;
	frees[idx].data = dest;

	return (void*)(sp->ptr);
}
static void* getter_any(go_any* a, int arg_size, int idx, free_fn *frees){
	frees[idx].fn = NULL;
	return (void*)(a->ptr);
}
static void* getter_struct(go_any* a, int arg_size, int idx, free_fn *frees){
	frees[idx].fn = NULL;
	return (void*)(a->ptr);

	// void* dest = malloc(arg_size);
	// memcpy(dest, *((void**)(a->ptr)), arg_size);

	// frees[idx].fn = free_ptr;
	// frees[idx].data = dest;

	// return dest;
}

static void* (*getters_map[])(go_any* a, int arg_size, int idx, free_fn *frees) = {getter_ptr, getter_str, getter_any, getter_struct};


static void* cccif_call(ffi_cif* cif, void* fnp,
	int argc, go_any* args, int* getters, ffi_type** cins_type){

	void** ins = (void**)((int64_t*)args + argc*2);
	free_fn* frees = (free_fn*)((int64_t*)ins + argc);

	int i = 0;
	for(i=0; i<argc; i++){
		ins[i] = getters_map[getters[i]](&args[i], cins_type[i]->size, i, frees);
	}

	void* out = (int64_t*)args + argc*5;
	ffi_call(cif, fnp, out, ins);

	for(i=0; i<argc; i++){
		if(frees[i].fn != NULL){
			frees[i].fn(frees[i].data);
		}
	}
	return out;
}
*/
import "C"
import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
	"unsafe"
)

//#region C

type Long C.slong
type Longlong C.slonglong
type ULong C.ulong
type ULonglong C.ulonglong
type Void_ C._void_t_
type SizeT C.size_t
type Unsigned C.unsigned
type Enum C._enum_t_

//#endregion

type uptr = unsafe.Pointer
type iptr = uintptr

type Abi = C.ffi_abi

const (
	AbiDefault Abi = C.FFI_DEFAULT_ABI
)

type Type = C.ffi_type

func U8() *Type   { return &C.ffi_type_uint8 }
func I8() *Type   { return &C.ffi_type_sint8 }
func U16() *Type  { return &C.ffi_type_uint16 }
func I16() *Type  { return &C.ffi_type_sint16 }
func U32() *Type  { return &C.ffi_type_uint32 }
func I32() *Type  { return &C.ffi_type_sint32 }
func U64() *Type  { return &C.ffi_type_uint64 }
func I64() *Type  { return &C.ffi_type_sint64 }
func F32() *Type  { return &C.ffi_type_float }
func F64() *Type  { return &C.ffi_type_double }
func Void() *Type { return &C.ffi_type_void }
func Ptr() *Type  { return &C.ffi_type_pointer }
func Str() *Type  { return ffi_type_str }
func Bool() *Type { return ffi_type_bool }

var ffi_type_str *C.ffi_type
var ffi_type_bool *C.ffi_type

var cTypeSize = uint64(unsafe.Sizeof(Type{}))

func init() {
	ffi_type_str = (*C.ffi_type)(CMalloc(cTypeSize))
	ffi_type_str._type = C.ffi_type_pointer._type
	ffi_type_str.alignment = C.ffi_type_pointer.alignment
	ffi_type_str.size = C.ffi_type_pointer.size
	ffi_type_str.elements = C.ffi_type_pointer.elements

	ffi_type_bool = (*C.ffi_type)(CMalloc(cTypeSize))
	ffi_type_bool._type = C.ffi_type_sint32._type
	ffi_type_bool.alignment = C.ffi_type_sint32.alignment
	ffi_type_bool.size = C.ffi_type_sint32.size
	ffi_type_bool.elements = C.ffi_type_sint32.elements
}

// #region stdlib

func CStr(s string) uptr                  { return uptr(C.CString(s)) }
func CMalloc(size uint64) uptr            { return C.malloc(C.size_t(size)) }
func CFree(ptr uptr)                      { C.free(ptr) }
func CMemset(ptr uptr, c int32, n uint64) { C.memset(ptr, C.int(c), C.size_t(n)) }
func CMemcpy(dst, src uptr, n uint64)     { C.memcpy(dst, src, C.size_t(n)) }
func CStrLen(ptr uptr) uint64             { return uint64(C.strlen((*C.char)(ptr))) }

// #endregion

// #region cnf

func libsFind(lib string) (string, error) {

	var dirs []string = []string{"./"}
	dirs = append(dirs, strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))...)

	if runtime.GOOS == "linux" {
		dirs = append(dirs,
			"/lib", "/lib64", "/usr/lib", "/usr/lib64",
			"/usr/local/lib",
		)
		if ld := os.Getenv("LD_LIBRARY_PATH"); ld != "" {
			dirs = append(dirs, strings.Split(ld, string(os.PathListSeparator))...)
		}
	}

	for _, dir := range dirs {
		if dir == "" {
			continue
		}
		pattern := filepath.Join(dir, lib)
		files, err := filepath.Glob(pattern)
		if err != nil {
			return "", fmt.Errorf("pattern error in %s: %v", pattern, err)
		}
		if len(files) > 0 {
			// c.libsSave(lib, files[0])
			return files[0], nil
		}
	}
	return "", fmt.Errorf("no file matching pattern %q found", lib)
}

// #endregion

// #region dlfcn

func Open(fileName string) (handle iptr) {
	fn, err := libsFind(fileName)
	if err != nil {
		log.Println(err)
		return
	}

	cFileName := C.CString(fn)
	defer C.free(uptr(cFileName))
	handle = iptr(uptr(C.dlopen(cFileName, C.RTLD_LAZY|C.RTLD_GLOBAL)))
	if handle == 0 {
		log.Printf("library %s\n can not open", fn)
	}
	return
}
func Close(handle uptr) { C.dlclose(handle) }
func Symbol(name string) (addr uptr) {
	cName := C.CString(name)
	defer C.free(uptr(cName))
	return C.dlsym(nil, cName)
}
func Error() string { return C.GoString(C.dlerror()) }

// #endregion

// #region cif

type cCif = C.ffi_cif

var cCifSize = uint64(unsafe.Sizeof(cCif{}))

type Cif struct {
	cCif       *cCif
	cIns       []*Type
	cOut       *Type
	goInsGeter *[]func(a *any, free *[]func()) uptr
	goId       int

	cinsGeter []int32 // 0:ptr, 1:str, 2:any, 3:struct
}

func cifAlloc() *Cif {
	cif := new(Cif)
	cif.goInsGeter = new([]func(a *any, free *[]func()) uptr)
	cif.goId = pin(cif)
	return cif
}
func cifPrep(cif *Cif, abi Abi, out *Type, ins []*Type) {

	var cCif = (*cCif)(CMalloc(cCifSize))

	argc := len(ins)
	cif.cIns = make([]*Type, argc+1)
	cif.cIns[argc] = nil
	copy(cif.cIns, ins)

	(*cif.goInsGeter) = make([]func(a *any, free *[]func()) uptr, argc)
	cif.cinsGeter = make([]int32, argc)
	for i := 0; i < argc; i++ {
		switch ins[i] {
		case Ptr():
			(*cif.goInsGeter)[i] = func(i *any, fs *[]func()) uptr {
				a := *(*[2]uptr)(uptr(i))
				p := uptr(&a[1])
				return p
			}
			cif.cinsGeter[i] = 0
			continue
		case Str():
			(*cif.goInsGeter)[i] = func(i *any, fs *[]func()) uptr {
				s := CStr((*i).(string))
				(*fs) = append(*fs, func() { CFree(s) })
				return uptr(&s)
			}
			cif.cinsGeter[i] = 1
			continue

		default:
			// (*cif.goInsGeter)[i] = func(i *any, _ *[]func()) uptr {
			// 	return (*(*[2]uptr)(uptr(i)))[1]
			// }

			isStruct := false
			for t, v := range structType {
				if ins[i] != v {
					continue
				}

				size := t.Size()
				(*cif.goInsGeter)[i] = func(a *any, free *[]func()) uptr {

					val := reflect.New(t)
					val.Elem().Set(reflect.ValueOf(*a))

					dst := CMalloc(uint64(size) + 8)
					C.memcpy(dst, val.UnsafePointer(), C.size_t(size))

					(*free) = append((*free), func() { CFree(dst) })
					return dst
				}
				isStruct = true
				cif.cinsGeter[i] = 3
				break
			}
			if !isStruct {
				(*cif.goInsGeter)[i] = func(i *any, _ *[]func()) uptr {
					return (*(*[2]uptr)(uptr(i)))[1]
				}
				cif.cinsGeter[i] = 2
			}
		}
	}

	C.ffi_prep_cif(cCif, abi, C.uint(argc), out, &cif.cIns[0])

	cif.cCif = cCif
	cif.cOut = out
}
func cifPrepAny(cif *Cif, abi Abi, typ reflect.Type) {
	out, ins := cCifTypeIO(typ)
	cifPrep(cif, abi, out, ins)
}

var structType = map[reflect.Type]*Type{}

func CCifType(t reflect.Type) *Type {
	switch t.Kind() {
	case reflect.Bool:
		return Bool()
	case reflect.Int:
		return I64()
	case reflect.Int8:
		return I8()
	case reflect.Int16:
		return I16()
	case reflect.Int32:
		return I32()
	case reflect.Int64:
		return I64()
	case reflect.Uint:
		return U64()
	case reflect.Uint8:
		return U8()
	case reflect.Uint16:
		return U16()
	case reflect.Uint32:
		return U32()
	case reflect.Uint64:
		return U64()
	case reflect.Uintptr:
		return U64()
	case reflect.Float32:
		return F32()
	case reflect.Float64:
		return F64()
	case reflect.String:
		return Str()
	case reflect.Struct:

		v, ok := structType[t]
		if ok {
			return v
		}

		rt := (*Type)(CMalloc(cTypeSize))
		rt.size = C.size_t(t.Size())
		rt.alignment = C.ushort(t.Align())
		rt._type = C.FFI_TYPE_STRUCT //13

		num := t.NumField()
		ptr := CMalloc(uint64(num * 8))
		rt.elements = (**C.ffi_type)(ptr)

		lst := unsafe.Slice((**Type)(ptr), num)
		for i := 0; i < num; i++ {
			lst[i] = CCifType(t.Field(i).Type)
		}

		structType[t] = rt
		return rt
	}
	return Ptr()
}
func cCifTypeIO(typ reflect.Type) (out *Type, ins []*Type) {
	insNum := typ.NumIn()

	cifIns := make([]*Type, insNum)
	cifOut := Void()

	for i := 0; i < insNum; i++ {
		t := typ.In(i)
		cifIns[i] = CCifType(t)
	}
	if typ.NumOut() > 0 {
		cifOut = CCifType(typ.Out(0))
	}
	return cifOut, cifIns
}
func CifPrep(abi Abi, out *Type, ins []*Type) *Cif {
	cif := cifAlloc()
	cifPrep(cif, abi, out, ins)
	return cif
}
func CifPrepByGoFn(abi Abi, fn reflect.Type) *Cif {
	cif := cifAlloc()
	out, ins := cCifTypeIO(fn)
	cifPrep(cif, abi, out, ins)
	return cif
}
func (cif *Cif) Close() {
	if cif == nil {
		return
	}

	CFree(uptr(cif.cCif))

	clear(cif.cIns)
	cif.cIns = nil
	cif.cOut = nil

	clear(*cif.goInsGeter)
	unpin(cif.goId)
	cif.goId = -1
	cif = nil
}

func CallCif[T any](cif *Cif, fn uptr, args []any) T {
	argc := len(args)

	size := argc * (8*2 + 8 + 8*2)
	io := CMalloc(uint64(size + int(cif.cCif.rtype.size) + 8))
	copy(unsafe.Slice((*any)(io), argc), args)
	defer CFree(io)

	gets, ityp := (*C.int)(nil), (**C.ffi_type)(nil)
	if len(args) != 0 {
		gets = (*C.int)(&cif.cinsGeter[0])
		ityp = &cif.cIns[0]
	}

	r := C.cccif_call(cif.cCif, fn,
		C.int(argc), (*C.go_any)(io), gets, ityp)

	var ret T
	_ = args
	ret = *((*T)(r))
	return ret
}

// func CallCif[T any](cif *Cif, fn uptr, args []any) T {
// 	var argc = len(args)
// 	var clears = []func(){}
// 	defer func() {
// 		for _, f := range clears {
// 			f()
// 		}
// 		clear(clears)
// 	}()

// 	insLen := argc*8 + int(cif.cCif.rtype.size)
// 	ins := CMalloc(uint64(insLen))
// 	defer CFree(ins)

// 	lst := unsafe.Slice((*uptr)(ins), argc)
// 	aptrs := make([]*any, argc)
// 	for i := 0; i < argc; i++ {
// 		ptr := &args[i]
// 		aptrs[i] = ptr
// 		lst[i] = (*cif.goInsGeter)[i](ptr, &clears)
// 	}

// 	var out T
// 	cout := uptr(iptr(ins) + iptr(argc*8))
// 	C.ffi_call(cif.cCif, (*[0]byte)(uptr(fn)), cout, (*uptr)(ins))
// 	_, _ = aptrs, args

// 	out = *(*T)(cout)
// 	return out
// }

// #endregion

// #region closure

type cClosure = C.ffi_closure

type reflectcls struct {
	real   any
	fn     reflect.Value
	ins    []func(uptr) reflect.Value
	out    func(uptr, []reflect.Value)
	doCall func(out, ins uptr)
}

func goGetter(t reflect.Type) func(uptr) reflect.Value {
	switch t.Kind() {
	case reflect.String:
		return func(addr uptr) reflect.Value {
			p := (*[65535]byte)(*(*uptr)(addr))
			l := 0
			for i := 0; i < 65535; i++ {
				if p[i] == 0 {
					break
				}
				l++
			}
			return reflect.ValueOf(string(p[:l]))
		}
	case reflect.Bool:
		return func(addr uptr) reflect.Value { return reflect.ValueOf(*(*int32)(addr) != 0) }
		// case reflect.Struct:
		// 	size := C.size_t(t.Size())
		// 	return func(src uptr) reflect.Value {
		// 		addr := reflect.New(t)
		// 		dst := addr.UnsafePointer()
		// 		C.memcpy(dst, src, size)
		// 		return addr.Elem()
		// 	}
	}

	// size := C.size_t(t.Size())
	// return func(src uptr) reflect.Value {
	// 	addr := reflect.New(t)
	// 	dst := addr.UnsafePointer()
	// 	C.memcpy(dst, src, size)
	// 	return addr.Elem()
	// }

	return func(addr uptr) reflect.Value {
		return reflect.NewAt(t, addr).Elem()
		// r := reflect.New(t).Elem()
		// reflect.Copy(r, reflect.NewAt(t, addr).Elem())
		// return r
	}
}

// func GoGetterAny(t reflect.Type) func(uptr) any {
// 	switch t.Kind() {
// 	case reflect.String:
// 		return func(addr uptr) any {
// 			p := (*[65535]byte)(*(*uptr)(addr))
// 			l := 0
// 			for i := 0; i < 65535; i++ {
// 				if p[i] == 0 {
// 					break
// 				}
// 				l++
// 			}
// 			return string(p[:l])
// 		}
// 	case reflect.Bool:
// 		return func(addr uptr) any { return (*(*int32)(addr) != 0) }

// 	case reflect.Struct:
// 		size := C.size_t(t.Size())
// 		return func(src uptr) any {
// 			addr := reflect.New(t)
// 			dst := addr.UnsafePointer()
// 			C.memcpy(dst, src, size)
// 			return addr.Elem().Interface()
// 		}
// 	}

//		return func(addr uptr) any {
//			return reflect.NewAt(t, addr).Elem().Interface()
//		}
//	}

func goSetter(fn reflect.Type) func(u uptr, vs []reflect.Value) {
	if fn.NumOut() == 0 {
		return func(u uptr, vs []reflect.Value) {
		}
	}
	switch fn.Out(0).Kind() {
	case reflect.Bool:
		return func(u uptr, vs []reflect.Value) {
			if vs[0].Bool() {
				*(*int32)(u) = 1
			} else {
				*(*int32)(u) = 0
			}
		}
	case reflect.String:
		return func(u uptr, vs []reflect.Value) {
			*(*uptr)(u) = uptr(C.CString(vs[0].String()))
		}
		// case reflect.Struct:
		// 	size := fn.Out(0).Size()
		// 	return func(u uptr, vs []reflect.Value) {

		// 		vs[0].Interface()

		// 		// addr := reflect.New(typ)
		// 		// dst := addr.UnsafePointer()
		// 		// C.memcpy(dst, src, size)
		// 		// return addr.Elem()

		// 		reflect.NewAt(fn.Out(0), u).Elem().Set(vs[0])
		// 	}

	}
	return func(u uptr, vs []reflect.Value) {
		reflect.NewAt(fn.Out(0), u).Elem().Set(vs[0])
	}
}

var cClosureSize = unsafe.Sizeof(cClosure{})

var codeHold = make(map[uptr]*Closure)
var codeHoldLock = sync.RWMutex{}

type Closure struct {
	cClosure         *cClosure
	cCode            unsafe.Pointer
	cCif             *Cif
	cID              *int
	goCls            *reflectcls
	goID             int
	goCloseTimestamp int64
	goCallingCount   int64
}

var zeroreflect reflect.Value = reflect.ValueOf(0)

func (cls *Closure) Code() unsafe.Pointer { return cls.cCode }
func (cls *Closure) Close() {
	if cls == nil {
		return
	}

	unpin(cls.goID)

	if cls.cClosure != nil {
		C.ffi_closure_free(uptr(cls.cClosure))
		cls.cClosure = nil
	}

	codeHoldLock.Lock()
	delete(codeHold, cls.cCode)
	codeHoldLock.Unlock()

	cls.cCode = nil
	if cls.cCif != nil {
		cls.cCif.Close()
		cls.cCif = nil
	}
	if cls.goCls != nil {
		cls.goCls.doCall = nil
		cls.goCls.out = nil
		cls.goCls.real = nil
		cls.goCls.fn = zeroreflect
		clear(cls.goCls.ins)
		cls.goCls.ins = nil
		cls.goCls.out = nil
		cls.goCls = nil
	}
	// cgoFree(uptr(cls.cID), 8)
	CFree(uptr(cls.cID))

	cls = nil
}
func ClosureAlloc() *Closure {
	code := new(uptr)
	cc := (*cClosure)(C.ffi_closure_alloc(C.size_t(cClosureSize), code))
	if cc == nil {
		return nil
	}

	cls := new(Closure)
	cls.cClosure = cc
	cls.cCode = *code
	cls.goCls = new(reflectcls)
	cls.goID = pin(cls)
	cls.cID = (*int)(CMalloc(8))
	(*cls.cID) = cls.goID
	cls.goCloseTimestamp = 0

	codeHoldLock.Lock()
	codeHold[cls.cCode] = cls
	codeHoldLock.Unlock()

	return cls
}
func (cls *Closure) ClosureLoc(fn any) {
	cls.goCls.real = fn
	cls.goCls.fn = reflect.ValueOf(fn)

	typ := cls.goCls.fn.Type()
	typNum := typ.NumIn()

	cls.goCls.ins = make([]func(uptr) reflect.Value, typNum)
	for i := 0; i < typNum; i++ {
		cls.goCls.ins[i] = goGetter(typ.In(i))
	}
	cls.goCls.out = goSetter(typ)

	if typNum == 0 {
		cls.goCls.doCall = func(out, ins uptr) {
			ret := cls.goCls.fn.Call(nil)
			cls.goCls.out(out, ret)
		}
	} else {
		cls.goCls.doCall = func(out, ins uptr) {
			argsPtr := unsafe.Slice((*uptr)(ins), typNum)
			argsVal := make([]reflect.Value, typNum)

			for i := 0; i < int(typNum); i++ {
				argsVal[i] = cls.goCls.ins[i](argsPtr[i])
			}
			ret := cls.goCls.fn.Call(argsVal)
			cls.goCls.out(out, ret)
		}
	}

	cif := cifAlloc()
	cifPrepAny(cif, AbiDefault, typ)
	cls.cCif = cif

	C.ffi_prep_closure_loc(cls.cClosure, cif.cCif,
		(*[0]byte)(C.closure_agent), uptr(cls.cID), cls.cCode)
}
func (cls *Closure) RestLocRaw(fn func(out, ins uptr)) {
	cls.goCls.doCall = fn
}
func ClosureCloseByCode(code uptr) {
	if code == nil {
		return
	}

	codeHoldLock.RLock()
	cls, ok := codeHold[code]
	codeHoldLock.RUnlock()

	if !ok {
		return
	}
	cls.Close()
}

var clsCloseQue = []*Closure{}
var clsCloseQueLock = sync.RWMutex{}

func ClosureCloseByCodeLate(code ...uptr) {
	clss := []*Closure{}
	t := time.Now().Unix()

	codeHoldLock.RLock()
	for i := range code {
		cls := codeHold[code[i]]
		cls.goCloseTimestamp = t
		clss = append(clss, cls)
	}
	codeHoldLock.RUnlock()

	clsCloseQueLock.Lock()
	clsCloseQue = append(clsCloseQue, clss...)
	clsCloseQueLock.Unlock()
}

func init() {
	go func() {
		for {
			time.Sleep(time.Second * 10)

			if len(clsCloseQue) == 0 {
				continue
			}

			keeps := []*Closure{}
			closes := []*Closure{}
			now := time.Now().Unix()

			clsCloseQueLock.Lock()
			for i := range clsCloseQue {
				cls := clsCloseQue[i]
				// if now-cls.goCloseTimestamp < 2 && atomic.LoadInt64(&cls.goCallingCount) == 0 {
				if now-cls.goCloseTimestamp < 2 {
					keeps = append(keeps, cls)
					continue
				}
				closes = append(closes, cls)
			}
			clsCloseQue = keeps
			clsCloseQueLock.Unlock()

			for i := range closes {
				closes[i].Close()
			}
		}

		// for {
		// 	time.Sleep(time.Second * 5)
		// }
	}()
}

//export closure_agent
func closure_agent(_, out, ins uptr, cId uptr) {
	cls := pinQuery[*Closure](*(*int)(cId))

	// atomic.AddInt64(&cls.goCallingCount, 1)
	// defer atomic.AddInt64(&cls.goCallingCount, -1)

	cls.goCls.doCall(out, ins)
	// outs := cls.goCls.doCall(ins)
	// cls.goCls.out(out, outs)

	// argc := cls.cCif.cCif.nargs
	// {
	// 	var outs []reflect.Value
	// 	if argc == 0 {
	// 		outs = cls.goCls.fn.Call(nil)
	// 	} else {
	// 		argsPtr := unsafe.Slice((*uptr)(ins), argc)
	// 		argsVal := make([]reflect.Value, argc)

	// 		for i := 0; i < int(argc); i++ {
	// 			argsVal[i] = cls.goCls.ins[i](argsPtr[i])
	// 		}
	// 		outs = cls.goCls.fn.Call(argsVal)
	// 	}
	// 	cls.goCls.out(out, outs)
	// }

}

func BindFunc[Ret any](funcAddr any, cPtr uptr) {

	typ := reflect.TypeOf(funcAddr).Elem()
	argc := typ.NumIn()
	cif := cifAlloc()
	cifPrepAny(cif, AbiDefault, typ)
	retfn := func(r Ret) []reflect.Value { return []reflect.Value{} }
	if typ.NumOut() > 0 {
		retfn = func(r Ret) []reflect.Value { return []reflect.Value{reflect.ValueOf(r)} }
	}

	rFunc := reflect.MakeFunc(typ, func(args []reflect.Value) (results []reflect.Value) {
		ins := make([]any, argc)
		for i := 0; i < argc; i++ {
			ins[i] = args[i].Interface()
		}

		r := CallCif[Ret](cif, cPtr, ins)
		return retfn(r)
	})
	reflect.ValueOf(funcAddr).Elem().Set(rFunc)
}

// #endregion

// #region holder

type hold struct {
	list  []any
	empty []int
}

var holder *hold
var holderLock = sync.RWMutex{}

func init() {
	holder = new(hold)
	holder.list = []any{}
	holder.empty = []int{}
}
func pin(a any) (id int) {
	holderLock.Lock()
	defer holderLock.Unlock()

	if len(holder.empty) == 0 {
		id = len(holder.list)
		holder.list = append(holder.list, a)
		return
	}

	size := len(holder.empty)
	id = holder.empty[size-1]
	holder.empty = holder.empty[:size-1]
	holder.list[id] = a
	return
}
func unpin(id int) {
	holderLock.Lock()
	defer holderLock.Unlock()

	holder.list[id] = nil
	holder.empty = append(holder.empty, id)
}
func pinQuery[T any](id int) T {
	holderLock.RLock()
	defer holderLock.RUnlock()
	return (holder.list[id]).(T)
}

// #endregion
