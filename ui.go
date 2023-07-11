package main

import (
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUI() {
	// get ticker price
	currentPrice, symbol := app.getTickerText()

	// put price information in container
	tickerContent := container.NewGridWithColumns(
		2,
		symbol,
		currentPrice,
	)

	app.TickerContainer = tickerContent

	// get toolbar
	toolbar := app.getToolbar()
	app.Toolbar = toolbar

	// get app tabs
	priceTabContent := app.pricesTab()
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent), // change default icons??
		container.NewTabItemWithIcon("Holdings", theme.InfoIcon(), canvas.NewText("Holdings Info Here", nil)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// add container to window
	finalContent := container.NewVBox(tickerContent, toolbar, tabs)

	app.MainWindow.SetContent(finalContent)

	// update tickers every thirty seconds
	go func() {
		for range time.Tick(time.Second * 30) {
			app.refreshTickerContent()
		}
	}()
}
