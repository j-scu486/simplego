// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: items.sql

package testapp

import (
	"context"
	"database/sql"
	"time"
)

const createItem = `-- name: CreateItem :exec
INSERT INTO goweb.items (
  name, price, quantity, onSale, created_at, updated_at, deleted_at
) VALUES (
  ?, ?, ?, ?, NOW(), NOW(), NULL
)
`

type CreateItemParams struct {
	Name     string
	Price    float64
	Quantity uint32
	Onsale   int8
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) error {
	_, err := q.db.ExecContext(ctx, createItem,
		arg.Name,
		arg.Price,
		arg.Quantity,
		arg.Onsale,
	)
	return err
}

const getItem = `-- name: GetItem :one
SELECT i.id, i.created_at, i.updated_at, i.deleted_at, i.name, i.price, i.quantity, i.onsale, s.id, s.created_at, s.updated_at, s.deleted_at, s.name, s.owner
FROM goweb.items i
JOIN goweb.stores_items si ON i.id = si.item_id
JOIN goweb.stores s ON si.store_id = s.id
WHERE i.name LIKE ?
LIMIT 1
`

type GetItemRow struct {
	ID          uint32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
	Name        string
	Price       float64
	Quantity    uint32
	Onsale      int8
	ID_2        uint32
	CreatedAt_2 time.Time
	UpdatedAt_2 time.Time
	DeletedAt_2 sql.NullTime
	Name_2      string
	Owner       StoresOwner
}

func (q *Queries) GetItem(ctx context.Context, name string) (GetItemRow, error) {
	row := q.db.QueryRowContext(ctx, getItem, name)
	var i GetItemRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
		&i.Price,
		&i.Quantity,
		&i.Onsale,
		&i.ID_2,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
		&i.DeletedAt_2,
		&i.Name_2,
		&i.Owner,
	)
	return i, err
}

const lastInsertedId = `-- name: LastInsertedId :one
SELECT LAST_INSERT_ID() AS id
`

func (q *Queries) LastInsertedId(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, lastInsertedId)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const storeItemCreate = `-- name: StoreItemCreate :exec
INSERT INTO goweb.stores_items (
  store_id, item_id
) VALUES (
  ?, ?
)
`

type StoreItemCreateParams struct {
	StoreID uint32
	ItemID  uint32
}

func (q *Queries) StoreItemCreate(ctx context.Context, arg StoreItemCreateParams) error {
	_, err := q.db.ExecContext(ctx, storeItemCreate, arg.StoreID, arg.ItemID)
	return err
}
