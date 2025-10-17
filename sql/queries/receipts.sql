-- name: CreateReceipt :exec
INSERT INTO receipts (
    created_at, updated_at, truck_reg, trailer_reg, gross, tare, net, doc_type, grain_type 
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
    $7
);

-- name: GetAllReceipts :many
SELECT * FROM receipts;

-- name: GetCurrentInventory :one
SELECT SUM(net)
FROM receipts;