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
WITH purchases_summary AS (
    SELECT p.*, 
    COALESCE(SUM(r.gross - r.tare)::NUMERIC(12, 3), 0)::NUMERIC(12, 3) AS expedited
    FROM purchases p
    LEFT JOIN receipts r
    ON p.id = r.purchase_id
    GROUP BY p.id
)
SELECT *, 
       (quantity - expedited)::NUMERIC(12, 3) AS leftover
FROM purchases_summary
ORDER BY id;

-- name: GetPurchaseById :one
SELECT * FROM purchases
WHERE id = $1;