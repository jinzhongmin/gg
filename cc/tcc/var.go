package tcc

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
)

type uptr = unsafe.Pointer
type iptr = uintptr

func init() { cc.Open("libtcc*") }

var (
	tcc_relocate   = cc.DlFunc[func(*State, uptr) int32, int32]{Name: "tcc_relocate"}
	tcc_get_symbol = cc.DlFunc[func(*State, cc.String) uptr, uptr]{Name: "tcc_get_symbol"}

	tcc_new            = cc.DlFunc[func() *State, *State]{Name: "tcc_new"}
	tcc_delete         = cc.DlFunc[func(*State), cc.Void]{Name: "tcc_delete"}
	tcc_set_lib_path   = cc.DlFunc[func(*State, cc.String), cc.Void]{Name: "tcc_set_lib_path"}
	tcc_set_error_func = cc.DlFunc[func(*State, uptr, uptr), cc.Void]{Name: "tcc_set_error_func"}
	tcc_set_options    = cc.DlFunc[func(*State, cc.String), cc.Void]{Name: "tcc_set_options"}

	tcc_add_include_path    = cc.DlFunc[func(*State, cc.String) int32, int32]{Name: "tcc_add_include_path"}
	tcc_add_sysinclude_path = cc.DlFunc[func(*State, cc.String) int32, int32]{Name: "tcc_add_sysinclude_path"}
	tcc_define_symbol       = cc.DlFunc[func(*State, cc.String, cc.String), cc.Void]{Name: "tcc_define_symbol"}
	tcc_undefine_symbol     = cc.DlFunc[func(*State, cc.String), cc.Void]{Name: "tcc_undefine_symbol"}

	tcc_add_file       = cc.DlFunc[func(*State, cc.String) int32, int32]{Name: "tcc_add_file"}
	tcc_compile_string = cc.DlFunc[func(*State, cc.String) int32, int32]{Name: "tcc_compile_string"}

	tcc_set_output_type  = cc.DlFunc[func(*State, OutputType) int32, int32]{Name: "tcc_set_output_type"}
	tcc_add_library_path = cc.DlFunc[func(*State, cc.String) int32, int32]{Name: "tcc_add_library_path"}
	tcc_add_library      = cc.DlFunc[func(*State, cc.String) int32, int32]{Name: "tcc_add_library"}
	tcc_add_symbol       = cc.DlFunc[func(*State, cc.String, cc.String) int32, int32]{Name: "tcc_add_symbol"}
	tcc_output_file      = cc.DlFunc[func(*State, cc.String) int32, int32]{Name: "tcc_output_file"}
	tcc_run              = cc.DlFunc[func(*State, int32, cc.Strings) int32, int32]{Name: "tcc_run"}
)
