-- name: GetSalesCommission :one
SELECT * FROM sales_commissions
WHERE sc_id = ? LIMIT 1;

-- name: ListSalesCommissions :many
SELECT * FROM sales_commissions
ORDER BY sc_id ASC;

-- name: CreateSalesCommission :execresult
INSERT INTO sales_commissions (sc_commissions_discount_id , sc_group_id
) VALUES ( 
  ? , ? 
);

-- name: UpdateSalesCommission :exec
UPDATE sales_commissions SET sc_commissions_discount_id = ? , sc_group_id = ?
WHERE sc_id = ?;

-- name: DeleteSalesCommission :exec
DELETE FROM sales_commissions
WHERE sc_id = ?;