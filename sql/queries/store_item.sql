-- name: StoreItemCreate :exec
INSERT INTO goweb.stores_items (
  store_id, item_id
) VALUES (
  ?, ?
);