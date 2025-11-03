package gio

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/glib/glib"
	"github.com/jinzhongmin/gg/glib/gobject"
)

type uptr = unsafe.Pointer
type iptr = uintptr

func vcbu(fn interface{}) uptr  { return cc.Cbk(fn) }
func anyptr(a interface{}) uptr { return (*(*[2]uptr)(uptr(&a)))[1] }

func init() { cc.Open("libgio-2*") }

var (

	// #region Action
	g_action_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "g_action_get_type"}
	g_action_get_name           = cc.DlFunc[func(action *GAction) string, string]{Name: "g_action_get_name"}
	g_action_get_parameter_type = cc.DlFunc[func(action *GAction) *glib.GVariantType, *glib.GVariantType]{Name: "g_action_get_parameter_type"}
	g_action_get_state_type     = cc.DlFunc[func(action *GAction) *glib.GVariantType, *glib.GVariantType]{Name: "g_action_get_state_type"}
	g_action_get_state_hint     = cc.DlFunc[func(action *GAction) *glib.GVariant, *glib.GVariant]{Name: "g_action_get_state_hint"}
	g_action_get_enabled        = cc.DlFunc[func(action *GAction) bool, bool]{Name: "g_action_get_enabled"}
	g_action_get_state          = cc.DlFunc[func(action *GAction) *glib.GVariant, *glib.GVariant]{Name: "g_action_get_state"}
	g_action_change_state       = cc.DlFunc[func(action *GAction, value *glib.GVariant), cc.Void]{Name: "g_action_change_state"}
	g_action_activate           = cc.DlFunc[func(action *GAction, parameter *glib.GVariant), cc.Void]{Name: "g_action_activate"}
	g_action_name_is_valid      = cc.DlFunc[func(actionName string) bool, bool]{Name: "g_action_name_is_valid"}
	// #endregion

	// #region ActionGroup
	g_action_group_get_type                  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "g_action_group_get_type"}
	g_action_group_has_action                = cc.DlFunc[func(actionGroup *GActionGroup, actionName string) bool, bool]{Name: "g_action_group_has_action"}
	g_action_group_list_actions              = cc.DlFunc[func(actionGroup *GActionGroup) cc.Strings, cc.Strings]{Name: "g_action_group_list_actions"}
	g_action_group_get_action_parameter_type = cc.DlFunc[func(actionGroup *GActionGroup, actionName string) *glib.GVariantType, *glib.GVariantType]{Name: "g_action_group_get_action_parameter_type"}
	g_action_group_get_action_state_type     = cc.DlFunc[func(actionGroup *GActionGroup, actionName string) *glib.GVariantType, *glib.GVariantType]{Name: "g_action_group_get_action_state_type"}
	g_action_group_get_action_state_hint     = cc.DlFunc[func(actionGroup *GActionGroup, actionName string) *glib.GVariant, *glib.GVariant]{Name: "g_action_group_get_action_state_hint"}
	// #endregion

	// #region Application
	g_application_get_type                            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "g_application_get_type"}
	g_application_id_is_valid                         = cc.DlFunc[func(applicationID string) bool, bool]{Name: "g_application_id_is_valid"}
	g_application_new                                 = cc.DlFunc[func(applicationID string, flags GApplicationFlags) *GApplication, *GApplication]{Name: "g_application_new"}
	g_application_get_application_id                  = cc.DlFunc[func(application *GApplication) string, string]{Name: "g_application_get_application_id"}
	g_application_set_application_id                  = cc.DlFunc[func(application *GApplication, applicationID string), cc.Void]{Name: "g_application_set_application_id"}
	g_application_get_version                         = cc.DlFunc[func(application *GApplication) string, string]{Name: "g_application_get_version"}
	g_application_set_version                         = cc.DlFunc[func(application *GApplication, version string), cc.Void]{Name: "g_application_set_version"}
	g_application_get_dbus_connection                 = cc.DlFunc[func(application *GApplication) uptr, uptr]{Name: "g_application_get_dbus_connection"}
	g_application_get_dbus_object_path                = cc.DlFunc[func(application *GApplication) string, string]{Name: "g_application_get_dbus_object_path"}
	g_application_get_inactivity_timeout              = cc.DlFunc[func(application *GApplication) uint32, uint32]{Name: "g_application_get_inactivity_timeout"}
	g_application_set_inactivity_timeout              = cc.DlFunc[func(application *GApplication, timeout uint32), cc.Void]{Name: "g_application_set_inactivity_timeout"}
	g_application_get_flags                           = cc.DlFunc[func(application *GApplication) uint32, uint32]{Name: "g_application_get_flags"}
	g_application_set_flags                           = cc.DlFunc[func(application *GApplication, flags uint32), cc.Void]{Name: "g_application_set_flags"}
	g_application_get_resource_base_path              = cc.DlFunc[func(application *GApplication) string, string]{Name: "g_application_get_resource_base_path"}
	g_application_set_resource_base_path              = cc.DlFunc[func(application *GApplication, resourcePath string), cc.Void]{Name: "g_application_set_resource_base_path"}
	g_application_set_action_group                    = cc.DlFunc[func(application *GApplication, actionGroup uptr), cc.Void]{Name: "g_application_set_action_group"}
	g_application_add_main_option_entries             = cc.DlFunc[func(application *GApplication, entries uptr), cc.Void]{Name: "g_application_add_main_option_entries"}
	g_application_add_main_option                     = cc.DlFunc[func(application *GApplication, longName string, shortName byte, flags uint32, arg uint32, description string, argDescription string), cc.Void]{Name: "g_application_add_main_option"}
	g_application_add_option_group                    = cc.DlFunc[func(application *GApplication, group uptr), cc.Void]{Name: "g_application_add_option_group"}
	g_application_set_option_context_parameter_string = cc.DlFunc[func(application *GApplication, parameterString string), cc.Void]{Name: "g_application_set_option_context_parameter_string"}
	g_application_set_option_context_summary          = cc.DlFunc[func(application *GApplication, summary string), cc.Void]{Name: "g_application_set_option_context_summary"}
	g_application_set_option_context_description      = cc.DlFunc[func(application *GApplication, description string), cc.Void]{Name: "g_application_set_option_context_description"}
	g_application_get_is_registered                   = cc.DlFunc[func(application *GApplication) bool, bool]{Name: "g_application_get_is_registered"}
	g_application_get_is_remote                       = cc.DlFunc[func(application *GApplication) bool, bool]{Name: "g_application_get_is_remote"}
	g_application_register                            = cc.DlFunc[func(application *GApplication, cancellable uptr, error **uptr) bool, bool]{Name: "g_application_register"}
	g_application_hold                                = cc.DlFunc[func(application *GApplication), cc.Void]{Name: "g_application_hold"}
	g_application_release                             = cc.DlFunc[func(application *GApplication), cc.Void]{Name: "g_application_release"}
	g_application_activate                            = cc.DlFunc[func(application *GApplication), cc.Void]{Name: "g_application_activate"}
	g_application_open                                = cc.DlFunc[func(application *GApplication, files **uptr, nFiles int32, hint string), cc.Void]{Name: "g_application_open"}
	g_application_run                                 = cc.DlFunc[func(application *GApplication, argc int32, argv cc.Strings) int32, int32]{Name: "g_application_run"}
	g_application_quit                                = cc.DlFunc[func(application *GApplication), cc.Void]{Name: "g_application_quit"}
	g_application_get_default                         = cc.DlFunc[func() *GApplicationObj, *GApplicationObj]{Name: "g_application_get_default"}
	g_application_set_default                         = cc.DlFunc[func(application *GApplication), cc.Void]{Name: "g_application_set_default"}
	g_application_mark_busy                           = cc.DlFunc[func(application *GApplication), cc.Void]{Name: "g_application_mark_busy"}
	g_application_unmark_busy                         = cc.DlFunc[func(application *GApplication), cc.Void]{Name: "g_application_unmark_busy"}
	g_application_get_is_busy                         = cc.DlFunc[func(application *GApplication) bool, bool]{Name: "g_application_get_is_busy"}
	g_application_send_notification                   = cc.DlFunc[func(application *GApplication, id string, notification uptr), cc.Void]{Name: "g_application_send_notification"}
	g_application_withdraw_notification               = cc.DlFunc[func(application *GApplication, id string), cc.Void]{Name: "g_application_withdraw_notification"}
	g_application_bind_busy_property                  = cc.DlFunc[func(application *GApplication, object uptr, property string), cc.Void]{Name: "g_application_bind_busy_property"}
	g_application_unbind_busy_property                = cc.DlFunc[func(application *GApplication, object uptr, property string), cc.Void]{Name: "g_application_unbind_busy_property"}
	// #endregion

	// #region Cancellable
	g_cancellable_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "g_cancellable_get_type"}
	g_cancellable_new                    = cc.DlFunc[func() *GCancellable, *GCancellable]{Name: "g_cancellable_new"}
	g_cancellable_is_cancelled           = cc.DlFunc[func(cancellable *GCancellable) bool, bool]{Name: "g_cancellable_is_cancelled"}
	g_cancellable_set_error_if_cancelled = cc.DlFunc[func(cancellable *GCancellable, err **uptr) bool, bool]{Name: "g_cancellable_set_error_if_cancelled"}
	g_cancellable_get_fd                 = cc.DlFunc[func(cancellable *GCancellable) int32, int32]{Name: "g_cancellable_get_fd"}
	g_cancellable_make_pollfd            = cc.DlFunc[func(cancellable *GCancellable, pollfd *glib.GPollFD) bool, bool]{Name: "g_cancellable_make_pollfd"}
	g_cancellable_release_fd             = cc.DlFunc[func(cancellable *GCancellable), cc.Void]{Name: "g_cancellable_release_fd"}
	g_cancellable_source_new             = cc.DlFunc[func(cancellable *GCancellable) *glib.GSource, *glib.GSource]{Name: "g_cancellable_source_new"}
	g_cancellable_get_current            = cc.DlFunc[func() *GCancellable, *GCancellable]{Name: "g_cancellable_get_current"}
	g_cancellable_push_current           = cc.DlFunc[func(cancellable *GCancellable), cc.Void]{Name: "g_cancellable_push_current"}
	g_cancellable_pop_current            = cc.DlFunc[func(cancellable *GCancellable), cc.Void]{Name: "g_cancellable_pop_current"}
	g_cancellable_reset                  = cc.DlFunc[func(cancellable *GCancellable), cc.Void]{Name: "g_cancellable_reset"}
	g_cancellable_connect                = cc.DlFunc[func(cancellable *GCancellable, callback uptr, data uptr, dataDestroyFunc uptr) uint64, uint64]{Name: "g_cancellable_connect"}
	g_cancellable_disconnect             = cc.DlFunc[func(cancellable *GCancellable, handlerID uint64), cc.Void]{Name: "g_cancellable_disconnect"}
	g_cancellable_cancel                 = cc.DlFunc[func(cancellable *GCancellable), cc.Void]{Name: "g_cancellable_cancel"}
	// #endregion

	// #region File
	// #endregion

	// #region FileInfo
	g_file_info_get_type                  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "g_file_info_get_type"}
	g_file_info_new                       = cc.DlFunc[func() *GFileInfo, *GFileInfo]{Name: "g_file_info_new"}
	g_file_info_dup                       = cc.DlFunc[func(info *GFileInfo) *GFileInfo, *GFileInfo]{Name: "g_file_info_dup"}
	g_file_info_copy_into                 = cc.DlFunc[func(src, dest *GFileInfo), cc.Void]{Name: "g_file_info_copy_into"}
	g_file_info_has_attribute             = cc.DlFunc[func(info *GFileInfo, attribute cc.String) bool, bool]{Name: "g_file_info_has_attribute"}
	g_file_info_has_namespace             = cc.DlFunc[func(info *GFileInfo, name_space cc.String) bool, bool]{Name: "g_file_info_has_namespace"}
	g_file_info_list_attributes           = cc.DlFunc[func(info *GFileInfo, name_space cc.String) cc.Strings, cc.Strings]{Name: "g_file_info_list_attributes"}
	g_file_info_get_attribute_data        = cc.DlFunc[func(info *GFileInfo, attribute cc.String, typ *GFileAttributeType, value_pp *uptr, status *GFileAttributeStatus) bool, bool]{Name: "g_file_info_get_attribute_data"}
	g_file_info_get_attribute_type        = cc.DlFunc[func(info *GFileInfo, attribute cc.String) GFileAttributeType, GFileAttributeType]{Name: "g_file_info_get_attribute_type"}
	g_file_info_remove_attribute          = cc.DlFunc[func(info *GFileInfo, attribute cc.String), cc.Void]{Name: "g_file_info_remove_attribute"}
	g_file_info_get_attribute_status      = cc.DlFunc[func(info *GFileInfo, attribute cc.String) GFileAttributeStatus, GFileAttributeStatus]{Name: "g_file_info_get_attribute_status"}
	g_file_info_set_attribute_status      = cc.DlFunc[func(info *GFileInfo, attribute cc.String, status GFileAttributeStatus) bool, bool]{Name: "g_file_info_set_attribute_status"}
	g_file_info_get_attribute_as_string   = cc.DlFunc[func(info *GFileInfo, attribute cc.String) cc.String, cc.String]{Name: "g_file_info_get_attribute_as_string"}
	g_file_info_get_attribute_string      = cc.DlFunc[func(info *GFileInfo, attribute cc.String) cc.String, cc.String]{Name: "g_file_info_get_attribute_string"}
	g_file_info_get_attribute_byte_string = cc.DlFunc[func(info *GFileInfo, attribute cc.String) cc.String, cc.String]{Name: "g_file_info_get_attribute_byte_string"}
	g_file_info_get_attribute_boolean     = cc.DlFunc[func(info *GFileInfo, attribute cc.String) bool, bool]{Name: "g_file_info_get_attribute_boolean"}
	g_file_info_get_attribute_uint32      = cc.DlFunc[func(info *GFileInfo, attribute cc.String) uint32, uint32]{Name: "g_file_info_get_attribute_uint32"}
	g_file_info_get_attribute_int32       = cc.DlFunc[func(info *GFileInfo, attribute cc.String) int32, int32]{Name: "g_file_info_get_attribute_int32"}
	g_file_info_get_attribute_uint64      = cc.DlFunc[func(info *GFileInfo, attribute cc.String) uint64, uint64]{Name: "g_file_info_get_attribute_uint64"}
	g_file_info_get_attribute_int64       = cc.DlFunc[func(info *GFileInfo, attribute cc.String) int64, int64]{Name: "g_file_info_get_attribute_int64"}
	g_file_info_get_attribute_object      = cc.DlFunc[func(info *GFileInfo, attribute cc.String) uptr, uptr]{Name: "g_file_info_get_attribute_object"}
	g_file_info_get_attribute_stringv     = cc.DlFunc[func(info *GFileInfo, attribute cc.String) cc.Strings, cc.Strings]{Name: "g_file_info_get_attribute_stringv"}
	g_file_info_get_attribute_file_path   = cc.DlFunc[func(info *GFileInfo, attribute cc.String) cc.String, cc.String]{Name: "g_file_info_get_attribute_file_path"}

	g_file_info_set_attribute             = cc.DlFunc[func(info *GFileInfo, attribute cc.String, typ GFileAttributeType, value_p uptr), cc.Void]{Name: "g_file_info_set_attribute"}
	g_file_info_set_attribute_string      = cc.DlFunc[func(info *GFileInfo, attribute cc.String, attr_value cc.String), cc.Void]{Name: "g_file_info_set_attribute_string"}
	g_file_info_set_attribute_byte_string = cc.DlFunc[func(info *GFileInfo, attribute cc.String, attr_value cc.String), cc.Void]{Name: "g_file_info_set_attribute_byte_string"}
	g_file_info_set_attribute_boolean     = cc.DlFunc[func(info *GFileInfo, attribute cc.String, attr_value bool), cc.Void]{Name: "g_file_info_set_attribute_boolean"}
	g_file_info_set_attribute_uint32      = cc.DlFunc[func(info *GFileInfo, attribute cc.String, attr_value uint32), cc.Void]{Name: "g_file_info_set_attribute_uint32"}
	g_file_info_set_attribute_int32       = cc.DlFunc[func(info *GFileInfo, attribute cc.String, attr_value int32), cc.Void]{Name: "g_file_info_set_attribute_int32"}
	g_file_info_set_attribute_uint64      = cc.DlFunc[func(info *GFileInfo, attribute cc.String, attr_value uint64), cc.Void]{Name: "g_file_info_set_attribute_uint64"}
	g_file_info_set_attribute_int64       = cc.DlFunc[func(info *GFileInfo, attribute cc.String, attr_value int64), cc.Void]{Name: "g_file_info_set_attribute_int64"}
	g_file_info_set_attribute_object      = cc.DlFunc[func(info *GFileInfo, attribute cc.String, attr_value *gobject.GObject), cc.Void]{Name: "g_file_info_set_attribute_object"}
	g_file_info_set_attribute_stringv     = cc.DlFunc[func(info *GFileInfo, attribute cc.String, attr_value cc.Strings), cc.Void]{Name: "g_file_info_set_attribute_stringv"}
	g_file_info_set_attribute_file_path   = cc.DlFunc[func(info *GFileInfo, attribute cc.String, attr_value cc.String), cc.Void]{Name: "g_file_info_set_attribute_file_path"}

	g_file_info_clear_status = cc.DlFunc[func(info *GFileInfo), cc.Void]{Name: "g_file_info_clear_status"}

	g_file_info_get_deletion_date          = cc.DlFunc[func(info *GFileInfo) uptr, uptr]{Name: "g_file_info_get_deletion_date"}
	g_file_info_get_file_type              = cc.DlFunc[func(info *GFileInfo) GFileType, GFileType]{Name: "g_file_info_get_file_type"}
	g_file_info_get_is_hidden              = cc.DlFunc[func(info *GFileInfo) bool, bool]{Name: "g_file_info_get_is_hidden"}
	g_file_info_get_is_backup              = cc.DlFunc[func(info *GFileInfo) bool, bool]{Name: "g_file_info_get_is_backup"}
	g_file_info_get_is_symlink             = cc.DlFunc[func(info *GFileInfo) bool, bool]{Name: "g_file_info_get_is_symlink"}
	g_file_info_get_name                   = cc.DlFunc[func(info *GFileInfo) cc.String, cc.String]{Name: "g_file_info_get_name"}
	g_file_info_get_display_name           = cc.DlFunc[func(info *GFileInfo) cc.String, cc.String]{Name: "g_file_info_get_display_name"}
	g_file_info_get_edit_name              = cc.DlFunc[func(info *GFileInfo) cc.String, cc.String]{Name: "g_file_info_get_edit_name"}
	g_file_info_get_icon                   = cc.DlFunc[func(info *GFileInfo) *GIcon, *GIcon]{Name: "g_file_info_get_icon"}
	g_file_info_get_symbolic_icon          = cc.DlFunc[func(info *GFileInfo) *GIcon, *GIcon]{Name: "g_file_info_get_symbolic_icon"}
	g_file_info_get_content_type           = cc.DlFunc[func(info *GFileInfo) cc.String, cc.String]{Name: "g_file_info_get_content_type"}
	g_file_info_get_size                   = cc.DlFunc[func(info *GFileInfo) int64, int64]{Name: "g_file_info_get_size"}
	g_file_info_get_modification_date_time = cc.DlFunc[func(info *GFileInfo) *glib.GDateTime, *glib.GDateTime]{Name: "g_file_info_get_modification_date_time"}
	g_file_info_get_access_date_time       = cc.DlFunc[func(info *GFileInfo) *glib.GDateTime, *glib.GDateTime]{Name: "g_file_info_get_access_date_time"}
	g_file_info_get_creation_date_time     = cc.DlFunc[func(info *GFileInfo) *glib.GDateTime, *glib.GDateTime]{Name: "g_file_info_get_creation_date_time"}
	g_file_info_get_symlink_target         = cc.DlFunc[func(info *GFileInfo) cc.String, cc.String]{Name: "g_file_info_get_symlink_target"}
	g_file_info_get_etag                   = cc.DlFunc[func(info *GFileInfo) cc.String, cc.String]{Name: "g_file_info_get_etag"}
	g_file_info_get_sort_order             = cc.DlFunc[func(info *GFileInfo) int32, int32]{Name: "g_file_info_get_sort_order"}

	g_file_info_set_attribute_mask   = cc.DlFunc[func(info *GFileInfo, mask *GFileAttributeMatcher), cc.Void]{Name: "g_file_info_set_attribute_mask"}
	g_file_info_unset_attribute_mask = cc.DlFunc[func(info *GFileInfo), cc.Void]{Name: "g_file_info_unset_attribute_mask"}

	g_file_info_set_file_type              = cc.DlFunc[func(info *GFileInfo, typ GFileType), cc.Void]{Name: "g_file_info_set_file_type"}
	g_file_info_set_is_hidden              = cc.DlFunc[func(info *GFileInfo, is_hidden bool), cc.Void]{Name: "g_file_info_set_is_hidden"}
	g_file_info_set_is_symlink             = cc.DlFunc[func(info *GFileInfo, is_symlink bool), cc.Void]{Name: "g_file_info_set_is_symlink"}
	g_file_info_set_name                   = cc.DlFunc[func(info *GFileInfo, name cc.String), cc.Void]{Name: "g_file_info_set_name"}
	g_file_info_set_display_name           = cc.DlFunc[func(info *GFileInfo, display_name cc.String), cc.Void]{Name: "g_file_info_set_display_name"}
	g_file_info_set_edit_name              = cc.DlFunc[func(info *GFileInfo, edit_name cc.String), cc.Void]{Name: "g_file_info_set_edit_name"}
	g_file_info_set_icon                   = cc.DlFunc[func(info *GFileInfo, icon *GIcon), cc.Void]{Name: "g_file_info_set_icon"}
	g_file_info_set_symbolic_icon          = cc.DlFunc[func(info *GFileInfo, icon *GIcon), cc.Void]{Name: "g_file_info_set_symbolic_icon"}
	g_file_info_set_content_type           = cc.DlFunc[func(info *GFileInfo, content_type cc.String), cc.Void]{Name: "g_file_info_set_content_type"}
	g_file_info_set_size                   = cc.DlFunc[func(info *GFileInfo, size int64), cc.Void]{Name: "g_file_info_set_size"}
	g_file_info_set_modification_date_time = cc.DlFunc[func(info *GFileInfo, mtime *glib.GDateTime), cc.Void]{Name: "g_file_info_set_modification_date_time"}
	g_file_info_set_access_date_time       = cc.DlFunc[func(info *GFileInfo, atime *glib.GDateTime), cc.Void]{Name: "g_file_info_set_access_date_time"}
	g_file_info_set_creation_date_time     = cc.DlFunc[func(info *GFileInfo, ctime *glib.GDateTime), cc.Void]{Name: "g_file_info_set_creation_date_time"}
	g_file_info_set_symlink_target         = cc.DlFunc[func(info *GFileInfo, symlink_target cc.String), cc.Void]{Name: "g_file_info_set_symlink_target"}
	g_file_info_set_sort_order             = cc.DlFunc[func(info *GFileInfo, sort_order int32), cc.Void]{Name: "g_file_info_set_sort_order"}

	g_file_attribute_matcher_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "g_file_attribute_matcher_get_type"}
	g_file_attribute_matcher_new                 = cc.DlFunc[func(attributes cc.String) *GFileAttributeMatcher, *GFileAttributeMatcher]{Name: "g_file_attribute_matcher_new"}
	g_file_attribute_matcher_ref                 = cc.DlFunc[func(matcher *GFileAttributeMatcher) *GFileAttributeMatcher, *GFileAttributeMatcher]{Name: "g_file_attribute_matcher_ref"}
	g_file_attribute_matcher_unref               = cc.DlFunc[func(matcher *GFileAttributeMatcher), cc.Void]{Name: "g_file_attribute_matcher_unref"}
	g_file_attribute_matcher_subtract            = cc.DlFunc[func(matcher, subtract *GFileAttributeMatcher) *GFileAttributeMatcher, *GFileAttributeMatcher]{Name: "g_file_attribute_matcher_subtract"}
	g_file_attribute_matcher_matches             = cc.DlFunc[func(matcher *GFileAttributeMatcher, attribute cc.String) bool, bool]{Name: "g_file_attribute_matcher_matches"}
	g_file_attribute_matcher_matches_only        = cc.DlFunc[func(matcher *GFileAttributeMatcher, attribute cc.String) bool, bool]{Name: "g_file_attribute_matcher_matches_only"}
	g_file_attribute_matcher_enumerate_namespace = cc.DlFunc[func(matcher *GFileAttributeMatcher, ns cc.String) bool, bool]{Name: "g_file_attribute_matcher_enumerate_namespace"}
	g_file_attribute_matcher_enumerate_next      = cc.DlFunc[func(matcher *GFileAttributeMatcher) cc.String, cc.String]{Name: "g_file_attribute_matcher_enumerate_next"}
	g_file_attribute_matcher_to_string           = cc.DlFunc[func(matcher *GFileAttributeMatcher) cc.String, cc.String]{Name: "g_file_attribute_matcher_to_string"}
	// #endregion

	// #region FileIOStream
	g_file_io_stream_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "g_file_io_stream_get_type"}
	g_file_io_stream_query_info        = cc.DlFunc[func(stream *GFileIOStream, attributes cc.String, cancellable *GCancellable, err **glib.GError) *GFileInfo, *GFileInfo]{Name: "g_file_io_stream_query_info"}
	g_file_io_stream_query_info_async  = cc.DlFunc[func(stream *GFileIOStream, attributes cc.String, io_priority int32, cancellable *GCancellable, callback uptr, user_data uptr), cc.Void]{Name: "g_file_io_stream_query_info_async"}
	g_file_io_stream_query_info_finish = cc.DlFunc[func(stream *GFileIOStream, result *GAsyncResult, err **glib.GError) *GFileInfo, *GFileInfo]{Name: "g_file_io_stream_query_info_finish"}
	g_file_io_stream_get_etag          = cc.DlFunc[func(stream *GFileIOStream) cc.String, cc.String]{Name: "g_file_io_stream_get_etag"}
	// #endregion

	// #region Icon
	g_icon_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "g_icon_get_type"}
	g_icon_hash           = cc.DlFunc[func(icon *GIcon) uint32, uint32]{Name: "g_icon_hash"}
	g_icon_equal          = cc.DlFunc[func(icon1, icon2 *GIcon) bool, bool]{Name: "g_icon_equal"}
	g_icon_to_string      = cc.DlFunc[func(icon *GIcon) cc.String, cc.String]{Name: "g_icon_to_string"}
	g_icon_new_for_string = cc.DlFunc[func(str cc.String, err **glib.GError) *GIcon, *GIcon]{Name: "g_icon_new_for_string"}
	g_icon_serialize      = cc.DlFunc[func(icon *GIcon) *glib.GVariant, *glib.GVariant]{Name: "g_icon_serialize"}
	g_icon_deserialize    = cc.DlFunc[func(value *glib.GVariant) *GIcon, *GIcon]{Name: "g_icon_deserialize"}
	// #endregion

	// #region IOStream
	g_io_stream_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "g_io_stream_get_type"}
	g_io_stream_get_input_stream  = cc.DlFunc[func(stream *GIOStream) *GInputStream, *GInputStream]{Name: "g_io_stream_get_input_stream"}
	g_io_stream_get_output_stream = cc.DlFunc[func(stream *GIOStream) *GOutputStream, *GOutputStream]{Name: "g_io_stream_get_output_stream"}
	g_io_stream_splice_async      = cc.DlFunc[func(stream1, stream2 *GIOStream, flags GIOStreamSpliceFlags, io_priority int32, cancellable *GCancellable, callback uptr, user_data uptr), cc.Void]{Name: "g_io_stream_splice_async"}
	g_io_stream_splice_finish     = cc.DlFunc[func(result *GAsyncResult, err **glib.GError) bool, bool]{Name: "g_io_stream_splice_finish"}
	g_io_stream_close             = cc.DlFunc[func(stream *GIOStream, cancellable *GCancellable, err **glib.GError) bool, bool]{Name: "g_io_stream_close"}
	g_io_stream_close_async       = cc.DlFunc[func(stream *GIOStream, io_priority int32, cancellable *GCancellable, callback uptr, user_data uptr), cc.Void]{Name: "g_io_stream_close_async"}
	g_io_stream_close_finish      = cc.DlFunc[func(stream *GIOStream, result *GAsyncResult, err **glib.GError) bool, bool]{Name: "g_io_stream_close_finish"}
	g_io_stream_is_closed         = cc.DlFunc[func(stream *GIOStream) bool, bool]{Name: "g_io_stream_is_closed"}
	g_io_stream_has_pending       = cc.DlFunc[func(stream *GIOStream) bool, bool]{Name: "g_io_stream_has_pending"}
	g_io_stream_set_pending       = cc.DlFunc[func(stream *GIOStream, err **glib.GError) bool, bool]{Name: "g_io_stream_set_pending"}
	g_io_stream_clear_pending     = cc.DlFunc[func(stream *GIOStream), cc.Void]{Name: "g_io_stream_clear_pending"}
	// #endregion

	// #region ListModel
	g_list_model_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "g_list_model_get_type"}
	g_list_model_get_item_type = cc.DlFunc[func(list uptr) gobject.GType, gobject.GType]{Name: "g_list_model_get_item_type"}
	g_list_model_get_n_items   = cc.DlFunc[func(list uptr) uint32, uint32]{Name: "g_list_model_get_n_items"}
	g_list_model_get_item      = cc.DlFunc[func(list uptr, position uint32) unsafe.Pointer, unsafe.Pointer]{Name: "g_list_model_get_item"}
	g_list_model_get_object    = cc.DlFunc[func(list uptr, position uint32) *gobject.GObject, *gobject.GObject]{Name: "g_list_model_get_object"}
	g_list_model_items_changed = cc.DlFunc[func(list uptr, position, removed, added uint32), cc.Void]{Name: "g_list_model_items_changed"}
	// #endregion

	// #region ListStore
	g_list_store_get_type                  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "g_list_store_get_type"}
	g_list_store_new                       = cc.DlFunc[func(itemType gobject.GType) uptr, uptr]{Name: "g_list_store_new"}
	g_list_store_insert                    = cc.DlFunc[func(store uptr, position uint32, item uptr), cc.Void]{Name: "g_list_store_insert"}
	g_list_store_insert_sorted             = cc.DlFunc[func(store uptr, item uptr, compareFunc uptr, userData uptr) uint32, uint32]{Name: "g_list_store_insert_sorted"}
	g_list_store_sort                      = cc.DlFunc[func(store uptr, compareFunc uptr, userData uptr), cc.Void]{Name: "g_list_store_sort"}
	g_list_store_append                    = cc.DlFunc[func(store uptr, item uptr), cc.Void]{Name: "g_list_store_append"}
	g_list_store_remove                    = cc.DlFunc[func(store uptr, position uint32), cc.Void]{Name: "g_list_store_remove"}
	g_list_store_remove_all                = cc.DlFunc[func(store uptr), cc.Void]{Name: "g_list_store_remove_all"}
	g_list_store_splice                    = cc.DlFunc[func(store uptr, position, nRemovals uint32, additions uptr, nAdditions uint32), cc.Void]{Name: "g_list_store_splice"}
	g_list_store_find                      = cc.DlFunc[func(store uptr, item uptr, position *uint32) bool, bool]{Name: "g_list_store_find"}
	g_list_store_find_with_equal_func      = cc.DlFunc[func(store uptr, item uptr, equalFunc uptr, position *uint32) bool, bool]{Name: "g_list_store_find_with_equal_func"}
	g_list_store_find_with_equal_func_full = cc.DlFunc[func(store uptr, item uptr, equalFunc uptr, userData uptr, position *uint32) bool, bool]{Name: "g_list_store_find_with_equal_func_full"}
	// #endregion

)
