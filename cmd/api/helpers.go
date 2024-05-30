package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"goweb/testapp"
	"net/http"
)

func toTinyInt(b bool) int8 {
	if b {
		return 1
	} else {
		return 0
	}
}

func (app *application) connectDB() (*testapp.Queries, context.Context, *sql.DB) {
	ctx := context.Background()

	db, dberr := sql.Open("mysql", "goweb:goweb@(mysql:3306)/goweb?parseTime=true")
	if dberr != nil {
		app.logger.Error(dberr.Error())
	}
	queries := testapp.New(db)

	return queries, ctx, db
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.Marshal(data)

	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	return nil
}
