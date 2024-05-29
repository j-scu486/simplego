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
	items, storeErr := queries.GetStoreItems(ctx, uint32(id))

	if storeErr != nil {
		msg := fmt.Sprintf("Could not find store with id %d", id)

		http.Error(w, msg, http.StatusNotFound)
		app.logger.Error(storeErr.Error())

		return
	}

	jsonErr := app.writeJSON(w, http.StatusOK, items, nil)

	if jsonErr != nil {
		app.logger.Error(jsonErr.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

	db.Close()
}

func (app *application) showStoreHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil {
		http.Error(w, "Invalid 'id'", http.StatusBadRequest)
		return
	}

	queries, ctx, db := app.connectDB()
	items, storeErr := queries.GetStore(ctx, uint32(id))

	if storeErr != nil {
		msg := fmt.Sprintf("Could not find store with id %d", id)

		http.Error(w, msg, http.StatusNotFound)
		app.logger.Error(storeErr.Error())

		return
	}

	jsonErr := app.writeJSON(w, http.StatusOK, items, nil)

	if jsonErr != nil {
		app.logger.Error(jsonErr.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

	db.Close()
}

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

	err = validate.Struct(store)

	if err != nil {
		app.writeJSON(w, http.StatusInternalServerError, err, nil)
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
