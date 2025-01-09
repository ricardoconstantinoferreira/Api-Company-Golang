package products

import (
	companyDB "company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := companyDB.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	productId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Impossible change to int", http.StatusInternalServerError)
		return
	}

	var product structs.Products
	err = json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "Impossible change to decode", http.StatusInternalServerError)
		return
	}

	error := updateProductsById(db, productId, product)
	if error != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Product updated successfully")
}

func updateProductsById(db *sql.DB, productId int, product structs.Products) error {
	query := "update products set description = ?, sku = ?, price = ?, user_id = ?, stock_id = ?, qtde = ? where id = ?"
	_, err := db.Exec(query, product.Description, product.Sku, product.Price, product.User, product.Stock, product.Qtde, productId)

	if err != nil {
		return err
	}

	return nil
}

func UpdateQtdeProductsById(db *sql.DB, productId int, qtde int) error {
	query := "update products set qtde = ? where id = ?"
	_, err := db.Exec(query, qtde, productId)

	if err != nil {
		return err
	}

	return nil
}
