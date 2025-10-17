-- name: CreateReceipt :exec
INSERT INTO receipts (
    created_at, updated_at, truck_reg, trailer_reg, gross, tare, net 
    )
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5
);

-- name: GetAllReceipts :many
SELECT * FROM receipts;