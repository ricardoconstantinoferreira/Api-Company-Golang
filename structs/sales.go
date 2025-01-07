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
	Qtde      int     `json:"qtde,omitempty"`
	TotalItem float64 `json:"total_item,omitempty"`
}

type ReportAll struct {
	Id          int     `json:"id,omitempty"`
	Employee    string  `json:"name,omitempty"`
	GrandTotal  float64 `json:"grand_total,omitempty"`
	Description string  `json:"description,omitempty"`
	PriceItem   float64 `json:"price_item,omitempty"`
	TotalItem   float64 `json:"total_item,omitempty"`
}
