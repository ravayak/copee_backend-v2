package energies

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListEnergiesCosts() ([]mysql.EnergiesCost, error) {
	ctx := context.Background()
	energiesCosts, err := mysql.QueriesDb.ListEnergiesCosts(ctx)
	return energiesCosts, err
}

func GetEnergiesCost(energyCostId int32) (mysql.EnergiesCost, error) {
	ctx := context.Background()
	energiesCost, err := mysql.QueriesDb.GetEnergiesCost(ctx, energyCostId)
	return energiesCost, err
}

func InsertEnergiesCost(energiesCostParam mysql.CreateEnergiesCostParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateEnergiesCost(ctx, energiesCostParam)
	return res, err
}

func UpdateEnergiesCost(energiesCostParam mysql.UpdateEnergiesCostParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateEnergiesCost(ctx, energiesCostParam)
	return err
}

func DeleteEnergiesCost(energyCostId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteEnergiesCost(ctx, energyCostId)
	return err
}
