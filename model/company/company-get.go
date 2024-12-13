package company

import (
	"company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetListCompanyHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	companys, err := getListCompany(db)

	if err != nil {
		http.Error(w, "No one company find", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companys)
}

func GetListCompanyByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	companyId, err := strconv.Atoi(id)
	company, err := getCompanyById(db, companyId)

	if err != nil {
		http.Error(w, "Company not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(company)

}

func getListCompany(db *sql.DB) (*map[int]structs.Company, error) {

	results, err := db.Query("select id, corporatereason, legalname, cnpj, mei, email, simple, address from company")

	if err != nil {
		panic(err.Error())
	}

	resultado := map[int]structs.Company{}

	company := structs.Company{}
	cont := 0

	for results.Next() {

		err = results.Scan(&company.Id, &company.CorporateReason, &company.LegalName, &company.CNPJ,
			&company.MEI, &company.Email, &company.Simple, &company.Address)

		if err != nil {
			panic(err.Error())
		}

		resultado[cont] = company
		cont += 1
	}

	return &resultado, nil
}

func getCompanyById(db *sql.DB, id int) (*structs.Company, error) {

	query := "select id, corporatereason, legalname, cnpj, mei, email, simple, address from company where id = ?"
	result := db.QueryRow(query, id)

	company := &structs.Company{}

	err := result.Scan(&company.Id, &company.CorporateReason, &company.LegalName, &company.CNPJ,
		&company.MEI, &company.Email, &company.Simple, &company.Address)

	if err != nil {
		return nil, err
	}

	return company, nil
}
