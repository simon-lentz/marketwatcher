package main

import "fyne.io/fyne/v2"

func (app *Config) refreshTickerContent() {
	symbol, currentPrice := app.getTickerText()
	app.TickerContainer.Objects = []fyne.CanvasObject{symbol, currentPrice}
	app.TickerContainer.Refresh()

	img := app.getImageAsset()
	app.PriceImgContainer.Objects = []fyne.CanvasObject{img}
	app.PriceImgContainer.Refresh()
}
