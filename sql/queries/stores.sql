-- name: GetStore :one
SELECT * FROM goweb.stores
WHERE id = ?
LIMIT 1;

-- name: GetStoreItems :many
SELECT i.*, s.id AS store_id
FROM goweb.items i
JOIN goweb.stores_items si ON i.id = si.item_id
JOIN goweb.stores s ON si.store_id = s.id
WHERE s.id = ?;

-- name: CreateStore :exec
INSERT INTO goweb.stores (
  name, owner, created_at, updated_at, deleted_at
)
VALUES (
  ?, ?, NOW(), NOW(), NULL
);