package employee

import (
	"company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func GetEmployeeAndCompanyByEmployeeId(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	employeeId, err := strconv.Atoi(id)

	employee, err := getEmployeeAndCompanyByEmployeeId(db, employeeId)

	if err != nil {
		http.Error(w, "Company not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

func getEmployeeAndCompanyByEmployeeId(db *sql.DB, employeeId int) (*structs.Employee, error) {
	query := "select e.id, e.name, e.document, e.positionjob, c.corporatereason from employee e " +
		"inner join company c on c.id = e.company_id " +
		"where e.id = ?"

	result := db.QueryRow(query, employeeId)

	employee := &structs.Employee{}
	err := result.Scan(&employee.Id, &employee.Name, &employee.Document, &employee.PositionJob, &employee.CompanyCorporativeReason)

	if err != nil {
		return nil, err
	}

	return employee, nil
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
