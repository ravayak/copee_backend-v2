package commissions

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListSalesCommissions() ([]mysql.SalesCommission, error) {
	ctx := context.Background()
	com, err := mysql.QueriesDb.ListSalesCommissions(ctx)
	return com, err
}

func GetSalesCommission(logId int32) (mysql.SalesCommission, error) {
	ctx := context.Background()
	com, err := mysql.QueriesDb.GetSalesCommission(ctx, logId)
	return com, err
}

func InsertSalesCommission(comParams mysql.CreateSalesCommissionParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateSalesCommission(ctx, comParams)
	return res, err
}

func UpdateSalesCommission(comParams mysql.UpdateSalesCommissionParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateSalesCommission(ctx, comParams)
	return err
}

func DeleteSalesCommission(groupId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteSalesCommission(ctx, groupId)
	return err
}
