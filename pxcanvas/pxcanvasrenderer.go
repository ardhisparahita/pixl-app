package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	PxCanvas     *PxCanvas
	canvasImage  *canvas.Image
	canvasBorder []canvas.Line
}

// WidgetRenderer interface implementation
func (renderer *PxCanvasRenderer) MinSize() fyne.Size {
	return renderer.PxCanvas.DrawingArea
}

// WidgetRenderer interface implementation
func (renderer *PxCanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5)
	for i := 0; i < len(renderer.canvasBorder); i++ {
		objects = append(objects, &renderer.canvasBorder[i])
	}
	objects = append(objects, renderer.canvasImage)
	return objects
}

// WidgetRenderer interface implementation
func (renderer *PxCanvasRenderer) Destroy() {

}

// WidgetRenderer interface implementation
func (renderer *PxCanvasRenderer) Layout(size fyne.Size) {
	renderer.LayoutCanvas(size)
	renderer.Layoutborder(size)
}

// WidgetRenderer interface implementation
func (renderer *PxCanvasRenderer) Refresh() {
	if renderer.PxCanvas.reloadImage {
		renderer.canvasImage = canvas.NewImageFromImage(renderer.PxCanvas.pixelData)
		renderer.canvasImage.ScaleMode = canvas.ImageScalePixels
		renderer.canvasImage.FillMode = canvas.ImageFillContain
		renderer.PxCanvas.reloadImage = false
	}
	renderer.Layout(renderer.PxCanvas.Size())
	canvas.Refresh(renderer.canvasImage)
}

func (renderer *PxCanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgPxWidth := renderer.PxCanvas.PxCols
	imgPxHeight := renderer.PxCanvas.PxRows
	pxSize := renderer.PxCanvas.PxSize
	renderer.canvasImage.Move(fyne.NewPos(renderer.PxCanvas.CanvasOffset.X, renderer.PxCanvas.CanvasOffset.Y))
	renderer.canvasImage.Resize(fyne.NewSize(float32(imgPxWidth*pxSize), float32(imgPxHeight*pxSize)))
}

func (renderer *PxCanvasRenderer) Layoutborder(size fyne.Size) {
	offset := renderer.PxCanvas.CanvasOffset
	imgHeight := renderer.canvasImage.Size().Height
	imgWidth := renderer.canvasImage.Size().Width

	left := &renderer.canvasBorder[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imgHeight)

	top := &renderer.canvasBorder[1]
	top.Position1 = fyne.NewPos(offset.X, offset.Y)
	top.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y)

	right := &renderer.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X+imgWidth, offset.Y)
	right.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)

	bottom := &renderer.canvasBorder[3]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y+imgHeight)
	bottom.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)
}
