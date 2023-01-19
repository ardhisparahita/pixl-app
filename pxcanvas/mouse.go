package pxcanvas

import (
	"ardhisparahita.io/pixl/pxcanvas/brush"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (PxCanvas *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	PxCanvas.scale(int(ev.Scrolled.DY))
	PxCanvas.Refresh()
}

func (PxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	if x, y := PxCanvas.MouseToCanvasXY(ev); x != nil && y != nil {
		brush.TryBrush(PxCanvas.appState, PxCanvas, ev)
		cursor := brush.Cursor(PxCanvas.PxCanvasConfig, PxCanvas.appState.Brushtype, ev, *x, *y)
		PxCanvas.renderer.SetCursor(cursor)
	} else {
		PxCanvas.renderer.SetCursor(make([]fyne.CanvasObject, 0))
	}
	PxCanvas.TryPan(PxCanvas.mouseState.previousCoord, ev)
	PxCanvas.Refresh()
	PxCanvas.mouseState.previousCoord = &ev.PointEvent
}

func (PxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {

}

func (PxCanvas *PxCanvas) MouseOut() {

}

func (PxCanvas *PxCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(PxCanvas.appState, PxCanvas, ev)
}

func (PxCanvas *PxCanvas) MouseUp(ev *desktop.MouseEvent) {

}
