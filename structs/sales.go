package structs

type Sales struct {
	Id         int          `json:"id,omitempty"`
	Employee   int          `json:"employee_id,omitempty"`
	GrandTotal float64      `json:"grand_total,omitempty"`
	SalesItems []SalesItems `json:"sales_items"`
}

type SalesItems struct {
	ProductId int     `json:"product_id,omitempty"`
	PriceItem float64 `json:"price_item,omitempty"`
}
