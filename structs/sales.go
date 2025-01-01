package structs

type Sales struct {
	Id         int     `json:"id,omitempty"`
	Employee   int     `json:"employee_id,omitempty"`
	GrandTotal float64 `json:"grand_total,omitempty"`
}
