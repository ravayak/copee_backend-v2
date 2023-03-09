package datatest

import (
	"database/sql"
	"time"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

// Create user data test
var CreateUserParams = mysql.CreateUserParams{
	UserFirstname:       sql.NullString{String: "Test", Valid: true},
	UserLastname:        sql.NullString{String: "Test", Valid: true},
	UserEmail:           sql.NullString{String: "test@test.test", Valid: true},
	UserPhone:           sql.NullString{String: "", Valid: false},
	UserIsActive:        sql.NullBool{Bool: true, Valid: true},
	UserDateCreated:     time.Now(),
	UserRecruitmentDate: sql.NullTime{Valid: false},
}

var CreateSharedUserParams = mysql.CreateUserParams{
	UserFirstname:       sql.NullString{String: "Test", Valid: true},
	UserLastname:        sql.NullString{String: "Test", Valid: true},
	UserEmail:           sql.NullString{String: "test-shared-user@test.test", Valid: true},
	UserPhone:           sql.NullString{String: "", Valid: false},
	UserIsActive:        sql.NullBool{Bool: true, Valid: true},
	UserDateCreated:     time.Now(),
	UserRecruitmentDate: sql.NullTime{Valid: false},
}

var CreateEnergiesParams = mysql.CreateEnergiesCostParams{
	EncType: sql.NullString{String: "Test", Valid: true},
	EncDate: sql.NullTime{Time: time.Now(), Valid: true},
	EncCost: sql.NullString{String: "10.21", Valid: true},
}

var CreateGroupParams = mysql.CreateGroupParams{
	GName:        sql.NullString{String: "Test", Valid: true},
	GDescription: sql.NullString{String: "Test", Valid: true},
}

var CreateEquipmentProductParams = mysql.CreateEquipmentProductParams{
	EpProductID:         sql.NullInt32{Int32: 1, Valid: true},
	EpPrice:             sql.NullString{String: "10.21", Valid: true},
	EpVat:               sql.NullString{String: "10.21", Valid: true},
	EpInstallationPrice: sql.NullString{String: "10.21", Valid: true},
}

func GetUpdateEquipmentCatalogueParams(catalogueId, productId int32) mysql.UpdateEquipmentCatalogueParams {
	return mysql.UpdateEquipmentCatalogueParams{
		EcID:          catalogueId,
		EcProductID:   sql.NullInt32{Int32: productId, Valid: true},
		EcQuantity:    sql.NullInt32{Int32: 2, Valid: true},
		EcDescription: sql.NullString{String: "Test updated", Valid: true},
		EcDateAdded:   sql.NullTime{Time: time.Now(), Valid: true},
		EcDateUpdated: sql.NullTime{Time: time.Now(), Valid: true},
	}
}

func GetEquipmentCatalogueParams(productId int32) mysql.CreateEquipmentCatalogueParams {
	return mysql.CreateEquipmentCatalogueParams{
		EcProductID:   sql.NullInt32{Int32: productId, Valid: true},
		EcQuantity:    sql.NullInt32{Int32: 1, Valid: true},
		EcDescription: sql.NullString{String: "Test", Valid: true},
		EcDateAdded:   sql.NullTime{Time: time.Now(), Valid: true},
		EcDateUpdated: sql.NullTime{Time: time.Now(), Valid: true},
	}
}

func GetUpdateEquipmentProductParams(epId int32) mysql.UpdateEquipmentProductParams {
	return mysql.UpdateEquipmentProductParams{
		EpID:                epId,
		EpProductID:         sql.NullInt32{Int32: 1, Valid: true},
		EpPrice:             sql.NullString{String: "11.00", Valid: true},
		EpVat:               sql.NullString{String: "11.00", Valid: true},
		EpInstallationPrice: sql.NullString{String: "11.00", Valid: true},
	}
}

var CreateClientParams = mysql.CreateClientParams{
	ClientLastName:         sql.NullString{String: "Test", Valid: true},
	ClientFirstName:        sql.NullString{String: "Test", Valid: true},
	ClientEmail:            sql.NullString{String: "test@test.test", Valid: true},
	ClientFiscalYearIncome: sql.NullInt32{Int32: 100000, Valid: true},
}

func GetClientInstallationCreateParams(clientId, userId, sharedUserId int32) mysql.CreateClientInstallationParams {
	return mysql.CreateClientInstallationParams{
		CiClientID:     sql.NullInt32{Int32: clientId, Valid: true},
		CiUserID:       sql.NullInt32{Int32: userId, Valid: true},
		CiSharedUserID: sql.NullInt32{Int32: sharedUserId, Valid: true},
		CiCreationDate: sql.NullTime{Time: time.Now(), Valid: true},
		CiUpdateDate:   sql.NullTime{Time: time.Now(), Valid: true},
		CiStatus:       sql.NullString{String: "Test", Valid: true},
		CiComments:     sql.NullString{String: "Test", Valid: true},
	}
}

func GetClientLoanCreateParams(clientId int32) mysql.CreateClientLoanParams {
	return mysql.CreateClientLoanParams{
		ClClientID:         sql.NullInt32{Int32: clientId, Valid: true},
		ClAmount:           sql.NullFloat64{Float64: 100000, Valid: true},
		ClInstallments:     sql.NullFloat64{Float64: 100., Valid: true},
		ClClientProvision:  sql.NullFloat64{Float64: 1000, Valid: true},
		ClClientPrepayment: sql.NullFloat64{Float64: 0, Valid: true},
		ClInsured:          sql.NullInt32{Int32: 0, Valid: true},
		ClFundingAgency:    sql.NullString{String: "Test", Valid: true},
	}
}

func GetClientHomeParamsCreateParams(clientId, geoId, homeEqmtId, homeBillsId int32) mysql.CreateClientHomeParams {
	return mysql.CreateClientHomeParams{
		ChClientID:         sql.NullInt32{Int32: clientId, Valid: true},
		ChGeoID:            sql.NullInt32{Int32: geoId, Valid: true},
		ChConstructionYear: sql.NullInt32{Int32: 2010, Valid: true},
		ChArea:             sql.NullInt32{Int32: 100, Valid: true},
		ChResidents:        sql.NullInt32{Int32: 2, Valid: true},
		ChRoofPositionning: sql.NullString{String: "Sud", Valid: true},
		ChRoofSlope:        sql.NullInt32{Int32: 45, Valid: true},
		ChLabel:            sql.NullString{String: "A", Valid: true},
		ChTr:               sql.NullInt32{Int32: 100, Valid: true},
		ChHuc:              sql.NullFloat64{Float64: 100, Valid: true},
		ChIsolation:        sql.NullString{String: "A", Valid: true},
	}
}

func GetCompanyCreateParams(geoId int32) mysql.CreateCompanyParams {
	return mysql.CreateCompanyParams{
		CGeoID:        sql.NullInt32{Int32: geoId, Valid: true},
		CName:         sql.NullString{String: "Company 1", Valid: true},
		CRcs:          sql.NullString{String: "Type 1", Valid: true},
		CSiret:        sql.NullString{String: "Siret 1", Valid: true},
		CSiren:        sql.NullString{String: "Siren 1", Valid: true},
		CIntraEuVat:   sql.NullString{String: "IntraEuVat 1", Valid: true},
		CLegalStatus:  sql.NullString{String: "LegalStatus 1", Valid: true},
		CCreationDate: sql.NullString{String: "CreationDate 1", Valid: true},
		CCapital:      sql.NullInt32{Int32: 1000, Valid: true},
	}
}

var CreateCompanyParams = mysql.CreateCompanyParams{
	CName:         sql.NullString{String: "Company 1", Valid: true},
	CRcs:          sql.NullString{String: "Type 1", Valid: true},
	CSiret:        sql.NullString{String: "Siret 1", Valid: true},
	CSiren:        sql.NullString{String: "Siren 1", Valid: true},
	CIntraEuVat:   sql.NullString{String: "IntraEuVat 1", Valid: true},
	CLegalStatus:  sql.NullString{String: "LegalStatus 1", Valid: true},
	CCreationDate: sql.NullString{String: "CreationDate 1", Valid: true},
	CCapital:      sql.NullInt32{Int32: 1000, Valid: true},
}

var GeoCreateParams = mysql.CreateGeoParams{
	ChgZone:       sql.NullString{String: "h2", Valid: true},
	ChgLatitude:   sql.NullFloat64{Float64: 45.5, Valid: true},
	ChgLongitude:  sql.NullFloat64{Float64: 45.5, Valid: true},
	ChgAltitude:   sql.NullInt32{Int32: 45, Valid: true},
	ChgDepartment: sql.NullInt32{Int32: 66, Valid: true},
	ChgCity:       sql.NullString{String: "Paris", Valid: true},
	ChgAddress:    sql.NullString{String: "1 rue de la paix", Valid: true},
	ChgPostcode:   sql.NullInt32{Int32: 75001, Valid: true},
}

// Update data test
func GetGroupUpdateParams(groupId int32) mysql.UpdateGroupParams {
	return mysql.UpdateGroupParams{
		GName:        sql.NullString{String: "Test updated", Valid: true},
		GDescription: sql.NullString{String: "Test updated", Valid: true},
		GID:          groupId,
	}
}

// Client Home Bills Create
func GetClientHomeParamsBillsParams(homeId int32) mysql.CreateClientHomeBillParams {
	return mysql.CreateClientHomeBillParams{
		ChbHomeID:      sql.NullInt32{Int32: homeId, Valid: true},
		ChbElectricity: sql.NullInt32{Int32: 50, Valid: true},
		ChbNaturalGas:  sql.NullInt32{Int32: 50, Valid: true},
		ChbPropaneGas:  sql.NullInt32{Int32: 50, Valid: true},
		ChbWood:        sql.NullInt32{Int32: 50, Valid: true},
		ChbHeatingOil:  sql.NullInt32{Int32: 50, Valid: true},
		ChbYear:        sql.NullInt32{Int32: 2009, Valid: true},
	}
}

func GetUpdateClientHomeBillsParams(homeId, homeBillsId int32) mysql.UpdateClientHomeBillParams {
	return mysql.UpdateClientHomeBillParams{
		ChbID:          homeBillsId,
		ChbHomeID:      sql.NullInt32{Int32: homeId, Valid: true},
		ChbElectricity: sql.NullInt32{Int32: 60, Valid: true},
		ChbNaturalGas:  sql.NullInt32{Int32: 60, Valid: true},
		ChbPropaneGas:  sql.NullInt32{Int32: 60, Valid: true},
		ChbWood:        sql.NullInt32{Int32: 60, Valid: true},
		ChbHeatingOil:  sql.NullInt32{Int32: 60, Valid: true},
		ChbYear:        sql.NullInt32{Int32: 2010, Valid: true},
	}
}

func GetClientHomeEquipmentParams(homeId int32) mysql.CreateClientHomeEquipmentParams {
	return mysql.CreateClientHomeEquipmentParams{
		CheHomeID:        sql.NullInt32{Int32: homeId, Valid: true},
		CheEquipmentType: sql.NullString{String: "Test", Valid: true},
		CheDescription:   sql.NullString{String: "Test", Valid: true},
	}
}

func GetClientHomeEquipmentUpdateParams(eqmtId, homeId int32) mysql.UpdateClientHomeEquipmentParams {
	return mysql.UpdateClientHomeEquipmentParams{
		CheHomeID:        sql.NullInt32{Int32: homeId, Valid: true},
		CheID:            eqmtId,
		CheEquipmentType: sql.NullString{String: "Test updated", Valid: true},
		CheDescription:   sql.NullString{String: "Test updated", Valid: true},
	}
}

func GetClientHomeParams(clientId, geoId int32) mysql.CreateClientHomeParams {
	return mysql.CreateClientHomeParams{
		ChClientID:         sql.NullInt32{Int32: int32(clientId), Valid: true},
		ChGeoID:            sql.NullInt32{Int32: int32(geoId), Valid: true},
		ChConstructionYear: sql.NullInt32{Int32: 2000, Valid: true},
		ChArea:             sql.NullInt32{Int32: 100, Valid: true},
		ChResidents:        sql.NullInt32{Int32: 4, Valid: true},
		ChRoofPositionning: sql.NullString{String: "Nord", Valid: true},
		ChRoofSlope:        sql.NullInt32{Int32: 45, Valid: true},
		ChLabel:            sql.NullString{String: "A", Valid: true},
		ChTr:               sql.NullInt32{Int32: 100, Valid: true},
		ChHuc:              sql.NullFloat64{Float64: 100.0, Valid: true},
		ChIsolation:        sql.NullString{String: "Plaquo", Valid: true},
	}
}

func GetClientInstallationFileParams(clientInstallationId int32) mysql.CreateClientInstallationFileParams {
	return mysql.CreateClientInstallationFileParams{
		CifInstallationID:   sql.NullInt32{Int32: clientInstallationId, Valid: true},
		CifFileType:         sql.NullString{String: "Test", Valid: true},
		CifFileUrl:          sql.NullString{String: "Test", Valid: true},
		CifFileCreationDate: sql.NullTime{Time: time.Now(), Valid: true},
	}
}

func GetClientInstallationPaymentCreateParams(clientInstallationId int32) mysql.CreateClientInstallationPaymentParams {
	return mysql.CreateClientInstallationPaymentParams{
		CipAmount:            sql.NullFloat64{Float64: 100.0, Valid: true},
		CipType:              sql.NullString{String: "Test", Valid: true},
		CipDate:              sql.NullTime{Time: time.Now(), Valid: true},
		CipInstallationID:    sql.NullInt32{Int32: int32(clientInstallationId), Valid: true},
		CipDescription:       sql.NullString{String: "Test", Valid: true},
		CipTransactionNumber: sql.NullString{String: "Test", Valid: true},
		CipVatRate:           sql.NullFloat64{Float64: 100.0, Valid: true},
	}
}

func GetClientInstallationPaymentUpdateParams(clientInstallationPaymentId, clientInstallationId int32) mysql.UpdateClientInstallationPaymentParams {
	return mysql.UpdateClientInstallationPaymentParams{
		CipAmount:            sql.NullFloat64{Float64: 100.0, Valid: true},
		CipType:              sql.NullString{String: "Test updated", Valid: true},
		CipDate:              sql.NullTime{Time: time.Now(), Valid: true},
		CipInstallationID:    sql.NullInt32{Int32: int32(clientInstallationId), Valid: true},
		CipDescription:       sql.NullString{String: "Test updated", Valid: true},
		CipTransactionNumber: sql.NullString{String: "Test updated", Valid: true},
		CipVatRate:           sql.NullFloat64{Float64: 100.0, Valid: true},
		CipID:                clientInstallationPaymentId,
	}
}

func GetUpdateClientInstallationFileParams(clientInstallationFileId, clientInstallationId int32) mysql.UpdateClientInstallationFileParams {
	return mysql.UpdateClientInstallationFileParams{
		CifInstallationID:   sql.NullInt32{Int32: clientInstallationId, Valid: true},
		CifFileType:         sql.NullString{String: "Test updated", Valid: true},
		CifFileUrl:          sql.NullString{String: "Test updated", Valid: true},
		CifFileCreationDate: sql.NullTime{Time: time.Now(), Valid: true},
		CifID:               clientInstallationFileId,
	}
}
