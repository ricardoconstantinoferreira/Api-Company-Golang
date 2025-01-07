package sales

import (
	"company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetSalesDetailsById(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	salesId, err := strconv.Atoi(id)

	details, err := getSalesDetailsById(db, salesId)

	if err != nil {
		http.Error(w, "Sales not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(details)
}

func getSalesDetailsById(db *sql.DB, salesId int) (*map[int]structs.ReportAll, error) {
	query := "select s.id, e.name, s.grand_total, p.description, si.price_item, si.total_item from sales s " +
		"inner join employee e on e.id = s.employee_id " +
		"inner join sales_items si on si.sales_id = s.id " +
		"inner join products p on p.id = si.product_id  " +
		"where s.id = ?"

	results, err := db.Query(query, salesId)

	if err != nil {
		panic(err.Error())
	}

	rowsResults := map[int]structs.ReportAll{}
	reportAll := &structs.ReportAll{}
	cont := 0

	for results.Next() {
		err := results.Scan(&reportAll.Id, &reportAll.Employee, &reportAll.GrandTotal, &reportAll.Description, &reportAll.PriceItem,
			&reportAll.TotalItem)

		if err != nil {
			panic(err.Error())
		}

		rowsResults[cont] = *reportAll
		cont += 1
	}

	return &rowsResults, nil

}
