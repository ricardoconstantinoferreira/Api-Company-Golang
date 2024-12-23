package employee

import (
	companyDB "company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	db := companyDB.Validate(w)
	defer db.Close()

	var employee structs.Employee
	json.NewDecoder(r.Body).Decode(&employee)

	if employee.Password != employee.ConfirmPassword {
		http.Error(w, "Please, passwords must be the same", http.StatusInternalServerError)
		return
	}

	err := createEmployee(db, employee)

	if err != nil {
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Employee created successfully!!!")
}

func createEmployee(db *sql.DB, employee structs.Employee) error {

	password, error := companyDB.HashPassword(employee.Password)

	if error != nil {
		panic("Error hash password")
	}

	query := "insert into employee (id, name, document, positionjob, company_id, password ) values (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, 0, employee.Name, employee.Document, employee.PositionJob, employee.Company, password)

	if err != nil {
		return err
	}

	return nil
}
