package employee

import (
	"company/db"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteEmployeeByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	employeeId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Error employeeId", http.StatusBadRequest)
		return
	}

	employee := deleteEmployeeById(db, employeeId)

	if employee != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Employee deleted successfully")
}

func deleteEmployeeById(db *sql.DB, employeeId int) error {
	query := "delete from employee where id = ?"
	_, err := db.Exec(query, employeeId)

	if err != nil {
		return err
	}

	return nil
}
