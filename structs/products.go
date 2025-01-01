package structs

type Products struct {
	Id          int     `json:"id,omitempty"`
	Description string  `json:"description,omitempty"`
	Sku         string  `json:"sku,omitempty"`
	Price       float64 `json:"price,omitempty"`
	User        int     `json:"user_id,omitempty"`
}