package ui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/simon-lentz/marketwatcher/repo"
)

func (app *Config) holdingsTab() *fyne.Container {
	app.Holdings = app.GetHoldingsSlice()
	app.HoldingsTable = app.GetHoldingsTable()

	holdingsCtr := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, app.HoldingsTable),
	)
	// the container has to allow for the central element (the table) to expand so we have to use Border and prob GridWrap too

	return holdingsCtr
}

func (app *Config) GetHoldingsTable() *widget.Table {
	table := widget.NewTable(
		func() (int, int) { // get table dimensions
			return len(app.Holdings), len(app.Holdings[0])
		},
		func() fyne.CanvasObject { // create container object template
			cellCtr := container.NewVBox(widget.NewLabel(""))
			return cellCtr
		},
		func(i widget.TableCellID, o fyne.CanvasObject) { // populate individual cells
			if i.Col != (len(app.Holdings[0])-1) || i.Row == 0 {
				// populate cell with textual info
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(app.Holdings[i.Row][i.Col].(string)),
				}
			} else {
				// last cell - put in the button
				btn := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Delete?", "", func(deleted bool) {
						if deleted {
							id, _ := strconv.Atoi(app.Holdings[i.Row][0].(string)) // get id of delete target
							if err := app.DB.DeleteHolding(int64(id)); err != nil {
								app.ErrorLog.Println(err)
							}
						}
						app.refreshHoldingsTable() // refresh holding table after mutation
					}, app.MainWindow)
				})
				btn.Importance = widget.HighImportance // .ButtonImportance(int) ??
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					btn,
				} // cast current cell to the container object
			}
		})

	// since the cols are not uniform the width has to be specified
	colWidths := []float32{50, 200, 200, 200, 110}
	for i := 0; i < len(colWidths); i++ {
		table.SetColumnWidth(i, colWidths[i])
	}
	return table
}

// build the data structure to contain all the db entries
func (app *Config) GetHoldingsSlice() [][]interface{} {
	var slice [][]interface{}
	holdings, err := app.CurrentHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	// this allows for the addition of the "Delete?" button to mutate a holding in-app
	slice = append(slice, []interface{}{"ID", "Units", "Value", ""})
	for _, h := range holdings {
		var currentRow []interface{} // tried to use any here but it didn't work ??

		currentRow = append(currentRow, strconv.FormatInt(h.ID, 10))
		currentRow = append(currentRow, fmt.Sprintf("%d units", h.Units))
		currentRow = append(currentRow, fmt.Sprintf("$%.2f", float64(h.Value)))
		currentRow = append(currentRow, widget.NewButton("Delete", func() {}))

		slice = append(slice, currentRow)
	}
	return slice
}

// accesses the local db
func (app *Config) CurrentHoldings() ([]repo.Holding, error) {
	holdings, err := app.DB.AllHoldings()
	if err != nil {
		return nil, err
	}
	return holdings, nil
}
