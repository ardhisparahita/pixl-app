package main

import (
	"image/color"

	"ardhisparahita.io/pixl/apptype"
	"ardhisparahita.io/pixl/swatch"
	"ardhisparahita.io/pixl/ui"
	"fyne.io/fyne/v2/app"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixlWIndow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlWIndow.ShowAndRun()
}
