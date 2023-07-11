package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/simon-lentz/marketwatcher/repo"
)

func (app *Config) connectDB() (*sql.DB, error) {
	path := ""

	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
		app.InfoLog.Println("DB located at:", path)
	} else {
		path = app.App.Storage().RootURI().Path() + "/sql.db"
		app.InfoLog.Println("DB located at:", path)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *Config) initDB(db *sql.DB) {
	app.DB = repo.NewSQLiteRepo(db)

	if err := app.DB.Migrate(); err != nil {
		app.ErrorLog.Println(err)
		log.Panic() // app will fail to run without db
	}
}
