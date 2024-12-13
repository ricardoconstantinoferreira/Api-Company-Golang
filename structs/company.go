package structs

type Company struct {
	Id              int    `json:"id,omitempty"`
	CorporateReason string `json:"corporatereason,omitempty"`
	LegalName       string `json:"legalname,omitempty"`
	CNPJ            string `json:"cnpj,omitempty"`
	MEI             bool   `json:"mei,omitempty"`
	Email           string `json:"email,omitempty"`
	Simple          bool   `json:"simple,omitempty"`
	Address         string `json:"address,omitempty"`
}
