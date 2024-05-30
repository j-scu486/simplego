package main

import (
	"fmt"
	"goweb/testapp"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) showStoreItemsHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil {
		http.Error(w, "Invalid 'id'", http.StatusBadRequest)
		return
	}

	queries, ctx, db := app.connectDB()
	defer db.Close()
	items, storeErr := queries.GetStoreItems(ctx, uint32(id))

	if storeErr != nil {
		msg := fmt.Sprintf("Could not find store with id %d", id)

		http.Error(w, msg, http.StatusNotFound)
		app.logger.Error(storeErr.Error())

		return
	}

	if jsonErr := app.writeJSON(w, http.StatusOK, items, nil); jsonErr != nil {
		app.logger.Error(jsonErr.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (app *application) showStoreHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	queries, ctx, db := app.connectDB()
	defer db.Close()
	items, storeErr := queries.GetStore(ctx, uint32(id))

	if storeErr != nil {
		msg := fmt.Sprintf("Could not find store with id %d", id)

		http.Error(w, msg, http.StatusNotFound)
		app.logger.Error(storeErr.Error())

		return
	}

	if jsonErr := app.writeJSON(w, http.StatusOK, items, nil); jsonErr != nil {
		app.logger.Error(jsonErr.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (app *application) createStoreHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	store := &Store{
		Name:  r.FormValue("name"),
		Owner: r.FormValue("owner"),
	}

	if err := validate.Struct(store); err != nil {
		app.logger.Error(err.Error())
		app.writeJSON(w, http.StatusBadRequest, err, nil)
		return
	}

	queries, ctx, db := app.connectDB()
	defer db.Close()

	err := queries.CreateStore(ctx, testapp.CreateStoreParams{
		Name:  store.Name,
		Owner: testapp.StoresOwner(store.Owner),
	})

	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Could not create store", http.StatusInternalServerError)
		return
	}

	app.writeJSON(w, http.StatusCreated, store, nil)
}
