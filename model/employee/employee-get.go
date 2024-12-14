package employee

import (
	"company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetListEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	employees, err := getListEmployee(db)

	if err != nil {
		http.Error(w, "No one company find", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func getListEmployee(db *sql.DB) (*map[int]structs.Employee, error) {
	query := "select e.id, e.name, e.document, e.positionjob, c.corporatereason from employee e " +
		"inner join company c on c.id = e.company_id"

	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	rowsEmployee := map[int]structs.Employee{}
	employee := &structs.Employee{}
	cont := 0

	for results.Next() {

		err := results.Scan(&employee.Id, &employee.Name, &employee.Document, &employee.PositionJob, &employee.CompanyCorporativeReason)

		if err != nil {
			panic(err.Error())
		}

		rowsEmployee[cont] = *employee
		cont += 1
	}

	return &rowsEmployee, nil

}
