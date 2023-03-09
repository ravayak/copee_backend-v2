-- name: GetSalesCommissionDiscount :one
SELECT * FROM sales_commissions_discount
WHERE sc_id = ? LIMIT 1;

-- name: ListSalesCommissionsDiscounts :many
SELECT * FROM sales_commissions_discount
ORDER BY sc_id ASC;

-- name: CreateSalesCommissionDiscount :execresult
INSERT INTO sales_commissions_discount ( sc_discount_pct , sc_com_pct  
) VALUES ( 
  ? , ? 
);

-- name: DeleteSalesCommissionDiscount :exec
DELETE FROM sales_commissions_discount
WHERE sc_id = ?;

-- name: UpdateSalesCommissionDiscount :exec
UPDATE sales_commissions_discount SET sc_discount_pct = ? , sc_com_pct = ?  
WHERE sc_id = ?;