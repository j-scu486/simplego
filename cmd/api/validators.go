package main

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type Item struct {
	Name     string  `validate:"required,max=40"`
	Price    float64 `validate:"required,numeric,gt=0"`
	Quantity int     `validate:"required,number"`
	OnSale   bool    `validate:"required,boolean"`
}

type Store struct {
	Name  string `validate:"required,max=40"`
	Owner string `validate:"required,oneof='state' 'private'"`
}

func validateStore(s *Store) error {
	var validate = validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(s)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
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
