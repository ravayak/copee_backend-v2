package companies

type Company struct {
	CID           int32  `json:"c_id"`
	CGeoID        int32  `json:"c_geo_id"`
	CName         string `json:"c_name"`
	CRcs          string `json:"c_rcs"`
	CSiret        string `json:"c_siret"`
	CSiren        string `json:"c_siren"`
	CIntraEuVat   string `json:"c_intra_eu_vat"`
	CLegalStatus  string `json:"c_legal_status"`
	CCreationDate string `json:"c_creation_date"`
	CCapital      int32  `json:"c_capital"`
}
