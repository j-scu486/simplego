-- name: GetItem :one
SELECT i.*, s.name AS store_name, s.owner
FROM goweb.items i
JOIN goweb.stores_items si ON i.id = si.item_id
JOIN goweb.stores s ON si.store_id = s.id
WHERE i.name LIKE ?
LIMIT 1;

-- name: GetAllItems :many
SELECT i.*, s.name AS store_name, s.owner
FROM goweb.items i
LEFT JOIN goweb.stores_items si ON i.id = si.item_id
LEFT JOIN goweb.stores s ON si.store_id = s.id;

-- name: CreateItem :exec
INSERT INTO goweb.items (
  name, price, quantity, onSale, created_at, updated_at, deleted_at
) VALUES (
  ?, ?, ?, ?, NOW(), NOW(), NULL
);

-- name: StoreItemCreate :exec
INSERT INTO goweb.stores_items (
  store_id, item_id
) VALUES (
  ?, ?
);

-- name: LastInsertedId :one
SELECT LAST_INSERT_ID() AS id;
