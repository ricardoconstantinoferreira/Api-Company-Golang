package main

import (
	"company/auth"
	loginHandler "company/handler"
	companyModel "company/model/company"
	employeeModel "company/model/employee"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/login", loginHandler.LoginHandler).Methods("POST")

	privateRouter := r.PathPrefix("/").Subrouter()
	privateRouter.Use(auth.AuthMiddleware)

	r.HandleFunc("/create-company", companyModel.CreateCompanyHandler).Methods("POST")
	privateRouter.HandleFunc("/get-all-company", companyModel.GetListCompanyHandler).Methods("GET")
	privateRouter.HandleFunc("/get-company-by-id/{id}", companyModel.GetListCompanyByIdHandler).Methods("GET")
	privateRouter.HandleFunc("/update-company-by-id/{id}", companyModel.UpdateCompanyByIdHandler).Methods("PUT")
	privateRouter.HandleFunc("/delete-company-by-id/{id}", companyModel.DeleteCompanyByIdHandler).Methods("DELETE")

	r.HandleFunc("/create-employee", employeeModel.CreateEmployeeHandler).Methods("POST")
	privateRouter.HandleFunc("/get-all-employee", employeeModel.GetListEmployeeHandler).Methods("GET")
	privateRouter.HandleFunc("/get-employee-by-id/{id}", employeeModel.GetEmployeeAndCompanyByEmployeeId).Methods("GET")
	privateRouter.HandleFunc("/update-employee-by-id/{id}", employeeModel.UpdateEmployeeByIdHandler).Methods("PUT")
	privateRouter.HandleFunc("/delete-employee-by-id/{id}", employeeModel.DeleteEmployeeByIdHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", r))
}
