package gdk

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
	"github.com/jinzhongmin/gg/glib/gio"
	"github.com/jinzhongmin/gg/glib/gobject"
)

type uptr = unsafe.Pointer

func init() { cc.Open("libgtk-4*") }

var (
	gdk_display_get_default = cc.DlFunc[func() *Display, *Display]{Name: "gdk_display_get_default"}

	// #region AppLaunchContext
	gdk_app_launch_context_get_type      = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_app_launch_context_get_type"}
	gdk_app_launch_context_new           = cc.DlFunc[func() *AppLaunchContext, *AppLaunchContext]{Name: "gdk_app_launch_context_new"}
	gdk_app_launch_context_get_display   = cc.DlFunc[func(*AppLaunchContext) *Display, *Display]{Name: "gdk_app_launch_context_get_display"}
	gdk_app_launch_context_set_desktop   = cc.DlFunc[func(*AppLaunchContext, int32), cc.Void]{Name: "gdk_app_launch_context_set_desktop"}
	gdk_app_launch_context_set_timestamp = cc.DlFunc[func(*AppLaunchContext, uint32), cc.Void]{Name: "gdk_app_launch_context_set_timestamp"}
	gdk_app_launch_context_set_icon      = cc.DlFunc[func(*AppLaunchContext, *gio.GIcon), cc.Void]{Name: "gdk_app_launch_context_set_icon"}
	gdk_app_launch_context_set_icon_name = cc.DlFunc[func(*AppLaunchContext, cc.String), cc.Void]{Name: "gdk_app_launch_context_set_icon_name"}
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
	gdk_button_event_get_button              = cc.DlFunc[func(event *Event) uint32, uint32]{Name: "gdk_button_event_get_button"}
	gdk_scroll_event_get_type                = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_scroll_event_get_type"}
	gdk_scroll_event_get_direction           = cc.DlFunc[func(event *Event) ScrollDirection, ScrollDirection]{Name: "gdk_scroll_event_get_direction"}
	gdk_scroll_event_get_deltas              = cc.DlFunc[func(event *Event, delta_x, delta_y *float64), cc.Void]{Name: "gdk_scroll_event_get_deltas"}
	gdk_scroll_event_get_unit                = cc.DlFunc[func(event *Event) ScrollUnit, ScrollUnit]{Name: "gdk_scroll_event_get_unit"}
	gdk_scroll_event_is_stop                 = cc.DlFunc[func(event *Event) bool, bool]{Name: "gdk_scroll_event_is_stop"}
	gdk_key_event_get_type                   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_key_event_get_type"}
	gdk_key_event_get_keyval                 = cc.DlFunc[func(event *Event) uint32, uint32]{Name: "gdk_key_event_get_keyval"}
	gdk_key_event_get_keycode                = cc.DlFunc[func(event *Event) uint32, uint32]{Name: "gdk_key_event_get_keycode"}
	gdk_key_event_get_consumed_modifiers     = cc.DlFunc[func(event *Event) ModifierType, ModifierType]{Name: "gdk_key_event_get_consumed_modifiers"}
	gdk_key_event_get_layout                 = cc.DlFunc[func(event *Event) uint32, uint32]{Name: "gdk_key_event_get_layout"}
	gdk_key_event_get_level                  = cc.DlFunc[func(event *Event) uint32, uint32]{Name: "gdk_key_event_get_level"}
	gdk_key_event_is_modifier                = cc.DlFunc[func(event *Event) bool, bool]{Name: "gdk_key_event_is_modifier"}
	gdk_focus_event_get_type                 = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_focus_event_get_type"}
	gdk_focus_event_get_in                   = cc.DlFunc[func(event *Event) bool, bool]{Name: "gdk_focus_event_get_in"}
	gdk_touch_event_get_type                 = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_touch_event_get_type"}
	gdk_touch_event_get_emulating_pointer    = cc.DlFunc[func(event *Event) bool, bool]{Name: "gdk_touch_event_get_emulating_pointer"}
	gdk_crossing_event_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_crossing_event_get_type"}
	gdk_crossing_event_get_mode              = cc.DlFunc[func(event *Event) CrossingMode, CrossingMode]{Name: "gdk_crossing_event_get_mode"}
	gdk_crossing_event_get_detail            = cc.DlFunc[func(event *Event) NotifyType, NotifyType]{Name: "gdk_crossing_event_get_detail"}
	gdk_crossing_event_get_focus             = cc.DlFunc[func(event *Event) bool, bool]{Name: "gdk_crossing_event_get_focus"}
	gdk_touchpad_event_get_type              = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_touchpad_event_get_type"}
	gdk_touchpad_event_get_gesture_phase     = cc.DlFunc[func(event *Event) TouchpadGesturePhase, TouchpadGesturePhase]{Name: "gdk_touchpad_event_get_gesture_phase"}
	gdk_touchpad_event_get_n_fingers         = cc.DlFunc[func(event *Event) uint32, uint32]{Name: "gdk_touchpad_event_get_n_fingers"}
	gdk_touchpad_event_get_deltas            = cc.DlFunc[func(event *Event, dx, dy *float64), cc.Void]{Name: "gdk_touchpad_event_get_deltas"}
	gdk_touchpad_event_get_pinch_angle_delta = cc.DlFunc[func(event *Event) float64, float64]{Name: "gdk_touchpad_event_get_pinch_angle_delta"}
	gdk_touchpad_event_get_pinch_scale       = cc.DlFunc[func(event *Event) float64, float64]{Name: "gdk_touchpad_event_get_pinch_scale"}
	gdk_pad_event_get_type                   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_pad_event_get_type"}
	gdk_pad_event_get_button                 = cc.DlFunc[func(event *Event) uint32, uint32]{Name: "gdk_pad_event_get_button"}
	gdk_pad_event_get_axis_value             = cc.DlFunc[func(event *Event, index *uint32, value *float64), cc.Void]{Name: "gdk_pad_event_get_axis_value"}
	gdk_pad_event_get_group_mode             = cc.DlFunc[func(event *Event, group, mode *uint32), cc.Void]{Name: "gdk_pad_event_get_group_mode"}
	gdk_dnd_event_get_type                   = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_dnd_event_get_type"}
	gdk_dnd_event_get_drop                   = cc.DlFunc[func(event *Event) *Drop, *Drop]{Name: "gdk_dnd_event_get_drop"}
	gdk_grab_broken_event_get_type           = cc.DlFunc[func() gobject.GType, gobject.GType]{Name: "gdk_grab_broken_event_get_type"}
	gdk_grab_broken_event_get_grab_surface   = cc.DlFunc[func(event *Event) *Surface, *Surface]{Name: "gdk_grab_broken_event_get_grab_surface"}
	gdk_grab_broken_event_get_implicit       = cc.DlFunc[func(event *Event) bool, bool]{Name: "gdk_grab_broken_event_get_implicit"}
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

)
