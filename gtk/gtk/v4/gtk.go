package gtk

import (
	"errors"
	"runtime"
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

func cbool(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

// #region AboutDialog

type AboutDialog struct {
	Window

	Root
	ShortcutManager
}

func GTypeAboutDialog() gobject.GType { return gtk_about_dialog_get_type.Fn()() }
func NewAboutDialog() *AboutDialog    { return gtk_about_dialog_new.Fn()() }
func ShowAboutDialog(parent WindowIface, firstPropertyName string, args ...interface{}) {
	cs := cc.NewString(firstPropertyName)
	defer cs.Free()
	gtk_show_about_dialog.FnVa()(GetWindowIface(parent), cs, args...)
}
func (a *AboutDialog) GetProgramName() string {
	return gtk_about_dialog_get_program_name.Fn()(a).String()
}
func (a *AboutDialog) SetProgramName(name string) {
	cs := cc.NewString(name)
	defer cs.Free()
	gtk_about_dialog_set_program_name.Fn()(a, cs)
}
func (a *AboutDialog) GetVersion() string { return gtk_about_dialog_get_version.Fn()(a).String() }
func (a *AboutDialog) SetVersion(version string) {
	cs := cc.NewString(version)
	defer cs.Free()
	gtk_about_dialog_set_version.Fn()(a, cs)
}
func (a *AboutDialog) GetCopyright() string { return gtk_about_dialog_get_copyright.Fn()(a).String() }
func (a *AboutDialog) SetCopyright(copyright string) {
	cs := cc.NewString(copyright)
	defer cs.Free()
	gtk_about_dialog_set_copyright.Fn()(a, cs)
}
func (a *AboutDialog) GetComments() string { return gtk_about_dialog_get_comments.Fn()(a).String() }
func (a *AboutDialog) SetComments(comments string) {
	cs := cc.NewString(comments)
	defer cs.Free()
	gtk_about_dialog_set_comments.Fn()(a, cs)
}
func (a *AboutDialog) GetLicense() string { return gtk_about_dialog_get_license.Fn()(a).String() }
func (a *AboutDialog) SetLicense(license string) {
	cs := cc.NewString(license)
	defer cs.Free()
	gtk_about_dialog_set_license.Fn()(a, cs)
}
func (a *AboutDialog) SetLicenseType(licenseType License) {
	gtk_about_dialog_set_license_type.Fn()(a, licenseType)
}
func (a *AboutDialog) GetLicenseType() License { return gtk_about_dialog_get_license_type.Fn()(a) }
func (a *AboutDialog) GetWrapLicense() bool    { return gtk_about_dialog_get_wrap_license.Fn()(a) != 0 }
func (a *AboutDialog) SetWrapLicense(wrapLicense bool) {
	gtk_about_dialog_set_wrap_license.Fn()(a, cbool(wrapLicense))
}
func (a *AboutDialog) GetSystemInformation() string {
	return gtk_about_dialog_get_system_information.Fn()(a).String()
}
func (a *AboutDialog) SetSystemInformation(systemInformation string) {
	cs := cc.NewString(systemInformation)
	defer cs.Free()
	gtk_about_dialog_set_system_information.Fn()(a, cs)
}
func (a *AboutDialog) GetWebsite() string { return gtk_about_dialog_get_website.Fn()(a).String() }
func (a *AboutDialog) SetWebsite(website string) {
	cs := cc.NewString(website)
	defer cs.Free()
	gtk_about_dialog_set_website.Fn()(a, cs)
}
func (a *AboutDialog) GetWebsiteLabel() string {
	return gtk_about_dialog_get_website_label.Fn()(a).String()
}
func (a *AboutDialog) SetWebsiteLabel(websiteLabel string) {
	cs := cc.NewString(websiteLabel)
	defer cs.Free()
	gtk_about_dialog_set_website_label.Fn()(a, cs)
}
func (a *AboutDialog) GetAuthors() []string { return gtk_about_dialog_get_authors.Fn()(a).Strings() }
func (a *AboutDialog) SetAuthors(authors []string) {
	cpp := cc.NewStrings(authors)
	defer cpp.Free()
	gtk_about_dialog_set_authors.Fn()(a, cpp)
}
func (a *AboutDialog) GetDocumenters() []string {
	return gtk_about_dialog_get_documenters.Fn()(a).Strings()
}
func (a *AboutDialog) SetDocumenters(documenters []string) {
	cpp := cc.NewStrings(documenters)
	defer cpp.Free()
	gtk_about_dialog_set_documenters.Fn()(a, cpp)
}
func (a *AboutDialog) GetArtists() []string { return gtk_about_dialog_get_artists.Fn()(a).Strings() }
func (a *AboutDialog) SetArtists(artists []string) {
	cpp := cc.NewStrings(artists)
	defer cpp.Free()
	gtk_about_dialog_set_artists.Fn()(a, cpp)
}
func (a *AboutDialog) GetTranslatorCredits() string {
	return gtk_about_dialog_get_translator_credits.Fn()(a).String()
}
func (a *AboutDialog) SetTranslatorCredits(translatorCredits string) {
	cs := cc.NewString(translatorCredits)
	defer cs.Free()
	gtk_about_dialog_set_translator_credits.Fn()(a, cs)
}
func (a *AboutDialog) GetLogo() *gdk.Paintable     { return gtk_about_dialog_get_logo.Fn()(a) }
func (a *AboutDialog) SetLogo(logo *gdk.Paintable) { gtk_about_dialog_set_logo.Fn()(a, logo) }
func (a *AboutDialog) GetLogoIconName() string {
	return gtk_about_dialog_get_logo_icon_name.Fn()(a).String()
}
func (a *AboutDialog) SetLogoIconName(iconName string) {
	cs := cc.NewString(iconName)
	defer cs.Free()
	gtk_about_dialog_set_logo_icon_name.Fn()(a, cs)
}
func (a *AboutDialog) AddCreditSection(sectionName string, people []string) {
	cs, cpp := cc.NewString(sectionName), cc.NewStrings(people)
	defer cpp.Free()
	defer cs.Free()
	gtk_about_dialog_add_credit_section.Fn()(a, cs, cpp)
}
func (a *AboutDialog) ConnectActivateLink(sig func(a *AboutDialog, link string, _ uptr) bool) uint64 {
	return a.SignalConnect("activate-link", sig, nil)
}

// #endregion

// #region Accelerator

func AcceleratorValid(keyval uint32, modifiers gdk.ModifierType) bool {
	return gtk_accelerator_valid.Fn()(keyval, modifiers) != 0
}
func AcceleratorParse(accelerator string) (key uint32, mods gdk.ModifierType, ok bool) {
	acl := cc.NewString(accelerator)
	defer acl.Free()
	ok = gtk_accelerator_parse.Fn()(acl, &key, &mods) != 0
	return
}

//	func AcceleratorParseWithKeycode(accelerator string, display *gdk.Display) (key uint32, codes []uint32, mods gdk.ModifierType, ok bool) {
//		var codesPtr *uint32
//		ok = gtk_accelerator_parse_with_keycode(accelerator, display, &key, &codesPtr, &mods)
//		if codesPtr != nil {
//			codes = cc.Uint32PtrToSlice(codesPtr)
//		}
//		return
//	}

func AcceleratorName(key uint32, mods gdk.ModifierType) string {
	return gtk_accelerator_name.Fn()(key, mods).TakeString()
}
func AcceleratorNameWithKeycode(display *gdk.Display, key uint32, keycode uint32, mods gdk.ModifierType) string {
	return gtk_accelerator_name_with_keycode.Fn()(display, key, keycode, mods).TakeString()
}
func AcceleratorGetLabel(key uint32, mods gdk.ModifierType) string {
	return gtk_accelerator_get_label.Fn()(key, mods).TakeString()
}
func AcceleratorGetLabelWithKeycode(display *gdk.Display, key uint32, keycode uint32, mods gdk.ModifierType) string {
	return gtk_accelerator_get_label_with_keycode.Fn()(display, key, keycode, mods).TakeString()
}
func AcceleratorGetDefaultModMask() gdk.ModifierType {
	return gtk_accelerator_get_default_mod_mask.Fn()()
}

// #endregion

// #region Accessible

type AccessibleIfaceObj struct {
	GIface gobject.GTypeInterface

	GetAtContext             cc.Func //GtkATContext * (* get_at_context) (GtkAccessible *self);
	GetPlatformState         cc.Func //gboolean (* get_platform_state) (GtkAccessible  *self,GtkAccessiblePlatformState  state);
	GetAccessibleParent      cc.Func //GtkAccessible * (* get_accessible_parent) (GtkAccessible *self);
	GetFirstAccessibleChild  cc.Func //GtkAccessible * (* get_first_accessible_child) (GtkAccessible *self);
	GetNextAccessibleSibling cc.Func //GtkAccessible * (* get_next_accessible_sibling) (GtkAccessible *self);
	GetBounds                cc.Func // gboolean (* get_bounds) (GtkAccessible *self,int *x,int *y,int *width,int *height);
}

type AccessibleIface interface {
	GetAccessibleIface() *Accessible
}

func GetAccessibleIface(a AccessibleIface) *Accessible {
	if anyptr(a) == nil {
		return nil
	}
	return a.GetAccessibleIface()
}
func (a *Accessible) GetAccessibleIface() *Accessible { return a }

type Accessible struct{}

func (a *Accessible) GetATContext() *ATContext { return gtk_accessible_get_at_context.Fn()(a) }
func (a *Accessible) GetPlatformState(state AccessiblePlatformState) bool {
	return gtk_accessible_get_platform_state.Fn()(a, state) != 0
}
func (a *Accessible) GetAccessibleParent() *Accessible {
	return gtk_accessible_get_accessible_parent.Fn()(a)
}
func (a *Accessible) SetAccessibleParent(parent, nextSibling AccessibleIface) {
	gtk_accessible_set_accessible_parent.Fn()(a, GetAccessibleIface(parent), GetAccessibleIface(parent))
}
func (a *Accessible) GetFirstAccessibleChild() *Accessible {
	return gtk_accessible_get_first_accessible_child.Fn()(a)
}
func (a *Accessible) GetNextAccessibleSibling() *Accessible {
	return gtk_accessible_get_next_accessible_sibling.Fn()(a)
}
func (a *Accessible) UpdateNextAccessibleSibling(newSibling AccessibleIface) {
	gtk_accessible_update_next_accessible_sibling.Fn()(a, GetAccessibleIface(newSibling))
}
func (a *Accessible) GetBounds() (x, y, width, height int32, ok bool) {
	ok = gtk_accessible_get_bounds.Fn()(a, &x, &y, &width, &height) != 0
	return
}
func (a *Accessible) GetAccessibleRole() AccessibleRole {
	return gtk_accessible_get_accessible_role.Fn()(a)
}
func (a *Accessible) UpdateState(firstState AccessibleState, args ...interface{}) {
	gtk_accessible_update_state.FnVa()(a, firstState, args...)
}

func (a *Accessible) UpdateProperty(firstProperty AccessibleProperty, args ...interface{}) {
	gtk_accessible_update_property.FnVa()(a, firstProperty, args...)
}
func (a *Accessible) UpdateRelation(firstRelation AccessibleRelation, args ...interface{}) {
	gtk_accessible_update_relation.FnVa()(a, firstRelation, args...)
}
func (a *Accessible) UpdateStateValue(states []AccessibleState, values []gobject.GValue) {
	gtk_accessible_update_state_value.Fn()(a, int32(len(states)), carry(states), carry(values))
}

func (a *Accessible) UpdatePropertyValue(properties []AccessibleProperty, values []gobject.GValue) {
	gtk_accessible_update_property_value.Fn()(a, int32(len(properties)), carry(properties), carry(values))
}

func (a *Accessible) UpdateRelationValue(relations []AccessibleRelation, values []gobject.GValue) {
	gtk_accessible_update_relation_value.Fn()(a, int32(len(relations)), carry(relations), carry(values))
}
func (a *Accessible) ResetState(state AccessibleState) {
	gtk_accessible_reset_state.Fn()(a, state)
}
func (a *Accessible) ResetProperty(property AccessibleProperty) {
	gtk_accessible_reset_property.Fn()(a, property)
}
func (a *Accessible) ResetRelation(relation AccessibleRelation) {
	gtk_accessible_reset_relation.Fn()(a, relation)
}
func (a *Accessible) StateInitValue(state AccessibleState, value *gobject.GValue) {
	gtk_accessible_state_init_value.Fn()(state, value)
}
func (a *Accessible) PropertyInitValue(property AccessibleProperty, value *gobject.GValue) {
	gtk_accessible_property_init_value.Fn()(property, value)
}
func (a *Accessible) RelationInitValue(relation AccessibleRelation, value *gobject.GValue) {
	gtk_accessible_relation_init_value.Fn()(relation, value)
}

type AccessibleList struct{}

func GTypeAccessibleList() gobject.GType { return gtk_accessible_list_get_type.Fn()() }
func (a *AccessibleList) GetObjects() *glib.GList[Accessible] {
	return gtk_accessible_list_get_objects.Fn()(a)
}
func NewAccessibleListFromList(list []AccessibleIface) *AccessibleList {
	lst := glib.NewGList[Accessible]()
	defer lst.Free()
	for _, a := range list {
		lst = lst.Append(GetAccessibleIface(a))
	}
	return gtk_accessible_list_new_from_list.Fn()(lst)
}
func NewAccessibleListFromArray(accessibles []AccessibleIface) *AccessibleList {
	p := (**Accessible)(nil)
	n := len(accessibles)
	if n > 0 {
		lst := make([]*Accessible, n)
		for i := 0; i < n; i++ {
			lst[i] = GetAccessibleIface(accessibles[i])
		}
		p = &lst[0]
	}
	return gtk_accessible_list_new_from_array.Fn()(p, uint64(len(accessibles)))
}
func (a *Accessible) Announce(message string, priority AccessibleAnnouncementPriority) {
	msg := cc.NewString(message)
	defer msg.Free()
	gtk_accessible_announce.Fn()(a, msg, priority)
}
func (a *Accessible) UpdatePlatformState(state AccessiblePlatformState) {
	gtk_accessible_update_platform_state.Fn()(a, state)
}

// #endregion

// #region AccessibleRange

type AccessibleRangeIfaceObj struct {
	GIface          gobject.GTypeInterface
	SetCurrentValue cc.Func //gboolean (* set_current_value) (GtkAccessibleRange *self,double value);
}

type AccessibleRangeIface interface {
	GetAccessibleRangeIface() *AccessibleRange
}

func GetAccessibleRangeIface(iface AccessibleRangeIface) *AccessibleRange {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetAccessibleRangeIface()
}

func (a *AccessibleRange) GetAccessibleRangeIface() *AccessibleRange {
	return a.GetAccessibleRangeIface()
}

type AccessibleRange struct{ Accessible }

func GTypeAccessibleRange() gobject.GType { return gtk_accessible_range_get_type.Fn()() }
func (a *AccessibleRange) SetCurrentValue(value float64) bool {
	return gtk_accessible_range_set_current_value.Fn()(a, value) != 0
}

// #endregion

// #region AccessibleTextRange

type AccessibleTextRange struct {
	Start  uint64
	Length uint64
}

type AccessibleTextIfaceObj struct {
	GIface gobject.GTypeInterface

	GetContents          cc.Func // GBytes * (* get_contents) (GtkAccessibleText *self, unsigned int start, unsigned int end);
	GetContentsAt        cc.Func // GBytes * (* get_contents_at) (GtkAccessibleText *self, unsigned int offset, GtkAccessibleTextGranularity granularity, unsigned int *start, unsigned int *end);
	GetCaretPosition     cc.Func // unsigned int (* get_caret_position) (GtkAccessibleText *self);
	GetSelection         cc.Func // gboolean (* get_selection) (GtkAccessibleText *self, gsize *n_ranges, GtkAccessibleTextRange **ranges);
	GetAttributes        cc.Func // gboolean (* get_attributes) (GtkAccessibleText *self, unsigned int offset, gsize *n_ranges, GtkAccessibleTextRange **ranges, char ***attribute_names, char ***attribute_values);
	GetDefaultAttributes cc.Func // void (* get_default_attributes) (GtkAccessibleText *self, char ***attribute_names, char ***attribute_values);
	GetExtents           cc.Func //gboolean (* get_extents) (GtkAccessibleText *self, unsigned int start, unsigned int end, graphene_rect_t *extents);
	GetOffset            cc.Func // gboolean (* get_offset) (GtkAccessibleText *self, const graphene_point_t *point, unsigned int *offset);
}

type AccessibleTextIface interface {
	GetAccessibleTextIface() *AccessibleText
}

func GetAccessibleTextIface(iface AccessibleTextIface) *AccessibleText {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetAccessibleTextIface()
}

func (a *AccessibleText) GetAccessibleTextIface() *AccessibleText { return a }

type AccessibleText struct{}

func GTypeAccessibleText() gobject.GType        { return gtk_accessible_text_get_type.Fn()() }
func (a *AccessibleText) UpdateCaretPosition()  { gtk_accessible_text_update_caret_position.Fn()(a) }
func (a *AccessibleText) UpdateSelectionBound() { gtk_accessible_text_update_selection_bound.Fn()(a) }
func (a *AccessibleText) UpdateContents(change AccessibleTextContentChange, start, end uint32) {
	gtk_accessible_text_update_contents.Fn()(a, change, start, end)
}
func (a *AccessibleText) GetContents(start, end uint32) *glib.GBytes {
	return gtk_accessible_text_get_contents.Fn()(a, start, end)
}
func (a *AccessibleText) GetContentsAt(offset uint32, granularity AccessibleTextGranularity) (*glib.GBytes, uint32, uint32) {
	var start, end uint32
	bytes := gtk_accessible_text_get_contents_at.Fn()(a, offset, granularity, &start, &end)
	return bytes, start, end
}
func (a *AccessibleText) GetCaretPosition() uint32 {
	return gtk_accessible_text_get_caret_position.Fn()(a)
}
func (a *AccessibleText) GetSelection() ([]AccessibleTextRange, bool) {
	var nRanges uint64
	var ranges *AccessibleTextRange
	ok := gtk_accessible_text_get_selection.Fn()(a, &nRanges, &ranges) != 0
	if !ok || nRanges == 0 {
		return nil, ok
	}
	return (*[1 << 30]AccessibleTextRange)(unsafe.Pointer(ranges))[:nRanges:nRanges], ok
}
func (a *AccessibleText) GetAttributes(offset uint32) ([]AccessibleTextRange, []string, []string, bool) {
	var nRanges uint64
	var ranges *AccessibleTextRange
	var names, values cc.Strings
	ok := gtk_accessible_text_get_attributes.Fn()(a, offset, &nRanges, &ranges, &names, &values) != 0
	if !ok || nRanges == 0 {
		return nil, nil, nil, ok
	}
	return (*[1 << 30]AccessibleTextRange)(unsafe.Pointer(ranges))[:nRanges],
		names.Strings(), values.Strings(), ok
}
func (a *AccessibleText) GetDefaultAttributes() ([]string, []string) {
	var names, values cc.Strings
	gtk_accessible_text_get_default_attributes.Fn()(a, &names, &values)
	return names.Strings(), values.Strings()
}
func (a *AccessibleText) GetExtents(start, end uint32) (*graphene.Rect, bool) {
	var rect graphene.Rect
	ok := gtk_accessible_text_get_extents.Fn()(a, start, end, &rect) != 0
	return &rect, ok
}
func (a *AccessibleText) GetOffset(point *graphene.Point) (uint32, bool) {
	var offset uint32
	ok := gtk_accessible_text_get_offset.Fn()(a, point, &offset) != 0
	return offset, ok
}

// #endregion

// #region Actionable

type ActionableIfaceObj struct {
	GIface gobject.GTypeInterface

	GetActionName        cc.Func // const char * (* get_action_name)             (GtkActionable *actionable);
	SetActionName        cc.Func // void (* set_action_name) (GtkActionable *actionable, const char *action_name);
	GetActionTargetValue cc.Func // GVariant * (* get_action_target_value) (GtkActionable *actionable);
	SetActionTargetValue cc.Func // void (* set_action_target_value) (GtkActionable *actionable, GVariant *target_value);
}

type Actionable struct{}

func GTypeActionable() gobject.GType        { return gtk_actionable_get_type.Fn()() }
func (a *Actionable) GetActionName() string { return gtk_actionable_get_action_name.Fn()(a).String() }
func (a *Actionable) SetActionName(name string) {
	cs := cc.NewString(name)
	defer cs.Free()
	gtk_actionable_set_action_name.Fn()(a, cs)
}
func (a *Actionable) GetActionTargetValue() *glib.GVariant {
	return gtk_actionable_get_action_target_value.Fn()(a)
}
func (a *Actionable) SetActionTargetValue(value *glib.GVariant) {
	gtk_actionable_set_action_target_value.Fn()(a, value)
}
func (a *Actionable) SetActionTarget(format string, args ...interface{}) {
	cs := cc.NewString(format)
	defer cs.Free()
	gtk_actionable_set_action_target.FnVa()(a, cs, args...)
}
func (a *Actionable) SetDetailedActionName(name string) {
	cs := cc.NewString(name)
	defer cs.Free()
	gtk_actionable_set_detailed_action_name.Fn()(a, cs)
}

// #endregion

// #region ActionBar

type ActionBar struct{ Widget }

func GTypeActionBar() gobject.GType           { return gtk_action_bar_get_type.Fn()() }
func NewActionBar() *ActionBar                { return gtk_action_bar_new.Fn()() }
func (a *ActionBar) GetCenterWidget() *Widget { return gtk_action_bar_get_center_widget.Fn()(a) }
func (a *ActionBar) SetCenterWidget(widget WidgetIface) {
	gtk_action_bar_set_center_widget.Fn()(a, GetWidgetIface(widget))
}
func (a *ActionBar) PackStart(child WidgetIface) {
	gtk_action_bar_pack_start.Fn()(a, GetWidgetIface(child))
}
func (a *ActionBar) PackEnd(child WidgetIface) {
	gtk_action_bar_pack_end.Fn()(a, GetWidgetIface(child))
}
func (a *ActionBar) Remove(child WidgetIface)  { gtk_action_bar_remove.Fn()(a, GetWidgetIface(child)) }
func (a *ActionBar) SetRevealed(revealed bool) { gtk_action_bar_set_revealed.Fn()(a, cbool(revealed)) }
func (a *ActionBar) GetRevealed() bool         { return gtk_action_bar_get_revealed.Fn()(a) != 0 }

// #endregion

// #region Adjustment

type AdjustmentObj struct {
	Parent gobject.GObjectObj
}
type AdjustmentClass struct {
	Parent gobject.GObjectClass

	Changed      cc.Func //void (* changed) (GtkAdjustment *adjustment);
	ValueChanged cc.Func //void (* value_changed) (GtkAdjustment *adjustment);
	_            [4]uptr //reserved
}

type Adjustment struct {
	gobjCore
}

func GTypeAdjustment() gobject.GType { return gtk_adjustment_get_type.Fn()() }
func NewAdjustment(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) *Adjustment {
	return gtk_adjustment_new.Fn()(value, lower, upper, stepIncrement, pageIncrement, pageSize)
}
func (a *Adjustment) ClampPage(lower, upper float64) { gtk_adjustment_clamp_page.Fn()(a, lower, upper) }
func (a *Adjustment) GetValue() float64              { return gtk_adjustment_get_value.Fn()(a) }
func (a *Adjustment) SetValue(value float64)         { gtk_adjustment_set_value.Fn()(a, value) }
func (a *Adjustment) GetLower() float64              { return gtk_adjustment_get_lower.Fn()(a) }
func (a *Adjustment) SetLower(lower float64)         { gtk_adjustment_set_lower.Fn()(a, lower) }
func (a *Adjustment) GetUpper() float64              { return gtk_adjustment_get_upper.Fn()(a) }
func (a *Adjustment) SetUpper(upper float64)         { gtk_adjustment_set_upper.Fn()(a, upper) }
func (a *Adjustment) GetStepIncrement() float64      { return gtk_adjustment_get_step_increment.Fn()(a) }
func (a *Adjustment) SetStepIncrement(step float64)  { gtk_adjustment_set_step_increment.Fn()(a, step) }
func (a *Adjustment) GetPageIncrement() float64      { return gtk_adjustment_get_page_increment.Fn()(a) }
func (a *Adjustment) SetPageIncrement(page float64)  { gtk_adjustment_set_page_increment.Fn()(a, page) }
func (a *Adjustment) GetPageSize() float64           { return gtk_adjustment_get_page_size.Fn()(a) }
func (a *Adjustment) SetPageSize(size float64)       { gtk_adjustment_set_page_size.Fn()(a, size) }
func (a *Adjustment) Configure(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) {
	gtk_adjustment_configure.Fn()(a, value, lower, upper, stepIncrement, pageIncrement, pageSize)
}
func (a *Adjustment) GetMinimumIncrement() float64 {
	return gtk_adjustment_get_minimum_increment.Fn()(a)
}
func (a *Adjustment) ConnectChanged(sig func(a *Adjustment)) uint64 {
	return a.SignalConnect("changed", sig, nil)
}
func (a *Adjustment) ConnectValueChanged(sig func(a *Adjustment)) uint64 {
	return a.SignalConnect("value-changed", sig, nil)
}

// #endregion

// #region AlertDialog

type AlertDialog struct{ gobjCore }

func GTypeAlertDialog() gobject.GType { return gtk_alert_dialog_get_type.Fn()() }
func NewAlertDialog(format string, args ...interface{}) *AlertDialog {
	cs := cc.NewString(format)
	defer cs.Free()
	return gtk_alert_dialog_new.FnVa()(cs, args...)
}
func (a *AlertDialog) GetModal() bool      { return gtk_alert_dialog_get_modal.Fn()(a) != 0 }
func (a *AlertDialog) SetModal(modal bool) { gtk_alert_dialog_set_modal.Fn()(a, cbool(modal)) }
func (a *AlertDialog) GetMessage() string  { return gtk_alert_dialog_get_message.Fn()(a).String() }
func (a *AlertDialog) SetMessage(msg string) {
	cs := cc.NewString(msg)
	defer cs.Free()
	gtk_alert_dialog_set_message.Fn()(a, cs)
}
func (a *AlertDialog) GetDetail() string { return gtk_alert_dialog_get_detail.Fn()(a).String() }
func (a *AlertDialog) SetDetail(detail string) {
	cs := cc.NewString(detail)
	defer cs.Free()
	gtk_alert_dialog_set_detail.Fn()(a, cs)
}
func (a *AlertDialog) GetButtons() []string { return gtk_alert_dialog_get_buttons.Fn()(a).Strings() }
func (a *AlertDialog) SetButtons(labels []string) {
	cpp := cc.NewStrings(labels)
	defer cpp.Free()
	gtk_alert_dialog_set_buttons.Fn()(a, cpp)
}
func (a *AlertDialog) GetCancelButton() int32 { return gtk_alert_dialog_get_cancel_button.Fn()(a) }
func (a *AlertDialog) SetCancelButton(button int32) {
	gtk_alert_dialog_set_cancel_button.Fn()(a, button)
}
func (a *AlertDialog) GetDefaultButton() int32 { return gtk_alert_dialog_get_default_button.Fn()(a) }
func (a *AlertDialog) SetDefaultButton(button int32) {
	gtk_alert_dialog_set_default_button.Fn()(a, button)
}
func (a *AlertDialog) Choose(parent WindowIface, cancellable *gio.GCancellable,
	callback func(dlg *AlertDialog, res *gio.GAsyncResult)) {
	cbk := cc.Cbk(callback)
	a.WeakRef(func(obj *gobject.GObject) { cc.CbkCloseLate(cbk) })
	gtk_alert_dialog_choose.Fn()(a, GetWindowIface(parent), cancellable, cbk, nil)
}
func (a *AlertDialog) ChooseFinish(result *gio.GAsyncResult) (button int32, err error) {
	var gerr *glib.GError
	return gtk_alert_dialog_choose_finish.Fn()(a, result, &gerr), gerr.TakeError()
}
func (a *AlertDialog) Show(parent WindowIface) { gtk_alert_dialog_show.Fn()(a, GetWindowIface(parent)) }

// #endregion

// #region Application

type ApplicationObj struct {
	Parent gio.GApplicationObj
}

type ApplicationClass struct {
	Parent gio.GApplicationClass

	WindowAdded   cc.Func // void (*window_added) (GtkApplication *application,GtkWindow *window);
	WindowRemoved cc.Func // void (*window_removed) (GtkApplication *application,GtkWindow *window);

	_ [8]uptr //reserved
}

type Application struct {
	gio.GApplication
	gio.GActionGroup
	gio.GActionMap
}

func GTypeApplication() gobject.GType { return gtk_application_get_type.Fn()() }
func NewApplication(applicationID string, flags gio.GApplicationFlags) *Application {
	cId := cc.NewString(applicationID)
	defer cId.Free()
	return gtk_application_new.Fn()(cId, flags)
}
func (a *Application) AddWindow(window WindowIface) {
	gtk_application_add_window.Fn()(a, GetWindowIface(window))
}
func (a *Application) RemoveWindow(window WindowIface) {
	gtk_application_remove_window.Fn()(a, GetWindowIface(window))
}
func (a *Application) GetWindows() *glib.GList[Window] {
	return gtk_application_get_windows.Fn()(a)
}
func (a *Application) GetMenubar() *gio.GMenuModel { return gtk_application_get_menubar.Fn()(a) }
func (a *Application) SetMenubar(menubar *gio.GMenuModel) {
	gtk_application_set_menubar.Fn()(a, menubar)
}

func (a *Application) Inhibit(window WindowIface, flags ApplicationInhibitFlags, reason string) uint32 {
	cs := cc.NewString(reason)
	defer cs.Free()
	return gtk_application_inhibit.Fn()(a, GetWindowIface(window), flags, cs)
}
func (a *Application) Uninhibit(cookie uint32) { gtk_application_uninhibit.Fn()(a, cookie) }
func (a *Application) GetWindowByID(id uint32) *Window {
	return gtk_application_get_window_by_id.Fn()(a, id)
}
func (a *Application) GetActiveWindow() *Window { return gtk_application_get_active_window.Fn()(a) }
func (a *Application) ListActionDescriptions() []string {
	descriptions := gtk_application_list_action_descriptions.Fn()(a)
	defer descriptions.Free()
	return descriptions.Strings()
}
func (a *Application) GetAccelsForAction(detailedActionName string) []string {
	cs := cc.NewString(detailedActionName)
	defer cs.Free()
	accels := gtk_application_get_accels_for_action.Fn()(a, cs)
	defer accels.Free()
	return accels.Strings()
}
func (a *Application) GetActionsForAccel(accel string) []string {
	cs := cc.NewString(accel)
	defer cs.Free()
	actions := gtk_application_get_actions_for_accel.Fn()(a, cs)
	defer actions.Free()
	return actions.Strings()
}

func (a *Application) SetAccelsForAction(detailedActionName string, accels []string) {
	cs := cc.NewString(detailedActionName)
	defer cs.Free()
	accelPtr := cc.NewStrings(accels)
	defer accelPtr.Free()
	gtk_application_set_accels_for_action.Fn()(a, cs, accelPtr)
}
func (a *Application) GetMenuByID(id string) *gio.GMenu {
	cs := cc.NewString(id)
	defer cs.Free()
	return gtk_application_get_menu_by_id.Fn()(a, cs)
}
func (a *Application) ConnectQueryEnd(sig func(a *Application)) uint64 {
	return a.SignalConnect("query-end", sig, nil)
}
func (a *Application) ConnectWindowAdded(sig func(a *Application, w *Window)) uint64 {
	return a.SignalConnect("window-added", sig, nil)
}
func (a *Application) ConnectWindowRemoved(sig func(a *Application)) uint64 {
	return a.SignalConnect("window-removed", sig, nil)
}

// #endregion

// #region ApplicationWindow

type ApplicationWindowObj struct {
	Parent WindowObj
}
type ApplicationWindowClass struct {
	Parent WindowClass
	_      [8]uptr //reserved
}

type ApplicationWindow struct {
	Window

	gio.GActionGroup
	gio.GActionMap
}

func GTypeApplicationWindow() gobject.GType { return gtk_application_window_get_type.Fn()() }
func NewApplicationWindow(app *Application) *ApplicationWindow {
	return gtk_application_window_new.Fn()(app)
}
func (w *ApplicationWindow) SetShowMenubar(show bool) {
	gtk_application_window_set_show_menubar.Fn()(w, cbool(show))
}
func (w *ApplicationWindow) GetShowMenubar() bool {
	return gtk_application_window_get_show_menubar.Fn()(w) != 0
}
func (w *ApplicationWindow) GetID() uint32 { return gtk_application_window_get_id.Fn()(w) }

// #endregion

// #region AspectFrame

type AspectFrame struct{ Widget }

func GTypeAspectFrame() gobject.GType { return gtk_aspect_frame_get_type.Fn()() }
func NewAspectFrame(xalign, yalign, ratio float32, obeyChild bool) *AspectFrame {
	return gtk_aspect_frame_new.Fn()(xalign, yalign, ratio, cbool(obeyChild))
}
func (f *AspectFrame) SetXalign(xalign float32) { gtk_aspect_frame_set_xalign.Fn()(f, xalign) }
func (f *AspectFrame) GetXalign() float32       { return gtk_aspect_frame_get_xalign.Fn()(f) }
func (f *AspectFrame) SetYalign(yalign float32) { gtk_aspect_frame_set_yalign.Fn()(f, yalign) }
func (f *AspectFrame) GetYalign() float32       { return gtk_aspect_frame_get_yalign.Fn()(f) }
func (f *AspectFrame) SetRatio(ratio float32)   { gtk_aspect_frame_set_ratio.Fn()(f, ratio) }
func (f *AspectFrame) GetRatio() float32        { return gtk_aspect_frame_get_ratio.Fn()(f) }
func (f *AspectFrame) SetObeyChild(obeyChild bool) {
	gtk_aspect_frame_set_obey_child.Fn()(f, cbool(obeyChild))
}
func (f *AspectFrame) GetObeyChild() bool { return gtk_aspect_frame_get_obey_child.Fn()(f) != 0 }
func (f *AspectFrame) SetChild(child WidgetIface) {
	gtk_aspect_frame_set_child.Fn()(f, GetWidgetIface(child))
}
func (f *AspectFrame) GetChild() *Widget { return gtk_aspect_frame_get_child.Fn()(f) }

// #endregion

// #region ATContext

type ATContext struct{ gobjCore }

func (a *ATContext) GetAccessible() *Accessible { return gtk_at_context_get_accessible.Fn()(a) }
func (a *ATContext) GetAccessibleRole() AccessibleRole {
	return gtk_at_context_get_accessible_role.Fn()(a)
}
func CreateATContext(accessibleRole AccessibleRole, accessible AccessibleIface, display *gdk.Display) *ATContext {
	return gtk_at_context_create.Fn()(accessibleRole, GetAccessibleIface(accessible), display)
}
func (a *ATContext) ConnectStateChange(sig func(a *ATContext)) uint64 {
	return a.SignalConnect("state-change", sig, nil)
}

// #endregion

// #region BinLayout

type BinLayout struct {
	gobjCore
	LayoutManager
}

func GTypeBinLayout() gobject.GType { return gtk_bin_layout_get_type.Fn()() }
func NewBinLayout() *BinLayout      { return gtk_bin_layout_new.Fn()() }

// #endregion

// #region Bitset

type Bitset struct{}

func GTypeBitset() gobject.GType             { return gtk_bitset_get_type.Fn()() }
func (b *Bitset) Ref() *Bitset               { return gtk_bitset_ref.Fn()(b) }
func (b *Bitset) Unref()                     { gtk_bitset_unref.Fn()(b) }
func (b *Bitset) Contains(value uint32) bool { return gtk_bitset_contains.Fn()(b, value) != 0 }
func (b *Bitset) IsEmpty() bool              { return gtk_bitset_is_empty.Fn()(b) != 0 }
func (b *Bitset) Equals(other *Bitset) bool  { return gtk_bitset_equals.Fn()(b, other) != 0 }
func (b *Bitset) GetSize() uint64            { return gtk_bitset_get_size.Fn()(b) }
func (b *Bitset) GetSizeInRange(first, last uint32) uint64 {
	return gtk_bitset_get_size_in_range.Fn()(b, first, last)
}
func (b *Bitset) GetNth(nth uint32) uint32          { return gtk_bitset_get_nth.Fn()(b, nth) }
func (b *Bitset) GetMinimum() uint32                { return gtk_bitset_get_minimum.Fn()(b) }
func (b *Bitset) GetMaximum() uint32                { return gtk_bitset_get_maximum.Fn()(b) }
func NewBitsetEmpty() *Bitset                       { return gtk_bitset_new_empty.Fn()() }
func (b *Bitset) Copy() *Bitset                     { return gtk_bitset_copy.Fn()(b) }
func NewBitsetRange(start, nItems uint32) *Bitset   { return gtk_bitset_new_range.Fn()(start, nItems) }
func (b *Bitset) RemoveAll()                        { gtk_bitset_remove_all.Fn()(b) }
func (b *Bitset) Add(value uint32) bool             { return gtk_bitset_add.Fn()(b, value) != 0 }
func (b *Bitset) Remove(value uint32) bool          { return gtk_bitset_remove.Fn()(b, value) != 0 }
func (b *Bitset) AddRange(start, nItems uint32)     { gtk_bitset_add_range.Fn()(b, start, nItems) }
func (b *Bitset) RemoveRange(start, nItems uint32)  { gtk_bitset_remove_range.Fn()(b, start, nItems) }
func (b *Bitset) AddRangeClosed(first, last uint32) { gtk_bitset_add_range_closed.Fn()(b, first, last) }
func (b *Bitset) RemoveRangeClosed(first, last uint32) {
	gtk_bitset_remove_range_closed.Fn()(b, first, last)
}
func (b *Bitset) AddRectangle(start, width, height, stride uint32) {
	gtk_bitset_add_rectangle.Fn()(b, start, width, height, stride)
}
func (b *Bitset) RemoveRectangle(start, width, height, stride uint32) {
	gtk_bitset_remove_rectangle.Fn()(b, start, width, height, stride)
}
func (b *Bitset) Union(other *Bitset)      { gtk_bitset_union.Fn()(b, other) }
func (b *Bitset) Intersect(other *Bitset)  { gtk_bitset_intersect.Fn()(b, other) }
func (b *Bitset) Subtract(other *Bitset)   { gtk_bitset_subtract.Fn()(b, other) }
func (b *Bitset) Difference(other *Bitset) { gtk_bitset_difference.Fn()(b, other) }
func (b *Bitset) ShiftLeft(amount uint32)  { gtk_bitset_shift_left.Fn()(b, amount) }
func (b *Bitset) ShiftRight(amount uint32) { gtk_bitset_shift_right.Fn()(b, amount) }
func (b *Bitset) Splice(position, removed, added uint32) {
	gtk_bitset_splice.Fn()(b, position, removed, added)
}

type BitsetIter struct{ _ [10]uptr }

func GTypeBitsetIter() gobject.GType { return gtk_bitset_iter_get_type.Fn()() }
func (i *BitsetIter) InitFirst(set *Bitset) (value uint32, ok bool) {
	ok = gtk_bitset_iter_init_first.Fn()(i, set, &value) != 0
	return
}
func (i *BitsetIter) InitLast(set *Bitset) (value uint32, ok bool) {
	ok = gtk_bitset_iter_init_last.Fn()(i, set, &value) != 0
	return
}
func (i *BitsetIter) InitAt(set *Bitset, target uint32) (value uint32, ok bool) {
	ok = gtk_bitset_iter_init_at.Fn()(i, set, target, &value) != 0
	return
}
func (i *BitsetIter) Next() (value uint32, ok bool) {
	ok = gtk_bitset_iter_next.Fn()(i, &value) != 0
	return
}
func (i *BitsetIter) Previous() (value uint32, ok bool) {
	ok = gtk_bitset_iter_previous.Fn()(i, &value) != 0
	return
}
func (i *BitsetIter) GetValue() uint32 { return gtk_bitset_iter_get_value.Fn()(i) }
func (i *BitsetIter) IsValid() bool    { return gtk_bitset_iter_is_valid.Fn()(i) != 0 }

// #endregion

// #region BookmarkList

type BookmarkList struct {
	gobjCore
	gio.GListModel
}

func GTypeBookmarkList() gobject.GType { return gtk_bookmark_list_get_type.Fn()() }
func NewBookmarkList(filename, attributes string) *BookmarkList {
	f, a := cc.NewString(filename), cc.NewString(attributes)
	defer f.Free()
	defer a.Free()
	return gtk_bookmark_list_new.Fn()(f, a)
}
func (b *BookmarkList) GetFilename() string { return gtk_bookmark_list_get_filename.Fn()(b).String() }
func (b *BookmarkList) SetAttributes(attrs string) {
	a := cc.NewString(attrs)
	defer a.Free()
	gtk_bookmark_list_set_attributes.Fn()(b, a)
}
func (b *BookmarkList) GetAttributes() string {
	return gtk_bookmark_list_get_attributes.Fn()(b).String()
}
func (b *BookmarkList) SetIOPriority(prio int32) { gtk_bookmark_list_set_io_priority.Fn()(b, prio) }
func (b *BookmarkList) GetIOPriority() int32     { return gtk_bookmark_list_get_io_priority.Fn()(b) }
func (b *BookmarkList) IsLoading() bool          { return gtk_bookmark_list_is_loading.Fn()(b) != 0 }

// #endregion

// #region BoolFilter

type BoolFilter struct{ Filter }

func GTypeBoolFilter() gobject.GType { return gtk_bool_filter_get_type.Fn()() }
func NewBoolFilter(expression ExpressionIface) *BoolFilter {
	return gtk_bool_filter_new.Fn()(GetExpressionIface(expression))
}
func (b *BoolFilter) GetExpression() *Expression { return gtk_bool_filter_get_expression.Fn()(b) }
func (b *BoolFilter) SetExpression(expr ExpressionIface) {
	gtk_bool_filter_set_expression.Fn()(b, GetExpressionIface(expr))
}
func (b *BoolFilter) GetInvert() bool       { return gtk_bool_filter_get_invert.Fn()(b) != 0 }
func (b *BoolFilter) SetInvert(invert bool) { gtk_bool_filter_set_invert.Fn()(b, cbool(invert)) }

// #endregion

// #region Border

type Border struct {
	Left, Right, Top, Bottom int16
}

func GTypeBorder() gobject.GType { return gtk_border_get_type.Fn()() }
func NewBorder() *Border         { return gtk_border_new.Fn()() }
func (b *Border) Copy() *Border  { return gtk_border_copy.Fn()(b) }
func (b *Border) Free()          { gtk_border_free.Fn()(b) }

// #endregion

// #region Box

type BoxObj struct{ Parent WidgetObj }
type BoxClass struct {
	Parent WidgetClass
	_      [8]uptr //padding
}

type BoxIface interface{ GetBoxIface() *Box }

func GetBoxIface(iface BoxIface) *Box {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetBoxIface()
}
func (b *Box) GetBoxIface() *Box { return b }

type Box struct {
	Widget
	Orientable
}

func GTypeBox() gobject.GType { return gtk_box_get_type.Fn()() }
func NewBox(orientation Orientation, spacing int32) *Box {
	return gtk_box_new.Fn()(orientation, spacing)
}
func (b *Box) SetHomogeneous(homogeneous bool) { gtk_box_set_homogeneous.Fn()(b, cbool(homogeneous)) }
func (b *Box) GetHomogeneous() bool            { return gtk_box_get_homogeneous.Fn()(b) != 0 }
func (b *Box) SetSpacing(spacing int32)        { gtk_box_set_spacing.Fn()(b, spacing) }
func (b *Box) GetSpacing() int32               { return gtk_box_get_spacing.Fn()(b) }
func (b *Box) SetBaselinePosition(position BaselinePosition) {
	gtk_box_set_baseline_position.Fn()(b, position)
}
func (b *Box) GetBaselinePosition() BaselinePosition { return gtk_box_get_baseline_position.Fn()(b) }
func (b *Box) SetBaselineChild(child int32)          { gtk_box_set_baseline_child.Fn()(b, child) }
func (b *Box) GetBaselineChild() int32               { return gtk_box_get_baseline_child.Fn()(b) }
func (b *Box) Append(child WidgetIface)              { gtk_box_append.Fn()(b, GetWidgetIface(child)) }
func (b *Box) Prepend(child WidgetIface)             { gtk_box_prepend.Fn()(b, GetWidgetIface(child)) }
func (b *Box) Remove(child WidgetIface)              { gtk_box_remove.Fn()(b, GetWidgetIface(child)) }
func (b *Box) InsertChildAfter(child, sibling WidgetIface) {
	gtk_box_insert_child_after.Fn()(b, GetWidgetIface(child), GetWidgetIface(sibling))
}
func (b *Box) ReorderChildAfter(child, sibling WidgetIface) {
	gtk_box_reorder_child_after.Fn()(b, GetWidgetIface(child), GetWidgetIface(sibling))
}

// #endregion

// #region BoxLayout

type BoxLayout struct {
	LayoutManager
	Orientable
}

func GTypeBoxLayout() gobject.GType                   { return gtk_box_layout_get_type.Fn()() }
func NewBoxLayout(orientation Orientation) *BoxLayout { return gtk_box_layout_new.Fn()(orientation) }
func (b *BoxLayout) SetHomogeneous(homogeneous bool) {
	gtk_box_layout_set_homogeneous.Fn()(b, cbool(homogeneous))
}
func (b *BoxLayout) GetHomogeneous() bool      { return gtk_box_layout_get_homogeneous.Fn()(b) != 0 }
func (b *BoxLayout) SetSpacing(spacing uint32) { gtk_box_layout_set_spacing.Fn()(b, spacing) }
func (b *BoxLayout) GetSpacing() uint32        { return gtk_box_layout_get_spacing.Fn()(b) }
func (b *BoxLayout) SetBaselinePosition(position BaselinePosition) {
	gtk_box_layout_set_baseline_position.Fn()(b, position)
}
func (b *BoxLayout) GetBaselinePosition() BaselinePosition {
	return gtk_box_layout_get_baseline_position.Fn()(b)
}
func (b *BoxLayout) SetBaselineChild(child int32) { gtk_box_layout_set_baseline_child.Fn()(b, child) }
func (b *BoxLayout) GetBaselineChild() int32      { return gtk_box_layout_get_baseline_child.Fn()(b) }

// #endregion

// #region Buildable

type BuildableIfaceObj struct {
	GIface gobject.GTypeInterface

	SetID                cc.Func // void (* set_id) (GtkBuildable *buildable, const char *id);
	GetID                cc.Func // const char * (* get_id) (GtkBuildable *buildable);
	AddChild             cc.Func // void (* add_child) (GtkBuildable *buildable,GtkBuilder *builder,GObject *child,const char *type);
	SetBuildableProperty cc.Func // void (* set_buildable_property) (GtkBuildable *buildable,GtkBuilder *builder,const char *name,const GValue *value);
	ConstructChild       cc.Func // GObject * (* construct_child) (GtkBuildable *buildable,GtkBuilder *builder,const char *name);
	CustomTagStart       cc.Func // gboolean (* custom_tag_start) (GtkBuildable *buildable, GtkBuilder *builder, GObject *child, const char *tagname, GtkBuildableParser *parser, gpointer *data);
	CustomTagEnd         cc.Func // void (* custom_tag_end) (GtkBuildable *buildable, GtkBuilder *builder, GObject *child, const char *tagname, gpointer data);
	CustomFinished       cc.Func // void (* custom_finished) (GtkBuildable *buildable, GtkBuilder *builder, GObject *child, const char *tagname, gpointer data);
	ParserFinished       cc.Func // void (* parser_finished) (GtkBuildable *buildable, GtkBuilder *builder);
	GetInternalChild     cc.Func //GObject * (* get_internal_child) (GtkBuildable *buildable, GtkBuilder *builder, const char *childname);
}

type BuildableIface interface {
	GetBuildableIface() *Buildable
}

func GetBuildableIface(b BuildableIface) *Buildable {
	if anyptr(b) == nil {
		return nil
	}
	return b.GetBuildableIface()
}

func (b *Buildable) GetBuildableIface() *Buildable { return b }

type Buildable struct{}

type BuildableParser struct {
	StartElement cc.Func // void (*start_element)(GtkBuildableParseContext *context,const char *element_name,const char**attribute_names,const char**attribute_values,gpointeruser_data,GError**error);
	EndElement   cc.Func // void (*end_element)(GtkBuildableParseContext *context,const char *element_name,gpointeruser_data,GError**error);
	Text         cc.Func // void (*text) (GtkBuildableParseContext *context,const char *text,gsize text_len,gpointeruser_data,GError**error);
	Error        cc.Func //void (*error)(GtkBuildableParseContext *context,GError *error,gpointer user_data);
	_            [4]uptr
}
type BuildableParseContext struct{}

func GTypeBuildable() gobject.GType         { return gtk_buildable_get_type.Fn()() }
func (b *Buildable) GetBuildableID() string { return gtk_buildable_get_buildable_id.Fn()(b).String() }
func (c *BuildableParseContext) Push(parser *BuildableParser, userData uptr) {
	gtk_buildable_parse_context_push.Fn()(c, parser, userData)
}
func (c *BuildableParseContext) Pop() uptr { return gtk_buildable_parse_context_pop.Fn()(c) }
func (c *BuildableParseContext) GetElement() string {
	return gtk_buildable_parse_context_get_element.Fn()(c).String()
}
func (c *BuildableParseContext) GetElementStack() *glib.GPtrArray[cc.String] {
	return gtk_buildable_parse_context_get_element_stack.Fn()(c)
}
func (c *BuildableParseContext) GetPosition() (lineNumber, charNumber int32) {
	gtk_buildable_parse_context_get_position.Fn()(c, &lineNumber, &charNumber)
	return
}

// #endregion

// #region Builder

type Builder struct{ gobjCore }

func BuilderErrorQuark() glib.GQuark { return gtk_builder_error_quark.Fn()() }
func GTypeBuilder() gobject.GType    { return gtk_builder_get_type.Fn()() }
func NewBuilder() *Builder           { return gtk_builder_new.Fn()() }
func (b *Builder) AddFromFile(filename string) error {
	var gerr *glib.GError
	cstr := cc.NewString(filename)
	defer cstr.Free()
	ok := gtk_builder_add_from_file.Fn()(b, cstr, &gerr) != 0
	if !ok {
		defer gerr.Free()
		return gerr.Error()
	}
	return nil
}
func (b *Builder) AddFromResource(resourcePath string) error {
	var gerr *glib.GError
	cstr := cc.NewString(resourcePath)
	defer cstr.Free()
	ok := gtk_builder_add_from_resource.Fn()(b, cstr, &gerr) != 0
	if !ok {
		defer gerr.Free()
		return gerr.Error()
	}
	return nil
}
func (b *Builder) AddFromString(buffer string) error {
	var gerr *glib.GError
	buf, len := cc.NewStringRef(buffer)
	gtk_builder_add_from_string.Fn()(b, buf, len, &gerr)
	return gerr.TakeError()
}
func (b *Builder) AddObjectsFromFile(filename string, objectIds []string) error {
	var gerr *glib.GError
	cs, cpp := cc.NewString(filename), cc.NewStrings(objectIds)
	defer cpp.Free()
	defer cs.Free()
	gtk_builder_add_objects_from_file.Fn()(b, cs, cpp, &gerr)
	return gerr.TakeError()
}
func (b *Builder) AddObjectsFromResource(resourcePath string, objectIds []string) error {
	var gerr *glib.GError
	cs, cpp := cc.NewString(resourcePath), cc.NewStrings(objectIds)
	defer cpp.Free()
	defer cs.Free()
	gtk_builder_add_objects_from_resource.Fn()(b, cs, cpp, &gerr)
	return gerr.TakeError()
}
func (b *Builder) AddObjectsFromString(buffer string, objectIds []string) error {
	var gerr *glib.GError
	buf, len := cc.NewStringRef(buffer)
	cpp := cc.NewStrings(objectIds)
	defer cpp.Free()
	gtk_builder_add_objects_from_string.Fn()(b, buf, len, cpp, &gerr)
	return gerr.TakeError()
}
func (b *Builder) GetObject(name string) *gobject.GObject {
	cstr := cc.NewString(name)
	defer cstr.Free()
	return gtk_builder_get_object.Fn()(b, cstr)
}
func (b *Builder) GetObjects() *glib.GSList[gobject.GObject] {
	return (*glib.GSList[gobject.GObject])(gtk_builder_get_objects.Fn()(b))
}
func (b *Builder) ExposeObject(name string, object *gobject.GObject) {
	cstr := cc.NewString(name)
	defer cstr.Free()
	gtk_builder_expose_object.Fn()(b, cstr, object)
}
func (b *Builder) GetCurrentObject() *gobject.GObject { return gtk_builder_get_current_object.Fn()(b) }
func (b *Builder) SetCurrentObject(currentObject GobjIface) {
	gtk_builder_set_current_object.Fn()(b, gobjGet(currentObject))
}

func (b *Builder) SetTranslationDomain(domain string) {
	cstr := cc.NewString(domain)
	defer cstr.Free()
	gtk_builder_set_translation_domain.Fn()(b, cstr)
}
func (b *Builder) GetTranslationDomain() string {
	return gtk_builder_get_translation_domain.Fn()(b).String()
}
func (b *Builder) GetScope() *BuilderScope      { return gtk_builder_get_scope.Fn()(b) }
func (b *Builder) SetScope(scope *BuilderScope) { gtk_builder_set_scope.Fn()(b, scope) }
func (b *Builder) GetTypeFromName(typeName string) gobject.GType {
	cstr := cc.NewString(typeName)
	defer cstr.Free()
	return gtk_builder_get_type_from_name.Fn()(b, cstr)
}
func (b *Builder) ValueFromString(pspec *gobject.GParamSpec, str string, value *gobject.GValue) error {
	var gerr *glib.GError
	cstr := cc.NewString(str)
	defer cstr.Free()
	gtk_builder_value_from_string.Fn()(b, pspec, cstr, value, &gerr)
	return gerr.TakeError()
}
func (b *Builder) ValueFromStringType(typ gobject.GType, str string, value *gobject.GValue) error {
	var gerr *glib.GError
	cstr := cc.NewString(str)
	defer cstr.Free()
	gtk_builder_value_from_string_type.Fn()(b, typ, cstr, value, &gerr)
	return gerr.TakeError()
}

func NewBuilderFromFile(filename string) *Builder {
	cstr := cc.NewString(filename)
	defer cstr.Free()
	return gtk_builder_new_from_file.Fn()(cstr)
}
func NewBuilderFromResource(resourcePath string) *Builder {
	cstr := cc.NewString(resourcePath)
	defer cstr.Free()
	return gtk_builder_new_from_resource.Fn()(cstr)
}
func NewBuilderFromString(str string, length int64) *Builder {
	cstr := cc.NewString(str)
	defer cstr.Free()
	return gtk_builder_new_from_string.Fn()(cstr, length)
}

// func (b *Builder) CreateClosure(functionName string, flags BuilderClosureFlags, object *gobject.GObject) (*gobject.GClosure, error) {
// 	var gerr *glib.GError
// 	closure := gtk_builder_create_closure(b, functionName, flags, object, &gerr)
// 	if gerr != nil {
// 		defer gerr.Free()
// 		return nil, errors.New(gerr.Message.String())
// 	}
// 	return closure, nil
// }

// #endregion

// #region BuilderListItemFactory

type BuilderListItemFactory struct{ ListItemFactory }

func GTypeBuilderListItemFactory() gobject.GType {
	return gtk_builder_list_item_factory_get_type.Fn()()
}
func NewBuilderListItemFactoryFromBytes(scope *BuilderScope, bytes []byte) *BuilderListItemFactory {
	bts := glib.NewGBytes(bytes)
	defer bts.Unref()
	return gtk_builder_list_item_factory_new_from_bytes.Fn()(scope, bts)
}
func NewBuilderListItemFactoryFromResource(scope *BuilderScope, resourcePath string) *BuilderListItemFactory {
	cstr := cc.NewString(resourcePath)
	defer cstr.Free()
	return gtk_builder_list_item_factory_new_from_resource.Fn()(scope, cstr)
}
func (b *BuilderListItemFactory) GetBytes() *glib.GBytes {
	return gtk_builder_list_item_factory_get_bytes.Fn()(b)
}
func (b *BuilderListItemFactory) GetResource() string {
	return gtk_builder_list_item_factory_get_resource.Fn()(b).String()
}
func (b *BuilderListItemFactory) GetScope() *BuilderScope {
	return gtk_builder_list_item_factory_get_scope.Fn()(b)
}

// #endregion

// #region BuilderScope

type BuilderScopeIfaceObj struct {
	GIface gobject.GTypeInterface

	GetTypeFromName     cc.Func // GType (* get_type_from_name) (GtkBuilderScope *self, GtkBuilder *builder, const char *type_name);
	GetTypeFromFunction cc.Func // GType (* get_type_from_function) (GtkBuilderScope *self, GtkBuilder *builder, const char *function_name);
	CreateClosure       cc.Func // GClosure * (* create_closure) (GtkBuilderScope *self, GtkBuilder *builder, const char *function_name, GtkBuilderClosureFlags flags, GObject *object, GError **error);
}
type BuilderCScopeClass struct {
	Parent gobject.GObjectClass
}

type BuilderCScope struct {
	gobjCore
	BuilderScope
}
type BuilderScope struct{}

func GTypeBuilderCScope() gobject.GType { return gtk_builder_cscope_get_type.Fn()() }
func NewBuilderCScope() *BuilderCScope  { return gtk_builder_cscope_new.Fn()() }
func (b *BuilderCScope) AddCallbackSymbol(callbackName string, callbackSymbol func()) {
	cbn := cc.NewString(callbackName)
	defer cbn.Free()

	cbk := cc.Cbk(callbackSymbol)
	b.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	gtk_builder_cscope_add_callback_symbol.Fn()(b, cbn, cbk)
}

//	func (b *BuilderCScope) AddCallbackSymbols(firstCallbackName string, firstCallbackSymbol uintptr, args ...interface{}) {
//		gtk_builder_cscope_add_callback_symbols(b, firstCallbackName, firstCallbackSymbol, args...)
//	}

func (b *BuilderCScope) LookupCallbackSymbol(callbackName string) uintptr {
	cb := cc.NewString(callbackName)
	defer cb.Free()
	return gtk_builder_cscope_lookup_callback_symbol.Fn()(b, cb)
}

// #endregion

// #region Button

type ButtonObj struct{ Parent WidgetObj }
type ButtonClass struct {
	Parent   WidgetClass
	Clicked  cc.Func // void (* clicked)  (GtkButton *button);
	Activate cc.Func // void (* activate) (GtkButton *button);
	_        [8]uptr //padding
}
type ButtonIface interface {
	GetButtonIface() *Button
}

func GetButtonIface(iface ButtonIface) *Button {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetButtonIface()
}
func (b *Button) GetButtonIface() *Button { return b }

type Button struct {
	Widget
	Actionable
}

func GTypeButton() gobject.GType { return gtk_button_get_type.Fn()() }
func NewButton() *Button         { return gtk_button_new.Fn()() }
func NewButtonWithLabel(label string) *Button {
	cLabel := cc.NewString(label)
	defer cLabel.Free()
	return gtk_button_new_with_label.Fn()(cLabel)
}
func NewButtonFromIconName(iconName string) *Button {
	str := cc.NewString(iconName)
	defer str.Free()
	return gtk_button_new_from_icon_name.Fn()(str)
}
func NewButtonWithMnemonic(label string) *Button {
	str := cc.NewString(label)
	defer str.Free()
	return gtk_button_new_with_mnemonic.Fn()(str)
}
func (b *Button) SetHasFrame(hasFrame bool) { gtk_button_set_has_frame.Fn()(b, cbool(hasFrame)) }
func (b *Button) GetHasFrame() bool         { return gtk_button_get_has_frame.Fn()(b) != 0 }
func (b *Button) SetLabel(label string) {
	str := cc.NewString(label)
	defer str.Free()
	gtk_button_set_label.Fn()(b, str)
}
func (b *Button) GetLabel() string { return gtk_button_get_label.Fn()(b).String() }
func (b *Button) SetUseUnderline(useUnderline bool) {
	gtk_button_set_use_underline.Fn()(b, cbool(useUnderline))
}
func (b *Button) GetUseUnderline() bool { return gtk_button_get_use_underline.Fn()(b) != 0 }

func (b *Button) SetIconName(iconName string) {
	str := cc.NewString(iconName)
	defer str.Free()
	gtk_button_set_icon_name.Fn()(b, str)
}
func (b *Button) GetIconName() string         { return gtk_button_get_icon_name.Fn()(b).String() }
func (b *Button) SetChild(child WidgetIface)  { gtk_button_set_child.Fn()(b, GetWidgetIface(child)) }
func (b *Button) GetChild() *Widget           { return gtk_button_get_child.Fn()(b) }
func (b *Button) SetCanShrink(canShrink bool) { gtk_button_set_can_shrink.Fn()(b, cbool(canShrink)) }
func (b *Button) GetCanShrink() bool          { return gtk_button_get_can_shrink.Fn()(b) != 0 }
func (b *Button) ConnectActivate(sig func(b *Button)) uint64 {
	return b.SignalConnect("activate", sig, nil)
}
func (b *Button) ConnectClicked(sig func(b *Button)) uint64 {
	return b.SignalConnect("clicked", sig, nil)
}

// #endregion

// #region Calendar

type Calendar struct{ Widget }

func GTypeCalendar() gobject.GType                 { return gtk_calendar_get_type.Fn()() }
func NewCalendar() *Calendar                       { return gtk_calendar_new.Fn()() }
func (c *Calendar) SelectDay(date *glib.GDateTime) { gtk_calendar_select_day.Fn()(c, date) }
func (c *Calendar) MarkDay(day uint32)             { gtk_calendar_mark_day.Fn()(c, day) }
func (c *Calendar) UnmarkDay(day uint32)           { gtk_calendar_unmark_day.Fn()(c, day) }
func (c *Calendar) ClearMarks()                    { gtk_calendar_clear_marks.Fn()(c) }
func (c *Calendar) SetShowWeekNumbers(value bool) {
	gtk_calendar_set_show_week_numbers.Fn()(c, cbool(value))
}
func (c *Calendar) GetShowWeekNumbers() bool   { return gtk_calendar_get_show_week_numbers.Fn()(c) != 0 }
func (c *Calendar) SetShowHeading(value bool)  { gtk_calendar_set_show_heading.Fn()(c, cbool(value)) }
func (c *Calendar) GetShowHeading() bool       { return gtk_calendar_get_show_heading.Fn()(c) != 0 }
func (c *Calendar) SetShowDayNames(value bool) { gtk_calendar_set_show_day_names.Fn()(c, cbool(value)) }
func (c *Calendar) GetShowDayNames() bool      { return gtk_calendar_get_show_day_names.Fn()(c) != 0 }
func (c *Calendar) SetDay(day int32)           { gtk_calendar_set_day.Fn()(c, day) }
func (c *Calendar) GetDay() int32              { return gtk_calendar_get_day.Fn()(c) }
func (c *Calendar) SetMonth(month int32)       { gtk_calendar_set_month.Fn()(c, month) }
func (c *Calendar) GetMonth() int32            { return gtk_calendar_get_month.Fn()(c) }
func (c *Calendar) SetYear(year int32)         { gtk_calendar_set_year.Fn()(c, year) }
func (c *Calendar) GetYear() int32             { return gtk_calendar_get_year.Fn()(c) }

// func (c *Calendar) GetDate() *glib.GDateTime       { return gtk_calendar_get_date(c) }

func (c *Calendar) GetDayIsMarked(day uint32) bool {
	return gtk_calendar_get_day_is_marked.Fn()(c, day) != 0
}
func (c *Calendar) ConnectDaySelected(sig func(c *Calendar)) {
	c.SignalConnect("day-selected", sig, nil)
}
func (c *Calendar) ConnectNextMonth(sig func(c *Calendar)) {
	c.SignalConnect("next-month", sig, nil)
}
func (c *Calendar) ConnectNextYear(sig func(c *Calendar)) {
	c.SignalConnect("next-year", sig, nil)
}
func (c *Calendar) ConnectPrevMonth(sig func(c *Calendar)) {
	c.SignalConnect("prev-month", sig, nil)
}
func (c *Calendar) ConnectPrevYear(sig func(c *Calendar)) {
	c.SignalConnect("prev-year", sig, nil)
}

// #endregion

// #region CenterBox

type CenterBox struct {
	Widget

	Orientable
}

func GTypeCenterBox() gobject.GType { return gtk_center_box_get_type.Fn()() }
func NewCenterBox() *CenterBox      { return gtk_center_box_new.Fn()() }
func (cb *CenterBox) SetStartWidget(child WidgetIface) {
	gtk_center_box_set_start_widget.Fn()(cb, GetWidgetIface(child))
}
func (cb *CenterBox) SetCenterWidget(child WidgetIface) {
	gtk_center_box_set_center_widget.Fn()(cb, GetWidgetIface(child))
}
func (cb *CenterBox) SetEndWidget(child WidgetIface) {
	gtk_center_box_set_end_widget.Fn()(cb, GetWidgetIface(child))
}
func (cb *CenterBox) GetStartWidget() *Widget  { return gtk_center_box_get_start_widget.Fn()(cb) }
func (cb *CenterBox) GetCenterWidget() *Widget { return gtk_center_box_get_center_widget.Fn()(cb) }
func (cb *CenterBox) GetEndWidget() *Widget    { return gtk_center_box_get_end_widget.Fn()(cb) }
func (cb *CenterBox) SetBaselinePosition(position BaselinePosition) {
	gtk_center_box_set_baseline_position.Fn()(cb, position)
}
func (cb *CenterBox) GetBaselinePosition() BaselinePosition {
	return gtk_center_box_get_baseline_position.Fn()(cb)
}
func (cb *CenterBox) SetShrinkCenterLast(shrinkCenterLast bool) {
	gtk_center_box_set_shrink_center_last.Fn()(cb, cbool(shrinkCenterLast))
}

func (cb *CenterBox) GetShrinkCenterLast() bool {
	return gtk_center_box_get_shrink_center_last.Fn()(cb) != 0
}

// #endregion

// #region CenterLayout

type CenterLayout struct {
	gobjCore
	LayoutManager
}

func GTypeCenterLayout() gobject.GType { return gtk_center_layout_get_type.Fn()() } // 假设已有对应的类型获取函数
func NewCenterLayout() *CenterLayout   { return gtk_center_layout_new.Fn()() }
func (cl *CenterLayout) SetOrientation(orientation Orientation) {
	gtk_center_layout_set_orientation.Fn()(cl, orientation)
}
func (cl *CenterLayout) GetOrientation() Orientation {
	return gtk_center_layout_get_orientation.Fn()(cl)
}
func (cl *CenterLayout) SetBaselinePosition(baselinePosition BaselinePosition) {
	gtk_center_layout_set_baseline_position.Fn()(cl, baselinePosition)
}
func (cl *CenterLayout) GetBaselinePosition() BaselinePosition {
	return gtk_center_layout_get_baseline_position.Fn()(cl)
}
func (cl *CenterLayout) SetStartWidget(widget WidgetIface) {
	gtk_center_layout_set_start_widget.Fn()(cl, widget.GetWidgetIface())
}
func (cl *CenterLayout) GetStartWidget() *Widget {
	return gtk_center_layout_get_start_widget.Fn()(cl)
}
func (cl *CenterLayout) SetCenterWidget(widget WidgetIface) {
	gtk_center_layout_set_center_widget.Fn()(cl, widget.GetWidgetIface())
}
func (cl *CenterLayout) GetCenterWidget() *Widget {
	return gtk_center_layout_get_center_widget.Fn()(cl)
}
func (cl *CenterLayout) SetEndWidget(widget WidgetIface) {
	gtk_center_layout_set_end_widget.Fn()(cl, widget.GetWidgetIface())
}
func (cl *CenterLayout) GetEndWidget() *Widget {
	return gtk_center_layout_get_end_widget.Fn()(cl)
}
func (cl *CenterLayout) SetShrinkCenterLast(shrinkCenterLast bool) {
	gtk_center_layout_set_shrink_center_last.Fn()(cl, cbool(shrinkCenterLast))
}
func (cl *CenterLayout) GetShrinkCenterLast() bool {
	return gtk_center_layout_get_shrink_center_last.Fn()(cl) != 0
}

// #endregion

// #region CheckButton

type CheckButtonObj struct{ Parent WidgetObj }
type CheckButtonClass struct {
	Parent WidgetClass

	Toggled  cc.Func //void (* toggled) (GtkCheckButton *check_button);
	Activate cc.Func //void (* activate) (GtkCheckButton *check_button);

	_ [7]uptr //padding
}

type CheckButton struct {
	Widget

	Actionable
}

func GTypeCheckButton() gobject.GType { return gtk_check_button_get_type.Fn()() }
func NewCheckButton() *CheckButton    { return gtk_check_button_new.Fn()() }
func NewCheckButtonWithLabel(label string) *CheckButton {
	str := cc.NewString(label)
	defer str.Free()
	return gtk_check_button_new_with_label.Fn()(str)
}
func NewCheckButtonWithMnemonic(label string) *CheckButton {
	str := cc.NewString(label)
	defer str.Free()
	return gtk_check_button_new_with_mnemonic.Fn()(str)
}
func (b *CheckButton) SetInconsistent(inconsistent bool) {
	gtk_check_button_set_inconsistent.Fn()(b, cbool(inconsistent))
}
func (b *CheckButton) GetInconsistent() bool  { return gtk_check_button_get_inconsistent.Fn()(b) != 0 }
func (b *CheckButton) GetActive() bool        { return gtk_check_button_get_active.Fn()(b) != 0 }
func (b *CheckButton) SetActive(setting bool) { gtk_check_button_set_active.Fn()(b, cbool(setting)) }
func (b *CheckButton) GetLabel() string       { return gtk_check_button_get_label.Fn()(b).String() }
func (b *CheckButton) SetLabel(label string) {
	str := cc.NewString(label)
	defer str.Free()
	gtk_check_button_set_label.Fn()(b, str)
}
func (b *CheckButton) SetGroup(group *CheckButton) { gtk_check_button_set_group.Fn()(b, group) }
func (b *CheckButton) GetUseUnderline() bool       { return gtk_check_button_get_use_underline.Fn()(b) != 0 }
func (b *CheckButton) SetUseUnderline(setting bool) {
	gtk_check_button_set_use_underline.Fn()(b, cbool(setting))
}
func (b *CheckButton) GetChild() *Widget { return gtk_check_button_get_child.Fn()(b) }
func (b *CheckButton) SetChild(child WidgetIface) {
	gtk_check_button_set_child.Fn()(b, child.GetWidgetIface())
}
func (b *CheckButton) ConnectActivate(sig func(b *CheckButton)) {
	b.SignalConnect("activate", sig, nil)
}
func (b *CheckButton) ConnectToggled(sig func(b *CheckButton)) {
	b.SignalConnect("toggled", sig, nil)
}

// #endregion

// #region ColorDialog

type ColorDialog struct{ gobjCore }

func GTypeColorDialog() gobject.GType   { return gtk_color_dialog_get_type.Fn()() }
func NewColorDialog() *ColorDialog      { return gtk_color_dialog_new.Fn()() }
func (d *ColorDialog) GetTitle() string { return gtk_color_dialog_get_title.Fn()(d).String() }
func (d *ColorDialog) SetTitle(title string) {
	str := cc.NewString(title)
	defer str.Free()
	gtk_color_dialog_set_title.Fn()(d, str)
}
func (d *ColorDialog) GetModal() bool      { return gtk_color_dialog_get_modal.Fn()(d) != 0 }
func (d *ColorDialog) SetModal(modal bool) { gtk_color_dialog_set_modal.Fn()(d, cbool(modal)) }
func (d *ColorDialog) GetWithAlpha() bool  { return gtk_color_dialog_get_with_alpha.Fn()(d) != 0 }
func (d *ColorDialog) SetWithAlpha(withAlpha bool) {
	gtk_color_dialog_set_with_alpha.Fn()(d, cbool(withAlpha))
}
func (d *ColorDialog) ChooseRGBA(parent WindowIface, initialColor *gdk.RGBA, cancellable *gio.GCancellable,
	callback func(d *ColorDialog, res *gio.GAsyncResult)) {
	cbk := cc.Cbk(callback)
	d.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	gtk_color_dialog_choose_rgba.Fn()(d, GetWindowIface(parent), initialColor, cancellable, cbk, nil)
}
func (d *ColorDialog) ChooseRGBAFinish(result *gio.GAsyncResult) (rgba *gdk.RGBA, err error) {
	var gerr *glib.GError
	return gtk_color_dialog_choose_rgba_finish.Fn()(d, result, &gerr), gerr.TakeError()
}

// #endregion

// #region ColorDialogButton

type ColorDialogButton struct{ Widget }

func GTypeColorDialogButton() gobject.GType { return gtk_color_dialog_button_get_type.Fn()() }
func NewColorDialogButton(dialog *ColorDialog) *ColorDialogButton {
	return gtk_color_dialog_button_new.Fn()(dialog)
}
func (b *ColorDialogButton) GetDialog() *ColorDialog {
	return gtk_color_dialog_button_get_dialog.Fn()(b)
}
func (b *ColorDialogButton) SetDialog(dialog *ColorDialog) {
	gtk_color_dialog_button_set_dialog.Fn()(b, dialog)
}
func (b *ColorDialogButton) GetRGBA() *gdk.RGBA {
	return gtk_color_dialog_button_get_rgba.Fn()(b)
}
func (b *ColorDialogButton) SetRGBA(color *gdk.RGBA) {
	gtk_color_dialog_button_set_rgba.Fn()(b, color)
}
func (b *ColorDialogButton) ConnectActivate(sig func(b *ColorDialogButton)) uint64 {
	return b.SignalConnect("activate", sig, nil)
}

func HSVToRGB(h, s, v float32) (r, g, b float32) { gtk_hsv_to_rgb.Fn()(h, s, v, &r, &g, &b); return }
func RGBToHSV(r, g, b float32) (h, s, v float32) { gtk_rgb_to_hsv.Fn()(r, g, b, &h, &s, &v); return }

// #endregion

// #region ColumnView

type ColumnView struct {
	Widget

	Scrollable
}

func GTypeColumnView() gobject.GType { return gtk_column_view_get_type.Fn()() }
func NewColumnView(model SelectionModelIface) *ColumnView {
	return gtk_column_view_new.Fn()(GetSelectionModelIface(model))
}
func (v *ColumnView) GetColumns() *gio.GListModel { return gtk_column_view_get_columns.Fn()(v) }
func (v *ColumnView) AppendColumn(column *ColumnViewColumn) {
	gtk_column_view_append_column.Fn()(v, column)
}
func (v *ColumnView) RemoveColumn(column *ColumnViewColumn) {
	gtk_column_view_remove_column.Fn()(v, column)
}
func (v *ColumnView) InsertColumn(position uint32, column *ColumnViewColumn) {
	gtk_column_view_insert_column.Fn()(v, position, column)
}
func (v *ColumnView) GetModel() SelectionModelIface { return gtk_column_view_get_model.Fn()(v) }
func (v *ColumnView) SetModel(model SelectionModelIface) {
	gtk_column_view_set_model.Fn()(v, GetSelectionModelIface(model))
}
func (v *ColumnView) GetShowRowSeparators() bool {
	return gtk_column_view_get_show_row_separators.Fn()(v) != 0
}
func (v *ColumnView) SetShowRowSeparators(showRowSeparators bool) {
	gtk_column_view_set_show_row_separators.Fn()(v, cbool(showRowSeparators))
}
func (v *ColumnView) GetShowColumnSeparators() bool {
	return gtk_column_view_get_show_column_separators.Fn()(v) != 0
}
func (v *ColumnView) SetShowColumnSeparators(showColumnSeparators bool) {
	gtk_column_view_set_show_column_separators.Fn()(v, cbool(showColumnSeparators))
}
func (v *ColumnView) GetSorter() *Sorter { return gtk_column_view_get_sorter.Fn()(v) }
func (v *ColumnView) SortByColumn(column *ColumnViewColumn, direction SortType) {
	gtk_column_view_sort_by_column.Fn()(v, column, direction)
}
func (v *ColumnView) SetSingleClickActivate(singleClickActivate bool) {
	gtk_column_view_set_single_click_activate.Fn()(v, cbool(singleClickActivate))
}
func (v *ColumnView) GetSingleClickActivate() bool {
	return gtk_column_view_get_single_click_activate.Fn()(v) != 0
}
func (v *ColumnView) SetReorderable(reorderable bool) {
	gtk_column_view_set_reorderable.Fn()(v, cbool(reorderable))
}
func (v *ColumnView) GetReorderable() bool { return gtk_column_view_get_reorderable.Fn()(v) != 0 }
func (v *ColumnView) SetEnableRubberband(enableRubberband bool) {
	gtk_column_view_set_enable_rubberband.Fn()(v, cbool(enableRubberband))
}
func (v *ColumnView) GetEnableRubberband() bool {
	return gtk_column_view_get_enable_rubberband.Fn()(v) != 0
}
func (v *ColumnView) SetTabBehavior(tabBehavior ListTabBehavior) {
	gtk_column_view_set_tab_behavior.Fn()(v, tabBehavior)
}
func (v *ColumnView) GetTabBehavior() ListTabBehavior {
	return gtk_column_view_get_tab_behavior.Fn()(v)
}
func (v *ColumnView) SetRowFactory(factory ListItemFactoryIface) {
	gtk_column_view_set_row_factory.Fn()(v, GetListItemFactoryIface(factory))
}
func (v *ColumnView) GetRowFactory() *ListItemFactory { return gtk_column_view_get_row_factory.Fn()(v) }
func (v *ColumnView) SetHeaderFactory(factory ListItemFactoryIface) {
	gtk_column_view_set_header_factory.Fn()(v, GetListItemFactoryIface(factory))
}
func (v *ColumnView) GetHeaderFactory() *ListItemFactory {
	return gtk_column_view_get_header_factory.Fn()(v)
}
func (v *ColumnView) ScrollTo(pos uint32, column *ColumnViewColumn, flags ListScrollFlags, scroll *ScrollInfo) {
	gtk_column_view_scroll_to.Fn()(v, pos, column, flags, scroll)
}
func (v *ColumnView) ConnectActivate(sig func(v *ColumnView, pos uint32)) uint64 {
	return v.SignalConnect("activate", sig, nil)
}

// #endregion

// #region ColumnViewCell

type ColumnViewCell struct{ ListItem }

func GTypeColumnViewCell() gobject.GType      { return gtk_column_view_cell_get_type.Fn()() }
func (c *ColumnViewCell) GetItem() uptr       { return gtk_column_view_cell_get_item.Fn()(c) }
func (c *ColumnViewCell) GetPosition() uint32 { return gtk_column_view_cell_get_position.Fn()(c) }
func (c *ColumnViewCell) GetSelected() bool   { return gtk_column_view_cell_get_selected.Fn()(c) != 0 }
func (c *ColumnViewCell) GetFocusable() bool  { return gtk_column_view_cell_get_focusable.Fn()(c) != 0 }
func (c *ColumnViewCell) SetFocusable(focusable bool) {
	gtk_column_view_cell_set_focusable.Fn()(c, cbool(focusable))
}
func (c *ColumnViewCell) SetChild(child WidgetIface) {
	gtk_column_view_cell_set_child.Fn()(c, GetWidgetIface(child))
}
func (c *ColumnViewCell) GetChild() *Widget { return gtk_column_view_cell_get_child.Fn()(c) }

// #endregion

// #region ColumnViewColumn

type ColumnViewColumn struct {
	gobjCore
}

func GTypeColumnViewColumn() gobject.GType { return gtk_column_view_column_get_type.Fn()() }
func NewColumnViewColumn(title string, factory ListItemFactoryIface) *ColumnViewColumn {
	cTitle := cc.NewString(title)
	defer cTitle.Free()
	return gtk_column_view_column_new.Fn()(cTitle, GetListItemFactoryIface(factory))
}
func (c *ColumnViewColumn) GetColumnView() *ColumnView {
	return gtk_column_view_column_get_column_view.Fn()(c)
}
func (c *ColumnViewColumn) SetFactory(factory ListItemFactoryIface) {
	gtk_column_view_column_set_factory.Fn()(c, GetListItemFactoryIface(factory))
}
func (c *ColumnViewColumn) GetFactory() *ListItemFactory {
	return gtk_column_view_column_get_factory.Fn()(c)
}
func (c *ColumnViewColumn) SetTitle(title string) {
	cTitle := cc.NewString(title)
	defer cTitle.Free()
	gtk_column_view_column_set_title.Fn()(c, cTitle)
}
func (c *ColumnViewColumn) GetTitle() string {
	return gtk_column_view_column_get_title.Fn()(c).String()
}
func (c *ColumnViewColumn) SetSorter(sorter *Sorter) {
	gtk_column_view_column_set_sorter.Fn()(c, sorter)
}
func (c *ColumnViewColumn) GetSorter() *Sorter { return gtk_column_view_column_get_sorter.Fn()(c) }
func (c *ColumnViewColumn) SetVisible(visible bool) {
	gtk_column_view_column_set_visible.Fn()(c, cbool(visible))
}
func (c *ColumnViewColumn) GetVisible() bool { return gtk_column_view_column_get_visible.Fn()(c) != 0 }
func (c *ColumnViewColumn) SetHeaderMenu(menu *gio.GMenuModel) {
	gtk_column_view_column_set_header_menu.Fn()(c, menu)
}
func (c *ColumnViewColumn) GetHeaderMenu() *gio.GMenuModel {
	return gtk_column_view_column_get_header_menu.Fn()(c)
}
func (c *ColumnViewColumn) SetFixedWidth(fixedWidth int32) {
	gtk_column_view_column_set_fixed_width.Fn()(c, fixedWidth)
}
func (c *ColumnViewColumn) GetFixedWidth() int32 {
	return gtk_column_view_column_get_fixed_width.Fn()(c)
}
func (c *ColumnViewColumn) SetResizable(resizable bool) {
	gtk_column_view_column_set_resizable.Fn()(c, cbool(resizable))
}
func (c *ColumnViewColumn) GetResizable() bool {
	return gtk_column_view_column_get_resizable.Fn()(c) != 0
}
func (c *ColumnViewColumn) SetExpand(expand bool) {
	gtk_column_view_column_set_expand.Fn()(c, cbool(expand))
}
func (c *ColumnViewColumn) GetExpand() bool { return gtk_column_view_column_get_expand.Fn()(c) != 0 }
func (c *ColumnViewColumn) SetId(id string) {
	cId := cc.NewString(id)
	defer cId.Free()
	gtk_column_view_column_set_id.Fn()(c, cId)
}
func (c *ColumnViewColumn) GetId() string { return gtk_column_view_column_get_id.Fn()(c).String() }

// #endregion

// #region ColumnViewRow

type ColumnViewRow struct{ gobjCore }

func GTypeColumnViewRow() gobject.GType      { return gtk_column_view_row_get_type.Fn()() }
func (r *ColumnViewRow) GetItem() uptr       { return gtk_column_view_row_get_item.Fn()(r) }
func (r *ColumnViewRow) GetPosition() uint32 { return gtk_column_view_row_get_position.Fn()(r) }
func (r *ColumnViewRow) GetSelected() bool   { return gtk_column_view_row_get_selected.Fn()(r) != 0 }
func (r *ColumnViewRow) GetSelectable() bool { return gtk_column_view_row_get_selectable.Fn()(r) != 0 }
func (r *ColumnViewRow) SetSelectable(selectable bool) {
	gtk_column_view_row_set_selectable.Fn()(r, cbool(selectable))
}
func (r *ColumnViewRow) GetActivatable() bool {
	return gtk_column_view_row_get_activatable.Fn()(r) != 0
}
func (r *ColumnViewRow) SetActivatable(activatable bool) {
	gtk_column_view_row_set_activatable.Fn()(r, cbool(activatable))
}
func (r *ColumnViewRow) GetFocusable() bool { return gtk_column_view_row_get_focusable.Fn()(r) != 0 }
func (r *ColumnViewRow) SetFocusable(focusable bool) {
	gtk_column_view_row_set_focusable.Fn()(r, cbool(focusable))
}
func (r *ColumnViewRow) GetAccessibleDescription() string {
	return gtk_column_view_row_get_accessible_description.Fn()(r).String()
}
func (r *ColumnViewRow) SetAccessibleDescription(description string) {
	cDescription := cc.NewString(description)
	defer cDescription.Free()
	gtk_column_view_row_set_accessible_description.Fn()(r, cDescription)
}
func (r *ColumnViewRow) GetAccessibleLabel() string {
	return gtk_column_view_row_get_accessible_label.Fn()(r).String()
}
func (r *ColumnViewRow) SetAccessibleLabel(label string) {
	cLabel := cc.NewString(label)
	defer cLabel.Free()
	gtk_column_view_row_set_accessible_label.Fn()(r, cLabel)
}

// #endregion

// #region ColumnViewSorter

type ColumnViewSorter struct{ Sorter }

func GTypeColumnViewSorter() gobject.GType { return gtk_column_view_sorter_get_type.Fn()() }
func (s *ColumnViewSorter) GetPrimarySortColumn() *ColumnViewColumn {
	return gtk_column_view_sorter_get_primary_sort_column.Fn()(s)
}
func (s *ColumnViewSorter) GetPrimarySortOrder() SortType {
	return gtk_column_view_sorter_get_primary_sort_order.Fn()(s)
}
func (s *ColumnViewSorter) GetNSortColumns() uint32 {
	return gtk_column_view_sorter_get_n_sort_columns.Fn()(s)
}
func (s *ColumnViewSorter) GetNthSortColumn(position uint32) (*ColumnViewColumn, SortType) {
	var sortOrder SortType
	column := gtk_column_view_sorter_get_nth_sort_column.Fn()(s, position, &sortOrder)
	return column, sortOrder
}

// #endregion

// #region Constraint

type ConstraintTargetIface interface {
	GetConstraintTargetIface() *ConstraintTarget
}

func GetConstraintTargetIface(t ConstraintTargetIface) *ConstraintTarget {
	if anyptr(t) == nil {
		return nil
	}
	return t.GetConstraintTargetIface()
}

func (t *ConstraintTarget) GetConstraintTargetIface() *ConstraintTarget { return t }

type ConstraintTarget struct{}

type Constraint struct{ gobjCore }

func NewConstraint(target ConstraintTargetIface, targetAttribute ConstraintAttribute,
	relation ConstraintRelation, source ConstraintTargetIface, sourceAttribute ConstraintAttribute, multiplier float64, constant float64, strength ConstraintStrength) *Constraint {
	return gtk_constraint_new.Fn()(GetConstraintTargetIface(target), targetAttribute, relation, GetConstraintTargetIface(source), sourceAttribute, multiplier, constant, strength)
}
func NewConstantConstraint(target ConstraintTargetIface, targetAttribute ConstraintAttribute,
	relation ConstraintRelation, constant float64, strength ConstraintStrength) *Constraint {
	return gtk_constraint_new_constant.Fn()(GetConstraintTargetIface(target), targetAttribute, relation, constant, strength)
}
func (c *Constraint) GetTarget() *ConstraintTarget { return gtk_constraint_get_target.Fn()(c) }
func (c *Constraint) GetTargetAttribute() ConstraintAttribute {
	return gtk_constraint_get_target_attribute.Fn()(c)
}
func (c *Constraint) GetSource() *ConstraintTarget { return gtk_constraint_get_source.Fn()(c) }
func (c *Constraint) GetSourceAttribute() ConstraintAttribute {
	return gtk_constraint_get_source_attribute.Fn()(c)
}
func (c *Constraint) GetRelation() ConstraintRelation { return gtk_constraint_get_relation.Fn()(c) }
func (c *Constraint) GetMultiplier() float64          { return gtk_constraint_get_multiplier.Fn()(c) }
func (c *Constraint) GetConstant() float64            { return gtk_constraint_get_constant.Fn()(c) }
func (c *Constraint) GetStrength() int32              { return gtk_constraint_get_strength.Fn()(c) }
func (c *Constraint) IsRequired() bool                { return gtk_constraint_is_required.Fn()(c) != 0 }
func (c *Constraint) IsAttached() bool                { return gtk_constraint_is_attached.Fn()(c) != 0 }
func (c *Constraint) IsConstant() bool                { return gtk_constraint_is_constant.Fn()(c) != 0 }

// #endregion

// #region ConstraintGuide

type ConstraintGuide struct{ gobjCore }

func GTypeConstraintGuide() gobject.GType  { return gtk_constraint_guide_get_type.Fn()() }
func NewConstraintGuide() *ConstraintGuide { return gtk_constraint_guide_new.Fn()() }
func (g *ConstraintGuide) SetMinSize(width, height int32) {
	gtk_constraint_guide_set_min_size.Fn()(g, width, height)
}
func (g *ConstraintGuide) GetMinSize() (width, height int32) {
	gtk_constraint_guide_get_min_size.Fn()(g, &width, &height)
	return
}
func (g *ConstraintGuide) SetNatSize(width, height int32) {
	gtk_constraint_guide_set_nat_size.Fn()(g, width, height)
}
func (g *ConstraintGuide) GetNatSize() (width, height int32) {
	gtk_constraint_guide_get_nat_size.Fn()(g, &width, &height)
	return
}
func (g *ConstraintGuide) SetMaxSize(width, height int32) {
	gtk_constraint_guide_set_max_size.Fn()(g, width, height)
}
func (g *ConstraintGuide) GetMaxSize() (width, height int32) {
	gtk_constraint_guide_get_max_size.Fn()(g, &width, &height)
	return
}
func (g *ConstraintGuide) GetStrength() ConstraintStrength {
	return gtk_constraint_guide_get_strength.Fn()(g)
}
func (g *ConstraintGuide) SetStrength(strength ConstraintStrength) {
	gtk_constraint_guide_set_strength.Fn()(g, strength)
}
func (g *ConstraintGuide) SetName(name string) {
	n := cc.NewString(name)
	defer n.Free()
	gtk_constraint_guide_set_name.Fn()(g, n)
}
func (g *ConstraintGuide) GetName() string { return gtk_constraint_guide_get_name.Fn()(g).String() }

// #endregion

// #region ConstraintLayout

type ConstraintLayout struct {
	gobjCore
	LayoutManager
}

func GTypeConstraintLayout() gobject.GType       { return gtk_constraint_layout_get_type.Fn()() }
func ConstraintVflParserErrorQuark() glib.GQuark { return gtk_constraint_vfl_parser_error_quark.Fn()() }
func NewConstraintLayout() *ConstraintLayout     { return gtk_constraint_layout_new.Fn()() }
func (l *ConstraintLayout) AddConstraint(constraint *Constraint) {
	gtk_constraint_layout_add_constraint.Fn()(l, constraint)
}
func (l *ConstraintLayout) RemoveConstraint(constraint *Constraint) {
	gtk_constraint_layout_remove_constraint.Fn()(l, constraint)
}
func (l *ConstraintLayout) AddGuide(guide *ConstraintGuide) {
	gtk_constraint_layout_add_guide.Fn()(l, guide)
}
func (l *ConstraintLayout) RemoveGuide(guide *ConstraintGuide) {
	gtk_constraint_layout_remove_guide.Fn()(l, guide)
}

// #endregion

// #region CssProvider

type CssProvider struct {
	gobjCore

	StyleProvider
}
type CssProviderObj struct{ ParentInstance gobject.GObjectObj }

func GTypeCssProvider() gobject.GType   { return gtk_css_provider_get_type.Fn()() }
func NewCssProvider() *CssProvider      { return gtk_css_provider_new.Fn()() }
func (p *CssProvider) String() string   { return gtk_css_provider_to_string.Fn()(p).TakeString() }
func (p *CssProvider) ToString() string { return gtk_css_provider_to_string.Fn()(p).TakeString() }
func (p *CssProvider) LoadFromString(str string) {
	cstr := cc.NewString(str)
	defer cstr.Free()
	gtk_css_provider_load_from_string.Fn()(p, cstr)
}
func (p *CssProvider) LoadFromPath(path string) {
	cPath := cc.NewString(path)
	defer cPath.Free()
	gtk_css_provider_load_from_path.Fn()(p, cPath)
}
func (p *CssProvider) LoadFromResource(resourcePath string) {
	cResourcePath := cc.NewString(resourcePath)
	defer cResourcePath.Free()
	gtk_css_provider_load_from_resource.Fn()(p, cResourcePath)
}
func (p *CssProvider) LoadNamed(name, variant string) {
	cn, cv := cc.NewString(name), cc.NewString(variant)
	defer cn.Free()
	defer cv.Free()
	gtk_css_provider_load_named.Fn()(p, cv, cv)
}
func (p *CssProvider) AddProviderForDisplay(dsp *gdk.Display, priority StyleProviderPriority) {
	gtk_style_context_add_provider_for_display.Fn()(dsp, GetStyleProviderIface(p), priority)
}
func (p *CssProvider) AddProviderForDefaultDisplay(priority StyleProviderPriority) {
	gtk_style_context_add_provider_for_display.Fn()(gdk.GetDisplayDefault(), GetStyleProviderIface(p), priority)
}
func AddCssProviderForDefaultDisplay(filepath string, priority StyleProviderPriority) {
	css := NewCssProvider()
	css.LoadFromPath(filepath)
	gtk_style_context_add_provider_for_display.Fn()(gdk.GetDisplayDefault(), GetStyleProviderIface(css), priority)
}

// #endregion

// #region CustomFilter

type CustomFilter struct{ Filter }

func GTypeCustomFilter() gobject.GType { return gtk_custom_filter_get_type.Fn()() }
func NewCustomFilter[T any](matchFunc func(item *T) bool) *CustomFilter {
	var mf, des uptr
	if matchFunc != nil {
		mf = cc.CbkRaw[func(item, _ uptr) int32](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 2)
			r := matchFunc(*(**T)(is[0]))
			if r {
				*(*int32)(out) = 1
				return
			}
			*(*int32)(out) = 0
		})
		des = cc.Cbk(func(d uptr) { cc.CbkClose(mf); cc.CbkCloseLate(des) })
	}
	return gtk_custom_filter_new.Fn()(mf, nil, des)
}

func (f *CustomFilter) SetFilterFunc(matchFunc func(item uptr) bool) {
	var mf, des uptr
	if matchFunc != nil {
		mf = cc.CbkRaw[func(item, _ uptr) int32](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 2)
			r := matchFunc(*(*uptr)(is[0]))
			if r {
				*(*int32)(out) = 1
				return
			}
			*(*int32)(out) = 0
		})

		des = cc.Cbk(func(d uptr) { cc.CbkClose(mf); cc.CbkCloseLate(des) })
	}
	gtk_custom_filter_set_filter_func.Fn()(f, mf, nil, des)
}

// #endregion

// #region CustomLayout

type CustomLayout struct{ LayoutManager }

func GTypeCustomLayout() gobject.GType { return gtk_custom_layout_get_type.Fn()() }
func NewCustomLayout(
	requestMode func(wd *Widget) SizeRequestMode,
	measure func(wd *Widget, orientation Orientation, forSize int32, minimum *int32, natural *int32, minimumBaseline *int32, naturalBaseline *int32),
	allocate func(wd *Widget, width int32, height int32, baseline int32)) *CustomLayout {

	var rm, me, al uptr
	if requestMode != nil {
		rm = cc.CbkRaw[func(wd *Widget) SizeRequestMode](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 1)
			*(*SizeRequestMode)(out) = requestMode(*(**Widget)(is[0]))
		})
	}
	if measure != nil {
		me = cc.CbkRaw[func(wd *Widget, orientation Orientation, forSize int32,
			minimum, natural, minimumBaseline, naturalBaseline *int32)](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 7)

			wd := *(**Widget)(is[0])
			orientation := *(*Orientation)(is[1])
			forSize := *(*int32)(is[2])
			minimum := *(**int32)(is[3])
			natural := *(**int32)(is[4])
			minimumBaseline := *(**int32)(is[5])
			naturalBaseline := *(**int32)(is[6])

			measure(wd, orientation, forSize, minimum, natural, minimumBaseline, naturalBaseline)
		})
	}
	if allocate != nil {
		al = cc.CbkRaw[func(wd *Widget, width, height, baseline int32)](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 4)
			allocate(*(**Widget)(is[0]), *(*int32)(is[1]), *(*int32)(is[2]), *(*int32)(is[3]))
		})
	}
	lyt := gtk_custom_layout_new.Fn()(rm, me, al)
	lyt.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(rm); cc.CbkClose(me); cc.CbkClose(al) })
	return lyt
}

// #endregion

// #region CustomSorter

type CustomSorter struct{ Sorter }

func GTypeCustomSorter() gobject.GType { return gtk_custom_sorter_get_type.Fn()() }
func NewCustomSorter[T any](sortFunc func(a, b *T) int32) *CustomSorter {
	var sf, des uptr
	if sortFunc != nil {
		sf = cc.CbkRaw[func(item, _ uptr) int32](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 3)
			*(*int32)(out) = sortFunc(*(**T)(is[0]), *(**T)(is[1]))
		})
		des = cc.Cbk(func(d uptr) { cc.CbkClose(sf); cc.CbkCloseLate(des) })
	}
	return gtk_custom_sorter_new.Fn()(sf, nil, des)
}
func (s *CustomSorter) SetSortFunc(sortFunc func(a, b uptr) int32) {
	var sf, des uptr
	if sortFunc != nil {
		sf = cc.CbkRaw[func(item, _ uptr) int32](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 3)
			*(*int32)(out) = sortFunc(*(*uptr)(is[0]), *(*uptr)(is[1]))
		})
		des = cc.Cbk(func(d uptr) { cc.CbkClose(sf); cc.CbkCloseLate(des) })
	}
	gtk_custom_sorter_set_sort_func.Fn()(s, sf, nil, des)
}

// #endregion

// #region DebugFlags

func GetDebugFlags() DebugFlags      { return gtk_get_debug_flags.Fn()() }
func SetDebugFlags(flags DebugFlags) { gtk_set_debug_flags.Fn()(flags) }

// #endregion

// #region DialogError

func DialogErrorQuark() glib.GQuark { return gtk_dialog_error_quark.Fn()() }

// #endregion

// #region DirectoryList

type DirectoryList struct {
	gobjCore
	gio.GListModel
}

func GTypeDirectoryList() gobject.GType { return gtk_directory_list_get_type.Fn()() }
func NewDirectoryList(attributes string, file gio.GFileIface) *DirectoryList {
	cAttributes := cc.NewString(attributes)
	defer cAttributes.Free()
	return gtk_directory_list_new.Fn()(cAttributes, gio.GetGFileIface(file))
}
func (d *DirectoryList) SetFile(file gio.GFileIface) {
	gtk_directory_list_set_file.Fn()(d, gio.GetGFileIface(file))
}
func (d *DirectoryList) GetFile() *gio.GFile {
	return gtk_directory_list_get_file.Fn()(d)
}
func (d *DirectoryList) SetAttributes(attributes string) {
	cAttributes := cc.NewString(attributes)
	defer cAttributes.Free()
	gtk_directory_list_set_attributes.Fn()(d, cAttributes)
}
func (d *DirectoryList) GetAttributes() string {
	return gtk_directory_list_get_attributes.Fn()(d).String()
}
func (d *DirectoryList) SetIOPriority(ioPriority int32) {
	gtk_directory_list_set_io_priority.Fn()(d, ioPriority)
}
func (d *DirectoryList) GetIOPriority() int32 { return gtk_directory_list_get_io_priority.Fn()(d) }
func (d *DirectoryList) IsLoading() bool      { return gtk_directory_list_is_loading.Fn()(d) != 0 }
func (d *DirectoryList) GetError() error {
	gerr := gtk_directory_list_get_error.Fn()(d)
	if gerr == nil {
		return nil
	}
	return errors.New(gerr.Message.RefString())
}
func (d *DirectoryList) SetMonitored(monitored bool) {
	gtk_directory_list_set_monitored.Fn()(d, cbool(monitored))
}
func (d *DirectoryList) GetMonitored() bool { return gtk_directory_list_get_monitored.Fn()(d) != 0 }

// #endregion

// #region DragIcon

type DragIcon struct {
	Widget
	Root
}

func GTypeDragIcon() gobject.GType                { return gtk_drag_icon_get_type.Fn()() }
func GetDragIconForDrag(drag *gdk.Drag) *DragIcon { return gtk_drag_icon_get_for_drag.Fn()(drag) }
func (d *DragIcon) SetChild(child WidgetIface) {
	gtk_drag_icon_set_child.Fn()(d, GetWidgetIface(child))
}
func (d *DragIcon) GetChild() *Widget { return gtk_drag_icon_get_child.Fn()(d) }
func SetDragIconFromPaintable(drag *gdk.Drag, paintable *gdk.Paintable, hotX, hotY int32) {
	gtk_drag_icon_set_from_paintable.Fn()(drag, paintable, hotX, hotY)
}

// func DragIconCreateWidgetForValue(value *gobject.GValue) *DragIcon {
func DragIconCreateWidgetForValue(a interface{}) *DragIcon {
	v := gobject.NewGValue(a)
	defer v.Unset()
	return gtk_drag_icon_create_widget_for_value.Fn()(v)
}

// #endregion

// #region DragSource

type DragSource struct{ GestureSingle }

func GTypeDragSource() gobject.GType { return gtk_drag_source_get_type.Fn()() }
func NewDragSource() *DragSource     { return gtk_drag_source_new.Fn()() }
func (d *DragSource) SetContent(content gdk.ContentProviderIface) {
	gtk_drag_source_set_content.Fn()(d, gdk.GetContentProviderIface(content))
}
func (d *DragSource) GetContent() *gdk.ContentProvider  { return gtk_drag_source_get_content.Fn()(d) }
func (d *DragSource) SetActions(actions gdk.DragAction) { gtk_drag_source_set_actions.Fn()(d, actions) }
func (d *DragSource) GetActions() gdk.DragAction        { return gtk_drag_source_get_actions.Fn()(d) }
func (d *DragSource) SetIcon(paintable *gdk.Paintable, hotX, hotY int32) {
	gtk_drag_source_set_icon.Fn()(d, paintable, hotX, hotY)
}
func (d *DragSource) DragCancel()        { gtk_drag_source_drag_cancel.Fn()(d) }
func (d *DragSource) GetDrag() *gdk.Drag { return gtk_drag_source_get_drag.Fn()(d) }
func (w *Widget) DragCheckThreshold(startX, startY, currentX, currentY int32) bool {
	return gtk_drag_check_threshold.Fn()(w, startX, startY, currentX, currentY) != 0
}
func (c *DragSource) ConnectDragBegin(sig func(c *DragSource, drag *gdk.Drag)) uint64 {
	return c.SignalConnect("drag-begin", sig, nil)
}
func (c *DragSource) ConnectDragCancel(
	sig func(c *DragSource, drag *gdk.Drag, reason *gdk.DragCancelReason, _ uptr) (handled bool)) uint64 {
	return c.SignalConnect("drag-cancel", sig, nil)
}
func (c *DragSource) ConnectDragEnd(
	sig func(c *DragSource, drag *gdk.Drag, deleteData bool)) uint64 {
	return c.SignalConnect("drag-end", sig, nil)
}
func (c *DragSource) ConnectPrepare(
	sig func(c *DragSource, x, y float64, _ uptr) *gdk.ContentProvider) uint64 {
	return c.SignalConnect("prepare", sig, nil)
}

// #endregion

// #region DrawingArea

type DrawingAreaObj struct {
	Parent WidgetObj
}
type DrawingAreaClass struct {
	Parent WidgetClass
	Resize cc.Func

	_ [8]uptr // padding
}

type DrawingArea struct {
	Widget
}

func GTypeDrawingArea() gobject.GType              { return gtk_drawing_area_get_type.Fn()() }
func NewDrawingArea() *DrawingArea                 { return gtk_drawing_area_new.Fn()() }
func (d *DrawingArea) SetContentWidth(width int32) { gtk_drawing_area_set_content_width.Fn()(d, width) }
func (d *DrawingArea) GetContentWidth() int32      { return gtk_drawing_area_get_content_width.Fn()(d) }
func (d *DrawingArea) SetContentHeight(height int32) {
	gtk_drawing_area_set_content_height.Fn()(d, height)
}
func (d *DrawingArea) GetContentHeight() int32 { return gtk_drawing_area_get_content_height.Fn()(d) }
func (d *DrawingArea) SetDrawFunc(drawFunc func(da *DrawingArea, cr *cairo.Context, width, height int32)) {

	df := cc.CbkRaw[func(da *DrawingArea, cr *cairo.Context, width, height int32, _ uptr)](func(out, ins uptr) {
		is := unsafe.Slice((*uptr)(ins), 5)
		drawFunc(*(**DrawingArea)(is[0]), *(**cairo.Context)(is[1]),
			*(*int32)(is[2]), *(*int32)(is[3]))
	})
	var des uptr
	des = cc.Cbk(func(_ uptr) { cc.CbkClose(df); cc.CbkCloseLate(des) })

	gtk_drawing_area_set_draw_func.Fn()(d, df, nil, des)
}
func (d *DrawingArea) ConnectResize(fn func(da *DrawingArea, w, h int32)) uint64 {
	return d.SignalConnect("resize", fn, nil)
}

// #endregion

// #region DropControllerMotion

type DropControllerMotion struct{ EventController }

func GTypeDropControllerMotion() gobject.GType { return gtk_drop_controller_motion_get_type.Fn()() }
func NewDropControllerMotion() *DropControllerMotion {
	return gtk_drop_controller_motion_new.Fn()()
}
func (c *DropControllerMotion) ContainsPointer() bool {
	return gtk_drop_controller_motion_contains_pointer.Fn()(c) != 0
}
func (c *DropControllerMotion) GetDrop() *gdk.Drop {
	return gtk_drop_controller_motion_get_drop.Fn()(c)
}
func (c *DropControllerMotion) IsPointer() bool {
	return gtk_drop_controller_motion_is_pointer.Fn()(c) != 0
}
func (c *DropControllerMotion) ConnectEnter(
	sig func(c *DropControllerMotion, x, y float64)) uint64 {
	return c.SignalConnect("enter", sig, nil)
}
func (c *DropControllerMotion) ConnectLeave(
	sig func(c *DropControllerMotion)) uint64 {
	return c.SignalConnect("leave", sig, nil)
}
func (c *DropControllerMotion) ConnectMotion(
	sig func(c *DropControllerMotion, x, y float64)) uint64 {
	return c.SignalConnect("motion", sig, nil)
}

// #endregion

// #region DropDown

type DropDown struct{ Widget }

func GTypeDropDown() gobject.GType { return gtk_drop_down_get_type.Fn()() }
func NewDropDown(model gio.GListModelIface, expression ExpressionIface) *DropDown {
	return gtk_drop_down_new.Fn()(gio.GetGListModelIface(model), GetExpressionIface(expression))
}
func NewDropDownFromStrings(strings []string) *DropDown {
	cpp := cc.NewStrings(strings)
	defer cpp.Free()
	return gtk_drop_down_new_from_strings.Fn()(cpp)
}
func (dd *DropDown) SetModel(model gio.GListModelIface) {
	gtk_drop_down_set_model.Fn()(dd, gio.GetGListModelIface(model))
}
func (dd *DropDown) GetModel() *gio.GListModel {
	return gtk_drop_down_get_model.Fn()(dd)
}
func (dd *DropDown) SetSelected(position uint32) {
	gtk_drop_down_set_selected.Fn()(dd, position)
}
func (dd *DropDown) GetSelected() uint32 {
	return gtk_drop_down_get_selected.Fn()(dd)
}
func (dd *DropDown) GetSelectedItem() uptr {
	return gtk_drop_down_get_selected_item.Fn()(dd)
}
func (dd *DropDown) SetFactory(factory *ListItemFactory) {
	gtk_drop_down_set_factory.Fn()(dd, factory)
}
func (dd *DropDown) GetFactory() *ListItemFactory {
	return gtk_drop_down_get_factory.Fn()(dd)
}
func (dd *DropDown) SetListFactory(factory *ListItemFactory) {
	gtk_drop_down_set_list_factory.Fn()(dd, factory)
}
func (dd *DropDown) GetListFactory() *ListItemFactory {
	return gtk_drop_down_get_list_factory.Fn()(dd)
}
func (dd *DropDown) SetHeaderFactory(factory *ListItemFactory) {
	gtk_drop_down_set_header_factory.Fn()(dd, factory)
}
func (dd *DropDown) GetHeaderFactory() *ListItemFactory {
	return gtk_drop_down_get_header_factory.Fn()(dd)
}
func (dd *DropDown) SetExpression(expression ExpressionIface) {
	gtk_drop_down_set_expression.Fn()(dd, GetExpressionIface(expression))
}
func (dd *DropDown) GetExpression() *Expression {
	return gtk_drop_down_get_expression.Fn()(dd)
}
func (dd *DropDown) SetEnableSearch(enableSearch bool) {
	gtk_drop_down_set_enable_search.Fn()(dd, cbool(enableSearch))
}
func (dd *DropDown) GetEnableSearch() bool {
	return gtk_drop_down_get_enable_search.Fn()(dd) != 0
}
func (dd *DropDown) SetShowArrow(showArrow bool) {
	gtk_drop_down_set_show_arrow.Fn()(dd, cbool(showArrow))
}
func (dd *DropDown) GetShowArrow() bool {
	return gtk_drop_down_get_show_arrow.Fn()(dd) != 0
}
func (dd *DropDown) SetSearchMatchMode(matchMode StringFilterMatchMode) {
	gtk_drop_down_set_search_match_mode.Fn()(dd, matchMode)
}
func (dd *DropDown) GetSearchMatchMode() StringFilterMatchMode {
	return gtk_drop_down_get_search_match_mode.Fn()(dd)
}
func (dd *DropDown) ConnectActivate(fn func(dd *DropDown)) uint64 {
	return dd.SignalConnect("activate", fn, nil)
}

// #endregion

// #region DropTarget

type DropTarget struct{ EventController }

func GTypeDropTarget() gobject.GType { return gtk_drop_target_get_type.Fn()() }
func NewDropTarget(gtype gobject.GType, actions gdk.DragAction) *DropTarget {
	return gtk_drop_target_new.Fn()(gtype, actions)
}
func (t *DropTarget) SetGTypes(types []gobject.GType) {
	gtk_drop_target_set_gtypes.Fn()(t, carry(types), uint64(len(types)))
}
func (t *DropTarget) GetGTypes() (types []gobject.GType) {
	var nTypes uint64
	cTypes := gtk_drop_target_get_gtypes.Fn()(t, &nTypes)
	if cTypes != nil && nTypes > 0 {
		types = (*[1 << 30]gobject.GType)(unsafe.Pointer(cTypes))[:nTypes:nTypes]
	}
	return
}
func (t *DropTarget) GetFormats() *gdk.ContentFormats {
	return gtk_drop_target_get_formats.Fn()(t)
}

func (t *DropTarget) SetActions(actions gdk.DragAction) { gtk_drop_target_set_actions.Fn()(t, actions) }
func (t *DropTarget) GetActions() gdk.DragAction        { return gtk_drop_target_get_actions.Fn()(t) }
func (t *DropTarget) SetPreload(preload bool)           { gtk_drop_target_set_preload.Fn()(t, cbool(preload)) }
func (t *DropTarget) GetPreload() bool                  { return gtk_drop_target_get_preload.Fn()(t) != 0 }
func (t *DropTarget) GetCurrentDrop() *gdk.Drop         { return gtk_drop_target_get_current_drop.Fn()(t) }
func (t *DropTarget) GetValue() *gobject.GValue         { return gtk_drop_target_get_value.Fn()(t) }
func (t *DropTarget) Reject()                           { gtk_drop_target_reject.Fn()(t) }
func (c *DropTarget) ConnectAccept(sig func(c *DropTarget, drop *gdk.Drop, _ uptr) (accepted bool)) uint64 {
	return c.SignalConnect("accept", sig, nil)
}
func (c *DropTarget) ConnectDrop(sig func(c *DropTarget, val *gobject.GValue, x, y float64, _ uptr) (droped bool)) uint64 {
	return c.SignalConnect("drop", sig, nil)
}
func (c *DropTarget) ConnectEnter(sig func(c *DropTarget, x, y float64, _ uptr) gdk.DragAction) uint64 {
	return c.SignalConnect("enter", sig, nil)
}
func (c *DropTarget) ConnectLeave(sig func(c *DropTarget)) uint64 {
	return c.SignalConnect("leave", sig, nil)
}
func (c *DropTarget) ConnectMotion(sig func(c *DropTarget, x, y float64, _ uptr) gdk.DragAction) uint64 {
	return c.SignalConnect("motion", sig, nil)
}

// #endregion

// #region DropTargetAsync

type DropTargetAsync struct{ EventController }

func GTypeDropTargetAsync() gobject.GType { return gtk_drop_target_async_get_type.Fn()() }
func NewDropTargetAsync(formats *gdk.ContentFormats, actions gdk.DragAction) *DropTargetAsync {
	return gtk_drop_target_async_new.Fn()(formats, actions)
}
func (t *DropTargetAsync) SetFormats(formats *gdk.ContentFormats) {
	gtk_drop_target_async_set_formats.Fn()(t, formats)
}
func (t *DropTargetAsync) GetFormats() *gdk.ContentFormats {
	return gtk_drop_target_async_get_formats.Fn()(t)
}
func (t *DropTargetAsync) SetActions(actions gdk.DragAction) {
	gtk_drop_target_async_set_actions.Fn()(t, actions)
}
func (t *DropTargetAsync) GetActions() gdk.DragAction {
	return gtk_drop_target_async_get_actions.Fn()(t)
}
func (t *DropTargetAsync) RejectDrop(drop *gdk.Drop) {
	gtk_drop_target_async_reject_drop.Fn()(t, drop)
}
func (c *DropTargetAsync) ConnectAccept(sig func(c *DropTargetAsync, drop *gdk.Drop, _ uptr) (accepted bool)) uint64 {
	return c.SignalConnect("accept", sig, nil)
}
func (c *DropTargetAsync) ConnectDragEnter(sig func(c *DropTargetAsync, drop gdk.Drop, x, y float64, _ uptr) gdk.DragAction) uint64 {
	return c.SignalConnect("drag-enter", sig, nil)
}
func (c *DropTargetAsync) ConnectDragLeave(sig func(c *DropTargetAsync, drop gdk.Drop)) uint64 {
	return c.SignalConnect("drag-leave", sig, nil)
}
func (c *DropTargetAsync) ConnectDragMotion(sig func(c *DropTargetAsync, drop gdk.Drop, x, y float64, _ uptr) gdk.DragAction) uint64 {
	return c.SignalConnect("drag-motion", sig, nil)
}
func (c *DropTargetAsync) ConnectDrop(sig func(c *DropTargetAsync, drop gdk.Drop, x, y float64, _ uptr) (accepted bool)) uint64 {
	return c.SignalConnect("drop", sig, nil)
}

// #endregion

// #region Editable

type EditableIfaceObj struct {
	GIface gobject.GTypeInterface

	/* signals */

	InsertText cc.Func // void (* insert_text)(GtkEditable *editable,const char  *text,int length,int *position);
	DeleteText cc.Func // void (* delete_text)(GtkEditable *editable,int start_pos,int end_pos);
	Changed    cc.Func // void (* changed)(GtkEditable *editable);

	/* vtable */

	GetText            cc.Func //const char * (* get_text)(GtkEditable *editable);
	DoInsertText       cc.Func //void(* do_insert_text)(GtkEditable* editable,const char *text,int length,int *position);
	DoDeleteText       cc.Func //void(* do_delete_text)(GtkEditable* editable,int start_pos,int end_pos);
	GetSelectionBounds cc.Func // gboolean (* get_selection_bounds) (GtkEditable *editable,int *start_pos,int *end_pos);
	SetSelectionBounds cc.Func // void     (* set_selection_bounds) (GtkEditable *editable,int start_pos,int end_pos);
	GetDelegate        cc.Func // GtkEditable * (* get_delegate) (GtkEditable *editable);
}

type EditableIface interface {
	GetEditableIface() *Editable
}

func GetEditableIface(iface EditableIface) *Editable {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetEditableIface()
}

func (e *Editable) GetEditableIface() *Editable { return e }

type Editable struct{}

func GTypeEditable() gobject.GType  { return gtk_editable_get_type.Fn()() }
func (e *Editable) GetText() string { return gtk_editable_get_text.Fn()(e).String() }
func (e *Editable) SetText(text string) {
	t := cc.NewString(text)
	defer t.Free()
	gtk_editable_set_text.Fn()(e, t)
}
func (e *Editable) GetChars(startPos, endPos int32) string {
	return gtk_editable_get_chars.Fn()(e, startPos, endPos).TakeString()
}
func (e *Editable) InsertText(text string, length int32, pos int32) (new_pos int32) {
	t := cc.NewString(text)
	defer t.Free()
	new_pos = pos
	gtk_editable_insert_text.Fn()(e, t, length, &new_pos)
	return
}
func (e *Editable) DeleteText(startPos, endPos int32) {
	gtk_editable_delete_text.Fn()(e, startPos, endPos)
}
func (e *Editable) GetSelectionBounds() (startPos, endPos int32, ok bool) {
	ok = gtk_editable_get_selection_bounds.Fn()(e, &startPos, &endPos) != 0
	return
}
func (e *Editable) DeleteSelection() { gtk_editable_delete_selection.Fn()(e) }
func (e *Editable) SelectRegion(startPos, endPos int32) {
	gtk_editable_select_region.Fn()(e, startPos, endPos)
}
func (e *Editable) SetPosition(position int32)    { gtk_editable_set_position.Fn()(e, position) }
func (e *Editable) GetPosition() int32            { return gtk_editable_get_position.Fn()(e) }
func (e *Editable) GetEditable() bool             { return gtk_editable_get_editable.Fn()(e) != 0 }
func (e *Editable) SetEditable(isEditable bool)   { gtk_editable_set_editable.Fn()(e, cbool(isEditable)) }
func (e *Editable) GetAlignment() float32         { return gtk_editable_get_alignment.Fn()(e) }
func (e *Editable) SetAlignment(xalign float32)   { gtk_editable_set_alignment.Fn()(e, xalign) }
func (e *Editable) GetWidthChars() int32          { return gtk_editable_get_width_chars.Fn()(e) }
func (e *Editable) SetWidthChars(nChars int32)    { gtk_editable_set_width_chars.Fn()(e, nChars) }
func (e *Editable) GetMaxWidthChars() int32       { return gtk_editable_get_max_width_chars.Fn()(e) }
func (e *Editable) SetMaxWidthChars(nChars int32) { gtk_editable_set_max_width_chars.Fn()(e, nChars) }
func (e *Editable) GetEnableUndo() bool           { return gtk_editable_get_enable_undo.Fn()(e) != 0 }
func (e *Editable) SetEnableUndo(enableUndo bool) {
	gtk_editable_set_enable_undo.Fn()(e, cbool(enableUndo))
}
func (e *Editable) ConnectChanged(sig func(e *Editable)) uint64 {
	return toGobj(e).SignalConnect("changed", sig, nil)
}
func (e *Editable) ConnectDeleteText(sig func(e *Editable, startPos, endPos int32)) uint64 {
	return toGobj(e).SignalConnect("delete-text", sig, nil)
}
func (e *Editable) ConnectInsertText(sig func(e *Editable, new_text cc.String, new_textlength int32, pos *int32)) uint64 {
	return toGobj(e).SignalConnect("insert-text", sig, nil)
}

func EditableInstallProperties(eb *gobject.GObjectClass, firstProp uint32) uint32 {
	return gtk_editable_install_properties.Fn()(eb, firstProp)
}

func (e *Editable) GetDelegate() *Editable { return gtk_editable_get_delegate.Fn()(e) }
func (e *Editable) InitDelegate()          { gtk_editable_init_delegate.Fn()(e) }
func (e *Editable) FinishDelegate()        { gtk_editable_finish_delegate.Fn()(e) }

func EditableDelegateSetProperty(obj GobjIface, propId uint32,
	value *gobject.GValue, pspec *gobject.GParamSpec) bool {
	return gtk_editable_delegate_set_property.Fn()(
		gobjGet(obj), propId, value, pspec) != 0
}
func EditableDelegateGetProperty(obj GobjIface, propId uint32,
	value *gobject.GValue, pspec *gobject.GParamSpec) bool {
	return gtk_editable_delegate_get_property.Fn()(
		gobjGet(obj), propId, value, pspec) != 0
}
func (e *Editable) DelegateGetAccessiblePlatformState(state AccessiblePlatformState) bool {
	return gtk_editable_delegate_get_accessible_platform_state.Fn()(e, state) != 0
}

// #endregion

// #region EditableLabel

type EditableLabel struct {
	Widget
	Editable
}

func GTypeEditableLabel() gobject.GType { return g_editable_label_get_type.Fn()() }
func NewEditableLabel(str string) *EditableLabel {
	cstr := cc.NewString(str)
	defer cstr.Free()
	return g_editable_label_new.Fn()(cstr)
}
func (label *EditableLabel) GetEditing() bool { return g_editable_label_get_editing.Fn()(label) != 0 }
func (label *EditableLabel) StartEditing()    { g_editable_label_start_editing.Fn()(label) }
func (label *EditableLabel) StopEditing(commit bool) {
	g_editable_label_stop_editing.Fn()(label, cbool(commit))
}

// #endregion

// #region EmojiChooser

type EmojiChooser struct {
	Popover
	Native
	ShortcutManager
}

func GTypeEmojiChooser() gobject.GType { return gtk_emoji_chooser_get_type.Fn()() }
func NewEmojiChooser() *EmojiChooser   { return gtk_emoji_chooser_new.Fn()() }
func (ec *EmojiChooser) ConnectEmojiPicked(sig func(ec *EmojiChooser, emoji string)) uint64 {
	return ec.SignalConnect("emoji-picked", sig, nil)
}

// #endregion

// #region Entry

type EntryObj struct {
	Parent WidgetObj
}
type EntryClass struct {
	Parent WidgetClass

	Activate cc.Func // void (* activate)(GtkEntry *entry);
	_        [8]uptr // padding
}

type Entry struct {
	Widget
	Editable
}

func GTypeEntry() gobject.GType { return gtk_entry_get_type.Fn()() }
func NewEntry() *Entry          { return gtk_entry_new.Fn()() }
func NewEntryWithBuffer(buffer EntryBufferIface) *Entry {
	return gtk_entry_new_with_buffer.Fn()(GetEntryBufferIface(buffer))
}
func (e *Entry) GetBuffer() *EntryBuffer       { return gtk_entry_get_buffer.Fn()(e) }
func (e *Entry) SetBuffer(buffer *EntryBuffer) { gtk_entry_set_buffer.Fn()(e, buffer) }
func (e *Entry) SetVisibility(visible bool)    { gtk_entry_set_visibility.Fn()(e, cbool(visible)) }
func (e *Entry) GetVisibility() bool           { return gtk_entry_get_visibility.Fn()(e) != 0 }
func (e *Entry) SetInvisibleChar(ch uint32)    { gtk_entry_set_invisible_char.Fn()(e, ch) }
func (e *Entry) GetInvisibleChar() uint32      { return gtk_entry_get_invisible_char.Fn()(e) }
func (e *Entry) UnsetInvisibleChar()           { gtk_entry_unset_invisible_char.Fn()(e) }
func (e *Entry) SetHasFrame(setting bool)      { gtk_entry_set_has_frame.Fn()(e, cbool(setting)) }
func (e *Entry) GetHasFrame() bool             { return gtk_entry_get_has_frame.Fn()(e) != 0 }
func (e *Entry) SetOverwriteMode(overwrite bool) {
	gtk_entry_set_overwrite_mode.Fn()(e, cbool(overwrite))
}
func (e *Entry) GetOverwriteMode() bool { return gtk_entry_get_overwrite_mode.Fn()(e) != 0 }
func (e *Entry) SetMaxLength(max int32) { gtk_entry_set_max_length.Fn()(e, max) }
func (e *Entry) GetMaxLength() int32    { return gtk_entry_get_max_length.Fn()(e) }
func (e *Entry) GetTextLength() uint16  { return gtk_entry_get_text_length.Fn()(e) }
func (e *Entry) SetActivatesDefault(setting bool) {
	gtk_entry_set_activates_default.Fn()(e, cbool(setting))
}
func (e *Entry) GetActivatesDefault() bool   { return gtk_entry_get_activates_default.Fn()(e) != 0 }
func (e *Entry) SetAlignment(xalign float32) { gtk_entry_set_alignment.Fn()(e, xalign) }
func (e *Entry) GetAlignment() float32       { return gtk_entry_get_alignment.Fn()(e) }
func (e *Entry) SetProgressFraction(fraction float64) {
	gtk_entry_set_progress_fraction.Fn()(e, fraction)
}
func (e *Entry) GetProgressFraction() float64 { return gtk_entry_get_progress_fraction.Fn()(e) }
func (e *Entry) SetProgressPulseStep(fraction float64) {
	gtk_entry_set_progress_pulse_step.Fn()(e, fraction)
}
func (e *Entry) GetProgressPulseStep() float64 { return gtk_entry_get_progress_pulse_step.Fn()(e) }
func (e *Entry) ProgressPulse()                { gtk_entry_progress_pulse.Fn()(e) }
func (e *Entry) GetPlaceholderText() string    { return gtk_entry_get_placeholder_text.Fn()(e).String() }
func (e *Entry) SetPlaceholderText(text string) {
	cText := cc.NewString(text)
	defer cText.Free()
	gtk_entry_set_placeholder_text.Fn()(e, cText)
}
func (e *Entry) SetIconFromPaintable(iconPos EntryIconPosition, paintable *gdk.Paintable) {
	gtk_entry_set_icon_from_paintable.Fn()(e, iconPos, paintable)
}
func (e *Entry) SetIconFromIconName(iconPos EntryIconPosition, iconName string) {
	cIconName := cc.NewString(iconName)
	defer cIconName.Free()
	gtk_entry_set_icon_from_icon_name.Fn()(e, iconPos, cIconName)
}
func (e *Entry) SetIconFromGIcon(iconPos EntryIconPosition, icon *gio.GIcon) {
	gtk_entry_set_icon_from_gicon.Fn()(e, iconPos, icon)
}
func (e *Entry) GetIconStorageType(iconPos EntryIconPosition) ImageType {
	return gtk_entry_get_icon_storage_type.Fn()(e, iconPos)
}
func (e *Entry) GetIconPaintable(iconPos EntryIconPosition) *gdk.Paintable {
	return gtk_entry_get_icon_paintable.Fn()(e, iconPos)
}
func (e *Entry) GetIconName(iconPos EntryIconPosition) string {
	return gtk_entry_get_icon_name.Fn()(e, iconPos).String()
}
func (e *Entry) GetIconGIcon(iconPos EntryIconPosition) *gio.GIcon {
	return gtk_entry_get_icon_gicon.Fn()(e, iconPos)
}
func (e *Entry) SetIconActivatable(iconPos EntryIconPosition, activatable bool) {
	gtk_entry_set_icon_activatable.Fn()(e, iconPos, cbool(activatable))
}
func (e *Entry) GetIconActivatable(iconPos EntryIconPosition) bool {
	return gtk_entry_get_icon_activatable.Fn()(e, iconPos) != 0
}
func (e *Entry) SetIconSensitive(iconPos EntryIconPosition, sensitive bool) {
	gtk_entry_set_icon_sensitive.Fn()(e, iconPos, cbool(sensitive))
}
func (e *Entry) GetIconSensitive(iconPos EntryIconPosition) bool {
	return gtk_entry_get_icon_sensitive.Fn()(e, iconPos) != 0
}
func (e *Entry) GetIconAtPos(x, y int32) int32 {
	return gtk_entry_get_icon_at_pos.Fn()(e, x, y)
}
func (e *Entry) SetIconTooltipText(iconPos EntryIconPosition, tooltip string) {
	cTooltip := cc.NewString(tooltip)
	defer cTooltip.Free()
	gtk_entry_set_icon_tooltip_text.Fn()(e, iconPos, cTooltip)
}
func (e *Entry) GetIconTooltipText(iconPos EntryIconPosition) string {
	return gtk_entry_get_icon_tooltip_text.Fn()(e, iconPos).TakeString()
}
func (e *Entry) SetIconTooltipMarkup(iconPos EntryIconPosition, tooltip string) {
	cTooltip := cc.NewString(tooltip)
	defer cTooltip.Free()
	gtk_entry_set_icon_tooltip_markup.Fn()(e, iconPos, cTooltip)
}
func (e *Entry) GetIconTooltipMarkup(iconPos EntryIconPosition) string {
	return gtk_entry_get_icon_tooltip_markup.Fn()(e, iconPos).TakeString()
}
func (e *Entry) SetIconDragSource(iconPos EntryIconPosition, provider gdk.ContentProviderIface, actions gdk.DragAction) {
	gtk_entry_set_icon_drag_source.Fn()(e, iconPos, gdk.GetContentProviderIface(provider), actions)
}
func (e *Entry) GetCurrentIconDragSource() int32 {
	return gtk_entry_get_current_icon_drag_source.Fn()(e)
}
func (e *Entry) GetIconArea(iconPos EntryIconPosition, iconArea *gdk.Rectangle) {
	gtk_entry_get_icon_area.Fn()(e, iconPos, iconArea)
}
func (e *Entry) ResetIMContext() {
	gtk_entry_reset_im_context.Fn()(e)
}
func (e *Entry) SetInputPurpose(purpose InputPurpose) {
	gtk_entry_set_input_purpose.Fn()(e, purpose)
}
func (e *Entry) GetInputPurpose() InputPurpose {
	return gtk_entry_get_input_purpose.Fn()(e)
}
func (e *Entry) SetInputHints(hints InputHints) {
	gtk_entry_set_input_hints.Fn()(e, hints)
}
func (e *Entry) SetTabs(tabs *pango.TabArray) {
	gtk_entry_set_tabs.Fn()(e, tabs)
}
func (e *Entry) GetTabs() *pango.TabArray {
	return gtk_entry_get_tabs.Fn()(e)
}
func (e *Entry) GrabFocusWithoutSelecting() bool {
	return gtk_entry_grab_focus_without_selecting.Fn()(e) != 0
}
func (e *Entry) SetExtraMenu(model *gio.GMenuModel) {
	gtk_entry_set_extra_menu.Fn()(e, model)
}
func (e *Entry) GetExtraMenu() *gio.GMenuModel {
	return gtk_entry_get_extra_menu.Fn()(e)
}
func (e *Entry) ConnectActivate(f func(*Entry)) uint64 {
	return e.SignalConnect("activate", f, nil)
}
func (e *Entry) ConnectIconPress(f func(*Entry, EntryIconPosition)) uint64 {
	return e.SignalConnect("icon-press", f, nil)
}
func (e *Entry) ConnectIconRelease(f func(*Entry, EntryIconPosition)) uint64 {
	return e.SignalConnect("icon-release", f, nil)
}

// #endregion

// #region EntryBuffer

type EntryBufferObj struct {
	Parent gobject.GObjectObj
}
type EntryBufferClass struct {
	Parent gobject.GObjectClass

	/* Signals */
	InsertedText cc.Func //void (* inserted_text) (GtkEntryBuffer *buffer,guint pos,const char* chars,guint n_chars)
	DeletedText  cc.Func //void (* deleted_text) (GtkEntryBuffer *buffer,guint pos,n_chars guint)

	/* Virtual Methods */
	GetText    cc.Func //const char* (* get_text) (GtkEntryBuffer *buffer,gsize n_bytes);
	GetLength  cc.Func //guint (* get_length) (GtkEntryBuffer *buffer);
	InsertText cc.Func //guint (* insert_text) (GtkEntryBuffer *buffer,guint pos,const char* chars,guint n_chars);
	DeleteText cc.Func //guint (* delete_text) (GtkEntryBuffer *buffer,guint pos,guint n_chars);

	_ [8]cc.Func //padding
}

type EntryBufferIface interface {
	GetEntryBufferIface() *EntryBuffer
}

func GetEntryBufferIface(iface EntryBufferIface) *EntryBuffer {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetEntryBufferIface()
}

func (c *EntryBuffer) GetEntryBufferIface() *EntryBuffer { return c }

type EntryBuffer struct{ gobjCore }

func GTypeEntryBuffer() gobject.GType { return gtk_entry_buffer_get_type.Fn()() }
func NewEntryBuffer(initialChars string, nInitialChars int32) *EntryBuffer {
	cChars := cc.NewString(initialChars)
	defer cChars.Free()
	return gtk_entry_buffer_new.Fn()(cChars, nInitialChars)
}
func (b *EntryBuffer) GetBytes() uint64  { return gtk_entry_buffer_get_bytes.Fn()(b) }
func (b *EntryBuffer) GetLength() uint32 { return gtk_entry_buffer_get_length.Fn()(b) }
func (b *EntryBuffer) GetText() string   { return gtk_entry_buffer_get_text.Fn()(b).String() }
func (b *EntryBuffer) SetText(chars string, nChars int32) {
	cChars := cc.NewString(chars)
	defer cChars.Free()
	gtk_entry_buffer_set_text.Fn()(b, cChars, nChars)
}
func (b *EntryBuffer) SetMaxLength(maxLength int32) {
	gtk_entry_buffer_set_max_length.Fn()(b, maxLength)
}
func (b *EntryBuffer) GetMaxLength() int32 {
	return gtk_entry_buffer_get_max_length.Fn()(b)
}
func (b *EntryBuffer) InsertText(position uint32, chars string, nChars int32) uint32 {
	cChars := cc.NewString(chars)
	defer cChars.Free()
	return gtk_entry_buffer_insert_text.Fn()(b, position, cChars, nChars)
}
func (b *EntryBuffer) DeleteText(position uint32, nChars int32) uint32 {
	return gtk_entry_buffer_delete_text.Fn()(b, position, nChars)
}
func (b *EntryBuffer) EmitInsertedText(position uint32, chars string, nChars uint32) {
	cChars := cc.NewString(chars)
	defer cChars.Free()
	gtk_entry_buffer_emit_inserted_text.Fn()(b, position, cChars, nChars)
}
func (b *EntryBuffer) EmitDeletedText(position, nChars uint32) {
	gtk_entry_buffer_emit_deleted_text.Fn()(b, position, nChars)
}
func (b *EntryBuffer) ConnectInsertedText(
	sig func(b *EntryBuffer, position uint32, chars cc.String, nChars uint32)) uint64 {
	return b.SignalConnect("inserted-text", sig, nil)
}
func (b *EntryBuffer) ConnectDeletedText(
	sig func(b *EntryBuffer, position, nChars uint32)) uint64 {
	return b.SignalConnect("deleted-text", sig, nil)
}

// #endregion

// #region EventController

type EventControllerIface interface {
	GetEventControllerIface() *EventController
}

func GetEventControllerIface(iface EventControllerIface) *EventController {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetEventControllerIface()
}
func (c *EventController) GetEventControllerIface() *EventController { return c }

type EventController struct{ gobjCore }

func GTypeEventController() gobject.GType     { return gtk_event_controller_get_type.Fn()() }
func (c *EventController) GetWidget() *Widget { return gtk_event_controller_get_widget.Fn()(c) }
func (c *EventController) Reset()             { gtk_event_controller_reset.Fn()(c) }
func (c *EventController) GetPropagationPhase() PropagationPhase {
	return gtk_event_controller_get_propagation_phase.Fn()(c)
}
func (c *EventController) SetPropagationPhase(phase PropagationPhase) {
	gtk_event_controller_set_propagation_phase.Fn()(c, phase)
}
func (c *EventController) GetPropagationLimit() PropagationLimit {
	return gtk_event_controller_get_propagation_limit.Fn()(c)
}
func (c *EventController) SetPropagationLimit(limit PropagationLimit) {
	gtk_event_controller_set_propagation_limit.Fn()(c, limit)
}
func (c *EventController) GetName() string { return gtk_event_controller_get_name.Fn()(c).String() }
func (c *EventController) SetName(name string) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_event_controller_set_name.Fn()(c, cName)
}
func (c *EventController) SetStaticName(name string) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_event_controller_set_static_name.Fn()(c, cName)
}
func (c *EventController) GetCurrentEvent() *gdk.Event {
	return gtk_event_controller_get_current_event.Fn()(c)
}
func (c *EventController) GetCurrentEventTime() uint32 {
	return gtk_event_controller_get_current_event_time.Fn()(c)
}
func (c *EventController) GetCurrentEventDevice() *gdk.Device {
	return gtk_event_controller_get_current_event_device.Fn()(c)
}
func (c *EventController) GetCurrentEventState() gdk.ModifierType {
	return gtk_event_controller_get_current_event_state.Fn()(c)
}

// #endregion

// #region EventControllerFocus

type EventControllerFocus struct{ EventController }

func GTypeEventControllerFocus() gobject.GType { return gtk_event_controller_focus_get_type.Fn()() }
func NewEventControllerFocus() *EventControllerFocus {
	return (*EventControllerFocus)(gtk_event_controller_focus_new.Fn()())
}
func (c *EventControllerFocus) ContainsFocus() bool {
	return gtk_event_controller_focus_contains_focus.Fn()(c) != 0
}
func (c *EventControllerFocus) IsFocus() bool {
	return gtk_event_controller_focus_is_focus.Fn()(c) != 0
}
func (c *EventControllerFocus) ConnectEnter(sig func(c *EventControllerFocus)) uint64 {
	return c.SignalConnect("enter", sig, nil)
}
func (c *EventControllerFocus) ConnectLeave(sig func(c *EventControllerFocus)) uint64 {
	return c.SignalConnect("leave", sig, nil)
}

// #endregion

// #region EventControllerKey

type EventControllerKey struct{ EventController }

func GTypeEventControllerKey() gobject.GType     { return gtk_event_controller_key_get_type.Fn()() }
func NewEventControllerKey() *EventControllerKey { return gtk_event_controller_key_new.Fn()() }
func (c *EventControllerKey) SetIMContext(imContext *IMContext) {
	gtk_event_controller_key_set_im_context.Fn()(c, imContext)
}
func (c *EventControllerKey) GetIMContext() *IMContext {
	return gtk_event_controller_key_get_im_context.Fn()(c)
}
func (c *EventControllerKey) Forward(widget WidgetIface) bool {
	return gtk_event_controller_key_forward.Fn()(c, GetWidgetIface(widget)) != 0
}
func (c *EventControllerKey) GetGroup() uint32 { return gtk_event_controller_key_get_group.Fn()(c) }

func (c *EventControllerKey) ConnectImUpdate(sig func(c *EventControllerKey)) uint64 {
	return c.SignalConnect("im-update", sig, nil)
}
func (c *EventControllerKey) ConnectKeyPressed(
	sig func(c *EventControllerKey, keyVal, keyCode uint32, state gdk.ModifierType) (ok bool)) uint64 {
	return c.SignalConnect("key-pressed", sig, nil)
}
func (c *EventControllerKey) ConnectKeyReleased(
	sig func(c *EventControllerKey, keyVal, keyCode uint32, state gdk.ModifierType)) uint64 {
	return c.SignalConnect("key-released", sig, nil)
}
func (c *EventControllerKey) ConnectKeyModifiers(
	sig func(c *EventControllerKey, state gdk.ModifierType)) uint64 {
	return c.SignalConnect("modifiers", sig, nil)
}

// #endregion

// #region EventControllerLegacy

type EventControllerLegacy struct{ EventController }

func GTypeEventControllerLegacy() gobject.GType {
	return gtk_event_controller_legacy_get_type.Fn()()
}
func NewEventControllerLegacy() *EventControllerLegacy {
	return gtk_event_controller_legacy_new.Fn()()
}
func (c *EventControllerLegacy) ConnectEvent(
	sig func(c *EventControllerLegacy, event *gdk.Event)) uint64 {
	return c.SignalConnect("event", sig, nil)
}

// #endregion

// #region EventControllerMotion

type EventControllerMotion struct{ EventController }

func GTypeEventControllerMotion() gobject.GType { return gtk_event_controller_motion_get_type.Fn()() }
func NewEventControllerMotion() *EventControllerMotion {
	return gtk_event_controller_motion_new.Fn()()
}
func (c *EventControllerMotion) ContainsPointer() bool {
	return gtk_event_controller_motion_contains_pointer.Fn()(c) != 0
}
func (c *EventControllerMotion) IsPointer() bool {
	return gtk_event_controller_motion_is_pointer.Fn()(c) != 0
}
func (c *EventControllerMotion) ConnectEnter(sig func(c *EventControllerMotion, x, y float64)) uint64 {
	return c.SignalConnect("enter", sig, nil)
}
func (c *EventControllerMotion) ConnectLeave(sig func(c *EventControllerMotion)) uint64 {
	return c.SignalConnect("leave", sig, nil)
}
func (c *EventControllerMotion) ConnectMotion(sig func(c *EventControllerMotion, x, y float64)) uint64 {
	return c.SignalConnect("motion", sig, nil)
}

// #endregion

// #region EventControllerScroll

type EventControllerScroll struct{ EventController }

func GTypeEventControllerScroll() gobject.GType { return gtk_event_controller_scroll_get_type.Fn()() }
func NewEventControllerScroll(flags EventControllerScrollFlags) *EventControllerScroll {
	return gtk_event_controller_scroll_new.Fn()(flags)
}
func (c *EventControllerScroll) SetFlags(flags EventControllerScrollFlags) {
	gtk_event_controller_scroll_set_flags.Fn()(c, flags)
}
func (c *EventControllerScroll) GetFlags() EventControllerScrollFlags {
	return gtk_event_controller_scroll_get_flags.Fn()(c)
}
func (c *EventControllerScroll) GetUnit() gdk.ScrollUnit {
	return gtk_event_controller_scroll_get_unit.Fn()(c)
}
func (c *EventControllerScroll) ConnectDecelerate(
	sig func(c *EventControllerScroll, velX, velY float64)) uint64 {
	return c.SignalConnect("decelerate", sig, nil)
}
func (c *EventControllerScroll) ConnectScroll(
	sig func(c *EventControllerScroll, dx, dy float64) (ok bool)) uint64 {
	return c.SignalConnect("scroll", sig, nil)
}
func (c *EventControllerScroll) ConnectScrollBegin(sig func(c *EventControllerScroll)) uint64 {
	return c.SignalConnect("scroll-begin", sig, nil)
}
func (c *EventControllerScroll) ConnectScrollEnd(sig func(c *EventControllerScroll)) uint64 {
	return c.SignalConnect("scroll-end", sig, nil)
}

// #endregion

// #region Expander

type Expander struct{ Widget }

func GTypeExpander() gobject.GType { return gtk_expander_get_type.Fn()() }
func NewExpander(label string) *Expander {
	cLabel := cc.NewString(label)
	defer cLabel.Free()
	return gtk_expander_new.Fn()(cLabel)
}
func NewExpanderWithMnemonic(label string) *Expander {
	cLabel := cc.NewString(label)
	defer cLabel.Free()
	return gtk_expander_new_with_mnemonic.Fn()(cLabel)
}
func (e *Expander) SetExpanded(expanded bool) { gtk_expander_set_expanded.Fn()(e, cbool(expanded)) }
func (e *Expander) GetExpanded() bool         { return gtk_expander_get_expanded.Fn()(e) != 0 }
func (e *Expander) SetLabel(label string) {
	cLabel := cc.NewString(label)
	defer cLabel.Free()
	gtk_expander_set_label.Fn()(e, cLabel)
}
func (e *Expander) GetLabel() string { return gtk_expander_get_label.Fn()(e).String() }
func (e *Expander) SetUseUnderline(useUnderline bool) {
	gtk_expander_set_use_underline.Fn()(e, cbool(useUnderline))
}
func (e *Expander) GetUseUnderline() bool {
	return gtk_expander_get_use_underline.Fn()(e) != 0
}
func (e *Expander) SetUseMarkup(useMarkup bool) {
	gtk_expander_set_use_markup.Fn()(e, cbool(useMarkup))
}
func (e *Expander) GetUseMarkup() bool { return gtk_expander_get_use_markup.Fn()(e) != 0 }
func (e *Expander) SetLabelWidget(labelWidget WidgetIface) {
	gtk_expander_set_label_widget.Fn()(e, GetWidgetIface(labelWidget))
}
func (e *Expander) GetLabelWidget() *Widget {
	return gtk_expander_get_label_widget.Fn()(e)
}
func (e *Expander) SetResizeToplevel(resizeToplevel bool) {
	gtk_expander_set_resize_toplevel.Fn()(e, cbool(resizeToplevel))
}
func (e *Expander) GetResizeToplevel() bool    { return gtk_expander_get_resize_toplevel.Fn()(e) != 0 }
func (e *Expander) SetChild(child WidgetIface) { gtk_expander_set_child.Fn()(e, GetWidgetIface(child)) }
func (e *Expander) GetChild() *Widget          { return gtk_expander_get_child.Fn()(e) }
func (e *Expander) ConnectActivate(sig func(e *Expander)) uint64 {
	return e.SignalConnect("activate", sig, nil)
}

// #endregion

// #region Expression

type ExpressionIface interface {
	GetExpressionIface() *Expression
}

func GetExpressionIface(iface ExpressionIface) *Expression {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetExpressionIface()
}
func (e *Expression) GetExpressionIface() *Expression { return e }

type Expression struct{}

func GTypeExpression() gobject.GType              { return gtk_expression_get_type.Fn()() }
func (e *Expression) Ref() *Expression            { return gtk_expression_ref.Fn()(e) }
func (e *Expression) Unref()                      { gtk_expression_unref.Fn()(e) }
func (e *Expression) GetValueType() gobject.GType { return gtk_expression_get_value_type.Fn()(e) }
func (e *Expression) IsStatic() bool              { return gtk_expression_is_static.Fn()(e) != 0 }
func (e *Expression) Evaluate(gobj GobjIface, value *gobject.GValue) bool {
	return gtk_expression_evaluate.Fn()(e, gobjGet(gobj), value) != 0
}
func (e *Expression) Watch(gobj GobjIface, notify func(_ uptr), userDestroy func(_ uptr)) *ExpressionWatch {
	var nf, des uptr
	if notify != nil {
		nf = cc.CbkRaw[func(_ uptr)](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 1)
			notify(*(*uptr)(is[0]))
		})
	}
	if des != nil {
		des = cc.Cbk(func(d uptr) {
			userDestroy(d)
			cc.CbkClose(nf)
			cc.CbkCloseLate(des)
		})
	}

	return gtk_expression_watch.Fn()(e, gobjGet(gobj), nf, nil, des)
}
func (e *Expression) Bind(gobj GobjIface, property string, this GobjIface) *ExpressionWatch {
	cProperty := cc.NewString(property)
	defer cProperty.Free()
	return gtk_expression_bind.Fn()(e, gobjGet(gobj), cProperty, gobjGet(this))
}

type ExpressionWatch struct{}

func GTypeExpressionWatch() gobject.GType        { return gtk_expression_watch_get_type.Fn()() }
func (w *ExpressionWatch) Ref() *ExpressionWatch { return gtk_expression_watch_ref.Fn()(w) }
func (w *ExpressionWatch) Unref()                { gtk_expression_watch_unref.Fn()(w) }
func (w *ExpressionWatch) Evaluate(value *gobject.GValue) bool {
	return gtk_expression_watch_evaluate.Fn()(w, value) != 0
}
func (w *ExpressionWatch) Unwatch() { gtk_expression_watch_unwatch.Fn()(w) }

type PropertyExpression struct{ Expression }

func GTypePropertyExpression() gobject.GType { return gtk_property_expression_get_type.Fn()() }
func NewPropertyExpression(thisType gobject.GType, expression *Expression, propertyName string) *PropertyExpression {
	cPropertyName := cc.NewString(propertyName)
	defer cPropertyName.Free()
	return gtk_property_expression_new.Fn()(thisType, expression, cPropertyName)
}
func NewPropertyExpressionForPspec(expression ExpressionIface, pspec *gobject.GParamSpec) *PropertyExpression {
	return gtk_property_expression_new_for_pspec.Fn()(GetExpressionIface(expression), pspec)
}
func (e *PropertyExpression) GetPropertyExpression() *Expression {
	return gtk_property_expression_get_expression.Fn()(e)
}
func (e *PropertyExpression) GetPropertyPspec() *gobject.GParamSpec {
	return gtk_property_expression_get_pspec.Fn()(e)
}

type ConstantExpression struct{ Expression }

func GTypeConstantExpression() gobject.GType { return gtk_constant_expression_get_type.Fn()() }
func NewConstantExpression(valueType gobject.GType, args ...interface{}) *ConstantExpression {
	return gtk_constant_expression_new.FnVa()(valueType, args...)
}
func NewConstantExpressionForValue(value *gobject.GValue) *ConstantExpression {
	return gtk_constant_expression_new_for_value.Fn()(value)
}
func (e *ConstantExpression) GetConstantValue() *gobject.GValue {
	return gtk_constant_expression_get_value.Fn()(e)
}

type ObjectExpression struct{ Expression }

func GTypeObjectExpression() gobject.GType { return gtk_object_expression_get_type.Fn()() }
func NewObjectExpression(object GobjIface) *ObjectExpression {
	return gtk_object_expression_new.Fn()(gobjGet(object))
}
func (e *ObjectExpression) GetObject() *gobject.GObject {
	return gtk_object_expression_get_object.Fn()(e)
}

// #endregion

// #region FileDialog

type FileDialog struct{ gobjCore }

func GTypeFileDialog() gobject.GType    { return gtk_file_dialog_get_type.Fn()() }
func NewFileDialog() *FileDialog        { return gtk_file_dialog_new.Fn()() }
func (fd *FileDialog) GetTitle() string { return gtk_file_dialog_get_title.Fn()(fd).String() }
func (fd *FileDialog) SetTitle(title string) {
	cTitle := cc.NewString(title)
	defer cTitle.Free()
	gtk_file_dialog_set_title.Fn()(fd, cTitle)
}
func (fd *FileDialog) GetModal() bool              { return gtk_file_dialog_get_modal.Fn()(fd) != 0 }
func (fd *FileDialog) SetModal(modal bool)         { gtk_file_dialog_set_modal.Fn()(fd, cbool(modal)) }
func (fd *FileDialog) GetFilters() *gio.GListModel { return gtk_file_dialog_get_filters.Fn()(fd) }
func (fd *FileDialog) SetFilters(filters gio.GListModelIface) {
	gtk_file_dialog_set_filters.Fn()(fd, gio.GetGListModelIface(filters))
}
func (fd *FileDialog) GetDefaultFilter() *FileFilter {
	return gtk_file_dialog_get_default_filter.Fn()(fd)
}
func (fd *FileDialog) SetDefaultFilter(filter *FileFilter) {
	gtk_file_dialog_set_default_filter.Fn()(fd, filter)
}
func (fd *FileDialog) GetInitialFolder() *gio.GFile {
	return gtk_file_dialog_get_initial_folder.Fn()(fd)
}
func (fd *FileDialog) SetInitialFolder(folder *gio.GFile) {
	gtk_file_dialog_set_initial_folder.Fn()(fd, folder)
}
func (fd *FileDialog) GetInitialName() string {
	return gtk_file_dialog_get_initial_name.Fn()(fd).String()
}
func (fd *FileDialog) SetInitialName(name string) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_file_dialog_set_initial_name.Fn()(fd, cName)
}
func (fd *FileDialog) GetInitialFile() *gio.GFile {
	return gtk_file_dialog_get_initial_file.Fn()(fd)
}
func (fd *FileDialog) SetInitialFile(file *gio.GFile) {
	gtk_file_dialog_set_initial_file.Fn()(fd, file)
}
func (fd *FileDialog) GetAcceptLabel() string {
	return gtk_file_dialog_get_accept_label.Fn()(fd).String()
}
func (fd *FileDialog) SetAcceptLabel(acceptLabel string) {
	cAcceptLabel := cc.NewString(acceptLabel)
	defer cAcceptLabel.Free()
	gtk_file_dialog_set_accept_label.Fn()(fd, cAcceptLabel)
}
func (fd *FileDialog) Open(parent WindowIface,
	cancellable *gio.GCancellable, callback func(fd *FileDialog, res *gio.GAsyncResult)) {
	cb := cc.Cbk(callback)
	fd.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	gtk_file_dialog_open.Fn()(fd, GetWindowIface(parent), cancellable, cb, nil)
}
func (fd *FileDialog) OpenFinish(result *gio.GAsyncResult) (*gio.GFile, error) {
	var gerr *glib.GError
	return gtk_file_dialog_open_finish.Fn()(fd, result, &gerr), gerr.TakeError()
}
func (fd *FileDialog) SelectFolder(parent WindowIface, cancellable *gio.GCancellable,
	callback func(fd *FileDialog, res *gio.GAsyncResult)) {
	cb := cc.Cbk(callback)
	fd.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	gtk_file_dialog_select_folder.Fn()(fd, GetWindowIface(parent), cancellable, cb, nil)
}
func (fd *FileDialog) SelectFolderFinish(result *gio.GAsyncResult) (*gio.GFile, error) {
	var gerr *glib.GError
	folder := gtk_file_dialog_select_folder_finish.Fn()(fd, result, &gerr)
	if gerr != nil {
		defer gerr.Free()
		return nil, gerr.Error()
	}
	return folder, nil
}
func (fd *FileDialog) Save(parent WindowIface, cancellable *gio.GCancellable,
	callback func(fd *FileDialog, res *gio.GAsyncResult)) {
	cb := cc.Cbk(callback)
	fd.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	gtk_file_dialog_save.Fn()(fd, GetWindowIface(parent), cancellable, cb, nil)
}
func (fd *FileDialog) SaveFinish(result *gio.GAsyncResult) (*gio.GFile, error) {
	var gerr *glib.GError
	file := gtk_file_dialog_save_finish.Fn()(fd, result, &gerr)
	if gerr != nil {
		defer gerr.Free()
		return nil, gerr.Error()
	}
	return file, nil
}
func (fd *FileDialog) OpenMultiple(parent WindowIface, cancellable *gio.GCancellable,
	callback func(fd *FileDialog, res *gio.GAsyncResult)) {
	cb := cc.Cbk(callback)
	fd.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	gtk_file_dialog_open_multiple.Fn()(fd, GetWindowIface(parent), cancellable, cb, nil)
}
func (fd *FileDialog) OpenMultipleFinish(result *gio.GAsyncResult) (*gio.GListModel, error) {
	var gerr *glib.GError
	files := gtk_file_dialog_open_multiple_finish.Fn()(fd, result, &gerr)
	if gerr != nil {
		defer gerr.Free()
		return nil, gerr.Error()
	}
	return files, nil
}
func (fd *FileDialog) SelectMultipleFolders(parent WindowIface, cancellable *gio.GCancellable,
	callback func(fd *FileDialog, res *gio.GAsyncResult)) {
	cb := cc.Cbk(callback)
	fd.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	gtk_file_dialog_select_multiple_folders.Fn()(fd, GetWindowIface(parent), cancellable, cb, nil)
}
func (fd *FileDialog) SelectMultipleFoldersFinish(result *gio.GAsyncResult) (*gio.GListModel, error) {
	var gerr *glib.GError
	folders := gtk_file_dialog_select_multiple_folders_finish.Fn()(fd, result, &gerr)
	if gerr != nil {
		defer gerr.Free()
		return nil, gerr.Error()
	}
	return folders, nil
}
func (fd *FileDialog) OpenTextFile(parent WindowIface, cancellable *gio.GCancellable,
	callback func(fd *FileDialog, res *gio.GAsyncResult)) {
	cb := cc.Cbk(callback)
	fd.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	gtk_file_dialog_open_text_file.Fn()(fd, GetWindowIface(parent), cancellable, cb, nil)
}
func (fd *FileDialog) OpenTextFileFinish(result *gio.GAsyncResult) (*gio.GFile, string, error) {
	var gerr *glib.GError
	var encoding cc.String
	return gtk_file_dialog_open_text_file_finish.Fn()(fd, result, &encoding, &gerr), encoding.TakeString(), gerr.TakeError()
}
func (fd *FileDialog) OpenMultipleTextFiles(parent WindowIface, cancellable *gio.GCancellable,
	callback func(fd *FileDialog, res *gio.GAsyncResult)) {
	cb := cc.Cbk(callback)
	fd.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	gtk_file_dialog_open_multiple_text_files.Fn()(fd, GetWindowIface(parent), cancellable, cb, nil)
}
func (fd *FileDialog) OpenMultipleTextFilesFinish(result *gio.GAsyncResult) (*gio.GListModel, string, error) {
	var gerr *glib.GError
	var encoding cc.String
	return gtk_file_dialog_open_multiple_text_files_finish.Fn()(fd, result, &encoding, &gerr), encoding.TakeString(), gerr.TakeError()
}
func (fd *FileDialog) SaveTextFile(parent WindowIface, cancellable *gio.GCancellable,
	callback func(fd *FileDialog, res *gio.GAsyncResult)) {
	cb := cc.Cbk(callback)
	fd.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	gtk_file_dialog_save_text_file.Fn()(fd, GetWindowIface(parent), cancellable, cb, nil)
}
func (fd *FileDialog) SaveTextFileFinish(result *gio.GAsyncResult) (*gio.GFile, string, string, error) {
	var gerr *glib.GError
	var encoding, lineEnding cc.String
	return gtk_file_dialog_save_text_file_finish.Fn()(fd, result, &encoding, &lineEnding, &gerr), encoding.String(), lineEnding.String(), gerr.TakeError()
}

// #endregion

// #region FileFilter

type FileFilter struct {
	Filter
	Buildable
}

func GTypeFileFilter() gobject.GType { return gtk_file_filter_get_type.Fn()() }

func NewFileFilter() *FileFilter { return gtk_file_filter_new.Fn()() }
func (f *FileFilter) SetName(name string) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_file_filter_set_name.Fn()(f, cName)
}
func (f *FileFilter) GetName() string { return gtk_file_filter_get_name.Fn()(f).String() }
func (f *FileFilter) AddMimeType(mimeType string) {
	cMimeType := cc.NewString(mimeType)
	defer cMimeType.Free()
	gtk_file_filter_add_mime_type.Fn()(f, cMimeType)
}
func (f *FileFilter) AddPattern(pattern string) {
	cPattern := cc.NewString(pattern)
	defer cPattern.Free()
	gtk_file_filter_add_pattern.Fn()(f, cPattern)
}
func (f *FileFilter) AddSuffix(suffix string) {
	cSuffix := cc.NewString(suffix)
	defer cSuffix.Free()
	gtk_file_filter_add_suffix.Fn()(f, cSuffix)
}
func (f *FileFilter) AddPixbufFormats()          { gtk_file_filter_add_pixbuf_formats.Fn()(f) }
func (f *FileFilter) GetAttributes() []string    { return gtk_file_filter_get_attributes.Fn()(f).Ref() }
func (f *FileFilter) ToGVariant() *glib.GVariant { return gtk_file_filter_to_gvariant.Fn()(f) }
func NewFileFilterFromGVariant(variant *glib.GVariant) *FileFilter {
	return gtk_file_filter_new_from_gvariant.Fn()(variant)
}

// #endregion

// #region FileLauncher

type FileLauncher struct{ gobjCore }

func GTypeFileLauncher() gobject.GType              { return gtk_file_launcher_get_type.Fn()() }
func NewFileLauncher(file *gio.GFile) *FileLauncher { return gtk_file_launcher_new.Fn()(file) }
func (fl *FileLauncher) GetFile() *gio.GFile        { return gtk_file_launcher_get_file.Fn()(fl) }
func (fl *FileLauncher) SetFile(file *gio.GFile)    { gtk_file_launcher_set_file.Fn()(fl, file) }
func (fl *FileLauncher) GetAlwaysAsk() bool         { return gtk_file_launcher_get_always_ask.Fn()(fl) != 0 }
func (fl *FileLauncher) SetAlwaysAsk(alwaysAsk bool) {
	gtk_file_launcher_set_always_ask.Fn()(fl, cbool(alwaysAsk))
}
func (fl *FileLauncher) GetWritable() bool { return gtk_file_launcher_get_writable.Fn()(fl) != 0 }
func (fl *FileLauncher) SetWritable(writable bool) {
	gtk_file_launcher_set_writable.Fn()(fl, cbool(writable))
}
func (fl *FileLauncher) Launch(parent WindowIface, cancellable *gio.GCancellable,
	callback func(fl *FileLauncher, res *gio.GAsyncResult)) {
	cb := cc.Cbk(callback)
	fl.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	gtk_file_launcher_launch.Fn()(fl, GetWindowIface(parent), cancellable, cb, nil)
}
func (fl *FileLauncher) LaunchFinish(result *gio.GAsyncResult) error {
	var gerr *glib.GError
	gtk_file_launcher_launch_finish.Fn()(fl, result, &gerr)
	return gerr.TakeError()
}
func (fl *FileLauncher) OpenContainingFolder(parent WindowIface, cancellable *gio.GCancellable,
	callback func(fl *FileLauncher, res *gio.GAsyncResult)) {
	cb := cc.Cbk(callback)
	fl.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	gtk_file_launcher_open_containing_folder.Fn()(fl, GetWindowIface(parent), cancellable, cb, nil)
}
func (fl *FileLauncher) OpenContainingFolderFinish(result *gio.GAsyncResult) error {
	var gerr *glib.GError
	gtk_file_launcher_open_containing_folder_finish.Fn()(fl, result, &gerr)
	return gerr.TakeError()
}

// #endregion

// #region Filter

type FilterObj struct {
	Parent gobject.GObjectObj
}
type FilterClass struct {
	Parent gobject.GObjectClass

	Match         cc.Func
	GetStricyness cc.Func
	_             [8]uptr //reserved
}

type FilterIface interface {
	GetFilterIface() *Filter
}

func GetFilterIface(iface FilterIface) *Filter {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetFilterIface()
}
func (f *Filter) GetFilterIface() *Filter { return f }

type Filter struct{ gobjCore }

func (f *Filter) Match(item GobjIface) bool   { return gtk_filter_match.Fn()(f, gobjGet(item)) != 0 }
func (f *Filter) GetStrictness() FilterMatch  { return gtk_filter_get_strictness.Fn()(f) }
func (f *Filter) Changed(change FilterChange) { gtk_filter_changed.Fn()(f, change) }

// #endregion

// #region FilterListModel

type FilterListModel struct {
	gobjCore
	SectionModel
}

func GTypeFilterListModel() gobject.GType { return gtk_filter_list_model_get_type.Fn()() }
func NewFilterListModel(model gio.GListModelIface, filter FilterIface) *FilterListModel {
	return gtk_filter_list_model_new.Fn()(gio.GetGListModelIface(model), GetFilterIface(filter))
}
func (flm *FilterListModel) SetFilter(filter FilterIface) {
	gtk_filter_list_model_set_filter.Fn()(flm, GetFilterIface(filter))
}
func (flm *FilterListModel) GetFilter() *Filter {
	return gtk_filter_list_model_get_filter.Fn()(flm)
}
func (flm *FilterListModel) SetModel(model gio.GListModelIface) {
	gtk_filter_list_model_set_model.Fn()(flm, gio.GetGListModelIface(model))
}
func (flm *FilterListModel) GetModel() *gio.GListModel {
	return gtk_filter_list_model_get_model.Fn()(flm)
}
func (flm *FilterListModel) SetIncremental(incremental bool) {
	gtk_filter_list_model_set_incremental.Fn()(flm, cbool(incremental))
}
func (flm *FilterListModel) GetIncremental() bool {
	return gtk_filter_list_model_get_incremental.Fn()(flm) != 0
}
func (flm *FilterListModel) GetPending() uint32 {
	return gtk_filter_list_model_get_pending.Fn()(flm)
}

// #endregion

// #region Fixed

type FixedObj struct{ Parent WidgetObj }
type FixedClass struct {
	Parent WidgetClass
	_      [8]uptr //padding
}

type Fixed struct{ Widget }

func GTypeFixed() gobject.GType { return gtk_fixed_get_type.Fn()() }
func NewFixed() *Fixed          { return gtk_fixed_new.Fn()() }
func (f *Fixed) Put(widget WidgetIface, x, y float64) {
	gtk_fixed_put.Fn()(f, GetWidgetIface(widget), x, y)
}
func (f *Fixed) Remove(widget WidgetIface) { gtk_fixed_remove.Fn()(f, GetWidgetIface(widget)) }
func (f *Fixed) Move(widget WidgetIface, x, y float64) {
	gtk_fixed_move.Fn()(f, GetWidgetIface(widget), x, y)
}
func (f *Fixed) GetChildPosition(widget WidgetIface) (x, y float64) {
	gtk_fixed_get_child_position.Fn()(f, GetWidgetIface(widget), &x, &y)
	return
}
func (f *Fixed) SetChildTransform(widget WidgetIface, transform *gsk.Transform) {
	gtk_fixed_set_child_transform.Fn()(f, GetWidgetIface(widget), transform)
}
func (f *Fixed) GetChildTransform(widget WidgetIface) *gsk.Transform {
	return gtk_fixed_get_child_transform.Fn()(f, GetWidgetIface(widget))
}

// #endregion

// #region FixedLayout

type FixedLayout struct {
	gobjCore
	LayoutManager
}
type FixedLayoutChild struct{ LayoutChild }

func GTypeFixedLayout() gobject.GType { return gtk_fixed_layout_get_type.Fn()() }
func NewFixedLayout() *FixedLayout    { return gtk_fixed_layout_new.Fn()() }
func (child *FixedLayoutChild) SetTransform(transform *gsk.Transform) {
	gtk_fixed_layout_child_set_transform.Fn()(child, transform)
}
func (child *FixedLayoutChild) GetTransform() *gsk.Transform {
	return gtk_fixed_layout_child_get_transform.Fn()(child)
}

// #endregion

// #region FlattenListModel

type FlattenListModel struct {
	gobjCore
	SectionModel
}

func GTypeFlattenListModel() gobject.GType { return gtk_flatten_list_model_get_type.Fn()() }
func NewFlattenListModel(model gio.GListModelIface) *FlattenListModel {
	return gtk_flatten_list_model_new.Fn()(gio.GetGListModelIface(model))
}
func (flm *FlattenListModel) SetModel(model gio.GListModelIface) {
	gtk_flatten_list_model_set_model.Fn()(flm, gio.GetGListModelIface(model))
}
func (flm *FlattenListModel) GetModel() *gio.GListModel {
	return gtk_flatten_list_model_get_model.Fn()(flm)
}
func (flm *FlattenListModel) GetModelForItem(position uint32) *gio.GListModel {
	return gtk_flatten_list_model_get_model_for_item.Fn()(flm, position)
}

// #endregion

// #region FlowBox

type FlowBoxChildObj struct {
	Parent WidgetObj
}
type FlowBoxChildClass struct {
	Parent WidgetClass

	Activate cc.Func //  void (*)(GtkFlowBoxChild*)
	_        [8]uptr //padding
}

type FlowBoxChild struct{ Widget }

func GTypeFlowBoxChild() gobject.GType { return gtk_flow_box_child_get_type.Fn()() }
func NewFlowBoxChild() *FlowBoxChild   { return gtk_flow_box_child_new.Fn()() }
func (f *FlowBoxChild) SetChild(child WidgetIface) {
	gtk_flow_box_child_set_child.Fn()(f, GetWidgetIface(child))
}
func (f *FlowBoxChild) GetChild() *Widget { return gtk_flow_box_child_get_child.Fn()(f) }
func (f *FlowBoxChild) GetIndex() int32   { return gtk_flow_box_child_get_index.Fn()(f) }
func (f *FlowBoxChild) IsSelected() bool  { return gtk_flow_box_child_is_selected.Fn()(f) != 0 }
func (f *FlowBoxChild) Changed()          { gtk_flow_box_child_changed.Fn()(f) }
func (f *FlowBoxChild) ConnectActivate(sig func(*FlowBoxChild)) uint64 {
	return f.SignalConnect("activate", sig, nil)
}

type FlowBox struct {
	Widget
	Orientable
}

func GTypeFlowBox() gobject.GType { return gtk_flow_box_get_type.Fn()() }
func NewFlowBox() *FlowBox        { return gtk_flow_box_new.Fn()() }
func (f *FlowBox) BindModel(model gio.GListModelIface, createWidgetFunc func(item *gobject.GObject) *Widget) {
	var cf, des uptr
	if createWidgetFunc != nil {
		cc.CbkRaw[func(item *gobject.GObject, _ uptr) *Widget](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 3)
			*(**Widget)(out) = createWidgetFunc(*(**gobject.GObject)(is[0]))
		})
		des = cc.Cbk(func(d uptr) { cc.CbkClose(cf) })
	}
	gtk_flow_box_bind_model.Fn()(f, gio.GetGListModelIface(model), cf, nil, des)
}
func (f *FlowBox) SetHomogeneous(homogeneous bool) {
	gtk_flow_box_set_homogeneous.Fn()(f, cbool(homogeneous))
}
func (f *FlowBox) GetHomogeneous() bool            { return gtk_flow_box_get_homogeneous.Fn()(f) != 0 }
func (f *FlowBox) SetColumnSpacing(spacing uint32) { gtk_flow_box_set_column_spacing.Fn()(f, spacing) }
func (f *FlowBox) GetColumnSpacing() uint32        { return gtk_flow_box_get_column_spacing.Fn()(f) }
func (f *FlowBox) SetMinChildrenPerLine(nChildren uint32) {
	gtk_flow_box_set_min_children_per_line.Fn()(f, nChildren)
}
func (f *FlowBox) GetMinChildrenPerLine() uint32 {
	return gtk_flow_box_get_min_children_per_line.Fn()(f)
}
func (f *FlowBox) SetMaxChildrenPerLine(nChildren uint32) {
	gtk_flow_box_set_max_children_per_line.Fn()(f, nChildren)
}
func (f *FlowBox) GetMaxChildrenPerLine() uint32 {
	return gtk_flow_box_get_max_children_per_line.Fn()(f)
}
func (f *FlowBox) SetActivateOnSingleClick(single bool) {
	gtk_flow_box_set_activate_on_single_click.Fn()(f, cbool(single))
}
func (f *FlowBox) GetActivateOnSingleClick() bool {
	return gtk_flow_box_get_activate_on_single_click.Fn()(f) != 0
}
func (f *FlowBox) Prepend(child WidgetIface) { gtk_flow_box_prepend.Fn()(f, GetWidgetIface(child)) }
func (f *FlowBox) Append(child WidgetIface)  { gtk_flow_box_append.Fn()(f, GetWidgetIface(child)) }
func (f *FlowBox) Insert(widget WidgetIface, position int32) {
	gtk_flow_box_insert.Fn()(f, GetWidgetIface(widget), position)
}
func (f *FlowBox) Remove(widget WidgetIface) { gtk_flow_box_remove.Fn()(f, GetWidgetIface(widget)) }
func (f *FlowBox) RemoveAll()                { gtk_flow_box_remove_all.Fn()(f) }
func (f *FlowBox) GetChildAtIndex(idx int32) *FlowBoxChild {
	return gtk_flow_box_get_child_at_index.Fn()(f, idx)
}
func (f *FlowBox) GetChildAtPos(x, y int32) *FlowBoxChild {
	return gtk_flow_box_get_child_at_pos.Fn()(f, x, y)
}
func (f *FlowBox) SelectedForeach(
	foreachFunc func(box *FlowBox, child *FlowBoxChild)) {
	var ff uptr
	if foreachFunc != nil {
		ff = cc.CbkRaw[func(box *FlowBox, child *FlowBoxChild, _ uptr)](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 3)
			foreachFunc(*(**FlowBox)(is[0]), *(**FlowBoxChild)(is[1]))
		})
		defer cc.CbkClose(ff)
	}

	gtk_flow_box_selected_foreach.Fn()(f, ff, nil)
}
func (f *FlowBox) GetSelectedChildren() *glib.GList[Widget] {
	return gtk_flow_box_get_selected_children.Fn()(f)
}
func (f *FlowBox) SelectChild(child *FlowBoxChild)     { gtk_flow_box_select_child.Fn()(f, child) }
func (f *FlowBox) UnselectChild(child *FlowBoxChild)   { gtk_flow_box_unselect_child.Fn()(f, child) }
func (f *FlowBox) SelectAll()                          { gtk_flow_box_select_all.Fn()(f) }
func (f *FlowBox) UnselectAll()                        { gtk_flow_box_unselect_all.Fn()(f) }
func (f *FlowBox) SetSelectionMode(mode SelectionMode) { gtk_flow_box_set_selection_mode.Fn()(f, mode) }
func (f *FlowBox) GetSelectionMode() SelectionMode     { return gtk_flow_box_get_selection_mode.Fn()(f) }
func (f *FlowBox) AetHAdjustment(adj *Adjustment) {
	gtk_flow_box_set_hadjustment.Fn()(f, adj)
}
func (f *FlowBox) AetVAdjustment(adj *Adjustment) {
	gtk_flow_box_set_vadjustment.Fn()(f, adj)
}
func (f *FlowBox) SetFilterFunc(filterFunc func(child *FlowBoxChild) bool) {
	var cf, des uptr
	if filterFunc != nil {
		cc.CbkRaw[func(child *FlowBoxChild, _ uptr) int32](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 2)
			if filterFunc(*(**FlowBoxChild)(is[0])) {
				*(*int32)(out) = 1
			} else {
				*(*int32)(out) = 0
			}
		})
		des = cc.Cbk(func(d uptr) { cc.CbkClose(cf); cc.CbkCloseLate(des) })
	}

	gtk_flow_box_set_filter_func.Fn()(f, cf, nil, des)
}
func (f *FlowBox) InvalidateFilter() { gtk_flow_box_invalidate_filter.Fn()(f) }
func (f *FlowBox) SetSortFunc(sortFunc func(child1, child2 *FlowBoxChild) int32) {
	var sf, des uptr
	if sortFunc != nil {
		cc.CbkRaw[func(child1, child2 *FlowBoxChild, _ uptr) int32](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 2)
			*(*int32)(out) = sortFunc(*(**FlowBoxChild)(is[0]), (*(**FlowBoxChild)(is[1])))
		})
		des = cc.Cbk(func(d uptr) { cc.CbkClose(sf); cc.CbkCloseLate(des) })
	}
	gtk_flow_box_set_sort_func.Fn()(f, sf, nil, des)
}
func (f *FlowBox) InvalidateSort() { gtk_flow_box_invalidate_sort.Fn()(f) }

func (f *FlowBox) ConnectActivateCursorChild(sig func(*FlowBox)) uint64 {
	return f.SignalConnect("activate-cursor-child", sig, nil)
}
func (f *FlowBox) ConnectChildActivated(sig func(*FlowBox, *FlowBoxChild)) uint64 {
	return f.SignalConnect("child-activated", sig, nil)
}
func (f *FlowBox) ConnectMoveCursor(
	sig func(fb *FlowBox, step MovementStep, count int32, extend, modify bool) bool) uint64 {
	return f.SignalConnect("move-cursor", sig, nil)
}
func (f *FlowBox) ConnectSelectAll(sig func(*FlowBox)) uint64 {
	return f.SignalConnect("select-all", sig, nil)
}
func (f *FlowBox) ConnectSelectedChildrenChanged(sig func(*FlowBox)) uint64 {
	return f.SignalConnect("selected-children-changed", sig, nil)
}
func (f *FlowBox) ConnectToggleCursorChild(sig func(*FlowBox)) uint64 {
	return f.SignalConnect("toggle-cursor-child", sig, nil)
}
func (f *FlowBox) ConnectUnselectAll(sig func(*FlowBox)) uint64 {
	return f.SignalConnect("unselect-all", sig, nil)
}

// #endregion

// #region FontDialog

type FontDialog struct{ gobjCore }

func GTypeFontDialog() gobject.GType { return gtk_font_dialog_get_type.Fn()() }
func NewFontDialog() *FontDialog     { return gtk_font_dialog_new.Fn()() }
func (fd *FontDialog) GetTitle() string {
	return gtk_font_dialog_get_title.Fn()(fd).String()
}
func (fd *FontDialog) SetTitle(title string) {
	chs := cc.NewString(title)
	defer chs.Free()
	gtk_font_dialog_set_title.Fn()(fd, chs)
}
func (fd *FontDialog) GetModal() bool               { return gtk_font_dialog_get_modal.Fn()(fd) != 0 }
func (fd *FontDialog) SetModal(modal bool)          { gtk_font_dialog_set_modal.Fn()(fd, cbool(modal)) }
func (fd *FontDialog) GetLanguage() *pango.Language { return gtk_font_dialog_get_language.Fn()(fd) }
func (fd *FontDialog) SetLanguage(language *pango.Language) {
	gtk_font_dialog_set_language.Fn()(fd, language)
}
func (fd *FontDialog) GetFontMap() *pango.FontMap { return gtk_font_dialog_get_font_map.Fn()(fd) }
func (fd *FontDialog) SetFontMap(fontmap *pango.FontMap) {
	gtk_font_dialog_set_font_map.Fn()(fd, fontmap)
}
func (fd *FontDialog) GetFilter() *Filter { return gtk_font_dialog_get_filter.Fn()(fd) }
func (fd *FontDialog) SetFilter(filter FilterIface) {
	gtk_font_dialog_set_filter.Fn()(fd, GetFilterIface(filter))
}
func (fd *FontDialog) ChooseFamily(parent WindowIface, initialValue *pango.FontFamily,
	cancellable *gio.GCancellable, callback func(fd *FontDialog, res *gio.GAsyncResult)) {
	cbk := cc.Cbk(callback)
	fd.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	gtk_font_dialog_choose_family.Fn()(fd, GetWindowIface(parent), initialValue, cancellable, cbk, nil)
}
func (fd *FontDialog) ChooseFamilyFinish(result *gio.GAsyncResult) (*pango.FontFamily, error) {
	var gerr *glib.GError
	return gtk_font_dialog_choose_family_finish.Fn()(fd, result, &gerr), gerr.TakeError()
}
func (fd *FontDialog) ChooseFace(parent WindowIface, initialValue *pango.FontFace,
	cancellable *gio.GCancellable, callback func(fd *FontDialog, res *gio.GAsyncResult)) {
	cbk := cc.Cbk(callback)
	fd.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	gtk_font_dialog_choose_face.Fn()(fd, GetWindowIface(parent), initialValue, cancellable, cbk, nil)
}
func (fd *FontDialog) ChooseFaceFinish(result *gio.GAsyncResult) (*pango.FontFace, error) {
	var gerr *glib.GError
	return gtk_font_dialog_choose_face_finish.Fn()(fd, result, &gerr), gerr.TakeError()
}
func (fd *FontDialog) ChooseFont(parent WindowIface, initialValue *pango.FontDescription,
	cancellable *gio.GCancellable, callback func(fd *FontDialog, res *gio.GAsyncResult)) {
	cbk := cc.Cbk(callback)
	fd.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	gtk_font_dialog_choose_font.Fn()(fd, GetWindowIface(parent), initialValue, cancellable, cbk, nil)
}
func (fd *FontDialog) ChooseFontFinish(result *gio.GAsyncResult) (*pango.FontDescription, error) {
	var gerr *glib.GError
	return gtk_font_dialog_choose_font_finish.Fn()(fd, result, &gerr), gerr.TakeError()
}
func (fd *FontDialog) ChooseFontAndFeatures(parent WindowIface, initialValue *pango.FontDescription,
	cancellable *gio.GCancellable, callback func(fd *FontDialog, res *gio.GAsyncResult)) {
	cbk := cc.Cbk(callback)
	fd.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	gtk_font_dialog_choose_font_and_features.Fn()(fd, GetWindowIface(parent), initialValue, cancellable, cbk, nil)
}
func (fd *FontDialog) ChooseFontAndFeaturesFinish(result *gio.GAsyncResult) (desc *pango.FontDescription, features string, lang *pango.Language, err error) {
	var gerr *glib.GError
	var fontDesc *pango.FontDescription
	var fontFeatures cc.String
	var language *pango.Language
	ok := gtk_font_dialog_choose_font_and_features_finish.Fn()(fd, result, &fontDesc, &fontFeatures, &language, &gerr) != 0
	if !ok {
		err = gerr.TakeError()
	}
	desc = fontDesc
	features = fontFeatures.TakeString()
	lang = language
	return
}

// #endregion

// #region FontDialogButton

type FontDialogButton struct{ Widget }

func GTypeFontDialogButton() gobject.GType { return gtk_font_dialog_button_get_type.Fn()() }
func NewFontDialogButton(dialog *FontDialog) *FontDialogButton {
	return gtk_font_dialog_button_new.Fn()(dialog)
}
func (btn *FontDialogButton) GetDialog() *FontDialog {
	return gtk_font_dialog_button_get_dialog.Fn()(btn)
}
func (btn *FontDialogButton) SetDialog(dialog *FontDialog) {
	gtk_font_dialog_button_set_dialog.Fn()(btn, dialog)
}
func (btn *FontDialogButton) GetLevel() FontLevel { return gtk_font_dialog_button_get_level.Fn()(btn) }
func (btn *FontDialogButton) SetLevel(level FontLevel) {
	gtk_font_dialog_button_set_level.Fn()(btn, level)
}
func (btn *FontDialogButton) GetFontDesc() *pango.FontDescription {
	return gtk_font_dialog_button_get_font_desc.Fn()(btn)
}
func (btn *FontDialogButton) SetFontDesc(fontDesc *pango.FontDescription) {
	gtk_font_dialog_button_set_font_desc.Fn()(btn, fontDesc)
}
func (btn *FontDialogButton) GetFontFeatures() string {
	return gtk_font_dialog_button_get_font_features.Fn()(btn).String()
}
func (btn *FontDialogButton) SetFontFeatures(features string) {
	chs := cc.NewString(features)
	defer chs.Free()
	gtk_font_dialog_button_set_font_features.Fn()(btn, chs)
}
func (btn *FontDialogButton) GetLanguage() *pango.Language {
	return gtk_font_dialog_button_get_language.Fn()(btn)
}
func (btn *FontDialogButton) SetLanguage(language *pango.Language) {
	gtk_font_dialog_button_set_language.Fn()(btn, language)
}
func (btn *FontDialogButton) GetUseFont() bool {
	return gtk_font_dialog_button_get_use_font.Fn()(btn) != 0
}
func (btn *FontDialogButton) SetUseFont(useFont bool) {
	gtk_font_dialog_button_set_use_font.Fn()(btn, cbool(useFont))
}
func (btn *FontDialogButton) GetUseSize() bool {
	return gtk_font_dialog_button_get_use_size.Fn()(btn) != 0
}
func (btn *FontDialogButton) SetUseSize(useSize bool) {
	gtk_font_dialog_button_set_use_size.Fn()(btn, cbool(useSize))
}
func (btn *FontDialogButton) ConnectActivate(sig func(*FontDialogButton)) uint64 {
	return btn.SignalConnect("activate", sig, nil)
}

// #endregion

// #region Frame

type FrameObj struct {
	Parent WidgetObj
}

type FrameClass struct {
	Parent                 WidgetClass
	ComputeChildAllocation cc.Func //void (*) (GtkFrame *frame, GtkAllocation *allocation);
	_                      [8]uptr //padding
}
type Frame struct{ Widget }

func GTypeFrame() gobject.GType { return gtk_frame_get_type.Fn()() }
func NewFrame(label string) *Frame {
	chs := cc.NewString(label)
	defer chs.Free()
	return gtk_frame_new.Fn()(chs)
}
func (f *Frame) SetLabel(label string) {
	chs := cc.NewString(label)
	defer chs.Free()
	gtk_frame_set_label.Fn()(f, chs)
}
func (f *Frame) GetLabel() string { return gtk_frame_get_label.Fn()(f).String() }
func (f *Frame) SetLabelWidget(labelWidget WidgetIface) {
	gtk_frame_set_label_widget.Fn()(f, GetWidgetIface(labelWidget))
}
func (f *Frame) GetLabelWidget() *Widget      { return gtk_frame_get_label_widget.Fn()(f) }
func (f *Frame) SetLabelAlign(xalign float32) { gtk_frame_set_label_align.Fn()(f, xalign) }
func (f *Frame) GetLabelAlign() float32       { return gtk_frame_get_label_align.Fn()(f) }
func (f *Frame) SetChild(child WidgetIface)   { gtk_frame_set_child.Fn()(f, GetWidgetIface(child)) }
func (f *Frame) GetChild() *Widget            { return gtk_frame_get_child.Fn()(f) }

// #endregion

// #region Gesture

type GestureIface interface {
	GetGestureIface() *Gesture
}

func GetGestureIface(iface GestureIface) *Gesture {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetGestureIface()
}

func (g *Gesture) GetGestureIface() *Gesture { return g }

type Gesture struct{ EventController }

func GTypeGesture() gobject.GType         { return gtk_gesture_get_type.Fn()() }
func (g *Gesture) GetDevice() *gdk.Device { return gtk_gesture_get_device.Fn()(g) }
func (g *Gesture) SetState(state EventSequenceState) bool {
	return gtk_gesture_set_state.Fn()(g, state) != 0
}
func (g *Gesture) GetSequenceState(sequence *gdk.EventSequence) EventSequenceState {
	return gtk_gesture_get_sequence_state.Fn()(g, sequence)
}
func (g *Gesture) SetSequenceState(sequence *gdk.EventSequence, state EventSequenceState) bool {
	return gtk_gesture_set_sequence_state.Fn()(g, sequence, state) != 0
}
func (g *Gesture) GetSequences() *glib.GList[gdk.EventSequence] {
	return gtk_gesture_get_sequences.Fn()(g)
}
func (g *Gesture) GetLastUpdatedSequence() *gdk.EventSequence {
	return gtk_gesture_get_last_updated_sequence.Fn()(g)
}
func (g *Gesture) HandlesSequence(sequence *gdk.EventSequence) bool {
	return gtk_gesture_handles_sequence.Fn()(g, sequence) != 0
}
func (g *Gesture) GetLastEvent(sequence *gdk.EventSequence) *gdk.Event {
	return gtk_gesture_get_last_event.Fn()(g, sequence)
}
func (g *Gesture) GetPoint(sequence *gdk.EventSequence) (x, y float64, ok bool) {
	ok = gtk_gesture_get_point.Fn()(g, sequence, &x, &y) != 0
	return
}
func (g *Gesture) GetBoundingBox(rect *gdk.Rectangle) bool {
	return gtk_gesture_get_bounding_box.Fn()(g, rect) != 0
}
func (g *Gesture) GetBoundingBoxCenter() (x, y float64, ok bool) {
	ok = gtk_gesture_get_bounding_box_center.Fn()(g, &x, &y) != 0
	return
}
func (g *Gesture) IsActive() bool     { return gtk_gesture_is_active.Fn()(g) != 0 }
func (g *Gesture) IsRecognized() bool { return gtk_gesture_is_recognized.Fn()(g) != 0 }
func (g *Gesture) Group(groupGesture GestureIface) {
	gtk_gesture_group.Fn()(GetGestureIface(groupGesture), g)
}
func (g *Gesture) Ungroup() { gtk_gesture_ungroup.Fn()(g) }
func (g *Gesture) GetGroup() *glib.GList[Gesture] {
	return gtk_gesture_get_group.Fn()(g)
}
func (g *Gesture) IsGroupedWith(other GestureIface) bool {
	return gtk_gesture_is_grouped_with.Fn()(g, GetGestureIface(other)) != 0
}
func (g *Gesture) ConnectBegin(sig func(c *Gesture, seq *gdk.EventSequence)) uint64 {
	return g.SignalConnect("begin", sig, nil)
}
func (g *Gesture) ConnectCancel(sig func(c *Gesture, seq *gdk.EventSequence)) uint64 {
	return g.SignalConnect("cancel", sig, nil)
}
func (g *Gesture) ConnectSequenceStateChanged(
	sig func(c *Gesture, seq *gdk.EventSequence, sta EventSequenceState)) uint64 {
	return g.SignalConnect("sequence-state-changed", sig, nil)
}
func (g *Gesture) ConnectUpdate(sig func(c *Gesture, seq *gdk.EventSequence)) uint64 {
	return g.SignalConnect("update", sig, nil)
}

// #endregion

// #region GestureClick

type GestureClick struct{ GestureSingle }

func GTypeGestureClick() gobject.GType { return gtk_gesture_click_get_type.Fn()() }
func NewGestureClick() *GestureClick   { return gtk_gesture_click_new.Fn()() }
func (g *GestureClick) ConnectPressed(sig func(c *GestureClick, nPress int32, x, y float64)) uint64 {
	return g.SignalConnect("pressed", sig, nil)
}
func (g *GestureClick) ConnectReleased(sig func(c *GestureClick, nPress int32, x, y float64)) uint64 {
	return g.SignalConnect("released", sig, nil)
}
func (g *GestureClick) ConnectStopped(sig func(c *GestureClick)) uint64 {
	return g.SignalConnect("stopped", sig, nil)
}
func (g *GestureClick) ConnectUnpairedRelease(
	sig func(c *GestureClick, x, y float64, sequence *gdk.EventSequence)) uint64 {
	return g.SignalConnect("unpaired-release", sig, nil)
}

// #endregion

// #region GestureDrag

type GestureDragIface interface {
	GetGestureDragIface() *GestureDrag
}

func GetGestureDragIface(iface GestureDragIface) *GestureDrag {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetGestureDragIface()
}

func (g *GestureDrag) GetGestureDragIface() *GestureDrag { return g }

type GestureDrag struct{ GestureSingle }

func GTypeGestureDrag() gobject.GType { return gtk_gesture_drag_get_type.Fn()() }
func NewGestureDrag() *GestureDrag    { return gtk_gesture_drag_new.Fn()() }
func (g *GestureDrag) GetStartPoint() (x, y float64, ok bool) {
	ok = gtk_gesture_drag_get_start_point.Fn()(g, &x, &y) != 0
	return
}
func (g *GestureDrag) GetOffset() (x, y float64, ok bool) {
	ok = gtk_gesture_drag_get_offset.Fn()(g, &x, &y) != 0
	return
}
func (g *GestureDrag) ConnectDragBegin(sig func(c *GestureDrag, startX, startY float64)) uint64 {
	return g.SignalConnect("drag-begin", sig, nil)
}
func (g *GestureDrag) ConnectDragEnd(sig func(c *GestureDrag, startX, startY float64)) uint64 {
	return g.SignalConnect("drag-end", sig, nil)
}
func (g *GestureDrag) ConnectDragUpdate(sig func(c *GestureDrag, startX, startY float64)) uint64 {
	return g.SignalConnect("drag-update", sig, nil)
}

// #endregion

// #region GestureLongPress

type GestureLongPress struct{ GestureSingle }

func GTypeGestureLongPress() gobject.GType   { return gtk_gesture_long_press_get_type.Fn()() }
func NewGestureLongPress() *GestureLongPress { return gtk_gesture_long_press_new.Fn()() }
func (g *GestureLongPress) SetDelayFactor(delayFactor float64) {
	gtk_gesture_long_press_set_delay_factor.Fn()(g, delayFactor)
}
func (g *GestureLongPress) GetDelayFactor() float64 {
	return gtk_gesture_long_press_get_delay_factor.Fn()(g)
}
func (g *GestureLongPress) ConnectCancelled(sig func(c *GestureLongPress)) uint64 {
	return g.SignalConnect("cancelled", sig, nil)
}
func (g *GestureLongPress) ConnectPressed(sig func(c *GestureLongPress, x, y float64)) uint64 {
	return g.SignalConnect("pressed", sig, nil)
}

// #endregion

// #region GesturePan

type GesturePan struct{ GestureDrag }

func GTypeGesturePan() gobject.GType                    { return gtk_gesture_pan_get_type.Fn()() }
func NewGesturePan(orientation Orientation) *GesturePan { return gtk_gesture_pan_new.Fn()(orientation) }
func (g *GesturePan) GetOrientation() Orientation       { return gtk_gesture_pan_get_orientation.Fn()(g) }
func (g *GesturePan) SetOrientation(orientation Orientation) {
	gtk_gesture_pan_set_orientation.Fn()(g, orientation)
}

func (g *GesturePan) ConnectPan(sig func(g *GesturePan, dir PanDirection, offset float64)) uint64 {
	return g.SignalConnect("pan", sig, nil)
}

// #endregion

// #region GestureRotate

type GestureRotate struct{ Gesture }

func GTypeGestureRotate() gobject.GType         { return gtk_gesture_rotate_get_type.Fn()() }
func NewGestureRotate() *GestureRotate          { return gtk_gesture_rotate_new.Fn()() }
func (g *GestureRotate) GetAngleDelta() float64 { return gtk_gesture_rotate_get_angle_delta.Fn()(g) }
func (g *GestureRotate) ConnectAngleChanged(sig func(c *GestureRotate, angle, angle_data float64)) uint64 {
	return g.SignalConnect("angle-changed", sig, nil)
}

// #endregion

// #region GestureSingle

type GestureSingleIface interface {
	GetGestureSingleIface() *GestureSingle
}

func GetGestureSingleIface(iface GestureSingleIface) *GestureSingle {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetGestureSingleIface()
}

func (g *GestureSingle) GetGestureSingleIface() *GestureSingle { return g }

type GestureSingle struct{ Gesture }

func GTypeGestureSingle() gobject.GType     { return gtk_gesture_single_get_type.Fn()() }
func (g *GestureSingle) GetTouchOnly() bool { return gtk_gesture_single_get_touch_only.Fn()(g) != 0 }
func (g *GestureSingle) SetTouchOnly(touchOnly bool) {
	gtk_gesture_single_set_touch_only.Fn()(g, cbool(touchOnly))
}
func (g *GestureSingle) GetExclusive() bool { return gtk_gesture_single_get_exclusive.Fn()(g) != 0 }
func (g *GestureSingle) SetExclusive(exclusive bool) {
	gtk_gesture_single_set_exclusive.Fn()(g, cbool(exclusive))
}
func (g *GestureSingle) GetButton() uint32       { return gtk_gesture_single_get_button.Fn()(g) }
func (g *GestureSingle) SetButton(button uint32) { gtk_gesture_single_set_button.Fn()(g, button) }
func (g *GestureSingle) GetCurrentButton() uint32 {
	return gtk_gesture_single_get_current_button.Fn()(g)
}
func (g *GestureSingle) GetCurrentSequence() *gdk.EventSequence {
	return gtk_gesture_single_get_current_sequence.Fn()(g)
}

// #endregion

// #region GestureStylus

type GestureStylus struct{ GestureSingle }

func GTypeGestureStylus() gobject.GType      { return gtk_gesture_stylus_get_type.Fn()() }
func NewGestureStylus() *GestureStylus       { return gtk_gesture_stylus_new.Fn()() }
func (g *GestureStylus) GetStylusOnly() bool { return gtk_gesture_stylus_get_stylus_only.Fn()(g) != 0 }
func (g *GestureStylus) SetStylusOnly(stylusOnly bool) {
	gtk_gesture_stylus_set_stylus_only.Fn()(g, cbool(stylusOnly))
}
func (g *GestureStylus) GetAxis(axis gdk.AxisUse) (value float64, ok bool) {
	ok = gtk_gesture_stylus_get_axis.Fn()(g, axis, &value) != 0
	return
}
func (g *GestureStylus) GetAxes(axes []gdk.AxisUse) (values []float64, ok bool) {
	var cValues *float64
	ok = gtk_gesture_stylus_get_axes.Fn()(g, carry(axes), &cValues) != 0
	if ok && cValues != nil {
		values = (*[1 << 30]float64)(uptr(cValues))[:len(axes):len(axes)]
	}
	return
}
func (g *GestureStylus) GetBacklog() (backlog []*gdk.TimeCoord, nElems uint32, ok bool) {
	var cBacklog *gdk.TimeCoord
	ok = gtk_gesture_stylus_get_backlog.Fn()(g, &cBacklog, &nElems) != 0
	if ok && cBacklog != nil {
		backlog = (*[1 << 30]*gdk.TimeCoord)(uptr(cBacklog))[:nElems:nElems]
	}
	return
}
func (g *GestureStylus) GetDeviceTool() *gdk.DeviceTool {
	return gtk_gesture_stylus_get_device_tool.Fn()(g)
}
func (g *GestureStylus) ConnectDown(sig func(c *GestureStylus, x, y float64)) uint64 {
	return g.SignalConnect("down", sig, nil)
}
func (g *GestureStylus) ConnectMotion(sig func(c *GestureStylus, x, y float64)) uint64 {
	return g.SignalConnect("motion", sig, nil)
}
func (g *GestureStylus) ConnectProximity(sig func(c *GestureStylus, x, y float64)) uint64 {
	return g.SignalConnect("proximity", sig, nil)
}
func (g *GestureStylus) ConnectUp(sig func(c *GestureStylus, x, y float64)) uint64 {
	return g.SignalConnect("up", sig, nil)
}

// #endregion

// #region GestureSwipe

type GestureSwipe struct{ GestureSingle }

func GTypeGestureSwipe() gobject.GType { return gtk_gesture_swipe_get_type.Fn()() }
func NewGestureSwipe() *GestureSwipe   { return gtk_gesture_swipe_new.Fn()() }
func (g *GestureSwipe) GetVelocity() (velocityX, velocityY float64, ok bool) {
	ok = gtk_gesture_swipe_get_velocity.Fn()(g, &velocityX, &velocityY) != 0
	return
}
func (g *GestureSwipe) ConnectSwipe(sig func(c *GestureSwipe, velocityX, velocityY float64)) uint64 {
	return g.SignalConnect("swipe", sig, nil)
}

// #endregion

// #region GestureZoom

type GestureZoom struct{ Gesture }

func GTypeGestureZoom() gobject.GType         { return gtk_gesture_zoom_get_type.Fn()() }
func NewGestureZoom() *GestureZoom            { return gtk_gesture_zoom_new.Fn()() }
func (g *GestureZoom) GetScaleDelta() float64 { return gtk_gesture_zoom_get_scale_delta.Fn()(g) }
func (g *GestureZoom) ConnectScaleChanged(sig func(c *GestureZoom, scale float64)) uint64 {
	return g.SignalConnect("scale-changed", sig, nil)
}

// #endregion

// #region GLArea

type GLAreaObj struct {
	Parent WidgetObj
}
type GLAreaClass struct {
	Parent        WidgetClass
	Render        cc.Func //gboolean(*)(GtkGLArea *area,GdkGLContext     *context);
	Resize        cc.Func //void(*)(GtkGLArea *area,int width,int height);
	CreateContext cc.Func // GdkGLContext * (*)(GtkGLArea*area);

	_ [8]uptr //padding

}
type GLArea struct{ Widget }

func GTypeGLArea() gobject.GType                { return gtk_gl_area_get_type.Fn()() }
func NewGLArea() *GLArea                        { return gtk_gl_area_new.Fn()() }
func (a *GLArea) SetAllowedAPIs(apis gdk.GLAPI) { gtk_gl_area_set_allowed_apis.Fn()(a, apis) }
func (a *GLArea) GetAllowedAPIs() gdk.GLAPI     { return gtk_gl_area_get_allowed_apis.Fn()(a) }
func (a *GLArea) GetAPI() gdk.GLAPI             { return gtk_gl_area_get_api.Fn()(a) }
func (a *GLArea) SetRequiredVersion(major, minor int32) {
	gtk_gl_area_set_required_version.Fn()(a, major, minor)
}
func (a *GLArea) GetRequiredVersion() (major, minor int32) {
	gtk_gl_area_get_required_version.Fn()(a, &major, &minor)
	return
}
func (a *GLArea) GetHasDepthBuffer() bool { return gtk_gl_area_get_has_depth_buffer.Fn()(a) != 0 }
func (a *GLArea) SetHasDepthBuffer(hasDepthBuffer bool) {
	gtk_gl_area_set_has_depth_buffer.Fn()(a, cbool(hasDepthBuffer))
}
func (a *GLArea) GetHasStencilBuffer() bool { return gtk_gl_area_get_has_stencil_buffer.Fn()(a) != 0 }
func (a *GLArea) SetHasStencilBuffer(hasStencilBuffer bool) {
	gtk_gl_area_set_has_stencil_buffer.Fn()(a, cbool(hasStencilBuffer))
}
func (a *GLArea) GetAutoRender() bool { return gtk_gl_area_get_auto_render.Fn()(a) != 0 }
func (a *GLArea) SetAutoRender(autoRender bool) {
	gtk_gl_area_set_auto_render.Fn()(a, cbool(autoRender))
}
func (a *GLArea) QueueRender()               { gtk_gl_area_queue_render.Fn()(a) }
func (a *GLArea) GetContext() *gdk.GLContext { return gtk_gl_area_get_context.Fn()(a) }
func (a *GLArea) MakeCurrent()               { gtk_gl_area_make_current.Fn()(a) }
func (a *GLArea) AttachBuffers()             { gtk_gl_area_attach_buffers.Fn()(a) }
func (a *GLArea) SetError(domain glib.GQuark, code int32, msg string) {
	e := glib.NewGError(domain, code, msg)
	defer e.Free()
	gtk_gl_area_set_error.Fn()(a, e)
}
func (a *GLArea) GetError() *glib.GError {
	return gtk_gl_area_get_error.Fn()(a)
}

func (g *GLArea) ConnectCreateContext(sig func(g *GLArea, _ uptr) gdk.GLContext) uint64 {
	return g.SignalConnect("create-context", sig, nil)
}
func (g *GLArea) ConnectRender(sig func(g *GLArea, ctx *gdk.GLContext, _ uptr) bool) uint64 {
	return g.SignalConnect("render", sig, nil)
}
func (g *GLArea) ConnectResize(sig func(g *GLArea, w, h int32)) uint64 {
	return g.SignalConnect("resize", sig, nil)
}

// #endregion

// #region GraphicsOffload

type GraphicsOffload struct{ Widget }

func GTypeGraphicsOffload() gobject.GType { return gtk_graphics_offload_get_type.Fn()() }
func NewGraphicsOffload(child WidgetIface) *GraphicsOffload {
	return gtk_graphics_offload_new.Fn()(GetWidgetIface(child))
}
func (g *GraphicsOffload) SetChild(child WidgetIface) {
	gtk_graphics_offload_set_child.Fn()(g, GetWidgetIface(child))
}
func (g *GraphicsOffload) GetChild() *Widget { return gtk_graphics_offload_get_child.Fn()(g) }
func (g *GraphicsOffload) SetEnabled(enabled GraphicsOffloadEnabled) {
	gtk_graphics_offload_set_enabled.Fn()(g, enabled)
}
func (g *GraphicsOffload) GetEnabled() GraphicsOffloadEnabled {
	return gtk_graphics_offload_get_enabled.Fn()(g)
}
func (g *GraphicsOffload) SetBlackBackground(value bool) {
	gtk_graphics_offload_set_black_background.Fn()(g, cbool(value))
}
func (g *GraphicsOffload) GetBlackBackground() bool {
	return gtk_graphics_offload_get_black_background.Fn()(g) != 0
}

// #endregion

// #region Grid

type GridObj struct{ Parent WidgetObj }
type GridClass struct {
	Parent WidgetClass

	_ [8]uptr //padding
}

type Grid struct {
	Widget
	Orientable
}

func GTypeGrid() gobject.GType { return gtk_grid_get_type.Fn()() }
func NewGrid() *Grid           { return gtk_grid_new.Fn()() }
func (g *Grid) Attach(child WidgetIface, column, row, width, height int32) {
	gtk_grid_attach.Fn()(g, GetWidgetIface(child), column, row, width, height)
}
func (g *Grid) AttachNextTo(child WidgetIface, sibling WidgetIface, side PositionType, width, height int32) {
	gtk_grid_attach_next_to.Fn()(g, GetWidgetIface(child), GetWidgetIface(sibling), side, width, height)
}
func (g *Grid) GetChildAt(column, row int32) *Widget {
	return gtk_grid_get_child_at.Fn()(g, column, row)
}
func (g *Grid) Remove(child WidgetIface)    { gtk_grid_remove.Fn()(g, GetWidgetIface(child)) }
func (g *Grid) InsertRow(position int32)    { gtk_grid_insert_row.Fn()(g, position) }
func (g *Grid) InsertColumn(position int32) { gtk_grid_insert_column.Fn()(g, position) }
func (g *Grid) RemoveRow(position int32)    { gtk_grid_remove_row.Fn()(g, position) }
func (g *Grid) RemoveColumn(position int32) { gtk_grid_remove_column.Fn()(g, position) }
func (g *Grid) InsertNextTo(sibling WidgetIface, side PositionType) {
	gtk_grid_insert_next_to.Fn()(g, GetWidgetIface(sibling), side)
}
func (g *Grid) SetRowHomogeneous(homogeneous bool) {
	gtk_grid_set_row_homogeneous.Fn()(g, cbool(homogeneous))
}
func (g *Grid) GetRowHomogeneous() bool      { return gtk_grid_get_row_homogeneous.Fn()(g) != 0 }
func (g *Grid) SetRowSpacing(spacing uint32) { gtk_grid_set_row_spacing.Fn()(g, spacing) }
func (g *Grid) GetRowSpacing() uint32        { return gtk_grid_get_row_spacing.Fn()(g) }
func (g *Grid) SetColumnHomogeneous(homogeneous bool) {
	gtk_grid_set_column_homogeneous.Fn()(g, cbool(homogeneous))
}
func (g *Grid) GetColumnHomogeneous() bool      { return gtk_grid_get_column_homogeneous.Fn()(g) != 0 }
func (g *Grid) SetColumnSpacing(spacing uint32) { gtk_grid_set_column_spacing.Fn()(g, spacing) }

func (g *Grid) GetColumnSpacing() uint32 { return gtk_grid_get_column_spacing.Fn()(g) }
func (g *Grid) SetRowBaselinePosition(row int32, pos BaselinePosition) {
	gtk_grid_set_row_baseline_position.Fn()(g, row, pos)
}
func (g *Grid) GetRowBaselinePosition(row int32) BaselinePosition {
	return gtk_grid_get_row_baseline_position.Fn()(g, row)
}
func (g *Grid) SetBaselineRow(row int32) { gtk_grid_set_baseline_row.Fn()(g, row) }
func (g *Grid) GetBaselineRow() int32    { return gtk_grid_get_baseline_row.Fn()(g) }
func (g *Grid) QueryChild(child WidgetIface) (column, row, width, height int32) {
	gtk_grid_query_child.Fn()(g, GetWidgetIface(child), &column, &row, &width, &height)
	return
}

// #endregion

// #region GridLayout

type GridLayout struct {
	gobjCore
	LayoutManager
}

func GTypeGridLayout() gobject.GType { return gtk_grid_layout_get_type.Fn()() }
func NewGridLayout() *GridLayout     { return gtk_grid_layout_new.Fn()() }
func (g *GridLayout) SetRowHomogeneous(homogeneous bool) {
	gtk_grid_layout_set_row_homogeneous.Fn()(g, cbool(homogeneous))
}
func (g *GridLayout) GetRowHomogeneous() bool {
	return gtk_grid_layout_get_row_homogeneous.Fn()(g) != 0
}
func (g *GridLayout) SetRowSpacing(spacing uint32) { gtk_grid_layout_set_row_spacing.Fn()(g, spacing) }
func (g *GridLayout) GetRowSpacing() uint32        { return gtk_grid_layout_get_row_spacing.Fn()(g) }
func (g *GridLayout) SetColumnHomogeneous(homogeneous bool) {
	gtk_grid_layout_set_column_homogeneous.Fn()(g, cbool(homogeneous))
}
func (g *GridLayout) GetColumnHomogeneous() bool {
	return gtk_grid_layout_get_column_homogeneous.Fn()(g) != 0
}
func (g *GridLayout) SetColumnSpacing(spacing uint32) {
	gtk_grid_layout_set_column_spacing.Fn()(g, spacing)
}
func (g *GridLayout) GetColumnSpacing() uint32 { return gtk_grid_layout_get_column_spacing.Fn()(g) }
func (g *GridLayout) SetRowBaselinePosition(row int32, pos BaselinePosition) {
	gtk_grid_layout_set_row_baseline_position.Fn()(g, row, pos)
}
func (g *GridLayout) GetRowBaselinePosition(row int32) BaselinePosition {
	return gtk_grid_layout_get_row_baseline_position.Fn()(g, row)
}
func (g *GridLayout) SetBaselineRow(row int32) { gtk_grid_layout_set_baseline_row.Fn()(g, row) }
func (g *GridLayout) GetBaselineRow() int32    { return gtk_grid_layout_get_baseline_row.Fn()(g) }

type GridLayoutChild struct{ LayoutChild }

func (c *GridLayoutChild) SetRow(row int32)       { gtk_grid_layout_child_set_row.Fn()(c, row) }
func (c *GridLayoutChild) GetRow() int32          { return gtk_grid_layout_child_get_row.Fn()(c) }
func (c *GridLayoutChild) SetColumn(column int32) { gtk_grid_layout_child_set_column.Fn()(c, column) }
func (c *GridLayoutChild) GetColumn() int32       { return gtk_grid_layout_child_get_column.Fn()(c) }
func (c *GridLayoutChild) SetColumnSpan(span int32) {
	gtk_grid_layout_child_set_column_span.Fn()(c, span)
}
func (c *GridLayoutChild) GetColumnSpan() int32  { return gtk_grid_layout_child_get_column_span.Fn()(c) }
func (c *GridLayoutChild) SetRowSpan(span int32) { gtk_grid_layout_child_set_row_span.Fn()(c, span) }
func (c *GridLayoutChild) GetRowSpan() int32     { return gtk_grid_layout_child_get_row_span.Fn()(c) }

// #endregion

// #region GridView

type GridView struct{ ListBase }

func GTypeGridView() gobject.GType { return gtk_grid_view_get_type.Fn()() }
func NewGridView(model SelectionModelIface, factory ListItemFactoryIface) *GridView {
	return gtk_grid_view_new.Fn()(GetSelectionModelIface(model), GetListItemFactoryIface(factory))
}
func (g *GridView) GetModel() *SelectionModel { return gtk_grid_view_get_model.Fn()(g) }
func (g *GridView) SetModel(model SelectionModelIface) {
	gtk_grid_view_set_model.Fn()(g, GetSelectionModelIface(model))
}
func (g *GridView) SetFactory(factory ListItemFactoryIface) {
	gtk_grid_view_set_factory.Fn()(g, GetListItemFactoryIface(factory))
}
func (g *GridView) GetFactory() *ListItemFactory { return gtk_grid_view_get_factory.Fn()(g) }
func (g *GridView) GetMinColumns() uint32        { return gtk_grid_view_get_min_columns.Fn()(g) }
func (g *GridView) SetMinColumns(minColumns uint32) {
	gtk_grid_view_set_min_columns.Fn()(g, minColumns)
}
func (g *GridView) GetMaxColumns() uint32 { return gtk_grid_view_get_max_columns.Fn()(g) }
func (g *GridView) SetMaxColumns(maxColumns uint32) {
	gtk_grid_view_set_max_columns.Fn()(g, maxColumns)
}
func (g *GridView) SetEnableRubberband(enableRubberband bool) {
	gtk_grid_view_set_enable_rubberband.Fn()(g, cbool(enableRubberband))
}
func (g *GridView) GetEnableRubberband() bool {
	return gtk_grid_view_get_enable_rubberband.Fn()(g) != 0
}
func (g *GridView) SetTabBehavior(tabBehavior ListTabBehavior) {
	gtk_grid_view_set_tab_behavior.Fn()(g, tabBehavior)
}
func (g *GridView) GetTabBehavior() ListTabBehavior { return gtk_grid_view_get_tab_behavior.Fn()(g) }
func (g *GridView) SetSingleClickActivate(singleClickActivate bool) {
	gtk_grid_view_set_single_click_activate.Fn()(g, cbool(singleClickActivate))
}
func (g *GridView) GetSingleClickActivate() bool {
	return gtk_grid_view_get_single_click_activate.Fn()(g) != 0
}
func (g *GridView) ScrollTo(pos uint32, flags ListScrollFlags, scroll *ScrollInfo) {
	gtk_grid_view_scroll_to.Fn()(g, pos, flags, scroll)
}
func (g *GridView) ConnectActivate(sig func(g *GridView, position uint32)) uint64 {
	return g.SignalConnect("activate", sig, nil)
}

// #endregion

// #region HeaderBar

type HeaderBar struct{ Widget }

func GTypeHeaderBar() gobject.GType { return gtk_header_bar_get_type.Fn()() }
func NewHeaderBar() *HeaderBar      { return gtk_header_bar_new.Fn()() }

func (h *HeaderBar) SetTitleWidget(widget WidgetIface) {
	gtk_header_bar_set_title_widget.Fn()(h, GetWidgetIface(widget))
}
func (h *HeaderBar) GetTitleWidget() *Widget { return gtk_header_bar_get_title_widget.Fn()(h) }
func (h *HeaderBar) PackStart(child WidgetIface) {
	gtk_header_bar_pack_start.Fn()(h, GetWidgetIface(child))
}
func (h *HeaderBar) PackEnd(child WidgetIface) {
	gtk_header_bar_pack_end.Fn()(h, GetWidgetIface(child))
}
func (h *HeaderBar) Remove(child WidgetIface) {
	gtk_header_bar_remove.Fn()(h, GetWidgetIface(child))
}
func (h *HeaderBar) GetShowTitleButtons() bool {
	return gtk_header_bar_get_show_title_buttons.Fn()(h) != 0
}
func (h *HeaderBar) SetShowTitleButtons(setting bool) {
	gtk_header_bar_set_show_title_buttons.Fn()(h, cbool(setting))
}
func (h *HeaderBar) SetDecorationLayout(layout string) {
	cLayout := cc.NewString(layout)
	defer cLayout.Free()
	gtk_header_bar_set_decoration_layout.Fn()(h, cLayout)
}
func (h *HeaderBar) GetDecorationLayout() string {
	return gtk_header_bar_get_decoration_layout.Fn()(h).String()
}
func (h *HeaderBar) GetUseNativeControls() bool {
	return gtk_header_bar_get_use_native_controls.Fn()(h) != 0
}
func (h *HeaderBar) SetUseNativeControls(setting bool) {
	gtk_header_bar_set_use_native_controls.Fn()(h, cbool(setting))
}

// #endregion

// #region IconTheme

type IconTheme struct{ gobjCore }

func IconThemeErrorQuark() glib.GQuark { return gtk_icon_theme_error_quark.Fn()() }
func GTypeIconTheme() gobject.GType    { return gtk_icon_theme_get_type.Fn()() }
func NewIconTheme() *IconTheme         { return gtk_icon_theme_new.Fn()() }
func GetIconThemeForDisplay(display *gdk.Display) *IconTheme {
	return gtk_icon_theme_get_for_display.Fn()(display)
}
func (it *IconTheme) GetDisplay() *gdk.Display { return gtk_icon_theme_get_display.Fn()(it) }
func (it *IconTheme) SetSearchPath(paths []string) {
	cpp := cc.NewStrings(paths)
	defer cpp.Free()
	gtk_icon_theme_set_search_path.Fn()(it, cpp)
}
func (it *IconTheme) GetSearchPath() []string {
	return gtk_icon_theme_get_search_path.Fn()(it).Ref()
}
func (it *IconTheme) AddSearchPath(path string) {
	cpath := cc.NewString(path)
	defer cpath.Free()
	gtk_icon_theme_add_search_path.Fn()(it, cpath)
}
func (it *IconTheme) SetResourcePath(paths []string) {
	cpp := cc.NewStrings(paths)
	defer cpp.Free()
	gtk_icon_theme_set_resource_path.Fn()(it, cpp)
}
func (it *IconTheme) GetResourcePath() []string {
	return gtk_icon_theme_get_resource_path.Fn()(it).Strings()
}
func (it *IconTheme) AddResourcePath(path string) {
	cpath := cc.NewString(path)
	defer cpath.Free()
	gtk_icon_theme_add_resource_path.Fn()(it, cpath)
}

func (it *IconTheme) SetThemeName(themeName string) {
	cTheme := cc.NewString(themeName)
	defer cTheme.Free()
	gtk_icon_theme_set_theme_name.Fn()(it, cTheme)
}
func (it *IconTheme) GetThemeName() string {
	return gtk_icon_theme_get_theme_name.Fn()(it).String()
}
func (it *IconTheme) HasIcon(iconName string) bool {
	cIcon := cc.NewString(iconName)
	defer cIcon.Free()
	return gtk_icon_theme_has_icon.Fn()(it, cIcon) != 0
}
func (it *IconTheme) HasGicon(gicon *gio.GIcon) bool {
	return gtk_icon_theme_has_gicon.Fn()(it, gicon) != 0
}
func (it *IconTheme) GetIconSizes(iconName string) (sizes []int32) {
	cIcon := cc.NewString(iconName)
	defer cIcon.Free()

	ptr := gtk_icon_theme_get_icon_sizes.Fn()(it, cIcon)
	if ptr == nil {
		return nil
	}

	is := *((*[1024]int32)(uptr(ptr)))

	for i := 0; i < 1024; i++ {
		if is[i] == 0 {
			break
		}
		sizes = append(sizes, is[i])
	}
	return sizes
}
func (it *IconTheme) LookupIcon(iconName string, fallbacks []string, size, scale int32, direction TextDirection, flags IconLookupFlags) *IconPaintable {
	cIcon := cc.NewString(iconName)
	defer cIcon.Free()
	cFallbacks := cc.NewStrings(fallbacks)
	defer cFallbacks.Free()
	return gtk_icon_theme_lookup_icon.Fn()(it, cIcon, cFallbacks, size, scale, direction, flags)
}
func (it *IconTheme) LookupByGicon(gicon *gio.GIcon, size, scale int32, direction TextDirection, flags IconLookupFlags) *IconPaintable {
	return gtk_icon_theme_lookup_by_gicon.Fn()(it, gicon, size, scale, direction, flags)
}
func IconPaintableNewForFile(file *gio.GFile, size, scale int32) *IconPaintable {
	return gtk_icon_paintable_new_for_file.Fn()(file, size, scale)
}
func (it *IconTheme) GetIconNames() []string {
	return gtk_icon_theme_get_icon_names.Fn()(it).Strings()
}
func (it *IconTheme) ConnectChanged(sig func(i *IconTheme)) uint64 {
	return it.SignalConnect("changed", sig, nil)
}

type IconPaintable struct{ gobjCore }

func GTypeIconPaintable() gobject.GType       { return gtk_icon_paintable_get_type.Fn()() }
func (ip *IconPaintable) GetFile() *gio.GFile { return gtk_icon_paintable_get_file.Fn()(ip) }
func (ip *IconPaintable) GetIconName() string {
	return gtk_icon_paintable_get_icon_name.Fn()(ip).String()
}
func (ip *IconPaintable) IsSymbolic() bool { return gtk_icon_paintable_is_symbolic.Fn()(ip) != 0 }

// #endregion

// #region Image

type Image struct{ Widget }

func GTypeImage() gobject.GType { return gtk_image_get_type.Fn()() }
func NewImage() *Image          { return gtk_image_new.Fn()() }
func NewImageFromFile(filename string) *Image {
	cFilename := cc.NewString(filename)
	defer cFilename.Free()
	return gtk_image_new_from_file.Fn()(cFilename)
}
func NewImageFromResource(resourcePath string) *Image {
	cResource := cc.NewString(resourcePath)
	defer cResource.Free()
	return gtk_image_new_from_resource.Fn()(cResource)
}

// func NewImageFromPixbuf(pixbuf *gdk.Pixbuf) *Image { return gtk_image_new_from_pixbuf(pixbuf) }

func NewImageFromPaintable(paintable *gdk.Paintable) *Image {
	return gtk_image_new_from_paintable.Fn()(paintable)
}
func NewImageFromIconName(iconName string) *Image {
	cIcon := cc.NewString(iconName)
	defer cIcon.Free()
	return gtk_image_new_from_icon_name.Fn()(cIcon)
}
func NewImageFromGicon(icon *gio.GIcon) *Image { return gtk_image_new_from_gicon.Fn()(icon) }

func (img *Image) Clear() { gtk_image_clear.Fn()(img) }
func (img *Image) SetFromFile(filename string) {
	cFilename := cc.NewString(filename)
	defer cFilename.Free()
	gtk_image_set_from_file.Fn()(img, cFilename)
}
func (img *Image) SetFromResource(resourcePath string) {
	cResource := cc.NewString(resourcePath)
	defer cResource.Free()
	gtk_image_set_from_resource.Fn()(img, cResource)
}

// func (img *Image) SetFromPixbuf(pixbuf *gdk.Pixbuf) { gtk_image_set_from_pixbuf(img, pixbuf) }

func (img *Image) SetFromPaintable(paintable *gdk.Paintable) {
	gtk_image_set_from_paintable.Fn()(img, paintable)
}
func (img *Image) SetFromIconName(iconName string) {
	cIcon := cc.NewString(iconName)
	defer cIcon.Free()
	gtk_image_set_from_icon_name.Fn()(img, cIcon)
}
func (img *Image) SetFromGicon(icon *gio.GIcon)  { gtk_image_set_from_gicon.Fn()(img, icon) }
func (img *Image) SetPixelSize(pixelSize int32)  { gtk_image_set_pixel_size.Fn()(img, pixelSize) }
func (img *Image) SetIconSize(iconSize IconSize) { gtk_image_set_icon_size.Fn()(img, iconSize) }
func (img *Image) GetStorageType() ImageType     { return gtk_image_get_storage_type.Fn()(img) }
func (img *Image) GetPaintable() *gdk.Paintable  { return gtk_image_get_paintable.Fn()(img) }
func (img *Image) GetIconName() string           { return gtk_image_get_icon_name.Fn()(img).String() }
func (img *Image) GetGicon() *gio.GIcon          { return gtk_image_get_gicon.Fn()(img) }
func (img *Image) GetPixelSize() int32           { return gtk_image_get_pixel_size.Fn()(img) }
func (img *Image) GetIconSize() IconSize         { return gtk_image_get_icon_size.Fn()(img) }

// #endregion

// #region IMContext

type IMContextObj struct {
	Parent gobject.GObjectObj
}
type IMContextClass struct {
	/*< private >*/
	Parent gobject.GObjectClass

	/*< public >*/
	/* Signals */
	PreeditStart        cc.Func // void (*preedit_start) (GtkIMContext *context);
	PreeditEnd          cc.Func // void (*preedit_end) (GtkIMContext *context);
	PreeditChanged      cc.Func // void (*preedit_changed) (GtkIMContext *context);
	Commit              cc.Func // void (*commit) (GtkIMContext *context, const char *str);
	RetrieveSurrounding cc.Func // gboolean (*retrieve_surrounding) (GtkIMContext *context);
	DeleteSurrounding   cc.Func // gboolean (*delete_surrounding) (GtkIMContext *context,int offset,int n_chars);

	/* Virtual functions */
	SetClientWidget             cc.Func // void (*set_client_widget) (GtkIMContext *context, GtkWidget *widget);
	GetPreeditString            cc.Func // void (*get_preedit_string) (GtkIMContext *context, char **str, PangoAttrList **attrs, int *cursor_pos);
	FilterKeypress              cc.Func // gboolean (*filter_keypress) (GtkIMContext *context,GdkEvent *event);
	FocusIn                     cc.Func // void (*focus_in) (GtkIMContext *context);
	FocusOut                    cc.Func // void (*focus_out) (GtkIMContext *context);
	Reset                       cc.Func // void (*reset) (GtkIMContext *context);
	SetCursorLocation           cc.Func // void (*set_cursor_location) (GtkIMContext *context, GdkRectangle *area);
	SetUsePreedit               cc.Func // void (*set_use_preedit) (GtkIMContext *context, gboolean use_preedit);
	SetSurrounding              cc.Func // void (*set_surrounding) (GtkIMContext *context, const char *text, int len, int cursor_index);
	GetSurrounding              cc.Func // gboolean (*get_surrounding) (GtkIMContext *context, char **text, int *cursor_index);
	SetSurroundingWithSelection cc.Func // void (*set_surrounding_with_selection)(GtkIMContext *context, const char *text, int len, int cursor_index,int anchor_index);
	GetSurroundingWithSelection cc.Func // gboolean (*get_surrounding_with_selection)(GtkIMContext *context, char **text, int *cursor_index, int *anchor_index);

	/*< private >*/
	ActivateOSK          cc.Func // void (* activate_osk) (GtkIMContext *context);
	ActivateOSKWithEvent cc.Func // gboolean (* activate_osk_with_event) (GtkIMContext *context,GdkEvent *event);

	/* Padding for future expansion */
	_ [3]uintptr // padding
}

type IMContextIface interface {
	GetIMContextIface() *IMContext
}

func GetIMContextIface(iface IMContextIface) *IMContext {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetIMContextIface()
}
func (c *IMContext) GetIMContextIface() *IMContext { return c }

type IMContext struct{ gobjCore }

func GTypeIMContext() gobject.GType { return gtk_im_context_get_type.Fn()() }
func (c *IMContext) SetClientWidget(widget WidgetIface) {
	gtk_im_context_set_client_widget.Fn()(c, GetWidgetIface(widget))
}
func (c *IMContext) GetPreeditString() (str string, attrs *pango.AttrList, cursorPos int32) {
	var cStr cc.String
	var cAttrs *pango.AttrList
	gtk_im_context_get_preedit_string.Fn()(c, &cStr, &cAttrs, &cursorPos)
	str = cStr.TakeString()
	attrs = cAttrs

	defer cStr.Free()
	return
}
func (c *IMContext) FilterKeypress(event *gdk.Event) bool {
	return gtk_im_context_filter_keypress.Fn()(c, event) != 0
}
func (c *IMContext) FilterKey(press bool, surface *gdk.Surface, device *gdk.Device, time uint32, keycode uint32, state gdk.ModifierType, group int32) bool {
	return gtk_im_context_filter_key.Fn()(c, cbool(press), surface, device, time, keycode, state, group) != 0
}
func (c *IMContext) FocusIn()  { gtk_im_context_focus_in.Fn()(c) }
func (c *IMContext) FocusOut() { gtk_im_context_focus_out.Fn()(c) }
func (c *IMContext) Reset()    { gtk_im_context_reset.Fn()(c) }
func (c *IMContext) SetCursorLocation(area *gdk.Rectangle) {
	gtk_im_context_set_cursor_location.Fn()(c, area)
}
func (c *IMContext) SetUsePreedit(usePreedit bool) {
	gtk_im_context_set_use_preedit.Fn()(c, cbool(usePreedit))
}
func (c *IMContext) SetSurroundingWithSelection(text string, cursorIndex, anchorIndex int32) {
	t := cc.NewString(text)
	defer t.Free()
	gtk_im_context_set_surrounding_with_selection.Fn()(c, t, int32(len(text)), cursorIndex, anchorIndex)
}
func (c *IMContext) GetSurroundingWithSelection() (text string, cursorIndex, anchorIndex int32, ok bool) {
	var cText cc.String
	ok = gtk_im_context_get_surrounding_with_selection.Fn()(c, &cText, &cursorIndex, &anchorIndex) != 0
	text = cText.TakeString()
	return
}
func (c *IMContext) DeleteSurrounding(offset, nChars int32) bool {
	return gtk_im_context_delete_surrounding.Fn()(c, offset, nChars) != 0
}
func (c *IMContext) ActivateOSK(event *gdk.Event) bool {
	return gtk_im_context_activate_osk.Fn()(c, event) != 0
}
func (c *IMContext) ConnectCommit(sig func(c *IMContext, str cc.String)) uint64 {
	return c.SignalConnect("commit", sig, nil)
}
func (c *IMContext) ConnectDeleteSurrounding(sig func(c *IMContext, offset, nChars int32)) uint64 {
	return c.SignalConnect("delete-surrounding", sig, nil)
}
func (c *IMContext) ConnectPreeditChanged(sig func(c *IMContext)) uint64 {
	return c.SignalConnect("preedit-changed", sig, nil)
}
func (c *IMContext) ConnectPreeditEnd(sig func(c *IMContext)) uint64 {
	return c.SignalConnect("preedit-end", sig, nil)
}
func (c *IMContext) ConnectPreeditStart(sig func(c *IMContext)) uint64 {
	return c.SignalConnect("preedit-start", sig, nil)
}
func (c *IMContext) ConnectRetrieveSurrounding(sig func(c *IMContext)) uint64 {
	return c.SignalConnect("retrieve-surrounding", sig, nil)
}

// #endregion

// #region IMContextSimple

type IMContextSimpleObj struct {
	Parent IMContextObj

	Priv uptr
}

type IMContextSimpleClass struct {
	Parent IMContextClass
}

type IMContextSimple struct{ IMContext }

func GTypeIMContextSimple() gobject.GType  { return gtk_im_context_simple_get_type.Fn()() }
func NewIMContextSimple() *IMContextSimple { return gtk_im_context_simple_new.Fn()() }
func (c *IMContextSimple) AddTable(data []uint16, maxSeqLen int32, nSeqs int32) {
	gtk_im_context_simple_add_table.Fn()(c, carry(data), maxSeqLen, nSeqs)
}
func (c *IMContextSimple) AddComposeFile(composeFile string) {
	cfile := cc.NewString(composeFile)
	defer cfile.Free()
	gtk_im_context_simple_add_compose_file.Fn()(c, cfile)
}

// #endregion

// #region IMMulticontext

type IMMulticontextObj struct {
	Parent IMContextObj
	Priv   uptr
}

type IMMulticontextClass struct {
	Parent IMContextClass
	_      [4]cc.Func
}

type IMMulticontext struct{ IMContext }

func GTypeIMMulticontext() gobject.GType { return gtk_im_multicontext_get_type.Fn()() }
func NewIMMulticontext() *IMMulticontext { return gtk_im_multicontext_new.Fn()() }
func (c *IMMulticontext) GetContextID() string {
	return gtk_im_multicontext_get_context_id.Fn()(c).String()
}
func (c *IMMulticontext) SetContextID(contextId string) {
	cid := cc.NewString(contextId)
	defer cid.Free()
	gtk_im_multicontext_set_context_id.Fn()(c, cid)
}

// #endregion

// #region Inscription

type Inscription struct{ Widget }

func GTypeInscription() gobject.GType { return gtk_inscription_get_type.Fn()() }
func NewInscription(text string) *Inscription {
	cText := cc.NewString(text)
	defer cText.Free()
	return gtk_inscription_new.Fn()(cText)
}
func (i *Inscription) GetText() string { return gtk_inscription_get_text.Fn()(i).String() }
func (i *Inscription) SetText(text string) {
	cText := cc.NewString(text)
	defer cText.Free()
	gtk_inscription_set_text.Fn()(i, cText)
}
func (i *Inscription) GetAttributes() *pango.AttrList { return gtk_inscription_get_attributes.Fn()(i) }
func (i *Inscription) SetAttributes(attrs *pango.AttrList) {
	gtk_inscription_set_attributes.Fn()(i, attrs)
}
func (i *Inscription) SetMarkup(markup string) {
	cMarkup := cc.NewString(markup)
	defer cMarkup.Free()
	gtk_inscription_set_markup.Fn()(i, cMarkup)
}
func (i *Inscription) GetTextOverflow() InscriptionOverflow {
	return gtk_inscription_get_text_overflow.Fn()(i)
}
func (i *Inscription) SetTextOverflow(overflow InscriptionOverflow) {
	gtk_inscription_set_text_overflow.Fn()(i, overflow)
}
func (i *Inscription) GetWrapMode() pango.WrapMode { return gtk_inscription_get_wrap_mode.Fn()(i) }
func (i *Inscription) SetWrapMode(wrapMode pango.WrapMode) {
	gtk_inscription_set_wrap_mode.Fn()(i, wrapMode)
}
func (i *Inscription) GetMinChars() uint32         { return gtk_inscription_get_min_chars.Fn()(i) }
func (i *Inscription) SetMinChars(minChars uint32) { gtk_inscription_set_min_chars.Fn()(i, minChars) }
func (i *Inscription) GetNatChars() uint32         { return gtk_inscription_get_nat_chars.Fn()(i) }
func (i *Inscription) SetNatChars(natChars uint32) { gtk_inscription_set_nat_chars.Fn()(i, natChars) }
func (i *Inscription) GetMinLines() uint32         { return gtk_inscription_get_min_lines.Fn()(i) }
func (i *Inscription) SetMinLines(minLines uint32) { gtk_inscription_set_min_lines.Fn()(i, minLines) }
func (i *Inscription) GetNatLines() uint32         { return gtk_inscription_get_nat_lines.Fn()(i) }
func (i *Inscription) SetNatLines(natLines uint32) { gtk_inscription_set_nat_lines.Fn()(i, natLines) }
func (i *Inscription) GetXAlign() float32          { return gtk_inscription_get_xalign.Fn()(i) }
func (i *Inscription) SetXAlign(xalign float32)    { gtk_inscription_set_xalign.Fn()(i, xalign) }
func (i *Inscription) GetYAlign() float32          { return gtk_inscription_get_yalign.Fn()(i) }
func (i *Inscription) SetYAlign(yalign float32)    { gtk_inscription_set_yalign.Fn()(i, yalign) }

// #endregion

// #region Label

type Label struct{ Widget }

func GTypeLabel() gobject.GType { return gtk_label_get_type.Fn()() }
func NewLabel(str string) *Label {
	cstr := cc.NewString(str)
	defer cstr.Free()
	return gtk_label_new.Fn()(cstr)
}
func NewLabelWithMnemonic(str string) *Label {
	cstr := cc.NewString(str)
	defer cstr.Free()
	return gtk_label_new_with_mnemonic.Fn()(cstr)
}
func (l *Label) SetText(str string) {
	cstr := cc.NewString(str)
	defer cstr.Free()
	gtk_label_set_text.Fn()(l, cstr)
}
func (l *Label) GetText() string                     { return gtk_label_get_text.Fn()(l).String() }
func (l *Label) SetAttributes(attrs *pango.AttrList) { gtk_label_set_attributes.Fn()(l, attrs) }
func (l *Label) GetAttributes() *pango.AttrList      { return gtk_label_get_attributes.Fn()(l) }
func (l *Label) SetLabel(str string) {
	cstr := cc.NewString(str)
	defer cstr.Free()
	gtk_label_set_label.Fn()(l, cstr)
}
func (l *Label) GetLabel() string { return gtk_label_get_label.Fn()(l).String() }
func (l *Label) SetMarkup(str string) {
	cstr := cc.NewString(str)
	defer cstr.Free()
	gtk_label_set_markup.Fn()(l, cstr)
}
func (l *Label) SetUseMarkup(setting bool)    { gtk_label_set_use_markup.Fn()(l, cbool(setting)) }
func (l *Label) GetUseMarkup() bool           { return gtk_label_get_use_markup.Fn()(l) != 0 }
func (l *Label) SetUseUnderline(setting bool) { gtk_label_set_use_underline.Fn()(l, cbool(setting)) }
func (l *Label) GetUseUnderline() bool        { return gtk_label_get_use_underline.Fn()(l) != 0 }

func (l *Label) SetMarkupWithMnemonic(str string) {
	cstr := cc.NewString(str)
	defer cstr.Free()
	gtk_label_set_markup_with_mnemonic.Fn()(l, cstr)
}
func (l *Label) GetMnemonicKeyval() uint32 { return gtk_label_get_mnemonic_keyval.Fn()(l) }
func (l *Label) SetMnemonicWidget(widget WidgetIface) {
	gtk_label_set_mnemonic_widget.Fn()(l, GetWidgetIface(widget))
}
func (l *Label) GetMnemonicWidget() *Widget { return gtk_label_get_mnemonic_widget.Fn()(l) }
func (l *Label) SetTextWithMnemonic(str string) {
	cstr := cc.NewString(str)
	defer cstr.Free()
	gtk_label_set_text_with_mnemonic.Fn()(l, cstr)
}
func (l *Label) SetJustify(jtype Justification)        { gtk_label_set_justify.Fn()(l, jtype) }
func (l *Label) GetJustify() Justification             { return gtk_label_get_justify.Fn()(l) }
func (l *Label) SetEllipsize(mode pango.EllipsizeMode) { gtk_label_set_ellipsize.Fn()(l, mode) }
func (l *Label) GetEllipsize() pango.EllipsizeMode     { return gtk_label_get_ellipsize.Fn()(l) }
func (l *Label) SetWidthChars(nChars int32)            { gtk_label_set_width_chars.Fn()(l, nChars) }
func (l *Label) GetWidthChars() int32                  { return gtk_label_get_width_chars.Fn()(l) }
func (l *Label) SetMaxWidthChars(nChars int32)         { gtk_label_set_max_width_chars.Fn()(l, nChars) }
func (l *Label) GetMaxWidthChars() int32               { return gtk_label_get_max_width_chars.Fn()(l) }
func (l *Label) SetLines(lines int32)                  { gtk_label_set_lines.Fn()(l, lines) }
func (l *Label) GetLines() int32                       { return gtk_label_get_lines.Fn()(l) }
func (l *Label) SetWrap(wrap bool)                     { gtk_label_set_wrap.Fn()(l, cbool(wrap)) }
func (l *Label) GetWrap() bool                         { return gtk_label_get_wrap.Fn()(l) != 0 }
func (l *Label) SetWrapMode(wrapMode pango.WrapMode)   { gtk_label_set_wrap_mode.Fn()(l, wrapMode) }
func (l *Label) GetWrapMode() pango.WrapMode           { return gtk_label_get_wrap_mode.Fn()(l) }
func (l *Label) SetNaturalWrapMode(wrapMode NaturalWrapMode) {
	gtk_label_set_natural_wrap_mode.Fn()(l, wrapMode)
}
func (l *Label) GetNaturalWrapMode() NaturalWrapMode { return gtk_label_get_natural_wrap_mode.Fn()(l) }

func (l *Label) SetSelectable(setting bool) { gtk_label_set_selectable.Fn()(l, cbool(setting)) }
func (l *Label) GetSelectable() bool        { return gtk_label_get_selectable.Fn()(l) != 0 }
func (l *Label) SelectRegion(startOffset, endOffset int32) {
	gtk_label_select_region.Fn()(l, startOffset, endOffset)
}
func (l *Label) GetSelectionBounds() (start, end int32, ok bool) {
	ok = gtk_label_get_selection_bounds.Fn()(l, &start, &end) != 0
	return
}
func (l *Label) GetLayout() *pango.Layout { return gtk_label_get_layout.Fn()(l) }
func (l *Label) GetLayoutOffsets() (x, y int32) {
	gtk_label_get_layout_offsets.Fn()(l, &x, &y)
	return
}
func (l *Label) SetSingleLineMode(singleLineMode bool) {
	gtk_label_set_single_line_mode.Fn()(l, cbool(singleLineMode))
}
func (l *Label) GetSingleLineMode() bool            { return gtk_label_get_single_line_mode.Fn()(l) != 0 }
func (l *Label) GetCurrentURI() string              { return gtk_label_get_current_uri.Fn()(l).String() }
func (l *Label) SetXAlign(xalign float32)           { gtk_label_set_xalign.Fn()(l, xalign) }
func (l *Label) GetXAlign() float32                 { return gtk_label_get_xalign.Fn()(l) }
func (l *Label) SetYAlign(yalign float32)           { gtk_label_set_yalign.Fn()(l, yalign) }
func (l *Label) GetYAlign() float32                 { return gtk_label_get_yalign.Fn()(l) }
func (l *Label) SetExtraMenu(model *gio.GMenuModel) { gtk_label_set_extra_menu.Fn()(l, model) }
func (l *Label) GetExtraMenu() *gio.GMenuModel      { return gtk_label_get_extra_menu.Fn()(l) }
func (l *Label) SetTabs(tabs *pango.TabArray)       { gtk_label_set_tabs.Fn()(l, tabs) }
func (l *Label) GetTabs() *pango.TabArray           { return gtk_label_get_tabs.Fn()(l) }
func (l *Label) ConnectActivateCurrentLink(sig func(l *Label)) uint64 {
	return l.SignalConnect("activate-current-link", sig, nil)
}
func (l *Label) ConnectActivateLink(sig func(l *Label, uri cc.String, _ uptr) bool) uint64 {
	return l.SignalConnect("activate-link", sig, nil)
}
func (l *Label) ConnectCopyClipboard(sig func(l *Label)) uint64 {
	return l.SignalConnect("copy-clipboard", sig, nil)
}
func (l *Label) ConnectMoveCursor(sig func(l *Label, step MovementStep, count int32, extendSelection bool)) uint64 {
	return l.SignalConnect("move-cursor", sig, nil)
}

// #endregion

// #region LayoutChild

type LayoutChild struct{ gobjCore }
type LayoutChildClass struct{ Parent gobject.GObjectClass }

func GTypeLayoutChild() gobject.GType { return gtk_layout_child_get_type.Fn()() }
func (l *LayoutChild) GetLayoutManager() *LayoutManager {
	return gtk_layout_child_get_layout_manager.Fn()(l)
}
func (l *LayoutChild) GetChildWidget() *Widget { return gtk_layout_child_get_child_widget.Fn()(l) }

// #endregion

// #region LayoutManager

type LayoutManagerClass struct {
	Parent gobject.GObjectClass

	GetRequestMode    cc.Func // GtkSizeRequestMode (* get_request_mode) (GtkLayoutManager *manager, GtkWidget *widget);
	Measure           cc.Func // void (* measure) (GtkLayoutManager *manager, GtkWidget *widget, GtkOrientation orientation, int for_size, int *minimum, int *natural, int *minimum_baseline, int *natural_baseline);
	Allocate          cc.Func // void (* allocate) (GtkLayoutManager *manager, GtkWidget *widget, int width, int height, int baseline);
	LayoutChildType   gobject.GType
	CreateLayoutChild cc.Func // GtkLayoutChild * (* create_layout_child) (GtkLayoutManager *manager, GtkWidget *widget, GtkWidget *for_child);
	Root              cc.Func // void (* root) (GtkLayoutManager *manager);
	Unroot            cc.Func // void (* unroot) (GtkLayoutManager *manager);

	_ [16]uptr //_padding
}

type LayoutManagerIface interface {
	GetLayoutManagerIface() *LayoutManager
}

func GetLayoutManagerIface(iface LayoutManagerIface) *LayoutManager {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetLayoutManagerIface()
}

func (l *LayoutManager) GetLayoutManagerIface() *LayoutManager { return l }

type LayoutManager struct{ gobjCore }

func GTypeLayoutManager() gobject.GType { return gtk_layout_manager_get_type.Fn()() }
func (l *LayoutManager) Measure(widget WidgetIface, orientation Orientation, for_size int) (minimum, natural, minimum_baseline, natural_baseline int) {
	gtk_layout_manager_measure.Fn()(l, GetWidgetIface(widget), orientation, for_size, &minimum, &natural, &minimum_baseline, &natural_baseline)
	return
}
func (l *LayoutManager) Allocate(widget WidgetIface, width, height, baseline int) {
	gtk_layout_manager_allocate.Fn()(l, GetWidgetIface(widget), width, height, baseline)
}
func (l *LayoutManager) GetRequestMode() SizeRequestMode {
	return gtk_layout_manager_get_request_mode.Fn()(l)
}
func (l *LayoutManager) GetWidget() *Widget {
	return gtk_layout_manager_get_widget.Fn()(l)
}
func (l *LayoutManager) LayoutChanged() {
	gtk_layout_manager_layout_changed.Fn()(l)
}
func (l *LayoutManager) GetLayoutChild(child WidgetIface) *LayoutChild {
	return gtk_layout_manager_get_layout_child.Fn()(l, GetWidgetIface(child))
}

// #endregion

// #region LevelBar

type LevelBar struct {
	Widget
	Orientable
}

func GTypeLevelBar() gobject.GType { return gtk_level_bar_get_type.Fn()() }
func NewLevelBar() *LevelBar       { return gtk_level_bar_new.Fn()() }
func NewLevelBarForInterval(minValue, maxValue float64) *LevelBar {
	return gtk_level_bar_new_for_interval.Fn()(minValue, maxValue)
}
func (l *LevelBar) SetMode(mode LevelBarMode) { gtk_level_bar_set_mode.Fn()(l, mode) }
func (l *LevelBar) GetMode() LevelBarMode     { return gtk_level_bar_get_mode.Fn()(l) }
func (l *LevelBar) SetValue(value float64)    { gtk_level_bar_set_value.Fn()(l, value) }
func (l *LevelBar) GetValue() float64         { return gtk_level_bar_get_value.Fn()(l) }
func (l *LevelBar) SetMinValue(value float64) { gtk_level_bar_set_min_value.Fn()(l, value) }
func (l *LevelBar) GetMinValue() float64      { return gtk_level_bar_get_min_value.Fn()(l) }
func (l *LevelBar) SetMaxValue(value float64) { gtk_level_bar_set_max_value.Fn()(l, value) }
func (l *LevelBar) GetMaxValue() float64      { return gtk_level_bar_get_max_value.Fn()(l) }
func (l *LevelBar) SetInverted(inverted bool) { gtk_level_bar_set_inverted.Fn()(l, cbool(inverted)) }
func (l *LevelBar) GetInverted() bool         { return gtk_level_bar_get_inverted.Fn()(l) != 0 }
func (l *LevelBar) AddOffsetValue(name string, value float64) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_level_bar_add_offset_value.Fn()(l, cName, value)
}
func (l *LevelBar) RemoveOffsetValue(name string) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_level_bar_remove_offset_value.Fn()(l, cName)
}
func (l *LevelBar) GetOffsetValue(name string) (float64, bool) {
	cName := cc.NewString(name)
	defer cName.Free()
	var value float64
	ok := gtk_level_bar_get_offset_value.Fn()(l, cName, &value) != 0
	return value, ok
}
func (l *LevelBar) ConnectOffsetChanged(sig func(l *LevelBar, name string)) uint64 {
	return l.SignalConnect("offset-changed", sig, nil)
}

// #endregion

// #region LinkButton

type LinkButton struct {
	Button
	Actionable
}

func GTypeLinkButton() gobject.GType { return gtk_link_button_get_type.Fn()() }
func NewLinkButton(uri string) *LinkButton {
	cUri := cc.NewString(uri)
	defer cUri.Free()
	return gtk_link_button_new.Fn()(cUri)
}
func NewLinkButtonWithLabel(uri, label string) *LinkButton {
	cUri, cLabel := cc.NewString(uri), cc.NewString(label)
	defer cUri.Free()
	defer cLabel.Free()
	return gtk_link_button_new_with_label.Fn()(cUri, cLabel)
}
func (lb *LinkButton) GetURI() string { return gtk_link_button_get_uri.Fn()(lb).String() }
func (lb *LinkButton) SetURI(uri string) {
	cUri := cc.NewString(uri)
	defer cUri.Free()
	gtk_link_button_set_uri.Fn()(lb, cUri)
}
func (lb *LinkButton) GetVisited() bool        { return gtk_link_button_get_visited.Fn()(lb) != 0 }
func (lb *LinkButton) SetVisited(visited bool) { gtk_link_button_set_visited.Fn()(lb, cbool(visited)) }

func (lb *LinkButton) ConnectActivateLink(sig func(lb *LinkButton, _ uptr) bool) uint64 {
	return lb.SignalConnect("activate-link", sig, nil)
}

// #endregion

// #region ListBase

type ListBaseIface interface {
	GetListBaseIface() *ListBase
}

func GetListBaseIface(iface ListBaseIface) *ListBase {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetListBaseIface()
}

func (lb *ListBase) GetListBaseIface() *ListBase { return lb }

type ListBase struct {
	Widget
	Orientable
	Scrollable
}

func GTypeListBase() gobject.GType { return gtk_list_base_get_type.Fn()() }

// #endregion

// #region ListBox

type ListBoxRowObj struct {
	Parent WidgetObj
}
type ListBoxRowClass struct {
	Parent   WidgetClass
	Activate cc.Func // void (* activate) (GtkListBoxRow *row);
	_        [8]uptr
}
type ListBoxRow struct{ Widget }

type ListBox struct{ Widget }

func GTypeListBoxRow() gobject.GType { return gtk_list_box_row_get_type.Fn()() }
func NewListBoxRow() *ListBoxRow     { return gtk_list_box_row_new.Fn()() }
func (r *ListBoxRow) SetChild(child WidgetIface) {
	gtk_list_box_row_set_child.Fn()(r, GetWidgetIface(child))
}
func (r *ListBoxRow) GetChild() *Widget  { return gtk_list_box_row_get_child.Fn()(r) }
func (r *ListBoxRow) GetHeader() *Widget { return gtk_list_box_row_get_header.Fn()(r) }
func (r *ListBoxRow) SetHeader(header WidgetIface) {
	gtk_list_box_row_set_header.Fn()(r, GetWidgetIface(header))
}
func (r *ListBoxRow) GetIndex() int32  { return gtk_list_box_row_get_index.Fn()(r) }
func (r *ListBoxRow) Changed()         { gtk_list_box_row_changed.Fn()(r) }
func (r *ListBoxRow) IsSelected() bool { return gtk_list_box_row_is_selected.Fn()(r) != 0 }
func (r *ListBoxRow) SetSelectable(selectable bool) {
	gtk_list_box_row_set_selectable.Fn()(r, cbool(selectable))
}
func (r *ListBoxRow) GetSelectable() bool { return gtk_list_box_row_get_selectable.Fn()(r) != 0 }
func (r *ListBoxRow) SetActivatable(activatable bool) {
	gtk_list_box_row_set_activatable.Fn()(r, cbool(activatable))
}
func (r *ListBoxRow) GetActivatable() bool { return gtk_list_box_row_get_activatable.Fn()(r) != 0 }
func (r *ListBoxRow) ConnectActivate(sig func(r *ListBoxRow)) uint64 {
	return r.SignalConnect("activate", sig, nil)
}

func GTypeListBox() gobject.GType            { return gtk_list_box_get_type.Fn()() }
func (b *ListBox) Prepend(child WidgetIface) { gtk_list_box_prepend.Fn()(b, GetWidgetIface(child)) }
func (b *ListBox) Append(child WidgetIface)  { gtk_list_box_append.Fn()(b, GetWidgetIface(child)) }
func (b *ListBox) Insert(child WidgetIface, position int32) {
	gtk_list_box_insert.Fn()(b, GetWidgetIface(child), position)
}
func (b *ListBox) Remove(child WidgetIface) { gtk_list_box_remove.Fn()(b, GetWidgetIface(child)) }
func (b *ListBox) RemoveAll()               { gtk_list_box_remove_all.Fn()(b) }

func (b *ListBox) GetSelectedRow() *ListBoxRow { return gtk_list_box_get_selected_row.Fn()(b) }
func (b *ListBox) GetRowAtIndex(index int32) *ListBoxRow {
	return gtk_list_box_get_row_at_index.Fn()(b, index)
}
func (b *ListBox) GetRowAtY(y int32) *ListBoxRow { return gtk_list_box_get_row_at_y.Fn()(b, y) }
func (b *ListBox) SelectRow(row *ListBoxRow)     { gtk_list_box_select_row.Fn()(b, row) }
func (b *ListBox) SetPlaceholder(placeholder WidgetIface) {
	gtk_list_box_set_placeholder.Fn()(b, GetWidgetIface(placeholder))
}
func (b *ListBox) SetAdjustment(adjustment *Adjustment) {
	gtk_list_box_set_adjustment.Fn()(b, adjustment)
}
func (b *ListBox) GetAdjustment() *Adjustment { return gtk_list_box_get_adjustment.Fn()(b) }
func (b *ListBox) SelectedForeach(fn func(box *ListBox, row *ListBoxRow)) {
	var cb uptr
	if fn != nil {
		cb = cc.CbkRaw[func(box *ListBox, row *ListBoxRow)](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 2)
			fn(*(**ListBox)(is[0]), *(**ListBoxRow)(is[1]))
		})
		defer cc.CbkClose(cb)
	}
	gtk_list_box_selected_foreach.Fn()(b, cb, nil)
}
func (b *ListBox) GetSelectedRows() *glib.GList[ListBoxRow] {
	return gtk_list_box_get_selected_rows.Fn()(b)
}
func (b *ListBox) UnselectRow(row *ListBoxRow)         { gtk_list_box_unselect_row.Fn()(b, row) }
func (b *ListBox) SelectAll()                          { gtk_list_box_select_all.Fn()(b) }
func (b *ListBox) UnselectAll()                        { gtk_list_box_unselect_all.Fn()(b) }
func (b *ListBox) SetSelectionMode(mode SelectionMode) { gtk_list_box_set_selection_mode.Fn()(b, mode) }
func (b *ListBox) GetSelectionMode() SelectionMode     { return gtk_list_box_get_selection_mode.Fn()(b) }
func (b *ListBox) SetFilterFunc(filterFunc func(row *ListBoxRow) bool) {
	var cb, des uptr
	if filterFunc != nil {
		cb = cc.CbkRaw[func(row *ListBoxRow, _ uptr) int32](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 2)
			if filterFunc(*(**ListBoxRow)(is[0])) {
				*(*int32)(out) = 1
			} else {
				*(*int32)(out) = 0
			}
		})
		des = cc.Cbk(func(_ uptr) { cc.CbkClose(cb); cc.CbkCloseLate(des) })
	}

	gtk_list_box_set_filter_func.Fn()(b, cb, nil, des)
}
func (b *ListBox) SetHeaderFunc(updateHeader func(row *ListBoxRow, before *ListBoxRow)) {
	var cb, des uptr
	if updateHeader != nil {
		cb = cc.CbkRaw[func(row *ListBoxRow, before *ListBoxRow, _ uptr)](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 2)
			updateHeader(*(**ListBoxRow)(is[0]), *(**ListBoxRow)(is[0]))
		})
		des = cc.Cbk(func(_ uptr) { cc.CbkClose(cb); cc.CbkCloseLate(des) })
	}

	gtk_list_box_set_header_func.Fn()(b, cb, nil, des)
}
func (b *ListBox) InvalidateFilter()  { gtk_list_box_invalidate_filter.Fn()(b) }
func (b *ListBox) InvalidateSort()    { gtk_list_box_invalidate_sort.Fn()(b) }
func (b *ListBox) InvalidateHeaders() { gtk_list_box_invalidate_headers.Fn()(b) }
func (b *ListBox) SetSortFunc(sortFunc func(row1, row2 *ListBoxRow) int32) {
	var cb, des uptr
	if sortFunc != nil {
		cb = cc.CbkRaw[func(row1, row2 *ListBoxRow, _ uptr) int32](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 3)
			*(*int32)(out) = sortFunc(*(**ListBoxRow)(is[0]), *(**ListBoxRow)(is[0]))
		})
		des = cc.Cbk(func(_ uptr) { cc.CbkClose(cb); cc.CbkCloseLate(des) })
	}

	gtk_list_box_set_sort_func.Fn()(b, cb, nil, des)
}
func (b *ListBox) SetActivateOnSingleClick(single bool) {
	gtk_list_box_set_activate_on_single_click.Fn()(b, cbool(single))
}
func (b *ListBox) GetActivateOnSingleClick() bool {
	return gtk_list_box_get_activate_on_single_click.Fn()(b) != 0
}
func (b *ListBox) DragUnhighlightRow()              { gtk_list_box_drag_unhighlight_row.Fn()(b) }
func (b *ListBox) DragHighlightRow(row *ListBoxRow) { gtk_list_box_drag_highlight_row.Fn()(b, row) }

func NewListBox() *ListBox { return gtk_list_box_new.Fn()() }
func (b *ListBox) BindModel(model gio.GListModelIface, createWidgetFunc func(item *gobject.GObject) *Widget) {
	var cb, des uptr
	if createWidgetFunc != nil {
		cc.CbkRaw[func(item *gobject.GObject, _ uptr) *Widget](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 3)
			*(**Widget)(out) = createWidgetFunc(*(**gobject.GObject)(is[0]))
		})
		des = cc.Cbk(func(d uptr) { cc.CbkClose(cb) })
	}
	gtk_list_box_bind_model.Fn()(b, gio.GetGListModelIface(model), cb, nil, des)
}
func (b *ListBox) SetShowSeparators(showSeparators bool) {
	gtk_list_box_set_show_separators.Fn()(b, cbool(showSeparators))
}
func (b *ListBox) GetShowSeparators() bool { return gtk_list_box_get_show_separators.Fn()(b) != 0 }
func (b *ListBox) SetTabBehavior(behavior ListTabBehavior) {
	gtk_list_box_set_tab_behavior.Fn()(b, behavior)
}
func (b *ListBox) GetTabBehavior() ListTabBehavior { return gtk_list_box_get_tab_behavior.Fn()(b) }
func (b *ListBox) ConnectActivateCursorRow(sig func(b *ListBox)) uint64 {
	return b.SignalConnect("activate-cursor-row", sig, nil)
}
func (b *ListBox) ConnectMoveCursor(
	sig func(r *ListBox, step MovementStep, count int32, extend, modify bool)) uint64 {
	return b.SignalConnect("move-cursor", sig, nil)
}
func (b *ListBox) ConnectRowActivated(sig func(b *ListBox, r *ListBoxRow)) uint64 {
	return b.SignalConnect("row-activated", sig, nil)
}
func (b *ListBox) ConnectRowSelected(sig func(b *ListBox, r *ListBoxRow)) uint64 {
	return b.SignalConnect("row-selected", sig, nil)
}
func (b *ListBox) ConnectSelectAll(sig func(b *ListBox)) uint64 {
	return b.SignalConnect("select-all", sig, nil)
}
func (b *ListBox) ConnectSelectedRowsChanged(sig func(b *ListBox)) uint64 {
	return b.SignalConnect("selected-rows-changed", sig, nil)
}
func (b *ListBox) ConnectToggleCursorRow(sig func(b *ListBox)) uint64 {
	return b.SignalConnect("toggle-cursor-row", sig, nil)
}
func (b *ListBox) ConnectUnselectAll(sig func(b *ListBox)) uint64 {
	return b.SignalConnect("unselect-all", sig, nil)
}

// #endregion

// #region ListHeader

type ListHeader struct{ gobjCore }

func GTypeListHeader() gobject.GType             { return gtk_list_header_get_type.Fn()() }
func (lh *ListHeader) GetItem() *gobject.GObject { return gtk_list_header_get_item.Fn()(lh) }
func (lh *ListHeader) GetStart() uint32          { return gtk_list_header_get_start.Fn()(lh) }
func (lh *ListHeader) GetEnd() uint32            { return gtk_list_header_get_end.Fn()(lh) }
func (lh *ListHeader) GetNItems() uint32         { return gtk_list_header_get_n_items.Fn()(lh) }
func (lh *ListHeader) SetChild(child WidgetIface) {
	gtk_list_header_set_child.Fn()(lh, GetWidgetIface(child))
}
func (lh *ListHeader) GetChild() *Widget { return gtk_list_header_get_child.Fn()(lh) }

// #endregion

// #region ListItem

type ListItemIface interface {
	GetListItemIface() *ListItem
}

func GetListItemIface(iface ListItemIface) *ListItem {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetListItemIface()
}
func (l *ListItem) GetListItemIface() *ListItem { return l }

type ListItem struct{ gobjCore }

func GTypeListItem() gobject.GType      { return gtk_list_item_get_type.Fn()() }
func (l *ListItem) GetItem() uptr       { return gtk_list_item_get_item.Fn()(l) }
func (l *ListItem) GetPosition() uint32 { return gtk_list_item_get_position.Fn()(l) }
func (l *ListItem) GetSelected() bool   { return gtk_list_item_get_selected.Fn()(l) != 0 }
func (l *ListItem) GetSelectable() bool { return gtk_list_item_get_selectable.Fn()(l) != 0 }
func (l *ListItem) SetSelectable(selectable bool) {
	gtk_list_item_set_selectable.Fn()(l, cbool(selectable))
}
func (l *ListItem) GetActivatable() bool { return gtk_list_item_get_activatable.Fn()(l) != 0 }
func (l *ListItem) SetActivatable(activatable bool) {
	gtk_list_item_set_activatable.Fn()(l, cbool(activatable))
}
func (l *ListItem) GetFocusable() bool { return gtk_list_item_get_focusable.Fn()(l) != 0 }
func (l *ListItem) SetFocusable(focusable bool) {
	gtk_list_item_set_focusable.Fn()(l, cbool(focusable))
}
func (l *ListItem) SetChild(child WidgetIface) {
	gtk_list_item_set_child.Fn()(l, GetWidgetIface(child))
}
func (l *ListItem) GetChild() *Widget { return gtk_list_item_get_child.Fn()(l) }
func (l *ListItem) SetAccessibleDescription(description string) {
	des := cc.NewString(description)
	defer des.Free()
	gtk_list_item_set_accessible_description.Fn()(l, des)
}
func (l *ListItem) GetAccessibleDescription() string {
	return gtk_list_item_get_accessible_description.Fn()(l).String()
}
func (l *ListItem) SetAccessibleLabel(label string) {
	lab := cc.NewString(label)
	defer lab.Free()
	gtk_list_item_set_accessible_label.Fn()(l, lab)
}
func (l *ListItem) GetAccessibleLabel() string {
	return gtk_list_item_get_accessible_label.Fn()(l).String()
}

// #endregion

// #region ListItemFactory

type ListItemFactoryIface interface {
	GetListItemFactoryIface() *ListItemFactory
}

func GetListItemFactoryIface(iface ListItemFactoryIface) *ListItemFactory {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetListItemFactoryIface()
}
func (l *ListItemFactory) GetListItemFactoryIface() *ListItemFactory { return l }

type ListItemFactory struct {
	gobjCore
}

func GTypeListItemFactory() gobject.GType { return gtk_list_item_factory_get_type.Fn()() }

// #endregion

// #region ListView

type ListView struct{ ListBase }

func GTypeListView() gobject.GType { return gtk_list_view_get_type.Fn()() }
func NewListView(model SelectionModelIface, factory ListItemFactoryIface) *ListView {
	return gtk_list_view_new.Fn()(GetSelectionModelIface(model), GetListItemFactoryIface(factory))
}
func (lv *ListView) GetModel() *SelectionModel { return gtk_list_view_get_model.Fn()(lv) }
func (lv *ListView) SetModel(model SelectionModelIface) {
	gtk_list_view_set_model.Fn()(lv, GetSelectionModelIface(model))
}
func (lv *ListView) SetFactory(factory ListItemFactoryIface) {
	gtk_list_view_set_factory.Fn()(lv, GetListItemFactoryIface(factory))
}
func (lv *ListView) GetFactory() *ListItemFactory { return gtk_list_view_get_factory.Fn()(lv) }
func (lv *ListView) SetHeaderFactory(factory ListItemFactoryIface) {
	gtk_list_view_set_header_factory.Fn()(lv, GetListItemFactoryIface(factory))
}
func (lv *ListView) GetHeaderFactory() *ListItemFactory {
	return gtk_list_view_get_header_factory.Fn()(lv)
}
func (lv *ListView) SetShowSeparators(showSeparators bool) {
	gtk_list_view_set_show_separators.Fn()(lv, cbool(showSeparators))
}
func (lv *ListView) GetShowSeparators() bool { return gtk_list_view_get_show_separators.Fn()(lv) != 0 }
func (lv *ListView) SetSingleClickActivate(singleClickActivate bool) {
	gtk_list_view_set_single_click_activate.Fn()(lv, cbool(singleClickActivate))
}
func (lv *ListView) GetSingleClickActivate() bool {
	return gtk_list_view_get_single_click_activate.Fn()(lv) != 0
}
func (lv *ListView) SetEnableRubberband(enableRubberband bool) {
	gtk_list_view_set_enable_rubberband.Fn()(lv, cbool(enableRubberband))
}
func (lv *ListView) GetEnableRubberband() bool {
	return gtk_list_view_get_enable_rubberband.Fn()(lv) != 0
}
func (lv *ListView) SetTabBehavior(tabBehavior ListTabBehavior) {
	gtk_list_view_set_tab_behavior.Fn()(lv, tabBehavior)
}
func (lv *ListView) GetTabBehavior() ListTabBehavior { return gtk_list_view_get_tab_behavior.Fn()(lv) }
func (lv *ListView) ScrollTo(pos uint32, flags ListScrollFlags, scroll *ScrollInfo) {
	gtk_list_view_scroll_to.Fn()(lv, pos, flags, scroll)
}
func (lv *ListView) ConnectActivate(sig func(lv *ListView, pos uint32)) uint64 {
	return lv.SignalConnect("activate", sig, nil)
}

// #endregion

// #region Main

func Init()               { gtk_init.Fn()() }
func InitCheck() bool     { return gtk_init_check.Fn()() != 0 }
func IsInitialized() bool { return gtk_is_initialized.Fn()() != 0 }
func InitAbiCheck(numChecks int32, sizeofWindow, sizeofBox uintptr) {
	gtk_init_abi_check.Fn()(numChecks, sizeofWindow, sizeofBox)
}
func InitCheckAbiCheck(numChecks int32, sizeofWindow, sizeofBox uintptr) bool {
	return gtk_init_check_abi_check.Fn()(numChecks, sizeofWindow, sizeofBox) != 0
}
func DisableSetlocale()                   { gtk_disable_setlocale.Fn()() }
func DisablePortals()                     { gtk_disable_portals.Fn()() }
func GetDefaultLanguage() *pango.Language { return gtk_get_default_language.Fn()() }
func GetLocaleDirection() TextDirection   { return gtk_get_locale_direction.Fn()() }

// #endregion

// #region MapListModel

type MapListModel struct {
	gobjCore
	SectionModel
}

func GTypeMapListModel() gobject.GType { return gtk_map_list_model_get_type.Fn()() }
func NewMapListModel[A, B any](model gio.GListModelIface, mapFunc func(item *A) *B) *MapListModel {
	var cb, des uptr
	if mapFunc != nil {
		cb = cc.CbkRaw[func(item, _ uptr) uptr](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 2)
			*(**B)(out) = mapFunc(*(**A)(is[0]))
		})
		des = cc.Cbk(func(_ uptr) { cc.CbkClose(cb); cc.CbkCloseLate(des) })
	}
	return gtk_map_list_model_new.Fn()(gio.GetGListModelIface(model), cb, nil, des)
}
func (m *MapListModel) SetMapFunc(mapFunc func(item uptr) uptr) {

	var cb, des uptr
	if mapFunc != nil {
		cb = cc.CbkRaw[func(item, _ uptr) uptr](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 2)
			*(*uptr)(out) = mapFunc(*(*uptr)(is[0]))
		})
		des = cc.Cbk(func(_ uptr) { cc.CbkClose(cb); cc.CbkCloseLate(des) })
	}
	gtk_map_list_model_set_map_func.Fn()(m, cb, nil, des)
}
func (m *MapListModel) SetModel(model gio.GListModelIface) {
	gtk_map_list_model_set_model.Fn()(m, gio.GetGListModelIface(model))
}
func (m *MapListModel) GetModel() *gio.GListModel { return gtk_map_list_model_get_model.Fn()(m) }
func (m *MapListModel) HasMap() bool              { return gtk_map_list_model_has_map.Fn()(m) != 0 }

// #endregion

// #region MediaControls

type MediaControls struct{ Widget }

func GTypeMediaControls() gobject.GType { return gtk_media_controls_get_type.Fn()() }
func NewMediaControls(stream MediaStreamIface) *MediaControls {
	return gtk_media_controls_new.Fn()(GetMediaStreamIface(stream))
}
func (c *MediaControls) GetMediaStream() *MediaStream {
	return gtk_media_controls_get_media_stream.Fn()(c)
}
func (c *MediaControls) SetMediaStream(stream MediaStreamIface) {
	gtk_media_controls_set_media_stream.Fn()(c, GetMediaStreamIface(stream))
}

// #endregion

// #region MediaFile

type MediaFileClass struct {
	Parent MediaStreamClass

	Open  cc.Func //void(* open)(GtkMediaFile *self);
	Close cc.Func //void(* close)(GtkMediaFile *self);

	_ [8]cc.Func //reserved
}

type MediaFile struct {
	MediaStream
	gdk.Paintable
}

func GTypeMediaFile() gobject.GType { return gtk_media_file_get_type.Fn()() }
func NewMediaFile() *MediaStream    { return gtk_media_file_new.Fn()() }
func NewMediaFileForFilename(filename string) *MediaStream {
	cFilename := cc.NewString(filename)
	defer cFilename.Free()
	return gtk_media_file_new_for_filename.Fn()(cFilename)
}
func NewMediaFileForResource(resourcePath string) *MediaStream {
	cResource := cc.NewString(resourcePath)
	defer cResource.Free()
	return gtk_media_file_new_for_resource.Fn()(cResource)
}
func NewMediaFileForFile(file *gio.GFile) *MediaStream {
	return gtk_media_file_new_for_file.Fn()(file)
}
func NewMediaFileForInputStream(stream gio.GInputStreamIface) *MediaStream {
	return gtk_media_file_new_for_input_stream.Fn()(gio.GetGInputStreamIface(stream))
}
func (mf *MediaFile) Clear() { gtk_media_file_clear.Fn()(mf) }
func (mf *MediaFile) SetFilename(filename string) {
	cFilename := cc.NewString(filename)
	defer cFilename.Free()
	gtk_media_file_set_filename.Fn()(mf, cFilename)
}
func (mf *MediaFile) SetResource(resourcePath string) {
	cResource := cc.NewString(resourcePath)
	defer cResource.Free()
	gtk_media_file_set_resource.Fn()(mf, cResource)
}
func (mf *MediaFile) SetFile(file *gio.GFile) { gtk_media_file_set_file.Fn()(mf, file) }
func (mf *MediaFile) GetFile() *gio.GFile     { return gtk_media_file_get_file.Fn()(mf) }
func (mf *MediaFile) SetInputStream(stream gio.GInputStreamIface) {
	gtk_media_file_set_input_stream.Fn()(mf, gio.GetGInputStreamIface(stream))
}
func (mf *MediaFile) GetInputStream() *gio.GInputStream {
	return gtk_media_file_get_input_stream.Fn()(mf)
}

// #endregion

// #region MediaStream

type MediaStreamClass struct {
	Parent gobject.GObjectClass

	Play        cc.Func // gboolean (* play) (GtkMediaStream *self);
	Pause       cc.Func // void (* pause) (GtkMediaStream *self);
	Seek        cc.Func // void (* seek) (GtkMediaStream *self,gint64 timestamp);
	UpdateAudio cc.Func // void (* update_audio) (GtkMediaStream *self,gboolean muted,double volume);
	Realize     cc.Func // void (* realize) (GtkMediaStream *self,GdkSurface *surface);
	Unrealize   cc.Func // void (* unrealize) (GtkMediaStream *self,GdkSurface *surface);

	_ [8]cc.Func //reserved
}

type MediaStreamIface interface {
	GetMediaStreamIface() *MediaStream
}

func GetMediaStreamIface(iface MediaStreamIface) *MediaStream {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetMediaStreamIface()
}

func (m *MediaStream) GetMediaStreamIface() *MediaStream { return m }

type MediaStream struct {
	gobjCore
}

func GTypeMediaStream() gobject.GType          { return gtk_media_stream_get_type.Fn()() }
func (ms *MediaStream) IsPrepared() bool       { return gtk_media_stream_is_prepared.Fn()(ms) != 0 }
func (ms *MediaStream) GetError() *glib.GError { return gtk_media_stream_get_error.Fn()(ms) }
func (ms *MediaStream) HasAudio() bool         { return gtk_media_stream_has_audio.Fn()(ms) != 0 }
func (ms *MediaStream) HasVideo() bool         { return gtk_media_stream_has_video.Fn()(ms) != 0 }
func (ms *MediaStream) Play()                  { gtk_media_stream_play.Fn()(ms) }
func (ms *MediaStream) Pause()                 { gtk_media_stream_pause.Fn()(ms) }
func (ms *MediaStream) GetPlaying() bool       { return gtk_media_stream_get_playing.Fn()(ms) != 0 }
func (ms *MediaStream) SetPlaying(playing bool) {
	gtk_media_stream_set_playing.Fn()(ms, cbool(playing))
}
func (ms *MediaStream) GetEnded() bool                 { return gtk_media_stream_get_ended.Fn()(ms) != 0 }
func (ms *MediaStream) GetTimestamp() int64            { return gtk_media_stream_get_timestamp.Fn()(ms) }
func (ms *MediaStream) GetDuration() int64             { return gtk_media_stream_get_duration.Fn()(ms) }
func (ms *MediaStream) IsSeekable() bool               { return gtk_media_stream_is_seekable.Fn()(ms) != 0 }
func (ms *MediaStream) IsSeeking() bool                { return gtk_media_stream_is_seeking.Fn()(ms) != 0 }
func (ms *MediaStream) Seek(timestamp int64)           { gtk_media_stream_seek.Fn()(ms, timestamp) }
func (ms *MediaStream) GetLoop() bool                  { return gtk_media_stream_get_loop.Fn()(ms) != 0 }
func (ms *MediaStream) SetLoop(loop bool)              { gtk_media_stream_set_loop.Fn()(ms, cbool(loop)) }
func (ms *MediaStream) GetMuted() bool                 { return gtk_media_stream_get_muted.Fn()(ms) != 0 }
func (ms *MediaStream) SetMuted(muted bool)            { gtk_media_stream_set_muted.Fn()(ms, cbool(muted)) }
func (ms *MediaStream) GetVolume() float64             { return gtk_media_stream_get_volume.Fn()(ms) }
func (ms *MediaStream) SetVolume(volume float64)       { gtk_media_stream_set_volume.Fn()(ms, volume) }
func (ms *MediaStream) Realize(surface *gdk.Surface)   { gtk_media_stream_realize.Fn()(ms, surface) }
func (ms *MediaStream) Unrealize(surface *gdk.Surface) { gtk_media_stream_unrealize.Fn()(ms, surface) }
func (ms *MediaStream) StreamPrepared(hasAudio, hasVideo, seekable bool, duration int64) {
	gtk_media_stream_stream_prepared.Fn()(ms, cbool(hasAudio), cbool(hasVideo), cbool(seekable), duration)
}
func (ms *MediaStream) StreamUnprepared()      { gtk_media_stream_stream_unprepared.Fn()(ms) }
func (ms *MediaStream) Update(timestamp int64) { gtk_media_stream_update.Fn()(ms, timestamp) }
func (ms *MediaStream) StreamEnded()           { gtk_media_stream_stream_ended.Fn()(ms) }
func (ms *MediaStream) SeekSuccess()           { gtk_media_stream_seek_success.Fn()(ms) }
func (ms *MediaStream) SeekFailed()            { gtk_media_stream_seek_failed.Fn()(ms) }

// func (ms *MediaStream) GError(err *glib.GError) { gtk_media_stream_gerror(ms, err) }

func (ms *MediaStream) Error(domain glib.GQuark, code int32, format string, args ...interface{}) {
	cFormat := cc.NewString(format)
	defer cFormat.Free()
	gtk_media_stream_error.FnVa()(ms, domain, code, cFormat, args...)
}

// #endregion

// #region MenuButton

type MenuButton struct{ Widget }

func GTypeMenuButton() gobject.GType { return gtk_menu_button_get_type.Fn()() }
func NewMenuButton() *MenuButton     { return gtk_menu_button_new.Fn()() }
func (mb *MenuButton) SetPopover(popover *Popover) {
	gtk_menu_button_set_popover.Fn()(mb, popover)
}
func (mb *MenuButton) GetPopover() *Popover { return gtk_menu_button_get_popover.Fn()(mb) }
func (mb *MenuButton) SetDirection(direction ArrowType) {
	gtk_menu_button_set_direction.Fn()(mb, direction)
}
func (mb *MenuButton) GetDirection() ArrowType { return gtk_menu_button_get_direction.Fn()(mb) }
func (mb *MenuButton) SetMenuModel(menuModel *gio.GMenuModel) {
	gtk_menu_button_set_menu_model.Fn()(mb, menuModel)
}
func (mb *MenuButton) GetMenuModel() *gio.GMenuModel { return gtk_menu_button_get_menu_model.Fn()(mb) }
func (mb *MenuButton) SetIconName(iconName string) {
	cIcon := cc.NewString(iconName)
	defer cIcon.Free()
	gtk_menu_button_set_icon_name.Fn()(mb, cIcon)
}
func (mb *MenuButton) GetIconName() string { return gtk_menu_button_get_icon_name.Fn()(mb).String() }
func (mb *MenuButton) SetAlwaysShowArrow(alwaysShowArrow bool) {
	gtk_menu_button_set_always_show_arrow.Fn()(mb, cbool(alwaysShowArrow))
}
func (mb *MenuButton) GetAlwaysShowArrow() bool {
	return gtk_menu_button_get_always_show_arrow.Fn()(mb) != 0
}
func (mb *MenuButton) SetLabel(label string) {
	cLabel := cc.NewString(label)
	defer cLabel.Free()
	gtk_menu_button_set_label.Fn()(mb, cLabel)
}
func (mb *MenuButton) GetLabel() string { return gtk_menu_button_get_label.Fn()(mb).String() }
func (mb *MenuButton) SetUseUnderline(useUnderline bool) {
	gtk_menu_button_set_use_underline.Fn()(mb, cbool(useUnderline))
}
func (mb *MenuButton) GetUseUnderline() bool { return gtk_menu_button_get_use_underline.Fn()(mb) != 0 }
func (mb *MenuButton) SetHasFrame(hasFrame bool) {
	gtk_menu_button_set_has_frame.Fn()(mb, cbool(hasFrame))
}
func (mb *MenuButton) GetHasFrame() bool { return gtk_menu_button_get_has_frame.Fn()(mb) != 0 }
func (mb *MenuButton) Popup()            { gtk_menu_button_popup.Fn()(mb) }
func (mb *MenuButton) Popdown()          { gtk_menu_button_popdown.Fn()(mb) }
func (mb *MenuButton) SetCreatePopupFunc(fn func(mb *MenuButton)) {
	var cb, des uptr
	if fn != nil {
		cb = cc.Cbk(fn)
		des = cc.Cbk(func(_ uptr) { cc.CbkClose(cb); cc.CbkCloseLate(des) })
	}
	gtk_menu_button_set_create_popup_func.Fn()(mb, cb, nil, nil)
}
func (mb *MenuButton) SetPrimary(primary bool) { gtk_menu_button_set_primary.Fn()(mb, cbool(primary)) }
func (mb *MenuButton) GetPrimary() bool        { return gtk_menu_button_get_primary.Fn()(mb) != 0 }
func (mb *MenuButton) SetChild(child WidgetIface) {
	gtk_menu_button_set_child.Fn()(mb, GetWidgetIface(child))
}
func (mb *MenuButton) GetChild() *Widget     { return gtk_menu_button_get_child.Fn()(mb) }
func (mb *MenuButton) SetActive(active bool) { gtk_menu_button_set_active.Fn()(mb, cbool(active)) }
func (mb *MenuButton) GetActive() bool       { return gtk_menu_button_get_active.Fn()(mb) != 0 }
func (mb *MenuButton) SetCanShrink(canShrink bool) {
	gtk_menu_button_set_can_shrink.Fn()(mb, cbool(canShrink))
}
func (mb *MenuButton) GetCanShrink() bool { return gtk_menu_button_get_can_shrink.Fn()(mb) != 0 }
func (mb *MenuButton) ConnectActivate(sig func(mb *MenuButton)) uint64 {
	return mb.SignalConnect("activate", sig, nil)
}

// #endregion

// #region MountOperation

type MountOperationObj struct {
	Parent gio.GMountOperationObj
	Priv   uptr
}

type MountOperationClass struct {
	Parent gio.GMountOperationClass
	_      [4]uptr //reserved
}

type MountOperation struct {
	gio.GMountOperation
}

func GTypeMountOperation() gobject.GType { return gtk_mount_operation_get_type.Fn()() }
func NewMountOperation(parent WindowIface) *MountOperation {
	return gtk_mount_operation_new.Fn()(GetWindowIface(parent))
}
func (op *MountOperation) IsShowing() bool { return gtk_mount_operation_is_showing.Fn()(op) != 0 }
func (op *MountOperation) SetParent(parent WindowIface) {
	gtk_mount_operation_set_parent.Fn()(op, GetWindowIface(parent))
}
func (op *MountOperation) GetParent() *Window { return gtk_mount_operation_get_parent.Fn()(op) }
func (op *MountOperation) SetDisplay(display *gdk.Display) {
	gtk_mount_operation_set_display.Fn()(op, display)
}
func (op *MountOperation) GetDisplay() *gdk.Display { return gtk_mount_operation_get_display.Fn()(op) }

// #endregion

// #region MultiFilter

type MultiFilterIface interface {
	GetMultiFilterIface() *MultiFilter
}

func GetMultiFilterIface(iface MultiFilterIface) *MultiFilter {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetMultiFilterIface()
}
func (m *MultiFilter) GetMultiFilterIface() *MultiFilter { return m }

type MultiFilter struct {
	Filter
	gio.GListModel
	Buildable
}

type AnyFilter struct{ MultiFilter }
type EveryFilter struct{ MultiFilter }

func GTypeMultiFilter() gobject.GType { return gtk_multi_filter_get_type.Fn()() }
func (m *MultiFilter) Append(filter FilterIface) {
	gtk_multi_filter_append.Fn()(m, GetFilterIface(filter))
}
func (m *MultiFilter) Remove(position uint32) { gtk_multi_filter_remove.Fn()(m, position) }
func GTypeAnyFilter() gobject.GType           { return gtk_any_filter_get_type.Fn()() }
func NewAnyFilter() *AnyFilter                { return gtk_any_filter_new.Fn()() }
func GTypeEveryFilter() gobject.GType         { return gtk_every_filter_get_type.Fn()() }
func NewEveryFilter() *EveryFilter            { return gtk_every_filter_new.Fn()() }

// #endregion

// #region MultiSelection

type MultiSelection struct {
	gobjCore

	gio.GListModel
	SectionModel
	SelectionModel
}

func GTypeMultiSelection() gobject.GType { return gtk_multi_selection_get_type.Fn()() }
func NewMultiSelection(model gio.GListModelIface) *MultiSelection {
	return gtk_multi_selection_new.Fn()(gio.GetGListModelIface(model))
}
func (m *MultiSelection) GetModel() *gio.GListModel { return gtk_multi_selection_get_model.Fn()(m) }
func (m *MultiSelection) SetModel(model gio.GListModelIface) {
	gtk_multi_selection_set_model.Fn()(m, gio.GetGListModelIface(model))
}

// #endregion

// #region MultiSorter

type MultiSorter struct {
	Sorter
	gio.GListModel
	Buildable
}

func GTypeMultiSorter() gobject.GType          { return gtk_multi_sorter_get_type.Fn()() }
func NewMultiSorter() *MultiSorter             { return gtk_multi_sorter_new.Fn()() }
func (ms *MultiSorter) Append(s SorterIface)   { gtk_multi_sorter_append.Fn()(ms, GetSorterIface(s)) }
func (ms *MultiSorter) Remove(position uint32) { gtk_multi_sorter_remove.Fn()(ms, position) }

// #endregion

// #region Native

type Native struct{}

func (n *Native) Realize()   { gtk_native_realize.Fn()(n) }
func (n *Native) Unrealize() { gtk_native_unrealize.Fn()(n) }
func GetNativeForSurface(surface *gdk.Surface) *Native {
	return gtk_native_get_for_surface.Fn()(surface)
}
func (n *Native) GetSurface() *gdk.Surface   { return gtk_native_get_surface.Fn()(n) }
func (n *Native) GetRenderer() *gsk.Renderer { return gtk_native_get_renderer.Fn()(n) }
func (n *Native) GetSurfaceTransform() (x, y float64) {
	gtk_native_get_surface_transform.Fn()(n, &x, &y)
	return
}

// #endregion

// #region NativeDialog

type NativeDialogClass struct {
	Parent gobject.GObjectClass

	Response cc.Func // void (* response) (GtkNativeDialog *self, int response_id);
	Show     cc.Func // void (* show) (GtkNativeDialog *self);
	Hide     cc.Func // void (* hide) (GtkNativeDialog *self);

	_ [4]cc.Func // reserved
}

type NativeDialog struct {
	gobjCore
}

func GTypeNativeDialog() gobject.GType       { return gtk_native_dialog_get_type.Fn()() }
func (nd *NativeDialog) Show()               { gtk_native_dialog_show.Fn()(nd) }
func (nd *NativeDialog) Hide()               { gtk_native_dialog_hide.Fn()(nd) }
func (nd *NativeDialog) Destroy()            { gtk_native_dialog_destroy.Fn()(nd) }
func (nd *NativeDialog) GetVisible() bool    { return gtk_native_dialog_get_visible.Fn()(nd) != 0 }
func (nd *NativeDialog) SetModal(modal bool) { gtk_native_dialog_set_modal.Fn()(nd, cbool(modal)) }
func (nd *NativeDialog) GetModal() bool      { return gtk_native_dialog_get_modal.Fn()(nd) != 0 }
func (nd *NativeDialog) SetTitle(title string) {
	cTitle := cc.NewString(title)
	defer cTitle.Free()
	gtk_native_dialog_set_title.Fn()(nd, cTitle)
}
func (nd *NativeDialog) GetTitle() string { return gtk_native_dialog_get_title.Fn()(nd).String() }
func (nd *NativeDialog) SetTransientFor(parent WindowIface) {
	gtk_native_dialog_set_transient_for.Fn()(nd, GetWindowIface(parent))
}
func (nd *NativeDialog) GetTransientFor() *Window {
	return gtk_native_dialog_get_transient_for.Fn()(nd)
}
func (nd *NativeDialog) ConnectResponse(sig func(nd *NativeDialog, responseId int32)) uint64 {
	return nd.SignalConnect("response", sig, nil)
}

// #endregion

// #region NoSelection

type NoSelection struct {
	gobjCore
	gio.GListModel
	SectionModel
	SelectionModel
}

func GTypeNoSelection() gobject.GType { return gtk_no_selection_get_type.Fn()() }
func NewNoSelection(model gio.GListModelIface) *NoSelection {
	return gtk_no_selection_new.Fn()(gio.GetGListModelIface(model))
}
func (ns *NoSelection) GetModel() *gio.GListModel { return gtk_no_selection_get_model.Fn()(ns) }
func (ns *NoSelection) SetModel(model gio.GListModelIface) {
	gtk_no_selection_set_model.Fn()(ns, gio.GetGListModelIface(model))
}

// #endregion

// #region Notebook

type Notebook struct{ Widget }

func GTypeNotebook() gobject.GType { return gtk_notebook_get_type.Fn()() }
func NewNotebook() *Notebook       { return gtk_notebook_new.Fn()() }
func (nb *Notebook) AppendPage(child, tabLabel WidgetIface) int32 {
	return gtk_notebook_append_page.Fn()(nb, GetWidgetIface(child), GetWidgetIface(tabLabel))
}
func (nb *Notebook) AppendPageMenu(child, tabLabel, menuLabel WidgetIface) int32 {
	return gtk_notebook_append_page_menu.Fn()(nb, GetWidgetIface(child), GetWidgetIface(tabLabel), GetWidgetIface(menuLabel))
}
func (nb *Notebook) PrependPage(child, tabLabel WidgetIface) int32 {
	return gtk_notebook_prepend_page.Fn()(nb, GetWidgetIface(child), GetWidgetIface(tabLabel))
}
func (nb *Notebook) PrependPageMenu(child, tabLabel, menuLabel WidgetIface) int32 {
	return gtk_notebook_prepend_page_menu.Fn()(nb, GetWidgetIface(child), GetWidgetIface(tabLabel), GetWidgetIface(menuLabel))
}
func (nb *Notebook) InsertPage(child, tabLabel WidgetIface, position int32) int32 {
	return gtk_notebook_insert_page.Fn()(nb, GetWidgetIface(child), GetWidgetIface(tabLabel), position)
}
func (nb *Notebook) InsertPageMenu(child, tabLabel, menuLabel WidgetIface, position int32) int32 {
	return gtk_notebook_insert_page_menu.Fn()(nb, GetWidgetIface(child), GetWidgetIface(tabLabel), GetWidgetIface(menuLabel), position)
}
func (nb *Notebook) RemovePage(pageNum int32) { gtk_notebook_remove_page.Fn()(nb, pageNum) }
func (nb *Notebook) SetGroupName(groupName string) {
	cGroupName := cc.NewString(groupName)
	defer cGroupName.Free()
	gtk_notebook_set_group_name.Fn()(nb, cGroupName)
}
func (nb *Notebook) GetGroupName() string { return gtk_notebook_get_group_name.Fn()(nb).String() }

func (nb *Notebook) GetCurrentPage() int32 { return gtk_notebook_get_current_page.Fn()(nb) }
func (nb *Notebook) GetNthPage(pageNum int32) *Widget {
	return gtk_notebook_get_nth_page.Fn()(nb, pageNum)
}
func (nb *Notebook) GetNPages() int32 { return gtk_notebook_get_n_pages.Fn()(nb) }
func (nb *Notebook) PageNum(child WidgetIface) int32 {
	return gtk_notebook_page_num.Fn()(nb, GetWidgetIface(child))
}
func (nb *Notebook) SetCurrentPage(pageNum int32) { gtk_notebook_set_current_page.Fn()(nb, pageNum) }
func (nb *Notebook) NextPage()                    { gtk_notebook_next_page.Fn()(nb) }
func (nb *Notebook) PrevPage()                    { gtk_notebook_prev_page.Fn()(nb) }
func (nb *Notebook) SetShowBorder(showBorder bool) {
	gtk_notebook_set_show_border.Fn()(nb, cbool(showBorder))
}
func (nb *Notebook) GetShowBorder() bool        { return gtk_notebook_get_show_border.Fn()(nb) != 0 }
func (nb *Notebook) SetShowTabs(showTabs bool)  { gtk_notebook_set_show_tabs.Fn()(nb, cbool(showTabs)) }
func (nb *Notebook) GetShowTabs() bool          { return gtk_notebook_get_show_tabs.Fn()(nb) != 0 }
func (nb *Notebook) SetTabPos(pos PositionType) { gtk_notebook_set_tab_pos.Fn()(nb, pos) }
func (nb *Notebook) GetTabPos() PositionType    { return gtk_notebook_get_tab_pos.Fn()(nb) }
func (nb *Notebook) SetScrollable(scrollable bool) {
	gtk_notebook_set_scrollable.Fn()(nb, cbool(scrollable))
}
func (nb *Notebook) GetScrollable() bool { return gtk_notebook_get_scrollable.Fn()(nb) != 0 }
func (nb *Notebook) PopupEnable()        { gtk_notebook_popup_enable.Fn()(nb) }
func (nb *Notebook) PopupDisable()       { gtk_notebook_popup_disable.Fn()(nb) }
func (nb *Notebook) GetTabLabel(child WidgetIface) *Widget {
	return gtk_notebook_get_tab_label.Fn()(nb, GetWidgetIface(child))
}
func (nb *Notebook) SetTabLabel(child, tabLabel WidgetIface) {
	gtk_notebook_set_tab_label.Fn()(nb, GetWidgetIface(child), GetWidgetIface(tabLabel))
}
func (nb *Notebook) SetTabLabelText(child WidgetIface, tabText string) {
	cTabText := cc.NewString(tabText)
	defer cTabText.Free()
	gtk_notebook_set_tab_label_text.Fn()(nb, GetWidgetIface(child), cTabText)
}
func (nb *Notebook) GetTabLabelText(child WidgetIface) string {
	return gtk_notebook_get_tab_label_text.Fn()(nb, GetWidgetIface(child)).String()
}
func (nb *Notebook) GetMenuLabel(child WidgetIface) *Widget {
	return gtk_notebook_get_menu_label.Fn()(nb, GetWidgetIface(child))
}
func (nb *Notebook) SetMenuLabel(child, menuLabel WidgetIface) {
	gtk_notebook_set_menu_label.Fn()(nb, GetWidgetIface(child), GetWidgetIface(menuLabel))
}
func (nb *Notebook) SetMenuLabelText(child WidgetIface, menuText string) {
	cMenuText := cc.NewString(menuText)
	defer cMenuText.Free()
	gtk_notebook_set_menu_label_text.Fn()(nb, GetWidgetIface(child), cMenuText)
}
func (nb *Notebook) GetMenuLabelText(child WidgetIface) string {
	return gtk_notebook_get_menu_label_text.Fn()(nb, GetWidgetIface(child)).String()
}
func (nb *Notebook) ReorderChild(child WidgetIface, position int32) {
	gtk_notebook_reorder_child.Fn()(nb, GetWidgetIface(child), position)
}
func (nb *Notebook) GetTabReorderable(child WidgetIface) bool {
	return gtk_notebook_get_tab_reorderable.Fn()(nb, GetWidgetIface(child)) != 0
}
func (nb *Notebook) SetTabReorderable(child WidgetIface, reorderable bool) {
	gtk_notebook_set_tab_reorderable.Fn()(nb, GetWidgetIface(child), cbool(reorderable))
}
func (nb *Notebook) GetTabDetachable(child WidgetIface) bool {
	return gtk_notebook_get_tab_detachable.Fn()(nb, GetWidgetIface(child)) != 0
}
func (nb *Notebook) SetTabDetachable(child WidgetIface, detachable bool) {
	gtk_notebook_set_tab_detachable.Fn()(nb, GetWidgetIface(child), cbool(detachable))
}
func (nb *Notebook) DetachTab(child WidgetIface) {
	gtk_notebook_detach_tab.Fn()(nb, GetWidgetIface(child))
}
func (nb *Notebook) GetActionWidget(packType PackType) *Widget {
	return gtk_notebook_get_action_widget.Fn()(nb, packType)
}
func (nb *Notebook) SetActionWidget(widget WidgetIface, packType PackType) {
	gtk_notebook_set_action_widget.Fn()(nb, GetWidgetIface(widget), packType)
}
func (nb *Notebook) ConnectChangeCurrentPage(sig func(nb *Notebook, page int32)) uint64 {
	return nb.SignalConnect("change-current-page", sig, nil)
}
func (nb *Notebook) ConnectCreateWindow(sig func(nb *Notebook, page *Widget, _ uptr) *Notebook) uint64 {
	return nb.SignalConnect("create-window", sig, nil)
}
func (nb *Notebook) ConnectFocusTab(sig func(nb *Notebook, tab *NotebookTab, _ uptr) bool) uint64 {
	return nb.SignalConnect("focus-tab", sig, nil)
}
func (nb *Notebook) ConnectMoveFocusOut(sig func(nb *Notebook, dir DirectionType)) uint64 {
	return nb.SignalConnect("move-focus-out", sig, nil)
}
func (nb *Notebook) ConnectPageAdded(sig func(nb *Notebook, child *Widget, pageNum uint32)) uint64 {
	return nb.SignalConnect("page-added", sig, nil)
}
func (nb *Notebook) ConnectPageRemoved(sig func(nb *Notebook, child *Widget, pageNum uint32)) uint64 {
	return nb.SignalConnect("page-removed", sig, nil)
}
func (nb *Notebook) ConnectPageReordered(sig func(nb *Notebook, child *Widget, pageNum uint32)) uint64 {
	return nb.SignalConnect("page-reordered", sig, nil)
}
func (nb *Notebook) ConnectReorderTab(
	sig func(nb *Notebook, dir DirectionType, moveToLast bool, _ uptr) bool) uint64 {
	return nb.SignalConnect("reorder-tab", sig, nil)
}
func (nb *Notebook) ConnectSelectPage(
	sig func(nb *Notebook, dir DirectionType, moveFocus bool, _ uptr) bool) uint64 {
	return nb.SignalConnect("select-page", sig, nil)
}
func (nb *Notebook) ConnectSwitchPage(sig func(nb *Notebook, page *Widget, pageNum uint32)) uint64 {
	return nb.SignalConnect("switch-page", sig, nil)
}

type NotebookPage struct{ gobjCore }

func GTypeNotebookPage() gobject.GType { return gtk_notebook_page_get_type.Fn()() }
func (nb *Notebook) GetPage(child WidgetIface) *NotebookPage {
	return gtk_notebook_get_page.Fn()(nb, GetWidgetIface(child))
}
func (page *NotebookPage) GetChild() *Widget   { return gtk_notebook_page_get_child.Fn()(page) }
func (nb *Notebook) GetPages() *gio.GListModel { return gtk_notebook_get_pages.Fn()(nb) }

// #endregion

// #region NumericSorter

type NumericSorter struct{ Sorter }

func GTypeNumericSorter() gobject.GType { return gtk_numeric_sorter_get_type.Fn()() }
func NewNumericSorter(expression ExpressionIface) *NumericSorter {
	return gtk_numeric_sorter_new.Fn()(GetExpressionIface(expression))
}
func (ns *NumericSorter) GetExpression() *Expression {
	return gtk_numeric_sorter_get_expression.Fn()(ns)
}
func (ns *NumericSorter) SetExpression(expression ExpressionIface) {
	gtk_numeric_sorter_set_expression.Fn()(ns, GetExpressionIface(expression))
}
func (ns *NumericSorter) GetSortOrder() SortType { return gtk_numeric_sorter_get_sort_order.Fn()(ns) }
func (ns *NumericSorter) SetSortOrder(sortOrder SortType) {
	gtk_numeric_sorter_set_sort_order.Fn()(ns, sortOrder)
}

// #endregion

// #region Orientable

type OrientableIfaceObj struct {
	Parent gobject.GTypeInterface
}

type Orientable struct{}

func GTypeOrientable() gobject.GType { return gtk_orientable_get_type.Fn()() }
func (o *Orientable) SetOrientation(orientation Orientation) {
	gtk_orientable_set_orientation.Fn()(o, orientation)
}
func (o *Orientable) GetOrientation() Orientation { return gtk_orientable_get_orientation.Fn()(o) }

// #endregion

// #region Overlay

type Overlay struct{ Widget }

func GTypeOverlay() gobject.GType { return gtk_overlay_get_type.Fn()() }
func NewOverlay() *Overlay        { return gtk_overlay_new.Fn()() }
func (o *Overlay) AddOverlay(widget WidgetIface) {
	gtk_overlay_add_overlay.Fn()(o, GetWidgetIface(widget))
}
func (o *Overlay) RemoveOverlay(widget WidgetIface) {
	gtk_overlay_remove_overlay.Fn()(o, GetWidgetIface(widget))
}
func (o *Overlay) SetChild(child WidgetIface) { gtk_overlay_set_child.Fn()(o, GetWidgetIface(child)) }
func (o *Overlay) GetChild() *Widget          { return gtk_overlay_get_child.Fn()(o) }
func (o *Overlay) GetMeasureOverlay(widget WidgetIface) bool {
	return gtk_overlay_get_measure_overlay.Fn()(o, GetWidgetIface(widget)) != 0
}
func (o *Overlay) SetMeasureOverlay(widget WidgetIface, measure bool) {
	gtk_overlay_set_measure_overlay.Fn()(o, GetWidgetIface(widget), cbool(measure))
}
func (o *Overlay) GetClipOverlay(widget WidgetIface) bool {
	return gtk_overlay_get_clip_overlay.Fn()(o, GetWidgetIface(widget)) != 0
}
func (o *Overlay) SetClipOverlay(widget WidgetIface, clipOverlay bool) {
	gtk_overlay_set_clip_overlay.Fn()(o, GetWidgetIface(widget), cbool(clipOverlay))
}
func (o *Overlay) ConnectGetChildPosition(
	sig func(o *Overlay, widget *Widget, allocation *gdk.Rectangle, _ uptr) (setedAllocation bool)) uint64 {
	return o.SignalConnect("get-child-position", sig, nil)
}

// #endregion

// #region OverlayLayout

type OverlayLayout struct{ LayoutManager }

type OverlayLayoutChild struct{ LayoutChild }

func GTypeOverlayLayout() gobject.GType { return gtk_overlay_layout_get_type.Fn()() }
func NewOverlayLayout() *OverlayLayout  { return gtk_overlay_layout_new.Fn()() }
func (c *OverlayLayoutChild) SetMeasure(measure bool) {
	gtk_overlay_layout_child_set_measure.Fn()(c, cbool(measure))
}
func (c *OverlayLayoutChild) GetMeasure() bool {
	return gtk_overlay_layout_child_get_measure.Fn()(c) != 0
}
func (c *OverlayLayoutChild) SetClipOverlay(clipOverlay bool) {
	gtk_overlay_layout_child_set_clip_overlay.Fn()(c, cbool(clipOverlay))
}
func (c *OverlayLayoutChild) GetClipOverlay() bool {
	return gtk_overlay_layout_child_get_clip_overlay.Fn()(c) != 0
}

// #endregion

// #region PadController

type PadActionEntry struct {
	Type       PadActionType
	Index      int32
	Mode       int32
	Label      cc.String
	ActionName cc.String
}

type PadController struct{ EventController }

func GTypePadController() gobject.GType { return gtk_pad_controller_get_type.Fn()() }
func NewPadController(group *gio.GActionGroup, pad *gdk.Device) *PadController {
	return gtk_pad_controller_new.Fn()(group, pad)
}
func (c *PadController) SetActionEntries(entries []PadActionEntry) {
	gtk_pad_controller_set_action_entries.Fn()(c, carry(entries), int32(len(entries)))
}
func (c *PadController) SetAction(actionType PadActionType, index, mode int32, label, actionName string) {
	cLabel, cActionName := cc.NewString(label), cc.NewString(actionName)
	defer cLabel.Free()
	defer cActionName.Free()
	gtk_pad_controller_set_action.Fn()(c, actionType, index, mode, cLabel, cActionName)
}

// #endregion

// #region Paned

type Paned struct {
	Widget
	AccessibleRange
	Orientable
}

func GTypePaned() gobject.GType               { return gtk_paned_get_type.Fn()() }
func NewPaned(orientation Orientation) *Paned { return gtk_paned_new.Fn()(orientation) }
func (p *Paned) SetStartChild(child WidgetIface) {
	gtk_paned_set_start_child.Fn()(p, GetWidgetIface(child))
}
func (p *Paned) GetStartChild() *Widget { return gtk_paned_get_start_child.Fn()(p) }
func (p *Paned) SetResizeStartChild(resize bool) {
	gtk_paned_set_resize_start_child.Fn()(p, cbool(resize))
}
func (p *Paned) GetResizeStartChild() bool { return gtk_paned_get_resize_start_child.Fn()(p) != 0 }
func (p *Paned) SetEndChild(child WidgetIface) {
	gtk_paned_set_end_child.Fn()(p, GetWidgetIface(child))
}
func (p *Paned) GetEndChild() *Widget { return gtk_paned_get_end_child.Fn()(p) }
func (p *Paned) SetShrinkStartChild(resize bool) {
	gtk_paned_set_shrink_start_child.Fn()(p, cbool(resize))
}
func (p *Paned) GetShrinkStartChild() bool     { return gtk_paned_get_shrink_start_child.Fn()(p) != 0 }
func (p *Paned) SetResizeEndChild(resize bool) { gtk_paned_set_resize_end_child.Fn()(p, cbool(resize)) }
func (p *Paned) GetResizeEndChild() bool       { return gtk_paned_get_resize_end_child.Fn()(p) != 0 }
func (p *Paned) SetShrinkEndChild(resize bool) { gtk_paned_set_shrink_end_child.Fn()(p, cbool(resize)) }
func (p *Paned) GetShrinkEndChild() bool       { return gtk_paned_get_shrink_end_child.Fn()(p) != 0 }
func (p *Paned) GetPosition() int32            { return gtk_paned_get_position.Fn()(p) }
func (p *Paned) SetPosition(position int32)    { gtk_paned_set_position.Fn()(p, position) }
func (p *Paned) SetWideHandle(wide bool)       { gtk_paned_set_wide_handle.Fn()(p, cbool(wide)) }
func (p *Paned) GetWideHandle() bool           { return gtk_paned_get_wide_handle.Fn()(p) != 0 }

func (p *Paned) ConnectAcceptPosition(sig func(p *Paned, _ uptr) bool) uint64 {
	return p.SignalConnect("accept-position", sig, nil)
}
func (p *Paned) ConnectCancelPosition(sig func(p *Paned, _ uptr) bool) uint64 {
	return p.SignalConnect("cancel-position", sig, nil)
}
func (p *Paned) ConnectCycleChildFocus(sig func(p *Paned, reversed bool, _ uptr) bool) uint64 {
	return p.SignalConnect("cycle-child-focus", sig, nil)
}
func (p *Paned) ConnectCycleHandleFocus(sig func(p *Paned, reversed bool, _ uptr) bool) uint64 {
	return p.SignalConnect("cycle-handle-focus", sig, nil)
}
func (p *Paned) ConnectMoveHandle(sig func(p *Paned, scrollType *ScrollType, _ uptr) bool) uint64 {
	return p.SignalConnect("move-handle", sig, nil)
}
func (p *Paned) ConnectToggleHandleFocus(sig func(p *Paned, _ uptr) bool) uint64 {
	return p.SignalConnect("toggle-handle-focus", sig, nil)
}

// #endregion

// #region PasswordEntry

type PasswordEntry struct {
	Widget
	Editable
}

func GTypePasswordEntry() gobject.GType { return gtk_password_entry_get_type.Fn()() }
func NewPasswordEntry() *PasswordEntry  { return gtk_password_entry_new.Fn()() }
func (e *PasswordEntry) SetShowPeekIcon(showPeekIcon bool) {
	gtk_password_entry_set_show_peek_icon.Fn()(e, cbool(showPeekIcon))
}
func (e *PasswordEntry) GetShowPeekIcon() bool {
	return gtk_password_entry_get_show_peek_icon.Fn()(e) != 0
}
func (e *PasswordEntry) SetExtraMenu(model *gio.GMenuModel) {
	gtk_password_entry_set_extra_menu.Fn()(e, model)
}
func (e *PasswordEntry) GetExtraMenu() *gio.GMenuModel {
	return gtk_password_entry_get_extra_menu.Fn()(e)
}

func (p *PasswordEntry) ConnectActivate(sig func(p *PasswordEntry)) uint64 {
	return p.SignalConnect("activate", sig, nil)
}

// #endregion

// #region PasswordEntryBuffer

type PasswordEntryBuffer struct{ EntryBuffer }

func GTypePasswordEntryBuffer() gobject.GType      { return gtk_password_entry_buffer_get_type.Fn()() }
func NewPasswordEntryBuffer() *PasswordEntryBuffer { return gtk_password_entry_buffer_new.Fn()() }

// #endregion

// #region Picture

type Picture struct{ Widget }

func GTypePicture() gobject.GType { return gtk_picture_get_type.Fn()() }
func NewPicture() *Picture        { return gtk_picture_new.Fn()() }
func NewPictureForPaintable(paintable *gdk.Paintable) *Picture {
	return gtk_picture_new_for_paintable.Fn()(paintable)
}
func NewPictureForFile(file *gio.GFile) *Picture { return gtk_picture_new_for_file.Fn()(file) }
func NewPictureForFilename(filename string) *Picture {
	cFilename := cc.NewString(filename)
	defer cFilename.Free()
	return gtk_picture_new_for_filename.Fn()(cFilename)
}
func NewPictureForResource(resourcePath string) *Picture {
	cResource := cc.NewString(resourcePath)
	defer cResource.Free()
	return gtk_picture_new_for_resource.Fn()(cResource)
}
func (p *Picture) SetPaintable(paintable *gdk.Paintable) {
	gtk_picture_set_paintable.Fn()(p, paintable)
}
func (p *Picture) GetPaintable() *gdk.Paintable { return gtk_picture_get_paintable.Fn()(p) }
func (p *Picture) SetFile(file *gio.GFile)      { gtk_picture_set_file.Fn()(p, file) }
func (p *Picture) GetFile() *gio.GFile          { return gtk_picture_get_file.Fn()(p) }
func (p *Picture) SetFilename(filename string) {
	cFilename := cc.NewString(filename)
	defer cFilename.Free()
	gtk_picture_set_filename.Fn()(p, cFilename)
}
func (p *Picture) SetResource(resourcePath string) {
	cResource := cc.NewString(resourcePath)
	defer cResource.Free()
	gtk_picture_set_resource.Fn()(p, cResource)
}
func (p *Picture) SetCanShrink(canShrink bool) { gtk_picture_set_can_shrink.Fn()(p, cbool(canShrink)) }
func (p *Picture) GetCanShrink() bool          { return gtk_picture_get_can_shrink.Fn()(p) != 0 }
func (p *Picture) SetContentFit(contentFit ContentFit) {
	gtk_picture_set_content_fit.Fn()(p, contentFit)
}
func (p *Picture) GetContentFit() ContentFit { return gtk_picture_get_content_fit.Fn()(p) }
func (p *Picture) SetAlternativeText(alternativeText string) {
	cAlt := cc.NewString(alternativeText)
	defer cAlt.Free()
	gtk_picture_set_alternative_text.Fn()(p, cAlt)
}
func (p *Picture) GetAlternativeText() string {
	return gtk_picture_get_alternative_text.Fn()(p).String()
}

// #endregion

// #region Popover

type PopoverObj struct {
	Parent Widget
}

type PopoverClass struct {
	Parent WidgetClass

	Closed          cc.Func // void (*)(GtkPopover*)
	ActivateDefault cc.Func // void (*)(GtkPopover*)

	_ [8]uptr //reserved
}

type PopoverIface interface {
	GetPopoverIface() *Popover
}

func GetPopoverIface(iface PopoverIface) *Popover {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetPopoverIface()
}
func (p *Popover) GetPopoverIface() *Popover { return p }

type Popover struct {
	Widget
	Native
	ShortcutManager
}

func GTypePopover() gobject.GType { return gtk_popover_get_type.Fn()() }
func NewPopover() *Popover        { return gtk_popover_new.Fn()() }
func (p *Popover) SetChild(child WidgetIface) {
	gtk_popover_set_child.Fn()(p, GetWidgetIface(child))
}
func (p *Popover) GetChild() *Widget { return gtk_popover_get_child.Fn()(p) }
func (p *Popover) SetPointingTo(rect *gdk.Rectangle) {
	gtk_popover_set_pointing_to.Fn()(p, rect)
}
func (p *Popover) GetPointingTo(rect *gdk.Rectangle) bool {
	return gtk_popover_get_pointing_to.Fn()(p, rect) != 0
}
func (p *Popover) SetPosition(position PositionType) {
	gtk_popover_set_position.Fn()(p, position)
}
func (p *Popover) GetPosition() PositionType {
	return gtk_popover_get_position.Fn()(p)
}
func (p *Popover) SetAutohide(autohide bool) {
	gtk_popover_set_autohide.Fn()(p, cbool(autohide))
}
func (p *Popover) GetAutohide() bool {
	return gtk_popover_get_autohide.Fn()(p) != 0
}
func (p *Popover) SetHasArrow(hasArrow bool) {
	gtk_popover_set_has_arrow.Fn()(p, cbool(hasArrow))
}
func (p *Popover) GetHasArrow() bool {
	return gtk_popover_get_has_arrow.Fn()(p) != 0
}
func (p *Popover) SetMnemonicsVisible(mnemonicsVisible bool) {
	gtk_popover_set_mnemonics_visible.Fn()(p, cbool(mnemonicsVisible))
}
func (p *Popover) GetMnemonicsVisible() bool {
	return gtk_popover_get_mnemonics_visible.Fn()(p) != 0
}
func (p *Popover) Popup()   { gtk_popover_popup.Fn()(p) }
func (p *Popover) Popdown() { gtk_popover_popdown.Fn()(p) }
func (p *Popover) SetOffset(xOffset, yOffset int32) {
	gtk_popover_set_offset.Fn()(p, xOffset, yOffset)
}
func (p *Popover) GetOffset() (xOffset, yOffset int32) {
	gtk_popover_get_offset.Fn()(p, &xOffset, &yOffset)
	return
}
func (p *Popover) SetCascadePopdown(cascadePopdown bool) {
	gtk_popover_set_cascade_popdown.Fn()(p, cbool(cascadePopdown))
}
func (p *Popover) GetCascadePopdown() bool {
	return gtk_popover_get_cascade_popdown.Fn()(p) != 0
}
func (p *Popover) SetDefaultWidget(widget WidgetIface) {
	gtk_popover_set_default_widget.Fn()(p, GetWidgetIface(widget))
}
func (p *Popover) Present() { gtk_popover_present.Fn()(p) }
func (p *Popover) ConnectActivateDefault(f func(p *Popover)) uint64 {
	return p.SignalConnect("activate-default", f, nil)
}
func (p *Popover) ConnectClosed(f func(p *Popover)) uint64 {
	return p.SignalConnect("closed", f, nil)
}

// #endregion

// #region PopoverMenu

type PopoverMenu struct{ Popover }

func GTypePopoverMenu() gobject.GType { return gtk_popover_menu_get_type.Fn()() }
func NewPopoverMenuFromModel(model *gio.GMenuModel) *PopoverMenu {
	return gtk_popover_menu_new_from_model.Fn()(model)
}
func NewPopoverMenuFromModelFull(model *gio.GMenuModel, flags PopoverMenuFlags) *PopoverMenu {
	return gtk_popover_menu_new_from_model_full.Fn()(model, flags)
}
func (p *PopoverMenu) SetMenuModel(model *gio.GMenuModel) {
	gtk_popover_menu_set_menu_model.Fn()(p, model)
}
func (p *PopoverMenu) GetMenuModel() *gio.GMenuModel   { return gtk_popover_menu_get_menu_model.Fn()(p) }
func (p *PopoverMenu) SetFlags(flags PopoverMenuFlags) { gtk_popover_menu_set_flags.Fn()(p, flags) }
func (p *PopoverMenu) GetFlags() PopoverMenuFlags      { return gtk_popover_menu_get_flags.Fn()(p) }
func (p *PopoverMenu) AddChild(child WidgetIface, id string) bool {
	cId := cc.NewString(id)
	defer cId.Free()
	return gtk_popover_menu_add_child.Fn()(p, GetWidgetIface(child), cId) != 0
}
func (p *PopoverMenu) RemoveChild(child WidgetIface) bool {
	return gtk_popover_menu_remove_child.Fn()(p, GetWidgetIface(child)) != 0
}

// #endregion

// #region PopoverMenuBar

type PopoverMenuBar struct{ Widget }

func GTypePopoverMenuBar() gobject.GType { return gtk_popover_menu_bar_get_type.Fn()() }
func NewPopoverMenuBarFromModel(model *gio.GMenuModel) *PopoverMenuBar {
	return gtk_popover_menu_bar_new_from_model.Fn()(model)
}
func (b *PopoverMenuBar) SetMenuModel(model *gio.GMenuModel) {
	gtk_popover_menu_bar_set_menu_model.Fn()(b, model)
}
func (b *PopoverMenuBar) GetMenuModel() *gio.GMenuModel {
	return gtk_popover_menu_bar_get_menu_model.Fn()(b)
}
func (b *PopoverMenuBar) AddChild(child WidgetIface, id string) bool {
	cId := cc.NewString(id)
	defer cId.Free()
	return gtk_popover_menu_bar_add_child.Fn()(b, GetWidgetIface(child), cId) != 0
}
func (b *PopoverMenuBar) RemoveChild(child WidgetIface) bool {
	return gtk_popover_menu_bar_remove_child.Fn()(b, GetWidgetIface(child)) != 0
}

// #endregion

// #region PrintDialog

type PrintSetup struct{}
type PrintSettings struct{ gobjCore }
type PrintDialog struct{ gobjCore }
type PageSetup struct{ gobjCore }

func GTypePrintSetup() gobject.GType   { return gtk_print_setup_get_type.Fn()() }
func (s *PrintSetup) Ref() *PrintSetup { return gtk_print_setup_ref.Fn()(s) }
func (s *PrintSetup) Unref()           { gtk_print_setup_unref.Fn()(s) }
func (s *PrintSetup) GetPrintSettings() *PrintSettings {
	return gtk_print_setup_get_print_settings.Fn()(s)
}
func (s *PrintSetup) GetPageSetup() *PageSetup { return gtk_print_setup_get_page_setup.Fn()(s) }

func GTypePrintDialog() gobject.GType   { return gtk_print_dialog_get_type.Fn()() }
func NewPrintDialog() *PrintDialog      { return gtk_print_dialog_new.Fn()() }
func (d *PrintDialog) GetTitle() string { return gtk_print_dialog_get_title.Fn()(d).String() }
func (d *PrintDialog) SetTitle(title string) {
	cTitle := cc.NewString(title)
	defer cTitle.Free()
	gtk_print_dialog_set_title.Fn()(d, cTitle)
}
func (d *PrintDialog) GetAcceptLabel() string {
	return gtk_print_dialog_get_accept_label.Fn()(d).String()
}
func (d *PrintDialog) SetAcceptLabel(acceptLabel string) {
	cLabel := cc.NewString(acceptLabel)
	defer cLabel.Free()
	gtk_print_dialog_set_accept_label.Fn()(d, cLabel)
}
func (d *PrintDialog) GetModal() bool           { return gtk_print_dialog_get_modal.Fn()(d) != 0 }
func (d *PrintDialog) SetModal(modal bool)      { gtk_print_dialog_set_modal.Fn()(d, cbool(modal)) }
func (d *PrintDialog) GetPageSetup() *PageSetup { return gtk_print_dialog_get_page_setup.Fn()(d) }
func (d *PrintDialog) SetPageSetup(pageSetup *PageSetup) {
	gtk_print_dialog_set_page_setup.Fn()(d, pageSetup)
}
func (d *PrintDialog) GetPrintSettings() *PrintSettings {
	return gtk_print_dialog_get_print_settings.Fn()(d)
}
func (d *PrintDialog) SetPrintSettings(printSettings *PrintSettings) {
	gtk_print_dialog_set_print_settings.Fn()(d, printSettings)
}
func (d *PrintDialog) Setup(parent WindowIface, cancellable *gio.GCancellable,
	callback func(d *PrintDialog, res *gio.GAsyncResult)) {
	cb := cc.Cbk(callback)
	d.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	gtk_print_dialog_setup.Fn()(d, GetWindowIface(parent), cancellable, cb, nil)
}
func (d *PrintDialog) SetupFinish(result *gio.GAsyncResult) (*PrintSetup, error) {
	var gerr *glib.GError
	return gtk_print_dialog_setup_finish.Fn()(d, result, &gerr), gerr.TakeError()
}
func (d *PrintDialog) Print(parent WindowIface, setup *PrintSetup, cancellable *gio.GCancellable,
	callback func(d *PrintDialog, res *gio.GAsyncResult)) {
	cb := cc.Cbk(callback)
	d.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	gtk_print_dialog_print.Fn()(d, GetWindowIface(parent), setup, cancellable, cb, nil)
}
func (d *PrintDialog) PrintFinish(result *gio.GAsyncResult) (*gio.GOutputStream, error) {
	var gerr *glib.GError
	return gtk_print_dialog_print_finish.Fn()(d, result, &gerr), gerr.TakeError()
}
func (d *PrintDialog) PrintFile(parent WindowIface, setup *PrintSetup, file *gio.GFile, cancellable *gio.GCancellable,
	callback func(d *PrintDialog, res *gio.GAsyncResult)) {
	cb := cc.Cbk(callback)
	d.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	gtk_print_dialog_print_file.Fn()(d, GetWindowIface(parent), setup, file, cancellable, cb, nil)
}
func (d *PrintDialog) PrintFileFinish(result *gio.GAsyncResult) error {
	var gerr *glib.GError
	gtk_print_dialog_print_file_finish.Fn()(d, result, &gerr)
	return gerr.TakeError()
}

// #endregion

// #region ProgressBar

type ProgressBar struct {
	Widget
	AccessibleRange
	Orientable
}

func GTypeProgressBar() gobject.GType { return gtk_progress_bar_get_type.Fn()() }
func NewProgressBar() *ProgressBar    { return gtk_progress_bar_new.Fn()() }
func (p *ProgressBar) Pulse()         { gtk_progress_bar_pulse.Fn()(p) }
func (p *ProgressBar) SetText(text string) {
	cText := cc.NewString(text)
	defer cText.Free()
	gtk_progress_bar_set_text.Fn()(p, cText)
}
func (p *ProgressBar) SetFraction(fraction float64) { gtk_progress_bar_set_fraction.Fn()(p, fraction) }
func (p *ProgressBar) SetPulseStep(fraction float64) {
	gtk_progress_bar_set_pulse_step.Fn()(p, fraction)
}
func (p *ProgressBar) SetInverted(inverted bool) {
	gtk_progress_bar_set_inverted.Fn()(p, cbool(inverted))
}
func (p *ProgressBar) GetText() string       { return gtk_progress_bar_get_text.Fn()(p).String() }
func (p *ProgressBar) GetFraction() float64  { return gtk_progress_bar_get_fraction.Fn()(p) }
func (p *ProgressBar) GetPulseStep() float64 { return gtk_progress_bar_get_pulse_step.Fn()(p) }
func (p *ProgressBar) GetInverted() bool     { return gtk_progress_bar_get_inverted.Fn()(p) != 0 }
func (p *ProgressBar) SetEllipsize(mode pango.EllipsizeMode) {
	gtk_progress_bar_set_ellipsize.Fn()(p, mode)
}
func (p *ProgressBar) GetEllipsize() pango.EllipsizeMode {
	return gtk_progress_bar_get_ellipsize.Fn()(p)
}
func (p *ProgressBar) SetShowText(showText bool) {
	gtk_progress_bar_set_show_text.Fn()(p, cbool(showText))
}
func (p *ProgressBar) GetShowText() bool { return gtk_progress_bar_get_show_text.Fn()(p) != 0 }

// #endregion

// #region Range

type RangeObj struct {
	Parent WidgetObj
}

type RangeClass struct {
	Parent WidgetClass

	ValueChanged   cc.Func //void (* value_changed) (GtkRange *range);
	AdjustBounds   cc.Func //void (* adjust_bounds) (GtkRange *range,double new_value);
	MoveSlider     cc.Func //void (* move_slider) (GtkRange *range,GtkScrollType scroll);
	GetRangeBorder cc.Func //void (* get_range_border) (GtkRange *range, GtkBorder *border_);
	ChangeValue    cc.Func //gboolean (* change_value) (GtkRange *range,GtkScrollType scroll,double new_value);

	_ [8]uptr // padding
}

type Range struct {
	Widget
	Accessible
	AccessibleRange
	Orientable
}

func GTypeRange() gobject.GType                { return gtk_range_get_type.Fn()() }
func (r *Range) SetAdjustment(adj *Adjustment) { gtk_range_set_adjustment.Fn()(r, adj) }
func (r *Range) GetAdjustment() *Adjustment    { return gtk_range_get_adjustment.Fn()(r) }
func (r *Range) SetInverted(setting bool)      { gtk_range_set_inverted.Fn()(r, cbool(setting)) }
func (r *Range) GetInverted() bool             { return gtk_range_get_inverted.Fn()(r) != 0 }
func (r *Range) SetFlippable(flippable bool)   { gtk_range_set_flippable.Fn()(r, cbool(flippable)) }
func (r *Range) GetFlippable() bool            { return gtk_range_get_flippable.Fn()(r) != 0 }
func (r *Range) SetSliderSizeFixed(sizeFixed bool) {
	gtk_range_set_slider_size_fixed.Fn()(r, cbool(sizeFixed))
}
func (r *Range) GetSliderSizeFixed() bool { return gtk_range_get_slider_size_fixed.Fn()(r) != 0 }
func (r *Range) GetRangeRect() (rangeRect gdk.Rectangle) {
	gtk_range_get_range_rect.Fn()(r, &rangeRect)
	return
}
func (r *Range) GetSliderRange() (sliderStart, sliderEnd int32) {
	gtk_range_get_slider_range.Fn()(r, &sliderStart, &sliderEnd)
	return
}
func (r *Range) SetIncrements(step, page float64) { gtk_range_set_increments.Fn()(r, step, page) }
func (r *Range) SetRange(min, max float64)        { gtk_range_set_range.Fn()(r, min, max) }
func (r *Range) SetValue(value float64)           { gtk_range_set_value.Fn()(r, value) }
func (r *Range) GetValue() float64                { return gtk_range_get_value.Fn()(r) }
func (r *Range) SetShowFillLevel(showFillLevel bool) {
	gtk_range_set_show_fill_level.Fn()(r, cbool(showFillLevel))
}
func (r *Range) GetShowFillLevel() bool { return gtk_range_get_show_fill_level.Fn()(r) != 0 }
func (r *Range) SetRestrictToFillLevel(restrict bool) {
	gtk_range_set_restrict_to_fill_level.Fn()(r, cbool(restrict))
}
func (r *Range) GetRestrictToFillLevel() bool {
	return gtk_range_get_restrict_to_fill_level.Fn()(r) != 0
}
func (r *Range) SetFillLevel(fillLevel float64)   { gtk_range_set_fill_level.Fn()(r, fillLevel) }
func (r *Range) GetFillLevel() float64            { return gtk_range_get_fill_level.Fn()(r) }
func (r *Range) SetRoundDigits(roundDigits int32) { gtk_range_set_round_digits.Fn()(r, roundDigits) }
func (r *Range) GetRoundDigits() int32            { return gtk_range_get_round_digits.Fn()(r) }
func (r *Range) ConnectAdjustBounds(sig func(r *Range, value float64)) uint64 {
	return r.SignalConnect("adjust-bounds", sig, nil)
}
func (r *Range) ConnectChangeValue(
	sig func(r *Range, scroll ScrollType, value float64, _ uptr) bool) uint64 {
	return r.SignalConnect("change-value", sig, nil)
}
func (r *Range) ConnectMoveSlider(sig func(r *Range, step ScrollType)) uint64 {
	return r.SignalConnect("move-slider", sig, nil)
}
func (r *Range) ConnectValueChanged(sig func(r *Range)) uint64 {
	return r.SignalConnect("value-changed", sig, nil)
}

// #endregion

// #region RecentManager

type RecentData struct {
	DisplayName cc.String
	Description cc.String
	MimeType    cc.String
	AppName     cc.String
	AppExec     cc.String
	Groups      cc.Strings
	IsPrivate   int32
}

type RecentManagerObj struct {
	Parent gobject.GObjectObj
	Priv   uptr
}

type RecentManagerClass struct {
	Parent  gobject.GObjectClass
	Changed cc.Func    // void (*changed) (GtkRecentManager *manager);
	_       [4]cc.Func //recent
}

type RecentManager struct{ gobjCore }

func RecentManagerErrorQuark() glib.GQuark    { return gtk_recent_manager_error_quark.Fn()() }
func GTypeRecentManager() gobject.GType       { return gtk_recent_manager_get_type.Fn()() }
func NewRecentManager() *RecentManager        { return gtk_recent_manager_new.Fn()() }
func GetDefaultRecentManager() *RecentManager { return gtk_recent_manager_get_default.Fn()() }
func (m *RecentManager) AddItem(uri string) bool {
	cUri := cc.NewString(uri)
	defer cUri.Free()
	return gtk_recent_manager_add_item.Fn()(m, cUri) != 0
}
func (m *RecentManager) AddFull(uri string, recentData *RecentData) bool {
	cUri := cc.NewString(uri)
	defer cUri.Free()
	return gtk_recent_manager_add_full.Fn()(m, cUri, recentData) != 0
}
func (m *RecentManager) RemoveItem(uri string) error {
	cUri := cc.NewString(uri)
	defer cUri.Free()
	var gerr *glib.GError
	gtk_recent_manager_remove_item.Fn()(m, cUri, &gerr)
	return gerr.TakeError()
}
func (m *RecentManager) LookupItem(uri string) (*RecentInfo, error) {
	cUri := cc.NewString(uri)
	defer cUri.Free()
	var gerr *glib.GError
	return gtk_recent_manager_lookup_item.Fn()(m, cUri, &gerr), gerr.TakeError()
}
func (m *RecentManager) HasItem(uri string) bool {
	cUri := cc.NewString(uri)
	defer cUri.Free()
	return gtk_recent_manager_has_item.Fn()(m, cUri) != 0
}
func (m *RecentManager) MoveItem(uri, newUri string) error {
	cUri := cc.NewString(uri)
	cNewUri := cc.NewString(newUri)
	defer cUri.Free()
	defer cNewUri.Free()
	var gerr *glib.GError
	gtk_recent_manager_move_item.Fn()(m, cUri, cNewUri, &gerr)
	return gerr.TakeError()
}
func (m *RecentManager) GetItems() *glib.GList[RecentInfo] {
	return gtk_recent_manager_get_items.Fn()(m)
}
func (m *RecentManager) PurgeItems() (int32, error) {
	var gerr *glib.GError
	return gtk_recent_manager_purge_items.Fn()(m, &gerr), gerr.TakeError()
}

type RecentInfo struct{}

func GTypeRecentInfo() gobject.GType   { return gtk_recent_info_get_type.Fn()() }
func (i *RecentInfo) Ref() *RecentInfo { return gtk_recent_info_ref.Fn()(i) }
func (i *RecentInfo) Unref()           { gtk_recent_info_unref.Fn()(i) }
func (i *RecentInfo) GetUri() string   { return gtk_recent_info_get_uri.Fn()(i).String() }
func (i *RecentInfo) GetDisplayName() string {
	return gtk_recent_info_get_display_name.Fn()(i).String()
}
func (i *RecentInfo) GetDescription() string       { return gtk_recent_info_get_description.Fn()(i).String() }
func (i *RecentInfo) GetMimeType() string          { return gtk_recent_info_get_mime_type.Fn()(i).String() }
func (i *RecentInfo) GetAdded() *glib.GDateTime    { return gtk_recent_info_get_added.Fn()(i) }
func (i *RecentInfo) GetModified() *glib.GDateTime { return gtk_recent_info_get_modified.Fn()(i) }
func (i *RecentInfo) GetVisited() *glib.GDateTime  { return gtk_recent_info_get_visited.Fn()(i) }
func (i *RecentInfo) GetPrivateHint() bool         { return gtk_recent_info_get_private_hint.Fn()(i) != 0 }
func (i *RecentInfo) GetApplicationInfo(appName string) (appExec string, count uint32, stamp *glib.GDateTime, ok bool) {
	cAppName := cc.NewString(appName)
	defer cAppName.Free()
	var cAppExec cc.String
	var cCount uint32
	var cStamp *glib.GDateTime
	ok = gtk_recent_info_get_application_info.Fn()(i, cAppName, &cAppExec, &cCount, &cStamp) != 0
	appExec = cAppExec.String()
	stamp = cStamp
	count = cCount
	return
}
func (i *RecentInfo) CreateAppInfo(appName string) (*gio.GAppInfo, error) {
	cAppName := cc.NewString(appName)
	defer cAppName.Free()
	var gerr *glib.GError
	return gtk_recent_info_create_app_info.Fn()(i, cAppName, &gerr), gerr.TakeError()
}
func (i *RecentInfo) GetApplications() []string {
	var length uint64
	ptr := gtk_recent_info_get_applications.Fn()(i, &length)
	return ptr.Ref()
}
func (i *RecentInfo) LastApplication() string {
	return gtk_recent_info_last_application.Fn()(i).TakeString()
}
func (i *RecentInfo) HasApplication(appName string) bool {
	cAppName := cc.NewString(appName)
	defer cAppName.Free()
	return gtk_recent_info_has_application.Fn()(i, cAppName) != 0
}
func (i *RecentInfo) GetGroups() []string {
	var length uint64
	ptr := gtk_recent_info_get_groups.Fn()(i, &length)
	return ptr.Strings()
}
func (i *RecentInfo) HasGroup(groupName string) bool {
	cGroupName := cc.NewString(groupName)
	defer cGroupName.Free()
	return gtk_recent_info_has_group.Fn()(i, cGroupName) != 0
}
func (i *RecentInfo) GetGicon() *gio.GIcon { return gtk_recent_info_get_gicon.Fn()(i) }
func (i *RecentInfo) GetShortName() string {
	return gtk_recent_info_get_short_name.Fn()(i).TakeString()
}
func (i *RecentInfo) GetUriDisplay() string {
	return gtk_recent_info_get_uri_display.Fn()(i).TakeString()
}
func (i *RecentInfo) GetAge() int32                { return gtk_recent_info_get_age.Fn()(i) }
func (i *RecentInfo) IsLocal() bool                { return gtk_recent_info_is_local.Fn()(i) != 0 }
func (i *RecentInfo) Exists() bool                 { return gtk_recent_info_exists.Fn()(i) != 0 }
func (i *RecentInfo) Match(other *RecentInfo) bool { return gtk_recent_info_match.Fn()(i, other) != 0 }
func (r *RecentManager) ConnectChanged(sig func(r *RecentManager)) uint64 {
	return r.SignalConnect("changed", sig, nil)
}

// #endregion

// #region Revealer

type Revealer struct{ Widget }

func GTypeRevealer() gobject.GType       { return gtk_revealer_get_type.Fn()() }
func NewRevealer() *Revealer             { return gtk_revealer_new.Fn()() }
func (r *Revealer) GetRevealChild() bool { return gtk_revealer_get_reveal_child.Fn()(r) != 0 }
func (r *Revealer) SetRevealChild(revealChild bool) {
	gtk_revealer_set_reveal_child.Fn()(r, cbool(revealChild))
}
func (r *Revealer) GetChildRevealed() bool { return gtk_revealer_get_child_revealed.Fn()(r) != 0 }
func (r *Revealer) GetTransitionDuration() uint32 {
	return gtk_revealer_get_transition_duration.Fn()(r)
}
func (r *Revealer) SetTransitionDuration(duration uint32) {
	gtk_revealer_set_transition_duration.Fn()(r, duration)
}
func (r *Revealer) SetTransitionType(transition RevealerTransitionType) {
	gtk_revealer_set_transition_type.Fn()(r, transition)
}
func (r *Revealer) GetTransitionType() RevealerTransitionType {
	return gtk_revealer_get_transition_type.Fn()(r)
}
func (r *Revealer) SetChild(child WidgetIface) { gtk_revealer_set_child.Fn()(r, GetWidgetIface(child)) }
func (r *Revealer) GetChild() *Widget          { return gtk_revealer_get_child.Fn()(r) }

// #endregion

// #region Root

type Root struct{ Native }

func (r *Root) GetDisplay() *gdk.Display { return gtk_root_get_display.Fn()(r) }
func (r *Root) SetFocus(focus *Widget)   { gtk_root_set_focus.Fn()(r, focus) }
func (r *Root) GetFocus() *Widget        { return gtk_root_get_focus.Fn()(r) }
func (r *Root) AsWindow() *Window        { return (*Window)(uptr(r)) }

// #endregion

// #region Scale

type ScaleObj struct{ Parent RangeObj }

type ScaleClass struct {
	Parent           RangeClass
	GetLayoutOffsets cc.Func //void (* get_layout_offsets) (GtkScale *scale,int *x,int *y);

	_ [8]uptr //padding
}

type Scale struct{ Range }

func GTypeScale() gobject.GType { return gtk_scale_get_type.Fn()() }
func NewScale(orientation Orientation, adjustment *Adjustment) *Scale {
	return gtk_scale_new.Fn()(orientation, adjustment)
}
func NewScaleWithRange(orientation Orientation, min, max, step float64) *Scale {
	return gtk_scale_new_with_range.Fn()(orientation, min, max, step)
}
func (s *Scale) SetDigits(digits int32)         { gtk_scale_set_digits.Fn()(s, digits) }
func (s *Scale) GetDigits() int32               { return gtk_scale_get_digits.Fn()(s) }
func (s *Scale) SetDrawValue(drawValue bool)    { gtk_scale_set_draw_value.Fn()(s, cbool(drawValue)) }
func (s *Scale) GetDrawValue() bool             { return gtk_scale_get_draw_value.Fn()(s) != 0 }
func (s *Scale) SetHasOrigin(hasOrigin bool)    { gtk_scale_set_has_origin.Fn()(s, cbool(hasOrigin)) }
func (s *Scale) GetHasOrigin() bool             { return gtk_scale_get_has_origin.Fn()(s) != 0 }
func (s *Scale) SetValuePos(pos PositionType)   { gtk_scale_set_value_pos.Fn()(s, pos) }
func (s *Scale) GetValuePos() PositionType      { return gtk_scale_get_value_pos.Fn()(s) }
func (s *Scale) GetLayout() *pango.Layout       { return gtk_scale_get_layout.Fn()(s) }
func (s *Scale) GetLayoutOffsets() (x, y int32) { gtk_scale_get_layout_offsets.Fn()(s, &x, &y); return }
func (s *Scale) AddMark(value float64, position PositionType, markup string) {
	cMarkup := cc.NewString(markup)
	defer cMarkup.Free()
	gtk_scale_add_mark.Fn()(s, value, position, cMarkup)
}
func (s *Scale) ClearMarks() { gtk_scale_clear_marks.Fn()(s) }
func (s *Scale) SetFormatValueFunc(fn func(scale *Scale, value float64) string) {
	var cb, des uptr
	if fn != nil {
		cb = cc.CbkRaw[func(scale *Scale, value float64, _ uptr) cc.String](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 2)
			r := fn(*(**Scale)(is[0]), *(*float64)(is[2]))
			*(*cc.String)(out) = cc.NewString(r)
		})
		des = cc.Cbk(func(_ uptr) { cc.CbkClose(cb); cc.CbkCloseLate(des) })
	}
	gtk_scale_set_format_value_func.Fn()(s, cb, nil, des)
}

// #endregion

// #region ScaleButton

type ScaleButtonObj struct {
	Parent WidgetObj
}

type ScaleButtonClass struct {
	Parent       WidgetClass
	ValueChanged cc.Func //void	(* value_changed) (GtkScaleButton *button,double value);

	_ [8]uptr // padding
}

type ScaleButton struct {
	Widget
	AccessibleRange
	Orientable
}

func GTypeScaleButton() gobject.GType { return gtk_scale_button_get_type.Fn()() }
func NewScaleButton(min, max, step float64, icons []string) *ScaleButton {
	cIcons := cc.NewStrings(icons)
	defer cIcons.Free()
	return gtk_scale_button_new.Fn()(min, max, step, cIcons)
}
func (b *ScaleButton) SetIcons(icons []string) {
	cIcons := cc.NewStrings(icons)
	defer cIcons.Free()
	gtk_scale_button_set_icons.Fn()(b, cIcons)
}
func (b *ScaleButton) GetValue() float64          { return gtk_scale_button_get_value.Fn()(b) }
func (b *ScaleButton) SetValue(value float64)     { gtk_scale_button_set_value.Fn()(b, value) }
func (b *ScaleButton) GetAdjustment() *Adjustment { return gtk_scale_button_get_adjustment.Fn()(b) }
func (b *ScaleButton) SetAdjustment(adjustment *Adjustment) {
	gtk_scale_button_set_adjustment.Fn()(b, adjustment)
}
func (b *ScaleButton) GetPlusButton() *Widget  { return gtk_scale_button_get_plus_button.Fn()(b) }
func (b *ScaleButton) GetMinusButton() *Widget { return gtk_scale_button_get_minus_button.Fn()(b) }
func (b *ScaleButton) GetPopup() *Widget       { return gtk_scale_button_get_popup.Fn()(b) }
func (b *ScaleButton) GetActive() bool         { return gtk_scale_button_get_active.Fn()(b) != 0 }
func (b *ScaleButton) GetHasFrame() bool       { return gtk_scale_button_get_has_frame.Fn()(b) != 0 }
func (b *ScaleButton) SetHasFrame(hasFrame bool) {
	gtk_scale_button_set_has_frame.Fn()(b, cbool(hasFrame))
}

func (r *ScaleButton) ConnectPopdown(sig func(r *ScaleButton)) uint64 {
	return r.SignalConnect("popdown", sig, nil)
}
func (r *ScaleButton) ConnectPopup(sig func(r *ScaleButton)) uint64 {
	return r.SignalConnect("popup", sig, nil)
}
func (r *ScaleButton) ConnectValueChanged(sig func(r *ScaleButton, val float64)) uint64 {
	return r.SignalConnect("value-changed", sig, nil)
}

// #endregion

// #region Scrollable

type ScrollableIfaceObj struct {
	GIface    gobject.GTypeInterface
	GetBorder cc.Func //gboolean (* get_border) (GtkScrollable *scrollable,GtkBorder *border);
}

type ScrollableIface interface {
	GetScrollableIface() *Scrollable
}

func GetScrollableIface(iface ScrollableIface) *Scrollable {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetScrollableIface()
}

func (s *Scrollable) GetScrollableIface() *Scrollable { return s }

type Scrollable struct{}

func GTypeScrollable() gobject.GType              { return gtk_scrollable_get_type.Fn()() }
func (s *Scrollable) GetHAdjustment() *Adjustment { return gtk_scrollable_get_hadjustment.Fn()(s) }
func (s *Scrollable) SetHAdjustment(hadjustment *Adjustment) {
	gtk_scrollable_set_hadjustment.Fn()(s, hadjustment)
}
func (s *Scrollable) GetVAdjustment() *Adjustment { return gtk_scrollable_get_vadjustment.Fn()(s) }
func (s *Scrollable) SetVAdjustment(vadjustment *Adjustment) {
	gtk_scrollable_set_vadjustment.Fn()(s, vadjustment)
}
func (s *Scrollable) GetHScrollPolicy() ScrollablePolicy {
	return gtk_scrollable_get_hscroll_policy.Fn()(s)
}
func (s *Scrollable) SetHScrollPolicy(policy ScrollablePolicy) {
	gtk_scrollable_set_hscroll_policy.Fn()(s, policy)
}
func (s *Scrollable) GetVScrollPolicy() ScrollablePolicy {
	return gtk_scrollable_get_vscroll_policy.Fn()(s)
}
func (s *Scrollable) SetVScrollPolicy(policy ScrollablePolicy) {
	gtk_scrollable_set_vscroll_policy.Fn()(s, policy)
}
func (s *Scrollable) GetBorder() (border Border, ok bool) {
	ok = gtk_scrollable_get_border.Fn()(s, &border) != 0
	return
}

// #endregion

// #region Scrollbar

type Scrollbar struct {
	Widget
	AccessibleRange
	Orientable
}

func GTypeScrollbar() gobject.GType { return gtk_scrollbar_get_type.Fn()() }
func NewScrollbar(orientation Orientation, adjustment *Adjustment) *Scrollbar {
	return gtk_scrollbar_new.Fn()(orientation, adjustment)
}
func (s *Scrollbar) SetAdjustment(adjustment *Adjustment) {
	gtk_scrollbar_set_adjustment.Fn()(s, adjustment)
}
func (s *Scrollbar) GetAdjustment() *Adjustment {
	return gtk_scrollbar_get_adjustment.Fn()(s)
}

// #endregion

// #region ScrolledWindow

type ScrolledWindow struct{ Widget }

func GTypeScrolledWindow() gobject.GType { return gtk_scrolled_window_get_type.Fn()() }
func NewScrolledWindow() *ScrolledWindow { return gtk_scrolled_window_new.Fn()() }
func (w *ScrolledWindow) SetHAdjustment(hadjustment *Adjustment) {
	gtk_scrolled_window_set_hadjustment.Fn()(w, hadjustment)
}
func (w *ScrolledWindow) SetVAdjustment(vadjustment *Adjustment) {
	gtk_scrolled_window_set_vadjustment.Fn()(w, vadjustment)
}
func (w *ScrolledWindow) GetHAdjustment() *Adjustment {
	return gtk_scrolled_window_get_hadjustment.Fn()(w)
}
func (w *ScrolledWindow) GetVAdjustment() *Adjustment {
	return gtk_scrolled_window_get_vadjustment.Fn()(w)
}
func (w *ScrolledWindow) GetHScrollbar() *Scrollbar {
	return gtk_scrolled_window_get_hscrollbar.Fn()(w)
}
func (w *ScrolledWindow) GetVScrollbar() *Scrollbar {
	return gtk_scrolled_window_get_vscrollbar.Fn()(w)
}
func (w *ScrolledWindow) SetPolicy(hscrollbarPolicy, vscrollbarPolicy PolicyType) {
	gtk_scrolled_window_set_policy.Fn()(w, hscrollbarPolicy, vscrollbarPolicy)
}
func (w *ScrolledWindow) GetPolicy() (hscrollbarPolicy, vscrollbarPolicy PolicyType) {
	gtk_scrolled_window_get_policy.Fn()(w, &hscrollbarPolicy, &vscrollbarPolicy)
	return
}
func (w *ScrolledWindow) SetPlacement(placement CornerType) {
	gtk_scrolled_window_set_placement.Fn()(w, placement)
}
func (w *ScrolledWindow) UnsetPlacement()          { gtk_scrolled_window_unset_placement.Fn()(w) }
func (w *ScrolledWindow) GetPlacement() CornerType { return gtk_scrolled_window_get_placement.Fn()(w) }
func (w *ScrolledWindow) SetHasFrame(hasFrame bool) {
	gtk_scrolled_window_set_has_frame.Fn()(w, cbool(hasFrame))
}
func (w *ScrolledWindow) GetHasFrame() bool { return gtk_scrolled_window_get_has_frame.Fn()(w) != 0 }
func (w *ScrolledWindow) GetMinContentWidth() int32 {
	return gtk_scrolled_window_get_min_content_width.Fn()(w)
}
func (w *ScrolledWindow) SetMinContentWidth(width int32) {
	gtk_scrolled_window_set_min_content_width.Fn()(w, width)
}
func (w *ScrolledWindow) GetMinContentHeight() int32 {
	return gtk_scrolled_window_get_min_content_height.Fn()(w)
}
func (w *ScrolledWindow) SetMinContentHeight(height int32) {
	gtk_scrolled_window_set_min_content_height.Fn()(w, height)
}
func (w *ScrolledWindow) SetKineticScrolling(kineticScrolling bool) {
	gtk_scrolled_window_set_kinetic_scrolling.Fn()(w, cbool(kineticScrolling))
}
func (w *ScrolledWindow) GetKineticScrolling() bool {
	return gtk_scrolled_window_get_kinetic_scrolling.Fn()(w) != 0
}
func (w *ScrolledWindow) SetOverlayScrolling(overlayScrolling bool) {
	gtk_scrolled_window_set_overlay_scrolling.Fn()(w, cbool(overlayScrolling))
}
func (w *ScrolledWindow) GetOverlayScrolling() bool {
	return gtk_scrolled_window_get_overlay_scrolling.Fn()(w) != 0
}
func (w *ScrolledWindow) SetMaxContentWidth(width int32) {
	gtk_scrolled_window_set_max_content_width.Fn()(w, width)
}
func (w *ScrolledWindow) GetMaxContentWidth() int32 {
	return gtk_scrolled_window_get_max_content_width.Fn()(w)
}
func (w *ScrolledWindow) SetMaxContentHeight(height int32) {
	gtk_scrolled_window_set_max_content_height.Fn()(w, height)
}
func (w *ScrolledWindow) GetMaxContentHeight() int32 {
	return gtk_scrolled_window_get_max_content_height.Fn()(w)
}
func (w *ScrolledWindow) SetPropagateNaturalWidth(propagate bool) {
	gtk_scrolled_window_set_propagate_natural_width.Fn()(w, cbool(propagate))
}
func (w *ScrolledWindow) GetPropagateNaturalWidth() bool {
	return gtk_scrolled_window_get_propagate_natural_width.Fn()(w) != 0
}
func (w *ScrolledWindow) SetPropagateNaturalHeight(propagate bool) {
	gtk_scrolled_window_set_propagate_natural_height.Fn()(w, cbool(propagate))
}
func (w *ScrolledWindow) GetPropagateNaturalHeight() bool {
	return gtk_scrolled_window_get_propagate_natural_height.Fn()(w) != 0
}
func (w *ScrolledWindow) SetChild(child WidgetIface) {
	gtk_scrolled_window_set_child.Fn()(w, GetWidgetIface(child))
}
func (w *ScrolledWindow) GetChild() *Widget { return gtk_scrolled_window_get_child.Fn()(w) }

func (w *ScrolledWindow) ConnectEdgeOvershot(sig func(w *ScrolledWindow, pos PositionType)) uint64 {
	return w.SignalConnect("edge-overshot", sig, nil)
}
func (w *ScrolledWindow) ConnectEdgeReached(sig func(w *ScrolledWindow, pos PositionType)) uint64 {
	return w.SignalConnect("edge-reached", sig, nil)
}
func (w *ScrolledWindow) ConnectMoveFocusOut(sig func(w *ScrolledWindow, dir DirectionType)) uint64 {
	return w.SignalConnect("move-focus-out", sig, nil)
}
func (w *ScrolledWindow) ConnectScrollChild(
	sig func(w *ScrolledWindow, scroll ScrollType, horizontal bool, _ uptr) bool) uint64 {
	return w.SignalConnect("scroll-child", sig, nil)
}

// #endregion

// #region ScrollInfo

type ScrollInfo struct{}

func GTypeScrollInfo() gobject.GType   { return gtk_scroll_info_get_type.Fn()() }
func NewScrollInfo() *ScrollInfo       { return gtk_scroll_info_new.Fn()() }
func (s *ScrollInfo) Ref() *ScrollInfo { return gtk_scroll_info_ref.Fn()(s) }
func (s *ScrollInfo) Unref()           { gtk_scroll_info_unref.Fn()(s) }
func (s *ScrollInfo) SetEnableHorizontal(horizontal bool) {
	gtk_scroll_info_set_enable_horizontal.Fn()(s, cbool(horizontal))
}
func (s *ScrollInfo) GetEnableHorizontal() bool {
	return gtk_scroll_info_get_enable_horizontal.Fn()(s) != 0
}
func (s *ScrollInfo) SetEnableVertical(vertical bool) {
	gtk_scroll_info_set_enable_vertical.Fn()(s, cbool(vertical))
}
func (s *ScrollInfo) GetEnableVertical() bool {
	return gtk_scroll_info_get_enable_vertical.Fn()(s) != 0
}

// #endregion

// #region SearchBar

type SearchBar struct{ Widget }

func GTypeSearchBar() gobject.GType               { return gtk_search_bar_get_type.Fn()() }
func NewSearchBar() *SearchBar                    { return gtk_search_bar_new.Fn()() }
func (b *SearchBar) ConnectEntry(entry *Editable) { gtk_search_bar_connect_entry.Fn()(b, entry) }
func (b *SearchBar) GetSearchMode() bool          { return gtk_search_bar_get_search_mode.Fn()(b) != 0 }
func (b *SearchBar) SetSearchMode(searchMode bool) {
	gtk_search_bar_set_search_mode.Fn()(b, cbool(searchMode))
}
func (b *SearchBar) GetShowCloseButton() bool {
	return gtk_search_bar_get_show_close_button.Fn()(b) != 0
}
func (b *SearchBar) SetShowCloseButton(visible bool) {
	gtk_search_bar_set_show_close_button.Fn()(b, cbool(visible))
}
func (b *SearchBar) SetKeyCaptureWidget(widget WidgetIface) {
	gtk_search_bar_set_key_capture_widget.Fn()(b, GetWidgetIface(widget))
}
func (b *SearchBar) GetKeyCaptureWidget() *Widget {
	return gtk_search_bar_get_key_capture_widget.Fn()(b)
}
func (b *SearchBar) SetChild(child WidgetIface) {
	gtk_search_bar_set_child.Fn()(b, GetWidgetIface(child))
}
func (b *SearchBar) GetChild() *Widget { return gtk_search_bar_get_child.Fn()(b) }

// #endregion

// #region SearchEntry

type SearchEntry struct {
	Widget
	Editable
}

func GTypeSearchEntry() gobject.GType { return gtk_search_entry_get_type.Fn()() }
func NewSearchEntry() *SearchEntry    { return gtk_search_entry_new.Fn()() }
func (e *SearchEntry) SetKeyCaptureWidget(widget WidgetIface) {
	gtk_search_entry_set_key_capture_widget.Fn()(e, GetWidgetIface(widget))
}
func (e *SearchEntry) GetKeyCaptureWidget() *Widget {
	return gtk_search_entry_get_key_capture_widget.Fn()(e)
}
func (e *SearchEntry) SetSearchDelay(delay uint32) {
	gtk_search_entry_set_search_delay.Fn()(e, delay)
}
func (e *SearchEntry) GetSearchDelay() uint32 {
	return gtk_search_entry_get_search_delay.Fn()(e)
}
func (e *SearchEntry) SetPlaceholderText(text string) {
	cText := cc.NewString(text)
	defer cText.Free()
	gtk_search_entry_set_placeholder_text.Fn()(e, cText)
}
func (e *SearchEntry) GetPlaceholderText() string {
	return gtk_search_entry_get_placeholder_text.Fn()(e).String()
}
func (e *SearchEntry) SetInputPurpose(purpose InputPurpose) {
	gtk_search_entry_set_input_purpose.Fn()(e, purpose)
}
func (e *SearchEntry) GetInputPurpose() InputPurpose {
	return gtk_search_entry_get_input_purpose.Fn()(e)
}
func (e *SearchEntry) SetInputHints(hints InputHints) {
	gtk_search_entry_set_input_hints.Fn()(e, hints)
}
func (e *SearchEntry) GetInputHints() InputHints {
	return gtk_search_entry_get_input_hints.Fn()(e)
}
func (s *SearchEntry) ConnectActivate(sig func(s *SearchEntry)) uint64 {
	return s.SignalConnect("activate", sig, nil)
}
func (s *SearchEntry) ConnectMextMatch(sig func(s *SearchEntry)) uint64 {
	return s.SignalConnect("next-match", sig, nil)
}
func (s *SearchEntry) ConnectPreviousMatch(sig func(s *SearchEntry)) uint64 {
	return s.SignalConnect("previous-match", sig, nil)
}
func (s *SearchEntry) ConnectSearchChanged(sig func(s *SearchEntry)) uint64 {
	return s.SignalConnect("search-changed", sig, nil)
}
func (s *SearchEntry) ConnectSearchStarted(sig func(s *SearchEntry)) uint64 {
	return s.SignalConnect("search-started", sig, nil)
}
func (s *SearchEntry) ConnectStopSearch(sig func(s *SearchEntry)) uint64 {
	return s.SignalConnect("stop-search", sig, nil)
}

// #endregion

// #region SectionModel

type SectionModelIfaceObj struct {
	GIface     gobject.GTypeInterface
	GetSection cc.Func
}

type SectionModelIface interface {
	GetSectionModelIface() *SectionModel
}

func GetSectionModelIface(iface SectionModelIface) *SectionModel {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetSectionModelIface()
}
func (s *SectionModel) GetSectionModelIface() *SectionModel { return s }

type SectionModel struct{ gio.GListModel }

func GTypeSectionModel() gobject.GType { return gtk_section_model_get_type.Fn()() }
func (m *SectionModel) GetSection(position uint32) (start, end uint32) {
	gtk_section_model_get_section.Fn()(m, position, &start, &end)
	return
}
func (m *SectionModel) SectionsChanged(position, nItems uint32) {
	gtk_section_model_sections_changed.Fn()(m, position, nItems)
}
func (m *SectionModel) ConnectSectionsChanged(fn func(m *SectionModel, pos, nItems uint32)) {
	toGobj(m).SignalConnect("sections-changed", fn, nil)
}

// #endregion

// #region SelectionFilterModel

type SelectionFilterModel struct {
	gobjCore
	gio.GListModel
}

func NewSelectionFilterModel(model SelectionModelIface) *SelectionFilterModel {
	return gtk_selection_filter_model_new.Fn()(GetSelectionModelIface(model))
}
func (m *SelectionFilterModel) SetModel(model SelectionModelIface) {
	gtk_selection_filter_model_set_model.Fn()(m, GetSelectionModelIface(model))
}
func (m *SelectionFilterModel) GetModel() *SelectionModel {
	return gtk_selection_filter_model_get_model.Fn()(m)
}

// #endregion

// #region SelectionModel

type SelectionModelIfaceObj struct {
	Parent gobject.GTypeInterface

	IsSelected          cc.Func // gboolean (* is_selected)  (GtkSelectionModel *model,guint  position);
	GetSelectionInRange cc.Func // GtkBitset * (* get_selection_in_range) (GtkSelectionModel *model,guint  position,guint  n_items);
	SelectItem          cc.Func // gboolean (* select_item)  (GtkSelectionModel *model,guint  position,gboolean unselect_rest);
	UnselectItem        cc.Func // gboolean (* unselect_item)  (GtkSelectionModel *model,guint  position);
	SelectRange         cc.Func // gboolean (* select_range)  (GtkSelectionModel *model,guint  position,guint  n_items,gboolean unselect_rest);
	UnselectRange       cc.Func // gboolean (* unselect_range)  (GtkSelectionModel *model,guint  position,guint  n_items);
	SelectAll           cc.Func // gboolean (* select_all)  (GtkSelectionModel *model);
	UnselectAll         cc.Func // gboolean (* unselect_all)  (GtkSelectionModel *model);
	SetSelection        cc.Func // gboolean (* set_selection)  (GtkSelectionModel *model,GtkBitset *selected,GtkBitset *mask);
}

type SelectionModelIface interface {
	GetSelectionModelIface() *SelectionModel
}

func GetSelectionModelIface(iface SelectionModelIface) *SelectionModel {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetSelectionModelIface()
}
func (s *SelectionModel) GetSelectionModelIface() *SelectionModel { return s }

type SelectionModel struct {
	gio.GListModel
}

func GTypeSelectionModel() gobject.GType { return gtk_selection_model_get_type.Fn()() }
func (m *SelectionModel) IsSelected(position uint32) bool {
	return gtk_selection_model_is_selected.Fn()(m, position) != 0
}
func (m *SelectionModel) GetSelection() *Bitset {
	return gtk_selection_model_get_selection.Fn()(m)
}
func (m *SelectionModel) GetSelectionInRange(position, nItems uint32) *Bitset {
	return gtk_selection_model_get_selection_in_range.Fn()(m, position, nItems)
}
func (m *SelectionModel) SelectItem(position uint32, unselectRest bool) bool {
	return gtk_selection_model_select_item.Fn()(m, position, cbool(unselectRest)) != 0
}
func (m *SelectionModel) UnselectItem(position uint32) bool {
	return gtk_selection_model_unselect_item.Fn()(m, position) != 0
}
func (m *SelectionModel) SelectRange(position, nItems uint32, unselectRest bool) bool {
	return gtk_selection_model_select_range.Fn()(m, position, nItems, cbool(unselectRest)) != 0
}
func (m *SelectionModel) UnselectRange(position, nItems uint32) bool {
	return gtk_selection_model_unselect_range.Fn()(m, position, nItems) != 0
}
func (m *SelectionModel) SelectAll() bool {
	return gtk_selection_model_select_all.Fn()(m) != 0
}
func (m *SelectionModel) UnselectAll() bool {
	return gtk_selection_model_unselect_all.Fn()(m) != 0
}
func (m *SelectionModel) SetSelection(selected, mask *Bitset) bool {
	return gtk_selection_model_set_selection.Fn()(m, selected, mask) != 0
}
func (m *SelectionModel) SelectionChanged(position, nItems uint32) {
	gtk_selection_model_selection_changed.Fn()(m, position, nItems)
}
func (m *SelectionModel) ConnectSelectionChanged(fn func(m *SelectionModel, pos, nItems uint32)) {
	toGobj(m).SignalConnect("selection-changed", fn, nil)
}

// #endregion

// #region Separator

type Separator struct {
	Widget
	Orientable
}

func GTypeSeparator() gobject.GType                   { return gtk_separator_get_type.Fn()() }
func NewSeparator(orientation Orientation) *Separator { return gtk_separator_new.Fn()(orientation) }

// #endregion

// #region Settings

type Settings struct {
	gobjCore
	StyleProvider
}

func GTypeSettings() gobject.GType  { return gtk_settings_get_type.Fn()() }
func GetDefaultSettings() *Settings { return gtk_settings_get_default.Fn()() }
func GetSettingsForDisplay(display *gdk.Display) *Settings {
	return gtk_settings_get_for_display.Fn()(display)
}
func (s *Settings) ResetProperty(name string) {
	n := cc.NewString(name)
	defer n.Free()
	gtk_settings_reset_property.Fn()(s, n)
}

// #endregion

// #region Shortcut

type Shortcut struct{ gobjCore }

func NewShortcut(trigger ShortcutTriggerIface, action ShortcutActionIface) *Shortcut {
	return gtk_shortcut_new.Fn()(GetShortcutTrigger(trigger), GetShortcutActionIface(action))
}
func NewShortcutWithArguments(trigger ShortcutTriggerIface, action ShortcutActionIface, format string, args ...interface{}) *Shortcut {
	f := cc.NewString(format)
	defer f.Free()
	return gtk_shortcut_new_with_arguments.FnVa()(GetShortcutTrigger(trigger), GetShortcutActionIface(action), f, args...)
}
func (s *Shortcut) GetTrigger() *ShortcutTrigger {
	return gtk_shortcut_get_trigger.Fn()(s)
}
func (s *Shortcut) SetTrigger(trigger ShortcutTriggerIface) {
	gtk_shortcut_set_trigger.Fn()(s, GetShortcutTrigger(trigger))
}
func (s *Shortcut) GetAction() *ShortcutAction {
	return gtk_shortcut_get_action.Fn()(s)
}
func (s *Shortcut) SetAction(act ShortcutActionIface) {
	gtk_shortcut_set_action.Fn()(s, GetShortcutActionIface(act))
}
func (s *Shortcut) GetArguments() *glib.GVariant {
	return gtk_shortcut_get_arguments.Fn()(s)
}
func (s *Shortcut) SetArguments(args *glib.GVariant) {
	gtk_shortcut_set_arguments.Fn()(s, args)
}

// #endregion

// #region ShortcutAction

type ShortcutActionIface interface {
	GetShortcutActionIface() *ShortcutAction
}

func GetShortcutActionIface(iface ShortcutActionIface) *ShortcutAction {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetShortcutActionIface()
}

func (s *ShortcutAction) GetShortcutActionIface() *ShortcutAction { return s }

type ShortcutAction struct{ gobjCore }

func (s *ShortcutAction) ToString() string {
	return gtk_shortcut_action_to_string.Fn()(s).TakeString()
}

//	func (s *ShortcutAction) Print(gstr *glib.GString) {
//		gtk_shortcut_action_print(s, gstr)
//	}

func (s *ShortcutAction) Activate(flags ShortcutActionFlags, widget *Widget, args *glib.GVariant) bool {
	return gtk_shortcut_action_activate.Fn()(s, flags, widget, args) != 0
}

func GTypeShortcutAction() gobject.GType { return gtk_shortcut_action_get_type.Fn()() }
func ShortcutActionParseString(str string) *ShortcutAction {
	cstr := cc.NewString(str)
	defer cstr.Free()
	return gtk_shortcut_action_parse_string.Fn()(cstr)
}

type NothingAction struct{ ShortcutAction }
type CallbackAction struct{ ShortcutAction }
type MnemonicAction struct{ ShortcutAction }
type ActivateAction struct{ ShortcutAction }
type SignalAction struct{ ShortcutAction }
type NamedAction struct{ ShortcutAction }

func GTypeNothingAction() gobject.GType  { return gtk_nothing_action_get_type.Fn()() }
func GTypeCallbackAction() gobject.GType { return gtk_callback_action_get_type.Fn()() }
func GTypeMnemonicAction() gobject.GType { return gtk_mnemonic_action_get_type.Fn()() }
func GTypeActivateAction() gobject.GType { return gtk_activate_action_get_type.Fn()() }
func GTypeSignalAction() gobject.GType   { return gtk_signal_action_get_type.Fn()() }
func GTypeNamedAction() gobject.GType    { return gtk_named_action_get_type.Fn()() }

func GetNothingAction() *NothingAction { return gtk_nothing_action_get.Fn()() }
func NewCallbackAction[T any](callback func(w *Widget, args *glib.GVariant, data *T) bool,
	data *T, destroyNotify func(data *T)) *CallbackAction {
	var cb, dn uptr
	if callback != nil {
		cb = cc.CbkRaw[func(w *Widget, a cc.Ptr, d uptr) int32](func(out, ins uptr) {
			ps := cc.Slice((*uptr)(ins), 3)
			if callback(*(**Widget)(ps[0]), *(**glib.GVariant)(ps[1]), *(**T)(ps[2])) {
				*(*int32)(out) = 1
			} else {
				*(*int32)(out) = 0
			}
		})
	}
	if destroyNotify != nil {
		dn = cc.Cbk(func(data *T) {
			if destroyNotify != nil {
				destroyNotify(data)
			}
			cc.CbkClose(cb)
			cc.CbkCloseLate(dn)
		})
	}
	return gtk_callback_action_new.Fn()(cb, uptr(data), dn)
}
func GetMnemonicAction() *MnemonicAction {
	return gtk_mnemonic_action_get.Fn()()
}
func GetActivateAction() *ActivateAction {
	return gtk_activate_action_get.Fn()()
}
func NewSignalAction(signalName string) *SignalAction {
	cSignalName := cc.NewString(signalName)
	defer cSignalName.Free()
	return gtk_signal_action_new.Fn()(cSignalName)
}
func (s *SignalAction) GetSignalName() string {
	return gtk_signal_action_get_signal_name.Fn()(s).String()
}
func NewNamedAction(name string) *NamedAction {
	cName := cc.NewString(name)
	defer cName.Free()
	return gtk_named_action_new.Fn()(cName)
}
func (n *NamedAction) GetActionName() string {
	return gtk_named_action_get_action_name.Fn()(n).String()
}

// #endregion

// #region ShortcutController

type ShortcutController struct {
	EventController
	Buildable
	gio.GListModel
}

func GTypeShortcutController() gobject.GType     { return gtk_shortcut_controller_get_type.Fn()() }
func NewShortcutController() *ShortcutController { return gtk_shortcut_controller_new.Fn()() }
func NewShortcutControllerForModel(model gio.GListModelIface) *ShortcutController {
	return gtk_shortcut_controller_new_for_model.Fn()(gio.GetGListModelIface(model))
}
func (c *ShortcutController) SetMnemonicsModifiers(modifiers gdk.ModifierType) {
	gtk_shortcut_controller_set_mnemonics_modifiers.Fn()(c, modifiers)
}
func (c *ShortcutController) GetMnemonicsModifiers() gdk.ModifierType {
	return gtk_shortcut_controller_get_mnemonics_modifiers.Fn()(c)
}
func (c *ShortcutController) SetScope(scope ShortcutScope) {
	gtk_shortcut_controller_set_scope.Fn()(c, scope)
}
func (c *ShortcutController) GetScope() ShortcutScope {
	return gtk_shortcut_controller_get_scope.Fn()(c)
}
func (c *ShortcutController) AddShortcut(shortcut *Shortcut) {
	gtk_shortcut_controller_add_shortcut.Fn()(c, shortcut)
}
func (c *ShortcutController) RemoveShortcut(shortcut *Shortcut) {
	gtk_shortcut_controller_remove_shortcut.Fn()(c, shortcut)
}

// #endregion

// #region ShortcutManager

type ShortcutManagerIfaceObj struct {
	GIface           gobject.GTypeInterface
	AddController    cc.Func // void (* add_controller) (GtkShortcutManager *self, GtkShortcutController *controller);
	RemoveController cc.Func // void (* remove_controller) (GtkShortcutManager *self, GtkShortcutController *controller);
}
type ShortcutManager struct{}

func GTypeShortcutManager() gobject.GType { return gtk_shortcut_manager_get_type.Fn()() }

// #endregion

// #region ShortcutTrigger

type ShortcutTriggerIface interface {
	GetShortcutTriggerIface() *ShortcutTrigger
}

func GetShortcutTrigger(iface ShortcutTriggerIface) *ShortcutTrigger {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetShortcutTriggerIface()
}
func (t *ShortcutTrigger) GetShortcutTriggerIface() *ShortcutTrigger { return t }

type ShortcutTrigger struct{ gobjCore }

func GTypeShortcutTrigger() gobject.GType { return gtk_shortcut_trigger_get_type.Fn()() }
func ParseShortcutTrigger(str string) *ShortcutTrigger {
	cstr := cc.NewString(str)
	defer cstr.Free()
	return gtk_shortcut_trigger_parse_string.Fn()(cstr)
}
func (t *ShortcutTrigger) ToString() string {
	return gtk_shortcut_trigger_to_string.Fn()(t).TakeString()
}
func (t *ShortcutTrigger) ToLabel(display *gdk.Display) string {
	return gtk_shortcut_trigger_to_label.Fn()(t, display).TakeString()
}
func (t *ShortcutTrigger) Hash() uint {
	return gtk_shortcut_trigger_hash.Fn()(t)
}
func (t *ShortcutTrigger) Equal(other ShortcutTriggerIface) bool {
	return gtk_shortcut_trigger_equal.Fn()(t, GetShortcutTrigger(other)) != 0
}
func (t *ShortcutTrigger) Compare(other ShortcutTriggerIface) int {
	return gtk_shortcut_trigger_compare.Fn()(t, GetShortcutTrigger(other))
}
func (t *ShortcutTrigger) Trigger(event *gdk.Event, enableMnemonics bool) gdk.KeyMatch {
	return gtk_shortcut_trigger_trigger.Fn()(t, event, cbool(enableMnemonics))
}

type NeverTrigger struct{ ShortcutTrigger }
type KeyvalTrigger struct{ ShortcutTrigger }
type MnemonicTrigger struct{ ShortcutTrigger }
type AlternativeTrigger struct{ ShortcutTrigger }

func GTypeNeverTrigger() gobject.GType       { return gtk_never_trigger_get_type.Fn()() }
func GTypeKeyvalTrigger() gobject.GType      { return gtk_keyval_trigger_get_type.Fn()() }
func GTypeMnemonicTrigger() gobject.GType    { return gtk_mnemonic_trigger_get_type.Fn()() }
func GTypeAlternativeTrigger() gobject.GType { return gtk_alternative_trigger_get_type.Fn()() }

func GetNeverTrigger() *NeverTrigger {
	return gtk_never_trigger_get.Fn()()
}
func NewKeyvalTrigger(keyval uint, modifiers gdk.ModifierType) *KeyvalTrigger {
	return gtk_keyval_trigger_new.Fn()(keyval, modifiers)
}
func (k *KeyvalTrigger) GetModifiers() gdk.ModifierType {
	return gtk_keyval_trigger_get_modifiers.Fn()(k)
}
func (k *KeyvalTrigger) GetKeyval() uint {
	return gtk_keyval_trigger_get_keyval.Fn()(k)
}
func NewMnemonicTrigger(keyval uint) *MnemonicTrigger {
	return gtk_mnemonic_trigger_new.Fn()(keyval)
}
func (m *MnemonicTrigger) GetKeyval() uint {
	return gtk_mnemonic_trigger_get_keyval.Fn()(m)
}
func NewAlternativeTrigger(first, second ShortcutTriggerIface) *AlternativeTrigger {
	return gtk_alternative_trigger_new.Fn()(GetShortcutTrigger(first), GetShortcutTrigger(second))
}
func (a *AlternativeTrigger) GetFirst() *ShortcutTrigger {
	return gtk_alternative_trigger_get_first.Fn()(a)
}
func (a *AlternativeTrigger) GetSecond() *ShortcutTrigger {
	return gtk_alternative_trigger_get_second.Fn()(a)
}

// #endregion

// #region SignalListItemFactory

type SignalListItemFactory struct{ ListItemFactory }

func GTypeSignalListItemFactory() gobject.GType { return gtk_signal_list_item_factory_get_type.Fn()() }
func NewSignalListItemFactory() *SignalListItemFactory {
	return gtk_signal_list_item_factory_new.Fn()()
}
func (l *SignalListItemFactory) ConnectBind(fn func(l *SignalListItemFactory, item *ListItem)) {
	l.SignalConnect("bind", fn, nil)
}
func (l *SignalListItemFactory) ConnectSetup(fn func(l *SignalListItemFactory, item *ListItem)) {
	l.SignalConnect("setup", fn, nil)
}
func (l *SignalListItemFactory) ConnectTeardown(fn func(l *SignalListItemFactory, item *ListItem)) {
	l.SignalConnect("teardown", fn, nil)
}
func (l *SignalListItemFactory) ConnectUnbind(fn func(l *SignalListItemFactory, item *ListItem)) {
	l.SignalConnect("unbind", fn, nil)
}

// #endregion

// #region SingleSelection

type SingleSelection struct {
	gobjCore
	gio.GListModel
	SectionModel
	SelectionModel
}

func NewSingleSelection(model gio.GListModelIface) *SingleSelection {
	return gtk_single_selection_new.Fn()(gio.GetGListModelIface(model))
}
func (s *SingleSelection) GetModel() *gio.GListModel {
	return gtk_single_selection_get_model.Fn()(s)
}
func (s *SingleSelection) SetModel(model *gio.GListModel) {
	gtk_single_selection_set_model.Fn()(s, model)
}
func (s *SingleSelection) GetSelected() uint32 {
	return gtk_single_selection_get_selected.Fn()(s)
}
func (s *SingleSelection) SetSelected(position uint32) {
	gtk_single_selection_set_selected.Fn()(s, position)
}
func (s *SingleSelection) GetSelectedItem() uptr {
	return gtk_single_selection_get_selected_item.Fn()(s)
}
func (s *SingleSelection) GetAutoselect() bool {
	return gtk_single_selection_get_autoselect.Fn()(s) != 0
}
func (s *SingleSelection) SetAutoselect(autoselect bool) {
	gtk_single_selection_set_autoselect.Fn()(s, cbool(autoselect))
}
func (s *SingleSelection) GetCanUnselect() bool {
	return gtk_single_selection_get_can_unselect.Fn()(s) != 0
}
func (s *SingleSelection) SetCanUnselect(canUnselect bool) {
	gtk_single_selection_set_can_unselect.Fn()(s, cbool(canUnselect))
}

// #endregion

// #region SizeGroup

type SizeGroupObj struct {
	Parent gobject.GObjectObj
}

type SizeGroup struct{ gobjCore }

func GTypeSizeGroup() gobject.GType              { return gtk_size_group_get_type.Fn()() }
func NewSizeGroup(mode SizeGroupMode) *SizeGroup { return gtk_size_group_new.Fn()(mode) }
func (sg *SizeGroup) SetMode(mode SizeGroupMode) { gtk_size_group_set_mode.Fn()(sg, mode) }
func (sg *SizeGroup) GetMode() SizeGroupMode     { return gtk_size_group_get_mode.Fn()(sg) }
func (sg *SizeGroup) AddWidget(widget WidgetIface) {
	gtk_size_group_add_widget.Fn()(sg, GetWidgetIface(widget))
}
func (sg *SizeGroup) RemoveWidget(widget WidgetIface) {
	gtk_size_group_remove_widget.Fn()(sg, GetWidgetIface(widget))
}
func (sg *SizeGroup) GetWidgets() *glib.GSList[Widget] {
	return (*glib.GSList[Widget])(gtk_size_group_get_widgets.Fn()(sg))
}

// #endregion

// #region RequestedSize

type RequestedSize struct {
	Data        uptr
	MinimumSize int32
	NaturalSize int32
}

func DistributeNaturalAllocation(extraSpace int32, nRequestedSizes uint32, sizes *RequestedSize) int32 {
	return gtk_distribute_natural_allocation.Fn()(extraSpace, nRequestedSizes, sizes)
}

// #endregion

// #region SliceListModel

type SliceListModel struct {
	SectionModel
}

func GTypeSliceListModel() gobject.GType { return gtk_slice_list_model_get_type.Fn()() }
func NewSliceListModel(model gio.GListModelIface, offset, size uint32) *SliceListModel {
	return gtk_slice_list_model_new.Fn()(gio.GetGListModelIface(model), offset, size)
}
func (m *SliceListModel) SetModel(model gio.GListModelIface) {
	gtk_slice_list_model_set_model.Fn()(m, gio.GetGListModelIface(model))
}
func (m *SliceListModel) GetModel() *gio.GListModel { return gtk_slice_list_model_get_model.Fn()(m) }
func (m *SliceListModel) SetOffset(offset uint32)   { gtk_slice_list_model_set_offset.Fn()(m, offset) }
func (m *SliceListModel) GetOffset() uint32         { return gtk_slice_list_model_get_offset.Fn()(m) }
func (m *SliceListModel) SetSize(size uint32)       { gtk_slice_list_model_set_size.Fn()(m, size) }
func (m *SliceListModel) GetSize() uint32           { return gtk_slice_list_model_get_size.Fn()(m) }

// #endregion

// #region Snapshot

type Snapshot struct{ gdk.Snapshot }

func GTypeSnapshot() gobject.GType { return gtk_snapshot_get_type.Fn()() }
func NewSnapshot() *Snapshot       { return gtk_snapshot_new.Fn()() }
func (s *Snapshot) FreeToNode() *gsk.RenderNode {
	return gtk_snapshot_free_to_node.Fn()(s)
}
func (s *Snapshot) FreeToPaintable(size *graphene.Size) *gdk.Paintable {
	return gtk_snapshot_free_to_paintable.Fn()(s, size)
}
func (s *Snapshot) ToNode() *gsk.RenderNode {
	return gtk_snapshot_to_node.Fn()(s)
}
func (s *Snapshot) ToPaintable(size *graphene.Size) *gdk.Paintable {
	return gtk_snapshot_to_paintable.Fn()(s, size)
}
func (s *Snapshot) PushDebug(message string, args ...interface{}) {
	msg := cc.NewString(message)
	defer msg.Free()
	gtk_snapshot_push_debug.FnVa()(s, msg, args...)
}
func (s *Snapshot) PushOpacity(opacity float64) {
	gtk_snapshot_push_opacity.Fn()(s, opacity)
}
func (s *Snapshot) PushBlur(radius float64) {
	gtk_snapshot_push_blur.Fn()(s, radius)
}
func (s *Snapshot) PushColorMatrix(colorMatrix *graphene.Matrix, colorOffset *graphene.Vec4) {
	gtk_snapshot_push_color_matrix.Fn()(s, colorMatrix, colorOffset)
}
func (s *Snapshot) PushRepeat(bounds, childBounds *graphene.Rect) {
	gtk_snapshot_push_repeat.Fn()(s, bounds, childBounds)
}
func (s *Snapshot) PushClip(bounds *graphene.Rect) { gtk_snapshot_push_clip.Fn()(s, bounds) }
func (s *Snapshot) PushRoundedClip(bounds *gsk.RoundedRect) {
	gtk_snapshot_push_rounded_clip.Fn()(s, bounds)
}
func (s *Snapshot) PushFill(path *gsk.Path, fillRule gsk.FillRule) {
	gtk_snapshot_push_fill.Fn()(s, path, fillRule)
}
func (s *Snapshot) PushStroke(path *gsk.Path, stroke *gsk.Stroke) {
	gtk_snapshot_push_stroke.Fn()(s, path, stroke)
}
func (s *Snapshot) PushShadow(shadow *gsk.Shadow, nShadows uint64) {
	gtk_snapshot_push_shadow.Fn()(s, shadow, nShadows)
}
func (s *Snapshot) PushBlend(blendMode gsk.BlendMode)  { gtk_snapshot_push_blend.Fn()(s, blendMode) }
func (s *Snapshot) PushMask(maskMode gsk.MaskMode)     { gtk_snapshot_push_mask.Fn()(s, maskMode) }
func (s *Snapshot) PushCrossFade(progress float64)     { gtk_snapshot_push_cross_fade.Fn()(s, progress) }
func (s *Snapshot) Pop()                               { gtk_snapshot_pop.Fn()(s) }
func (s *Snapshot) Save()                              { gtk_snapshot_save.Fn()(s) }
func (s *Snapshot) Restore()                           { gtk_snapshot_restore.Fn()(s) }
func (s *Snapshot) Transform(transform *gsk.Transform) { gtk_snapshot_transform.Fn()(s, transform) }
func (s *Snapshot) TransformMatrix(matrix *graphene.Matrix) {
	gtk_snapshot_transform_matrix.Fn()(s, matrix)
}
func (s *Snapshot) Translate(point *graphene.Point)     { gtk_snapshot_translate.Fn()(s, point) }
func (s *Snapshot) Translate3D(point *graphene.Point3D) { gtk_snapshot_translate_3d.Fn()(s, point) }
func (s *Snapshot) Rotate(angle float32) {
	gtk_snapshot_rotate.Fn()(s, angle)
}
func (s *Snapshot) Rotate3D(angle float32, axis *graphene.Vec3) {
	gtk_snapshot_rotate_3d.Fn()(s, angle, axis)
}
func (s *Snapshot) Scale(factorX, factorY float32) {
	gtk_snapshot_scale.Fn()(s, factorX, factorY)
}
func (s *Snapshot) Scale3D(factorX, factorY, factorZ float32) {
	gtk_snapshot_scale_3d.Fn()(s, factorX, factorY, factorZ)
}
func (s *Snapshot) Perspective(depth float32) {
	gtk_snapshot_perspective.Fn()(s, depth)
}
func (s *Snapshot) AppendNode(node *gsk.RenderNode) {
	gtk_snapshot_append_node.Fn()(s, node)
}

func (s *Snapshot) AppendCairo(bounds *graphene.Rect) *cairo.Context {
	return gtk_snapshot_append_cairo.Fn()(s, bounds)
}
func (s *Snapshot) AppendTexture(texture *gdk.Texture, bounds *graphene.Rect) {
	gtk_snapshot_append_texture.Fn()(s, texture, bounds)
}
func (s *Snapshot) AppendScaledTexture(texture *gdk.Texture, filter gsk.ScalingFilter, bounds *graphene.Rect) {
	gtk_snapshot_append_scaled_texture.Fn()(s, texture, filter, bounds)
}
func (s *Snapshot) AppendColor(color *gdk.RGBA, bounds *graphene.Rect) {
	gtk_snapshot_append_color.Fn()(s, color, bounds)
}
func (s *Snapshot) AppendLinearGradient(bounds *graphene.Rect, startPoint, endPoint *graphene.Point, stops *gsk.ColorStop, nStops uint64) {
	gtk_snapshot_append_linear_gradient.Fn()(s, bounds, startPoint, endPoint, stops, nStops)
}

func (s *Snapshot) AppendRepeatingLinearGradient(bounds *graphene.Rect, startPoint, endPoint *graphene.Point, stops *gsk.ColorStop, nStops uint64) {
	gtk_snapshot_append_repeating_linear_gradient.Fn()(s, bounds, startPoint, endPoint, stops, nStops)
}
func (s *Snapshot) AppendRadialGradient(bounds *graphene.Rect, center *graphene.Point, hradius, vradius, start, end float32, stops *gsk.ColorStop, nStops uint64) {
	gtk_snapshot_append_radial_gradient.Fn()(s, bounds, center, hradius, vradius, start, end, stops, nStops)
}
func (s *Snapshot) AppendRepeatingRadialGradient(bounds *graphene.Rect, center *graphene.Point, hradius, vradius, start, end float32, stops *gsk.ColorStop, nStops uint64) {
	gtk_snapshot_append_repeating_radial_gradient.Fn()(s, bounds, center, hradius, vradius, start, end, stops, nStops)
}
func (s *Snapshot) AppendConicGradient(bounds *graphene.Rect, center *graphene.Point, rotation float32, stops *gsk.ColorStop, nStops uint64) {
	gtk_snapshot_append_conic_gradient.Fn()(s, bounds, center, rotation, stops, nStops)
}
func (s *Snapshot) AppendBorder(outline *gsk.RoundedRect, borderWidth [4]float32, borderColor [4]gdk.RGBA) {
	gtk_snapshot_append_border.Fn()(s, outline, &borderWidth[0], &borderColor[0])
}
func (s *Snapshot) AppendInsetShadow(outline *gsk.RoundedRect, color *gdk.RGBA, dx, dy, spread, blurRadius float32) {
	gtk_snapshot_append_inset_shadow.Fn()(s, outline, color, dx, dy, spread, blurRadius)
}
func (s *Snapshot) AppendOutsetShadow(outline *gsk.RoundedRect, color *gdk.RGBA, dx, dy, spread, blurRadius float32) {
	gtk_snapshot_append_outset_shadow.Fn()(s, outline, color, dx, dy, spread, blurRadius)
}
func (s *Snapshot) AppendLayout(layout *pango.Layout, color *gdk.RGBA) {
	gtk_snapshot_append_layout.Fn()(s, layout, color)
}
func (s *Snapshot) AppendFill(path *gsk.Path, fillRule gsk.FillRule, color *gdk.RGBA) {
	gtk_snapshot_append_fill.Fn()(s, path, fillRule, color)
}
func (s *Snapshot) AppendStroke(path *gsk.Path, stroke *gsk.Stroke, color *gdk.RGBA) {
	gtk_snapshot_append_stroke.Fn()(s, path, stroke, color)
}

// #endregion

// #region Sorter

type SorterClass struct {
	Parent gobject.GObjectClass

	Compare  cc.Func
	GetOrder cc.Func
	_        [8]cc.Func //reserved
}

type SorterIface interface {
	GetSorterIface() *Sorter
}

func GetSorterIface(iface SorterIface) *Sorter {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetSorterIface()
}
func (s *Sorter) GetSorterIface() *Sorter { return s }

type Sorter struct{ gobjCore }

func GTypeSorter() gobject.GType                     { return gtk_sorter_get_type.Fn()() }
func (s *Sorter) Compare(item1, item2 uptr) Ordering { return gtk_sorter_compare.Fn()(s, item1, item2) }
func (s *Sorter) GetOrder() SorterOrder              { return gtk_sorter_get_order.Fn()(s) }
func (s *Sorter) Changed(change SorterChange)        { gtk_sorter_changed.Fn()(s, change) }
func (s *Sorter) ConnectChanged(fn func(s *Sorter, change SorterChange)) {
	s.SignalConnect("changed", fn, nil)
}

// #endregion

// #region SortListModel

type SortListModel struct{ SectionModel }

func GTypeSortListModel() gobject.GType { return gtk_sort_list_model_get_type.Fn()() }
func NewSortListModel(model gio.GListModelIface, sorter SorterIface) *SortListModel {
	return gtk_sort_list_model_new.Fn()(gio.GetGListModelIface(model), GetSorterIface(sorter))
}
func (m *SortListModel) SetSorter(sorter SorterIface) {
	gtk_sort_list_model_set_sorter.Fn()(m, GetSorterIface(sorter))
}
func (m *SortListModel) GetSorter() *Sorter { return gtk_sort_list_model_get_sorter.Fn()(m) }
func (m *SortListModel) SetSectionSorter(sorter SorterIface) {
	gtk_sort_list_model_set_section_sorter.Fn()(m, GetSorterIface(sorter))
}
func (m *SortListModel) GetSectionSorter() *Sorter {
	return gtk_sort_list_model_get_section_sorter.Fn()(m)
}
func (m *SortListModel) SetModel(model gio.GListModelIface) {
	gtk_sort_list_model_set_model.Fn()(m, gio.GetGListModelIface(model))
}
func (m *SortListModel) GetModel() *gio.GListModel { return gtk_sort_list_model_get_model.Fn()(m) }
func (m *SortListModel) SetIncremental(incremental bool) {
	gtk_sort_list_model_set_incremental.Fn()(m, cbool(incremental))
}
func (m *SortListModel) GetIncremental() bool {
	return gtk_sort_list_model_get_incremental.Fn()(m) != 0
}
func (m *SortListModel) GetPending() uint32 { return gtk_sort_list_model_get_pending.Fn()(m) }

// #endregion

// #region SpinButton

type SpinButton struct {
	Widget
	AccessibleRange
	Editable
	Orientable
}

func GTypeSpinButton() gobject.GType { return gtk_spin_button_get_type.Fn()() }
func NewSpinButton(adjustment *Adjustment, climbRate float64, digits uint32) *SpinButton {
	return gtk_spin_button_new.Fn()(adjustment, climbRate, digits)
}
func NewSpinButtonWithRange(min, max, step float64) *SpinButton {
	return gtk_spin_button_new_with_range.Fn()(min, max, step)
}
func (s *SpinButton) Configure(adjustment *Adjustment, climbRate float64, digits uint32) {
	gtk_spin_button_configure.Fn()(s, adjustment, climbRate, digits)
}
func (s *SpinButton) SetActivatesDefault(activatesDefault bool) {
	gtk_spin_button_set_activates_default.Fn()(s, cbool(activatesDefault))
}
func (s *SpinButton) GetActivatesDefault() bool {
	return gtk_spin_button_get_activates_default.Fn()(s) != 0
}
func (s *SpinButton) SetAdjustment(adjustment *Adjustment) {
	gtk_spin_button_set_adjustment.Fn()(s, adjustment)
}
func (s *SpinButton) GetAdjustment() *Adjustment {
	return gtk_spin_button_get_adjustment.Fn()(s)
}
func (s *SpinButton) SetDigits(digits uint32) { gtk_spin_button_set_digits.Fn()(s, digits) }
func (s *SpinButton) GetDigits() uint32       { return gtk_spin_button_get_digits.Fn()(s) }
func (s *SpinButton) SetIncrements(step, page float64) {
	gtk_spin_button_set_increments.Fn()(s, step, page)
}
func (s *SpinButton) GetIncrements() (step, page float64) {
	gtk_spin_button_get_increments.Fn()(s, &step, &page)
	return
}
func (s *SpinButton) SetRange(min, max float64) { gtk_spin_button_set_range.Fn()(s, min, max) }
func (s *SpinButton) GetRange() (min, max float64) {
	gtk_spin_button_get_range.Fn()(s, &min, &max)
	return
}
func (s *SpinButton) GetValue() float64      { return gtk_spin_button_get_value.Fn()(s) }
func (s *SpinButton) GetValueAsInt() int32   { return gtk_spin_button_get_value_as_int.Fn()(s) }
func (s *SpinButton) SetValue(value float64) { gtk_spin_button_set_value.Fn()(s, value) }
func (s *SpinButton) SetUpdatePolicy(policy SpinButtonUpdatePolicy) {
	gtk_spin_button_set_update_policy.Fn()(s, policy)
}
func (s *SpinButton) GetUpdatePolicy() SpinButtonUpdatePolicy {
	return gtk_spin_button_get_update_policy.Fn()(s)
}
func (s *SpinButton) SetNumeric(numeric bool) { gtk_spin_button_set_numeric.Fn()(s, cbool(numeric)) }
func (s *SpinButton) GetNumeric() bool        { return gtk_spin_button_get_numeric.Fn()(s) != 0 }
func (s *SpinButton) Spin(direction SpinType, increment float64) {
	gtk_spin_button_spin.Fn()(s, direction, increment)
}
func (s *SpinButton) SetWrap(wrap bool) { gtk_spin_button_set_wrap.Fn()(s, cbool(wrap)) }
func (s *SpinButton) GetWrap() bool     { return gtk_spin_button_get_wrap.Fn()(s) != 0 }
func (s *SpinButton) SetSnapToTicks(snapToTicks bool) {
	gtk_spin_button_set_snap_to_ticks.Fn()(s, cbool(snapToTicks))
}
func (s *SpinButton) GetSnapToTicks() bool { return gtk_spin_button_get_snap_to_ticks.Fn()(s) != 0 }
func (s *SpinButton) SetClimbRate(climbRate float64) {
	gtk_spin_button_set_climb_rate.Fn()(s, climbRate)
}
func (s *SpinButton) GetClimbRate() float64 { return gtk_spin_button_get_climb_rate.Fn()(s) }
func (s *SpinButton) Update()               { gtk_spin_button_update.Fn()(s) }
func (s *SpinButton) ConnectActivate(sig func(s *SpinButton)) uint64 {
	return s.SignalConnect("activate", sig, nil)
}
func (s *SpinButton) ConnectChangeValue(sig func(s *SpinButton, scroll ScrollType)) uint64 {
	return s.SignalConnect("change-value", sig, nil)
}
func (s *SpinButton) ConnectInput(sig func(s *SpinButton, newValue *float64, _ uptr) bool) uint64 {
	return s.SignalConnect("input", sig, nil)
}
func (s *SpinButton) ConnectOutput(sig func(s *SpinButton, _ uptr) bool) uint64 {
	return s.SignalConnect("output", sig, nil)
}
func (s *SpinButton) ConnectValueChanged(sig func(s *SpinButton)) uint64 {
	return s.SignalConnect("value-changed", sig, nil)
}
func (s *SpinButton) ConnectWrapped(sig func(s *SpinButton)) uint64 {
	return s.SignalConnect("wrapped", sig, nil)
}

// #endregion

// #region Spinner

type Spinner struct{ Widget }

func GTypeSpinner() gobject.GType            { return gtk_spinner_get_type.Fn()() }
func NewSpinner() *Spinner                   { return gtk_spinner_new.Fn()() }
func (s *Spinner) Start()                    { gtk_spinner_start.Fn()(s) }
func (s *Spinner) Stop()                     { gtk_spinner_stop.Fn()(s) }
func (s *Spinner) SetSpinning(spinning bool) { gtk_spinner_set_spinning.Fn()(s, cbool(spinning)) }
func (s *Spinner) GetSpinning() bool         { return gtk_spinner_get_spinning.Fn()(s) != 0 }

// #endregion

// #region Stack

type Stack struct{ Widget }

type StackPage struct{ gobjCore }

func GTypeStackPage() gobject.GType          { return gtk_stack_page_get_type.Fn()() }
func (p *StackPage) GetChild() *Widget       { return gtk_stack_page_get_child.Fn()(p) }
func (p *StackPage) GetVisible() bool        { return gtk_stack_page_get_visible.Fn()(p) != 0 }
func (p *StackPage) SetVisible(visible bool) { gtk_stack_page_set_visible.Fn()(p, cbool(visible)) }
func (p *StackPage) GetNeedsAttention() bool { return gtk_stack_page_get_needs_attention.Fn()(p) != 0 }
func (p *StackPage) SetNeedsAttention(setting bool) {
	gtk_stack_page_set_needs_attention.Fn()(p, cbool(setting))
}
func (p *StackPage) GetUseUnderline() bool { return gtk_stack_page_get_use_underline.Fn()(p) != 0 }
func (p *StackPage) SetUseUnderline(setting bool) {
	gtk_stack_page_set_use_underline.Fn()(p, cbool(setting))
}
func (p *StackPage) GetName() string { return gtk_stack_page_get_name.Fn()(p).String() }
func (p *StackPage) SetName(setting string) {
	cSetting := cc.NewString(setting)
	defer cSetting.Free()
	gtk_stack_page_set_name.Fn()(p, cSetting)
}
func (p *StackPage) GetTitle() string { return gtk_stack_page_get_title.Fn()(p).String() }
func (p *StackPage) SetTitle(setting string) {
	cSetting := cc.NewString(setting)
	defer cSetting.Free()
	gtk_stack_page_set_title.Fn()(p, cSetting)
}
func (p *StackPage) GetIconName() string { return gtk_stack_page_get_icon_name.Fn()(p).String() }
func (p *StackPage) SetIconName(setting string) {
	cSetting := cc.NewString(setting)
	defer cSetting.Free()
	gtk_stack_page_set_icon_name.Fn()(p, cSetting)
}

func GTypeStack() gobject.GType { return gtk_stack_get_type.Fn()() }
func NewStack() *Stack          { return gtk_stack_new.Fn()() }
func (s *Stack) AddChild(child WidgetIface) *StackPage {
	return gtk_stack_add_child.Fn()(s, GetWidgetIface(child))
}
func (s *Stack) AddNamed(child WidgetIface, name string) *StackPage {
	cName := cc.NewString(name)
	defer cName.Free()
	return gtk_stack_add_named.Fn()(s, GetWidgetIface(child), cName)
}
func (s *Stack) AddTitled(child WidgetIface, name, title string) *StackPage {
	cName := cc.NewString(name)
	defer cName.Free()
	cTitle := cc.NewString(title)
	defer cTitle.Free()
	return gtk_stack_add_titled.Fn()(s, GetWidgetIface(child), cName, cTitle)
}
func (s *Stack) Remove(child WidgetIface) { gtk_stack_remove.Fn()(s, GetWidgetIface(child)) }
func (s *Stack) GetPage(child WidgetIface) *StackPage {
	return gtk_stack_get_page.Fn()(s, GetWidgetIface(child))
}
func (s *Stack) GetChildByName(name string) *Widget {
	cName := cc.NewString(name)
	defer cName.Free()
	return gtk_stack_get_child_by_name.Fn()(s, cName)
}
func (s *Stack) SetVisibleChild(child WidgetIface) {
	gtk_stack_set_visible_child.Fn()(s, GetWidgetIface(child))
}
func (s *Stack) GetVisibleChild() *Widget { return gtk_stack_get_visible_child.Fn()(s) }
func (s *Stack) SetVisibleChildName(name string) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_stack_set_visible_child_name.Fn()(s, cName)
}
func (s *Stack) GetVisibleChildName() string {
	return gtk_stack_get_visible_child_name.Fn()(s).String()
}
func (s *Stack) SetVisibleChildFull(name string, transition StackTransitionType) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_stack_set_visible_child_full.Fn()(s, cName, transition)
}
func (s *Stack) SetHHomogeneous(hhomogeneous bool) {
	gtk_stack_set_hhomogeneous.Fn()(s, cbool(hhomogeneous))
}
func (s *Stack) GetHHomogeneous() bool { return gtk_stack_get_hhomogeneous.Fn()(s) != 0 }
func (s *Stack) SetVHomogeneous(vhomogeneous bool) {
	gtk_stack_set_vhomogeneous.Fn()(s, cbool(vhomogeneous))
}
func (s *Stack) GetVHomogeneous() bool { return gtk_stack_get_vhomogeneous.Fn()(s) != 0 }
func (s *Stack) SetTransitionDuration(duration uint32) {
	gtk_stack_set_transition_duration.Fn()(s, duration)
}
func (s *Stack) GetTransitionDuration() uint32 { return gtk_stack_get_transition_duration.Fn()(s) }
func (s *Stack) SetTransitionType(transition StackTransitionType) {
	gtk_stack_set_transition_type.Fn()(s, transition)
}
func (s *Stack) GetTransitionType() StackTransitionType {
	return gtk_stack_get_transition_type.Fn()(s)
}
func (s *Stack) GetTransitionRunning() bool { return gtk_stack_get_transition_running.Fn()(s) != 0 }
func (s *Stack) SetInterpolateSize(interpolateSize bool) {
	gtk_stack_set_interpolate_size.Fn()(s, cbool(interpolateSize))
}
func (s *Stack) GetInterpolateSize() bool  { return gtk_stack_get_interpolate_size.Fn()(s) != 0 }
func (s *Stack) GetPages() *SelectionModel { return gtk_stack_get_pages.Fn()(s) }

// #endregion

// #region StackSidebar

type StackSidebar struct{ Widget }

func GTypeStackSidebar() gobject.GType        { return gtk_stack_sidebar_get_type.Fn()() }
func NewStackSidebar() *StackSidebar          { return gtk_stack_sidebar_new.Fn()() }
func (s *StackSidebar) SetStack(stack *Stack) { gtk_stack_sidebar_set_stack.Fn()(s, stack) }
func (s *StackSidebar) GetStack() *Stack      { return gtk_stack_sidebar_get_stack.Fn()(s) }

// #endregion

// #region StackSwitcher

type StackSwitcher struct {
	Widget
	Orientable
}

func GTypeStackSwitcher() gobject.GType        { return gtk_stack_switcher_get_type.Fn()() }
func NewStackSwitcher() *StackSwitcher         { return gtk_stack_switcher_new.Fn()() }
func (s *StackSwitcher) SetStack(stack *Stack) { gtk_stack_switcher_set_stack.Fn()(s, stack) }
func (s *StackSwitcher) GetStack() *Stack      { return gtk_stack_switcher_get_stack.Fn()(s) }

// #endregion

// #region StringFilter

type StringFilter struct{ Filter }

func GTypeStringFilter() gobject.GType { return gtk_string_filter_get_type.Fn()() }
func NewStringFilter(expression ExpressionIface) *StringFilter {
	return gtk_string_filter_new.Fn()(GetExpressionIface(expression))
}
func (f *StringFilter) GetSearch() string { return gtk_string_filter_get_search.Fn()(f).String() }
func (f *StringFilter) SetSearch(search string) {
	cSearch := cc.NewString(search)
	defer cSearch.Free()
	gtk_string_filter_set_search.Fn()(f, cSearch)
}
func (f *StringFilter) GetExpression() *Expression { return gtk_string_filter_get_expression.Fn()(f) }
func (f *StringFilter) SetExpression(expression ExpressionIface) {
	gtk_string_filter_set_expression.Fn()(f, GetExpressionIface(expression))
}
func (f *StringFilter) GetIgnoreCase() bool { return gtk_string_filter_get_ignore_case.Fn()(f) != 0 }
func (f *StringFilter) SetIgnoreCase(ignoreCase bool) {
	gtk_string_filter_set_ignore_case.Fn()(f, cbool(ignoreCase))
}
func (f *StringFilter) GetMatchMode() StringFilterMatchMode {
	return gtk_string_filter_get_match_mode.Fn()(f)
}
func (f *StringFilter) SetMatchMode(mode StringFilterMatchMode) {
	gtk_string_filter_set_match_mode.Fn()(f, mode)
}

// #endregion

// #region StringList

type StringObject struct{ gobjCore }

func GTypeStringObject() gobject.GType { return gtk_string_object_get_type.Fn()() }
func NewStringObject(str string) *StringObject {
	s := cc.NewString(str)
	defer s.Free()
	return gtk_string_object_new.Fn()(s)
}
func (s *StringObject) GetString() string { return gtk_string_object_get_string.Fn()(s).String() }
func (s *StringObject) String() string    { return gtk_string_object_get_string.Fn()(s).String() }

type StringList struct {
	gobjCore
	gio.GListModel
	Buildable
}

func GTypeStringList() gobject.GType { return gtk_string_list_get_type.Fn()() }
func NewStringList(strings []string) *StringList {
	strPtr := cc.NewStrings(strings)
	defer strPtr.Free()
	return gtk_string_list_new.Fn()(strPtr)
}
func (s *StringList) Append(str string) {
	cstr := cc.NewString(str)
	defer cstr.Free()
	gtk_string_list_append.Fn()(s, cstr)
}
func (s *StringList) Take(str string)        { gtk_string_list_take.Fn()(s, cc.NewString(str)) }
func (s *StringList) Remove(position uint32) { gtk_string_list_remove.Fn()(s, position) }
func (s *StringList) Splice(position, nRemovals uint32, additions []string) {
	addPtr := cc.NewStrings(additions)
	defer addPtr.Free()
	gtk_string_list_splice.Fn()(s, position, nRemovals, addPtr)
}
func (s *StringList) GetString(position uint32) string {
	return gtk_string_list_get_string.Fn()(s, position).RefString()
}
func (s *StringList) Find(str string) uint32 {
	cstr := cc.NewString(str)
	defer cstr.Free()
	return gtk_string_list_find.Fn()(s, cstr)
}

// #endregion

// #region StringSorter

type StringSorter struct{ Sorter }

func GTypeStringSorter() gobject.GType { return gtk_string_sorter_get_type.Fn()() }
func NewStringSorter(expression ExpressionIface) *StringSorter {
	return gtk_string_sorter_new.Fn()(GetExpressionIface(expression))
}
func (s *StringSorter) GetExpression() *Expression { return gtk_string_sorter_get_expression.Fn()(s) }
func (s *StringSorter) SetExpression(expression ExpressionIface) {
	gtk_string_sorter_set_expression.Fn()(s, GetExpressionIface(expression))
}
func (s *StringSorter) GetIgnoreCase() bool { return gtk_string_sorter_get_ignore_case.Fn()(s) != 0 }
func (s *StringSorter) SetIgnoreCase(ignoreCase bool) {
	gtk_string_sorter_set_ignore_case.Fn()(s, cbool(ignoreCase))
}

// #endregion

// #region StyleProvider

type StyleProviderPriority uint32

const (
	StyleProviderPriorityFallback    StyleProviderPriority = 1
	StyleProviderPriorityTheme       StyleProviderPriority = 200
	StyleProviderPrioritySettings    StyleProviderPriority = 400
	StyleProviderPriorityApplication StyleProviderPriority = 600
	StyleProviderPriorityUser        StyleProviderPriority = 800
)

type StyleProviderIface interface{ GetStyleProviderIface() *StyleProvider }

func GetStyleProviderIface(ifac StyleProviderIface) *StyleProvider {
	if anyptr(ifac) == nil {
		return nil
	}
	return ifac.GetStyleProviderIface()
}

func (provider *StyleProvider) GetStyleProviderIface() *StyleProvider { return provider }

type StyleProvider struct{}

func GTypeStyleProvider() gobject.GType { return gtk_style_provider_get_type.Fn()() }
func AddProviderForDisplay(display *gdk.Display, provider StyleProviderIface, priority StyleProviderPriority) {
	gtk_style_context_add_provider_for_display.Fn()(display, GetStyleProviderIface(provider), priority)
}
func RemoveProviderForDisplay(display *gdk.Display, provider StyleProviderIface) {
	gtk_style_context_remove_provider_for_display.Fn()(display, GetStyleProviderIface(provider))
}

// #endregion

// #region Switch

type Switch struct {
	Widget
	Buildable
}

func GTypeSwitch() gobject.GType           { return gtk_switch_get_type.Fn()() }
func NewSwitch() *Switch                   { return gtk_switch_new.Fn()() }
func (sw *Switch) SetActive(isActive bool) { gtk_switch_set_active.Fn()(sw, cbool(isActive)) }
func (sw *Switch) GetActive() bool         { return gtk_switch_get_active.Fn()(sw) != 0 }
func (sw *Switch) SetState(state bool)     { gtk_switch_set_state.Fn()(sw, cbool(state)) }
func (sw *Switch) GetState() bool          { return gtk_switch_get_state.Fn()(sw) != 0 }
func (s *Switch) ConnectActivate(sig func(s *Switch)) uint64 {
	return s.SignalConnect("activate", sig, nil)
}
func (s *Switch) ConnectStateSet(sig func(s *Switch, sate bool, _ uptr) bool) uint64 {
	return s.SignalConnect("state-set", sig, nil)
}

// #endregion

// #region SymbolicPaintable

type SymbolicPaintableIfaceObj struct {
	GIface           gobject.GTypeInterface
	SnapshotSymbolic cc.Func //void (* snapshot_symbolic) (GtkSymbolicPaintable *paintable, GdkSnapshot *snapshot, double width, double height, const GdkRGBA *colors, gsize n_colors);
}

type SymbolicPaintableIface interface {
	GetSymbolicPaintableIface() *SymbolicPaintable
}

func GetSymbolicPaintableIface(iface SymbolicPaintableIface) *SymbolicPaintable {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetSymbolicPaintableIface()
}

type SymbolicPaintable struct{ gdk.Paintable }

func (p *SymbolicPaintable) SnapshotSymbolic(snapshot *Snapshot, width, height float64, colors *gdk.RGBA, nColors uint64) {
	gtk_symbolic_paintable_snapshot_symbolic.Fn()(p, snapshot, width, height, colors, nColors)
}

// #endregion

// #region TestAtContext

func TestAccessibleHasRole(a AccessibleIface, role AccessibleRole) bool {
	return gtk_test_accessible_has_role.Fn()(GetAccessibleIface(a), role) != 0
}
func TestAccessibleHasProperty(a AccessibleIface, property AccessibleProperty) bool {
	return gtk_test_accessible_has_property.Fn()(GetAccessibleIface(a), property) != 0
}
func TestAccessibleHasRelation(a AccessibleIface, relation AccessibleRelation) bool {
	return gtk_test_accessible_has_relation.Fn()(GetAccessibleIface(a), relation) != 0
}
func TestAccessibleHasState(a AccessibleIface, state AccessibleState) bool {
	return gtk_test_accessible_has_state.Fn()(GetAccessibleIface(a), state) != 0
}
func TestAccessibleCheckProperty(a AccessibleIface, property AccessibleProperty, args ...interface{}) string {
	return gtk_test_accessible_check_property.FnVa()(GetAccessibleIface(a), property, args...).TakeString()
}
func TestAccessibleCheckRelation(a AccessibleIface, relation AccessibleRelation, args ...interface{}) string {
	return gtk_test_accessible_check_relation.FnVa()(GetAccessibleIface(a), relation, args...).TakeString()
}
func TestAccessibleCheckState(a AccessibleIface, state AccessibleState, args ...interface{}) string {
	return gtk_test_accessible_check_state.FnVa()(GetAccessibleIface(a), state, args...).TakeString()
}
func TestAccessibleAssertionMessageRole(domain, file, function, expr string, line int32, a AccessibleIface, expected, actual AccessibleRole) {
	cDomain := cc.NewString(domain)
	defer cDomain.Free()
	cFile := cc.NewString(file)
	defer cFile.Free()
	cFunc := cc.NewString(function)
	defer cFunc.Free()
	cExpr := cc.NewString(expr)
	defer cExpr.Free()
	gtk_test_accessible_assertion_message_role.Fn()(cDomain, cFile, cFunc, cExpr, line, GetAccessibleIface(a), expected, actual)
}

// #endregion

// #region TestUtils
func TestInit(argcp *int32, argvp *cc.Strings, args ...interface{}) {
	gtk_test_init.FnVa()(argcp, argvp, args...)
}
func TestRegisterAllTypes() { gtk_test_register_all_types.Fn()() }
func TestListAllTypes() []gobject.GType {
	var n uint32
	return cc.Slice(gtk_test_list_all_types.Fn()(&n), n)
}
func TestWidgetWaitForDraw(w WidgetIface) { gtk_test_widget_wait_for_draw.Fn()(GetWidgetIface(w)) }

// #endregion

// #region Text

type TextObj struct {
	Parent WidgetObj
}

type Text struct {
	Widget
	AccessibleText
	Editable
}

func GTypeText() gobject.GType                    { return gtk_text_get_type.Fn()() }
func NewText() *Text                              { return gtk_text_new.Fn()() }
func NewTextWithBuffer(buffer *EntryBuffer) *Text { return gtk_text_new_with_buffer.Fn()(buffer) }
func (t *Text) GetBuffer() *EntryBuffer           { return gtk_text_get_buffer.Fn()(t) }
func (t *Text) SetBuffer(buffer *EntryBuffer)     { gtk_text_set_buffer.Fn()(t, buffer) }
func (t *Text) SetVisibility(visible bool)        { gtk_text_set_visibility.Fn()(t, cbool(visible)) }
func (t *Text) GetVisibility() bool               { return gtk_text_get_visibility.Fn()(t) != 0 }
func (t *Text) SetInvisibleChar(ch uint32)        { gtk_text_set_invisible_char.Fn()(t, ch) }
func (t *Text) GetInvisibleChar() uint32          { return gtk_text_get_invisible_char.Fn()(t) }
func (t *Text) UnsetInvisibleChar()               { gtk_text_unset_invisible_char.Fn()(t) }
func (t *Text) SetOverwriteMode(overwrite bool) {
	gtk_text_set_overwrite_mode.Fn()(t, cbool(overwrite))
}
func (t *Text) GetOverwriteMode() bool    { return gtk_text_get_overwrite_mode.Fn()(t) != 0 }
func (t *Text) SetMaxLength(length int32) { gtk_text_set_max_length.Fn()(t, length) }
func (t *Text) GetMaxLength() int32       { return gtk_text_get_max_length.Fn()(t) }
func (t *Text) GetTextLength() uint16     { return gtk_text_get_text_length.Fn()(t) }
func (t *Text) SetActivatesDefault(activates bool) {
	gtk_text_set_activates_default.Fn()(t, cbool(activates))
}
func (t *Text) GetActivatesDefault() bool  { return gtk_text_get_activates_default.Fn()(t) != 0 }
func (t *Text) GetPlaceholderText() string { return gtk_text_get_placeholder_text.Fn()(t).String() }
func (t *Text) SetPlaceholderText(text string) {
	cText := cc.NewString(text)
	defer cText.Free()
	gtk_text_set_placeholder_text.Fn()(t, cText)
}
func (t *Text) SetInputPurpose(purpose InputPurpose) { gtk_text_set_input_purpose.Fn()(t, purpose) }
func (t *Text) GetInputPurpose() InputPurpose        { return gtk_text_get_input_purpose.Fn()(t) }
func (t *Text) SetInputHints(hints InputHints)       { gtk_text_set_input_hints.Fn()(t, hints) }
func (t *Text) GetInputHints() InputHints            { return gtk_text_get_input_hints.Fn()(t) }
func (t *Text) SetAttributes(attrs *pango.AttrList)  { gtk_text_set_attributes.Fn()(t, attrs) }
func (t *Text) GetAttributes() *pango.AttrList       { return gtk_text_get_attributes.Fn()(t) }
func (t *Text) SetTabs(tabs *pango.TabArray)         { gtk_text_set_tabs.Fn()(t, tabs) }
func (t *Text) GetTabs() *pango.TabArray             { return gtk_text_get_tabs.Fn()(t) }
func (t *Text) GrabFocusWithoutSelecting() bool {
	return gtk_text_grab_focus_without_selecting.Fn()(t) != 0
}
func (t *Text) SetExtraMenu(model *gio.GMenuModel) { gtk_text_set_extra_menu.Fn()(t, model) }
func (t *Text) GetExtraMenu() *gio.GMenuModel      { return gtk_text_get_extra_menu.Fn()(t) }
func (t *Text) SetEnableEmojiCompletion(enable bool) {
	gtk_text_set_enable_emoji_completion.Fn()(t, cbool(enable))
}
func (t *Text) GetEnableEmojiCompletion() bool {
	return gtk_text_get_enable_emoji_completion.Fn()(t) != 0
}
func (t *Text) SetPropagateTextWidth(propagate bool) {
	gtk_text_set_propagate_text_width.Fn()(t, cbool(propagate))
}
func (t *Text) GetPropagateTextWidth() bool { return gtk_text_get_propagate_text_width.Fn()(t) != 0 }
func (t *Text) SetTruncateMultiline(truncate bool) {
	gtk_text_set_truncate_multiline.Fn()(t, cbool(truncate))
}
func (t *Text) GetTruncateMultiline() bool { return gtk_text_get_truncate_multiline.Fn()(t) != 0 }
func (t *Text) ComputeCursorExtents(position uint64, strong, weak *graphene.Rect) {
	gtk_text_compute_cursor_extents.Fn()(t, position, strong, weak)
}
func (s *Text) ConnectActivate(sig func(s *Text)) uint64 {
	return s.SignalConnect("activate", sig, nil)
}
func (s *Text) ConnectBackspace(sig func(s *Text)) uint64 {
	return s.SignalConnect("backspace", sig, nil)
}
func (s *Text) ConnectCopyClipboard(sig func(s *Text)) uint64 {
	return s.SignalConnect("copy-clipboard", sig, nil)
}
func (s *Text) ConnectCutClipboard(sig func(s *Text)) uint64 {
	return s.SignalConnect("cut-clipboard", sig, nil)
}
func (s *Text) ConnectDeleteFromCursor(sig func(s *Text, typ *DeleteType, count int32)) uint64 {
	return s.SignalConnect("delete-from-cursor", sig, nil)
}
func (s *Text) ConnectInsertAtCursor(sig func(s *Text, str cc.String)) uint64 {
	return s.SignalConnect("insert-at-cursor", sig, nil)
}
func (s *Text) ConnectInsertEmoji(sig func(s *Text)) uint64 {
	return s.SignalConnect("insert-emoji", sig, nil)
}
func (s *Text) ConnectMoveCursor(
	sig func(s *Text, step MovementStep, count int32, extend bool)) uint64 {
	return s.SignalConnect("move-cursor", sig, nil)
}
func (s *Text) ConnectPasteClipboard(sig func(s *Text)) uint64 {
	return s.SignalConnect("paste-clipboard", sig, nil)
}
func (s *Text) ConnectPreeditChanged(sig func(s *Text, preedit cc.String)) uint64 {
	return s.SignalConnect("preedit-changed", sig, nil)
}
func (s *Text) ConnectToggleOverwrite(sig func(s *Text)) uint64 {
	return s.SignalConnect("toggle-overwrite", sig, nil)
}

// #endregion

// #region TextBuffer

type TextBufferObj struct {
	Parent gobject.GObjectObj
	Priv   uptr
}

type TextBufferClass struct {
	Parent gobject.GObjectClass

	InsertText         cc.Func // void (* insert_text) (GtkTextBuffer *buffer,GtkTextIter *pos,const char *new_text,int new_text_length);
	InsertPaintable    cc.Func // void (* insert_paintable) (GtkTextBuffer *buffer,GtkTextIter *iter,GdkPaintable *paintable);
	InsertChild_anchor cc.Func // void (* insert_child_anchor) (GtkTextBuffer *buffer,GtkTextIter *iter,GtkTextChildAnchor *anchor);
	DeleteRange        cc.Func // void (* delete_range) (GtkTextBuffer *buffer,GtkTextIter *start,GtkTextIter *end);
	Changed            cc.Func // void (* changed) (GtkTextBuffer *buffer);
	ModifiedChanged    cc.Func // void (* modified_changed) (GtkTextBuffer *buffer);
	MarkSet            cc.Func // void (* mark_set) (GtkTextBuffer *buffer,const GtkTextIter *location,GtkTextMark *mark);
	MarkDeleted        cc.Func // void (* mark_deleted) (GtkTextBuffer *buffer,GtkTextMark *mark);
	ApplyTag           cc.Func // void (* apply_tag) (GtkTextBuffer *buffer,GtkTextTag *tag,const GtkTextIter *start,const GtkTextIter *end);
	RemoveTag          cc.Func // void (* remove_tag) (GtkTextBuffer *buffer,GtkTextTag *tag,const GtkTextIter *start,const GtkTextIter *end);
	BeginUserAction    cc.Func // void (* begin_user_action) (GtkTextBuffer *buffer);
	EndUserAction      cc.Func // void (* end_user_action) (GtkTextBuffer *buffer);
	PasteDone          cc.Func // void (* paste_done) (GtkTextBuffer *buffer,GdkClipboard *clipboard);
	Undo               cc.Func // void (* undo) (GtkTextBuffer *buffer);
	Redo               cc.Func // void (* redo) (GtkTextBuffer *buffer);

	_ [4]cc.Func //reserved
}
type TextBuffer struct{ gobjCore }

func GTypeTextBuffer() gobject.GType                { return gtk_text_buffer_get_type.Fn()() }
func NewTextBuffer(table *TextTagTable) *TextBuffer { return gtk_text_buffer_new.Fn()(table) }
func (b *TextBuffer) GetLineCount() int32           { return gtk_text_buffer_get_line_count.Fn()(b) }
func (b *TextBuffer) GetCharCount() int32           { return gtk_text_buffer_get_char_count.Fn()(b) }
func (b *TextBuffer) GetTagTable() *TextTagTable    { return gtk_text_buffer_get_tag_table.Fn()(b) }
func (b *TextBuffer) SetText(text string, length int32) {
	cText := cc.NewString(text)
	defer cText.Free()
	gtk_text_buffer_set_text.Fn()(b, cText, length)
}
func (b *TextBuffer) Insert(iter *TextIter, text string, length int32) {
	cText := cc.NewString(text)
	defer cText.Free()
	gtk_text_buffer_insert.Fn()(b, iter, cText, length)
}
func (b *TextBuffer) InsertAtCursor(text string, length int32) {
	cText := cc.NewString(text)
	defer cText.Free()
	gtk_text_buffer_insert_at_cursor.Fn()(b, cText, length)
}
func (b *TextBuffer) InsertInteractive(iter *TextIter, text string, length int32, defaultEditable bool) bool {
	cText := cc.NewString(text)
	defer cText.Free()
	return gtk_text_buffer_insert_interactive.Fn()(b, iter, cText, length, cbool(defaultEditable)) != 0
}
func (b *TextBuffer) InsertInteractiveAtCursor(text string, length int32, defaultEditable bool) bool {
	cText := cc.NewString(text)
	defer cText.Free()
	return gtk_text_buffer_insert_interactive_at_cursor.Fn()(b, cText, length, cbool(defaultEditable)) != 0
}
func (b *TextBuffer) InsertRange(iter *TextIter, start, end *TextIter) {
	gtk_text_buffer_insert_range.Fn()(b, iter, start, end)
}
func (b *TextBuffer) InsertRangeInteractive(iter *TextIter, start, end *TextIter, defaultEditable bool) bool {
	return gtk_text_buffer_insert_range_interactive.Fn()(b, iter, start, end, cbool(defaultEditable)) != 0
}
func (b *TextBuffer) InsertWithTags(iter *TextIter, text string, length int32, firstTag *TextTag, args ...interface{}) {
	cText := cc.NewString(text)
	defer cText.Free()
	gtk_text_buffer_insert_with_tags.FnVa()(b, iter, cText, length, firstTag, args...)
}
func (b *TextBuffer) InsertWithTagsByName(iter *TextIter, text string, length int32, firstTagName string, args ...interface{}) {
	cText := cc.NewString(text)
	defer cText.Free()
	cTagName := cc.NewString(firstTagName)
	defer cTagName.Free()
	gtk_text_buffer_insert_with_tags_by_name.FnVa()(b, iter, cText, length, cTagName, args...)
}
func (b *TextBuffer) InsertMarkup(iter *TextIter, markup string, length int32) {
	cMarkup := cc.NewString(markup)
	defer cMarkup.Free()
	gtk_text_buffer_insert_markup.Fn()(b, iter, cMarkup, length)
}
func (b *TextBuffer) Delete(start, end *TextIter) { gtk_text_buffer_delete.Fn()(b, start, end) }
func (b *TextBuffer) DeleteInteractive(startIter, endIter *TextIter, defaultEditable bool) bool {
	return gtk_text_buffer_delete_interactive.Fn()(b, startIter, endIter, cbool(defaultEditable)) != 0
}
func (b *TextBuffer) Backspace(iter *TextIter, interactive, defaultEditable bool) bool {
	return gtk_text_buffer_backspace.Fn()(b, iter, cbool(interactive), cbool(defaultEditable)) != 0
}
func (b *TextBuffer) GetText(start, end *TextIter, includeHiddenChars bool) string {
	return gtk_text_buffer_get_text.Fn()(b, start, end, cbool(includeHiddenChars)).TakeString()
}
func (b *TextBuffer) GetSlice(start, end *TextIter, includeHiddenChars bool) string {
	return gtk_text_buffer_get_slice.Fn()(b, start, end, cbool(includeHiddenChars)).TakeString()
}
func (b *TextBuffer) InsertPaintable(iter *TextIter, paintable *gdk.Paintable) {
	gtk_text_buffer_insert_paintable.Fn()(b, iter, paintable)
}
func (b *TextBuffer) InsertChildAnchor(iter *TextIter, anchor *TextChildAnchor) {
	gtk_text_buffer_insert_child_anchor.Fn()(b, iter, anchor)
}
func (b *TextBuffer) CreateChildAnchor(iter *TextIter) *TextChildAnchor {
	return gtk_text_buffer_create_child_anchor.Fn()(b, iter)
}
func (b *TextBuffer) AddMark(mark *TextMark, where *TextIter) {
	gtk_text_buffer_add_mark.Fn()(b, mark, where)
}
func (b *TextBuffer) CreateMark(markName string, where *TextIter, leftGravity bool) *TextMark {
	cMarkName := cc.NewString(markName)
	defer cMarkName.Free()
	return gtk_text_buffer_create_mark.Fn()(b, cMarkName, where, cbool(leftGravity))
}
func (b *TextBuffer) MoveMark(mark *TextMark, where *TextIter) {
	gtk_text_buffer_move_mark.Fn()(b, mark, where)
}
func (b *TextBuffer) DeleteMark(mark *TextMark) {
	gtk_text_buffer_delete_mark.Fn()(b, mark)
}
func (b *TextBuffer) GetMark(name string) *TextMark {
	cName := cc.NewString(name)
	defer cName.Free()
	return gtk_text_buffer_get_mark.Fn()(b, cName)
}
func (b *TextBuffer) MoveMarkByName(name string, where *TextIter) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_text_buffer_move_mark_by_name.Fn()(b, cName, where)
}
func (b *TextBuffer) DeleteMarkByName(name string) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_text_buffer_delete_mark_by_name.Fn()(b, cName)
}
func (b *TextBuffer) GetInsert() *TextMark { return gtk_text_buffer_get_insert.Fn()(b) }
func (b *TextBuffer) GetSelectionBound() *TextMark {
	return gtk_text_buffer_get_selection_bound.Fn()(b)
}
func (b *TextBuffer) PlaceCursor(where *TextIter) { gtk_text_buffer_place_cursor.Fn()(b, where) }
func (b *TextBuffer) SelectRange(ins, bound *TextIter) {
	gtk_text_buffer_select_range.Fn()(b, ins, bound)
}

func (b *TextBuffer) ApplyTag(tag *TextTag, start, end *TextIter) {
	gtk_text_buffer_apply_tag.Fn()(b, tag, start, end)
}
func (b *TextBuffer) RemoveTag(tag *TextTag, start, end *TextIter) {
	gtk_text_buffer_remove_tag.Fn()(b, tag, start, end)
}
func (b *TextBuffer) ApplyTagByName(name string, start, end *TextIter) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_text_buffer_apply_tag_by_name.Fn()(b, cName, start, end)
}
func (b *TextBuffer) RemoveTagByName(name string, start, end *TextIter) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_text_buffer_remove_tag_by_name.Fn()(b, cName, start, end)
}
func (b *TextBuffer) RemoveAllTags(start, end *TextIter) {
	gtk_text_buffer_remove_all_tags.Fn()(b, start, end)
}
func (b *TextBuffer) CreateTag(tagName string, firstPropertyName string, args ...interface{}) *TextTag {
	cTagName := cc.NewString(tagName)
	defer cTagName.Free()
	cPropName := cc.NewString(firstPropertyName)
	defer cPropName.Free()
	return gtk_text_buffer_create_tag.FnVa()(b, cTagName, cPropName, args...)
}
func (b *TextBuffer) GetIterAtLineOffset(iter *TextIter, lineNumber, charOffset int32) bool {
	return gtk_text_buffer_get_iter_at_line_offset.Fn()(b, iter, lineNumber, charOffset) != 0
}
func (b *TextBuffer) GetIterAtLineIndex(iter *TextIter, lineNumber, byteIndex int32) bool {
	return gtk_text_buffer_get_iter_at_line_index.Fn()(b, iter, lineNumber, byteIndex) != 0
}
func (b *TextBuffer) GetIterAtOffset(iter *TextIter, charOffset int32) {
	gtk_text_buffer_get_iter_at_offset.Fn()(b, iter, charOffset)
}
func (b *TextBuffer) GetIterAtLine(iter *TextIter, lineNumber int32) bool {
	return gtk_text_buffer_get_iter_at_line.Fn()(b, iter, lineNumber) != 0
}
func (b *TextBuffer) GetStartIter(iter *TextIter) {
	gtk_text_buffer_get_start_iter.Fn()(b, iter)
}
func (b *TextBuffer) GetEndIter(iter *TextIter) {
	gtk_text_buffer_get_end_iter.Fn()(b, iter)
}
func (b *TextBuffer) GetBounds(start, end *TextIter) {
	gtk_text_buffer_get_bounds.Fn()(b, start, end)
}
func (b *TextBuffer) GetIterAtMark(iter *TextIter, mark *TextMark) {
	gtk_text_buffer_get_iter_at_mark.Fn()(b, iter, mark)
}
func (b *TextBuffer) GetIterAtChildAnchor(iter *TextIter, anchor *TextChildAnchor) {
	gtk_text_buffer_get_iter_at_child_anchor.Fn()(b, iter, anchor)
}
func (b *TextBuffer) GetModified() bool {
	return gtk_text_buffer_get_modified.Fn()(b) != 0
}
func (b *TextBuffer) SetModified(setting bool) {
	gtk_text_buffer_set_modified.Fn()(b, cbool(setting))
}
func (b *TextBuffer) GetHasSelection() bool {
	return gtk_text_buffer_get_has_selection.Fn()(b) != 0
}
func (b *TextBuffer) AddSelectionClipboard(clipboard *gdk.Clipboard) {
	gtk_text_buffer_add_selection_clipboard.Fn()(b, clipboard)
}
func (b *TextBuffer) RemoveSelectionClipboard(clipboard *gdk.Clipboard) {
	gtk_text_buffer_remove_selection_clipboard.Fn()(b, clipboard)
}
func (b *TextBuffer) CutClipboard(clipboard *gdk.Clipboard, defaultEditable bool) {
	gtk_text_buffer_cut_clipboard.Fn()(b, clipboard, cbool(defaultEditable))
}
func (b *TextBuffer) CopyClipboard(clipboard *gdk.Clipboard) {
	gtk_text_buffer_copy_clipboard.Fn()(b, clipboard)
}
func (b *TextBuffer) PasteClipboard(clipboard *gdk.Clipboard, overrideLocation *TextIter, defaultEditable bool) {
	gtk_text_buffer_paste_clipboard.Fn()(b, clipboard, overrideLocation, cbool(defaultEditable))
}
func (b *TextBuffer) GetSelectionBounds(start, end *TextIter) bool {
	return gtk_text_buffer_get_selection_bounds.Fn()(b, start, end) != 0
}
func (b *TextBuffer) DeleteSelection(interactive, defaultEditable bool) bool {
	return gtk_text_buffer_delete_selection.Fn()(b, cbool(interactive), cbool(defaultEditable)) != 0
}
func (b *TextBuffer) GetSelectionContent() *gdk.ContentProvider {
	return gtk_text_buffer_get_selection_content.Fn()(b)
}
func (b *TextBuffer) GetCanUndo() bool {
	return gtk_text_buffer_get_can_undo.Fn()(b) != 0
}
func (b *TextBuffer) GetCanRedo() bool {
	return gtk_text_buffer_get_can_redo.Fn()(b) != 0
}
func (b *TextBuffer) GetEnableUndo() bool {
	return gtk_text_buffer_get_enable_undo.Fn()(b) != 0
}
func (b *TextBuffer) SetEnableUndo(enableUndo bool) {
	gtk_text_buffer_set_enable_undo.Fn()(b, cbool(enableUndo))
}
func (b *TextBuffer) GetMaxUndoLevels() uint32 {
	return gtk_text_buffer_get_max_undo_levels.Fn()(b)
}
func (b *TextBuffer) SetMaxUndoLevels(maxUndoLevels uint32) {
	gtk_text_buffer_set_max_undo_levels.Fn()(b, maxUndoLevels)
}
func (b *TextBuffer) Undo() {
	gtk_text_buffer_undo.Fn()(b)
}
func (b *TextBuffer) Redo() {
	gtk_text_buffer_redo.Fn()(b)
}
func (b *TextBuffer) BeginIrreversibleAction() {
	gtk_text_buffer_begin_irreversible_action.Fn()(b)
}
func (b *TextBuffer) EndIrreversibleAction() {
	gtk_text_buffer_end_irreversible_action.Fn()(b)
}
func (b *TextBuffer) BeginUserAction() {
	gtk_text_buffer_begin_user_action.Fn()(b)
}
func (b *TextBuffer) EndUserAction() {
	gtk_text_buffer_end_user_action.Fn()(b)
}
func (b *TextBuffer) AddCommitNotify(flags uint32,
	commitNotify func(buf *TextBuffer, flag TextBufferNotifyFlags, pos, len uint32)) uint32 {
	var cb, des uptr
	if commitNotify != nil {
		cb = cc.CbkRaw[func(*TextBuffer, TextBufferNotifyFlags, uint32, uint32, uptr)](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 4)
			commitNotify(*(**TextBuffer)(is[0]), *(*TextBufferNotifyFlags)(is[1]), *(*uint32)(is[2]), *(*uint32)(is[3]))
		})
		des = cc.Cbk(func(d uptr) { cc.CbkClose(cb); cc.CbkCloseLate(des) })
	}

	return gtk_text_buffer_add_commit_notify.Fn()(b, flags, cb, nil, des)
}
func (b *TextBuffer) RemoveCommitNotify(handler uint32) {
	gtk_text_buffer_remove_commit_notify.Fn()(b, handler)
}
func (s *TextBuffer) ConnectApplyTag(
	sig func(s *TextBuffer, tag *TextTag, start, end *TextIter)) uint64 {
	return s.SignalConnect("apply-tag", sig, nil)
}
func (s *TextBuffer) ConnectBeginUserAction(sig func(s *TextBuffer)) uint64 {
	return s.SignalConnect("begin-user-action", sig, nil)
}
func (s *TextBuffer) ConnectChanged(sig func(s *TextBuffer)) uint64 {
	return s.SignalConnect("changed", sig, nil)
}
func (s *TextBuffer) ConnectDeleteRange(
	sig func(s *TextBuffer, start, end *TextIter)) uint64 {
	return s.SignalConnect("delete-range", sig, nil)
}
func (s *TextBuffer) ConnectEndUserAction(sig func(s *TextBuffer)) uint64 {
	return s.SignalConnect("end-user-action", sig, nil)
}
func (s *TextBuffer) ConnectInsertChildAnchor(
	sig func(s *TextBuffer, location *TextIter, anchor *TextChildAnchor)) uint64 {
	return s.SignalConnect("insert-child-anchor", sig, nil)
}
func (s *TextBuffer) ConnectInsertPaintable(
	sig func(s *TextBuffer, location *TextIter, paintable *gdk.Paintable)) uint64 {
	return s.SignalConnect("insert-paintable", sig, nil)
}
func (s *TextBuffer) ConnectInsertText(
	sig func(s *TextBuffer, location *TextIter, text cc.String, len int32)) uint64 {
	return s.SignalConnect("insert-text", sig, nil)
}
func (s *TextBuffer) ConnectMarkDeleted(
	sig func(s *TextBuffer, mark *TextMark)) uint64 {
	return s.SignalConnect("mark-deleted", sig, nil)
}
func (s *TextBuffer) ConnectMarkSet(
	sig func(s *TextBuffer, location *TextIter, mark *TextMark)) uint64 {
	return s.SignalConnect("mark-set", sig, nil)
}
func (s *TextBuffer) ConnectModifiedChanged(sig func(s *TextBuffer)) uint64 {
	return s.SignalConnect("modified-changed", sig, nil)
}
func (s *TextBuffer) ConnectPasteDone(
	sig func(s *TextBuffer, clipboard *gdk.Clipboard)) uint64 {
	return s.SignalConnect("paste-done", sig, nil)
}
func (s *TextBuffer) ConnectRedo(sig func(s *TextBuffer)) uint64 {
	return s.SignalConnect("redo", sig, nil)
}
func (s *TextBuffer) ConnectRemoveTag(
	sig func(s *TextBuffer, tag *TextTag, start, end *TextIter)) uint64 {
	return s.SignalConnect("remove-tag", sig, nil)
}
func (s *TextBuffer) ConnectUndo(sig func(s *TextBuffer)) uint64 {
	return s.SignalConnect("undo", sig, nil)
}

// #endregion

// #region TextChildAnchor

type TextChildAnchorObj struct {
	Parent  gobject.GObjectObj
	Segment uptr
}

type TextChildAnchorClass struct {
	Parent gobject.GObjectClass
	_      [4]uptr //reserved
}

type TextChildAnchor struct{ gobjCore }

func GTypeTextChildAnchor() gobject.GType  { return gtk_text_child_anchor_get_type.Fn()() }
func NewTextChildAnchor() *TextChildAnchor { return gtk_text_child_anchor_new.Fn()() }
func NewTextChildAnchorWithReplacement(character string) *TextChildAnchor {
	cChar := cc.NewString(character)
	defer cChar.Free()
	return gtk_text_child_anchor_new_with_replacement.Fn()(cChar)
}
func (a *TextChildAnchor) GetWidgets() []*Widget {
	var outLen uint32
	ptr := gtk_text_child_anchor_get_widgets.Fn()(a, &outLen)
	if ptr == nil || outLen == 0 {
		return nil
	}
	wds := make([]*Widget, outLen)
	lst := (*[1 << 16]*Widget)(unsafe.Pointer(ptr))[:outLen:outLen]
	for i := uint32(0); i < outLen; i++ {
		wds[i] = lst[i]
	}
	cc.Free(uptr(ptr))
	return wds
}
func (a *TextChildAnchor) GetDeleted() bool { return gtk_text_child_anchor_get_deleted.Fn()(a) != 0 }

// #endregion

// #region TextIter

type TextIter struct {
	Dummy1 uptr
	Dummy2 uptr

	Dummy3 int32
	Dummy4 int32
	Dummy5 int32
	Dummy6 int32
	Dummy7 int32
	Dummy8 int32

	Dummy9  uptr
	Dummy10 uptr

	Dummy11 int32
	Dummy12 int32

	Dummy13 int32
	Dummy14 uptr
}

func GTypeTextIter() gobject.GType              { return gtk_text_iter_get_type.Fn()() }
func (i *TextIter) GetBuffer() *TextBuffer      { return gtk_text_iter_get_buffer.Fn()(i) }
func (i *TextIter) Copy() *TextIter             { return gtk_text_iter_copy.Fn()(i) }
func (i *TextIter) Free()                       { gtk_text_iter_free.Fn()(i) }
func (i *TextIter) Assign(other *TextIter)      { gtk_text_iter_assign.Fn()(i, other) }
func (i *TextIter) GetOffset() int32            { return gtk_text_iter_get_offset.Fn()(i) }
func (i *TextIter) GetLine() int32              { return gtk_text_iter_get_line.Fn()(i) }
func (i *TextIter) GetLineOffset() int32        { return gtk_text_iter_get_line_offset.Fn()(i) }
func (i *TextIter) GetLineIndex() int32         { return gtk_text_iter_get_line_index.Fn()(i) }
func (i *TextIter) GetVisibleLineOffset() int32 { return gtk_text_iter_get_visible_line_offset.Fn()(i) }
func (i *TextIter) GetVisibleLineIndex() int32  { return gtk_text_iter_get_visible_line_index.Fn()(i) }
func (i *TextIter) GetChar() uint32             { return gtk_text_iter_get_char.Fn()(i) }
func (start *TextIter) GetSlice(end *TextIter) string {
	return gtk_text_iter_get_slice.Fn()(start, end).TakeString()
}
func (start *TextIter) GetText(end *TextIter) string {
	return gtk_text_iter_get_text.Fn()(start, end).TakeString()
}
func (start *TextIter) GetVisibleSlice(end *TextIter) string {
	return gtk_text_iter_get_visible_slice.Fn()(start, end).TakeString()
}
func (start *TextIter) GetVisibleText(end *TextIter) string {
	return gtk_text_iter_get_visible_text.Fn()(start, end).TakeString()
}
func (i *TextIter) GetPaintable() *gdk.Paintable { return gtk_text_iter_get_paintable.Fn()(i) }
func (i *TextIter) GetMarks() *glib.GSList[TextMark] {
	return (*glib.GSList[TextMark])(gtk_text_iter_get_marks.Fn()(i))
}
func (i *TextIter) GetChildAnchor() *TextChildAnchor { return gtk_text_iter_get_child_anchor.Fn()(i) }
func (i *TextIter) GetToggledTags(toggledOn bool) *glib.GSList[TextTag] {
	return (*glib.GSList[TextTag])(gtk_text_iter_get_toggled_tags.Fn()(i, cbool(toggledOn)))
}
func (i *TextIter) StartsTag(tag *TextTag) bool  { return gtk_text_iter_starts_tag.Fn()(i, tag) != 0 }
func (i *TextIter) EndsTag(tag *TextTag) bool    { return gtk_text_iter_ends_tag.Fn()(i, tag) != 0 }
func (i *TextIter) TogglesTag(tag *TextTag) bool { return gtk_text_iter_toggles_tag.Fn()(i, tag) != 0 }
func (i *TextIter) HasTag(tag *TextTag) bool     { return gtk_text_iter_has_tag.Fn()(i, tag) != 0 }
func (i *TextIter) GetTags() *glib.GSList[TextTag] {
	return (*glib.GSList[TextTag])(gtk_text_iter_get_tags.Fn()(i))
}
func (i *TextIter) Editable(defaultSetting bool) bool {
	return gtk_text_iter_editable.Fn()(i, cbool(defaultSetting)) != 0
}
func (i *TextIter) CanInsert(defaultEditability bool) bool {
	return gtk_text_iter_can_insert.Fn()(i, cbool(defaultEditability)) != 0
}
func (i *TextIter) StartsWord() bool             { return gtk_text_iter_starts_word.Fn()(i) != 0 }
func (i *TextIter) EndsWord() bool               { return gtk_text_iter_ends_word.Fn()(i) != 0 }
func (i *TextIter) InsideWord() bool             { return gtk_text_iter_inside_word.Fn()(i) != 0 }
func (i *TextIter) StartsSentence() bool         { return gtk_text_iter_starts_sentence.Fn()(i) != 0 }
func (i *TextIter) EndsSentence() bool           { return gtk_text_iter_ends_sentence.Fn()(i) != 0 }
func (i *TextIter) InsideSentence() bool         { return gtk_text_iter_inside_sentence.Fn()(i) != 0 }
func (i *TextIter) StartsLine() bool             { return gtk_text_iter_starts_line.Fn()(i) != 0 }
func (i *TextIter) EndsLine() bool               { return gtk_text_iter_ends_line.Fn()(i) != 0 }
func (i *TextIter) IsCursorPosition() bool       { return gtk_text_iter_is_cursor_position.Fn()(i) != 0 }
func (i *TextIter) GetCharsInLine() int32        { return gtk_text_iter_get_chars_in_line.Fn()(i) }
func (i *TextIter) GetBytesInLine() int32        { return gtk_text_iter_get_bytes_in_line.Fn()(i) }
func (i *TextIter) GetLanguage() *pango.Language { return gtk_text_iter_get_language.Fn()(i) }
func (i *TextIter) IsEnd() bool                  { return gtk_text_iter_is_end.Fn()(i) != 0 }
func (i *TextIter) IsStart() bool                { return gtk_text_iter_is_start.Fn()(i) != 0 }
func (i *TextIter) ForwardChar() bool            { return gtk_text_iter_forward_char.Fn()(i) != 0 }
func (i *TextIter) BackwardChar() bool           { return gtk_text_iter_backward_char.Fn()(i) != 0 }
func (i *TextIter) ForwardChars(count int32) bool {
	return gtk_text_iter_forward_chars.Fn()(i, count) != 0
}
func (i *TextIter) BackwardChars(count int32) bool {
	return gtk_text_iter_backward_chars.Fn()(i, count) != 0
}
func (i *TextIter) ForwardLine() bool  { return gtk_text_iter_forward_line.Fn()(i) != 0 }
func (i *TextIter) BackwardLine() bool { return gtk_text_iter_backward_line.Fn()(i) != 0 }
func (i *TextIter) ForwardLines(count int32) bool {
	return gtk_text_iter_forward_lines.Fn()(i, count) != 0
}
func (i *TextIter) BackwardLines(count int32) bool {
	return gtk_text_iter_backward_lines.Fn()(i, count) != 0
}
func (i *TextIter) ForwardWordEnd() bool    { return gtk_text_iter_forward_word_end.Fn()(i) != 0 }
func (i *TextIter) BackwardWordStart() bool { return gtk_text_iter_backward_word_start.Fn()(i) != 0 }
func (i *TextIter) ForwardWordEnds(count int32) bool {
	return gtk_text_iter_forward_word_ends.Fn()(i, count) != 0
}
func (i *TextIter) BackwardWordStarts(count int32) bool {
	return gtk_text_iter_backward_word_starts.Fn()(i, count) != 0
}
func (i *TextIter) ForwardVisibleLine() bool { return gtk_text_iter_forward_visible_line.Fn()(i) != 0 }
func (i *TextIter) BackwardVisibleLine() bool {
	return gtk_text_iter_backward_visible_line.Fn()(i) != 0
}
func (i *TextIter) ForwardVisibleLines(count int32) bool {
	return gtk_text_iter_forward_visible_lines.Fn()(i, count) != 0
}
func (i *TextIter) BackwardVisibleLines(count int32) bool {
	return gtk_text_iter_backward_visible_lines.Fn()(i, count) != 0
}
func (i *TextIter) ForwardVisibleWordEnd() bool {
	return gtk_text_iter_forward_visible_word_end.Fn()(i) != 0
}
func (i *TextIter) BackwardVisibleWordStart() bool {
	return gtk_text_iter_backward_visible_word_start.Fn()(i) != 0
}
func (i *TextIter) ForwardVisibleWordEnds(count int32) bool {
	return gtk_text_iter_forward_visible_word_ends.Fn()(i, count) != 0
}
func (i *TextIter) BackwardVisibleWordStarts(count int32) bool {
	return gtk_text_iter_backward_visible_word_starts.Fn()(i, count) != 0
}
func (i *TextIter) ForwardSentenceEnd() bool { return gtk_text_iter_forward_sentence_end.Fn()(i) != 0 }
func (i *TextIter) BackwardSentenceStart() bool {
	return gtk_text_iter_backward_sentence_start.Fn()(i) != 0
}
func (i *TextIter) ForwardSentenceEnds(count int32) bool {
	return gtk_text_iter_forward_sentence_ends.Fn()(i, count) != 0
}
func (i *TextIter) BackwardSentenceStarts(count int32) bool {
	return gtk_text_iter_backward_sentence_starts.Fn()(i, count) != 0
}
func (i *TextIter) ForwardCursorPosition() bool {
	return gtk_text_iter_forward_cursor_position.Fn()(i) != 0
}
func (i *TextIter) BackwardCursorPosition() bool {
	return gtk_text_iter_backward_cursor_position.Fn()(i) != 0
}
func (i *TextIter) ForwardCursorPositions(count int32) bool {
	return gtk_text_iter_forward_cursor_positions.Fn()(i, count) != 0
}
func (i *TextIter) BackwardCursorPositions(count int32) bool {
	return gtk_text_iter_backward_cursor_positions.Fn()(i, count) != 0
}
func (i *TextIter) ForwardVisibleCursorPosition() bool {
	return gtk_text_iter_forward_visible_cursor_position.Fn()(i) != 0
}
func (i *TextIter) BackwardVisibleCursorPosition() bool {
	return gtk_text_iter_backward_visible_cursor_position.Fn()(i) != 0
}
func (i *TextIter) ForwardVisibleCursorPositions(count int32) bool {
	return gtk_text_iter_forward_visible_cursor_positions.Fn()(i, count) != 0
}
func (i *TextIter) BackwardVisibleCursorPositions(count int32) bool {
	return gtk_text_iter_backward_visible_cursor_positions.Fn()(i, count) != 0
}
func (i *TextIter) SetOffset(charOffset int32)     { gtk_text_iter_set_offset.Fn()(i, charOffset) }
func (i *TextIter) SetLine(lineNumber int32)       { gtk_text_iter_set_line.Fn()(i, lineNumber) }
func (i *TextIter) SetLineOffset(charOnLine int32) { gtk_text_iter_set_line_offset.Fn()(i, charOnLine) }
func (i *TextIter) SetLineIndex(byteOnLine int32)  { gtk_text_iter_set_line_index.Fn()(i, byteOnLine) }
func (i *TextIter) ForwardToEnd()                  { gtk_text_iter_forward_to_end.Fn()(i) }
func (i *TextIter) ForwardToLineEnd() bool         { return gtk_text_iter_forward_to_line_end.Fn()(i) != 0 }
func (i *TextIter) SetVisibleLineOffset(charOnLine int32) {
	gtk_text_iter_set_visible_line_offset.Fn()(i, charOnLine)
}
func (i *TextIter) SetVisibleLineIndex(byteOnLine int32) {
	gtk_text_iter_set_visible_line_index.Fn()(i, byteOnLine)
}
func (i *TextIter) ForwardToTagToggle(tag *TextTag) bool {
	return gtk_text_iter_forward_to_tag_toggle.Fn()(i, tag) != 0
}
func (i *TextIter) BackwardToTagToggle(tag *TextTag) bool {
	return gtk_text_iter_backward_to_tag_toggle.Fn()(i, tag) != 0
}
func (i *TextIter) ForwardFindChar(pred func(ch uint32) bool, limit *TextIter) bool {
	var cb uptr
	if pred != nil {
		cb = cc.CbkRaw[func(ch uint32, _ uptr) int32](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 2)
			if pred(*(*uint32)(is[0])) {
				*(*int32)(out) = 1
			} else {
				*(*int32)(out) = 0
			}
		})
		defer cc.CbkClose(cb)
	}
	return gtk_text_iter_forward_find_char.Fn()(i, cb, nil, limit) != 0
}
func (i *TextIter) BackwardFindChar(pred func(ch uint32) bool, limit *TextIter) bool {
	var cb uptr
	if pred != nil {
		cb = cc.CbkRaw[func(ch uint32, _ uptr) int32](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 2)
			if pred(*(*uint32)(is[0])) {
				*(*int32)(out) = 1
			} else {
				*(*int32)(out) = 0
			}
		})
		defer cc.CbkClose(cb)
	}
	return gtk_text_iter_backward_find_char.Fn()(i, cb, nil, limit) != 0
}
func (i *TextIter) ForwardSearch(str string, flags uint32, matchStart, matchEnd *TextIter, limit *TextIter) bool {
	cStr := cc.NewString(str)
	defer cStr.Free()
	return gtk_text_iter_forward_search.Fn()(i, cStr, flags, matchStart, matchEnd, limit) != 0
}
func (i *TextIter) BackwardSearch(str string, flags uint32, matchStart, matchEnd *TextIter, limit *TextIter) bool {
	cStr := cc.NewString(str)
	defer cStr.Free()
	return gtk_text_iter_backward_search.Fn()(i, cStr, flags, matchStart, matchEnd, limit) != 0
}
func (lhs *TextIter) Equal(rhs *TextIter) bool    { return gtk_text_iter_equal.Fn()(lhs, rhs) != 0 }
func (lhs *TextIter) Compare(rhs *TextIter) int32 { return gtk_text_iter_compare.Fn()(lhs, rhs) }
func (i *TextIter) InRange(start, end *TextIter) bool {
	return gtk_text_iter_in_range.Fn()(i, start, end) != 0
}
func (i *TextIter) Order(second *TextIter) { gtk_text_iter_order.Fn()(i, second) }

// #endregion

// #region TextMark

type TextMarkObj struct {
	Parent  gobject.GObjectObj
	Segment uptr
}
type TextMarkClass struct {
	Parent gobject.GObjectClass
	_      [8]uptr //padding
}

type TextMark struct{ gobjCore }

func GTypeTextMark() gobject.GType { return gtk_text_mark_get_type.Fn()() }
func NewTextMark(name string, leftGravity bool) *TextMark {
	cName := cc.NewString(name)
	defer cName.Free()
	return gtk_text_mark_new.Fn()(cName, cbool(leftGravity))
}
func (m *TextMark) SetVisible(setting bool) { gtk_text_mark_set_visible.Fn()(m, cbool(setting)) }
func (m *TextMark) GetVisible() bool        { return gtk_text_mark_get_visible.Fn()(m) != 0 }
func (m *TextMark) GetName() string         { return gtk_text_mark_get_name.Fn()(m).String() }
func (m *TextMark) GetDeleted() bool        { return gtk_text_mark_get_deleted.Fn()(m) != 0 }
func (m *TextMark) GetBuffer() *TextBuffer  { return gtk_text_mark_get_buffer.Fn()(m) }
func (m *TextMark) GetLeftGravity() bool    { return gtk_text_mark_get_left_gravity.Fn()(m) != 0 }

// #endregion

// #region TextTag

type TextTagObj struct {
	Parent gobject.GObjectObj
	Priv   uptr
}
type TextTagClass struct {
	Parent gobject.GObjectClass
	_      [8]uptr //padding
}

type TextTag struct{ gobjCore }

func GTypeTextTag() gobject.GType { return gtk_text_tag_get_type.Fn()() }
func NewTextTag(name string) *TextTag {
	cName := cc.NewString(name)
	defer cName.Free()
	return gtk_text_tag_new.Fn()(cName)
}
func (t *TextTag) GetPriority() int32         { return gtk_text_tag_get_priority.Fn()(t) }
func (t *TextTag) SetPriority(priority int32) { gtk_text_tag_set_priority.Fn()(t, priority) }
func (t *TextTag) Changed(sizeChanged bool)   { gtk_text_tag_changed.Fn()(t, cbool(sizeChanged)) }

// #endregion

// #region TextTagTable

type TextTagTable struct {
	gobjCore
	Buildable
}

func GTypeTextTagTable() gobject.GType        { return gtk_text_tag_table_get_type.Fn()() }
func NewTextTagTable() *TextTagTable          { return gtk_text_tag_table_new.Fn()() }
func (t *TextTagTable) Add(tag *TextTag) bool { return gtk_text_tag_table_add.Fn()(t, tag) != 0 }
func (t *TextTagTable) Remove(tag *TextTag)   { gtk_text_tag_table_remove.Fn()(t, tag) }
func (t *TextTagTable) Lookup(name string) *TextTag {
	cName := cc.NewString(name)
	defer cName.Free()
	return gtk_text_tag_table_lookup.Fn()(t, cName)
}
func (t *TextTagTable) Foreach(fn func(table *TextTagTable, tag *TextTag)) {
	var cb uptr
	if fn != nil {
		cb = cc.CbkRaw[func(table *TextTagTable, tag *TextTag, _ uptr)](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 3)
			fn(*(**TextTagTable)(is[0]), *(**TextTag)(is[1]))
		})
		defer cc.CbkClose(cb)
	}
	gtk_text_tag_table_foreach.Fn()(t, cb, nil)
}
func (t *TextTagTable) GetSize() int32 { return gtk_text_tag_table_get_size.Fn()(t) }

func (s *TextTagTable) ConnectTagAdded(sig func(s *TextTagTable, tag *TextTag)) uint64 {
	return s.SignalConnect("tag-added", sig, nil)
}
func (s *TextTagTable) ConnectTagChanged(
	sig func(s *TextTagTable, tag *TextTag, sizeChanged bool)) uint64 {
	return s.SignalConnect("tag-changed", sig, nil)
}
func (s *TextTagTable) ConnectTagRemoved(
	sig func(s *TextTagTable, tag *TextTag)) uint64 {
	return s.SignalConnect("tag-removed", sig, nil)
}

// #endregion

// #region TextView

type TextViewObj struct {
	Parent WidgetObj
	Priv   uptr
}

type TextViewClass struct {
	Parent WidgetClass

	MoveCursor       cc.Func // void (* move_cursor) (GtkTextView *text_view, GtkMovementStep step, int count, gboolean extend_selection);
	SetAnchor        cc.Func // void (* set_anchor) (GtkTextView *text_view);
	InsertAtCursor   cc.Func // void (* insert_at_cursor) (GtkTextView *text_view, const char *str);
	DeleteFromCursor cc.Func // void (* delete_from_cursor) (GtkTextView *text_view, GtkDeleteType type, int count);
	Backspace        cc.Func // void (* backspace) (GtkTextView *text_view);
	CutClipboard     cc.Func // void (* cut_clipboard) (GtkTextView *text_view);
	CopyClipboard    cc.Func // void (* copy_clipboard) (GtkTextView *text_view);
	PasteClipboard   cc.Func // void (* paste_clipboard) (GtkTextView *text_view);
	ToggleOverwrite  cc.Func // void (* toggle_overwrite) (GtkTextView *text_view);
	CreateBuffer     cc.Func // GtkTextBuffer * (* create_buffer) (GtkTextView *text_view);
	SnapshotLayer    cc.Func // void (* snapshot_layer) (GtkTextView *text_view, GtkTextViewLayer layer,GtkSnapshot *snapshot);
	ExtendSelection  cc.Func // gboolean (* extend_selection) (GtkTextView *text_view, GtkTextExtendSelection granularity, const GtkTextIter *location, GtkTextIter *start, GtkTextIter *end);
	InsertEmoji      cc.Func // void (* insert_emoji) (GtkTextView *text_view);

	_ [8]uptr //padding

}

type TextView struct {
	gobjCore
	AccessibleText
	Scrollable
}

func GTypeTextView() gobject.GType { return gtk_text_view_get_type.Fn()() }
func NewTextView() *TextView       { return gtk_text_view_new.Fn()() }
func NewTextViewWithBuffer(buffer *TextBuffer) *TextView {
	return gtk_text_view_new_with_buffer.Fn()(buffer)
}
func (v *TextView) SetBuffer(buffer *TextBuffer) { gtk_text_view_set_buffer.Fn()(v, buffer) }
func (v *TextView) GetBuffer() *TextBuffer       { return gtk_text_view_get_buffer.Fn()(v) }
func (v *TextView) ScrollToIter(iter *TextIter, withinMargin float64, useAlign bool, xalign, yalign float64) bool {
	return gtk_text_view_scroll_to_iter.Fn()(v, iter, withinMargin, cbool(useAlign), xalign, yalign) != 0
}
func (v *TextView) ScrollToMark(mark *TextMark, withinMargin float64, useAlign bool, xalign, yalign float64) {
	gtk_text_view_scroll_to_mark.Fn()(v, mark, withinMargin, cbool(useAlign), xalign, yalign)
}
func (v *TextView) ScrollMarkOnscreen(mark *TextMark) {
	gtk_text_view_scroll_mark_onscreen.Fn()(v, mark)
}
func (v *TextView) MoveMarkOnscreen(mark *TextMark) bool {
	return gtk_text_view_move_mark_onscreen.Fn()(v, mark) != 0
}
func (v *TextView) PlaceCursorOnscreen() bool {
	return gtk_text_view_place_cursor_onscreen.Fn()(v) != 0
}
func (v *TextView) GetVisibleRect(visibleRect *gdk.Rectangle) {
	gtk_text_view_get_visible_rect.Fn()(v, visibleRect)
}
func (v *TextView) GetVisibleOffset() (xOffset, yOffset float64) {
	gtk_text_view_get_visible_offset.Fn()(v, &xOffset, &yOffset)
	return
}
func (v *TextView) SetCursorVisible(setting bool) {
	gtk_text_view_set_cursor_visible.Fn()(v, cbool(setting))
}
func (v *TextView) GetCursorVisible() bool { return gtk_text_view_get_cursor_visible.Fn()(v) != 0 }
func (v *TextView) ResetCursorBlink()      { gtk_text_view_reset_cursor_blink.Fn()(v) }
func (v *TextView) GetCursorLocations(iter *TextIter, strong, weak *gdk.Rectangle) {
	gtk_text_view_get_cursor_locations.Fn()(v, iter, strong, weak)
}
func (v *TextView) GetIterLocation(iter *TextIter, location *gdk.Rectangle) {
	gtk_text_view_get_iter_location.Fn()(v, iter, location)
}
func (v *TextView) GetIterAtLocation(iter *TextIter, x, y int32) bool {
	return gtk_text_view_get_iter_at_location.Fn()(v, iter, x, y) != 0
}
func (v *TextView) GetIterAtPosition(iter *TextIter, trailing *int32, x, y int32) bool {
	return gtk_text_view_get_iter_at_position.Fn()(v, iter, trailing, x, y) != 0
}
func (v *TextView) GetLineYRange(iter *TextIter, y, height *int32) {
	gtk_text_view_get_line_yrange.Fn()(v, iter, y, height)
}
func (v *TextView) GetLineAtY(targetIter *TextIter, y int32, lineTop *int32) {
	gtk_text_view_get_line_at_y.Fn()(v, targetIter, y, lineTop)
}
func (v *TextView) BufferToWindowCoords(win int32, bufferX, bufferY int32, windowX, windowY *int32) {
	gtk_text_view_buffer_to_window_coords.Fn()(v, win, bufferX, bufferY, windowX, windowY)
}
func (v *TextView) WindowToBufferCoords(win int32, windowX, windowY int32, bufferX, bufferY *int32) {
	gtk_text_view_window_to_buffer_coords.Fn()(v, win, windowX, windowY, bufferX, bufferY)
}
func (v *TextView) ForwardDisplayLine(iter *TextIter) bool {
	return gtk_text_view_forward_display_line.Fn()(v, iter) != 0
}
func (v *TextView) BackwardDisplayLine(iter *TextIter) bool {
	return gtk_text_view_backward_display_line.Fn()(v, iter) != 0
}
func (v *TextView) ForwardDisplayLineEnd(iter *TextIter) bool {
	return gtk_text_view_forward_display_line_end.Fn()(v, iter) != 0
}
func (v *TextView) BackwardDisplayLineStart(iter *TextIter) bool {
	return gtk_text_view_backward_display_line_start.Fn()(v, iter) != 0
}
func (v *TextView) StartsDisplayLine(iter *TextIter) bool {
	return gtk_text_view_starts_display_line.Fn()(v, iter) != 0
}
func (v *TextView) MoveVisually(iter *TextIter, count int32) bool {
	return gtk_text_view_move_visually.Fn()(v, iter, count) != 0
}
func (v *TextView) IMContextFilterKeypress(event *gdk.Event) bool {
	return gtk_text_view_im_context_filter_keypress.Fn()(v, event) != 0
}
func (v *TextView) ResetIMContext()             { gtk_text_view_reset_im_context.Fn()(v) }
func (v *TextView) GetGutter(win int32) *Widget { return gtk_text_view_get_gutter.Fn()(v, win) }
func (v *TextView) SetGutter(win int32, widget WidgetIface) {
	gtk_text_view_set_gutter.Fn()(v, win, GetWidgetIface(widget))
}
func (v *TextView) AddChildAtAnchor(child WidgetIface, anchor *TextChildAnchor) {
	gtk_text_view_add_child_at_anchor.Fn()(v, GetWidgetIface(child), anchor)
}
func (v *TextView) AddOverlay(child WidgetIface, xpos, ypos int32) {
	gtk_text_view_add_overlay.Fn()(v, GetWidgetIface(child), xpos, ypos)
}
func (v *TextView) MoveOverlay(child WidgetIface, xpos, ypos int32) {
	gtk_text_view_move_overlay.Fn()(v, GetWidgetIface(child), xpos, ypos)
}
func (v *TextView) Remove(child WidgetIface)      { gtk_text_view_remove.Fn()(v, GetWidgetIface(child)) }
func (v *TextView) SetWrapMode(wrapMode WrapMode) { gtk_text_view_set_wrap_mode.Fn()(v, wrapMode) }
func (v *TextView) GetWrapMode() WrapMode         { return gtk_text_view_get_wrap_mode.Fn()(v) }
func (v *TextView) SetEditable(setting bool)      { gtk_text_view_set_editable.Fn()(v, cbool(setting)) }
func (v *TextView) GetEditable() bool             { return gtk_text_view_get_editable.Fn()(v) != 0 }
func (v *TextView) SetOverwrite(overwrite bool) {
	gtk_text_view_set_overwrite.Fn()(v, cbool(overwrite))
}
func (v *TextView) GetOverwrite() bool { return gtk_text_view_get_overwrite.Fn()(v) != 0 }
func (v *TextView) SetAcceptsTab(acceptsTab bool) {
	gtk_text_view_set_accepts_tab.Fn()(v, cbool(acceptsTab))
}
func (v *TextView) GetAcceptsTab() bool { return gtk_text_view_get_accepts_tab.Fn()(v) != 0 }
func (v *TextView) SetPixelsAboveLines(pixelsAboveLines int32) {
	gtk_text_view_set_pixels_above_lines.Fn()(v, pixelsAboveLines)
}
func (v *TextView) GetPixelsAboveLines() int32 { return gtk_text_view_get_pixels_above_lines.Fn()(v) }
func (v *TextView) SetPixelsBelowLines(pixelsBelowLines int32) {
	gtk_text_view_set_pixels_below_lines.Fn()(v, pixelsBelowLines)
}
func (v *TextView) GetPixelsBelowLines() int32 { return gtk_text_view_get_pixels_below_lines.Fn()(v) }
func (v *TextView) SetPixelsInsideWrap(pixelsInsideWrap int32) {
	gtk_text_view_set_pixels_inside_wrap.Fn()(v, pixelsInsideWrap)
}
func (v *TextView) GetPixelsInsideWrap() int32       { return gtk_text_view_get_pixels_inside_wrap.Fn()(v) }
func (v *TextView) SetJustification(j Justification) { gtk_text_view_set_justification.Fn()(v, j) }
func (v *TextView) GetJustification() Justification  { return gtk_text_view_get_justification.Fn()(v) }
func (v *TextView) SetLeftMargin(leftMargin int32)   { gtk_text_view_set_left_margin.Fn()(v, leftMargin) }
func (v *TextView) GetLeftMargin() int32             { return gtk_text_view_get_left_margin.Fn()(v) }
func (v *TextView) SetRightMargin(rightMargin int32) {
	gtk_text_view_set_right_margin.Fn()(v, rightMargin)
}
func (v *TextView) GetRightMargin() int32        { return gtk_text_view_get_right_margin.Fn()(v) }
func (v *TextView) SetTopMargin(topMargin int32) { gtk_text_view_set_top_margin.Fn()(v, topMargin) }
func (v *TextView) GetTopMargin() int32          { return gtk_text_view_get_top_margin.Fn()(v) }
func (v *TextView) SetBottomMargin(bottomMargin int32) {
	gtk_text_view_set_bottom_margin.Fn()(v, bottomMargin)
}
func (v *TextView) GetBottomMargin() int32       { return gtk_text_view_get_bottom_margin.Fn()(v) }
func (v *TextView) SetIndent(indent int32)       { gtk_text_view_set_indent.Fn()(v, indent) }
func (v *TextView) GetIndent() int32             { return gtk_text_view_get_indent.Fn()(v) }
func (v *TextView) SetTabs(tabs *pango.TabArray) { gtk_text_view_set_tabs.Fn()(v, tabs) }
func (v *TextView) GetTabs() *pango.TabArray     { return gtk_text_view_get_tabs.Fn()(v) }
func (v *TextView) SetInputPurpose(purpose InputPurpose) {
	gtk_text_view_set_input_purpose.Fn()(v, purpose)
}
func (v *TextView) GetInputPurpose() InputPurpose  { return gtk_text_view_get_input_purpose.Fn()(v) }
func (v *TextView) SetInputHints(hints InputHints) { gtk_text_view_set_input_hints.Fn()(v, hints) }
func (v *TextView) GetInputHints() InputHints      { return gtk_text_view_get_input_hints.Fn()(v) }
func (v *TextView) SetMonospace(monospace bool) {
	gtk_text_view_set_monospace.Fn()(v, cbool(monospace))
}
func (v *TextView) GetMonospace() bool                 { return gtk_text_view_get_monospace.Fn()(v) != 0 }
func (v *TextView) SetExtraMenu(model *gio.GMenuModel) { gtk_text_view_set_extra_menu.Fn()(v, model) }
func (v *TextView) GetExtraMenu() *gio.GMenuModel      { return gtk_text_view_get_extra_menu.Fn()(v) }
func (v *TextView) GetRtlContext() *pango.Context      { return gtk_text_view_get_rtl_context.Fn()(v) }
func (v *TextView) GetLtrContext() *pango.Context      { return gtk_text_view_get_ltr_context.Fn()(v) }

func (v *TextView) ConnectBackspace(sig func(v *TextView)) uint64 {
	return v.SignalConnect("backspace", sig, nil)
}
func (v *TextView) ConnectCopyClipboard(sig func(v *TextView)) uint64 {
	return v.SignalConnect("copy-clipboard", sig, nil)
}
func (v *TextView) ConnectCutClipboard(sig func(v *TextView)) uint64 {
	return v.SignalConnect("cut-clipboard", sig, nil)
}
func (v *TextView) ConnectDeleteFromCursor(sig func(v *TextView, typ *DeleteType, count int32)) uint64 {
	return v.SignalConnect("delete-from-cursor", sig, nil)
}
func (v *TextView) ConnectExtendSelection(
	sig func(v *TextView, granularity TextExtendSelection, location, start, end *TextIter, _ uptr) bool) uint64 {
	return v.SignalConnect("extend-selection", sig, nil)
}
func (v *TextView) ConnectInsertAtCursor(sig func(v *TextView, str cc.String)) uint64 {
	return v.SignalConnect("insert-at-cursor", sig, nil)
}
func (v *TextView) ConnectInsertEmoji(sig func(v *TextView)) uint64 {
	return v.SignalConnect("insert-emoji", sig, nil)
}
func (v *TextView) ConnectMoveCursor(sig func(v *TextView, step MovementStep, count int32, extendSelection bool)) uint64 {
	return v.SignalConnect("move-cursor", sig, nil)
}
func (v *TextView) ConnectMoveViewport(sig func(v *TextView, step MovementStep, count int32)) uint64 {
	return v.SignalConnect("move-viewport", sig, nil)
}
func (v *TextView) ConnectPasteClipboard(sig func(v *TextView)) uint64 {
	return v.SignalConnect("paste-clipboard", sig, nil)
}
func (v *TextView) ConnectPreeditChanged(sig func(v *TextView, preedit cc.String)) uint64 {
	return v.SignalConnect("preedit-changed", sig, nil)
}
func (v *TextView) ConnectSelectAll(sig func(v *TextView, selected bool)) uint64 {
	return v.SignalConnect("select-all", sig, nil)
}
func (v *TextView) ConnectSetAnchor(sig func(v *TextView)) uint64 {
	return v.SignalConnect("set-anchor", sig, nil)
}
func (v *TextView) ConnectToggleCursorVisible(sig func(v *TextView)) uint64 {
	return v.SignalConnect("toggle-cursor-visible", sig, nil)
}
func (v *TextView) ConnectToggleOverwrite(sig func(v *TextView)) uint64 {
	return v.SignalConnect("toggle-overwrite", sig, nil)
}

// #endregion

// #region ToggleButton

type ToggleButtonObj struct {
	Parent ButtonObj
}

type ToggleButtonClass struct {
	Parent  ButtonClass
	Toggled cc.Func //void (* toggled) (GtkToggleButton *toggle_button);
	_       [8]uptr //padding
}

type ToggleButton struct {
	Button
	Actionable
}

func GTypeToggleButton() gobject.GType { return gtk_toggle_button_get_type.Fn()() }
func NewToggleButton() *ToggleButton   { return gtk_toggle_button_new.Fn()() }
func NewToggleButtonWithLabel(label string) *ToggleButton {
	cLabel := cc.NewString(label)
	defer cLabel.Free()
	return gtk_toggle_button_new_with_label.Fn()(cLabel)
}
func NewToggleButtonWithMnemonic(label string) *ToggleButton {
	cLabel := cc.NewString(label)
	defer cLabel.Free()
	return gtk_toggle_button_new_with_mnemonic.Fn()(cLabel)
}
func (tb *ToggleButton) SetActive(isActive bool) {
	gtk_toggle_button_set_active.Fn()(tb, cbool(isActive))
}
func (tb *ToggleButton) GetActive() bool              { return gtk_toggle_button_get_active.Fn()(tb) != 0 }
func (tb *ToggleButton) Toggled()                     { gtk_toggle_button_toggled.Fn()(tb) }
func (tb *ToggleButton) SetGroup(group *ToggleButton) { gtk_toggle_button_set_group.Fn()(tb, group) }
func (b *ToggleButton) ConnectToggled(sig func(b *ToggleButton)) uint64 {
	return b.SignalConnect("toggled", sig, nil)
}

// #endregion

// #region Tooltip

type Tooltip struct{ gobjCore }

func GTypeTooltip() gobject.GType { return gtk_tooltip_get_type.Fn()() }
func (tt *Tooltip) SetMarkup(markup string) {
	cMarkup := cc.NewString(markup)
	defer cMarkup.Free()
	gtk_tooltip_set_markup.Fn()(tt, cMarkup)
}
func (tt *Tooltip) SetText(text string) {
	cText := cc.NewString(text)
	defer cText.Free()
	gtk_tooltip_set_text.Fn()(tt, cText)
}
func (tt *Tooltip) SetIcon(paintable *gdk.Paintable) {
	gtk_tooltip_set_icon.Fn()(tt, paintable)
}
func (tt *Tooltip) SetIconFromIconName(iconName string) {
	cIconName := cc.NewString(iconName)
	defer cIconName.Free()
	gtk_tooltip_set_icon_from_icon_name.Fn()(tt, cIconName)
}
func (tt *Tooltip) SetIconFromGicon(gicon *gio.GIcon) {
	gtk_tooltip_set_icon_from_gicon.Fn()(tt, gicon)
}
func (tt *Tooltip) SetCustom(customWidget WidgetIface) {
	gtk_tooltip_set_custom.Fn()(tt, GetWidgetIface(customWidget))
}
func (tt *Tooltip) SetTipArea(rect *gdk.Rectangle) {
	gtk_tooltip_set_tip_area.Fn()(tt, rect)
}

// #endregion

// #region TreeExpander

type TreeExpander struct{ Widget }

func GTypeTreeExpander() gobject.GType     { return gtk_tree_expander_get_type.Fn()() }
func NewTreeExpander() *TreeExpander       { return gtk_tree_expander_new.Fn()() }
func (te *TreeExpander) GetChild() *Widget { return gtk_tree_expander_get_child.Fn()(te) }
func (te *TreeExpander) SetChild(child WidgetIface) {
	gtk_tree_expander_set_child.Fn()(te, GetWidgetIface(child))
}
func (te *TreeExpander) GetItem() uptr {
	return gtk_tree_expander_get_item.Fn()(te)
}
func (te *TreeExpander) GetListRow() *TreeListRow {
	return gtk_tree_expander_get_list_row.Fn()(te)
}
func (te *TreeExpander) SetListRow(listRow *TreeListRow) {
	gtk_tree_expander_set_list_row.Fn()(te, listRow)
}
func (te *TreeExpander) GetIndentForDepth() bool {
	return gtk_tree_expander_get_indent_for_depth.Fn()(te) != 0
}
func (te *TreeExpander) SetIndentForDepth(indentForDepth bool) {
	gtk_tree_expander_set_indent_for_depth.Fn()(te, cbool(indentForDepth))
}
func (te *TreeExpander) GetIndentForIcon() bool {
	return gtk_tree_expander_get_indent_for_icon.Fn()(te) != 0
}
func (te *TreeExpander) SetIndentForIcon(indentForIcon bool) {
	gtk_tree_expander_set_indent_for_icon.Fn()(te, cbool(indentForIcon))
}
func (te *TreeExpander) GetHideExpander() bool {
	return gtk_tree_expander_get_hide_expander.Fn()(te) != 0
}
func (te *TreeExpander) SetHideExpander(hideExpander bool) {
	gtk_tree_expander_set_hide_expander.Fn()(te, cbool(hideExpander))
}

// #endregion

// #region TreeListModel

type TreeListModel struct {
	gobjCore
	gio.GListModel
}

type TreeListRow struct{ gobjCore }

func GTypeTreeListModel() gobject.GType { return gtk_tree_list_model_get_type.Fn()() }
func NewTreeListModel[T any](root gio.GListModelIface, passthrough, autoexpand bool,
	createFunc func(item *T) *gio.GListModel) *TreeListModel {
	var cb, des uptr
	if createFunc != nil {
		cb = cc.CbkRaw[func(item, _ uptr) uptr](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 2)
			*(**gio.GListModel)(out) = createFunc(*(**T)(is[0]))
		})
		des = cc.Cbk(func(_ uptr) { cc.CbkClose(cb); cc.CbkCloseLate(des) })
	}

	return gtk_tree_list_model_new.Fn()(gio.GetGListModelIface(root), cbool(passthrough), cbool(autoexpand), cb, nil, des)
}
func (tlm *TreeListModel) GetModel() *gio.GListModel {
	return gtk_tree_list_model_get_model.Fn()(tlm)
}
func (tlm *TreeListModel) GetPassthrough() bool {
	return gtk_tree_list_model_get_passthrough.Fn()(tlm) != 0
}
func (tlm *TreeListModel) SetAutoexpand(autoexpand bool) {
	gtk_tree_list_model_set_autoexpand.Fn()(tlm, cbool(autoexpand))
}
func (tlm *TreeListModel) GetAutoexpand() bool {
	return gtk_tree_list_model_get_autoexpand.Fn()(tlm) != 0
}
func (tlm *TreeListModel) GetChildRow(position uint32) *TreeListRow {
	return gtk_tree_list_model_get_child_row.Fn()(tlm, position)
}
func (tlm *TreeListModel) GetRow(position uint32) *TreeListRow {
	return gtk_tree_list_model_get_row.Fn()(tlm, position)
}

func GTypeTreeListRow() gobject.GType { return gtk_tree_list_row_get_type.Fn()() }
func (row *TreeListRow) GetItem() uptr {
	return gtk_tree_list_row_get_item.Fn()(row)
}
func (row *TreeListRow) SetExpanded(expanded bool) {
	gtk_tree_list_row_set_expanded.Fn()(row, cbool(expanded))
}
func (row *TreeListRow) GetExpanded() bool {
	return gtk_tree_list_row_get_expanded.Fn()(row) != 0
}
func (row *TreeListRow) IsExpandable() bool {
	return gtk_tree_list_row_is_expandable.Fn()(row) != 0
}
func (row *TreeListRow) GetPosition() uint32 {
	return gtk_tree_list_row_get_position.Fn()(row)
}
func (row *TreeListRow) GetDepth() uint32 {
	return gtk_tree_list_row_get_depth.Fn()(row)
}
func (row *TreeListRow) GetChildren() *gio.GListModel {
	return gtk_tree_list_row_get_children.Fn()(row)
}
func (row *TreeListRow) GetParent() *TreeListRow {
	return gtk_tree_list_row_get_parent.Fn()(row)
}
func (row *TreeListRow) GetChildRow(position uint32) *TreeListRow {
	return gtk_tree_list_row_get_child_row.Fn()(row, position)
}

// #endregion

// #region TreeListRowSorter

type TreeListRowSorter struct{ Sorter }

func GTypeTreeListRowSorter() gobject.GType { return gtk_tree_list_row_sorter_get_type.Fn()() }
func NewTreeListRowSorter(sorter SorterIface) *TreeListRowSorter {
	return gtk_tree_list_row_sorter_new.Fn()(GetSorterIface(sorter))
}
func (s *TreeListRowSorter) GetSorter() *Sorter {
	return gtk_tree_list_row_sorter_get_sorter.Fn()(s)
}
func (s *TreeListRowSorter) SetSorter(sorter *Sorter) {
	gtk_tree_list_row_sorter_set_sorter.Fn()(s, sorter)
}

// #endregion

// #region UriLauncher

type UriLauncher struct{ gobjCore }

func GTypeUriLauncher() gobject.GType { return gtk_uri_launcher_get_type.Fn()() }
func NewUriLauncher(uri string) *UriLauncher {
	cUri := cc.NewString(uri)
	defer cUri.Free()
	return gtk_uri_launcher_new.Fn()(cUri)
}
func (ul *UriLauncher) GetUri() string {
	return gtk_uri_launcher_get_uri.Fn()(ul).String()
}
func (ul *UriLauncher) SetUri(uri string) {
	cUri := cc.NewString(uri)
	defer cUri.Free()
	gtk_uri_launcher_set_uri.Fn()(ul, cUri)
}
func (ul *UriLauncher) Launch(parent WindowIface, cancellable *gio.GCancellable,
	callback func(ul *UriLauncher, res *gio.GAsyncResult)) {
	cb := cc.Cbk(callback)
	ul.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	gtk_uri_launcher_launch.Fn()(ul, GetWindowIface(parent), cancellable, cb, nil)
}
func (ul *UriLauncher) LaunchFinish(result *gio.GAsyncResult) error {
	var gerr *glib.GError
	gtk_uri_launcher_launch_finish.Fn()(ul, result, &gerr)
	return gerr.TakeError()
}

// #endregion

// #region Version

func GetMajorVersion() uint32 { return gtk_get_major_version.Fn()() }
func GetMinorVersion() uint32 { return gtk_get_minor_version.Fn()() }
func GetMicroVersion() uint32 { return gtk_get_micro_version.Fn()() }
func GetBinaryAge() uint32    { return gtk_get_binary_age.Fn()() }
func GetInterfaceAge() uint32 { return gtk_get_interface_age.Fn()() }
func CheckVersion(requiredMajor, requiredMinor, requiredMicro uint32) string {
	return gtk_check_version.Fn()(requiredMajor, requiredMinor, requiredMicro).String()
}

// #endregion

// #region Video

type Video struct{ Widget }

func GTypeVideo() gobject.GType { return gtk_video_get_type.Fn()() }
func NewVideo() *Video          { return gtk_video_new.Fn()() }
func NewVideoForMediaStream(stream *MediaStream) *Video {
	return gtk_video_new_for_media_stream.Fn()(stream)
}
func NewVideoForFile(file *gio.GFile) *Video { return gtk_video_new_for_file.Fn()(file) }
func NewVideoForFilename(filename string) *Video {
	cFilename := cc.NewString(filename)
	defer cFilename.Free()
	return gtk_video_new_for_filename.Fn()(cFilename)
}
func NewVideoForResource(resourcePath string) *Video {
	cResource := cc.NewString(resourcePath)
	defer cResource.Free()
	return gtk_video_new_for_resource.Fn()(cResource)
}
func (v *Video) GetMediaStream() *MediaStream       { return gtk_video_get_media_stream.Fn()(v) }
func (v *Video) SetMediaStream(stream *MediaStream) { gtk_video_set_media_stream.Fn()(v, stream) }
func (v *Video) GetFile() *gio.GFile                { return gtk_video_get_file.Fn()(v) }
func (v *Video) SetFile(file *gio.GFile)            { gtk_video_set_file.Fn()(v, file) }
func (v *Video) SetFilename(filename string) {
	cFilename := cc.NewString(filename)
	defer cFilename.Free()
	gtk_video_set_filename.Fn()(v, cFilename)
}
func (v *Video) SetResource(resourcePath string) {
	cResource := cc.NewString(resourcePath)
	defer cResource.Free()
	gtk_video_set_resource.Fn()(v, cResource)
}
func (v *Video) GetAutoplay() bool         { return gtk_video_get_autoplay.Fn()(v) != 0 }
func (v *Video) SetAutoplay(autoplay bool) { gtk_video_set_autoplay.Fn()(v, cbool(autoplay)) }
func (v *Video) GetLoop() bool             { return gtk_video_get_loop.Fn()(v) != 0 }
func (v *Video) SetLoop(loop bool)         { gtk_video_set_loop.Fn()(v, cbool(loop)) }
func (v *Video) GetGraphicsOffload() GraphicsOffloadEnabled {
	return gtk_video_get_graphics_offload.Fn()(v)
}
func (v *Video) SetGraphicsOffload(enabled GraphicsOffloadEnabled) {
	gtk_video_set_graphics_offload.Fn()(v, enabled)
}

// #endregion

// #region Viewport

type Viewport struct {
	Widget
	Scrollable
}

func GTypeViewport() gobject.GType { return gtk_viewport_get_type.Fn()() }
func NewViewport(hadj, vadj *Adjustment) *Viewport {
	return gtk_viewport_new.Fn()(hadj, vadj)
}
func (vp *Viewport) GetScrollToFocus() bool { return gtk_viewport_get_scroll_to_focus.Fn()(vp) != 0 }
func (vp *Viewport) SetScrollToFocus(scrollToFocus bool) {
	gtk_viewport_set_scroll_to_focus.Fn()(vp, cbool(scrollToFocus))
}
func (vp *Viewport) SetChild(child WidgetIface) {
	gtk_viewport_set_child.Fn()(vp, GetWidgetIface(child))
}
func (vp *Viewport) GetChild() *Widget { return gtk_viewport_get_child.Fn()(vp) }
func (vp *Viewport) ScrollTo(descendant WidgetIface, scroll *ScrollInfo) {
	gtk_viewport_scroll_to.Fn()(vp, GetWidgetIface(descendant), scroll)
}

// #endregion

// #region Widget

type WidgetObj struct {
	Parent gobject.GInitiallyUnownedObj
	Priv   *uptr //GtkWidgetPrivate
}
type WidgetClass struct {
	ParentClass gobject.GInitiallyUnownedClass // 对应 GInitiallyUnownedClass

	Show      cc.Func // void (* show) (GtkWidget *widget);
	Hide      cc.Func // void (* hide) (GtkWidget *widget);
	Map       cc.Func // void (* map)  (GtkWidget *widget);
	Unmap     cc.Func // void (* unmap) (GtkWidget *widget);
	Realize   cc.Func // void (* realize) (GtkWidget *widget);
	Unrealize cc.Func // void (* unrealize) (GtkWidget *widget);
	Root      cc.Func // void (* root) (GtkWidget *widget);
	Unroot    cc.Func // void (* unroot) (GtkWidget *widget);

	SizeAllocate         cc.Func // void (* size_allocate) (GtkWidget *widget,int width,int height,int baseline);
	StateFlagsChanged    cc.Func // void (* state_flags_changed) (GtkWidget *widget,GtkStateFlags previous_state_flags);
	DirectionChanged     cc.Func // void (* direction_changed) (GtkWidget *widget,GtkTextDirection previous_direction);
	GetRequestMode       cc.Func // GtkSizeRequestMode (* get_request_mode) (GtkWidget *widget);
	Measure              cc.Func // void (* measure) (GtkWidget *widget, GtkOrientation orientation, int for_size, int *minimum, int *natural, int *minimum_baseline, int *natural_baseline);
	MnemonicActivate     cc.Func // gboolean (* mnemonic_activate) (GtkWidget *widget, gboolean group_cycling);
	GrabFocus            cc.Func // gboolean (* grab_focus) (GtkWidget *widget);
	Focus                cc.Func // gboolean (* focus) (GtkWidget *widget, GtkDirectionType direction);
	SetFocusChild        cc.Func // void (* set_focus_child) (GtkWidget *widget, GtkWidget *child);
	MoveFocus            cc.Func // void (* move_focus) (GtkWidget *widget, GtkDirectionType direction);
	KeynavFailed         cc.Func // gboolean (* keynav_failed) (GtkWidget *widget, GtkDirectionType direction);
	QueryTooltip         cc.Func // gboolean (* query_tooltip) (GtkWidget *widget, int x, int y, gboolean keyboard_tooltip, GtkTooltip *tooltip);
	ComputeExpand        cc.Func // void (* compute_expand) (GtkWidget *widget, gboolean *hexpand_p, gboolean *vexpand_p);
	CssChanged           cc.Func // void (* css_changed) (GtkWidget *widget, GtkCssStyleChange *change);
	SystemSettingChanged cc.Func // void (* system_setting_changed) (GtkWidget *widget, GtkSystemSetting settings);
	Snapshot             cc.Func // void (* snapshot) (GtkWidget *widget, GtkSnapshot *snapshot);
	Contains             cc.Func // gboolean (* contains) (GtkWidget *widget, double x, double y);

	Priv    uptr // 对应 GtkWidgetClassPrivate
	Padding [8]uptr
}
type Widget struct {
	gobjCore

	//implements

	Accessible
	Buildable
	ConstraintTarget
}

func WidgetCast[T WidgetIface](wd *Widget) T { return *(*T)(uptr(&wd)) }

type WidgetIface interface{ GetWidgetIface() *Widget }

func GetWidgetIface(w WidgetIface) *Widget {
	if anyptr(w) == nil {
		return nil
	}
	return w.GetWidgetIface()
}

func (w *Widget) GetWidgetIface() *Widget { return w }

func GTypeWidget() gobject.GType { return gtk_widget_get_type.Fn()() }
func (w *Widget) Unparent()      { gtk_widget_unparent.Fn()(w) }
func (w *Widget) Show()          { gtk_widget_show.Fn()(w) }
func (w *Widget) Hide()          { gtk_widget_hide.Fn()(w) }
func (w *Widget) Map()           { gtk_widget_map.Fn()(w) }
func (w *Widget) Unmap()         { gtk_widget_unmap.Fn()(w) }
func (w *Widget) Realize()       { gtk_widget_realize.Fn()(w) }
func (w *Widget) Unrealize()     { gtk_widget_unrealize.Fn()(w) }

type Allocation gdk.Rectangle

func (w *Widget) QueueDraw()                     { gtk_widget_queue_draw.Fn()(w) }
func (w *Widget) QueueResize()                   { gtk_widget_queue_resize.Fn()(w) }
func (w *Widget) QueueAllocate()                 { gtk_widget_queue_allocate.Fn()(w) }
func (w *Widget) GetFrameClock() *gdk.FrameClock { return gtk_widget_get_frame_clock.Fn()(w) }
func (w *Widget) SizeAllocate(allocation *Allocation, baseline int) {
	gtk_widget_size_allocate.Fn()(w, allocation, baseline)
}
func (w *Widget) Allocate(width, height, baseline int, transform *gsk.Transform) {
	gtk_widget_allocate.Fn()(w, width, height, baseline, transform)
}

type Requisition struct {
	Width  int32
	Height int32
}

func (w *Widget) GetRequestMode() SizeRequestMode {
	return gtk_widget_get_request_mode.Fn()(w)
}
func (w *Widget) Measure(orientation Orientation, for_size int) (minimum, natural, minimum_baseline, natural_baseline int) {
	gtk_widget_measure.Fn()(w, orientation, for_size, &minimum, &natural, &minimum_baseline, &natural_baseline)
	return
}
func (w *Widget) GetPreferredSize() (minimum_size, natural_size Requisition) {
	gtk_widget_get_preferred_size.Fn()(w, &minimum_size, &natural_size)
	return
}
func (w *Widget) SetLayoutManager(layoutManager LayoutManagerIface) {
	gtk_widget_set_layout_manager.Fn()(w, GetLayoutManagerIface(layoutManager))
}
func (w *Widget) GetLayoutManager() *LayoutManager {
	return gtk_widget_get_layout_manager.Fn()(w)
}
func (w *WidgetClass) SetLayoutManagerType(gtype gobject.GType) {
	gtk_widget_class_set_layout_manager_type.Fn()(w, gtype)
}
func (w *WidgetClass) GetLayoutManagerType() gobject.GType {
	return gtk_widget_class_get_layout_manager_type.Fn()(w)
}
func (w *WidgetClass) AddBinding(keyval uint, mods gdk.ModifierType,
	callback func(w *Widget, args *glib.GVariant, data uptr) bool, format string, args ...interface{}) {

	var cb uptr
	if callback != nil {
		cb = cc.CbkRaw[func(w *Widget, a cc.Ptr, d uptr) int32](func(out, ins uptr) {
			ps := cc.Slice((*uptr)(ins), 3)
			if callback(*(**Widget)(ps[0]), *(**glib.GVariant)(ps[1]), *(*uptr)(ps[2])) {
				*(*int32)(out) = 1
			} else {
				*(*int32)(out) = 0
			}
		})
	}

	f := cc.NewString(format)
	defer f.Free()
	gtk_widget_class_add_binding.FnVa()(w, keyval, mods, cb, f, args...)
}
func (w *WidgetClass) AddBindingSignal(keyval uint, mods gdk.ModifierType, signal, format string, args ...interface{}) {
	s, f := cc.NewString(signal), cc.NewString(format)
	defer s.Free()
	defer f.Free()
	gtk_widget_class_add_binding_signal.FnVa()(w, keyval, mods, s, f, args...)
}
func (w *WidgetClass) AddBindingAction(keyval uint, mods gdk.ModifierType, actionName, format string, args ...interface{}) {
	a, f := cc.NewString(actionName), cc.NewString(format)
	defer a.Free()
	defer f.Free()
	gtk_widget_class_add_binding_action.FnVa()(w, keyval, mods, a, f, args...)
}
func (w *WidgetClass) AddShortcut(shortcut *Shortcut) {
	gtk_widget_class_add_shortcut.Fn()(w, shortcut)
}
func (w *WidgetClass) SetActivateSignal(signalId uint) {
	gtk_widget_class_set_activate_signal.Fn()(w, signalId)
}
func (w *WidgetClass) SetActivateSignalFromName(signalName string) {
	cSignalName := cc.NewString(signalName)
	defer cSignalName.Free()
	gtk_widget_class_set_activate_signal_from_name.Fn()(w, cSignalName)
}
func (w *WidgetClass) GetActivateSignal() uint { return gtk_widget_class_get_activate_signal.Fn()(w) }
func (w *Widget) MnemonicActivate(groupCycling bool) bool {
	return gtk_widget_mnemonic_activate.Fn()(w, cbool(groupCycling)) != 0
}
func (w *Widget) Activate() bool              { return gtk_widget_activate.Fn()(w) != 0 }
func (w *Widget) SetCanFocus(canFocus bool)   { gtk_widget_set_can_focus.Fn()(w, cbool(canFocus)) }
func (w *Widget) GetCanFocus() bool           { return gtk_widget_get_can_focus.Fn()(w) != 0 }
func (w *Widget) SetFocusable(focusable bool) { gtk_widget_set_focusable.Fn()(w, cbool(focusable)) }
func (w *Widget) GetFocusable() bool          { return gtk_widget_get_focusable.Fn()(w) != 0 }
func (w *Widget) HasFocus() bool              { return gtk_widget_has_focus.Fn()(w) != 0 }
func (w *Widget) IsFocus() bool               { return gtk_widget_is_focus.Fn()(w) != 0 }
func (w *Widget) HasVisibleFocus() bool       { return gtk_widget_has_visible_focus.Fn()(w) != 0 }
func (w *Widget) GrabFocus() bool             { return gtk_widget_grab_focus.Fn()(w) != 0 }
func (w *Widget) SetFocusOnClick(focusOnClick bool) {
	gtk_widget_set_focus_on_click.Fn()(w, cbool(focusOnClick))
}
func (w *Widget) GetFocusOnClick() bool       { return gtk_widget_get_focus_on_click.Fn()(w) != 0 }
func (w *Widget) SetCanTarget(canTarget bool) { gtk_widget_set_can_target.Fn()(w, cbool(canTarget)) }
func (w *Widget) GetCanTarget() bool          { return gtk_widget_get_can_target.Fn()(w) != 0 }
func (w *Widget) HasDefault() bool            { return gtk_widget_has_default.Fn()(w) != 0 }
func (w *Widget) SetReceivesDefault(receivesDefault bool) {
	gtk_widget_set_receives_default.Fn()(w, cbool(receivesDefault))
}
func (w *Widget) GetReceivesDefault() bool { return gtk_widget_get_receives_default.Fn()(w) != 0 }
func (w *Widget) SetName(name string) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_widget_set_name.Fn()(w, cName)
}
func (w *Widget) GetName() string { return gtk_widget_get_name.Fn()(w).String() }
func (w *Widget) SetStateFlags(flags StateFlags, clear bool) {
	gtk_widget_set_state_flags.Fn()(w, flags, cbool(clear))
}
func (w *Widget) UnsetStateFlags(flags StateFlags) { gtk_widget_unset_state_flags.Fn()(w, flags) }
func (w *Widget) GetStateFlags() StateFlags        { return gtk_widget_get_state_flags.Fn()(w) }
func (w *Widget) SetSensitive(sensitive bool)      { gtk_widget_set_sensitive.Fn()(w, cbool(sensitive)) }
func (w *Widget) GetSensitive() bool               { return gtk_widget_get_sensitive.Fn()(w) != 0 }
func (w *Widget) IsSensitive() bool                { return gtk_widget_is_sensitive.Fn()(w) != 0 }
func (w *Widget) SetVisible(visible bool)          { gtk_widget_set_visible.Fn()(w, cbool(visible)) }
func (w *Widget) GetVisible() bool                 { return gtk_widget_get_visible.Fn()(w) != 0 }
func (w *Widget) IsVisible() bool                  { return gtk_widget_is_visible.Fn()(w) != 0 }
func (w *Widget) IsDrawable() bool                 { return gtk_widget_is_drawable.Fn()(w) != 0 }
func (w *Widget) GetRealized() bool                { return gtk_widget_get_realized.Fn()(w) != 0 }
func (w *Widget) GetMapped() bool                  { return gtk_widget_get_mapped.Fn()(w) != 0 }
func (w *Widget) SetParent(parent WidgetIface)     { gtk_widget_set_parent.Fn()(w, GetWidgetIface(parent)) }
func (w *Widget) GetParent() *Widget               { return gtk_widget_get_parent.Fn()(w) }
func (w *Widget) GetRoot() *Root                   { return gtk_widget_get_root.Fn()(w) }
func (w *Widget) GetNative() *Native               { return gtk_widget_get_native.Fn()(w) }
func (w *Widget) SetChildVisible(visible bool)     { gtk_widget_set_child_visible.Fn()(w, cbool(visible)) }
func (w *Widget) GetChildVisible() bool            { return gtk_widget_get_child_visible.Fn()(w) != 0 }
func (w *Widget) GetAllocatedWidth() int32         { return gtk_widget_get_allocated_width.Fn()(w) }
func (w *Widget) GetAllocatedHeight() int32        { return gtk_widget_get_allocated_height.Fn()(w) }
func (w *Widget) GetAllocatedBaseline() int32      { return gtk_widget_get_allocated_baseline.Fn()(w) }
func (w *Widget) GetAllocation() (alc Allocation)  { gtk_widget_get_allocation.Fn()(w, &alc); return }
func (w *Widget) ComputeTransform(target WidgetIface) (mat *graphene.Matrix, ok bool) {
	mat = graphene.MatrixAlloc()
	runtime.SetFinalizer(mat, func(m *graphene.Matrix) { m.Free() })
	ok = gtk_widget_compute_transform.Fn()(w, GetWidgetIface(target), mat) != 0
	return
}
func (w *Widget) ComputeBounds(target WidgetIface) (outBounds graphene.Rect, ok bool) {
	ok = gtk_widget_compute_bounds.Fn()(w, GetWidgetIface(target), &outBounds) != 0
	return
}
func (w *Widget) ComputePoint(target WidgetIface, point graphene.Point) (outPoint graphene.Point, ok bool) {
	ok = gtk_widget_compute_point.Fn()(w, GetWidgetIface(target), &point, &outPoint) != 0
	return
}
func (w *Widget) GetWidth() int32    { return gtk_widget_get_width.Fn()(w) }
func (w *Widget) GetHeight() int32   { return gtk_widget_get_height.Fn()(w) }
func (w *Widget) GetBaseline() int32 { return gtk_widget_get_baseline.Fn()(w) }
func (w *Widget) GetSize(orientation Orientation) int32 {
	return gtk_widget_get_size.Fn()(w, orientation)
}
func (w *Widget) ChildFocus(direction DirectionType) bool {
	return gtk_widget_child_focus.Fn()(w, direction) != 0
}
func (w *Widget) KeynavFailed(direction DirectionType) bool {
	return gtk_widget_keynav_failed.Fn()(w, direction) != 0
}
func (w *Widget) ErrorBell() { gtk_widget_error_bell.Fn()(w) }
func (w *Widget) SetSizeRequest(width, height int32) {
	gtk_widget_set_size_request.Fn()(w, width, height)
}
func (w *Widget) GetSizeRequest() (width, height int32) {
	gtk_widget_get_size_request.Fn()(w, &width, &height)
	return
}
func (w *Widget) SetOpacity(opacity float64)    { gtk_widget_set_opacity.Fn()(w, opacity) }
func (w *Widget) GetOpacity() float64           { return gtk_widget_get_opacity.Fn()(w) }
func (w *Widget) SetOverflow(overflow Overflow) { gtk_widget_set_overflow.Fn()(w, overflow) }
func (w *Widget) GetOverflow() Overflow         { return gtk_widget_get_overflow.Fn()(w) }
func (w *Widget) GetAncestor(widgetType gobject.GType) *Widget {
	return gtk_widget_get_ancestor.Fn()(w, widgetType)
}
func (w *Widget) GetScaleFactor() int32        { return gtk_widget_get_scale_factor.Fn()(w) }
func (w *Widget) GetDisplay() *gdk.Display     { return gtk_widget_get_display.Fn()(w) }
func (w *Widget) GetSettings() *Settings       { return gtk_widget_get_settings.Fn()(w) }
func (w *Widget) GetClipboard() *gdk.Clipboard { return gtk_widget_get_clipboard.Fn()(w) }
func (w *Widget) GetPrimaryClipboard() *gdk.Clipboard {
	return gtk_widget_get_primary_clipboard.Fn()(w)
}
func (w *Widget) GetHExpand() bool       { return gtk_widget_get_hexpand.Fn()(w) != 0 }
func (w *Widget) SetHExpand(expand bool) { gtk_widget_set_hexpand.Fn()(w, cbool(expand)) }
func (w *Widget) GetHExpandSet() bool    { return gtk_widget_get_hexpand_set.Fn()(w) != 0 }
func (w *Widget) SetHExpandSet(set bool) { gtk_widget_set_hexpand_set.Fn()(w, cbool(set)) }
func (w *Widget) GetVExpand() bool       { return gtk_widget_get_vexpand.Fn()(w) != 0 }
func (w *Widget) SetVExpand(expand bool) { gtk_widget_set_vexpand.Fn()(w, cbool(expand)) }
func (w *Widget) GetVExpandSet() bool    { return gtk_widget_get_vexpand_set.Fn()(w) != 0 }
func (w *Widget) SetVExpandSet(set bool) { gtk_widget_set_vexpand_set.Fn()(w, cbool(set)) }
func (w *Widget) ComputeExpand(orientation Orientation) bool {
	return gtk_widget_compute_expand.Fn()(w, orientation) != 0
}
func (w *Widget) GetHAlign() Align             { return gtk_widget_get_halign.Fn()(w) }
func (w *Widget) SetHAlign(align Align)        { gtk_widget_set_halign.Fn()(w, align) }
func (w *Widget) GetVAlign() Align             { return gtk_widget_get_valign.Fn()(w) }
func (w *Widget) SetVAlign(align Align)        { gtk_widget_set_valign.Fn()(w, align) }
func (w *Widget) GetMarginStart() int32        { return gtk_widget_get_margin_start.Fn()(w) }
func (w *Widget) SetMarginStart(margin int32)  { gtk_widget_set_margin_start.Fn()(w, margin) }
func (w *Widget) GetMarginEnd() int32          { return gtk_widget_get_margin_end.Fn()(w) }
func (w *Widget) SetMarginEnd(margin int32)    { gtk_widget_set_margin_end.Fn()(w, margin) }
func (w *Widget) GetMarginTop() int32          { return gtk_widget_get_margin_top.Fn()(w) }
func (w *Widget) SetMarginTop(margin int32)    { gtk_widget_set_margin_top.Fn()(w, margin) }
func (w *Widget) GetMarginBottom() int32       { return gtk_widget_get_margin_bottom.Fn()(w) }
func (w *Widget) SetMarginBottom(margin int32) { gtk_widget_set_margin_bottom.Fn()(w, margin) }
func (w *Widget) IsAncestor(ancestor WidgetIface) bool {
	return gtk_widget_is_ancestor.Fn()(w, GetWidgetIface(ancestor)) != 0
}
func (w *Widget) TranslateCoordinates(destWidget WidgetIface, srcX, srcY float64) (destX, destY float64, ok bool) {
	ok = gtk_widget_translate_coordinates.Fn()(w, GetWidgetIface(destWidget), srcX, srcY, &destX, &destY) != 0
	return
}
func (w *Widget) Contains(x, y float64) bool { return gtk_widget_contains.Fn()(w, x, y) != 0 }
func (w *Widget) Pick(x, y float64, flags PickFlags) *Widget {
	return gtk_widget_pick.Fn()(w, x, y, flags)
}
func (w *Widget) AddController(controller EventControllerIface) {
	gtk_widget_add_controller.Fn()(w, GetEventControllerIface(controller))
}
func (w *Widget) RemoveController(controller EventControllerIface) {
	gtk_widget_remove_controller.Fn()(w, GetEventControllerIface(controller))
}
func (w *Widget) CreatePangoContext() *pango.Context { return gtk_widget_create_pango_context.Fn()(w) }
func (w *Widget) GetPangoContext() *pango.Context    { return gtk_widget_get_pango_context.Fn()(w) }
func (w *Widget) SetFontOptions(options *cairo.FontOptions) {
	gtk_widget_set_font_options.Fn()(w, options)
}
func (w *Widget) GetFontOptions() *cairo.FontOptions { return gtk_widget_get_font_options.Fn()(w) }
func (w *Widget) CreatePangoLayout(text string) *pango.Layout {
	cText := cc.NewString(text)
	defer cText.Free()
	return gtk_widget_create_pango_layout.Fn()(w, cText)
}

func (w *Widget) SetDirection(dir TextDirection) { gtk_widget_set_direction.Fn()(w, dir) }
func (w *Widget) GetDirection() TextDirection    { return gtk_widget_get_direction.Fn()(w) }
func SetDefaultDirection(dir TextDirection)      { gtk_widget_set_default_direction.Fn()(dir) }
func GetDefaultDirection() TextDirection         { return gtk_widget_get_default_direction.Fn()() }
func (w *Widget) SetCursor(cursor *gdk.Cursor)   { gtk_widget_set_cursor.Fn()(w, cursor) }
func (w *Widget) SetCursorFromName(name string) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_widget_set_cursor_from_name.Fn()(w, cName)
}
func (w *Widget) GetCursor() *gdk.Cursor { return gtk_widget_get_cursor.Fn()(w) }
func (w *Widget) ListMnemonicLabels() *glib.GList[Widget] {
	return gtk_widget_list_mnemonic_labels.Fn()(w)
}
func (w *Widget) AddMnemonicLabel(label WidgetIface) {
	gtk_widget_add_mnemonic_label.Fn()(w, GetWidgetIface(label))
}
func (w *Widget) RemoveMnemonicLabel(label WidgetIface) {
	gtk_widget_remove_mnemonic_label.Fn()(w, GetWidgetIface(label))
}

func (w *Widget) TriggerTooltipQuery() { gtk_widget_trigger_tooltip_query.Fn()(w) }
func (w *Widget) SetTooltipText(text string) {
	cText := cc.NewString(text)
	defer cText.Free()
	gtk_widget_set_tooltip_text.Fn()(w, cText)
}
func (w *Widget) GetTooltipText() string { return gtk_widget_get_tooltip_text.Fn()(w).String() }
func (w *Widget) SetTooltipMarkup(markup string) {
	cMarkup := cc.NewString(markup)
	defer cMarkup.Free()
	gtk_widget_set_tooltip_markup.Fn()(w, cMarkup)
}
func (w *Widget) GetTooltipMarkup() string { return gtk_widget_get_tooltip_markup.Fn()(w).String() }
func (w *Widget) SetHasTooltip(hasTooltip bool) {
	gtk_widget_set_has_tooltip.Fn()(w, cbool(hasTooltip))
}
func (w *Widget) GetHasTooltip() bool     { return gtk_widget_get_has_tooltip.Fn()(w) != 0 }
func GTypeRequisition() gobject.GType     { return gtk_requisition_get_type.Fn()() }
func NewRequisition() *Requisition        { return gtk_requisition_new.Fn()() }
func (r *Requisition) Copy() *Requisition { return gtk_requisition_copy.Fn()(r) }
func (r *Requisition) Free()              { gtk_requisition_free.Fn()(r) }
func (w *Widget) InDestruction() bool     { return gtk_widget_in_destruction.Fn()(w) != 0 }
func (wc *WidgetClass) SetCssName(name string) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_widget_class_set_css_name.Fn()(wc, cName)
}
func (wc *WidgetClass) GetCssName() string { return gtk_widget_class_get_css_name.Fn()(wc).String() }
func (w *Widget) AddTickCallback(callback func(wd *Widget, fc *gdk.FrameClock) (Continue bool)) uint32 {
	var cb, des uptr
	if callback != nil {
		cb = cc.CbkRaw[func(wd *Widget, fc *gdk.FrameClock, _ uptr) int32](func(out, ins uptr) {
			is := unsafe.Slice((*uptr)(ins), 3)
			r := callback(*(**Widget)(is[0]), *(**gdk.FrameClock)(is[1]))
			if r {
				*(*int32)(out) = 1
			} else {
				*(*int32)(out) = 0
			}
		})
		des = cc.Cbk(func(_ uptr) { cc.CbkClose(cb); cc.CbkCloseLate(des) })
	}
	return gtk_widget_add_tick_callback.Fn()(w, cb, nil, des)
}
func (w *Widget) RemoveTickCallback(id uint32) { gtk_widget_remove_tick_callback.Fn()(w, id) }
func (w *Widget) InsertActionGroup(name string, group *gio.GActionGroup) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_widget_insert_action_group.Fn()(w, cName, group)
}
func (w *Widget) ActivateAction(name string, formatString string, args ...interface{}) bool {
	cName, cformatString := cc.NewString(name), cc.NewString(formatString)
	defer cName.Free()
	defer cformatString.Free()
	return gtk_widget_activate_action.FnVa()(w, cName, cformatString, args...) != 0
}
func (w *Widget) ActivateActionVariant(name string, variant *glib.GVariant) bool {
	cName := cc.NewString(name)
	defer cName.Free()
	return gtk_widget_activate_action_variant.Fn()(w, cName, variant) != 0
}

func (w *Widget) ActivateDefault()                  { gtk_widget_activate_default.Fn()(w) }
func (w *Widget) SetFontMap(fontMap *pango.FontMap) { gtk_widget_set_font_map.Fn()(w, fontMap) }
func (w *Widget) GetFontMap() *pango.FontMap        { return gtk_widget_get_font_map.Fn()(w) }
func (w *Widget) GetFirstChild() *Widget            { return gtk_widget_get_first_child.Fn()(w) }
func (w *Widget) GetLastChild() *Widget             { return gtk_widget_get_last_child.Fn()(w) }
func (w *Widget) GetNextSibling() *Widget           { return gtk_widget_get_next_sibling.Fn()(w) }
func (w *Widget) GetPrevSibling() *Widget           { return gtk_widget_get_prev_sibling.Fn()(w) }

func (w *Widget) ObserveChildren() []*Widget {
	return gio.GListModelList[*Widget](gtk_widget_observe_children.Fn()(w))
}
func (w *Widget) ObserveControllers() []*EventController {
	return gio.GListModelList[*EventController](gtk_widget_observe_controllers.Fn()(w))
}

func (w *Widget) InsertAfter(parent, previousSibling WidgetIface) {
	gtk_widget_insert_after.Fn()(w, GetWidgetIface(parent), GetWidgetIface(previousSibling))
}
func (w *Widget) InsertBefore(parent, nextSibling WidgetIface) {
	gtk_widget_insert_before.Fn()(w, GetWidgetIface(parent), GetWidgetIface(nextSibling))
}
func (w *Widget) SetFocusChild(child WidgetIface) {
	gtk_widget_set_focus_child.Fn()(w, GetWidgetIface(child))
}
func (w *Widget) GetFocusChild() *Widget { return gtk_widget_get_focus_child.Fn()(w) }
func (w *Widget) SnapshotChild(child WidgetIface, snapshot *Snapshot) {
	gtk_widget_snapshot_child.Fn()(w, GetWidgetIface(child), snapshot)
}
func (w *Widget) ShouldLayout() bool {
	return gtk_widget_should_layout.Fn()(w) != 0
}
func (w *Widget) GetCssName() string {
	return gtk_widget_get_css_name.Fn()(w).String()
}
func (w *Widget) AddCssClass(cssClass string) {
	ccc := cc.NewString(cssClass)
	defer ccc.Free()
	gtk_widget_add_css_class.Fn()(w, ccc)
}
func (w *Widget) RemoveCssClass(cssClass string) {
	ccc := cc.NewString(cssClass)
	defer ccc.Free()
	gtk_widget_remove_css_class.Fn()(w, ccc)
}
func (w *Widget) HasCssClass(cssClass string) bool {
	ccc := cc.NewString(cssClass)
	defer ccc.Free()
	return gtk_widget_has_css_class.Fn()(w, ccc) != 0
}
func (w *Widget) GetCssClasses() []string {
	classes := gtk_widget_get_css_classes.Fn()(w)
	defer classes.Free()
	return classes.Strings()
}
func (w *Widget) SetCssClasses(classes []string) {
	cls := cc.NewStrings(classes)
	defer cls.Free()
	gtk_widget_set_css_classes.Fn()(w, cls)
}
func (w *Widget) GetColor() (color gdk.RGBA) { gtk_widget_get_color.Fn()(w, &color); return }
func (wc *WidgetClass) InstallAction(actionName, parameterType string,
	activate func(widget *Widget, actionName string, variant *glib.GVariant)) {
	cName, cParameter := cc.NewString(actionName), cc.NewString(parameterType)
	defer cName.Free()
	defer cParameter.Free()

	act := cc.CbkRaw[func(widget *Widget, actionName cc.String, variant *glib.GVariant)](func(out, ins uptr) {
		is := unsafe.Slice((*uptr)(ins), 3)
		activate(*(**Widget)(is[0]), (*(*cc.String)(is[1])).String(), *(**glib.GVariant)(is[2]))
	})

	gtk_widget_class_install_action.Fn()(wc, cName, cParameter, act)
}
func (wc *WidgetClass) InstallPropertyAction(actionName, propertyName string) {
	cAct, cProp := cc.NewString(actionName), cc.NewString(propertyName)
	defer cAct.Free()
	defer cProp.Free()
	gtk_widget_class_install_property_action.Fn()(wc, cAct, cProp)
}
func (wc *WidgetClass) QueryAction(index uint) (owner gobject.GType, actionName string, parameterType *glib.GVariantType, propertyName string, ok bool) {
	var cActionName, cPropertyName cc.String
	ok = gtk_widget_class_query_action.Fn()(wc, index, &owner, &cActionName, &parameterType, &cPropertyName) != 0
	actionName = cActionName.String()
	propertyName = cActionName.String()
	return
}
func (w *Widget) SetActionEnabled(actionName string, enabled bool) {
	cAct := cc.NewString(actionName)
	defer cAct.Free()
	gtk_widget_action_set_enabled.Fn()(w, cAct, cbool(enabled))
}
func (wc *WidgetClass) SetAccessibleRole(accessibleRole AccessibleRole) {
	gtk_widget_class_set_accessible_role.Fn()(wc, accessibleRole)
}
func (wc *WidgetClass) GetAccessibleRole() AccessibleRole {
	return gtk_widget_class_get_accessible_role.Fn()(wc)
}

func (w *Widget) ConnectDestroy(sig func(w *Widget)) uint64 {
	return w.SignalConnect("destroy", sig, nil)
}
func (w *Widget) ConnectDirectionChanged(sig func(w *Widget, td TextDirection)) uint64 {
	return w.SignalConnect("direction-changed", sig, nil)
}
func (w *Widget) ConnectHide(sig func(w *Widget, td TextDirection)) uint64 {
	return w.SignalConnect("hide", sig, nil)
}
func (w *Widget) ConnectKeynavFailed(sig func(w *Widget, dir DirectionType, _ uptr) bool) uint64 {
	return w.SignalConnect("keynav-failed", sig, nil)
}
func (w *Widget) ConnectMap(sig func(w *Widget)) uint64 {
	return w.SignalConnect("map", sig, nil)
}
func (w *Widget) ConnectMnemonicActivate(sig func(w *Widget, groupCycling bool, _ uptr) bool) uint64 {
	return w.SignalConnect("mnemonic-activate", sig, nil)
}
func (w *Widget) ConnectMoveFocus(sig func(w *Widget, dir DirectionType)) uint64 {
	return w.SignalConnect("move-focus", sig, nil)
}
func (w *Widget) ConnectQueryTooltip(
	sig func(w *Widget, x, y int32, keyboardMode bool, tooltip *Tooltip, _ uptr) bool) uint64 {
	return w.SignalConnect("query-tooltip", sig, nil)
}
func (w *Widget) ConnectRealize(sig func(w *Widget)) uint64 {
	return w.SignalConnect("realize", sig, nil)
}
func (w *Widget) ConnectShow(sig func(w *Widget)) uint64 {
	return w.SignalConnect("show", sig, nil)
}
func (w *Widget) ConnectStateFlagsChanged(sig func(w *Widget, flag StateFlags)) uint64 {
	return w.SignalConnect("state-flags-changed", sig, nil)
}
func (w *Widget) ConnectUnrealize(sig func(w *Widget)) uint64 {
	return w.SignalConnect("unrealize", sig, nil)
}

// #endregion

// #region WidgetPaintable

type WidgetPaintable struct {
	gobjCore
	gdk.Paintable
}

func GTypeWidgetPaintable() gobject.GType { return gtk_widget_paintable_get_type.Fn()() }
func NewWidgetPaintable(widget WidgetIface) *WidgetPaintable {
	return gtk_widget_paintable_new.Fn()(GetWidgetIface(widget))
}
func (wp *WidgetPaintable) GetWidget() *Widget { return gtk_widget_paintable_get_widget.Fn()(wp) }
func (wp *WidgetPaintable) SetWidget(widget WidgetIface) {
	gtk_widget_paintable_set_widget.Fn()(wp, GetWidgetIface(widget))
}

// #endregion

// #region Window

type WindowClass struct {
	Parent WidgetClass

	ActivateFocus   cc.Func // void (* activate_focus) (GtkWindow *window);
	ActivateDefault cc.Func // void (* activate_default) (GtkWindow *window);
	KeysChanged     cc.Func // void	 (* keys_changed) (GtkWindow *window);
	EnableDebugging cc.Func // gboolean (* enable_debugging) (GtkWindow *window,gboolean toggle);
	CloseRequest    cc.Func // gboolean (* close_request) (GtkWindow *window);
	_               [8]uptr // padding
}
type WindowObj struct{ Parent WidgetObj }
type WindowIface interface {
	GetWindowIface() *Window
}

func GetWindowIface(w WindowIface) *Window {
	if anyptr(w) == nil {
		return nil
	}
	return w.GetWindowIface()
}

func (w *Window) GetWindowIface() *Window { return w }

type Window struct {
	Widget

	Root
	ShortcutManager
}

func GTypeWindow() gobject.GType { return gtk_window_get_type.Fn()() }
func NewWindow() *Window         { return gtk_window_new.Fn()() }
func (w *Window) SetTitle(title string) {
	cTitle := cc.NewString(title)
	defer cTitle.Free()
	gtk_window_set_title.Fn()(w, cTitle)
}
func (w *Window) GetTitle() string { return gtk_window_get_title.Fn()(w).String() }
func (w *Window) SetStartupID(startupID string) {
	cStartupID := cc.NewString(startupID)
	defer cStartupID.Free()
	gtk_window_set_startup_id.Fn()(w, cStartupID)
}
func (w *Window) SetFocus(focus WidgetIface) { gtk_window_set_focus.Fn()(w, GetWidgetIface(focus)) }
func (w *Window) GetFocus() *Widget          { return gtk_window_get_focus.Fn()(w) }
func (w *Window) SetDefaultWidget(defaultWidget WidgetIface) {
	gtk_window_set_default_widget.Fn()(w, GetWidgetIface(defaultWidget))
}
func (w *Window) GetDefaultWidget() *Widget { return gtk_window_get_default_widget.Fn()(w) }
func (w *Window) SetTransientFor(parent WindowIface) {
	gtk_window_set_transient_for.Fn()(w, GetWindowIface(parent))
}
func (w *Window) GetTransientFor() *Window { return gtk_window_get_transient_for.Fn()(w) }
func (w *Window) SetDestroyWithParent(setting bool) {
	gtk_window_set_destroy_with_parent.Fn()(w, cbool(setting))
}
func (w *Window) GetDestroyWithParent() bool  { return gtk_window_get_destroy_with_parent.Fn()(w) != 0 }
func (w *Window) SetHideOnClose(setting bool) { gtk_window_set_hide_on_close.Fn()(w, cbool(setting)) }
func (w *Window) GetHideOnClose() bool        { return gtk_window_get_hide_on_close.Fn()(w) != 0 }
func (w *Window) SetMnemonicsVisible(setting bool) {
	gtk_window_set_mnemonics_visible.Fn()(w, cbool(setting))
}
func (w *Window) GetMnemonicsVisible() bool       { return gtk_window_get_mnemonics_visible.Fn()(w) != 0 }
func (w *Window) SetFocusVisible(setting bool)    { gtk_window_set_focus_visible.Fn()(w, cbool(setting)) }
func (w *Window) GetFocusVisible() bool           { return gtk_window_get_focus_visible.Fn()(w) != 0 }
func (w *Window) SetResizable(resizable bool)     { gtk_window_set_resizable.Fn()(w, cbool(resizable)) }
func (w *Window) GetResizable() bool              { return gtk_window_get_resizable.Fn()(w) != 0 }
func (w *Window) SetDisplay(display *gdk.Display) { gtk_window_set_display.Fn()(w, display) }
func (w *Window) IsActive() bool                  { return gtk_window_is_active.Fn()(w) != 0 }
func (w *Window) SetDecorated(setting bool)       { gtk_window_set_decorated.Fn()(w, cbool(setting)) }
func (w *Window) GetDecorated() bool              { return gtk_window_get_decorated.Fn()(w) != 0 }
func (w *Window) SetDeletable(setting bool)       { gtk_window_set_deletable.Fn()(w, cbool(setting)) }
func (w *Window) GetDeletable() bool              { return gtk_window_get_deletable.Fn()(w) != 0 }
func (w *Window) SetIconName(name string) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_window_set_icon_name.Fn()(w, cName)
}
func (w *Window) GetIconName() string { return gtk_window_get_icon_name.Fn()(w).String() }
func SetDefaultIconName(name string) {
	cName := cc.NewString(name)
	defer cName.Free()
	gtk_window_set_default_icon_name.Fn()(cName)
}
func GetDefaultIconName() string { return gtk_window_get_default_icon_name.Fn()().String() }
func SetAutoStartupNotification(setting bool) {
	gtk_window_set_auto_startup_notification.Fn()(cbool(setting))
}
func (w *Window) SetModal(modal bool) { gtk_window_set_modal.Fn()(w, cbool(modal)) }
func (w *Window) GetModal() bool      { return gtk_window_get_modal.Fn()(w) != 0 }
func GetToplevels() []*Window {
	// lst := gtk_window_get_toplevels()
	// num := lst.GetNItems()
	// if num == 0 {
	// 	return nil
	// }
	// ws := make([]*Window, num)
	// for i := uint32(0); i < num; i++ {
	// 	ws[i] = (*Window)(lst.GetItem(i))
	// }
	return gio.GListModelList[*Window](gtk_window_get_toplevels.Fn()())
}
func ListToplevels() *glib.GList[Widget] {
	return gtk_window_list_toplevels.Fn()()
}
func (w *Window) Present()      { gtk_window_present.Fn()(w) }
func (w *Window) Minimize()     { gtk_window_minimize.Fn()(w) }
func (w *Window) Unminimize()   { gtk_window_unminimize.Fn()(w) }
func (w *Window) Maximize()     { gtk_window_maximize.Fn()(w) }
func (w *Window) Unmaximize()   { gtk_window_unmaximize.Fn()(w) }
func (w *Window) Fullscreen()   { gtk_window_fullscreen.Fn()(w) }
func (w *Window) Unfullscreen() { gtk_window_unfullscreen.Fn()(w) }
func (w *Window) FullscreenOnMonitor(monitor *gdk.Monitor) {
	gtk_window_fullscreen_on_monitor.Fn()(w, monitor)
}
func (w *Window) Close() { gtk_window_close.Fn()(w) }
func (w *Window) SetDefaultSize(width, height int32) {
	gtk_window_set_default_size.Fn()(w, width, height)
}
func (w *Window) GetDefaultSize() (width, height int32) {
	gtk_window_get_default_size.Fn()(w, &width, &height)
	return
}
func (w *Window) GetGroup() *WindowGroup       { return gtk_window_get_group.Fn()(w) }
func (w *Window) HasGroup() bool               { return gtk_window_has_group.Fn()(w) != 0 }
func (w *Window) GetApplication() *Application { return gtk_window_get_application.Fn()(w) }
func (w *Window) SetApplication(application *Application) {
	gtk_window_set_application.Fn()(w, application)
}

func (w *Window) SetChild(child WidgetIface) { gtk_window_set_child.Fn()(w, GetWidgetIface(child)) }
func (w *Window) GetChild() *Widget          { return gtk_window_get_child.Fn()(w) }
func (w *Window) SetTitlebar(titlebar WidgetIface) {
	gtk_window_set_titlebar.Fn()(w, GetWidgetIface(titlebar))
}
func (w *Window) GetTitlebar() *Widget { return gtk_window_get_titlebar.Fn()(w) }
func (w *Window) IsMaximized() bool    { return gtk_window_is_maximized.Fn()(w) != 0 }
func (w *Window) IsFullscreen() bool   { return gtk_window_is_fullscreen.Fn()(w) != 0 }
func (w *Window) IsSuspended() bool    { return gtk_window_is_suspended.Fn()(w) != 0 }
func (w *Window) Destroy()             { gtk_window_destroy.Fn()(w) }

func SetInteractiveDebugging(enable bool) { gtk_window_set_interactive_debugging.Fn()(cbool(enable)) }
func (w *Window) SetHandleMenubarAccel(handleMenubarAccel bool) {
	gtk_window_set_handle_menubar_accel.Fn()(w, cbool(handleMenubarAccel))
}
func (w *Window) GetHandleMenubarAccel() bool {
	return gtk_window_get_handle_menubar_accel.Fn()(w) != 0
}
func (w *Window) ConnectActivateDefault(sig func(w *Window)) uint64 {
	return w.SignalConnect("activate-default", sig, nil)
}
func (w *Window) ConnectCloseRequest(sig func(w *Window, _ uptr) (stopClose bool)) uint64 {
	return w.SignalConnect("close-request", sig, nil)
}
func (w *Window) ConnectEnableDebugging(sig func(w *Window, toggle bool, _ uptr) bool) uint64 {
	return w.SignalConnect("enable-debugging", sig, nil)
}

// func (w *Window) ConnectEnableDebugging(sig func(w *Window, toggle bool, _ uptr) bool) uint64 {
// 	return w.SignalConnect("enable-debugging", sig, nil)
// }

// #endregion

// #region WindowControls

type WindowControls struct{ Widget }

func GTypeWindowControls() gobject.GType              { return gtk_window_controls_get_type.Fn()() }
func NewWindowControls(side PackType) *WindowControls { return gtk_window_controls_new.Fn()(side) }
func (wc *WindowControls) GetSide() PackType          { return gtk_window_controls_get_side.Fn()(wc) }
func (wc *WindowControls) SetSide(side PackType)      { gtk_window_controls_set_side.Fn()(wc, side) }
func (wc *WindowControls) GetDecorationLayout() string {
	return gtk_window_controls_get_decoration_layout.Fn()(wc).String()
}
func (wc *WindowControls) SetDecorationLayout(layout string) {
	cLayout := cc.NewString(layout)
	defer cLayout.Free()
	gtk_window_controls_set_decoration_layout.Fn()(wc, cLayout)
}
func (wc *WindowControls) GetUseNativeControls() bool {
	return gtk_window_controls_get_use_native_controls.Fn()(wc) != 0
}
func (wc *WindowControls) SetUseNativeControls(setting bool) {
	gtk_window_controls_set_use_native_controls.Fn()(wc, cbool(setting))
}
func (wc *WindowControls) GetEmpty() bool { return gtk_window_controls_get_empty.Fn()(wc) != 0 }

// #endregion

// #region WindowGroup

type WindowGroupObj struct {
	Parent gobject.GObjectObj
}
type WindowGroupClass struct {
	Parent gobject.GObjectClass
	_      [4]uptr // padding
}

type WindowGroup struct{ gobjCore }

func GTypeWindowGroup() gobject.GType { return gtk_window_group_get_type.Fn()() }
func NewWindowGroup() *WindowGroup    { return gtk_window_group_new.Fn()() }
func (wg *WindowGroup) AddWindow(window WindowIface) {
	gtk_window_group_add_window.Fn()(wg, GetWindowIface(window))
}
func (wg *WindowGroup) RemoveWindow(window WindowIface) {
	gtk_window_group_remove_window.Fn()(wg, GetWindowIface(window))
}
func (wg *WindowGroup) ListWindows() *glib.GList[Window] {
	return gtk_window_group_list_windows.Fn()(wg)
}

// #endregion

// #region WindowHandle

type WindowHandle struct{ Widget }

func GTypeWindowHandle() gobject.GType     { return gtk_window_handle_get_type.Fn()() }
func NewWindowHandle() *WindowHandle       { return gtk_window_handle_new.Fn()() }
func (wh *WindowHandle) GetChild() *Widget { return gtk_window_handle_get_child.Fn()(wh) }
func (wh *WindowHandle) SetChild(child WidgetIface) {
	gtk_window_handle_set_child.Fn()(wh, GetWidgetIface(child))
}

// #endregion
