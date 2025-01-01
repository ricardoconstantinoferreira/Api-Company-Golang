package structs

type SalesItems struct {
	Id      int `json:"id,omitempty"`
	Product int `json:"product_id,omitempty"`
	Sales   int `json:"sales_id,omitempty"`
}
