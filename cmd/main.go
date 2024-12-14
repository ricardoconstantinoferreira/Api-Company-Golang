package main

import (
	companyModel "company/model/company"
	employeeModel "company/model/employee"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/create-company", companyModel.CreateCompanyHandler).Methods("POST")
	r.HandleFunc("/get-all-company", companyModel.GetListCompanyHandler).Methods("GET")
	r.HandleFunc("/get-company-by-id/{id}", companyModel.GetListCompanyByIdHandler).Methods("GET")
	r.HandleFunc("/update-company-by-id/{id}", companyModel.UpdateCompanyByIdHandler).Methods("PUT")
	r.HandleFunc("/delete-company-by-id/{id}", companyModel.DeleteCompanyByIdHandler).Methods("DELETE")

	r.HandleFunc("/create-employee", employeeModel.CreateEmployeeHandler).Methods("POST")
	r.HandleFunc("/get-all-employee", employeeModel.GetListEmployeeHandler).Methods("GET")
	r.HandleFunc("/update-employee-by-id/{id}", employeeModel.UpdateEmployeeByIdHandler).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", r))
}
