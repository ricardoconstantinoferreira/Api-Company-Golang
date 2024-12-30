package products

import (
	companyDB "company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProductsHandler(w http.ResponseWriter, r *http.Request) {
	db := companyDB.Validate(w)
	defer db.Close()

	var products structs.Products
	json.NewDecoder(r.Body).Decode(&products)

	err := createProducts(db, products)

	if err != nil {
		http.Error(w, "Failed to create products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Product created sucessfully!!!")
}

func createProducts(db *sql.DB, products structs.Products) error {
	query := "insert into products (id, description, sku, price, user_id) values (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, 0, products.Description, products.Sku, products.Price, products.User)

	if err != nil {
		return err
	}

	return nil
}
