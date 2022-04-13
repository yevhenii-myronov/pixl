package ui

import (
	"fyne.io/fyne/v2"
	"github.com/wizard2014/pixl/apptype"
	"github.com/wizard2014/pixl/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
