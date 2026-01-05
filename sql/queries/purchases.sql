-- name: CreatePurchase :exec
INSERT INTO purchases (
    created_at, updated_at, suplier, price, quantity, grain_type 
    )
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4
);

-- name: GetAllPurchases :many
SELECT * FROM purchases;