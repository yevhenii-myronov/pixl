package ui

import (
	"fyne.io/fyne/v2"
	"github.com/yevhenii-myronov/pixl/apptype"
	"github.com/yevhenii-myronov/pixl/pxcanvas"
	"github.com/yevhenii-myronov/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
