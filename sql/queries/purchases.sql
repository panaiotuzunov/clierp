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
SELECT p.*, SUM(r.net) AS expedited 
FROM purchases p
LEFT JOIN receipts r
ON p.id = r.purchase_id
GROUP BY p.id;

-- name: GetPurchaseById :one
SELECT * FROM purchases
WHERE id = $1;