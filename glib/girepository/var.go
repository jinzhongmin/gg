package girepository

import (
	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/glib/glib"
	"github.com/jinzhongmin/gg/glib/gobject"
)

func init() { cc.Open("libgirepository-2*") }

var (
	gi_arg_info_get_direction          = cc.DlFunc[func(*ArgInfo) Direction, Direction]{Name: "gi_arg_info_get_direction"}
	gi_arg_info_is_return_value        = cc.DlFunc[func(*ArgInfo) int32, int32]{Name: "gi_arg_info_is_return_value"}
	gi_arg_info_is_optional            = cc.DlFunc[func(*ArgInfo) int32, int32]{Name: "gi_arg_info_is_optional"}
	gi_arg_info_is_caller_allocates    = cc.DlFunc[func(*ArgInfo) int32, int32]{Name: "gi_arg_info_is_caller_allocates"}
	gi_arg_info_may_be_null            = cc.DlFunc[func(*ArgInfo) int32, int32]{Name: "gi_arg_info_may_be_null"}
	gi_arg_info_is_skip                = cc.DlFunc[func(*ArgInfo) int32, int32]{Name: "gi_arg_info_is_skip"}
	gi_arg_info_get_ownership_transfer = cc.DlFunc[func(*ArgInfo) Transfer, Transfer]{Name: "gi_arg_info_get_ownership_transfer"}
	gi_arg_info_get_scope              = cc.DlFunc[func(*ArgInfo) ScopeType, ScopeType]{Name: "gi_arg_info_get_scope"}
	gi_arg_info_get_closure_index      = cc.DlFunc[func(*ArgInfo, *uint32) int32, int32]{Name: "gi_arg_info_get_closure_index"}
	gi_arg_info_get_destroy_index      = cc.DlFunc[func(*ArgInfo, *uint32) int32, int32]{Name: "gi_arg_info_get_destroy_index"}
	gi_arg_info_get_type_info          = cc.DlFunc[func(*ArgInfo) *TypeInfo, *TypeInfo]{Name: "gi_arg_info_get_type_info"}
	gi_arg_info_load_type_info         = cc.DlFunc[func(*ArgInfo, *TypeInfo), cc.Void]{Name: "gi_arg_info_load_type_info"}

	gi_base_info_ref                = cc.DlFunc[func(*BaseInfo) *BaseInfo, *BaseInfo]{Name: "gi_base_info_ref"}
	gi_base_info_unref              = cc.DlFunc[func(*BaseInfo), cc.Void]{Name: "gi_base_info_unref"}
	gi_base_info_clear              = cc.DlFunc[func(*BaseInfo), cc.Void]{Name: "gi_base_info_clear"}
	gi_base_info_get_name           = cc.DlFunc[func(*BaseInfo) cc.String, cc.String]{Name: "gi_base_info_get_name"}
	gi_base_info_get_namespace      = cc.DlFunc[func(*BaseInfo) cc.String, cc.String]{Name: "gi_base_info_get_namespace"}
	gi_base_info_is_deprecated      = cc.DlFunc[func(*BaseInfo) int32, int32]{Name: "gi_base_info_is_deprecated"}
	gi_base_info_get_attribute      = cc.DlFunc[func(*BaseInfo, cc.String) cc.String, cc.String]{Name: "gi_base_info_get_attribute"}
	gi_base_info_iterate_attributes = cc.DlFunc[func(*BaseInfo, *AttributeIter, *cc.String, *cc.String) int32, int32]{Name: "gi_base_info_iterate_attributes"}
	gi_base_info_get_container      = cc.DlFunc[func(*BaseInfo) *BaseInfo, *BaseInfo]{Name: "gi_base_info_get_container"}
	gi_base_info_get_typelib        = cc.DlFunc[func(*BaseInfo) *Typelib, *Typelib]{Name: "gi_base_info_get_typelib"}
	gi_base_info_equal              = cc.DlFunc[func(*BaseInfo, *BaseInfo) int32, int32]{Name: "gi_base_info_equal"}

	gi_callable_info_is_method                 = cc.DlFunc[func(*CallableInfo) int32, int32]{Name: "gi_callable_info_is_method"}
	gi_callable_info_can_throw_gerror          = cc.DlFunc[func(*CallableInfo) int32, int32]{Name: "gi_callable_info_can_throw_gerror"}
	gi_callable_info_get_return_type           = cc.DlFunc[func(*CallableInfo) *TypeInfo, *TypeInfo]{Name: "gi_callable_info_get_return_type"}
	gi_callable_info_load_return_type          = cc.DlFunc[func(*CallableInfo, *TypeInfo), cc.Void]{Name: "gi_callable_info_load_return_type"}
	gi_callable_info_get_return_attribute      = cc.DlFunc[func(*CallableInfo, cc.String) cc.String, cc.String]{Name: "gi_callable_info_get_return_attribute"}
	gi_callable_info_iterate_return_attributes = cc.DlFunc[func(*CallableInfo, *AttributeIter, *cc.String, *cc.String) int32, int32]{Name: "gi_callable_info_iterate_return_attributes"}
	gi_callable_info_get_caller_owns           = cc.DlFunc[func(*CallableInfo) Transfer, Transfer]{Name: "gi_callable_info_get_caller_owns"}
	gi_callable_info_may_return_null           = cc.DlFunc[func(*CallableInfo) int32, int32]{Name: "gi_callable_info_may_return_null"}
	gi_callable_info_skip_return               = cc.DlFunc[func(*CallableInfo) int32, int32]{Name: "gi_callable_info_skip_return"}
	gi_callable_info_get_n_args                = cc.DlFunc[func(*CallableInfo) uint32, uint32]{Name: "gi_callable_info_get_n_args"}
	gi_callable_info_get_arg                   = cc.DlFunc[func(*CallableInfo, uint32) *ArgInfo, *ArgInfo]{Name: "gi_callable_info_get_arg"}
	gi_callable_info_load_arg                  = cc.DlFunc[func(*CallableInfo, uint32, *ArgInfo), cc.Void]{Name: "gi_callable_info_load_arg"}
	// gi_callable_info_invoke                          func(*CallableInfo, uptr, uptr, uint, uptr, uint, uptr, **glib.GError) int32
	gi_callable_info_get_instance_ownership_transfer = cc.DlFunc[func(*CallableInfo) Transfer, Transfer]{Name: "gi_callable_info_get_instance_ownership_transfer"}

	gi_constant_info_get_type_info = cc.DlFunc[func(*ConstantInfo) *TypeInfo, *TypeInfo]{Name: "gi_constant_info_get_type_info"}
	gi_constant_info_free_value    = cc.DlFunc[func(*ConstantInfo, *Argument), cc.Void]{Name: "gi_constant_info_free_value"}
	gi_constant_info_get_value     = cc.DlFunc[func(*ConstantInfo, *Argument) uint, uint]{Name: "gi_constant_info_get_value"}

	gi_enum_info_get_n_values     = cc.DlFunc[func(*EnumInfo) uint32, uint32]{Name: "gi_enum_info_get_n_values"}
	gi_enum_info_get_value        = cc.DlFunc[func(*EnumInfo, uint32) *ValueInfo, *ValueInfo]{Name: "gi_enum_info_get_value"}
	gi_enum_info_get_n_methods    = cc.DlFunc[func(*EnumInfo) uint32, uint32]{Name: "gi_enum_info_get_n_methods"}
	gi_enum_info_get_method       = cc.DlFunc[func(*EnumInfo, uint32) *FunctionInfo, *FunctionInfo]{Name: "gi_enum_info_get_method"}
	gi_enum_info_get_storage_type = cc.DlFunc[func(*EnumInfo) TypeTag, TypeTag]{Name: "gi_enum_info_get_storage_type"}
	gi_enum_info_get_error_domain = cc.DlFunc[func(*EnumInfo) cc.String, cc.String]{Name: "gi_enum_info_get_error_domain"}

	gi_field_info_get_flags     = cc.DlFunc[func(*FieldInfo) FieldInfoFlags, FieldInfoFlags]{Name: "gi_field_info_get_flags"}
	gi_field_info_get_size      = cc.DlFunc[func(*FieldInfo) uint, uint]{Name: "gi_field_info_get_size"}
	gi_field_info_get_offset    = cc.DlFunc[func(*FieldInfo) uint, uint]{Name: "gi_field_info_get_offset"}
	gi_field_info_get_type_info = cc.DlFunc[func(*FieldInfo) *TypeInfo, *TypeInfo]{Name: "gi_field_info_get_type_info"}
	gi_field_info_get_field     = cc.DlFunc[func(*FieldInfo, uptr, *Argument) int32, int32]{Name: "gi_field_info_get_field"}
	gi_field_info_set_field     = cc.DlFunc[func(*FieldInfo, uptr, *Argument) int32, int32]{Name: "gi_field_info_set_field"}

	gi_function_info_get_symbol   = cc.DlFunc[func(*FunctionInfo) cc.String, cc.String]{Name: "gi_function_info_get_symbol"}
	gi_function_info_get_flags    = cc.DlFunc[func(*FunctionInfo) FunctionInfoFlags, FunctionInfoFlags]{Name: "gi_function_info_get_flags"}
	gi_function_info_get_property = cc.DlFunc[func(*FunctionInfo) *PropertyInfo, *PropertyInfo]{Name: "gi_function_info_get_property"}
	gi_function_info_get_vfunc    = cc.DlFunc[func(*FunctionInfo) *VFuncInfo, *VFuncInfo]{Name: "gi_function_info_get_vfunc"}
	// gi_function_info_invoke       func(*FunctionInfo, *Argument, uint, *Argument, uint, *Argument, **glib.GError) int32
	gi_invoke_error_quark = cc.DlFunc[func() glib.GQuark, glib.GQuark]{Name: "gi_invoke_error_quark"}

	gi_interface_info_get_n_prerequisites = cc.DlFunc[func(*InterfaceInfo) uint32, uint32]{Name: "gi_interface_info_get_n_prerequisites"}
	gi_interface_info_get_prerequisite    = cc.DlFunc[func(*InterfaceInfo, uint32) *BaseInfo, *BaseInfo]{Name: "gi_interface_info_get_prerequisite"}
	gi_interface_info_get_n_properties    = cc.DlFunc[func(*InterfaceInfo) uint32, uint32]{Name: "gi_interface_info_get_n_properties"}
	gi_interface_info_get_property        = cc.DlFunc[func(*InterfaceInfo, uint32) *PropertyInfo, *PropertyInfo]{Name: "gi_interface_info_get_property"}
	gi_interface_info_get_n_methods       = cc.DlFunc[func(*InterfaceInfo) uint32, uint32]{Name: "gi_interface_info_get_n_methods"}
	gi_interface_info_get_method          = cc.DlFunc[func(*InterfaceInfo, uint32) *FunctionInfo, *FunctionInfo]{Name: "gi_interface_info_get_method"}
	gi_interface_info_find_method         = cc.DlFunc[func(*InterfaceInfo, cc.String) *FunctionInfo, *FunctionInfo]{Name: "gi_interface_info_find_method"}
	gi_interface_info_get_n_signals       = cc.DlFunc[func(*InterfaceInfo) uint32, uint32]{Name: "gi_interface_info_get_n_signals"}
	gi_interface_info_get_signal          = cc.DlFunc[func(*InterfaceInfo, uint32) *SignalInfo, *SignalInfo]{Name: "gi_interface_info_get_signal"}
	gi_interface_info_find_signal         = cc.DlFunc[func(*InterfaceInfo, cc.String) *SignalInfo, *SignalInfo]{Name: "gi_interface_info_find_signal"}
	gi_interface_info_get_n_vfuncs        = cc.DlFunc[func(*InterfaceInfo) uint32, uint32]{Name: "gi_interface_info_get_n_vfuncs"}
	gi_interface_info_get_vfunc           = cc.DlFunc[func(*InterfaceInfo, uint32) *VFuncInfo, *VFuncInfo]{Name: "gi_interface_info_get_vfunc"}
	gi_interface_info_find_vfunc          = cc.DlFunc[func(*InterfaceInfo, cc.String) *VFuncInfo, *VFuncInfo]{Name: "gi_interface_info_find_vfunc"}
	gi_interface_info_get_n_constants     = cc.DlFunc[func(*InterfaceInfo) uint32, uint32]{Name: "gi_interface_info_get_n_constants"}
	gi_interface_info_get_constant        = cc.DlFunc[func(*InterfaceInfo, uint32) *ConstantInfo, *ConstantInfo]{Name: "gi_interface_info_get_constant"}
	gi_interface_info_get_iface_struct    = cc.DlFunc[func(*InterfaceInfo) *StructInfo, *StructInfo]{Name: "gi_interface_info_get_iface_struct"}

	gi_object_info_get_type_name                  = cc.DlFunc[func(*ObjectInfo) cc.String, cc.String]{Name: "gi_object_info_get_type_name"}
	gi_object_info_get_type_init_function_name    = cc.DlFunc[func(*ObjectInfo) cc.String, cc.String]{Name: "gi_object_info_get_type_init_function_name"}
	gi_object_info_get_abstract                   = cc.DlFunc[func(*ObjectInfo) int32, int32]{Name: "gi_object_info_get_abstract"}
	gi_object_info_get_final                      = cc.DlFunc[func(*ObjectInfo) int32, int32]{Name: "gi_object_info_get_final"}
	gi_object_info_get_fundamental                = cc.DlFunc[func(*ObjectInfo) int32, int32]{Name: "gi_object_info_get_fundamental"}
	gi_object_info_get_parent                     = cc.DlFunc[func(*ObjectInfo) *ObjectInfo, *ObjectInfo]{Name: "gi_object_info_get_parent"}
	gi_object_info_get_n_interfaces               = cc.DlFunc[func(*ObjectInfo) uint32, uint32]{Name: "gi_object_info_get_n_interfaces"}
	gi_object_info_get_interface                  = cc.DlFunc[func(*ObjectInfo, uint32) *InterfaceInfo, *InterfaceInfo]{Name: "gi_object_info_get_interface"}
	gi_object_info_get_n_fields                   = cc.DlFunc[func(*ObjectInfo) uint32, uint32]{Name: "gi_object_info_get_n_fields"}
	gi_object_info_get_field                      = cc.DlFunc[func(*ObjectInfo, uint32) *FieldInfo, *FieldInfo]{Name: "gi_object_info_get_field"}
	gi_object_info_get_n_properties               = cc.DlFunc[func(*ObjectInfo) uint32, uint32]{Name: "gi_object_info_get_n_properties"}
	gi_object_info_get_property                   = cc.DlFunc[func(*ObjectInfo, uint32) *PropertyInfo, *PropertyInfo]{Name: "gi_object_info_get_property"}
	gi_object_info_get_n_methods                  = cc.DlFunc[func(*ObjectInfo) uint32, uint32]{Name: "gi_object_info_get_n_methods"}
	gi_object_info_get_method                     = cc.DlFunc[func(*ObjectInfo, uint32) *FunctionInfo, *FunctionInfo]{Name: "gi_object_info_get_method"}
	gi_object_info_find_method                    = cc.DlFunc[func(*ObjectInfo, cc.String) *FunctionInfo, *FunctionInfo]{Name: "gi_object_info_find_method"}
	gi_object_info_find_method_using_interfaces   = cc.DlFunc[func(*ObjectInfo, cc.String, **BaseInfo) *FunctionInfo, *FunctionInfo]{Name: "gi_object_info_find_method_using_interfaces"}
	gi_object_info_get_n_signals                  = cc.DlFunc[func(*ObjectInfo) uint32, uint32]{Name: "gi_object_info_get_n_signals"}
	gi_object_info_get_signal                     = cc.DlFunc[func(*ObjectInfo, uint32) *SignalInfo, *SignalInfo]{Name: "gi_object_info_get_signal"}
	gi_object_info_find_signal                    = cc.DlFunc[func(*ObjectInfo, cc.String) *SignalInfo, *SignalInfo]{Name: "gi_object_info_find_signal"}
	gi_object_info_get_n_vfuncs                   = cc.DlFunc[func(*ObjectInfo) uint32, uint32]{Name: "gi_object_info_get_n_vfuncs"}
	gi_object_info_get_vfunc                      = cc.DlFunc[func(*ObjectInfo, uint32) *VFuncInfo, *VFuncInfo]{Name: "gi_object_info_get_vfunc"}
	gi_object_info_find_vfunc                     = cc.DlFunc[func(*ObjectInfo, cc.String) *VFuncInfo, *VFuncInfo]{Name: "gi_object_info_find_vfunc"}
	gi_object_info_find_vfunc_using_interfaces    = cc.DlFunc[func(*ObjectInfo, cc.String, **BaseInfo) *VFuncInfo, *VFuncInfo]{Name: "gi_object_info_find_vfunc_using_interfaces"}
	gi_object_info_get_n_constants                = cc.DlFunc[func(*ObjectInfo) uint32, uint32]{Name: "gi_object_info_get_n_constants"}
	gi_object_info_get_constant                   = cc.DlFunc[func(*ObjectInfo, uint32) *ConstantInfo, *ConstantInfo]{Name: "gi_object_info_get_constant"}
	gi_object_info_get_class_struct               = cc.DlFunc[func(*ObjectInfo) *StructInfo, *StructInfo]{Name: "gi_object_info_get_class_struct"}
	gi_object_info_get_ref_function_name          = cc.DlFunc[func(*ObjectInfo) cc.String, cc.String]{Name: "gi_object_info_get_ref_function_name"}
	gi_object_info_get_ref_function_pointer       = cc.DlFunc[func(*ObjectInfo) uptr, uptr]{Name: "gi_object_info_get_ref_function_pointer"}
	gi_object_info_get_unref_function_name        = cc.DlFunc[func(*ObjectInfo) cc.String, cc.String]{Name: "gi_object_info_get_unref_function_name"}
	gi_object_info_get_unref_function_pointer     = cc.DlFunc[func(*ObjectInfo) uptr, uptr]{Name: "gi_object_info_get_unref_function_pointer"}
	gi_object_info_get_set_value_function_name    = cc.DlFunc[func(*ObjectInfo) cc.String, cc.String]{Name: "gi_object_info_get_set_value_function_name"}
	gi_object_info_get_set_value_function_pointer = cc.DlFunc[func(*ObjectInfo) uptr, uptr]{Name: "gi_object_info_get_set_value_function_pointer"}
	gi_object_info_get_get_value_function_name    = cc.DlFunc[func(*ObjectInfo) cc.String, cc.String]{Name: "gi_object_info_get_get_value_function_name"}
	gi_object_info_get_get_value_function_pointer = cc.DlFunc[func(*ObjectInfo) uptr, uptr]{Name: "gi_object_info_get_get_value_function_pointer"}

	gi_property_info_get_flags              = cc.DlFunc[func(*PropertyInfo) int32, int32]{Name: "gi_property_info_get_flags"}
	gi_property_info_get_type_info          = cc.DlFunc[func(*PropertyInfo) *TypeInfo, *TypeInfo]{Name: "gi_property_info_get_type_info"}
	gi_property_info_get_ownership_transfer = cc.DlFunc[func(*PropertyInfo) int32, int32]{Name: "gi_property_info_get_ownership_transfer"}
	gi_property_info_get_setter             = cc.DlFunc[func(*PropertyInfo) *FunctionInfo, *FunctionInfo]{Name: "gi_property_info_get_setter"}
	gi_property_info_get_getter             = cc.DlFunc[func(*PropertyInfo) *FunctionInfo, *FunctionInfo]{Name: "gi_property_info_get_getter"}

	gi_registered_type_info_get_type_name               = cc.DlFunc[func(*RegisteredTypeInfo) cc.String, cc.String]{Name: "gi_registered_type_info_get_type_name"}
	gi_registered_type_info_get_type_init_function_name = cc.DlFunc[func(*RegisteredTypeInfo) cc.String, cc.String]{Name: "gi_registered_type_info_get_type_init_function_name"}
	gi_registered_type_info_get_g_type                  = cc.DlFunc[func(*RegisteredTypeInfo) gobject.GType, gobject.GType]{Name: "gi_registered_type_info_get_g_type"}
	gi_registered_type_info_is_boxed                    = cc.DlFunc[func(*RegisteredTypeInfo) int32, int32]{Name: "gi_registered_type_info_is_boxed"}

	gi_repository_new                         = cc.DlFunc[func() *Repository, *Repository]{Name: "gi_repository_new"}
	gi_repository_prepend_search_path         = cc.DlFunc[func(*Repository, cc.String), cc.Void]{Name: "gi_repository_prepend_search_path"}
	gi_repository_prepend_library_path        = cc.DlFunc[func(*Repository, cc.String), cc.Void]{Name: "gi_repository_prepend_library_path"}
	gi_repository_get_search_path             = cc.DlFunc[func(*Repository, *uint64) cc.Strings, cc.Strings]{Name: "gi_repository_get_search_path"}
	gi_repository_get_library_path            = cc.DlFunc[func(*Repository, *uint64) cc.Strings, cc.Strings]{Name: "gi_repository_get_library_path"}
	gi_repository_load_typelib                = cc.DlFunc[func(*Repository, *Typelib, RepositoryLoadFlags, **glib.GError) cc.String, cc.String]{Name: "gi_repository_load_typelib"}
	gi_repository_is_registered               = cc.DlFunc[func(*Repository, cc.String, cc.String) int32, int32]{Name: "gi_repository_is_registered"}
	gi_repository_find_by_name                = cc.DlFunc[func(*Repository, cc.String, cc.String) *BaseInfo, *BaseInfo]{Name: "gi_repository_find_by_name"}
	gi_repository_enumerate_versions          = cc.DlFunc[func(*Repository, cc.String, *uint64) cc.Strings, cc.Strings]{Name: "gi_repository_enumerate_versions"}
	gi_repository_require                     = cc.DlFunc[func(*Repository, cc.String, cc.String, RepositoryLoadFlags, **glib.GError) *Typelib, *Typelib]{Name: "gi_repository_require"}
	gi_repository_require_private             = cc.DlFunc[func(*Repository, cc.String, cc.String, cc.String, RepositoryLoadFlags, **glib.GError) *Typelib, *Typelib]{Name: "gi_repository_require_private"}
	gi_repository_get_immediate_dependencies  = cc.DlFunc[func(*Repository, cc.String, *uint64) cc.Strings, cc.Strings]{Name: "gi_repository_get_immediate_dependencies"}
	gi_repository_get_dependencies            = cc.DlFunc[func(*Repository, cc.String, *uint64) cc.Strings, cc.Strings]{Name: "gi_repository_get_dependencies"}
	gi_repository_get_loaded_namespaces       = cc.DlFunc[func(*Repository, *uint64) cc.Strings, cc.Strings]{Name: "gi_repository_get_loaded_namespaces"}
	gi_repository_find_by_gtype               = cc.DlFunc[func(*Repository, gobject.GType) *BaseInfo, *BaseInfo]{Name: "gi_repository_find_by_gtype"}
	gi_repository_get_object_gtype_interfaces = cc.DlFunc[func(*Repository, gobject.GType, *uint, ***InterfaceInfo), cc.Void]{Name: "gi_repository_get_object_gtype_interfaces"}
	gi_repository_get_n_infos                 = cc.DlFunc[func(*Repository, cc.String) uint32, uint32]{Name: "gi_repository_get_n_infos"}
	gi_repository_get_info                    = cc.DlFunc[func(*Repository, cc.String, uint32) *BaseInfo, *BaseInfo]{Name: "gi_repository_get_info"}
	gi_repository_find_by_error_domain        = cc.DlFunc[func(*Repository, glib.GQuark) *EnumInfo, *EnumInfo]{Name: "gi_repository_find_by_error_domain"}
	gi_repository_get_typelib_path            = cc.DlFunc[func(*Repository, cc.String) cc.String, cc.String]{Name: "gi_repository_get_typelib_path"}
	gi_repository_get_shared_libraries        = cc.DlFunc[func(*Repository, cc.String, *uint64) cc.Strings, cc.Strings]{Name: "gi_repository_get_shared_libraries"}
	gi_repository_get_c_prefix                = cc.DlFunc[func(*Repository, cc.String) cc.String, cc.String]{Name: "gi_repository_get_c_prefix"}
	gi_repository_get_version                 = cc.DlFunc[func(*Repository, cc.String) cc.String, cc.String]{Name: "gi_repository_get_version"}
	gi_repository_get_option_group            = cc.DlFunc[func() *glib.GOptionGroup, *glib.GOptionGroup]{Name: "gi_repository_get_option_group"}
	gi_repository_dump                        = cc.DlFunc[func(cc.String, cc.String, **glib.GError) int32, int32]{Name: "gi_repository_dump"}
	gi_repository_error_quark                 = cc.DlFunc[func() glib.GQuark, glib.GQuark]{Name: "gi_repository_error_quark"}

	gi_signal_info_get_flags         = cc.DlFunc[func(*SignalInfo) int32, int32]{Name: "gi_signal_info_get_flags"}
	gi_signal_info_get_class_closure = cc.DlFunc[func(*SignalInfo) *VFuncInfo, *VFuncInfo]{Name: "gi_signal_info_get_class_closure"}
	gi_signal_info_true_stops_emit   = cc.DlFunc[func(*SignalInfo) int32, int32]{Name: "gi_signal_info_true_stops_emit"}

	gi_struct_info_get_n_fields           = cc.DlFunc[func(*StructInfo) uint32, uint32]{Name: "gi_struct_info_get_n_fields"}
	gi_struct_info_get_field              = cc.DlFunc[func(*StructInfo, uint32) *FieldInfo, *FieldInfo]{Name: "gi_struct_info_get_field"}
	gi_struct_info_find_field             = cc.DlFunc[func(*StructInfo, cc.String) *FieldInfo, *FieldInfo]{Name: "gi_struct_info_find_field"}
	gi_struct_info_get_n_methods          = cc.DlFunc[func(*StructInfo) uint32, uint32]{Name: "gi_struct_info_get_n_methods"}
	gi_struct_info_get_method             = cc.DlFunc[func(*StructInfo, uint32) *FunctionInfo, *FunctionInfo]{Name: "gi_struct_info_get_method"}
	gi_struct_info_find_method            = cc.DlFunc[func(*StructInfo, cc.String) *FunctionInfo, *FunctionInfo]{Name: "gi_struct_info_find_method"}
	gi_struct_info_get_size               = cc.DlFunc[func(*StructInfo) uint, uint]{Name: "gi_struct_info_get_size"}
	gi_struct_info_get_alignment          = cc.DlFunc[func(*StructInfo) uint, uint]{Name: "gi_struct_info_get_alignment"}
	gi_struct_info_is_gtype_struct        = cc.DlFunc[func(*StructInfo) int32, int32]{Name: "gi_struct_info_is_gtype_struct"}
	gi_struct_info_is_foreign             = cc.DlFunc[func(*StructInfo) int32, int32]{Name: "gi_struct_info_is_foreign"}
	gi_struct_info_get_copy_function_name = cc.DlFunc[func(*StructInfo) cc.String, cc.String]{Name: "gi_struct_info_get_copy_function_name"}
	gi_struct_info_get_free_function_name = cc.DlFunc[func(*StructInfo) cc.String, cc.String]{Name: "gi_struct_info_get_free_function_name"}

	gi_type_tag_to_string                   = cc.DlFunc[func(int32) cc.String, cc.String]{Name: "gi_type_tag_to_string"}
	gi_type_info_is_pointer                 = cc.DlFunc[func(*TypeInfo) int32, int32]{Name: "gi_type_info_is_pointer"}
	gi_type_info_get_tag                    = cc.DlFunc[func(*TypeInfo) int32, int32]{Name: "gi_type_info_get_tag"}
	gi_type_info_get_param_type             = cc.DlFunc[func(*TypeInfo, uint32) *TypeInfo, *TypeInfo]{Name: "gi_type_info_get_param_type"}
	gi_type_info_get_interface              = cc.DlFunc[func(*TypeInfo) *BaseInfo, *BaseInfo]{Name: "gi_type_info_get_interface"}
	gi_type_info_get_array_length_index     = cc.DlFunc[func(*TypeInfo, *uint32) int32, int32]{Name: "gi_type_info_get_array_length_index"}
	gi_type_info_get_array_fixed_size       = cc.DlFunc[func(*TypeInfo, *uint) int32, int32]{Name: "gi_type_info_get_array_fixed_size"}
	gi_type_info_is_zero_terminated         = cc.DlFunc[func(*TypeInfo) int32, int32]{Name: "gi_type_info_is_zero_terminated"}
	gi_type_info_get_array_type             = cc.DlFunc[func(*TypeInfo) int32, int32]{Name: "gi_type_info_get_array_type"}
	gi_type_info_get_storage_type           = cc.DlFunc[func(*TypeInfo) int32, int32]{Name: "gi_type_info_get_storage_type"}
	gi_type_info_argument_from_hash_pointer = cc.DlFunc[func(*TypeInfo, uptr, *Argument), cc.Void]{Name: "gi_type_info_argument_from_hash_pointer"}
	gi_type_info_hash_pointer_from_argument = cc.DlFunc[func(*TypeInfo, *Argument) uptr, uptr]{Name: "gi_type_info_hash_pointer_from_argument"}
	gi_type_tag_argument_from_hash_pointer  = cc.DlFunc[func(int32, uptr, *Argument), cc.Void]{Name: "gi_type_tag_argument_from_hash_pointer"}
	gi_type_tag_hash_pointer_from_argument  = cc.DlFunc[func(int32, *Argument) uptr, uptr]{Name: "gi_type_tag_hash_pointer_from_argument"}

	gi_union_info_get_n_fields             = cc.DlFunc[func(*UnionInfo) uint32, uint32]{Name: "gi_union_info_get_n_fields"}
	gi_union_info_get_field                = cc.DlFunc[func(*UnionInfo, uint32) *FieldInfo, *FieldInfo]{Name: "gi_union_info_get_field"}
	gi_union_info_get_n_methods            = cc.DlFunc[func(*UnionInfo) uint32, uint32]{Name: "gi_union_info_get_n_methods"}
	gi_union_info_get_method               = cc.DlFunc[func(*UnionInfo, uint32) *FunctionInfo, *FunctionInfo]{Name: "gi_union_info_get_method"}
	gi_union_info_is_discriminated         = cc.DlFunc[func(*UnionInfo) int32, int32]{Name: "gi_union_info_is_discriminated"}
	gi_union_info_get_discriminator_offset = cc.DlFunc[func(*UnionInfo, *uint) int32, int32]{Name: "gi_union_info_get_discriminator_offset"}
	gi_union_info_get_discriminator_type   = cc.DlFunc[func(*UnionInfo) *TypeInfo, *TypeInfo]{Name: "gi_union_info_get_discriminator_type"}
	gi_union_info_get_discriminator        = cc.DlFunc[func(*UnionInfo, uint) *ConstantInfo, *ConstantInfo]{Name: "gi_union_info_get_discriminator"}
	gi_union_info_find_method              = cc.DlFunc[func(*UnionInfo, cc.String) *FunctionInfo, *FunctionInfo]{Name: "gi_union_info_find_method"}
	gi_union_info_get_size                 = cc.DlFunc[func(*UnionInfo) uint, uint]{Name: "gi_union_info_get_size"}
	gi_union_info_get_alignment            = cc.DlFunc[func(*UnionInfo) uint, uint]{Name: "gi_union_info_get_alignment"}
	gi_union_info_get_copy_function_name   = cc.DlFunc[func(*UnionInfo) cc.String, cc.String]{Name: "gi_union_info_get_copy_function_name"}
	gi_union_info_get_free_function_name   = cc.DlFunc[func(*UnionInfo) cc.String, cc.String]{Name: "gi_union_info_get_free_function_name"}

	gi_value_info_get_value = cc.DlFunc[func(*ValueInfo) int64, int64]{Name: "gi_value_info_get_value"}

	gi_vfunc_info_get_flags   = cc.DlFunc[func(*VFuncInfo) int32, int32]{Name: "gi_vfunc_info_get_flags"}
	gi_vfunc_info_get_offset  = cc.DlFunc[func(*VFuncInfo) uint, uint]{Name: "gi_vfunc_info_get_offset"}
	gi_vfunc_info_get_signal  = cc.DlFunc[func(*VFuncInfo) *SignalInfo, *SignalInfo]{Name: "gi_vfunc_info_get_signal"}
	gi_vfunc_info_get_invoker = cc.DlFunc[func(*VFuncInfo) *FunctionInfo, *FunctionInfo]{Name: "gi_vfunc_info_get_invoker"}
	gi_vfunc_info_get_address = cc.DlFunc[func(*VFuncInfo, gobject.GType, **glib.GError) uptr, uptr]{Name: "gi_vfunc_info_get_address"}

	gi_typelib_new_from_bytes = cc.DlFunc[func(uptr, **glib.GError) *Typelib, *Typelib]{Name: "gi_typelib_new_from_bytes"}
	gi_typelib_ref            = cc.DlFunc[func(*Typelib) *Typelib, *Typelib]{Name: "gi_typelib_ref"}
	gi_typelib_unref          = cc.DlFunc[func(*Typelib), cc.Void]{Name: "gi_typelib_unref"}
	gi_typelib_symbol         = cc.DlFunc[func(*Typelib, cc.String, *uptr) int32, int32]{Name: "gi_typelib_symbol"}
	gi_typelib_get_namespace  = cc.DlFunc[func(*Typelib) cc.String, cc.String]{Name: "gi_typelib_get_namespace"}

	gi_callable_info_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_callable_info_get_type"}
	gi_function_info_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_function_info_get_type"}
	gi_callback_info_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_callback_info_get_type"}
	gi_registered_type_info_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_registered_type_info_get_type"}
	gi_struct_info_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_struct_info_get_type"}
	gi_union_info_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_union_info_get_type"}
	gi_enum_info_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_enum_info_get_type"}
	gi_flags_info_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_flags_info_get_type"}
	gi_object_info_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_object_info_get_type"}
	gi_interface_info_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_interface_info_get_type"}
	gi_constant_info_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_constant_info_get_type"}
	gi_value_info_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_value_info_get_type"}
	gi_signal_info_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_signal_info_get_type"}
	gi_vfunc_info_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_vfunc_info_get_type"}
	gi_property_info_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_property_info_get_type"}
	gi_field_info_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gi_field_info_get_type"}
)
