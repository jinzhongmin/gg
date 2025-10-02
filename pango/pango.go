package pango

import (
	"github.com/jinzhongmin/gg/glib/gobject"
)

type Context struct{ gobject.GObject }
type FontFace struct{ gobject.GObject }
type FontFamily struct{ gobject.GObject }
type Layout struct{ gobject.GObject }
type FontMap struct{ gobject.GObject }

type AttrList struct{}
type FontDescription struct{}
type TabArray struct{}
type Language struct{}
