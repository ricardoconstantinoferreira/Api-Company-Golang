package structs

type Employee struct {
	Id                       int    `json:"id,omitempty"`
	Name                     string `json:"name,omitempty"`
	Document                 string `json:"document,omitempty"`
	PositionJob              string `json:"positionjob,omitempty"`
	Company                  int    `json:"company_id,omitempty"`
	CompanyCorporativeReason string `json:"corporatereason,omitempty"`
	Username                 string `json:"username,omitempty"`
	Password                 string `json:"password,omitempty"`
	ConfirmPassword          string `json:"confirm_password,omitempty"`
}
