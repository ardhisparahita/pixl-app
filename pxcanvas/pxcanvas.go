package pxcanvas

import (
	"image"
	"image/color"

	"ardhisparahita.io/pixl/apptype"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer    *PxCanvasRenderer
	PixelData   image.Image
	mouseState  PxCanvasMouseState
	appState    *apptype.State
	reloadImage bool
}

func (PxCanvas *PxCanvas) Bounds() image.Rectangle {
	x0 := int(PxCanvas.CanvasOffset.X)
	y0 := int(PxCanvas.CanvasOffset.Y)
	x1 := int(PxCanvas.PxCols*PxCanvas.PxSize + int(PxCanvas.CanvasOffset.X))
	y1 := int(PxCanvas.PxRows*PxCanvas.PxSize + int(PxCanvas.CanvasOffset.Y))

	return image.Rect(x0, y0, x1, y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= float32(bounds.Min.X) &&
		pos.X < float32(bounds.Max.X) &&
		pos.Y >= float32(bounds.Min.Y) &&
		pos.Y < float32(bounds.Max.Y) {
		return true
	}
	return false
}

func NewBlankImage(cols, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			img.Set(x, y, c)
		}
	}
	return img
}

func NewPxCanvas(state *apptype.State, config apptype.PxCanvasConfig) *PxCanvas {
	PxCanvas := &PxCanvas{
		PxCanvasConfig: config,
		appState:       state,
	}
	PxCanvas.PixelData = NewBlankImage(config.PxCols, config.PxRows, color.NRGBA{128, 128, 128, 255})
	PxCanvas.ExtendBaseWidget(PxCanvas)
	return PxCanvas
}

func (pxCanvas *PxCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &PxCanvasRenderer{
		pxCanvas:     pxCanvas,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
	}
	pxCanvas.renderer = renderer
	return renderer
}

func (PxCanvas *PxCanvas) TryPan(previousCoord *fyne.PointEvent, ev *desktop.MouseEvent) {
	if previousCoord != nil && ev.Button == desktop.MouseButtonTertiary {
		PxCanvas.Pan(*previousCoord, ev.PointEvent)
	}
}

// Brushable interface
func (PxCanvas *PxCanvas) SetColor(c color.Color, x, y int) {
	if nrgba, ok := PxCanvas.PixelData.(*image.NRGBA); ok {
		nrgba.Set(x, y, c)
	}
	if rgba, ok := PxCanvas.PixelData.(*image.RGBA); ok {
		rgba.Set(x, y, c)
	}
	PxCanvas.Refresh()
}

func (PxCanvas *PxCanvas) MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int) {
	bounds := PxCanvas.Bounds()
	if !InBounds(ev.Position, bounds) {
		return nil, nil
	}

	pxSize := float32(PxCanvas.PxSize)
	xOffset := PxCanvas.CanvasOffset.X
	yOffset := PxCanvas.CanvasOffset.Y

	x := int((ev.Position.X - xOffset) / pxSize)
	y := int((ev.Position.Y - yOffset) / pxSize)

	return &x, &y
}

func (PxCanvas *PxCanvas) LoadImage(img image.Image) {
	dimensions := img.Bounds()

	PxCanvas.PxCanvasConfig.PxCols = dimensions.Dx()
	PxCanvas.PxCanvasConfig.PxRows = dimensions.Dy()

	PxCanvas.PixelData = img
	PxCanvas.reloadImage = true
	PxCanvas.Refresh()
}

func (PxCanvas *PxCanvas) NewDrawing(cols, rows int) {
	PxCanvas.appState.SetFilePath("")
	PxCanvas.PxCols = cols
	PxCanvas.PxRows = rows
	pixelData := NewBlankImage(cols, rows, color.NRGBA{128, 128, 128, 255})
	PxCanvas.LoadImage(pixelData)
}
