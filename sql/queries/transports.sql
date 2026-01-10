-- name: CreateTransport :exec
INSERT INTO transports (
    created_at, updated_at, truck_reg, trailer_reg, net, grain_type, sale_id, purchase_id
    )
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
);

-- name: GetAllTransports :many
SELECT t.*, p.suplier, s.client
FROM transports t
LEFT JOIN purchases p
ON t.purchase_id = p.id
LEFT JOIN sales s
ON t.sale_id = s.id;