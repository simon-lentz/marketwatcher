package main

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) getToolbar() *widget.Toolbar {
	toolbar := widget.NewToolbar( // We don't have to pass fyne.Window to the func because it is accessed by the new toolbar funcs
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			app.refreshTickerContent()
		}),
	)

	return toolbar
}
