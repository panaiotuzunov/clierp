-- name: CreateSale :exec
INSERT INTO sales (
    created_at, updated_at, client, price, quantity, grain_type 
    )
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4
);

-- name: GetAllSales :many
SELECT s.*, 
    COALESCE(SUM(r.gross - r.tare), 0)::NUMERIC(12, 3) AS expedited_receipts,
    COALESCE(SUM(t.net), 0)::NUMERIC(12, 3) AS expedited_transports
FROM sales s
LEFT JOIN receipts r
ON s.id = r.sale_id
LEFT JOIN transports t
ON s.id = t.sale_id
GROUP BY s.id
ORDER BY s.id;

-- name: GetSaleById :one
SELECT * FROM sales
WHERE id = $1;

-- name: GetSalesByGrainType :many
SELECT s.*, 
    COALESCE(SUM(r.gross - r.tare), 0)::NUMERIC(12, 3) AS expedited_receipts,
    COALESCE(SUM(t.net), 0)::NUMERIC(12, 3) AS expedited_transports
FROM sales s
LEFT JOIN receipts r
ON s.id = r.sale_id
LEFT JOIN transports t
ON s.id = t.sale_id
WHERE s.grain_type = $1
GROUP BY s.id
ORDER BY s.id;

-- name: GetSaleByIdandGrainType :one
SELECT * FROM sales
WHERE id = $1 AND
grain_type = $2;
