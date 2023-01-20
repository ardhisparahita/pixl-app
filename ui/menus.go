package ui

import (
	"errors"
	"image/png"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func saveFileDialog(app *AppInit) {
	dialog.ShowFileSave(func(uri fyne.URIWriteCloser, e error) {
		if uri == nil {
			return
		} else {
			err := png.Encode(uri, app.PixlCanvas.PixelData)
			if err != nil {
				dialog.ShowError(err, app.PixlWIndow)
				return
			}
			app.State.SetFilePath(uri.URI().Path())
		}
	}, app.PixlWIndow)
}

func BuildSaveAsMenus(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save as...", func() {
		saveFileDialog(app)
	})
}

func BuildSaveMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		if app.State.FilePath == "" {
			saveFileDialog(app)
		} else {
			tryClose := func(fh *os.File) {
				err := fh.Close()
				if err != nil {
					dialog.ShowError(err, app.PixlWIndow)
				}

			}
			fh, err := os.Create(app.State.FilePath)
			defer tryClose(fh)

			if err != nil {
				dialog.ShowError(err, app.PixlWIndow)
				return
			}
			err = png.Encode(fh, app.PixlCanvas.PixelData)
			if err != nil {
				dialog.ShowError(err, app.PixlWIndow)
				return
			}
		}
	})
}

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
		widthEntry.Validator = sizeValidator

		heightEntry := widget.NewEntry()
		heightEntry.Validator = sizeValidator

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
				app.PixlCanvas.NewDrawing(pixelWidth, pixelHeight)
			}
		}, app.PixlWIndow)
	})
}

func BuildMenus(app *AppInit) *fyne.Menu {
	return fyne.NewMenu(
		"File",
		BuildNewMenu(app),
		BuildSaveMenu(app),
		BuildSaveAsMenus(app),
	)
}

func SetupMenus(app *AppInit) {
	menus := BuildMenus(app)
	mainMenu := fyne.NewMainMenu(menus)
	app.PixlWIndow.SetMainMenu(mainMenu)
}
