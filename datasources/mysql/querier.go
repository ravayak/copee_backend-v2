// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package mysql

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (sql.Result, error)
	CreateClient(ctx context.Context, arg CreateClientParams) (sql.Result, error)
	CreateClientAss(ctx context.Context, arg CreateClientAssParams) (sql.Result, error)
	CreateClientHome(ctx context.Context, arg CreateClientHomeParams) (sql.Result, error)
	CreateClientHomeBill(ctx context.Context, arg CreateClientHomeBillParams) (sql.Result, error)
	CreateClientHomeEquipment(ctx context.Context, arg CreateClientHomeEquipmentParams) (sql.Result, error)
	CreateClientInstallation(ctx context.Context, arg CreateClientInstallationParams) (sql.Result, error)
	CreateClientInstallationFile(ctx context.Context, arg CreateClientInstallationFileParams) (sql.Result, error)
	CreateClientInstallationPayment(ctx context.Context, arg CreateClientInstallationPaymentParams) (sql.Result, error)
	CreateClientInstallationProduct(ctx context.Context, arg CreateClientInstallationProductParams) (sql.Result, error)
	CreateClientLoan(ctx context.Context, arg CreateClientLoanParams) (sql.Result, error)
	CreateCompany(ctx context.Context, arg CreateCompanyParams) (sql.Result, error)
	CreateEnergiesCost(ctx context.Context, arg CreateEnergiesCostParams) (sql.Result, error)
	CreateEquipmentCatalogue(ctx context.Context, arg CreateEquipmentCatalogueParams) (sql.Result, error)
	CreateEquipmentDesignations(ctx context.Context, arg CreateEquipmentDesignationsParams) (sql.Result, error)
	CreateEquipmentInstaller(ctx context.Context, arg CreateEquipmentInstallerParams) (sql.Result, error)
	CreateEquipmentProduct(ctx context.Context, arg CreateEquipmentProductParams) (sql.Result, error)
	CreateEquipmentServices(ctx context.Context, arg CreateEquipmentServicesParams) (sql.Result, error)
	CreateGeo(ctx context.Context, arg CreateGeoParams) (sql.Result, error)
	CreateGroup(ctx context.Context, arg CreateGroupParams) (sql.Result, error)
	CreateGroupPrivilege(ctx context.Context, gpPrivilege sql.NullString) (sql.Result, error)
	CreateGroupPrivilegeLnk(ctx context.Context, arg CreateGroupPrivilegeLnkParams) (sql.Result, error)
	CreateGroupUserLnk(ctx context.Context, arg CreateGroupUserLnkParams) (sql.Result, error)
	CreateGroupUserLnkByGroupName(ctx context.Context, arg CreateGroupUserLnkByGroupNameParams) (sql.Result, error)
	CreateLog(ctx context.Context, arg CreateLogParams) (sql.Result, error)
	CreateSalesCommission(ctx context.Context, arg CreateSalesCommissionParams) (sql.Result, error)
	CreateSalesCommissionDiscount(ctx context.Context, arg CreateSalesCommissionDiscountParams) (sql.Result, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error)
	DeleteCategory(ctx context.Context, ecID int32) error
	DeleteClient(ctx context.Context, clientID int32) error
	DeleteClientAss(ctx context.Context, caID int32) error
	DeleteClientHome(ctx context.Context, chID int32) error
	DeleteClientHomeBill(ctx context.Context, chbID int32) error
	DeleteClientHomeEquipment(ctx context.Context, cheID int32) error
	DeleteClientInstallation(ctx context.Context, ciID int32) error
	DeleteClientInstallationFile(ctx context.Context, cifID int32) error
	DeleteClientInstallationPayment(ctx context.Context, cipID int32) error
	DeleteClientInstallationProduct(ctx context.Context, cipID int32) error
	DeleteClientLoan(ctx context.Context, clID int32) error
	DeleteCompany(ctx context.Context, cID int32) error
	DeleteEnergiesCost(ctx context.Context, encID int32) error
	DeleteEquipmentCatalogue(ctx context.Context, ecID int32) error
	DeleteEquipmentDesignation(ctx context.Context, eaID int32) error
	DeleteEquipmentInstaller(ctx context.Context, eiID int32) error
	DeleteEquipmentProduct(ctx context.Context, epID int32) error
	DeleteEquipmentService(ctx context.Context, esID int32) error
	DeleteGeo(ctx context.Context, chgID int32) error
	DeleteGroup(ctx context.Context, gID int32) error
	DeleteGroupPrivilege(ctx context.Context, gpPrivilegeID int32) error
	DeleteGroupPrivilegeLnk(ctx context.Context, gplID int32) error
	DeleteGroupUserLnk(ctx context.Context, guID int32) error
	DeleteLog(ctx context.Context, logID int32) error
	DeleteSalesCommission(ctx context.Context, scID int32) error
	DeleteSalesCommissionDiscount(ctx context.Context, scID int32) error
	DeleteUser(ctx context.Context, userID int32) error
	GetCategory(ctx context.Context, ecID int32) (EquipmentsCategory, error)
	GetClient(ctx context.Context, clientID int32) (Client, error)
	GetClientAss(ctx context.Context, caID int32) (ClientsAss, error)
	GetClientHomeEquipmentParams(ctx context.Context, cheID int32) (ClientsHomesEquipment, error)
	GetClientHomeParams(ctx context.Context, chID int32) (ClientsHome, error)
	GetClientHomeParamsBill(ctx context.Context, chbID int32) (ClientsHomesBill, error)
	GetClientInstallation(ctx context.Context, ciID int32) (ClientsInstallation, error)
	GetClientInstallationFile(ctx context.Context, cifID int32) (ClientsInstallationsFile, error)
	GetClientInstallationPayment(ctx context.Context, cipID int32) (ClientsInstallationsPayment, error)
	GetClientInstallationProduct(ctx context.Context, cipID int32) (ClientsInstallationsProduct, error)
	GetClientLoan(ctx context.Context, clID int32) (ClientsLoan, error)
	GetCompany(ctx context.Context, cID int32) (Company, error)
	GetEnergiesCost(ctx context.Context, encID int32) (EnergiesCost, error)
	GetEquipmentCatalogue(ctx context.Context, ecID int32) (EquipmentsCatalogue, error)
	GetEquipmentDesignation(ctx context.Context, eaID int32) (EquipmentsDesignation, error)
	GetEquipmentInstaller(ctx context.Context, eiID int32) (EquipmentsInstaller, error)
	GetEquipmentProduct(ctx context.Context, epID int32) (EquipmentsProduct, error)
	GetEquipmentService(ctx context.Context, esID int32) (EquipmentsService, error)
	GetGeo(ctx context.Context, chgID int32) (Geo, error)
	GetGroup(ctx context.Context, gID int32) (Group, error)
	GetGroupByName(ctx context.Context, gName sql.NullString) (Group, error)
	GetGroupPrivilege(ctx context.Context, gpPrivilegeID int32) (GroupsPrivilege, error)
	GetGroupPrivilegeLnk(ctx context.Context, gplID int32) (GroupsPrivilegesLnk, error)
	GetGroupPrivilegesByUserId(ctx context.Context, userID int32) ([]GroupsPrivilege, error)
	GetGroupUserLnk(ctx context.Context, guID int32) (GroupsUsersLnk, error)
	GetLog(ctx context.Context, logID int32) (Log, error)
	GetSalesCommission(ctx context.Context, scID int32) (SalesCommission, error)
	GetSalesCommissionDiscount(ctx context.Context, scID int32) (SalesCommissionsDiscount, error)
	GetUser(ctx context.Context, userID int32) (User, error)
	GetUserByEmail(ctx context.Context, userEmail sql.NullString) (User, error)
	ListCategories(ctx context.Context) ([]EquipmentsCategory, error)
	ListClientInstallationProducts(ctx context.Context) ([]ClientsInstallationsProduct, error)
	ListClients(ctx context.Context) ([]Client, error)
	ListClientsAss(ctx context.Context) ([]ClientsAss, error)
	ListClientsByUser(ctx context.Context, ciUserID sql.NullInt32) ([]Client, error)
	ListClientsHomes(ctx context.Context) ([]ClientsHome, error)
	ListClientsHomesBills(ctx context.Context) ([]ClientsHomesBill, error)
	ListClientsHomesEquipments(ctx context.Context) ([]ClientsHomesEquipment, error)
	ListClientsIdByUser(ctx context.Context, ciUserID sql.NullInt32) ([]int32, error)
	ListClientsInstallation(ctx context.Context) ([]ClientsInstallation, error)
	ListClientsInstallationFiles(ctx context.Context) ([]ClientsInstallationsFile, error)
	ListClientsInstallationsPayments(ctx context.Context) ([]ClientsInstallationsPayment, error)
	ListClientsLoans(ctx context.Context) ([]ClientsLoan, error)
	ListCompanies(ctx context.Context) ([]Company, error)
	ListEnergiesCosts(ctx context.Context) ([]EnergiesCost, error)
	ListEquipmentCatalogue(ctx context.Context) ([]EquipmentsCatalogue, error)
	ListEquipmentDesignations(ctx context.Context) ([]EquipmentsDesignation, error)
	ListEquipmentInstallers(ctx context.Context) ([]EquipmentsInstaller, error)
	ListEquipmentProducts(ctx context.Context) ([]EquipmentsProduct, error)
	ListEquipmentServices(ctx context.Context) ([]EquipmentsService, error)
	ListGeo(ctx context.Context) ([]Geo, error)
	ListGroupPrivilegesLnk(ctx context.Context) ([]GroupsPrivilegesLnk, error)
	ListGroups(ctx context.Context) ([]Group, error)
	ListGroupsByUser(ctx context.Context, guUserID sql.NullInt32) ([]Group, error)
	ListGroupsPrivileges(ctx context.Context) ([]GroupsPrivilege, error)
	ListGroupsUsersLnk(ctx context.Context) ([]GroupsUsersLnk, error)
	ListLogs(ctx context.Context) ([]Log, error)
	ListSalesCommissions(ctx context.Context) ([]SalesCommission, error)
	ListSalesCommissionsDiscounts(ctx context.Context) ([]SalesCommissionsDiscount, error)
	ListUsers(ctx context.Context) ([]User, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error
	UpdateClient(ctx context.Context, arg UpdateClientParams) error
	UpdateClientAss(ctx context.Context, arg UpdateClientAssParams) error
	UpdateClientHome(ctx context.Context, arg UpdateClientHomeParams) error
	UpdateClientHomeBill(ctx context.Context, arg UpdateClientHomeBillParams) error
	UpdateClientHomeEquipment(ctx context.Context, arg UpdateClientHomeEquipmentParams) error
	UpdateClientInstallation(ctx context.Context, arg UpdateClientInstallationParams) error
	UpdateClientInstallationFile(ctx context.Context, arg UpdateClientInstallationFileParams) error
	UpdateClientInstallationPayment(ctx context.Context, arg UpdateClientInstallationPaymentParams) error
	UpdateClientInstallationProduct(ctx context.Context, arg UpdateClientInstallationProductParams) error
	UpdateClientLoan(ctx context.Context, arg UpdateClientLoanParams) error
	UpdateCompany(ctx context.Context, arg UpdateCompanyParams) error
	UpdateEnergiesCost(ctx context.Context, arg UpdateEnergiesCostParams) error
	UpdateEquipmentCatalogue(ctx context.Context, arg UpdateEquipmentCatalogueParams) error
	UpdateEquipmentDesignation(ctx context.Context, arg UpdateEquipmentDesignationParams) error
	UpdateEquipmentInstaller(ctx context.Context, arg UpdateEquipmentInstallerParams) error
	UpdateEquipmentProduct(ctx context.Context, arg UpdateEquipmentProductParams) error
	UpdateEquipmentService(ctx context.Context, arg UpdateEquipmentServiceParams) error
	UpdateGeo(ctx context.Context, arg UpdateGeoParams) error
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) error
	UpdateGroupPrivilege(ctx context.Context, arg UpdateGroupPrivilegeParams) error
	UpdateGroupPrivilegeLnk(ctx context.Context, arg UpdateGroupPrivilegeLnkParams) error
	UpdateGroupUserLnk(ctx context.Context, arg UpdateGroupUserLnkParams) error
	UpdateLog(ctx context.Context, arg UpdateLogParams) error
	UpdateSalesCommission(ctx context.Context, arg UpdateSalesCommissionParams) error
	UpdateSalesCommissionDiscount(ctx context.Context, arg UpdateSalesCommissionDiscountParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
