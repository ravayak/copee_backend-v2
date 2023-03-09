package catalogues

import (
	"database/sql"

	"github.com/ravayak/copee_backend/apis/domain/equipments/categories"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	CategoriesService categoriesServiceInterface = &categoriesService{}
)

type categoriesService struct{}
type categoriesServiceInterface interface {
	InsertCategory(catParams mysql.CreateCategoryParams) (sql.Result, error)
	UpdateCategory(catParams mysql.UpdateCategoryParams) error
	DeleteCategory(catId int32) error
	GetCategory(catId int32) (mysql.EquipmentsCategory, error)
	ListCategories() ([]mysql.EquipmentsCategory, error)
}

func (is *categoriesService) InsertCategory(catParams mysql.CreateCategoryParams) (sql.Result, error) {
	return categories.InsertEquipmentCategory(catParams)
}

func (is *categoriesService) UpdateCategory(catParams mysql.UpdateCategoryParams) error {
	return categories.UpdateEquipmentCategory(catParams)
}

func (is *categoriesService) DeleteCategory(catId int32) error {
	return categories.DeleteEquipmentCategory(catId)
}

func (is *categoriesService) GetCategory(catId int32) (mysql.EquipmentsCategory, error) {
	return categories.GetEquipmentCategory(catId)
}

func (is *categoriesService) ListCategories() ([]mysql.EquipmentsCategory, error) {
	return categories.ListEquipmentCategories()
}
