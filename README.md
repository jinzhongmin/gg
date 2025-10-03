# GG - GTK4 GUI Framework

A  GTK4 GUI framework for Go that leverages libffi and minimal CGO to achieve significantly faster compilation times compared to traditional Go GTK bindings.

## Environment Setup

### Windows with MSYS2

Add the MSYS2 UCRT bin directory to your system PATH to ensure all required DLLs and executables are accessible, and ensure you have the following dependencies installed via MSYS2 pacman:

``` bash
pacman -S mingw-w64-ucrt-x86_64-toolchain mingw-w64-ucrt-x86_64-gtk4 mingw-w64-ucrt-x86_64-mimalloc
```

## Example

### Hello world

``` go
package main

import (
    "unsafe"

    "github.com/jinzhongmin/gg/glib/gio"
    "github.com/jinzhongmin/gg/gtk/gtk/v4"
)

func main() {
    app := gtk.NewApplication("com.github.jinzhongmin.helloworld", gio.GApplicationFlagsDefaultFlags)
    app.ConnectActivate(func(a *gio.GApplication) {
        app := (*gtk.Application)(unsafe.Pointer(a))

        window := gtk.NewApplicationWindow(app)
        button := gtk.NewButtonWithLabel("hello world")
        button.ConnectClicked(func(b *gtk.Button) {
            window.Close()
        })

        window.SetChild(button)
        window.Present()
    })
    app.Run(nil)
}

```

### Build

``` bash
go build -tags ccrelease  # Using ccrelease tag to minimize reflect usage for potentially better performance
```

## Basic Usage

``` go
cc.Open("libgtk-4.dll") // Manually load the GTK dynamic library if you encounter failures when calling GTK functions
```

If you find that the existing APIs don't match your local GTK APIs, you can manually implement the required APIs:

```go

// *gtk.Label is equivalent to the corresponding C pointer - Go pointers are unwrapped raw pointers

// cc.DlFunc[Go function prototype, return type of Go function]{Name: symbol}
// Use cc.Void if the return type is void
var gtk_label_new = cc.DlFunc[func(str cc.String) *gtk.Label, *gtk.Label]{Name: "gtk_label_new"}

func newLabel(str string) *gtk.Label {
    cstr := cc.NewString(str) // string can be replaced with cc.String
    defer cstr.Free()
    return gtk_label_new.Fn()(cstr)
}
```

How to Bind Go Variable Lifetime with Corresponding C-side Variable Lifetime

``` go
package main

import (
    "fmt"
    "unsafe"

    "github.com/jinzhongmin/gg/glib/gio"
    "github.com/jinzhongmin/gg/gtk/gtk/v4"
)

type MyButton struct {
    *gtk.Button
    cur int
}

func newMyButton() *MyButton {
    btn := new(MyButton)
    btn.Button = gtk.NewButton()
    btn.cur = 0
    btn.SetLabel(fmt.Sprintf("%d", btn.cur))

    btn.Button.Pin(btn) // Bind btn's lifetime to btn.Button - it will exist as long as btn.Button exists
    return btn
}

func (btn *MyButton) Inc() {
    btn.cur += 1
    btn.SetLabel(fmt.Sprintf("%d", btn.cur))
}

func main() {
    app := gtk.NewApplication("com.github.jinzhongmin.helloworld", gio.GApplicationFlagsDefaultFlags)
    app.ConnectActivate(func(a *gio.GApplication) {
        app := (*gtk.Application)(unsafe.Pointer(a))

        window := gtk.NewApplicationWindow(app)

        button := newMyButton() // button lifetime begins
        button.ConnectClicked(func(b *gtk.Button) {
            obj, ok := b.GetPinned() // Use this to retrieve the Go variable pinned by gtk.Button
            if !ok {
                return
            }
            btn := obj.(*MyButton)
            btn.Inc()
        })
        // Although button's lifetime would normally end here, the Pin() call in newMyButton() 
        // ensures the variable persists until gtk.Button is destroyed

        window.SetChild(button)
        window.Present()
    })
    app.Run(nil)
}
```

How to Use ImGui in GTK Applications

First, compile [libdcimgui](https://github.com/jinzhongmin/libdcimgui) to obtain the dynamic library file, then place it in either your application directory or system PATH.

```go
package main

import (
    "math"
    "unsafe"

    "github.com/jinzhongmin/gg/cc"
    "github.com/jinzhongmin/gg/dcimgui"
    ig "github.com/jinzhongmin/gg/dcimgui"
    plt "github.com/jinzhongmin/gg/dcimgui/dcimplot"
    "github.com/jinzhongmin/gg/glib/gio"
    "github.com/jinzhongmin/gg/gtk/gtk/v4"
    "github.com/jinzhongmin/gg/gtk/gtk/v4/imgtk"
)

type uptr = unsafe.Pointer

// Convert Go string to ImGui required type
// Example: sr("hello\x00") - requires \x00 termination
// This approach avoids frequent CGO string conversion memory allocations
func sr(str string) cc.String { return *(*cc.String)(uptr(&str)) }

func build[T any](obj T, builders ...func(T)) T {
    for _, fn := range builders {
        fn(obj)
    }
    return obj
}

func main() {

    app := gtk.NewApplication("com.github.jinzhongmin.imgtk", 0)
    app.ConnectActivate(func(a *gio.GApplication) {
        build(gtk.NewWindow(), func(win *gtk.Window) { // Window self setup

            win.SetApplication(app)
            win.SetTitle("mcv")
            win.SetDefaultSize(900, 600)

        }, func(win *gtk.Window) { // Window layout

            var font *ig.ImFont
            var implotCtx *plt.ImPlotContext

            var sin cc.Func // Corresponds to C callback function
            var cos uptr    // uptr and cc.Func are essentially equivalent

            // C prototype: ImPlotPoint (*ImPlotGetter)(int idx, void* user_data);
            // The following function binds sin to a Go function matching the C prototype
            // This binding uses reflect in generated code, which may have slightly lower performance
            sin.Bind(func(idx int32, _ uptr) (pt plt.ImPlotPoint) {
                x := float64(math.Pi) / 100 * float64(idx)
                y := math.Sin(x)

                pt.X = x
                pt.Y = y
                return
            })

            // C prototype: ImPlotPoint (*ImPlotGetter)(int idx, void* user_data);
            // The following function binds cos to a Go function matching the C prototype
            // This binding does not use reflect in generated code
            cos = cc.CbkRaw[func(idx int32, _ uptr) plt.ImPlotPoint](func(out, ins uptr) {
                // Input parameters need manual conversion:
                // out: *ImPlotPoint
                // ins: []unsafe.Pointer{*int32, *unsafe.Pointer}

                pt := (*plt.ImPlotPoint)(out)

                inPtrs := unsafe.Slice((*uptr)(ins), 2)
                idx := *(*int32)(inPtrs[0])

                x := float64(math.Pi) / 100 * float64(idx)
                y := math.Cos(x)

                pt.X = x
                pt.Y = y
            })

            vbox := build(gtk.NewBox(gtk.OrientationVertical, 0), func(vbox *gtk.Box) { // vbox configuration

                vbox.SetHExpand(true)
                vbox.SetVExpand(true)

            }, func(vbox *gtk.Box) { // Child 1: ImArea

                var customColor plt.ImPlotColormap

                // For C-side char** parameter types, use cc.Strings with pre-allocation
                labels := cc.NewStrings([]string{
                    "PASS", "NG",
                })

                // Use imgtk.NewImArea to create an ImGui control in GTK
                ia := build(imgtk.NewImArea(func(ia *imgtk.ImArea) {
                    // ImGui rendering process below

                    if customColor == 0 {
                        customColor = plt.AddColormapImVec4(sr("pn\x00"), []ig.ImVec4{
                            {93.0 / 255.0, 197.0 / 255.0, 108.0 / 255.0, 1},
                            {1, 0, 0, 1},
                        }, true)
                    }

                    dcimgui.StyleColorsLight(nil)
                    plt.SetCurrentContext(implotCtx)

                    if ig.Begin(sr("PieChart example\x00"), nil, 0) {

                        if plt.BeginPlot(sr("Production Line PASS Statistics\x00"), ig.ImVec2{-1, -1}, 0) {

                            plt.SetupAxis(plt.ImAxis_X1, 0,
                                plt.ImPlotAxisFlags_NoLabel|plt.ImPlotAxisFlags_NoTickMarks|plt.ImPlotAxisFlags_NoTickLabels)
                            plt.SetupAxesLimits(-110, 110, -110, 110, plt.ImPlotCond_Always)
                            plt.SetupAxisFormat(plt.ImAxis_X1, sr("\x00"))
                            plt.SetupAxisFormat(plt.ImAxis_Y1, sr("\x00"))
                            plt.SetupLegend(plt.ImPlotLocation_SouthEast,
                                plt.ImPlotLegendFlags_Outside|plt.ImPlotLegendFlags_Horizontal)

                            plt.PushColormap(customColor)
                            ig.PushFontSize(nil, 18)
                            plt.PlotPieChart(labels, ig.ImGuiDataType_Double, []float64{0.8 * 100, 0.1 * 100}, 0, 0, 100, sr("%.1f%%\x00"), 90, 0)
                            ig.PopFont()
                            plt.PopColormap(1)

                            plt.EndPlot()
                        }

                    }
                    ig.End()

                    if ig.Begin(sr("callback\x00"), nil, 0) {
                        if plt.BeginPlot(sr("callback\x00"), ig.ImVec2{-1, -1}, 0) {
                            plt.PlotLineG(sr("sin\x00"), sin, 200, 0, nil)
                            plt.PlotLineG(sr("cos\x00"), cc.Func(cos), 200, 0, nil)
                            plt.EndPlot()
                        }
                    }
                    ig.End()

                }), func(ia *imgtk.ImArea) {
                    ia.ConnectImInit(func(igc *ig.ImGuiContext, igi *ig.ImGuiIO) {
                        // Load fonts during ImGui initialization
                        font = igi.Fonts.AddFontFromFileTTF(sr("./DingTalk_JinBuTi.ttf\x00"), 16.5, nil, nil)
                        
                        // ImPlot context needs to be created manually
                        implotCtx = plt.CreateContext()
                    })
                    ia.ConnectImDestroy(func(igc *ig.ImGuiContext, igi *ig.ImGuiIO) {
                        // Cleanup during destruction
                        if implotCtx != nil {
                            implotCtx.Destroy()
                        }
                        if font != nil {
                            igi.Fonts.RemoveFont(font)
                        }

                    })

                })
                ia.SetHExpand(true)
                ia.SetVExpand(true)
                vbox.Append(ia)
            })
            vbox.SetHExpand(true)
            vbox.SetVExpand(true)
            win.SetChild(vbox)

        }).Present()

    })
    app.Run(nil)
}
``
