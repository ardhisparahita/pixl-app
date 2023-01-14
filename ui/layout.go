package ui

func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)

	app.PixlWIndow.SetContent(swatchesContainer)
}