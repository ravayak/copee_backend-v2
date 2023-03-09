package clientscontroller

import (
	"database/sql"

	// "errors"
	// "fmt"
	// "reflect"

	"github.com/gin-gonic/gin"
	clients "github.com/ravayak/copee_backend/apis/domain/clients"
	data "github.com/ravayak/copee_backend/apis/domain/clients/data"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

// setClientParams sets the client params for create and update
func SetClientParams(clientParamsType string, clientToEdit clients.Client, c *gin.Context) (interface{}, error) {
	if err := c.ShouldBindJSON(&clientToEdit); err != nil {
		return nil, err
	}
	var clientParams interface{}
	if clientParamsType == data.CREATE_CLIENT_PARAMS {
		clientParams = mysql.CreateClientParams{
			ClientLastName:         sql.NullString{String: clientToEdit.ClientLastName, Valid: true},
			ClientFirstName:        sql.NullString{String: clientToEdit.ClientFirstName, Valid: true},
			ClientEmail:            sql.NullString{String: clientToEdit.ClientEmail, Valid: true},
			ClientPhone:            sql.NullString{String: clientToEdit.ClientPhone, Valid: true},
			ClientFiscalYearIncome: sql.NullInt32{Int32: clientToEdit.ClientFiscalYearIncome, Valid: true},
		}
	}

	if clientParamsType == data.UPDATE_CLIENT_PARAMS {
		clientParams = mysql.UpdateClientParams{
			ClientLastName:         sql.NullString{String: clientToEdit.ClientLastName, Valid: true},
			ClientFirstName:        sql.NullString{String: clientToEdit.ClientFirstName, Valid: true},
			ClientEmail:            sql.NullString{String: clientToEdit.ClientEmail, Valid: true},
			ClientPhone:            sql.NullString{String: clientToEdit.ClientPhone, Valid: true},
			ClientFiscalYearIncome: sql.NullInt32{Int32: clientToEdit.ClientFiscalYearIncome, Valid: true},
			ClientID:               clientToEdit.ClientID,
		}
	}

	return clientParams, nil
}

// func convertToMysqlType(source interface{}, destination interface{}) error {
// 	sValue := reflect.ValueOf(source)

// 	if sValue.Kind() != reflect.Struct {
// 		return errors.New("supplied type is not a struct")
// 	}
// 	numFields := sValue.NumField()

// 	for i := 0; i < numFields; i++ {
// 		f := sValue.Elem().Field(i)
// 		field := sValue.Field(i)
// 		fieldType := sValue.Type().Field(i)
// 		fieldName := sValue.Type().Field(i).Name
// 		fmt.Printf("%s: %v |  type: %v\n", fieldName, field.Interface(), fieldType.Type.Kind())
// 		if fieldType.Type.Kind() == reflect.Int32 {
// 			if field.Int() == 0 {
// 				nullInt := sql.NullInt32{Int32: 0, Valid: false}
// 			} else {
// 				nullInt := sql.NullInt32{Int32: int32(field.Int()), Valid: true}
// 			}
// 		}
// 		if fieldType.Type.Kind() == reflect.String {
// 			if field.String() == "" {
// 				nullString := sql.NullString{String: "", Valid: false}
// 			} else {
// 				nullString := sql.NullString{String: field.String(), Valid: true}
// 			}
// 		}
// 	}

// 	fmt.Println("s")
// 	fmt.Println(s)
// 	for i := 0; i < numFields; i++ {
// 		field := sValue.Field(i)
// 		fieldType := sValue.Type().Field(i)
// 		fieldName := sValue.Type().Field(i).Name
// 		fmt.Printf("%s: %v |  type: %v\n", fieldName, field.Interface(), fieldType.Type.Kind())
// 	}
// 	return nil
// }
