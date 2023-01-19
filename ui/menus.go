package ui

import (
	"errors"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/widget"
)

func BuildNewMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		sizeValidator := func(s string) error {
			width, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("must be a positive integer")
			}
			if width <= 0 {
				return errors.New("must be > 0")
			}
			return nil
		}
		widthEntry := widget.NewEntry()
		widthEntry.validator = sizeValidator

		heightEntry := widget.NewEntry()
		heightEntry.validator = sizeValidator

		widthFormEntry := widget.NewFormItem("Width", widthEntry)
		heightFormEntry := widget.NewFormItem("Height", heightEntry)

		formItems := []*widget.FormItem{widthFormEntry, heightFormEntry}

		dialog.ShowForm("New Image", "Create", "Cancel", formItems, func(ok bool) {
			if ok {
				pixelWidth := 0
				pixelHeight := 0

				if widthEntry.Validate() != nil {
					dialog.ShowError(errors.New("invalid width"), app.PixlWIndow)
				} else {
					pixelWidth, _ = strconv.Atoi(widthEntry.Text)
				}
				if heightEntry.Validate() != nil {
					dialog.ShowError(errors.New("invalid height"), app.PixlWIndow)
				} else {
					pixelHeight, _ = strconv.Atoi(heightEntry.Text)
				}
				app.PixlCanvas.NewDrawing(pixelHeight, pixelWidth)
			}
		})

	})
}
