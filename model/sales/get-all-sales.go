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

func GetAllSales(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	reports, err := getAllSales(db)

	if err != nil {
		http.Error(w, "No one report find", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}

func GetSalesById(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	salesId, err := strconv.Atoi(id)

	reports, err := getSalesById(db, salesId)

	if err != nil {
		http.Error(w, "No one report find", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}

func getSalesById(db *sql.DB, salesId int) (*structs.ReportAll, error) {
	query := "select s.id, e.name, s.grand_total  from sales s " +
		"inner join employee e on e.id = s.employee_id " +
		"where s.id = ?"
	result := db.QueryRow(query, salesId)

	report := &structs.ReportAll{}

	err := result.Scan(&report.Id, &report.Employee, &report.GrandTotal)

	if err != nil {
		return nil, err
	}

	return report, nil
}

func getAllSales(db *sql.DB) (*map[int]structs.ReportAll, error) {
	query := "select s.id, e.name, s.grand_total  from sales s " +
		"inner join employee e on e.id = s.employee_id"

	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	rowsResults := map[int]structs.ReportAll{}
	reportAll := &structs.ReportAll{}
	cont := 0

	for results.Next() {
		err := results.Scan(&reportAll.Id, &reportAll.Employee, &reportAll.GrandTotal)

		if err != nil {
			panic(err.Error())
		}

		rowsResults[cont] = *reportAll
		cont += 1
	}

	return &rowsResults, nil
}
