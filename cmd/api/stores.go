package main

import (
	"goweb/testapp"
	"log"
	"net/http"
)

func (app *application) createStoreHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	store := &Store{
		Name:  r.FormValue("name"),
		Owner: r.FormValue("owner"),
	}

	err = validateStore(store)

	if err != nil {
		log.Fatal("Invalid Store")
		return
	}

	queries, ctx, db := app.connectDB()

	err = queries.CreateStore(ctx, testapp.CreateStoreParams{
		Name:  store.Name,
		Owner: testapp.StoresOwner(store.Owner),
	})

	if err != nil {
		app.writeJSON(w, http.StatusInternalServerError, err, nil)
	}

	db.Close()
}
