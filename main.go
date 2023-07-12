package main

import (
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	_ "github.com/glebarez/go-sqlite" // embedded db for local storage
	ui "github.com/simon-lentz/marketwatcher/ui"
)

func main() {
	var myApp ui.Config

	//create a fyne app
	fyneApp := app.NewWithID("placeholder")
	myApp.App = fyneApp
	myApp.HTTPClient = &http.Client{}

	// create loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to db
	db, err := myApp.ConnectDB()
	if err != nil {
		log.Panic(err) // if db can't be written then the app will fail to run
	}

	// create a db repo
	myApp.InitDB(db)
	fyneApp.Preferences().StringWithFallback("Currency", "USD")

	// create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("MarketWatcher")
	myApp.MainWindow.Resize(fyne.Size{Height: 800, Width: 1000})
	myApp.MainWindow.SetFixedSize(true) // to be removed later for dynamic ui sizing
	myApp.MainWindow.SetMaster()

	myApp.MakeUI()

	// show and run the application
	myApp.MainWindow.ShowAndRun()
}
