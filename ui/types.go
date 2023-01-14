package ui

import (
	"fyne.io/fyne/v2"
	"ardhisparahita.io/pixl/apptype"
	"ardhisparahita.io/pixl/swatch"
)

type AppInit struct {
	PixlWIndow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}