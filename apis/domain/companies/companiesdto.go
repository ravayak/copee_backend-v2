package companies

type Company struct {
	CID           int32  `json:"id"`
	CGeoID        int32  `json:"geo_id"`
	CName         string `json:"name"`
	CRcs          string `json:"rcs"`
	CSiret        string `json:"siret"`
	CSiren        string `json:"siren"`
	CIntraEuVat   string `json:"intra_eu_vat"`
	CLegalStatus  string `json:"legal_status"`
	CCreationDate string `json:"creation_date"`
	CCapital      int32  `json:"capital"`
}
