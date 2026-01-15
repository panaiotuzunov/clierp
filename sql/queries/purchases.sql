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
SELECT p.*, 
    COALESCE(SUM(r.gross - r.tare), 0)::NUMERIC(12, 3) AS expedited_receipts,
    COALESCE(SUM(t.net), 0)::NUMERIC(12, 3) AS expedited_transports
FROM purchases p
LEFT JOIN receipts r
ON p.id = r.purchase_id
LEFT JOIN transports t
ON p.id = t.purchase_id
GROUP BY p.id
ORDER BY p.id;

-- name: GetPurchaseById :one
SELECT * FROM purchases
WHERE id = $1;