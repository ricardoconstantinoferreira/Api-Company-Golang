package structs

type LoginEmployee struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
