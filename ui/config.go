package ui

import (
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/simon-lentz/marketwatcher/repo"
)

type Config struct {
	App               fyne.App
	InfoLog           *log.Logger
	ErrorLog          *log.Logger
	DB                repo.Repo
	MainWindow        fyne.Window
	TickerContainer   *fyne.Container // update with refresh button
	Toolbar           *widget.Toolbar
	PriceImgContainer *fyne.Container // update with refresh button
	Holdings          [][]interface{}
	HoldingsTable     *widget.Table
	HoldingUnits      *widget.Entry
	HoldingValue      *widget.Entry
	HTTPClient        *http.Client
}
