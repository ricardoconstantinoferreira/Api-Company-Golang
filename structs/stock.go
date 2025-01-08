package structs

type Stock struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

type StockProduct struct {
	Id      int    `json:"id,omitempty"`
	StockId int    `json:"stock_id,omitempty"`
	Product string `json:"description,omitempty"`
	Qtde    int    `json:"qtde,omitempty"`
}
