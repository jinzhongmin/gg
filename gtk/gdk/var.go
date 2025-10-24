package gdk

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cairo"
	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/glib/gio"
	"github.com/jinzhongmin/gg/glib/glib"
	"github.com/jinzhongmin/gg/glib/gobject"
	"github.com/jinzhongmin/gg/pango"
)

type uptr = unsafe.Pointer

func vcbu(fn interface{}) uptr { return cc.Cbk(fn) }
func init()                    { cc.Open("libgtk-4*") }
func carry[T any](arry []T) *T {
	if len(arry) == 0 {
		return nil
	}
	return &arry[0]
}
func anyptr(a interface{}) uptr { return (*(*[2]uptr)(uptr(&a)))[1] }

var (

	// #region AppLaunchContext
	gdk_app_launch_context_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_app_launch_context_get_type"}
	gdk_app_launch_context_new           = cc.DlFunc[func() *AppLaunchContext, *AppLaunchContext]{Name: "gdk_app_launch_context_new"}
	gdk_app_launch_context_get_display   = cc.DlFunc[func(*AppLaunchContext) *Display, *Display]{Name: "gdk_app_launch_context_get_display"}
	gdk_app_launch_context_set_desktop   = cc.DlFunc[func(*AppLaunchContext, int32), cc.Void]{Name: "gdk_app_launch_context_set_desktop"}
	gdk_app_launch_context_set_timestamp = cc.DlFunc[func(*AppLaunchContext, uint32), cc.Void]{Name: "gdk_app_launch_context_set_timestamp"}
	gdk_app_launch_context_set_icon      = cc.DlFunc[func(*AppLaunchContext, *gio.GIcon), cc.Void]{Name: "gdk_app_launch_context_set_icon"}
	gdk_app_launch_context_set_icon_name = cc.DlFunc[func(*AppLaunchContext, cc.String), cc.Void]{Name: "gdk_app_launch_context_set_icon_name"}
	// #endregion

	// #region CICPParams
	gdk_cicp_params_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_cicp_params_get_type"}
	gdk_cicp_params_new                     = cc.DlFunc[func() *CicpParams, *CicpParams]{Name: "gdk_cicp_params_new"}
	gdk_cicp_params_get_color_primaries     = cc.DlFunc[func(*CicpParams) uint32, uint32]{Name: "gdk_cicp_params_get_color_primaries"}
	gdk_cicp_params_set_color_primaries     = cc.DlFunc[func(*CicpParams, uint32), cc.Void]{Name: "gdk_cicp_params_set_color_primaries"}
	gdk_cicp_params_get_transfer_function   = cc.DlFunc[func(*CicpParams) uint32, uint32]{Name: "gdk_cicp_params_get_transfer_function"}
	gdk_cicp_params_set_transfer_function   = cc.DlFunc[func(*CicpParams, uint32), cc.Void]{Name: "gdk_cicp_params_set_transfer_function"}
	gdk_cicp_params_get_matrix_coefficients = cc.DlFunc[func(*CicpParams) uint32, uint32]{Name: "gdk_cicp_params_get_matrix_coefficients"}
	gdk_cicp_params_set_matrix_coefficients = cc.DlFunc[func(*CicpParams, uint32), cc.Void]{Name: "gdk_cicp_params_set_matrix_coefficients"}
	gdk_cicp_params_get_range               = cc.DlFunc[func(*CicpParams) CicpRange, CicpRange]{Name: "gdk_cicp_params_get_range"}
	gdk_cicp_params_set_range               = cc.DlFunc[func(*CicpParams, CicpRange), cc.Void]{Name: "gdk_cicp_params_set_range"}
	gdk_cicp_params_build_color_state       = cc.DlFunc[func(*CicpParams, **error) *ColorState, *ColorState]{Name: "gdk_cicp_params_build_color_state"}
	// #endregion

	// #region Clipboard
	gdk_clipboard_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_clipboard_get_type"}
	gdk_clipboard_get_display         = cc.DlFunc[func(*Clipboard) *Display, *Display]{Name: "gdk_clipboard_get_display"}
	gdk_clipboard_get_formats         = cc.DlFunc[func(*Clipboard) *ContentFormats, *ContentFormats]{Name: "gdk_clipboard_get_formats"}
	gdk_clipboard_is_local            = cc.DlFunc[func(*Clipboard) int32, int32]{Name: "gdk_clipboard_is_local"}
	gdk_clipboard_get_content         = cc.DlFunc[func(*Clipboard) *ContentProvider, *ContentProvider]{Name: "gdk_clipboard_get_content"}
	gdk_clipboard_store_async         = cc.DlFunc[func(*Clipboard, int32, *gio.GCancellable, uptr, uptr), cc.Void]{Name: "gdk_clipboard_store_async"}
	gdk_clipboard_store_finish        = cc.DlFunc[func(*Clipboard, *gio.GAsyncResult, **glib.GError) int32, int32]{Name: "gdk_clipboard_store_finish"}
	gdk_clipboard_read_async          = cc.DlFunc[func(*Clipboard, cc.Strings, int32, *gio.GCancellable, uptr, uptr), cc.Void]{Name: "gdk_clipboard_read_async"}
	gdk_clipboard_read_finish         = cc.DlFunc[func(*Clipboard, *gio.GAsyncResult, *cc.String, **glib.GError) *gio.GInputStream, *gio.GInputStream]{Name: "gdk_clipboard_read_finish"}
	gdk_clipboard_read_value_async    = cc.DlFunc[func(*Clipboard, gobject.GType, int32, *gio.GCancellable, uptr, uptr), cc.Void]{Name: "gdk_clipboard_read_value_async"}
	gdk_clipboard_read_value_finish   = cc.DlFunc[func(*Clipboard, *gio.GAsyncResult, **glib.GError) *gobject.GValue, *gobject.GValue]{Name: "gdk_clipboard_read_value_finish"}
	gdk_clipboard_read_texture_async  = cc.DlFunc[func(*Clipboard, *gio.GCancellable, uptr, uptr), cc.Void]{Name: "gdk_clipboard_read_texture_async"}
	gdk_clipboard_read_texture_finish = cc.DlFunc[func(*Clipboard, *gio.GAsyncResult, **glib.GError) *Texture, *Texture]{Name: "gdk_clipboard_read_texture_finish"}
	gdk_clipboard_read_text_async     = cc.DlFunc[func(*Clipboard, *gio.GCancellable, uptr, uptr), cc.Void]{Name: "gdk_clipboard_read_text_async"}
	gdk_clipboard_read_text_finish    = cc.DlFunc[func(*Clipboard, *gio.GAsyncResult, **glib.GError) cc.String, cc.String]{Name: "gdk_clipboard_read_text_finish"}
	gdk_clipboard_set_content         = cc.DlFunc[func(*Clipboard, *ContentProvider) int32, int32]{Name: "gdk_clipboard_set_content"}
	gdk_clipboard_set                 = cc.DlFunc[func(*Clipboard, gobject.GType, ...interface{}), cc.Void]{Name: "gdk_clipboard_set", Va: true}
	gdk_clipboard_set_value           = cc.DlFunc[func(*Clipboard, *gobject.GValue), cc.Void]{Name: "gdk_clipboard_set_value"}
	gdk_clipboard_set_text            = cc.DlFunc[func(*Clipboard, cc.String), cc.Void]{Name: "gdk_clipboard_set_text"}
	gdk_clipboard_set_texture         = cc.DlFunc[func(*Clipboard, *Texture), cc.Void]{Name: "gdk_clipboard_set_texture"}
	// #endregion

	// #region ColorState
	gdk_color_state_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_color_state_get_type"}
	gdk_color_state_ref                = cc.DlFunc[func(*ColorState) *ColorState, *ColorState]{Name: "gdk_color_state_ref"}
	gdk_color_state_unref              = cc.DlFunc[func(*ColorState), cc.Void]{Name: "gdk_color_state_unref"}
	gdk_color_state_get_srgb           = cc.DlFunc[func() *ColorState, *ColorState]{Name: "gdk_color_state_get_srgb"}
	gdk_color_state_get_srgb_linear    = cc.DlFunc[func() *ColorState, *ColorState]{Name: "gdk_color_state_get_srgb_linear"}
	gdk_color_state_get_rec2100_pq     = cc.DlFunc[func() *ColorState, *ColorState]{Name: "gdk_color_state_get_rec2100_pq"}
	gdk_color_state_get_rec2100_linear = cc.DlFunc[func() *ColorState, *ColorState]{Name: "gdk_color_state_get_rec2100_linear"}
	gdk_color_state_get_oklab          = cc.DlFunc[func() *ColorState, *ColorState]{Name: "gdk_color_state_get_oklab"}
	gdk_color_state_get_oklch          = cc.DlFunc[func() *ColorState, *ColorState]{Name: "gdk_color_state_get_oklch"}
	gdk_color_state_equal              = cc.DlFunc[func(*ColorState, *ColorState) int32, int32]{Name: "gdk_color_state_equal"}
	gdk_color_state_create_cicp_params = cc.DlFunc[func(*ColorState) *CicpParams, *CicpParams]{Name: "gdk_color_state_create_cicp_params"}
	// #endregion

	// #region ContentDeserializer
	gdk_content_deserializer_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_content_deserializer_get_type"}
	gdk_content_deserializer_get_mime_type           = cc.DlFunc[func(*ContentDeserializer) cc.String, cc.String]{Name: "gdk_content_deserializer_get_mime_type"}
	gdk_content_deserializer_get_gtype               = cc.DlFunc[func(*ContentDeserializer) gobject.GType, gobject.GType]{Name: "gdk_content_deserializer_get_gtype"}
	gdk_content_deserializer_get_value               = cc.DlFunc[func(*ContentDeserializer) *gobject.GValue, *gobject.GValue]{Name: "gdk_content_deserializer_get_value"}
	gdk_content_deserializer_get_input_stream        = cc.DlFunc[func(*ContentDeserializer) *gio.GInputStream, *gio.GInputStream]{Name: "gdk_content_deserializer_get_input_stream"}
	gdk_content_deserializer_get_priority            = cc.DlFunc[func(*ContentDeserializer) int32, int32]{Name: "gdk_content_deserializer_get_priority"}
	gdk_content_deserializer_get_cancellable         = cc.DlFunc[func(*ContentDeserializer) *gio.GCancellable, *gio.GCancellable]{Name: "gdk_content_deserializer_get_cancellable"}
	gdk_content_deserializer_get_user_data           = cc.DlFunc[func(*ContentDeserializer) uptr, uptr]{Name: "gdk_content_deserializer_get_user_data"}
	gdk_content_deserializer_set_task_data           = cc.DlFunc[func(*ContentDeserializer, uptr, uptr), cc.Void]{Name: "gdk_content_deserializer_set_task_data"}
	gdk_content_deserializer_get_task_data           = cc.DlFunc[func(*ContentDeserializer) uptr, uptr]{Name: "gdk_content_deserializer_get_task_data"}
	gdk_content_deserializer_return_success          = cc.DlFunc[func(*ContentDeserializer), cc.Void]{Name: "gdk_content_deserializer_return_success"}
	gdk_content_deserializer_return_error            = cc.DlFunc[func(*ContentDeserializer, *glib.GError), cc.Void]{Name: "gdk_content_deserializer_return_error"}
	gdk_content_formats_union_deserialize_gtypes     = cc.DlFunc[func(*ContentFormats) *ContentFormats, *ContentFormats]{Name: "gdk_content_formats_union_deserialize_gtypes"}
	gdk_content_formats_union_deserialize_mime_types = cc.DlFunc[func(*ContentFormats) *ContentFormats, *ContentFormats]{Name: "gdk_content_formats_union_deserialize_mime_types"}
	gdk_content_register_deserializer                = cc.DlFunc[func(cc.String, gobject.GType, uptr, uptr, uptr), cc.Void]{Name: "gdk_content_register_deserializer"}
	gdk_content_deserialize_async                    = cc.DlFunc[func(*gio.GInputStream, cc.String, gobject.GType, int32, *gio.GCancellable, uptr, uptr), cc.Void]{Name: "gdk_content_deserialize_async"}
	gdk_content_deserialize_finish                   = cc.DlFunc[func(*gio.GAsyncResult, *gobject.GValue, **glib.GError) int32, int32]{Name: "gdk_content_deserialize_finish"}
	// #endregion

	// #region ContentFormats
	gdk_intern_mime_type                        = cc.DlFunc[func(cc.String) cc.String, cc.String]{Name: "gdk_intern_mime_type"}
	gdk_content_formats_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_content_formats_get_type"}
	gdk_content_formats_new                     = cc.DlFunc[func(cc.Strings, uint32) *ContentFormats, *ContentFormats]{Name: "gdk_content_formats_new"}
	gdk_content_formats_new_for_gtype           = cc.DlFunc[func(gobject.GType) *ContentFormats, *ContentFormats]{Name: "gdk_content_formats_new_for_gtype"}
	gdk_content_formats_parse                   = cc.DlFunc[func(cc.String) *ContentFormats, *ContentFormats]{Name: "gdk_content_formats_parse"}
	gdk_content_formats_ref                     = cc.DlFunc[func(*ContentFormats) *ContentFormats, *ContentFormats]{Name: "gdk_content_formats_ref"}
	gdk_content_formats_unref                   = cc.DlFunc[func(*ContentFormats), cc.Void]{Name: "gdk_content_formats_unref"}
	gdk_content_formats_print                   = cc.DlFunc[func(*ContentFormats, *glib.GString), cc.Void]{Name: "gdk_content_formats_print"}
	gdk_content_formats_to_string               = cc.DlFunc[func(*ContentFormats) cc.String, cc.String]{Name: "gdk_content_formats_to_string"}
	gdk_content_formats_get_gtypes              = cc.DlFunc[func(*ContentFormats, *uint64) *gobject.GType, *gobject.GType]{Name: "gdk_content_formats_get_gtypes"}
	gdk_content_formats_get_mime_types          = cc.DlFunc[func(*ContentFormats, *uint64) cc.Strings, cc.Strings]{Name: "gdk_content_formats_get_mime_types"}
	gdk_content_formats_union                   = cc.DlFunc[func(*ContentFormats, *ContentFormats) *ContentFormats, *ContentFormats]{Name: "gdk_content_formats_union"}
	gdk_content_formats_match                   = cc.DlFunc[func(*ContentFormats, *ContentFormats) int32, int32]{Name: "gdk_content_formats_match"}
	gdk_content_formats_match_gtype             = cc.DlFunc[func(*ContentFormats, *ContentFormats) gobject.GType, gobject.GType]{Name: "gdk_content_formats_match_gtype"}
	gdk_content_formats_match_mime_type         = cc.DlFunc[func(*ContentFormats, *ContentFormats) cc.String, cc.String]{Name: "gdk_content_formats_match_mime_type"}
	gdk_content_formats_contain_gtype           = cc.DlFunc[func(*ContentFormats, gobject.GType) int32, int32]{Name: "gdk_content_formats_contain_gtype"}
	gdk_content_formats_contain_mime_type       = cc.DlFunc[func(*ContentFormats, cc.String) int32, int32]{Name: "gdk_content_formats_contain_mime_type"}
	gdk_content_formats_is_empty                = cc.DlFunc[func(*ContentFormats) int32, int32]{Name: "gdk_content_formats_is_empty"}
	gdk_content_formats_builder_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_content_formats_builder_get_type"}
	gdk_content_formats_builder_new             = cc.DlFunc[func() *ContentFormatsBuilder, *ContentFormatsBuilder]{Name: "gdk_content_formats_builder_new"}
	gdk_content_formats_builder_ref             = cc.DlFunc[func(*ContentFormatsBuilder) *ContentFormatsBuilder, *ContentFormatsBuilder]{Name: "gdk_content_formats_builder_ref"}
	gdk_content_formats_builder_unref           = cc.DlFunc[func(*ContentFormatsBuilder), cc.Void]{Name: "gdk_content_formats_builder_unref"}
	gdk_content_formats_builder_free_to_formats = cc.DlFunc[func(*ContentFormatsBuilder) *ContentFormats, *ContentFormats]{Name: "gdk_content_formats_builder_free_to_formats"}
	gdk_content_formats_builder_to_formats      = cc.DlFunc[func(*ContentFormatsBuilder) *ContentFormats, *ContentFormats]{Name: "gdk_content_formats_builder_to_formats"}
	gdk_content_formats_builder_add_formats     = cc.DlFunc[func(*ContentFormatsBuilder, *ContentFormats), cc.Void]{Name: "gdk_content_formats_builder_add_formats"}
	gdk_content_formats_builder_add_mime_type   = cc.DlFunc[func(*ContentFormatsBuilder, cc.String), cc.Void]{Name: "gdk_content_formats_builder_add_mime_type"}
	gdk_content_formats_builder_add_gtype       = cc.DlFunc[func(*ContentFormatsBuilder, gobject.GType), cc.Void]{Name: "gdk_content_formats_builder_add_gtype"}
	gdk_file_list_get_type                      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_file_list_get_type"}
	gdk_file_list_get_files                     = cc.DlFunc[func(*FileList) uptr, uptr]{Name: "gdk_file_list_get_files"}
	gdk_file_list_new_from_list                 = cc.DlFunc[func(uptr) *FileList, *FileList]{Name: "gdk_file_list_new_from_list"}
	gdk_file_list_new_from_array                = cc.DlFunc[func(**gio.GFile, uint64) *FileList, *FileList]{Name: "gdk_file_list_new_from_array"}
	// #endregion

	// #region ContentProvider
	gdk_content_provider_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_content_provider_get_type"}
	gdk_content_provider_ref_formats            = cc.DlFunc[func(*ContentProvider) *ContentFormats, *ContentFormats]{Name: "gdk_content_provider_ref_formats"}
	gdk_content_provider_ref_storable_formats   = cc.DlFunc[func(*ContentProvider) *ContentFormats, *ContentFormats]{Name: "gdk_content_provider_ref_storable_formats"}
	gdk_content_provider_content_changed        = cc.DlFunc[func(*ContentProvider), cc.Void]{Name: "gdk_content_provider_content_changed"}
	gdk_content_provider_write_mime_type_async  = cc.DlFunc[func(*ContentProvider, cc.String, *gio.GOutputStream, int32, *gio.GCancellable, uptr, uptr), cc.Void]{Name: "gdk_content_provider_write_mime_type_async"}
	gdk_content_provider_write_mime_type_finish = cc.DlFunc[func(*ContentProvider, *gio.GAsyncResult, **glib.GError) int32, int32]{Name: "gdk_content_provider_write_mime_type_finish"}
	gdk_content_provider_get_value              = cc.DlFunc[func(*ContentProvider, *gobject.GValue, **glib.GError) int32, int32]{Name: "gdk_content_provider_get_value"}
	// #endregion

	// #region ContentProvider
	gdk_content_provider_new_for_value = cc.DlFunc[func(*gobject.GValue) *ContentProvider, *ContentProvider]{Name: "gdk_content_provider_new_for_value"}
	gdk_content_provider_new_typed     = cc.DlFunc[func(gobject.GType, ...interface{}) *ContentProvider, *ContentProvider]{Name: "gdk_content_provider_new_typed", Va: true}
	gdk_content_provider_new_union     = cc.DlFunc[func(**ContentProvider, uint64) *ContentProvider, *ContentProvider]{Name: "gdk_content_provider_new_union"}
	gdk_content_provider_new_for_bytes = cc.DlFunc[func(cc.String, *glib.GBytes) *ContentProvider, *ContentProvider]{Name: "gdk_content_provider_new_for_bytes"}
	// #endregion

	// #region ContentSerializer
	gdk_content_serializer_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_content_serializer_get_type"}
	gdk_content_serializer_get_mime_type           = cc.DlFunc[func(*ContentSerializer) cc.String, cc.String]{Name: "gdk_content_serializer_get_mime_type"}
	gdk_content_serializer_get_gtype               = cc.DlFunc[func(*ContentSerializer) gobject.GType, gobject.GType]{Name: "gdk_content_serializer_get_gtype"}
	gdk_content_serializer_get_value               = cc.DlFunc[func(*ContentSerializer) *gobject.GValue, *gobject.GValue]{Name: "gdk_content_serializer_get_value"}
	gdk_content_serializer_get_output_stream       = cc.DlFunc[func(*ContentSerializer) *gio.GOutputStream, *gio.GOutputStream]{Name: "gdk_content_serializer_get_output_stream"}
	gdk_content_serializer_get_priority            = cc.DlFunc[func(*ContentSerializer) int32, int32]{Name: "gdk_content_serializer_get_priority"}
	gdk_content_serializer_get_cancellable         = cc.DlFunc[func(*ContentSerializer) *gio.GCancellable, *gio.GCancellable]{Name: "gdk_content_serializer_get_cancellable"}
	gdk_content_serializer_get_user_data           = cc.DlFunc[func(*ContentSerializer) uptr, uptr]{Name: "gdk_content_serializer_get_user_data"}
	gdk_content_serializer_set_task_data           = cc.DlFunc[func(*ContentSerializer, uptr, uptr), cc.Void]{Name: "gdk_content_serializer_set_task_data"}
	gdk_content_serializer_get_task_data           = cc.DlFunc[func(*ContentSerializer) uptr, uptr]{Name: "gdk_content_serializer_get_task_data"}
	gdk_content_serializer_return_success          = cc.DlFunc[func(*ContentSerializer), cc.Void]{Name: "gdk_content_serializer_return_success"}
	gdk_content_serializer_return_error            = cc.DlFunc[func(*ContentSerializer, *glib.GError), cc.Void]{Name: "gdk_content_serializer_return_error"}
	gdk_content_formats_union_serialize_gtypes     = cc.DlFunc[func(*ContentFormats) *ContentFormats, *ContentFormats]{Name: "gdk_content_formats_union_serialize_gtypes"}
	gdk_content_formats_union_serialize_mime_types = cc.DlFunc[func(*ContentFormats) *ContentFormats, *ContentFormats]{Name: "gdk_content_formats_union_serialize_mime_types"}
	gdk_content_register_serializer                = cc.DlFunc[func(gobject.GType, cc.String, uptr, uptr, uptr), cc.Void]{Name: "gdk_content_register_serializer"}
	gdk_content_serialize_async                    = cc.DlFunc[func(*gio.GOutputStream, cc.String, *gobject.GValue, int32, *gio.GCancellable, uptr, uptr), cc.Void]{Name: "gdk_content_serialize_async"}
	gdk_content_serialize_finish                   = cc.DlFunc[func(*gio.GAsyncResult, **glib.GError) int32, int32]{Name: "gdk_content_serialize_finish"}
	// #endregion

	// #region Cursor
	gdk_cursor_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_cursor_get_type"}
	gdk_cursor_new_from_texture  = cc.DlFunc[func(texture *Texture, hotspot_x, hotspot_y int32, fallback *Cursor) *Cursor, *Cursor]{Name: "gdk_cursor_new_from_texture"}
	gdk_cursor_new_from_name     = cc.DlFunc[func(name cc.String, fallback *Cursor) *Cursor, *Cursor]{Name: "gdk_cursor_new_from_name"}
	gdk_cursor_new_from_callback = cc.DlFunc[func(callback uptr, data uptr, destroy uptr, fallback *Cursor) *Cursor, *Cursor]{Name: "gdk_cursor_new_from_callback"}
	gdk_cursor_get_fallback      = cc.DlFunc[func(cursor *Cursor) *Cursor, *Cursor]{Name: "gdk_cursor_get_fallback"}
	gdk_cursor_get_name          = cc.DlFunc[func(cursor *Cursor) cc.String, cc.String]{Name: "gdk_cursor_get_name"}
	gdk_cursor_get_texture       = cc.DlFunc[func(cursor *Cursor) *Texture, *Texture]{Name: "gdk_cursor_get_texture"}
	gdk_cursor_get_hotspot_x     = cc.DlFunc[func(cursor *Cursor) int32, int32]{Name: "gdk_cursor_get_hotspot_x"}
	gdk_cursor_get_hotspot_y     = cc.DlFunc[func(cursor *Cursor) int32, int32]{Name: "gdk_cursor_get_hotspot_y"}
	// #endregion

	// #region Device
	gdk_device_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_device_get_type"}
	gdk_device_get_name                = cc.DlFunc[func(*Device) cc.String, cc.String]{Name: "gdk_device_get_name"}
	gdk_device_get_vendor_id           = cc.DlFunc[func(*Device) cc.String, cc.String]{Name: "gdk_device_get_vendor_id"}
	gdk_device_get_product_id          = cc.DlFunc[func(*Device) cc.String, cc.String]{Name: "gdk_device_get_product_id"}
	gdk_device_get_display             = cc.DlFunc[func(*Device) *Display, *Display]{Name: "gdk_device_get_display"}
	gdk_device_get_seat                = cc.DlFunc[func(*Device) *Seat, *Seat]{Name: "gdk_device_get_seat"}
	gdk_device_get_device_tool         = cc.DlFunc[func(*Device) *DeviceTool, *DeviceTool]{Name: "gdk_device_get_device_tool"}
	gdk_device_get_source              = cc.DlFunc[func(*Device) InputSource, InputSource]{Name: "gdk_device_get_source"}
	gdk_device_get_has_cursor          = cc.DlFunc[func(*Device) int32, int32]{Name: "gdk_device_get_has_cursor"}
	gdk_device_get_num_touches         = cc.DlFunc[func(*Device) uint32, uint32]{Name: "gdk_device_get_num_touches"}
	gdk_device_get_modifier_state      = cc.DlFunc[func(*Device) ModifierType, ModifierType]{Name: "gdk_device_get_modifier_state"}
	gdk_device_get_direction           = cc.DlFunc[func(*Device) pango.Direction, pango.Direction]{Name: "gdk_device_get_direction"}
	gdk_device_has_bidi_layouts        = cc.DlFunc[func(*Device) int32, int32]{Name: "gdk_device_has_bidi_layouts"}
	gdk_device_get_caps_lock_state     = cc.DlFunc[func(*Device) int32, int32]{Name: "gdk_device_get_caps_lock_state"}
	gdk_device_get_num_lock_state      = cc.DlFunc[func(*Device) int32, int32]{Name: "gdk_device_get_num_lock_state"}
	gdk_device_get_scroll_lock_state   = cc.DlFunc[func(*Device) int32, int32]{Name: "gdk_device_get_scroll_lock_state"}
	gdk_device_get_surface_at_position = cc.DlFunc[func(*Device, *float64, *float64) *Surface, *Surface]{Name: "gdk_device_get_surface_at_position"}
	gdk_device_get_timestamp           = cc.DlFunc[func(*Device) uint32, uint32]{Name: "gdk_device_get_timestamp"}
	gdk_device_get_active_layout_index = cc.DlFunc[func(*Device) int32, int32]{Name: "gdk_device_get_active_layout_index"}
	gdk_device_get_layout_names        = cc.DlFunc[func(*Device) cc.Strings, cc.Strings]{Name: "gdk_device_get_layout_names"}
	// #endregion

	// #region DevicePad
	gdk_device_pad_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_device_pad_get_type"}
	gdk_device_pad_get_n_groups      = cc.DlFunc[func(*DevicePad) int32, int32]{Name: "gdk_device_pad_get_n_groups"}
	gdk_device_pad_get_group_n_modes = cc.DlFunc[func(*DevicePad, int32) int32, int32]{Name: "gdk_device_pad_get_group_n_modes"}
	gdk_device_pad_get_n_features    = cc.DlFunc[func(*DevicePad, DevicePadFeature) int32, int32]{Name: "gdk_device_pad_get_n_features"}
	gdk_device_pad_get_feature_group = cc.DlFunc[func(*DevicePad, DevicePadFeature, int32) int32, int32]{Name: "gdk_device_pad_get_feature_group"}
	// #endregion

	// #region DeviceTool
	gdk_device_tool_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_device_tool_get_type"}
	gdk_device_tool_get_serial      = cc.DlFunc[func(*DeviceTool) uint64, uint64]{Name: "gdk_device_tool_get_serial"}
	gdk_device_tool_get_hardware_id = cc.DlFunc[func(*DeviceTool) uint64, uint64]{Name: "gdk_device_tool_get_hardware_id"}
	gdk_device_tool_get_tool_type   = cc.DlFunc[func(*DeviceTool) DeviceToolType, DeviceToolType]{Name: "gdk_device_tool_get_tool_type"}
	gdk_device_tool_get_axes        = cc.DlFunc[func(*DeviceTool) AxisFlags, AxisFlags]{Name: "gdk_device_tool_get_axes"}
	// #endregion

	// #region Display
	gdk_display_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_display_get_type"}
	gdk_display_open                   = cc.DlFunc[func(cc.String) *Display, *Display]{Name: "gdk_display_open"}
	gdk_display_get_name               = cc.DlFunc[func(*Display) cc.String, cc.String]{Name: "gdk_display_get_name"}
	gdk_display_device_is_grabbed      = cc.DlFunc[func(*Display, *Device) int32, int32]{Name: "gdk_display_device_is_grabbed"}
	gdk_display_beep                   = cc.DlFunc[func(*Display), cc.Void]{Name: "gdk_display_beep"}
	gdk_display_sync                   = cc.DlFunc[func(*Display), cc.Void]{Name: "gdk_display_sync"}
	gdk_display_flush                  = cc.DlFunc[func(*Display), cc.Void]{Name: "gdk_display_flush"}
	gdk_display_close                  = cc.DlFunc[func(*Display), cc.Void]{Name: "gdk_display_close"}
	gdk_display_is_closed              = cc.DlFunc[func(*Display) int32, int32]{Name: "gdk_display_is_closed"}
	gdk_display_is_composited          = cc.DlFunc[func(*Display) int32, int32]{Name: "gdk_display_is_composited"}
	gdk_display_is_rgba                = cc.DlFunc[func(*Display) int32, int32]{Name: "gdk_display_is_rgba"}
	gdk_display_supports_shadow_width  = cc.DlFunc[func(*Display) int32, int32]{Name: "gdk_display_supports_shadow_width"}
	gdk_display_supports_input_shapes  = cc.DlFunc[func(*Display) int32, int32]{Name: "gdk_display_supports_input_shapes"}
	gdk_display_prepare_gl             = cc.DlFunc[func(*Display, **glib.GError) int32, int32]{Name: "gdk_display_prepare_gl"}
	gdk_display_create_gl_context      = cc.DlFunc[func(*Display, **glib.GError) *GLContext, *GLContext]{Name: "gdk_display_create_gl_context"}
	gdk_display_get_default            = cc.DlFunc[func() *Display, *Display]{Name: "gdk_display_get_default"}
	gdk_display_get_clipboard          = cc.DlFunc[func(*Display) *Clipboard, *Clipboard]{Name: "gdk_display_get_clipboard"}
	gdk_display_get_primary_clipboard  = cc.DlFunc[func(*Display) *Clipboard, *Clipboard]{Name: "gdk_display_get_primary_clipboard"}
	gdk_display_get_app_launch_context = cc.DlFunc[func(*Display) *AppLaunchContext, *AppLaunchContext]{Name: "gdk_display_get_app_launch_context"}
	gdk_display_get_default_seat       = cc.DlFunc[func(*Display) *Seat, *Seat]{Name: "gdk_display_get_default_seat"}
	gdk_display_list_seats             = cc.DlFunc[func(*Display) *glib.GList[Seat], *glib.GList[Seat]]{Name: "gdk_display_list_seats"}
	gdk_display_get_monitors           = cc.DlFunc[func(*Display) *gio.GListModel, *gio.GListModel]{Name: "gdk_display_get_monitors"}
	gdk_display_get_monitor_at_surface = cc.DlFunc[func(*Display, *Surface) *Monitor, *Monitor]{Name: "gdk_display_get_monitor_at_surface"}
	gdk_display_map_keyval             = cc.DlFunc[func(*Display, uint32, **KeymapKey, *int32) int32, int32]{Name: "gdk_display_map_keyval"}
	gdk_display_map_keycode            = cc.DlFunc[func(*Display, uint32, **KeymapKey, **uint32, *int32) int32, int32]{Name: "gdk_display_map_keycode"}
	gdk_display_translate_key          = cc.DlFunc[func(*Display, uint32, ModifierType, int32, *uint32, *int32, *int32, *ModifierType) int32, int32]{Name: "gdk_display_translate_key"}
	gdk_display_get_setting            = cc.DlFunc[func(*Display, cc.String, *gobject.GValue) int32, int32]{Name: "gdk_display_get_setting"}
	gdk_display_get_dmabuf_formats     = cc.DlFunc[func(*Display) *DmabufFormats, *DmabufFormats]{Name: "gdk_display_get_dmabuf_formats"}
	// #endregion

	// #region DisplayManager
	gdk_display_manager_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_display_manager_get_type"}
	gdk_display_manager_get                 = cc.DlFunc[func() *DisplayManager, *DisplayManager]{Name: "gdk_display_manager_get"}
	gdk_display_manager_get_default_display = cc.DlFunc[func(*DisplayManager) *Display, *Display]{Name: "gdk_display_manager_get_default_display"}
	gdk_display_manager_set_default_display = cc.DlFunc[func(*DisplayManager, *Display), cc.Void]{Name: "gdk_display_manager_set_default_display"}
	gdk_display_manager_list_displays       = cc.DlFunc[func(*DisplayManager) uptr, uptr]{Name: "gdk_display_manager_list_displays"}
	gdk_display_manager_open_display        = cc.DlFunc[func(*DisplayManager, cc.String) *Display, *Display]{Name: "gdk_display_manager_open_display"}
	gdk_set_allowed_backends                = cc.DlFunc[func(cc.String), cc.Void]{Name: "gdk_set_allowed_backends"}
	// #endregion

	// #region DmabufFormats
	gdk_dmabuf_formats_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_dmabuf_formats_get_type"}
	gdk_dmabuf_formats_ref           = cc.DlFunc[func(*DmabufFormats) *DmabufFormats, *DmabufFormats]{Name: "gdk_dmabuf_formats_ref"}
	gdk_dmabuf_formats_unref         = cc.DlFunc[func(*DmabufFormats), cc.Void]{Name: "gdk_dmabuf_formats_unref"}
	gdk_dmabuf_formats_get_n_formats = cc.DlFunc[func(*DmabufFormats) uint64, uint64]{Name: "gdk_dmabuf_formats_get_n_formats"}
	gdk_dmabuf_formats_get_format    = cc.DlFunc[func(*DmabufFormats, uint64, *uint32, *uint64), cc.Void]{Name: "gdk_dmabuf_formats_get_format"}
	gdk_dmabuf_formats_contains      = cc.DlFunc[func(*DmabufFormats, uint32, uint64) int32, int32]{Name: "gdk_dmabuf_formats_contains"}
	gdk_dmabuf_formats_equal         = cc.DlFunc[func(*DmabufFormats, *DmabufFormats) int32, int32]{Name: "gdk_dmabuf_formats_equal"}
	// #endregion

	// #region DmabufTexture
	gdk_dmabuf_texture_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_dmabuf_texture_get_type"}
	gdk_dmabuf_error_quark      = cc.DlFunc[func() glib.GQuark, glib.GQuark]{Name: "gdk_dmabuf_error_quark"}
	// #endregion

	// #region DmabufTextureBuilder
	gdk_dmabuf_texture_builder_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_dmabuf_texture_builder_get_type"}
	gdk_dmabuf_texture_builder_new                = cc.DlFunc[func() *DmabufTextureBuilder, *DmabufTextureBuilder]{Name: "gdk_dmabuf_texture_builder_new"}
	gdk_dmabuf_texture_builder_get_display        = cc.DlFunc[func(*DmabufTextureBuilder) *Display, *Display]{Name: "gdk_dmabuf_texture_builder_get_display"}
	gdk_dmabuf_texture_builder_set_display        = cc.DlFunc[func(*DmabufTextureBuilder, *Display), cc.Void]{Name: "gdk_dmabuf_texture_builder_set_display"}
	gdk_dmabuf_texture_builder_get_width          = cc.DlFunc[func(*DmabufTextureBuilder) uint32, uint32]{Name: "gdk_dmabuf_texture_builder_get_width"}
	gdk_dmabuf_texture_builder_set_width          = cc.DlFunc[func(*DmabufTextureBuilder, uint32), cc.Void]{Name: "gdk_dmabuf_texture_builder_set_width"}
	gdk_dmabuf_texture_builder_get_height         = cc.DlFunc[func(*DmabufTextureBuilder) uint32, uint32]{Name: "gdk_dmabuf_texture_builder_get_height"}
	gdk_dmabuf_texture_builder_set_height         = cc.DlFunc[func(*DmabufTextureBuilder, uint32), cc.Void]{Name: "gdk_dmabuf_texture_builder_set_height"}
	gdk_dmabuf_texture_builder_get_fourcc         = cc.DlFunc[func(*DmabufTextureBuilder) uint32, uint32]{Name: "gdk_dmabuf_texture_builder_get_fourcc"}
	gdk_dmabuf_texture_builder_set_fourcc         = cc.DlFunc[func(*DmabufTextureBuilder, uint32), cc.Void]{Name: "gdk_dmabuf_texture_builder_set_fourcc"}
	gdk_dmabuf_texture_builder_get_modifier       = cc.DlFunc[func(*DmabufTextureBuilder) uint64, uint64]{Name: "gdk_dmabuf_texture_builder_get_modifier"}
	gdk_dmabuf_texture_builder_set_modifier       = cc.DlFunc[func(*DmabufTextureBuilder, uint64), cc.Void]{Name: "gdk_dmabuf_texture_builder_set_modifier"}
	gdk_dmabuf_texture_builder_get_premultiplied  = cc.DlFunc[func(*DmabufTextureBuilder) int32, int32]{Name: "gdk_dmabuf_texture_builder_get_premultiplied"}
	gdk_dmabuf_texture_builder_set_premultiplied  = cc.DlFunc[func(*DmabufTextureBuilder, bool), cc.Void]{Name: "gdk_dmabuf_texture_builder_set_premultiplied"}
	gdk_dmabuf_texture_builder_get_n_planes       = cc.DlFunc[func(*DmabufTextureBuilder) uint32, uint32]{Name: "gdk_dmabuf_texture_builder_get_n_planes"}
	gdk_dmabuf_texture_builder_set_n_planes       = cc.DlFunc[func(*DmabufTextureBuilder, uint32), cc.Void]{Name: "gdk_dmabuf_texture_builder_set_n_planes"}
	gdk_dmabuf_texture_builder_get_fd             = cc.DlFunc[func(*DmabufTextureBuilder, uint32) int32, int32]{Name: "gdk_dmabuf_texture_builder_get_fd"}
	gdk_dmabuf_texture_builder_set_fd             = cc.DlFunc[func(*DmabufTextureBuilder, uint32, int32), cc.Void]{Name: "gdk_dmabuf_texture_builder_set_fd"}
	gdk_dmabuf_texture_builder_get_stride         = cc.DlFunc[func(*DmabufTextureBuilder, uint32) uint32, uint32]{Name: "gdk_dmabuf_texture_builder_get_stride"}
	gdk_dmabuf_texture_builder_set_stride         = cc.DlFunc[func(*DmabufTextureBuilder, uint32, uint32), cc.Void]{Name: "gdk_dmabuf_texture_builder_set_stride"}
	gdk_dmabuf_texture_builder_get_offset         = cc.DlFunc[func(*DmabufTextureBuilder, uint32) uint32, uint32]{Name: "gdk_dmabuf_texture_builder_get_offset"}
	gdk_dmabuf_texture_builder_set_offset         = cc.DlFunc[func(*DmabufTextureBuilder, uint32, uint32), cc.Void]{Name: "gdk_dmabuf_texture_builder_set_offset"}
	gdk_dmabuf_texture_builder_get_color_state    = cc.DlFunc[func(*DmabufTextureBuilder) *ColorState, *ColorState]{Name: "gdk_dmabuf_texture_builder_get_color_state"}
	gdk_dmabuf_texture_builder_set_color_state    = cc.DlFunc[func(*DmabufTextureBuilder, *ColorState), cc.Void]{Name: "gdk_dmabuf_texture_builder_set_color_state"}
	gdk_dmabuf_texture_builder_get_update_texture = cc.DlFunc[func(*DmabufTextureBuilder) *Texture, *Texture]{Name: "gdk_dmabuf_texture_builder_get_update_texture"}
	gdk_dmabuf_texture_builder_set_update_texture = cc.DlFunc[func(*DmabufTextureBuilder, *Texture), cc.Void]{Name: "gdk_dmabuf_texture_builder_set_update_texture"}
	gdk_dmabuf_texture_builder_get_update_region  = cc.DlFunc[func(*DmabufTextureBuilder) *cairo.Region, *cairo.Region]{Name: "gdk_dmabuf_texture_builder_get_update_region"}
	gdk_dmabuf_texture_builder_set_update_region  = cc.DlFunc[func(*DmabufTextureBuilder, *cairo.Region), cc.Void]{Name: "gdk_dmabuf_texture_builder_set_update_region"}
	gdk_dmabuf_texture_builder_build              = cc.DlFunc[func(*DmabufTextureBuilder, uptr, uptr, **glib.GError) *Texture, *Texture]{Name: "gdk_dmabuf_texture_builder_build"}
	// #endregion

	// #region Drag
	gdk_drag_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_drag_get_type"}
	gdk_drag_get_display         = cc.DlFunc[func(*Drag) *Display, *Display]{Name: "gdk_drag_get_display"}
	gdk_drag_get_device          = cc.DlFunc[func(*Drag) *Device, *Device]{Name: "gdk_drag_get_device"}
	gdk_drag_get_formats         = cc.DlFunc[func(*Drag) *ContentFormats, *ContentFormats]{Name: "gdk_drag_get_formats"}
	gdk_drag_get_actions         = cc.DlFunc[func(*Drag) DragAction, DragAction]{Name: "gdk_drag_get_actions"}
	gdk_drag_get_selected_action = cc.DlFunc[func(*Drag) DragAction, DragAction]{Name: "gdk_drag_get_selected_action"}
	gdk_drag_action_is_unique    = cc.DlFunc[func(DragAction) int32, int32]{Name: "gdk_drag_action_is_unique"}
	gdk_drag_begin               = cc.DlFunc[func(*Surface, *Device, *ContentProvider, DragAction, float64, float64) *Drag, *Drag]{Name: "gdk_drag_begin"}
	gdk_drag_drop_done           = cc.DlFunc[func(*Drag, bool), cc.Void]{Name: "gdk_drag_drop_done"}
	gdk_drag_get_drag_surface    = cc.DlFunc[func(*Drag) *Surface, *Surface]{Name: "gdk_drag_get_drag_surface"}
	gdk_drag_set_hotspot         = cc.DlFunc[func(*Drag, int32, int32), cc.Void]{Name: "gdk_drag_set_hotspot"}
	gdk_drag_get_content         = cc.DlFunc[func(*Drag) *ContentProvider, *ContentProvider]{Name: "gdk_drag_get_content"}
	gdk_drag_get_surface         = cc.DlFunc[func(*Drag) *Surface, *Surface]{Name: "gdk_drag_get_surface"}
	// #endregion

	// #region DragSurface
	gdk_drag_surface_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_drag_surface_get_type"}
	gdk_drag_surface_present  = cc.DlFunc[func(*DragSurface, int32, int32) int32, int32]{Name: "gdk_drag_surface_present"}
	// #endregion

	// #region DragSurfaceSize
	gdk_drag_surface_size_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_drag_surface_size_get_type"}
	gdk_drag_surface_size_set_size = cc.DlFunc[func(*DragSurfaceSize, int32, int32), cc.Void]{Name: "gdk_drag_surface_size_set_size"}
	// #endregion

	// #region DrawContext
	gdk_draw_context_get_type    = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_draw_context_get_type"}
	gdk_draw_context_get_display = cc.DlFunc[func(*DrawContext) *Display, *Display]{Name: "gdk_draw_context_get_display"}
	gdk_draw_context_get_surface = cc.DlFunc[func(*DrawContext) *Surface, *Surface]{Name: "gdk_draw_context_get_surface"}
	// #endregion

	// #region Drop
	gdk_drop_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_drop_get_type"}
	gdk_drop_get_display       = cc.DlFunc[func(*Drop) *Display, *Display]{Name: "gdk_drop_get_display"}
	gdk_drop_get_device        = cc.DlFunc[func(*Drop) *Device, *Device]{Name: "gdk_drop_get_device"}
	gdk_drop_get_surface       = cc.DlFunc[func(*Drop) *Surface, *Surface]{Name: "gdk_drop_get_surface"}
	gdk_drop_get_formats       = cc.DlFunc[func(*Drop) *ContentFormats, *ContentFormats]{Name: "gdk_drop_get_formats"}
	gdk_drop_get_actions       = cc.DlFunc[func(*Drop) DragAction, DragAction]{Name: "gdk_drop_get_actions"}
	gdk_drop_get_drag          = cc.DlFunc[func(*Drop) *Drag, *Drag]{Name: "gdk_drop_get_drag"}
	gdk_drop_status            = cc.DlFunc[func(*Drop, DragAction, DragAction), cc.Void]{Name: "gdk_drop_status"}
	gdk_drop_finish            = cc.DlFunc[func(*Drop, DragAction), cc.Void]{Name: "gdk_drop_finish"}
	gdk_drop_read_async        = cc.DlFunc[func(*Drop, cc.Strings, int32, *gio.GCancellable, uptr, uptr), cc.Void]{Name: "gdk_drop_read_async"}
	gdk_drop_read_finish       = cc.DlFunc[func(*Drop, *gio.GAsyncResult, *cc.String, **glib.GError) *gio.GInputStream, *gio.GInputStream]{Name: "gdk_drop_read_finish"}
	gdk_drop_read_value_async  = cc.DlFunc[func(*Drop, gobject.GType, int32, *gio.GCancellable, uptr, uptr), cc.Void]{Name: "gdk_drop_read_value_async"}
	gdk_drop_read_value_finish = cc.DlFunc[func(*Drop, *gio.GAsyncResult, **glib.GError) *gobject.GValue, *gobject.GValue]{Name: "gdk_drop_read_value_finish"}
	// #endregion

	// #region Event
	gdk_event_get_type                       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_event_get_type"}
	gdk_event_sequence_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_event_sequence_get_type"}
	gdk_event_ref                            = cc.DlFunc[func(event *Event) *Event, *Event]{Name: "gdk_event_ref"}
	gdk_event_unref                          = cc.DlFunc[func(event *Event), cc.Void]{Name: "gdk_event_unref"}
	gdk_event_get_event_type                 = cc.DlFunc[func(event *Event) EventType, EventType]{Name: "gdk_event_get_event_type"}
	gdk_event_get_surface                    = cc.DlFunc[func(event *Event) *Surface, *Surface]{Name: "gdk_event_get_surface"}
	gdk_event_get_seat                       = cc.DlFunc[func(event *Event) *Seat, *Seat]{Name: "gdk_event_get_seat"}
	gdk_event_get_device                     = cc.DlFunc[func(event *Event) *Device, *Device]{Name: "gdk_event_get_device"}
	gdk_event_get_device_tool                = cc.DlFunc[func(event *Event) *DeviceTool, *DeviceTool]{Name: "gdk_event_get_device_tool"}
	gdk_event_get_time                       = cc.DlFunc[func(event *Event) uint32, uint32]{Name: "gdk_event_get_time"}
	gdk_event_get_display                    = cc.DlFunc[func(event *Event) *Display, *Display]{Name: "gdk_event_get_display"}
	gdk_event_get_event_sequence             = cc.DlFunc[func(event *Event) *EventSequence, *EventSequence]{Name: "gdk_event_get_event_sequence"}
	gdk_event_get_modifier_state             = cc.DlFunc[func(event *Event) ModifierType, ModifierType]{Name: "gdk_event_get_modifier_state"}
	gdk_event_get_position                   = cc.DlFunc[func(event *Event, x, y *float64) bool, bool]{Name: "gdk_event_get_position"}
	gdk_event_get_axes                       = cc.DlFunc[func(event *Event, axes **float64, n_axes *uint32) bool, bool]{Name: "gdk_event_get_axes"}
	gdk_event_get_axis                       = cc.DlFunc[func(event *Event, axis_use AxisUse, value *float64) bool, bool]{Name: "gdk_event_get_axis"}
	gdk_event_get_history                    = cc.DlFunc[func(event *Event, out_n_coords *uint32) *TimeCoord, *TimeCoord]{Name: "gdk_event_get_history"}
	gdk_event_get_pointer_emulated           = cc.DlFunc[func(event *Event) bool, bool]{Name: "gdk_event_get_pointer_emulated"}
	gdk_button_event_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_button_event_get_type"}
	gdk_button_event_get_button              = cc.DlFunc[func(event *ButtonEvent) uint32, uint32]{Name: "gdk_button_event_get_button"}
	gdk_scroll_event_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_scroll_event_get_type"}
	gdk_scroll_event_get_direction           = cc.DlFunc[func(event *ScrollEvent) ScrollDirection, ScrollDirection]{Name: "gdk_scroll_event_get_direction"}
	gdk_scroll_event_get_deltas              = cc.DlFunc[func(event *ScrollEvent, delta_x, delta_y *float64), cc.Void]{Name: "gdk_scroll_event_get_deltas"}
	gdk_scroll_event_get_unit                = cc.DlFunc[func(event *ScrollEvent) ScrollUnit, ScrollUnit]{Name: "gdk_scroll_event_get_unit"}
	gdk_scroll_event_is_stop                 = cc.DlFunc[func(event *ScrollEvent) bool, bool]{Name: "gdk_scroll_event_is_stop"}
	gdk_key_event_get_type                   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_key_event_get_type"}
	gdk_key_event_get_keyval                 = cc.DlFunc[func(event *KeyEvent) uint32, uint32]{Name: "gdk_key_event_get_keyval"}
	gdk_key_event_get_keycode                = cc.DlFunc[func(event *KeyEvent) uint32, uint32]{Name: "gdk_key_event_get_keycode"}
	gdk_key_event_get_consumed_modifiers     = cc.DlFunc[func(event *KeyEvent) ModifierType, ModifierType]{Name: "gdk_key_event_get_consumed_modifiers"}
	gdk_key_event_get_layout                 = cc.DlFunc[func(event *KeyEvent) uint32, uint32]{Name: "gdk_key_event_get_layout"}
	gdk_key_event_get_level                  = cc.DlFunc[func(event *KeyEvent) uint32, uint32]{Name: "gdk_key_event_get_level"}
	gdk_key_event_is_modifier                = cc.DlFunc[func(event *KeyEvent) bool, bool]{Name: "gdk_key_event_is_modifier"}
	gdk_focus_event_get_type                 = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_focus_event_get_type"}
	gdk_focus_event_get_in                   = cc.DlFunc[func(event *FocusEvent) bool, bool]{Name: "gdk_focus_event_get_in"}
	gdk_touch_event_get_type                 = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_touch_event_get_type"}
	gdk_touch_event_get_emulating_pointer    = cc.DlFunc[func(event *TouchEvent) bool, bool]{Name: "gdk_touch_event_get_emulating_pointer"}
	gdk_crossing_event_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_crossing_event_get_type"}
	gdk_crossing_event_get_mode              = cc.DlFunc[func(event *CrossingEvent) CrossingMode, CrossingMode]{Name: "gdk_crossing_event_get_mode"}
	gdk_crossing_event_get_detail            = cc.DlFunc[func(event *CrossingEvent) NotifyType, NotifyType]{Name: "gdk_crossing_event_get_detail"}
	gdk_crossing_event_get_focus             = cc.DlFunc[func(event *CrossingEvent) bool, bool]{Name: "gdk_crossing_event_get_focus"}
	gdk_touchpad_event_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_touchpad_event_get_type"}
	gdk_touchpad_event_get_gesture_phase     = cc.DlFunc[func(event *TouchpadEvent) TouchpadGesturePhase, TouchpadGesturePhase]{Name: "gdk_touchpad_event_get_gesture_phase"}
	gdk_touchpad_event_get_n_fingers         = cc.DlFunc[func(event *TouchpadEvent) uint32, uint32]{Name: "gdk_touchpad_event_get_n_fingers"}
	gdk_touchpad_event_get_deltas            = cc.DlFunc[func(event *TouchpadEvent, dx, dy *float64), cc.Void]{Name: "gdk_touchpad_event_get_deltas"}
	gdk_touchpad_event_get_pinch_angle_delta = cc.DlFunc[func(event *TouchpadEvent) float64, float64]{Name: "gdk_touchpad_event_get_pinch_angle_delta"}
	gdk_touchpad_event_get_pinch_scale       = cc.DlFunc[func(event *TouchpadEvent) float64, float64]{Name: "gdk_touchpad_event_get_pinch_scale"}
	gdk_pad_event_get_type                   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_pad_event_get_type"}
	gdk_pad_event_get_button                 = cc.DlFunc[func(event *PadEvent) uint32, uint32]{Name: "gdk_pad_event_get_button"}
	gdk_pad_event_get_axis_value             = cc.DlFunc[func(event *PadEvent, index *uint32, value *float64), cc.Void]{Name: "gdk_pad_event_get_axis_value"}
	gdk_pad_event_get_group_mode             = cc.DlFunc[func(event *PadEvent, group, mode *uint32), cc.Void]{Name: "gdk_pad_event_get_group_mode"}
	gdk_dnd_event_get_type                   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_dnd_event_get_type"}
	gdk_dnd_event_get_drop                   = cc.DlFunc[func(event *DNDEvent) *Drop, *Drop]{Name: "gdk_dnd_event_get_drop"}
	gdk_grab_broken_event_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_grab_broken_event_get_type"}
	gdk_grab_broken_event_get_grab_surface   = cc.DlFunc[func(event *GrabBrokenEvent) *Surface, *Surface]{Name: "gdk_grab_broken_event_get_grab_surface"}
	gdk_grab_broken_event_get_implicit       = cc.DlFunc[func(event *GrabBrokenEvent) bool, bool]{Name: "gdk_grab_broken_event_get_implicit"}
	gdk_motion_event_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_motion_event_get_type"}
	gdk_delete_event_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_delete_event_get_type"}
	gdk_proximity_event_get_type             = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_proximity_event_get_type"}
	gdk_event_triggers_context_menu          = cc.DlFunc[func(event *Event) bool, bool]{Name: "gdk_event_triggers_context_menu"}
	gdk_events_get_distance                  = cc.DlFunc[func(event1, event2 *Event, distance *float64) bool, bool]{Name: "gdk_events_get_distance"}
	gdk_events_get_angle                     = cc.DlFunc[func(event1, event2 *Event, angle *float64) bool, bool]{Name: "gdk_events_get_angle"}
	gdk_events_get_center                    = cc.DlFunc[func(event1, event2 *Event, x, y *float64) bool, bool]{Name: "gdk_events_get_center"}
	gdk_key_event_matches                    = cc.DlFunc[func(event *Event, keyval uint32, modifiers ModifierType) KeyMatch, KeyMatch]{Name: "gdk_key_event_matches"}
	gdk_key_event_get_match                  = cc.DlFunc[func(event *Event, keyval *uint32, modifiers *ModifierType) bool, bool]{Name: "gdk_key_event_get_match"}
	// #endregion

	// #region FrameClock
	gdk_frame_clock_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_frame_clock_get_type"}
	gdk_frame_clock_get_frame_time      = cc.DlFunc[func(*FrameClock) int64, int64]{Name: "gdk_frame_clock_get_frame_time"}
	gdk_frame_clock_request_phase       = cc.DlFunc[func(*FrameClock, FrameClockPhase), cc.Void]{Name: "gdk_frame_clock_request_phase"}
	gdk_frame_clock_begin_updating      = cc.DlFunc[func(*FrameClock), cc.Void]{Name: "gdk_frame_clock_begin_updating"}
	gdk_frame_clock_end_updating        = cc.DlFunc[func(*FrameClock), cc.Void]{Name: "gdk_frame_clock_end_updating"}
	gdk_frame_clock_get_frame_counter   = cc.DlFunc[func(*FrameClock) int64, int64]{Name: "gdk_frame_clock_get_frame_counter"}
	gdk_frame_clock_get_history_start   = cc.DlFunc[func(*FrameClock) int64, int64]{Name: "gdk_frame_clock_get_history_start"}
	gdk_frame_clock_get_timings         = cc.DlFunc[func(*FrameClock, int64) *FrameTimings, *FrameTimings]{Name: "gdk_frame_clock_get_timings"}
	gdk_frame_clock_get_current_timings = cc.DlFunc[func(*FrameClock) *FrameTimings, *FrameTimings]{Name: "gdk_frame_clock_get_current_timings"}
	gdk_frame_clock_get_refresh_info    = cc.DlFunc[func(*FrameClock, int64, *int64, *int64), cc.Void]{Name: "gdk_frame_clock_get_refresh_info"}
	gdk_frame_clock_get_fps             = cc.DlFunc[func(*FrameClock) float64, float64]{Name: "gdk_frame_clock_get_fps"}
	// #endregion

	// #region FrameTimings
	gdk_frame_timings_get_type                        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_frame_timings_get_type"}
	gdk_frame_timings_ref                             = cc.DlFunc[func(*FrameTimings) *FrameTimings, *FrameTimings]{Name: "gdk_frame_timings_ref"}
	gdk_frame_timings_unref                           = cc.DlFunc[func(*FrameTimings), cc.Void]{Name: "gdk_frame_timings_unref"}
	gdk_frame_timings_get_frame_counter               = cc.DlFunc[func(*FrameTimings) int64, int64]{Name: "gdk_frame_timings_get_frame_counter"}
	gdk_frame_timings_get_complete                    = cc.DlFunc[func(*FrameTimings) int32, int32]{Name: "gdk_frame_timings_get_complete"}
	gdk_frame_timings_get_frame_time                  = cc.DlFunc[func(*FrameTimings) int64, int64]{Name: "gdk_frame_timings_get_frame_time"}
	gdk_frame_timings_get_presentation_time           = cc.DlFunc[func(*FrameTimings) int64, int64]{Name: "gdk_frame_timings_get_presentation_time"}
	gdk_frame_timings_get_refresh_interval            = cc.DlFunc[func(*FrameTimings) int64, int64]{Name: "gdk_frame_timings_get_refresh_interval"}
	gdk_frame_timings_get_predicted_presentation_time = cc.DlFunc[func(*FrameTimings) int64, int64]{Name: "gdk_frame_timings_get_predicted_presentation_time"}
	// #endregion

	// #region GLContext
	gdk_gl_error_quark                    = cc.DlFunc[func() glib.GQuark, glib.GQuark]{Name: "gdk_gl_error_quark"}
	gdk_gl_context_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_gl_context_get_type"}
	gdk_gl_context_get_display            = cc.DlFunc[func(*GLContext) *Display, *Display]{Name: "gdk_gl_context_get_display"}
	gdk_gl_context_get_surface            = cc.DlFunc[func(*GLContext) *Surface, *Surface]{Name: "gdk_gl_context_get_surface"}
	gdk_gl_context_get_shared_context     = cc.DlFunc[func(*GLContext) *GLContext, *GLContext]{Name: "gdk_gl_context_get_shared_context"}
	gdk_gl_context_get_version            = cc.DlFunc[func(*GLContext, *int, *int), cc.Void]{Name: "gdk_gl_context_get_version"}
	gdk_gl_context_is_legacy              = cc.DlFunc[func(*GLContext) int32, int32]{Name: "gdk_gl_context_is_legacy"}
	gdk_gl_context_is_shared              = cc.DlFunc[func(*GLContext, *GLContext) int32, int32]{Name: "gdk_gl_context_is_shared"}
	gdk_gl_context_set_required_version   = cc.DlFunc[func(*GLContext, int, int), cc.Void]{Name: "gdk_gl_context_set_required_version"}
	gdk_gl_context_get_required_version   = cc.DlFunc[func(*GLContext, *int, *int), cc.Void]{Name: "gdk_gl_context_get_required_version"}
	gdk_gl_context_set_debug_enabled      = cc.DlFunc[func(*GLContext, bool), cc.Void]{Name: "gdk_gl_context_set_debug_enabled"}
	gdk_gl_context_get_debug_enabled      = cc.DlFunc[func(*GLContext) int32, int32]{Name: "gdk_gl_context_get_debug_enabled"}
	gdk_gl_context_set_forward_compatible = cc.DlFunc[func(*GLContext, bool), cc.Void]{Name: "gdk_gl_context_set_forward_compatible"}
	gdk_gl_context_get_forward_compatible = cc.DlFunc[func(*GLContext) int32, int32]{Name: "gdk_gl_context_get_forward_compatible"}
	gdk_gl_context_set_allowed_apis       = cc.DlFunc[func(*GLContext, GLAPI), cc.Void]{Name: "gdk_gl_context_set_allowed_apis"}
	gdk_gl_context_get_allowed_apis       = cc.DlFunc[func(*GLContext) GLAPI, GLAPI]{Name: "gdk_gl_context_get_allowed_apis"}
	gdk_gl_context_get_api                = cc.DlFunc[func(*GLContext) GLAPI, GLAPI]{Name: "gdk_gl_context_get_api"}
	gdk_gl_context_set_use_es             = cc.DlFunc[func(*GLContext, int), cc.Void]{Name: "gdk_gl_context_set_use_es"}
	gdk_gl_context_get_use_es             = cc.DlFunc[func(*GLContext) int32, int32]{Name: "gdk_gl_context_get_use_es"}
	gdk_gl_context_realize                = cc.DlFunc[func(*GLContext, **glib.GError) int32, int32]{Name: "gdk_gl_context_realize"}
	gdk_gl_context_make_current           = cc.DlFunc[func(*GLContext), cc.Void]{Name: "gdk_gl_context_make_current"}
	gdk_gl_context_get_current            = cc.DlFunc[func() *GLContext, *GLContext]{Name: "gdk_gl_context_get_current"}
	gdk_gl_context_clear_current          = cc.DlFunc[func(), cc.Void]{Name: "gdk_gl_context_clear_current"}
	// #endregion

	// #region GLTexture
	gdk_gl_texture_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_gl_texture_get_type"}
	gdk_gl_texture_new      = cc.DlFunc[func(*GLContext, uint32, int32, int32, uptr, uptr) *Texture, *Texture]{Name: "gdk_gl_texture_new"}
	gdk_gl_texture_release  = cc.DlFunc[func(*Texture), cc.Void]{Name: "gdk_gl_texture_release"}
	// #endregion

	// #region GLTextureBuilder
	gdk_gl_texture_builder_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_gl_texture_builder_get_type"}
	gdk_gl_texture_builder_new                = cc.DlFunc[func() *GLTextureBuilder, *GLTextureBuilder]{Name: "gdk_gl_texture_builder_new"}
	gdk_gl_texture_builder_get_context        = cc.DlFunc[func(*GLTextureBuilder) *GLContext, *GLContext]{Name: "gdk_gl_texture_builder_get_context"}
	gdk_gl_texture_builder_set_context        = cc.DlFunc[func(*GLTextureBuilder, *GLContext), cc.Void]{Name: "gdk_gl_texture_builder_set_context"}
	gdk_gl_texture_builder_get_id             = cc.DlFunc[func(*GLTextureBuilder) uint32, uint32]{Name: "gdk_gl_texture_builder_get_id"}
	gdk_gl_texture_builder_set_id             = cc.DlFunc[func(*GLTextureBuilder, uint32), cc.Void]{Name: "gdk_gl_texture_builder_set_id"}
	gdk_gl_texture_builder_get_width          = cc.DlFunc[func(*GLTextureBuilder) int32, int32]{Name: "gdk_gl_texture_builder_get_width"}
	gdk_gl_texture_builder_set_width          = cc.DlFunc[func(*GLTextureBuilder, int32), cc.Void]{Name: "gdk_gl_texture_builder_set_width"}
	gdk_gl_texture_builder_get_height         = cc.DlFunc[func(*GLTextureBuilder) int32, int32]{Name: "gdk_gl_texture_builder_get_height"}
	gdk_gl_texture_builder_set_height         = cc.DlFunc[func(*GLTextureBuilder, int32), cc.Void]{Name: "gdk_gl_texture_builder_set_height"}
	gdk_gl_texture_builder_get_format         = cc.DlFunc[func(*GLTextureBuilder) MemoryFormat, MemoryFormat]{Name: "gdk_gl_texture_builder_get_format"}
	gdk_gl_texture_builder_set_format         = cc.DlFunc[func(*GLTextureBuilder, MemoryFormat), cc.Void]{Name: "gdk_gl_texture_builder_set_format"}
	gdk_gl_texture_builder_get_has_mipmap     = cc.DlFunc[func(*GLTextureBuilder) bool, bool]{Name: "gdk_gl_texture_builder_get_has_mipmap"}
	gdk_gl_texture_builder_set_has_mipmap     = cc.DlFunc[func(*GLTextureBuilder, bool), cc.Void]{Name: "gdk_gl_texture_builder_set_has_mipmap"}
	gdk_gl_texture_builder_get_sync           = cc.DlFunc[func(*GLTextureBuilder) uptr, uptr]{Name: "gdk_gl_texture_builder_get_sync"}
	gdk_gl_texture_builder_set_sync           = cc.DlFunc[func(*GLTextureBuilder, uptr), cc.Void]{Name: "gdk_gl_texture_builder_set_sync"}
	gdk_gl_texture_builder_get_color_state    = cc.DlFunc[func(*GLTextureBuilder) *ColorState, *ColorState]{Name: "gdk_gl_texture_builder_get_color_state"}
	gdk_gl_texture_builder_set_color_state    = cc.DlFunc[func(*GLTextureBuilder, *ColorState), cc.Void]{Name: "gdk_gl_texture_builder_set_color_state"}
	gdk_gl_texture_builder_get_update_texture = cc.DlFunc[func(*GLTextureBuilder) *Texture, *Texture]{Name: "gdk_gl_texture_builder_get_update_texture"}
	gdk_gl_texture_builder_set_update_texture = cc.DlFunc[func(*GLTextureBuilder, *Texture), cc.Void]{Name: "gdk_gl_texture_builder_set_update_texture"}
	gdk_gl_texture_builder_get_update_region  = cc.DlFunc[func(*GLTextureBuilder) *cairo.Region, *cairo.Region]{Name: "gdk_gl_texture_builder_get_update_region"}
	gdk_gl_texture_builder_set_update_region  = cc.DlFunc[func(*GLTextureBuilder, *cairo.Region), cc.Void]{Name: "gdk_gl_texture_builder_set_update_region"}
	gdk_gl_texture_builder_build              = cc.DlFunc[func(*GLTextureBuilder, uptr, uptr) *Texture, *Texture]{Name: "gdk_gl_texture_builder_build"}
	// #endregion

	// #region Keyval
	gdk_keyval_name         = cc.DlFunc[func(uint32) cc.String, cc.String]{Name: "gdk_keyval_name"}
	gdk_keyval_from_name    = cc.DlFunc[func(cc.String) uint32, uint32]{Name: "gdk_keyval_from_name"}
	gdk_keyval_convert_case = cc.DlFunc[func(uint32, *uint32, *uint32), cc.Void]{Name: "gdk_keyval_convert_case"}
	gdk_keyval_to_upper     = cc.DlFunc[func(uint32) uint32, uint32]{Name: "gdk_keyval_to_upper"}
	gdk_keyval_to_lower     = cc.DlFunc[func(uint32) uint32, uint32]{Name: "gdk_keyval_to_lower"}
	gdk_keyval_is_upper     = cc.DlFunc[func(uint32) bool, bool]{Name: "gdk_keyval_is_upper"}
	gdk_keyval_is_lower     = cc.DlFunc[func(uint32) bool, bool]{Name: "gdk_keyval_is_lower"}
	gdk_keyval_to_unicode   = cc.DlFunc[func(uint32) uint32, uint32]{Name: "gdk_keyval_to_unicode"}
	gdk_unicode_to_keyval   = cc.DlFunc[func(uint32) uint32, uint32]{Name: "gdk_unicode_to_keyval"}
	// #endregion

	// #region MemoryTexture
	gdk_memory_texture_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_memory_texture_get_type"}
	gdk_memory_texture_new      = cc.DlFunc[func(int32, int32, MemoryFormat, *glib.GBytes, uint64) *MemoryTexture, *MemoryTexture]{Name: "gdk_memory_texture_new"}
	// #endregion

	// #region MemoryTextureBuilder
	gdk_memory_texture_builder_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_memory_texture_builder_get_type"}
	gdk_memory_texture_builder_new                = cc.DlFunc[func() *MemoryTextureBuilder, *MemoryTextureBuilder]{Name: "gdk_memory_texture_builder_new"}
	gdk_memory_texture_builder_get_bytes          = cc.DlFunc[func(*MemoryTextureBuilder) *glib.GBytes, *glib.GBytes]{Name: "gdk_memory_texture_builder_get_bytes"}
	gdk_memory_texture_builder_set_bytes          = cc.DlFunc[func(*MemoryTextureBuilder, *glib.GBytes), cc.Void]{Name: "gdk_memory_texture_builder_set_bytes"}
	gdk_memory_texture_builder_get_stride         = cc.DlFunc[func(*MemoryTextureBuilder) uint64, uint64]{Name: "gdk_memory_texture_builder_get_stride"}
	gdk_memory_texture_builder_set_stride         = cc.DlFunc[func(*MemoryTextureBuilder, uint64), cc.Void]{Name: "gdk_memory_texture_builder_set_stride"}
	gdk_memory_texture_builder_get_width          = cc.DlFunc[func(*MemoryTextureBuilder) int32, int32]{Name: "gdk_memory_texture_builder_get_width"}
	gdk_memory_texture_builder_set_width          = cc.DlFunc[func(*MemoryTextureBuilder, int32), cc.Void]{Name: "gdk_memory_texture_builder_set_width"}
	gdk_memory_texture_builder_get_height         = cc.DlFunc[func(*MemoryTextureBuilder) int32, int32]{Name: "gdk_memory_texture_builder_get_height"}
	gdk_memory_texture_builder_set_height         = cc.DlFunc[func(*MemoryTextureBuilder, int32), cc.Void]{Name: "gdk_memory_texture_builder_set_height"}
	gdk_memory_texture_builder_get_format         = cc.DlFunc[func(*MemoryTextureBuilder) MemoryFormat, MemoryFormat]{Name: "gdk_memory_texture_builder_get_format"}
	gdk_memory_texture_builder_set_format         = cc.DlFunc[func(*MemoryTextureBuilder, MemoryFormat), cc.Void]{Name: "gdk_memory_texture_builder_set_format"}
	gdk_memory_texture_builder_get_color_state    = cc.DlFunc[func(*MemoryTextureBuilder) *ColorState, *ColorState]{Name: "gdk_memory_texture_builder_get_color_state"}
	gdk_memory_texture_builder_set_color_state    = cc.DlFunc[func(*MemoryTextureBuilder, *ColorState), cc.Void]{Name: "gdk_memory_texture_builder_set_color_state"}
	gdk_memory_texture_builder_get_update_texture = cc.DlFunc[func(*MemoryTextureBuilder) *Texture, *Texture]{Name: "gdk_memory_texture_builder_get_update_texture"}
	gdk_memory_texture_builder_set_update_texture = cc.DlFunc[func(*MemoryTextureBuilder, *Texture), cc.Void]{Name: "gdk_memory_texture_builder_set_update_texture"}
	gdk_memory_texture_builder_get_update_region  = cc.DlFunc[func(*MemoryTextureBuilder) *cairo.Region, *cairo.Region]{Name: "gdk_memory_texture_builder_get_update_region"}
	gdk_memory_texture_builder_set_update_region  = cc.DlFunc[func(*MemoryTextureBuilder, *cairo.Region), cc.Void]{Name: "gdk_memory_texture_builder_set_update_region"}
	gdk_memory_texture_builder_build              = cc.DlFunc[func(*MemoryTextureBuilder) *Texture, *Texture]{Name: "gdk_memory_texture_builder_build"}
	// #endregion

	// #region Monitor
	gdk_monitor_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_monitor_get_type"}
	gdk_monitor_get_display         = cc.DlFunc[func(*Monitor) *Display, *Display]{Name: "gdk_monitor_get_display"}
	gdk_monitor_get_geometry        = cc.DlFunc[func(*Monitor, *Rectangle), cc.Void]{Name: "gdk_monitor_get_geometry"}
	gdk_monitor_get_width_mm        = cc.DlFunc[func(*Monitor) int32, int32]{Name: "gdk_monitor_get_width_mm"}
	gdk_monitor_get_height_mm       = cc.DlFunc[func(*Monitor) int32, int32]{Name: "gdk_monitor_get_height_mm"}
	gdk_monitor_get_manufacturer    = cc.DlFunc[func(*Monitor) cc.String, cc.String]{Name: "gdk_monitor_get_manufacturer"}
	gdk_monitor_get_model           = cc.DlFunc[func(*Monitor) cc.String, cc.String]{Name: "gdk_monitor_get_model"}
	gdk_monitor_get_connector       = cc.DlFunc[func(*Monitor) cc.String, cc.String]{Name: "gdk_monitor_get_connector"}
	gdk_monitor_get_scale_factor    = cc.DlFunc[func(*Monitor) int32, int32]{Name: "gdk_monitor_get_scale_factor"}
	gdk_monitor_get_scale           = cc.DlFunc[func(*Monitor) float64, float64]{Name: "gdk_monitor_get_scale"}
	gdk_monitor_get_refresh_rate    = cc.DlFunc[func(*Monitor) int32, int32]{Name: "gdk_monitor_get_refresh_rate"}
	gdk_monitor_get_subpixel_layout = cc.DlFunc[func(*Monitor) SubpixelLayout, SubpixelLayout]{Name: "gdk_monitor_get_subpixel_layout"}
	gdk_monitor_is_valid            = cc.DlFunc[func(*Monitor) bool, bool]{Name: "gdk_monitor_is_valid"}
	gdk_monitor_get_description     = cc.DlFunc[func(*Monitor) cc.String, cc.String]{Name: "gdk_monitor_get_description"}
	// #endregion

	// #region Paintable
	gdk_paintable_get_type                   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_paintable_get_type"}
	gdk_paintable_snapshot                   = cc.DlFunc[func(*Paintable, *Snapshot, float64, float64), cc.Void]{Name: "gdk_paintable_snapshot"}
	gdk_paintable_get_current_image          = cc.DlFunc[func(*Paintable) *Paintable, *Paintable]{Name: "gdk_paintable_get_current_image"}
	gdk_paintable_get_flags                  = cc.DlFunc[func(*Paintable) PaintableFlags, PaintableFlags]{Name: "gdk_paintable_get_flags"}
	gdk_paintable_get_intrinsic_width        = cc.DlFunc[func(*Paintable) int32, int32]{Name: "gdk_paintable_get_intrinsic_width"}
	gdk_paintable_get_intrinsic_height       = cc.DlFunc[func(*Paintable) int32, int32]{Name: "gdk_paintable_get_intrinsic_height"}
	gdk_paintable_get_intrinsic_aspect_ratio = cc.DlFunc[func(*Paintable) float64, float64]{Name: "gdk_paintable_get_intrinsic_aspect_ratio"}
	gdk_paintable_compute_concrete_size      = cc.DlFunc[func(*Paintable, float64, float64, float64, float64, *float64, *float64), cc.Void]{Name: "gdk_paintable_compute_concrete_size"}
	gdk_paintable_invalidate_contents        = cc.DlFunc[func(*Paintable), cc.Void]{Name: "gdk_paintable_invalidate_contents"}
	gdk_paintable_invalidate_size            = cc.DlFunc[func(*Paintable), cc.Void]{Name: "gdk_paintable_invalidate_size"}
	gdk_paintable_new_empty                  = cc.DlFunc[func(int32, int32) *Paintable, *Paintable]{Name: "gdk_paintable_new_empty"}
	// #endregion

	// #region Pango
	gdk_pango_layout_line_get_clip_region = cc.DlFunc[func(*pango.LayoutLine, int32, int32, *int32, int32) *cairo.Region, *cairo.Region]{Name: "gdk_pango_layout_line_get_clip_region"}
	gdk_pango_layout_get_clip_region      = cc.DlFunc[func(*pango.Layout, int32, int32, *int32, int32) *cairo.Region, *cairo.Region]{Name: "gdk_pango_layout_get_clip_region"}
	// #endregion

	// #region Popup
	gdk_popup_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_popup_get_type"}
	gdk_popup_present            = cc.DlFunc[func(*Popup, int32, int32, *PopupLayout) int32, int32]{Name: "gdk_popup_present"}
	gdk_popup_get_surface_anchor = cc.DlFunc[func(*Popup) Gravity, Gravity]{Name: "gdk_popup_get_surface_anchor"}
	gdk_popup_get_rect_anchor    = cc.DlFunc[func(*Popup) Gravity, Gravity]{Name: "gdk_popup_get_rect_anchor"}
	gdk_popup_get_parent         = cc.DlFunc[func(*Popup) *Surface, *Surface]{Name: "gdk_popup_get_parent"}
	gdk_popup_get_position_x     = cc.DlFunc[func(*Popup) int32, int32]{Name: "gdk_popup_get_position_x"}
	gdk_popup_get_position_y     = cc.DlFunc[func(*Popup) int32, int32]{Name: "gdk_popup_get_position_y"}
	gdk_popup_get_autohide       = cc.DlFunc[func(*Popup) int32, int32]{Name: "gdk_popup_get_autohide"}
	// #endregion

	// #region PopupLayout
	gdk_popup_layout_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_popup_layout_get_type"}
	gdk_popup_layout_new                = cc.DlFunc[func(*Rectangle, Gravity, Gravity) *PopupLayout, *PopupLayout]{Name: "gdk_popup_layout_new"}
	gdk_popup_layout_ref                = cc.DlFunc[func(*PopupLayout) *PopupLayout, *PopupLayout]{Name: "gdk_popup_layout_ref"}
	gdk_popup_layout_unref              = cc.DlFunc[func(*PopupLayout), cc.Void]{Name: "gdk_popup_layout_unref"}
	gdk_popup_layout_copy               = cc.DlFunc[func(*PopupLayout) *PopupLayout, *PopupLayout]{Name: "gdk_popup_layout_copy"}
	gdk_popup_layout_equal              = cc.DlFunc[func(*PopupLayout, *PopupLayout) int32, int32]{Name: "gdk_popup_layout_equal"}
	gdk_popup_layout_set_anchor_rect    = cc.DlFunc[func(*PopupLayout, *Rectangle), cc.Void]{Name: "gdk_popup_layout_set_anchor_rect"}
	gdk_popup_layout_get_anchor_rect    = cc.DlFunc[func(*PopupLayout) *Rectangle, *Rectangle]{Name: "gdk_popup_layout_get_anchor_rect"}
	gdk_popup_layout_set_rect_anchor    = cc.DlFunc[func(*PopupLayout, Gravity), cc.Void]{Name: "gdk_popup_layout_set_rect_anchor"}
	gdk_popup_layout_get_rect_anchor    = cc.DlFunc[func(*PopupLayout) Gravity, Gravity]{Name: "gdk_popup_layout_get_rect_anchor"}
	gdk_popup_layout_set_surface_anchor = cc.DlFunc[func(*PopupLayout, Gravity), cc.Void]{Name: "gdk_popup_layout_set_surface_anchor"}
	gdk_popup_layout_get_surface_anchor = cc.DlFunc[func(*PopupLayout) Gravity, Gravity]{Name: "gdk_popup_layout_get_surface_anchor"}
	gdk_popup_layout_set_anchor_hints   = cc.DlFunc[func(*PopupLayout, AnchorHints), cc.Void]{Name: "gdk_popup_layout_set_anchor_hints"}
	gdk_popup_layout_get_anchor_hints   = cc.DlFunc[func(*PopupLayout) AnchorHints, AnchorHints]{Name: "gdk_popup_layout_get_anchor_hints"}
	gdk_popup_layout_set_offset         = cc.DlFunc[func(*PopupLayout, int32, int32), cc.Void]{Name: "gdk_popup_layout_set_offset"}
	gdk_popup_layout_get_offset         = cc.DlFunc[func(*PopupLayout, *int32, *int32), cc.Void]{Name: "gdk_popup_layout_get_offset"}
	gdk_popup_layout_set_shadow_width   = cc.DlFunc[func(*PopupLayout, int32, int32, int32, int32), cc.Void]{Name: "gdk_popup_layout_set_shadow_width"}
	gdk_popup_layout_get_shadow_width   = cc.DlFunc[func(*PopupLayout, *int32, *int32, *int32, *int32), cc.Void]{Name: "gdk_popup_layout_get_shadow_width"}
	// #endregion

	// #region Rectangle
	gdk_rectangle_intersect      = cc.DlFunc[func(*Rectangle, *Rectangle, *Rectangle) int32, int32]{Name: "gdk_rectangle_intersect"}
	gdk_rectangle_union          = cc.DlFunc[func(*Rectangle, *Rectangle, *Rectangle), cc.Void]{Name: "gdk_rectangle_union"}
	gdk_rectangle_equal          = cc.DlFunc[func(*Rectangle, *Rectangle) int32, int32]{Name: "gdk_rectangle_equal"}
	gdk_rectangle_contains_point = cc.DlFunc[func(*Rectangle, int32, int32) int32, int32]{Name: "gdk_rectangle_contains_point"}
	gdk_rectangle_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_rectangle_get_type"}
	// #endregion

	// #region RGBA
	gdk_rgba_get_type  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_rgba_get_type"}
	gdk_rgba_copy      = cc.DlFunc[func(*RGBA) *RGBA, *RGBA]{Name: "gdk_rgba_copy"}
	gdk_rgba_free      = cc.DlFunc[func(*RGBA), cc.Void]{Name: "gdk_rgba_free"}
	gdk_rgba_is_clear  = cc.DlFunc[func(*RGBA) int32, int32]{Name: "gdk_rgba_is_clear"}
	gdk_rgba_is_opaque = cc.DlFunc[func(*RGBA) int32, int32]{Name: "gdk_rgba_is_opaque"}
	gdk_rgba_hash      = cc.DlFunc[func(*RGBA) uint32, uint32]{Name: "gdk_rgba_hash"}
	gdk_rgba_equal     = cc.DlFunc[func(*RGBA, *RGBA) int32, int32]{Name: "gdk_rgba_equal"}
	gdk_rgba_parse     = cc.DlFunc[func(*RGBA, cc.String) int32, int32]{Name: "gdk_rgba_parse"}
	gdk_rgba_to_string = cc.DlFunc[func(*RGBA) cc.String, cc.String]{Name: "gdk_rgba_to_string"}
	// #endregion

	// #region Seat
	gdk_seat_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_seat_get_type"}
	gdk_seat_get_display      = cc.DlFunc[func(*Seat) *Display, *Display]{Name: "gdk_seat_get_display"}
	gdk_seat_get_capabilities = cc.DlFunc[func(*Seat) SeatCapabilities, SeatCapabilities]{Name: "gdk_seat_get_capabilities"}
	gdk_seat_get_devices      = cc.DlFunc[func(*Seat, SeatCapabilities) *glib.GList[Device], *glib.GList[Device]]{Name: "gdk_seat_get_devices"}
	gdk_seat_get_tools        = cc.DlFunc[func(*Seat) *glib.GList[DeviceTool], *glib.GList[DeviceTool]]{Name: "gdk_seat_get_tools"}
	gdk_seat_get_pointer      = cc.DlFunc[func(*Seat) *Device, *Device]{Name: "gdk_seat_get_pointer"}
	gdk_seat_get_keyboard     = cc.DlFunc[func(*Seat) *Device, *Device]{Name: "gdk_seat_get_keyboard"}
	// #endregion

	// #region Surface
	gdk_surface_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_surface_get_type"}
	gdk_surface_new_toplevel           = cc.DlFunc[func(*Display) *Surface, *Surface]{Name: "gdk_surface_new_toplevel"}
	gdk_surface_new_popup              = cc.DlFunc[func(*Surface, bool) *Surface, *Surface]{Name: "gdk_surface_new_popup"}
	gdk_surface_destroy                = cc.DlFunc[func(*Surface), cc.Void]{Name: "gdk_surface_destroy"}
	gdk_surface_is_destroyed           = cc.DlFunc[func(*Surface) int32, int32]{Name: "gdk_surface_is_destroyed"}
	gdk_surface_get_display            = cc.DlFunc[func(*Surface) *Display, *Display]{Name: "gdk_surface_get_display"}
	gdk_surface_hide                   = cc.DlFunc[func(*Surface), cc.Void]{Name: "gdk_surface_hide"}
	gdk_surface_set_input_region       = cc.DlFunc[func(*Surface, *cairo.Region), cc.Void]{Name: "gdk_surface_set_input_region"}
	gdk_surface_get_mapped             = cc.DlFunc[func(*Surface) int32, int32]{Name: "gdk_surface_get_mapped"}
	gdk_surface_set_cursor             = cc.DlFunc[func(*Surface, *Cursor), cc.Void]{Name: "gdk_surface_set_cursor"}
	gdk_surface_get_cursor             = cc.DlFunc[func(*Surface) *Cursor, *Cursor]{Name: "gdk_surface_get_cursor"}
	gdk_surface_set_device_cursor      = cc.DlFunc[func(*Surface, *Device, *Cursor), cc.Void]{Name: "gdk_surface_set_device_cursor"}
	gdk_surface_get_device_cursor      = cc.DlFunc[func(*Surface, *Device) *Cursor, *Cursor]{Name: "gdk_surface_get_device_cursor"}
	gdk_surface_get_width              = cc.DlFunc[func(*Surface) int32, int32]{Name: "gdk_surface_get_width"}
	gdk_surface_get_height             = cc.DlFunc[func(*Surface) int32, int32]{Name: "gdk_surface_get_height"}
	gdk_surface_translate_coordinates  = cc.DlFunc[func(*Surface, *Surface, *float64, *float64) int32, int32]{Name: "gdk_surface_translate_coordinates"}
	gdk_surface_get_scale_factor       = cc.DlFunc[func(*Surface) int32, int32]{Name: "gdk_surface_get_scale_factor"}
	gdk_surface_get_scale              = cc.DlFunc[func(*Surface) float64, float64]{Name: "gdk_surface_get_scale"}
	gdk_surface_get_device_position    = cc.DlFunc[func(*Surface, *Device, *float64, *float64, *ModifierType) int32, int32]{Name: "gdk_surface_get_device_position"}
	gdk_surface_create_similar_surface = cc.DlFunc[func(*Surface, cairo.Content, int32, int32) *cairo.Surface, *cairo.Surface]{Name: "gdk_surface_create_similar_surface"}
	gdk_surface_beep                   = cc.DlFunc[func(*Surface), cc.Void]{Name: "gdk_surface_beep"}
	gdk_surface_queue_render           = cc.DlFunc[func(*Surface), cc.Void]{Name: "gdk_surface_queue_render"}
	gdk_surface_request_layout         = cc.DlFunc[func(*Surface), cc.Void]{Name: "gdk_surface_request_layout"}
	gdk_surface_get_frame_clock        = cc.DlFunc[func(*Surface) *FrameClock, *FrameClock]{Name: "gdk_surface_get_frame_clock"}
	gdk_surface_set_opaque_region      = cc.DlFunc[func(*Surface, *cairo.Region), cc.Void]{Name: "gdk_surface_set_opaque_region"}
	gdk_surface_create_cairo_context   = cc.DlFunc[func(*Surface) *cairo.Context, *cairo.Context]{Name: "gdk_surface_create_cairo_context"}
	gdk_surface_create_gl_context      = cc.DlFunc[func(*Surface, **glib.GError) *GLContext, *GLContext]{Name: "gdk_surface_create_gl_context"}
	// #endregion

	// #region Texture
	gdk_texture_error_quark        = cc.DlFunc[func() glib.GQuark, glib.GQuark]{Name: "gdk_texture_error_quark"}
	gdk_texture_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_texture_get_type"}
	gdk_texture_new_for_pixbuf     = cc.DlFunc[func(*Pixbuf) *Texture, *Texture]{Name: "gdk_texture_new_for_pixbuf"}
	gdk_texture_new_from_resource  = cc.DlFunc[func(cc.String) *Texture, *Texture]{Name: "gdk_texture_new_from_resource"}
	gdk_texture_new_from_file      = cc.DlFunc[func(*gio.GFile, **glib.GError) *Texture, *Texture]{Name: "gdk_texture_new_from_file"}
	gdk_texture_new_from_filename  = cc.DlFunc[func(cc.String, **glib.GError) *Texture, *Texture]{Name: "gdk_texture_new_from_filename"}
	gdk_texture_new_from_bytes     = cc.DlFunc[func(*glib.GBytes, **glib.GError) *Texture, *Texture]{Name: "gdk_texture_new_from_bytes"}
	gdk_texture_get_width          = cc.DlFunc[func(*Texture) int32, int32]{Name: "gdk_texture_get_width"}
	gdk_texture_get_height         = cc.DlFunc[func(*Texture) int32, int32]{Name: "gdk_texture_get_height"}
	gdk_texture_get_format         = cc.DlFunc[func(*Texture) MemoryFormat, MemoryFormat]{Name: "gdk_texture_get_format"}
	gdk_texture_get_color_state    = cc.DlFunc[func(*Texture) *ColorState, *ColorState]{Name: "gdk_texture_get_color_state"}
	gdk_texture_download           = cc.DlFunc[func(*Texture, uptr, uint64), cc.Void]{Name: "gdk_texture_download"}
	gdk_texture_save_to_png        = cc.DlFunc[func(*Texture, cc.String) int32, int32]{Name: "gdk_texture_save_to_png"}
	gdk_texture_save_to_png_bytes  = cc.DlFunc[func(*Texture) *glib.GBytes, *glib.GBytes]{Name: "gdk_texture_save_to_png_bytes"}
	gdk_texture_save_to_tiff       = cc.DlFunc[func(*Texture, cc.String) int32, int32]{Name: "gdk_texture_save_to_tiff"}
	gdk_texture_save_to_tiff_bytes = cc.DlFunc[func(*Texture) *glib.GBytes, *glib.GBytes]{Name: "gdk_texture_save_to_tiff_bytes"}
	// #endregion

	// #region TextureDownloader
	gdk_texture_downloader_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_texture_downloader_get_type"}
	gdk_texture_downloader_new             = cc.DlFunc[func(*Texture) *TextureDownloader, *TextureDownloader]{Name: "gdk_texture_downloader_new"}
	gdk_texture_downloader_copy            = cc.DlFunc[func(*TextureDownloader) *TextureDownloader, *TextureDownloader]{Name: "gdk_texture_downloader_copy"}
	gdk_texture_downloader_free            = cc.DlFunc[func(*TextureDownloader), cc.Void]{Name: "gdk_texture_downloader_free"}
	gdk_texture_downloader_set_texture     = cc.DlFunc[func(*TextureDownloader, *Texture), cc.Void]{Name: "gdk_texture_downloader_set_texture"}
	gdk_texture_downloader_get_texture     = cc.DlFunc[func(*TextureDownloader) *Texture, *Texture]{Name: "gdk_texture_downloader_get_texture"}
	gdk_texture_downloader_set_format      = cc.DlFunc[func(*TextureDownloader, MemoryFormat), cc.Void]{Name: "gdk_texture_downloader_set_format"}
	gdk_texture_downloader_get_format      = cc.DlFunc[func(*TextureDownloader) MemoryFormat, MemoryFormat]{Name: "gdk_texture_downloader_get_format"}
	gdk_texture_downloader_set_color_state = cc.DlFunc[func(*TextureDownloader, *ColorState), cc.Void]{Name: "gdk_texture_downloader_set_color_state"}
	gdk_texture_downloader_get_color_state = cc.DlFunc[func(*TextureDownloader) *ColorState, *ColorState]{Name: "gdk_texture_downloader_get_color_state"}
	gdk_texture_downloader_download_into   = cc.DlFunc[func(*TextureDownloader, uptr, uint64), cc.Void]{Name: "gdk_texture_downloader_download_into"}
	gdk_texture_downloader_download_bytes  = cc.DlFunc[func(*TextureDownloader, *uint64) *glib.GBytes, *glib.GBytes]{Name: "gdk_texture_downloader_download_bytes"}
	// #endregion

	// #region Toplevel
	gdk_toplevel_get_type                  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_toplevel_get_type"}
	gdk_toplevel_present                   = cc.DlFunc[func(*Toplevel, *ToplevelLayout), cc.Void]{Name: "gdk_toplevel_present"}
	gdk_toplevel_minimize                  = cc.DlFunc[func(*Toplevel) int32, int32]{Name: "gdk_toplevel_minimize"}
	gdk_toplevel_lower                     = cc.DlFunc[func(*Toplevel) int32, int32]{Name: "gdk_toplevel_lower"}
	gdk_toplevel_focus                     = cc.DlFunc[func(*Toplevel, uint32), cc.Void]{Name: "gdk_toplevel_focus"}
	gdk_toplevel_get_state                 = cc.DlFunc[func(*Toplevel) ToplevelState, ToplevelState]{Name: "gdk_toplevel_get_state"}
	gdk_toplevel_set_title                 = cc.DlFunc[func(*Toplevel, cc.String), cc.Void]{Name: "gdk_toplevel_set_title"}
	gdk_toplevel_set_startup_id            = cc.DlFunc[func(*Toplevel, cc.String), cc.Void]{Name: "gdk_toplevel_set_startup_id"}
	gdk_toplevel_set_transient_for         = cc.DlFunc[func(*Toplevel, *Surface), cc.Void]{Name: "gdk_toplevel_set_transient_for"}
	gdk_toplevel_set_modal                 = cc.DlFunc[func(*Toplevel, bool), cc.Void]{Name: "gdk_toplevel_set_modal"}
	gdk_toplevel_set_icon_list             = cc.DlFunc[func(*Toplevel, uptr), cc.Void]{Name: "gdk_toplevel_set_icon_list"}
	gdk_toplevel_show_window_menu          = cc.DlFunc[func(*Toplevel, *Event) int32, int32]{Name: "gdk_toplevel_show_window_menu"}
	gdk_toplevel_set_decorated             = cc.DlFunc[func(*Toplevel, bool), cc.Void]{Name: "gdk_toplevel_set_decorated"}
	gdk_toplevel_set_deletable             = cc.DlFunc[func(*Toplevel, bool), cc.Void]{Name: "gdk_toplevel_set_deletable"}
	gdk_toplevel_supports_edge_constraints = cc.DlFunc[func(*Toplevel) int32, int32]{Name: "gdk_toplevel_supports_edge_constraints"}
	gdk_toplevel_inhibit_system_shortcuts  = cc.DlFunc[func(*Toplevel, *Event), cc.Void]{Name: "gdk_toplevel_inhibit_system_shortcuts"}
	gdk_toplevel_restore_system_shortcuts  = cc.DlFunc[func(*Toplevel), cc.Void]{Name: "gdk_toplevel_restore_system_shortcuts"}
	gdk_toplevel_begin_resize              = cc.DlFunc[func(*Toplevel, SurfaceEdge, *Device, int32, float64, float64, uint32), cc.Void]{Name: "gdk_toplevel_begin_resize"}
	gdk_toplevel_begin_move                = cc.DlFunc[func(*Toplevel, *Device, int32, float64, float64, uint32), cc.Void]{Name: "gdk_toplevel_begin_move"}
	gdk_toplevel_titlebar_gesture          = cc.DlFunc[func(*Toplevel, TitlebarGesture) int32, int32]{Name: "gdk_toplevel_titlebar_gesture"}
	// #endregion

	// #region ToplevelLayout
	gdk_toplevel_layout_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_toplevel_layout_get_type"}
	gdk_toplevel_layout_new                    = cc.DlFunc[func() *ToplevelLayout, *ToplevelLayout]{Name: "gdk_toplevel_layout_new"}
	gdk_toplevel_layout_ref                    = cc.DlFunc[func(*ToplevelLayout) *ToplevelLayout, *ToplevelLayout]{Name: "gdk_toplevel_layout_ref"}
	gdk_toplevel_layout_unref                  = cc.DlFunc[func(*ToplevelLayout), cc.Void]{Name: "gdk_toplevel_layout_unref"}
	gdk_toplevel_layout_copy                   = cc.DlFunc[func(*ToplevelLayout) *ToplevelLayout, *ToplevelLayout]{Name: "gdk_toplevel_layout_copy"}
	gdk_toplevel_layout_equal                  = cc.DlFunc[func(*ToplevelLayout, *ToplevelLayout) int32, int32]{Name: "gdk_toplevel_layout_equal"}
	gdk_toplevel_layout_set_maximized          = cc.DlFunc[func(*ToplevelLayout, bool), cc.Void]{Name: "gdk_toplevel_layout_set_maximized"}
	gdk_toplevel_layout_set_fullscreen         = cc.DlFunc[func(*ToplevelLayout, bool, *Monitor), cc.Void]{Name: "gdk_toplevel_layout_set_fullscreen"}
	gdk_toplevel_layout_get_maximized          = cc.DlFunc[func(*ToplevelLayout, *bool) int32, int32]{Name: "gdk_toplevel_layout_get_maximized"}
	gdk_toplevel_layout_get_fullscreen         = cc.DlFunc[func(*ToplevelLayout, *bool) int32, int32]{Name: "gdk_toplevel_layout_get_fullscreen"}
	gdk_toplevel_layout_get_fullscreen_monitor = cc.DlFunc[func(*ToplevelLayout) *Monitor, *Monitor]{Name: "gdk_toplevel_layout_get_fullscreen_monitor"}
	gdk_toplevel_layout_set_resizable          = cc.DlFunc[func(*ToplevelLayout, bool), cc.Void]{Name: "gdk_toplevel_layout_set_resizable"}
	gdk_toplevel_layout_get_resizable          = cc.DlFunc[func(*ToplevelLayout) int32, int32]{Name: "gdk_toplevel_layout_get_resizable"}
	// #endregion

	// #region ToplevelSize
	gdk_toplevel_size_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_toplevel_size_get_type"}
	gdk_toplevel_size_get_bounds       = cc.DlFunc[func(*ToplevelSize, *int, *int), cc.Void]{Name: "gdk_toplevel_size_get_bounds"}
	gdk_toplevel_size_set_size         = cc.DlFunc[func(*ToplevelSize, int, int), cc.Void]{Name: "gdk_toplevel_size_set_size"}
	gdk_toplevel_size_set_min_size     = cc.DlFunc[func(*ToplevelSize, int, int), cc.Void]{Name: "gdk_toplevel_size_set_min_size"}
	gdk_toplevel_size_set_shadow_width = cc.DlFunc[func(*ToplevelSize, int, int, int, int), cc.Void]{Name: "gdk_toplevel_size_set_shadow_width"}
	// #endregion

)
