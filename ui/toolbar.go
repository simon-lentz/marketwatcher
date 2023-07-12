package ui

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/simon-lentz/marketwatcher/repo"
)

var currency = "USD"

func (app *Config) GetToolbar() *widget.Toolbar {
	toolbar := widget.NewToolbar( // We don't have to pass fyne.Window to the func because it is accessed by the new toolbar funcs
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			app.AddHoldingsDialog()
		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			app.refreshTickerContent()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			preferencesWindow := app.showPreferences()
			preferencesWindow.Resize(fyne.NewSize(500, 400))
			preferencesWindow.Show()
		}),
	)

	return toolbar
}

func (app *Config) AddHoldingsDialog() dialog.Dialog {
	addUnitsEntry := widget.NewEntry()
	addValueEntry := widget.NewEntry()

	app.HoldingUnits = addUnitsEntry
	app.HoldingValue = addValueEntry

	/*
		dateValidator := func(d string) error {
			if _, error := time.Parse("2023-07-11", d); err != nil {
				return err
			}
			return nil
		}

		addDateEntry.Validator = dateValidator
	*/
	intValidator := func(s string) error { // maybe use a generic ordered check or something
		if _, err := strconv.Atoi(s); err != nil {
			return err
		}
		return nil
	}
	addUnitsEntry.Validator = intValidator

	floatValidator := func(s string) error {
		if _, err := strconv.ParseFloat(s, 32); err != nil {
			return err
		}
		return nil
	}
	addValueEntry.Validator = floatValidator

	addForm := dialog.NewForm(
		"Add Holding?",
		"Add",
		"Cancel",
		[]*widget.FormItem{
			{Text: "Units Purchased", Widget: addUnitsEntry},
			{Text: "Holding Value", Widget: addValueEntry},
		},
		func(valid bool) {
			if valid {
				units, _ := strconv.Atoi(addUnitsEntry.Text)
				value, _ := strconv.ParseFloat(addValueEntry.Text, 32)

				if _, err := app.DB.InsertHolding(repo.Holding{
					Units: int64(units), // change that later
					Value: value,
				}); err != nil {
					app.ErrorLog.Println(err)
				}
				app.refreshHoldingsTable()
			}
		}, app.MainWindow)

	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()

	return addForm
}

func (app *Config) showPreferences() fyne.Window {
	w := app.App.NewWindow("Preferences") // realistically using a new window for preferences is dumb just wanted to see if it would work, clicking it opens more and more new windows
	label := widget.NewLabel("Preferred Currency")
	cur := widget.NewSelect([]string{"USD", "YUAN", "GBP", "EURO"}, func(value string) {
		currency = value
		app.App.Preferences().SetString("Currency", value)
	})
	cur.Selected = currency

	btn := widget.NewButton("Save", func() {
		w.Close()
		app.refreshHoldingsTable()
	})
	btn.Importance = widget.HighImportance

	w.SetContent(container.NewVBox(label, cur, btn))

	return w
}
