package gio

import (
	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/glib/glib"
	"github.com/jinzhongmin/gg/glib/gobject"
)

// #region Action

type GActionIfaceObj struct {
	GIface gobject.GTypeInterface

	GetName          uptr
	GetParameterType uptr
	GetStateType     uptr
	GetStateHint     uptr
	GetEnabled       uptr
	GetState         uptr
	ChangeState      uptr
	Activate         uptr
}
type GAction struct{}

func GTypeGAction() gobject.GType { return g_action_get_type.Fn()() }

func (action *GAction) GetName() string {
	return g_action_get_name.Fn()(action)
}
func (action *GAction) GetParameterType() *glib.GVariantType {
	return g_action_get_parameter_type.Fn()(action)
}
func (action *GAction) GetStateType() *glib.GVariantType {
	return g_action_get_state_type.Fn()(action)
}
func (action *GAction) GetStateHint() *glib.GVariant {
	return g_action_get_state_hint.Fn()(action)
}
func (action *GAction) GetEnabled() bool {
	return g_action_get_enabled.Fn()(action)
}
func (action *GAction) GetState() *glib.GVariant {
	return g_action_get_state.Fn()(action)
}
func (action *GAction) ChangeState(value *glib.GVariant) {
	g_action_change_state.Fn()(action, value)
}
func (action *GAction) Activate(parameter *glib.GVariant) {
	g_action_activate.Fn()(action, parameter)
}
func ActionNameIsValid(actionName string) bool {
	return g_action_name_is_valid.Fn()(actionName)
}

// #endregion

// #region ActionGroup
type GActionGroupIfaceObj struct {
	GIface gobject.GTypeInterface

	HasAction              uptr
	ListActions            uptr
	GetActionEnabled       uptr
	GetActionParameterType uptr
	GetActionStateType     uptr
	GetActionStateHint     uptr
	GetActionState         uptr
	ChangeActionState      uptr
	ActivateAction         uptr

	ActionAdded          uptr
	ActionRemoved        uptr
	ActionEnabledChanged uptr
	ActionStateChanged   uptr

	QueryAction uptr
}
type GActionGroup struct{}

func GTypeGActionGroup() gobject.GType {
	return g_action_group_get_type.Fn()()
}
func (group *GActionGroup) HasAction(actionName string) bool {
	return g_action_group_has_action.Fn()(group, actionName)
}
func (group *GActionGroup) ListActions() []string {
	cActions := g_action_group_list_actions.Fn()(group)
	return cActions.Strings()
}
func (group *GActionGroup) GetActionParameterType(actionName string) *glib.GVariantType {
	return g_action_group_get_action_parameter_type.Fn()(group, actionName)
}
func (group *GActionGroup) GetActionStateType(actionName string) *glib.GVariantType {
	return g_action_group_get_action_state_type.Fn()(group, actionName)
}
func (group *GActionGroup) GetActionStateHint(actionName string) *glib.GVariant {
	return g_action_group_get_action_state_hint.Fn()(group, actionName)
}

// #endregion

// #region ActionMap

type GActionMap struct{}

// #endregion

// #region AppInfo

type GAppInfo struct{}

// #endregion

// #region AppLaunchContext

type GAppLaunchContext struct{ gobject.GObject }

// #endregion

// #region Application

type GApplicationObj struct {
	Parent gobject.GObjectObj
	Priv   uptr
}
type GApplicationClass struct {
	ParentClass gobject.GObjectClass

	Startup            uptr
	Activate           uptr
	Open               uptr
	CommandLine        uptr
	LocalCommandLine   uptr
	BeforeEmit         uptr
	AfterEmit          uptr
	AddPlatformData    uptr
	QuitMainloop       uptr
	RunMainloop        uptr
	Shutdown           uptr
	DbusRegister       uptr
	DbusUnregister     uptr
	HandleLocalOptions uptr
	NameLost           uptr
	Padding            [7]uptr
}
type GApplication struct{ gobject.GObject }

func GTypeGApplication() gobject.GType {
	return g_application_get_type.Fn()()
}
func GApplicationIDisValid(id string) bool {
	return g_application_id_is_valid.Fn()(id)
}
func NewGApplication(appID string, flag GApplicationFlags) *GApplication {
	return g_application_new.Fn()(appID, flag)
}
func (app *GApplication) GetApplicationID() string {
	return g_application_get_application_id.Fn()(app)
}
func (app *GApplication) SetApplicationID(applicationID string) {
	g_application_set_application_id.Fn()(app, applicationID)
}

func (app *GApplication) GetVersion() string {
	return g_application_get_version.Fn()(app)
}

func (app *GApplication) SetVersion(version string) {
	g_application_set_version.Fn()(app, version)
}

func (app *GApplication) GetDBusConnection() uptr {
	return g_application_get_dbus_connection.Fn()(app)
}

func (app *GApplication) GetDBusObjectPath() string {
	return g_application_get_dbus_object_path.Fn()(app)
}

func (app *GApplication) GetInactivityTimeout() uint32 {
	return g_application_get_inactivity_timeout.Fn()(app)
}

func (app *GApplication) SetInactivityTimeout(timeout uint32) {
	g_application_set_inactivity_timeout.Fn()(app, timeout)
}

func (app *GApplication) GetFlags() uint32 {
	return g_application_get_flags.Fn()(app)
}

func (app *GApplication) SetFlags(flags uint32) {
	g_application_set_flags.Fn()(app, flags)
}

func (app *GApplication) GetResourceBasePath() string {
	return g_application_get_resource_base_path.Fn()(app)
}

func (app *GApplication) SetResourceBasePath(resourcePath string) {
	g_application_set_resource_base_path.Fn()(app, resourcePath)
}

func (app *GApplication) SetActionGroup(actionGroup uptr) {
	g_application_set_action_group.Fn()(app, actionGroup)
}

func (app *GApplication) AddMainOptionEntries(entries uptr) {
	g_application_add_main_option_entries.Fn()(app, entries)
}

func (app *GApplication) AddMainOption(longName string, shortName byte, flags uint32, arg uint32, description string, argDescription string) {
	g_application_add_main_option.Fn()(app, longName, shortName, flags, arg, description, argDescription)
}

func (app *GApplication) AddOptionGroup(group uptr) {
	g_application_add_option_group.Fn()(app, group)
}

func (app *GApplication) SetOptionContextParameterString(parameterString string) {
	g_application_set_option_context_parameter_string.Fn()(app, parameterString)
}

func (app *GApplication) SetOptionContextSummary(summary string) {
	g_application_set_option_context_summary.Fn()(app, summary)
}

func (app *GApplication) SetOptionContextDescription(description string) {
	g_application_set_option_context_description.Fn()(app, description)
}

func (app *GApplication) GetIsRegistered() bool {
	return g_application_get_is_registered.Fn()(app)
}

func (app *GApplication) GetIsRemote() bool {
	return g_application_get_is_remote.Fn()(app)
}

func (app *GApplication) Register(cancellable uptr, error **uptr) bool {
	return g_application_register.Fn()(app, cancellable, error)
}

func (app *GApplication) Hold() {
	g_application_hold.Fn()(app)
}

func (app *GApplication) Release() {
	g_application_release.Fn()(app)
}

func (app *GApplication) Activate() {
	g_application_activate.Fn()(app)
}

func (app *GApplication) Open(files **uptr, nFiles int32, hint string) {
	g_application_open.Fn()(app, files, nFiles, hint)
}

func (app *GApplication) Run(args []string) int32 {
	argc := int32(len(args))
	argv := cc.NewStrings(args)
	return g_application_run.Fn()(app, argc, argv)
}

func (app *GApplication) Quit() {
	g_application_quit.Fn()(app)
}

func (app *GApplication) MarkBusy() {
	g_application_mark_busy.Fn()(app)
}

func (app *GApplication) UnmarkBusy() {
	g_application_unmark_busy.Fn()(app)
}

func (app *GApplication) GetIsBusy() bool {
	return g_application_get_is_busy.Fn()(app)
}

func (app *GApplication) SendNotification(id string, notification uptr) {
	g_application_send_notification.Fn()(app, id, notification)
}

func (app *GApplication) WithdrawNotification(id string) {
	g_application_withdraw_notification.Fn()(app, id)
}

func (app *GApplication) BindBusyProperty(object uptr, property string) {
	g_application_bind_busy_property.Fn()(app, object, property)
}

func (app *GApplication) UnbindBusyProperty(object uptr, property string) {
	g_application_unbind_busy_property.Fn()(app, object, property)
}
func (app *GApplication) ConnectActivate(sig func(a *GApplication)) uint64 {
	return app.SignalConnect("activate", sig, nil)
}
func (app *GApplication) ConnectCommandLine(
	sig func(a *GApplication, commandLine, _ uptr) int32) uint64 {
	return app.SignalConnect("command-line", sig, nil)
}
func (app *GApplication) ConnectHandleLocalOptions(
	sig func(a *GApplication, optionsDict, _ uptr) int32) uint64 {
	return app.SignalConnect("handle-local-options", sig, nil)
}
func (app *GApplication) ConnectNameLost(
	sig func(a *GApplication, _ uptr) bool) uint64 {
	return app.SignalConnect("name-lost", sig, nil)
}
func (app *GApplication) ConnectOpen(
	sig func(a *GApplication, files uptr, nFiles int32, hint string)) uint64 {
	return app.SignalConnect("open", sig, nil)
}
func (app *GApplication) ConnectShutdown(sig func(a *GApplication)) uint64 {
	return app.SignalConnect("shutdown", sig, nil)
}
func (app *GApplication) ConnectStartup(sig func(a *GApplication)) uint64 {
	return app.SignalConnect("startup", sig, nil)
}

// #endregion

// #region AsyncResult

type GAsyncResult struct{ gobject.GObjectCore }

// #endregion

// #region Cancellable

type GCancellableObj struct {
	Parent gobject.GObjectObj
	Priv   uptr
}
type GCancellableClass struct {
	Parent    gobject.GObjectClass
	Cancelled cc.Func //void (* cancelled) (GCancellable *cancellable);
	_         [5]uptr //reserved
}

type GCancellable struct{ gobject.GObjectCore }

func GTypeGCancellable() gobject.GType    { return g_cancellable_get_type.Fn()() }
func NewGCancellable() *GCancellable      { return g_cancellable_new.Fn()() }
func (c *GCancellable) IsCancelled() bool { return g_cancellable_is_cancelled.Fn()(c) }
func (c *GCancellable) SetErrorIfCancelled(err **uptr) bool {
	return g_cancellable_set_error_if_cancelled.Fn()(c, err)
}
func (c *GCancellable) GetFd() int32 { return g_cancellable_get_fd.Fn()(c) }
func (c *GCancellable) MakePollfd(pollfd *glib.GPollFD) bool {
	return g_cancellable_make_pollfd.Fn()(c, pollfd)
}
func (c *GCancellable) ReleaseFd()               { g_cancellable_release_fd.Fn()(c) }
func (c *GCancellable) SourceNew() *glib.GSource { return g_cancellable_source_new.Fn()(c) }
func GetCurrentGCancellable() *GCancellable      { return g_cancellable_get_current.Fn()() }
func PushCurrentGCancellable(c *GCancellable)    { g_cancellable_push_current.Fn()(c) }
func PopCurrentGCancellable()                    { g_cancellable_pop_current.Fn()(nil) }
func (c *GCancellable) Reset()                   { g_cancellable_reset.Fn()(c) }
func (c *GCancellable) Connect(callback func(cancellable *GCancellable, data uptr), data interface{}, dataDestroyFunc func(uptr)) uint64 {
	return g_cancellable_connect.Fn()(c, vcbu(callback), anyptr(data), vcbu(dataDestroyFunc))
}
func (c *GCancellable) Disconnect(handlerID uint64) { g_cancellable_disconnect.Fn()(c, handlerID) }
func (c *GCancellable) Cancel()                     { g_cancellable_cancel.Fn()(c) }
func (c *GCancellable) ConnectCancelled(sig func(c *GCancellable)) {
	c.SignalConnect("cancelled", sig, nil)
}

// #endregion

// #region File

type GFileIfaceObj struct {
	GIface gobject.GTypeInterface

	Dup                    cc.Func // GFile *(*dup)(GFile *file);
	Hash                   cc.Func // guint (*hash)(GFile *file);
	Equal                  cc.Func // gboolean (*equal)(GFile *file1, GFile *file2);
	IsNative               cc.Func // gboolean (*is_native)(GFile *file);
	HasUriScheme           cc.Func // gboolean (*has_uri_scheme)(GFile *file, const char *uri_scheme);
	GetUriScheme           cc.Func // char *(*get_uri_scheme)(GFile *file);
	GetBasename            cc.Func // char *(*get_basename)(GFile *file);
	GetPath                cc.Func // char *(*get_path)(GFile *file);
	GetUri                 cc.Func // char *(*get_uri)(GFile *file);
	GetParseName           cc.Func // char *(*get_parse_name)(GFile *file);
	GetParent              cc.Func // GFile *(*get_parent)(GFile *file);
	PrefixMatches          cc.Func // gboolean (*prefix_matches)(GFile *prefix, GFile *file);
	GetRelativePath        cc.Func // char *(*get_relative_path)(GFile *parent, GFile *descendant);
	ResolveRelativePath    cc.Func // GFile *(*resolve_relative_path)(GFile *file, const char *relative_path);
	GetChildForDisplayName cc.Func // GFile *(*get_child_for_display_name)(GFile *file, const char *display_name, GError **error);

	EnumerateChildren       cc.Func // GFileEnumerator *(*enumerate_children)(GFile *file, const char *attributes, GFileQueryInfoFlags flags, GCancellable *cancellable, GError **error);
	EnumerateChildrenAsync  cc.Func // void (*enumerate_children_async)(GFile *file, const char *attributes, GFileQueryInfoFlags flags, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	EnumerateChildrenFinish cc.Func // GFileEnumerator *(*enumerate_children_finish)(GFile *file, GAsyncResult *res, GError **error);

	QueryInfo       cc.Func // GFileInfo *(*query_info)(GFile *file, const char *attributes, GFileQueryInfoFlags flags, GCancellable *cancellable, GError **error);
	QueryInfoAsync  cc.Func // void (*query_info_async)(GFile *file, const char *attributes, GFileQueryInfoFlags flags, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	QueryInfoFinish cc.Func // GFileInfo *(*query_info_finish)(GFile *file, GAsyncResult *res, GError **error);

	QueryFilesystemInfo       cc.Func // GFileInfo *(*query_filesystem_info)(GFile *file, const char *attributes, GCancellable *cancellable, GError **error);
	QueryFilesystemInfoAsync  cc.Func // void (*query_filesystem_info_async)(GFile *file, const char *attributes, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	QueryFilesystemInfoFinish cc.Func // GFileInfo *(*query_filesystem_info_finish)(GFile *file, GAsyncResult *res, GError **error);

	FindEnclosingMount       cc.Func // GMount *(*find_enclosing_mount)(GFile *file, GCancellable *cancellable, GError **error);
	FindEnclosingMountAsync  cc.Func // void (*find_enclosing_mount_async)(GFile *file, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	FindEnclosingMountFinish cc.Func // GMount *(*find_enclosing_mount_finish)(GFile *file, GAsyncResult *res, GError **error);

	SetDisplayName       cc.Func // GFile *(*set_display_name)(GFile *file, const char *display_name, GCancellable *cancellable, GError **error);
	SetDisplayNameAsync  cc.Func // void (*set_display_name_async)(GFile *file, const char *display_name, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	SetDisplayNameFinish cc.Func // GFile *(*set_display_name_finish)(GFile *file, GAsyncResult *res, GError **error);

	QuerySettableAttributes        cc.Func // GFileAttributeInfoList *(*query_settable_attributes)(GFile *file, GCancellable *cancellable, GError **error);
	_QuerySettableAttributesAsync  cc.Func // void (*_query_settable_attributes_async)(void);
	_QuerySettableAttributesFinish cc.Func // void (*_query_settable_attributes_finish)(void);

	QueryWritableNamespaces        cc.Func // GFileAttributeInfoList *(*query_writable_namespaces)(GFile *file, GCancellable *cancellable, GError **error);
	_QueryWritableNamespacesAsync  cc.Func // void (*_query_writable_namespaces_async)(void);
	_QueryWritableNamespacesFinish cc.Func // void (*_query_writable_namespaces_finish)(void);

	SetAttribute          cc.Func // gboolean (*set_attribute)(GFile *file, const char *attribute, GFileAttributeType type, gpointer value_p, GFileQueryInfoFlags flags, GCancellable *cancellable, GError **error);
	SetAttributesFromInfo cc.Func // gboolean (*set_attributes_from_info)(GFile *file, GFileInfo *info, GFileQueryInfoFlags flags, GCancellable *cancellable, GError **error);
	SetAttributesAsync    cc.Func // void (*set_attributes_async)(GFile *file, GFileInfo *info, GFileQueryInfoFlags flags, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	SetAttributesFinish   cc.Func // gboolean (*set_attributes_finish)(GFile *file, GAsyncResult *result, GFileInfo **info, GError **error);

	ReadFn     cc.Func // GFileInputStream *(*read_fn)(GFile *file, GCancellable *cancellable, GError **error);
	ReadAsync  cc.Func // void (*read_async)(GFile *file, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	ReadFinish cc.Func // GFileInputStream *(*read_finish)(GFile *file, GAsyncResult *res, GError **error);

	AppendTo       cc.Func // GFileOutputStream *(*append_to)(GFile *file, GFileCreateFlags flags, GCancellable *cancellable, GError **error);
	AppendToAsync  cc.Func // void (*append_to_async)(GFile *file, GFileCreateFlags flags, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	AppendToFinish cc.Func // GFileOutputStream *(*append_to_finish)(GFile *file, GAsyncResult *res, GError **error);

	Create       cc.Func // GFileOutputStream *(*create)(GFile *file, GFileCreateFlags flags, GCancellable *cancellable, GError **error);
	CreateAsync  cc.Func // void (*create_async)(GFile *file, GFileCreateFlags flags, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	CreateFinish cc.Func // GFileOutputStream *(*create_finish)(GFile *file, GAsyncResult *res, GError **error);

	Replace       cc.Func // GFileOutputStream *(*replace)(GFile *file, const char *etag, gboolean make_backup, GFileCreateFlags flags, GCancellable *cancellable, GError **error);
	ReplaceAsync  cc.Func // void (*replace_async)(GFile *file, const char *etag, gboolean make_backup, GFileCreateFlags flags, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	ReplaceFinish cc.Func // GFileOutputStream *(*replace_finish)(GFile *file, GAsyncResult *res, GError **error);

	DeleteFile       cc.Func // gboolean (*delete_file)(GFile *file, GCancellable *cancellable, GError **error);
	DeleteFileAsync  cc.Func // void (*delete_file_async)(GFile *file, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	DeleteFileFinish cc.Func // gboolean (*delete_file_finish)(GFile *file, GAsyncResult *result, GError **error);

	Trash       cc.Func // gboolean (*trash)(GFile *file, GCancellable *cancellable, GError **error);
	TrashAsync  cc.Func // void (*trash_async)(GFile *file, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	TrashFinish cc.Func // gboolean (*trash_finish)(GFile *file, GAsyncResult *result, GError **error);

	MakeDirectory       cc.Func // gboolean (*make_directory)(GFile *file, GCancellable *cancellable, GError **error);
	MakeDirectoryAsync  cc.Func // void (*make_directory_async)(GFile *file, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	MakeDirectoryFinish cc.Func // gboolean (*make_directory_finish)(GFile *file, GAsyncResult *result, GError **error);

	MakeSymbolicLink       cc.Func // gboolean (*make_symbolic_link)(GFile *file, const char *symlink_value, GCancellable *cancellable, GError **error);
	MakeSymbolicLinkAsync  cc.Func // void (*make_symbolic_link_async)(GFile *file, const char *symlink_value, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	MakeSymbolicLinkFinish cc.Func // gboolean (*make_symbolic_link_finish)(GFile *file, GAsyncResult *result, GError **error);

	Copy       cc.Func // gboolean (*copy)(GFile *source, GFile *destination, GFileCopyFlags flags, GCancellable *cancellable, GFileProgressCallback progress_callback, gpointer progress_callback_data, GError **error);
	CopyAsync  cc.Func // void (*copy_async)(GFile *source, GFile *destination, GFileCopyFlags flags, int io_priority, GCancellable *cancellable, GFileProgressCallback progress_callback, gpointer progress_callback_data, GAsyncReadyCallback callback, gpointer user_data);
	CopyFinish cc.Func // gboolean (*copy_finish)(GFile *file, GAsyncResult *res, GError **error);

	Move       cc.Func // gboolean (*move)(GFile *source, GFile *destination, GFileCopyFlags flags, GCancellable *cancellable, GFileProgressCallback progress_callback, gpointer progress_callback_data, GError **error);
	MoveAsync  cc.Func // void (*move_async)(GFile *source, GFile *destination, GFileCopyFlags flags, int io_priority, GCancellable *cancellable, GFileProgressCallback progress_callback, gpointer progress_callback_data, GAsyncReadyCallback callback, gpointer user_data);
	MoveFinish cc.Func // gboolean (*move_finish)(GFile *file, GAsyncResult *result, GError **error);

	MountMountable       cc.Func // void (*mount_mountable)(GFile *file, GMountMountFlags flags, GMountOperation *mount_operation, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	MountMountableFinish cc.Func // GFile *(*mount_mountable_finish)(GFile *file, GAsyncResult *result, GError **error);

	UnmountMountable       cc.Func // void (*unmount_mountable)(GFile *file, GMountUnmountFlags flags, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	UnmountMountableFinish cc.Func // gboolean (*unmount_mountable_finish)(GFile *file, GAsyncResult *result, GError **error);

	EjectMountable       cc.Func // void (*eject_mountable)(GFile *file, GMountUnmountFlags flags, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	EjectMountableFinish cc.Func // gboolean (*eject_mountable_finish)(GFile *file, GAsyncResult *result, GError **error);

	MountEnclosingVolume       cc.Func // void (*mount_enclosing_volume)(GFile *location, GMountMountFlags flags, GMountOperation *mount_operation, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	MountEnclosingVolumeFinish cc.Func // gboolean (*mount_enclosing_volume_finish)(GFile *location, GAsyncResult *result, GError **error);

	MonitorDir  cc.Func // GFileMonitor *(*monitor_dir)(GFile *file, GFileMonitorFlags flags, GCancellable *cancellable, GError **error);
	MonitorFile cc.Func // GFileMonitor *(*monitor_file)(GFile *file, GFileMonitorFlags flags, GCancellable *cancellable, GError **error);

	OpenReadwrite       cc.Func // GFileIOStream *(*open_readwrite)(GFile *file, GCancellable *cancellable, GError **error);
	OpenReadwriteAsync  cc.Func // void (*open_readwrite_async)(GFile *file, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	OpenReadwriteFinish cc.Func // GFileIOStream *(*open_readwrite_finish)(GFile *file, GAsyncResult *res, GError **error);

	CreateReadwrite       cc.Func // GFileIOStream *(*create_readwrite)(GFile *file, GFileCreateFlags flags, GCancellable *cancellable, GError **error);
	CreateReadwriteAsync  cc.Func // void (*create_readwrite_async)(GFile *file, GFileCreateFlags flags, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	CreateReadwriteFinish cc.Func // GFileIOStream *(*create_readwrite_finish)(GFile *file, GAsyncResult *res, GError **error);

	ReplaceReadwrite       cc.Func // GFileIOStream *(*replace_readwrite)(GFile *file, const char *etag, gboolean make_backup, GFileCreateFlags flags, GCancellable *cancellable, GError **error);
	ReplaceReadwriteAsync  cc.Func // void (*replace_readwrite_async)(GFile *file, const char *etag, gboolean make_backup, GFileCreateFlags flags, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	ReplaceReadwriteFinish cc.Func // GFileIOStream *(*replace_readwrite_finish)(GFile *file, GAsyncResult *res, GError **error);

	StartMountable       cc.Func // void (*start_mountable)(GFile *file, GDriveStartFlags flags, GMountOperation *start_operation, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	StartMountableFinish cc.Func // gboolean (*start_mountable_finish)(GFile *file, GAsyncResult *result, GError **error);

	StopMountable       cc.Func // void (*stop_mountable)(GFile *file, GMountUnmountFlags flags, GMountOperation *mount_operation, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	StopMountableFinish cc.Func // gboolean (*stop_mountable_finish)(GFile *file, GAsyncResult *result, GError **error);

	SupportsThreadContexts bool // gboolean supports_thread_contexts;

	UnmountMountableWithOperation       cc.Func // void (*unmount_mountable_with_operation)(GFile *file, GMountUnmountFlags flags, GMountOperation *mount_operation, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	UnmountMountableWithOperationFinish cc.Func // gboolean (*unmount_mountable_with_operation_finish)(GFile *file, GAsyncResult *result, GError **error);

	EjectMountableWithOperation       cc.Func // void (*eject_mountable_with_operation)(GFile *file, GMountUnmountFlags flags, GMountOperation *mount_operation, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	EjectMountableWithOperationFinish cc.Func // gboolean (*eject_mountable_with_operation_finish)(GFile *file, GAsyncResult *result, GError **error);

	PollMountable       cc.Func // void (*poll_mountable)(GFile *file, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	PollMountableFinish cc.Func // gboolean (*poll_mountable_finish)(GFile *file, GAsyncResult *result, GError **error);

	MeasureDiskUsage       cc.Func // gboolean (*measure_disk_usage)(GFile *file, GFileMeasureFlags flags, GCancellable *cancellable, GFileMeasureProgressCallback progress_callback, gpointer progress_data, guint64 *disk_usage, guint64 *num_dirs, guint64 *num_files, GError **error);
	MeasureDiskUsageAsync  cc.Func // void (*measure_disk_usage_async)(GFile *file, GFileMeasureFlags flags, gint io_priority, GCancellable *cancellable, GFileMeasureProgressCallback progress_callback, gpointer progress_data, GAsyncReadyCallback callback, gpointer user_data);
	MeasureDiskUsageFinish cc.Func // gboolean (*measure_disk_usage_finish)(GFile *file, GAsyncResult *result, guint64 *disk_usage, guint64 *num_dirs, guint64 *num_files, GError **error);

	QueryExists cc.Func // gboolean (*query_exists)(GFile *file, GCancellable *cancellable);
}

type GFileIface interface {
	GetGFileIface() *GFile
}

func GetGFileIface(iface GFileIface) *GFile {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetGFileIface()
}
func (iface *GFile) GetGFileIface() *GFile { return iface }

type GFile struct{ gobject.GObjectCore }

// #endregion

// #region FileInfo

type GFileInfo struct{ gobject.GObjectCore }

func GTypeFileInfo() gobject.GType            { return g_file_info_get_type.Fn()() }
func NewFileInfo() *GFileInfo                 { return g_file_info_new.Fn()() }
func (f *GFileInfo) Dup() *GFileInfo          { return g_file_info_dup.Fn()(f) }
func (f *GFileInfo) CopyInto(dest *GFileInfo) { g_file_info_copy_into.Fn()(f, dest) }
func (f *GFileInfo) HasAttribute(attribute string) bool {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_info_has_attribute.Fn()(f, cstr)
}
func (f *GFileInfo) HasNamespace(nameSpace string) bool {
	cstr := cc.NewString(nameSpace)
	defer cstr.Free()
	return g_file_info_has_namespace.Fn()(f, cstr)
}
func (f *GFileInfo) ListAttributes(nameSpace string) []string {
	cstr := cc.NewString(nameSpace)
	defer cstr.Free()
	ss := g_file_info_list_attributes.Fn()(f, cstr)
	defer ss.Free()
	return ss.Strings()
}
func (f *GFileInfo) GetAttributeData(attribute string) (typ GFileAttributeType, value uptr, status GFileAttributeStatus, ok bool) {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	ok = g_file_info_get_attribute_data.Fn()(f, cstr, &typ, &value, &status)
	return
}
func (f *GFileInfo) GetAttributeType(attribute string) GFileAttributeType {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_info_get_attribute_type.Fn()(f, cstr)
}
func (f *GFileInfo) RemoveAttribute(attribute string) {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	g_file_info_remove_attribute.Fn()(f, cstr)
}
func (f *GFileInfo) GetAttributeStatus(attribute string) GFileAttributeStatus {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_info_get_attribute_status.Fn()(f, cstr)
}
func (f *GFileInfo) SetAttributeStatus(attribute string, status GFileAttributeStatus) bool {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_info_set_attribute_status.Fn()(f, cstr, status)
}
func (f *GFileInfo) GetAttributeAsString(attribute string) string {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_info_get_attribute_as_string.Fn()(f, cstr).String()
}
func (f *GFileInfo) GetAttributeString(attribute string) string {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_info_get_attribute_string.Fn()(f, cstr).String()
}
func (f *GFileInfo) GetAttributeByteString(attribute string) string {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_info_get_attribute_byte_string.Fn()(f, cstr).String()
}
func (f *GFileInfo) GetAttributeBoolean(attribute string) bool {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_info_get_attribute_boolean.Fn()(f, cstr)
}
func (f *GFileInfo) GetAttributeUint32(attribute string) uint32 {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_info_get_attribute_uint32.Fn()(f, cstr)
}
func (f *GFileInfo) GetAttributeInt32(attribute string) int32 {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_info_get_attribute_int32.Fn()(f, cstr)
}
func (f *GFileInfo) GetAttributeUint64(attribute string) uint64 {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_info_get_attribute_uint64.Fn()(f, cstr)
}
func (f *GFileInfo) GetAttributeInt64(attribute string) int64 {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_info_get_attribute_int64.Fn()(f, cstr)
}
func (f *GFileInfo) GetAttributeObject(attribute string) uptr {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_info_get_attribute_object.Fn()(f, cstr)
}
func (f *GFileInfo) GetAttributeStringv(attribute string) []string {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	ss := g_file_info_get_attribute_stringv.Fn()(f, cstr)
	defer ss.Free()
	return ss.Strings()
}
func (f *GFileInfo) GetAttributeFilePath(attribute string) string {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_info_get_attribute_file_path.Fn()(f, cstr).String()
}

func (f *GFileInfo) SetAttribute(attribute string, typ GFileAttributeType, valueAddr interface{}) {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	g_file_info_set_attribute.Fn()(f, cstr, typ, anyptr(valueAddr))
}
func (f *GFileInfo) SetAttributeString(attribute string, attr_value string) {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	cval := cc.NewString(attr_value)
	defer cval.Free()
	g_file_info_set_attribute_string.Fn()(f, cstr, cval)
}
func (f *GFileInfo) SetAttributeByteString(attribute string, attr_value string) {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	cval := cc.NewString(attr_value)
	defer cval.Free()
	g_file_info_set_attribute_byte_string.Fn()(f, cstr, cval)
}
func (f *GFileInfo) SetAttributeBoolean(attribute string, attr_value bool) {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	g_file_info_set_attribute_boolean.Fn()(f, cstr, attr_value)
}
func (f *GFileInfo) SetAttributeUint32(attribute string, attr_value uint32) {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	g_file_info_set_attribute_uint32.Fn()(f, cstr, attr_value)
}
func (f *GFileInfo) SetAttributeInt32(attribute string, attr_value int32) {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	g_file_info_set_attribute_int32.Fn()(f, cstr, attr_value)
}
func (f *GFileInfo) SetAttributeUint64(attribute string, attr_value uint64) {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	g_file_info_set_attribute_uint64.Fn()(f, cstr, attr_value)
}
func (f *GFileInfo) SetAttributeInt64(attribute string, attr_value int64) {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	g_file_info_set_attribute_int64.Fn()(f, cstr, attr_value)
}
func (f *GFileInfo) SetAttributeObject(attribute string, attr_value *gobject.GObject) {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	g_file_info_set_attribute_object.Fn()(f, cstr, attr_value)
}
func (f *GFileInfo) SetAttributeStringv(attribute string, attr_value []string) {
	cstr, cstrs := cc.NewString(attribute), cc.NewStrings(attr_value)
	defer cstr.Free()
	defer cstrs.Free()
	g_file_info_set_attribute_stringv.Fn()(f, cstr, cstrs)
}
func (f *GFileInfo) SetAttributeFilePath(attribute string, attr_value string) {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	cval := cc.NewString(attr_value)
	defer cval.Free()
	g_file_info_set_attribute_file_path.Fn()(f, cstr, cval)
}

func (f *GFileInfo) ClearStatus() { g_file_info_clear_status.Fn()(f) }

func (f *GFileInfo) GetDeletionDate() uptr   { return g_file_info_get_deletion_date.Fn()(f) }
func (f *GFileInfo) GetFileType() GFileType  { return g_file_info_get_file_type.Fn()(f) }
func (f *GFileInfo) GetIsHidden() bool       { return g_file_info_get_is_hidden.Fn()(f) }
func (f *GFileInfo) GetIsBackup() bool       { return g_file_info_get_is_backup.Fn()(f) }
func (f *GFileInfo) GetIsSymlink() bool      { return g_file_info_get_is_symlink.Fn()(f) }
func (f *GFileInfo) GetName() string         { return g_file_info_get_name.Fn()(f).String() }
func (f *GFileInfo) GetDisplayName() string  { return g_file_info_get_display_name.Fn()(f).String() }
func (f *GFileInfo) GetEditName() string     { return g_file_info_get_edit_name.Fn()(f).String() }
func (f *GFileInfo) GetIcon() *GIcon         { return g_file_info_get_icon.Fn()(f) }
func (f *GFileInfo) GetSymbolicIcon() *GIcon { return g_file_info_get_symbolic_icon.Fn()(f) }
func (f *GFileInfo) GetContentType() string  { return g_file_info_get_content_type.Fn()(f).String() }
func (f *GFileInfo) GetSize() int64          { return g_file_info_get_size.Fn()(f) }
func (f *GFileInfo) GetModificationDateTime() *glib.GDateTime {
	return g_file_info_get_modification_date_time.Fn()(f)
}
func (f *GFileInfo) GetAccessDateTime() *glib.GDateTime {
	return g_file_info_get_access_date_time.Fn()(f)
}
func (f *GFileInfo) GetCreationDateTime() *glib.GDateTime {
	return g_file_info_get_creation_date_time.Fn()(f)
}
func (f *GFileInfo) GetSymlinkTarget() string { return g_file_info_get_symlink_target.Fn()(f).String() }
func (f *GFileInfo) GetEtag() string          { return g_file_info_get_etag.Fn()(f).String() }
func (f *GFileInfo) GetSortOrder() int32      { return g_file_info_get_sort_order.Fn()(f) }

func (f *GFileInfo) SetAttributeMask(mask *GFileAttributeMatcher) {
	g_file_info_set_attribute_mask.Fn()(f, mask)
}
func (f *GFileInfo) UnsetAttributeMask()          { g_file_info_unset_attribute_mask.Fn()(f) }
func (f *GFileInfo) SetFileType(typ GFileType)    { g_file_info_set_file_type.Fn()(f, typ) }
func (f *GFileInfo) SetIsHidden(is_hidden bool)   { g_file_info_set_is_hidden.Fn()(f, is_hidden) }
func (f *GFileInfo) SetIsSymlink(is_symlink bool) { g_file_info_set_is_symlink.Fn()(f, is_symlink) }
func (f *GFileInfo) SetName(name string) {
	cstr := cc.NewString(name)
	defer cstr.Free()
	g_file_info_set_name.Fn()(f, cstr)
}
func (f *GFileInfo) SetDisplayName(display_name string) {
	cstr := cc.NewString(display_name)
	defer cstr.Free()
	g_file_info_set_display_name.Fn()(f, cstr)
}
func (f *GFileInfo) SetEditName(edit_name string) {
	cstr := cc.NewString(edit_name)
	defer cstr.Free()
	g_file_info_set_edit_name.Fn()(f, cstr)
}
func (f *GFileInfo) SetIcon(icon *GIcon)         { g_file_info_set_icon.Fn()(f, icon) }
func (f *GFileInfo) SetSymbolicIcon(icon *GIcon) { g_file_info_set_symbolic_icon.Fn()(f, icon) }
func (f *GFileInfo) SetContentType(content_type string) {
	cstr := cc.NewString(content_type)
	defer cstr.Free()
	g_file_info_set_content_type.Fn()(f, cstr)
}
func (f *GFileInfo) SetSize(size int64) { g_file_info_set_size.Fn()(f, size) }
func (f *GFileInfo) SetModificationDateTime(mtime *glib.GDateTime) {
	g_file_info_set_modification_date_time.Fn()(f, mtime)
}
func (f *GFileInfo) SetAccessDateTime(atime *glib.GDateTime) {
	g_file_info_set_access_date_time.Fn()(f, atime)
}
func (f *GFileInfo) SetCreationDateTime(ctime *glib.GDateTime) {
	g_file_info_set_creation_date_time.Fn()(f, ctime)
}
func (f *GFileInfo) SetSymlinkTarget(symlink_target string) {
	cstr := cc.NewString(symlink_target)
	defer cstr.Free()
	g_file_info_set_symlink_target.Fn()(f, cstr)
}
func (f *GFileInfo) SetSortOrder(sort_order int32) { g_file_info_set_sort_order.Fn()(f, sort_order) }

type GFileAttributeMatcher struct{}

func GTypeFileAttributeMatcher() gobject.GType { return g_file_attribute_matcher_get_type.Fn()() }
func NewFileAttributeMatcher(attributes string) *GFileAttributeMatcher {
	cstr := cc.NewString(attributes)
	defer cstr.Free()
	return g_file_attribute_matcher_new.Fn()(cstr)
}
func (m *GFileAttributeMatcher) Ref() *GFileAttributeMatcher {
	return g_file_attribute_matcher_ref.Fn()(m)
}
func (m *GFileAttributeMatcher) Unref() { g_file_attribute_matcher_unref.Fn()(m) }
func (m *GFileAttributeMatcher) Subtract(subtract *GFileAttributeMatcher) *GFileAttributeMatcher {
	return g_file_attribute_matcher_subtract.Fn()(m, subtract)
}
func (m *GFileAttributeMatcher) Matches(attribute string) bool {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_attribute_matcher_matches.Fn()(m, cstr)
}
func (m *GFileAttributeMatcher) MatchesOnly(attribute string) bool {
	cstr := cc.NewString(attribute)
	defer cstr.Free()
	return g_file_attribute_matcher_matches_only.Fn()(m, cstr)
}
func (m *GFileAttributeMatcher) EnumerateNamespace(ns string) bool {
	cstr := cc.NewString(ns)
	defer cstr.Free()
	return g_file_attribute_matcher_enumerate_namespace.Fn()(m, cstr)
}
func (m *GFileAttributeMatcher) EnumerateNext() string {
	return g_file_attribute_matcher_enumerate_next.Fn()(m).String()
}
func (m *GFileAttributeMatcher) ToString() string {
	return g_file_attribute_matcher_to_string.Fn()(m).String()
}

// #endregion

// #region FileInputStream

type GFileInputStreamObj struct {
	Parent GInputStream
	Priv   uptr
}
type GFileInputStreamClass struct {
	Parent GInputStreamClass

	Tell            cc.Func // goffset (* tell) (GFileInputStream *stream);
	CanSeek         cc.Func // gboolean (* can_seek) (GFileInputStream *stream);
	Seek            cc.Func // gboolean (* seek)	 (GFileInputStream *stream, goffset offset, GSeekType type, GCancellable *cancellable, GError **error);
	QueryInfo       cc.Func // GFileInfo * (* query_info) (GFileInputStream *stream, const char *attributes, GCancellable *cancellable, GError **error);
	QueryInfoAsync  cc.Func // void (* query_info_async) (GFileInputStream *stream, const char *attributes, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	QueryInfoFinish cc.Func // GFileInfo * (* query_info_finish) (GFileInputStream *stream, GAsyncResult *result, GError **error);

	_ [5]cc.Func //reserved
}

type GFileInputStream struct{ GInputStream }

// #endregion

// #region FileIOStream

type GFileIOStreamObj struct {
	Parent GIOStreamObj
	Priv   uptr
}

type GFileIOStreamClass struct {
	Parent GIOStreamClass

	Tell            cc.Func // goffset (* tell) (GFileIOStream *stream);
	CanSeek         cc.Func // gboolean (* can_seek) (GFileIOStream *stream);
	Seek            cc.Func // gboolean (* seek)	 (GFileIOStream *stream, goffset offset, GSeekType type, GCancellable *cancellable, GError **error);
	CanTruncate     cc.Func // gboolean (* can_truncate) (GFileIOStream *stream);
	TruncateFn      cc.Func // gboolean (* truncate_fn) (GFileIOStream *stream, goffset size, GCancellable *cancellable, GError **error);
	QueryInfo       cc.Func // GFileInfo * (* query_info) (GFileIOStream *stream, const char *attributes, GCancellable *cancellable, GError **error);
	QueryInfoAsync  cc.Func // void (* query_info_async) (GFileIOStream *stream, const char *attributes, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	QueryInfoFinish cc.Func // GFileInfo * (* query_info_finish) (GFileIOStream *stream, GAsyncResult *result, GError **error);
	GetEtag         cc.Func // char * (* get_etag) (GFileIOStream *stream);

	_ [5]uptr //reserved
}

type GFileIOStream struct{ GIOStream }

func GTypeGFileIOStream() gobject.GType { return g_file_io_stream_get_type.Fn()() }

func (s *GFileIOStream) QueryInfo(attributes string, cancellable *GCancellable) (*GFileInfo, error) {
	cstr := cc.NewString(attributes)
	defer cstr.Free()
	var gerr *glib.GError
	inf := g_file_io_stream_query_info.Fn()(s, cstr, cancellable, &gerr)
	if gerr != nil {
		return nil, gerr.Error()
	}
	return inf, nil
}
func (s *GFileIOStream) QueryInfoAsync(attributes string, io_priority int32,
	cancellable *GCancellable, callback func(src *gobject.GObject, res GAsyncResult, _ uptr), user_data interface{}) {
	cstr := cc.NewString(attributes)
	defer cstr.Free()
	g_file_io_stream_query_info_async.Fn()(s, cstr, io_priority, cancellable, vcbu(callback), anyptr(user_data))
}
func (s *GFileIOStream) QueryInfoFinish(result *GAsyncResult) (*GFileInfo, error) {
	var gerr *glib.GError
	inf := g_file_io_stream_query_info_finish.Fn()(s, result, &gerr)
	if gerr != nil {
		return nil, gerr.Error()
	}
	return inf, nil
}
func (s *GFileIOStream) GetEtag() string { return g_file_io_stream_get_etag.Fn()(s).String() }

// #endregion

// #region FileOutputStream

type GFileOutputStreamObj struct {
	Parent GOutputStreamObj
	Priv   uptr
}
type GFileOutputStreamClass struct {
	Parent GOutputStreamClass

	Tell            cc.Func // goffset (* tell) (GFileOutputStream *stream);
	CanSeek         cc.Func // gboolean (* can_seek) (GFileOutputStream *stream);
	Seek            cc.Func // gboolean (* seek)	 (GFileOutputStream *stream, goffset offset, GSeekType type, GCancellable *cancellable, GError **error);
	CanTruncate     cc.Func // gboolean (* can_truncate) (GFileOutputStream *stream);
	TruncateFn      cc.Func // gboolean (* truncate_fn) (GFileOutputStream *stream, goffset size, GCancellable *cancellable, GError **error);
	QueryInfo       cc.Func // GFileInfo * (* query_info) (GFileOutputStream *stream, const char *attributes, GCancellable *cancellable, GError **error);
	QueryInfoAsync  cc.Func // void (* query_info_async) (GFileOutputStream *stream, const char *attributes, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	QueryInfoFinish cc.Func // GFileInfo * (* query_info_finish) (GFileOutputStream *stream, GAsyncResult *result, GError **error);
	GetEtag         cc.Func // char * (* get_etag) (GFileOutputStream *stream);

	_ [5]cc.Func //reserved
}

type GFileOutputStream struct{ GOutputStream }

// #endregion

// #region Icon
type GIconIfaceObj struct {
	GIface     gobject.GTypeInterface
	Hash       cc.Func // guint      (* hash)        (GIcon *icon);
	Equal      cc.Func // gboolean   (* equal)       (GIcon *icon1, GIcon *icon2);
	ToTokens   cc.Func // gboolean   (* to_tokens)   (GIcon *icon, GPtrArray *tokens, gint *out_version);
	FromTokens cc.Func // GIcon *    (* from_tokens) (char **tokens, gint num_tokens, gint version, GError **error);
	Serialize  cc.Func // GVariant * (* serialize)   (GIcon *icon);
}

type GIconIface interface {
	GetGIconIface() *GIcon
}

func GetGIconIface(iface GIconIface) *GIcon {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetGIconIface()
}

func (icon *GIcon) GetGIconIface() *GIcon { return icon }

type GIcon struct{ gobject.GObjectCore }

func GTypeGIcon() gobject.GType { return g_icon_get_type.Fn()() }

func (icon *GIcon) Hash() uint32                { return g_icon_hash.Fn()(icon) }
func (icon *GIcon) Equal(icon2 GIconIface) bool { return g_icon_equal.Fn()(icon, GetGIconIface(icon2)) }
func (icon *GIcon) ToString() string            { return g_icon_to_string.Fn()(icon).String() }
func NewGIconForString(str string) (*GIcon, error) {
	cstr := cc.NewString(str)
	defer cstr.Free()
	var gerr *glib.GError
	icon := g_icon_new_for_string.Fn()(cstr, &gerr)
	if gerr != nil {
		defer gerr.Free()
		return nil, gerr.Error()
	}
	return icon, nil
}
func (icon *GIcon) Serialize() *glib.GVariant         { return g_icon_serialize.Fn()(icon) }
func NewGIconDeserialize(value *glib.GVariant) *GIcon { return g_icon_deserialize.Fn()(value) }

// #endregion

// #region IOStream

type GIOStreamObj struct {
	Parent gobject.GObjectObj
	Priv   uptr
}

type GIOStreamClass struct {
	Parent gobject.GObjectClass

	GetInputStream  cc.Func // GInputStream *  (*get_input_stream)  (GIOStream *stream);
	GetOutputStream cc.Func // GOutputStream * (*get_output_stream) (GIOStream *stream);

	CloseFn     cc.Func // gboolean (* close_fn)	 (GIOStream *stream,GCancellable *cancellable,GError **error);
	CloseAsync  cc.Func // void (* close_async) (GIOStream *stream,int io_priority,GCancellable *cancellable,GAsyncReadyCallback callback,gpointer user_data);
	CloseFinish cc.Func // gboolean (* close_finish) (GIOStream *stream,GAsyncResult *result,GError **error);

	_ [10]cc.Func
}

type GIOStreamIface interface {
	GetGIOStreamIface() *GIOStream
}

func GetGIOStreamIface(iface GIOStreamIface) *GIOStream {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetGIOStreamIface()
}

func (s *GIOStream) GetGIOStreamIface() *GIOStream { return s }

type GIOStream struct{ gobject.GObjectCore }

func GTypeGIOStream() gobject.GType { return g_io_stream_get_type.Fn()() }

func (s *GIOStream) GetInputStream() *GInputStream   { return g_io_stream_get_input_stream.Fn()(s) }
func (s *GIOStream) GetOutputStream() *GOutputStream { return g_io_stream_get_output_stream.Fn()(s) }
func (s *GIOStream) SpliceAsync(stream2 GIOStreamIface, flags GIOStreamSpliceFlags, io_priority int32,
	cancellable *GCancellable, callback func(src *gobject.GObject, res *GAsyncResult, user_data uptr), user_data interface{}) {
	g_io_stream_splice_async.Fn()(s, GetGIOStreamIface(stream2), flags, io_priority, cancellable, vcbu(callback), anyptr(user_data))
}
func GIOStreamSpliceFinish(result *GAsyncResult) error {
	var gerr *glib.GError
	ok := g_io_stream_splice_finish.Fn()(result, &gerr)
	if ok {
		return nil
	}
	defer gerr.Free()
	return gerr.Error()

}
func (s *GIOStream) Close(cancellable *GCancellable, err **uptr) error {
	var gerr *glib.GError
	ok := g_io_stream_close.Fn()(s, cancellable, &gerr)
	if ok {
		return nil
	}
	defer gerr.Free()
	return gerr.Error()
}
func (s *GIOStream) CloseAsync(io_priority int32, cancellable *GCancellable,
	callback func(src *gobject.GObject, res *GAsyncResult, user_data uptr), user_data interface{}) {
	g_io_stream_close_async.Fn()(s, io_priority, cancellable, vcbu(callback), anyptr(user_data))
}
func (s *GIOStream) CloseFinish(result *GAsyncResult) error {
	var gerr *glib.GError
	ok := g_io_stream_close_finish.Fn()(s, result, &gerr)
	if ok {
		return nil
	}
	defer gerr.Free()
	return gerr.Error()
}
func (s *GIOStream) IsClosed() bool   { return g_io_stream_is_closed.Fn()(s) }
func (s *GIOStream) HasPending() bool { return g_io_stream_has_pending.Fn()(s) }
func (s *GIOStream) SetPending(err **uptr) error {
	var gerr *glib.GError
	ok := g_io_stream_set_pending.Fn()(s, &gerr)
	if ok {
		return nil
	}
	defer gerr.Free()
	return gerr.Error()
}
func (s *GIOStream) ClearPending() { g_io_stream_clear_pending.Fn()(s) }

// #endregion

// #region InputStream

type GInputStreamObj struct {
	Parent gobject.GObjectObj
	Priv   *uptr
}
type GInputStreamClass struct {
	Parent gobject.GObjectClass

	ReadFn  cc.Func // gssize (* read_fn) (GInputStream *stream,void *buffer,gsize count,GCancellable *cancellable,GError **error);
	Skip    cc.Func // gssize (* skip) (GInputStream *stream,gsize count,GCancellable *cancellable,GError **error);
	CloseFn cc.Func // gboolean (* close_fn)	 (GInputStream *stream,GCancellable *cancellable,GError **error);

	ReadAsync   cc.Func // void (* read_async) (GInputStream *stream,void *buffer,gsize count,int io_priority,GCancellable *cancellable,GAsyncReadyCallback callback,gpointer user_data);
	ReadFinish  cc.Func // gssize (* read_finish) (GInputStream *stream,GAsyncResult *result,GError **error);
	SkipAsync   cc.Func // void (* skip_async) (GInputStream *stream,gsize count,int io_priority,GCancellable *cancellable,GAsyncReadyCallback callback,gpointer user_data);
	SkipFinish  cc.Func // gssize (* skip_finish) (GInputStream *stream,GAsyncResult *result,GError **error);
	CloseAsync  cc.Func // void (* close_async) (GInputStream *stream,int io_priority,GCancellable *cancellable,GAsyncReadyCallback callback,gpointer user_data);
	CloseFinish cc.Func // gboolean (* close_finish) (GInputStream *stream,GAsyncResult *result,GError **error);

	_ [5]cc.Func
}

type GInputStreamIface interface {
	GetGInputStreamIface() *GInputStream
}

func GetGInputStreamIface(iface GInputStreamIface) *GInputStream {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetGInputStreamIface()
}

func (stream *GInputStream) GetGInputStreamIface() *GInputStream { return stream }

type GInputStream struct{ gobject.GObjectCore }

// #endregion

// #region ListModel

type GListModelIfaceObj struct {
	GIface      gobject.GTypeInterface
	GetItemType cc.Func // func(*GListModel) gobject.GType
	GetNItems   cc.Func // func(*GListModel) uint32
	GetItem     cc.Func // func(*GListModel, uint32) uptr
}

type GListModelIface interface {
	GetGListModelIface() *GListModel
}

func GetGListModelIface(iface GListModelIface) *GListModel {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetGListModelIface()
}

func (list *GListModel) GetGListModelIface() *GListModel { return list }

type GListModel struct{}

func GTypeGListModel() gobject.GType { return g_list_model_get_type.Fn()() }

func (list *GListModel) GetItemType() gobject.GType {
	return g_list_model_get_item_type.Fn()(list)
}
func (list *GListModel) GetNItems() uint32 {
	return g_list_model_get_n_items.Fn()(list)
}
func (list *GListModel) GetItem(position uint32) uptr {
	return g_list_model_get_item.Fn()(list, position)
}
func (list *GListModel) GetObject(position uint32) *gobject.GObject {
	return g_list_model_get_object.Fn()(list, position)
}
func (list *GListModel) ItemsChanged(position, removed, added uint32) {
	g_list_model_items_changed.Fn()(list, position, removed, added)
}
func (list *GListModel) ConnectItemsChanged(
	sig func(list *GListModel, position, removed, added uint32)) uint64 {
	return (*gobject.GObject)(uptr(list)).SignalConnect("items-changed", sig, nil)
}

func GListModelList[T any](l *GListModel) []T {
	num := l.GetNItems()
	if num == 0 {
		return nil
	}

	slice := make([]T, num)

	var t T
	switch any(t).(type) {
	case string:
		for i := uint32(0); i < num; i++ {
			a := ((cc.String)(l.GetItem(i))).String()
			slice[i] = any(a).(T)
		}
	default:
		for i := uint32(0); i < num; i++ {
			a := l.GetItem(i)
			slice[i] = *(*T)(uptr(&a))
		}
	}

	return slice
}

// #endregion

// #region ListStore

type GListStore struct {
	GListModel
}

func GTypeGListStore() gobject.GType { return g_list_store_get_type.Fn()() }

func NewGListStore(itemType gobject.GType) *GListStore {
	return g_list_store_new.Fn()(itemType)
}
func (s *GListStore) Insert(position uint32, item uptr) { g_list_store_insert.Fn()(s, position, item) }
func (s *GListStore) InsertSorted(item uptr, compareFunc uptr, userData uptr) uint32 {
	return g_list_store_insert_sorted.Fn()(s, item, compareFunc, userData)
}
func (s *GListStore) Sort(compareFunc uptr, userData uptr) {
	g_list_store_sort.Fn()(s, compareFunc, userData)
}
func (s *GListStore) Append(itemPtr interface{}) { g_list_store_append.Fn()(s, anyptr(itemPtr)) }
func (s *GListStore) Remove(position uint32)     { g_list_store_remove.Fn()(s, position) }
func (s *GListStore) RemoveAll()                 { g_list_store_remove_all.Fn()(s) }
func (s *GListStore) Splice(position, nRemovals uint32, additions *uptr, nAdditions uint32) {
	g_list_store_splice.Fn()(s, position, nRemovals, additions, nAdditions)
}
func (s *GListStore) Find(item uptr, position *uint32) bool {
	return g_list_store_find.Fn()(s, item, position)
}
func (s *GListStore) FindWithEqualFunc(item uptr, equalFunc uptr, position *uint32) bool {
	return g_list_store_find_with_equal_func.Fn()(s, item, equalFunc, position)
}
func (s *GListStore) FindWithEqualFuncFull(item uptr, equalFunc uptr, userData uptr, position *uint32) bool {
	return g_list_store_find_with_equal_func_full.Fn()(s, item, equalFunc, userData, position)
}

// #endregion

// #region Menu

type GMenu struct{ GMenuModel }

// #endregion

// #region MenuModel

type GMenuModelObj struct {
	Parent gobject.GObject
	Priv   uptr
}
type GMenuModelClass struct {
	Parent gobject.GObjectClass

	IsMutable         cc.Func
	GetNItems         cc.Func
	GetItemAttributes cc.Func

	IterateItemAttributes cc.Func
	GetItemAttributeValue cc.Func

	GetItemLinks     cc.Func
	IterateItemLinks cc.Func
	GetItemLink      cc.Func
}
type GMenuModel struct{ gobject.GObjectCore }

// #endregion

// #region MountOperation

type GMountOperationObj struct {
	Parent gobject.GObjectObj
	Priv   uptr
}

type GMountOperationClass struct {
	Parent gobject.GObjectClass

	AskPassword cc.Func //void (* ask_password) (GMountOperation *op,const char *message,const char *default_user,const char *default_domain, GAskPasswordFlags flags);

	AskQuestion cc.Func // void (* ask_question) (GMountOperation *op,const char* message,const char*choices[]);
	Reply       cc.Func // void (* reply)(GMountOperation *op, GMountOperationResult result);
	Aborted     cc.Func // void (* aborted)(GMountOperation *op);

	ShowProcesses       cc.Func // void (* show_processes)(GMountOperation *op,const gchar *message,GArray *processes,const gchar *choices[]);
	ShowUnmountProgress cc.Func // void (* show_unmount_progress)(GMountOperation *op,const gchar *message,gint64 time_left,gint64 bytes_left);

	_ [8]uptr //reserved
}

type GMountOperationIface interface {
	GetGMountOperationIface() *GMountOperation
}

func GetGMountOperationIface(iface GMountOperationIface) *GMountOperation {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetGMountOperationIface()
}

func (m *GMountOperation) GetGMountOperationIface() *GMountOperation { return m }

type GMountOperation struct{ gobject.GObjectCore }

func (m *GMountOperation) ConnectAborted(sig func(m *GMountOperation)) uint64 {
	return m.SignalConnect("aborted", sig, nil)
}
func (m *GMountOperation) ConnectAskPassword(
	sig func(m *GMountOperation, msg, defaultUser, defaultDomain cc.String, flags GAskPasswordFlags)) uint64 {
	return m.SignalConnect("ask-password", sig, nil)
}
func (m *GMountOperation) ConnectAskQuestion(
	sig func(m *GMountOperation, msg cc.String, choices cc.Strings)) uint64 {
	return m.SignalConnect("ask-question", sig, nil)
}
func (m *GMountOperation) ConnectReply(sig func(m *GMountOperation, result GMountOperationResult)) uint64 {
	return m.SignalConnect("reply", sig, nil)
}
func (m *GMountOperation) ConnectShowProcesses(
	sig func(m *GMountOperation, msg cc.String, processed *cc.Ptr, choices cc.Strings)) uint64 {
	return m.SignalConnect("show-processes", sig, nil)
}
func (m *GMountOperation) ConnectShowUnmountProgress(
	sig func(m *GMountOperation, msg cc.String, timeLeft, BytesLeft int64)) uint64 {
	return m.SignalConnect("show-unmount-progress", sig, nil)
}

// #endregion

// #region OutputStream

type GOutputStreamObj struct {
	Parent gobject.GObjectObj
	Priv   uptr
}

type GOutputStreamClass struct {
	Parent gobject.GObjectClass

	WriteFn cc.Func // gssize (* write_fn) (GOutputStream *stream,const void *buffer,gsize count,GCancellable *cancellable,GError **error);
	Splice  cc.Func // gssize (* splice) (GOutputStream *stream,GInputStream *source,GOutputStreamSpliceFlags flags,GCancellable *cancellable,GError **error);
	Flush   cc.Func // gboolean (* flush)	 (GOutputStream *stream,GCancellable *cancellable,GError **error);
	CloseFn cc.Func // gboolean (* close_fn) (GOutputStream *stream,GCancellable *cancellable,GError **error);

	WriteAsync   cc.Func // void (* write_async) (GOutputStream *stream,const void *buffer,gsize count,int io_priority,GCancellable *cancellable,GAsyncReadyCallback callback,gpointer user_data);
	WriteFinish  cc.Func // gssize (* write_finish) (GOutputStream *stream,GAsyncResult *result,GError **error);
	SpliceAsync  cc.Func // void (* splice_async) (GOutputStream *stream,GInputStream *source,GOutputStreamSpliceFlags flags,int io_priority,GCancellable *cancellable,GAsyncReadyCallback callback,gpointer user_data);
	SpliceFinish cc.Func // gssize (* splice_finish) (GOutputStream *stream,GAsyncResult *result,GError **error);
	FlushAsync   cc.Func // void (* flush_async) (GOutputStream *stream,int io_priority,GCancellable *cancellable,GAsyncReadyCallback callback,gpointer user_data);
	FlushFinish  cc.Func // gboolean (* flush_finish) (GOutputStream *stream,GAsyncResult *result,GError **error);
	CloseAsync   cc.Func // void (* close_async) (GOutputStream *stream,int io_priority,GCancellable *cancellable,GAsyncReadyCallback callback,gpointer user_data);
	CloseFinish  cc.Func // gboolean (* close_finish) (GOutputStream *stream,GAsyncResult *result,GError **error);

	WritevFn     cc.Func // gboolean (* writev_fn) (GOutputStream *stream,const GOutputVector *vectors,gsize n_vectors,gsize *bytes_written,GCancellable *cancellable,GError **error);
	WritevAsync  cc.Func // void (* writev_async) (GOutputStream *stream,const GOutputVector *vectors,gsize n_vectors,int io_priority,GCancellable *cancellable,GAsyncReadyCallback callback,gpointer user_data);
	WritevFinish cc.Func // gboolean (* writev_finish) (GOutputStream *stream,GAsyncResult *result,gsize *bytes_written,GError **error);

	_ [5]cc.Func // reserved
}

type GOutputStreamIface interface {
	GetGOutputStreamIface() *GOutputStream
}

func GetGOutputStreamIface(iface GOutputStreamIface) *GOutputStream {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetGOutputStreamIface()
}

func (o *GOutputStream) GetGOutputStreamIface() *GOutputStream { return o }

type GOutputStream struct{ gobject.GObjectCore }

// #endregion
