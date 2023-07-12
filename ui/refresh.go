package ui

import "fyne.io/fyne/v2"

func (app *Config) refreshTickerContent() {
	symbol, currentPrice := app.GetTickerText()
	app.TickerContainer.Objects = []fyne.CanvasObject{symbol, currentPrice}
	app.TickerContainer.Refresh()

	img := app.getImageAsset()
	app.PriceImgContainer.Objects = []fyne.CanvasObject{img}
	app.PriceImgContainer.Refresh()
}

func (app *Config) refreshHoldingsTable() {
	app.Holdings = app.GetHoldingsSlice()
	app.HoldingsTable.Refresh()
}
