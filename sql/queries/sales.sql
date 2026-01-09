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
SELECT s.*, COALESCE(SUM(r.net), 0)::INT AS expedited 
FROM sales s
LEFT JOIN receipts r
ON s.id = r.sale_id
GROUP BY s.id
ORDER BY s.id;

-- name: GetSaleById :one
SELECT * FROM sales
WHERE id = $1;