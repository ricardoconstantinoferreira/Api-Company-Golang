package company

import (
	"company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateCompanyByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	companyId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Impossible change to int", http.StatusInternalServerError)
		return
	}

	var company structs.Company
	err = json.NewDecoder(r.Body).Decode(&company)

	if err != nil {
		http.Error(w, "Impossible decode", http.StatusInternalServerError)
		return
	}

	error := updateCompanyById(db, companyId, company)
	if error != nil {
		http.Error(w, "Company not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Company updated successfully")
}

func updateCompanyById(db *sql.DB, companyId int, company structs.Company) error {
	query := "update company set corporatereason = ?, legalname = ?, cnpj = ?, mei = ?, email = ?, simple = ?, address = ? where id = ?"
	_, err := db.Exec(query, company.CorporateReason, company.LegalName, company.CNPJ, company.MEI, company.Email,
		company.Simple, company.Address, companyId)

	if err != nil {
		return err
	}

	return nil
}
