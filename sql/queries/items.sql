-- name: GetItem :one
SELECT i.*, s.*
FROM goweb.items i
JOIN goweb.stores_items si ON i.id = si.item_id
JOIN goweb.stores s ON si.store_id = s.id
WHERE i.id = ?
LIMIT 1;

-- name: CreateItem :exec
INSERT INTO goweb.items (
  name, price, quantity, onSale, created_at, updated_at, deleted_at
) VALUES (
  ?, ?, ?, ?, NOW(), NOW(), NULL
);
