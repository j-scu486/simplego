package main

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
