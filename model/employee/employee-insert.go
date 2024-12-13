package employee

import (
	"company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	var employee structs.Employee
	json.NewDecoder(r.Body).Decode(&employee)

	err := createEmployee(db, employee)

	if err != nil {
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully!!!")
}

func createEmployee(db *sql.DB, employee structs.Employee) error {
	query := "insert into employee (id, name, document, positionjob, company_id ) values (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, 0, employee.Name, employee.Document, employee.PositionJob, employee.Company.Id)

	if err != nil {
		return err
	}

	return nil
}
