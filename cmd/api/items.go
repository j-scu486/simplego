package main

import (
	"encoding/json"
	"fmt"
	"goweb/testapp"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type ItemFormData struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	OnSale   bool    `json:"onsale"`
	Stores   []int   `json:"stores"`
}

func (app *application) createItemHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	var itemFormData ItemFormData
	json.NewDecoder(r.Body).Decode(&itemFormData)

	item := &Item{
		Name:     itemFormData.Name,
		Price:    itemFormData.Price,
		Quantity: itemFormData.Quantity,
		OnSale:   itemFormData.OnSale,
	}

	err = validate.Struct(item)

	if err != nil {
		log.Panic(err)
		return
	}

	queries, ctx, db := app.connectDB()

	err = queries.CreateItem(ctx, testapp.CreateItemParams{
		Name:     item.Name,
		Price:    item.Price,
		Quantity: uint32(item.Quantity),
		Onsale:   toTinyInt(item.OnSale),
	})

	if err != nil {
		log.Panic(err)
	}

	if len(itemFormData.Stores) > 0 {
		id, _ := queries.LastInsertedId(ctx)

		for _, v := range itemFormData.Stores {
			err := queries.StoreItemCreate(ctx, testapp.StoreItemCreateParams{
				StoreID: uint32(v),
				ItemID:  uint32(id),
			})

			if err != nil {
				log.Panic(err)
			}
		}
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
