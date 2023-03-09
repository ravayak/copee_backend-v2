package products

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListClientInstallationProducts() ([]mysql.ClientsInstallationsProduct, error) {
	ctx := context.Background()
	product, err := mysql.QueriesDb.ListClientInstallationProducts(ctx)
	return product, err
}

func GetClientInstallationProduct(productId int32) (mysql.ClientsInstallationsProduct, error) {
	ctx := context.Background()
	designation, err := mysql.QueriesDb.GetClientInstallationProduct(ctx, productId)
	return designation, err
}

func InsertClientInstallationProduct(productParams mysql.CreateClientInstallationProductParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateClientInstallationProduct(ctx, productParams)
	return res, err
}

func UpdateClientInstallationProduct(productParams mysql.UpdateClientInstallationProductParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateClientInstallationProduct(ctx, productParams)
	return err
}

func DeleteClientInstallationProduct(productId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteClientInstallationProduct(ctx, productId)
	return err
}
