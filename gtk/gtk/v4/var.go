package gtk

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cairo"
	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/glib/gio"
	"github.com/jinzhongmin/gg/glib/glib"
	"github.com/jinzhongmin/gg/glib/gobject"
	"github.com/jinzhongmin/gg/graphene"
	"github.com/jinzhongmin/gg/gtk/gdk"
	"github.com/jinzhongmin/gg/gtk/gsk"
	"github.com/jinzhongmin/gg/pango"
)

type uptr = unsafe.Pointer
type iptr = uintptr

// func vcbu(fn interface{}) UPtr  { return cc.Cbk(fn) }

func anyptr(a interface{}) uptr { return (*(*[2]uptr)(uptr(&a)))[1] }
func carry[T any](arry []T) *T {
	if len(arry) == 0 {
		return nil
	}
	return &arry[0]
}
func toGobj(a interface{}) *gobject.GObject           { return (*gobject.GObject)(anyptr(a)) }
func gobjGet(a gobject.GObjectIface) *gobject.GObject { return gobject.GetGObjectIface(a) }

type gobjCore = gobject.GObjectCore
type GobjIface = gobject.GObjectIface

// func init() { cc.Open("libgtk-4*") }

var (
	// #region AboutDialog
	gtk_about_dialog_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_about_dialog_get_type"}
	gtk_about_dialog_new                    = cc.DlFunc[func() *AboutDialog, *AboutDialog]{Name: "gtk_about_dialog_new"}
	gtk_show_about_dialog                   = cc.DlFunc[func(parent *Window, firstPropertyName cc.String, args ...interface{}), cc.Void]{Name: "gtk_show_about_dialog", Va: true}
	gtk_about_dialog_get_program_name       = cc.DlFunc[func(about *AboutDialog) cc.String, cc.String]{Name: "gtk_about_dialog_get_program_name"}
	gtk_about_dialog_set_program_name       = cc.DlFunc[func(about *AboutDialog, name cc.String), cc.Void]{Name: "gtk_about_dialog_set_program_name"}
	gtk_about_dialog_get_version            = cc.DlFunc[func(about *AboutDialog) cc.String, cc.String]{Name: "gtk_about_dialog_get_version"}
	gtk_about_dialog_set_version            = cc.DlFunc[func(about *AboutDialog, version cc.String), cc.Void]{Name: "gtk_about_dialog_set_version"}
	gtk_about_dialog_get_copyright          = cc.DlFunc[func(about *AboutDialog) cc.String, cc.String]{Name: "gtk_about_dialog_get_copyright"}
	gtk_about_dialog_set_copyright          = cc.DlFunc[func(about *AboutDialog, copyright cc.String), cc.Void]{Name: "gtk_about_dialog_set_copyright"}
	gtk_about_dialog_get_comments           = cc.DlFunc[func(about *AboutDialog) cc.String, cc.String]{Name: "gtk_about_dialog_get_comments"}
	gtk_about_dialog_set_comments           = cc.DlFunc[func(about *AboutDialog, comments cc.String), cc.Void]{Name: "gtk_about_dialog_set_comments"}
	gtk_about_dialog_get_license            = cc.DlFunc[func(about *AboutDialog) cc.String, cc.String]{Name: "gtk_about_dialog_get_license"}
	gtk_about_dialog_set_license            = cc.DlFunc[func(about *AboutDialog, license cc.String), cc.Void]{Name: "gtk_about_dialog_set_license"}
	gtk_about_dialog_set_license_type       = cc.DlFunc[func(about *AboutDialog, licenseType License), cc.Void]{Name: "gtk_about_dialog_set_license_type"}
	gtk_about_dialog_get_license_type       = cc.DlFunc[func(about *AboutDialog) License, License]{Name: "gtk_about_dialog_get_license_type"}
	gtk_about_dialog_get_wrap_license       = cc.DlFunc[func(about *AboutDialog) int32, int32]{Name: "gtk_about_dialog_get_wrap_license"}
	gtk_about_dialog_set_wrap_license       = cc.DlFunc[func(about *AboutDialog, wrapLicense int32), cc.Void]{Name: "gtk_about_dialog_set_wrap_license"}
	gtk_about_dialog_get_system_information = cc.DlFunc[func(about *AboutDialog) cc.String, cc.String]{Name: "gtk_about_dialog_get_system_information"}
	gtk_about_dialog_set_system_information = cc.DlFunc[func(about *AboutDialog, systemInformation cc.String), cc.Void]{Name: "gtk_about_dialog_set_system_information"}
	gtk_about_dialog_get_website            = cc.DlFunc[func(about *AboutDialog) cc.String, cc.String]{Name: "gtk_about_dialog_get_website"}
	gtk_about_dialog_set_website            = cc.DlFunc[func(about *AboutDialog, website cc.String), cc.Void]{Name: "gtk_about_dialog_set_website"}
	gtk_about_dialog_get_website_label      = cc.DlFunc[func(about *AboutDialog) cc.String, cc.String]{Name: "gtk_about_dialog_get_website_label"}
	gtk_about_dialog_set_website_label      = cc.DlFunc[func(about *AboutDialog, websiteLabel cc.String), cc.Void]{Name: "gtk_about_dialog_set_website_label"}
	gtk_about_dialog_get_authors            = cc.DlFunc[func(about *AboutDialog) cc.Strings, cc.Strings]{Name: "gtk_about_dialog_get_authors"}
	gtk_about_dialog_set_authors            = cc.DlFunc[func(about *AboutDialog, authors cc.Strings), cc.Void]{Name: "gtk_about_dialog_set_authors"}
	gtk_about_dialog_get_documenters        = cc.DlFunc[func(about *AboutDialog) cc.Strings, cc.Strings]{Name: "gtk_about_dialog_get_documenters"}
	gtk_about_dialog_set_documenters        = cc.DlFunc[func(about *AboutDialog, documenters cc.Strings), cc.Void]{Name: "gtk_about_dialog_set_documenters"}
	gtk_about_dialog_get_artists            = cc.DlFunc[func(about *AboutDialog) cc.Strings, cc.Strings]{Name: "gtk_about_dialog_get_artists"}
	gtk_about_dialog_set_artists            = cc.DlFunc[func(about *AboutDialog, artists cc.Strings), cc.Void]{Name: "gtk_about_dialog_set_artists"}
	gtk_about_dialog_get_translator_credits = cc.DlFunc[func(about *AboutDialog) cc.String, cc.String]{Name: "gtk_about_dialog_get_translator_credits"}
	gtk_about_dialog_set_translator_credits = cc.DlFunc[func(about *AboutDialog, translatorCredits cc.String), cc.Void]{Name: "gtk_about_dialog_set_translator_credits"}
	gtk_about_dialog_get_logo               = cc.DlFunc[func(about *AboutDialog) *gdk.Paintable, *gdk.Paintable]{Name: "gtk_about_dialog_get_logo"}
	gtk_about_dialog_set_logo               = cc.DlFunc[func(about *AboutDialog, logo *gdk.Paintable), cc.Void]{Name: "gtk_about_dialog_set_logo"}
	gtk_about_dialog_get_logo_icon_name     = cc.DlFunc[func(about *AboutDialog) cc.String, cc.String]{Name: "gtk_about_dialog_get_logo_icon_name"}
	gtk_about_dialog_set_logo_icon_name     = cc.DlFunc[func(about *AboutDialog, iconName cc.String), cc.Void]{Name: "gtk_about_dialog_set_logo_icon_name"}
	gtk_about_dialog_add_credit_section     = cc.DlFunc[func(about *AboutDialog, sectionName cc.String, people cc.Strings), cc.Void]{Name: "gtk_about_dialog_add_credit_section"}
	// #endregion

	// #region Accelerator
	gtk_accelerator_valid                  = cc.DlFunc[func(keyval uint32, modifiers gdk.ModifierType) int32, int32]{Name: "gtk_accelerator_valid"}
	gtk_accelerator_parse                  = cc.DlFunc[func(accelerator cc.String, acceleratorKey *uint32, acceleratorMods *gdk.ModifierType) int32, int32]{Name: "gtk_accelerator_parse"}
	gtk_accelerator_parse_with_keycode     = cc.DlFunc[func(accelerator cc.String, display *gdk.Display, acceleratorKey *uint32, acceleratorCodes **uint32, acceleratorMods *gdk.ModifierType) int32, int32]{Name: "gtk_accelerator_parse_with_keycode"}
	gtk_accelerator_name                   = cc.DlFunc[func(acceleratorKey uint32, acceleratorMods gdk.ModifierType) cc.String, cc.String]{Name: "gtk_accelerator_name"}
	gtk_accelerator_name_with_keycode      = cc.DlFunc[func(display *gdk.Display, acceleratorKey uint32, keycode uint32, acceleratorMods gdk.ModifierType) cc.String, cc.String]{Name: "gtk_accelerator_name_with_keycode"}
	gtk_accelerator_get_label              = cc.DlFunc[func(acceleratorKey uint32, acceleratorMods gdk.ModifierType) cc.String, cc.String]{Name: "gtk_accelerator_get_label"}
	gtk_accelerator_get_label_with_keycode = cc.DlFunc[func(display *gdk.Display, acceleratorKey uint32, keycode uint32, acceleratorMods gdk.ModifierType) cc.String, cc.String]{Name: "gtk_accelerator_get_label_with_keycode"}
	gtk_accelerator_get_default_mod_mask   = cc.DlFunc[func() gdk.ModifierType, gdk.ModifierType]{Name: "gtk_accelerator_get_default_mod_mask"}
	// #endregion

	// #region Accessible
	gtk_accessible_get_at_context                 = cc.DlFunc[func(self *Accessible) *ATContext, *ATContext]{Name: "gtk_accessible_get_at_context"}
	gtk_accessible_get_platform_state             = cc.DlFunc[func(self *Accessible, state AccessiblePlatformState) int32, int32]{Name: "gtk_accessible_get_platform_state"}
	gtk_accessible_get_accessible_parent          = cc.DlFunc[func(self *Accessible) *Accessible, *Accessible]{Name: "gtk_accessible_get_accessible_parent"}
	gtk_accessible_set_accessible_parent          = cc.DlFunc[func(self *Accessible, parent *Accessible, nextSibling *Accessible), cc.Void]{Name: "gtk_accessible_set_accessible_parent"}
	gtk_accessible_get_first_accessible_child     = cc.DlFunc[func(self *Accessible) *Accessible, *Accessible]{Name: "gtk_accessible_get_first_accessible_child"}
	gtk_accessible_get_next_accessible_sibling    = cc.DlFunc[func(self *Accessible) *Accessible, *Accessible]{Name: "gtk_accessible_get_next_accessible_sibling"}
	gtk_accessible_update_next_accessible_sibling = cc.DlFunc[func(self *Accessible, newSibling *Accessible), cc.Void]{Name: "gtk_accessible_update_next_accessible_sibling"}
	gtk_accessible_get_bounds                     = cc.DlFunc[func(self *Accessible, x, y, width, height *int32) int32, int32]{Name: "gtk_accessible_get_bounds"}
	gtk_accessible_get_accessible_role            = cc.DlFunc[func(a *Accessible) AccessibleRole, AccessibleRole]{Name: "gtk_accessible_get_accessible_role"}
	gtk_accessible_update_state                   = cc.DlFunc[func(a *Accessible, firstState AccessibleState, args ...interface{}), cc.Void]{Name: "gtk_accessible_update_state", Va: true}
	gtk_accessible_update_property                = cc.DlFunc[func(a *Accessible, firstProperty AccessibleProperty, args ...interface{}), cc.Void]{Name: "gtk_accessible_update_property", Va: true}
	gtk_accessible_update_relation                = cc.DlFunc[func(a *Accessible, firstRelation AccessibleRelation, args ...interface{}), cc.Void]{Name: "gtk_accessible_update_relation", Va: true}
	gtk_accessible_update_state_value             = cc.DlFunc[func(a *Accessible, nStates int32, states *AccessibleState, values *gobject.GValue), cc.Void]{Name: "gtk_accessible_update_state_value"}
	gtk_accessible_update_property_value          = cc.DlFunc[func(a *Accessible, nProperties int32, properties *AccessibleProperty, values *gobject.GValue), cc.Void]{Name: "gtk_accessible_update_property_value"}
	gtk_accessible_update_relation_value          = cc.DlFunc[func(a *Accessible, nRelations int32, relations *AccessibleRelation, values *gobject.GValue), cc.Void]{Name: "gtk_accessible_update_relation_value"}
	gtk_accessible_reset_state                    = cc.DlFunc[func(a *Accessible, state AccessibleState), cc.Void]{Name: "gtk_accessible_reset_state"}
	gtk_accessible_reset_property                 = cc.DlFunc[func(a *Accessible, property AccessibleProperty), cc.Void]{Name: "gtk_accessible_reset_property"}
	gtk_accessible_reset_relation                 = cc.DlFunc[func(a *Accessible, relation AccessibleRelation), cc.Void]{Name: "gtk_accessible_reset_relation"}
	gtk_accessible_state_init_value               = cc.DlFunc[func(state AccessibleState, value *gobject.GValue), cc.Void]{Name: "gtk_accessible_state_init_value"}
	gtk_accessible_property_init_value            = cc.DlFunc[func(property AccessibleProperty, value *gobject.GValue), cc.Void]{Name: "gtk_accessible_property_init_value"}
	gtk_accessible_relation_init_value            = cc.DlFunc[func(relation AccessibleRelation, value *gobject.GValue), cc.Void]{Name: "gtk_accessible_relation_init_value"}
	gtk_accessible_list_get_type                  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_accessible_list_get_type"}
	gtk_accessible_list_get_objects               = cc.DlFunc[func(accessibleList *AccessibleList) *glib.GList[Accessible], *glib.GList[Accessible]]{Name: "gtk_accessible_list_get_objects"}
	gtk_accessible_list_new_from_list             = cc.DlFunc[func(list *glib.GList[Accessible]) *AccessibleList, *AccessibleList]{Name: "gtk_accessible_list_new_from_list"}
	gtk_accessible_list_new_from_array            = cc.DlFunc[func(accessibles **Accessible, nAccessibles uint64) *AccessibleList, *AccessibleList]{Name: "gtk_accessible_list_new_from_array"}
	gtk_accessible_announce                       = cc.DlFunc[func(a *Accessible, message cc.String, priority AccessibleAnnouncementPriority), cc.Void]{Name: "gtk_accessible_announce"}
	gtk_accessible_update_platform_state          = cc.DlFunc[func(a *Accessible, state AccessiblePlatformState), cc.Void]{Name: "gtk_accessible_update_platform_state"}
	// #endregion

	// #region AccessibleRange
	gtk_accessible_range_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_accessible_range_get_type"}
	gtk_accessible_range_set_current_value = cc.DlFunc[func(*AccessibleRange, float64) int32, int32]{Name: "gtk_accessible_range_set_current_value"}
	// #endregion

	// #region AccessibleTextRange
	gtk_accessible_text_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_accessible_text_get_type"}
	gtk_accessible_text_update_caret_position  = cc.DlFunc[func(*AccessibleText), cc.Void]{Name: "gtk_accessible_text_update_caret_position"}
	gtk_accessible_text_update_selection_bound = cc.DlFunc[func(*AccessibleText), cc.Void]{Name: "gtk_accessible_text_update_selection_bound"}
	gtk_accessible_text_update_contents        = cc.DlFunc[func(*AccessibleText, AccessibleTextContentChange, uint32, uint32), cc.Void]{Name: "gtk_accessible_text_update_contents"}
	gtk_accessible_text_get_contents           = cc.DlFunc[func(*AccessibleText, uint32, uint32) *glib.GBytes, *glib.GBytes]{Name: "gtk_accessible_text_get_contents"}
	gtk_accessible_text_get_contents_at        = cc.DlFunc[func(*AccessibleText, uint32, AccessibleTextGranularity, *uint32, *uint32) *glib.GBytes, *glib.GBytes]{Name: "gtk_accessible_text_get_contents_at"}
	gtk_accessible_text_get_caret_position     = cc.DlFunc[func(*AccessibleText) uint32, uint32]{Name: "gtk_accessible_text_get_caret_position"}
	gtk_accessible_text_get_selection          = cc.DlFunc[func(*AccessibleText, *uint64, **AccessibleTextRange) int32, int32]{Name: "gtk_accessible_text_get_selection"}
	gtk_accessible_text_get_attributes         = cc.DlFunc[func(*AccessibleText, uint32, *uint64, **AccessibleTextRange, *cc.Strings, *cc.Strings) int32, int32]{Name: "gtk_accessible_text_get_attributes"}
	gtk_accessible_text_get_default_attributes = cc.DlFunc[func(*AccessibleText, *cc.Strings, *cc.Strings), cc.Void]{Name: "gtk_accessible_text_get_default_attributes"}
	gtk_accessible_text_get_extents            = cc.DlFunc[func(*AccessibleText, uint32, uint32, *graphene.Rect) int32, int32]{Name: "gtk_accessible_text_get_extents"}
	gtk_accessible_text_get_offset             = cc.DlFunc[func(*AccessibleText, *graphene.Point, *uint32) int32, int32]{Name: "gtk_accessible_text_get_offset"}
	// #endregion

	// #region Actionable
	gtk_actionable_get_type                 = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_actionable_get_type"}
	gtk_actionable_get_action_name          = cc.DlFunc[func(*Actionable) cc.String, cc.String]{Name: "gtk_actionable_get_action_name"}
	gtk_actionable_set_action_name          = cc.DlFunc[func(*Actionable, cc.String), cc.Void]{Name: "gtk_actionable_set_action_name"}
	gtk_actionable_get_action_target_value  = cc.DlFunc[func(*Actionable) *glib.GVariant, *glib.GVariant]{Name: "gtk_actionable_get_action_target_value"}
	gtk_actionable_set_action_target_value  = cc.DlFunc[func(*Actionable, *glib.GVariant), cc.Void]{Name: "gtk_actionable_set_action_target_value"}
	gtk_actionable_set_action_target        = cc.DlFunc[func(*Actionable, cc.String, ...interface{}), cc.Void]{Name: "gtk_actionable_set_action_target", Va: true}
	gtk_actionable_set_detailed_action_name = cc.DlFunc[func(*Actionable, cc.String), cc.Void]{Name: "gtk_actionable_set_detailed_action_name"}
	// #endregion

	// #region ActionBar
	gtk_action_bar_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_action_bar_get_type"}
	gtk_action_bar_new               = cc.DlFunc[func() *ActionBar, *ActionBar]{Name: "gtk_action_bar_new"}
	gtk_action_bar_get_center_widget = cc.DlFunc[func(*ActionBar) *Widget, *Widget]{Name: "gtk_action_bar_get_center_widget"}
	gtk_action_bar_set_center_widget = cc.DlFunc[func(*ActionBar, *Widget), cc.Void]{Name: "gtk_action_bar_set_center_widget"}
	gtk_action_bar_pack_start        = cc.DlFunc[func(*ActionBar, *Widget), cc.Void]{Name: "gtk_action_bar_pack_start"}
	gtk_action_bar_pack_end          = cc.DlFunc[func(*ActionBar, *Widget), cc.Void]{Name: "gtk_action_bar_pack_end"}
	gtk_action_bar_remove            = cc.DlFunc[func(*ActionBar, *Widget), cc.Void]{Name: "gtk_action_bar_remove"}
	gtk_action_bar_set_revealed      = cc.DlFunc[func(*ActionBar, int32), cc.Void]{Name: "gtk_action_bar_set_revealed"}
	gtk_action_bar_get_revealed      = cc.DlFunc[func(*ActionBar) int32, int32]{Name: "gtk_action_bar_get_revealed"}
	// #endregion

	// #region Adjustment
	gtk_adjustment_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_adjustment_get_type"}
	gtk_adjustment_new                   = cc.DlFunc[func(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) *Adjustment, *Adjustment]{Name: "gtk_adjustment_new"}
	gtk_adjustment_clamp_page            = cc.DlFunc[func(adjustment *Adjustment, lower, upper float64), cc.Void]{Name: "gtk_adjustment_clamp_page"}
	gtk_adjustment_get_value             = cc.DlFunc[func(adjustment *Adjustment) float64, float64]{Name: "gtk_adjustment_get_value"}
	gtk_adjustment_set_value             = cc.DlFunc[func(adjustment *Adjustment, value float64), cc.Void]{Name: "gtk_adjustment_set_value"}
	gtk_adjustment_get_lower             = cc.DlFunc[func(adjustment *Adjustment) float64, float64]{Name: "gtk_adjustment_get_lower"}
	gtk_adjustment_set_lower             = cc.DlFunc[func(adjustment *Adjustment, lower float64), cc.Void]{Name: "gtk_adjustment_set_lower"}
	gtk_adjustment_get_upper             = cc.DlFunc[func(adjustment *Adjustment) float64, float64]{Name: "gtk_adjustment_get_upper"}
	gtk_adjustment_set_upper             = cc.DlFunc[func(adjustment *Adjustment, upper float64), cc.Void]{Name: "gtk_adjustment_set_upper"}
	gtk_adjustment_get_step_increment    = cc.DlFunc[func(adjustment *Adjustment) float64, float64]{Name: "gtk_adjustment_get_step_increment"}
	gtk_adjustment_set_step_increment    = cc.DlFunc[func(adjustment *Adjustment, stepIncrement float64), cc.Void]{Name: "gtk_adjustment_set_step_increment"}
	gtk_adjustment_get_page_increment    = cc.DlFunc[func(adjustment *Adjustment) float64, float64]{Name: "gtk_adjustment_get_page_increment"}
	gtk_adjustment_set_page_increment    = cc.DlFunc[func(adjustment *Adjustment, pageIncrement float64), cc.Void]{Name: "gtk_adjustment_set_page_increment"}
	gtk_adjustment_get_page_size         = cc.DlFunc[func(adjustment *Adjustment) float64, float64]{Name: "gtk_adjustment_get_page_size"}
	gtk_adjustment_set_page_size         = cc.DlFunc[func(adjustment *Adjustment, pageSize float64), cc.Void]{Name: "gtk_adjustment_set_page_size"}
	gtk_adjustment_configure             = cc.DlFunc[func(adjustment *Adjustment, value, lower, upper, stepIncrement, pageIncrement, pageSize float64), cc.Void]{Name: "gtk_adjustment_configure"}
	gtk_adjustment_get_minimum_increment = cc.DlFunc[func(adjustment *Adjustment) float64, float64]{Name: "gtk_adjustment_get_minimum_increment"}
	// #endregion

	// #region AlertDialog
	gtk_alert_dialog_new                = cc.DlFunc[func(format cc.String, args ...interface{}) *AlertDialog, *AlertDialog]{Name: "gtk_alert_dialog_new", Va: true}
	gtk_alert_dialog_get_modal          = cc.DlFunc[func(self *AlertDialog) int32, int32]{Name: "gtk_alert_dialog_get_modal"}
	gtk_alert_dialog_set_modal          = cc.DlFunc[func(self *AlertDialog, modal int32), cc.Void]{Name: "gtk_alert_dialog_set_modal"}
	gtk_alert_dialog_get_message        = cc.DlFunc[func(self *AlertDialog) cc.String, cc.String]{Name: "gtk_alert_dialog_get_message"}
	gtk_alert_dialog_set_message        = cc.DlFunc[func(self *AlertDialog, message cc.String), cc.Void]{Name: "gtk_alert_dialog_set_message"}
	gtk_alert_dialog_get_detail         = cc.DlFunc[func(self *AlertDialog) cc.String, cc.String]{Name: "gtk_alert_dialog_get_detail"}
	gtk_alert_dialog_set_detail         = cc.DlFunc[func(self *AlertDialog, detail cc.String), cc.Void]{Name: "gtk_alert_dialog_set_detail"}
	gtk_alert_dialog_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_alert_dialog_get_type"}
	gtk_alert_dialog_get_buttons        = cc.DlFunc[func(self *AlertDialog) cc.Strings, cc.Strings]{Name: "gtk_alert_dialog_get_buttons"}
	gtk_alert_dialog_set_buttons        = cc.DlFunc[func(self *AlertDialog, labels cc.Strings), cc.Void]{Name: "gtk_alert_dialog_set_buttons"}
	gtk_alert_dialog_get_cancel_button  = cc.DlFunc[func(self *AlertDialog) int32, int32]{Name: "gtk_alert_dialog_get_cancel_button"}
	gtk_alert_dialog_set_cancel_button  = cc.DlFunc[func(self *AlertDialog, button int32), cc.Void]{Name: "gtk_alert_dialog_set_cancel_button"}
	gtk_alert_dialog_get_default_button = cc.DlFunc[func(self *AlertDialog) int32, int32]{Name: "gtk_alert_dialog_get_default_button"}
	gtk_alert_dialog_set_default_button = cc.DlFunc[func(self *AlertDialog, button int32), cc.Void]{Name: "gtk_alert_dialog_set_default_button"}
	gtk_alert_dialog_choose             = cc.DlFunc[func(self *AlertDialog, parent *Window, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_alert_dialog_choose"}
	gtk_alert_dialog_choose_finish      = cc.DlFunc[func(self *AlertDialog, result *gio.GAsyncResult, err **glib.GError) int32, int32]{Name: "gtk_alert_dialog_choose_finish"}
	gtk_alert_dialog_show               = cc.DlFunc[func(self *AlertDialog, parent *Window), cc.Void]{Name: "gtk_alert_dialog_show"}
	// #endregion

	// #region Application
	gtk_application_get_type                 = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_application_get_type"}
	gtk_application_new                      = cc.DlFunc[func(applicationID cc.String, flags gio.GApplicationFlags) *Application, *Application]{Name: "gtk_application_new"}
	gtk_application_add_window               = cc.DlFunc[func(application *Application, window *Window), cc.Void]{Name: "gtk_application_add_window"}
	gtk_application_remove_window            = cc.DlFunc[func(application *Application, window *Window), cc.Void]{Name: "gtk_application_remove_window"}
	gtk_application_get_windows              = cc.DlFunc[func(application *Application) *glib.GList[Window], *glib.GList[Window]]{Name: "gtk_application_get_windows"}
	gtk_application_get_menubar              = cc.DlFunc[func(application *Application) *gio.GMenuModel, *gio.GMenuModel]{Name: "gtk_application_get_menubar"}
	gtk_application_set_menubar              = cc.DlFunc[func(application *Application, menubar *gio.GMenuModel), cc.Void]{Name: "gtk_application_set_menubar"}
	gtk_application_inhibit                  = cc.DlFunc[func(application *Application, window *Window, flags ApplicationInhibitFlags, reason cc.String) uint32, uint32]{Name: "gtk_application_inhibit"}
	gtk_application_uninhibit                = cc.DlFunc[func(application *Application, cookie uint32), cc.Void]{Name: "gtk_application_uninhibit"}
	gtk_application_get_window_by_id         = cc.DlFunc[func(application *Application, id uint32) *Window, *Window]{Name: "gtk_application_get_window_by_id"}
	gtk_application_get_active_window        = cc.DlFunc[func(application *Application) *Window, *Window]{Name: "gtk_application_get_active_window"}
	gtk_application_list_action_descriptions = cc.DlFunc[func(application *Application) cc.Strings, cc.Strings]{Name: "gtk_application_list_action_descriptions"}
	gtk_application_get_accels_for_action    = cc.DlFunc[func(application *Application, detailedActionName cc.String) cc.Strings, cc.Strings]{Name: "gtk_application_get_accels_for_action"}
	gtk_application_get_actions_for_accel    = cc.DlFunc[func(application *Application, accel cc.String) cc.Strings, cc.Strings]{Name: "gtk_application_get_actions_for_accel"}
	gtk_application_set_accels_for_action    = cc.DlFunc[func(application *Application, detailedActionName cc.String, accels cc.Strings), cc.Void]{Name: "gtk_application_set_accels_for_action"}
	gtk_application_get_menu_by_id           = cc.DlFunc[func(application *Application, id cc.String) *gio.GMenu, *gio.GMenu]{Name: "gtk_application_get_menu_by_id"}
	// #endregion

	// #region ApplicationWindow
	gtk_application_window_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_application_window_get_type"}
	gtk_application_window_new              = cc.DlFunc[func(application *Application) *ApplicationWindow, *ApplicationWindow]{Name: "gtk_application_window_new"}
	gtk_application_window_set_show_menubar = cc.DlFunc[func(window *ApplicationWindow, show int32), cc.Void]{Name: "gtk_application_window_set_show_menubar"}
	gtk_application_window_get_show_menubar = cc.DlFunc[func(window *ApplicationWindow) int32, int32]{Name: "gtk_application_window_get_show_menubar"}
	gtk_application_window_get_id           = cc.DlFunc[func(window *ApplicationWindow) uint32, uint32]{Name: "gtk_application_window_get_id"}
	// #endregion

	// #region AspectFrame
	gtk_aspect_frame_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_aspect_frame_get_type"}
	gtk_aspect_frame_new            = cc.DlFunc[func(xalign, yalign, ratio float32, obeyChild int32) *AspectFrame, *AspectFrame]{Name: "gtk_aspect_frame_new"}
	gtk_aspect_frame_set_xalign     = cc.DlFunc[func(self *AspectFrame, xalign float32), cc.Void]{Name: "gtk_aspect_frame_set_xalign"}
	gtk_aspect_frame_get_xalign     = cc.DlFunc[func(self *AspectFrame) float32, float32]{Name: "gtk_aspect_frame_get_xalign"}
	gtk_aspect_frame_set_yalign     = cc.DlFunc[func(self *AspectFrame, yalign float32), cc.Void]{Name: "gtk_aspect_frame_set_yalign"}
	gtk_aspect_frame_get_yalign     = cc.DlFunc[func(self *AspectFrame) float32, float32]{Name: "gtk_aspect_frame_get_yalign"}
	gtk_aspect_frame_set_ratio      = cc.DlFunc[func(self *AspectFrame, ratio float32), cc.Void]{Name: "gtk_aspect_frame_set_ratio"}
	gtk_aspect_frame_get_ratio      = cc.DlFunc[func(self *AspectFrame) float32, float32]{Name: "gtk_aspect_frame_get_ratio"}
	gtk_aspect_frame_set_obey_child = cc.DlFunc[func(self *AspectFrame, obeyChild int32), cc.Void]{Name: "gtk_aspect_frame_set_obey_child"}
	gtk_aspect_frame_get_obey_child = cc.DlFunc[func(self *AspectFrame) int32, int32]{Name: "gtk_aspect_frame_get_obey_child"}
	gtk_aspect_frame_set_child      = cc.DlFunc[func(self *AspectFrame, child *Widget), cc.Void]{Name: "gtk_aspect_frame_set_child"}
	gtk_aspect_frame_get_child      = cc.DlFunc[func(self *AspectFrame) *Widget, *Widget]{Name: "gtk_aspect_frame_get_child"}
	// #endregion

	// #region ATContext
	gtk_at_context_get_accessible      = cc.DlFunc[func(self *ATContext) *Accessible, *Accessible]{Name: "gtk_at_context_get_accessible"}
	gtk_at_context_get_accessible_role = cc.DlFunc[func(self *ATContext) AccessibleRole, AccessibleRole]{Name: "gtk_at_context_get_accessible_role"}
	gtk_at_context_create              = cc.DlFunc[func(accessibleRole AccessibleRole, accessible *Accessible, display *gdk.Display) *ATContext, *ATContext]{Name: "gtk_at_context_create"}
	// #endregion

	// #region BinLayout
	gtk_bin_layout_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_bin_layout_get_type"}
	gtk_bin_layout_new      = cc.DlFunc[func() *BinLayout, *BinLayout]{Name: "gtk_bin_layout_new"}
	// #endregion

	// #region Bitset
	gtk_bitset_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_bitset_get_type"}
	gtk_bitset_ref                 = cc.DlFunc[func(self *Bitset) *Bitset, *Bitset]{Name: "gtk_bitset_ref"}
	gtk_bitset_unref               = cc.DlFunc[func(self *Bitset), cc.Void]{Name: "gtk_bitset_unref"}
	gtk_bitset_contains            = cc.DlFunc[func(self *Bitset, value uint32) int32, int32]{Name: "gtk_bitset_contains"}
	gtk_bitset_is_empty            = cc.DlFunc[func(self *Bitset) int32, int32]{Name: "gtk_bitset_is_empty"}
	gtk_bitset_equals              = cc.DlFunc[func(self *Bitset, other *Bitset) int32, int32]{Name: "gtk_bitset_equals"}
	gtk_bitset_get_size            = cc.DlFunc[func(self *Bitset) uint64, uint64]{Name: "gtk_bitset_get_size"}
	gtk_bitset_get_size_in_range   = cc.DlFunc[func(self *Bitset, first, last uint32) uint64, uint64]{Name: "gtk_bitset_get_size_in_range"}
	gtk_bitset_get_nth             = cc.DlFunc[func(self *Bitset, nth uint32) uint32, uint32]{Name: "gtk_bitset_get_nth"}
	gtk_bitset_get_minimum         = cc.DlFunc[func(self *Bitset) uint32, uint32]{Name: "gtk_bitset_get_minimum"}
	gtk_bitset_get_maximum         = cc.DlFunc[func(self *Bitset) uint32, uint32]{Name: "gtk_bitset_get_maximum"}
	gtk_bitset_new_empty           = cc.DlFunc[func() *Bitset, *Bitset]{Name: "gtk_bitset_new_empty"}
	gtk_bitset_copy                = cc.DlFunc[func(self *Bitset) *Bitset, *Bitset]{Name: "gtk_bitset_copy"}
	gtk_bitset_new_range           = cc.DlFunc[func(start, nItems uint32) *Bitset, *Bitset]{Name: "gtk_bitset_new_range"}
	gtk_bitset_remove_all          = cc.DlFunc[func(self *Bitset), cc.Void]{Name: "gtk_bitset_remove_all"}
	gtk_bitset_add                 = cc.DlFunc[func(self *Bitset, value uint32) int32, int32]{Name: "gtk_bitset_add"}
	gtk_bitset_remove              = cc.DlFunc[func(self *Bitset, value uint32) int32, int32]{Name: "gtk_bitset_remove"}
	gtk_bitset_add_range           = cc.DlFunc[func(self *Bitset, start, nItems uint32), cc.Void]{Name: "gtk_bitset_add_range"}
	gtk_bitset_remove_range        = cc.DlFunc[func(self *Bitset, start, nItems uint32), cc.Void]{Name: "gtk_bitset_remove_range"}
	gtk_bitset_add_range_closed    = cc.DlFunc[func(self *Bitset, first, last uint32), cc.Void]{Name: "gtk_bitset_add_range_closed"}
	gtk_bitset_remove_range_closed = cc.DlFunc[func(self *Bitset, first, last uint32), cc.Void]{Name: "gtk_bitset_remove_range_closed"}
	gtk_bitset_add_rectangle       = cc.DlFunc[func(self *Bitset, start, width, height, stride uint32), cc.Void]{Name: "gtk_bitset_add_rectangle"}
	gtk_bitset_remove_rectangle    = cc.DlFunc[func(self *Bitset, start, width, height, stride uint32), cc.Void]{Name: "gtk_bitset_remove_rectangle"}
	gtk_bitset_union               = cc.DlFunc[func(self, other *Bitset), cc.Void]{Name: "gtk_bitset_union"}
	gtk_bitset_intersect           = cc.DlFunc[func(self, other *Bitset), cc.Void]{Name: "gtk_bitset_intersect"}
	gtk_bitset_subtract            = cc.DlFunc[func(self, other *Bitset), cc.Void]{Name: "gtk_bitset_subtract"}
	gtk_bitset_difference          = cc.DlFunc[func(self, other *Bitset), cc.Void]{Name: "gtk_bitset_difference"}
	gtk_bitset_shift_left          = cc.DlFunc[func(self *Bitset, amount uint32), cc.Void]{Name: "gtk_bitset_shift_left"}
	gtk_bitset_shift_right         = cc.DlFunc[func(self *Bitset, amount uint32), cc.Void]{Name: "gtk_bitset_shift_right"}
	gtk_bitset_splice              = cc.DlFunc[func(self *Bitset, position, removed, added uint32), cc.Void]{Name: "gtk_bitset_splice"}
	gtk_bitset_iter_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_bitset_iter_get_type"}
	gtk_bitset_iter_init_first     = cc.DlFunc[func(iter *BitsetIter, set *Bitset, value *uint32) int32, int32]{Name: "gtk_bitset_iter_init_first"}
	gtk_bitset_iter_init_last      = cc.DlFunc[func(iter *BitsetIter, set *Bitset, value *uint32) int32, int32]{Name: "gtk_bitset_iter_init_last"}
	gtk_bitset_iter_init_at        = cc.DlFunc[func(iter *BitsetIter, set *Bitset, target uint32, value *uint32) int32, int32]{Name: "gtk_bitset_iter_init_at"}
	gtk_bitset_iter_next           = cc.DlFunc[func(iter *BitsetIter, value *uint32) int32, int32]{Name: "gtk_bitset_iter_next"}
	gtk_bitset_iter_previous       = cc.DlFunc[func(iter *BitsetIter, value *uint32) int32, int32]{Name: "gtk_bitset_iter_previous"}
	gtk_bitset_iter_get_value      = cc.DlFunc[func(iter *BitsetIter) uint32, uint32]{Name: "gtk_bitset_iter_get_value"}
	gtk_bitset_iter_is_valid       = cc.DlFunc[func(iter *BitsetIter) int32, int32]{Name: "gtk_bitset_iter_is_valid"}
	// #endregion

	// #region BookmarkList
	gtk_bookmark_list_new             = cc.DlFunc[func(filename, attributes cc.String) *BookmarkList, *BookmarkList]{Name: "gtk_bookmark_list_new"}
	gtk_bookmark_list_get_filename    = cc.DlFunc[func(self *BookmarkList) cc.String, cc.String]{Name: "gtk_bookmark_list_get_filename"}
	gtk_bookmark_list_set_attributes  = cc.DlFunc[func(self *BookmarkList, attributes cc.String), cc.Void]{Name: "gtk_bookmark_list_set_attributes"}
	gtk_bookmark_list_get_attributes  = cc.DlFunc[func(self *BookmarkList) cc.String, cc.String]{Name: "gtk_bookmark_list_get_attributes"}
	gtk_bookmark_list_set_io_priority = cc.DlFunc[func(self *BookmarkList, ioPriority int32), cc.Void]{Name: "gtk_bookmark_list_set_io_priority"}
	gtk_bookmark_list_get_io_priority = cc.DlFunc[func(self *BookmarkList) int32, int32]{Name: "gtk_bookmark_list_get_io_priority"}
	gtk_bookmark_list_is_loading      = cc.DlFunc[func(self *BookmarkList) int32, int32]{Name: "gtk_bookmark_list_is_loading"}
	gtk_bookmark_list_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_bookmark_list_get_type"}
	// #endregion

	// #region BoolFilter
	gtk_bool_filter_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_bool_filter_get_type"}
	gtk_bool_filter_new            = cc.DlFunc[func(expression *Expression) *BoolFilter, *BoolFilter]{Name: "gtk_bool_filter_new"}
	gtk_bool_filter_get_expression = cc.DlFunc[func(self *BoolFilter) *Expression, *Expression]{Name: "gtk_bool_filter_get_expression"}
	gtk_bool_filter_set_expression = cc.DlFunc[func(self *BoolFilter, expression *Expression), cc.Void]{Name: "gtk_bool_filter_set_expression"}
	gtk_bool_filter_get_invert     = cc.DlFunc[func(self *BoolFilter) int32, int32]{Name: "gtk_bool_filter_get_invert"}
	gtk_bool_filter_set_invert     = cc.DlFunc[func(self *BoolFilter, invert int32), cc.Void]{Name: "gtk_bool_filter_set_invert"}
	// #endregion

	// #region Border
	gtk_border_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_border_get_type"}
	gtk_border_new      = cc.DlFunc[func() *Border, *Border]{Name: "gtk_border_new"}
	gtk_border_copy     = cc.DlFunc[func(border *Border) *Border, *Border]{Name: "gtk_border_copy"}
	gtk_border_free     = cc.DlFunc[func(border *Border), cc.Void]{Name: "gtk_border_free"}
	// #endregion

	// #region Box
	gtk_box_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_box_get_type"}
	gtk_box_new                   = cc.DlFunc[func(orientation Orientation, spacing int32) *Box, *Box]{Name: "gtk_box_new"}
	gtk_box_set_homogeneous       = cc.DlFunc[func(box *Box, homogeneous int32), cc.Void]{Name: "gtk_box_set_homogeneous"}
	gtk_box_get_homogeneous       = cc.DlFunc[func(box *Box) int32, int32]{Name: "gtk_box_get_homogeneous"}
	gtk_box_set_spacing           = cc.DlFunc[func(box *Box, spacing int32), cc.Void]{Name: "gtk_box_set_spacing"}
	gtk_box_get_spacing           = cc.DlFunc[func(box *Box) int32, int32]{Name: "gtk_box_get_spacing"}
	gtk_box_set_baseline_position = cc.DlFunc[func(box *Box, position BaselinePosition), cc.Void]{Name: "gtk_box_set_baseline_position"}
	gtk_box_get_baseline_position = cc.DlFunc[func(box *Box) BaselinePosition, BaselinePosition]{Name: "gtk_box_get_baseline_position"}
	gtk_box_set_baseline_child    = cc.DlFunc[func(box *Box, child int32), cc.Void]{Name: "gtk_box_set_baseline_child"}
	gtk_box_get_baseline_child    = cc.DlFunc[func(box *Box) int32, int32]{Name: "gtk_box_get_baseline_child"}
	gtk_box_append                = cc.DlFunc[func(box *Box, child *Widget), cc.Void]{Name: "gtk_box_append"}
	gtk_box_prepend               = cc.DlFunc[func(box *Box, child *Widget), cc.Void]{Name: "gtk_box_prepend"}
	gtk_box_remove                = cc.DlFunc[func(box *Box, child *Widget), cc.Void]{Name: "gtk_box_remove"}
	gtk_box_insert_child_after    = cc.DlFunc[func(box *Box, child *Widget, sibling *Widget), cc.Void]{Name: "gtk_box_insert_child_after"}
	gtk_box_reorder_child_after   = cc.DlFunc[func(box *Box, child *Widget, sibling *Widget), cc.Void]{Name: "gtk_box_reorder_child_after"}
	// #endregion

	// #region BoxLayout
	gtk_box_layout_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_box_layout_get_type"}
	gtk_box_layout_new                   = cc.DlFunc[func(orientation Orientation) *BoxLayout, *BoxLayout]{Name: "gtk_box_layout_new"}
	gtk_box_layout_set_homogeneous       = cc.DlFunc[func(boxLayout *BoxLayout, homogeneous int32), cc.Void]{Name: "gtk_box_layout_set_homogeneous"}
	gtk_box_layout_get_homogeneous       = cc.DlFunc[func(boxLayout *BoxLayout) int32, int32]{Name: "gtk_box_layout_get_homogeneous"}
	gtk_box_layout_set_spacing           = cc.DlFunc[func(boxLayout *BoxLayout, spacing uint32), cc.Void]{Name: "gtk_box_layout_set_spacing"}
	gtk_box_layout_get_spacing           = cc.DlFunc[func(boxLayout *BoxLayout) uint32, uint32]{Name: "gtk_box_layout_get_spacing"}
	gtk_box_layout_set_baseline_position = cc.DlFunc[func(boxLayout *BoxLayout, position BaselinePosition), cc.Void]{Name: "gtk_box_layout_set_baseline_position"}
	gtk_box_layout_get_baseline_position = cc.DlFunc[func(boxLayout *BoxLayout) BaselinePosition, BaselinePosition]{Name: "gtk_box_layout_get_baseline_position"}
	gtk_box_layout_set_baseline_child    = cc.DlFunc[func(boxLayout *BoxLayout, child int32), cc.Void]{Name: "gtk_box_layout_set_baseline_child"}
	gtk_box_layout_get_baseline_child    = cc.DlFunc[func(boxLayout *BoxLayout) int32, int32]{Name: "gtk_box_layout_get_baseline_child"}
	// #endregion

	// #region Buildable
	gtk_buildable_get_type                        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_buildable_get_type"}
	gtk_buildable_get_buildable_id                = cc.DlFunc[func(b *Buildable) cc.String, cc.String]{Name: "gtk_buildable_get_buildable_id"}
	gtk_buildable_parse_context_push              = cc.DlFunc[func(context *BuildableParseContext, parser *BuildableParser, userData uptr), cc.Void]{Name: "gtk_buildable_parse_context_push"}
	gtk_buildable_parse_context_pop               = cc.DlFunc[func(context *BuildableParseContext) uptr, uptr]{Name: "gtk_buildable_parse_context_pop"}
	gtk_buildable_parse_context_get_element       = cc.DlFunc[func(context *BuildableParseContext) cc.String, cc.String]{Name: "gtk_buildable_parse_context_get_element"}
	gtk_buildable_parse_context_get_element_stack = cc.DlFunc[func(context *BuildableParseContext) *glib.GPtrArray[cc.String], *glib.GPtrArray[cc.String]]{Name: "gtk_buildable_parse_context_get_element_stack"}
	gtk_buildable_parse_context_get_position      = cc.DlFunc[func(context *BuildableParseContext, lineNumber, charNumber *int32), cc.Void]{Name: "gtk_buildable_parse_context_get_position"}
	// #endregion

	// #region Builder
	gtk_builder_error_quark               = cc.DlFunc[func() glib.GQuark, glib.GQuark]{Name: "gtk_builder_error_quark"}
	gtk_builder_get_type                  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_builder_get_type"}
	gtk_builder_new                       = cc.DlFunc[func() *Builder, *Builder]{Name: "gtk_builder_new"}
	gtk_builder_add_from_file             = cc.DlFunc[func(builder *Builder, filename cc.String, err **glib.GError) int32, int32]{Name: "gtk_builder_add_from_file"}
	gtk_builder_add_from_resource         = cc.DlFunc[func(builder *Builder, resourcePath cc.String, err **glib.GError) int32, int32]{Name: "gtk_builder_add_from_resource"}
	gtk_builder_add_from_string           = cc.DlFunc[func(builder *Builder, buffer cc.String, length int64, err **glib.GError) int32, int32]{Name: "gtk_builder_add_from_string"}
	gtk_builder_add_objects_from_file     = cc.DlFunc[func(builder *Builder, filename cc.String, objectIds cc.Strings, err **glib.GError) int32, int32]{Name: "gtk_builder_add_objects_from_file"}
	gtk_builder_add_objects_from_resource = cc.DlFunc[func(builder *Builder, resourcePath cc.String, objectIds cc.Strings, err **glib.GError) int32, int32]{Name: "gtk_builder_add_objects_from_resource"}
	gtk_builder_add_objects_from_string   = cc.DlFunc[func(builder *Builder, buffer cc.String, length int64, objectIds cc.Strings, err **glib.GError) int32, int32]{Name: "gtk_builder_add_objects_from_string"}
	gtk_builder_get_object                = cc.DlFunc[func(builder *Builder, name cc.String) *gobject.GObject, *gobject.GObject]{Name: "gtk_builder_get_object"}
	gtk_builder_get_objects               = cc.DlFunc[func(builder *Builder) uptr, uptr]{Name: "gtk_builder_get_objects"}
	gtk_builder_expose_object             = cc.DlFunc[func(builder *Builder, name cc.String, object *gobject.GObject), cc.Void]{Name: "gtk_builder_expose_object"}
	gtk_builder_get_current_object        = cc.DlFunc[func(builder *Builder) *gobject.GObject, *gobject.GObject]{Name: "gtk_builder_get_current_object"}
	gtk_builder_set_current_object        = cc.DlFunc[func(builder *Builder, currentObject *gobject.GObject), cc.Void]{Name: "gtk_builder_set_current_object"}
	gtk_builder_set_translation_domain    = cc.DlFunc[func(builder *Builder, domain cc.String), cc.Void]{Name: "gtk_builder_set_translation_domain"}
	gtk_builder_get_translation_domain    = cc.DlFunc[func(builder *Builder) cc.String, cc.String]{Name: "gtk_builder_get_translation_domain"}
	gtk_builder_get_scope                 = cc.DlFunc[func(builder *Builder) *BuilderScope, *BuilderScope]{Name: "gtk_builder_get_scope"}
	gtk_builder_set_scope                 = cc.DlFunc[func(builder *Builder, scope *BuilderScope), cc.Void]{Name: "gtk_builder_set_scope"}
	gtk_builder_get_type_from_name        = cc.DlFunc[func(builder *Builder, typeName cc.String) gobject.GType, gobject.GType]{Name: "gtk_builder_get_type_from_name"}
	gtk_builder_value_from_string         = cc.DlFunc[func(builder *Builder, pspec *gobject.GParamSpec, str cc.String, value *gobject.GValue, gerr **glib.GError) int32, int32]{Name: "gtk_builder_value_from_string"}
	gtk_builder_value_from_string_type    = cc.DlFunc[func(builder *Builder, typ gobject.GType, str cc.String, value *gobject.GValue, gerr **glib.GError) int32, int32]{Name: "gtk_builder_value_from_string_type"}
	gtk_builder_new_from_file             = cc.DlFunc[func(filename cc.String) *Builder, *Builder]{Name: "gtk_builder_new_from_file"}
	gtk_builder_new_from_resource         = cc.DlFunc[func(resourcePath cc.String) *Builder, *Builder]{Name: "gtk_builder_new_from_resource"}
	gtk_builder_new_from_string           = cc.DlFunc[func(str cc.String, length int64) *Builder, *Builder]{Name: "gtk_builder_new_from_string"}
	// gtk_builder_create_closure         func(builder *Builder, functionName string, flags BuilderClosureFlags, object *gobject.GObject) (*gobject.GClosure, error)
	// #endregion

	// #region BuilderListItemFactory
	gtk_builder_list_item_factory_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_builder_list_item_factory_get_type"}
	gtk_builder_list_item_factory_new_from_bytes    = cc.DlFunc[func(scope *BuilderScope, bytes *glib.GBytes) *BuilderListItemFactory, *BuilderListItemFactory]{Name: "gtk_builder_list_item_factory_new_from_bytes"}
	gtk_builder_list_item_factory_new_from_resource = cc.DlFunc[func(scope *BuilderScope, resourcePath cc.String) *BuilderListItemFactory, *BuilderListItemFactory]{Name: "gtk_builder_list_item_factory_new_from_resource"}
	gtk_builder_list_item_factory_get_bytes         = cc.DlFunc[func(self *BuilderListItemFactory) *glib.GBytes, *glib.GBytes]{Name: "gtk_builder_list_item_factory_get_bytes"}
	gtk_builder_list_item_factory_get_resource      = cc.DlFunc[func(self *BuilderListItemFactory) cc.String, cc.String]{Name: "gtk_builder_list_item_factory_get_resource"}
	gtk_builder_list_item_factory_get_scope         = cc.DlFunc[func(self *BuilderListItemFactory) *BuilderScope, *BuilderScope]{Name: "gtk_builder_list_item_factory_get_scope"}
	// #endregion

	// #region BuilderScope
	gtk_builder_cscope_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_builder_cscope_get_type"}
	gtk_builder_cscope_new                 = cc.DlFunc[func() *BuilderCScope, *BuilderCScope]{Name: "gtk_builder_cscope_new"}
	gtk_builder_cscope_add_callback_symbol = cc.DlFunc[func(self *BuilderCScope, callbackName cc.String, callbackSymbol uptr), cc.Void]{Name: "gtk_builder_cscope_add_callback_symbol"}
	// gtk_builder_cscope_add_callback_symbols   func(self *BuilderCScope, firstCallbackName string, firstCallbackSymbol uptr, args ...interface{})
	gtk_builder_cscope_lookup_callback_symbol = cc.DlFunc[func(self *BuilderCScope, callbackName cc.String) uintptr, uintptr]{Name: "gtk_builder_cscope_lookup_callback_symbol"}
	// #endregion

	// #region Button
	gtk_button_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_button_get_type"}
	gtk_button_new                = cc.DlFunc[func() *Button, *Button]{Name: "gtk_button_new"}
	gtk_button_new_with_label     = cc.DlFunc[func(label cc.String) *Button, *Button]{Name: "gtk_button_new_with_label"}
	gtk_button_new_from_icon_name = cc.DlFunc[func(iconName cc.String) *Button, *Button]{Name: "gtk_button_new_from_icon_name"}
	gtk_button_new_with_mnemonic  = cc.DlFunc[func(label cc.String) *Button, *Button]{Name: "gtk_button_new_with_mnemonic"}
	gtk_button_set_has_frame      = cc.DlFunc[func(button *Button, hasFrame int32), cc.Void]{Name: "gtk_button_set_has_frame"}
	gtk_button_get_has_frame      = cc.DlFunc[func(button *Button) int32, int32]{Name: "gtk_button_get_has_frame"}
	gtk_button_set_label          = cc.DlFunc[func(button *Button, label cc.String), cc.Void]{Name: "gtk_button_set_label"}
	gtk_button_get_label          = cc.DlFunc[func(button *Button) cc.String, cc.String]{Name: "gtk_button_get_label"}
	gtk_button_set_use_underline  = cc.DlFunc[func(button *Button, useUnderline int32), cc.Void]{Name: "gtk_button_set_use_underline"}
	gtk_button_get_use_underline  = cc.DlFunc[func(button *Button) int32, int32]{Name: "gtk_button_get_use_underline"}
	gtk_button_set_icon_name      = cc.DlFunc[func(button *Button, iconName cc.String), cc.Void]{Name: "gtk_button_set_icon_name"}
	gtk_button_get_icon_name      = cc.DlFunc[func(button *Button) cc.String, cc.String]{Name: "gtk_button_get_icon_name"}
	gtk_button_set_child          = cc.DlFunc[func(button *Button, child *Widget), cc.Void]{Name: "gtk_button_set_child"}
	gtk_button_get_child          = cc.DlFunc[func(button *Button) *Widget, *Widget]{Name: "gtk_button_get_child"}
	gtk_button_set_can_shrink     = cc.DlFunc[func(button *Button, canShrink int32), cc.Void]{Name: "gtk_button_set_can_shrink"}
	gtk_button_get_can_shrink     = cc.DlFunc[func(button *Button) int32, int32]{Name: "gtk_button_get_can_shrink"}
	// #endregion

	// #region Calendar
	gtk_calendar_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_calendar_get_type"}
	gtk_calendar_new                   = cc.DlFunc[func() *Calendar, *Calendar]{Name: "gtk_calendar_new"}
	gtk_calendar_select_day            = cc.DlFunc[func(self *Calendar, date *glib.GDateTime), cc.Void]{Name: "gtk_calendar_select_day"}
	gtk_calendar_mark_day              = cc.DlFunc[func(calendar *Calendar, day uint32), cc.Void]{Name: "gtk_calendar_mark_day"}
	gtk_calendar_unmark_day            = cc.DlFunc[func(calendar *Calendar, day uint32), cc.Void]{Name: "gtk_calendar_unmark_day"}
	gtk_calendar_clear_marks           = cc.DlFunc[func(calendar *Calendar), cc.Void]{Name: "gtk_calendar_clear_marks"}
	gtk_calendar_set_show_week_numbers = cc.DlFunc[func(self *Calendar, value int32), cc.Void]{Name: "gtk_calendar_set_show_week_numbers"}
	gtk_calendar_get_show_week_numbers = cc.DlFunc[func(self *Calendar) int32, int32]{Name: "gtk_calendar_get_show_week_numbers"}
	gtk_calendar_set_show_heading      = cc.DlFunc[func(self *Calendar, value int32), cc.Void]{Name: "gtk_calendar_set_show_heading"}
	gtk_calendar_get_show_heading      = cc.DlFunc[func(self *Calendar) int32, int32]{Name: "gtk_calendar_get_show_heading"}
	gtk_calendar_set_show_day_names    = cc.DlFunc[func(self *Calendar, value int32), cc.Void]{Name: "gtk_calendar_set_show_day_names"}
	gtk_calendar_get_show_day_names    = cc.DlFunc[func(self *Calendar) int32, int32]{Name: "gtk_calendar_get_show_day_names"}
	gtk_calendar_set_day               = cc.DlFunc[func(self *Calendar, day int32), cc.Void]{Name: "gtk_calendar_set_day"}
	gtk_calendar_get_day               = cc.DlFunc[func(self *Calendar) int32, int32]{Name: "gtk_calendar_get_day"}
	gtk_calendar_set_month             = cc.DlFunc[func(self *Calendar, month int32), cc.Void]{Name: "gtk_calendar_set_month"}
	gtk_calendar_get_month             = cc.DlFunc[func(self *Calendar) int32, int32]{Name: "gtk_calendar_get_month"}
	gtk_calendar_set_year              = cc.DlFunc[func(self *Calendar, year int32), cc.Void]{Name: "gtk_calendar_set_year"}
	gtk_calendar_get_year              = cc.DlFunc[func(self *Calendar) int32, int32]{Name: "gtk_calendar_get_year"}
	gtk_calendar_get_date              = cc.DlFunc[func(self *Calendar) *glib.GDateTime, *glib.GDateTime]{Name: "gtk_calendar_get_date"}
	gtk_calendar_get_day_is_marked     = cc.DlFunc[func(calendar *Calendar, day uint32) int32, int32]{Name: "gtk_calendar_get_day_is_marked"}
	// #endregion

	// #region CenterBox
	gtk_center_box_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_center_box_get_type"}
	gtk_center_box_new                    = cc.DlFunc[func() *CenterBox, *CenterBox]{Name: "gtk_center_box_new"}
	gtk_center_box_set_start_widget       = cc.DlFunc[func(self *CenterBox, child *Widget), cc.Void]{Name: "gtk_center_box_set_start_widget"}
	gtk_center_box_set_center_widget      = cc.DlFunc[func(self *CenterBox, child *Widget), cc.Void]{Name: "gtk_center_box_set_center_widget"}
	gtk_center_box_set_end_widget         = cc.DlFunc[func(self *CenterBox, child *Widget), cc.Void]{Name: "gtk_center_box_set_end_widget"}
	gtk_center_box_get_start_widget       = cc.DlFunc[func(self *CenterBox) *Widget, *Widget]{Name: "gtk_center_box_get_start_widget"}
	gtk_center_box_get_center_widget      = cc.DlFunc[func(self *CenterBox) *Widget, *Widget]{Name: "gtk_center_box_get_center_widget"}
	gtk_center_box_get_end_widget         = cc.DlFunc[func(self *CenterBox) *Widget, *Widget]{Name: "gtk_center_box_get_end_widget"}
	gtk_center_box_set_baseline_position  = cc.DlFunc[func(self *CenterBox, position BaselinePosition), cc.Void]{Name: "gtk_center_box_set_baseline_position"}
	gtk_center_box_get_baseline_position  = cc.DlFunc[func(self *CenterBox) BaselinePosition, BaselinePosition]{Name: "gtk_center_box_get_baseline_position"}
	gtk_center_box_set_shrink_center_last = cc.DlFunc[func(self *CenterBox, shrinkCenterLast int32), cc.Void]{Name: "gtk_center_box_set_shrink_center_last"}
	gtk_center_box_get_shrink_center_last = cc.DlFunc[func(self *CenterBox) int32, int32]{Name: "gtk_center_box_get_shrink_center_last"}
	// #endregion

	// #region CenterLayout
	gtk_center_layout_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_center_layout_get_type"}
	gtk_center_layout_new                    = cc.DlFunc[func() *CenterLayout, *CenterLayout]{Name: "gtk_center_layout_new"}
	gtk_center_layout_set_orientation        = cc.DlFunc[func(self *CenterLayout, orientation Orientation), cc.Void]{Name: "gtk_center_layout_set_orientation"}
	gtk_center_layout_get_orientation        = cc.DlFunc[func(self *CenterLayout) Orientation, Orientation]{Name: "gtk_center_layout_get_orientation"}
	gtk_center_layout_set_baseline_position  = cc.DlFunc[func(self *CenterLayout, baselinePosition BaselinePosition), cc.Void]{Name: "gtk_center_layout_set_baseline_position"}
	gtk_center_layout_get_baseline_position  = cc.DlFunc[func(self *CenterLayout) BaselinePosition, BaselinePosition]{Name: "gtk_center_layout_get_baseline_position"}
	gtk_center_layout_set_start_widget       = cc.DlFunc[func(self *CenterLayout, widget *Widget), cc.Void]{Name: "gtk_center_layout_set_start_widget"}
	gtk_center_layout_get_start_widget       = cc.DlFunc[func(self *CenterLayout) *Widget, *Widget]{Name: "gtk_center_layout_get_start_widget"}
	gtk_center_layout_set_center_widget      = cc.DlFunc[func(self *CenterLayout, widget *Widget), cc.Void]{Name: "gtk_center_layout_set_center_widget"}
	gtk_center_layout_get_center_widget      = cc.DlFunc[func(*CenterLayout) *Widget, *Widget]{Name: "gtk_center_layout_get_center_widget"}
	gtk_center_layout_set_end_widget         = cc.DlFunc[func(*CenterLayout, *Widget), cc.Void]{Name: "gtk_center_layout_set_end_widget"}
	gtk_center_layout_get_end_widget         = cc.DlFunc[func(*CenterLayout) *Widget, *Widget]{Name: "gtk_center_layout_get_end_widget"}
	gtk_center_layout_set_shrink_center_last = cc.DlFunc[func(*CenterLayout, int32), cc.Void]{Name: "gtk_center_layout_set_shrink_center_last"}
	gtk_center_layout_get_shrink_center_last = cc.DlFunc[func(*CenterLayout) int32, int32]{Name: "gtk_center_layout_get_shrink_center_last"}
	// #endregion

	// #region CheckButton
	gtk_check_button_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_check_button_get_type"}
	gtk_check_button_new               = cc.DlFunc[func() *CheckButton, *CheckButton]{Name: "gtk_check_button_new"}
	gtk_check_button_new_with_label    = cc.DlFunc[func(label cc.String) *CheckButton, *CheckButton]{Name: "gtk_check_button_new_with_label"}
	gtk_check_button_new_with_mnemonic = cc.DlFunc[func(label cc.String) *CheckButton, *CheckButton]{Name: "gtk_check_button_new_with_mnemonic"}
	gtk_check_button_set_inconsistent  = cc.DlFunc[func(checkButton *CheckButton, inconsistent int32), cc.Void]{Name: "gtk_check_button_set_inconsistent"}
	gtk_check_button_get_inconsistent  = cc.DlFunc[func(checkButton *CheckButton) int32, int32]{Name: "gtk_check_button_get_inconsistent"}
	gtk_check_button_get_active        = cc.DlFunc[func(checkButton *CheckButton) int32, int32]{Name: "gtk_check_button_get_active"}
	gtk_check_button_set_active        = cc.DlFunc[func(checkButton *CheckButton, setting int32), cc.Void]{Name: "gtk_check_button_set_active"}
	gtk_check_button_get_label         = cc.DlFunc[func(*CheckButton) cc.String, cc.String]{Name: "gtk_check_button_get_label"}
	gtk_check_button_set_label         = cc.DlFunc[func(*CheckButton, cc.String), cc.Void]{Name: "gtk_check_button_set_label"}
	gtk_check_button_set_group         = cc.DlFunc[func(*CheckButton, *CheckButton), cc.Void]{Name: "gtk_check_button_set_group"}
	gtk_check_button_get_use_underline = cc.DlFunc[func(*CheckButton) int32, int32]{Name: "gtk_check_button_get_use_underline"}
	gtk_check_button_set_use_underline = cc.DlFunc[func(*CheckButton, int32), cc.Void]{Name: "gtk_check_button_set_use_underline"}
	gtk_check_button_get_child         = cc.DlFunc[func(*CheckButton) *Widget, *Widget]{Name: "gtk_check_button_get_child"}
	gtk_check_button_set_child         = cc.DlFunc[func(*CheckButton, *Widget), cc.Void]{Name: "gtk_check_button_set_child"}
	// #endregion

	// #region ColorDialog
	gtk_color_dialog_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_color_dialog_get_type"}
	gtk_color_dialog_new                = cc.DlFunc[func() *ColorDialog, *ColorDialog]{Name: "gtk_color_dialog_new"}
	gtk_color_dialog_get_title          = cc.DlFunc[func(*ColorDialog) cc.String, cc.String]{Name: "gtk_color_dialog_get_title"}
	gtk_color_dialog_set_title          = cc.DlFunc[func(*ColorDialog, cc.String), cc.Void]{Name: "gtk_color_dialog_set_title"}
	gtk_color_dialog_get_modal          = cc.DlFunc[func(*ColorDialog) int32, int32]{Name: "gtk_color_dialog_get_modal"}
	gtk_color_dialog_set_modal          = cc.DlFunc[func(*ColorDialog, int32), cc.Void]{Name: "gtk_color_dialog_set_modal"}
	gtk_color_dialog_get_with_alpha     = cc.DlFunc[func(*ColorDialog) int32, int32]{Name: "gtk_color_dialog_get_with_alpha"}
	gtk_color_dialog_set_with_alpha     = cc.DlFunc[func(*ColorDialog, int32), cc.Void]{Name: "gtk_color_dialog_set_with_alpha"}
	gtk_color_dialog_choose_rgba        = cc.DlFunc[func(*ColorDialog, *Window, *gdk.RGBA, *gio.GCancellable, uptr, uptr), cc.Void]{Name: "gtk_color_dialog_choose_rgba"}
	gtk_color_dialog_choose_rgba_finish = cc.DlFunc[func(*ColorDialog, *gio.GAsyncResult, **glib.GError) *gdk.RGBA, *gdk.RGBA]{Name: "gtk_color_dialog_choose_rgba_finish"}
	// #endregion

	// #region ColorDialogButton
	gtk_color_dialog_button_get_type   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_color_dialog_button_get_type"}
	gtk_color_dialog_button_new        = cc.DlFunc[func(dialog *ColorDialog) *ColorDialogButton, *ColorDialogButton]{Name: "gtk_color_dialog_button_new"}
	gtk_color_dialog_button_get_dialog = cc.DlFunc[func(self *ColorDialogButton) *ColorDialog, *ColorDialog]{Name: "gtk_color_dialog_button_get_dialog"}
	gtk_color_dialog_button_set_dialog = cc.DlFunc[func(self *ColorDialogButton, dialog *ColorDialog), cc.Void]{Name: "gtk_color_dialog_button_set_dialog"}
	gtk_color_dialog_button_get_rgba   = cc.DlFunc[func(self *ColorDialogButton) *gdk.RGBA, *gdk.RGBA]{Name: "gtk_color_dialog_button_get_rgba"}
	gtk_color_dialog_button_set_rgba   = cc.DlFunc[func(self *ColorDialogButton, color *gdk.RGBA), cc.Void]{Name: "gtk_color_dialog_button_set_rgba"}
	gtk_hsv_to_rgb                     = cc.DlFunc[func(h, s, v float32, r, g, b *float32), cc.Void]{Name: "gtk_hsv_to_rgb"}
	gtk_rgb_to_hsv                     = cc.DlFunc[func(r, g, b float32, h, s, v *float32), cc.Void]{Name: "gtk_rgb_to_hsv"}
	// #endregion

	// #region ColumnView
	gtk_column_view_get_type                   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_column_view_get_type"}
	gtk_column_view_new                        = cc.DlFunc[func(model *SelectionModel) *ColumnView, *ColumnView]{Name: "gtk_column_view_new"}
	gtk_column_view_get_columns                = cc.DlFunc[func(self *ColumnView) *gio.GListModel, *gio.GListModel]{Name: "gtk_column_view_get_columns"}
	gtk_column_view_append_column              = cc.DlFunc[func(self *ColumnView, column *ColumnViewColumn), cc.Void]{Name: "gtk_column_view_append_column"}
	gtk_column_view_remove_column              = cc.DlFunc[func(self *ColumnView, column *ColumnViewColumn), cc.Void]{Name: "gtk_column_view_remove_column"}
	gtk_column_view_insert_column              = cc.DlFunc[func(self *ColumnView, position uint32, column *ColumnViewColumn), cc.Void]{Name: "gtk_column_view_insert_column"}
	gtk_column_view_get_model                  = cc.DlFunc[func(self *ColumnView) *SelectionModel, *SelectionModel]{Name: "gtk_column_view_get_model"}
	gtk_column_view_set_model                  = cc.DlFunc[func(self *ColumnView, model *SelectionModel), cc.Void]{Name: "gtk_column_view_set_model"}
	gtk_column_view_get_show_row_separators    = cc.DlFunc[func(self *ColumnView) int32, int32]{Name: "gtk_column_view_get_show_row_separators"}
	gtk_column_view_set_show_row_separators    = cc.DlFunc[func(self *ColumnView, showRowSeparators int32), cc.Void]{Name: "gtk_column_view_set_show_row_separators"}
	gtk_column_view_get_show_column_separators = cc.DlFunc[func(self *ColumnView) int32, int32]{Name: "gtk_column_view_get_show_column_separators"}
	gtk_column_view_set_show_column_separators = cc.DlFunc[func(self *ColumnView, showColumnSeparators int32), cc.Void]{Name: "gtk_column_view_set_show_column_separators"}
	gtk_column_view_get_sorter                 = cc.DlFunc[func(self *ColumnView) *Sorter, *Sorter]{Name: "gtk_column_view_get_sorter"}
	gtk_column_view_sort_by_column             = cc.DlFunc[func(self *ColumnView, column *ColumnViewColumn, direction SortType), cc.Void]{Name: "gtk_column_view_sort_by_column"}
	gtk_column_view_set_single_click_activate  = cc.DlFunc[func(self *ColumnView, singleClickActivate int32), cc.Void]{Name: "gtk_column_view_set_single_click_activate"}
	gtk_column_view_get_single_click_activate  = cc.DlFunc[func(self *ColumnView) int32, int32]{Name: "gtk_column_view_get_single_click_activate"}
	gtk_column_view_set_reorderable            = cc.DlFunc[func(self *ColumnView, reorderable int32), cc.Void]{Name: "gtk_column_view_set_reorderable"}
	gtk_column_view_get_reorderable            = cc.DlFunc[func(self *ColumnView) int32, int32]{Name: "gtk_column_view_get_reorderable"}
	gtk_column_view_set_enable_rubberband      = cc.DlFunc[func(self *ColumnView, enableRubberband int32), cc.Void]{Name: "gtk_column_view_set_enable_rubberband"}
	gtk_column_view_get_enable_rubberband      = cc.DlFunc[func(self *ColumnView) int32, int32]{Name: "gtk_column_view_get_enable_rubberband"}
	gtk_column_view_set_tab_behavior           = cc.DlFunc[func(self *ColumnView, tabBehavior ListTabBehavior), cc.Void]{Name: "gtk_column_view_set_tab_behavior"}
	gtk_column_view_get_tab_behavior           = cc.DlFunc[func(self *ColumnView) ListTabBehavior, ListTabBehavior]{Name: "gtk_column_view_get_tab_behavior"}
	gtk_column_view_set_row_factory            = cc.DlFunc[func(self *ColumnView, factory *ListItemFactory), cc.Void]{Name: "gtk_column_view_set_row_factory"}
	gtk_column_view_get_row_factory            = cc.DlFunc[func(self *ColumnView) *ListItemFactory, *ListItemFactory]{Name: "gtk_column_view_get_row_factory"}
	gtk_column_view_set_header_factory         = cc.DlFunc[func(self *ColumnView, factory *ListItemFactory), cc.Void]{Name: "gtk_column_view_set_header_factory"}
	gtk_column_view_get_header_factory         = cc.DlFunc[func(self *ColumnView) *ListItemFactory, *ListItemFactory]{Name: "gtk_column_view_get_header_factory"}
	gtk_column_view_scroll_to                  = cc.DlFunc[func(self *ColumnView, pos uint32, column *ColumnViewColumn, flags ListScrollFlags, scroll *ScrollInfo), cc.Void]{Name: "gtk_column_view_scroll_to"}
	// #endregion

	// #region ColumnViewCell
	gtk_column_view_cell_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_column_view_cell_get_type"}
	gtk_column_view_cell_get_item      = cc.DlFunc[func(*ColumnViewCell) uptr, uptr]{Name: "gtk_column_view_cell_get_item"}
	gtk_column_view_cell_get_position  = cc.DlFunc[func(*ColumnViewCell) uint32, uint32]{Name: "gtk_column_view_cell_get_position"}
	gtk_column_view_cell_get_selected  = cc.DlFunc[func(*ColumnViewCell) int32, int32]{Name: "gtk_column_view_cell_get_selected"}
	gtk_column_view_cell_get_focusable = cc.DlFunc[func(*ColumnViewCell) int32, int32]{Name: "gtk_column_view_cell_get_focusable"}
	gtk_column_view_cell_set_focusable = cc.DlFunc[func(*ColumnViewCell, int32), cc.Void]{Name: "gtk_column_view_cell_set_focusable"}
	gtk_column_view_cell_set_child     = cc.DlFunc[func(*ColumnViewCell, *Widget), cc.Void]{Name: "gtk_column_view_cell_set_child"}
	gtk_column_view_cell_get_child     = cc.DlFunc[func(*ColumnViewCell) *Widget, *Widget]{Name: "gtk_column_view_cell_get_child"}
	// #endregion

	// #region ColumnViewColumn
	gtk_column_view_column_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_column_view_column_get_type"}
	gtk_column_view_column_new             = cc.DlFunc[func(title cc.String, factory *ListItemFactory) *ColumnViewColumn, *ColumnViewColumn]{Name: "gtk_column_view_column_new"}
	gtk_column_view_column_get_column_view = cc.DlFunc[func(self *ColumnViewColumn) *ColumnView, *ColumnView]{Name: "gtk_column_view_column_get_column_view"}
	gtk_column_view_column_set_factory     = cc.DlFunc[func(self *ColumnViewColumn, factory *ListItemFactory), cc.Void]{Name: "gtk_column_view_column_set_factory"}
	gtk_column_view_column_get_factory     = cc.DlFunc[func(self *ColumnViewColumn) *ListItemFactory, *ListItemFactory]{Name: "gtk_column_view_column_get_factory"}
	gtk_column_view_column_set_title       = cc.DlFunc[func(self *ColumnViewColumn, title cc.String), cc.Void]{Name: "gtk_column_view_column_set_title"}
	gtk_column_view_column_get_title       = cc.DlFunc[func(self *ColumnViewColumn) cc.String, cc.String]{Name: "gtk_column_view_column_get_title"}
	gtk_column_view_column_set_sorter      = cc.DlFunc[func(self *ColumnViewColumn, sorter *Sorter), cc.Void]{Name: "gtk_column_view_column_set_sorter"}
	gtk_column_view_column_get_sorter      = cc.DlFunc[func(self *ColumnViewColumn) *Sorter, *Sorter]{Name: "gtk_column_view_column_get_sorter"}
	gtk_column_view_column_set_visible     = cc.DlFunc[func(self *ColumnViewColumn, visible int32), cc.Void]{Name: "gtk_column_view_column_set_visible"}
	gtk_column_view_column_get_visible     = cc.DlFunc[func(self *ColumnViewColumn) int32, int32]{Name: "gtk_column_view_column_get_visible"}
	gtk_column_view_column_set_header_menu = cc.DlFunc[func(self *ColumnViewColumn, menu *gio.GMenuModel), cc.Void]{Name: "gtk_column_view_column_set_header_menu"}
	gtk_column_view_column_get_header_menu = cc.DlFunc[func(self *ColumnViewColumn) *gio.GMenuModel, *gio.GMenuModel]{Name: "gtk_column_view_column_get_header_menu"}
	gtk_column_view_column_set_fixed_width = cc.DlFunc[func(self *ColumnViewColumn, fixedWidth int32), cc.Void]{Name: "gtk_column_view_column_set_fixed_width"}
	gtk_column_view_column_get_fixed_width = cc.DlFunc[func(self *ColumnViewColumn) int32, int32]{Name: "gtk_column_view_column_get_fixed_width"}
	gtk_column_view_column_set_resizable   = cc.DlFunc[func(self *ColumnViewColumn, resizable int32), cc.Void]{Name: "gtk_column_view_column_set_resizable"}
	gtk_column_view_column_get_resizable   = cc.DlFunc[func(self *ColumnViewColumn) int32, int32]{Name: "gtk_column_view_column_get_resizable"}
	gtk_column_view_column_set_expand      = cc.DlFunc[func(self *ColumnViewColumn, expand int32), cc.Void]{Name: "gtk_column_view_column_set_expand"}
	gtk_column_view_column_get_expand      = cc.DlFunc[func(self *ColumnViewColumn) int32, int32]{Name: "gtk_column_view_column_get_expand"}
	gtk_column_view_column_set_id          = cc.DlFunc[func(self *ColumnViewColumn, id cc.String), cc.Void]{Name: "gtk_column_view_column_set_id"}
	gtk_column_view_column_get_id          = cc.DlFunc[func(self *ColumnViewColumn) cc.String, cc.String]{Name: "gtk_column_view_column_get_id"}
	// #endregion

	// #region ColumnViewRow
	gtk_column_view_row_get_type                   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_column_view_row_get_type"}
	gtk_column_view_row_get_item                   = cc.DlFunc[func(*ColumnViewRow) uptr, uptr]{Name: "gtk_column_view_row_get_item"}
	gtk_column_view_row_get_position               = cc.DlFunc[func(*ColumnViewRow) uint32, uint32]{Name: "gtk_column_view_row_get_position"}
	gtk_column_view_row_get_selected               = cc.DlFunc[func(*ColumnViewRow) int32, int32]{Name: "gtk_column_view_row_get_selected"}
	gtk_column_view_row_get_selectable             = cc.DlFunc[func(*ColumnViewRow) int32, int32]{Name: "gtk_column_view_row_get_selectable"}
	gtk_column_view_row_set_selectable             = cc.DlFunc[func(*ColumnViewRow, int32), cc.Void]{Name: "gtk_column_view_row_set_selectable"}
	gtk_column_view_row_get_activatable            = cc.DlFunc[func(*ColumnViewRow) int32, int32]{Name: "gtk_column_view_row_get_activatable"}
	gtk_column_view_row_set_activatable            = cc.DlFunc[func(*ColumnViewRow, int32), cc.Void]{Name: "gtk_column_view_row_set_activatable"}
	gtk_column_view_row_get_focusable              = cc.DlFunc[func(*ColumnViewRow) int32, int32]{Name: "gtk_column_view_row_get_focusable"}
	gtk_column_view_row_set_focusable              = cc.DlFunc[func(*ColumnViewRow, int32), cc.Void]{Name: "gtk_column_view_row_set_focusable"}
	gtk_column_view_row_get_accessible_description = cc.DlFunc[func(*ColumnViewRow) cc.String, cc.String]{Name: "gtk_column_view_row_get_accessible_description"}
	gtk_column_view_row_set_accessible_description = cc.DlFunc[func(*ColumnViewRow, cc.String), cc.Void]{Name: "gtk_column_view_row_set_accessible_description"}
	gtk_column_view_row_get_accessible_label       = cc.DlFunc[func(*ColumnViewRow) cc.String, cc.String]{Name: "gtk_column_view_row_get_accessible_label"}
	gtk_column_view_row_set_accessible_label       = cc.DlFunc[func(*ColumnViewRow, cc.String), cc.Void]{Name: "gtk_column_view_row_set_accessible_label"}
	// #endregion

	// #region ColumnViewSorter
	gtk_column_view_sorter_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_column_view_sorter_get_type"}
	gtk_column_view_sorter_get_primary_sort_column = cc.DlFunc[func(self *ColumnViewSorter) *ColumnViewColumn, *ColumnViewColumn]{Name: "gtk_column_view_sorter_get_primary_sort_column"}
	gtk_column_view_sorter_get_primary_sort_order  = cc.DlFunc[func(self *ColumnViewSorter) SortType, SortType]{Name: "gtk_column_view_sorter_get_primary_sort_order"}
	gtk_column_view_sorter_get_n_sort_columns      = cc.DlFunc[func(self *ColumnViewSorter) uint32, uint32]{Name: "gtk_column_view_sorter_get_n_sort_columns"}
	gtk_column_view_sorter_get_nth_sort_column     = cc.DlFunc[func(self *ColumnViewSorter, position uint32, sortOrder *SortType) *ColumnViewColumn, *ColumnViewColumn]{Name: "gtk_column_view_sorter_get_nth_sort_column"}
	// #endregion

	// #region Constraint
	gtk_constraint_new                  = cc.DlFunc[func(target *ConstraintTarget, targetAttribute ConstraintAttribute, relation ConstraintRelation, source *ConstraintTarget, sourceAttribute ConstraintAttribute, multiplier float64, constant float64, strength ConstraintStrength) *Constraint, *Constraint]{Name: "gtk_constraint_new"}
	gtk_constraint_new_constant         = cc.DlFunc[func(target *ConstraintTarget, targetAttribute ConstraintAttribute, relation ConstraintRelation, constant float64, strength ConstraintStrength) *Constraint, *Constraint]{Name: "gtk_constraint_new_constant"}
	gtk_constraint_get_target           = cc.DlFunc[func(constraint *Constraint) *ConstraintTarget, *ConstraintTarget]{Name: "gtk_constraint_get_target"}
	gtk_constraint_get_target_attribute = cc.DlFunc[func(constraint *Constraint) ConstraintAttribute, ConstraintAttribute]{Name: "gtk_constraint_get_target_attribute"}
	gtk_constraint_get_source           = cc.DlFunc[func(constraint *Constraint) *ConstraintTarget, *ConstraintTarget]{Name: "gtk_constraint_get_source"}
	gtk_constraint_get_source_attribute = cc.DlFunc[func(constraint *Constraint) ConstraintAttribute, ConstraintAttribute]{Name: "gtk_constraint_get_source_attribute"}
	gtk_constraint_get_relation         = cc.DlFunc[func(constraint *Constraint) ConstraintRelation, ConstraintRelation]{Name: "gtk_constraint_get_relation"}
	gtk_constraint_get_multiplier       = cc.DlFunc[func(constraint *Constraint) float64, float64]{Name: "gtk_constraint_get_multiplier"}
	gtk_constraint_get_constant         = cc.DlFunc[func(constraint *Constraint) float64, float64]{Name: "gtk_constraint_get_constant"}
	gtk_constraint_get_strength         = cc.DlFunc[func(constraint *Constraint) int32, int32]{Name: "gtk_constraint_get_strength"}
	gtk_constraint_is_required          = cc.DlFunc[func(constraint *Constraint) int32, int32]{Name: "gtk_constraint_is_required"}
	gtk_constraint_is_attached          = cc.DlFunc[func(constraint *Constraint) int32, int32]{Name: "gtk_constraint_is_attached"}
	gtk_constraint_is_constant          = cc.DlFunc[func(constraint *Constraint) int32, int32]{Name: "gtk_constraint_is_constant"}
	// #endregion

	// #region ConstraintGuide
	gtk_constraint_guide_get_type     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_constraint_guide_get_type"}
	gtk_constraint_guide_new          = cc.DlFunc[func() *ConstraintGuide, *ConstraintGuide]{Name: "gtk_constraint_guide_new"}
	gtk_constraint_guide_set_min_size = cc.DlFunc[func(*ConstraintGuide, int32, int32), cc.Void]{Name: "gtk_constraint_guide_set_min_size"}
	gtk_constraint_guide_get_min_size = cc.DlFunc[func(*ConstraintGuide, *int32, *int32), cc.Void]{Name: "gtk_constraint_guide_get_min_size"}
	gtk_constraint_guide_set_nat_size = cc.DlFunc[func(*ConstraintGuide, int32, int32), cc.Void]{Name: "gtk_constraint_guide_set_nat_size"}
	gtk_constraint_guide_get_nat_size = cc.DlFunc[func(*ConstraintGuide, *int32, *int32), cc.Void]{Name: "gtk_constraint_guide_get_nat_size"}
	gtk_constraint_guide_set_max_size = cc.DlFunc[func(*ConstraintGuide, int32, int32), cc.Void]{Name: "gtk_constraint_guide_set_max_size"}
	gtk_constraint_guide_get_max_size = cc.DlFunc[func(*ConstraintGuide, *int32, *int32), cc.Void]{Name: "gtk_constraint_guide_get_max_size"}
	gtk_constraint_guide_get_strength = cc.DlFunc[func(*ConstraintGuide) ConstraintStrength, ConstraintStrength]{Name: "gtk_constraint_guide_get_strength"}
	gtk_constraint_guide_set_strength = cc.DlFunc[func(*ConstraintGuide, ConstraintStrength), cc.Void]{Name: "gtk_constraint_guide_set_strength"}
	gtk_constraint_guide_set_name     = cc.DlFunc[func(*ConstraintGuide, cc.String), cc.Void]{Name: "gtk_constraint_guide_set_name"}
	gtk_constraint_guide_get_name     = cc.DlFunc[func(*ConstraintGuide) cc.String, cc.String]{Name: "gtk_constraint_guide_get_name"}
	// #endregion

	// #region ConstraintLayout
	gtk_constraint_layout_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_constraint_layout_get_type"}
	gtk_constraint_vfl_parser_error_quark   = cc.DlFunc[func() glib.GQuark, glib.GQuark]{Name: "gtk_constraint_vfl_parser_error_quark"}
	gtk_constraint_layout_new               = cc.DlFunc[func() *ConstraintLayout, *ConstraintLayout]{Name: "gtk_constraint_layout_new"}
	gtk_constraint_layout_add_constraint    = cc.DlFunc[func(*ConstraintLayout, *Constraint), cc.Void]{Name: "gtk_constraint_layout_add_constraint"}
	gtk_constraint_layout_remove_constraint = cc.DlFunc[func(*ConstraintLayout, *Constraint), cc.Void]{Name: "gtk_constraint_layout_remove_constraint"}
	gtk_constraint_layout_add_guide         = cc.DlFunc[func(*ConstraintLayout, *ConstraintGuide), cc.Void]{Name: "gtk_constraint_layout_add_guide"}
	gtk_constraint_layout_remove_guide      = cc.DlFunc[func(*ConstraintLayout, *ConstraintGuide), cc.Void]{Name: "gtk_constraint_layout_remove_guide"}
	// #endregion

	// #region CssProvider
	gtk_css_provider_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_css_provider_get_type"}
	gtk_css_provider_new                = cc.DlFunc[func() *CssProvider, *CssProvider]{Name: "gtk_css_provider_new"}
	gtk_css_provider_to_string          = cc.DlFunc[func(provider *CssProvider) cc.String, cc.String]{Name: "gtk_css_provider_to_string"}
	gtk_css_provider_load_from_data     = cc.DlFunc[func(provider *CssProvider, data uptr, length int64), cc.Void]{Name: "gtk_css_provider_load_from_data"}
	gtk_css_provider_load_from_string   = cc.DlFunc[func(provider *CssProvider, str cc.String), cc.Void]{Name: "gtk_css_provider_load_from_string"}
	gtk_css_provider_load_from_path     = cc.DlFunc[func(provider *CssProvider, path cc.String), cc.Void]{Name: "gtk_css_provider_load_from_path"}
	gtk_css_provider_load_from_resource = cc.DlFunc[func(provider *CssProvider, resourcePath cc.String), cc.Void]{Name: "gtk_css_provider_load_from_resource"}
	gtk_css_provider_load_named         = cc.DlFunc[func(provider *CssProvider, name cc.String, variant cc.String), cc.Void]{Name: "gtk_css_provider_load_named"}
	// #endregion

	// #region CustomFilter
	gtk_custom_filter_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_custom_filter_get_type"}
	gtk_custom_filter_new             = cc.DlFunc[func(matchFunc, userData uptr, userDestroy uptr) *CustomFilter, *CustomFilter]{Name: "gtk_custom_filter_new"}
	gtk_custom_filter_set_filter_func = cc.DlFunc[func(cf *CustomFilter, uptr, userData, userDestroy uptr), cc.Void]{Name: "gtk_custom_filter_set_filter_func"}
	// #endregion

	// #region CustomLayout
	gtk_custom_layout_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_custom_layout_get_type"}
	gtk_custom_layout_new      = cc.DlFunc[func(requestMode uptr, measure uptr, allocate uptr) *CustomLayout, *CustomLayout]{Name: "gtk_custom_layout_new"}
	// #endregion

	// #region CustomSorter
	gtk_custom_sorter_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_custom_sorter_get_type"}
	gtk_custom_sorter_new           = cc.DlFunc[func(sortFunc uptr, userData uptr, userDestroy uptr) *CustomSorter, *CustomSorter]{Name: "gtk_custom_sorter_new"}
	gtk_custom_sorter_set_sort_func = cc.DlFunc[func(sorter *CustomSorter, sortFunc uptr, userData uptr, userDestroy uptr), cc.Void]{Name: "gtk_custom_sorter_set_sort_func"}
	// #endregion

	// #region DebugFlags
	gtk_get_debug_flags = cc.DlFunc[func() DebugFlags, DebugFlags]{Name: "gtk_get_debug_flags"}
	gtk_set_debug_flags = cc.DlFunc[func(flags DebugFlags), cc.Void]{Name: "gtk_set_debug_flags"}
	// #endregion

	// #region DialogError
	gtk_dialog_error_quark = cc.DlFunc[func() glib.GQuark, glib.GQuark]{Name: "gtk_dialog_error_quark"}
	// #endregion

	// #region DirectoryList
	gtk_directory_list_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_directory_list_get_type"}
	gtk_directory_list_new             = cc.DlFunc[func(attributes cc.String, file *gio.GFile) *DirectoryList, *DirectoryList]{Name: "gtk_directory_list_new"}
	gtk_directory_list_set_file        = cc.DlFunc[func(dl *DirectoryList, file *gio.GFile), cc.Void]{Name: "gtk_directory_list_set_file"}
	gtk_directory_list_get_file        = cc.DlFunc[func(dl *DirectoryList) *gio.GFile, *gio.GFile]{Name: "gtk_directory_list_get_file"}
	gtk_directory_list_set_attributes  = cc.DlFunc[func(dl *DirectoryList, attributes cc.String), cc.Void]{Name: "gtk_directory_list_set_attributes"}
	gtk_directory_list_get_attributes  = cc.DlFunc[func(dl *DirectoryList) cc.String, cc.String]{Name: "gtk_directory_list_get_attributes"}
	gtk_directory_list_set_io_priority = cc.DlFunc[func(dl *DirectoryList, ioPriority int32), cc.Void]{Name: "gtk_directory_list_set_io_priority"}
	gtk_directory_list_get_io_priority = cc.DlFunc[func(dl *DirectoryList) int32, int32]{Name: "gtk_directory_list_get_io_priority"}
	gtk_directory_list_is_loading      = cc.DlFunc[func(dl *DirectoryList) int32, int32]{Name: "gtk_directory_list_is_loading"}
	gtk_directory_list_get_error       = cc.DlFunc[func(dl *DirectoryList) *glib.GError, *glib.GError]{Name: "gtk_directory_list_get_error"}
	gtk_directory_list_set_monitored   = cc.DlFunc[func(dl *DirectoryList, monitored int32), cc.Void]{Name: "gtk_directory_list_set_monitored"}
	gtk_directory_list_get_monitored   = cc.DlFunc[func(dl *DirectoryList) int32, int32]{Name: "gtk_directory_list_get_monitored"}
	// #endregion

	// #region DragIcon
	gtk_drag_icon_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_drag_icon_get_type"}
	gtk_drag_icon_get_for_drag            = cc.DlFunc[func(drag *gdk.Drag) *DragIcon, *DragIcon]{Name: "gtk_drag_icon_get_for_drag"}
	gtk_drag_icon_set_child               = cc.DlFunc[func(di *DragIcon, child *Widget), cc.Void]{Name: "gtk_drag_icon_set_child"}
	gtk_drag_icon_get_child               = cc.DlFunc[func(di *DragIcon) *Widget, *Widget]{Name: "gtk_drag_icon_get_child"}
	gtk_drag_icon_set_from_paintable      = cc.DlFunc[func(drag *gdk.Drag, paintable *gdk.Paintable, hotX, hotY int32), cc.Void]{Name: "gtk_drag_icon_set_from_paintable"}
	gtk_drag_icon_create_widget_for_value = cc.DlFunc[func(value *gobject.GValue) *DragIcon, *DragIcon]{Name: "gtk_drag_icon_create_widget_for_value"}
	// #endregion

	// #region DragSource
	gtk_drag_source_get_type    = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_drag_source_get_type"}
	gtk_drag_source_new         = cc.DlFunc[func() *DragSource, *DragSource]{Name: "gtk_drag_source_new"}
	gtk_drag_source_set_content = cc.DlFunc[func(source *DragSource, content *gdk.ContentProvider), cc.Void]{Name: "gtk_drag_source_set_content"}
	gtk_drag_source_get_content = cc.DlFunc[func(source *DragSource) *gdk.ContentProvider, *gdk.ContentProvider]{Name: "gtk_drag_source_get_content"}
	gtk_drag_source_set_actions = cc.DlFunc[func(source *DragSource, actions gdk.DragAction), cc.Void]{Name: "gtk_drag_source_set_actions"}
	gtk_drag_source_get_actions = cc.DlFunc[func(source *DragSource) gdk.DragAction, gdk.DragAction]{Name: "gtk_drag_source_get_actions"}
	gtk_drag_source_set_icon    = cc.DlFunc[func(source *DragSource, paintable *gdk.Paintable, hotX, hotY int32), cc.Void]{Name: "gtk_drag_source_set_icon"}
	gtk_drag_source_drag_cancel = cc.DlFunc[func(source *DragSource), cc.Void]{Name: "gtk_drag_source_drag_cancel"}
	gtk_drag_source_get_drag    = cc.DlFunc[func(source *DragSource) *gdk.Drag, *gdk.Drag]{Name: "gtk_drag_source_get_drag"}
	gtk_drag_check_threshold    = cc.DlFunc[func(widget *Widget, startX, startY, currentX, currentY int32) int32, int32]{Name: "gtk_drag_check_threshold"}
	// #endregion

	// #region DrawingArea
	gtk_drawing_area_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_drawing_area_get_type"}
	gtk_drawing_area_new                = cc.DlFunc[func() *DrawingArea, *DrawingArea]{Name: "gtk_drawing_area_new"}
	gtk_drawing_area_set_content_width  = cc.DlFunc[func(da *DrawingArea, width int32), cc.Void]{Name: "gtk_drawing_area_set_content_width"}
	gtk_drawing_area_get_content_width  = cc.DlFunc[func(da *DrawingArea) int32, int32]{Name: "gtk_drawing_area_get_content_width"}
	gtk_drawing_area_set_content_height = cc.DlFunc[func(da *DrawingArea, height int32), cc.Void]{Name: "gtk_drawing_area_set_content_height"}
	gtk_drawing_area_get_content_height = cc.DlFunc[func(da *DrawingArea) int32, int32]{Name: "gtk_drawing_area_get_content_height"}
	gtk_drawing_area_set_draw_func      = cc.DlFunc[func(da *DrawingArea, drawFunc uptr, userData uptr, destroy uptr), cc.Void]{Name: "gtk_drawing_area_set_draw_func"}
	// #endregion

	// #region DropControllerMotion
	gtk_drop_controller_motion_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_drop_controller_motion_get_type"}
	gtk_drop_controller_motion_new              = cc.DlFunc[func() *DropControllerMotion, *DropControllerMotion]{Name: "gtk_drop_controller_motion_new"}
	gtk_drop_controller_motion_contains_pointer = cc.DlFunc[func(controller *DropControllerMotion) int32, int32]{Name: "gtk_drop_controller_motion_contains_pointer"}
	gtk_drop_controller_motion_get_drop         = cc.DlFunc[func(controller *DropControllerMotion) *gdk.Drop, *gdk.Drop]{Name: "gtk_drop_controller_motion_get_drop"}
	gtk_drop_controller_motion_is_pointer       = cc.DlFunc[func(controller *DropControllerMotion) int32, int32]{Name: "gtk_drop_controller_motion_is_pointer"}
	// #endregion

	// #region DropDown
	gtk_drop_down_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_drop_down_get_type"}
	gtk_drop_down_new                   = cc.DlFunc[func(model *gio.GListModel, expression *Expression) *DropDown, *DropDown]{Name: "gtk_drop_down_new"}
	gtk_drop_down_new_from_strings      = cc.DlFunc[func(strings cc.Strings) *DropDown, *DropDown]{Name: "gtk_drop_down_new_from_strings"}
	gtk_drop_down_set_model             = cc.DlFunc[func(dd *DropDown, model *gio.GListModel), cc.Void]{Name: "gtk_drop_down_set_model"}
	gtk_drop_down_get_model             = cc.DlFunc[func(dd *DropDown) *gio.GListModel, *gio.GListModel]{Name: "gtk_drop_down_get_model"}
	gtk_drop_down_set_selected          = cc.DlFunc[func(dd *DropDown, position uint32), cc.Void]{Name: "gtk_drop_down_set_selected"}
	gtk_drop_down_get_selected          = cc.DlFunc[func(dd *DropDown) uint32, uint32]{Name: "gtk_drop_down_get_selected"}
	gtk_drop_down_get_selected_item     = cc.DlFunc[func(dd *DropDown) uptr, uptr]{Name: "gtk_drop_down_get_selected_item"}
	gtk_drop_down_set_factory           = cc.DlFunc[func(dd *DropDown, factory *ListItemFactory), cc.Void]{Name: "gtk_drop_down_set_factory"}
	gtk_drop_down_get_factory           = cc.DlFunc[func(dd *DropDown) *ListItemFactory, *ListItemFactory]{Name: "gtk_drop_down_get_factory"}
	gtk_drop_down_set_list_factory      = cc.DlFunc[func(dd *DropDown, factory *ListItemFactory), cc.Void]{Name: "gtk_drop_down_set_list_factory"}
	gtk_drop_down_get_list_factory      = cc.DlFunc[func(dd *DropDown) *ListItemFactory, *ListItemFactory]{Name: "gtk_drop_down_get_list_factory"}
	gtk_drop_down_set_header_factory    = cc.DlFunc[func(dd *DropDown, factory *ListItemFactory), cc.Void]{Name: "gtk_drop_down_set_header_factory"}
	gtk_drop_down_get_header_factory    = cc.DlFunc[func(dd *DropDown) *ListItemFactory, *ListItemFactory]{Name: "gtk_drop_down_get_header_factory"}
	gtk_drop_down_set_expression        = cc.DlFunc[func(dd *DropDown, expression *Expression), cc.Void]{Name: "gtk_drop_down_set_expression"}
	gtk_drop_down_get_expression        = cc.DlFunc[func(dd *DropDown) *Expression, *Expression]{Name: "gtk_drop_down_get_expression"}
	gtk_drop_down_set_enable_search     = cc.DlFunc[func(dd *DropDown, enableSearch int32), cc.Void]{Name: "gtk_drop_down_set_enable_search"}
	gtk_drop_down_get_enable_search     = cc.DlFunc[func(dd *DropDown) int32, int32]{Name: "gtk_drop_down_get_enable_search"}
	gtk_drop_down_set_show_arrow        = cc.DlFunc[func(dd *DropDown, showArrow int32), cc.Void]{Name: "gtk_drop_down_set_show_arrow"}
	gtk_drop_down_get_show_arrow        = cc.DlFunc[func(dd *DropDown) int32, int32]{Name: "gtk_drop_down_get_show_arrow"}
	gtk_drop_down_set_search_match_mode = cc.DlFunc[func(dd *DropDown, matchMode StringFilterMatchMode), cc.Void]{Name: "gtk_drop_down_set_search_match_mode"}
	gtk_drop_down_get_search_match_mode = cc.DlFunc[func(dd *DropDown) StringFilterMatchMode, StringFilterMatchMode]{Name: "gtk_drop_down_get_search_match_mode"}
	// #endregion

	// #region DropTarget
	gtk_drop_target_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_drop_target_get_type"}
	gtk_drop_target_new              = cc.DlFunc[func(gtype gobject.GType, actions gdk.DragAction) *DropTarget, *DropTarget]{Name: "gtk_drop_target_new"}
	gtk_drop_target_set_gtypes       = cc.DlFunc[func(target *DropTarget, types *gobject.GType, nTypes uint64), cc.Void]{Name: "gtk_drop_target_set_gtypes"}
	gtk_drop_target_get_gtypes       = cc.DlFunc[func(target *DropTarget, nTypes *uint64) *gobject.GType, *gobject.GType]{Name: "gtk_drop_target_get_gtypes"}
	gtk_drop_target_get_formats      = cc.DlFunc[func(target *DropTarget) *gdk.ContentFormats, *gdk.ContentFormats]{Name: "gtk_drop_target_get_formats"}
	gtk_drop_target_set_actions      = cc.DlFunc[func(target *DropTarget, actions gdk.DragAction), cc.Void]{Name: "gtk_drop_target_set_actions"}
	gtk_drop_target_get_actions      = cc.DlFunc[func(target *DropTarget) gdk.DragAction, gdk.DragAction]{Name: "gtk_drop_target_get_actions"}
	gtk_drop_target_set_preload      = cc.DlFunc[func(target *DropTarget, preload int32), cc.Void]{Name: "gtk_drop_target_set_preload"}
	gtk_drop_target_get_preload      = cc.DlFunc[func(target *DropTarget) int32, int32]{Name: "gtk_drop_target_get_preload"}
	gtk_drop_target_get_current_drop = cc.DlFunc[func(target *DropTarget) *gdk.Drop, *gdk.Drop]{Name: "gtk_drop_target_get_current_drop"}
	gtk_drop_target_get_value        = cc.DlFunc[func(target *DropTarget) *gobject.GValue, *gobject.GValue]{Name: "gtk_drop_target_get_value"}
	gtk_drop_target_reject           = cc.DlFunc[func(target *DropTarget), cc.Void]{Name: "gtk_drop_target_reject"}
	// #endregion

	// #region DropTargetAsync
	gtk_drop_target_async_get_type    = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_drop_target_async_get_type"}
	gtk_drop_target_async_new         = cc.DlFunc[func(formats *gdk.ContentFormats, actions gdk.DragAction) *DropTargetAsync, *DropTargetAsync]{Name: "gtk_drop_target_async_new"}
	gtk_drop_target_async_set_formats = cc.DlFunc[func(target *DropTargetAsync, formats *gdk.ContentFormats), cc.Void]{Name: "gtk_drop_target_async_set_formats"}
	gtk_drop_target_async_get_formats = cc.DlFunc[func(target *DropTargetAsync) *gdk.ContentFormats, *gdk.ContentFormats]{Name: "gtk_drop_target_async_get_formats"}
	gtk_drop_target_async_set_actions = cc.DlFunc[func(target *DropTargetAsync, actions gdk.DragAction), cc.Void]{Name: "gtk_drop_target_async_set_actions"}
	gtk_drop_target_async_get_actions = cc.DlFunc[func(target *DropTargetAsync) gdk.DragAction, gdk.DragAction]{Name: "gtk_drop_target_async_get_actions"}
	gtk_drop_target_async_reject_drop = cc.DlFunc[func(target *DropTargetAsync, drop *gdk.Drop), cc.Void]{Name: "gtk_drop_target_async_reject_drop"}
	// #endregion

	// #region Editable
	gtk_editable_get_type                               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_editable_get_type"}
	gtk_editable_get_text                               = cc.DlFunc[func(e *Editable) cc.String, cc.String]{Name: "gtk_editable_get_text"}
	gtk_editable_set_text                               = cc.DlFunc[func(e *Editable, text cc.String), cc.Void]{Name: "gtk_editable_set_text"}
	gtk_editable_get_chars                              = cc.DlFunc[func(e *Editable, startPos, endPos int32) cc.String, cc.String]{Name: "gtk_editable_get_chars"}
	gtk_editable_insert_text                            = cc.DlFunc[func(e *Editable, text cc.String, length int32, position *int32), cc.Void]{Name: "gtk_editable_insert_text"}
	gtk_editable_delete_text                            = cc.DlFunc[func(e *Editable, startPos, endPos int32), cc.Void]{Name: "gtk_editable_delete_text"}
	gtk_editable_get_selection_bounds                   = cc.DlFunc[func(e *Editable, startPos, endPos *int32) int32, int32]{Name: "gtk_editable_get_selection_bounds"}
	gtk_editable_delete_selection                       = cc.DlFunc[func(e *Editable), cc.Void]{Name: "gtk_editable_delete_selection"}
	gtk_editable_select_region                          = cc.DlFunc[func(e *Editable, startPos, endPos int32), cc.Void]{Name: "gtk_editable_select_region"}
	gtk_editable_set_position                           = cc.DlFunc[func(e *Editable, position int32), cc.Void]{Name: "gtk_editable_set_position"}
	gtk_editable_get_position                           = cc.DlFunc[func(e *Editable) int32, int32]{Name: "gtk_editable_get_position"}
	gtk_editable_get_editable                           = cc.DlFunc[func(e *Editable) int32, int32]{Name: "gtk_editable_get_editable"}
	gtk_editable_set_editable                           = cc.DlFunc[func(e *Editable, isEditable int32), cc.Void]{Name: "gtk_editable_set_editable"}
	gtk_editable_get_alignment                          = cc.DlFunc[func(e *Editable) float32, float32]{Name: "gtk_editable_get_alignment"}
	gtk_editable_set_alignment                          = cc.DlFunc[func(e *Editable, xalign float32), cc.Void]{Name: "gtk_editable_set_alignment"}
	gtk_editable_get_width_chars                        = cc.DlFunc[func(e *Editable) int32, int32]{Name: "gtk_editable_get_width_chars"}
	gtk_editable_set_width_chars                        = cc.DlFunc[func(e *Editable, nChars int32), cc.Void]{Name: "gtk_editable_set_width_chars"}
	gtk_editable_get_max_width_chars                    = cc.DlFunc[func(e *Editable) int32, int32]{Name: "gtk_editable_get_max_width_chars"}
	gtk_editable_set_max_width_chars                    = cc.DlFunc[func(e *Editable, nChars int32), cc.Void]{Name: "gtk_editable_set_max_width_chars"}
	gtk_editable_get_enable_undo                        = cc.DlFunc[func(e *Editable) int32, int32]{Name: "gtk_editable_get_enable_undo"}
	gtk_editable_set_enable_undo                        = cc.DlFunc[func(e *Editable, enableUndo int32), cc.Void]{Name: "gtk_editable_set_enable_undo"}
	gtk_editable_install_properties                     = cc.DlFunc[func(objectClass *gobject.GObjectClass, firstProp uint32) uint32, uint32]{Name: "gtk_editable_install_properties"}
	gtk_editable_get_delegate                           = cc.DlFunc[func(editable *Editable) *Editable, *Editable]{Name: "gtk_editable_get_delegate"}
	gtk_editable_init_delegate                          = cc.DlFunc[func(editable *Editable), cc.Void]{Name: "gtk_editable_init_delegate"}
	gtk_editable_finish_delegate                        = cc.DlFunc[func(editable *Editable), cc.Void]{Name: "gtk_editable_finish_delegate"}
	gtk_editable_delegate_set_property                  = cc.DlFunc[func(object *gobject.GObject, propId uint32, value *gobject.GValue, pspec *gobject.GParamSpec) int32, int32]{Name: "gtk_editable_delegate_set_property"}
	gtk_editable_delegate_get_property                  = cc.DlFunc[func(object *gobject.GObject, propId uint32, value *gobject.GValue, pspec *gobject.GParamSpec) int32, int32]{Name: "gtk_editable_delegate_get_property"}
	gtk_editable_delegate_get_accessible_platform_state = cc.DlFunc[func(editable *Editable, state AccessiblePlatformState) int32, int32]{Name: "gtk_editable_delegate_get_accessible_platform_state"}
	// #endregion

	// #region EditableLabel
	g_editable_label_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "g_editable_label_get_type"}
	g_editable_label_new           = cc.DlFunc[func(str cc.String) *EditableLabel, *EditableLabel]{Name: "g_editable_label_new"}
	g_editable_label_get_editing   = cc.DlFunc[func(label *EditableLabel) int32, int32]{Name: "g_editable_label_get_editing"}
	g_editable_label_start_editing = cc.DlFunc[func(label *EditableLabel), cc.Void]{Name: "g_editable_label_start_editing"}
	g_editable_label_stop_editing  = cc.DlFunc[func(label *EditableLabel, commit int32), cc.Void]{Name: "g_editable_label_stop_editing"}
	// #endregion

	// #region EmojiChooser
	gtk_emoji_chooser_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_emoji_chooser_get_type"}
	gtk_emoji_chooser_new      = cc.DlFunc[func() *EmojiChooser, *EmojiChooser]{Name: "gtk_emoji_chooser_new"}
	// #endregion

	// #region Entry
	gtk_entry_get_type                     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_entry_get_type"}
	gtk_entry_new                          = cc.DlFunc[func() *Entry, *Entry]{Name: "gtk_entry_new"}
	gtk_entry_new_with_buffer              = cc.DlFunc[func(buffer *EntryBuffer) *Entry, *Entry]{Name: "gtk_entry_new_with_buffer"}
	gtk_entry_get_buffer                   = cc.DlFunc[func(e *Entry) *EntryBuffer, *EntryBuffer]{Name: "gtk_entry_get_buffer"}
	gtk_entry_set_buffer                   = cc.DlFunc[func(e *Entry, buffer *EntryBuffer), cc.Void]{Name: "gtk_entry_set_buffer"}
	gtk_entry_set_visibility               = cc.DlFunc[func(e *Entry, visible int32), cc.Void]{Name: "gtk_entry_set_visibility"}
	gtk_entry_get_visibility               = cc.DlFunc[func(e *Entry) int32, int32]{Name: "gtk_entry_get_visibility"}
	gtk_entry_set_invisible_char           = cc.DlFunc[func(e *Entry, ch uint32), cc.Void]{Name: "gtk_entry_set_invisible_char"}
	gtk_entry_get_invisible_char           = cc.DlFunc[func(e *Entry) uint32, uint32]{Name: "gtk_entry_get_invisible_char"}
	gtk_entry_unset_invisible_char         = cc.DlFunc[func(e *Entry), cc.Void]{Name: "gtk_entry_unset_invisible_char"}
	gtk_entry_set_has_frame                = cc.DlFunc[func(e *Entry, setting int32), cc.Void]{Name: "gtk_entry_set_has_frame"}
	gtk_entry_get_has_frame                = cc.DlFunc[func(e *Entry) int32, int32]{Name: "gtk_entry_get_has_frame"}
	gtk_entry_set_overwrite_mode           = cc.DlFunc[func(e *Entry, overwrite int32), cc.Void]{Name: "gtk_entry_set_overwrite_mode"}
	gtk_entry_get_overwrite_mode           = cc.DlFunc[func(e *Entry) int32, int32]{Name: "gtk_entry_get_overwrite_mode"}
	gtk_entry_set_max_length               = cc.DlFunc[func(e *Entry, max int32), cc.Void]{Name: "gtk_entry_set_max_length"}
	gtk_entry_get_max_length               = cc.DlFunc[func(e *Entry) int32, int32]{Name: "gtk_entry_get_max_length"}
	gtk_entry_get_text_length              = cc.DlFunc[func(e *Entry) uint16, uint16]{Name: "gtk_entry_get_text_length"}
	gtk_entry_set_activates_default        = cc.DlFunc[func(e *Entry, setting int32), cc.Void]{Name: "gtk_entry_set_activates_default"}
	gtk_entry_get_activates_default        = cc.DlFunc[func(e *Entry) int32, int32]{Name: "gtk_entry_get_activates_default"}
	gtk_entry_set_alignment                = cc.DlFunc[func(e *Entry, xalign float32), cc.Void]{Name: "gtk_entry_set_alignment"}
	gtk_entry_get_alignment                = cc.DlFunc[func(e *Entry) float32, float32]{Name: "gtk_entry_get_alignment"}
	gtk_entry_set_progress_fraction        = cc.DlFunc[func(e *Entry, fraction float64), cc.Void]{Name: "gtk_entry_set_progress_fraction"}
	gtk_entry_get_progress_fraction        = cc.DlFunc[func(e *Entry) float64, float64]{Name: "gtk_entry_get_progress_fraction"}
	gtk_entry_set_progress_pulse_step      = cc.DlFunc[func(e *Entry, fraction float64), cc.Void]{Name: "gtk_entry_set_progress_pulse_step"}
	gtk_entry_get_progress_pulse_step      = cc.DlFunc[func(e *Entry) float64, float64]{Name: "gtk_entry_get_progress_pulse_step"}
	gtk_entry_progress_pulse               = cc.DlFunc[func(e *Entry), cc.Void]{Name: "gtk_entry_progress_pulse"}
	gtk_entry_get_placeholder_text         = cc.DlFunc[func(e *Entry) cc.String, cc.String]{Name: "gtk_entry_get_placeholder_text"}
	gtk_entry_set_placeholder_text         = cc.DlFunc[func(e *Entry, text cc.String), cc.Void]{Name: "gtk_entry_set_placeholder_text"}
	gtk_entry_set_icon_from_paintable      = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition, paintable *gdk.Paintable), cc.Void]{Name: "gtk_entry_set_icon_from_paintable"}
	gtk_entry_set_icon_from_icon_name      = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition, iconName cc.String), cc.Void]{Name: "gtk_entry_set_icon_from_icon_name"}
	gtk_entry_set_icon_from_gicon          = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition, icon *gio.GIcon), cc.Void]{Name: "gtk_entry_set_icon_from_gicon"}
	gtk_entry_get_icon_storage_type        = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition) ImageType, ImageType]{Name: "gtk_entry_get_icon_storage_type"}
	gtk_entry_get_icon_paintable           = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition) *gdk.Paintable, *gdk.Paintable]{Name: "gtk_entry_get_icon_paintable"}
	gtk_entry_get_icon_name                = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition) cc.String, cc.String]{Name: "gtk_entry_get_icon_name"}
	gtk_entry_get_icon_gicon               = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition) *gio.GIcon, *gio.GIcon]{Name: "gtk_entry_get_icon_gicon"}
	gtk_entry_set_icon_activatable         = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition, activatable int32), cc.Void]{Name: "gtk_entry_set_icon_activatable"}
	gtk_entry_get_icon_activatable         = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition) int32, int32]{Name: "gtk_entry_get_icon_activatable"}
	gtk_entry_set_icon_sensitive           = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition, sensitive int32), cc.Void]{Name: "gtk_entry_set_icon_sensitive"}
	gtk_entry_get_icon_sensitive           = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition) int32, int32]{Name: "gtk_entry_get_icon_sensitive"}
	gtk_entry_get_icon_at_pos              = cc.DlFunc[func(e *Entry, x, y int32) int32, int32]{Name: "gtk_entry_get_icon_at_pos"}
	gtk_entry_set_icon_tooltip_text        = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition, tooltip cc.String), cc.Void]{Name: "gtk_entry_set_icon_tooltip_text"}
	gtk_entry_get_icon_tooltip_text        = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition) cc.String, cc.String]{Name: "gtk_entry_get_icon_tooltip_text"}
	gtk_entry_set_icon_tooltip_markup      = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition, tooltip cc.String), cc.Void]{Name: "gtk_entry_set_icon_tooltip_markup"}
	gtk_entry_get_icon_tooltip_markup      = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition) cc.String, cc.String]{Name: "gtk_entry_get_icon_tooltip_markup"}
	gtk_entry_set_icon_drag_source         = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition, provider *gdk.ContentProvider, actions gdk.DragAction), cc.Void]{Name: "gtk_entry_set_icon_drag_source"}
	gtk_entry_get_current_icon_drag_source = cc.DlFunc[func(e *Entry) int32, int32]{Name: "gtk_entry_get_current_icon_drag_source"}
	gtk_entry_get_icon_area                = cc.DlFunc[func(e *Entry, iconPos EntryIconPosition, iconArea *gdk.Rectangle), cc.Void]{Name: "gtk_entry_get_icon_area"}
	gtk_entry_reset_im_context             = cc.DlFunc[func(e *Entry), cc.Void]{Name: "gtk_entry_reset_im_context"}
	gtk_entry_set_input_purpose            = cc.DlFunc[func(e *Entry, purpose InputPurpose), cc.Void]{Name: "gtk_entry_set_input_purpose"}
	gtk_entry_get_input_purpose            = cc.DlFunc[func(e *Entry) InputPurpose, InputPurpose]{Name: "gtk_entry_get_input_purpose"}
	gtk_entry_set_input_hints              = cc.DlFunc[func(e *Entry, hints InputHints), cc.Void]{Name: "gtk_entry_set_input_hints"}
	gtk_entry_get_input_hints              = cc.DlFunc[func(e *Entry) InputHints, InputHints]{Name: "gtk_entry_get_input_hints"}
	gtk_entry_set_attributes               = cc.DlFunc[func(e *Entry, attrs *pango.AttrList), cc.Void]{Name: "gtk_entry_set_attributes"}
	gtk_entry_get_attributes               = cc.DlFunc[func(e *Entry) *pango.AttrList, *pango.AttrList]{Name: "gtk_entry_get_attributes"}
	gtk_entry_set_tabs                     = cc.DlFunc[func(e *Entry, tabs *pango.TabArray), cc.Void]{Name: "gtk_entry_set_tabs"}
	gtk_entry_get_tabs                     = cc.DlFunc[func(e *Entry) *pango.TabArray, *pango.TabArray]{Name: "gtk_entry_get_tabs"}
	gtk_entry_grab_focus_without_selecting = cc.DlFunc[func(e *Entry) int32, int32]{Name: "gtk_entry_grab_focus_without_selecting"}
	gtk_entry_set_extra_menu               = cc.DlFunc[func(e *Entry, model *gio.GMenuModel), cc.Void]{Name: "gtk_entry_set_extra_menu"}
	gtk_entry_get_extra_menu               = cc.DlFunc[func(e *Entry) *gio.GMenuModel, *gio.GMenuModel]{Name: "gtk_entry_get_extra_menu"}
	// #endregion

	// #region EntryBuffer
	gtk_entry_buffer_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_entry_buffer_get_type"}
	gtk_entry_buffer_new                = cc.DlFunc[func(initialChars cc.String, nInitialChars int32) *EntryBuffer, *EntryBuffer]{Name: "gtk_entry_buffer_new"}
	gtk_entry_buffer_get_bytes          = cc.DlFunc[func(b *EntryBuffer) uint64, uint64]{Name: "gtk_entry_buffer_get_bytes"}
	gtk_entry_buffer_get_length         = cc.DlFunc[func(b *EntryBuffer) uint32, uint32]{Name: "gtk_entry_buffer_get_length"}
	gtk_entry_buffer_get_text           = cc.DlFunc[func(b *EntryBuffer) cc.String, cc.String]{Name: "gtk_entry_buffer_get_text"}
	gtk_entry_buffer_set_text           = cc.DlFunc[func(b *EntryBuffer, chars cc.String, nChars int32), cc.Void]{Name: "gtk_entry_buffer_set_text"}
	gtk_entry_buffer_set_max_length     = cc.DlFunc[func(b *EntryBuffer, maxLength int32), cc.Void]{Name: "gtk_entry_buffer_set_max_length"}
	gtk_entry_buffer_get_max_length     = cc.DlFunc[func(b *EntryBuffer) int32, int32]{Name: "gtk_entry_buffer_get_max_length"}
	gtk_entry_buffer_insert_text        = cc.DlFunc[func(b *EntryBuffer, position uint32, chars cc.String, nChars int32) uint32, uint32]{Name: "gtk_entry_buffer_insert_text"}
	gtk_entry_buffer_delete_text        = cc.DlFunc[func(b *EntryBuffer, position uint32, nChars int32) uint32, uint32]{Name: "gtk_entry_buffer_delete_text"}
	gtk_entry_buffer_emit_inserted_text = cc.DlFunc[func(b *EntryBuffer, position uint32, chars cc.String, nChars uint32), cc.Void]{Name: "gtk_entry_buffer_emit_inserted_text"}
	gtk_entry_buffer_emit_deleted_text  = cc.DlFunc[func(b *EntryBuffer, position, nChars uint32), cc.Void]{Name: "gtk_entry_buffer_emit_deleted_text"}
	// #endregion

	// #region EventController
	gtk_event_controller_get_type                 = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_event_controller_get_type"}
	gtk_event_controller_get_widget               = cc.DlFunc[func(controller *EventController) *Widget, *Widget]{Name: "gtk_event_controller_get_widget"}
	gtk_event_controller_reset                    = cc.DlFunc[func(controller *EventController), cc.Void]{Name: "gtk_event_controller_reset"}
	gtk_event_controller_get_propagation_phase    = cc.DlFunc[func(controller *EventController) PropagationPhase, PropagationPhase]{Name: "gtk_event_controller_get_propagation_phase"}
	gtk_event_controller_set_propagation_phase    = cc.DlFunc[func(controller *EventController, phase PropagationPhase), cc.Void]{Name: "gtk_event_controller_set_propagation_phase"}
	gtk_event_controller_get_propagation_limit    = cc.DlFunc[func(controller *EventController) PropagationLimit, PropagationLimit]{Name: "gtk_event_controller_get_propagation_limit"}
	gtk_event_controller_set_propagation_limit    = cc.DlFunc[func(controller *EventController, limit PropagationLimit), cc.Void]{Name: "gtk_event_controller_set_propagation_limit"}
	gtk_event_controller_get_name                 = cc.DlFunc[func(controller *EventController) cc.String, cc.String]{Name: "gtk_event_controller_get_name"}
	gtk_event_controller_set_name                 = cc.DlFunc[func(controller *EventController, name cc.String), cc.Void]{Name: "gtk_event_controller_set_name"}
	gtk_event_controller_set_static_name          = cc.DlFunc[func(controller *EventController, name cc.String), cc.Void]{Name: "gtk_event_controller_set_static_name"}
	gtk_event_controller_get_current_event        = cc.DlFunc[func(controller *EventController) *gdk.Event, *gdk.Event]{Name: "gtk_event_controller_get_current_event"}
	gtk_event_controller_get_current_event_time   = cc.DlFunc[func(controller *EventController) uint32, uint32]{Name: "gtk_event_controller_get_current_event_time"}
	gtk_event_controller_get_current_event_device = cc.DlFunc[func(controller *EventController) *gdk.Device, *gdk.Device]{Name: "gtk_event_controller_get_current_event_device"}
	gtk_event_controller_get_current_event_state  = cc.DlFunc[func(controller *EventController) gdk.ModifierType, gdk.ModifierType]{Name: "gtk_event_controller_get_current_event_state"}
	// #endregion

	// #region EventControllerFocus
	gtk_event_controller_focus_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_event_controller_focus_get_type"}
	gtk_event_controller_focus_new            = cc.DlFunc[func() *EventControllerFocus, *EventControllerFocus]{Name: "gtk_event_controller_focus_new"}
	gtk_event_controller_focus_contains_focus = cc.DlFunc[func(controller *EventControllerFocus) int32, int32]{Name: "gtk_event_controller_focus_contains_focus"}
	gtk_event_controller_focus_is_focus       = cc.DlFunc[func(controller *EventControllerFocus) int32, int32]{Name: "gtk_event_controller_focus_is_focus"}
	// #endregion

	// #region EventControllerKey
	gtk_event_controller_key_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_event_controller_key_get_type"}
	gtk_event_controller_key_new            = cc.DlFunc[func() *EventControllerKey, *EventControllerKey]{Name: "gtk_event_controller_key_new"}
	gtk_event_controller_key_set_im_context = cc.DlFunc[func(controller *EventControllerKey, imContext *IMContext), cc.Void]{Name: "gtk_event_controller_key_set_im_context"}
	gtk_event_controller_key_get_im_context = cc.DlFunc[func(controller *EventControllerKey) *IMContext, *IMContext]{Name: "gtk_event_controller_key_get_im_context"}
	gtk_event_controller_key_forward        = cc.DlFunc[func(controller *EventControllerKey, widget *Widget) int32, int32]{Name: "gtk_event_controller_key_forward"}
	gtk_event_controller_key_get_group      = cc.DlFunc[func(controller *EventControllerKey) uint32, uint32]{Name: "gtk_event_controller_key_get_group"}
	// #endregion

	// #region EventControllerLegacy
	gtk_event_controller_legacy_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_event_controller_legacy_get_type"}
	gtk_event_controller_legacy_new      = cc.DlFunc[func() *EventControllerLegacy, *EventControllerLegacy]{Name: "gtk_event_controller_legacy_new"}
	// #endregion

	// #region EventControllerMotion
	gtk_event_controller_motion_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_event_controller_motion_get_type"}
	gtk_event_controller_motion_new              = cc.DlFunc[func() *EventControllerMotion, *EventControllerMotion]{Name: "gtk_event_controller_motion_new"}
	gtk_event_controller_motion_contains_pointer = cc.DlFunc[func(controller *EventControllerMotion) int32, int32]{Name: "gtk_event_controller_motion_contains_pointer"}
	gtk_event_controller_motion_is_pointer       = cc.DlFunc[func(controller *EventControllerMotion) int32, int32]{Name: "gtk_event_controller_motion_is_pointer"}
	// #endregion

	// #region EventControllerScroll
	gtk_event_controller_scroll_get_type  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_event_controller_scroll_get_type"}
	gtk_event_controller_scroll_new       = cc.DlFunc[func(flags EventControllerScrollFlags) *EventControllerScroll, *EventControllerScroll]{Name: "gtk_event_controller_scroll_new"}
	gtk_event_controller_scroll_set_flags = cc.DlFunc[func(scroll *EventControllerScroll, flags EventControllerScrollFlags), cc.Void]{Name: "gtk_event_controller_scroll_set_flags"}
	gtk_event_controller_scroll_get_flags = cc.DlFunc[func(scroll *EventControllerScroll) EventControllerScrollFlags, EventControllerScrollFlags]{Name: "gtk_event_controller_scroll_get_flags"}
	gtk_event_controller_scroll_get_unit  = cc.DlFunc[func(scroll *EventControllerScroll) gdk.ScrollUnit, gdk.ScrollUnit]{Name: "gtk_event_controller_scroll_get_unit"}
	// #endregion

	// #region Expander
	gtk_expander_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_expander_get_type"}
	gtk_expander_new                 = cc.DlFunc[func(label cc.String) *Expander, *Expander]{Name: "gtk_expander_new"}
	gtk_expander_new_with_mnemonic   = cc.DlFunc[func(label cc.String) *Expander, *Expander]{Name: "gtk_expander_new_with_mnemonic"}
	gtk_expander_set_expanded        = cc.DlFunc[func(e *Expander, expanded int32), cc.Void]{Name: "gtk_expander_set_expanded"}
	gtk_expander_get_expanded        = cc.DlFunc[func(e *Expander) int32, int32]{Name: "gtk_expander_get_expanded"}
	gtk_expander_set_label           = cc.DlFunc[func(e *Expander, label cc.String), cc.Void]{Name: "gtk_expander_set_label"}
	gtk_expander_get_label           = cc.DlFunc[func(e *Expander) cc.String, cc.String]{Name: "gtk_expander_get_label"}
	gtk_expander_set_use_underline   = cc.DlFunc[func(e *Expander, useUnderline int32), cc.Void]{Name: "gtk_expander_set_use_underline"}
	gtk_expander_get_use_underline   = cc.DlFunc[func(e *Expander) int32, int32]{Name: "gtk_expander_get_use_underline"}
	gtk_expander_set_use_markup      = cc.DlFunc[func(e *Expander, useMarkup int32), cc.Void]{Name: "gtk_expander_set_use_markup"}
	gtk_expander_get_use_markup      = cc.DlFunc[func(e *Expander) int32, int32]{Name: "gtk_expander_get_use_markup"}
	gtk_expander_set_label_widget    = cc.DlFunc[func(e *Expander, labelWidget *Widget), cc.Void]{Name: "gtk_expander_set_label_widget"}
	gtk_expander_get_label_widget    = cc.DlFunc[func(e *Expander) *Widget, *Widget]{Name: "gtk_expander_get_label_widget"}
	gtk_expander_set_resize_toplevel = cc.DlFunc[func(e *Expander, resizeToplevel int32), cc.Void]{Name: "gtk_expander_set_resize_toplevel"}
	gtk_expander_get_resize_toplevel = cc.DlFunc[func(e *Expander) int32, int32]{Name: "gtk_expander_get_resize_toplevel"}
	gtk_expander_set_child           = cc.DlFunc[func(e *Expander, child *Widget), cc.Void]{Name: "gtk_expander_set_child"}
	gtk_expander_get_child           = cc.DlFunc[func(e *Expander) *Widget, *Widget]{Name: "gtk_expander_get_child"}
	// #endregion

	// #region Expression
	gtk_expression_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_expression_get_type"}
	gtk_expression_ref                     = cc.DlFunc[func(self *Expression) *Expression, *Expression]{Name: "gtk_expression_ref"}
	gtk_expression_unref                   = cc.DlFunc[func(self *Expression), cc.Void]{Name: "gtk_expression_unref"}
	gtk_expression_get_value_type          = cc.DlFunc[func(self *Expression) gobject.GType, gobject.GType]{Name: "gtk_expression_get_value_type"}
	gtk_expression_is_static               = cc.DlFunc[func(self *Expression) int32, int32]{Name: "gtk_expression_is_static"}
	gtk_expression_evaluate                = cc.DlFunc[func(self *Expression, this_ *gobject.GObject, value *gobject.GValue) int32, int32]{Name: "gtk_expression_evaluate"}
	gtk_expression_watch                   = cc.DlFunc[func(self *Expression, this_ *gobject.GObject, notify uptr, userData uptr, userDestroy uptr) *ExpressionWatch, *ExpressionWatch]{Name: "gtk_expression_watch"}
	gtk_expression_bind                    = cc.DlFunc[func(self *Expression, target *gobject.GObject, property cc.String, this *gobject.GObject) *ExpressionWatch, *ExpressionWatch]{Name: "gtk_expression_bind"}
	gtk_expression_watch_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_expression_watch_get_type"}
	gtk_expression_watch_ref               = cc.DlFunc[func(watch *ExpressionWatch) *ExpressionWatch, *ExpressionWatch]{Name: "gtk_expression_watch_ref"}
	gtk_expression_watch_unref             = cc.DlFunc[func(watch *ExpressionWatch), cc.Void]{Name: "gtk_expression_watch_unref"}
	gtk_expression_watch_evaluate          = cc.DlFunc[func(watch *ExpressionWatch, value *gobject.GValue) int32, int32]{Name: "gtk_expression_watch_evaluate"}
	gtk_expression_watch_unwatch           = cc.DlFunc[func(watch *ExpressionWatch), cc.Void]{Name: "gtk_expression_watch_unwatch"}
	gtk_property_expression_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_property_expression_get_type"}
	gtk_property_expression_new            = cc.DlFunc[func(thisType gobject.GType, expression *Expression, propertyName cc.String) *PropertyExpression, *PropertyExpression]{Name: "gtk_property_expression_new"}
	gtk_property_expression_new_for_pspec  = cc.DlFunc[func(expression *Expression, pspec *gobject.GParamSpec) *PropertyExpression, *PropertyExpression]{Name: "gtk_property_expression_new_for_pspec"}
	gtk_property_expression_get_expression = cc.DlFunc[func(expression *PropertyExpression) *Expression, *Expression]{Name: "gtk_property_expression_get_expression"}
	gtk_property_expression_get_pspec      = cc.DlFunc[func(expression *PropertyExpression) *gobject.GParamSpec, *gobject.GParamSpec]{Name: "gtk_property_expression_get_pspec"}
	gtk_constant_expression_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_constant_expression_get_type"}
	gtk_constant_expression_new            = cc.DlFunc[func(valueType gobject.GType, args ...interface{}) *ConstantExpression, *ConstantExpression]{Name: "gtk_constant_expression_new", Va: true}
	gtk_constant_expression_new_for_value  = cc.DlFunc[func(value *gobject.GValue) *ConstantExpression, *ConstantExpression]{Name: "gtk_constant_expression_new_for_value"}
	gtk_constant_expression_get_value      = cc.DlFunc[func(expression *ConstantExpression) *gobject.GValue, *gobject.GValue]{Name: "gtk_constant_expression_get_value"}
	gtk_object_expression_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_object_expression_get_type"}
	gtk_object_expression_new              = cc.DlFunc[func(object *gobject.GObject) *ObjectExpression, *ObjectExpression]{Name: "gtk_object_expression_new"}
	gtk_object_expression_get_object       = cc.DlFunc[func(expression *ObjectExpression) *gobject.GObject, *gobject.GObject]{Name: "gtk_object_expression_get_object"}
	// #endregion

	// #region FileDialog
	gtk_file_dialog_get_type                        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_file_dialog_get_type"}
	gtk_file_dialog_new                             = cc.DlFunc[func() *FileDialog, *FileDialog]{Name: "gtk_file_dialog_new"}
	gtk_file_dialog_get_title                       = cc.DlFunc[func(fd *FileDialog) cc.String, cc.String]{Name: "gtk_file_dialog_get_title"}
	gtk_file_dialog_set_title                       = cc.DlFunc[func(fd *FileDialog, title cc.String), cc.Void]{Name: "gtk_file_dialog_set_title"}
	gtk_file_dialog_get_modal                       = cc.DlFunc[func(fd *FileDialog) int32, int32]{Name: "gtk_file_dialog_get_modal"}
	gtk_file_dialog_set_modal                       = cc.DlFunc[func(fd *FileDialog, modal int32), cc.Void]{Name: "gtk_file_dialog_set_modal"}
	gtk_file_dialog_get_filters                     = cc.DlFunc[func(fd *FileDialog) *gio.GListModel, *gio.GListModel]{Name: "gtk_file_dialog_get_filters"}
	gtk_file_dialog_set_filters                     = cc.DlFunc[func(fd *FileDialog, filters *gio.GListModel), cc.Void]{Name: "gtk_file_dialog_set_filters"}
	gtk_file_dialog_get_default_filter              = cc.DlFunc[func(fd *FileDialog) *FileFilter, *FileFilter]{Name: "gtk_file_dialog_get_default_filter"}
	gtk_file_dialog_set_default_filter              = cc.DlFunc[func(fd *FileDialog, filter *FileFilter), cc.Void]{Name: "gtk_file_dialog_set_default_filter"}
	gtk_file_dialog_get_initial_folder              = cc.DlFunc[func(fd *FileDialog) *gio.GFile, *gio.GFile]{Name: "gtk_file_dialog_get_initial_folder"}
	gtk_file_dialog_set_initial_folder              = cc.DlFunc[func(fd *FileDialog, folder *gio.GFile), cc.Void]{Name: "gtk_file_dialog_set_initial_folder"}
	gtk_file_dialog_get_initial_name                = cc.DlFunc[func(fd *FileDialog) cc.String, cc.String]{Name: "gtk_file_dialog_get_initial_name"}
	gtk_file_dialog_set_initial_name                = cc.DlFunc[func(fd *FileDialog, name cc.String), cc.Void]{Name: "gtk_file_dialog_set_initial_name"}
	gtk_file_dialog_get_initial_file                = cc.DlFunc[func(fd *FileDialog) *gio.GFile, *gio.GFile]{Name: "gtk_file_dialog_get_initial_file"}
	gtk_file_dialog_set_initial_file                = cc.DlFunc[func(fd *FileDialog, file *gio.GFile), cc.Void]{Name: "gtk_file_dialog_set_initial_file"}
	gtk_file_dialog_get_accept_label                = cc.DlFunc[func(fd *FileDialog) cc.String, cc.String]{Name: "gtk_file_dialog_get_accept_label"}
	gtk_file_dialog_set_accept_label                = cc.DlFunc[func(fd *FileDialog, acceptLabel cc.String), cc.Void]{Name: "gtk_file_dialog_set_accept_label"}
	gtk_file_dialog_open                            = cc.DlFunc[func(fd *FileDialog, parent *Window, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_file_dialog_open"}
	gtk_file_dialog_open_finish                     = cc.DlFunc[func(fd *FileDialog, result *gio.GAsyncResult, err **glib.GError) *gio.GFile, *gio.GFile]{Name: "gtk_file_dialog_open_finish"}
	gtk_file_dialog_select_folder                   = cc.DlFunc[func(fd *FileDialog, parent *Window, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_file_dialog_select_folder"}
	gtk_file_dialog_select_folder_finish            = cc.DlFunc[func(fd *FileDialog, result *gio.GAsyncResult, err **glib.GError) *gio.GFile, *gio.GFile]{Name: "gtk_file_dialog_select_folder_finish"}
	gtk_file_dialog_save                            = cc.DlFunc[func(fd *FileDialog, parent *Window, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_file_dialog_save"}
	gtk_file_dialog_save_finish                     = cc.DlFunc[func(fd *FileDialog, result *gio.GAsyncResult, err **glib.GError) *gio.GFile, *gio.GFile]{Name: "gtk_file_dialog_save_finish"}
	gtk_file_dialog_open_multiple                   = cc.DlFunc[func(fd *FileDialog, parent *Window, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_file_dialog_open_multiple"}
	gtk_file_dialog_open_multiple_finish            = cc.DlFunc[func(fd *FileDialog, result *gio.GAsyncResult, err **glib.GError) *gio.GListModel, *gio.GListModel]{Name: "gtk_file_dialog_open_multiple_finish"}
	gtk_file_dialog_select_multiple_folders         = cc.DlFunc[func(fd *FileDialog, parent *Window, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_file_dialog_select_multiple_folders"}
	gtk_file_dialog_select_multiple_folders_finish  = cc.DlFunc[func(fd *FileDialog, result *gio.GAsyncResult, err **glib.GError) *gio.GListModel, *gio.GListModel]{Name: "gtk_file_dialog_select_multiple_folders_finish"}
	gtk_file_dialog_open_text_file                  = cc.DlFunc[func(fd *FileDialog, parent *Window, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_file_dialog_open_text_file"}
	gtk_file_dialog_open_text_file_finish           = cc.DlFunc[func(fd *FileDialog, result *gio.GAsyncResult, encoding *cc.String, err **glib.GError) *gio.GFile, *gio.GFile]{Name: "gtk_file_dialog_open_text_file_finish"}
	gtk_file_dialog_open_multiple_text_files        = cc.DlFunc[func(fd *FileDialog, parent *Window, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_file_dialog_open_multiple_text_files"}
	gtk_file_dialog_open_multiple_text_files_finish = cc.DlFunc[func(fd *FileDialog, result *gio.GAsyncResult, encoding *cc.String, err **glib.GError) *gio.GListModel, *gio.GListModel]{Name: "gtk_file_dialog_open_multiple_text_files_finish"}
	gtk_file_dialog_save_text_file                  = cc.DlFunc[func(fd *FileDialog, parent *Window, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_file_dialog_save_text_file"}
	gtk_file_dialog_save_text_file_finish           = cc.DlFunc[func(fd *FileDialog, result *gio.GAsyncResult, encoding *cc.String, lineEnding *cc.String, err **glib.GError) *gio.GFile, *gio.GFile]{Name: "gtk_file_dialog_save_text_file_finish"}
	// #endregion

	// #region FileFilter
	gtk_file_filter_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_file_filter_get_type"}
	gtk_file_filter_new                = cc.DlFunc[func() *FileFilter, *FileFilter]{Name: "gtk_file_filter_new"}
	gtk_file_filter_set_name           = cc.DlFunc[func(f *FileFilter, name cc.String), cc.Void]{Name: "gtk_file_filter_set_name"}
	gtk_file_filter_get_name           = cc.DlFunc[func(f *FileFilter) cc.String, cc.String]{Name: "gtk_file_filter_get_name"}
	gtk_file_filter_add_mime_type      = cc.DlFunc[func(f *FileFilter, mimeType cc.String), cc.Void]{Name: "gtk_file_filter_add_mime_type"}
	gtk_file_filter_add_pattern        = cc.DlFunc[func(f *FileFilter, pattern cc.String), cc.Void]{Name: "gtk_file_filter_add_pattern"}
	gtk_file_filter_add_suffix         = cc.DlFunc[func(f *FileFilter, suffix cc.String), cc.Void]{Name: "gtk_file_filter_add_suffix"}
	gtk_file_filter_add_pixbuf_formats = cc.DlFunc[func(f *FileFilter), cc.Void]{Name: "gtk_file_filter_add_pixbuf_formats"}
	gtk_file_filter_get_attributes     = cc.DlFunc[func(f *FileFilter) cc.Strings, cc.Strings]{Name: "gtk_file_filter_get_attributes"}
	gtk_file_filter_to_gvariant        = cc.DlFunc[func(f *FileFilter) *glib.GVariant, *glib.GVariant]{Name: "gtk_file_filter_to_gvariant"}
	gtk_file_filter_new_from_gvariant  = cc.DlFunc[func(variant *glib.GVariant) *FileFilter, *FileFilter]{Name: "gtk_file_filter_new_from_gvariant"}
	// #endregion

	// #region FileLauncher
	gtk_file_launcher_get_type                      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_file_launcher_get_type"}
	gtk_file_launcher_new                           = cc.DlFunc[func(file *gio.GFile) *FileLauncher, *FileLauncher]{Name: "gtk_file_launcher_new"}
	gtk_file_launcher_get_file                      = cc.DlFunc[func(fl *FileLauncher) *gio.GFile, *gio.GFile]{Name: "gtk_file_launcher_get_file"}
	gtk_file_launcher_set_file                      = cc.DlFunc[func(fl *FileLauncher, file *gio.GFile), cc.Void]{Name: "gtk_file_launcher_set_file"}
	gtk_file_launcher_get_always_ask                = cc.DlFunc[func(fl *FileLauncher) int32, int32]{Name: "gtk_file_launcher_get_always_ask"}
	gtk_file_launcher_set_always_ask                = cc.DlFunc[func(fl *FileLauncher, alwaysAsk int32), cc.Void]{Name: "gtk_file_launcher_set_always_ask"}
	gtk_file_launcher_get_writable                  = cc.DlFunc[func(fl *FileLauncher) int32, int32]{Name: "gtk_file_launcher_get_writable"}
	gtk_file_launcher_set_writable                  = cc.DlFunc[func(fl *FileLauncher, writable int32), cc.Void]{Name: "gtk_file_launcher_set_writable"}
	gtk_file_launcher_launch                        = cc.DlFunc[func(fl *FileLauncher, parent *Window, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_file_launcher_launch"}
	gtk_file_launcher_launch_finish                 = cc.DlFunc[func(fl *FileLauncher, result *gio.GAsyncResult, err **glib.GError) int32, int32]{Name: "gtk_file_launcher_launch_finish"}
	gtk_file_launcher_open_containing_folder        = cc.DlFunc[func(fl *FileLauncher, parent *Window, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_file_launcher_open_containing_folder"}
	gtk_file_launcher_open_containing_folder_finish = cc.DlFunc[func(fl *FileLauncher, result *gio.GAsyncResult, err **glib.GError) int32, int32]{Name: "gtk_file_launcher_open_containing_folder_finish"}
	// #endregion

	// #region Filter
	gtk_filter_match          = cc.DlFunc[func(self *Filter, item *gobject.GObject) int32, int32]{Name: "gtk_filter_match"}
	gtk_filter_get_strictness = cc.DlFunc[func(self *Filter) FilterMatch, FilterMatch]{Name: "gtk_filter_get_strictness"}
	gtk_filter_changed        = cc.DlFunc[func(self *Filter, change FilterChange), cc.Void]{Name: "gtk_filter_changed"}
	// #endregion

	// #region FilterListModel
	gtk_filter_list_model_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_filter_list_model_get_type"}
	gtk_filter_list_model_new             = cc.DlFunc[func(model *gio.GListModel, filter *Filter) *FilterListModel, *FilterListModel]{Name: "gtk_filter_list_model_new"}
	gtk_filter_list_model_set_filter      = cc.DlFunc[func(flm *FilterListModel, filter *Filter), cc.Void]{Name: "gtk_filter_list_model_set_filter"}
	gtk_filter_list_model_get_filter      = cc.DlFunc[func(flm *FilterListModel) *Filter, *Filter]{Name: "gtk_filter_list_model_get_filter"}
	gtk_filter_list_model_set_model       = cc.DlFunc[func(flm *FilterListModel, model *gio.GListModel), cc.Void]{Name: "gtk_filter_list_model_set_model"}
	gtk_filter_list_model_get_model       = cc.DlFunc[func(flm *FilterListModel) *gio.GListModel, *gio.GListModel]{Name: "gtk_filter_list_model_get_model"}
	gtk_filter_list_model_set_incremental = cc.DlFunc[func(flm *FilterListModel, incremental int32), cc.Void]{Name: "gtk_filter_list_model_set_incremental"}
	gtk_filter_list_model_get_incremental = cc.DlFunc[func(flm *FilterListModel) int32, int32]{Name: "gtk_filter_list_model_get_incremental"}
	gtk_filter_list_model_get_pending     = cc.DlFunc[func(flm *FilterListModel) uint32, uint32]{Name: "gtk_filter_list_model_get_pending"}
	// #endregion

	// #region Fixed
	gtk_fixed_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_fixed_get_type"}
	gtk_fixed_new                 = cc.DlFunc[func() *Fixed, *Fixed]{Name: "gtk_fixed_new"}
	gtk_fixed_put                 = cc.DlFunc[func(f *Fixed, widget *Widget, x, y float64), cc.Void]{Name: "gtk_fixed_put"}
	gtk_fixed_remove              = cc.DlFunc[func(f *Fixed, widget *Widget), cc.Void]{Name: "gtk_fixed_remove"}
	gtk_fixed_move                = cc.DlFunc[func(f *Fixed, widget *Widget, x, y float64), cc.Void]{Name: "gtk_fixed_move"}
	gtk_fixed_get_child_position  = cc.DlFunc[func(f *Fixed, widget *Widget, x, y *float64), cc.Void]{Name: "gtk_fixed_get_child_position"}
	gtk_fixed_set_child_transform = cc.DlFunc[func(f *Fixed, widget *Widget, transform *gsk.Transform), cc.Void]{Name: "gtk_fixed_set_child_transform"}
	gtk_fixed_get_child_transform = cc.DlFunc[func(f *Fixed, widget *Widget) *gsk.Transform, *gsk.Transform]{Name: "gtk_fixed_get_child_transform"}
	// #endregion

	// #region FixedLayout
	gtk_fixed_layout_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_fixed_layout_get_type"}
	gtk_fixed_layout_new                 = cc.DlFunc[func() *FixedLayout, *FixedLayout]{Name: "gtk_fixed_layout_new"}
	gtk_fixed_layout_child_set_transform = cc.DlFunc[func(child *FixedLayoutChild, transform *gsk.Transform), cc.Void]{Name: "gtk_fixed_layout_child_set_transform"}
	gtk_fixed_layout_child_get_transform = cc.DlFunc[func(child *FixedLayoutChild) *gsk.Transform, *gsk.Transform]{Name: "gtk_fixed_layout_child_get_transform"}
	// #endregion

	// #region FlattenListModel
	gtk_flatten_list_model_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_flatten_list_model_get_type"}
	gtk_flatten_list_model_new                = cc.DlFunc[func(model *gio.GListModel) *FlattenListModel, *FlattenListModel]{Name: "gtk_flatten_list_model_new"}
	gtk_flatten_list_model_set_model          = cc.DlFunc[func(flm *FlattenListModel, model *gio.GListModel), cc.Void]{Name: "gtk_flatten_list_model_set_model"}
	gtk_flatten_list_model_get_model          = cc.DlFunc[func(flm *FlattenListModel) *gio.GListModel, *gio.GListModel]{Name: "gtk_flatten_list_model_get_model"}
	gtk_flatten_list_model_get_model_for_item = cc.DlFunc[func(flm *FlattenListModel, position uint32) *gio.GListModel, *gio.GListModel]{Name: "gtk_flatten_list_model_get_model_for_item"}
	// #endregion

	// #region FlowBox
	gtk_flow_box_child_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_flow_box_child_get_type"}
	gtk_flow_box_child_new                    = cc.DlFunc[func() *FlowBoxChild, *FlowBoxChild]{Name: "gtk_flow_box_child_new"}
	gtk_flow_box_child_set_child              = cc.DlFunc[func(f *FlowBoxChild, child *Widget), cc.Void]{Name: "gtk_flow_box_child_set_child"}
	gtk_flow_box_child_get_child              = cc.DlFunc[func(f *FlowBoxChild) *Widget, *Widget]{Name: "gtk_flow_box_child_get_child"}
	gtk_flow_box_child_get_index              = cc.DlFunc[func(f *FlowBoxChild) int32, int32]{Name: "gtk_flow_box_child_get_index"}
	gtk_flow_box_child_is_selected            = cc.DlFunc[func(f *FlowBoxChild) int32, int32]{Name: "gtk_flow_box_child_is_selected"}
	gtk_flow_box_child_changed                = cc.DlFunc[func(f *FlowBoxChild), cc.Void]{Name: "gtk_flow_box_child_changed"}
	gtk_flow_box_get_type                     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_flow_box_get_type"}
	gtk_flow_box_new                          = cc.DlFunc[func() *FlowBox, *FlowBox]{Name: "gtk_flow_box_new"}
	gtk_flow_box_bind_model                   = cc.DlFunc[func(f *FlowBox, model *gio.GListModel, createWidgetFunc uptr, userData uptr, userDataFreeFunc uptr), cc.Void]{Name: "gtk_flow_box_bind_model"}
	gtk_flow_box_set_homogeneous              = cc.DlFunc[func(f *FlowBox, homogeneous int32), cc.Void]{Name: "gtk_flow_box_set_homogeneous"}
	gtk_flow_box_get_homogeneous              = cc.DlFunc[func(f *FlowBox) int32, int32]{Name: "gtk_flow_box_get_homogeneous"}
	gtk_flow_box_set_row_spacing              = cc.DlFunc[func(f *FlowBox, spacing uint32), cc.Void]{Name: "gtk_flow_box_set_row_spacing"}
	gtk_flow_box_get_row_spacing              = cc.DlFunc[func(f *FlowBox) uint32, uint32]{Name: "gtk_flow_box_get_row_spacing"}
	gtk_flow_box_set_column_spacing           = cc.DlFunc[func(f *FlowBox, spacing uint32), cc.Void]{Name: "gtk_flow_box_set_column_spacing"}
	gtk_flow_box_get_column_spacing           = cc.DlFunc[func(f *FlowBox) uint32, uint32]{Name: "gtk_flow_box_get_column_spacing"}
	gtk_flow_box_set_min_children_per_line    = cc.DlFunc[func(f *FlowBox, nChildren uint32), cc.Void]{Name: "gtk_flow_box_set_min_children_per_line"}
	gtk_flow_box_get_min_children_per_line    = cc.DlFunc[func(f *FlowBox) uint32, uint32]{Name: "gtk_flow_box_get_min_children_per_line"}
	gtk_flow_box_set_max_children_per_line    = cc.DlFunc[func(f *FlowBox, nChildren uint32), cc.Void]{Name: "gtk_flow_box_set_max_children_per_line"}
	gtk_flow_box_get_max_children_per_line    = cc.DlFunc[func(f *FlowBox) uint32, uint32]{Name: "gtk_flow_box_get_max_children_per_line"}
	gtk_flow_box_set_activate_on_single_click = cc.DlFunc[func(f *FlowBox, single int32), cc.Void]{Name: "gtk_flow_box_set_activate_on_single_click"}
	gtk_flow_box_get_activate_on_single_click = cc.DlFunc[func(f *FlowBox) int32, int32]{Name: "gtk_flow_box_get_activate_on_single_click"}
	gtk_flow_box_prepend                      = cc.DlFunc[func(f *FlowBox, child *Widget), cc.Void]{Name: "gtk_flow_box_prepend"}
	gtk_flow_box_append                       = cc.DlFunc[func(f *FlowBox, child *Widget), cc.Void]{Name: "gtk_flow_box_append"}
	gtk_flow_box_insert                       = cc.DlFunc[func(f *FlowBox, widget *Widget, position int32), cc.Void]{Name: "gtk_flow_box_insert"}
	gtk_flow_box_remove                       = cc.DlFunc[func(f *FlowBox, widget *Widget), cc.Void]{Name: "gtk_flow_box_remove"}
	gtk_flow_box_remove_all                   = cc.DlFunc[func(f *FlowBox), cc.Void]{Name: "gtk_flow_box_remove_all"}
	gtk_flow_box_get_child_at_index           = cc.DlFunc[func(f *FlowBox, idx int32) *FlowBoxChild, *FlowBoxChild]{Name: "gtk_flow_box_get_child_at_index"}
	gtk_flow_box_get_child_at_pos             = cc.DlFunc[func(f *FlowBox, x, y int32) *FlowBoxChild, *FlowBoxChild]{Name: "gtk_flow_box_get_child_at_pos"}
	gtk_flow_box_selected_foreach             = cc.DlFunc[func(f *FlowBox, foreachFunc, data uptr), cc.Void]{Name: "gtk_flow_box_selected_foreach"}
	gtk_flow_box_get_selected_children        = cc.DlFunc[func(f *FlowBox) *glib.GList[Widget], *glib.GList[Widget]]{Name: "gtk_flow_box_get_selected_children"}
	gtk_flow_box_select_child                 = cc.DlFunc[func(f *FlowBox, child *FlowBoxChild), cc.Void]{Name: "gtk_flow_box_select_child"}
	gtk_flow_box_unselect_child               = cc.DlFunc[func(f *FlowBox, child *FlowBoxChild), cc.Void]{Name: "gtk_flow_box_unselect_child"}
	gtk_flow_box_select_all                   = cc.DlFunc[func(f *FlowBox), cc.Void]{Name: "gtk_flow_box_select_all"}
	gtk_flow_box_unselect_all                 = cc.DlFunc[func(f *FlowBox), cc.Void]{Name: "gtk_flow_box_unselect_all"}
	gtk_flow_box_set_selection_mode           = cc.DlFunc[func(f *FlowBox, mode SelectionMode), cc.Void]{Name: "gtk_flow_box_set_selection_mode"}
	gtk_flow_box_get_selection_mode           = cc.DlFunc[func(f *FlowBox) SelectionMode, SelectionMode]{Name: "gtk_flow_box_get_selection_mode"}
	gtk_flow_box_set_hadjustment              = cc.DlFunc[func(f *FlowBox, adjustment *Adjustment), cc.Void]{Name: "gtk_flow_box_set_hadjustment"}
	gtk_flow_box_set_vadjustment              = cc.DlFunc[func(f *FlowBox, adjustment *Adjustment), cc.Void]{Name: "gtk_flow_box_set_vadjustment"}
	gtk_flow_box_set_filter_func              = cc.DlFunc[func(f *FlowBox, filterFunc uptr, userData uptr, destroy uptr), cc.Void]{Name: "gtk_flow_box_set_filter_func"}
	gtk_flow_box_invalidate_filter            = cc.DlFunc[func(f *FlowBox), cc.Void]{Name: "gtk_flow_box_invalidate_filter"}
	gtk_flow_box_set_sort_func                = cc.DlFunc[func(f *FlowBox, sortFunc uptr, userData uptr, destroy uptr), cc.Void]{Name: "gtk_flow_box_set_sort_func"}
	gtk_flow_box_invalidate_sort              = cc.DlFunc[func(f *FlowBox), cc.Void]{Name: "gtk_flow_box_invalidate_sort"}
	// #endregion

	// #region FontDialog
	gtk_font_dialog_get_type                        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_font_dialog_get_type"}
	gtk_font_dialog_new                             = cc.DlFunc[func() *FontDialog, *FontDialog]{Name: "gtk_font_dialog_new"}
	gtk_font_dialog_get_title                       = cc.DlFunc[func(fd *FontDialog) cc.String, cc.String]{Name: "gtk_font_dialog_get_title"}
	gtk_font_dialog_set_title                       = cc.DlFunc[func(fd *FontDialog, title cc.String), cc.Void]{Name: "gtk_font_dialog_set_title"}
	gtk_font_dialog_get_modal                       = cc.DlFunc[func(fd *FontDialog) int32, int32]{Name: "gtk_font_dialog_get_modal"}
	gtk_font_dialog_set_modal                       = cc.DlFunc[func(fd *FontDialog, modal int32), cc.Void]{Name: "gtk_font_dialog_set_modal"}
	gtk_font_dialog_get_language                    = cc.DlFunc[func(fd *FontDialog) *pango.Language, *pango.Language]{Name: "gtk_font_dialog_get_language"}
	gtk_font_dialog_set_language                    = cc.DlFunc[func(fd *FontDialog, language *pango.Language), cc.Void]{Name: "gtk_font_dialog_set_language"}
	gtk_font_dialog_get_font_map                    = cc.DlFunc[func(fd *FontDialog) *pango.FontMap, *pango.FontMap]{Name: "gtk_font_dialog_get_font_map"}
	gtk_font_dialog_set_font_map                    = cc.DlFunc[func(fd *FontDialog, fontmap *pango.FontMap), cc.Void]{Name: "gtk_font_dialog_set_font_map"}
	gtk_font_dialog_get_filter                      = cc.DlFunc[func(fd *FontDialog) *Filter, *Filter]{Name: "gtk_font_dialog_get_filter"}
	gtk_font_dialog_set_filter                      = cc.DlFunc[func(fd *FontDialog, filter *Filter), cc.Void]{Name: "gtk_font_dialog_set_filter"}
	gtk_font_dialog_choose_family                   = cc.DlFunc[func(fd *FontDialog, parent *Window, initialValue *pango.FontFamily, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_font_dialog_choose_family"}
	gtk_font_dialog_choose_family_finish            = cc.DlFunc[func(fd *FontDialog, result *gio.GAsyncResult, err **glib.GError) *pango.FontFamily, *pango.FontFamily]{Name: "gtk_font_dialog_choose_family_finish"}
	gtk_font_dialog_choose_face                     = cc.DlFunc[func(fd *FontDialog, parent *Window, initialValue *pango.FontFace, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_font_dialog_choose_face"}
	gtk_font_dialog_choose_face_finish              = cc.DlFunc[func(fd *FontDialog, result *gio.GAsyncResult, err **glib.GError) *pango.FontFace, *pango.FontFace]{Name: "gtk_font_dialog_choose_face_finish"}
	gtk_font_dialog_choose_font                     = cc.DlFunc[func(fd *FontDialog, parent *Window, initialValue *pango.FontDescription, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_font_dialog_choose_font"}
	gtk_font_dialog_choose_font_finish              = cc.DlFunc[func(fd *FontDialog, result *gio.GAsyncResult, err **glib.GError) *pango.FontDescription, *pango.FontDescription]{Name: "gtk_font_dialog_choose_font_finish"}
	gtk_font_dialog_choose_font_and_features        = cc.DlFunc[func(fd *FontDialog, parent *Window, initialValue *pango.FontDescription, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_font_dialog_choose_font_and_features"}
	gtk_font_dialog_choose_font_and_features_finish = cc.DlFunc[func(fd *FontDialog, result *gio.GAsyncResult, fontDesc **pango.FontDescription, fontFeatures *cc.String, language **pango.Language, err **glib.GError) int32, int32]{Name: "gtk_font_dialog_choose_font_and_features_finish"}
	// #endregion

	// #region FontDialogButton
	gtk_font_dialog_button_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_font_dialog_button_get_type"}
	gtk_font_dialog_button_new               = cc.DlFunc[func(dialog *FontDialog) *FontDialogButton, *FontDialogButton]{Name: "gtk_font_dialog_button_new"}
	gtk_font_dialog_button_get_dialog        = cc.DlFunc[func(btn *FontDialogButton) *FontDialog, *FontDialog]{Name: "gtk_font_dialog_button_get_dialog"}
	gtk_font_dialog_button_set_dialog        = cc.DlFunc[func(btn *FontDialogButton, dialog *FontDialog), cc.Void]{Name: "gtk_font_dialog_button_set_dialog"}
	gtk_font_dialog_button_get_level         = cc.DlFunc[func(btn *FontDialogButton) FontLevel, FontLevel]{Name: "gtk_font_dialog_button_get_level"}
	gtk_font_dialog_button_set_level         = cc.DlFunc[func(btn *FontDialogButton, level FontLevel), cc.Void]{Name: "gtk_font_dialog_button_set_level"}
	gtk_font_dialog_button_get_font_desc     = cc.DlFunc[func(btn *FontDialogButton) *pango.FontDescription, *pango.FontDescription]{Name: "gtk_font_dialog_button_get_font_desc"}
	gtk_font_dialog_button_set_font_desc     = cc.DlFunc[func(btn *FontDialogButton, fontDesc *pango.FontDescription), cc.Void]{Name: "gtk_font_dialog_button_set_font_desc"}
	gtk_font_dialog_button_get_font_features = cc.DlFunc[func(btn *FontDialogButton) cc.String, cc.String]{Name: "gtk_font_dialog_button_get_font_features"}
	gtk_font_dialog_button_set_font_features = cc.DlFunc[func(btn *FontDialogButton, features cc.String), cc.Void]{Name: "gtk_font_dialog_button_set_font_features"}
	gtk_font_dialog_button_get_language      = cc.DlFunc[func(btn *FontDialogButton) *pango.Language, *pango.Language]{Name: "gtk_font_dialog_button_get_language"}
	gtk_font_dialog_button_set_language      = cc.DlFunc[func(btn *FontDialogButton, language *pango.Language), cc.Void]{Name: "gtk_font_dialog_button_set_language"}
	gtk_font_dialog_button_get_use_font      = cc.DlFunc[func(btn *FontDialogButton) int32, int32]{Name: "gtk_font_dialog_button_get_use_font"}
	gtk_font_dialog_button_set_use_font      = cc.DlFunc[func(btn *FontDialogButton, useFont int32), cc.Void]{Name: "gtk_font_dialog_button_set_use_font"}
	gtk_font_dialog_button_get_use_size      = cc.DlFunc[func(btn *FontDialogButton) int32, int32]{Name: "gtk_font_dialog_button_get_use_size"}
	gtk_font_dialog_button_set_use_size      = cc.DlFunc[func(btn *FontDialogButton, useSize int32), cc.Void]{Name: "gtk_font_dialog_button_set_use_size"}
	// #endregion

	// #region Frame
	gtk_frame_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_frame_get_type"}
	gtk_frame_new              = cc.DlFunc[func(label cc.String) *Frame, *Frame]{Name: "gtk_frame_new"}
	gtk_frame_set_label        = cc.DlFunc[func(frame *Frame, label cc.String), cc.Void]{Name: "gtk_frame_set_label"}
	gtk_frame_get_label        = cc.DlFunc[func(frame *Frame) cc.String, cc.String]{Name: "gtk_frame_get_label"}
	gtk_frame_set_label_widget = cc.DlFunc[func(frame *Frame, labelWidget *Widget), cc.Void]{Name: "gtk_frame_set_label_widget"}
	gtk_frame_get_label_widget = cc.DlFunc[func(frame *Frame) *Widget, *Widget]{Name: "gtk_frame_get_label_widget"}
	gtk_frame_set_label_align  = cc.DlFunc[func(frame *Frame, xalign float32), cc.Void]{Name: "gtk_frame_set_label_align"}
	gtk_frame_get_label_align  = cc.DlFunc[func(frame *Frame) float32, float32]{Name: "gtk_frame_get_label_align"}
	gtk_frame_set_child        = cc.DlFunc[func(frame *Frame, child *Widget), cc.Void]{Name: "gtk_frame_set_child"}
	gtk_frame_get_child        = cc.DlFunc[func(frame *Frame) *Widget, *Widget]{Name: "gtk_frame_get_child"}
	// #endregion

	// #region Gesture
	gtk_gesture_get_type                  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_gesture_get_type"}
	gtk_gesture_get_device                = cc.DlFunc[func(gesture *Gesture) *gdk.Device, *gdk.Device]{Name: "gtk_gesture_get_device"}
	gtk_gesture_set_state                 = cc.DlFunc[func(gesture *Gesture, state EventSequenceState) int32, int32]{Name: "gtk_gesture_set_state"}
	gtk_gesture_get_sequence_state        = cc.DlFunc[func(gesture *Gesture, sequence *gdk.EventSequence) EventSequenceState, EventSequenceState]{Name: "gtk_gesture_get_sequence_state"}
	gtk_gesture_set_sequence_state        = cc.DlFunc[func(gesture *Gesture, sequence *gdk.EventSequence, state EventSequenceState) int32, int32]{Name: "gtk_gesture_set_sequence_state"}
	gtk_gesture_get_sequences             = cc.DlFunc[func(gesture *Gesture) *glib.GList[gdk.EventSequence], *glib.GList[gdk.EventSequence]]{Name: "gtk_gesture_get_sequences"}
	gtk_gesture_get_last_updated_sequence = cc.DlFunc[func(gesture *Gesture) *gdk.EventSequence, *gdk.EventSequence]{Name: "gtk_gesture_get_last_updated_sequence"}
	gtk_gesture_handles_sequence          = cc.DlFunc[func(gesture *Gesture, sequence *gdk.EventSequence) int32, int32]{Name: "gtk_gesture_handles_sequence"}
	gtk_gesture_get_last_event            = cc.DlFunc[func(gesture *Gesture, sequence *gdk.EventSequence) *gdk.Event, *gdk.Event]{Name: "gtk_gesture_get_last_event"}
	gtk_gesture_get_point                 = cc.DlFunc[func(gesture *Gesture, sequence *gdk.EventSequence, x, y *float64) int32, int32]{Name: "gtk_gesture_get_point"}
	gtk_gesture_get_bounding_box          = cc.DlFunc[func(gesture *Gesture, rect *gdk.Rectangle) int32, int32]{Name: "gtk_gesture_get_bounding_box"}
	gtk_gesture_get_bounding_box_center   = cc.DlFunc[func(gesture *Gesture, x, y *float64) int32, int32]{Name: "gtk_gesture_get_bounding_box_center"}
	gtk_gesture_is_active                 = cc.DlFunc[func(gesture *Gesture) int32, int32]{Name: "gtk_gesture_is_active"}
	gtk_gesture_is_recognized             = cc.DlFunc[func(gesture *Gesture) int32, int32]{Name: "gtk_gesture_is_recognized"}
	gtk_gesture_group                     = cc.DlFunc[func(groupGesture, gesture *Gesture), cc.Void]{Name: "gtk_gesture_group"}
	gtk_gesture_ungroup                   = cc.DlFunc[func(gesture *Gesture), cc.Void]{Name: "gtk_gesture_ungroup"}
	gtk_gesture_get_group                 = cc.DlFunc[func(gesture *Gesture) *glib.GList[Gesture], *glib.GList[Gesture]]{Name: "gtk_gesture_get_group"}
	gtk_gesture_is_grouped_with           = cc.DlFunc[func(gesture, other *Gesture) int32, int32]{Name: "gtk_gesture_is_grouped_with"}
	// #endregion

	// #region GestureClick
	gtk_gesture_click_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_gesture_click_get_type"}
	gtk_gesture_click_new      = cc.DlFunc[func() *GestureClick, *GestureClick]{Name: "gtk_gesture_click_new"}
	// #endregion

	// #region GestureDrag
	gtk_gesture_drag_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_gesture_drag_get_type"}
	gtk_gesture_drag_new             = cc.DlFunc[func() *GestureDrag, *GestureDrag]{Name: "gtk_gesture_drag_new"}
	gtk_gesture_drag_get_start_point = cc.DlFunc[func(gesture *GestureDrag, x, y *float64) int32, int32]{Name: "gtk_gesture_drag_get_start_point"}
	gtk_gesture_drag_get_offset      = cc.DlFunc[func(gesture *GestureDrag, x, y *float64) int32, int32]{Name: "gtk_gesture_drag_get_offset"}
	// #endregion

	// #region GestureLongPress
	gtk_gesture_long_press_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_gesture_long_press_get_type"}
	gtk_gesture_long_press_new              = cc.DlFunc[func() *GestureLongPress, *GestureLongPress]{Name: "gtk_gesture_long_press_new"}
	gtk_gesture_long_press_set_delay_factor = cc.DlFunc[func(gesture *GestureLongPress, delayFactor float64), cc.Void]{Name: "gtk_gesture_long_press_set_delay_factor"}
	gtk_gesture_long_press_get_delay_factor = cc.DlFunc[func(gesture *GestureLongPress) float64, float64]{Name: "gtk_gesture_long_press_get_delay_factor"}
	// #endregion

	// #region GesturePan
	gtk_gesture_pan_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_gesture_pan_get_type"}
	gtk_gesture_pan_new             = cc.DlFunc[func(orientation Orientation) *GesturePan, *GesturePan]{Name: "gtk_gesture_pan_new"}
	gtk_gesture_pan_get_orientation = cc.DlFunc[func(g *GesturePan) Orientation, Orientation]{Name: "gtk_gesture_pan_get_orientation"}
	gtk_gesture_pan_set_orientation = cc.DlFunc[func(g *GesturePan, orientation Orientation), cc.Void]{Name: "gtk_gesture_pan_set_orientation"}
	// #endregion

	// #region GestureRotate
	gtk_gesture_rotate_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_gesture_rotate_get_type"}
	gtk_gesture_rotate_new             = cc.DlFunc[func() *GestureRotate, *GestureRotate]{Name: "gtk_gesture_rotate_new"}
	gtk_gesture_rotate_get_angle_delta = cc.DlFunc[func(gesture *GestureRotate) float64, float64]{Name: "gtk_gesture_rotate_get_angle_delta"}
	// #endregion

	// #region GestureSingle
	gtk_gesture_single_get_type             = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_gesture_single_get_type"}
	gtk_gesture_single_get_touch_only       = cc.DlFunc[func(gesture *GestureSingle) int32, int32]{Name: "gtk_gesture_single_get_touch_only"}
	gtk_gesture_single_set_touch_only       = cc.DlFunc[func(gesture *GestureSingle, touchOnly int32), cc.Void]{Name: "gtk_gesture_single_set_touch_only"}
	gtk_gesture_single_get_exclusive        = cc.DlFunc[func(gesture *GestureSingle) int32, int32]{Name: "gtk_gesture_single_get_exclusive"}
	gtk_gesture_single_set_exclusive        = cc.DlFunc[func(gesture *GestureSingle, exclusive int32), cc.Void]{Name: "gtk_gesture_single_set_exclusive"}
	gtk_gesture_single_get_button           = cc.DlFunc[func(gesture *GestureSingle) uint32, uint32]{Name: "gtk_gesture_single_get_button"}
	gtk_gesture_single_set_button           = cc.DlFunc[func(gesture *GestureSingle, button uint32), cc.Void]{Name: "gtk_gesture_single_set_button"}
	gtk_gesture_single_get_current_button   = cc.DlFunc[func(gesture *GestureSingle) uint32, uint32]{Name: "gtk_gesture_single_get_current_button"}
	gtk_gesture_single_get_current_sequence = cc.DlFunc[func(gesture *GestureSingle) *gdk.EventSequence, *gdk.EventSequence]{Name: "gtk_gesture_single_get_current_sequence"}
	// #endregion

	// #region GestureStylus
	gtk_gesture_stylus_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_gesture_stylus_get_type"}
	gtk_gesture_stylus_new             = cc.DlFunc[func() *GestureStylus, *GestureStylus]{Name: "gtk_gesture_stylus_new"}
	gtk_gesture_stylus_get_stylus_only = cc.DlFunc[func(gesture *GestureStylus) int32, int32]{Name: "gtk_gesture_stylus_get_stylus_only"}
	gtk_gesture_stylus_set_stylus_only = cc.DlFunc[func(gesture *GestureStylus, stylusOnly int32), cc.Void]{Name: "gtk_gesture_stylus_set_stylus_only"}
	gtk_gesture_stylus_get_axis        = cc.DlFunc[func(gesture *GestureStylus, axis gdk.AxisUse, value *float64) int32, int32]{Name: "gtk_gesture_stylus_get_axis"}
	gtk_gesture_stylus_get_axes        = cc.DlFunc[func(gesture *GestureStylus, axes *gdk.AxisUse, values **float64) int32, int32]{Name: "gtk_gesture_stylus_get_axes"}
	gtk_gesture_stylus_get_backlog     = cc.DlFunc[func(gesture *GestureStylus, backlog **gdk.TimeCoord, nElems *uint32) int32, int32]{Name: "gtk_gesture_stylus_get_backlog"}
	gtk_gesture_stylus_get_device_tool = cc.DlFunc[func(gesture *GestureStylus) *gdk.DeviceTool, *gdk.DeviceTool]{Name: "gtk_gesture_stylus_get_device_tool"}
	// #endregion

	// #region GestureSwipe
	gtk_gesture_swipe_get_type     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_gesture_swipe_get_type"}
	gtk_gesture_swipe_new          = cc.DlFunc[func() *GestureSwipe, *GestureSwipe]{Name: "gtk_gesture_swipe_new"}
	gtk_gesture_swipe_get_velocity = cc.DlFunc[func(gesture *GestureSwipe, velocityX, velocityY *float64) int32, int32]{Name: "gtk_gesture_swipe_get_velocity"}
	// #endregion

	// #region GestureZoom
	gtk_gesture_zoom_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_gesture_zoom_get_type"}
	gtk_gesture_zoom_new             = cc.DlFunc[func() *GestureZoom, *GestureZoom]{Name: "gtk_gesture_zoom_new"}
	gtk_gesture_zoom_get_scale_delta = cc.DlFunc[func(gesture *GestureZoom) float64, float64]{Name: "gtk_gesture_zoom_get_scale_delta"}
	// #endregion

	// #region GLArea
	gtk_gl_area_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_gl_area_get_type"}
	gtk_gl_area_new                    = cc.DlFunc[func() *GLArea, *GLArea]{Name: "gtk_gl_area_new"}
	gtk_gl_area_set_allowed_apis       = cc.DlFunc[func(area *GLArea, apis gdk.GLAPI), cc.Void]{Name: "gtk_gl_area_set_allowed_apis"}
	gtk_gl_area_get_allowed_apis       = cc.DlFunc[func(area *GLArea) gdk.GLAPI, gdk.GLAPI]{Name: "gtk_gl_area_get_allowed_apis"}
	gtk_gl_area_get_api                = cc.DlFunc[func(area *GLArea) gdk.GLAPI, gdk.GLAPI]{Name: "gtk_gl_area_get_api"}
	gtk_gl_area_set_required_version   = cc.DlFunc[func(area *GLArea, major int32, minor int32), cc.Void]{Name: "gtk_gl_area_set_required_version"}
	gtk_gl_area_get_required_version   = cc.DlFunc[func(area *GLArea, major *int32, minor *int32), cc.Void]{Name: "gtk_gl_area_get_required_version"}
	gtk_gl_area_get_has_depth_buffer   = cc.DlFunc[func(area *GLArea) int32, int32]{Name: "gtk_gl_area_get_has_depth_buffer"}
	gtk_gl_area_set_has_depth_buffer   = cc.DlFunc[func(area *GLArea, hasDepthBuffer int32), cc.Void]{Name: "gtk_gl_area_set_has_depth_buffer"}
	gtk_gl_area_get_has_stencil_buffer = cc.DlFunc[func(area *GLArea) int32, int32]{Name: "gtk_gl_area_get_has_stencil_buffer"}
	gtk_gl_area_set_has_stencil_buffer = cc.DlFunc[func(area *GLArea, hasStencilBuffer int32), cc.Void]{Name: "gtk_gl_area_set_has_stencil_buffer"}
	gtk_gl_area_get_auto_render        = cc.DlFunc[func(area *GLArea) int32, int32]{Name: "gtk_gl_area_get_auto_render"}
	gtk_gl_area_set_auto_render        = cc.DlFunc[func(area *GLArea, autoRender int32), cc.Void]{Name: "gtk_gl_area_set_auto_render"}
	gtk_gl_area_queue_render           = cc.DlFunc[func(area *GLArea), cc.Void]{Name: "gtk_gl_area_queue_render"}
	gtk_gl_area_get_context            = cc.DlFunc[func(area *GLArea) *gdk.GLContext, *gdk.GLContext]{Name: "gtk_gl_area_get_context"}
	gtk_gl_area_make_current           = cc.DlFunc[func(area *GLArea), cc.Void]{Name: "gtk_gl_area_make_current"}
	gtk_gl_area_attach_buffers         = cc.DlFunc[func(area *GLArea), cc.Void]{Name: "gtk_gl_area_attach_buffers"}
	gtk_gl_area_set_error              = cc.DlFunc[func(area *GLArea, err *glib.GError), cc.Void]{Name: "gtk_gl_area_set_error"}
	gtk_gl_area_get_error              = cc.DlFunc[func(area *GLArea) *glib.GError, *glib.GError]{Name: "gtk_gl_area_get_error"}
	// #endregion

	// #region GraphicsOffload
	gtk_graphics_offload_get_type             = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_graphics_offload_get_type"}
	gtk_graphics_offload_new                  = cc.DlFunc[func(child *Widget) *GraphicsOffload, *GraphicsOffload]{Name: "gtk_graphics_offload_new"}
	gtk_graphics_offload_set_child            = cc.DlFunc[func(goff *GraphicsOffload, child *Widget), cc.Void]{Name: "gtk_graphics_offload_set_child"}
	gtk_graphics_offload_get_child            = cc.DlFunc[func(goff *GraphicsOffload) *Widget, *Widget]{Name: "gtk_graphics_offload_get_child"}
	gtk_graphics_offload_set_enabled          = cc.DlFunc[func(goff *GraphicsOffload, enabled GraphicsOffloadEnabled), cc.Void]{Name: "gtk_graphics_offload_set_enabled"}
	gtk_graphics_offload_get_enabled          = cc.DlFunc[func(goff *GraphicsOffload) GraphicsOffloadEnabled, GraphicsOffloadEnabled]{Name: "gtk_graphics_offload_get_enabled"}
	gtk_graphics_offload_set_black_background = cc.DlFunc[func(goff *GraphicsOffload, value int32), cc.Void]{Name: "gtk_graphics_offload_set_black_background"}
	gtk_graphics_offload_get_black_background = cc.DlFunc[func(goff *GraphicsOffload) int32, int32]{Name: "gtk_graphics_offload_get_black_background"}
	// #endregion

	// #region Grid
	gtk_grid_get_type                  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_grid_get_type"}
	gtk_grid_new                       = cc.DlFunc[func() *Grid, *Grid]{Name: "gtk_grid_new"}
	gtk_grid_attach                    = cc.DlFunc[func(grid *Grid, child *Widget, column, row, width, height int32), cc.Void]{Name: "gtk_grid_attach"}
	gtk_grid_attach_next_to            = cc.DlFunc[func(grid *Grid, child *Widget, sibling *Widget, side PositionType, width, height int32), cc.Void]{Name: "gtk_grid_attach_next_to"}
	gtk_grid_get_child_at              = cc.DlFunc[func(grid *Grid, column, row int32) *Widget, *Widget]{Name: "gtk_grid_get_child_at"}
	gtk_grid_remove                    = cc.DlFunc[func(grid *Grid, child *Widget), cc.Void]{Name: "gtk_grid_remove"}
	gtk_grid_insert_row                = cc.DlFunc[func(grid *Grid, position int32), cc.Void]{Name: "gtk_grid_insert_row"}
	gtk_grid_insert_column             = cc.DlFunc[func(grid *Grid, position int32), cc.Void]{Name: "gtk_grid_insert_column"}
	gtk_grid_remove_row                = cc.DlFunc[func(grid *Grid, position int32), cc.Void]{Name: "gtk_grid_remove_row"}
	gtk_grid_remove_column             = cc.DlFunc[func(grid *Grid, position int32), cc.Void]{Name: "gtk_grid_remove_column"}
	gtk_grid_insert_next_to            = cc.DlFunc[func(grid *Grid, sibling *Widget, side PositionType), cc.Void]{Name: "gtk_grid_insert_next_to"}
	gtk_grid_set_row_homogeneous       = cc.DlFunc[func(grid *Grid, homogeneous int32), cc.Void]{Name: "gtk_grid_set_row_homogeneous"}
	gtk_grid_get_row_homogeneous       = cc.DlFunc[func(grid *Grid) int32, int32]{Name: "gtk_grid_get_row_homogeneous"}
	gtk_grid_set_row_spacing           = cc.DlFunc[func(grid *Grid, spacing uint32), cc.Void]{Name: "gtk_grid_set_row_spacing"}
	gtk_grid_get_row_spacing           = cc.DlFunc[func(grid *Grid) uint32, uint32]{Name: "gtk_grid_get_row_spacing"}
	gtk_grid_set_column_homogeneous    = cc.DlFunc[func(grid *Grid, homogeneous int32), cc.Void]{Name: "gtk_grid_set_column_homogeneous"}
	gtk_grid_get_column_homogeneous    = cc.DlFunc[func(grid *Grid) int32, int32]{Name: "gtk_grid_get_column_homogeneous"}
	gtk_grid_set_column_spacing        = cc.DlFunc[func(grid *Grid, spacing uint32), cc.Void]{Name: "gtk_grid_set_column_spacing"}
	gtk_grid_get_column_spacing        = cc.DlFunc[func(grid *Grid) uint32, uint32]{Name: "gtk_grid_get_column_spacing"}
	gtk_grid_set_row_baseline_position = cc.DlFunc[func(grid *Grid, row int32, pos BaselinePosition), cc.Void]{Name: "gtk_grid_set_row_baseline_position"}
	gtk_grid_get_row_baseline_position = cc.DlFunc[func(grid *Grid, row int32) BaselinePosition, BaselinePosition]{Name: "gtk_grid_get_row_baseline_position"}
	gtk_grid_set_baseline_row          = cc.DlFunc[func(grid *Grid, row int32), cc.Void]{Name: "gtk_grid_set_baseline_row"}
	gtk_grid_get_baseline_row          = cc.DlFunc[func(grid *Grid) int32, int32]{Name: "gtk_grid_get_baseline_row"}
	gtk_grid_query_child               = cc.DlFunc[func(grid *Grid, child *Widget, column, row, width, height *int32), cc.Void]{Name: "gtk_grid_query_child"}
	// #endregion

	// #region GridLayout
	gtk_grid_layout_get_type                  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_grid_layout_get_type"}
	gtk_grid_layout_new                       = cc.DlFunc[func() *GridLayout, *GridLayout]{Name: "gtk_grid_layout_new"}
	gtk_grid_layout_set_row_homogeneous       = cc.DlFunc[func(g *GridLayout, homogeneous int32), cc.Void]{Name: "gtk_grid_layout_set_row_homogeneous"}
	gtk_grid_layout_get_row_homogeneous       = cc.DlFunc[func(g *GridLayout) int32, int32]{Name: "gtk_grid_layout_get_row_homogeneous"}
	gtk_grid_layout_set_row_spacing           = cc.DlFunc[func(g *GridLayout, spacing uint32), cc.Void]{Name: "gtk_grid_layout_set_row_spacing"}
	gtk_grid_layout_get_row_spacing           = cc.DlFunc[func(g *GridLayout) uint32, uint32]{Name: "gtk_grid_layout_get_row_spacing"}
	gtk_grid_layout_set_column_homogeneous    = cc.DlFunc[func(g *GridLayout, homogeneous int32), cc.Void]{Name: "gtk_grid_layout_set_column_homogeneous"}
	gtk_grid_layout_get_column_homogeneous    = cc.DlFunc[func(g *GridLayout) int32, int32]{Name: "gtk_grid_layout_get_column_homogeneous"}
	gtk_grid_layout_set_column_spacing        = cc.DlFunc[func(g *GridLayout, spacing uint32), cc.Void]{Name: "gtk_grid_layout_set_column_spacing"}
	gtk_grid_layout_get_column_spacing        = cc.DlFunc[func(g *GridLayout) uint32, uint32]{Name: "gtk_grid_layout_get_column_spacing"}
	gtk_grid_layout_set_row_baseline_position = cc.DlFunc[func(g *GridLayout, row int32, pos BaselinePosition), cc.Void]{Name: "gtk_grid_layout_set_row_baseline_position"}
	gtk_grid_layout_get_row_baseline_position = cc.DlFunc[func(g *GridLayout, row int32) BaselinePosition, BaselinePosition]{Name: "gtk_grid_layout_get_row_baseline_position"}
	gtk_grid_layout_set_baseline_row          = cc.DlFunc[func(g *GridLayout, row int32), cc.Void]{Name: "gtk_grid_layout_set_baseline_row"}
	gtk_grid_layout_get_baseline_row          = cc.DlFunc[func(g *GridLayout) int32, int32]{Name: "gtk_grid_layout_get_baseline_row"}
	gtk_grid_layout_child_set_row             = cc.DlFunc[func(child *GridLayoutChild, row int32), cc.Void]{Name: "gtk_grid_layout_child_set_row"}
	gtk_grid_layout_child_get_row             = cc.DlFunc[func(child *GridLayoutChild) int32, int32]{Name: "gtk_grid_layout_child_get_row"}
	gtk_grid_layout_child_set_column          = cc.DlFunc[func(child *GridLayoutChild, column int32), cc.Void]{Name: "gtk_grid_layout_child_set_column"}
	gtk_grid_layout_child_get_column          = cc.DlFunc[func(child *GridLayoutChild) int32, int32]{Name: "gtk_grid_layout_child_get_column"}
	gtk_grid_layout_child_set_column_span     = cc.DlFunc[func(child *GridLayoutChild, span int32), cc.Void]{Name: "gtk_grid_layout_child_set_column_span"}
	gtk_grid_layout_child_get_column_span     = cc.DlFunc[func(child *GridLayoutChild) int32, int32]{Name: "gtk_grid_layout_child_get_column_span"}
	gtk_grid_layout_child_set_row_span        = cc.DlFunc[func(child *GridLayoutChild, span int32), cc.Void]{Name: "gtk_grid_layout_child_set_row_span"}
	gtk_grid_layout_child_get_row_span        = cc.DlFunc[func(child *GridLayoutChild) int32, int32]{Name: "gtk_grid_layout_child_get_row_span"}
	// #endregion

	// #region GridView
	gtk_grid_view_get_type                  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_grid_view_get_type"}
	gtk_grid_view_new                       = cc.DlFunc[func(model *SelectionModel, factory *ListItemFactory) *GridView, *GridView]{Name: "gtk_grid_view_new"}
	gtk_grid_view_get_model                 = cc.DlFunc[func(g *GridView) *SelectionModel, *SelectionModel]{Name: "gtk_grid_view_get_model"}
	gtk_grid_view_set_model                 = cc.DlFunc[func(g *GridView, model *SelectionModel), cc.Void]{Name: "gtk_grid_view_set_model"}
	gtk_grid_view_set_factory               = cc.DlFunc[func(g *GridView, factory *ListItemFactory), cc.Void]{Name: "gtk_grid_view_set_factory"}
	gtk_grid_view_get_factory               = cc.DlFunc[func(g *GridView) *ListItemFactory, *ListItemFactory]{Name: "gtk_grid_view_get_factory"}
	gtk_grid_view_get_min_columns           = cc.DlFunc[func(g *GridView) uint32, uint32]{Name: "gtk_grid_view_get_min_columns"}
	gtk_grid_view_set_min_columns           = cc.DlFunc[func(g *GridView, minColumns uint32), cc.Void]{Name: "gtk_grid_view_set_min_columns"}
	gtk_grid_view_get_max_columns           = cc.DlFunc[func(g *GridView) uint32, uint32]{Name: "gtk_grid_view_get_max_columns"}
	gtk_grid_view_set_max_columns           = cc.DlFunc[func(g *GridView, maxColumns uint32), cc.Void]{Name: "gtk_grid_view_set_max_columns"}
	gtk_grid_view_set_enable_rubberband     = cc.DlFunc[func(g *GridView, enableRubberband int32), cc.Void]{Name: "gtk_grid_view_set_enable_rubberband"}
	gtk_grid_view_get_enable_rubberband     = cc.DlFunc[func(g *GridView) int32, int32]{Name: "gtk_grid_view_get_enable_rubberband"}
	gtk_grid_view_set_tab_behavior          = cc.DlFunc[func(g *GridView, tabBehavior ListTabBehavior), cc.Void]{Name: "gtk_grid_view_set_tab_behavior"}
	gtk_grid_view_get_tab_behavior          = cc.DlFunc[func(g *GridView) ListTabBehavior, ListTabBehavior]{Name: "gtk_grid_view_get_tab_behavior"}
	gtk_grid_view_set_single_click_activate = cc.DlFunc[func(g *GridView, singleClickActivate int32), cc.Void]{Name: "gtk_grid_view_set_single_click_activate"}
	gtk_grid_view_get_single_click_activate = cc.DlFunc[func(g *GridView) int32, int32]{Name: "gtk_grid_view_get_single_click_activate"}
	gtk_grid_view_scroll_to                 = cc.DlFunc[func(g *GridView, pos uint32, flags ListScrollFlags, scroll *ScrollInfo), cc.Void]{Name: "gtk_grid_view_scroll_to"}
	// #endregion

	// #region HeaderBar
	gtk_header_bar_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_header_bar_get_type"}
	gtk_header_bar_new                     = cc.DlFunc[func() *HeaderBar, *HeaderBar]{Name: "gtk_header_bar_new"}
	gtk_header_bar_set_title_widget        = cc.DlFunc[func(bar *HeaderBar, titleWidget *Widget), cc.Void]{Name: "gtk_header_bar_set_title_widget"}
	gtk_header_bar_get_title_widget        = cc.DlFunc[func(bar *HeaderBar) *Widget, *Widget]{Name: "gtk_header_bar_get_title_widget"}
	gtk_header_bar_pack_start              = cc.DlFunc[func(bar *HeaderBar, child *Widget), cc.Void]{Name: "gtk_header_bar_pack_start"}
	gtk_header_bar_pack_end                = cc.DlFunc[func(bar *HeaderBar, child *Widget), cc.Void]{Name: "gtk_header_bar_pack_end"}
	gtk_header_bar_remove                  = cc.DlFunc[func(bar *HeaderBar, child *Widget), cc.Void]{Name: "gtk_header_bar_remove"}
	gtk_header_bar_get_show_title_buttons  = cc.DlFunc[func(bar *HeaderBar) int32, int32]{Name: "gtk_header_bar_get_show_title_buttons"}
	gtk_header_bar_set_show_title_buttons  = cc.DlFunc[func(bar *HeaderBar, setting int32), cc.Void]{Name: "gtk_header_bar_set_show_title_buttons"}
	gtk_header_bar_set_decoration_layout   = cc.DlFunc[func(bar *HeaderBar, layout cc.String), cc.Void]{Name: "gtk_header_bar_set_decoration_layout"}
	gtk_header_bar_get_decoration_layout   = cc.DlFunc[func(bar *HeaderBar) cc.String, cc.String]{Name: "gtk_header_bar_get_decoration_layout"}
	gtk_header_bar_get_use_native_controls = cc.DlFunc[func(bar *HeaderBar) int32, int32]{Name: "gtk_header_bar_get_use_native_controls"}
	gtk_header_bar_set_use_native_controls = cc.DlFunc[func(bar *HeaderBar, setting int32), cc.Void]{Name: "gtk_header_bar_set_use_native_controls"}
	// #endregion

	// #region IconTheme
	gtk_icon_theme_error_quark       = cc.DlFunc[func() glib.GQuark, glib.GQuark]{Name: "gtk_icon_theme_error_quark"}
	gtk_icon_theme_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_icon_theme_get_type"}
	gtk_icon_theme_new               = cc.DlFunc[func() *IconTheme, *IconTheme]{Name: "gtk_icon_theme_new"}
	gtk_icon_theme_get_for_display   = cc.DlFunc[func(display *gdk.Display) *IconTheme, *IconTheme]{Name: "gtk_icon_theme_get_for_display"}
	gtk_icon_theme_get_display       = cc.DlFunc[func(iconTheme *IconTheme) *gdk.Display, *gdk.Display]{Name: "gtk_icon_theme_get_display"}
	gtk_icon_theme_set_search_path   = cc.DlFunc[func(iconTheme *IconTheme, path cc.Strings), cc.Void]{Name: "gtk_icon_theme_set_search_path"}
	gtk_icon_theme_get_search_path   = cc.DlFunc[func(iconTheme *IconTheme) cc.Strings, cc.Strings]{Name: "gtk_icon_theme_get_search_path"}
	gtk_icon_theme_add_search_path   = cc.DlFunc[func(iconTheme *IconTheme, path cc.String), cc.Void]{Name: "gtk_icon_theme_add_search_path"}
	gtk_icon_theme_set_resource_path = cc.DlFunc[func(iconTheme *IconTheme, path cc.Strings), cc.Void]{Name: "gtk_icon_theme_set_resource_path"}
	gtk_icon_theme_get_resource_path = cc.DlFunc[func(iconTheme *IconTheme) cc.Strings, cc.Strings]{Name: "gtk_icon_theme_get_resource_path"}
	gtk_icon_theme_add_resource_path = cc.DlFunc[func(iconTheme *IconTheme, path cc.String), cc.Void]{Name: "gtk_icon_theme_add_resource_path"}
	gtk_icon_theme_set_theme_name    = cc.DlFunc[func(iconTheme *IconTheme, themeName cc.String), cc.Void]{Name: "gtk_icon_theme_set_theme_name"}
	gtk_icon_theme_get_theme_name    = cc.DlFunc[func(iconTheme *IconTheme) cc.String, cc.String]{Name: "gtk_icon_theme_get_theme_name"}
	gtk_icon_theme_has_icon          = cc.DlFunc[func(iconTheme *IconTheme, iconName cc.String) int32, int32]{Name: "gtk_icon_theme_has_icon"}
	gtk_icon_theme_has_gicon         = cc.DlFunc[func(iconTheme *IconTheme, gicon *gio.GIcon) int32, int32]{Name: "gtk_icon_theme_has_gicon"}
	gtk_icon_theme_get_icon_sizes    = cc.DlFunc[func(iconTheme *IconTheme, iconName cc.String) *int32, *int32]{Name: "gtk_icon_theme_get_icon_sizes"}
	gtk_icon_theme_lookup_icon       = cc.DlFunc[func(iconTheme *IconTheme, iconName cc.String, fallbacks cc.Strings, size, scale int32, direction TextDirection, flags IconLookupFlags) *IconPaintable, *IconPaintable]{Name: "gtk_icon_theme_lookup_icon"}
	gtk_icon_theme_lookup_by_gicon   = cc.DlFunc[func(iconTheme *IconTheme, gicon *gio.GIcon, size, scale int32, direction TextDirection, flags IconLookupFlags) *IconPaintable, *IconPaintable]{Name: "gtk_icon_theme_lookup_by_gicon"}
	gtk_icon_paintable_new_for_file  = cc.DlFunc[func(file *gio.GFile, size, scale int32) *IconPaintable, *IconPaintable]{Name: "gtk_icon_paintable_new_for_file"}
	gtk_icon_theme_get_icon_names    = cc.DlFunc[func(iconTheme *IconTheme) cc.Strings, cc.Strings]{Name: "gtk_icon_theme_get_icon_names"}
	gtk_icon_paintable_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_icon_paintable_get_type"}
	gtk_icon_paintable_get_file      = cc.DlFunc[func(iconPaintable *IconPaintable) *gio.GFile, *gio.GFile]{Name: "gtk_icon_paintable_get_file"}
	gtk_icon_paintable_get_icon_name = cc.DlFunc[func(iconPaintable *IconPaintable) cc.String, cc.String]{Name: "gtk_icon_paintable_get_icon_name"}
	gtk_icon_paintable_is_symbolic   = cc.DlFunc[func(iconPaintable *IconPaintable) int32, int32]{Name: "gtk_icon_paintable_is_symbolic"}
	// #endregion

	// #region Image
	gtk_image_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_image_get_type"}
	gtk_image_new               = cc.DlFunc[func() *Image, *Image]{Name: "gtk_image_new"}
	gtk_image_new_from_file     = cc.DlFunc[func(filename cc.String) *Image, *Image]{Name: "gtk_image_new_from_file"}
	gtk_image_new_from_resource = cc.DlFunc[func(resourcePath cc.String) *Image, *Image]{Name: "gtk_image_new_from_resource"}
	// gtk_image_new_from_pixbuf    func(pixbuf *gdk.Pixbuf) *Image
	gtk_image_new_from_paintable = cc.DlFunc[func(paintable *gdk.Paintable) *Image, *Image]{Name: "gtk_image_new_from_paintable"}
	gtk_image_new_from_icon_name = cc.DlFunc[func(iconName cc.String) *Image, *Image]{Name: "gtk_image_new_from_icon_name"}
	gtk_image_new_from_gicon     = cc.DlFunc[func(icon *gio.GIcon) *Image, *Image]{Name: "gtk_image_new_from_gicon"}
	gtk_image_clear              = cc.DlFunc[func(image *Image), cc.Void]{Name: "gtk_image_clear"}
	gtk_image_set_from_file      = cc.DlFunc[func(image *Image, filename cc.String), cc.Void]{Name: "gtk_image_set_from_file"}
	gtk_image_set_from_resource  = cc.DlFunc[func(image *Image, resourcePath cc.String), cc.Void]{Name: "gtk_image_set_from_resource"}
	// gtk_image_set_from_pixbuf    func(image *Image, pixbuf *gdk.Pixbuf)
	gtk_image_set_from_paintable = cc.DlFunc[func(image *Image, paintable *gdk.Paintable), cc.Void]{Name: "gtk_image_set_from_paintable"}
	gtk_image_set_from_icon_name = cc.DlFunc[func(image *Image, iconName cc.String), cc.Void]{Name: "gtk_image_set_from_icon_name"}
	gtk_image_set_from_gicon     = cc.DlFunc[func(image *Image, icon *gio.GIcon), cc.Void]{Name: "gtk_image_set_from_gicon"}
	gtk_image_set_pixel_size     = cc.DlFunc[func(image *Image, pixelSize int32), cc.Void]{Name: "gtk_image_set_pixel_size"}
	gtk_image_set_icon_size      = cc.DlFunc[func(image *Image, iconSize IconSize), cc.Void]{Name: "gtk_image_set_icon_size"}
	gtk_image_get_storage_type   = cc.DlFunc[func(image *Image) ImageType, ImageType]{Name: "gtk_image_get_storage_type"}
	gtk_image_get_paintable      = cc.DlFunc[func(image *Image) *gdk.Paintable, *gdk.Paintable]{Name: "gtk_image_get_paintable"}
	gtk_image_get_icon_name      = cc.DlFunc[func(image *Image) cc.String, cc.String]{Name: "gtk_image_get_icon_name"}
	gtk_image_get_gicon          = cc.DlFunc[func(image *Image) *gio.GIcon, *gio.GIcon]{Name: "gtk_image_get_gicon"}
	gtk_image_get_pixel_size     = cc.DlFunc[func(image *Image) int32, int32]{Name: "gtk_image_get_pixel_size"}
	gtk_image_get_icon_size      = cc.DlFunc[func(image *Image) IconSize, IconSize]{Name: "gtk_image_get_icon_size"}
	// #endregion

	// #region IMContext
	gtk_im_context_get_type                       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_im_context_get_type"}
	gtk_im_context_set_client_widget              = cc.DlFunc[func(context *IMContext, widget *Widget), cc.Void]{Name: "gtk_im_context_set_client_widget"}
	gtk_im_context_get_preedit_string             = cc.DlFunc[func(context *IMContext, str *cc.String, attrs **pango.AttrList, cursorPos *int32), cc.Void]{Name: "gtk_im_context_get_preedit_string"}
	gtk_im_context_filter_keypress                = cc.DlFunc[func(context *IMContext, event *gdk.Event) int32, int32]{Name: "gtk_im_context_filter_keypress"}
	gtk_im_context_filter_key                     = cc.DlFunc[func(context *IMContext, press int32, surface *gdk.Surface, device *gdk.Device, time uint32, keycode uint32, state gdk.ModifierType, group int32) int32, int32]{Name: "gtk_im_context_filter_key"}
	gtk_im_context_focus_in                       = cc.DlFunc[func(context *IMContext), cc.Void]{Name: "gtk_im_context_focus_in"}
	gtk_im_context_focus_out                      = cc.DlFunc[func(context *IMContext), cc.Void]{Name: "gtk_im_context_focus_out"}
	gtk_im_context_reset                          = cc.DlFunc[func(context *IMContext), cc.Void]{Name: "gtk_im_context_reset"}
	gtk_im_context_set_cursor_location            = cc.DlFunc[func(context *IMContext, area *gdk.Rectangle), cc.Void]{Name: "gtk_im_context_set_cursor_location"}
	gtk_im_context_set_use_preedit                = cc.DlFunc[func(context *IMContext, usePreedit int32), cc.Void]{Name: "gtk_im_context_set_use_preedit"}
	gtk_im_context_set_surrounding_with_selection = cc.DlFunc[func(context *IMContext, text cc.String, length, cursorIndex, anchorIndex int32), cc.Void]{Name: "gtk_im_context_set_surrounding_with_selection"}
	gtk_im_context_get_surrounding_with_selection = cc.DlFunc[func(context *IMContext, text *cc.String, cursorIndex, anchorIndex *int32) int32, int32]{Name: "gtk_im_context_get_surrounding_with_selection"}
	gtk_im_context_delete_surrounding             = cc.DlFunc[func(context *IMContext, offset, nChars int32) int32, int32]{Name: "gtk_im_context_delete_surrounding"}
	gtk_im_context_activate_osk                   = cc.DlFunc[func(context *IMContext, event *gdk.Event) int32, int32]{Name: "gtk_im_context_activate_osk"}
	// #endregion

	// #region IMContextSimple
	gtk_im_context_simple_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_im_context_simple_get_type"}
	gtk_im_context_simple_new              = cc.DlFunc[func() *IMContextSimple, *IMContextSimple]{Name: "gtk_im_context_simple_new"}
	gtk_im_context_simple_add_table        = cc.DlFunc[func(ctx *IMContextSimple, data *uint16, maxSeqLen int32, nSeqs int32), cc.Void]{Name: "gtk_im_context_simple_add_table"}
	gtk_im_context_simple_add_compose_file = cc.DlFunc[func(ctx *IMContextSimple, composeFile cc.String), cc.Void]{Name: "gtk_im_context_simple_add_compose_file"}
	// #endregion

	// #region IMMulticontext
	gtk_im_multicontext_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_im_multicontext_get_type"}
	gtk_im_multicontext_new            = cc.DlFunc[func() *IMMulticontext, *IMMulticontext]{Name: "gtk_im_multicontext_new"}
	gtk_im_multicontext_get_context_id = cc.DlFunc[func(ctx *IMMulticontext) cc.String, cc.String]{Name: "gtk_im_multicontext_get_context_id"}
	gtk_im_multicontext_set_context_id = cc.DlFunc[func(ctx *IMMulticontext, contextId cc.String), cc.Void]{Name: "gtk_im_multicontext_set_context_id"}
	// #endregion

	// #region Inscription
	gtk_inscription_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_inscription_get_type"}
	gtk_inscription_new               = cc.DlFunc[func(text cc.String) *Inscription, *Inscription]{Name: "gtk_inscription_new"}
	gtk_inscription_get_text          = cc.DlFunc[func(ins *Inscription) cc.String, cc.String]{Name: "gtk_inscription_get_text"}
	gtk_inscription_set_text          = cc.DlFunc[func(ins *Inscription, text cc.String), cc.Void]{Name: "gtk_inscription_set_text"}
	gtk_inscription_get_attributes    = cc.DlFunc[func(ins *Inscription) *pango.AttrList, *pango.AttrList]{Name: "gtk_inscription_get_attributes"}
	gtk_inscription_set_attributes    = cc.DlFunc[func(ins *Inscription, attrs *pango.AttrList), cc.Void]{Name: "gtk_inscription_set_attributes"}
	gtk_inscription_set_markup        = cc.DlFunc[func(ins *Inscription, markup cc.String), cc.Void]{Name: "gtk_inscription_set_markup"}
	gtk_inscription_get_text_overflow = cc.DlFunc[func(ins *Inscription) InscriptionOverflow, InscriptionOverflow]{Name: "gtk_inscription_get_text_overflow"}
	gtk_inscription_set_text_overflow = cc.DlFunc[func(ins *Inscription, overflow InscriptionOverflow), cc.Void]{Name: "gtk_inscription_set_text_overflow"}
	gtk_inscription_get_wrap_mode     = cc.DlFunc[func(ins *Inscription) pango.WrapMode, pango.WrapMode]{Name: "gtk_inscription_get_wrap_mode"}
	gtk_inscription_set_wrap_mode     = cc.DlFunc[func(ins *Inscription, wrapMode pango.WrapMode), cc.Void]{Name: "gtk_inscription_set_wrap_mode"}
	gtk_inscription_get_min_chars     = cc.DlFunc[func(ins *Inscription) uint32, uint32]{Name: "gtk_inscription_get_min_chars"}
	gtk_inscription_set_min_chars     = cc.DlFunc[func(ins *Inscription, minChars uint32), cc.Void]{Name: "gtk_inscription_set_min_chars"}
	gtk_inscription_get_nat_chars     = cc.DlFunc[func(ins *Inscription) uint32, uint32]{Name: "gtk_inscription_get_nat_chars"}
	gtk_inscription_set_nat_chars     = cc.DlFunc[func(ins *Inscription, natChars uint32), cc.Void]{Name: "gtk_inscription_set_nat_chars"}
	gtk_inscription_get_min_lines     = cc.DlFunc[func(ins *Inscription) uint32, uint32]{Name: "gtk_inscription_get_min_lines"}
	gtk_inscription_set_min_lines     = cc.DlFunc[func(ins *Inscription, minLines uint32), cc.Void]{Name: "gtk_inscription_set_min_lines"}
	gtk_inscription_get_nat_lines     = cc.DlFunc[func(ins *Inscription) uint32, uint32]{Name: "gtk_inscription_get_nat_lines"}
	gtk_inscription_set_nat_lines     = cc.DlFunc[func(ins *Inscription, natLines uint32), cc.Void]{Name: "gtk_inscription_set_nat_lines"}
	gtk_inscription_get_xalign        = cc.DlFunc[func(ins *Inscription) float32, float32]{Name: "gtk_inscription_get_xalign"}
	gtk_inscription_set_xalign        = cc.DlFunc[func(ins *Inscription, xalign float32), cc.Void]{Name: "gtk_inscription_set_xalign"}
	gtk_inscription_get_yalign        = cc.DlFunc[func(ins *Inscription) float32, float32]{Name: "gtk_inscription_get_yalign"}
	gtk_inscription_set_yalign        = cc.DlFunc[func(ins *Inscription, yalign float32), cc.Void]{Name: "gtk_inscription_set_yalign"}
	// #endregion

	// #region Label
	gtk_label_get_type                 = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_label_get_type"}
	gtk_label_new                      = cc.DlFunc[func(str cc.String) *Label, *Label]{Name: "gtk_label_new"}
	gtk_label_new_with_mnemonic        = cc.DlFunc[func(str cc.String) *Label, *Label]{Name: "gtk_label_new_with_mnemonic"}
	gtk_label_set_text                 = cc.DlFunc[func(label *Label, str cc.String), cc.Void]{Name: "gtk_label_set_text"}
	gtk_label_get_text                 = cc.DlFunc[func(label *Label) cc.String, cc.String]{Name: "gtk_label_get_text"}
	gtk_label_set_attributes           = cc.DlFunc[func(label *Label, attrs *pango.AttrList), cc.Void]{Name: "gtk_label_set_attributes"}
	gtk_label_get_attributes           = cc.DlFunc[func(label *Label) *pango.AttrList, *pango.AttrList]{Name: "gtk_label_get_attributes"}
	gtk_label_set_label                = cc.DlFunc[func(label *Label, str cc.String), cc.Void]{Name: "gtk_label_set_label"}
	gtk_label_get_label                = cc.DlFunc[func(label *Label) cc.String, cc.String]{Name: "gtk_label_get_label"}
	gtk_label_set_markup               = cc.DlFunc[func(label *Label, str cc.String), cc.Void]{Name: "gtk_label_set_markup"}
	gtk_label_set_use_markup           = cc.DlFunc[func(label *Label, setting int32), cc.Void]{Name: "gtk_label_set_use_markup"}
	gtk_label_get_use_markup           = cc.DlFunc[func(label *Label) int32, int32]{Name: "gtk_label_get_use_markup"}
	gtk_label_set_use_underline        = cc.DlFunc[func(label *Label, setting int32), cc.Void]{Name: "gtk_label_set_use_underline"}
	gtk_label_get_use_underline        = cc.DlFunc[func(label *Label) int32, int32]{Name: "gtk_label_get_use_underline"}
	gtk_label_set_markup_with_mnemonic = cc.DlFunc[func(label *Label, str cc.String), cc.Void]{Name: "gtk_label_set_markup_with_mnemonic"}
	gtk_label_get_mnemonic_keyval      = cc.DlFunc[func(label *Label) uint32, uint32]{Name: "gtk_label_get_mnemonic_keyval"}
	gtk_label_set_mnemonic_widget      = cc.DlFunc[func(label *Label, widget *Widget), cc.Void]{Name: "gtk_label_set_mnemonic_widget"}
	gtk_label_get_mnemonic_widget      = cc.DlFunc[func(label *Label) *Widget, *Widget]{Name: "gtk_label_get_mnemonic_widget"}
	gtk_label_set_text_with_mnemonic   = cc.DlFunc[func(label *Label, str cc.String), cc.Void]{Name: "gtk_label_set_text_with_mnemonic"}
	gtk_label_set_justify              = cc.DlFunc[func(label *Label, jtype Justification), cc.Void]{Name: "gtk_label_set_justify"}
	gtk_label_get_justify              = cc.DlFunc[func(label *Label) Justification, Justification]{Name: "gtk_label_get_justify"}
	gtk_label_set_ellipsize            = cc.DlFunc[func(label *Label, mode pango.EllipsizeMode), cc.Void]{Name: "gtk_label_set_ellipsize"}
	gtk_label_get_ellipsize            = cc.DlFunc[func(label *Label) pango.EllipsizeMode, pango.EllipsizeMode]{Name: "gtk_label_get_ellipsize"}
	gtk_label_set_width_chars          = cc.DlFunc[func(label *Label, nChars int32), cc.Void]{Name: "gtk_label_set_width_chars"}
	gtk_label_get_width_chars          = cc.DlFunc[func(label *Label) int32, int32]{Name: "gtk_label_get_width_chars"}
	gtk_label_set_max_width_chars      = cc.DlFunc[func(label *Label, nChars int32), cc.Void]{Name: "gtk_label_set_max_width_chars"}
	gtk_label_get_max_width_chars      = cc.DlFunc[func(label *Label) int32, int32]{Name: "gtk_label_get_max_width_chars"}
	gtk_label_set_lines                = cc.DlFunc[func(label *Label, lines int32), cc.Void]{Name: "gtk_label_set_lines"}
	gtk_label_get_lines                = cc.DlFunc[func(label *Label) int32, int32]{Name: "gtk_label_get_lines"}
	gtk_label_set_wrap                 = cc.DlFunc[func(label *Label, wrap int32), cc.Void]{Name: "gtk_label_set_wrap"}
	gtk_label_get_wrap                 = cc.DlFunc[func(label *Label) int32, int32]{Name: "gtk_label_get_wrap"}
	gtk_label_set_wrap_mode            = cc.DlFunc[func(label *Label, wrapMode pango.WrapMode), cc.Void]{Name: "gtk_label_set_wrap_mode"}
	gtk_label_get_wrap_mode            = cc.DlFunc[func(label *Label) pango.WrapMode, pango.WrapMode]{Name: "gtk_label_get_wrap_mode"}
	gtk_label_set_natural_wrap_mode    = cc.DlFunc[func(label *Label, wrapMode NaturalWrapMode), cc.Void]{Name: "gtk_label_set_natural_wrap_mode"}
	gtk_label_get_natural_wrap_mode    = cc.DlFunc[func(label *Label) NaturalWrapMode, NaturalWrapMode]{Name: "gtk_label_get_natural_wrap_mode"}
	gtk_label_set_selectable           = cc.DlFunc[func(label *Label, setting int32), cc.Void]{Name: "gtk_label_set_selectable"}
	gtk_label_get_selectable           = cc.DlFunc[func(label *Label) int32, int32]{Name: "gtk_label_get_selectable"}
	gtk_label_select_region            = cc.DlFunc[func(label *Label, startOffset, endOffset int32), cc.Void]{Name: "gtk_label_select_region"}
	gtk_label_get_selection_bounds     = cc.DlFunc[func(label *Label, start, end *int32) int32, int32]{Name: "gtk_label_get_selection_bounds"}
	gtk_label_get_layout               = cc.DlFunc[func(label *Label) *pango.Layout, *pango.Layout]{Name: "gtk_label_get_layout"}
	gtk_label_get_layout_offsets       = cc.DlFunc[func(label *Label, x, y *int32), cc.Void]{Name: "gtk_label_get_layout_offsets"}
	gtk_label_set_single_line_mode     = cc.DlFunc[func(label *Label, singleLineMode int32), cc.Void]{Name: "gtk_label_set_single_line_mode"}
	gtk_label_get_single_line_mode     = cc.DlFunc[func(label *Label) int32, int32]{Name: "gtk_label_get_single_line_mode"}
	gtk_label_get_current_uri          = cc.DlFunc[func(label *Label) cc.String, cc.String]{Name: "gtk_label_get_current_uri"}
	gtk_label_set_xalign               = cc.DlFunc[func(label *Label, xalign float32), cc.Void]{Name: "gtk_label_set_xalign"}
	gtk_label_get_xalign               = cc.DlFunc[func(label *Label) float32, float32]{Name: "gtk_label_get_xalign"}
	gtk_label_set_yalign               = cc.DlFunc[func(label *Label, yalign float32), cc.Void]{Name: "gtk_label_set_yalign"}
	gtk_label_get_yalign               = cc.DlFunc[func(label *Label) float32, float32]{Name: "gtk_label_get_yalign"}
	gtk_label_set_extra_menu           = cc.DlFunc[func(label *Label, model *gio.GMenuModel), cc.Void]{Name: "gtk_label_set_extra_menu"}
	gtk_label_get_extra_menu           = cc.DlFunc[func(label *Label) *gio.GMenuModel, *gio.GMenuModel]{Name: "gtk_label_get_extra_menu"}
	gtk_label_set_tabs                 = cc.DlFunc[func(label *Label, tabs *pango.TabArray), cc.Void]{Name: "gtk_label_set_tabs"}
	gtk_label_get_tabs                 = cc.DlFunc[func(label *Label) *pango.TabArray, *pango.TabArray]{Name: "gtk_label_get_tabs"}
	// #endregion

	// #region LayoutChild
	gtk_layout_child_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_layout_child_get_type"}
	gtk_layout_child_get_layout_manager = cc.DlFunc[func(*LayoutChild) *LayoutManager, *LayoutManager]{Name: "gtk_layout_child_get_layout_manager"}
	gtk_layout_child_get_child_widget   = cc.DlFunc[func(*LayoutChild) *Widget, *Widget]{Name: "gtk_layout_child_get_child_widget"}
	// #endregion

	// #region LayoutManager
	gtk_layout_manager_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_layout_manager_get_type"}
	gtk_layout_manager_measure          = cc.DlFunc[func(*LayoutManager, *Widget, Orientation, int, *int, *int, *int, *int), cc.Void]{Name: "gtk_layout_manager_measure"}
	gtk_layout_manager_allocate         = cc.DlFunc[func(*LayoutManager, *Widget, int, int, int), cc.Void]{Name: "gtk_layout_manager_allocate"}
	gtk_layout_manager_get_request_mode = cc.DlFunc[func(*LayoutManager) SizeRequestMode, SizeRequestMode]{Name: "gtk_layout_manager_get_request_mode"}
	gtk_layout_manager_get_widget       = cc.DlFunc[func(*LayoutManager) *Widget, *Widget]{Name: "gtk_layout_manager_get_widget"}
	gtk_layout_manager_layout_changed   = cc.DlFunc[func(*LayoutManager), cc.Void]{Name: "gtk_layout_manager_layout_changed"}
	gtk_layout_manager_get_layout_child = cc.DlFunc[func(*LayoutManager, *Widget) *LayoutChild, *LayoutChild]{Name: "gtk_layout_manager_get_layout_child"}
	// #endregion

	// #region LevelBar
	gtk_level_bar_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_level_bar_get_type"}
	gtk_level_bar_new                 = cc.DlFunc[func() *LevelBar, *LevelBar]{Name: "gtk_level_bar_new"}
	gtk_level_bar_new_for_interval    = cc.DlFunc[func(minValue, maxValue float64) *LevelBar, *LevelBar]{Name: "gtk_level_bar_new_for_interval"}
	gtk_level_bar_set_mode            = cc.DlFunc[func(l *LevelBar, mode LevelBarMode), cc.Void]{Name: "gtk_level_bar_set_mode"}
	gtk_level_bar_get_mode            = cc.DlFunc[func(l *LevelBar) LevelBarMode, LevelBarMode]{Name: "gtk_level_bar_get_mode"}
	gtk_level_bar_set_value           = cc.DlFunc[func(l *LevelBar, value float64), cc.Void]{Name: "gtk_level_bar_set_value"}
	gtk_level_bar_get_value           = cc.DlFunc[func(l *LevelBar) float64, float64]{Name: "gtk_level_bar_get_value"}
	gtk_level_bar_set_min_value       = cc.DlFunc[func(l *LevelBar, value float64), cc.Void]{Name: "gtk_level_bar_set_min_value"}
	gtk_level_bar_get_min_value       = cc.DlFunc[func(l *LevelBar) float64, float64]{Name: "gtk_level_bar_get_min_value"}
	gtk_level_bar_set_max_value       = cc.DlFunc[func(l *LevelBar, value float64), cc.Void]{Name: "gtk_level_bar_set_max_value"}
	gtk_level_bar_get_max_value       = cc.DlFunc[func(l *LevelBar) float64, float64]{Name: "gtk_level_bar_get_max_value"}
	gtk_level_bar_set_inverted        = cc.DlFunc[func(l *LevelBar, inverted int32), cc.Void]{Name: "gtk_level_bar_set_inverted"}
	gtk_level_bar_get_inverted        = cc.DlFunc[func(l *LevelBar) int32, int32]{Name: "gtk_level_bar_get_inverted"}
	gtk_level_bar_add_offset_value    = cc.DlFunc[func(l *LevelBar, name cc.String, value float64), cc.Void]{Name: "gtk_level_bar_add_offset_value"}
	gtk_level_bar_remove_offset_value = cc.DlFunc[func(l *LevelBar, name cc.String), cc.Void]{Name: "gtk_level_bar_remove_offset_value"}
	gtk_level_bar_get_offset_value    = cc.DlFunc[func(l *LevelBar, name cc.String, value *float64) int32, int32]{Name: "gtk_level_bar_get_offset_value"}
	// #endregion

	// #region LinkButton
	gtk_link_button_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_link_button_get_type"}
	gtk_link_button_new            = cc.DlFunc[func(uri cc.String) *LinkButton, *LinkButton]{Name: "gtk_link_button_new"}
	gtk_link_button_new_with_label = cc.DlFunc[func(uri cc.String, label cc.String) *LinkButton, *LinkButton]{Name: "gtk_link_button_new_with_label"}
	gtk_link_button_get_uri        = cc.DlFunc[func(lb *LinkButton) cc.String, cc.String]{Name: "gtk_link_button_get_uri"}
	gtk_link_button_set_uri        = cc.DlFunc[func(lb *LinkButton, uri cc.String), cc.Void]{Name: "gtk_link_button_set_uri"}
	gtk_link_button_get_visited    = cc.DlFunc[func(lb *LinkButton) int32, int32]{Name: "gtk_link_button_get_visited"}
	gtk_link_button_set_visited    = cc.DlFunc[func(lb *LinkButton, visited int32), cc.Void]{Name: "gtk_link_button_set_visited"}
	// #endregion

	// #region ListBase
	gtk_list_base_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_list_base_get_type"}
	// #endregion

	// #region ListBox
	gtk_list_box_row_get_type                 = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_list_box_row_get_type"}
	gtk_list_box_row_new                      = cc.DlFunc[func() *ListBoxRow, *ListBoxRow]{Name: "gtk_list_box_row_new"}
	gtk_list_box_row_set_child                = cc.DlFunc[func(row *ListBoxRow, child *Widget), cc.Void]{Name: "gtk_list_box_row_set_child"}
	gtk_list_box_row_get_child                = cc.DlFunc[func(row *ListBoxRow) *Widget, *Widget]{Name: "gtk_list_box_row_get_child"}
	gtk_list_box_row_get_header               = cc.DlFunc[func(row *ListBoxRow) *Widget, *Widget]{Name: "gtk_list_box_row_get_header"}
	gtk_list_box_row_set_header               = cc.DlFunc[func(row *ListBoxRow, header *Widget), cc.Void]{Name: "gtk_list_box_row_set_header"}
	gtk_list_box_row_get_index                = cc.DlFunc[func(row *ListBoxRow) int32, int32]{Name: "gtk_list_box_row_get_index"}
	gtk_list_box_row_changed                  = cc.DlFunc[func(row *ListBoxRow), cc.Void]{Name: "gtk_list_box_row_changed"}
	gtk_list_box_row_is_selected              = cc.DlFunc[func(row *ListBoxRow) int32, int32]{Name: "gtk_list_box_row_is_selected"}
	gtk_list_box_row_set_selectable           = cc.DlFunc[func(row *ListBoxRow, selectable int32), cc.Void]{Name: "gtk_list_box_row_set_selectable"}
	gtk_list_box_row_get_selectable           = cc.DlFunc[func(row *ListBoxRow) int32, int32]{Name: "gtk_list_box_row_get_selectable"}
	gtk_list_box_row_set_activatable          = cc.DlFunc[func(row *ListBoxRow, activatable int32), cc.Void]{Name: "gtk_list_box_row_set_activatable"}
	gtk_list_box_row_get_activatable          = cc.DlFunc[func(row *ListBoxRow) int32, int32]{Name: "gtk_list_box_row_get_activatable"}
	gtk_list_box_get_type                     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_list_box_get_type"}
	gtk_list_box_prepend                      = cc.DlFunc[func(box *ListBox, child *Widget), cc.Void]{Name: "gtk_list_box_prepend"}
	gtk_list_box_append                       = cc.DlFunc[func(box *ListBox, child *Widget), cc.Void]{Name: "gtk_list_box_append"}
	gtk_list_box_insert                       = cc.DlFunc[func(box *ListBox, child *Widget, position int32), cc.Void]{Name: "gtk_list_box_insert"}
	gtk_list_box_remove                       = cc.DlFunc[func(box *ListBox, child *Widget), cc.Void]{Name: "gtk_list_box_remove"}
	gtk_list_box_remove_all                   = cc.DlFunc[func(box *ListBox), cc.Void]{Name: "gtk_list_box_remove_all"}
	gtk_list_box_get_selected_row             = cc.DlFunc[func(box *ListBox) *ListBoxRow, *ListBoxRow]{Name: "gtk_list_box_get_selected_row"}
	gtk_list_box_get_row_at_index             = cc.DlFunc[func(box *ListBox, index int32) *ListBoxRow, *ListBoxRow]{Name: "gtk_list_box_get_row_at_index"}
	gtk_list_box_get_row_at_y                 = cc.DlFunc[func(box *ListBox, y int32) *ListBoxRow, *ListBoxRow]{Name: "gtk_list_box_get_row_at_y"}
	gtk_list_box_select_row                   = cc.DlFunc[func(box *ListBox, row *ListBoxRow), cc.Void]{Name: "gtk_list_box_select_row"}
	gtk_list_box_set_placeholder              = cc.DlFunc[func(box *ListBox, placeholder *Widget), cc.Void]{Name: "gtk_list_box_set_placeholder"}
	gtk_list_box_set_adjustment               = cc.DlFunc[func(box *ListBox, adjustment *Adjustment), cc.Void]{Name: "gtk_list_box_set_adjustment"}
	gtk_list_box_get_adjustment               = cc.DlFunc[func(box *ListBox) *Adjustment, *Adjustment]{Name: "gtk_list_box_get_adjustment"}
	gtk_list_box_selected_foreach             = cc.DlFunc[func(box *ListBox, fn, data uptr), cc.Void]{Name: "gtk_list_box_selected_foreach"}
	gtk_list_box_get_selected_rows            = cc.DlFunc[func(box *ListBox) *glib.GList[ListBoxRow], *glib.GList[ListBoxRow]]{Name: "gtk_list_box_get_selected_rows"}
	gtk_list_box_unselect_row                 = cc.DlFunc[func(box *ListBox, row *ListBoxRow), cc.Void]{Name: "gtk_list_box_unselect_row"}
	gtk_list_box_select_all                   = cc.DlFunc[func(box *ListBox), cc.Void]{Name: "gtk_list_box_select_all"}
	gtk_list_box_unselect_all                 = cc.DlFunc[func(box *ListBox), cc.Void]{Name: "gtk_list_box_unselect_all"}
	gtk_list_box_set_selection_mode           = cc.DlFunc[func(box *ListBox, mode SelectionMode), cc.Void]{Name: "gtk_list_box_set_selection_mode"}
	gtk_list_box_get_selection_mode           = cc.DlFunc[func(box *ListBox) SelectionMode, SelectionMode]{Name: "gtk_list_box_get_selection_mode"}
	gtk_list_box_set_filter_func              = cc.DlFunc[func(box *ListBox, filterFunc uptr, userData uptr, destroy uptr), cc.Void]{Name: "gtk_list_box_set_filter_func"}
	gtk_list_box_set_header_func              = cc.DlFunc[func(box *ListBox, updateHeader uptr, userData uptr, destroy uptr), cc.Void]{Name: "gtk_list_box_set_header_func"}
	gtk_list_box_invalidate_filter            = cc.DlFunc[func(box *ListBox), cc.Void]{Name: "gtk_list_box_invalidate_filter"}
	gtk_list_box_invalidate_sort              = cc.DlFunc[func(box *ListBox), cc.Void]{Name: "gtk_list_box_invalidate_sort"}
	gtk_list_box_invalidate_headers           = cc.DlFunc[func(box *ListBox), cc.Void]{Name: "gtk_list_box_invalidate_headers"}
	gtk_list_box_set_sort_func                = cc.DlFunc[func(box *ListBox, sortFunc uptr, userData uptr, destroy uptr), cc.Void]{Name: "gtk_list_box_set_sort_func"}
	gtk_list_box_set_activate_on_single_click = cc.DlFunc[func(box *ListBox, single int32), cc.Void]{Name: "gtk_list_box_set_activate_on_single_click"}
	gtk_list_box_get_activate_on_single_click = cc.DlFunc[func(box *ListBox) int32, int32]{Name: "gtk_list_box_get_activate_on_single_click"}
	gtk_list_box_drag_unhighlight_row         = cc.DlFunc[func(box *ListBox), cc.Void]{Name: "gtk_list_box_drag_unhighlight_row"}
	gtk_list_box_drag_highlight_row           = cc.DlFunc[func(box *ListBox, row *ListBoxRow), cc.Void]{Name: "gtk_list_box_drag_highlight_row"}
	gtk_list_box_new                          = cc.DlFunc[func() *ListBox, *ListBox]{Name: "gtk_list_box_new"}
	gtk_list_box_bind_model                   = cc.DlFunc[func(box *ListBox, model *gio.GListModel, createWidgetFunc uptr, userData uptr, userDataFreeFunc uptr), cc.Void]{Name: "gtk_list_box_bind_model"}
	gtk_list_box_set_show_separators          = cc.DlFunc[func(box *ListBox, showSeparators int32), cc.Void]{Name: "gtk_list_box_set_show_separators"}
	gtk_list_box_get_show_separators          = cc.DlFunc[func(box *ListBox) int32, int32]{Name: "gtk_list_box_get_show_separators"}
	gtk_list_box_set_tab_behavior             = cc.DlFunc[func(box *ListBox, behavior ListTabBehavior), cc.Void]{Name: "gtk_list_box_set_tab_behavior"}
	gtk_list_box_get_tab_behavior             = cc.DlFunc[func(box *ListBox) ListTabBehavior, ListTabBehavior]{Name: "gtk_list_box_get_tab_behavior"}
	// #endregion

	// #region ListHeader
	gtk_list_header_get_type    = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_list_header_get_type"}
	gtk_list_header_get_item    = cc.DlFunc[func(lh *ListHeader) *gobject.GObject, *gobject.GObject]{Name: "gtk_list_header_get_item"}
	gtk_list_header_get_start   = cc.DlFunc[func(lh *ListHeader) uint32, uint32]{Name: "gtk_list_header_get_start"}
	gtk_list_header_get_end     = cc.DlFunc[func(lh *ListHeader) uint32, uint32]{Name: "gtk_list_header_get_end"}
	gtk_list_header_get_n_items = cc.DlFunc[func(lh *ListHeader) uint32, uint32]{Name: "gtk_list_header_get_n_items"}
	gtk_list_header_set_child   = cc.DlFunc[func(lh *ListHeader, child *Widget), cc.Void]{Name: "gtk_list_header_set_child"}
	gtk_list_header_get_child   = cc.DlFunc[func(lh *ListHeader) *Widget, *Widget]{Name: "gtk_list_header_get_child"}
	// #endregion

	// #region ListItem
	gtk_list_item_get_type                   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_list_item_get_type"}
	gtk_list_item_get_item                   = cc.DlFunc[func(*ListItem) uptr, uptr]{Name: "gtk_list_item_get_item"}
	gtk_list_item_get_position               = cc.DlFunc[func(*ListItem) uint32, uint32]{Name: "gtk_list_item_get_position"}
	gtk_list_item_get_selected               = cc.DlFunc[func(*ListItem) int32, int32]{Name: "gtk_list_item_get_selected"}
	gtk_list_item_get_selectable             = cc.DlFunc[func(*ListItem) int32, int32]{Name: "gtk_list_item_get_selectable"}
	gtk_list_item_set_selectable             = cc.DlFunc[func(*ListItem, int32), cc.Void]{Name: "gtk_list_item_set_selectable"}
	gtk_list_item_get_activatable            = cc.DlFunc[func(*ListItem) int32, int32]{Name: "gtk_list_item_get_activatable"}
	gtk_list_item_set_activatable            = cc.DlFunc[func(*ListItem, int32), cc.Void]{Name: "gtk_list_item_set_activatable"}
	gtk_list_item_get_focusable              = cc.DlFunc[func(*ListItem) int32, int32]{Name: "gtk_list_item_get_focusable"}
	gtk_list_item_set_focusable              = cc.DlFunc[func(*ListItem, int32), cc.Void]{Name: "gtk_list_item_set_focusable"}
	gtk_list_item_set_child                  = cc.DlFunc[func(*ListItem, *Widget), cc.Void]{Name: "gtk_list_item_set_child"}
	gtk_list_item_get_child                  = cc.DlFunc[func(*ListItem) *Widget, *Widget]{Name: "gtk_list_item_get_child"}
	gtk_list_item_set_accessible_description = cc.DlFunc[func(*ListItem, cc.String), cc.Void]{Name: "gtk_list_item_set_accessible_description"}
	gtk_list_item_get_accessible_description = cc.DlFunc[func(*ListItem) cc.String, cc.String]{Name: "gtk_list_item_get_accessible_description"}
	gtk_list_item_set_accessible_label       = cc.DlFunc[func(*ListItem, cc.String), cc.Void]{Name: "gtk_list_item_set_accessible_label"}
	gtk_list_item_get_accessible_label       = cc.DlFunc[func(*ListItem) cc.String, cc.String]{Name: "gtk_list_item_get_accessible_label"}
	// #endregion

	// #region ListItemFactory
	gtk_list_item_factory_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_list_item_factory_get_type"}
	// #endregion

	// #region ListView
	gtk_list_view_get_type                  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_list_view_get_type"}
	gtk_list_view_new                       = cc.DlFunc[func(model *SelectionModel, factory *ListItemFactory) *ListView, *ListView]{Name: "gtk_list_view_new"}
	gtk_list_view_get_model                 = cc.DlFunc[func(lv *ListView) *SelectionModel, *SelectionModel]{Name: "gtk_list_view_get_model"}
	gtk_list_view_set_model                 = cc.DlFunc[func(lv *ListView, model *SelectionModel), cc.Void]{Name: "gtk_list_view_set_model"}
	gtk_list_view_set_factory               = cc.DlFunc[func(lv *ListView, factory *ListItemFactory), cc.Void]{Name: "gtk_list_view_set_factory"}
	gtk_list_view_get_factory               = cc.DlFunc[func(lv *ListView) *ListItemFactory, *ListItemFactory]{Name: "gtk_list_view_get_factory"}
	gtk_list_view_set_header_factory        = cc.DlFunc[func(lv *ListView, factory *ListItemFactory), cc.Void]{Name: "gtk_list_view_set_header_factory"}
	gtk_list_view_get_header_factory        = cc.DlFunc[func(lv *ListView) *ListItemFactory, *ListItemFactory]{Name: "gtk_list_view_get_header_factory"}
	gtk_list_view_set_show_separators       = cc.DlFunc[func(lv *ListView, showSeparators int32), cc.Void]{Name: "gtk_list_view_set_show_separators"}
	gtk_list_view_get_show_separators       = cc.DlFunc[func(lv *ListView) int32, int32]{Name: "gtk_list_view_get_show_separators"}
	gtk_list_view_set_single_click_activate = cc.DlFunc[func(lv *ListView, singleClickActivate int32), cc.Void]{Name: "gtk_list_view_set_single_click_activate"}
	gtk_list_view_get_single_click_activate = cc.DlFunc[func(lv *ListView) int32, int32]{Name: "gtk_list_view_get_single_click_activate"}
	gtk_list_view_set_enable_rubberband     = cc.DlFunc[func(lv *ListView, enableRubberband int32), cc.Void]{Name: "gtk_list_view_set_enable_rubberband"}
	gtk_list_view_get_enable_rubberband     = cc.DlFunc[func(lv *ListView) int32, int32]{Name: "gtk_list_view_get_enable_rubberband"}
	gtk_list_view_set_tab_behavior          = cc.DlFunc[func(lv *ListView, tabBehavior ListTabBehavior), cc.Void]{Name: "gtk_list_view_set_tab_behavior"}
	gtk_list_view_get_tab_behavior          = cc.DlFunc[func(lv *ListView) ListTabBehavior, ListTabBehavior]{Name: "gtk_list_view_get_tab_behavior"}
	gtk_list_view_scroll_to                 = cc.DlFunc[func(lv *ListView, pos uint32, flags ListScrollFlags, scroll *ScrollInfo), cc.Void]{Name: "gtk_list_view_scroll_to"}
	// #endregion

	// #region Main
	gtk_init                 = cc.DlFunc[func(), cc.Void]{Name: "gtk_init"}
	gtk_init_check           = cc.DlFunc[func() int32, int32]{Name: "gtk_init_check"}
	gtk_is_initialized       = cc.DlFunc[func() int32, int32]{Name: "gtk_is_initialized"}
	gtk_init_abi_check       = cc.DlFunc[func(numChecks int32, sizeofWindow uintptr, sizeofBox uintptr), cc.Void]{Name: "gtk_init_abi_check"}
	gtk_init_check_abi_check = cc.DlFunc[func(numChecks int32, sizeofWindow uintptr, sizeofBox uintptr) int32, int32]{Name: "gtk_init_check_abi_check"}
	gtk_disable_setlocale    = cc.DlFunc[func(), cc.Void]{Name: "gtk_disable_setlocale"}
	gtk_disable_portals      = cc.DlFunc[func(), cc.Void]{Name: "gtk_disable_portals"}
	gtk_get_default_language = cc.DlFunc[func() *pango.Language, *pango.Language]{Name: "gtk_get_default_language"}
	gtk_get_locale_direction = cc.DlFunc[func() TextDirection, TextDirection]{Name: "gtk_get_locale_direction"}
	// #endregion

	// #region MapListModel
	gtk_map_list_model_get_type     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_map_list_model_get_type"}
	gtk_map_list_model_new          = cc.DlFunc[func(model *gio.GListModel, mapFunc uptr, userData uptr, userDestroy uptr) *MapListModel, *MapListModel]{Name: "gtk_map_list_model_new"}
	gtk_map_list_model_set_map_func = cc.DlFunc[func(m *MapListModel, mapFunc uptr, userData uptr, userDestroy uptr), cc.Void]{Name: "gtk_map_list_model_set_map_func"}
	gtk_map_list_model_set_model    = cc.DlFunc[func(m *MapListModel, model *gio.GListModel), cc.Void]{Name: "gtk_map_list_model_set_model"}
	gtk_map_list_model_get_model    = cc.DlFunc[func(m *MapListModel) *gio.GListModel, *gio.GListModel]{Name: "gtk_map_list_model_get_model"}
	gtk_map_list_model_has_map      = cc.DlFunc[func(m *MapListModel) int32, int32]{Name: "gtk_map_list_model_has_map"}
	// #endregion

	// #region MediaControls
	gtk_media_controls_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_media_controls_get_type"}
	gtk_media_controls_new              = cc.DlFunc[func(stream *MediaStream) *MediaControls, *MediaControls]{Name: "gtk_media_controls_new"}
	gtk_media_controls_get_media_stream = cc.DlFunc[func(controls *MediaControls) *MediaStream, *MediaStream]{Name: "gtk_media_controls_get_media_stream"}
	gtk_media_controls_set_media_stream = cc.DlFunc[func(controls *MediaControls, stream *MediaStream), cc.Void]{Name: "gtk_media_controls_set_media_stream"}
	// #endregion

	// #region MediaFile
	gtk_media_file_get_type             = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_media_file_get_type"}
	gtk_media_file_new                  = cc.DlFunc[func() *MediaStream, *MediaStream]{Name: "gtk_media_file_new"}
	gtk_media_file_new_for_filename     = cc.DlFunc[func(filename cc.String) *MediaStream, *MediaStream]{Name: "gtk_media_file_new_for_filename"}
	gtk_media_file_new_for_resource     = cc.DlFunc[func(resourcePath cc.String) *MediaStream, *MediaStream]{Name: "gtk_media_file_new_for_resource"}
	gtk_media_file_new_for_file         = cc.DlFunc[func(file *gio.GFile) *MediaStream, *MediaStream]{Name: "gtk_media_file_new_for_file"}
	gtk_media_file_new_for_input_stream = cc.DlFunc[func(stream *gio.GInputStream) *MediaStream, *MediaStream]{Name: "gtk_media_file_new_for_input_stream"}
	gtk_media_file_clear                = cc.DlFunc[func(mf *MediaFile), cc.Void]{Name: "gtk_media_file_clear"}
	gtk_media_file_set_filename         = cc.DlFunc[func(mf *MediaFile, filename cc.String), cc.Void]{Name: "gtk_media_file_set_filename"}
	gtk_media_file_set_resource         = cc.DlFunc[func(mf *MediaFile, resourcePath cc.String), cc.Void]{Name: "gtk_media_file_set_resource"}
	gtk_media_file_set_file             = cc.DlFunc[func(mf *MediaFile, file *gio.GFile), cc.Void]{Name: "gtk_media_file_set_file"}
	gtk_media_file_get_file             = cc.DlFunc[func(mf *MediaFile) *gio.GFile, *gio.GFile]{Name: "gtk_media_file_get_file"}
	gtk_media_file_set_input_stream     = cc.DlFunc[func(mf *MediaFile, stream *gio.GInputStream), cc.Void]{Name: "gtk_media_file_set_input_stream"}
	gtk_media_file_get_input_stream     = cc.DlFunc[func(mf *MediaFile) *gio.GInputStream, *gio.GInputStream]{Name: "gtk_media_file_get_input_stream"}
	// #endregion

	// #region MediaStream
	gtk_media_stream_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_media_stream_get_type"}
	gtk_media_stream_is_prepared       = cc.DlFunc[func(ms *MediaStream) int32, int32]{Name: "gtk_media_stream_is_prepared"}
	gtk_media_stream_get_error         = cc.DlFunc[func(ms *MediaStream) *glib.GError, *glib.GError]{Name: "gtk_media_stream_get_error"}
	gtk_media_stream_has_audio         = cc.DlFunc[func(ms *MediaStream) int32, int32]{Name: "gtk_media_stream_has_audio"}
	gtk_media_stream_has_video         = cc.DlFunc[func(ms *MediaStream) int32, int32]{Name: "gtk_media_stream_has_video"}
	gtk_media_stream_play              = cc.DlFunc[func(ms *MediaStream), cc.Void]{Name: "gtk_media_stream_play"}
	gtk_media_stream_pause             = cc.DlFunc[func(ms *MediaStream), cc.Void]{Name: "gtk_media_stream_pause"}
	gtk_media_stream_get_playing       = cc.DlFunc[func(ms *MediaStream) int32, int32]{Name: "gtk_media_stream_get_playing"}
	gtk_media_stream_set_playing       = cc.DlFunc[func(ms *MediaStream, playing int32), cc.Void]{Name: "gtk_media_stream_set_playing"}
	gtk_media_stream_get_ended         = cc.DlFunc[func(ms *MediaStream) int32, int32]{Name: "gtk_media_stream_get_ended"}
	gtk_media_stream_get_timestamp     = cc.DlFunc[func(ms *MediaStream) int64, int64]{Name: "gtk_media_stream_get_timestamp"}
	gtk_media_stream_get_duration      = cc.DlFunc[func(ms *MediaStream) int64, int64]{Name: "gtk_media_stream_get_duration"}
	gtk_media_stream_is_seekable       = cc.DlFunc[func(ms *MediaStream) int32, int32]{Name: "gtk_media_stream_is_seekable"}
	gtk_media_stream_is_seeking        = cc.DlFunc[func(ms *MediaStream) int32, int32]{Name: "gtk_media_stream_is_seeking"}
	gtk_media_stream_seek              = cc.DlFunc[func(ms *MediaStream, timestamp int64), cc.Void]{Name: "gtk_media_stream_seek"}
	gtk_media_stream_get_loop          = cc.DlFunc[func(ms *MediaStream) int32, int32]{Name: "gtk_media_stream_get_loop"}
	gtk_media_stream_set_loop          = cc.DlFunc[func(ms *MediaStream, loop int32), cc.Void]{Name: "gtk_media_stream_set_loop"}
	gtk_media_stream_get_muted         = cc.DlFunc[func(ms *MediaStream) int32, int32]{Name: "gtk_media_stream_get_muted"}
	gtk_media_stream_set_muted         = cc.DlFunc[func(ms *MediaStream, muted int32), cc.Void]{Name: "gtk_media_stream_set_muted"}
	gtk_media_stream_get_volume        = cc.DlFunc[func(ms *MediaStream) float64, float64]{Name: "gtk_media_stream_get_volume"}
	gtk_media_stream_set_volume        = cc.DlFunc[func(ms *MediaStream, volume float64), cc.Void]{Name: "gtk_media_stream_set_volume"}
	gtk_media_stream_realize           = cc.DlFunc[func(ms *MediaStream, surface *gdk.Surface), cc.Void]{Name: "gtk_media_stream_realize"}
	gtk_media_stream_unrealize         = cc.DlFunc[func(ms *MediaStream, surface *gdk.Surface), cc.Void]{Name: "gtk_media_stream_unrealize"}
	gtk_media_stream_stream_prepared   = cc.DlFunc[func(ms *MediaStream, hasAudio, hasVideo, seekable int32, duration int64), cc.Void]{Name: "gtk_media_stream_stream_prepared"}
	gtk_media_stream_stream_unprepared = cc.DlFunc[func(ms *MediaStream), cc.Void]{Name: "gtk_media_stream_stream_unprepared"}
	gtk_media_stream_update            = cc.DlFunc[func(ms *MediaStream, timestamp int64), cc.Void]{Name: "gtk_media_stream_update"}
	gtk_media_stream_stream_ended      = cc.DlFunc[func(ms *MediaStream), cc.Void]{Name: "gtk_media_stream_stream_ended"}
	gtk_media_stream_seek_success      = cc.DlFunc[func(ms *MediaStream), cc.Void]{Name: "gtk_media_stream_seek_success"}
	gtk_media_stream_seek_failed       = cc.DlFunc[func(ms *MediaStream), cc.Void]{Name: "gtk_media_stream_seek_failed"}
	gtk_media_stream_gerror            = cc.DlFunc[func(ms *MediaStream, err *glib.GError), cc.Void]{Name: "gtk_media_stream_gerror"}
	gtk_media_stream_error             = cc.DlFunc[func(ms *MediaStream, domain glib.GQuark, code int32, format cc.String, args ...interface{}), cc.Void]{Name: "gtk_media_stream_error", Va: true}
	// #endregion

	// #region MenuButton
	gtk_menu_button_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_menu_button_get_type"}
	gtk_menu_button_new                   = cc.DlFunc[func() *MenuButton, *MenuButton]{Name: "gtk_menu_button_new"}
	gtk_menu_button_set_popover           = cc.DlFunc[func(mb *MenuButton, popover *Popover), cc.Void]{Name: "gtk_menu_button_set_popover"}
	gtk_menu_button_get_popover           = cc.DlFunc[func(mb *MenuButton) *Popover, *Popover]{Name: "gtk_menu_button_get_popover"}
	gtk_menu_button_set_direction         = cc.DlFunc[func(mb *MenuButton, direction ArrowType), cc.Void]{Name: "gtk_menu_button_set_direction"}
	gtk_menu_button_get_direction         = cc.DlFunc[func(mb *MenuButton) ArrowType, ArrowType]{Name: "gtk_menu_button_get_direction"}
	gtk_menu_button_set_menu_model        = cc.DlFunc[func(mb *MenuButton, menuModel *gio.GMenuModel), cc.Void]{Name: "gtk_menu_button_set_menu_model"}
	gtk_menu_button_get_menu_model        = cc.DlFunc[func(mb *MenuButton) *gio.GMenuModel, *gio.GMenuModel]{Name: "gtk_menu_button_get_menu_model"}
	gtk_menu_button_set_icon_name         = cc.DlFunc[func(mb *MenuButton, iconName cc.String), cc.Void]{Name: "gtk_menu_button_set_icon_name"}
	gtk_menu_button_get_icon_name         = cc.DlFunc[func(mb *MenuButton) cc.String, cc.String]{Name: "gtk_menu_button_get_icon_name"}
	gtk_menu_button_set_always_show_arrow = cc.DlFunc[func(mb *MenuButton, alwaysShowArrow int32), cc.Void]{Name: "gtk_menu_button_set_always_show_arrow"}
	gtk_menu_button_get_always_show_arrow = cc.DlFunc[func(mb *MenuButton) int32, int32]{Name: "gtk_menu_button_get_always_show_arrow"}
	gtk_menu_button_set_label             = cc.DlFunc[func(mb *MenuButton, label cc.String), cc.Void]{Name: "gtk_menu_button_set_label"}
	gtk_menu_button_get_label             = cc.DlFunc[func(mb *MenuButton) cc.String, cc.String]{Name: "gtk_menu_button_get_label"}
	gtk_menu_button_set_use_underline     = cc.DlFunc[func(mb *MenuButton, useUnderline int32), cc.Void]{Name: "gtk_menu_button_set_use_underline"}
	gtk_menu_button_get_use_underline     = cc.DlFunc[func(mb *MenuButton) int32, int32]{Name: "gtk_menu_button_get_use_underline"}
	gtk_menu_button_set_has_frame         = cc.DlFunc[func(mb *MenuButton, hasFrame int32), cc.Void]{Name: "gtk_menu_button_set_has_frame"}
	gtk_menu_button_get_has_frame         = cc.DlFunc[func(mb *MenuButton) int32, int32]{Name: "gtk_menu_button_get_has_frame"}
	gtk_menu_button_popup                 = cc.DlFunc[func(mb *MenuButton), cc.Void]{Name: "gtk_menu_button_popup"}
	gtk_menu_button_popdown               = cc.DlFunc[func(mb *MenuButton), cc.Void]{Name: "gtk_menu_button_popdown"}
	gtk_menu_button_set_create_popup_func = cc.DlFunc[func(mb *MenuButton, fn uptr, userData uptr, destroy uptr), cc.Void]{Name: "gtk_menu_button_set_create_popup_func"}
	gtk_menu_button_set_primary           = cc.DlFunc[func(mb *MenuButton, primary int32), cc.Void]{Name: "gtk_menu_button_set_primary"}
	gtk_menu_button_get_primary           = cc.DlFunc[func(mb *MenuButton) int32, int32]{Name: "gtk_menu_button_get_primary"}
	gtk_menu_button_set_child             = cc.DlFunc[func(mb *MenuButton, child *Widget), cc.Void]{Name: "gtk_menu_button_set_child"}
	gtk_menu_button_get_child             = cc.DlFunc[func(mb *MenuButton) *Widget, *Widget]{Name: "gtk_menu_button_get_child"}
	gtk_menu_button_set_active            = cc.DlFunc[func(mb *MenuButton, active int32), cc.Void]{Name: "gtk_menu_button_set_active"}
	gtk_menu_button_get_active            = cc.DlFunc[func(mb *MenuButton) int32, int32]{Name: "gtk_menu_button_get_active"}
	gtk_menu_button_set_can_shrink        = cc.DlFunc[func(mb *MenuButton, canShrink int32), cc.Void]{Name: "gtk_menu_button_set_can_shrink"}
	gtk_menu_button_get_can_shrink        = cc.DlFunc[func(mb *MenuButton) int32, int32]{Name: "gtk_menu_button_get_can_shrink"}
	// #endregion

	// #region MountOperation
	gtk_mount_operation_get_type    = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_mount_operation_get_type"}
	gtk_mount_operation_new         = cc.DlFunc[func(parent *Window) *MountOperation, *MountOperation]{Name: "gtk_mount_operation_new"}
	gtk_mount_operation_is_showing  = cc.DlFunc[func(op *MountOperation) int32, int32]{Name: "gtk_mount_operation_is_showing"}
	gtk_mount_operation_set_parent  = cc.DlFunc[func(op *MountOperation, parent *Window), cc.Void]{Name: "gtk_mount_operation_set_parent"}
	gtk_mount_operation_get_parent  = cc.DlFunc[func(op *MountOperation) *Window, *Window]{Name: "gtk_mount_operation_get_parent"}
	gtk_mount_operation_set_display = cc.DlFunc[func(op *MountOperation, display *gdk.Display), cc.Void]{Name: "gtk_mount_operation_set_display"}
	gtk_mount_operation_get_display = cc.DlFunc[func(op *MountOperation) *gdk.Display, *gdk.Display]{Name: "gtk_mount_operation_get_display"}
	// #endregion

	// #region MultiFilter
	gtk_multi_filter_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_multi_filter_get_type"}
	gtk_multi_filter_append   = cc.DlFunc[func(m *MultiFilter, filter *Filter), cc.Void]{Name: "gtk_multi_filter_append"}
	gtk_multi_filter_remove   = cc.DlFunc[func(m *MultiFilter, position uint32), cc.Void]{Name: "gtk_multi_filter_remove"}
	gtk_any_filter_get_type   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_any_filter_get_type"}
	gtk_any_filter_new        = cc.DlFunc[func() *AnyFilter, *AnyFilter]{Name: "gtk_any_filter_new"}
	gtk_every_filter_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_every_filter_get_type"}
	gtk_every_filter_new      = cc.DlFunc[func() *EveryFilter, *EveryFilter]{Name: "gtk_every_filter_new"}
	// #endregion

	// #region MultiSorter
	gtk_multi_sorter_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_multi_sorter_get_type"}
	gtk_multi_sorter_new      = cc.DlFunc[func() *MultiSorter, *MultiSorter]{Name: "gtk_multi_sorter_new"}
	gtk_multi_sorter_append   = cc.DlFunc[func(ms *MultiSorter, s *Sorter), cc.Void]{Name: "gtk_multi_sorter_append"}
	gtk_multi_sorter_remove   = cc.DlFunc[func(ms *MultiSorter, position uint32), cc.Void]{Name: "gtk_multi_sorter_remove"}
	// #endregion

	// #region MultiSelection
	gtk_multi_selection_get_type  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_multi_selection_get_type"}
	gtk_multi_selection_new       = cc.DlFunc[func(model *gio.GListModel) *MultiSelection, *MultiSelection]{Name: "gtk_multi_selection_new"}
	gtk_multi_selection_get_model = cc.DlFunc[func(self *MultiSelection) *gio.GListModel, *gio.GListModel]{Name: "gtk_multi_selection_get_model"}
	gtk_multi_selection_set_model = cc.DlFunc[func(self *MultiSelection, model *gio.GListModel), cc.Void]{Name: "gtk_multi_selection_set_model"}
	// #endregion

	// #region Native
	gtk_native_realize               = cc.DlFunc[func(*Native), cc.Void]{Name: "gtk_native_realize"}
	gtk_native_unrealize             = cc.DlFunc[func(*Native), cc.Void]{Name: "gtk_native_unrealize"}
	gtk_native_get_for_surface       = cc.DlFunc[func(*gdk.Surface) *Native, *Native]{Name: "gtk_native_get_for_surface"}
	gtk_native_get_surface           = cc.DlFunc[func(*Native) *gdk.Surface, *gdk.Surface]{Name: "gtk_native_get_surface"}
	gtk_native_get_renderer          = cc.DlFunc[func(*Native) *gsk.Renderer, *gsk.Renderer]{Name: "gtk_native_get_renderer"}
	gtk_native_get_surface_transform = cc.DlFunc[func(*Native, *float64, *float64), cc.Void]{Name: "gtk_native_get_surface_transform"}
	// #endregion

	// #region NativeDialog
	gtk_native_dialog_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_native_dialog_get_type"}
	gtk_native_dialog_show              = cc.DlFunc[func(nd *NativeDialog), cc.Void]{Name: "gtk_native_dialog_show"}
	gtk_native_dialog_hide              = cc.DlFunc[func(nd *NativeDialog), cc.Void]{Name: "gtk_native_dialog_hide"}
	gtk_native_dialog_destroy           = cc.DlFunc[func(nd *NativeDialog), cc.Void]{Name: "gtk_native_dialog_destroy"}
	gtk_native_dialog_get_visible       = cc.DlFunc[func(nd *NativeDialog) int32, int32]{Name: "gtk_native_dialog_get_visible"}
	gtk_native_dialog_set_modal         = cc.DlFunc[func(nd *NativeDialog, modal int32), cc.Void]{Name: "gtk_native_dialog_set_modal"}
	gtk_native_dialog_get_modal         = cc.DlFunc[func(nd *NativeDialog) int32, int32]{Name: "gtk_native_dialog_get_modal"}
	gtk_native_dialog_set_title         = cc.DlFunc[func(nd *NativeDialog, title cc.String), cc.Void]{Name: "gtk_native_dialog_set_title"}
	gtk_native_dialog_get_title         = cc.DlFunc[func(nd *NativeDialog) cc.String, cc.String]{Name: "gtk_native_dialog_get_title"}
	gtk_native_dialog_set_transient_for = cc.DlFunc[func(nd *NativeDialog, parent *Window), cc.Void]{Name: "gtk_native_dialog_set_transient_for"}
	gtk_native_dialog_get_transient_for = cc.DlFunc[func(nd *NativeDialog) *Window, *Window]{Name: "gtk_native_dialog_get_transient_for"}
	// #endregion

	// #region NoSelection
	gtk_no_selection_get_type  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_no_selection_get_type"}
	gtk_no_selection_new       = cc.DlFunc[func(model *gio.GListModel) *NoSelection, *NoSelection]{Name: "gtk_no_selection_new"}
	gtk_no_selection_get_model = cc.DlFunc[func(ns *NoSelection) *gio.GListModel, *gio.GListModel]{Name: "gtk_no_selection_get_model"}
	gtk_no_selection_set_model = cc.DlFunc[func(ns *NoSelection, model *gio.GListModel), cc.Void]{Name: "gtk_no_selection_set_model"}
	// #endregion

	// #region Notebook
	gtk_notebook_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_notebook_get_type"}
	gtk_notebook_new                 = cc.DlFunc[func() *Notebook, *Notebook]{Name: "gtk_notebook_new"}
	gtk_notebook_append_page         = cc.DlFunc[func(nb *Notebook, child, tabLabel *Widget) int32, int32]{Name: "gtk_notebook_append_page"}
	gtk_notebook_append_page_menu    = cc.DlFunc[func(nb *Notebook, child, tabLabel, menuLabel *Widget) int32, int32]{Name: "gtk_notebook_append_page_menu"}
	gtk_notebook_prepend_page        = cc.DlFunc[func(nb *Notebook, child, tabLabel *Widget) int32, int32]{Name: "gtk_notebook_prepend_page"}
	gtk_notebook_prepend_page_menu   = cc.DlFunc[func(nb *Notebook, child, tabLabel, menuLabel *Widget) int32, int32]{Name: "gtk_notebook_prepend_page_menu"}
	gtk_notebook_insert_page         = cc.DlFunc[func(nb *Notebook, child, tabLabel *Widget, position int32) int32, int32]{Name: "gtk_notebook_insert_page"}
	gtk_notebook_insert_page_menu    = cc.DlFunc[func(nb *Notebook, child, tabLabel, menuLabel *Widget, position int32) int32, int32]{Name: "gtk_notebook_insert_page_menu"}
	gtk_notebook_remove_page         = cc.DlFunc[func(nb *Notebook, pageNum int32), cc.Void]{Name: "gtk_notebook_remove_page"}
	gtk_notebook_set_group_name      = cc.DlFunc[func(nb *Notebook, groupName cc.String), cc.Void]{Name: "gtk_notebook_set_group_name"}
	gtk_notebook_get_group_name      = cc.DlFunc[func(nb *Notebook) cc.String, cc.String]{Name: "gtk_notebook_get_group_name"}
	gtk_notebook_get_current_page    = cc.DlFunc[func(nb *Notebook) int32, int32]{Name: "gtk_notebook_get_current_page"}
	gtk_notebook_get_nth_page        = cc.DlFunc[func(nb *Notebook, pageNum int32) *Widget, *Widget]{Name: "gtk_notebook_get_nth_page"}
	gtk_notebook_get_n_pages         = cc.DlFunc[func(nb *Notebook) int32, int32]{Name: "gtk_notebook_get_n_pages"}
	gtk_notebook_page_num            = cc.DlFunc[func(nb *Notebook, child *Widget) int32, int32]{Name: "gtk_notebook_page_num"}
	gtk_notebook_set_current_page    = cc.DlFunc[func(nb *Notebook, pageNum int32), cc.Void]{Name: "gtk_notebook_set_current_page"}
	gtk_notebook_next_page           = cc.DlFunc[func(nb *Notebook), cc.Void]{Name: "gtk_notebook_next_page"}
	gtk_notebook_prev_page           = cc.DlFunc[func(nb *Notebook), cc.Void]{Name: "gtk_notebook_prev_page"}
	gtk_notebook_set_show_border     = cc.DlFunc[func(nb *Notebook, showBorder int32), cc.Void]{Name: "gtk_notebook_set_show_border"}
	gtk_notebook_get_show_border     = cc.DlFunc[func(nb *Notebook) int32, int32]{Name: "gtk_notebook_get_show_border"}
	gtk_notebook_set_show_tabs       = cc.DlFunc[func(nb *Notebook, showTabs int32), cc.Void]{Name: "gtk_notebook_set_show_tabs"}
	gtk_notebook_get_show_tabs       = cc.DlFunc[func(nb *Notebook) int32, int32]{Name: "gtk_notebook_get_show_tabs"}
	gtk_notebook_set_tab_pos         = cc.DlFunc[func(nb *Notebook, pos PositionType), cc.Void]{Name: "gtk_notebook_set_tab_pos"}
	gtk_notebook_get_tab_pos         = cc.DlFunc[func(nb *Notebook) PositionType, PositionType]{Name: "gtk_notebook_get_tab_pos"}
	gtk_notebook_set_scrollable      = cc.DlFunc[func(nb *Notebook, scrollable int32), cc.Void]{Name: "gtk_notebook_set_scrollable"}
	gtk_notebook_get_scrollable      = cc.DlFunc[func(nb *Notebook) int32, int32]{Name: "gtk_notebook_get_scrollable"}
	gtk_notebook_popup_enable        = cc.DlFunc[func(nb *Notebook), cc.Void]{Name: "gtk_notebook_popup_enable"}
	gtk_notebook_popup_disable       = cc.DlFunc[func(nb *Notebook), cc.Void]{Name: "gtk_notebook_popup_disable"}
	gtk_notebook_get_tab_label       = cc.DlFunc[func(nb *Notebook, child *Widget) *Widget, *Widget]{Name: "gtk_notebook_get_tab_label"}
	gtk_notebook_set_tab_label       = cc.DlFunc[func(nb *Notebook, child, tabLabel *Widget), cc.Void]{Name: "gtk_notebook_set_tab_label"}
	gtk_notebook_set_tab_label_text  = cc.DlFunc[func(nb *Notebook, child *Widget, tabText cc.String), cc.Void]{Name: "gtk_notebook_set_tab_label_text"}
	gtk_notebook_get_tab_label_text  = cc.DlFunc[func(nb *Notebook, child *Widget) cc.String, cc.String]{Name: "gtk_notebook_get_tab_label_text"}
	gtk_notebook_get_menu_label      = cc.DlFunc[func(nb *Notebook, child *Widget) *Widget, *Widget]{Name: "gtk_notebook_get_menu_label"}
	gtk_notebook_set_menu_label      = cc.DlFunc[func(nb *Notebook, child, menuLabel *Widget), cc.Void]{Name: "gtk_notebook_set_menu_label"}
	gtk_notebook_set_menu_label_text = cc.DlFunc[func(nb *Notebook, child *Widget, menuText cc.String), cc.Void]{Name: "gtk_notebook_set_menu_label_text"}
	gtk_notebook_get_menu_label_text = cc.DlFunc[func(nb *Notebook, child *Widget) cc.String, cc.String]{Name: "gtk_notebook_get_menu_label_text"}
	gtk_notebook_reorder_child       = cc.DlFunc[func(nb *Notebook, child *Widget, position int32), cc.Void]{Name: "gtk_notebook_reorder_child"}
	gtk_notebook_get_tab_reorderable = cc.DlFunc[func(nb *Notebook, child *Widget) int32, int32]{Name: "gtk_notebook_get_tab_reorderable"}
	gtk_notebook_set_tab_reorderable = cc.DlFunc[func(nb *Notebook, child *Widget, reorderable int32), cc.Void]{Name: "gtk_notebook_set_tab_reorderable"}
	gtk_notebook_get_tab_detachable  = cc.DlFunc[func(nb *Notebook, child *Widget) int32, int32]{Name: "gtk_notebook_get_tab_detachable"}
	gtk_notebook_set_tab_detachable  = cc.DlFunc[func(nb *Notebook, child *Widget, detachable int32), cc.Void]{Name: "gtk_notebook_set_tab_detachable"}
	gtk_notebook_detach_tab          = cc.DlFunc[func(nb *Notebook, child *Widget), cc.Void]{Name: "gtk_notebook_detach_tab"}
	gtk_notebook_get_action_widget   = cc.DlFunc[func(nb *Notebook, packType PackType) *Widget, *Widget]{Name: "gtk_notebook_get_action_widget"}
	gtk_notebook_set_action_widget   = cc.DlFunc[func(nb *Notebook, widget *Widget, packType PackType), cc.Void]{Name: "gtk_notebook_set_action_widget"}
	gtk_notebook_page_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_notebook_page_get_type"}
	gtk_notebook_get_page            = cc.DlFunc[func(nb *Notebook, child *Widget) *NotebookPage, *NotebookPage]{Name: "gtk_notebook_get_page"}
	gtk_notebook_page_get_child      = cc.DlFunc[func(page *NotebookPage) *Widget, *Widget]{Name: "gtk_notebook_page_get_child"}
	gtk_notebook_get_pages           = cc.DlFunc[func(nb *Notebook) *gio.GListModel, *gio.GListModel]{Name: "gtk_notebook_get_pages"}
	// #endregion

	// #region NumericSorter
	gtk_numeric_sorter_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_numeric_sorter_get_type"}
	gtk_numeric_sorter_new            = cc.DlFunc[func(expression *Expression) *NumericSorter, *NumericSorter]{Name: "gtk_numeric_sorter_new"}
	gtk_numeric_sorter_get_expression = cc.DlFunc[func(ns *NumericSorter) *Expression, *Expression]{Name: "gtk_numeric_sorter_get_expression"}
	gtk_numeric_sorter_set_expression = cc.DlFunc[func(ns *NumericSorter, expression *Expression), cc.Void]{Name: "gtk_numeric_sorter_set_expression"}
	gtk_numeric_sorter_get_sort_order = cc.DlFunc[func(ns *NumericSorter) SortType, SortType]{Name: "gtk_numeric_sorter_get_sort_order"}
	gtk_numeric_sorter_set_sort_order = cc.DlFunc[func(ns *NumericSorter, sortOrder SortType), cc.Void]{Name: "gtk_numeric_sorter_set_sort_order"}
	// #endregion

	// #region Orientable
	gtk_orientable_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_orientable_get_type"}
	gtk_orientable_set_orientation = cc.DlFunc[func(orientable *Orientable, orientation Orientation), cc.Void]{Name: "gtk_orientable_set_orientation"}
	gtk_orientable_get_orientation = cc.DlFunc[func(orientable *Orientable) Orientation, Orientation]{Name: "gtk_orientable_get_orientation"}
	// #endregion

	// #region Overlay
	gtk_overlay_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_overlay_get_type"}
	gtk_overlay_new                 = cc.DlFunc[func() *Overlay, *Overlay]{Name: "gtk_overlay_new"}
	gtk_overlay_add_overlay         = cc.DlFunc[func(overlay *Overlay, widget *Widget), cc.Void]{Name: "gtk_overlay_add_overlay"}
	gtk_overlay_remove_overlay      = cc.DlFunc[func(overlay *Overlay, widget *Widget), cc.Void]{Name: "gtk_overlay_remove_overlay"}
	gtk_overlay_set_child           = cc.DlFunc[func(overlay *Overlay, child *Widget), cc.Void]{Name: "gtk_overlay_set_child"}
	gtk_overlay_get_child           = cc.DlFunc[func(overlay *Overlay) *Widget, *Widget]{Name: "gtk_overlay_get_child"}
	gtk_overlay_get_measure_overlay = cc.DlFunc[func(overlay *Overlay, widget *Widget) int32, int32]{Name: "gtk_overlay_get_measure_overlay"}
	gtk_overlay_set_measure_overlay = cc.DlFunc[func(overlay *Overlay, widget *Widget, measure int32), cc.Void]{Name: "gtk_overlay_set_measure_overlay"}
	gtk_overlay_get_clip_overlay    = cc.DlFunc[func(overlay *Overlay, widget *Widget) int32, int32]{Name: "gtk_overlay_get_clip_overlay"}
	gtk_overlay_set_clip_overlay    = cc.DlFunc[func(overlay *Overlay, widget *Widget, clipOverlay int32), cc.Void]{Name: "gtk_overlay_set_clip_overlay"}
	// #endregion

	// #region OverlayLayout
	gtk_overlay_layout_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_overlay_layout_get_type"}
	gtk_overlay_layout_new                    = cc.DlFunc[func() *OverlayLayout, *OverlayLayout]{Name: "gtk_overlay_layout_new"}
	gtk_overlay_layout_child_set_measure      = cc.DlFunc[func(child *OverlayLayoutChild, measure int32), cc.Void]{Name: "gtk_overlay_layout_child_set_measure"}
	gtk_overlay_layout_child_get_measure      = cc.DlFunc[func(child *OverlayLayoutChild) int32, int32]{Name: "gtk_overlay_layout_child_get_measure"}
	gtk_overlay_layout_child_set_clip_overlay = cc.DlFunc[func(child *OverlayLayoutChild, clipOverlay int32), cc.Void]{Name: "gtk_overlay_layout_child_set_clip_overlay"}
	gtk_overlay_layout_child_get_clip_overlay = cc.DlFunc[func(child *OverlayLayoutChild) int32, int32]{Name: "gtk_overlay_layout_child_get_clip_overlay"}
	// #endregion

	// #region PadController
	gtk_pad_controller_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_pad_controller_get_type"}
	gtk_pad_controller_new                = cc.DlFunc[func(group *gio.GActionGroup, pad *gdk.Device) *PadController, *PadController]{Name: "gtk_pad_controller_new"}
	gtk_pad_controller_set_action_entries = cc.DlFunc[func(controller *PadController, entries *PadActionEntry, nEntries int32), cc.Void]{Name: "gtk_pad_controller_set_action_entries"}
	gtk_pad_controller_set_action         = cc.DlFunc[func(controller *PadController, actionType PadActionType, index int32, mode int32, label cc.String, actionName cc.String), cc.Void]{Name: "gtk_pad_controller_set_action"}
	// #endregion

	// #region Paned
	gtk_paned_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_paned_get_type"}
	gtk_paned_new                    = cc.DlFunc[func(orientation Orientation) *Paned, *Paned]{Name: "gtk_paned_new"}
	gtk_paned_set_start_child        = cc.DlFunc[func(paned *Paned, child *Widget), cc.Void]{Name: "gtk_paned_set_start_child"}
	gtk_paned_get_start_child        = cc.DlFunc[func(paned *Paned) *Widget, *Widget]{Name: "gtk_paned_get_start_child"}
	gtk_paned_set_resize_start_child = cc.DlFunc[func(paned *Paned, resize int32), cc.Void]{Name: "gtk_paned_set_resize_start_child"}
	gtk_paned_get_resize_start_child = cc.DlFunc[func(paned *Paned) int32, int32]{Name: "gtk_paned_get_resize_start_child"}
	gtk_paned_set_end_child          = cc.DlFunc[func(paned *Paned, child *Widget), cc.Void]{Name: "gtk_paned_set_end_child"}
	gtk_paned_get_end_child          = cc.DlFunc[func(paned *Paned) *Widget, *Widget]{Name: "gtk_paned_get_end_child"}
	gtk_paned_set_shrink_start_child = cc.DlFunc[func(paned *Paned, resize int32), cc.Void]{Name: "gtk_paned_set_shrink_start_child"}
	gtk_paned_get_shrink_start_child = cc.DlFunc[func(paned *Paned) int32, int32]{Name: "gtk_paned_get_shrink_start_child"}
	gtk_paned_set_resize_end_child   = cc.DlFunc[func(paned *Paned, resize int32), cc.Void]{Name: "gtk_paned_set_resize_end_child"}
	gtk_paned_get_resize_end_child   = cc.DlFunc[func(paned *Paned) int32, int32]{Name: "gtk_paned_get_resize_end_child"}
	gtk_paned_set_shrink_end_child   = cc.DlFunc[func(paned *Paned, resize int32), cc.Void]{Name: "gtk_paned_set_shrink_end_child"}
	gtk_paned_get_shrink_end_child   = cc.DlFunc[func(paned *Paned) int32, int32]{Name: "gtk_paned_get_shrink_end_child"}
	gtk_paned_get_position           = cc.DlFunc[func(paned *Paned) int32, int32]{Name: "gtk_paned_get_position"}
	gtk_paned_set_position           = cc.DlFunc[func(paned *Paned, position int32), cc.Void]{Name: "gtk_paned_set_position"}
	gtk_paned_set_wide_handle        = cc.DlFunc[func(paned *Paned, wide int32), cc.Void]{Name: "gtk_paned_set_wide_handle"}
	gtk_paned_get_wide_handle        = cc.DlFunc[func(paned *Paned) int32, int32]{Name: "gtk_paned_get_wide_handle"}
	// #endregion

	// #region PasswordEntry
	gtk_password_entry_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_password_entry_get_type"}
	gtk_password_entry_new                = cc.DlFunc[func() *PasswordEntry, *PasswordEntry]{Name: "gtk_password_entry_new"}
	gtk_password_entry_set_show_peek_icon = cc.DlFunc[func(entry *PasswordEntry, showPeekIcon int32), cc.Void]{Name: "gtk_password_entry_set_show_peek_icon"}
	gtk_password_entry_get_show_peek_icon = cc.DlFunc[func(entry *PasswordEntry) int32, int32]{Name: "gtk_password_entry_get_show_peek_icon"}
	gtk_password_entry_set_extra_menu     = cc.DlFunc[func(entry *PasswordEntry, model *gio.GMenuModel), cc.Void]{Name: "gtk_password_entry_set_extra_menu"}
	gtk_password_entry_get_extra_menu     = cc.DlFunc[func(entry *PasswordEntry) *gio.GMenuModel, *gio.GMenuModel]{Name: "gtk_password_entry_get_extra_menu"}
	// #endregion

	// #region PasswordEntryBuffer
	gtk_password_entry_buffer_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_password_entry_buffer_get_type"}
	gtk_password_entry_buffer_new      = cc.DlFunc[func() *PasswordEntryBuffer, *PasswordEntryBuffer]{Name: "gtk_password_entry_buffer_new"}
	// #endregion

	// #region Picture
	gtk_picture_get_type             = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_picture_get_type"}
	gtk_picture_new                  = cc.DlFunc[func() *Picture, *Picture]{Name: "gtk_picture_new"}
	gtk_picture_new_for_paintable    = cc.DlFunc[func(paintable *gdk.Paintable) *Picture, *Picture]{Name: "gtk_picture_new_for_paintable"}
	gtk_picture_new_for_file         = cc.DlFunc[func(file *gio.GFile) *Picture, *Picture]{Name: "gtk_picture_new_for_file"}
	gtk_picture_new_for_filename     = cc.DlFunc[func(filename cc.String) *Picture, *Picture]{Name: "gtk_picture_new_for_filename"}
	gtk_picture_new_for_resource     = cc.DlFunc[func(resourcePath cc.String) *Picture, *Picture]{Name: "gtk_picture_new_for_resource"}
	gtk_picture_set_paintable        = cc.DlFunc[func(pic *Picture, paintable *gdk.Paintable), cc.Void]{Name: "gtk_picture_set_paintable"}
	gtk_picture_get_paintable        = cc.DlFunc[func(pic *Picture) *gdk.Paintable, *gdk.Paintable]{Name: "gtk_picture_get_paintable"}
	gtk_picture_set_file             = cc.DlFunc[func(pic *Picture, file *gio.GFile), cc.Void]{Name: "gtk_picture_set_file"}
	gtk_picture_get_file             = cc.DlFunc[func(pic *Picture) *gio.GFile, *gio.GFile]{Name: "gtk_picture_get_file"}
	gtk_picture_set_filename         = cc.DlFunc[func(pic *Picture, filename cc.String), cc.Void]{Name: "gtk_picture_set_filename"}
	gtk_picture_set_resource         = cc.DlFunc[func(pic *Picture, resourcePath cc.String), cc.Void]{Name: "gtk_picture_set_resource"}
	gtk_picture_set_can_shrink       = cc.DlFunc[func(pic *Picture, canShrink int32), cc.Void]{Name: "gtk_picture_set_can_shrink"}
	gtk_picture_get_can_shrink       = cc.DlFunc[func(pic *Picture) int32, int32]{Name: "gtk_picture_get_can_shrink"}
	gtk_picture_set_content_fit      = cc.DlFunc[func(pic *Picture, contentFit ContentFit), cc.Void]{Name: "gtk_picture_set_content_fit"}
	gtk_picture_get_content_fit      = cc.DlFunc[func(pic *Picture) ContentFit, ContentFit]{Name: "gtk_picture_get_content_fit"}
	gtk_picture_set_alternative_text = cc.DlFunc[func(pic *Picture, alternativeText cc.String), cc.Void]{Name: "gtk_picture_set_alternative_text"}
	gtk_picture_get_alternative_text = cc.DlFunc[func(pic *Picture) cc.String, cc.String]{Name: "gtk_picture_get_alternative_text"}
	// #endregion

	// #region Popover
	gtk_popover_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_popover_get_type"}
	gtk_popover_new                   = cc.DlFunc[func() *Popover, *Popover]{Name: "gtk_popover_new"}
	gtk_popover_set_child             = cc.DlFunc[func(popover *Popover, child *Widget), cc.Void]{Name: "gtk_popover_set_child"}
	gtk_popover_get_child             = cc.DlFunc[func(popover *Popover) *Widget, *Widget]{Name: "gtk_popover_get_child"}
	gtk_popover_set_pointing_to       = cc.DlFunc[func(popover *Popover, rect *gdk.Rectangle), cc.Void]{Name: "gtk_popover_set_pointing_to"}
	gtk_popover_get_pointing_to       = cc.DlFunc[func(popover *Popover, rect *gdk.Rectangle) int32, int32]{Name: "gtk_popover_get_pointing_to"}
	gtk_popover_set_position          = cc.DlFunc[func(popover *Popover, position PositionType), cc.Void]{Name: "gtk_popover_set_position"}
	gtk_popover_get_position          = cc.DlFunc[func(popover *Popover) PositionType, PositionType]{Name: "gtk_popover_get_position"}
	gtk_popover_set_autohide          = cc.DlFunc[func(popover *Popover, autohide int32), cc.Void]{Name: "gtk_popover_set_autohide"}
	gtk_popover_get_autohide          = cc.DlFunc[func(popover *Popover) int32, int32]{Name: "gtk_popover_get_autohide"}
	gtk_popover_set_has_arrow         = cc.DlFunc[func(popover *Popover, hasArrow int32), cc.Void]{Name: "gtk_popover_set_has_arrow"}
	gtk_popover_get_has_arrow         = cc.DlFunc[func(popover *Popover) int32, int32]{Name: "gtk_popover_get_has_arrow"}
	gtk_popover_set_mnemonics_visible = cc.DlFunc[func(popover *Popover, mnemonicsVisible int32), cc.Void]{Name: "gtk_popover_set_mnemonics_visible"}
	gtk_popover_get_mnemonics_visible = cc.DlFunc[func(popover *Popover) int32, int32]{Name: "gtk_popover_get_mnemonics_visible"}
	gtk_popover_popup                 = cc.DlFunc[func(popover *Popover), cc.Void]{Name: "gtk_popover_popup"}
	gtk_popover_popdown               = cc.DlFunc[func(popover *Popover), cc.Void]{Name: "gtk_popover_popdown"}
	gtk_popover_set_offset            = cc.DlFunc[func(popover *Popover, xOffset int32, yOffset int32), cc.Void]{Name: "gtk_popover_set_offset"}
	gtk_popover_get_offset            = cc.DlFunc[func(popover *Popover, xOffset *int32, yOffset *int32), cc.Void]{Name: "gtk_popover_get_offset"}
	gtk_popover_set_cascade_popdown   = cc.DlFunc[func(popover *Popover, cascadePopdown int32), cc.Void]{Name: "gtk_popover_set_cascade_popdown"}
	gtk_popover_get_cascade_popdown   = cc.DlFunc[func(popover *Popover) int32, int32]{Name: "gtk_popover_get_cascade_popdown"}
	gtk_popover_set_default_widget    = cc.DlFunc[func(popover *Popover, widget *Widget), cc.Void]{Name: "gtk_popover_set_default_widget"}
	gtk_popover_present               = cc.DlFunc[func(popover *Popover), cc.Void]{Name: "gtk_popover_present"}
	// #endregion

	// #region PopoverMenu
	gtk_popover_menu_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_popover_menu_get_type"}
	gtk_popover_menu_new_from_model      = cc.DlFunc[func(model *gio.GMenuModel) *PopoverMenu, *PopoverMenu]{Name: "gtk_popover_menu_new_from_model"}
	gtk_popover_menu_new_from_model_full = cc.DlFunc[func(model *gio.GMenuModel, flags PopoverMenuFlags) *PopoverMenu, *PopoverMenu]{Name: "gtk_popover_menu_new_from_model_full"}
	gtk_popover_menu_set_menu_model      = cc.DlFunc[func(popover *PopoverMenu, model *gio.GMenuModel), cc.Void]{Name: "gtk_popover_menu_set_menu_model"}
	gtk_popover_menu_get_menu_model      = cc.DlFunc[func(popover *PopoverMenu) *gio.GMenuModel, *gio.GMenuModel]{Name: "gtk_popover_menu_get_menu_model"}
	gtk_popover_menu_set_flags           = cc.DlFunc[func(popover *PopoverMenu, flags PopoverMenuFlags), cc.Void]{Name: "gtk_popover_menu_set_flags"}
	gtk_popover_menu_get_flags           = cc.DlFunc[func(popover *PopoverMenu) PopoverMenuFlags, PopoverMenuFlags]{Name: "gtk_popover_menu_get_flags"}
	gtk_popover_menu_add_child           = cc.DlFunc[func(popover *PopoverMenu, child *Widget, id cc.String) int32, int32]{Name: "gtk_popover_menu_add_child"}
	gtk_popover_menu_remove_child        = cc.DlFunc[func(popover *PopoverMenu, child *Widget) int32, int32]{Name: "gtk_popover_menu_remove_child"}
	// #endregion

	// #region PopoverMenuBar
	gtk_popover_menu_bar_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_popover_menu_bar_get_type"}
	gtk_popover_menu_bar_new_from_model = cc.DlFunc[func(model *gio.GMenuModel) *PopoverMenuBar, *PopoverMenuBar]{Name: "gtk_popover_menu_bar_new_from_model"}
	gtk_popover_menu_bar_set_menu_model = cc.DlFunc[func(bar *PopoverMenuBar, model *gio.GMenuModel), cc.Void]{Name: "gtk_popover_menu_bar_set_menu_model"}
	gtk_popover_menu_bar_get_menu_model = cc.DlFunc[func(bar *PopoverMenuBar) *gio.GMenuModel, *gio.GMenuModel]{Name: "gtk_popover_menu_bar_get_menu_model"}
	gtk_popover_menu_bar_add_child      = cc.DlFunc[func(bar *PopoverMenuBar, child *Widget, id cc.String) int32, int32]{Name: "gtk_popover_menu_bar_add_child"}
	gtk_popover_menu_bar_remove_child   = cc.DlFunc[func(bar *PopoverMenuBar, child *Widget) int32, int32]{Name: "gtk_popover_menu_bar_remove_child"}
	// #endregion

	// #region PrintDialog
	gtk_print_setup_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_print_setup_get_type"}
	gtk_print_setup_ref                = cc.DlFunc[func(setup *PrintSetup) *PrintSetup, *PrintSetup]{Name: "gtk_print_setup_ref"}
	gtk_print_setup_unref              = cc.DlFunc[func(setup *PrintSetup), cc.Void]{Name: "gtk_print_setup_unref"}
	gtk_print_setup_get_print_settings = cc.DlFunc[func(setup *PrintSetup) *PrintSettings, *PrintSettings]{Name: "gtk_print_setup_get_print_settings"}
	gtk_print_setup_get_page_setup     = cc.DlFunc[func(setup *PrintSetup) *PageSetup, *PageSetup]{Name: "gtk_print_setup_get_page_setup"}

	gtk_print_dialog_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_print_dialog_get_type"}
	gtk_print_dialog_new                = cc.DlFunc[func() *PrintDialog, *PrintDialog]{Name: "gtk_print_dialog_new"}
	gtk_print_dialog_get_title          = cc.DlFunc[func(dialog *PrintDialog) cc.String, cc.String]{Name: "gtk_print_dialog_get_title"}
	gtk_print_dialog_set_title          = cc.DlFunc[func(dialog *PrintDialog, title cc.String), cc.Void]{Name: "gtk_print_dialog_set_title"}
	gtk_print_dialog_get_accept_label   = cc.DlFunc[func(dialog *PrintDialog) cc.String, cc.String]{Name: "gtk_print_dialog_get_accept_label"}
	gtk_print_dialog_set_accept_label   = cc.DlFunc[func(dialog *PrintDialog, acceptLabel cc.String), cc.Void]{Name: "gtk_print_dialog_set_accept_label"}
	gtk_print_dialog_get_modal          = cc.DlFunc[func(dialog *PrintDialog) int32, int32]{Name: "gtk_print_dialog_get_modal"}
	gtk_print_dialog_set_modal          = cc.DlFunc[func(dialog *PrintDialog, modal int32), cc.Void]{Name: "gtk_print_dialog_set_modal"}
	gtk_print_dialog_get_page_setup     = cc.DlFunc[func(dialog *PrintDialog) *PageSetup, *PageSetup]{Name: "gtk_print_dialog_get_page_setup"}
	gtk_print_dialog_set_page_setup     = cc.DlFunc[func(dialog *PrintDialog, pageSetup *PageSetup), cc.Void]{Name: "gtk_print_dialog_set_page_setup"}
	gtk_print_dialog_get_print_settings = cc.DlFunc[func(dialog *PrintDialog) *PrintSettings, *PrintSettings]{Name: "gtk_print_dialog_get_print_settings"}
	gtk_print_dialog_set_print_settings = cc.DlFunc[func(dialog *PrintDialog, printSettings *PrintSettings), cc.Void]{Name: "gtk_print_dialog_set_print_settings"}
	gtk_print_dialog_setup              = cc.DlFunc[func(dialog *PrintDialog, parent *Window, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_print_dialog_setup"}
	gtk_print_dialog_setup_finish       = cc.DlFunc[func(dialog *PrintDialog, result *gio.GAsyncResult, err **glib.GError) *PrintSetup, *PrintSetup]{Name: "gtk_print_dialog_setup_finish"}
	gtk_print_dialog_print              = cc.DlFunc[func(dialog *PrintDialog, parent *Window, setup *PrintSetup, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_print_dialog_print"}
	gtk_print_dialog_print_finish       = cc.DlFunc[func(dialog *PrintDialog, result *gio.GAsyncResult, err **glib.GError) *gio.GOutputStream, *gio.GOutputStream]{Name: "gtk_print_dialog_print_finish"}
	gtk_print_dialog_print_file         = cc.DlFunc[func(dialog *PrintDialog, parent *Window, setup *PrintSetup, file *gio.GFile, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_print_dialog_print_file"}
	gtk_print_dialog_print_file_finish  = cc.DlFunc[func(dialog *PrintDialog, result *gio.GAsyncResult, err **glib.GError) int32, int32]{Name: "gtk_print_dialog_print_file_finish"}
	// #endregion

	// #region ProgressBar
	gtk_progress_bar_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_progress_bar_get_type"}
	gtk_progress_bar_new            = cc.DlFunc[func() *ProgressBar, *ProgressBar]{Name: "gtk_progress_bar_new"}
	gtk_progress_bar_pulse          = cc.DlFunc[func(pbar *ProgressBar), cc.Void]{Name: "gtk_progress_bar_pulse"}
	gtk_progress_bar_set_text       = cc.DlFunc[func(pbar *ProgressBar, text cc.String), cc.Void]{Name: "gtk_progress_bar_set_text"}
	gtk_progress_bar_set_fraction   = cc.DlFunc[func(pbar *ProgressBar, fraction float64), cc.Void]{Name: "gtk_progress_bar_set_fraction"}
	gtk_progress_bar_set_pulse_step = cc.DlFunc[func(pbar *ProgressBar, fraction float64), cc.Void]{Name: "gtk_progress_bar_set_pulse_step"}
	gtk_progress_bar_set_inverted   = cc.DlFunc[func(pbar *ProgressBar, inverted int32), cc.Void]{Name: "gtk_progress_bar_set_inverted"}
	gtk_progress_bar_get_text       = cc.DlFunc[func(pbar *ProgressBar) cc.String, cc.String]{Name: "gtk_progress_bar_get_text"}
	gtk_progress_bar_get_fraction   = cc.DlFunc[func(pbar *ProgressBar) float64, float64]{Name: "gtk_progress_bar_get_fraction"}
	gtk_progress_bar_get_pulse_step = cc.DlFunc[func(pbar *ProgressBar) float64, float64]{Name: "gtk_progress_bar_get_pulse_step"}
	gtk_progress_bar_get_inverted   = cc.DlFunc[func(pbar *ProgressBar) int32, int32]{Name: "gtk_progress_bar_get_inverted"}
	gtk_progress_bar_set_ellipsize  = cc.DlFunc[func(pbar *ProgressBar, mode pango.EllipsizeMode), cc.Void]{Name: "gtk_progress_bar_set_ellipsize"}
	gtk_progress_bar_get_ellipsize  = cc.DlFunc[func(pbar *ProgressBar) pango.EllipsizeMode, pango.EllipsizeMode]{Name: "gtk_progress_bar_get_ellipsize"}
	gtk_progress_bar_set_show_text  = cc.DlFunc[func(pbar *ProgressBar, showText int32), cc.Void]{Name: "gtk_progress_bar_set_show_text"}
	gtk_progress_bar_get_show_text  = cc.DlFunc[func(pbar *ProgressBar) int32, int32]{Name: "gtk_progress_bar_get_show_text"}
	// #endregion

	// #region Range
	gtk_range_get_type                   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_range_get_type"}
	gtk_range_set_adjustment             = cc.DlFunc[func(r *Range, adj *Adjustment), cc.Void]{Name: "gtk_range_set_adjustment"}
	gtk_range_get_adjustment             = cc.DlFunc[func(r *Range) *Adjustment, *Adjustment]{Name: "gtk_range_get_adjustment"}
	gtk_range_set_inverted               = cc.DlFunc[func(r *Range, setting int32), cc.Void]{Name: "gtk_range_set_inverted"}
	gtk_range_get_inverted               = cc.DlFunc[func(r *Range) int32, int32]{Name: "gtk_range_get_inverted"}
	gtk_range_set_flippable              = cc.DlFunc[func(r *Range, flippable int32), cc.Void]{Name: "gtk_range_set_flippable"}
	gtk_range_get_flippable              = cc.DlFunc[func(r *Range) int32, int32]{Name: "gtk_range_get_flippable"}
	gtk_range_set_slider_size_fixed      = cc.DlFunc[func(r *Range, sizeFixed int32), cc.Void]{Name: "gtk_range_set_slider_size_fixed"}
	gtk_range_get_slider_size_fixed      = cc.DlFunc[func(r *Range) int32, int32]{Name: "gtk_range_get_slider_size_fixed"}
	gtk_range_get_range_rect             = cc.DlFunc[func(r *Range, rangeRect *gdk.Rectangle), cc.Void]{Name: "gtk_range_get_range_rect"}
	gtk_range_get_slider_range           = cc.DlFunc[func(r *Range, sliderStart *int32, sliderEnd *int32), cc.Void]{Name: "gtk_range_get_slider_range"}
	gtk_range_set_increments             = cc.DlFunc[func(r *Range, step, page float64), cc.Void]{Name: "gtk_range_set_increments"}
	gtk_range_set_range                  = cc.DlFunc[func(r *Range, min, max float64), cc.Void]{Name: "gtk_range_set_range"}
	gtk_range_set_value                  = cc.DlFunc[func(r *Range, value float64), cc.Void]{Name: "gtk_range_set_value"}
	gtk_range_get_value                  = cc.DlFunc[func(r *Range) float64, float64]{Name: "gtk_range_get_value"}
	gtk_range_set_show_fill_level        = cc.DlFunc[func(r *Range, showFillLevel int32), cc.Void]{Name: "gtk_range_set_show_fill_level"}
	gtk_range_get_show_fill_level        = cc.DlFunc[func(r *Range) int32, int32]{Name: "gtk_range_get_show_fill_level"}
	gtk_range_set_restrict_to_fill_level = cc.DlFunc[func(r *Range, restrict int32), cc.Void]{Name: "gtk_range_set_restrict_to_fill_level"}
	gtk_range_get_restrict_to_fill_level = cc.DlFunc[func(r *Range) int32, int32]{Name: "gtk_range_get_restrict_to_fill_level"}
	gtk_range_set_fill_level             = cc.DlFunc[func(r *Range, fillLevel float64), cc.Void]{Name: "gtk_range_set_fill_level"}
	gtk_range_get_fill_level             = cc.DlFunc[func(r *Range) float64, float64]{Name: "gtk_range_get_fill_level"}
	gtk_range_set_round_digits           = cc.DlFunc[func(r *Range, roundDigits int32), cc.Void]{Name: "gtk_range_set_round_digits"}
	gtk_range_get_round_digits           = cc.DlFunc[func(r *Range) int32, int32]{Name: "gtk_range_get_round_digits"}
	// #endregion

	// #region RecentManager
	gtk_recent_manager_error_quark = cc.DlFunc[func() glib.GQuark, glib.GQuark]{Name: "gtk_recent_manager_error_quark"}
	gtk_recent_manager_get_type    = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_recent_manager_get_type"}
	gtk_recent_manager_new         = cc.DlFunc[func() *RecentManager, *RecentManager]{Name: "gtk_recent_manager_new"}
	gtk_recent_manager_get_default = cc.DlFunc[func() *RecentManager, *RecentManager]{Name: "gtk_recent_manager_get_default"}
	gtk_recent_manager_add_item    = cc.DlFunc[func(manager *RecentManager, uri cc.String) int32, int32]{Name: "gtk_recent_manager_add_item"}
	gtk_recent_manager_add_full    = cc.DlFunc[func(manager *RecentManager, uri cc.String, recentData *RecentData) int32, int32]{Name: "gtk_recent_manager_add_full"}
	gtk_recent_manager_remove_item = cc.DlFunc[func(manager *RecentManager, uri cc.String, err **glib.GError) int32, int32]{Name: "gtk_recent_manager_remove_item"}
	gtk_recent_manager_lookup_item = cc.DlFunc[func(manager *RecentManager, uri cc.String, err **glib.GError) *RecentInfo, *RecentInfo]{Name: "gtk_recent_manager_lookup_item"}
	gtk_recent_manager_has_item    = cc.DlFunc[func(manager *RecentManager, uri cc.String) int32, int32]{Name: "gtk_recent_manager_has_item"}
	gtk_recent_manager_move_item   = cc.DlFunc[func(manager *RecentManager, uri cc.String, newUri cc.String, err **glib.GError) int32, int32]{Name: "gtk_recent_manager_move_item"}
	gtk_recent_manager_get_items   = cc.DlFunc[func(manager *RecentManager) *glib.GList[RecentInfo], *glib.GList[RecentInfo]]{Name: "gtk_recent_manager_get_items"}
	gtk_recent_manager_purge_items = cc.DlFunc[func(manager *RecentManager, err **glib.GError) int32, int32]{Name: "gtk_recent_manager_purge_items"}

	gtk_recent_info_get_type             = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_recent_info_get_type"}
	gtk_recent_info_ref                  = cc.DlFunc[func(info *RecentInfo) *RecentInfo, *RecentInfo]{Name: "gtk_recent_info_ref"}
	gtk_recent_info_unref                = cc.DlFunc[func(info *RecentInfo), cc.Void]{Name: "gtk_recent_info_unref"}
	gtk_recent_info_get_uri              = cc.DlFunc[func(info *RecentInfo) cc.String, cc.String]{Name: "gtk_recent_info_get_uri"}
	gtk_recent_info_get_display_name     = cc.DlFunc[func(info *RecentInfo) cc.String, cc.String]{Name: "gtk_recent_info_get_display_name"}
	gtk_recent_info_get_description      = cc.DlFunc[func(info *RecentInfo) cc.String, cc.String]{Name: "gtk_recent_info_get_description"}
	gtk_recent_info_get_mime_type        = cc.DlFunc[func(info *RecentInfo) cc.String, cc.String]{Name: "gtk_recent_info_get_mime_type"}
	gtk_recent_info_get_added            = cc.DlFunc[func(info *RecentInfo) *glib.GDateTime, *glib.GDateTime]{Name: "gtk_recent_info_get_added"}
	gtk_recent_info_get_modified         = cc.DlFunc[func(info *RecentInfo) *glib.GDateTime, *glib.GDateTime]{Name: "gtk_recent_info_get_modified"}
	gtk_recent_info_get_visited          = cc.DlFunc[func(info *RecentInfo) *glib.GDateTime, *glib.GDateTime]{Name: "gtk_recent_info_get_visited"}
	gtk_recent_info_get_private_hint     = cc.DlFunc[func(info *RecentInfo) int32, int32]{Name: "gtk_recent_info_get_private_hint"}
	gtk_recent_info_get_application_info = cc.DlFunc[func(info *RecentInfo, appName cc.String, appExec *cc.String, count *uint32, stamp **glib.GDateTime) int32, int32]{Name: "gtk_recent_info_get_application_info"}
	gtk_recent_info_create_app_info      = cc.DlFunc[func(info *RecentInfo, appName cc.String, err **glib.GError) *gio.GAppInfo, *gio.GAppInfo]{Name: "gtk_recent_info_create_app_info"}
	gtk_recent_info_get_applications     = cc.DlFunc[func(info *RecentInfo, length *uint64) cc.Strings, cc.Strings]{Name: "gtk_recent_info_get_applications"}
	gtk_recent_info_last_application     = cc.DlFunc[func(info *RecentInfo) cc.String, cc.String]{Name: "gtk_recent_info_last_application"}
	gtk_recent_info_has_application      = cc.DlFunc[func(info *RecentInfo, appName cc.String) int32, int32]{Name: "gtk_recent_info_has_application"}
	gtk_recent_info_get_groups           = cc.DlFunc[func(info *RecentInfo, length *uint64) cc.Strings, cc.Strings]{Name: "gtk_recent_info_get_groups"}
	gtk_recent_info_has_group            = cc.DlFunc[func(info *RecentInfo, groupName cc.String) int32, int32]{Name: "gtk_recent_info_has_group"}
	gtk_recent_info_get_gicon            = cc.DlFunc[func(info *RecentInfo) *gio.GIcon, *gio.GIcon]{Name: "gtk_recent_info_get_gicon"}
	gtk_recent_info_get_short_name       = cc.DlFunc[func(info *RecentInfo) cc.String, cc.String]{Name: "gtk_recent_info_get_short_name"}
	gtk_recent_info_get_uri_display      = cc.DlFunc[func(info *RecentInfo) cc.String, cc.String]{Name: "gtk_recent_info_get_uri_display"}
	gtk_recent_info_get_age              = cc.DlFunc[func(info *RecentInfo) int32, int32]{Name: "gtk_recent_info_get_age"}
	gtk_recent_info_is_local             = cc.DlFunc[func(info *RecentInfo) int32, int32]{Name: "gtk_recent_info_is_local"}
	gtk_recent_info_exists               = cc.DlFunc[func(info *RecentInfo) int32, int32]{Name: "gtk_recent_info_exists"}
	gtk_recent_info_match                = cc.DlFunc[func(infoA *RecentInfo, infoB *RecentInfo) int32, int32]{Name: "gtk_recent_info_match"}
	// #endregion

	// #region Revealer
	gtk_revealer_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_revealer_get_type"}
	gtk_revealer_new                     = cc.DlFunc[func() *Revealer, *Revealer]{Name: "gtk_revealer_new"}
	gtk_revealer_get_reveal_child        = cc.DlFunc[func(revealer *Revealer) int32, int32]{Name: "gtk_revealer_get_reveal_child"}
	gtk_revealer_set_reveal_child        = cc.DlFunc[func(revealer *Revealer, revealChild int32), cc.Void]{Name: "gtk_revealer_set_reveal_child"}
	gtk_revealer_get_child_revealed      = cc.DlFunc[func(revealer *Revealer) int32, int32]{Name: "gtk_revealer_get_child_revealed"}
	gtk_revealer_get_transition_duration = cc.DlFunc[func(revealer *Revealer) uint32, uint32]{Name: "gtk_revealer_get_transition_duration"}
	gtk_revealer_set_transition_duration = cc.DlFunc[func(revealer *Revealer, duration uint32), cc.Void]{Name: "gtk_revealer_set_transition_duration"}
	gtk_revealer_set_transition_type     = cc.DlFunc[func(revealer *Revealer, transition RevealerTransitionType), cc.Void]{Name: "gtk_revealer_set_transition_type"}
	gtk_revealer_get_transition_type     = cc.DlFunc[func(revealer *Revealer) RevealerTransitionType, RevealerTransitionType]{Name: "gtk_revealer_get_transition_type"}
	gtk_revealer_set_child               = cc.DlFunc[func(revealer *Revealer, child *Widget), cc.Void]{Name: "gtk_revealer_set_child"}
	gtk_revealer_get_child               = cc.DlFunc[func(revealer *Revealer) *Widget, *Widget]{Name: "gtk_revealer_get_child"}
	// #endregion

	// #region Root
	gtk_root_get_display = cc.DlFunc[func(*Root) *gdk.Display, *gdk.Display]{Name: "gtk_root_get_display"}
	gtk_root_set_focus   = cc.DlFunc[func(*Root, *Widget), cc.Void]{Name: "gtk_root_set_focus"}
	gtk_root_get_focus   = cc.DlFunc[func(*Root) *Widget, *Widget]{Name: "gtk_root_get_focus"}
	// #endregion

	// #region Scale
	gtk_scale_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_scale_get_type"}
	gtk_scale_new                   = cc.DlFunc[func(orientation Orientation, adjustment *Adjustment) *Scale, *Scale]{Name: "gtk_scale_new"}
	gtk_scale_new_with_range        = cc.DlFunc[func(orientation Orientation, min, max, step float64) *Scale, *Scale]{Name: "gtk_scale_new_with_range"}
	gtk_scale_set_digits            = cc.DlFunc[func(scale *Scale, digits int32), cc.Void]{Name: "gtk_scale_set_digits"}
	gtk_scale_get_digits            = cc.DlFunc[func(scale *Scale) int32, int32]{Name: "gtk_scale_get_digits"}
	gtk_scale_set_draw_value        = cc.DlFunc[func(scale *Scale, drawValue int32), cc.Void]{Name: "gtk_scale_set_draw_value"}
	gtk_scale_get_draw_value        = cc.DlFunc[func(scale *Scale) int32, int32]{Name: "gtk_scale_get_draw_value"}
	gtk_scale_set_has_origin        = cc.DlFunc[func(scale *Scale, hasOrigin int32), cc.Void]{Name: "gtk_scale_set_has_origin"}
	gtk_scale_get_has_origin        = cc.DlFunc[func(scale *Scale) int32, int32]{Name: "gtk_scale_get_has_origin"}
	gtk_scale_set_value_pos         = cc.DlFunc[func(scale *Scale, pos PositionType), cc.Void]{Name: "gtk_scale_set_value_pos"}
	gtk_scale_get_value_pos         = cc.DlFunc[func(scale *Scale) PositionType, PositionType]{Name: "gtk_scale_get_value_pos"}
	gtk_scale_get_layout            = cc.DlFunc[func(scale *Scale) *pango.Layout, *pango.Layout]{Name: "gtk_scale_get_layout"}
	gtk_scale_get_layout_offsets    = cc.DlFunc[func(scale *Scale, x, y *int32), cc.Void]{Name: "gtk_scale_get_layout_offsets"}
	gtk_scale_add_mark              = cc.DlFunc[func(scale *Scale, value float64, position PositionType, markup cc.String), cc.Void]{Name: "gtk_scale_add_mark"}
	gtk_scale_clear_marks           = cc.DlFunc[func(scale *Scale), cc.Void]{Name: "gtk_scale_clear_marks"}
	gtk_scale_set_format_value_func = cc.DlFunc[func(scale *Scale, fn uptr, userData uptr, destroyNotify uptr), cc.Void]{Name: "gtk_scale_set_format_value_func"}
	// #endregion

	// #region ScaleButton
	gtk_scale_button_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_scale_button_get_type"}
	gtk_scale_button_new              = cc.DlFunc[func(min, max, step float64, icons cc.Strings) *ScaleButton, *ScaleButton]{Name: "gtk_scale_button_new"}
	gtk_scale_button_set_icons        = cc.DlFunc[func(button *ScaleButton, icons cc.Strings), cc.Void]{Name: "gtk_scale_button_set_icons"}
	gtk_scale_button_get_value        = cc.DlFunc[func(button *ScaleButton) float64, float64]{Name: "gtk_scale_button_get_value"}
	gtk_scale_button_set_value        = cc.DlFunc[func(button *ScaleButton, value float64), cc.Void]{Name: "gtk_scale_button_set_value"}
	gtk_scale_button_get_adjustment   = cc.DlFunc[func(button *ScaleButton) *Adjustment, *Adjustment]{Name: "gtk_scale_button_get_adjustment"}
	gtk_scale_button_set_adjustment   = cc.DlFunc[func(button *ScaleButton, adjustment *Adjustment), cc.Void]{Name: "gtk_scale_button_set_adjustment"}
	gtk_scale_button_get_plus_button  = cc.DlFunc[func(button *ScaleButton) *Widget, *Widget]{Name: "gtk_scale_button_get_plus_button"}
	gtk_scale_button_get_minus_button = cc.DlFunc[func(button *ScaleButton) *Widget, *Widget]{Name: "gtk_scale_button_get_minus_button"}
	gtk_scale_button_get_popup        = cc.DlFunc[func(button *ScaleButton) *Widget, *Widget]{Name: "gtk_scale_button_get_popup"}
	gtk_scale_button_get_active       = cc.DlFunc[func(button *ScaleButton) int32, int32]{Name: "gtk_scale_button_get_active"}
	gtk_scale_button_get_has_frame    = cc.DlFunc[func(button *ScaleButton) int32, int32]{Name: "gtk_scale_button_get_has_frame"}
	gtk_scale_button_set_has_frame    = cc.DlFunc[func(button *ScaleButton, hasFrame int32), cc.Void]{Name: "gtk_scale_button_set_has_frame"}
	// #endregion

	// #region Scrollable
	gtk_scrollable_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_scrollable_get_type"}
	gtk_scrollable_get_hadjustment    = cc.DlFunc[func(scrollable *Scrollable) *Adjustment, *Adjustment]{Name: "gtk_scrollable_get_hadjustment"}
	gtk_scrollable_set_hadjustment    = cc.DlFunc[func(scrollable *Scrollable, hadjustment *Adjustment), cc.Void]{Name: "gtk_scrollable_set_hadjustment"}
	gtk_scrollable_get_vadjustment    = cc.DlFunc[func(scrollable *Scrollable) *Adjustment, *Adjustment]{Name: "gtk_scrollable_get_vadjustment"}
	gtk_scrollable_set_vadjustment    = cc.DlFunc[func(scrollable *Scrollable, vadjustment *Adjustment), cc.Void]{Name: "gtk_scrollable_set_vadjustment"}
	gtk_scrollable_get_hscroll_policy = cc.DlFunc[func(scrollable *Scrollable) ScrollablePolicy, ScrollablePolicy]{Name: "gtk_scrollable_get_hscroll_policy"}
	gtk_scrollable_set_hscroll_policy = cc.DlFunc[func(scrollable *Scrollable, policy ScrollablePolicy), cc.Void]{Name: "gtk_scrollable_set_hscroll_policy"}
	gtk_scrollable_get_vscroll_policy = cc.DlFunc[func(scrollable *Scrollable) ScrollablePolicy, ScrollablePolicy]{Name: "gtk_scrollable_get_vscroll_policy"}
	gtk_scrollable_set_vscroll_policy = cc.DlFunc[func(scrollable *Scrollable, policy ScrollablePolicy), cc.Void]{Name: "gtk_scrollable_set_vscroll_policy"}
	gtk_scrollable_get_border         = cc.DlFunc[func(scrollable *Scrollable, border *Border) int32, int32]{Name: "gtk_scrollable_get_border"}
	// #endregion

	// #region Scrollbar
	gtk_scrollbar_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_scrollbar_get_type"}
	gtk_scrollbar_new            = cc.DlFunc[func(orientation Orientation, adjustment *Adjustment) *Scrollbar, *Scrollbar]{Name: "gtk_scrollbar_new"}
	gtk_scrollbar_set_adjustment = cc.DlFunc[func(sb *Scrollbar, adjustment *Adjustment), cc.Void]{Name: "gtk_scrollbar_set_adjustment"}
	gtk_scrollbar_get_adjustment = cc.DlFunc[func(sb *Scrollbar) *Adjustment, *Adjustment]{Name: "gtk_scrollbar_get_adjustment"}
	// #endregion

	// #region ScrolledWindow
	gtk_scrolled_window_get_type                     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_scrolled_window_get_type"}
	gtk_scrolled_window_new                          = cc.DlFunc[func() *ScrolledWindow, *ScrolledWindow]{Name: "gtk_scrolled_window_new"}
	gtk_scrolled_window_set_hadjustment              = cc.DlFunc[func(w *ScrolledWindow, hadjustment *Adjustment), cc.Void]{Name: "gtk_scrolled_window_set_hadjustment"}
	gtk_scrolled_window_set_vadjustment              = cc.DlFunc[func(w *ScrolledWindow, vadjustment *Adjustment), cc.Void]{Name: "gtk_scrolled_window_set_vadjustment"}
	gtk_scrolled_window_get_hadjustment              = cc.DlFunc[func(w *ScrolledWindow) *Adjustment, *Adjustment]{Name: "gtk_scrolled_window_get_hadjustment"}
	gtk_scrolled_window_get_vadjustment              = cc.DlFunc[func(w *ScrolledWindow) *Adjustment, *Adjustment]{Name: "gtk_scrolled_window_get_vadjustment"}
	gtk_scrolled_window_get_hscrollbar               = cc.DlFunc[func(w *ScrolledWindow) *Scrollbar, *Scrollbar]{Name: "gtk_scrolled_window_get_hscrollbar"}
	gtk_scrolled_window_get_vscrollbar               = cc.DlFunc[func(w *ScrolledWindow) *Scrollbar, *Scrollbar]{Name: "gtk_scrolled_window_get_vscrollbar"}
	gtk_scrolled_window_set_policy                   = cc.DlFunc[func(w *ScrolledWindow, hscrollbarPolicy, vscrollbarPolicy PolicyType), cc.Void]{Name: "gtk_scrolled_window_set_policy"}
	gtk_scrolled_window_get_policy                   = cc.DlFunc[func(w *ScrolledWindow, hscrollbarPolicy, vscrollbarPolicy *PolicyType), cc.Void]{Name: "gtk_scrolled_window_get_policy"}
	gtk_scrolled_window_set_placement                = cc.DlFunc[func(w *ScrolledWindow, placement CornerType), cc.Void]{Name: "gtk_scrolled_window_set_placement"}
	gtk_scrolled_window_unset_placement              = cc.DlFunc[func(w *ScrolledWindow), cc.Void]{Name: "gtk_scrolled_window_unset_placement"}
	gtk_scrolled_window_get_placement                = cc.DlFunc[func(w *ScrolledWindow) CornerType, CornerType]{Name: "gtk_scrolled_window_get_placement"}
	gtk_scrolled_window_set_has_frame                = cc.DlFunc[func(w *ScrolledWindow, hasFrame int32), cc.Void]{Name: "gtk_scrolled_window_set_has_frame"}
	gtk_scrolled_window_get_has_frame                = cc.DlFunc[func(w *ScrolledWindow) int32, int32]{Name: "gtk_scrolled_window_get_has_frame"}
	gtk_scrolled_window_get_min_content_width        = cc.DlFunc[func(w *ScrolledWindow) int32, int32]{Name: "gtk_scrolled_window_get_min_content_width"}
	gtk_scrolled_window_set_min_content_width        = cc.DlFunc[func(w *ScrolledWindow, width int32), cc.Void]{Name: "gtk_scrolled_window_set_min_content_width"}
	gtk_scrolled_window_get_min_content_height       = cc.DlFunc[func(w *ScrolledWindow) int32, int32]{Name: "gtk_scrolled_window_get_min_content_height"}
	gtk_scrolled_window_set_min_content_height       = cc.DlFunc[func(w *ScrolledWindow, height int32), cc.Void]{Name: "gtk_scrolled_window_set_min_content_height"}
	gtk_scrolled_window_set_kinetic_scrolling        = cc.DlFunc[func(w *ScrolledWindow, kineticScrolling int32), cc.Void]{Name: "gtk_scrolled_window_set_kinetic_scrolling"}
	gtk_scrolled_window_get_kinetic_scrolling        = cc.DlFunc[func(w *ScrolledWindow) int32, int32]{Name: "gtk_scrolled_window_get_kinetic_scrolling"}
	gtk_scrolled_window_set_overlay_scrolling        = cc.DlFunc[func(w *ScrolledWindow, overlayScrolling int32), cc.Void]{Name: "gtk_scrolled_window_set_overlay_scrolling"}
	gtk_scrolled_window_get_overlay_scrolling        = cc.DlFunc[func(w *ScrolledWindow) int32, int32]{Name: "gtk_scrolled_window_get_overlay_scrolling"}
	gtk_scrolled_window_set_max_content_width        = cc.DlFunc[func(w *ScrolledWindow, width int32), cc.Void]{Name: "gtk_scrolled_window_set_max_content_width"}
	gtk_scrolled_window_get_max_content_width        = cc.DlFunc[func(w *ScrolledWindow) int32, int32]{Name: "gtk_scrolled_window_get_max_content_width"}
	gtk_scrolled_window_set_max_content_height       = cc.DlFunc[func(w *ScrolledWindow, height int32), cc.Void]{Name: "gtk_scrolled_window_set_max_content_height"}
	gtk_scrolled_window_get_max_content_height       = cc.DlFunc[func(w *ScrolledWindow) int32, int32]{Name: "gtk_scrolled_window_get_max_content_height"}
	gtk_scrolled_window_set_propagate_natural_width  = cc.DlFunc[func(w *ScrolledWindow, propagate int32), cc.Void]{Name: "gtk_scrolled_window_set_propagate_natural_width"}
	gtk_scrolled_window_get_propagate_natural_width  = cc.DlFunc[func(w *ScrolledWindow) int32, int32]{Name: "gtk_scrolled_window_get_propagate_natural_width"}
	gtk_scrolled_window_set_propagate_natural_height = cc.DlFunc[func(w *ScrolledWindow, propagate int32), cc.Void]{Name: "gtk_scrolled_window_set_propagate_natural_height"}
	gtk_scrolled_window_get_propagate_natural_height = cc.DlFunc[func(w *ScrolledWindow) int32, int32]{Name: "gtk_scrolled_window_get_propagate_natural_height"}
	gtk_scrolled_window_set_child                    = cc.DlFunc[func(w *ScrolledWindow, child *Widget), cc.Void]{Name: "gtk_scrolled_window_set_child"}
	gtk_scrolled_window_get_child                    = cc.DlFunc[func(w *ScrolledWindow) *Widget, *Widget]{Name: "gtk_scrolled_window_get_child"}
	// #endregion

	// #region ScrollInfo
	gtk_scroll_info_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_scroll_info_get_type"}
	gtk_scroll_info_new                   = cc.DlFunc[func() *ScrollInfo, *ScrollInfo]{Name: "gtk_scroll_info_new"}
	gtk_scroll_info_ref                   = cc.DlFunc[func(self *ScrollInfo) *ScrollInfo, *ScrollInfo]{Name: "gtk_scroll_info_ref"}
	gtk_scroll_info_unref                 = cc.DlFunc[func(self *ScrollInfo), cc.Void]{Name: "gtk_scroll_info_unref"}
	gtk_scroll_info_set_enable_horizontal = cc.DlFunc[func(self *ScrollInfo, horizontal int32), cc.Void]{Name: "gtk_scroll_info_set_enable_horizontal"}
	gtk_scroll_info_get_enable_horizontal = cc.DlFunc[func(self *ScrollInfo) int32, int32]{Name: "gtk_scroll_info_get_enable_horizontal"}
	gtk_scroll_info_set_enable_vertical   = cc.DlFunc[func(self *ScrollInfo, vertical int32), cc.Void]{Name: "gtk_scroll_info_set_enable_vertical"}
	gtk_scroll_info_get_enable_vertical   = cc.DlFunc[func(self *ScrollInfo) int32, int32]{Name: "gtk_scroll_info_get_enable_vertical"}
	// #endregion

	// #region SearchBar
	gtk_search_bar_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_search_bar_get_type"}
	gtk_search_bar_new                    = cc.DlFunc[func() *SearchBar, *SearchBar]{Name: "gtk_search_bar_new"}
	gtk_search_bar_connect_entry          = cc.DlFunc[func(bar *SearchBar, entry *Editable), cc.Void]{Name: "gtk_search_bar_connect_entry"}
	gtk_search_bar_get_search_mode        = cc.DlFunc[func(bar *SearchBar) int32, int32]{Name: "gtk_search_bar_get_search_mode"}
	gtk_search_bar_set_search_mode        = cc.DlFunc[func(bar *SearchBar, searchMode int32), cc.Void]{Name: "gtk_search_bar_set_search_mode"}
	gtk_search_bar_get_show_close_button  = cc.DlFunc[func(bar *SearchBar) int32, int32]{Name: "gtk_search_bar_get_show_close_button"}
	gtk_search_bar_set_show_close_button  = cc.DlFunc[func(bar *SearchBar, visible int32), cc.Void]{Name: "gtk_search_bar_set_show_close_button"}
	gtk_search_bar_set_key_capture_widget = cc.DlFunc[func(bar *SearchBar, widget *Widget), cc.Void]{Name: "gtk_search_bar_set_key_capture_widget"}
	gtk_search_bar_get_key_capture_widget = cc.DlFunc[func(bar *SearchBar) *Widget, *Widget]{Name: "gtk_search_bar_get_key_capture_widget"}
	gtk_search_bar_set_child              = cc.DlFunc[func(bar *SearchBar, child *Widget), cc.Void]{Name: "gtk_search_bar_set_child"}
	gtk_search_bar_get_child              = cc.DlFunc[func(bar *SearchBar) *Widget, *Widget]{Name: "gtk_search_bar_get_child"}
	// #endregion

	// #region SearchEntry
	gtk_search_entry_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_search_entry_get_type"}
	gtk_search_entry_new                    = cc.DlFunc[func() *SearchEntry, *SearchEntry]{Name: "gtk_search_entry_new"}
	gtk_search_entry_set_key_capture_widget = cc.DlFunc[func(entry *SearchEntry, widget *Widget), cc.Void]{Name: "gtk_search_entry_set_key_capture_widget"}
	gtk_search_entry_get_key_capture_widget = cc.DlFunc[func(entry *SearchEntry) *Widget, *Widget]{Name: "gtk_search_entry_get_key_capture_widget"}
	gtk_search_entry_set_search_delay       = cc.DlFunc[func(entry *SearchEntry, delay uint32), cc.Void]{Name: "gtk_search_entry_set_search_delay"}
	gtk_search_entry_get_search_delay       = cc.DlFunc[func(entry *SearchEntry) uint32, uint32]{Name: "gtk_search_entry_get_search_delay"}
	gtk_search_entry_set_placeholder_text   = cc.DlFunc[func(entry *SearchEntry, text cc.String), cc.Void]{Name: "gtk_search_entry_set_placeholder_text"}
	gtk_search_entry_get_placeholder_text   = cc.DlFunc[func(entry *SearchEntry) cc.String, cc.String]{Name: "gtk_search_entry_get_placeholder_text"}
	gtk_search_entry_set_input_purpose      = cc.DlFunc[func(entry *SearchEntry, purpose InputPurpose), cc.Void]{Name: "gtk_search_entry_set_input_purpose"}
	gtk_search_entry_get_input_purpose      = cc.DlFunc[func(entry *SearchEntry) InputPurpose, InputPurpose]{Name: "gtk_search_entry_get_input_purpose"}
	gtk_search_entry_set_input_hints        = cc.DlFunc[func(entry *SearchEntry, hints InputHints), cc.Void]{Name: "gtk_search_entry_set_input_hints"}
	gtk_search_entry_get_input_hints        = cc.DlFunc[func(entry *SearchEntry) InputHints, InputHints]{Name: "gtk_search_entry_get_input_hints"}
	// #endregion

	// #region SectionModel
	gtk_section_model_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_section_model_get_type"}
	gtk_section_model_get_section      = cc.DlFunc[func(self *SectionModel, position uint32, outStart, outEnd *uint32), cc.Void]{Name: "gtk_section_model_get_section"}
	gtk_section_model_sections_changed = cc.DlFunc[func(self *SectionModel, position, nItems uint32), cc.Void]{Name: "gtk_section_model_sections_changed"}
	// #endregion

	// #region SelectionFilterModel
	gtk_selection_filter_model_new       = cc.DlFunc[func(model *SelectionModel) *SelectionFilterModel, *SelectionFilterModel]{Name: "gtk_selection_filter_model_new"}
	gtk_selection_filter_model_set_model = cc.DlFunc[func(self *SelectionFilterModel, model *SelectionModel), cc.Void]{Name: "gtk_selection_filter_model_set_model"}
	gtk_selection_filter_model_get_model = cc.DlFunc[func(self *SelectionFilterModel) *SelectionModel, *SelectionModel]{Name: "gtk_selection_filter_model_get_model"}
	// #endregion

	// #region SelectionModel
	gtk_selection_model_get_type               = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_selection_model_get_type"}
	gtk_selection_model_is_selected            = cc.DlFunc[func(model *SelectionModel, position uint32) int32, int32]{Name: "gtk_selection_model_is_selected"}
	gtk_selection_model_get_selection          = cc.DlFunc[func(model *SelectionModel) *Bitset, *Bitset]{Name: "gtk_selection_model_get_selection"}
	gtk_selection_model_get_selection_in_range = cc.DlFunc[func(model *SelectionModel, position, nItems uint32) *Bitset, *Bitset]{Name: "gtk_selection_model_get_selection_in_range"}
	gtk_selection_model_select_item            = cc.DlFunc[func(model *SelectionModel, position uint32, unselectRest int32) int32, int32]{Name: "gtk_selection_model_select_item"}
	gtk_selection_model_unselect_item          = cc.DlFunc[func(model *SelectionModel, position uint32) int32, int32]{Name: "gtk_selection_model_unselect_item"}
	gtk_selection_model_select_range           = cc.DlFunc[func(model *SelectionModel, position, nItems uint32, unselectRest int32) int32, int32]{Name: "gtk_selection_model_select_range"}
	gtk_selection_model_unselect_range         = cc.DlFunc[func(model *SelectionModel, position, nItems uint32) int32, int32]{Name: "gtk_selection_model_unselect_range"}
	gtk_selection_model_select_all             = cc.DlFunc[func(model *SelectionModel) int32, int32]{Name: "gtk_selection_model_select_all"}
	gtk_selection_model_unselect_all           = cc.DlFunc[func(model *SelectionModel) int32, int32]{Name: "gtk_selection_model_unselect_all"}
	gtk_selection_model_set_selection          = cc.DlFunc[func(model *SelectionModel, selected, mask *Bitset) int32, int32]{Name: "gtk_selection_model_set_selection"}
	gtk_selection_model_selection_changed      = cc.DlFunc[func(model *SelectionModel, position, nItems uint32), cc.Void]{Name: "gtk_selection_model_selection_changed"}
	// #endregion

	// #region Separator
	gtk_separator_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_separator_get_type"}
	gtk_separator_new      = cc.DlFunc[func(orientation Orientation) *Separator, *Separator]{Name: "gtk_separator_new"}
	// #endregion

	// #region Settings
	gtk_settings_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_settings_get_type"}
	gtk_settings_get_default     = cc.DlFunc[func() *Settings, *Settings]{Name: "gtk_settings_get_default"}
	gtk_settings_get_for_display = cc.DlFunc[func(display *gdk.Display) *Settings, *Settings]{Name: "gtk_settings_get_for_display"}
	gtk_settings_reset_property  = cc.DlFunc[func(settings *Settings, name cc.String), cc.Void]{Name: "gtk_settings_reset_property"}
	// #endregion

	// #region Shortcut
	gtk_shortcut_new                = cc.DlFunc[func(*ShortcutTrigger, *ShortcutAction) *Shortcut, *Shortcut]{Name: "gtk_shortcut_new"}
	gtk_shortcut_new_with_arguments = cc.DlFunc[func(*ShortcutTrigger, *ShortcutAction, cc.String, ...interface{}) *Shortcut, *Shortcut]{Name: "gtk_shortcut_new_with_arguments", Va: true}
	gtk_shortcut_get_trigger        = cc.DlFunc[func(*Shortcut) *ShortcutTrigger, *ShortcutTrigger]{Name: "gtk_shortcut_get_trigger"}
	gtk_shortcut_set_trigger        = cc.DlFunc[func(*Shortcut, *ShortcutTrigger), cc.Void]{Name: "gtk_shortcut_set_trigger"}
	gtk_shortcut_get_action         = cc.DlFunc[func(*Shortcut) *ShortcutAction, *ShortcutAction]{Name: "gtk_shortcut_get_action"}
	gtk_shortcut_set_action         = cc.DlFunc[func(*Shortcut, *ShortcutAction), cc.Void]{Name: "gtk_shortcut_set_action"}
	gtk_shortcut_get_arguments      = cc.DlFunc[func(*Shortcut) *glib.GVariant, *glib.GVariant]{Name: "gtk_shortcut_get_arguments"}
	gtk_shortcut_set_arguments      = cc.DlFunc[func(*Shortcut, *glib.GVariant), cc.Void]{Name: "gtk_shortcut_set_arguments"}
	// #endregion

	// #region ShortcutAction
	gtk_shortcut_action_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_shortcut_action_get_type"}
	gtk_nothing_action_get_type       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_nothing_action_get_type"}
	gtk_callback_action_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_callback_action_get_type"}
	gtk_mnemonic_action_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_mnemonic_action_get_type"}
	gtk_activate_action_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_activate_action_get_type"}
	gtk_signal_action_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_signal_action_get_type"}
	gtk_named_action_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_named_action_get_type"}
	gtk_shortcut_action_to_string     = cc.DlFunc[func(*ShortcutAction) cc.String, cc.String]{Name: "gtk_shortcut_action_to_string"}
	gtk_shortcut_action_parse_string  = cc.DlFunc[func(cc.String) *ShortcutAction, *ShortcutAction]{Name: "gtk_shortcut_action_parse_string"}
	gtk_shortcut_action_print         = cc.DlFunc[func(*ShortcutAction, *glib.GString), cc.Void]{Name: "gtk_shortcut_action_print"}
	gtk_shortcut_action_activate      = cc.DlFunc[func(*ShortcutAction, ShortcutActionFlags, *Widget, *glib.GVariant) int32, int32]{Name: "gtk_shortcut_action_activate"}
	gtk_nothing_action_get            = cc.DlFunc[func() *NothingAction, *NothingAction]{Name: "gtk_nothing_action_get"}
	gtk_callback_action_new           = cc.DlFunc[func(callback uptr, data uptr, destroy uptr) *CallbackAction, *CallbackAction]{Name: "gtk_callback_action_new"}
	gtk_mnemonic_action_get           = cc.DlFunc[func() *MnemonicAction, *MnemonicAction]{Name: "gtk_mnemonic_action_get"}
	gtk_activate_action_get           = cc.DlFunc[func() *ActivateAction, *ActivateAction]{Name: "gtk_activate_action_get"}
	gtk_signal_action_new             = cc.DlFunc[func(signalName cc.String) *SignalAction, *SignalAction]{Name: "gtk_signal_action_new"}
	gtk_signal_action_get_signal_name = cc.DlFunc[func(*SignalAction) cc.String, cc.String]{Name: "gtk_signal_action_get_signal_name"}
	gtk_named_action_new              = cc.DlFunc[func(name cc.String) *NamedAction, *NamedAction]{Name: "gtk_named_action_new"}
	gtk_named_action_get_action_name  = cc.DlFunc[func(*NamedAction) cc.String, cc.String]{Name: "gtk_named_action_get_action_name"}
	// #endregion

	// #region ShortcutController
	gtk_shortcut_controller_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_shortcut_controller_get_type"}
	gtk_shortcut_controller_new                     = cc.DlFunc[func() *ShortcutController, *ShortcutController]{Name: "gtk_shortcut_controller_new"}
	gtk_shortcut_controller_new_for_model           = cc.DlFunc[func(model *gio.GListModel) *ShortcutController, *ShortcutController]{Name: "gtk_shortcut_controller_new_for_model"}
	gtk_shortcut_controller_set_mnemonics_modifiers = cc.DlFunc[func(ctrl *ShortcutController, modifiers gdk.ModifierType), cc.Void]{Name: "gtk_shortcut_controller_set_mnemonics_modifiers"}
	gtk_shortcut_controller_get_mnemonics_modifiers = cc.DlFunc[func(ctrl *ShortcutController) gdk.ModifierType, gdk.ModifierType]{Name: "gtk_shortcut_controller_get_mnemonics_modifiers"}
	gtk_shortcut_controller_set_scope               = cc.DlFunc[func(ctrl *ShortcutController, scope ShortcutScope), cc.Void]{Name: "gtk_shortcut_controller_set_scope"}
	gtk_shortcut_controller_get_scope               = cc.DlFunc[func(ctrl *ShortcutController) ShortcutScope, ShortcutScope]{Name: "gtk_shortcut_controller_get_scope"}
	gtk_shortcut_controller_add_shortcut            = cc.DlFunc[func(ctrl *ShortcutController, shortcut *Shortcut), cc.Void]{Name: "gtk_shortcut_controller_add_shortcut"}
	gtk_shortcut_controller_remove_shortcut         = cc.DlFunc[func(ctrl *ShortcutController, shortcut *Shortcut), cc.Void]{Name: "gtk_shortcut_controller_remove_shortcut"}
	// #endregion

	// #region ShortcutManager
	gtk_shortcut_manager_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_shortcut_manager_get_type"}
	// #endregion

	// #region ShortcutTrigger
	gtk_shortcut_trigger_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_shortcut_trigger_get_type"}
	gtk_never_trigger_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_never_trigger_get_type"}
	gtk_keyval_trigger_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_keyval_trigger_get_type"}
	gtk_mnemonic_trigger_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_mnemonic_trigger_get_type"}
	gtk_alternative_trigger_get_type   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_alternative_trigger_get_type"}
	gtk_shortcut_trigger_parse_string  = cc.DlFunc[func(cc.String) *ShortcutTrigger, *ShortcutTrigger]{Name: "gtk_shortcut_trigger_parse_string"}
	gtk_shortcut_trigger_to_string     = cc.DlFunc[func(*ShortcutTrigger) cc.String, cc.String]{Name: "gtk_shortcut_trigger_to_string"}
	gtk_shortcut_trigger_print         = cc.DlFunc[func(*ShortcutTrigger, *glib.GString), cc.Void]{Name: "gtk_shortcut_trigger_print"}
	gtk_shortcut_trigger_to_label      = cc.DlFunc[func(*ShortcutTrigger, *gdk.Display) cc.String, cc.String]{Name: "gtk_shortcut_trigger_to_label"}
	gtk_shortcut_trigger_print_label   = cc.DlFunc[func(*ShortcutTrigger, *gdk.Display, *glib.GString) int32, int32]{Name: "gtk_shortcut_trigger_print_label"}
	gtk_shortcut_trigger_hash          = cc.DlFunc[func(*ShortcutTrigger) uint, uint]{Name: "gtk_shortcut_trigger_hash"}
	gtk_shortcut_trigger_equal         = cc.DlFunc[func(*ShortcutTrigger, *ShortcutTrigger) int32, int32]{Name: "gtk_shortcut_trigger_equal"}
	gtk_shortcut_trigger_compare       = cc.DlFunc[func(*ShortcutTrigger, *ShortcutTrigger) int, int]{Name: "gtk_shortcut_trigger_compare"}
	gtk_shortcut_trigger_trigger       = cc.DlFunc[func(*ShortcutTrigger, *gdk.Event, int32) gdk.KeyMatch, gdk.KeyMatch]{Name: "gtk_shortcut_trigger_trigger"}
	gtk_never_trigger_get              = cc.DlFunc[func() *NeverTrigger, *NeverTrigger]{Name: "gtk_never_trigger_get"}
	gtk_keyval_trigger_new             = cc.DlFunc[func(uint, gdk.ModifierType) *KeyvalTrigger, *KeyvalTrigger]{Name: "gtk_keyval_trigger_new"}
	gtk_keyval_trigger_get_modifiers   = cc.DlFunc[func(*KeyvalTrigger) gdk.ModifierType, gdk.ModifierType]{Name: "gtk_keyval_trigger_get_modifiers"}
	gtk_keyval_trigger_get_keyval      = cc.DlFunc[func(*KeyvalTrigger) uint, uint]{Name: "gtk_keyval_trigger_get_keyval"}
	gtk_mnemonic_trigger_new           = cc.DlFunc[func(uint) *MnemonicTrigger, *MnemonicTrigger]{Name: "gtk_mnemonic_trigger_new"}
	gtk_mnemonic_trigger_get_keyval    = cc.DlFunc[func(*MnemonicTrigger) uint, uint]{Name: "gtk_mnemonic_trigger_get_keyval"}
	gtk_alternative_trigger_new        = cc.DlFunc[func(*ShortcutTrigger, *ShortcutTrigger) *AlternativeTrigger, *AlternativeTrigger]{Name: "gtk_alternative_trigger_new"}
	gtk_alternative_trigger_get_first  = cc.DlFunc[func(*AlternativeTrigger) *ShortcutTrigger, *ShortcutTrigger]{Name: "gtk_alternative_trigger_get_first"}
	gtk_alternative_trigger_get_second = cc.DlFunc[func(*AlternativeTrigger) *ShortcutTrigger, *ShortcutTrigger]{Name: "gtk_alternative_trigger_get_second"}

	// #endregion

	// #region SignalListItemFactory
	gtk_signal_list_item_factory_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_signal_list_item_factory_get_type"}
	gtk_signal_list_item_factory_new      = cc.DlFunc[func() *SignalListItemFactory, *SignalListItemFactory]{Name: "gtk_signal_list_item_factory_new"}
	// #endregion

	// #region SingleSelection
	gtk_single_selection_new               = cc.DlFunc[func(model *gio.GListModel) *SingleSelection, *SingleSelection]{Name: "gtk_single_selection_new"}
	gtk_single_selection_get_model         = cc.DlFunc[func(sel *SingleSelection) *gio.GListModel, *gio.GListModel]{Name: "gtk_single_selection_get_model"}
	gtk_single_selection_set_model         = cc.DlFunc[func(sel *SingleSelection, model *gio.GListModel), cc.Void]{Name: "gtk_single_selection_set_model"}
	gtk_single_selection_get_selected      = cc.DlFunc[func(sel *SingleSelection) uint32, uint32]{Name: "gtk_single_selection_get_selected"}
	gtk_single_selection_set_selected      = cc.DlFunc[func(sel *SingleSelection, position uint32), cc.Void]{Name: "gtk_single_selection_set_selected"}
	gtk_single_selection_get_selected_item = cc.DlFunc[func(sel *SingleSelection) uptr, uptr]{Name: "gtk_single_selection_get_selected_item"}
	gtk_single_selection_get_autoselect    = cc.DlFunc[func(sel *SingleSelection) int32, int32]{Name: "gtk_single_selection_get_autoselect"}
	gtk_single_selection_set_autoselect    = cc.DlFunc[func(sel *SingleSelection, autoselect int32), cc.Void]{Name: "gtk_single_selection_set_autoselect"}
	gtk_single_selection_get_can_unselect  = cc.DlFunc[func(sel *SingleSelection) int32, int32]{Name: "gtk_single_selection_get_can_unselect"}
	gtk_single_selection_set_can_unselect  = cc.DlFunc[func(sel *SingleSelection, canUnselect int32), cc.Void]{Name: "gtk_single_selection_set_can_unselect"}
	// #endregion

	// #region SizeGroup
	gtk_size_group_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_size_group_get_type"}
	gtk_size_group_new           = cc.DlFunc[func(mode SizeGroupMode) *SizeGroup, *SizeGroup]{Name: "gtk_size_group_new"}
	gtk_size_group_set_mode      = cc.DlFunc[func(sg *SizeGroup, mode SizeGroupMode), cc.Void]{Name: "gtk_size_group_set_mode"}
	gtk_size_group_get_mode      = cc.DlFunc[func(sg *SizeGroup) SizeGroupMode, SizeGroupMode]{Name: "gtk_size_group_get_mode"}
	gtk_size_group_add_widget    = cc.DlFunc[func(sg *SizeGroup, widget *Widget), cc.Void]{Name: "gtk_size_group_add_widget"}
	gtk_size_group_remove_widget = cc.DlFunc[func(sg *SizeGroup, widget *Widget), cc.Void]{Name: "gtk_size_group_remove_widget"}
	gtk_size_group_get_widgets   = cc.DlFunc[func(sg *SizeGroup) uptr, uptr]{Name: "gtk_size_group_get_widgets"}
	// #endregion

	// #region RequestedSize
	gtk_distribute_natural_allocation = cc.DlFunc[func(extraSpace int32, nRequestedSizes uint32, sizes *RequestedSize) int32, int32]{Name: "gtk_distribute_natural_allocation"}
	// #endregion

	// #region SliceListModel
	gtk_slice_list_model_get_type   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_slice_list_model_get_type"}
	gtk_slice_list_model_new        = cc.DlFunc[func(model *gio.GListModel, offset, size uint32) *SliceListModel, *SliceListModel]{Name: "gtk_slice_list_model_new"}
	gtk_slice_list_model_set_model  = cc.DlFunc[func(self *SliceListModel, model *gio.GListModel), cc.Void]{Name: "gtk_slice_list_model_set_model"}
	gtk_slice_list_model_get_model  = cc.DlFunc[func(self *SliceListModel) *gio.GListModel, *gio.GListModel]{Name: "gtk_slice_list_model_get_model"}
	gtk_slice_list_model_set_offset = cc.DlFunc[func(self *SliceListModel, offset uint32), cc.Void]{Name: "gtk_slice_list_model_set_offset"}
	gtk_slice_list_model_get_offset = cc.DlFunc[func(self *SliceListModel) uint32, uint32]{Name: "gtk_slice_list_model_get_offset"}
	gtk_slice_list_model_set_size   = cc.DlFunc[func(self *SliceListModel, size uint32), cc.Void]{Name: "gtk_slice_list_model_set_size"}
	gtk_slice_list_model_get_size   = cc.DlFunc[func(self *SliceListModel) uint32, uint32]{Name: "gtk_slice_list_model_get_size"}
	// #endregion

	// #region Snapshot
	gtk_snapshot_get_type                         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_snapshot_get_type"}
	gtk_snapshot_new                              = cc.DlFunc[func() *Snapshot, *Snapshot]{Name: "gtk_snapshot_new"}
	gtk_snapshot_free_to_node                     = cc.DlFunc[func(snapshot *Snapshot) *gsk.RenderNode, *gsk.RenderNode]{Name: "gtk_snapshot_free_to_node"}
	gtk_snapshot_free_to_paintable                = cc.DlFunc[func(snapshot *Snapshot, size *graphene.Size) *gdk.Paintable, *gdk.Paintable]{Name: "gtk_snapshot_free_to_paintable"}
	gtk_snapshot_to_node                          = cc.DlFunc[func(snapshot *Snapshot) *gsk.RenderNode, *gsk.RenderNode]{Name: "gtk_snapshot_to_node"}
	gtk_snapshot_to_paintable                     = cc.DlFunc[func(snapshot *Snapshot, size *graphene.Size) *gdk.Paintable, *gdk.Paintable]{Name: "gtk_snapshot_to_paintable"}
	gtk_snapshot_push_debug                       = cc.DlFunc[func(snapshot *Snapshot, message cc.String, args ...interface{}), cc.Void]{Name: "gtk_snapshot_push_debug", Va: true}
	gtk_snapshot_push_opacity                     = cc.DlFunc[func(snapshot *Snapshot, opacity float64), cc.Void]{Name: "gtk_snapshot_push_opacity"}
	gtk_snapshot_push_blur                        = cc.DlFunc[func(snapshot *Snapshot, radius float64), cc.Void]{Name: "gtk_snapshot_push_blur"}
	gtk_snapshot_push_color_matrix                = cc.DlFunc[func(snapshot *Snapshot, colorMatrix *graphene.Matrix, colorOffset *graphene.Vec4), cc.Void]{Name: "gtk_snapshot_push_color_matrix"}
	gtk_snapshot_push_repeat                      = cc.DlFunc[func(snapshot *Snapshot, bounds *graphene.Rect, childBounds *graphene.Rect), cc.Void]{Name: "gtk_snapshot_push_repeat"}
	gtk_snapshot_push_clip                        = cc.DlFunc[func(snapshot *Snapshot, bounds *graphene.Rect), cc.Void]{Name: "gtk_snapshot_push_clip"}
	gtk_snapshot_push_rounded_clip                = cc.DlFunc[func(snapshot *Snapshot, bounds *gsk.RoundedRect), cc.Void]{Name: "gtk_snapshot_push_rounded_clip"}
	gtk_snapshot_push_fill                        = cc.DlFunc[func(snapshot *Snapshot, path *gsk.Path, fillRule gsk.FillRule), cc.Void]{Name: "gtk_snapshot_push_fill"}
	gtk_snapshot_push_stroke                      = cc.DlFunc[func(snapshot *Snapshot, path *gsk.Path, stroke *gsk.Stroke), cc.Void]{Name: "gtk_snapshot_push_stroke"}
	gtk_snapshot_push_shadow                      = cc.DlFunc[func(snapshot *Snapshot, shadow *gsk.Shadow, nShadows uint64), cc.Void]{Name: "gtk_snapshot_push_shadow"}
	gtk_snapshot_push_blend                       = cc.DlFunc[func(snapshot *Snapshot, blendMode gsk.BlendMode), cc.Void]{Name: "gtk_snapshot_push_blend"}
	gtk_snapshot_push_mask                        = cc.DlFunc[func(snapshot *Snapshot, maskMode gsk.MaskMode), cc.Void]{Name: "gtk_snapshot_push_mask"}
	gtk_snapshot_push_cross_fade                  = cc.DlFunc[func(snapshot *Snapshot, progress float64), cc.Void]{Name: "gtk_snapshot_push_cross_fade"}
	gtk_snapshot_pop                              = cc.DlFunc[func(snapshot *Snapshot), cc.Void]{Name: "gtk_snapshot_pop"}
	gtk_snapshot_save                             = cc.DlFunc[func(snapshot *Snapshot), cc.Void]{Name: "gtk_snapshot_save"}
	gtk_snapshot_restore                          = cc.DlFunc[func(snapshot *Snapshot), cc.Void]{Name: "gtk_snapshot_restore"}
	gtk_snapshot_transform                        = cc.DlFunc[func(snapshot *Snapshot, transform *gsk.Transform), cc.Void]{Name: "gtk_snapshot_transform"}
	gtk_snapshot_transform_matrix                 = cc.DlFunc[func(snapshot *Snapshot, matrix *graphene.Matrix), cc.Void]{Name: "gtk_snapshot_transform_matrix"}
	gtk_snapshot_translate                        = cc.DlFunc[func(snapshot *Snapshot, point *graphene.Point), cc.Void]{Name: "gtk_snapshot_translate"}
	gtk_snapshot_translate_3d                     = cc.DlFunc[func(snapshot *Snapshot, point *graphene.Point3D), cc.Void]{Name: "gtk_snapshot_translate_3d"}
	gtk_snapshot_rotate                           = cc.DlFunc[func(snapshot *Snapshot, angle float32), cc.Void]{Name: "gtk_snapshot_rotate"}
	gtk_snapshot_rotate_3d                        = cc.DlFunc[func(snapshot *Snapshot, angle float32, axis *graphene.Vec3), cc.Void]{Name: "gtk_snapshot_rotate_3d"}
	gtk_snapshot_scale                            = cc.DlFunc[func(snapshot *Snapshot, factorX, factorY float32), cc.Void]{Name: "gtk_snapshot_scale"}
	gtk_snapshot_scale_3d                         = cc.DlFunc[func(snapshot *Snapshot, factorX, factorY, factorZ float32), cc.Void]{Name: "gtk_snapshot_scale_3d"}
	gtk_snapshot_perspective                      = cc.DlFunc[func(snapshot *Snapshot, depth float32), cc.Void]{Name: "gtk_snapshot_perspective"}
	gtk_snapshot_append_node                      = cc.DlFunc[func(snapshot *Snapshot, node *gsk.RenderNode), cc.Void]{Name: "gtk_snapshot_append_node"}
	gtk_snapshot_append_cairo                     = cc.DlFunc[func(snapshot *Snapshot, bounds *graphene.Rect) *cairo.Context, *cairo.Context]{Name: "gtk_snapshot_append_cairo"}
	gtk_snapshot_append_texture                   = cc.DlFunc[func(snapshot *Snapshot, texture *gdk.Texture, bounds *graphene.Rect), cc.Void]{Name: "gtk_snapshot_append_texture"}
	gtk_snapshot_append_scaled_texture            = cc.DlFunc[func(snapshot *Snapshot, texture *gdk.Texture, filter gsk.ScalingFilter, bounds *graphene.Rect), cc.Void]{Name: "gtk_snapshot_append_scaled_texture"}
	gtk_snapshot_append_color                     = cc.DlFunc[func(snapshot *Snapshot, color *gdk.RGBA, bounds *graphene.Rect), cc.Void]{Name: "gtk_snapshot_append_color"}
	gtk_snapshot_append_linear_gradient           = cc.DlFunc[func(snapshot *Snapshot, bounds *graphene.Rect, startPoint, endPoint *graphene.Point, stops *gsk.ColorStop, nStops uint64), cc.Void]{Name: "gtk_snapshot_append_linear_gradient"}
	gtk_snapshot_append_repeating_linear_gradient = cc.DlFunc[func(snapshot *Snapshot, bounds *graphene.Rect, startPoint, endPoint *graphene.Point, stops *gsk.ColorStop, nStops uint64), cc.Void]{Name: "gtk_snapshot_append_repeating_linear_gradient"}
	gtk_snapshot_append_radial_gradient           = cc.DlFunc[func(snapshot *Snapshot, bounds *graphene.Rect, center *graphene.Point, hradius, vradius, start, end float32, stops *gsk.ColorStop, nStops uint64), cc.Void]{Name: "gtk_snapshot_append_radial_gradient"}
	gtk_snapshot_append_repeating_radial_gradient = cc.DlFunc[func(snapshot *Snapshot, bounds *graphene.Rect, center *graphene.Point, hradius, vradius, start, end float32, stops *gsk.ColorStop, nStops uint64), cc.Void]{Name: "gtk_snapshot_append_repeating_radial_gradient"}
	gtk_snapshot_append_conic_gradient            = cc.DlFunc[func(snapshot *Snapshot, bounds *graphene.Rect, center *graphene.Point, rotation float32, stops *gsk.ColorStop, nStops uint64), cc.Void]{Name: "gtk_snapshot_append_conic_gradient"}
	gtk_snapshot_append_border                    = cc.DlFunc[func(snapshot *Snapshot, outline *gsk.RoundedRect, borderWidth *float32, borderColor *gdk.RGBA), cc.Void]{Name: "gtk_snapshot_append_border"}
	gtk_snapshot_append_inset_shadow              = cc.DlFunc[func(snapshot *Snapshot, outline *gsk.RoundedRect, color *gdk.RGBA, dx, dy, spread, blurRadius float32), cc.Void]{Name: "gtk_snapshot_append_inset_shadow"}
	gtk_snapshot_append_outset_shadow             = cc.DlFunc[func(snapshot *Snapshot, outline *gsk.RoundedRect, color *gdk.RGBA, dx, dy, spread, blurRadius float32), cc.Void]{Name: "gtk_snapshot_append_outset_shadow"}
	gtk_snapshot_append_layout                    = cc.DlFunc[func(snapshot *Snapshot, layout *pango.Layout, color *gdk.RGBA), cc.Void]{Name: "gtk_snapshot_append_layout"}
	gtk_snapshot_append_fill                      = cc.DlFunc[func(snapshot *Snapshot, path *gsk.Path, fillRule gsk.FillRule, color *gdk.RGBA), cc.Void]{Name: "gtk_snapshot_append_fill"}
	gtk_snapshot_append_stroke                    = cc.DlFunc[func(snapshot *Snapshot, path *gsk.Path, stroke *gsk.Stroke, color *gdk.RGBA), cc.Void]{Name: "gtk_snapshot_append_stroke"}
	// #endregion

	// #region Sorter
	gtk_sorter_get_type  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_sorter_get_type"}
	gtk_sorter_compare   = cc.DlFunc[func(self *Sorter, item1, item2 uptr) Ordering, Ordering]{Name: "gtk_sorter_compare"}
	gtk_sorter_get_order = cc.DlFunc[func(self *Sorter) SorterOrder, SorterOrder]{Name: "gtk_sorter_get_order"}
	gtk_sorter_changed   = cc.DlFunc[func(self *Sorter, change SorterChange), cc.Void]{Name: "gtk_sorter_changed"}
	// #endregion

	// #region SortListModel
	gtk_sort_list_model_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_sort_list_model_get_type"}
	gtk_sort_list_model_new                = cc.DlFunc[func(model *gio.GListModel, sorter *Sorter) *SortListModel, *SortListModel]{Name: "gtk_sort_list_model_new"}
	gtk_sort_list_model_set_sorter         = cc.DlFunc[func(m *SortListModel, sorter *Sorter), cc.Void]{Name: "gtk_sort_list_model_set_sorter"}
	gtk_sort_list_model_get_sorter         = cc.DlFunc[func(m *SortListModel) *Sorter, *Sorter]{Name: "gtk_sort_list_model_get_sorter"}
	gtk_sort_list_model_set_section_sorter = cc.DlFunc[func(m *SortListModel, sorter *Sorter), cc.Void]{Name: "gtk_sort_list_model_set_section_sorter"}
	gtk_sort_list_model_get_section_sorter = cc.DlFunc[func(m *SortListModel) *Sorter, *Sorter]{Name: "gtk_sort_list_model_get_section_sorter"}
	gtk_sort_list_model_set_model          = cc.DlFunc[func(m *SortListModel, model *gio.GListModel), cc.Void]{Name: "gtk_sort_list_model_set_model"}
	gtk_sort_list_model_get_model          = cc.DlFunc[func(m *SortListModel) *gio.GListModel, *gio.GListModel]{Name: "gtk_sort_list_model_get_model"}
	gtk_sort_list_model_set_incremental    = cc.DlFunc[func(m *SortListModel, incremental int32), cc.Void]{Name: "gtk_sort_list_model_set_incremental"}
	gtk_sort_list_model_get_incremental    = cc.DlFunc[func(m *SortListModel) int32, int32]{Name: "gtk_sort_list_model_get_incremental"}
	gtk_sort_list_model_get_pending        = cc.DlFunc[func(m *SortListModel) uint32, uint32]{Name: "gtk_sort_list_model_get_pending"}
	// #endregion

	// #region SpinButton
	gtk_spin_button_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_spin_button_get_type"}
	gtk_spin_button_configure             = cc.DlFunc[func(spin *SpinButton, adjustment *Adjustment, climbRate float64, digits uint32), cc.Void]{Name: "gtk_spin_button_configure"}
	gtk_spin_button_new                   = cc.DlFunc[func(adjustment *Adjustment, climbRate float64, digits uint32) *SpinButton, *SpinButton]{Name: "gtk_spin_button_new"}
	gtk_spin_button_new_with_range        = cc.DlFunc[func(min, max, step float64) *SpinButton, *SpinButton]{Name: "gtk_spin_button_new_with_range"}
	gtk_spin_button_set_activates_default = cc.DlFunc[func(spin *SpinButton, activatesDefault int32), cc.Void]{Name: "gtk_spin_button_set_activates_default"}
	gtk_spin_button_get_activates_default = cc.DlFunc[func(spin *SpinButton) int32, int32]{Name: "gtk_spin_button_get_activates_default"}
	gtk_spin_button_set_adjustment        = cc.DlFunc[func(spin *SpinButton, adjustment *Adjustment), cc.Void]{Name: "gtk_spin_button_set_adjustment"}
	gtk_spin_button_get_adjustment        = cc.DlFunc[func(spin *SpinButton) *Adjustment, *Adjustment]{Name: "gtk_spin_button_get_adjustment"}
	gtk_spin_button_set_digits            = cc.DlFunc[func(spin *SpinButton, digits uint32), cc.Void]{Name: "gtk_spin_button_set_digits"}
	gtk_spin_button_get_digits            = cc.DlFunc[func(spin *SpinButton) uint32, uint32]{Name: "gtk_spin_button_get_digits"}
	gtk_spin_button_set_increments        = cc.DlFunc[func(spin *SpinButton, step, page float64), cc.Void]{Name: "gtk_spin_button_set_increments"}
	gtk_spin_button_get_increments        = cc.DlFunc[func(spin *SpinButton, step, page *float64), cc.Void]{Name: "gtk_spin_button_get_increments"}
	gtk_spin_button_set_range             = cc.DlFunc[func(spin *SpinButton, min, max float64), cc.Void]{Name: "gtk_spin_button_set_range"}
	gtk_spin_button_get_range             = cc.DlFunc[func(spin *SpinButton, min, max *float64), cc.Void]{Name: "gtk_spin_button_get_range"}
	gtk_spin_button_get_value             = cc.DlFunc[func(spin *SpinButton) float64, float64]{Name: "gtk_spin_button_get_value"}
	gtk_spin_button_get_value_as_int      = cc.DlFunc[func(spin *SpinButton) int32, int32]{Name: "gtk_spin_button_get_value_as_int"}
	gtk_spin_button_set_value             = cc.DlFunc[func(spin *SpinButton, value float64), cc.Void]{Name: "gtk_spin_button_set_value"}
	gtk_spin_button_set_update_policy     = cc.DlFunc[func(spin *SpinButton, policy SpinButtonUpdatePolicy), cc.Void]{Name: "gtk_spin_button_set_update_policy"}
	gtk_spin_button_get_update_policy     = cc.DlFunc[func(spin *SpinButton) SpinButtonUpdatePolicy, SpinButtonUpdatePolicy]{Name: "gtk_spin_button_get_update_policy"}
	gtk_spin_button_set_numeric           = cc.DlFunc[func(spin *SpinButton, numeric int32), cc.Void]{Name: "gtk_spin_button_set_numeric"}
	gtk_spin_button_get_numeric           = cc.DlFunc[func(spin *SpinButton) int32, int32]{Name: "gtk_spin_button_get_numeric"}
	gtk_spin_button_spin                  = cc.DlFunc[func(spin *SpinButton, direction SpinType, increment float64), cc.Void]{Name: "gtk_spin_button_spin"}
	gtk_spin_button_set_wrap              = cc.DlFunc[func(spin *SpinButton, wrap int32), cc.Void]{Name: "gtk_spin_button_set_wrap"}
	gtk_spin_button_get_wrap              = cc.DlFunc[func(spin *SpinButton) int32, int32]{Name: "gtk_spin_button_get_wrap"}
	gtk_spin_button_set_snap_to_ticks     = cc.DlFunc[func(spin *SpinButton, snapToTicks int32), cc.Void]{Name: "gtk_spin_button_set_snap_to_ticks"}
	gtk_spin_button_get_snap_to_ticks     = cc.DlFunc[func(spin *SpinButton) int32, int32]{Name: "gtk_spin_button_get_snap_to_ticks"}
	gtk_spin_button_set_climb_rate        = cc.DlFunc[func(spin *SpinButton, climbRate float64), cc.Void]{Name: "gtk_spin_button_set_climb_rate"}
	gtk_spin_button_get_climb_rate        = cc.DlFunc[func(spin *SpinButton) float64, float64]{Name: "gtk_spin_button_get_climb_rate"}
	gtk_spin_button_update                = cc.DlFunc[func(spin *SpinButton), cc.Void]{Name: "gtk_spin_button_update"}
	// #endregion

	// #region Spinner
	gtk_spinner_get_type     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_spinner_get_type"}
	gtk_spinner_new          = cc.DlFunc[func() *Spinner, *Spinner]{Name: "gtk_spinner_new"}
	gtk_spinner_start        = cc.DlFunc[func(spinner *Spinner), cc.Void]{Name: "gtk_spinner_start"}
	gtk_spinner_stop         = cc.DlFunc[func(spinner *Spinner), cc.Void]{Name: "gtk_spinner_stop"}
	gtk_spinner_set_spinning = cc.DlFunc[func(spinner *Spinner, spinning int32), cc.Void]{Name: "gtk_spinner_set_spinning"}
	gtk_spinner_get_spinning = cc.DlFunc[func(spinner *Spinner) int32, int32]{Name: "gtk_spinner_get_spinning"}
	// #endregion

	// #region Stack
	gtk_stack_page_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_stack_page_get_type"}
	gtk_stack_page_get_child           = cc.DlFunc[func(page *StackPage) *Widget, *Widget]{Name: "gtk_stack_page_get_child"}
	gtk_stack_page_get_visible         = cc.DlFunc[func(page *StackPage) int32, int32]{Name: "gtk_stack_page_get_visible"}
	gtk_stack_page_set_visible         = cc.DlFunc[func(page *StackPage, visible int32), cc.Void]{Name: "gtk_stack_page_set_visible"}
	gtk_stack_page_get_needs_attention = cc.DlFunc[func(page *StackPage) int32, int32]{Name: "gtk_stack_page_get_needs_attention"}
	gtk_stack_page_set_needs_attention = cc.DlFunc[func(page *StackPage, setting int32), cc.Void]{Name: "gtk_stack_page_set_needs_attention"}
	gtk_stack_page_get_use_underline   = cc.DlFunc[func(page *StackPage) int32, int32]{Name: "gtk_stack_page_get_use_underline"}
	gtk_stack_page_set_use_underline   = cc.DlFunc[func(page *StackPage, setting int32), cc.Void]{Name: "gtk_stack_page_set_use_underline"}
	gtk_stack_page_get_name            = cc.DlFunc[func(page *StackPage) cc.String, cc.String]{Name: "gtk_stack_page_get_name"}
	gtk_stack_page_set_name            = cc.DlFunc[func(page *StackPage, setting cc.String), cc.Void]{Name: "gtk_stack_page_set_name"}
	gtk_stack_page_get_title           = cc.DlFunc[func(page *StackPage) cc.String, cc.String]{Name: "gtk_stack_page_get_title"}
	gtk_stack_page_set_title           = cc.DlFunc[func(page *StackPage, setting cc.String), cc.Void]{Name: "gtk_stack_page_set_title"}
	gtk_stack_page_get_icon_name       = cc.DlFunc[func(page *StackPage) cc.String, cc.String]{Name: "gtk_stack_page_get_icon_name"}
	gtk_stack_page_set_icon_name       = cc.DlFunc[func(page *StackPage, setting cc.String), cc.Void]{Name: "gtk_stack_page_set_icon_name"}
	gtk_stack_get_type                 = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_stack_get_type"}
	gtk_stack_new                      = cc.DlFunc[func() *Stack, *Stack]{Name: "gtk_stack_new"}
	gtk_stack_add_child                = cc.DlFunc[func(stack *Stack, child *Widget) *StackPage, *StackPage]{Name: "gtk_stack_add_child"}
	gtk_stack_add_named                = cc.DlFunc[func(stack *Stack, child *Widget, name cc.String) *StackPage, *StackPage]{Name: "gtk_stack_add_named"}
	gtk_stack_add_titled               = cc.DlFunc[func(stack *Stack, child *Widget, name, title cc.String) *StackPage, *StackPage]{Name: "gtk_stack_add_titled"}
	gtk_stack_remove                   = cc.DlFunc[func(stack *Stack, child *Widget), cc.Void]{Name: "gtk_stack_remove"}
	gtk_stack_get_page                 = cc.DlFunc[func(stack *Stack, child *Widget) *StackPage, *StackPage]{Name: "gtk_stack_get_page"}
	gtk_stack_get_child_by_name        = cc.DlFunc[func(stack *Stack, name cc.String) *Widget, *Widget]{Name: "gtk_stack_get_child_by_name"}
	gtk_stack_set_visible_child        = cc.DlFunc[func(stack *Stack, child *Widget), cc.Void]{Name: "gtk_stack_set_visible_child"}
	gtk_stack_get_visible_child        = cc.DlFunc[func(stack *Stack) *Widget, *Widget]{Name: "gtk_stack_get_visible_child"}
	gtk_stack_set_visible_child_name   = cc.DlFunc[func(stack *Stack, name cc.String), cc.Void]{Name: "gtk_stack_set_visible_child_name"}
	gtk_stack_get_visible_child_name   = cc.DlFunc[func(stack *Stack) cc.String, cc.String]{Name: "gtk_stack_get_visible_child_name"}
	gtk_stack_set_visible_child_full   = cc.DlFunc[func(stack *Stack, name cc.String, transition StackTransitionType), cc.Void]{Name: "gtk_stack_set_visible_child_full"}
	gtk_stack_set_hhomogeneous         = cc.DlFunc[func(stack *Stack, hhomogeneous int32), cc.Void]{Name: "gtk_stack_set_hhomogeneous"}
	gtk_stack_get_hhomogeneous         = cc.DlFunc[func(stack *Stack) int32, int32]{Name: "gtk_stack_get_hhomogeneous"}
	gtk_stack_set_vhomogeneous         = cc.DlFunc[func(stack *Stack, vhomogeneous int32), cc.Void]{Name: "gtk_stack_set_vhomogeneous"}
	gtk_stack_get_vhomogeneous         = cc.DlFunc[func(stack *Stack) int32, int32]{Name: "gtk_stack_get_vhomogeneous"}
	gtk_stack_set_transition_duration  = cc.DlFunc[func(stack *Stack, duration uint32), cc.Void]{Name: "gtk_stack_set_transition_duration"}
	gtk_stack_get_transition_duration  = cc.DlFunc[func(stack *Stack) uint32, uint32]{Name: "gtk_stack_get_transition_duration"}
	gtk_stack_set_transition_type      = cc.DlFunc[func(stack *Stack, transition StackTransitionType), cc.Void]{Name: "gtk_stack_set_transition_type"}
	gtk_stack_get_transition_type      = cc.DlFunc[func(stack *Stack) StackTransitionType, StackTransitionType]{Name: "gtk_stack_get_transition_type"}
	gtk_stack_get_transition_running   = cc.DlFunc[func(stack *Stack) int32, int32]{Name: "gtk_stack_get_transition_running"}
	gtk_stack_set_interpolate_size     = cc.DlFunc[func(stack *Stack, interpolateSize int32), cc.Void]{Name: "gtk_stack_set_interpolate_size"}
	gtk_stack_get_interpolate_size     = cc.DlFunc[func(stack *Stack) int32, int32]{Name: "gtk_stack_get_interpolate_size"}
	gtk_stack_get_pages                = cc.DlFunc[func(stack *Stack) *SelectionModel, *SelectionModel]{Name: "gtk_stack_get_pages"}
	// #endregion

	// #region StackSidebar
	gtk_stack_sidebar_get_type  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_stack_sidebar_get_type"}
	gtk_stack_sidebar_new       = cc.DlFunc[func() *StackSidebar, *StackSidebar]{Name: "gtk_stack_sidebar_new"}
	gtk_stack_sidebar_set_stack = cc.DlFunc[func(sidebar *StackSidebar, stack *Stack), cc.Void]{Name: "gtk_stack_sidebar_set_stack"}
	gtk_stack_sidebar_get_stack = cc.DlFunc[func(sidebar *StackSidebar) *Stack, *Stack]{Name: "gtk_stack_sidebar_get_stack"}
	// #endregion

	// #region StackSwitcher
	gtk_stack_switcher_get_type  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_stack_switcher_get_type"}
	gtk_stack_switcher_new       = cc.DlFunc[func() *StackSwitcher, *StackSwitcher]{Name: "gtk_stack_switcher_new"}
	gtk_stack_switcher_set_stack = cc.DlFunc[func(switcher *StackSwitcher, stack *Stack), cc.Void]{Name: "gtk_stack_switcher_set_stack"}
	gtk_stack_switcher_get_stack = cc.DlFunc[func(switcher *StackSwitcher) *Stack, *Stack]{Name: "gtk_stack_switcher_get_stack"}
	// #endregion

	// #region StringFilter
	gtk_string_filter_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_string_filter_get_type"}
	gtk_string_filter_new             = cc.DlFunc[func(expression *Expression) *StringFilter, *StringFilter]{Name: "gtk_string_filter_new"}
	gtk_string_filter_get_search      = cc.DlFunc[func(f *StringFilter) cc.String, cc.String]{Name: "gtk_string_filter_get_search"}
	gtk_string_filter_set_search      = cc.DlFunc[func(f *StringFilter, search cc.String), cc.Void]{Name: "gtk_string_filter_set_search"}
	gtk_string_filter_get_expression  = cc.DlFunc[func(f *StringFilter) *Expression, *Expression]{Name: "gtk_string_filter_get_expression"}
	gtk_string_filter_set_expression  = cc.DlFunc[func(f *StringFilter, expression *Expression), cc.Void]{Name: "gtk_string_filter_set_expression"}
	gtk_string_filter_get_ignore_case = cc.DlFunc[func(f *StringFilter) int32, int32]{Name: "gtk_string_filter_get_ignore_case"}
	gtk_string_filter_set_ignore_case = cc.DlFunc[func(f *StringFilter, ignoreCase int32), cc.Void]{Name: "gtk_string_filter_set_ignore_case"}
	gtk_string_filter_get_match_mode  = cc.DlFunc[func(f *StringFilter) StringFilterMatchMode, StringFilterMatchMode]{Name: "gtk_string_filter_get_match_mode"}
	gtk_string_filter_set_match_mode  = cc.DlFunc[func(f *StringFilter, mode StringFilterMatchMode), cc.Void]{Name: "gtk_string_filter_set_match_mode"}
	// #endregion

	// #region StringList
	gtk_string_object_get_type   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_string_object_get_type"}
	gtk_string_object_new        = cc.DlFunc[func(str cc.String) *StringObject, *StringObject]{Name: "gtk_string_object_new"}
	gtk_string_object_get_string = cc.DlFunc[func(self *StringObject) cc.String, cc.String]{Name: "gtk_string_object_get_string"}
	gtk_string_list_get_type     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_string_list_get_type"}
	gtk_string_list_new          = cc.DlFunc[func(strings cc.Strings) *StringList, *StringList]{Name: "gtk_string_list_new"}
	gtk_string_list_append       = cc.DlFunc[func(self *StringList, str cc.String), cc.Void]{Name: "gtk_string_list_append"}
	gtk_string_list_take         = cc.DlFunc[func(self *StringList, str cc.String), cc.Void]{Name: "gtk_string_list_take"}
	gtk_string_list_remove       = cc.DlFunc[func(self *StringList, position uint32), cc.Void]{Name: "gtk_string_list_remove"}
	gtk_string_list_splice       = cc.DlFunc[func(self *StringList, position, nRemovals uint32, additions cc.Strings), cc.Void]{Name: "gtk_string_list_splice"}
	gtk_string_list_get_string   = cc.DlFunc[func(self *StringList, position uint32) cc.String, cc.String]{Name: "gtk_string_list_get_string"}
	gtk_string_list_find         = cc.DlFunc[func(self *StringList, str cc.String) uint32, uint32]{Name: "gtk_string_list_find"}
	// #endregion

	// #region StringSorter
	gtk_string_sorter_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_string_sorter_get_type"}
	gtk_string_sorter_new             = cc.DlFunc[func(expression *Expression) *StringSorter, *StringSorter]{Name: "gtk_string_sorter_new"}
	gtk_string_sorter_get_expression  = cc.DlFunc[func(s *StringSorter) *Expression, *Expression]{Name: "gtk_string_sorter_get_expression"}
	gtk_string_sorter_set_expression  = cc.DlFunc[func(s *StringSorter, expression *Expression), cc.Void]{Name: "gtk_string_sorter_set_expression"}
	gtk_string_sorter_get_ignore_case = cc.DlFunc[func(s *StringSorter) int32, int32]{Name: "gtk_string_sorter_get_ignore_case"}
	gtk_string_sorter_set_ignore_case = cc.DlFunc[func(s *StringSorter, ignoreCase int32), cc.Void]{Name: "gtk_string_sorter_set_ignore_case"}
	// #endregion

	// #region StyleProvider
	gtk_style_provider_get_type                   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_style_provider_get_type"}
	gtk_style_context_add_provider_for_display    = cc.DlFunc[func(display *gdk.Display, provider *StyleProvider, priority StyleProviderPriority), cc.Void]{Name: "gtk_style_context_add_provider_for_display"}
	gtk_style_context_remove_provider_for_display = cc.DlFunc[func(display *gdk.Display, provider *StyleProvider), cc.Void]{Name: "gtk_style_context_remove_provider_for_display"}
	// #endregion

	// #region Switch
	gtk_switch_get_type   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_switch_get_type"}
	gtk_switch_new        = cc.DlFunc[func() *Switch, *Switch]{Name: "gtk_switch_new"}
	gtk_switch_set_active = cc.DlFunc[func(sw *Switch, isActive int32), cc.Void]{Name: "gtk_switch_set_active"}
	gtk_switch_get_active = cc.DlFunc[func(sw *Switch) int32, int32]{Name: "gtk_switch_get_active"}
	gtk_switch_set_state  = cc.DlFunc[func(sw *Switch, state int32), cc.Void]{Name: "gtk_switch_set_state"}
	gtk_switch_get_state  = cc.DlFunc[func(sw *Switch) int32, int32]{Name: "gtk_switch_get_state"}
	// #endregion

	// #region SymbolicPaintable
	gtk_symbolic_paintable_snapshot_symbolic = cc.DlFunc[func(p *SymbolicPaintable, snapshot *Snapshot, width, height float64, colors *gdk.RGBA, nColors uint64), cc.Void]{Name: "gtk_symbolic_paintable_snapshot_symbolic"}
	// #endregion

	// #region TestAtContext
	gtk_test_accessible_has_role               = cc.DlFunc[func(a *Accessible, role AccessibleRole) int32, int32]{Name: "gtk_test_accessible_has_role"}
	gtk_test_accessible_has_property           = cc.DlFunc[func(a *Accessible, property AccessibleProperty) int32, int32]{Name: "gtk_test_accessible_has_property"}
	gtk_test_accessible_has_relation           = cc.DlFunc[func(a *Accessible, relation AccessibleRelation) int32, int32]{Name: "gtk_test_accessible_has_relation"}
	gtk_test_accessible_has_state              = cc.DlFunc[func(a *Accessible, state AccessibleState) int32, int32]{Name: "gtk_test_accessible_has_state"}
	gtk_test_accessible_check_property         = cc.DlFunc[func(a *Accessible, property AccessibleProperty, args ...interface{}) cc.String, cc.String]{Name: "gtk_test_accessible_check_property", Va: true}
	gtk_test_accessible_check_relation         = cc.DlFunc[func(a *Accessible, relation AccessibleRelation, args ...interface{}) cc.String, cc.String]{Name: "gtk_test_accessible_check_relation", Va: true}
	gtk_test_accessible_check_state            = cc.DlFunc[func(a *Accessible, state AccessibleState, args ...interface{}) cc.String, cc.String]{Name: "gtk_test_accessible_check_state", Va: true}
	gtk_test_accessible_assertion_message_role = cc.DlFunc[func(domain, file, function, expr cc.String, line int32, a *Accessible, expected, actual AccessibleRole), cc.Void]{Name: "gtk_test_accessible_assertion_message_role"}
	// #endregion

	// #region TestUtils
	gtk_test_init                 = cc.DlFunc[func(argcp *int32, argvp *cc.Strings, args ...interface{}), cc.Void]{Name: "gtk_test_init", Va: true}
	gtk_test_register_all_types   = cc.DlFunc[func(), cc.Void]{Name: "gtk_test_register_all_types"}
	gtk_test_list_all_types       = cc.DlFunc[func(nTypes *uint32) *gobject.GType, *gobject.GType]{Name: "gtk_test_list_all_types"}
	gtk_test_widget_wait_for_draw = cc.DlFunc[func(w *Widget), cc.Void]{Name: "gtk_test_widget_wait_for_draw"}
	// #endregion

	// #region Text
	gtk_text_get_type                          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_text_get_type"}
	gtk_text_new                               = cc.DlFunc[func() *Text, *Text]{Name: "gtk_text_new"}
	gtk_text_new_with_buffer                   = cc.DlFunc[func(buffer *EntryBuffer) *Text, *Text]{Name: "gtk_text_new_with_buffer"}
	gtk_text_get_buffer                        = cc.DlFunc[func(t *Text) *EntryBuffer, *EntryBuffer]{Name: "gtk_text_get_buffer"}
	gtk_text_set_buffer                        = cc.DlFunc[func(t *Text, buffer *EntryBuffer), cc.Void]{Name: "gtk_text_set_buffer"}
	gtk_text_set_visibility                    = cc.DlFunc[func(t *Text, visible int32), cc.Void]{Name: "gtk_text_set_visibility"}
	gtk_text_get_visibility                    = cc.DlFunc[func(t *Text) int32, int32]{Name: "gtk_text_get_visibility"}
	gtk_text_set_invisible_char                = cc.DlFunc[func(t *Text, ch uint32), cc.Void]{Name: "gtk_text_set_invisible_char"}
	gtk_text_get_invisible_char                = cc.DlFunc[func(t *Text) uint32, uint32]{Name: "gtk_text_get_invisible_char"}
	gtk_text_unset_invisible_char              = cc.DlFunc[func(t *Text), cc.Void]{Name: "gtk_text_unset_invisible_char"}
	gtk_text_set_overwrite_mode                = cc.DlFunc[func(t *Text, overwrite int32), cc.Void]{Name: "gtk_text_set_overwrite_mode"}
	gtk_text_get_overwrite_mode                = cc.DlFunc[func(t *Text) int32, int32]{Name: "gtk_text_get_overwrite_mode"}
	gtk_text_set_max_length                    = cc.DlFunc[func(t *Text, length int32), cc.Void]{Name: "gtk_text_set_max_length"}
	gtk_text_get_max_length                    = cc.DlFunc[func(t *Text) int32, int32]{Name: "gtk_text_get_max_length"}
	gtk_text_get_text_length                   = cc.DlFunc[func(t *Text) uint16, uint16]{Name: "gtk_text_get_text_length"}
	gtk_text_set_activates_default             = cc.DlFunc[func(t *Text, activates int32), cc.Void]{Name: "gtk_text_set_activates_default"}
	gtk_text_get_activates_default             = cc.DlFunc[func(t *Text) int32, int32]{Name: "gtk_text_get_activates_default"}
	gtk_text_get_placeholder_text              = cc.DlFunc[func(t *Text) cc.String, cc.String]{Name: "gtk_text_get_placeholder_text"}
	gtk_text_set_placeholder_text              = cc.DlFunc[func(t *Text, text cc.String), cc.Void]{Name: "gtk_text_set_placeholder_text"}
	gtk_text_set_input_purpose                 = cc.DlFunc[func(t *Text, purpose InputPurpose), cc.Void]{Name: "gtk_text_set_input_purpose"}
	gtk_text_get_input_purpose                 = cc.DlFunc[func(t *Text) InputPurpose, InputPurpose]{Name: "gtk_text_get_input_purpose"}
	gtk_text_set_input_hints                   = cc.DlFunc[func(t *Text, hints InputHints), cc.Void]{Name: "gtk_text_set_input_hints"}
	gtk_text_get_input_hints                   = cc.DlFunc[func(t *Text) InputHints, InputHints]{Name: "gtk_text_get_input_hints"}
	gtk_text_set_attributes                    = cc.DlFunc[func(t *Text, attrs *pango.AttrList), cc.Void]{Name: "gtk_text_set_attributes"}
	gtk_text_get_attributes                    = cc.DlFunc[func(t *Text) *pango.AttrList, *pango.AttrList]{Name: "gtk_text_get_attributes"}
	gtk_text_set_tabs                          = cc.DlFunc[func(t *Text, tabs *pango.TabArray), cc.Void]{Name: "gtk_text_set_tabs"}
	gtk_text_get_tabs                          = cc.DlFunc[func(t *Text) *pango.TabArray, *pango.TabArray]{Name: "gtk_text_get_tabs"}
	gtk_text_grab_focus_without_selecting      = cc.DlFunc[func(t *Text) int32, int32]{Name: "gtk_text_grab_focus_without_selecting"}
	gtk_text_set_extra_menu                    = cc.DlFunc[func(t *Text, model *gio.GMenuModel), cc.Void]{Name: "gtk_text_set_extra_menu"}
	gtk_text_get_extra_menu                    = cc.DlFunc[func(t *Text) *gio.GMenuModel, *gio.GMenuModel]{Name: "gtk_text_get_extra_menu"}
	gtk_text_set_enable_emoji_completion       = cc.DlFunc[func(t *Text, enable int32), cc.Void]{Name: "gtk_text_set_enable_emoji_completion"}
	gtk_text_get_enable_emoji_completion       = cc.DlFunc[func(t *Text) int32, int32]{Name: "gtk_text_get_enable_emoji_completion"}
	gtk_text_set_propagate_text_width          = cc.DlFunc[func(t *Text, propagate int32), cc.Void]{Name: "gtk_text_set_propagate_text_width"}
	gtk_text_get_propagate_text_width          = cc.DlFunc[func(t *Text) int32, int32]{Name: "gtk_text_get_propagate_text_width"}
	gtk_text_set_truncate_multiline            = cc.DlFunc[func(t *Text, truncate int32), cc.Void]{Name: "gtk_text_set_truncate_multiline"}
	gtk_text_get_truncate_multiline            = cc.DlFunc[func(t *Text) int32, int32]{Name: "gtk_text_get_truncate_multiline"}
	gtk_text_compute_cursor_extents            = cc.DlFunc[func(t *Text, position uint64, strong, weak *graphene.Rect), cc.Void]{Name: "gtk_text_compute_cursor_extents"}
	gtk_text_buffer_apply_tag                  = cc.DlFunc[func(buffer *TextBuffer, tag *TextTag, start, end *TextIter), cc.Void]{Name: "gtk_text_buffer_apply_tag"}
	gtk_text_buffer_remove_tag                 = cc.DlFunc[func(buffer *TextBuffer, tag *TextTag, start, end *TextIter), cc.Void]{Name: "gtk_text_buffer_remove_tag"}
	gtk_text_buffer_apply_tag_by_name          = cc.DlFunc[func(buffer *TextBuffer, name cc.String, start, end *TextIter), cc.Void]{Name: "gtk_text_buffer_apply_tag_by_name"}
	gtk_text_buffer_remove_tag_by_name         = cc.DlFunc[func(buffer *TextBuffer, name cc.String, start, end *TextIter), cc.Void]{Name: "gtk_text_buffer_remove_tag_by_name"}
	gtk_text_buffer_remove_all_tags            = cc.DlFunc[func(buffer *TextBuffer, start, end *TextIter), cc.Void]{Name: "gtk_text_buffer_remove_all_tags"}
	gtk_text_buffer_create_tag                 = cc.DlFunc[func(buffer *TextBuffer, tagName cc.String, firstPropertyName cc.String, args ...interface{}) *TextTag, *TextTag]{Name: "gtk_text_buffer_create_tag", Va: true}
	gtk_text_buffer_get_iter_at_line_offset    = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, lineNumber, charOffset int32) int32, int32]{Name: "gtk_text_buffer_get_iter_at_line_offset"}
	gtk_text_buffer_get_iter_at_line_index     = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, lineNumber, byteIndex int32) int32, int32]{Name: "gtk_text_buffer_get_iter_at_line_index"}
	gtk_text_buffer_get_iter_at_offset         = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, charOffset int32), cc.Void]{Name: "gtk_text_buffer_get_iter_at_offset"}
	gtk_text_buffer_get_iter_at_line           = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, lineNumber int32) int32, int32]{Name: "gtk_text_buffer_get_iter_at_line"}
	gtk_text_buffer_get_start_iter             = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter), cc.Void]{Name: "gtk_text_buffer_get_start_iter"}
	gtk_text_buffer_get_end_iter               = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter), cc.Void]{Name: "gtk_text_buffer_get_end_iter"}
	gtk_text_buffer_get_bounds                 = cc.DlFunc[func(buffer *TextBuffer, start, end *TextIter), cc.Void]{Name: "gtk_text_buffer_get_bounds"}
	gtk_text_buffer_get_iter_at_mark           = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, mark *TextMark), cc.Void]{Name: "gtk_text_buffer_get_iter_at_mark"}
	gtk_text_buffer_get_iter_at_child_anchor   = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, anchor *TextChildAnchor), cc.Void]{Name: "gtk_text_buffer_get_iter_at_child_anchor"}
	gtk_text_buffer_get_modified               = cc.DlFunc[func(buffer *TextBuffer) int32, int32]{Name: "gtk_text_buffer_get_modified"}
	gtk_text_buffer_set_modified               = cc.DlFunc[func(buffer *TextBuffer, setting int32), cc.Void]{Name: "gtk_text_buffer_set_modified"}
	gtk_text_buffer_get_has_selection          = cc.DlFunc[func(buffer *TextBuffer) int32, int32]{Name: "gtk_text_buffer_get_has_selection"}
	gtk_text_buffer_add_selection_clipboard    = cc.DlFunc[func(buffer *TextBuffer, clipboard *gdk.Clipboard), cc.Void]{Name: "gtk_text_buffer_add_selection_clipboard"}
	gtk_text_buffer_remove_selection_clipboard = cc.DlFunc[func(buffer *TextBuffer, clipboard *gdk.Clipboard), cc.Void]{Name: "gtk_text_buffer_remove_selection_clipboard"}
	gtk_text_buffer_cut_clipboard              = cc.DlFunc[func(buffer *TextBuffer, clipboard *gdk.Clipboard, defaultEditable int32), cc.Void]{Name: "gtk_text_buffer_cut_clipboard"}
	gtk_text_buffer_copy_clipboard             = cc.DlFunc[func(buffer *TextBuffer, clipboard *gdk.Clipboard), cc.Void]{Name: "gtk_text_buffer_copy_clipboard"}
	gtk_text_buffer_paste_clipboard            = cc.DlFunc[func(buffer *TextBuffer, clipboard *gdk.Clipboard, overrideLocation *TextIter, defaultEditable int32), cc.Void]{Name: "gtk_text_buffer_paste_clipboard"}
	gtk_text_buffer_get_selection_bounds       = cc.DlFunc[func(buffer *TextBuffer, start, end *TextIter) int32, int32]{Name: "gtk_text_buffer_get_selection_bounds"}
	gtk_text_buffer_delete_selection           = cc.DlFunc[func(buffer *TextBuffer, interactive, defaultEditable int32) int32, int32]{Name: "gtk_text_buffer_delete_selection"}
	gtk_text_buffer_get_selection_content      = cc.DlFunc[func(buffer *TextBuffer) *gdk.ContentProvider, *gdk.ContentProvider]{Name: "gtk_text_buffer_get_selection_content"}
	gtk_text_buffer_get_can_undo               = cc.DlFunc[func(buffer *TextBuffer) int32, int32]{Name: "gtk_text_buffer_get_can_undo"}
	gtk_text_buffer_get_can_redo               = cc.DlFunc[func(buffer *TextBuffer) int32, int32]{Name: "gtk_text_buffer_get_can_redo"}
	gtk_text_buffer_get_enable_undo            = cc.DlFunc[func(buffer *TextBuffer) int32, int32]{Name: "gtk_text_buffer_get_enable_undo"}
	gtk_text_buffer_set_enable_undo            = cc.DlFunc[func(buffer *TextBuffer, enableUndo int32), cc.Void]{Name: "gtk_text_buffer_set_enable_undo"}
	gtk_text_buffer_get_max_undo_levels        = cc.DlFunc[func(buffer *TextBuffer) uint32, uint32]{Name: "gtk_text_buffer_get_max_undo_levels"}
	gtk_text_buffer_set_max_undo_levels        = cc.DlFunc[func(buffer *TextBuffer, maxUndoLevels uint32), cc.Void]{Name: "gtk_text_buffer_set_max_undo_levels"}
	gtk_text_buffer_undo                       = cc.DlFunc[func(buffer *TextBuffer), cc.Void]{Name: "gtk_text_buffer_undo"}
	gtk_text_buffer_redo                       = cc.DlFunc[func(buffer *TextBuffer), cc.Void]{Name: "gtk_text_buffer_redo"}
	gtk_text_buffer_begin_irreversible_action  = cc.DlFunc[func(buffer *TextBuffer), cc.Void]{Name: "gtk_text_buffer_begin_irreversible_action"}
	gtk_text_buffer_end_irreversible_action    = cc.DlFunc[func(buffer *TextBuffer), cc.Void]{Name: "gtk_text_buffer_end_irreversible_action"}
	gtk_text_buffer_begin_user_action          = cc.DlFunc[func(buffer *TextBuffer), cc.Void]{Name: "gtk_text_buffer_begin_user_action"}
	gtk_text_buffer_end_user_action            = cc.DlFunc[func(buffer *TextBuffer), cc.Void]{Name: "gtk_text_buffer_end_user_action"}
	gtk_text_buffer_add_commit_notify          = cc.DlFunc[func(buffer *TextBuffer, flags uint32, commitNotify uptr, userData uptr, destroy uptr) uint32, uint32]{Name: "gtk_text_buffer_add_commit_notify"}
	gtk_text_buffer_remove_commit_notify       = cc.DlFunc[func(buffer *TextBuffer, handler uint32), cc.Void]{Name: "gtk_text_buffer_remove_commit_notify"}
	// #endregion

	// #region TextBuffer
	gtk_text_buffer_get_type                     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_text_buffer_get_type"}
	gtk_text_buffer_new                          = cc.DlFunc[func(table *TextTagTable) *TextBuffer, *TextBuffer]{Name: "gtk_text_buffer_new"}
	gtk_text_buffer_get_line_count               = cc.DlFunc[func(buffer *TextBuffer) int32, int32]{Name: "gtk_text_buffer_get_line_count"}
	gtk_text_buffer_get_char_count               = cc.DlFunc[func(buffer *TextBuffer) int32, int32]{Name: "gtk_text_buffer_get_char_count"}
	gtk_text_buffer_get_tag_table                = cc.DlFunc[func(buffer *TextBuffer) *TextTagTable, *TextTagTable]{Name: "gtk_text_buffer_get_tag_table"}
	gtk_text_buffer_set_text                     = cc.DlFunc[func(buffer *TextBuffer, text cc.String, len int32), cc.Void]{Name: "gtk_text_buffer_set_text"}
	gtk_text_buffer_insert                       = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, text cc.String, len int32), cc.Void]{Name: "gtk_text_buffer_insert"}
	gtk_text_buffer_insert_at_cursor             = cc.DlFunc[func(buffer *TextBuffer, text cc.String, len int32), cc.Void]{Name: "gtk_text_buffer_insert_at_cursor"}
	gtk_text_buffer_insert_interactive           = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, text cc.String, len int32, defaultEditable int32) int32, int32]{Name: "gtk_text_buffer_insert_interactive"}
	gtk_text_buffer_insert_interactive_at_cursor = cc.DlFunc[func(buffer *TextBuffer, text cc.String, len int32, defaultEditable int32) int32, int32]{Name: "gtk_text_buffer_insert_interactive_at_cursor"}
	gtk_text_buffer_insert_range                 = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, start, end *TextIter), cc.Void]{Name: "gtk_text_buffer_insert_range"}
	gtk_text_buffer_insert_range_interactive     = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, start, end *TextIter, defaultEditable int32) int32, int32]{Name: "gtk_text_buffer_insert_range_interactive"}
	gtk_text_buffer_insert_with_tags             = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, text cc.String, len int32, firstTag *TextTag, args ...interface{}), cc.Void]{Name: "gtk_text_buffer_insert_with_tags", Va: true}
	gtk_text_buffer_insert_with_tags_by_name     = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, text cc.String, len int32, firstTagName cc.String, args ...interface{}), cc.Void]{Name: "gtk_text_buffer_insert_with_tags_by_name", Va: true}
	gtk_text_buffer_insert_markup                = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, markup cc.String, len int32), cc.Void]{Name: "gtk_text_buffer_insert_markup"}
	gtk_text_buffer_delete                       = cc.DlFunc[func(buffer *TextBuffer, start, end *TextIter), cc.Void]{Name: "gtk_text_buffer_delete"}
	gtk_text_buffer_delete_interactive           = cc.DlFunc[func(buffer *TextBuffer, startIter, endIter *TextIter, defaultEditable int32) int32, int32]{Name: "gtk_text_buffer_delete_interactive"}
	gtk_text_buffer_backspace                    = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, interactive, defaultEditable int32) int32, int32]{Name: "gtk_text_buffer_backspace"}
	gtk_text_buffer_get_text                     = cc.DlFunc[func(buffer *TextBuffer, start, end *TextIter, includeHiddenChars int32) cc.String, cc.String]{Name: "gtk_text_buffer_get_text"}
	gtk_text_buffer_get_slice                    = cc.DlFunc[func(buffer *TextBuffer, start, end *TextIter, includeHiddenChars int32) cc.String, cc.String]{Name: "gtk_text_buffer_get_slice"}
	gtk_text_buffer_insert_paintable             = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, paintable *gdk.Paintable), cc.Void]{Name: "gtk_text_buffer_insert_paintable"}
	gtk_text_buffer_insert_child_anchor          = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter, anchor *TextChildAnchor), cc.Void]{Name: "gtk_text_buffer_insert_child_anchor"}
	gtk_text_buffer_create_child_anchor          = cc.DlFunc[func(buffer *TextBuffer, iter *TextIter) *TextChildAnchor, *TextChildAnchor]{Name: "gtk_text_buffer_create_child_anchor"}
	gtk_text_buffer_add_mark                     = cc.DlFunc[func(buffer *TextBuffer, mark *TextMark, where *TextIter), cc.Void]{Name: "gtk_text_buffer_add_mark"}
	gtk_text_buffer_create_mark                  = cc.DlFunc[func(buffer *TextBuffer, markName cc.String, where *TextIter, leftGravity int32) *TextMark, *TextMark]{Name: "gtk_text_buffer_create_mark"}
	gtk_text_buffer_move_mark                    = cc.DlFunc[func(buffer *TextBuffer, mark *TextMark, where *TextIter), cc.Void]{Name: "gtk_text_buffer_move_mark"}
	gtk_text_buffer_delete_mark                  = cc.DlFunc[func(buffer *TextBuffer, mark *TextMark), cc.Void]{Name: "gtk_text_buffer_delete_mark"}
	gtk_text_buffer_get_mark                     = cc.DlFunc[func(buffer *TextBuffer, name cc.String) *TextMark, *TextMark]{Name: "gtk_text_buffer_get_mark"}
	gtk_text_buffer_move_mark_by_name            = cc.DlFunc[func(buffer *TextBuffer, name cc.String, where *TextIter), cc.Void]{Name: "gtk_text_buffer_move_mark_by_name"}
	gtk_text_buffer_delete_mark_by_name          = cc.DlFunc[func(buffer *TextBuffer, name cc.String), cc.Void]{Name: "gtk_text_buffer_delete_mark_by_name"}
	gtk_text_buffer_get_insert                   = cc.DlFunc[func(buffer *TextBuffer) *TextMark, *TextMark]{Name: "gtk_text_buffer_get_insert"}
	gtk_text_buffer_get_selection_bound          = cc.DlFunc[func(buffer *TextBuffer) *TextMark, *TextMark]{Name: "gtk_text_buffer_get_selection_bound"}
	gtk_text_buffer_place_cursor                 = cc.DlFunc[func(buffer *TextBuffer, where *TextIter), cc.Void]{Name: "gtk_text_buffer_place_cursor"}
	gtk_text_buffer_select_range                 = cc.DlFunc[func(buffer *TextBuffer, ins, bound *TextIter), cc.Void]{Name: "gtk_text_buffer_select_range"}
	// #endregion

	// #region TextChildAnchor
	gtk_text_child_anchor_get_type             = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_text_child_anchor_get_type"}
	gtk_text_child_anchor_new                  = cc.DlFunc[func() *TextChildAnchor, *TextChildAnchor]{Name: "gtk_text_child_anchor_new"}
	gtk_text_child_anchor_new_with_replacement = cc.DlFunc[func(character cc.String) *TextChildAnchor, *TextChildAnchor]{Name: "gtk_text_child_anchor_new_with_replacement"}
	gtk_text_child_anchor_get_widgets          = cc.DlFunc[func(anchor *TextChildAnchor, outLen *uint32) **Widget, **Widget]{Name: "gtk_text_child_anchor_get_widgets"}
	gtk_text_child_anchor_get_deleted          = cc.DlFunc[func(anchor *TextChildAnchor) int32, int32]{Name: "gtk_text_child_anchor_get_deleted"}
	// #endregion

	// #region TextIter
	gtk_text_iter_get_buffer                        = cc.DlFunc[func(iter *TextIter) *TextBuffer, *TextBuffer]{Name: "gtk_text_iter_get_buffer"}
	gtk_text_iter_copy                              = cc.DlFunc[func(iter *TextIter) *TextIter, *TextIter]{Name: "gtk_text_iter_copy"}
	gtk_text_iter_free                              = cc.DlFunc[func(iter *TextIter), cc.Void]{Name: "gtk_text_iter_free"}
	gtk_text_iter_assign                            = cc.DlFunc[func(iter *TextIter, other *TextIter), cc.Void]{Name: "gtk_text_iter_assign"}
	gtk_text_iter_get_type                          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_text_iter_get_type"}
	gtk_text_iter_get_offset                        = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_get_offset"}
	gtk_text_iter_get_line                          = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_get_line"}
	gtk_text_iter_get_line_offset                   = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_get_line_offset"}
	gtk_text_iter_get_line_index                    = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_get_line_index"}
	gtk_text_iter_get_visible_line_offset           = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_get_visible_line_offset"}
	gtk_text_iter_get_visible_line_index            = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_get_visible_line_index"}
	gtk_text_iter_get_char                          = cc.DlFunc[func(iter *TextIter) uint32, uint32]{Name: "gtk_text_iter_get_char"}
	gtk_text_iter_get_slice                         = cc.DlFunc[func(start, end *TextIter) cc.String, cc.String]{Name: "gtk_text_iter_get_slice"}
	gtk_text_iter_get_text                          = cc.DlFunc[func(start, end *TextIter) cc.String, cc.String]{Name: "gtk_text_iter_get_text"}
	gtk_text_iter_get_visible_slice                 = cc.DlFunc[func(start, end *TextIter) cc.String, cc.String]{Name: "gtk_text_iter_get_visible_slice"}
	gtk_text_iter_get_visible_text                  = cc.DlFunc[func(start, end *TextIter) cc.String, cc.String]{Name: "gtk_text_iter_get_visible_text"}
	gtk_text_iter_get_paintable                     = cc.DlFunc[func(iter *TextIter) *gdk.Paintable, *gdk.Paintable]{Name: "gtk_text_iter_get_paintable"}
	gtk_text_iter_get_marks                         = cc.DlFunc[func(iter *TextIter) uptr, uptr]{Name: "gtk_text_iter_get_marks"}
	gtk_text_iter_get_child_anchor                  = cc.DlFunc[func(iter *TextIter) *TextChildAnchor, *TextChildAnchor]{Name: "gtk_text_iter_get_child_anchor"}
	gtk_text_iter_get_toggled_tags                  = cc.DlFunc[func(iter *TextIter, toggledOn int32) uptr, uptr]{Name: "gtk_text_iter_get_toggled_tags"}
	gtk_text_iter_starts_tag                        = cc.DlFunc[func(iter *TextIter, tag *TextTag) int32, int32]{Name: "gtk_text_iter_starts_tag"}
	gtk_text_iter_ends_tag                          = cc.DlFunc[func(iter *TextIter, tag *TextTag) int32, int32]{Name: "gtk_text_iter_ends_tag"}
	gtk_text_iter_toggles_tag                       = cc.DlFunc[func(iter *TextIter, tag *TextTag) int32, int32]{Name: "gtk_text_iter_toggles_tag"}
	gtk_text_iter_has_tag                           = cc.DlFunc[func(iter *TextIter, tag *TextTag) int32, int32]{Name: "gtk_text_iter_has_tag"}
	gtk_text_iter_get_tags                          = cc.DlFunc[func(iter *TextIter) uptr, uptr]{Name: "gtk_text_iter_get_tags"}
	gtk_text_iter_editable                          = cc.DlFunc[func(iter *TextIter, defaultSetting int32) int32, int32]{Name: "gtk_text_iter_editable"}
	gtk_text_iter_can_insert                        = cc.DlFunc[func(iter *TextIter, defaultEditability int32) int32, int32]{Name: "gtk_text_iter_can_insert"}
	gtk_text_iter_starts_word                       = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_starts_word"}
	gtk_text_iter_ends_word                         = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_ends_word"}
	gtk_text_iter_inside_word                       = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_inside_word"}
	gtk_text_iter_starts_sentence                   = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_starts_sentence"}
	gtk_text_iter_ends_sentence                     = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_ends_sentence"}
	gtk_text_iter_inside_sentence                   = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_inside_sentence"}
	gtk_text_iter_starts_line                       = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_starts_line"}
	gtk_text_iter_ends_line                         = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_ends_line"}
	gtk_text_iter_is_cursor_position                = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_is_cursor_position"}
	gtk_text_iter_get_chars_in_line                 = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_get_chars_in_line"}
	gtk_text_iter_get_bytes_in_line                 = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_get_bytes_in_line"}
	gtk_text_iter_get_language                      = cc.DlFunc[func(iter *TextIter) *pango.Language, *pango.Language]{Name: "gtk_text_iter_get_language"}
	gtk_text_iter_is_end                            = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_is_end"}
	gtk_text_iter_is_start                          = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_is_start"}
	gtk_text_iter_forward_char                      = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_forward_char"}
	gtk_text_iter_backward_char                     = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_backward_char"}
	gtk_text_iter_forward_chars                     = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_forward_chars"}
	gtk_text_iter_backward_chars                    = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_backward_chars"}
	gtk_text_iter_forward_line                      = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_forward_line"}
	gtk_text_iter_backward_line                     = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_backward_line"}
	gtk_text_iter_forward_lines                     = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_forward_lines"}
	gtk_text_iter_backward_lines                    = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_backward_lines"}
	gtk_text_iter_forward_word_end                  = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_forward_word_end"}
	gtk_text_iter_backward_word_start               = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_backward_word_start"}
	gtk_text_iter_forward_word_ends                 = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_forward_word_ends"}
	gtk_text_iter_backward_word_starts              = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_backward_word_starts"}
	gtk_text_iter_forward_visible_line              = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_forward_visible_line"}
	gtk_text_iter_backward_visible_line             = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_backward_visible_line"}
	gtk_text_iter_forward_visible_lines             = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_forward_visible_lines"}
	gtk_text_iter_backward_visible_lines            = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_backward_visible_lines"}
	gtk_text_iter_forward_visible_word_end          = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_forward_visible_word_end"}
	gtk_text_iter_backward_visible_word_start       = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_backward_visible_word_start"}
	gtk_text_iter_forward_visible_word_ends         = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_forward_visible_word_ends"}
	gtk_text_iter_backward_visible_word_starts      = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_backward_visible_word_starts"}
	gtk_text_iter_forward_sentence_end              = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_forward_sentence_end"}
	gtk_text_iter_backward_sentence_start           = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_backward_sentence_start"}
	gtk_text_iter_forward_sentence_ends             = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_forward_sentence_ends"}
	gtk_text_iter_backward_sentence_starts          = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_backward_sentence_starts"}
	gtk_text_iter_forward_cursor_position           = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_forward_cursor_position"}
	gtk_text_iter_backward_cursor_position          = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_backward_cursor_position"}
	gtk_text_iter_forward_cursor_positions          = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_forward_cursor_positions"}
	gtk_text_iter_backward_cursor_positions         = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_backward_cursor_positions"}
	gtk_text_iter_forward_visible_cursor_position   = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_forward_visible_cursor_position"}
	gtk_text_iter_backward_visible_cursor_position  = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_backward_visible_cursor_position"}
	gtk_text_iter_forward_visible_cursor_positions  = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_forward_visible_cursor_positions"}
	gtk_text_iter_backward_visible_cursor_positions = cc.DlFunc[func(iter *TextIter, count int32) int32, int32]{Name: "gtk_text_iter_backward_visible_cursor_positions"}
	gtk_text_iter_set_offset                        = cc.DlFunc[func(iter *TextIter, charOffset int32), cc.Void]{Name: "gtk_text_iter_set_offset"}
	gtk_text_iter_set_line                          = cc.DlFunc[func(iter *TextIter, lineNumber int32), cc.Void]{Name: "gtk_text_iter_set_line"}
	gtk_text_iter_set_line_offset                   = cc.DlFunc[func(iter *TextIter, charOnLine int32), cc.Void]{Name: "gtk_text_iter_set_line_offset"}
	gtk_text_iter_set_line_index                    = cc.DlFunc[func(iter *TextIter, byteOnLine int32), cc.Void]{Name: "gtk_text_iter_set_line_index"}
	gtk_text_iter_forward_to_end                    = cc.DlFunc[func(iter *TextIter), cc.Void]{Name: "gtk_text_iter_forward_to_end"}
	gtk_text_iter_forward_to_line_end               = cc.DlFunc[func(iter *TextIter) int32, int32]{Name: "gtk_text_iter_forward_to_line_end"}
	gtk_text_iter_set_visible_line_offset           = cc.DlFunc[func(iter *TextIter, charOnLine int32), cc.Void]{Name: "gtk_text_iter_set_visible_line_offset"}
	gtk_text_iter_set_visible_line_index            = cc.DlFunc[func(iter *TextIter, byteOnLine int32), cc.Void]{Name: "gtk_text_iter_set_visible_line_index"}
	gtk_text_iter_forward_to_tag_toggle             = cc.DlFunc[func(iter *TextIter, tag *TextTag) int32, int32]{Name: "gtk_text_iter_forward_to_tag_toggle"}
	gtk_text_iter_backward_to_tag_toggle            = cc.DlFunc[func(iter *TextIter, tag *TextTag) int32, int32]{Name: "gtk_text_iter_backward_to_tag_toggle"}
	gtk_text_iter_forward_find_char                 = cc.DlFunc[func(iter *TextIter, pred uptr, userData uptr, limit *TextIter) int32, int32]{Name: "gtk_text_iter_forward_find_char"}
	gtk_text_iter_backward_find_char                = cc.DlFunc[func(iter *TextIter, pred uptr, userData uptr, limit *TextIter) int32, int32]{Name: "gtk_text_iter_backward_find_char"}
	gtk_text_iter_forward_search                    = cc.DlFunc[func(iter *TextIter, str cc.String, flags uint32, matchStart, matchEnd *TextIter, limit *TextIter) int32, int32]{Name: "gtk_text_iter_forward_search"}
	gtk_text_iter_backward_search                   = cc.DlFunc[func(iter *TextIter, str cc.String, flags uint32, matchStart, matchEnd *TextIter, limit *TextIter) int32, int32]{Name: "gtk_text_iter_backward_search"}
	gtk_text_iter_equal                             = cc.DlFunc[func(lhs *TextIter, rhs *TextIter) int32, int32]{Name: "gtk_text_iter_equal"}
	gtk_text_iter_compare                           = cc.DlFunc[func(lhs *TextIter, rhs *TextIter) int32, int32]{Name: "gtk_text_iter_compare"}
	gtk_text_iter_in_range                          = cc.DlFunc[func(iter *TextIter, start *TextIter, end *TextIter) int32, int32]{Name: "gtk_text_iter_in_range"}
	gtk_text_iter_order                             = cc.DlFunc[func(first *TextIter, second *TextIter), cc.Void]{Name: "gtk_text_iter_order"}
	// #endregion

	// #region TextTag
	gtk_text_tag_get_type     = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_text_tag_get_type"}
	gtk_text_tag_new          = cc.DlFunc[func(name cc.String) *TextTag, *TextTag]{Name: "gtk_text_tag_new"}
	gtk_text_tag_get_priority = cc.DlFunc[func(tag *TextTag) int32, int32]{Name: "gtk_text_tag_get_priority"}
	gtk_text_tag_set_priority = cc.DlFunc[func(tag *TextTag, priority int32), cc.Void]{Name: "gtk_text_tag_set_priority"}
	gtk_text_tag_changed      = cc.DlFunc[func(tag *TextTag, sizeChanged int32), cc.Void]{Name: "gtk_text_tag_changed"}
	// #endregion

	// #region TextMark
	gtk_text_mark_get_type         = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_text_mark_get_type"}
	gtk_text_mark_new              = cc.DlFunc[func(name cc.String, leftGravity int32) *TextMark, *TextMark]{Name: "gtk_text_mark_new"}
	gtk_text_mark_set_visible      = cc.DlFunc[func(mark *TextMark, setting int32), cc.Void]{Name: "gtk_text_mark_set_visible"}
	gtk_text_mark_get_visible      = cc.DlFunc[func(mark *TextMark) int32, int32]{Name: "gtk_text_mark_get_visible"}
	gtk_text_mark_get_name         = cc.DlFunc[func(mark *TextMark) cc.String, cc.String]{Name: "gtk_text_mark_get_name"}
	gtk_text_mark_get_deleted      = cc.DlFunc[func(mark *TextMark) int32, int32]{Name: "gtk_text_mark_get_deleted"}
	gtk_text_mark_get_buffer       = cc.DlFunc[func(mark *TextMark) *TextBuffer, *TextBuffer]{Name: "gtk_text_mark_get_buffer"}
	gtk_text_mark_get_left_gravity = cc.DlFunc[func(mark *TextMark) int32, int32]{Name: "gtk_text_mark_get_left_gravity"}
	// #endregion

	// #region TextTagTable
	gtk_text_tag_table_get_type = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_text_tag_table_get_type"}
	gtk_text_tag_table_new      = cc.DlFunc[func() *TextTagTable, *TextTagTable]{Name: "gtk_text_tag_table_new"}
	gtk_text_tag_table_add      = cc.DlFunc[func(table *TextTagTable, tag *TextTag) int32, int32]{Name: "gtk_text_tag_table_add"}
	gtk_text_tag_table_remove   = cc.DlFunc[func(table *TextTagTable, tag *TextTag), cc.Void]{Name: "gtk_text_tag_table_remove"}
	gtk_text_tag_table_lookup   = cc.DlFunc[func(table *TextTagTable, name cc.String) *TextTag, *TextTag]{Name: "gtk_text_tag_table_lookup"}
	gtk_text_tag_table_foreach  = cc.DlFunc[func(table *TextTagTable, fn uptr, data uptr), cc.Void]{Name: "gtk_text_tag_table_foreach"}
	gtk_text_tag_table_get_size = cc.DlFunc[func(table *TextTagTable) int32, int32]{Name: "gtk_text_tag_table_get_size"}
	// #endregion

	// #region TextView
	gtk_text_view_get_type                    = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_text_view_get_type"}
	gtk_text_view_new                         = cc.DlFunc[func() *TextView, *TextView]{Name: "gtk_text_view_new"}
	gtk_text_view_new_with_buffer             = cc.DlFunc[func(buffer *TextBuffer) *TextView, *TextView]{Name: "gtk_text_view_new_with_buffer"}
	gtk_text_view_set_buffer                  = cc.DlFunc[func(textView *TextView, buffer *TextBuffer), cc.Void]{Name: "gtk_text_view_set_buffer"}
	gtk_text_view_get_buffer                  = cc.DlFunc[func(textView *TextView) *TextBuffer, *TextBuffer]{Name: "gtk_text_view_get_buffer"}
	gtk_text_view_scroll_to_iter              = cc.DlFunc[func(textView *TextView, iter *TextIter, withinMargin float64, useAlign int32, xalign float64, yalign float64) int32, int32]{Name: "gtk_text_view_scroll_to_iter"}
	gtk_text_view_scroll_to_mark              = cc.DlFunc[func(textView *TextView, mark *TextMark, withinMargin float64, useAlign int32, xalign float64, yalign float64), cc.Void]{Name: "gtk_text_view_scroll_to_mark"}
	gtk_text_view_scroll_mark_onscreen        = cc.DlFunc[func(textView *TextView, mark *TextMark), cc.Void]{Name: "gtk_text_view_scroll_mark_onscreen"}
	gtk_text_view_move_mark_onscreen          = cc.DlFunc[func(textView *TextView, mark *TextMark) int32, int32]{Name: "gtk_text_view_move_mark_onscreen"}
	gtk_text_view_place_cursor_onscreen       = cc.DlFunc[func(textView *TextView) int32, int32]{Name: "gtk_text_view_place_cursor_onscreen"}
	gtk_text_view_get_visible_rect            = cc.DlFunc[func(textView *TextView, visibleRect *gdk.Rectangle), cc.Void]{Name: "gtk_text_view_get_visible_rect"}
	gtk_text_view_get_visible_offset          = cc.DlFunc[func(textView *TextView, xOffset *float64, yOffset *float64), cc.Void]{Name: "gtk_text_view_get_visible_offset"}
	gtk_text_view_set_cursor_visible          = cc.DlFunc[func(textView *TextView, setting int32), cc.Void]{Name: "gtk_text_view_set_cursor_visible"}
	gtk_text_view_get_cursor_visible          = cc.DlFunc[func(textView *TextView) int32, int32]{Name: "gtk_text_view_get_cursor_visible"}
	gtk_text_view_reset_cursor_blink          = cc.DlFunc[func(textView *TextView), cc.Void]{Name: "gtk_text_view_reset_cursor_blink"}
	gtk_text_view_get_cursor_locations        = cc.DlFunc[func(textView *TextView, iter *TextIter, strong, weak *gdk.Rectangle), cc.Void]{Name: "gtk_text_view_get_cursor_locations"}
	gtk_text_view_get_iter_location           = cc.DlFunc[func(textView *TextView, iter *TextIter, location *gdk.Rectangle), cc.Void]{Name: "gtk_text_view_get_iter_location"}
	gtk_text_view_get_iter_at_location        = cc.DlFunc[func(textView *TextView, iter *TextIter, x, y int32) int32, int32]{Name: "gtk_text_view_get_iter_at_location"}
	gtk_text_view_get_iter_at_position        = cc.DlFunc[func(textView *TextView, iter *TextIter, trailing *int32, x, y int32) int32, int32]{Name: "gtk_text_view_get_iter_at_position"}
	gtk_text_view_get_line_yrange             = cc.DlFunc[func(textView *TextView, iter *TextIter, y, height *int32), cc.Void]{Name: "gtk_text_view_get_line_yrange"}
	gtk_text_view_get_line_at_y               = cc.DlFunc[func(textView *TextView, targetIter *TextIter, y int32, lineTop *int32), cc.Void]{Name: "gtk_text_view_get_line_at_y"}
	gtk_text_view_buffer_to_window_coords     = cc.DlFunc[func(textView *TextView, win int32, bufferX, bufferY int32, windowX, windowY *int32), cc.Void]{Name: "gtk_text_view_buffer_to_window_coords"}
	gtk_text_view_window_to_buffer_coords     = cc.DlFunc[func(textView *TextView, win int32, windowX, windowY int32, bufferX, bufferY *int32), cc.Void]{Name: "gtk_text_view_window_to_buffer_coords"}
	gtk_text_view_forward_display_line        = cc.DlFunc[func(textView *TextView, iter *TextIter) int32, int32]{Name: "gtk_text_view_forward_display_line"}
	gtk_text_view_backward_display_line       = cc.DlFunc[func(textView *TextView, iter *TextIter) int32, int32]{Name: "gtk_text_view_backward_display_line"}
	gtk_text_view_forward_display_line_end    = cc.DlFunc[func(textView *TextView, iter *TextIter) int32, int32]{Name: "gtk_text_view_forward_display_line_end"}
	gtk_text_view_backward_display_line_start = cc.DlFunc[func(textView *TextView, iter *TextIter) int32, int32]{Name: "gtk_text_view_backward_display_line_start"}
	gtk_text_view_starts_display_line         = cc.DlFunc[func(textView *TextView, iter *TextIter) int32, int32]{Name: "gtk_text_view_starts_display_line"}
	gtk_text_view_move_visually               = cc.DlFunc[func(textView *TextView, iter *TextIter, count int32) int32, int32]{Name: "gtk_text_view_move_visually"}
	gtk_text_view_im_context_filter_keypress  = cc.DlFunc[func(textView *TextView, event *gdk.Event) int32, int32]{Name: "gtk_text_view_im_context_filter_keypress"}
	gtk_text_view_reset_im_context            = cc.DlFunc[func(textView *TextView), cc.Void]{Name: "gtk_text_view_reset_im_context"}
	gtk_text_view_get_gutter                  = cc.DlFunc[func(textView *TextView, win int32) *Widget, *Widget]{Name: "gtk_text_view_get_gutter"}
	gtk_text_view_set_gutter                  = cc.DlFunc[func(textView *TextView, win int32, widget *Widget), cc.Void]{Name: "gtk_text_view_set_gutter"}
	gtk_text_view_add_child_at_anchor         = cc.DlFunc[func(textView *TextView, child *Widget, anchor *TextChildAnchor), cc.Void]{Name: "gtk_text_view_add_child_at_anchor"}
	gtk_text_view_add_overlay                 = cc.DlFunc[func(textView *TextView, child *Widget, xpos, ypos int32), cc.Void]{Name: "gtk_text_view_add_overlay"}
	gtk_text_view_move_overlay                = cc.DlFunc[func(textView *TextView, child *Widget, xpos, ypos int32), cc.Void]{Name: "gtk_text_view_move_overlay"}
	gtk_text_view_remove                      = cc.DlFunc[func(textView *TextView, child *Widget), cc.Void]{Name: "gtk_text_view_remove"}
	gtk_text_view_set_wrap_mode               = cc.DlFunc[func(textView *TextView, wrapMode WrapMode), cc.Void]{Name: "gtk_text_view_set_wrap_mode"}
	gtk_text_view_get_wrap_mode               = cc.DlFunc[func(textView *TextView) WrapMode, WrapMode]{Name: "gtk_text_view_get_wrap_mode"}
	gtk_text_view_set_editable                = cc.DlFunc[func(textView *TextView, setting int32), cc.Void]{Name: "gtk_text_view_set_editable"}
	gtk_text_view_get_editable                = cc.DlFunc[func(textView *TextView) int32, int32]{Name: "gtk_text_view_get_editable"}
	gtk_text_view_set_overwrite               = cc.DlFunc[func(textView *TextView, overwrite int32), cc.Void]{Name: "gtk_text_view_set_overwrite"}
	gtk_text_view_get_overwrite               = cc.DlFunc[func(textView *TextView) int32, int32]{Name: "gtk_text_view_get_overwrite"}
	gtk_text_view_set_accepts_tab             = cc.DlFunc[func(textView *TextView, acceptsTab int32), cc.Void]{Name: "gtk_text_view_set_accepts_tab"}
	gtk_text_view_get_accepts_tab             = cc.DlFunc[func(textView *TextView) int32, int32]{Name: "gtk_text_view_get_accepts_tab"}
	gtk_text_view_set_pixels_above_lines      = cc.DlFunc[func(textView *TextView, pixelsAboveLines int32), cc.Void]{Name: "gtk_text_view_set_pixels_above_lines"}
	gtk_text_view_get_pixels_above_lines      = cc.DlFunc[func(textView *TextView) int32, int32]{Name: "gtk_text_view_get_pixels_above_lines"}
	gtk_text_view_set_pixels_below_lines      = cc.DlFunc[func(textView *TextView, pixelsBelowLines int32), cc.Void]{Name: "gtk_text_view_set_pixels_below_lines"}
	gtk_text_view_get_pixels_below_lines      = cc.DlFunc[func(textView *TextView) int32, int32]{Name: "gtk_text_view_get_pixels_below_lines"}
	gtk_text_view_set_pixels_inside_wrap      = cc.DlFunc[func(textView *TextView, pixelsInsideWrap int32), cc.Void]{Name: "gtk_text_view_set_pixels_inside_wrap"}
	gtk_text_view_get_pixels_inside_wrap      = cc.DlFunc[func(textView *TextView) int32, int32]{Name: "gtk_text_view_get_pixels_inside_wrap"}
	gtk_text_view_set_justification           = cc.DlFunc[func(textView *TextView, justification Justification), cc.Void]{Name: "gtk_text_view_set_justification"}
	gtk_text_view_get_justification           = cc.DlFunc[func(textView *TextView) Justification, Justification]{Name: "gtk_text_view_get_justification"}
	gtk_text_view_set_left_margin             = cc.DlFunc[func(textView *TextView, leftMargin int32), cc.Void]{Name: "gtk_text_view_set_left_margin"}
	gtk_text_view_get_left_margin             = cc.DlFunc[func(textView *TextView) int32, int32]{Name: "gtk_text_view_get_left_margin"}
	gtk_text_view_set_right_margin            = cc.DlFunc[func(textView *TextView, rightMargin int32), cc.Void]{Name: "gtk_text_view_set_right_margin"}
	gtk_text_view_get_right_margin            = cc.DlFunc[func(textView *TextView) int32, int32]{Name: "gtk_text_view_get_right_margin"}
	gtk_text_view_set_top_margin              = cc.DlFunc[func(textView *TextView, topMargin int32), cc.Void]{Name: "gtk_text_view_set_top_margin"}
	gtk_text_view_get_top_margin              = cc.DlFunc[func(textView *TextView) int32, int32]{Name: "gtk_text_view_get_top_margin"}
	gtk_text_view_set_bottom_margin           = cc.DlFunc[func(textView *TextView, bottomMargin int32), cc.Void]{Name: "gtk_text_view_set_bottom_margin"}
	gtk_text_view_get_bottom_margin           = cc.DlFunc[func(textView *TextView) int32, int32]{Name: "gtk_text_view_get_bottom_margin"}
	gtk_text_view_set_indent                  = cc.DlFunc[func(textView *TextView, indent int32), cc.Void]{Name: "gtk_text_view_set_indent"}
	gtk_text_view_get_indent                  = cc.DlFunc[func(textView *TextView) int32, int32]{Name: "gtk_text_view_get_indent"}
	gtk_text_view_set_tabs                    = cc.DlFunc[func(textView *TextView, tabs *pango.TabArray), cc.Void]{Name: "gtk_text_view_set_tabs"}
	gtk_text_view_get_tabs                    = cc.DlFunc[func(textView *TextView) *pango.TabArray, *pango.TabArray]{Name: "gtk_text_view_get_tabs"}
	gtk_text_view_set_input_purpose           = cc.DlFunc[func(textView *TextView, purpose InputPurpose), cc.Void]{Name: "gtk_text_view_set_input_purpose"}
	gtk_text_view_get_input_purpose           = cc.DlFunc[func(textView *TextView) InputPurpose, InputPurpose]{Name: "gtk_text_view_get_input_purpose"}
	gtk_text_view_set_input_hints             = cc.DlFunc[func(textView *TextView, hints InputHints), cc.Void]{Name: "gtk_text_view_set_input_hints"}
	gtk_text_view_get_input_hints             = cc.DlFunc[func(textView *TextView) InputHints, InputHints]{Name: "gtk_text_view_get_input_hints"}
	gtk_text_view_set_monospace               = cc.DlFunc[func(textView *TextView, monospace int32), cc.Void]{Name: "gtk_text_view_set_monospace"}
	gtk_text_view_get_monospace               = cc.DlFunc[func(textView *TextView) int32, int32]{Name: "gtk_text_view_get_monospace"}
	gtk_text_view_set_extra_menu              = cc.DlFunc[func(textView *TextView, model *gio.GMenuModel), cc.Void]{Name: "gtk_text_view_set_extra_menu"}
	gtk_text_view_get_extra_menu              = cc.DlFunc[func(textView *TextView) *gio.GMenuModel, *gio.GMenuModel]{Name: "gtk_text_view_get_extra_menu"}
	gtk_text_view_get_rtl_context             = cc.DlFunc[func(textView *TextView) *pango.Context, *pango.Context]{Name: "gtk_text_view_get_rtl_context"}
	gtk_text_view_get_ltr_context             = cc.DlFunc[func(textView *TextView) *pango.Context, *pango.Context]{Name: "gtk_text_view_get_ltr_context"}
	// #endregion

	// #region ToggleButton
	gtk_toggle_button_get_type          = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_toggle_button_get_type"}
	gtk_toggle_button_new               = cc.DlFunc[func() *ToggleButton, *ToggleButton]{Name: "gtk_toggle_button_new"}
	gtk_toggle_button_new_with_label    = cc.DlFunc[func(label cc.String) *ToggleButton, *ToggleButton]{Name: "gtk_toggle_button_new_with_label"}
	gtk_toggle_button_new_with_mnemonic = cc.DlFunc[func(label cc.String) *ToggleButton, *ToggleButton]{Name: "gtk_toggle_button_new_with_mnemonic"}
	gtk_toggle_button_set_active        = cc.DlFunc[func(tb *ToggleButton, isActive int32), cc.Void]{Name: "gtk_toggle_button_set_active"}
	gtk_toggle_button_get_active        = cc.DlFunc[func(tb *ToggleButton) int32, int32]{Name: "gtk_toggle_button_get_active"}
	gtk_toggle_button_toggled           = cc.DlFunc[func(tb *ToggleButton), cc.Void]{Name: "gtk_toggle_button_toggled"}
	gtk_toggle_button_set_group         = cc.DlFunc[func(tb *ToggleButton, group *ToggleButton), cc.Void]{Name: "gtk_toggle_button_set_group"}
	// #endregion

	// #region Tooltip
	gtk_tooltip_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_tooltip_get_type"}
	gtk_tooltip_set_markup              = cc.DlFunc[func(tt *Tooltip, markup cc.String), cc.Void]{Name: "gtk_tooltip_set_markup"}
	gtk_tooltip_set_text                = cc.DlFunc[func(tt *Tooltip, text cc.String), cc.Void]{Name: "gtk_tooltip_set_text"}
	gtk_tooltip_set_icon                = cc.DlFunc[func(tt *Tooltip, paintable *gdk.Paintable), cc.Void]{Name: "gtk_tooltip_set_icon"}
	gtk_tooltip_set_icon_from_icon_name = cc.DlFunc[func(tt *Tooltip, iconName cc.String), cc.Void]{Name: "gtk_tooltip_set_icon_from_icon_name"}
	gtk_tooltip_set_icon_from_gicon     = cc.DlFunc[func(tt *Tooltip, gicon *gio.GIcon), cc.Void]{Name: "gtk_tooltip_set_icon_from_gicon"}
	gtk_tooltip_set_custom              = cc.DlFunc[func(tt *Tooltip, customWidget *Widget), cc.Void]{Name: "gtk_tooltip_set_custom"}
	gtk_tooltip_set_tip_area            = cc.DlFunc[func(tt *Tooltip, rect *gdk.Rectangle), cc.Void]{Name: "gtk_tooltip_set_tip_area"}
	// #endregion

	// #region TreeExpander
	gtk_tree_expander_get_type             = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_tree_expander_get_type"}
	gtk_tree_expander_new                  = cc.DlFunc[func() *TreeExpander, *TreeExpander]{Name: "gtk_tree_expander_new"}
	gtk_tree_expander_get_child            = cc.DlFunc[func(te *TreeExpander) *Widget, *Widget]{Name: "gtk_tree_expander_get_child"}
	gtk_tree_expander_set_child            = cc.DlFunc[func(te *TreeExpander, child *Widget), cc.Void]{Name: "gtk_tree_expander_set_child"}
	gtk_tree_expander_get_item             = cc.DlFunc[func(te *TreeExpander) uptr, uptr]{Name: "gtk_tree_expander_get_item"}
	gtk_tree_expander_get_list_row         = cc.DlFunc[func(te *TreeExpander) *TreeListRow, *TreeListRow]{Name: "gtk_tree_expander_get_list_row"}
	gtk_tree_expander_set_list_row         = cc.DlFunc[func(te *TreeExpander, listRow *TreeListRow), cc.Void]{Name: "gtk_tree_expander_set_list_row"}
	gtk_tree_expander_get_indent_for_depth = cc.DlFunc[func(te *TreeExpander) int32, int32]{Name: "gtk_tree_expander_get_indent_for_depth"}
	gtk_tree_expander_set_indent_for_depth = cc.DlFunc[func(te *TreeExpander, indentForDepth int32), cc.Void]{Name: "gtk_tree_expander_set_indent_for_depth"}
	gtk_tree_expander_get_indent_for_icon  = cc.DlFunc[func(te *TreeExpander) int32, int32]{Name: "gtk_tree_expander_get_indent_for_icon"}
	gtk_tree_expander_set_indent_for_icon  = cc.DlFunc[func(te *TreeExpander, indentForIcon int32), cc.Void]{Name: "gtk_tree_expander_set_indent_for_icon"}
	gtk_tree_expander_get_hide_expander    = cc.DlFunc[func(te *TreeExpander) int32, int32]{Name: "gtk_tree_expander_get_hide_expander"}
	gtk_tree_expander_set_hide_expander    = cc.DlFunc[func(te *TreeExpander, hideExpander int32), cc.Void]{Name: "gtk_tree_expander_set_hide_expander"}
	// #endregion

	// #region TreeListModel
	gtk_tree_list_model_get_type        = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_tree_list_model_get_type"}
	gtk_tree_list_model_new             = cc.DlFunc[func(root *gio.GListModel, passthrough int32, autoexpand int32, createFunc uptr, userData uptr, userDestroy uptr) *TreeListModel, *TreeListModel]{Name: "gtk_tree_list_model_new"}
	gtk_tree_list_model_get_model       = cc.DlFunc[func(tlm *TreeListModel) *gio.GListModel, *gio.GListModel]{Name: "gtk_tree_list_model_get_model"}
	gtk_tree_list_model_get_passthrough = cc.DlFunc[func(tlm *TreeListModel) int32, int32]{Name: "gtk_tree_list_model_get_passthrough"}
	gtk_tree_list_model_set_autoexpand  = cc.DlFunc[func(tlm *TreeListModel, autoexpand int32), cc.Void]{Name: "gtk_tree_list_model_set_autoexpand"}
	gtk_tree_list_model_get_autoexpand  = cc.DlFunc[func(tlm *TreeListModel) int32, int32]{Name: "gtk_tree_list_model_get_autoexpand"}
	gtk_tree_list_model_get_child_row   = cc.DlFunc[func(tlm *TreeListModel, position uint32) *TreeListRow, *TreeListRow]{Name: "gtk_tree_list_model_get_child_row"}
	gtk_tree_list_model_get_row         = cc.DlFunc[func(tlm *TreeListModel, position uint32) *TreeListRow, *TreeListRow]{Name: "gtk_tree_list_model_get_row"}

	gtk_tree_list_row_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_tree_list_row_get_type"}
	gtk_tree_list_row_get_item      = cc.DlFunc[func(row *TreeListRow) uptr, uptr]{Name: "gtk_tree_list_row_get_item"}
	gtk_tree_list_row_set_expanded  = cc.DlFunc[func(row *TreeListRow, expanded int32), cc.Void]{Name: "gtk_tree_list_row_set_expanded"}
	gtk_tree_list_row_get_expanded  = cc.DlFunc[func(row *TreeListRow) int32, int32]{Name: "gtk_tree_list_row_get_expanded"}
	gtk_tree_list_row_is_expandable = cc.DlFunc[func(row *TreeListRow) int32, int32]{Name: "gtk_tree_list_row_is_expandable"}
	gtk_tree_list_row_get_position  = cc.DlFunc[func(row *TreeListRow) uint32, uint32]{Name: "gtk_tree_list_row_get_position"}
	gtk_tree_list_row_get_depth     = cc.DlFunc[func(row *TreeListRow) uint32, uint32]{Name: "gtk_tree_list_row_get_depth"}
	gtk_tree_list_row_get_children  = cc.DlFunc[func(row *TreeListRow) *gio.GListModel, *gio.GListModel]{Name: "gtk_tree_list_row_get_children"}
	gtk_tree_list_row_get_parent    = cc.DlFunc[func(row *TreeListRow) *TreeListRow, *TreeListRow]{Name: "gtk_tree_list_row_get_parent"}
	gtk_tree_list_row_get_child_row = cc.DlFunc[func(row *TreeListRow, position uint32) *TreeListRow, *TreeListRow]{Name: "gtk_tree_list_row_get_child_row"}
	// #endregion

	// #region TreeListRowSorter
	gtk_tree_list_row_sorter_get_type   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_tree_list_row_sorter_get_type"}
	gtk_tree_list_row_sorter_new        = cc.DlFunc[func(sorter *Sorter) *TreeListRowSorter, *TreeListRowSorter]{Name: "gtk_tree_list_row_sorter_new"}
	gtk_tree_list_row_sorter_get_sorter = cc.DlFunc[func(sorter *TreeListRowSorter) *Sorter, *Sorter]{Name: "gtk_tree_list_row_sorter_get_sorter"}
	gtk_tree_list_row_sorter_set_sorter = cc.DlFunc[func(sorter *TreeListRowSorter, s *Sorter), cc.Void]{Name: "gtk_tree_list_row_sorter_set_sorter"}
	// #endregion

	// #region UriLauncher
	gtk_uri_launcher_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_uri_launcher_get_type"}
	gtk_uri_launcher_new           = cc.DlFunc[func(uri cc.String) *UriLauncher, *UriLauncher]{Name: "gtk_uri_launcher_new"}
	gtk_uri_launcher_get_uri       = cc.DlFunc[func(ul *UriLauncher) cc.String, cc.String]{Name: "gtk_uri_launcher_get_uri"}
	gtk_uri_launcher_set_uri       = cc.DlFunc[func(ul *UriLauncher, uri cc.String), cc.Void]{Name: "gtk_uri_launcher_set_uri"}
	gtk_uri_launcher_launch        = cc.DlFunc[func(ul *UriLauncher, parent *Window, cancellable *gio.GCancellable, callback uptr, userData uptr), cc.Void]{Name: "gtk_uri_launcher_launch"}
	gtk_uri_launcher_launch_finish = cc.DlFunc[func(ul *UriLauncher, result *gio.GAsyncResult, err **glib.GError) int32, int32]{Name: "gtk_uri_launcher_launch_finish"}
	// #endregion

	// #region Version
	gtk_get_major_version = cc.DlFunc[func() uint32, uint32]{Name: "gtk_get_major_version"}
	gtk_get_minor_version = cc.DlFunc[func() uint32, uint32]{Name: "gtk_get_minor_version"}
	gtk_get_micro_version = cc.DlFunc[func() uint32, uint32]{Name: "gtk_get_micro_version"}
	gtk_get_binary_age    = cc.DlFunc[func() uint32, uint32]{Name: "gtk_get_binary_age"}
	gtk_get_interface_age = cc.DlFunc[func() uint32, uint32]{Name: "gtk_get_interface_age"}
	gtk_check_version     = cc.DlFunc[func(requiredMajor, requiredMinor, requiredMicro uint32) cc.String, cc.String]{Name: "gtk_check_version"}
	// #endregion

	// #region Video
	gtk_video_get_type             = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_video_get_type"}
	gtk_video_new                  = cc.DlFunc[func() *Video, *Video]{Name: "gtk_video_new"}
	gtk_video_new_for_media_stream = cc.DlFunc[func(stream *MediaStream) *Video, *Video]{Name: "gtk_video_new_for_media_stream"}
	gtk_video_new_for_file         = cc.DlFunc[func(file *gio.GFile) *Video, *Video]{Name: "gtk_video_new_for_file"}
	gtk_video_new_for_filename     = cc.DlFunc[func(filename cc.String) *Video, *Video]{Name: "gtk_video_new_for_filename"}
	gtk_video_new_for_resource     = cc.DlFunc[func(resourcePath cc.String) *Video, *Video]{Name: "gtk_video_new_for_resource"}
	gtk_video_get_media_stream     = cc.DlFunc[func(v *Video) *MediaStream, *MediaStream]{Name: "gtk_video_get_media_stream"}
	gtk_video_set_media_stream     = cc.DlFunc[func(v *Video, stream *MediaStream), cc.Void]{Name: "gtk_video_set_media_stream"}
	gtk_video_get_file             = cc.DlFunc[func(v *Video) *gio.GFile, *gio.GFile]{Name: "gtk_video_get_file"}
	gtk_video_set_file             = cc.DlFunc[func(v *Video, file *gio.GFile), cc.Void]{Name: "gtk_video_set_file"}
	gtk_video_set_filename         = cc.DlFunc[func(v *Video, filename cc.String), cc.Void]{Name: "gtk_video_set_filename"}
	gtk_video_set_resource         = cc.DlFunc[func(v *Video, resourcePath cc.String), cc.Void]{Name: "gtk_video_set_resource"}
	gtk_video_get_autoplay         = cc.DlFunc[func(v *Video) int32, int32]{Name: "gtk_video_get_autoplay"}
	gtk_video_set_autoplay         = cc.DlFunc[func(v *Video, autoplay int32), cc.Void]{Name: "gtk_video_set_autoplay"}
	gtk_video_get_loop             = cc.DlFunc[func(v *Video) int32, int32]{Name: "gtk_video_get_loop"}
	gtk_video_set_loop             = cc.DlFunc[func(v *Video, loop int32), cc.Void]{Name: "gtk_video_set_loop"}
	gtk_video_get_graphics_offload = cc.DlFunc[func(v *Video) GraphicsOffloadEnabled, GraphicsOffloadEnabled]{Name: "gtk_video_get_graphics_offload"}
	gtk_video_set_graphics_offload = cc.DlFunc[func(v *Video, enabled GraphicsOffloadEnabled), cc.Void]{Name: "gtk_video_set_graphics_offload"}
	// #endregion

	// #region Viewport
	gtk_viewport_get_type            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_viewport_get_type"}
	gtk_viewport_new                 = cc.DlFunc[func(hadj, vadj *Adjustment) *Viewport, *Viewport]{Name: "gtk_viewport_new"}
	gtk_viewport_get_scroll_to_focus = cc.DlFunc[func(vp *Viewport) int32, int32]{Name: "gtk_viewport_get_scroll_to_focus"}
	gtk_viewport_set_scroll_to_focus = cc.DlFunc[func(vp *Viewport, scrollToFocus int32), cc.Void]{Name: "gtk_viewport_set_scroll_to_focus"}
	gtk_viewport_set_child           = cc.DlFunc[func(vp *Viewport, child *Widget), cc.Void]{Name: "gtk_viewport_set_child"}
	gtk_viewport_get_child           = cc.DlFunc[func(vp *Viewport) *Widget, *Widget]{Name: "gtk_viewport_get_child"}
	gtk_viewport_scroll_to           = cc.DlFunc[func(vp *Viewport, descendant *Widget, scroll *ScrollInfo), cc.Void]{Name: "gtk_viewport_scroll_to"}
	// #endregion

	// #region Widget
	gtk_widget_get_type                            = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_widget_get_type"}
	gtk_widget_unparent                            = cc.DlFunc[func(*Widget), cc.Void]{Name: "gtk_widget_unparent"}
	gtk_widget_show                                = cc.DlFunc[func(*Widget), cc.Void]{Name: "gtk_widget_show"}
	gtk_widget_hide                                = cc.DlFunc[func(*Widget), cc.Void]{Name: "gtk_widget_hide"}
	gtk_widget_map                                 = cc.DlFunc[func(*Widget), cc.Void]{Name: "gtk_widget_map"}
	gtk_widget_unmap                               = cc.DlFunc[func(*Widget), cc.Void]{Name: "gtk_widget_unmap"}
	gtk_widget_realize                             = cc.DlFunc[func(*Widget), cc.Void]{Name: "gtk_widget_realize"}
	gtk_widget_unrealize                           = cc.DlFunc[func(*Widget), cc.Void]{Name: "gtk_widget_unrealize"}
	gtk_widget_queue_draw                          = cc.DlFunc[func(*Widget), cc.Void]{Name: "gtk_widget_queue_draw"}
	gtk_widget_queue_resize                        = cc.DlFunc[func(*Widget), cc.Void]{Name: "gtk_widget_queue_resize"}
	gtk_widget_queue_allocate                      = cc.DlFunc[func(*Widget), cc.Void]{Name: "gtk_widget_queue_allocate"}
	gtk_widget_get_frame_clock                     = cc.DlFunc[func(*Widget) *gdk.FrameClock, *gdk.FrameClock]{Name: "gtk_widget_get_frame_clock"}
	gtk_widget_size_allocate                       = cc.DlFunc[func(*Widget, *Allocation, int), cc.Void]{Name: "gtk_widget_size_allocate"}
	gtk_widget_allocate                            = cc.DlFunc[func(*Widget, int, int, int, *gsk.Transform), cc.Void]{Name: "gtk_widget_allocate"}
	gtk_widget_get_request_mode                    = cc.DlFunc[func(*Widget) SizeRequestMode, SizeRequestMode]{Name: "gtk_widget_get_request_mode"}
	gtk_widget_measure                             = cc.DlFunc[func(*Widget, Orientation, int, *int, *int, *int, *int), cc.Void]{Name: "gtk_widget_measure"}
	gtk_widget_get_preferred_size                  = cc.DlFunc[func(*Widget, *Requisition, *Requisition), cc.Void]{Name: "gtk_widget_get_preferred_size"}
	gtk_widget_set_layout_manager                  = cc.DlFunc[func(*Widget, *LayoutManager), cc.Void]{Name: "gtk_widget_set_layout_manager"}
	gtk_widget_get_layout_manager                  = cc.DlFunc[func(*Widget) *LayoutManager, *LayoutManager]{Name: "gtk_widget_get_layout_manager"}
	gtk_widget_class_set_layout_manager_type       = cc.DlFunc[func(*WidgetClass, gobject.GType), cc.Void]{Name: "gtk_widget_class_set_layout_manager_type"}
	gtk_widget_class_get_layout_manager_type       = cc.DlFunc[func(*WidgetClass) gobject.GType, gobject.GType]{Name: "gtk_widget_class_get_layout_manager_type"}
	gtk_widget_class_add_binding                   = cc.DlFunc[func(*WidgetClass, uint, gdk.ModifierType, uptr, cc.String, ...interface{}), cc.Void]{Name: "gtk_widget_class_add_binding", Va: true}
	gtk_widget_class_add_binding_signal            = cc.DlFunc[func(*WidgetClass, uint, gdk.ModifierType, cc.String, cc.String, ...interface{}), cc.Void]{Name: "gtk_widget_class_add_binding_signal", Va: true}
	gtk_widget_class_add_binding_action            = cc.DlFunc[func(*WidgetClass, uint, gdk.ModifierType, cc.String, cc.String, ...interface{}), cc.Void]{Name: "gtk_widget_class_add_binding_action", Va: true}
	gtk_widget_class_add_shortcut                  = cc.DlFunc[func(*WidgetClass, *Shortcut), cc.Void]{Name: "gtk_widget_class_add_shortcut"}
	gtk_widget_class_set_activate_signal           = cc.DlFunc[func(*WidgetClass, uint), cc.Void]{Name: "gtk_widget_class_set_activate_signal"}
	gtk_widget_class_set_activate_signal_from_name = cc.DlFunc[func(*WidgetClass, cc.String), cc.Void]{Name: "gtk_widget_class_set_activate_signal_from_name"}
	gtk_widget_class_get_activate_signal           = cc.DlFunc[func(*WidgetClass) uint, uint]{Name: "gtk_widget_class_get_activate_signal"}
	gtk_widget_mnemonic_activate                   = cc.DlFunc[func(*Widget, int32) int32, int32]{Name: "gtk_widget_mnemonic_activate"}
	gtk_widget_activate                            = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_activate"}
	gtk_widget_set_can_focus                       = cc.DlFunc[func(*Widget, int32), cc.Void]{Name: "gtk_widget_set_can_focus"}
	gtk_widget_get_can_focus                       = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_get_can_focus"}
	gtk_widget_set_focusable                       = cc.DlFunc[func(*Widget, int32), cc.Void]{Name: "gtk_widget_set_focusable"}
	gtk_widget_get_focusable                       = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_get_focusable"}
	gtk_widget_has_focus                           = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_has_focus"}
	gtk_widget_is_focus                            = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_is_focus"}
	gtk_widget_has_visible_focus                   = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_has_visible_focus"}
	gtk_widget_grab_focus                          = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_grab_focus"}
	gtk_widget_set_focus_on_click                  = cc.DlFunc[func(*Widget, int32), cc.Void]{Name: "gtk_widget_set_focus_on_click"}
	gtk_widget_get_focus_on_click                  = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_get_focus_on_click"}
	gtk_widget_set_can_target                      = cc.DlFunc[func(*Widget, int32), cc.Void]{Name: "gtk_widget_set_can_target"}
	gtk_widget_get_can_target                      = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_get_can_target"}
	gtk_widget_has_default                         = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_has_default"}
	gtk_widget_set_receives_default                = cc.DlFunc[func(*Widget, int32), cc.Void]{Name: "gtk_widget_set_receives_default"}
	gtk_widget_get_receives_default                = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_get_receives_default"}
	gtk_widget_set_name                            = cc.DlFunc[func(*Widget, cc.String), cc.Void]{Name: "gtk_widget_set_name"}
	gtk_widget_get_name                            = cc.DlFunc[func(*Widget) cc.String, cc.String]{Name: "gtk_widget_get_name"}
	gtk_widget_set_state_flags                     = cc.DlFunc[func(*Widget, StateFlags, int32), cc.Void]{Name: "gtk_widget_set_state_flags"}
	gtk_widget_unset_state_flags                   = cc.DlFunc[func(*Widget, StateFlags), cc.Void]{Name: "gtk_widget_unset_state_flags"}
	gtk_widget_get_state_flags                     = cc.DlFunc[func(*Widget) StateFlags, StateFlags]{Name: "gtk_widget_get_state_flags"}
	gtk_widget_set_sensitive                       = cc.DlFunc[func(*Widget, int32), cc.Void]{Name: "gtk_widget_set_sensitive"}
	gtk_widget_get_sensitive                       = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_get_sensitive"}
	gtk_widget_is_sensitive                        = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_is_sensitive"}
	gtk_widget_set_visible                         = cc.DlFunc[func(*Widget, int32), cc.Void]{Name: "gtk_widget_set_visible"}
	gtk_widget_get_visible                         = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_get_visible"}
	gtk_widget_is_visible                          = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_is_visible"}
	gtk_widget_is_drawable                         = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_is_drawable"}
	gtk_widget_get_realized                        = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_get_realized"}
	gtk_widget_get_mapped                          = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_get_mapped"}
	gtk_widget_set_parent                          = cc.DlFunc[func(*Widget, *Widget), cc.Void]{Name: "gtk_widget_set_parent"}
	gtk_widget_get_parent                          = cc.DlFunc[func(*Widget) *Widget, *Widget]{Name: "gtk_widget_get_parent"}
	gtk_widget_get_root                            = cc.DlFunc[func(*Widget) *Root, *Root]{Name: "gtk_widget_get_root"}
	gtk_widget_get_native                          = cc.DlFunc[func(*Widget) *Native, *Native]{Name: "gtk_widget_get_native"}
	gtk_widget_set_child_visible                   = cc.DlFunc[func(*Widget, int32), cc.Void]{Name: "gtk_widget_set_child_visible"}
	gtk_widget_get_child_visible                   = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_get_child_visible"}
	gtk_widget_get_allocated_width                 = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_get_allocated_width"}
	gtk_widget_get_allocated_height                = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_get_allocated_height"}
	gtk_widget_get_allocated_baseline              = cc.DlFunc[func(*Widget) int32, int32]{Name: "gtk_widget_get_allocated_baseline"}
	gtk_widget_get_allocation                      = cc.DlFunc[func(*Widget, *Allocation), cc.Void]{Name: "gtk_widget_get_allocation"}
	gtk_widget_compute_transform                   = cc.DlFunc[func(widget *Widget, target *Widget, outTransform *graphene.Matrix) int32, int32]{Name: "gtk_widget_compute_transform"}
	gtk_widget_compute_bounds                      = cc.DlFunc[func(widget *Widget, target *Widget, outBounds *graphene.Rect) int32, int32]{Name: "gtk_widget_compute_bounds"}
	gtk_widget_compute_point                       = cc.DlFunc[func(widget *Widget, target *Widget, point *graphene.Point, outPoint *graphene.Point) int32, int32]{Name: "gtk_widget_compute_point"}
	gtk_widget_get_width                           = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_get_width"}
	gtk_widget_get_height                          = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_get_height"}
	gtk_widget_get_baseline                        = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_get_baseline"}
	gtk_widget_get_size                            = cc.DlFunc[func(widget *Widget, orientation Orientation) int32, int32]{Name: "gtk_widget_get_size"}
	gtk_widget_child_focus                         = cc.DlFunc[func(widget *Widget, direction DirectionType) int32, int32]{Name: "gtk_widget_child_focus"}
	gtk_widget_keynav_failed                       = cc.DlFunc[func(widget *Widget, direction DirectionType) int32, int32]{Name: "gtk_widget_keynav_failed"}
	gtk_widget_error_bell                          = cc.DlFunc[func(widget *Widget), cc.Void]{Name: "gtk_widget_error_bell"}
	gtk_widget_set_size_request                    = cc.DlFunc[func(widget *Widget, width int32, height int32), cc.Void]{Name: "gtk_widget_set_size_request"}
	gtk_widget_get_size_request                    = cc.DlFunc[func(widget *Widget, width *int32, height *int32), cc.Void]{Name: "gtk_widget_get_size_request"}
	gtk_widget_set_opacity                         = cc.DlFunc[func(widget *Widget, opacity float64), cc.Void]{Name: "gtk_widget_set_opacity"}
	gtk_widget_get_opacity                         = cc.DlFunc[func(widget *Widget) float64, float64]{Name: "gtk_widget_get_opacity"}
	gtk_widget_set_overflow                        = cc.DlFunc[func(widget *Widget, overflow Overflow), cc.Void]{Name: "gtk_widget_set_overflow"}
	gtk_widget_get_overflow                        = cc.DlFunc[func(widget *Widget) Overflow, Overflow]{Name: "gtk_widget_get_overflow"}
	gtk_widget_get_ancestor                        = cc.DlFunc[func(widget *Widget, widgetType gobject.GType) *Widget, *Widget]{Name: "gtk_widget_get_ancestor"}
	gtk_widget_get_scale_factor                    = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_get_scale_factor"}
	gtk_widget_get_display                         = cc.DlFunc[func(widget *Widget) *gdk.Display, *gdk.Display]{Name: "gtk_widget_get_display"}
	gtk_widget_get_settings                        = cc.DlFunc[func(widget *Widget) *Settings, *Settings]{Name: "gtk_widget_get_settings"}
	gtk_widget_get_clipboard                       = cc.DlFunc[func(widget *Widget) *gdk.Clipboard, *gdk.Clipboard]{Name: "gtk_widget_get_clipboard"}
	gtk_widget_get_primary_clipboard               = cc.DlFunc[func(widget *Widget) *gdk.Clipboard, *gdk.Clipboard]{Name: "gtk_widget_get_primary_clipboard"}
	gtk_widget_get_hexpand                         = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_get_hexpand"}
	gtk_widget_set_hexpand                         = cc.DlFunc[func(widget *Widget, expand int32), cc.Void]{Name: "gtk_widget_set_hexpand"}
	gtk_widget_get_hexpand_set                     = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_get_hexpand_set"}
	gtk_widget_set_hexpand_set                     = cc.DlFunc[func(widget *Widget, set int32), cc.Void]{Name: "gtk_widget_set_hexpand_set"}
	gtk_widget_get_vexpand                         = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_get_vexpand"}
	gtk_widget_set_vexpand                         = cc.DlFunc[func(widget *Widget, expand int32), cc.Void]{Name: "gtk_widget_set_vexpand"}
	gtk_widget_get_vexpand_set                     = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_get_vexpand_set"}
	gtk_widget_set_vexpand_set                     = cc.DlFunc[func(widget *Widget, set int32), cc.Void]{Name: "gtk_widget_set_vexpand_set"}
	gtk_widget_compute_expand                      = cc.DlFunc[func(widget *Widget, orientation Orientation) int32, int32]{Name: "gtk_widget_compute_expand"}
	gtk_widget_get_halign                          = cc.DlFunc[func(widget *Widget) Align, Align]{Name: "gtk_widget_get_halign"}
	gtk_widget_set_halign                          = cc.DlFunc[func(widget *Widget, align Align), cc.Void]{Name: "gtk_widget_set_halign"}
	gtk_widget_get_valign                          = cc.DlFunc[func(widget *Widget) Align, Align]{Name: "gtk_widget_get_valign"}
	gtk_widget_set_valign                          = cc.DlFunc[func(widget *Widget, align Align), cc.Void]{Name: "gtk_widget_set_valign"}
	gtk_widget_get_margin_start                    = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_get_margin_start"}
	gtk_widget_set_margin_start                    = cc.DlFunc[func(widget *Widget, margin int32), cc.Void]{Name: "gtk_widget_set_margin_start"}
	gtk_widget_get_margin_end                      = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_get_margin_end"}
	gtk_widget_set_margin_end                      = cc.DlFunc[func(widget *Widget, margin int32), cc.Void]{Name: "gtk_widget_set_margin_end"}
	gtk_widget_get_margin_top                      = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_get_margin_top"}
	gtk_widget_set_margin_top                      = cc.DlFunc[func(widget *Widget, margin int32), cc.Void]{Name: "gtk_widget_set_margin_top"}
	gtk_widget_get_margin_bottom                   = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_get_margin_bottom"}
	gtk_widget_set_margin_bottom                   = cc.DlFunc[func(widget *Widget, margin int32), cc.Void]{Name: "gtk_widget_set_margin_bottom"}
	gtk_widget_is_ancestor                         = cc.DlFunc[func(widget *Widget, ancestor *Widget) int32, int32]{Name: "gtk_widget_is_ancestor"}
	gtk_widget_translate_coordinates               = cc.DlFunc[func(srcWidget *Widget, destWidget *Widget, srcX, srcY float64, destX, destY *float64) int32, int32]{Name: "gtk_widget_translate_coordinates"}
	gtk_widget_contains                            = cc.DlFunc[func(widget *Widget, x, y float64) int32, int32]{Name: "gtk_widget_contains"}
	gtk_widget_pick                                = cc.DlFunc[func(widget *Widget, x, y float64, flags PickFlags) *Widget, *Widget]{Name: "gtk_widget_pick"}
	gtk_widget_add_controller                      = cc.DlFunc[func(widget *Widget, controller *EventController), cc.Void]{Name: "gtk_widget_add_controller"}
	gtk_widget_remove_controller                   = cc.DlFunc[func(widget *Widget, controller *EventController), cc.Void]{Name: "gtk_widget_remove_controller"}
	gtk_widget_create_pango_context                = cc.DlFunc[func(widget *Widget) *pango.Context, *pango.Context]{Name: "gtk_widget_create_pango_context"}
	gtk_widget_get_pango_context                   = cc.DlFunc[func(widget *Widget) *pango.Context, *pango.Context]{Name: "gtk_widget_get_pango_context"}
	gtk_widget_set_font_options                    = cc.DlFunc[func(widget *Widget, options *cairo.FontOptions), cc.Void]{Name: "gtk_widget_set_font_options"}
	gtk_widget_get_font_options                    = cc.DlFunc[func(widget *Widget) *cairo.FontOptions, *cairo.FontOptions]{Name: "gtk_widget_get_font_options"}
	gtk_widget_create_pango_layout                 = cc.DlFunc[func(widget *Widget, text cc.String) *pango.Layout, *pango.Layout]{Name: "gtk_widget_create_pango_layout"}
	gtk_widget_set_direction                       = cc.DlFunc[func(widget *Widget, dir TextDirection), cc.Void]{Name: "gtk_widget_set_direction"}
	gtk_widget_get_direction                       = cc.DlFunc[func(widget *Widget) TextDirection, TextDirection]{Name: "gtk_widget_get_direction"}
	gtk_widget_set_default_direction               = cc.DlFunc[func(dir TextDirection), cc.Void]{Name: "gtk_widget_set_default_direction"}
	gtk_widget_get_default_direction               = cc.DlFunc[func() TextDirection, TextDirection]{Name: "gtk_widget_get_default_direction"}
	gtk_widget_set_cursor                          = cc.DlFunc[func(widget *Widget, cursor *gdk.Cursor), cc.Void]{Name: "gtk_widget_set_cursor"}
	gtk_widget_set_cursor_from_name                = cc.DlFunc[func(widget *Widget, name cc.String), cc.Void]{Name: "gtk_widget_set_cursor_from_name"}
	gtk_widget_get_cursor                          = cc.DlFunc[func(widget *Widget) *gdk.Cursor, *gdk.Cursor]{Name: "gtk_widget_get_cursor"}
	gtk_widget_list_mnemonic_labels                = cc.DlFunc[func(widget *Widget) *glib.GList[Widget], *glib.GList[Widget]]{Name: "gtk_widget_list_mnemonic_labels"}
	gtk_widget_add_mnemonic_label                  = cc.DlFunc[func(widget *Widget, label *Widget), cc.Void]{Name: "gtk_widget_add_mnemonic_label"}
	gtk_widget_remove_mnemonic_label               = cc.DlFunc[func(widget *Widget, label *Widget), cc.Void]{Name: "gtk_widget_remove_mnemonic_label"}
	gtk_widget_trigger_tooltip_query               = cc.DlFunc[func(widget *Widget), cc.Void]{Name: "gtk_widget_trigger_tooltip_query"}
	gtk_widget_set_tooltip_text                    = cc.DlFunc[func(widget *Widget, text cc.String), cc.Void]{Name: "gtk_widget_set_tooltip_text"}
	gtk_widget_get_tooltip_text                    = cc.DlFunc[func(widget *Widget) cc.String, cc.String]{Name: "gtk_widget_get_tooltip_text"}
	gtk_widget_set_tooltip_markup                  = cc.DlFunc[func(widget *Widget, markup cc.String), cc.Void]{Name: "gtk_widget_set_tooltip_markup"}
	gtk_widget_get_tooltip_markup                  = cc.DlFunc[func(widget *Widget) cc.String, cc.String]{Name: "gtk_widget_get_tooltip_markup"}
	gtk_widget_set_has_tooltip                     = cc.DlFunc[func(widget *Widget, hasTooltip int32), cc.Void]{Name: "gtk_widget_set_has_tooltip"}
	gtk_widget_get_has_tooltip                     = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_get_has_tooltip"}
	gtk_requisition_get_type                       = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_requisition_get_type"}
	gtk_requisition_new                            = cc.DlFunc[func() *Requisition, *Requisition]{Name: "gtk_requisition_new"}
	gtk_requisition_copy                           = cc.DlFunc[func(requisition *Requisition) *Requisition, *Requisition]{Name: "gtk_requisition_copy"}
	gtk_requisition_free                           = cc.DlFunc[func(requisition *Requisition), cc.Void]{Name: "gtk_requisition_free"}
	gtk_widget_in_destruction                      = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_in_destruction"}
	gtk_widget_class_set_css_name                  = cc.DlFunc[func(widgetClass *WidgetClass, name cc.String), cc.Void]{Name: "gtk_widget_class_set_css_name"}
	gtk_widget_class_get_css_name                  = cc.DlFunc[func(widgetClass *WidgetClass) cc.String, cc.String]{Name: "gtk_widget_class_get_css_name"}
	gtk_widget_add_tick_callback                   = cc.DlFunc[func(widget *Widget, callback, userData, notify uptr) uint32, uint32]{Name: "gtk_widget_add_tick_callback"}
	gtk_widget_remove_tick_callback                = cc.DlFunc[func(widget *Widget, id uint32), cc.Void]{Name: "gtk_widget_remove_tick_callback"}
	gtk_widget_insert_action_group                 = cc.DlFunc[func(widget *Widget, name cc.String, group *gio.GActionGroup), cc.Void]{Name: "gtk_widget_insert_action_group"}
	gtk_widget_activate_action                     = cc.DlFunc[func(widget *Widget, name cc.String, formatString cc.String, args ...interface{}) int32, int32]{Name: "gtk_widget_activate_action", Va: true}
	gtk_widget_activate_action_variant             = cc.DlFunc[func(widget *Widget, name cc.String, variant *glib.GVariant) int32, int32]{Name: "gtk_widget_activate_action_variant"}
	gtk_widget_activate_default                    = cc.DlFunc[func(widget *Widget), cc.Void]{Name: "gtk_widget_activate_default"}
	gtk_widget_set_font_map                        = cc.DlFunc[func(widget *Widget, fontMap *pango.FontMap), cc.Void]{Name: "gtk_widget_set_font_map"}
	gtk_widget_get_font_map                        = cc.DlFunc[func(widget *Widget) *pango.FontMap, *pango.FontMap]{Name: "gtk_widget_get_font_map"}
	gtk_widget_get_first_child                     = cc.DlFunc[func(widget *Widget) *Widget, *Widget]{Name: "gtk_widget_get_first_child"}
	gtk_widget_get_last_child                      = cc.DlFunc[func(widget *Widget) *Widget, *Widget]{Name: "gtk_widget_get_last_child"}
	gtk_widget_get_next_sibling                    = cc.DlFunc[func(widget *Widget) *Widget, *Widget]{Name: "gtk_widget_get_next_sibling"}
	gtk_widget_get_prev_sibling                    = cc.DlFunc[func(widget *Widget) *Widget, *Widget]{Name: "gtk_widget_get_prev_sibling"}
	gtk_widget_observe_children                    = cc.DlFunc[func(widget *Widget) *gio.GListModel, *gio.GListModel]{Name: "gtk_widget_observe_children"}
	gtk_widget_observe_controllers                 = cc.DlFunc[func(widget *Widget) *gio.GListModel, *gio.GListModel]{Name: "gtk_widget_observe_controllers"}
	gtk_widget_insert_after                        = cc.DlFunc[func(widget *Widget, parent *Widget, previousSibling *Widget), cc.Void]{Name: "gtk_widget_insert_after"}
	gtk_widget_insert_before                       = cc.DlFunc[func(widget *Widget, parent *Widget, nextSibling *Widget), cc.Void]{Name: "gtk_widget_insert_before"}
	gtk_widget_set_focus_child                     = cc.DlFunc[func(widget *Widget, child *Widget), cc.Void]{Name: "gtk_widget_set_focus_child"}
	gtk_widget_get_focus_child                     = cc.DlFunc[func(widget *Widget) *Widget, *Widget]{Name: "gtk_widget_get_focus_child"}
	gtk_widget_snapshot_child                      = cc.DlFunc[func(widget *Widget, child *Widget, snapshot *Snapshot), cc.Void]{Name: "gtk_widget_snapshot_child"}
	gtk_widget_should_layout                       = cc.DlFunc[func(widget *Widget) int32, int32]{Name: "gtk_widget_should_layout"}
	gtk_widget_get_css_name                        = cc.DlFunc[func(widget *Widget) cc.String, cc.String]{Name: "gtk_widget_get_css_name"}
	gtk_widget_add_css_class                       = cc.DlFunc[func(widget *Widget, cssClass cc.String), cc.Void]{Name: "gtk_widget_add_css_class"}
	gtk_widget_remove_css_class                    = cc.DlFunc[func(widget *Widget, cssClass cc.String), cc.Void]{Name: "gtk_widget_remove_css_class"}
	gtk_widget_has_css_class                       = cc.DlFunc[func(widget *Widget, cssClass cc.String) int32, int32]{Name: "gtk_widget_has_css_class"}
	gtk_widget_get_css_classes                     = cc.DlFunc[func(widget *Widget) cc.Strings, cc.Strings]{Name: "gtk_widget_get_css_classes"}
	gtk_widget_set_css_classes                     = cc.DlFunc[func(widget *Widget, classes cc.Strings), cc.Void]{Name: "gtk_widget_set_css_classes"}
	gtk_widget_get_color                           = cc.DlFunc[func(widget *Widget, color *gdk.RGBA), cc.Void]{Name: "gtk_widget_get_color"}
	gtk_widget_class_install_action                = cc.DlFunc[func(widgetClass *WidgetClass, actionName, parameterType cc.String, activate uptr), cc.Void]{Name: "gtk_widget_class_install_action"}
	gtk_widget_class_install_property_action       = cc.DlFunc[func(widgetClass *WidgetClass, actionName, propertyName cc.String), cc.Void]{Name: "gtk_widget_class_install_property_action"}
	gtk_widget_class_query_action                  = cc.DlFunc[func(widgetClass *WidgetClass, index uint, owner *gobject.GType, actionName *cc.String, parameterType **glib.GVariantType, propertyName *cc.String) int32, int32]{Name: "gtk_widget_class_query_action"}
	gtk_widget_action_set_enabled                  = cc.DlFunc[func(widget *Widget, actionName cc.String, enabled int32), cc.Void]{Name: "gtk_widget_action_set_enabled"}
	gtk_widget_class_set_accessible_role           = cc.DlFunc[func(widgetClass *WidgetClass, accessibleRole AccessibleRole), cc.Void]{Name: "gtk_widget_class_set_accessible_role"}
	gtk_widget_class_get_accessible_role           = cc.DlFunc[func(widgetClass *WidgetClass) AccessibleRole, AccessibleRole]{Name: "gtk_widget_class_get_accessible_role"}
	// #endregion

	// #region WidgetPaintable
	gtk_widget_paintable_get_type   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_widget_paintable_get_type"}
	gtk_widget_paintable_new        = cc.DlFunc[func(widget *Widget) *WidgetPaintable, *WidgetPaintable]{Name: "gtk_widget_paintable_new"}
	gtk_widget_paintable_get_widget = cc.DlFunc[func(wp *WidgetPaintable) *Widget, *Widget]{Name: "gtk_widget_paintable_get_widget"}
	gtk_widget_paintable_set_widget = cc.DlFunc[func(wp *WidgetPaintable, widget *Widget), cc.Void]{Name: "gtk_widget_paintable_set_widget"}
	// #endregion

	// #region Window
	gtk_window_get_type                      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_window_get_type"}
	gtk_window_new                           = cc.DlFunc[func() *Window, *Window]{Name: "gtk_window_new"}
	gtk_window_set_title                     = cc.DlFunc[func(window *Window, title cc.String), cc.Void]{Name: "gtk_window_set_title"}
	gtk_window_get_title                     = cc.DlFunc[func(window *Window) cc.String, cc.String]{Name: "gtk_window_get_title"}
	gtk_window_set_startup_id                = cc.DlFunc[func(window *Window, startupID cc.String), cc.Void]{Name: "gtk_window_set_startup_id"}
	gtk_window_set_focus                     = cc.DlFunc[func(window *Window, focus *Widget), cc.Void]{Name: "gtk_window_set_focus"}
	gtk_window_get_focus                     = cc.DlFunc[func(window *Window) *Widget, *Widget]{Name: "gtk_window_get_focus"}
	gtk_window_set_default_widget            = cc.DlFunc[func(window *Window, defaultWidget *Widget), cc.Void]{Name: "gtk_window_set_default_widget"}
	gtk_window_get_default_widget            = cc.DlFunc[func(window *Window) *Widget, *Widget]{Name: "gtk_window_get_default_widget"}
	gtk_window_set_transient_for             = cc.DlFunc[func(window *Window, parent *Window), cc.Void]{Name: "gtk_window_set_transient_for"}
	gtk_window_get_transient_for             = cc.DlFunc[func(window *Window) *Window, *Window]{Name: "gtk_window_get_transient_for"}
	gtk_window_set_destroy_with_parent       = cc.DlFunc[func(window *Window, setting int32), cc.Void]{Name: "gtk_window_set_destroy_with_parent"}
	gtk_window_get_destroy_with_parent       = cc.DlFunc[func(window *Window) int32, int32]{Name: "gtk_window_get_destroy_with_parent"}
	gtk_window_set_hide_on_close             = cc.DlFunc[func(window *Window, setting int32), cc.Void]{Name: "gtk_window_set_hide_on_close"}
	gtk_window_get_hide_on_close             = cc.DlFunc[func(window *Window) int32, int32]{Name: "gtk_window_get_hide_on_close"}
	gtk_window_set_mnemonics_visible         = cc.DlFunc[func(window *Window, setting int32), cc.Void]{Name: "gtk_window_set_mnemonics_visible"}
	gtk_window_get_mnemonics_visible         = cc.DlFunc[func(window *Window) int32, int32]{Name: "gtk_window_get_mnemonics_visible"}
	gtk_window_set_focus_visible             = cc.DlFunc[func(window *Window, setting int32), cc.Void]{Name: "gtk_window_set_focus_visible"}
	gtk_window_get_focus_visible             = cc.DlFunc[func(window *Window) int32, int32]{Name: "gtk_window_get_focus_visible"}
	gtk_window_set_resizable                 = cc.DlFunc[func(window *Window, resizable int32), cc.Void]{Name: "gtk_window_set_resizable"}
	gtk_window_get_resizable                 = cc.DlFunc[func(window *Window) int32, int32]{Name: "gtk_window_get_resizable"}
	gtk_window_set_display                   = cc.DlFunc[func(window *Window, display *gdk.Display), cc.Void]{Name: "gtk_window_set_display"}
	gtk_window_is_active                     = cc.DlFunc[func(window *Window) int32, int32]{Name: "gtk_window_is_active"}
	gtk_window_set_decorated                 = cc.DlFunc[func(window *Window, setting int32), cc.Void]{Name: "gtk_window_set_decorated"}
	gtk_window_get_decorated                 = cc.DlFunc[func(window *Window) int32, int32]{Name: "gtk_window_get_decorated"}
	gtk_window_set_deletable                 = cc.DlFunc[func(window *Window, setting int32), cc.Void]{Name: "gtk_window_set_deletable"}
	gtk_window_get_deletable                 = cc.DlFunc[func(window *Window) int32, int32]{Name: "gtk_window_get_deletable"}
	gtk_window_set_icon_name                 = cc.DlFunc[func(window *Window, name cc.String), cc.Void]{Name: "gtk_window_set_icon_name"}
	gtk_window_get_icon_name                 = cc.DlFunc[func(window *Window) cc.String, cc.String]{Name: "gtk_window_get_icon_name"}
	gtk_window_set_default_icon_name         = cc.DlFunc[func(name cc.String), cc.Void]{Name: "gtk_window_set_default_icon_name"}
	gtk_window_get_default_icon_name         = cc.DlFunc[func() cc.String, cc.String]{Name: "gtk_window_get_default_icon_name"}
	gtk_window_set_auto_startup_notification = cc.DlFunc[func(setting int32), cc.Void]{Name: "gtk_window_set_auto_startup_notification"}
	gtk_window_set_modal                     = cc.DlFunc[func(window *Window, modal int32), cc.Void]{Name: "gtk_window_set_modal"}
	gtk_window_get_modal                     = cc.DlFunc[func(window *Window) int32, int32]{Name: "gtk_window_get_modal"}
	gtk_window_get_toplevels                 = cc.DlFunc[func() *gio.GListModel, *gio.GListModel]{Name: "gtk_window_get_toplevels"}
	gtk_window_list_toplevels                = cc.DlFunc[func() *glib.GList[Widget], *glib.GList[Widget]]{Name: "gtk_window_list_toplevels"}
	gtk_window_present                       = cc.DlFunc[func(window *Window), cc.Void]{Name: "gtk_window_present"}
	gtk_window_minimize                      = cc.DlFunc[func(window *Window), cc.Void]{Name: "gtk_window_minimize"}
	gtk_window_unminimize                    = cc.DlFunc[func(window *Window), cc.Void]{Name: "gtk_window_unminimize"}
	gtk_window_maximize                      = cc.DlFunc[func(window *Window), cc.Void]{Name: "gtk_window_maximize"}
	gtk_window_unmaximize                    = cc.DlFunc[func(window *Window), cc.Void]{Name: "gtk_window_unmaximize"}
	gtk_window_fullscreen                    = cc.DlFunc[func(window *Window), cc.Void]{Name: "gtk_window_fullscreen"}
	gtk_window_unfullscreen                  = cc.DlFunc[func(window *Window), cc.Void]{Name: "gtk_window_unfullscreen"}
	gtk_window_fullscreen_on_monitor         = cc.DlFunc[func(window *Window, monitor *gdk.Monitor), cc.Void]{Name: "gtk_window_fullscreen_on_monitor"}
	gtk_window_close                         = cc.DlFunc[func(window *Window), cc.Void]{Name: "gtk_window_close"}
	gtk_window_set_default_size              = cc.DlFunc[func(window *Window, width int32, height int32), cc.Void]{Name: "gtk_window_set_default_size"}
	gtk_window_get_default_size              = cc.DlFunc[func(window *Window, width *int32, height *int32), cc.Void]{Name: "gtk_window_get_default_size"}
	gtk_window_get_group                     = cc.DlFunc[func(window *Window) *WindowGroup, *WindowGroup]{Name: "gtk_window_get_group"}
	gtk_window_has_group                     = cc.DlFunc[func(window *Window) int32, int32]{Name: "gtk_window_has_group"}
	gtk_window_get_application               = cc.DlFunc[func(window *Window) *Application, *Application]{Name: "gtk_window_get_application"}
	gtk_window_set_application               = cc.DlFunc[func(window *Window, application *Application), cc.Void]{Name: "gtk_window_set_application"}
	gtk_window_set_child                     = cc.DlFunc[func(window *Window, child *Widget), cc.Void]{Name: "gtk_window_set_child"}
	gtk_window_get_child                     = cc.DlFunc[func(window *Window) *Widget, *Widget]{Name: "gtk_window_get_child"}
	gtk_window_set_titlebar                  = cc.DlFunc[func(window *Window, titlebar *Widget), cc.Void]{Name: "gtk_window_set_titlebar"}
	gtk_window_get_titlebar                  = cc.DlFunc[func(window *Window) *Widget, *Widget]{Name: "gtk_window_get_titlebar"}
	gtk_window_is_maximized                  = cc.DlFunc[func(window *Window) int32, int32]{Name: "gtk_window_is_maximized"}
	gtk_window_is_fullscreen                 = cc.DlFunc[func(window *Window) int32, int32]{Name: "gtk_window_is_fullscreen"}
	gtk_window_is_suspended                  = cc.DlFunc[func(window *Window) int32, int32]{Name: "gtk_window_is_suspended"}
	gtk_window_destroy                       = cc.DlFunc[func(window *Window), cc.Void]{Name: "gtk_window_destroy"}
	gtk_window_set_interactive_debugging     = cc.DlFunc[func(enable int32), cc.Void]{Name: "gtk_window_set_interactive_debugging"}
	gtk_window_set_handle_menubar_accel      = cc.DlFunc[func(window *Window, handleMenubarAccel int32), cc.Void]{Name: "gtk_window_set_handle_menubar_accel"}
	gtk_window_get_handle_menubar_accel      = cc.DlFunc[func(window *Window) int32, int32]{Name: "gtk_window_get_handle_menubar_accel"}
	// #endregion

	// #region WindowControls
	gtk_window_controls_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_window_controls_get_type"}
	gtk_window_controls_new                     = cc.DlFunc[func(side PackType) *WindowControls, *WindowControls]{Name: "gtk_window_controls_new"}
	gtk_window_controls_get_side                = cc.DlFunc[func(wc *WindowControls) PackType, PackType]{Name: "gtk_window_controls_get_side"}
	gtk_window_controls_set_side                = cc.DlFunc[func(wc *WindowControls, side PackType), cc.Void]{Name: "gtk_window_controls_set_side"}
	gtk_window_controls_get_decoration_layout   = cc.DlFunc[func(wc *WindowControls) cc.String, cc.String]{Name: "gtk_window_controls_get_decoration_layout"}
	gtk_window_controls_set_decoration_layout   = cc.DlFunc[func(wc *WindowControls, layout cc.String), cc.Void]{Name: "gtk_window_controls_set_decoration_layout"}
	gtk_window_controls_get_use_native_controls = cc.DlFunc[func(wc *WindowControls) int32, int32]{Name: "gtk_window_controls_get_use_native_controls"}
	gtk_window_controls_set_use_native_controls = cc.DlFunc[func(wc *WindowControls, setting int32), cc.Void]{Name: "gtk_window_controls_set_use_native_controls"}
	gtk_window_controls_get_empty               = cc.DlFunc[func(wc *WindowControls) int32, int32]{Name: "gtk_window_controls_get_empty"}
	// #endregion

	// #region WindowGroup
	gtk_window_group_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_window_group_get_type"}
	gtk_window_group_new           = cc.DlFunc[func() *WindowGroup, *WindowGroup]{Name: "gtk_window_group_new"}
	gtk_window_group_add_window    = cc.DlFunc[func(windowGroup *WindowGroup, window *Window), cc.Void]{Name: "gtk_window_group_add_window"}
	gtk_window_group_remove_window = cc.DlFunc[func(windowGroup *WindowGroup, window *Window), cc.Void]{Name: "gtk_window_group_remove_window"}
	gtk_window_group_list_windows  = cc.DlFunc[func(windowGroup *WindowGroup) *glib.GList[Window], *glib.GList[Window]]{Name: "gtk_window_group_list_windows"}
	// #endregion

	// #region WindowHandle
	gtk_window_handle_get_type  = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gtk_window_handle_get_type"}
	gtk_window_handle_new       = cc.DlFunc[func() *WindowHandle, *WindowHandle]{Name: "gtk_window_handle_new"}
	gtk_window_handle_get_child = cc.DlFunc[func(wh *WindowHandle) *Widget, *Widget]{Name: "gtk_window_handle_get_child"}
	gtk_window_handle_set_child = cc.DlFunc[func(wh *WindowHandle, child *Widget), cc.Void]{Name: "gtk_window_handle_set_child"}
	// #endregion

)
