package ui

import (
	"ardhisparahita.io/pixl/apptype"
	"ardhisparahita.io/pixl/pxcanvas"
	"ardhisparahita.io/pixl/swatch"
	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWIndow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}