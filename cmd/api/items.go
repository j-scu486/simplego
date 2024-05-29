package main

import (
	"fmt"
	"goweb/testapp"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func (app *application) createItemHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	name := r.PostFormValue("name")
	priceStr := r.PostFormValue("price")
	quantityStr := r.PostFormValue("quantity")
	onSaleStr := r.PostFormValue("on_sale")

	if name == "" || priceStr == "" || quantityStr == "" || onSaleStr == "" {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	onSale, err := strconv.ParseBool(onSaleStr)
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	item := &Item{
		Name:     name,
		Price:    price,
		Quantity: quantity,
		OnSale:   onSale,
	}

	err = validateItem(item)

	if err != nil {
		log.Fatal("Invalid Item")
		return
	}

	queries, ctx, db := app.connectDB()

	err = queries.CreateItem(ctx, testapp.CreateItemParams{
		Name:     item.Name,
		Price:    item.Price,
		Quantity: uint32(item.Quantity),
		Onsale:   toTinyInt(onSale),
	})

	if err != nil {
		app.writeJSON(w, http.StatusInternalServerError, err, nil)
	}

	db.Close()
}

func (app *application) showItemHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		http.Error(w, "ID is not valid!", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid 'id' query parameter", http.StatusBadRequest)
		return
	}

	queries, ctx, db := app.connectDB()
	items, queryErr := queries.GetItem(ctx, uint32(id))

	if queryErr != nil {
		msg := fmt.Sprintf("Could not find item with id %d", id)

		http.Error(w, msg, http.StatusNotFound)
		app.logger.Error(queryErr.Error())

		return
	}

	jsonErr := app.writeJSON(w, http.StatusOK, items, nil)

	if jsonErr != nil {
		app.logger.Error(jsonErr.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

	db.Close()
}
