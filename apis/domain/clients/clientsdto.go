package clients

import "time"

type Client struct {
	ClientID               int32     `json:"client_id"`
	ClientLastName         string    `json:"client_last_name"`
	ClientFirstName        string    `json:"client_first_name"`
	ClientEmail            string    `json:"client_email"`
	ClientPhone            string    `json:"client_phone"`
	ClientFiscalYearIncome int32     `json:"client_fiscal_year_income"`
	ClientDateCreated      time.Time `json:"client_date_created"`
}
