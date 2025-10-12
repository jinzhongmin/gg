package pango

import (
	"unsafe"

	"github.com/jinzhongmin/gg/glib/gobject"
)

type uptr = unsafe.Pointer
type Context struct{ gobject.GObject }
type FontFace struct{ gobject.GObject }
type FontFamily struct{ gobject.GObject }
type Layout struct{ gobject.GObject }
type FontMap struct{ gobject.GObject }

type AttrList struct{}
type FontDescription struct{}
type TabArray struct{}
type Language struct{}
type LayoutRun struct{}

type LayoutLine struct {
	Layout     *Layout
	StartIndex int32
	Length     int32
	Runs       *LayoutRun
	_          uint32
}

type GlyphItem struct {
	Item         *Item
	Glyphs       *GlyphString
	YOffset      int32
	StartXOffset int32
	EndXOffset   int32
}

type GlyphString struct {
	NumGlyphs   int32
	Glyphs      GlyphInfo
	LogClusters *int32
}

type Glyph uint32

type GlyphInfo struct {
	Glyph    Glyph
	Geometry GlyphGeometry
	Attr     GlyphVisAttr
}

type GlyphUnit int32
type GlyphGeometry struct {
	Width   GlyphUnit
	XOffset GlyphUnit
	YOffset GlyphUnit
}
type GlyphVisAttr struct {
	flags uint32
}

type Item struct {
	Offset   int32
	Length   int32
	NumChars int32
	Analysis Analysis
}

type Analysis struct {
	ShapeEngine uptr
	LangEngine  uptr
	Font        *Font
	Level       uint8
	Gravity     uint8
	Flags       uint8
	Script      uint8
	Language    *Language
	ExtraAttrs  uptr
}

type Font struct{ gobject.GObjectCore }
