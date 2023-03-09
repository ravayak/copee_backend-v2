package discounts

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListSalesCommissionsDiscounts() ([]mysql.SalesCommissionsDiscount, error) {
	ctx := context.Background()
	comDiscount, err := mysql.QueriesDb.ListSalesCommissionsDiscounts(ctx)
	return comDiscount, err
}

func GetSalesCommissionsDiscount(cdId int32) (mysql.SalesCommissionsDiscount, error) {
	ctx := context.Background()
	cd, err := mysql.QueriesDb.GetSalesCommissionDiscount(ctx, cdId)
	return cd, err
}

func InsertSalesCommissionsDiscount(cdParams mysql.CreateSalesCommissionDiscountParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateSalesCommissionDiscount(ctx, cdParams)
	return res, err
}

func UpdateSalesCommissionsDiscount(cdParams mysql.UpdateSalesCommissionDiscountParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateSalesCommissionDiscount(ctx, cdParams)
	return err
}

func DeleteSalesCommissionsDiscount(groupId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteSalesCommissionDiscount(ctx, groupId)
	return err
}
