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

func cbool(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

type EventSequence struct{}
type Pixbuf struct{}

// #region AppLaunchContext

type AppLaunchContext struct{ gio.GAppLaunchContext }

func GTypeAppLaunchContext() gobject.GType       { return gdk_app_launch_context_get_type.Fn()() }
func NewAppLaunchContext() *AppLaunchContext     { return gdk_app_launch_context_new.Fn()() }
func (a *AppLaunchContext) GetDisplay() *Display { return gdk_app_launch_context_get_display.Fn()(a) }
func (a *AppLaunchContext) SetDesktop(desktop int32) {
	gdk_app_launch_context_set_desktop.Fn()(a, desktop)
}
func (a *AppLaunchContext) SetTimestamp(timestamp uint32) {
	gdk_app_launch_context_set_timestamp.Fn()(a, timestamp)
}
func (a *AppLaunchContext) SetIcon(icon *gio.GIcon) { gdk_app_launch_context_set_icon.Fn()(a, icon) }
func (a *AppLaunchContext) SetIconName(name string) {
	s := cc.NewString(name)
	defer s.Free()
	gdk_app_launch_context_set_icon_name.Fn()(a, s)
}

// #endregion

// #region ButtonEvent

type ButtonEvent struct{ Event }

func GTypeButtonEvent() gobject.GType    { return gdk_button_event_get_type.Fn()() }
func (e *ButtonEvent) GetButton() uint32 { return gdk_button_event_get_button.Fn()(e) }

// #endregion

// #region CICPParams

type CicpParams struct{ gobject.GObjectCore }

func GTypeCicpParams() gobject.GType { return gdk_cicp_params_get_type.Fn()() }

func NewCicpParams() *CicpParams                { return gdk_cicp_params_new.Fn()() }
func (p *CicpParams) GetColorPrimaries() uint32 { return gdk_cicp_params_get_color_primaries.Fn()(p) }
func (p *CicpParams) SetColorPrimaries(colorPrimaries uint32) {
	gdk_cicp_params_set_color_primaries.Fn()(p, colorPrimaries)
}
func (p *CicpParams) GetTransferFunction() uint32 {
	return gdk_cicp_params_get_transfer_function.Fn()(p)
}
func (p *CicpParams) SetTransferFunction(transferFunction uint32) {
	gdk_cicp_params_set_transfer_function.Fn()(p, transferFunction)
}
func (p *CicpParams) GetMatrixCoefficients() uint32 {
	return gdk_cicp_params_get_matrix_coefficients.Fn()(p)
}
func (p *CicpParams) SetMatrixCoefficients(matrixCoefficients uint32) {
	gdk_cicp_params_set_matrix_coefficients.Fn()(p, matrixCoefficients)
}
func (p *CicpParams) GetRange() CicpRange       { return gdk_cicp_params_get_range.Fn()(p) }
func (p *CicpParams) SetRange(range_ CicpRange) { gdk_cicp_params_set_range.Fn()(p, range_) }
func (p *CicpParams) BuildColorState() (*ColorState, error) {
	var err *error
	return gdk_cicp_params_build_color_state.Fn()(p, &err), *err
}

// #endregion

// #region Clipboard

type Clipboard struct{ gobject.GObjectCore }

func GTypeClipboard() gobject.GType { return gdk_clipboard_get_type.Fn()() }

func (cb *Clipboard) GetDisplay() *Display         { return gdk_clipboard_get_display.Fn()(cb) }
func (cb *Clipboard) GetFormats() *ContentFormats  { return gdk_clipboard_get_formats.Fn()(cb) }
func (cb *Clipboard) IsLocal() bool                { return gdk_clipboard_is_local.Fn()(cb) != 0 }
func (cb *Clipboard) GetContent() *ContentProvider { return gdk_clipboard_get_content.Fn()(cb) }

func (cb *Clipboard) StoreAsync(ioPriority int32, cancellable *gio.GCancellable, callback func(cb *Clipboard, res *gio.GAsyncResult)) {
	var cbk uptr
	if callback != nil {
		cbk = cc.Cbk(func(cb *Clipboard, res *gio.GAsyncResult, _ uptr) {
			callback(cb, res)
		})
		cb.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	}
	gdk_clipboard_store_async.Fn()(cb, ioPriority, cancellable, cbk, nil)
}
func (cb *Clipboard) StoreFinish(res *gio.GAsyncResult) error {
	var err *glib.GError
	ok := gdk_clipboard_store_finish.Fn()(cb, res, &err) != 0
	if ok {
		return nil
	}
	return err.TakeError()
}

func (cb *Clipboard) ReadAsync(mimeTypes []string, ioPriority int32, cancellable *gio.GCancellable, callback func(cb *Clipboard, res *gio.GAsyncResult)) {
	cMimeTypes := cc.NewStrings(mimeTypes)
	defer cMimeTypes.Free()
	var cbk uptr
	if cancellable != nil {
		cbk = cc.Cbk(func(cb *Clipboard, res *gio.GAsyncResult, _ uptr) {
			callback(cb, res)
		})
		cb.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	}

	gdk_clipboard_read_async.Fn()(cb, cMimeTypes, ioPriority, cancellable, cbk, nil)
}
func (cb *Clipboard) ReadFinish(res *gio.GAsyncResult) (mimeType string, stream *gio.GInputStream, err error) {
	var mt cc.String
	var gerr *glib.GError
	stream = gdk_clipboard_read_finish.Fn()(cb, res, &mt, &gerr)
	mimeType = mt.String()
	err = gerr.TakeError()
	return
}
func (cb *Clipboard) ReadValueAsync(typ gobject.GType, ioPriority int32, cancellable *gio.GCancellable, callback func(cb *Clipboard, res *gio.GAsyncResult)) {
	var cbk uptr
	if callback != nil {
		cbk = cc.Cbk(func(cb *Clipboard, res *gio.GAsyncResult, _ uptr) {
			callback(cb, res)
		})
		cb.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	}
	gdk_clipboard_read_value_async.Fn()(cb, typ, ioPriority, cancellable, cbk, nil)
}
func (cb *Clipboard) ReadValueFinish(res *gio.GAsyncResult) (val *gobject.GValue, err error) {
	var gerr *glib.GError
	val, err = gdk_clipboard_read_value_finish.Fn()(cb, res, &gerr), gerr.TakeError()
	return
}

func (cb *Clipboard) ReadTextureAsync(cancellable *gio.GCancellable, callback func(cb *Clipboard, res *gio.GAsyncResult)) {
	var cbk uptr
	if callback != nil {
		cbk = cc.Cbk(func(cb *Clipboard, res *gio.GAsyncResult, _ uptr) {
			callback(cb, res)
		})
		cb.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	}
	gdk_clipboard_read_texture_async.Fn()(cb, cancellable, cbk, nil)
}
func (cb *Clipboard) ReadTextureFinish(res *gio.GAsyncResult) (tex *Texture, err error) {
	var gerr *glib.GError
	tex, err = gdk_clipboard_read_texture_finish.Fn()(cb, res, &gerr), gerr.TakeError()
	return
}

func (cb *Clipboard) ReadTextAsync(cancellable *gio.GCancellable, callback func(cb *Clipboard, res *gio.GAsyncResult)) {
	var cbk uptr
	if callback != nil {
		cbk = cc.Cbk(func(cb *Clipboard, res *gio.GAsyncResult, _ uptr) {
			callback(cb, res)
		})
		cb.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	}
	gdk_clipboard_read_text_async.Fn()(cb, cancellable, cbk, nil)
}
func (cb *Clipboard) ReadTextFinish(res *gio.GAsyncResult) (txt string, err error) {
	var gerr *glib.GError
	txt, err = gdk_clipboard_read_text_finish.Fn()(cb, res, &gerr).TakeString(), gerr.TakeError()
	return
}

func (cb *Clipboard) SetContent(provider *ContentProvider) bool {
	return gdk_clipboard_set_content.Fn()(cb, provider) != 0
}
func (cb *Clipboard) Set(typ gobject.GType, args ...interface{}) {
	gdk_clipboard_set.FnVa()(cb, typ, args...)
}
func (cb *Clipboard) SetValue(value *gobject.GValue) { gdk_clipboard_set_value.Fn()(cb, value) }
func (cb *Clipboard) SetText(text string) {
	cText := cc.NewString(text)
	defer cText.Free()
	gdk_clipboard_set_text.Fn()(cb, cText)
}
func (cb *Clipboard) SetTexture(texture *Texture) { gdk_clipboard_set_texture.Fn()(cb, texture) }

func (cb *Clipboard) ConnectChanged(sig func(cb *Clipboard)) uint64 {
	return cb.SignalConnect("changed", sig, nil)
}

// #endregion

// #region ColorState

type ColorState struct{}

func GTypeColorState() gobject.GType { return gdk_color_state_get_type.Fn()() }

func (cs *ColorState) Ref() *ColorState       { return gdk_color_state_ref.Fn()(cs) }
func (cs *ColorState) Unref()                 { gdk_color_state_unref.Fn()(cs) }
func ColorStateGetSRGB() *ColorState          { return gdk_color_state_get_srgb.Fn()() }
func ColorStateGetSRGBLinear() *ColorState    { return gdk_color_state_get_srgb_linear.Fn()() }
func ColorStateGetRec2100Pq() *ColorState     { return gdk_color_state_get_rec2100_pq.Fn()() }
func ColorStateGetRec2100Linear() *ColorState { return gdk_color_state_get_rec2100_linear.Fn()() }
func ColorStateGetOklab() *ColorState         { return gdk_color_state_get_oklab.Fn()() }
func ColorStateGetOklch() *ColorState         { return gdk_color_state_get_oklch.Fn()() }
func (cs *ColorState) Equal(other *ColorState) bool {
	return gdk_color_state_equal.Fn()(cs, other) != 0
}
func (cs *ColorState) CreateCicpParams() *CicpParams {
	return gdk_color_state_create_cicp_params.Fn()(cs)
}

// #endregion

// #region ContentDeserializer

type ContentDeserializer struct {
	gobject.GObjectCore
	gio.GAsyncResult
}

func GTypeContentDeserializer() gobject.GType { return gdk_content_deserializer_get_type.Fn()() }

func (d *ContentDeserializer) GetMimeType() string {
	return gdk_content_deserializer_get_mime_type.Fn()(d).String()
}
func (d *ContentDeserializer) GetGType() gobject.GType {
	return gdk_content_deserializer_get_gtype.Fn()(d)
}
func (d *ContentDeserializer) GetValue() *gobject.GValue {
	return gdk_content_deserializer_get_value.Fn()(d)
}
func (d *ContentDeserializer) GetInputStream() *gio.GInputStream {
	return gdk_content_deserializer_get_input_stream.Fn()(d)
}
func (d *ContentDeserializer) GetPriority() int32 {
	return gdk_content_deserializer_get_priority.Fn()(d)
}
func (d *ContentDeserializer) GetCancellable() *gio.GCancellable {
	return gdk_content_deserializer_get_cancellable.Fn()(d)
}
func (d *ContentDeserializer) GetUserData() uptr {
	return gdk_content_deserializer_get_user_data.Fn()(d)
}
func (d *ContentDeserializer) SetTaskData(data uptr, notify func(uptr)) {
	var cb uptr
	if notify != nil {
		cb := cc.Cbk(notify)
		d.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cb) })
	}
	gdk_content_deserializer_set_task_data.Fn()(d, data, cb)
}
func (d *ContentDeserializer) GetTaskData() uptr {
	return gdk_content_deserializer_get_task_data.Fn()(d)
}
func (d *ContentDeserializer) ReturnSuccess() { gdk_content_deserializer_return_success.Fn()(d) }
func (d *ContentDeserializer) ReturnError(error *glib.GError) {
	gdk_content_deserializer_return_error.Fn()(d, error)
}

func (f *ContentFormats) UnionDeserializeGtypes() *ContentFormats {
	return gdk_content_formats_union_deserialize_gtypes.Fn()(f)
}
func (f *ContentFormats) UnionDeserializeMimeTypes() *ContentFormats {
	return gdk_content_formats_union_deserialize_mime_types.Fn()(f)
}

func ContentRegisterDeserializer(mimeType string, typ gobject.GType, deserialize func(*ContentDeserializer)) {
	cMimeType := cc.NewString(mimeType)
	defer cMimeType.Free()
	fn := cc.Cbk(deserialize)
	var des uptr
	des = cc.Cbk(func(uptr) {
		cc.CbkClose(fn)
		cc.CbkCloseLate(des)
	})
	gdk_content_register_deserializer.Fn()(cMimeType, typ, fn, nil, des)
}

func ContentDeserializeAsync(stream *gio.GInputStream, mimeType string, typ gobject.GType, ioPriority int32, cancellable *gio.GCancellable, callback func(res *gio.GAsyncResult)) {
	cMimeType := cc.NewString(mimeType)
	defer cMimeType.Free()
	var cbk uptr
	if callback != nil {
		cbk = cc.Cbk(func(res *gio.GAsyncResult, _ uptr) {
			callback(res)
		})
		stream.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	}

	gdk_content_deserialize_async.Fn()(stream, cMimeType, typ, ioPriority, cancellable, cbk, nil)
}
func ContentDeserializeFinish(result *gio.GAsyncResult, value *gobject.GValue) error {
	var err *glib.GError
	ok := gdk_content_deserialize_finish.Fn()(result, value, &err) != 0
	if ok {
		return nil
	}
	return err.TakeError()
}

// #endregion

// #region ContentFormats

func GTypeContentFormats() gobject.GType        { return gdk_content_formats_get_type.Fn()() }
func GTypeContentFormatsBuilder() gobject.GType { return gdk_content_formats_builder_get_type.Fn()() }
func GTypeFileList() gobject.GType              { return gdk_file_list_get_type.Fn()() }

func InternMimeType(s string) string {
	cs := cc.NewString(s)
	defer cs.Free()
	return gdk_intern_mime_type.Fn()(cs).String()
}

type ContentFormats struct{}

func ContentFormatsNew(mimeTypes []string) *ContentFormats {
	cs := cc.NewStrings(mimeTypes)
	defer cs.Free()
	return gdk_content_formats_new.Fn()(cs, uint32(len(mimeTypes)))
}
func ContentFormatsNewForGtype(typ gobject.GType) *ContentFormats {
	return gdk_content_formats_new_for_gtype.Fn()(typ)
}
func ContentFormatsParse(s string) *ContentFormats {
	cs := cc.NewString(s)
	defer cs.Free()
	return gdk_content_formats_parse.Fn()(cs)
}
func (f *ContentFormats) Ref() *ContentFormats    { return gdk_content_formats_ref.Fn()(f) }
func (f *ContentFormats) Unref()                  { gdk_content_formats_unref.Fn()(f) }
func (f *ContentFormats) Print(str *glib.GString) { gdk_content_formats_print.Fn()(f, str) }
func (f *ContentFormats) ToString() string        { return gdk_content_formats_to_string.Fn()(f).TakeString() }
func (f *ContentFormats) GetGtypes() (types []gobject.GType, nGtypes uint64) {
	var n uint64
	ptr := gdk_content_formats_get_gtypes.Fn()(f, &n)
	if ptr != nil && n > 0 {
		types = unsafe.Slice(ptr, n)
	}
	return types, n
}
func (f *ContentFormats) GetMimeTypes() (mimes []string) {
	var n uint64
	cs := gdk_content_formats_get_mime_types.Fn()(f, &n)
	if cs != 0 && n > 0 {
		mimes = cs.Strings(n)
	}
	return mimes
}
func (f *ContentFormats) Union(second *ContentFormats) *ContentFormats {
	return gdk_content_formats_union.Fn()(f, second)
}
func (f *ContentFormats) Match(second *ContentFormats) bool {
	return gdk_content_formats_match.Fn()(f, second) != 0
}
func (f *ContentFormats) MatchGtype(second *ContentFormats) gobject.GType {
	return gdk_content_formats_match_gtype.Fn()(f, second)
}
func (f *ContentFormats) MatchMimeType(second *ContentFormats) string {
	return gdk_content_formats_match_mime_type.Fn()(f, second).String()
}
func (f *ContentFormats) ContainGtype(typ gobject.GType) bool {
	return gdk_content_formats_contain_gtype.Fn()(f, typ) != 0
}
func (f *ContentFormats) ContainMimeType(mimeType string) bool {
	cs := cc.NewString(mimeType)
	defer cs.Free()
	return gdk_content_formats_contain_mime_type.Fn()(f, cs) != 0
}
func (f *ContentFormats) IsEmpty() bool { return gdk_content_formats_is_empty.Fn()(f) != 0 }

type ContentFormatsBuilder struct{}

func ContentFormatsBuilderNew() *ContentFormatsBuilder { return gdk_content_formats_builder_new.Fn()() }
func (b *ContentFormatsBuilder) Ref() *ContentFormatsBuilder {
	return gdk_content_formats_builder_ref.Fn()(b)
}
func (b *ContentFormatsBuilder) Unref() { gdk_content_formats_builder_unref.Fn()(b) }
func (b *ContentFormatsBuilder) FreeToFormats() *ContentFormats {
	return gdk_content_formats_builder_free_to_formats.Fn()(b)
}
func (b *ContentFormatsBuilder) ToFormats() *ContentFormats {
	return gdk_content_formats_builder_to_formats.Fn()(b)
}
func (b *ContentFormatsBuilder) AddFormats(formats *ContentFormats) {
	gdk_content_formats_builder_add_formats.Fn()(b, formats)
}
func (b *ContentFormatsBuilder) AddMimeType(mimeType string) {
	cs := cc.NewString(mimeType)
	defer cs.Free()
	gdk_content_formats_builder_add_mime_type.Fn()(b, cs)
}
func (b *ContentFormatsBuilder) AddGtype(typ gobject.GType) {
	gdk_content_formats_builder_add_gtype.Fn()(b, typ)
}

type FileList struct{}

func (fl *FileList) GetFiles() *glib.GSList[gio.GFile] {
	return (*glib.GSList[gio.GFile])(gdk_file_list_get_files.Fn()(fl))
}
func FileListNewFromList(files *glib.GSList[gio.GFile]) *FileList {
	return gdk_file_list_new_from_list.Fn()(uptr(files))
}
func FileListNewFromArray(files []*gio.GFile) *FileList {
	file, nFiles := (**gio.GFile)(nil), uint64(len(files))
	if nFiles > 0 {
		file = &files[0]
	}
	return gdk_file_list_new_from_array.Fn()(file, nFiles)
}

// #endregion

// #region ContentProvider

type ContentProviderObj struct {
	Parent gobject.GObject
}

type ContentProviderClass struct {
	Parent gobject.GObjectClass

	ContentChanged cc.Func // void (* content_changed) (GdkContentProvider *provider);

	AttachClipboard cc.Func // void (* attach_clipboard) (GdkContentProvider *provider, GdkClipboard *clipboard);
	DetachClipboard cc.Func // void (* detach_clipboard) (GdkContentProvider *provider,GdkClipboard *clipboard);

	RefFormats          cc.Func // GdkContentFormats * (* ref_formats) (GdkContentProvider *provider);
	RefStorableFormats  cc.Func // GdkContentFormats * (* ref_storable_formats) (GdkContentProvider *provider);
	WriteMimeTypeAsync  cc.Func // void (* write_mime_type_async) (GdkContentProvider *provider, const char *mime_type, GOutputStream *stream, int io_priority, GCancellable *cancellable, GAsyncReadyCallback callback, gpointer user_data);
	WriteMimeTypeFinish cc.Func // gboolean (* write_mime_type_finish) (GdkContentProvider *provider, GAsyncResult *result, GError **error);
	GetValue            cc.Func // gboolean (* get_value) (GdkContentProvider *provider, GValue *value, GError **error);

	_ [8]uptr
}

type ContentProvider struct{ gobject.GObjectCore }

type ContentProviderIface interface{ GetContentProviderIface() *ContentProvider }

func GetContentProviderIface(w ContentProviderIface) *ContentProvider {
	if anyptr(w) == nil {
		return nil
	}
	return w.GetContentProviderIface()
}
func (p *ContentProvider) GetContentProviderIface() *ContentProvider { return p }

func GTypeContentProvider() gobject.GType { return gdk_content_provider_get_type.Fn()() }

func (provider *ContentProvider) RefFormats() *ContentFormats {
	return gdk_content_provider_ref_formats.Fn()(provider)
}
func (provider *ContentProvider) RefStorableFormats() *ContentFormats {
	return gdk_content_provider_ref_storable_formats.Fn()(provider)
}
func (provider *ContentProvider) ContentChanged() {
	gdk_content_provider_content_changed.Fn()(provider)
}
func (provider *ContentProvider) WriteMimeTypeAsync(mime_type string, stream *gio.GOutputStream, io_priority int32, cancellable *gio.GCancellable, callback func(provider *ContentProvider, res *gio.GAsyncResult)) {
	var cbk uptr
	if callback != nil {
		cbk = cc.Cbk(func(provider *ContentProvider, res *gio.GAsyncResult, _ uptr) {
			callback(provider, res)
		})
		provider.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	}

	cMimeType := cc.NewString(mime_type)
	defer cMimeType.Free()
	gdk_content_provider_write_mime_type_async.Fn()(provider, cMimeType, stream, io_priority, cancellable, cbk, nil)
}
func (provider *ContentProvider) WriteMimeTypeFinish(res *gio.GAsyncResult) error {
	var err *glib.GError
	ok := gdk_content_provider_write_mime_type_finish.Fn()(provider, res, &err) != 0
	if ok {
		return nil
	}
	return err.TakeError()
}
func (provider *ContentProvider) GetValue() (value *gobject.GValue, err error) {
	value = new(gobject.GValue)
	var gerr *glib.GError
	ok := gdk_content_provider_get_value.Fn()(provider, value, &gerr) != 0
	if ok {
		err = gerr.TakeError()
		return
	}
	return
}
func (cp *ContentProvider) ConnectContentchanged(sig func(cp *ContentProvider)) uint64 {
	return cp.SignalConnect("content-changed", sig, nil)
}
func NewContentProviderForValue(value *gobject.GValue) *ContentProvider {
	return gdk_content_provider_new_for_value.Fn()(value)
}
func NewContentProviderTyped(typ gobject.GType, args ...interface{}) *ContentProvider {
	return gdk_content_provider_new_typed.FnVa()(typ, args...)
}
func NewContentProviderUnion(providers []*ContentProvider) *ContentProvider {
	p, n := (**ContentProvider)(nil), uint64(len(providers))
	if n > 0 {
		p = &providers[0]
	}
	return gdk_content_provider_new_union.Fn()(p, n)
}
func NewContentProviderForBytes(mimeType string, bytes *glib.GBytes) *ContentProvider {
	cMimeType := cc.NewString(mimeType)
	defer cMimeType.Free()
	return gdk_content_provider_new_for_bytes.Fn()(cMimeType, bytes)
}

// #endregion

// #region ContentSerializer

type ContentSerializer struct {
	gobject.GObjectCore
	gio.GAsyncResult
}

func GTypeContentSerializer() gobject.GType { return gdk_content_serializer_get_type.Fn()() }

func (s *ContentSerializer) GetMimeType() string {
	return gdk_content_serializer_get_mime_type.Fn()(s).String()
}
func (s *ContentSerializer) GetGType() gobject.GType { return gdk_content_serializer_get_gtype.Fn()(s) }
func (s *ContentSerializer) GetValue() *gobject.GValue {
	return gdk_content_serializer_get_value.Fn()(s)
}
func (s *ContentSerializer) GetOutputStream() *gio.GOutputStream {
	return gdk_content_serializer_get_output_stream.Fn()(s)
}
func (s *ContentSerializer) GetPriority() int32 { return gdk_content_serializer_get_priority.Fn()(s) }
func (s *ContentSerializer) GetCancellable() *gio.GCancellable {
	return gdk_content_serializer_get_cancellable.Fn()(s)
}
func (s *ContentSerializer) GetUserData() uptr { return gdk_content_serializer_get_user_data.Fn()(s) }
func (s *ContentSerializer) SetTaskData(data uptr, notify func(uptr)) {
	var n uptr
	if notify != nil {
		n = cc.Cbk(notify)
		s.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(n) })
	}
	gdk_content_serializer_set_task_data.Fn()(s, data, n)
}
func (s *ContentSerializer) GetTaskData() uptr { return gdk_content_serializer_get_task_data.Fn()(s) }
func (s *ContentSerializer) ReturnSuccess()    { gdk_content_serializer_return_success.Fn()(s) }
func (s *ContentSerializer) ReturnError(error *glib.GError) {
	gdk_content_serializer_return_error.Fn()(s, error)
}
func (f *ContentFormats) UnionSerializeGtypes() *ContentFormats {
	return gdk_content_formats_union_serialize_gtypes.Fn()(f)
}
func (f *ContentFormats) UnionSerializeMimeTypes() *ContentFormats {
	return gdk_content_formats_union_serialize_mime_types.Fn()(f)
}
func ContentRegisterSerializer(typ gobject.GType, mimeType string, serialize func(serial *ContentSerializer)) {
	cMimeType := cc.NewString(mimeType)
	defer cMimeType.Free()
	fn := cc.Cbk(serialize)
	var notify uptr
	if serialize != nil {
		notify = cc.Cbk(func(_ uptr) {
			cc.CbkClose(fn)
			cc.CbkCloseLate(notify)
		})
	}
	gdk_content_register_serializer.Fn()(typ, cMimeType, fn, nil, notify)
}
func ContentSerializeAsync(stream *gio.GOutputStream, mimeType string, value *gobject.GValue, ioPriority int32, cancellable *gio.GCancellable, callback func(res *gio.GAsyncResult)) {
	var cbk uptr
	if callback != nil {
		cbk = cc.Cbk(func(res *gio.GAsyncResult, _ uptr) {
			callback(res)
		})
		stream.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	}

	cMimeType := cc.NewString(mimeType)
	defer cMimeType.Free()
	gdk_content_serialize_async.Fn()(stream, cMimeType, value, ioPriority, cancellable, cbk, nil)
}
func ContentSerializeFinish(result *gio.GAsyncResult) error {
	var err *glib.GError
	ok := gdk_content_serialize_finish.Fn()(result, &err) != 0
	if ok {
		return nil
	}
	return err.TakeError()
}

// #endregion

// #region CrossingEvent

type CrossingEvent struct{ Event }

func GTypeCrossingEvent() gobject.GType                { return gdk_crossing_event_get_type.Fn()() }
func (e *CrossingEvent) GetCrossingMode() CrossingMode { return gdk_crossing_event_get_mode.Fn()(e) }
func (e *CrossingEvent) GetCrossingDetail() NotifyType { return gdk_crossing_event_get_detail.Fn()(e) }
func (e *CrossingEvent) GetCrossingFocus() bool        { return gdk_crossing_event_get_focus.Fn()(e) != 0 }

// #endregion

// #region Cursor

type Cursor struct{ gobject.GObject }

func GTypeCursor() gobject.GType { return gdk_cursor_get_type.Fn()() }

func NewCursorFromTexture(texture *Texture, hotspotX, hotspotY int32, fallback *Cursor) *Cursor {
	return gdk_cursor_new_from_texture.Fn()(texture, hotspotX, hotspotY, fallback)
}

func NewCursorFromName(name string, fallback *Cursor) *Cursor {
	cstr := cc.NewString(name)
	defer cstr.Free()
	return gdk_cursor_new_from_name.Fn()(cstr, fallback)
}
func NewCursorFromCallback(
	callback func(cursor *Cursor, cursorSize int32, scale float64, ouW, outH, hotspotX, hotspotY *int32, _ uptr), destroy uintptr, fallback *Cursor) *Cursor {
	cbk := cc.Cbk(callback)

	var desPtr uptr
	desPtr = cc.Cbk(func() {
		if callback != nil {
			cc.CbkClose(cbk)
		}
		cc.CbkCloseLate(desPtr)
	})
	return gdk_cursor_new_from_callback.Fn()(cbk, nil, desPtr, fallback)
}
func (c *Cursor) GetFallback() *Cursor { return gdk_cursor_get_fallback.Fn()(c) }
func (c *Cursor) GetName() string      { return gdk_cursor_get_name.Fn()(c).String() }
func (c *Cursor) GetTexture() *Texture { return gdk_cursor_get_texture.Fn()(c) }
func (c *Cursor) GetHotspotX() int32   { return gdk_cursor_get_hotspot_x.Fn()(c) }
func (c *Cursor) GetHotspotY() int32   { return gdk_cursor_get_hotspot_y.Fn()(c) }

// #endregion

// #region DeleteEvent

type DeleteEvent struct{ Event }

func GTypeDeleteEvent() gobject.GType { return gdk_delete_event_get_type.Fn()() }

// #endregion

// #region Device

type Device struct{ gobject.GObjectCore }

func GTypeDevice() gobject.GType { return gdk_device_get_type.Fn()() }

func (d *Device) GetName() string                { return gdk_device_get_name.Fn()(d).String() }
func (d *Device) GetVendorId() string            { return gdk_device_get_vendor_id.Fn()(d).String() }
func (d *Device) GetProductId() string           { return gdk_device_get_product_id.Fn()(d).String() }
func (d *Device) GetDisplay() *Display           { return gdk_device_get_display.Fn()(d) }
func (d *Device) GetSeat() *Seat                 { return gdk_device_get_seat.Fn()(d) }
func (d *Device) GetDeviceTool() *DeviceTool     { return gdk_device_get_device_tool.Fn()(d) }
func (d *Device) GetSource() InputSource         { return gdk_device_get_source.Fn()(d) }
func (d *Device) GetHasCursor() bool             { return gdk_device_get_has_cursor.Fn()(d) != 0 }
func (d *Device) GetNumTouches() uint32          { return gdk_device_get_num_touches.Fn()(d) }
func (d *Device) GetModifierState() ModifierType { return gdk_device_get_modifier_state.Fn()(d) }
func (d *Device) GetDirection() pango.Direction  { return gdk_device_get_direction.Fn()(d) }
func (d *Device) HasBidiLayouts() bool           { return gdk_device_has_bidi_layouts.Fn()(d) != 0 }
func (d *Device) GetCapsLockState() bool         { return gdk_device_get_caps_lock_state.Fn()(d) != 0 }
func (d *Device) GetNumLockState() bool          { return gdk_device_get_num_lock_state.Fn()(d) != 0 }
func (d *Device) GetScrollLockState() bool       { return gdk_device_get_scroll_lock_state.Fn()(d) != 0 }
func (d *Device) GetSurfaceAtPosition() (surface *Surface, x float64, y float64) {
	surface = gdk_device_get_surface_at_position.Fn()(d, &x, &y)
	return
}
func (d *Device) GetTimestamp() uint32        { return gdk_device_get_timestamp.Fn()(d) }
func (d *Device) GetActiveLayoutIndex() int32 { return gdk_device_get_active_layout_index.Fn()(d) }
func (d *Device) GetLayoutNames() []string    { return gdk_device_get_layout_names.Fn()(d).Strings() }

func (d *Device) ConnectChanged(sig func(d *Device)) uint64 {
	return d.SignalConnect("changed", sig, nil)
}
func (d *Device) ConnectToolChanged(sig func(d *Device, tool *DeviceTool)) uint64 {
	return d.SignalConnect("tool-changed", sig, nil)
}

type TimeCoord struct {
	Time  uint32
	Flags AxisFlags
	Axes  [12]float64
}

// #endregion

// #region DevicePad

type DevicePad struct{}

func GTypeDevicePad() gobject.GType { return gdk_device_pad_get_type.Fn()() }

func (pad *DevicePad) GetNGroups() int32 { return gdk_device_pad_get_n_groups.Fn()(pad) }
func (pad *DevicePad) GetGroupNModes(groupIdx int32) int32 {
	return gdk_device_pad_get_group_n_modes.Fn()(pad, groupIdx)
}
func (pad *DevicePad) GetNFeatures(feature DevicePadFeature) int32 {
	return gdk_device_pad_get_n_features.Fn()(pad, feature)
}
func (pad *DevicePad) GetFeatureGroup(feature DevicePadFeature, featureIdx int32) int32 {
	return gdk_device_pad_get_feature_group.Fn()(pad, feature, featureIdx)
}

// #endregion

// #region DeviceTool

type DeviceTool struct{ gobject.GObjectCore }

func GTypeDeviceTool() gobject.GType { return gdk_device_tool_get_type.Fn()() }

func (tool *DeviceTool) GetSerial() uint64           { return gdk_device_tool_get_serial.Fn()(tool) }
func (tool *DeviceTool) GetHardwareId() uint64       { return gdk_device_tool_get_hardware_id.Fn()(tool) }
func (tool *DeviceTool) GetToolType() DeviceToolType { return gdk_device_tool_get_tool_type.Fn()(tool) }
func (tool *DeviceTool) GetAxes() AxisFlags          { return gdk_device_tool_get_axes.Fn()(tool) }

// #endregion

// #region Display

type Display struct{ gobject.GObjectCore }

func GTypeDisplay() gobject.GType { return gdk_display_get_type.Fn()() }

func OpenDisplay(displayName string) *Display {
	cName := cc.NewString(displayName)
	defer cName.Free()
	return gdk_display_open.Fn()(cName)
}

func (d *Display) GetName() string { return gdk_display_get_name.Fn()(d).String() }
func (d *Display) DeviceIsGrabbed(device *Device) bool {
	return gdk_display_device_is_grabbed.Fn()(d, device) != 0
}
func (d *Display) Beep()                     { gdk_display_beep.Fn()(d) }
func (d *Display) Sync()                     { gdk_display_sync.Fn()(d) }
func (d *Display) Flush()                    { gdk_display_flush.Fn()(d) }
func (d *Display) Close()                    { gdk_display_close.Fn()(d) }
func (d *Display) IsClosed() bool            { return gdk_display_is_closed.Fn()(d) != 0 }
func (d *Display) IsComposited() bool        { return gdk_display_is_composited.Fn()(d) != 0 }
func (d *Display) IsRGBA() bool              { return gdk_display_is_rgba.Fn()(d) != 0 }
func (d *Display) SupportsShadowWidth() bool { return gdk_display_supports_shadow_width.Fn()(d) != 0 }
func (d *Display) SupportsInputShapes() bool { return gdk_display_supports_input_shapes.Fn()(d) != 0 }
func (d *Display) PrepareGL() error {
	var err *glib.GError
	ok := gdk_display_prepare_gl.Fn()(d, &err) != 0
	if ok {
		return nil
	}
	return err.TakeError()
}
func (d *Display) CreateGLContext() (ctx *GLContext, err error) {
	var gerr *glib.GError
	ctx, err = gdk_display_create_gl_context.Fn()(d, &gerr), gerr.TakeError()
	return
}
func GetDisplayDefault() *Display                  { return gdk_display_get_default.Fn()() }
func (d *Display) GetClipboard() *Clipboard        { return gdk_display_get_clipboard.Fn()(d) }
func (d *Display) GetPrimaryClipboard() *Clipboard { return gdk_display_get_primary_clipboard.Fn()(d) }
func (d *Display) GetAppLaunchContext() *AppLaunchContext {
	return gdk_display_get_app_launch_context.Fn()(d)
}
func (d *Display) GetDefaultSeat() *Seat        { return gdk_display_get_default_seat.Fn()(d) }
func (d *Display) ListSeats() *glib.GList[Seat] { return gdk_display_list_seats.Fn()(d) }
func (d *Display) GetMonitors() *gio.GListModel[Monitor] {
	return (*gio.GListModel[Monitor])(gdk_display_get_monitors.Fn()(d))
}
func (d *Display) GetMonitorAtSurface(surface *Surface) *Monitor {
	return gdk_display_get_monitor_at_surface.Fn()(d, surface)
}

type KeymapKey struct {
	KeyCode uint32
	Group   int32
	Level   int32
}

func (d *Display) MapKeyval(keyval uint32) (keys []KeymapKey, ok bool) {
	var cKeys *KeymapKey
	var nKeys int32
	res := gdk_display_map_keyval.Fn()(d, keyval, &cKeys, &nKeys)
	if res == 0 || nKeys <= 0 {
		return nil, false
	}

	keys = make([]KeymapKey, nKeys)
	copy(keys, unsafe.Slice(cKeys, nKeys))
	defer glib.GFree(cKeys)
	return
}

func (d *Display) MapKeycode(keycode uint32) (keys []KeymapKey, keyvals []uint32, ok bool) {
	var cKeys *KeymapKey
	var cKeyvals *uint32
	var nEntries int32
	res := gdk_display_map_keycode.Fn()(d, keycode, &cKeys, &cKeyvals, &nEntries)
	if res == 0 || nEntries <= 0 {
		return nil, nil, false
	}
	defer func() {
		glib.GFree(cKeys)
		glib.GFree(cKeyvals)
	}()

	keys = make([]KeymapKey, nEntries)
	copy(keys, unsafe.Slice(cKeys, nEntries))

	keyvals = make([]uint32, nEntries)
	copy(keyvals, unsafe.Slice(cKeyvals, nEntries))
	return keys, keyvals, true
}

func (d *Display) TranslateKey(keycode uint32, state ModifierType, group int32) (keyval uint32, effectiveGroup int32, level int32, consumed ModifierType, ok bool) {
	ok = gdk_display_translate_key.Fn()(d, keycode, state, group, &keyval, &effectiveGroup, &level, &consumed) != 0
	return
}

func (d *Display) GetSetting(name string) (ok bool, value *gobject.GValue) {
	cName := cc.NewString(name)
	defer cName.Free()
	v := gobject.GValue{}
	ok = gdk_display_get_setting.Fn()(d, cName, &v) != 0
	value = &v
	return
}

func (d *Display) GetDmabufFormats() *DmabufFormats { return gdk_display_get_dmabuf_formats.Fn()(d) }

func (d *Display) ConnectClosed(sig func(d *Display, isErr int32)) uint64 {
	return d.SignalConnect("closed", sig, nil)
}
func (d *Display) ConnectOpened(sig func(d *Display)) uint64 {
	return d.SignalConnect("opened", sig, nil)
}
func (d *Display) ConnectSeatAdded(sig func(d *Display, seat *Seat)) uint64 {
	return d.SignalConnect("seat-added", sig, nil)
}
func (d *Display) ConnectSeatRemoved(sig func(d *Display, seat *Seat)) uint64 {
	return d.SignalConnect("seat-removed", sig, nil)
}
func (d *Display) ConnectSettingChanged(sig func(d *Display, setting string)) uint64 {
	return d.SignalConnect("setting-changed", sig, nil)
}

// #endregion

// #region DisplayManager

type DisplayManager struct {
	gobject.GObjectCore
}

func GTypeDisplayManager() gobject.GType { return gdk_display_manager_get_type.Fn()() }
func DisplayManagerGet() *DisplayManager { return gdk_display_manager_get.Fn()() }
func (d *DisplayManager) GetDefaultDisplay() *Display {
	return gdk_display_manager_get_default_display.Fn()(d)
}
func (d *DisplayManager) SetDefaultDisplay(display *Display) {
	gdk_display_manager_set_default_display.Fn()(d, display)
}
func (d *DisplayManager) ListDisplays() *glib.GSList[Display] {
	return (*glib.GSList[Display])(gdk_display_manager_list_displays.Fn()(d))
}
func (d *DisplayManager) OpenDisplay(name string) *Display {
	cName := cc.NewString(name)
	defer cName.Free()
	return gdk_display_manager_open_display.Fn()(d, cName)
}
func SetAllowedBackends(backends string) {
	cBackends := cc.NewString(backends)
	defer cBackends.Free()
	gdk_set_allowed_backends.Fn()(cBackends)
}
func (d *DisplayManager) ConnectDisplayOpened(sig func(d *DisplayManager, display *Display)) uint64 {
	return d.SignalConnect("display-opened", sig, nil)
}

// #endregion

// #region DmabufFormats

type DmabufFormats struct{}

func GTypeDmabufFormats() gobject.GType      { return gdk_dmabuf_formats_get_type.Fn()() }
func (f *DmabufFormats) Ref() *DmabufFormats { return gdk_dmabuf_formats_ref.Fn()(f) }
func (f *DmabufFormats) Unref()              { gdk_dmabuf_formats_unref.Fn()(f) }
func (f *DmabufFormats) GetNFormats() uint64 { return gdk_dmabuf_formats_get_n_formats.Fn()(f) }
func (f *DmabufFormats) GetFormat(idx uint64) (fourcc uint32, modifier uint64) {
	gdk_dmabuf_formats_get_format.Fn()(f, idx, &fourcc, &modifier)
	return
}
func (f *DmabufFormats) Contains(fourcc uint32, modifier uint64) bool {
	return gdk_dmabuf_formats_contains.Fn()(f, fourcc, modifier) != 0
}
func (f *DmabufFormats) Equal(other *DmabufFormats) bool {
	return gdk_dmabuf_formats_equal.Fn()(f, other) != 0
}

// #endregion

// #region DmabufTexture

type DmabufTexture struct{ Texture }

func GTypeDmabufTexture() gobject.GType { return gdk_dmabuf_texture_get_type.Fn()() }
func DmabufErrorQuark() glib.GQuark     { return gdk_dmabuf_error_quark.Fn()() }

// #endregion

// #region DmabufTextureBuilder

type DmabufTextureBuilder struct {
	gobject.GObjectCore
}

func GTypeDmabufTextureBuilder() gobject.GType       { return gdk_dmabuf_texture_builder_get_type.Fn()() }
func NewDmabufTextureBuilder() *DmabufTextureBuilder { return gdk_dmabuf_texture_builder_new.Fn()() }
func (b *DmabufTextureBuilder) GetDisplay() *Display {
	return gdk_dmabuf_texture_builder_get_display.Fn()(b)
}
func (b *DmabufTextureBuilder) SetDisplay(display *Display) {
	gdk_dmabuf_texture_builder_set_display.Fn()(b, display)
}
func (b *DmabufTextureBuilder) GetWidth() uint32 { return gdk_dmabuf_texture_builder_get_width.Fn()(b) }
func (b *DmabufTextureBuilder) SetWidth(width uint32) {
	gdk_dmabuf_texture_builder_set_width.Fn()(b, width)
}
func (b *DmabufTextureBuilder) GetHeight() uint32 {
	return gdk_dmabuf_texture_builder_get_height.Fn()(b)
}
func (b *DmabufTextureBuilder) SetHeight(height uint32) {
	gdk_dmabuf_texture_builder_set_height.Fn()(b, height)
}
func (b *DmabufTextureBuilder) GetFourcc() uint32 {
	return gdk_dmabuf_texture_builder_get_fourcc.Fn()(b)
}
func (b *DmabufTextureBuilder) SetFourcc(fourcc uint32) {
	gdk_dmabuf_texture_builder_set_fourcc.Fn()(b, fourcc)
}
func (b *DmabufTextureBuilder) GetModifier() uint64 {
	return gdk_dmabuf_texture_builder_get_modifier.Fn()(b)
}
func (b *DmabufTextureBuilder) SetModifier(modifier uint64) {
	gdk_dmabuf_texture_builder_set_modifier.Fn()(b, modifier)
}
func (b *DmabufTextureBuilder) GetPremultiplied() bool {
	return gdk_dmabuf_texture_builder_get_premultiplied.Fn()(b) != 0
}
func (b *DmabufTextureBuilder) SetPremultiplied(premultiplied bool) {
	gdk_dmabuf_texture_builder_set_premultiplied.Fn()(b, cbool(premultiplied))
}
func (b *DmabufTextureBuilder) GetNPlanes() uint32 {
	return gdk_dmabuf_texture_builder_get_n_planes.Fn()(b)
}
func (b *DmabufTextureBuilder) SetNPlanes(nPlanes uint32) {
	gdk_dmabuf_texture_builder_set_n_planes.Fn()(b, nPlanes)
}
func (b *DmabufTextureBuilder) GetFd(plane uint32) int32 {
	return gdk_dmabuf_texture_builder_get_fd.Fn()(b, plane)
}
func (b *DmabufTextureBuilder) SetFd(plane uint32, fd int32) {
	gdk_dmabuf_texture_builder_set_fd.Fn()(b, plane, fd)
}
func (b *DmabufTextureBuilder) GetStride(plane uint32) uint32 {
	return gdk_dmabuf_texture_builder_get_stride.Fn()(b, plane)
}
func (b *DmabufTextureBuilder) SetStride(plane uint32, stride uint32) {
	gdk_dmabuf_texture_builder_set_stride.Fn()(b, plane, stride)
}
func (b *DmabufTextureBuilder) GetOffset(plane uint32) uint32 {
	return gdk_dmabuf_texture_builder_get_offset.Fn()(b, plane)
}
func (b *DmabufTextureBuilder) SetOffset(plane uint32, offset uint32) {
	gdk_dmabuf_texture_builder_set_offset.Fn()(b, plane, offset)
}
func (b *DmabufTextureBuilder) GetColorState() *ColorState {
	return gdk_dmabuf_texture_builder_get_color_state.Fn()(b)
}
func (b *DmabufTextureBuilder) SetColorState(colorState *ColorState) {
	gdk_dmabuf_texture_builder_set_color_state.Fn()(b, colorState)
}
func (b *DmabufTextureBuilder) GetUpdateTexture() *Texture {
	return gdk_dmabuf_texture_builder_get_update_texture.Fn()(b)
}
func (b *DmabufTextureBuilder) SetUpdateTexture(texture *Texture) {
	gdk_dmabuf_texture_builder_set_update_texture.Fn()(b, texture)
}
func (b *DmabufTextureBuilder) GetUpdateRegion() *cairo.Region {
	return gdk_dmabuf_texture_builder_get_update_region.Fn()(b)
}
func (b *DmabufTextureBuilder) SetUpdateRegion(region *cairo.Region) {
	gdk_dmabuf_texture_builder_set_update_region.Fn()(b, region)
}
func (b *DmabufTextureBuilder) Build(destroy func()) (tex *Texture, err error) {
	var des uptr
	if destroy != nil {
		des = cc.Cbk(func(_ uptr) { destroy() })
		b.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(des) })
	}
	var gerr *glib.GError
	tex, err = gdk_dmabuf_texture_builder_build.Fn()(b, des, nil, &gerr), gerr.TakeError()
	return
}

// #endregion

// #region DNDEvent

type DNDEvent struct{ Event }

func GTypeDNDEvent() gobject.GType    { return gdk_dnd_event_get_type.Fn()() }
func (e *DNDEvent) GetDNDDrop() *Drop { return gdk_dnd_event_get_drop.Fn()(e) }

// #endregion

// #region Drag

type Drag struct{ gobject.GObjectCore }

func GTypeDrag() gobject.GType                  { return gdk_drag_get_type.Fn()() }
func (d *Drag) GetDisplay() *Display            { return gdk_drag_get_display.Fn()(d) }
func (d *Drag) GetDevice() *Device              { return gdk_drag_get_device.Fn()(d) }
func (d *Drag) GetFormats() *ContentFormats     { return gdk_drag_get_formats.Fn()(d) }
func (d *Drag) GetActions() DragAction          { return gdk_drag_get_actions.Fn()(d) }
func (d *Drag) GetSelectedAction() DragAction   { return gdk_drag_get_selected_action.Fn()(d) }
func DragActionIsUnique(action DragAction) bool { return gdk_drag_action_is_unique.Fn()(action) != 0 }
func DragBegin(surface *Surface, device *Device, content *ContentProvider, actions DragAction, dx, dy float64) *Drag {
	return gdk_drag_begin.Fn()(surface, device, content, actions, dx, dy)
}
func (d *Drag) DropDone(success bool)        { gdk_drag_drop_done.Fn()(d, cbool(success)) }
func (d *Drag) GetDragSurface() *Surface     { return gdk_drag_get_drag_surface.Fn()(d) }
func (d *Drag) SetHotspot(hotX, hotY int32)  { gdk_drag_set_hotspot.Fn()(d, hotX, hotY) }
func (d *Drag) GetContent() *ContentProvider { return gdk_drag_get_content.Fn()(d) }
func (d *Drag) GetSurface() *Surface         { return gdk_drag_get_surface.Fn()(d) }

func (d *Drag) ConnectCancel(sig func(d *Drag, reason DragCancelReason)) uint64 {
	return d.SignalConnect("cancel", sig, nil)
}
func (d *Drag) ConnectDndFinished(sig func(d *Drag)) uint64 {
	return d.SignalConnect("dnd-finished", sig, nil)
}
func (d *Drag) ConnectDropPerformed(sig func(d *Drag)) uint64 {
	return d.SignalConnect("drop-performed", sig, nil)
}

// #endregion

// #region DragSurface

type DragSurface struct{ Surface }

func GTypeDragSurface() gobject.GType { return gdk_drag_surface_get_type.Fn()() }

type DragSurfaceIface interface{ GetDragSurfaceIface() *DragSurface }

func GetDragSurfaceIface(w DragSurfaceIface) *DragSurface {
	if anyptr(w) == nil {
		return nil
	}
	return w.GetDragSurfaceIface()
}
func (ds *DragSurface) GetDragSurfaceIface() *DragSurface { return ds }

func (ds *DragSurface) Present(width, height int32) bool {
	return gdk_drag_surface_present.Fn()(ds, width, height) != 0
}

func (ds *DragSurface) ConnectComputeSize(sig func(d *DragSurface, size *DragSurfaceSize)) uint64 {
	return ds.SignalConnect("compute-size", sig, nil)
}

// #endregion

// #region DragSurfaceSize

type DragSurfaceSize struct{}

func GTypeDragSurfaceSize() gobject.GType { return gdk_drag_surface_size_get_type.Fn()() }

func (size *DragSurfaceSize) SetSize(width, height int32) {
	gdk_drag_surface_size_set_size.Fn()(size, width, height)
}

// #endregion

// #region DrawContext

type DrawContext struct{ gobject.GObject }

func GTypeDrawContext() gobject.GType { return gdk_draw_context_get_type.Fn()() }

func (dc *DrawContext) GetDisplay() *Display { return gdk_draw_context_get_display.Fn()(dc) }
func (dc *DrawContext) GetSurface() *Surface { return gdk_draw_context_get_surface.Fn()(dc) }

// #endregion

// #region Drop

type Drop struct{ gobject.GObjectCore }

func GTypeDrop() gobject.GType { return gdk_drop_get_type.Fn()() }

func (d *Drop) GetDisplay() *Display                 { return gdk_drop_get_display.Fn()(d) }
func (d *Drop) GetDevice() *Device                   { return gdk_drop_get_device.Fn()(d) }
func (d *Drop) GetSurface() *Surface                 { return gdk_drop_get_surface.Fn()(d) }
func (d *Drop) GetFormats() *ContentFormats          { return gdk_drop_get_formats.Fn()(d) }
func (d *Drop) GetActions() DragAction               { return gdk_drop_get_actions.Fn()(d) }
func (d *Drop) GetDrag() *Drag                       { return gdk_drop_get_drag.Fn()(d) }
func (d *Drop) Status(actions, preferred DragAction) { gdk_drop_status.Fn()(d, actions, preferred) }
func (d *Drop) Finish(action DragAction)             { gdk_drop_finish.Fn()(d, action) }

func (d *Drop) ReadAsync(mimeTypes []string, ioPriority int32, cancellable *gio.GCancellable, callback func(d *Drop, res *gio.GAsyncResult)) {
	var cbk uptr
	if callback != nil {
		cbk = cc.Cbk(func(d *Drop, res *gio.GAsyncResult, _ uptr) {
			callback(d, res)
		})
		d.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	}

	ctypes := cc.NewStrings(mimeTypes)
	defer ctypes.Free()
	gdk_drop_read_async.Fn()(d, ctypes, ioPriority, cancellable, cbk, nil)
}

func (d *Drop) ReadFinish(res *gio.GAsyncResult) (stream *gio.GInputStream, mimeType string, err error) {
	var cMimeType cc.String
	var gerr *glib.GError
	stream = gdk_drop_read_finish.Fn()(d, res, &cMimeType, &gerr)
	mimeType, err = cMimeType.String(), gerr.TakeError()
	return
}

func (d *Drop) ReadValueAsync(typ gobject.GType, ioPriority int32, cancellable *gio.GCancellable, callback func(d *Drop, res *gio.GAsyncResult)) {
	var cbk uptr
	if callback != nil {
		cbk = cc.Cbk(func(d *Drop, res *gio.GAsyncResult, _ uptr) {
			callback(d, res)
		})
		d.WeakRef(func(obj *gobject.GObject) { cc.CbkClose(cbk) })
	}

	gdk_drop_read_value_async.Fn()(d, typ, ioPriority, cancellable, cbk, nil)
}

func (d *Drop) ReadValueFinish(res *gio.GAsyncResult) (value *gobject.GValue, err error) {
	var gerr *glib.GError
	value, err = gdk_drop_read_value_finish.Fn()(d, res, &gerr), gerr.TakeError()
	return
}

// #endregion

// #region Event

type Event struct{ gobject.GTypeInstance }

type EventIface interface{ GetEventIface() *Event }

func GetEventIface(w EventIface) *Event {
	if anyptr(w) == nil {
		return nil
	}
	return w.GetEventIface()
}
func (s *Event) GetEventIface() *Event { return s }

func EventCast[T EventIface](e EventIface) T {
	p := e.GetEventIface()
	return *(*T)(uptr(&(p)))
}

func GTypeEvent() gobject.GType                   { return gdk_event_get_type.Fn()() }
func GTypeEventSequence() gobject.GType           { return gdk_event_sequence_get_type.Fn()() }
func (e *Event) Ref() *Event                      { return gdk_event_ref.Fn()(e) }
func (e *Event) Unref()                           { gdk_event_unref.Fn()(e) }
func (e *Event) GetEventType() EventType          { return gdk_event_get_event_type.Fn()(e) }
func (e *Event) GetSurface() *Surface             { return gdk_event_get_surface.Fn()(e) }
func (e *Event) GetSeat() *Seat                   { return gdk_event_get_seat.Fn()(e) }
func (e *Event) GetDevice() *Device               { return gdk_event_get_device.Fn()(e) }
func (e *Event) GetDeviceTool() *DeviceTool       { return gdk_event_get_device_tool.Fn()(e) }
func (e *Event) GetTime() uint32                  { return gdk_event_get_time.Fn()(e) }
func (e *Event) GetDisplay() *Display             { return gdk_event_get_display.Fn()(e) }
func (e *Event) GetEventSequence() *EventSequence { return gdk_event_get_event_sequence.Fn()(e) }
func (e *Event) GetModifierState() ModifierType   { return gdk_event_get_modifier_state.Fn()(e) }
func (e *Event) GetPosition() (x, y float64, ok bool) {
	ok = gdk_event_get_position.Fn()(e, &x, &y) != 0
	return
}
func (e *Event) GetAxes() (axes []float64, ok bool) {
	var p *float64
	var n uint32
	ok = gdk_event_get_axes.Fn()(e, &p, &n) != 0
	axes = cc.Slice(p, n)
	return
}
func (e *Event) GetAxis(axisUse AxisUse) (value float64, ok bool) {
	ok = gdk_event_get_axis.Fn()(e, axisUse, &value) != 0
	return
}
func (e *Event) GetHistory() []TimeCoord {
	var outNCoords uint32
	coord := gdk_event_get_history.Fn()(e, &outNCoords)
	return cc.Slice(coord, outNCoords)
}
func (e *Event) GetPointerEmulated() bool { return gdk_event_get_pointer_emulated.Fn()(e) != 0 }

func (e *Event) TriggersContextMenu() bool { return gdk_event_triggers_context_menu.Fn()(e) != 0 }
func EventsGetDistance(e1, e2 *Event) (distance float64, ok bool) {
	ok = gdk_events_get_distance.Fn()(e1, e2, &distance) != 0
	return
}
func (e1 *Event) GetDistance(e2 *Event) (distance float64, ok bool) {
	ok = gdk_events_get_distance.Fn()(e1, e2, &distance) != 0
	return
}
func EventsGetAngle(e1, e2 *Event) (angle float64, ok bool) {
	ok = gdk_events_get_angle.Fn()(e1, e2, &angle) != 0
	return
}
func (e1 *Event) GetAngle(e2 *Event) (angle float64, ok bool) {
	ok = gdk_events_get_angle.Fn()(e1, e2, &angle) != 0
	return
}
func EventsGetCenter(e1, e2 *Event) (x, y float64, ok bool) {
	ok = gdk_events_get_center.Fn()(e1, e2, &x, &y) != 0
	return
}
func (e1 *Event) GetCenter(e2 *Event) (x, y float64, ok bool) {
	ok = gdk_events_get_center.Fn()(e1, e2, &x, &y) != 0
	return
}
func (e *Event) KeyEventMatches(keyval uint32, modifiers ModifierType) KeyMatch {
	return gdk_key_event_matches.Fn()(e, keyval, modifiers)
}
func (e *Event) KeyEventGetMatch() (keyval uint32, modifiers ModifierType, ok bool) {
	ok = gdk_key_event_get_match.Fn()(e, &keyval, &modifiers) != 0
	return
}

// #endregion

// #region FocusEvent

type FocusEvent struct{ Event }

func GTypeFocusEvent() gobject.GType   { return gdk_focus_event_get_type.Fn()() }
func (e *FocusEvent) GetFocusIn() bool { return gdk_focus_event_get_in.Fn()(e) != 0 }

// #endregion

// #region FrameClock

type FrameClock struct {
	gobject.GObjectCore
}

func GTypeFrameClock() gobject.GType { return gdk_frame_clock_get_type.Fn()() }

func (fc *FrameClock) GetFrameTime() int64 { return gdk_frame_clock_get_frame_time.Fn()(fc) }
func (fc *FrameClock) RequestPhase(phase FrameClockPhase) {
	gdk_frame_clock_request_phase.Fn()(fc, phase)
}
func (fc *FrameClock) BeginUpdating()         { gdk_frame_clock_begin_updating.Fn()(fc) }
func (fc *FrameClock) EndUpdating()           { gdk_frame_clock_end_updating.Fn()(fc) }
func (fc *FrameClock) GetFrameCounter() int64 { return gdk_frame_clock_get_frame_counter.Fn()(fc) }
func (fc *FrameClock) GetHistoryStart() int64 { return gdk_frame_clock_get_history_start.Fn()(fc) }
func (fc *FrameClock) GetTimings(frameCounter int64) *FrameTimings {
	return gdk_frame_clock_get_timings.Fn()(fc, frameCounter)
}
func (fc *FrameClock) GetCurrentTimings() *FrameTimings {
	return gdk_frame_clock_get_current_timings.Fn()(fc)
}
func (fc *FrameClock) GetRefreshInfo(baseTime int64) (refreshInterval, presentationTime int64) {
	gdk_frame_clock_get_refresh_info.Fn()(fc, baseTime, &refreshInterval, &presentationTime)
	return
}
func (fc *FrameClock) GetFps() float64 { return gdk_frame_clock_get_fps.Fn()(fc) }

func (fc *FrameClock) ConnectAfterPaint(sig func(fc *FrameClock)) uint64 {
	return fc.SignalConnect("after-paint", sig, nil)
}
func (fc *FrameClock) ConnectBeforePaint(sig func(fc *FrameClock)) uint64 {
	return fc.SignalConnect("before-paint", sig, nil)
}
func (fc *FrameClock) ConnectFlushEvents(sig func(fc *FrameClock)) uint64 {
	return fc.SignalConnect("flush-events", sig, nil)
}
func (fc *FrameClock) ConnectLayout(sig func(fc *FrameClock)) uint64 {
	return fc.SignalConnect("layout", sig, nil)
}
func (fc *FrameClock) ConnectPaint(sig func(fc *FrameClock)) uint64 {
	return fc.SignalConnect("paint", sig, nil)
}
func (fc *FrameClock) ConnectResumeEvents(sig func(fc *FrameClock)) uint64 {
	return fc.SignalConnect("resume-events", sig, nil)
}
func (fc *FrameClock) ConnectUpdate(sig func(fc *FrameClock)) uint64 {
	return fc.SignalConnect("update", sig, nil)
}

// #endregion

// #region FrameTimings

func GTypeFrameTimings() gobject.GType { return gdk_frame_timings_get_type.Fn()() }

type FrameTimings struct{}

func (ft *FrameTimings) Ref() *FrameTimings     { return gdk_frame_timings_ref.Fn()(ft) }
func (ft *FrameTimings) Unref()                 { gdk_frame_timings_unref.Fn()(ft) }
func (ft *FrameTimings) GetFrameCounter() int64 { return gdk_frame_timings_get_frame_counter.Fn()(ft) }
func (ft *FrameTimings) GetComplete() bool      { return gdk_frame_timings_get_complete.Fn()(ft) != 0 }
func (ft *FrameTimings) GetFrameTime() int64    { return gdk_frame_timings_get_frame_time.Fn()(ft) }
func (ft *FrameTimings) GetPresentationTime() int64 {
	return gdk_frame_timings_get_presentation_time.Fn()(ft)
}
func (ft *FrameTimings) GetRefreshInterval() int64 {
	return gdk_frame_timings_get_refresh_interval.Fn()(ft)
}
func (ft *FrameTimings) GetPredictedPresentationTime() int64 {
	return gdk_frame_timings_get_predicted_presentation_time.Fn()(ft)
}

// #endregion

// #region GLContext

func GdkGLErrorQuark() glib.GQuark  { return gdk_gl_error_quark.Fn()() }
func GTypeGLContext() gobject.GType { return gdk_gl_context_get_type.Fn()() }

type GLContext struct{ DrawContext }

func (ctx *GLContext) GetDisplay() *Display { return gdk_gl_context_get_display.Fn()(ctx) }
func (ctx *GLContext) GetSurface() *Surface { return gdk_gl_context_get_surface.Fn()(ctx) }
func (ctx *GLContext) GetSharedContext() *GLContext {
	return gdk_gl_context_get_shared_context.Fn()(ctx)
}
func (ctx *GLContext) GetVersion() (major, minor int) {
	gdk_gl_context_get_version.Fn()(ctx, &major, &minor)
	return
}
func (ctx *GLContext) IsLegacy() bool { return gdk_gl_context_is_legacy.Fn()(ctx) != 0 }
func (ctx *GLContext) IsShared(other *GLContext) bool {
	return gdk_gl_context_is_shared.Fn()(ctx, other) != 0
}
func (ctx *GLContext) SetRequiredVersion(major, minor int) {
	gdk_gl_context_set_required_version.Fn()(ctx, major, minor)
}
func (ctx *GLContext) GetRequiredVersion() (major, minor int) {
	gdk_gl_context_get_required_version.Fn()(ctx, &major, &minor)
	return
}
func (ctx *GLContext) SetDebugEnabled(enabled bool) {
	gdk_gl_context_set_debug_enabled.Fn()(ctx, cbool(enabled))
}
func (ctx *GLContext) GetDebugEnabled() bool { return gdk_gl_context_get_debug_enabled.Fn()(ctx) != 0 }
func (ctx *GLContext) SetForwardCompatible(compatible bool) {
	gdk_gl_context_set_forward_compatible.Fn()(ctx, cbool(compatible))
}
func (ctx *GLContext) GetForwardCompatible() bool {
	return gdk_gl_context_get_forward_compatible.Fn()(ctx) != 0
}
func (ctx *GLContext) SetAllowedAPIs(apis GLAPI) { gdk_gl_context_set_allowed_apis.Fn()(ctx, apis) }
func (ctx *GLContext) GetAllowedAPIs() GLAPI     { return gdk_gl_context_get_allowed_apis.Fn()(ctx) }
func (ctx *GLContext) GetAPI() GLAPI             { return gdk_gl_context_get_api.Fn()(ctx) }
func (ctx *GLContext) SetUseES(useEs int)        { gdk_gl_context_set_use_es.Fn()(ctx, useEs) }
func (ctx *GLContext) GetUseES() bool            { return gdk_gl_context_get_use_es.Fn()(ctx) != 0 }
func (ctx *GLContext) Realize() error {
	var err *glib.GError
	ok := gdk_gl_context_realize.Fn()(ctx, &err) != 0
	if ok {
		return nil
	}
	return err.TakeError()
}
func (ctx *GLContext) MakeCurrent()   { gdk_gl_context_make_current.Fn()(ctx) }
func GLContextGetCurrent() *GLContext { return gdk_gl_context_get_current.Fn()() }
func GLContextClearCurrent()          { gdk_gl_context_clear_current.Fn()() }

// #endregion

// #region GLTexture

type GLTexture struct{ Texture }

func GTypeGLTexture() gobject.GType { return gdk_gl_texture_get_type.Fn()() }
func NewGLTexture(context *GLContext, id uint32, width, height int32, destroy func()) *Texture {
	var des uptr
	if destroy != nil {
		des = cc.Cbk(func(_ uptr) {
			destroy()
			cc.CbkCloseLate(des)
		})
	}
	return gdk_gl_texture_new.Fn()(context, id, width, height, des, nil)
}
func (t *Texture) Release() { gdk_gl_texture_release.Fn()(t) }

// #endregion

// #region GLTextureBuilder

func GTypeGLTextureBuilder() gobject.GType { return gdk_gl_texture_builder_get_type.Fn()() }

type GLTextureBuilder struct {
	gobject.GObjectCore
}

func NewGLTextureBuilder() *GLTextureBuilder       { return gdk_gl_texture_builder_new.Fn()() }
func (b *GLTextureBuilder) GetContext() *GLContext { return gdk_gl_texture_builder_get_context.Fn()(b) }
func (b *GLTextureBuilder) SetContext(context *GLContext) {
	gdk_gl_texture_builder_set_context.Fn()(b, context)
}
func (b *GLTextureBuilder) GetId() uint32           { return gdk_gl_texture_builder_get_id.Fn()(b) }
func (b *GLTextureBuilder) SetId(id uint32)         { gdk_gl_texture_builder_set_id.Fn()(b, id) }
func (b *GLTextureBuilder) GetWidth() int32         { return gdk_gl_texture_builder_get_width.Fn()(b) }
func (b *GLTextureBuilder) SetWidth(width int32)    { gdk_gl_texture_builder_set_width.Fn()(b, width) }
func (b *GLTextureBuilder) GetHeight() int32        { return gdk_gl_texture_builder_get_height.Fn()(b) }
func (b *GLTextureBuilder) SetHeight(height int32)  { gdk_gl_texture_builder_set_height.Fn()(b, height) }
func (b *GLTextureBuilder) GetFormat() MemoryFormat { return gdk_gl_texture_builder_get_format.Fn()(b) }
func (b *GLTextureBuilder) SetFormat(format MemoryFormat) {
	gdk_gl_texture_builder_set_format.Fn()(b, format)
}
func (b *GLTextureBuilder) GetHasMipmap() bool {
	return gdk_gl_texture_builder_get_has_mipmap.Fn()(b) != 0
}
func (b *GLTextureBuilder) SetHasMipmap(hasMipmap bool) {
	gdk_gl_texture_builder_set_has_mipmap.Fn()(b, cbool(hasMipmap))
}
func (b *GLTextureBuilder) GetSync() uptr     { return gdk_gl_texture_builder_get_sync.Fn()(b) }
func (b *GLTextureBuilder) SetSync(sync uptr) { gdk_gl_texture_builder_set_sync.Fn()(b, sync) }
func (b *GLTextureBuilder) GetColorState() *ColorState {
	return gdk_gl_texture_builder_get_color_state.Fn()(b)
}
func (b *GLTextureBuilder) SetColorState(colorState *ColorState) {
	gdk_gl_texture_builder_set_color_state.Fn()(b, colorState)
}
func (b *GLTextureBuilder) GetUpdateTexture() *Texture {
	return gdk_gl_texture_builder_get_update_texture.Fn()(b)
}
func (b *GLTextureBuilder) SetUpdateTexture(texture *Texture) {
	gdk_gl_texture_builder_set_update_texture.Fn()(b, texture)
}
func (b *GLTextureBuilder) GetUpdateRegion() *cairo.Region {
	return gdk_gl_texture_builder_get_update_region.Fn()(b)
}
func (b *GLTextureBuilder) SetUpdateRegion(region *cairo.Region) {
	gdk_gl_texture_builder_set_update_region.Fn()(b, region)
}
func (b *GLTextureBuilder) Build(destroy func()) *Texture {
	var des uptr
	if destroy != nil {
		des = cc.Cbk(func(_ uptr) {
			destroy()
			cc.CbkCloseLate(des)
		})
	}
	return gdk_gl_texture_builder_build.Fn()(b, des, nil)
}

// #endregion

// #region GrabBrokenEvent

type GrabBrokenEvent struct{ Event }

func GTypeGrabBrokenEvent() gobject.GType { return gdk_grab_broken_event_get_type.Fn()() }
func (e *GrabBrokenEvent) GetGrabBrokenSurface() *Surface {
	return gdk_grab_broken_event_get_grab_surface.Fn()(e)
}
func (e *GrabBrokenEvent) GetGrabBrokenImplicit() bool {
	return gdk_grab_broken_event_get_implicit.Fn()(e) != 0
}

// #endregion

// #region KeyEvent

type KeyEvent struct{ Event }

func GTypeKeyEvent() gobject.GType     { return gdk_key_event_get_type.Fn()() }
func (e *KeyEvent) GetKeyval() uint32  { return gdk_key_event_get_keyval.Fn()(e) }
func (e *KeyEvent) GetKeycode() uint32 { return gdk_key_event_get_keycode.Fn()(e) }
func (e *KeyEvent) GetConsumedModifiers() ModifierType {
	return gdk_key_event_get_consumed_modifiers.Fn()(e)
}
func (e *KeyEvent) GetLayout() uint32 { return gdk_key_event_get_layout.Fn()(e) }
func (e *KeyEvent) GetLevel() uint32  { return gdk_key_event_get_level.Fn()(e) }
func (e *KeyEvent) IsModifier() bool  { return gdk_key_event_is_modifier.Fn()(e) != 0 }

// #endregion

// #region Keyval

func KeyvalName(keyval uint32) string { return gdk_keyval_name.Fn()(keyval).String() }
func KeyvalFromName(keyvalName string) uint32 {
	return gdk_keyval_from_name.Fn()(cc.NewString(keyvalName))
}
func KeyvalConvertCase(symbol uint32) (lower, upper uint32) {
	gdk_keyval_convert_case.Fn()(symbol, &lower, &upper)
	return
}
func KeyvalToUpper(keyval uint32) uint32   { return gdk_keyval_to_upper.Fn()(keyval) }
func KeyvalToLower(keyval uint32) uint32   { return gdk_keyval_to_lower.Fn()(keyval) }
func KeyvalIsUpper(keyval uint32) bool     { return gdk_keyval_is_upper.Fn()(keyval) != 0 }
func KeyvalIsLower(keyval uint32) bool     { return gdk_keyval_is_lower.Fn()(keyval) != 0 }
func KeyvalToUnicode(keyval uint32) uint32 { return gdk_keyval_to_unicode.Fn()(keyval) }
func UnicodeToKeyval(wc uint32) uint32     { return gdk_unicode_to_keyval.Fn()(wc) }

// #endregion

// #region MemoryTexture

type MemoryTexture struct {
	Texture
}

func GTypeMemoryTexture() gobject.GType { return gdk_memory_texture_get_type.Fn()() }
func NewMemoryTexture(width, height int32, format MemoryFormat, bytes *glib.GBytes, stride uint64) *MemoryTexture {
	return gdk_memory_texture_new.Fn()(width, height, format, bytes, stride)
}

// #endregion

// #region MemoryTextureBuilder

type MemoryTextureBuilder struct {
	gobject.GObjectCore
}

func GTypeMemoryTextureBuilder() gobject.GType { return gdk_memory_texture_builder_get_type.Fn()() }

func NewMemoryTextureBuilder() *MemoryTextureBuilder { return gdk_memory_texture_builder_new.Fn()() }
func (b *MemoryTextureBuilder) GetBytes() *glib.GBytes {
	return gdk_memory_texture_builder_get_bytes.Fn()(b)
}
func (b *MemoryTextureBuilder) SetBytes(bytes *glib.GBytes) {
	gdk_memory_texture_builder_set_bytes.Fn()(b, bytes)
}
func (b *MemoryTextureBuilder) GetStride() uint64 {
	return gdk_memory_texture_builder_get_stride.Fn()(b)
}
func (b *MemoryTextureBuilder) SetStride(stride uint64) {
	gdk_memory_texture_builder_set_stride.Fn()(b, stride)
}
func (b *MemoryTextureBuilder) GetWidth() int32 { return gdk_memory_texture_builder_get_width.Fn()(b) }
func (b *MemoryTextureBuilder) SetWidth(width int32) {
	gdk_memory_texture_builder_set_width.Fn()(b, width)
}
func (b *MemoryTextureBuilder) GetHeight() int32 {
	return gdk_memory_texture_builder_get_height.Fn()(b)
}
func (b *MemoryTextureBuilder) SetHeight(height int32) {
	gdk_memory_texture_builder_set_height.Fn()(b, height)
}
func (b *MemoryTextureBuilder) GetFormat() MemoryFormat {
	return gdk_memory_texture_builder_get_format.Fn()(b)
}
func (b *MemoryTextureBuilder) SetFormat(format MemoryFormat) {
	gdk_memory_texture_builder_set_format.Fn()(b, format)
}
func (b *MemoryTextureBuilder) GetColorState() *ColorState {
	return gdk_memory_texture_builder_get_color_state.Fn()(b)
}
func (b *MemoryTextureBuilder) SetColorState(colorState *ColorState) {
	gdk_memory_texture_builder_set_color_state.Fn()(b, colorState)
}
func (b *MemoryTextureBuilder) GetUpdateTexture() *Texture {
	return gdk_memory_texture_builder_get_update_texture.Fn()(b)
}
func (b *MemoryTextureBuilder) SetUpdateTexture(texture *Texture) {
	gdk_memory_texture_builder_set_update_texture.Fn()(b, texture)
}
func (b *MemoryTextureBuilder) GetUpdateRegion() *cairo.Region {
	return gdk_memory_texture_builder_get_update_region.Fn()(b)
}
func (b *MemoryTextureBuilder) SetUpdateRegion(region *cairo.Region) {
	gdk_memory_texture_builder_set_update_region.Fn()(b, region)
}
func (b *MemoryTextureBuilder) Build() *Texture { return gdk_memory_texture_builder_build.Fn()(b) }

// #endregion

// #region Monitor

type Monitor struct{ gobject.GObjectCore }

func GTypeMonitor() gobject.GType                    { return gdk_monitor_get_type.Fn()() }
func (m *Monitor) GetDisplay() *Display              { return gdk_monitor_get_display.Fn()(m) }
func (m *Monitor) GetGeometry(geometry *Rectangle)   { gdk_monitor_get_geometry.Fn()(m, geometry) }
func (m *Monitor) GetWidthMm() int32                 { return gdk_monitor_get_width_mm.Fn()(m) }
func (m *Monitor) GetHeightMm() int32                { return gdk_monitor_get_height_mm.Fn()(m) }
func (m *Monitor) GetManufacturer() string           { return gdk_monitor_get_manufacturer.Fn()(m).String() }
func (m *Monitor) GetModel() string                  { return gdk_monitor_get_model.Fn()(m).String() }
func (m *Monitor) GetConnector() string              { return gdk_monitor_get_connector.Fn()(m).String() }
func (m *Monitor) GetScaleFactor() int32             { return gdk_monitor_get_scale_factor.Fn()(m) }
func (m *Monitor) GetScale() float64                 { return gdk_monitor_get_scale.Fn()(m) }
func (m *Monitor) GetRefreshRate() int32             { return gdk_monitor_get_refresh_rate.Fn()(m) }
func (m *Monitor) GetSubpixelLayout() SubpixelLayout { return gdk_monitor_get_subpixel_layout.Fn()(m) }
func (m *Monitor) IsValid() bool                     { return gdk_monitor_is_valid.Fn()(m) != 0 }
func (m *Monitor) GetDescription() string            { return gdk_monitor_get_description.Fn()(m).String() }
func (m *Monitor) ConnectInvalidate(sig func(m *Monitor)) uint64 {
	return m.SignalConnect("invalidate", sig, nil)
}

// #endregion

// #region MotionEvent

type MotionEvent struct{ Event }

func GTypeMotionEvent() gobject.GType { return gdk_motion_event_get_type.Fn()() }

// #endregion

// #region PadEvent

type PadEvent struct{ Event }

func GTypePadEvent() gobject.GType       { return gdk_pad_event_get_type.Fn()() }
func (e *PadEvent) GetPadButton() uint32 { return gdk_pad_event_get_button.Fn()(e) }
func (e *PadEvent) GetPadAxisValue() (index uint32, value float64) {
	gdk_pad_event_get_axis_value.Fn()(e, &index, &value)
	return
}
func (e *PadEvent) GetPadGroupMode() (group, mode uint32) {
	gdk_pad_event_get_group_mode.Fn()(e, &group, &mode)
	return
}

// #endregion

// #region ProximityEvent

type ProximityEvent struct{ Event }

func GTypeProximityEvent() gobject.GType { return gdk_proximity_event_get_type.Fn()() }

// #endregion

// #region ScrollEvent

type ScrollEvent struct{ Event }

func GTypeScrollEvent() gobject.GType { return gdk_scroll_event_get_type.Fn()() }
func (e *ScrollEvent) GetScrollDirection() ScrollDirection {
	return gdk_scroll_event_get_direction.Fn()(e)
}
func (e *ScrollEvent) GetScrollDeltas() (deltaX, deltaY float64) {
	gdk_scroll_event_get_deltas.Fn()(e, &deltaX, &deltaY)
	return
}
func (e *ScrollEvent) GetScrollUnit() ScrollUnit { return gdk_scroll_event_get_unit.Fn()(e) }
func (e *ScrollEvent) IsScrollStop() bool        { return gdk_scroll_event_is_stop.Fn()(e) != 0 }

// #endregion

// #region Paintable

type PaintableIfaceObj struct {
	Parent gobject.GTypeInterface

	Snapshot                cc.Func // void (* snapshot) (GdkPaintable *paintable, GdkSnapshot *snapshot, double width, double height);
	GetCurrentImage         cc.Func // GdkPaintable * (* get_current_image) (GdkPaintable *paintable);
	GetFlags                cc.Func // GdkPaintableFlags (* get_flags) (GdkPaintable *paintable);
	GetIntrinsicWidth       cc.Func // int (* get_intrinsic_width) (GdkPaintable *paintable);
	GetIntrinsicHeight      cc.Func // int (* get_intrinsic_height) (GdkPaintable *paintable);
	GetIntrinsicAspectRatio cc.Func // double (* get_intrinsic_aspect_ratio) (GdkPaintable *paintable);
}

type PaintableIface interface {
	GetPaintableIface() *Paintable
}

func GetPaintableIface(iface PaintableIface) *Paintable {
	if anyptr(iface) == nil {
		return nil
	}
	return iface.GetPaintableIface()
}

func (p *Paintable) GetPaintableIface() *Paintable { return p }

type Paintable struct{}

func GTypePaintable() gobject.GType { return gdk_paintable_get_type.Fn()() }
func (p *Paintable) Snapshot(snapshot *Snapshot, width, height float64) {
	gdk_paintable_snapshot.Fn()(p, snapshot, width, height)
}
func (p *Paintable) GetCurrentImage() *Paintable { return gdk_paintable_get_current_image.Fn()(p) }
func (p *Paintable) GetFlags() PaintableFlags    { return gdk_paintable_get_flags.Fn()(p) }
func (p *Paintable) GetIntrinsicWidth() int32    { return gdk_paintable_get_intrinsic_width.Fn()(p) }
func (p *Paintable) GetIntrinsicHeight() int32   { return gdk_paintable_get_intrinsic_height.Fn()(p) }
func (p *Paintable) GetIntrinsicAspectRatio() float64 {
	return gdk_paintable_get_intrinsic_aspect_ratio.Fn()(p)
}
func (p *Paintable) ComputeConcreteSize(specifiedWidth, specifiedHeight, defaultWidth, defaultHeight float64) (concreteWidth, concreteHeight float64) {
	gdk_paintable_compute_concrete_size.Fn()(p, specifiedWidth, specifiedHeight, defaultWidth, defaultHeight, &concreteWidth, &concreteHeight)
	return
}
func (p *Paintable) InvalidateContents() { gdk_paintable_invalidate_contents.Fn()(p) }
func (p *Paintable) InvalidateSize()     { gdk_paintable_invalidate_size.Fn()(p) }
func NewPaintableEmpty(intrinsicWidth, intrinsicHeight int32) *Paintable {
	return gdk_paintable_new_empty.Fn()(intrinsicWidth, intrinsicHeight)
}
func (p *Paintable) ConnectInvalidateContents(sig func(p *Paintable)) uint64 {
	return (*gobject.GObject)(uptr(p)).SignalConnect("invalidate-contents", sig, nil)
}
func (p *Paintable) ConnectInvalidateSize(sig func(p *Paintable)) uint64 {
	return (*gobject.GObject)(uptr(p)).SignalConnect("invalidate-size", sig, nil)
}

// #endregion

// #region Pango

func PangoLayoutLineGetClipRegion(lin *pango.LayoutLine, xOrigin, yOrigin int32, indexRanges []int32) *cairo.Region {
	var ranges *int32
	num := int32(len(indexRanges))
	if num > 0 {
		ranges = &indexRanges[0]
	}
	return gdk_pango_layout_line_get_clip_region.Fn()(lin, xOrigin, yOrigin, ranges, num)
}
func PangoLayoutGetClipRegion(lyt *pango.Layout, xOrigin, yOrigin int32, indexRanges []int32) *cairo.Region {
	var ranges *int32
	num := int32(len(indexRanges))
	if num > 0 {
		ranges = &indexRanges[0]
	}
	return gdk_pango_layout_get_clip_region.Fn()(lyt, xOrigin, yOrigin, ranges, num)
}

// #endregion

// #region Popup

type Popup struct{ Surface }

func GTypePopup() gobject.GType { return gdk_popup_get_type.Fn()() }

func (popup *Popup) Present(width, height int32, layout *PopupLayout) bool {
	return gdk_popup_present.Fn()(popup, width, height, layout) != 0
}
func (popup *Popup) GetSurfaceAnchor() Gravity { return gdk_popup_get_surface_anchor.Fn()(popup) }
func (popup *Popup) GetRectAnchor() Gravity    { return gdk_popup_get_rect_anchor.Fn()(popup) }
func (popup *Popup) GetParent() *Surface       { return gdk_popup_get_parent.Fn()(popup) }
func (popup *Popup) GetPositionX() int32       { return gdk_popup_get_position_x.Fn()(popup) }
func (popup *Popup) GetPositionY() int32       { return gdk_popup_get_position_y.Fn()(popup) }
func (popup *Popup) GetAutohide() bool         { return gdk_popup_get_autohide.Fn()(popup) != 0 }

// #endregion

// #region PopupLayout

type PopupLayout struct{}

func GTypePopupLayout() gobject.GType { return gdk_popup_layout_get_type.Fn()() }

func NewPopupLayout(anchorRect *Rectangle, rectAnchor, surfaceAnchor Gravity) *PopupLayout {
	return gdk_popup_layout_new.Fn()(anchorRect, rectAnchor, surfaceAnchor)
}
func (layout *PopupLayout) Ref() *PopupLayout  { return gdk_popup_layout_ref.Fn()(layout) }
func (layout *PopupLayout) Unref()             { gdk_popup_layout_unref.Fn()(layout) }
func (layout *PopupLayout) Copy() *PopupLayout { return gdk_popup_layout_copy.Fn()(layout) }
func (layout *PopupLayout) Equal(other *PopupLayout) bool {
	return gdk_popup_layout_equal.Fn()(layout, other) != 0
}
func (layout *PopupLayout) SetAnchorRect(anchorRect *Rectangle) {
	gdk_popup_layout_set_anchor_rect.Fn()(layout, anchorRect)
}
func (layout *PopupLayout) GetAnchorRect() *Rectangle {
	return gdk_popup_layout_get_anchor_rect.Fn()(layout)
}
func (layout *PopupLayout) SetRectAnchor(anchor Gravity) {
	gdk_popup_layout_set_rect_anchor.Fn()(layout, anchor)
}
func (layout *PopupLayout) GetRectAnchor() Gravity {
	return gdk_popup_layout_get_rect_anchor.Fn()(layout)
}
func (layout *PopupLayout) SetSurfaceAnchor(anchor Gravity) {
	gdk_popup_layout_set_surface_anchor.Fn()(layout, anchor)
}
func (layout *PopupLayout) GetSurfaceAnchor() Gravity {
	return gdk_popup_layout_get_surface_anchor.Fn()(layout)
}
func (layout *PopupLayout) SetAnchorHints(anchorHints AnchorHints) {
	gdk_popup_layout_set_anchor_hints.Fn()(layout, anchorHints)
}
func (layout *PopupLayout) GetAnchorHints() AnchorHints {
	return gdk_popup_layout_get_anchor_hints.Fn()(layout)
}
func (layout *PopupLayout) SetOffset(dx, dy int32) { gdk_popup_layout_set_offset.Fn()(layout, dx, dy) }
func (layout *PopupLayout) GetOffset() (dx, dy int32) {
	gdk_popup_layout_get_offset.Fn()(layout, &dx, &dy)
	return
}
func (layout *PopupLayout) SetShadowWidth(left, right, top, bottom int32) {
	gdk_popup_layout_set_shadow_width.Fn()(layout, left, right, top, bottom)
}
func (layout *PopupLayout) GetShadowWidth() (left, right, top, bottom int32) {
	gdk_popup_layout_get_shadow_width.Fn()(layout, &left, &right, &top, &bottom)
	return
}

// #endregion

// #region Rectangle

type Rectangle struct{ X, Y, W, H int32 }

func (rect *Rectangle) Intersect(other *Rectangle) (ok bool, dest *Rectangle) {
	dest = new(Rectangle)
	ok = gdk_rectangle_intersect.Fn()(rect, other, dest) != 0
	return
}
func (rect *Rectangle) Union(other *Rectangle) (dest *Rectangle) {
	dest = new(Rectangle)
	gdk_rectangle_union.Fn()(rect, other, dest)
	return
}
func (rect *Rectangle) Equal(other *Rectangle) bool {
	return gdk_rectangle_equal.Fn()(rect, other) != 0
}
func (rect *Rectangle) ContainsPoint(x, y int32) bool {
	return gdk_rectangle_contains_point.Fn()(rect, x, y) != 0
}
func GTypeRectangle() gobject.GType { return gdk_rectangle_get_type.Fn()() }

// #endregion

// #region RGBA

type RGBA struct{ Red, Green, Blue, Alpha float32 }

func GTypeRGBA() gobject.GType            { return gdk_rgba_get_type.Fn()() }
func (rgba *RGBA) Copy() *RGBA            { return gdk_rgba_copy.Fn()(rgba) }
func (rgba *RGBA) Free()                  { gdk_rgba_free.Fn()(rgba) }
func (rgba *RGBA) IsClear() bool          { return gdk_rgba_is_clear.Fn()(rgba) != 0 }
func (rgba *RGBA) IsOpaque() bool         { return gdk_rgba_is_opaque.Fn()(rgba) != 0 }
func (rgba *RGBA) Hash() uint32           { return gdk_rgba_hash.Fn()((rgba)) }
func (rgba *RGBA) Equal(other *RGBA) bool { return gdk_rgba_equal.Fn()((rgba), (other)) != 0 }
func (rgba *RGBA) Parse(spec string) bool {
	cSpec := cc.NewString(spec)
	defer cSpec.Free()
	return gdk_rgba_parse.Fn()(rgba, cSpec) != 0
}
func (rgba *RGBA) ToString() string {
	return gdk_rgba_to_string.Fn()(rgba).TakeString()
}

// #endregion

// #region Seat

type Seat struct{ gobject.GObjectCore }

func GTypeSeat() gobject.GType                       { return gdk_seat_get_type.Fn()() }
func (seat *Seat) GetDisplay() *Display              { return gdk_seat_get_display.Fn()(seat) }
func (seat *Seat) GetCapabilities() SeatCapabilities { return gdk_seat_get_capabilities.Fn()(seat) }
func (seat *Seat) GetDevices(capabilities SeatCapabilities) *glib.GList[Device] {
	return gdk_seat_get_devices.Fn()(seat, capabilities)
}
func (seat *Seat) GetTools() *glib.GList[DeviceTool] { return gdk_seat_get_tools.Fn()(seat) }
func (seat *Seat) GetPointer() *Device               { return gdk_seat_get_pointer.Fn()(seat) }
func (seat *Seat) GetKeyboard() *Device              { return gdk_seat_get_keyboard.Fn()(seat) }

func (seat *Seat) ConnectDeviceAdded(sig func(seat *Seat, device *Device)) uint64 {
	return seat.SignalConnect("device-added", sig, nil)
}
func (seat *Seat) ConnectDeviceRemoved(sig func(seat *Seat, device *Device)) uint64 {
	return seat.SignalConnect("device-removed", sig, nil)
}
func (seat *Seat) ConnectToolAdded(sig func(seat *Seat, tool *DeviceTool)) uint64 {
	return seat.SignalConnect("tool-added", sig, nil)
}
func (seat *Seat) ConnectToolRemoved(sig func(seat *Seat, tool *DeviceTool)) uint64 {
	return seat.SignalConnect("tool-removed", sig, nil)
}

// #endregion

// #region Snapshot

type Snapshot struct{ gobject.GObjectCore }

// #endregion

// #region Surface

type Surface struct{ gobject.GObjectCore }

func GTypeSurface() gobject.GType                  { return gdk_surface_get_type.Fn()() }
func NewSurfaceToplevel(display *Display) *Surface { return gdk_surface_new_toplevel.Fn()(display) }
func NewSurfacePopup(parent *Surface, autohide bool) *Surface {
	return gdk_surface_new_popup.Fn()(parent, cbool(autohide))
}
func (surface *Surface) Destroy()             { gdk_surface_destroy.Fn()(surface) }
func (surface *Surface) IsDestroyed() bool    { return gdk_surface_is_destroyed.Fn()(surface) != 0 }
func (surface *Surface) GetDisplay() *Display { return gdk_surface_get_display.Fn()(surface) }
func (surface *Surface) Hide()                { gdk_surface_hide.Fn()(surface) }
func (surface *Surface) SetInputRegion(region *cairo.Region) {
	gdk_surface_set_input_region.Fn()(surface, region)
}
func (surface *Surface) GetMapped() bool          { return gdk_surface_get_mapped.Fn()(surface) != 0 }
func (surface *Surface) SetCursor(cursor *Cursor) { gdk_surface_set_cursor.Fn()(surface, cursor) }
func (surface *Surface) GetCursor() *Cursor       { return gdk_surface_get_cursor.Fn()(surface) }
func (surface *Surface) SetDeviceCursor(device *Device, cursor *Cursor) {
	gdk_surface_set_device_cursor.Fn()(surface, device, cursor)
}
func (surface *Surface) GetDeviceCursor(device *Device) *Cursor {
	return gdk_surface_get_device_cursor.Fn()(surface, device)
}
func (surface *Surface) GetWidth() int32  { return gdk_surface_get_width.Fn()(surface) }
func (surface *Surface) GetHeight() int32 { return gdk_surface_get_height.Fn()(surface) }
func (surface *Surface) TranslateCoordinates(to *Surface) (ok bool, x, y float64) {
	ok = gdk_surface_translate_coordinates.Fn()(surface, to, &x, &y) != 0
	return
}
func (surface *Surface) GetScaleFactor() int32 { return gdk_surface_get_scale_factor.Fn()(surface) }
func (surface *Surface) GetScale() float64     { return gdk_surface_get_scale.Fn()(surface) }
func (surface *Surface) GetDevicePosition(device *Device) (x, y float64, mask ModifierType, ok bool) {
	ok = gdk_surface_get_device_position.Fn()(surface, device, &x, &y, &mask) != 0
	return
}
func (surface *Surface) CreateSimilarSurface(content cairo.Content, width, height int32) *cairo.Surface {
	return gdk_surface_create_similar_surface.Fn()(surface, content, width, height)
}
func (surface *Surface) Beep()                      { gdk_surface_beep.Fn()(surface) }
func (surface *Surface) QueueRender()               { gdk_surface_queue_render.Fn()(surface) }
func (surface *Surface) RequestLayout()             { gdk_surface_request_layout.Fn()(surface) }
func (surface *Surface) GetFrameClock() *FrameClock { return gdk_surface_get_frame_clock.Fn()(surface) }
func (surface *Surface) SetOpaqueRegion(region *cairo.Region) {
	gdk_surface_set_opaque_region.Fn()(surface, region)
}
func (surface *Surface) CreateCairoContext() *cairo.Context {
	return gdk_surface_create_cairo_context.Fn()(surface)
}
func (surface *Surface) CreateGLContext() (ctx *GLContext, err error) {
	var gerr *glib.GError
	ctx, err = gdk_surface_create_gl_context.Fn()(surface, &gerr), gerr.TakeError()
	return
}

func (surface *Surface) ConnectEnterMonitor(sig func(surface *Surface, monitor *Monitor)) uint64 {
	return surface.SignalConnect("enter-monitor", sig, nil)
}
func (surface *Surface) ConnectEvent(sig func(surface *Surface, event *Event, _ uptr) bool) uint64 {
	return surface.SignalConnect("event", sig, nil)
}
func (surface *Surface) ConnectLayout(sig func(surface *Surface, width, height int32)) uint64 {
	return surface.SignalConnect("layout", sig, nil)
}
func (surface *Surface) ConnectLeaveMonitor(sig func(surface *Surface, monitor *Monitor)) uint64 {
	return surface.SignalConnect("leave-monitor", sig, nil)
}
func (surface *Surface) ConnectRender(sig func(surface *Surface, region cairo.Region, _ uptr) bool) uint64 {
	return surface.SignalConnect("render", sig, nil)
}

// #endregion

// #region Texture

type Texture struct {
	gobject.GObjectCore

	Paintable
	gio.GIcon
	gio.GLoadableIcon
}

func GdkTextureErrorQuark() glib.GQuark           { return gdk_texture_error_quark.Fn()() }
func GTypeTexture() gobject.GType                 { return gdk_texture_get_type.Fn()() }
func NewTextureForPixbuf(pixbuf *Pixbuf) *Texture { return gdk_texture_new_for_pixbuf.Fn()(pixbuf) }
func NewTextureFromResource(resourcePath string) *Texture {
	return gdk_texture_new_from_resource.Fn()(cc.NewString(resourcePath))
}
func NewTextureFromFile(file *gio.GFile) (tex *Texture, err error) {
	var gerr *glib.GError
	tex, err = gdk_texture_new_from_file.Fn()(file, &gerr), gerr.TakeError()
	return
}
func NewTextureFromFilename(path string) (tex *Texture, err error) {
	var gerr *glib.GError
	tex, err = gdk_texture_new_from_filename.Fn()(cc.NewString(path), &gerr), gerr.TakeError()
	return
}
func NewTextureFromBytes(bytes *glib.GBytes) (tex *Texture, err error) {
	var gerr *glib.GError
	tex, err = gdk_texture_new_from_bytes.Fn()(bytes, &gerr), gerr.TakeError()
	return
}
func (t *Texture) GetWidth() int32                   { return gdk_texture_get_width.Fn()(t) }
func (t *Texture) GetHeight() int32                  { return gdk_texture_get_height.Fn()(t) }
func (t *Texture) GetFormat() MemoryFormat           { return gdk_texture_get_format.Fn()(t) }
func (t *Texture) GetColorState() *ColorState        { return gdk_texture_get_color_state.Fn()(t) }
func (t *Texture) Download(data uptr, stride uint64) { gdk_texture_download.Fn()(t, data, stride) }
func (t *Texture) SaveToPNG(filename string) bool {
	return gdk_texture_save_to_png.Fn()(t, cc.NewString(filename)) != 0
}
func (t *Texture) SaveToPNGBytes() *glib.GBytes { return gdk_texture_save_to_png_bytes.Fn()(t) }
func (t *Texture) SaveToTIFF(filename string) bool {
	return gdk_texture_save_to_tiff.Fn()(t, cc.NewString(filename)) != 0
}
func (t *Texture) SaveToTIFFBytes() *glib.GBytes { return gdk_texture_save_to_tiff_bytes.Fn()(t) }

// #endregion

// #region TouchEvent

type TouchEvent struct{ Event }

func GTypeTouchEvent() gobject.GType { return gdk_touch_event_get_type.Fn()() }
func (e *TouchEvent) GetEmulatingPointer() bool {
	return gdk_touch_event_get_emulating_pointer.Fn()(e) != 0
}

// #endregion

// #region TouchpadEvent
type TouchpadEvent struct{ Event }

func GTypeTouchpadEvent() gobject.GType { return gdk_touchpad_event_get_type.Fn()() }
func (e *TouchpadEvent) GetTouchpadGesturePhase() TouchpadGesturePhase {
	return gdk_touchpad_event_get_gesture_phase.Fn()(e)
}
func (e *TouchpadEvent) GetTouchpadNFingers() uint32 { return gdk_touchpad_event_get_n_fingers.Fn()(e) }
func (e *TouchpadEvent) GetTouchpadDeltas() (dx, dy float64) {
	gdk_touchpad_event_get_deltas.Fn()(e, &dx, &dy)
	return
}
func (e *TouchpadEvent) GetTouchpadPinchAngleDelta() float64 {
	return gdk_touchpad_event_get_pinch_angle_delta.Fn()(e)
}
func (e *TouchpadEvent) GetTouchpadPinchScale() float64 {
	return gdk_touchpad_event_get_pinch_scale.Fn()(e)
}

// #endregion

// #region TextureDownloader

type TextureDownloader struct{}

func GTypeTextureDownloader() gobject.GType { return gdk_texture_downloader_get_type.Fn()() }
func NewTextureDownloader(texture *Texture) *TextureDownloader {
	return gdk_texture_downloader_new.Fn()(texture)
}
func (d *TextureDownloader) Copy() *TextureDownloader { return gdk_texture_downloader_copy.Fn()(d) }
func (d *TextureDownloader) Free()                    { gdk_texture_downloader_free.Fn()(d) }
func (d *TextureDownloader) SetTexture(texture *Texture) {
	gdk_texture_downloader_set_texture.Fn()(d, texture)
}
func (d *TextureDownloader) GetTexture() *Texture { return gdk_texture_downloader_get_texture.Fn()(d) }
func (d *TextureDownloader) SetFormat(format MemoryFormat) {
	gdk_texture_downloader_set_format.Fn()(d, format)
}
func (d *TextureDownloader) GetFormat() MemoryFormat {
	return gdk_texture_downloader_get_format.Fn()(d)
}
func (d *TextureDownloader) SetColorState(colorState *ColorState) {
	gdk_texture_downloader_set_color_state.Fn()(d, colorState)
}
func (d *TextureDownloader) GetColorState() *ColorState {
	return gdk_texture_downloader_get_color_state.Fn()(d)
}
func (d *TextureDownloader) DownloadInto(data uptr, stride uint64) {
	gdk_texture_downloader_download_into.Fn()(d, data, stride)
}
func (d *TextureDownloader) DownloadBytes() (bytes *glib.GBytes, stride uint64) {
	bytes = gdk_texture_downloader_download_bytes.Fn()(d, &stride)
	return
}

// #endregion

// #region Toplevel

type ToplevelIface interface{ GetToplevelIface() *Toplevel }

func GetToplevelIface(t ToplevelIface) *Toplevel {
	if anyptr(t) == nil {
		return nil
	}
	return t.GetToplevelIface()
}
func (t *Toplevel) GetToplevelIface() *Toplevel { return t }

type Toplevel struct{ Surface }

func GTypeToplevel() gobject.GType { return gdk_toplevel_get_type.Fn()() }

func (t *Toplevel) Present(layout *ToplevelLayout) { gdk_toplevel_present.Fn()(t, layout) }
func (t *Toplevel) Minimize() bool                 { return gdk_toplevel_minimize.Fn()(t) != 0 }
func (t *Toplevel) Lower() bool                    { return gdk_toplevel_lower.Fn()(t) != 0 }
func (t *Toplevel) Focus(timestamp uint32)         { gdk_toplevel_focus.Fn()(t, timestamp) }
func (t *Toplevel) GetState() ToplevelState        { return gdk_toplevel_get_state.Fn()(t) }
func (t *Toplevel) SetTitle(title string) {
	cTitle := cc.NewString(title)
	defer cTitle.Free()
	gdk_toplevel_set_title.Fn()(t, cTitle)
}
func (t *Toplevel) SetStartupId(startupId string) {
	cId := cc.NewString(startupId)
	defer cId.Free()
	gdk_toplevel_set_startup_id.Fn()(t, cId)
}
func (t *Toplevel) SetTransientFor(parent *Surface) { gdk_toplevel_set_transient_for.Fn()(t, parent) }
func (t *Toplevel) SetModal(modal bool)             { gdk_toplevel_set_modal.Fn()(t, cbool(modal)) }
func (t *Toplevel) SetIconList(surfaces uptr)       { gdk_toplevel_set_icon_list.Fn()(t, surfaces) }
func (t *Toplevel) ShowWindowMenu(event *Event) bool {
	return gdk_toplevel_show_window_menu.Fn()(t, event) != 0
}
func (t *Toplevel) SetDecorated(decorated bool) { gdk_toplevel_set_decorated.Fn()(t, cbool(decorated)) }
func (t *Toplevel) SetDeletable(deletable bool) { gdk_toplevel_set_deletable.Fn()(t, cbool(deletable)) }
func (t *Toplevel) SupportsEdgeConstraints() bool {
	return gdk_toplevel_supports_edge_constraints.Fn()(t) != 0
}
func (t *Toplevel) InhibitSystemShortcuts(event *Event) {
	gdk_toplevel_inhibit_system_shortcuts.Fn()(t, event)
}
func (t *Toplevel) RestoreSystemShortcuts() { gdk_toplevel_restore_system_shortcuts.Fn()(t) }
func (t *Toplevel) BeginResize(edge SurfaceEdge, device *Device, button int32, x, y float64, timestamp uint32) {
	gdk_toplevel_begin_resize.Fn()(t, edge, device, button, x, y, timestamp)
}
func (t *Toplevel) BeginMove(device *Device, button int32, x, y float64, timestamp uint32) {
	gdk_toplevel_begin_move.Fn()(t, device, button, x, y, timestamp)
}
func (t *Toplevel) TitlebarGesture(gesture TitlebarGesture) bool {
	return gdk_toplevel_titlebar_gesture.Fn()(t, gesture) != 0
}
func (t *Toplevel) ConnectComputeSize(sig func(t *Toplevel, size *ToplevelSize)) uint64 {
	return (*gobject.GObject)(uptr(t)).SignalConnect("compute-size", sig, nil)
}

// #endregion

// #region ToplevelLayout

type ToplevelLayout struct{}

func GTypeToplevelLayout() gobject.GType { return gdk_toplevel_layout_get_type.Fn()() }

func NewToplevelLayout() *ToplevelLayout        { return gdk_toplevel_layout_new.Fn()() }
func (l *ToplevelLayout) Ref() *ToplevelLayout  { return gdk_toplevel_layout_ref.Fn()(l) }
func (l *ToplevelLayout) Unref()                { gdk_toplevel_layout_unref.Fn()(l) }
func (l *ToplevelLayout) Copy() *ToplevelLayout { return gdk_toplevel_layout_copy.Fn()(l) }
func (l *ToplevelLayout) Equal(other *ToplevelLayout) bool {
	return gdk_toplevel_layout_equal.Fn()(l, other) != 0
}
func (l *ToplevelLayout) SetMaximized(maximized bool) {
	gdk_toplevel_layout_set_maximized.Fn()(l, cbool(maximized))
}
func (l *ToplevelLayout) SetFullscreen(fullscreen bool, monitor *Monitor) {
	gdk_toplevel_layout_set_fullscreen.Fn()(l, cbool(fullscreen), monitor)
}
func (l *ToplevelLayout) GetMaximized() (maximized bool, ok bool) {
	max := int32(0)
	ok, maximized = gdk_toplevel_layout_get_maximized.Fn()(l, &max) != 0, max != 0
	return
}
func (l *ToplevelLayout) GetFullscreen() (fullscreen bool, ok bool) {
	fs := int32(0)
	ok, fullscreen = gdk_toplevel_layout_get_fullscreen.Fn()(l, &fs) != 0, fs != 0
	return
}
func (l *ToplevelLayout) GetFullscreenMonitor() *Monitor {
	return gdk_toplevel_layout_get_fullscreen_monitor.Fn()(l)
}
func (l *ToplevelLayout) SetResizable(resizable bool) {
	gdk_toplevel_layout_set_resizable.Fn()(l, cbool(resizable))
}
func (l *ToplevelLayout) GetResizable() bool { return gdk_toplevel_layout_get_resizable.Fn()(l) != 0 }

// #endregion

// #region ToplevelSize

type ToplevelSize struct{}

func GTypeToplevelSize() gobject.GType { return gdk_toplevel_size_get_type.Fn()() }

func (s *ToplevelSize) GetBounds() (boundsWidth, boundsHeight int) {
	gdk_toplevel_size_get_bounds.Fn()(s, &boundsWidth, &boundsHeight)
	return
}
func (s *ToplevelSize) SetSize(width, height int) { gdk_toplevel_size_set_size.Fn()(s, width, height) }
func (s *ToplevelSize) SetMinSize(minWidth, minHeight int) {
	gdk_toplevel_size_set_min_size.Fn()(s, minWidth, minHeight)
}
func (s *ToplevelSize) SetShadowWidth(left, right, top, bottom int) {
	gdk_toplevel_size_set_shadow_width.Fn()(s, left, right, top, bottom)
}

// #endregion
