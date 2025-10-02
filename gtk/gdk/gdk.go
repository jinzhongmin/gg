package gdk

import (
	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/glib/gio"
	"github.com/jinzhongmin/gg/glib/gobject"
)

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

// #region Display

func GetDisplayDefault() *Display { return gdk_display_get_default.Fn()() }

// #endregion

// #region Event

type Event struct{ gobject.GTypeInstance }

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
	ok = gdk_event_get_position.Fn()(e, &x, &y)
	return
}
func (e *Event) GetAxes() (axes []float64, ok bool) {
	var p *float64
	var n uint32
	ok = gdk_event_get_axes.Fn()(e, &p, &n)
	axes = *(*[]float64)(cc.Slice(uptr(p), int64(n)))
	return
}
func (e *Event) GetAxis(axisUse AxisUse) (value float64, ok bool) {
	ok = gdk_event_get_axis.Fn()(e, axisUse, &value)
	return
}
func (e *Event) GetHistory() []TimeCoord {
	var outNCoords uint32
	coord := gdk_event_get_history.Fn()(e, &outNCoords)
	return *(*[]TimeCoord)(cc.Slice(uptr(coord), int64(outNCoords)))
}
func (e *Event) GetPointerEmulated() bool { return gdk_event_get_pointer_emulated.Fn()(e) }

func GTypeButtonEvent() gobject.GType { return gdk_button_event_get_type.Fn()() }
func (e *Event) GetButton() uint32    { return gdk_button_event_get_button.Fn()(e) }

func GTypeScrollEvent() gobject.GType                { return gdk_scroll_event_get_type.Fn()() }
func (e *Event) GetScrollDirection() ScrollDirection { return gdk_scroll_event_get_direction.Fn()(e) }
func (e *Event) GetScrollDeltas() (deltaX, deltaY float64) {
	gdk_scroll_event_get_deltas.Fn()(e, &deltaX, &deltaY)
	return
}
func (e *Event) GetScrollUnit() ScrollUnit { return gdk_scroll_event_get_unit.Fn()(e) }
func (e *Event) IsScrollStop() bool        { return gdk_scroll_event_is_stop.Fn()(e) }

func GTypeKeyEvent() gobject.GType  { return gdk_key_event_get_type.Fn()() }
func (e *Event) GetKeyval() uint32  { return gdk_key_event_get_keyval.Fn()(e) }
func (e *Event) GetKeycode() uint32 { return gdk_key_event_get_keycode.Fn()(e) }
func (e *Event) GetConsumedModifiers() ModifierType {
	return gdk_key_event_get_consumed_modifiers.Fn()(e)
}
func (e *Event) GetLayout() uint32 { return gdk_key_event_get_layout.Fn()(e) }
func (e *Event) GetLevel() uint32  { return gdk_key_event_get_level.Fn()(e) }
func (e *Event) IsModifier() bool  { return gdk_key_event_is_modifier.Fn()(e) }

func GTypeFocusEvent() gobject.GType { return gdk_focus_event_get_type.Fn()() }
func (e *Event) GetFocusIn() bool    { return gdk_focus_event_get_in.Fn()(e) }

func GTypeTouchEvent() gobject.GType       { return gdk_touch_event_get_type.Fn()() }
func (e *Event) GetEmulatingPointer() bool { return gdk_touch_event_get_emulating_pointer.Fn()(e) }

func GTypeCrossingEvent() gobject.GType        { return gdk_crossing_event_get_type.Fn()() }
func (e *Event) GetCrossingMode() CrossingMode { return gdk_crossing_event_get_mode.Fn()(e) }
func (e *Event) GetCrossingDetail() NotifyType { return gdk_crossing_event_get_detail.Fn()(e) }
func (e *Event) GetCrossingFocus() bool        { return gdk_crossing_event_get_focus.Fn()(e) }

func GTypeTouchpadEvent() gobject.GType { return gdk_touchpad_event_get_type.Fn()() }
func (e *Event) GetTouchpadGesturePhase() TouchpadGesturePhase {
	return gdk_touchpad_event_get_gesture_phase.Fn()(e)
}
func (e *Event) GetTouchpadNFingers() uint32 { return gdk_touchpad_event_get_n_fingers.Fn()(e) }
func (e *Event) GetTouchpadDeltas() (dx, dy float64) {
	gdk_touchpad_event_get_deltas.Fn()(e, &dx, &dy)
	return
}
func (e *Event) GetTouchpadPinchAngleDelta() float64 {
	return gdk_touchpad_event_get_pinch_angle_delta.Fn()(e)
}
func (e *Event) GetTouchpadPinchScale() float64 { return gdk_touchpad_event_get_pinch_scale.Fn()(e) }
func GTypePadEvent() gobject.GType              { return gdk_pad_event_get_type.Fn()() }
func (e *Event) GetPadButton() uint32           { return gdk_pad_event_get_button.Fn()(e) }
func (e *Event) GetPadAxisValue() (index uint32, value float64) {
	gdk_pad_event_get_axis_value.Fn()(e, &index, &value)
	return
}
func (e *Event) GetPadGroupMode() (group, mode uint32) {
	gdk_pad_event_get_group_mode.Fn()(e, &group, &mode)
	return
}
func GTypeDNDEvent() gobject.GType        { return gdk_dnd_event_get_type.Fn()() }
func (e *Event) GetDNDDrop() *Drop        { return gdk_dnd_event_get_drop.Fn()(e) }
func GTypeGrabBrokenEvent() gobject.GType { return gdk_grab_broken_event_get_type.Fn()() }
func (e *Event) GetGrabBrokenSurface() *Surface {
	return gdk_grab_broken_event_get_grab_surface.Fn()(e)
}
func (e *Event) GetGrabBrokenImplicit() bool { return gdk_grab_broken_event_get_implicit.Fn()(e) }
func GTypeMotionEvent() gobject.GType        { return gdk_motion_event_get_type.Fn()() }
func GTypeDeleteEvent() gobject.GType        { return gdk_delete_event_get_type.Fn()() }
func GTypeProximityEvent() gobject.GType     { return gdk_proximity_event_get_type.Fn()() }
func (e *Event) TriggersContextMenu() bool   { return gdk_event_triggers_context_menu.Fn()(e) }
func EventsGetDistance(e1, e2 *Event) (distance float64, ok bool) {
	ok = gdk_events_get_distance.Fn()(e1, e2, &distance)
	return
}
func (e1 *Event) GetDistance(e2 *Event) (distance float64, ok bool) {
	ok = gdk_events_get_distance.Fn()(e1, e2, &distance)
	return
}
func EventsGetAngle(e1, e2 *Event) (angle float64, ok bool) {
	ok = gdk_events_get_angle.Fn()(e1, e2, &angle)
	return
}
func (e1 *Event) GetAngle(e2 *Event) (angle float64, ok bool) {
	ok = gdk_events_get_angle.Fn()(e1, e2, &angle)
	return
}
func EventsGetCenter(e1, e2 *Event) (x, y float64, ok bool) {
	ok = gdk_events_get_center.Fn()(e1, e2, &x, &y)
	return
}
func (e1 *Event) GetCenter(e2 *Event) (x, y float64, ok bool) {
	ok = gdk_events_get_center.Fn()(e1, e2, &x, &y)
	return
}
func (e *Event) KeyEventMatches(keyval uint32, modifiers ModifierType) KeyMatch {
	return gdk_key_event_matches.Fn()(e, keyval, modifiers)
}
func (e *Event) KeyEventGetMatch() (keyval uint32, modifiers ModifierType, ok bool) {
	ok = gdk_key_event_get_match.Fn()(e, &keyval, &modifiers)
	return
}

// #endregion
