package tcc

import (
	"github.com/jinzhongmin/gg/cc"
)

type State struct{}

func New() *State { return tcc_new.Fn()() }

func (s *State) Delete() { tcc_delete.Fn()(s) }
func (s *State) SetLibPath(path string) {
	cstr := cc.NewString(path)
	defer cstr.Free()
	tcc_set_lib_path.Fn()(s, cstr)
}
func (s *State) SetErrorFunc(callback func(msg string)) {
	fn := func(_ uptr, cp cc.String) {
		callback(cp.String())
	}
	tcc_set_error_func.Fn()(s, nil, cc.Cbk(fn))
}
func (s *State) SetOptions(options string) {
	cstr := cc.NewString(options)
	defer cstr.Free()
	tcc_set_options.Fn()(s, cstr)
}

func (s *State) AddIncludePath(path string) int32 {
	cstr := cc.NewString(path)
	defer cstr.Free()
	return tcc_add_include_path.Fn()(s, cstr)
}
func (s *State) AddSysIncludePath(path string) int32 {
	cstr := cc.NewString(path)
	defer cstr.Free()
	return tcc_add_sysinclude_path.Fn()(s, cstr)
}
func (s *State) DefineSymbol(sym string, value string) {
	csym, cval := cc.NewString(sym), cc.NewString(value)
	defer csym.Free()
	defer cval.Free()
	tcc_define_symbol.Fn()(s, csym, cval)
}
func (s *State) UnDefineSymbol(sym string) {
	csym := cc.NewString(sym)
	defer csym.Free()
	tcc_undefine_symbol.Fn()(s, csym)
}

func (s *State) AddFile(filename string) int32 {
	cstr := cc.NewString(filename)
	defer cstr.Free()
	return tcc_add_file.Fn()(s, cstr)
}
func (s *State) CompileString(buf string) int32 {
	c := cc.NewString(buf)
	defer c.Free()
	return tcc_compile_string.Fn()(s, c)
}

type OutputType int32

const (
	_ OutputType = iota
	OutputMemory
	OutputExe
	OutputDll
	OutputObj
	OutputPreprocess
)

func (s *State) SetOutputType(t OutputType) int32 { return tcc_set_output_type.Fn()(s, t) }
func (s *State) AddLibraryPath(path string) int32 {
	cstr := cc.NewString(path)
	defer cstr.Free()
	return tcc_add_library_path.Fn()(s, cstr)
}
func (s *State) AddLibrary(libname string) int32 {
	cstr := cc.NewString(libname)
	defer cstr.Free()
	return tcc_add_library.Fn()(s, cstr)
}
func (s *State) AddSymbol(name, val string) int32 {
	cname, cval := cc.NewString(name), cc.NewString(val)
	defer cname.Free()
	defer cval.Free()
	return tcc_add_symbol.Fn()(s, cname, cval)
}
func (s *State) OutputFile(filename string) int32 {
	cstr := cc.NewString(filename)
	defer cstr.Free()
	return tcc_output_file.Fn()(s, cstr)
}

func (s *State) Run(args []string) int32 {
	strs := cc.NewStrings(args)
	defer strs.Free()
	return tcc_run.Fn()(s, int32(len(args)), strs)
}
func (s *State) Relocate(p uptr) int32 { return tcc_relocate.Fn()(s, p) }

func (s *State) GetSymbol(sym string) uptr {
	cstr := cc.NewString(sym)
	defer cstr.Free()
	return tcc_get_symbol.Fn()(s, cstr)
}

var RelocateAuto = uptr(iptr(1))
