package main

import (
	"encoding/json"
	"fmt"
	"goweb/testapp"
	"net/http"

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
	if err := r.ParseForm(); err != nil {
		app.logger.Error(err.Error())
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

	if err := validate.Struct(item); err != nil {
		app.logger.Error(err.Error())
		return
	}

	queries, ctx, db := app.connectDB()
	defer db.Close()

	err := queries.CreateItem(ctx, testapp.CreateItemParams{
		Name:     item.Name,
		Price:    item.Price,
		Quantity: uint32(item.Quantity),
		Onsale:   toTinyInt(item.OnSale),
	})

	if err != nil {
		app.logger.Error(err.Error())
	}

	if len(itemFormData.Stores) > 0 {
		id, _ := queries.LastInsertedId(ctx)

		for _, v := range itemFormData.Stores {
			err := queries.StoreItemCreate(ctx, testapp.StoreItemCreateParams{
				StoreID: uint32(v),
				ItemID:  uint32(id),
			})

			if err != nil {
				app.logger.Error(err.Error())
			}
		}
	}

	app.writeJSON(w, http.StatusCreated, item, nil)
}

func (app *application) showItemHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	queries, ctx, db := app.connectDB()
	defer db.Close()

	if name == "" {
		items, _ := queries.GetAllItems(ctx)

		if jsonErr := app.writeJSON(w, http.StatusOK, items, nil); jsonErr != nil {
			app.logger.Error(jsonErr.Error())
			http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		}
	} else {
		items, queryErr := queries.GetItem(ctx, `%`+name+`%`)

		if queryErr != nil {
			msg := fmt.Sprintf("Could not find item with name %s", name)

			http.Error(w, msg, http.StatusNotFound)
			app.logger.Error(queryErr.Error())

			return
		}

		if jsonErr := app.writeJSON(w, http.StatusOK, items, nil); jsonErr != nil {
			app.logger.Error(jsonErr.Error())
			http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		}
	}
}
