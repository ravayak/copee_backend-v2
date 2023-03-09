package energiescostsservice

import (
	"database/sql"

	energies "github.com/ravayak/copee_backend/apis/domain/energies/costs"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	EnergieCostService energiesCostServiceInterface = &energiesCostService{}
)

type energiesCostService struct{}

type energiesCostServiceInterface interface {
	InsertEnergieCost(energieCostParams mysql.CreateEnergiesCostParams) (sql.Result, error)
	UpdateEnergieCost(energieCostParams mysql.UpdateEnergiesCostParams) error
	DeleteEnergieCost(energieCostId int32) error
	GetEnergieCost(energieCostId int32) (mysql.EnergiesCost, error)
	ListEnergieCosts() ([]mysql.EnergiesCost, error)
}

func (cs *energiesCostService) InsertEnergieCost(energieCostParams mysql.CreateEnergiesCostParams) (sql.Result, error) {
	return energies.InsertEnergiesCost(energieCostParams)
}

func (cs *energiesCostService) UpdateEnergieCost(energieCostParams mysql.UpdateEnergiesCostParams) error {
	return energies.UpdateEnergiesCost(energieCostParams)
}

func (cs *energiesCostService) DeleteEnergieCost(energieCostId int32) error {
	return energies.DeleteEnergiesCost(energieCostId)
}

func (cs *energiesCostService) GetEnergieCost(energieCostId int32) (mysql.EnergiesCost, error) {
	return energies.GetEnergiesCost(energieCostId)
}

func (cs *energiesCostService) ListEnergieCosts() ([]mysql.EnergiesCost, error) {
	return energies.ListEnergiesCosts()
}
