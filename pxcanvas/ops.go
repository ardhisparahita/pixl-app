package pxcanvas

import "fyne.io/fyne/v2"

func (PxCanvas *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDiff := currentCoord.Position.X - previousCoord.Position.X
	yDiff := currentCoord.Position.Y - previousCoord.Position.Y

	PxCanvas.CanvasOffset.X += xDiff
	PxCanvas.CanvasOffset.Y += yDiff
	PxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) scale(direction int) {
	switch {
	case direction > 0:
		pxCanvas.PxSize += 1
	case direction < 0:
		if pxCanvas.PxSize > 2 {
			pxCanvas.PxSize -= 1
		}
	default:
		pxCanvas.PxSize = 10
	}
}
