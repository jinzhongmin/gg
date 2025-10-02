package imgtk

import (
	"time"
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
	ig "github.com/jinzhongmin/gg/dcimgui"
	"github.com/jinzhongmin/gg/gl"
	"github.com/jinzhongmin/gg/glib/glib"
	"github.com/jinzhongmin/gg/glib/gobject"
	"github.com/jinzhongmin/gg/gtk/gdk"
	"github.com/jinzhongmin/gg/gtk/gdk/gdkkey"
	"github.com/jinzhongmin/gg/gtk/gtk/v4"
)

type uptr = unsafe.Pointer

func div[T any](a T, fns ...func(T)) T {
	for _, fn := range fns {
		fn(a)
	}
	return a
}

type ImArea struct {
	*gtk.GLArea

	imCtx *ig.ImGuiContext

	imeAct      bool
	imeCtx      *gtk.IMMulticontext
	imeRct      gdk.Rectangle
	fpsReset    bool
	fpsCurrent  float64
	fpsFocusIn  float64
	fpsFocusOut float64

	fnRender   func(*ImArea)
	fnInits    []func(*ig.ImGuiContext, *ig.ImGuiIO)
	fnDestroys []func(*ig.ImGuiContext, *ig.ImGuiIO)

	inRender bool
}

func NewImArea(render func(*ImArea)) *ImArea {
	ia := &ImArea{}
	ia.GLArea = gtk.NewGLArea()
	ia.GLArea.Pin(ia)
	ia.fnRender = render

	ia.fpsFocusIn = 25
	ia.fpsFocusOut = 5

	ia.fnInits = make([]func(*ig.ImGuiContext, *ig.ImGuiIO), 0)

	div(ia, func(area *ImArea) {
		area.SetFocusable(true)

		flagQuit := false
		x, y, z, w := float32(0.45), float32(0.55), float32(0.60), float32(1.00)

		runCount := 0
		var restLoop func()
		restLoop = func() {
			runCount += 1

			ms := uint32(1000.0 / area.fpsCurrent)
			timeoutAdd(ms, func() (Continue bool) {

				if flagQuit || runCount > 1 {
					runCount -= 1
					return false
				}

				if area.fpsReset {
					area.fpsReset = false

					runCount -= 1
					restLoop()
					return false
				}

				if area.inRender {
					return true
				}

				area.inRender = true
				area.QueueRender()
				return true
			})

		}
		area.ConnectDestroy(func(w *gtk.Widget) {
			if area.imCtx == nil {
				return
			}

			ig.SetCurrentContext(area.imCtx)
			io := ig.GetIO()

			for _, fn := range area.fnDestroys {
				fn(area.imCtx, io)
			}

			ig.ImplOpenGL3Shutdown()
			area.imCtx.Destroy()
			area.imCtx = nil
			flagQuit = true
		})
		area.ConnectResize(func(g *gtk.GLArea, w, h int32) {
			if area.imCtx == nil {
				return
			}

			if w == 0 || h == 0 {
				return
			}

			ig.SetCurrentContext(area.imCtx)
			io := ig.GetIO()
			io.DisplaySize.X = float32(w)
			io.DisplaySize.Y = float32(h)

			area.fpsCurrent = 21
			area.fpsReset = true
			restLoop()
		})

		area.ConnectRealize(func(w *gtk.Widget) {

			area.imeCtx = gtk.NewIMMulticontext()
			area.imeCtx.SetClientWidget(area)
			area.imeCtx.ConnectCommit(func(c *gtk.IMContext, str cc.String) {
				if area.imCtx == nil || !area.imeAct {
					return
				}
				io := ig.GetIO()
				io.AddInputCharactersUTF8(str)
			})

			area.imCtx = ig.CreateContext(nil)
			ig.SetCurrentContext(area.imCtx)
			ig.ImplOpenGL3Init("#version 130")

			io := ig.GetPlatformIO()
			io.Platform_SetImeDataFn.Bind(
				func(ctx *ig.ImGuiContext, vp *ig.ImGuiViewport, data *ig.ImGuiPlatformImeData) {

					if data.WantVisible == 1 {
						area.imeAct = true
						area.imeCtx.FocusIn()

						area.imeRct.X = int32(data.InputPos.X)
						area.imeRct.Y = int32(data.InputPos.Y)
						area.imeRct.W = int32(-1)
						area.imeRct.H = int32(data.InputLineHeight + 10)

						area.imeCtx.SetCursorLocation(&area.imeRct)
					} else {
						area.imeAct = false
						area.imeCtx.Reset()
					}

				})

			igio := ig.GetIO()
			for _, fn := range area.fnInits {
				fn(area.imCtx, igio)
			}
		})

		connect(area, "render", cc.CbkRaw[func(g *gtk.GLArea, ctx *gdk.GLContext, _ uptr) bool](func(out, ins uptr) {
			ret := (*int32)(out)
			if area.imCtx == nil || flagQuit {
				*ret = 1
				return
			}

			ig.SetCurrentContext(area.imCtx)
			ig.ImplOpenGL3NewFrame()
			{
				ig.NewFrame()
				{
					area.fnRender(area)
				}
				ig.EndFrame()
				ig.Render()

				gl.ClearColor(x*w, y*w, z*w, w)
				gl.Clear(gl.COLOR_BUFFER_BIT)
				time.Sleep(time.Microsecond)
			}
			ig.ImplOpenGL3RenderDrawData(ig.GetDrawData())
			area.inRender = false
			*ret = 1
		}))

	}, func(area *ImArea) {

		// mouse move
		onMove := gtk.NewEventControllerMotion()
		area.AddController(onMove)
		connect(onMove, "motion", cc.CbkRaw[func(c *gtk.EventControllerMotion, x, y float64)](func(out, ins uptr) {
			if area.imCtx == nil {
				return
			}

			fs := (*[3]*float64)(ins)
			ig.SetCurrentContext(area.imCtx)
			io := ig.GetIO()
			io.AddMousePosEvent(float32(*fs[1]), float32(*fs[2]))
		}))

		onMove.ConnectEnter(func(c *gtk.EventControllerMotion, x, y float64) {
			area.fpsCurrent = area.fpsFocusIn
			area.fpsReset = true

			if area.imCtx == nil {
				return
			}
			ig.SetCurrentContext(area.imCtx)
			io := ig.GetIO()
			io.AddFocusEvent(true)
		})
		onMove.ConnectLeave(func(c *gtk.EventControllerMotion) {
			area.fpsCurrent = area.fpsFocusOut
			area.fpsReset = true

			if area.imCtx == nil {
				return
			}
			ig.SetCurrentContext(area.imCtx)
			io := ig.GetIO()
			io.AddFocusEvent(false)
		})

		// click
		onClick := gtk.NewGestureClick()
		area.AddController(onClick)
		onClick.SetButton(0)

		keyGet := func() func(c *gtk.GestureClick) int32 {
			var keyIdx = [10]int32{}
			keyIdx[1] = 0
			keyIdx[2] = 2
			keyIdx[3] = 1
			keyIdx[8] = 3
			keyIdx[9] = 4
			return func(c *gtk.GestureClick) int32 {
				evt := c.GetCurrentEvent()
				btn := evt.GetButton()
				return keyIdx[btn]
			}
		}()
		onClick.ConnectPressed(func(c *gtk.GestureClick, nPress int32, x, y float64) {
			if area.imCtx == nil {
				return
			}
			ig.SetCurrentContext(area.imCtx)
			io := ig.GetIO()
			io.AddMouseButtonEvent(keyGet(c), true)

		})
		onClick.ConnectReleased(func(c *gtk.GestureClick, nPress int32, x, y float64) {
			if area.imCtx == nil {
				return
			}
			ig.SetCurrentContext(area.imCtx)
			io := ig.GetIO()
			io.AddMouseButtonEvent(keyGet(c), false)
		})

		// key
		onKey := gtk.NewEventControllerKey()
		area.AddController(onKey)
		onKey.ConnectKeyPressed(func(c *gtk.EventControllerKey,
			keyVal, keyCode uint32, state gdk.ModifierType) (ok bool) {
			if area.imCtx == nil {
				return
			}
			ig.SetCurrentContext(area.imCtx)
			io := ig.GetIO()

			if v, ok := keyVals(keyVal); ok {
				io.AddKeyEvent(v, true)
			}
			k := keyCodes(keyCode)
			io.AddKeyEvent(k, true)

			if area.imeAct {
				evt := c.GetCurrentEvent()
				if !(area.imeCtx.FilterKeypress(evt)) {
					if keyVal > 32 && keyVal <= 126 {
						io.AddInputCharacter(keyVal)
					}
				}
			} else {
				if keyVal > 32 && keyVal <= 126 {
					io.AddInputCharacter(keyVal)
				}
			}

			return true
		})
		onKey.ConnectKeyReleased(func(c *gtk.EventControllerKey,
			keyVal, keyCode uint32, state gdk.ModifierType) {
			if area.imCtx == nil {
				return
			}
			ig.SetCurrentContext(area.imCtx)
			io := ig.GetIO()

			if v, ok := keyVals(keyVal); ok {
				io.AddKeyEvent(v, false)
			}
			io.AddKeyEvent(keyCodes(keyCode), false)

			if area.imeAct {
				evt := c.GetCurrentEvent()
				area.imeCtx.FilterKeypress(evt)
			}
		})

		// onScroll
		onScroll := gtk.NewEventControllerScroll(gtk.EventControllerScrollFlagsBothAxes)
		area.AddController(onScroll)
		onScroll.ConnectScroll(func(c *gtk.EventControllerScroll, dx, dy float64) (ok bool) {
			if area.imCtx == nil {
				return
			}
			ig.SetCurrentContext(area.imCtx)
			io := ig.GetIO()
			io.AddMouseWheelEvent(-float32(dx), -float32(dy))
			return true
		})

		area.ConnectHide(func(w *gtk.Widget, td gtk.TextDirection) {
			area.fpsCurrent = 1
			area.fpsReset = true
		})
		area.ConnectShow(func(w *gtk.Widget) {
			area.fpsCurrent = area.fpsFocusOut
			area.fpsReset = true
		})

	})

	return ia
}

func (ia *ImArea) ConfigFps(focusIn, focusOut float64) {
	ia.fpsFocusIn = focusIn
	ia.fpsFocusOut = focusOut
}
func (ia *ImArea) SetFps(fps float64) {
	ia.fpsCurrent = fps
	ia.fpsReset = true
}

func (ia *ImArea) ConnectImRender(render func(*ImArea)) {
	glib.IdleAdd(func() (Continue bool) { ia.fnRender = render; return false })
}
func (ia *ImArea) ConnectImInit(init func(*ig.ImGuiContext, *ig.ImGuiIO)) {
	ia.fnInits = append(ia.fnInits, init)
}
func (ia *ImArea) ConnectImDestroy(destroy func(*ig.ImGuiContext, *ig.ImGuiIO)) {
	ia.fnDestroys = append(ia.fnDestroys, destroy)
}

func keyVals(kv uint32) (ig.ImGuiKey, bool) {
	mp := map[uint32]ig.ImGuiKey{
		gdkkey.KEY_leftarrow:  ig.ImGuiKey_LeftArrow,
		gdkkey.KEY_rightarrow: ig.ImGuiKey_RightArrow,
		gdkkey.KEY_uparrow:    ig.ImGuiKey_UpArrow,
		gdkkey.KEY_downarrow:  ig.ImGuiKey_DownArrow,

		gdkkey.KEY_Shift_L:   ig.ImGuiKey_LeftShift,
		gdkkey.KEY_Shift_R:   ig.ImGuiKey_RightShift,
		gdkkey.KEY_Control_L: ig.ImGuiKey_LeftCtrl,
		gdkkey.KEY_Control_R: ig.ImGuiKey_RightCtrl,
		gdkkey.KEY_Alt_L:     ig.ImGuiKey_LeftAlt,
		gdkkey.KEY_Alt_R:     ig.ImGuiKey_RightAlt,

		gdkkey.KEY_F1:  ig.ImGuiKey_F1,
		gdkkey.KEY_F2:  ig.ImGuiKey_F2,
		gdkkey.KEY_F3:  ig.ImGuiKey_F3,
		gdkkey.KEY_F4:  ig.ImGuiKey_F4,
		gdkkey.KEY_F5:  ig.ImGuiKey_F5,
		gdkkey.KEY_F6:  ig.ImGuiKey_F6,
		gdkkey.KEY_F7:  ig.ImGuiKey_F7,
		gdkkey.KEY_F8:  ig.ImGuiKey_F8,
		gdkkey.KEY_F9:  ig.ImGuiKey_F9,
		gdkkey.KEY_F10: ig.ImGuiKey_F10,
		gdkkey.KEY_F11: ig.ImGuiKey_F11,
		gdkkey.KEY_F12: ig.ImGuiKey_F12,
		gdkkey.KEY_F13: ig.ImGuiKey_F13,
		gdkkey.KEY_F14: ig.ImGuiKey_F14,
		gdkkey.KEY_F15: ig.ImGuiKey_F15,
		gdkkey.KEY_F16: ig.ImGuiKey_F16,
		gdkkey.KEY_F17: ig.ImGuiKey_F17,
		gdkkey.KEY_F18: ig.ImGuiKey_F18,
		gdkkey.KEY_F19: ig.ImGuiKey_F19,
		gdkkey.KEY_F20: ig.ImGuiKey_F20,
		gdkkey.KEY_F21: ig.ImGuiKey_F21,
		gdkkey.KEY_F22: ig.ImGuiKey_F22,
		gdkkey.KEY_F23: ig.ImGuiKey_F23,
		gdkkey.KEY_F24: ig.ImGuiKey_F24,
	}
	v, ok := mp[kv]
	return v, ok
}
func keyCodes(kc uint32) ig.ImGuiKey {
	mp := map[byte]ig.ImGuiKey{
		// 0:   ig.ImGuiKey_None,
		// 8:   ig.ImGuiKey_Backspace,
		// 9:   ig.ImGuiKey_Tab,
		// 13:  ig.ImGuiKey_Enter,
		// 27:  ig.ImGuiKey_Escape,
		// 32:  ig.ImGuiKey_Space,
		// 127: ig.ImGuiKey_Delete,

		// 控制字符映射
		0:   ig.ImGuiKey_None,      // NUL
		1:   ig.ImGuiKey_None,      // SOH
		2:   ig.ImGuiKey_None,      // STX
		3:   ig.ImGuiKey_None,      // ETX
		4:   ig.ImGuiKey_None,      // EOT
		5:   ig.ImGuiKey_None,      // ENQ
		6:   ig.ImGuiKey_None,      // ACK
		7:   ig.ImGuiKey_None,      // BEL
		8:   ig.ImGuiKey_Backspace, // BS
		9:   ig.ImGuiKey_Tab,       // HT
		10:  ig.ImGuiKey_Enter,     // LF
		11:  ig.ImGuiKey_None,      // VT
		12:  ig.ImGuiKey_None,      // FF
		13:  ig.ImGuiKey_Enter,     // CR
		14:  ig.ImGuiKey_None,      // SO
		15:  ig.ImGuiKey_None,      // SI
		16:  ig.ImGuiKey_None,      // DLE
		17:  ig.ImGuiKey_None,      // DC1
		18:  ig.ImGuiKey_None,      // DC2
		19:  ig.ImGuiKey_None,      // DC3
		20:  ig.ImGuiKey_None,      // DC4
		21:  ig.ImGuiKey_None,      // NAK
		22:  ig.ImGuiKey_None,      // SYN
		23:  ig.ImGuiKey_None,      // ETB
		24:  ig.ImGuiKey_None,      // CAN
		25:  ig.ImGuiKey_None,      // EM
		26:  ig.ImGuiKey_None,      // SUB
		27:  ig.ImGuiKey_Escape,    // ESC
		28:  ig.ImGuiKey_None,      // FS
		29:  ig.ImGuiKey_None,      // GS
		30:  ig.ImGuiKey_None,      // RS
		31:  ig.ImGuiKey_None,      // US
		127: ig.ImGuiKey_Delete,    // DEL

		'0': ig.ImGuiKey_0,
		'1': ig.ImGuiKey_1,
		'2': ig.ImGuiKey_2,
		'3': ig.ImGuiKey_3,
		'4': ig.ImGuiKey_4,
		'5': ig.ImGuiKey_5,
		'6': ig.ImGuiKey_6,
		'7': ig.ImGuiKey_7,
		'8': ig.ImGuiKey_8,
		'9': ig.ImGuiKey_9,
		'A': ig.ImGuiKey_A,
		'B': ig.ImGuiKey_B,
		'C': ig.ImGuiKey_C,
		'D': ig.ImGuiKey_D,
		'E': ig.ImGuiKey_E,
		'F': ig.ImGuiKey_F,
		'G': ig.ImGuiKey_G,
		'H': ig.ImGuiKey_H,
		'I': ig.ImGuiKey_I,
		'J': ig.ImGuiKey_J,
		'K': ig.ImGuiKey_K,
		'L': ig.ImGuiKey_L,
		'M': ig.ImGuiKey_M,
		'N': ig.ImGuiKey_N,
		'O': ig.ImGuiKey_O,
		'P': ig.ImGuiKey_P,
		'Q': ig.ImGuiKey_Q,
		'R': ig.ImGuiKey_R,
		'S': ig.ImGuiKey_S,
		'T': ig.ImGuiKey_T,
		'U': ig.ImGuiKey_U,
		'V': ig.ImGuiKey_V,
		'W': ig.ImGuiKey_W,
		'X': ig.ImGuiKey_X,
		'Y': ig.ImGuiKey_Y,
		'Z': ig.ImGuiKey_Z,
		'a': ig.ImGuiKey_A,
		'b': ig.ImGuiKey_B,
		'c': ig.ImGuiKey_C,
		'd': ig.ImGuiKey_D,
		'e': ig.ImGuiKey_E,
		'f': ig.ImGuiKey_F,
		'g': ig.ImGuiKey_G,
		'h': ig.ImGuiKey_H,
		'i': ig.ImGuiKey_I,
		'j': ig.ImGuiKey_J,
		'k': ig.ImGuiKey_K,
		'l': ig.ImGuiKey_L,
		'm': ig.ImGuiKey_M,
		'n': ig.ImGuiKey_N,
		'o': ig.ImGuiKey_O,
		'p': ig.ImGuiKey_P,
		'q': ig.ImGuiKey_Q,
		'r': ig.ImGuiKey_R,
		's': ig.ImGuiKey_S,
		't': ig.ImGuiKey_T,
		'u': ig.ImGuiKey_U,
		'v': ig.ImGuiKey_V,
		'w': ig.ImGuiKey_W,
		'x': ig.ImGuiKey_X,
		'y': ig.ImGuiKey_Y,
		'z': ig.ImGuiKey_Z,
	}
	return mp[byte(kc)]
}

var g_timeout_add = cc.DlFunc[func(interval uint32, fn uptr, data uptr) uint32, uint32]{Name: "g_timeout_add"}

func timeoutAdd(ms uint32, fn func() (Continue bool)) uint32 {
	var fnp uptr
	fnp = cc.CbkRaw[func(uptr) uptr](func(out, ins uptr) {
		b := fn()
		if !b {
			cc.CbkCloseLate(fnp)
			*(*int32)(out) = 0
		} else {
			*(*int32)(out) = 1
		}
	})
	return g_timeout_add.Fn()(ms, fnp, nil)
}

func connect(obj gobject.GObjectIface, sig string, fn uptr) uint64 {
	return gobject.SignalConnectDataCfunc(obj, sig, fn, nil, func(data uptr) {
		cc.CbkClose(fn)
	}, 0)
}
