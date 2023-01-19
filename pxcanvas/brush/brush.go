package brush

import (
	"ardhisparahita.io/pixl/apptype"
	"fyne.io/fyne/v2/driver/desktop"
)

const (
	Pixel = iota
)

func TryBrush(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	switch {
	case appState.Brushtype == Pixel:
		return TryPaintPixel(appState, canvas, ev)
	default:
		return false
	}
}

func TryPaintPixel(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	x, y := canvas.MouseToCanvasXY(ev)
	if x != nil && y != nil && ev.Button == desktop.MouseButtonPrimary {
		canvas.SetColor(appState.BrushColor, *x, *y)
		return true
	}
	return false
}
