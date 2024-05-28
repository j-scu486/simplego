package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/items", app.showItemHandler)
	router.HandlerFunc(http.MethodPost, "/item", app.createItemHandler)

	// router.HandlerFunc(http.MethodPost, "/store", app.createStoreHandler)
	// router.HandlerFunc(http.MethodGet, "/store/:id", app.showStoreHandler)
	// router.HandlerFunc(http.MethodGet, "/store/:id/items", app.showStoreItemsHandler)

	return router
}
