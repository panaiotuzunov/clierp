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
WITH sales_summary AS (
    SELECT s.*, 
    COALESCE(ABS(SUM(r.gross - r.tare))::NUMERIC(12, 3), 0)::NUMERIC(12, 3) AS expedited
    FROM sales s
    LEFT JOIN receipts r
    ON s.id = r.sale_id
    GROUP BY s.id
)
SELECT *, 
       (quantity - expedited)::NUMERIC(12, 3) AS leftover
FROM sales_summary
ORDER BY id;

-- name: GetSaleById :one
SELECT * FROM sales
WHERE id = $1;

-- name: GetSalesByGrainType :many
WITH sales_summary AS (
    SELECT s.*, 
    COALESCE(ABS(SUM(r.gross - r.tare))::NUMERIC(12, 3), 0)::NUMERIC(12, 3) AS expedited
    FROM sales s
    LEFT JOIN receipts r
    ON s.id = r.sale_id
    WHERE s.grain_type = $1
    GROUP BY s.id
)
SELECT *, 
       (quantity - expedited)::NUMERIC(12, 3) AS leftover
FROM sales_summary
ORDER BY id;

-- name: GetSaleByIdandGrainType :one
SELECT * FROM sales
WHERE id = $1 AND
grain_type = $2;
