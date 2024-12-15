package company

import (
	"company/db"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteCompanyByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	companyId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Error companyId", http.StatusBadRequest)
		return
	}

	company := deleteCompanyById(db, companyId)

	if company != nil {
		http.Error(w, "Company not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Company deleted successfully")
}

func deleteCompanyById(db *sql.DB, id int) error {
	query := "delete from company where id = ?"
	_, err := db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
