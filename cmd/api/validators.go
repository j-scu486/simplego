package main

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type Item struct {
	Name     string  `validate:"required"`
	Price    float64 `validate:"required"`
	Quantity int     `validate:"required"`
	OnSale   bool    `validate:"required"`
}

func validateItem(i *Item) error {
	var validate = validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(i)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
