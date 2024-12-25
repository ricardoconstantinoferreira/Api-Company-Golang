package employee

import (
	companyDB "company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateEmployeeByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := companyDB.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	employeeId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Impossible change to int", http.StatusInternalServerError)
	}

	var employee structs.Employee
	err = json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		http.Error(w, "Impossible decode", http.StatusInternalServerError)
	}

	error := updateEmployeeById(db, employeeId, employee)
	if error != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Employee updated successfully")
}

func updateEmployeeById(db *sql.DB, employeeId int, employee structs.Employee) error {

	password, error := companyDB.HashPassword(employee.Password)

	if error != nil {
		panic("Error hash password")
	}

	query := "update employee set name = ?, document = ?, positionjob = ?, company_id = ?, username = ?, password = ? where id = ?"
	_, err := db.Exec(query, employee.Name, employee.Document, employee.PositionJob, employee.Company, employee.Username, password, employeeId)

	if err != nil {
		return err
	}

	return nil
}
