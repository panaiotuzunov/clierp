-- name: CreateReceipt :exec
INSERT INTO receipts (
    created_at, updated_at, truck_reg, trailer_reg, gross, tare, doc_type, grain_type, purchase_id, sale_id 
    )
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
);

-- name: GetAllReceipts :many
SELECT r.*, 
p.suplier, 
s.client, (r.gross - r.tare)::NUMERIC(12,3) AS net
FROM receipts r
LEFT JOIN purchases p
ON r.purchase_id = p.id
LEFT JOIN sales s
ON r.sale_id = s.id;

-- name: GetCurrentInventoryByType :many
SELECT grain_type, SUM(gross - tare)::NUMERIC(12,3) AS net
FROM receipts
GROUP BY(grain_type);