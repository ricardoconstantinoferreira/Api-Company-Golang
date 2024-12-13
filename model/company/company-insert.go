package company

import (
	"company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	var company structs.Company
	json.NewDecoder(r.Body).Decode(&company)

	err := createCompany(db, company)

	if err != nil {
		http.Error(w, "Failed create company", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Company created successfully!")

}

func createCompany(db *sql.DB, company structs.Company) error {
	query := "INSERT INTO company (id, corporatereason, legalname, cnpj, mei, email, simple, address) values (?,?,?,?,?,?,?,?)"
	_, err := db.Exec(query, 0, company.CorporateReason, company.LegalName, company.CNPJ,
		company.MEI, company.Email, company.Simple, company.Address)

	if err != nil {
		return err
	}

	return nil
}
