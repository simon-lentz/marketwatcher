package main

import (
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	_ "github.com/glebarez/go-sqlite" // embedded db for local storage
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

func main() {
	var myApp Config

	//create a fyne app
	fyneApp := app.NewWithID("placeholder")
	myApp.App = fyneApp
	myApp.HTTPClient = &http.Client{}

	// create loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to db
	db, err := myApp.connectDB()
	if err != nil {
		log.Panic(err) // if db can't be written then the app will fail to run
	}

	// create a db repo
	myApp.initDB(db)
	fyneApp.Preferences().StringWithFallback("Currency", "USD")

	// create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("MarketWatcher")
	myApp.MainWindow.Resize(fyne.Size{Height: 800, Width: 1000})
	myApp.MainWindow.SetFixedSize(true) // to be removed later for dynamic ui sizing
	myApp.MainWindow.SetMaster()

	myApp.makeUI()

	// show and run the application
	myApp.MainWindow.ShowAndRun()
}
