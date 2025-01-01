package products

import (
	"company/db"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	productId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Error productId", http.StatusBadRequest)
		return
	}

	product := deleteProductById(db, productId)

	if product != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Product deleted successfully")
}

func deleteProductById(db *sql.DB, productId int) error {
	query := "delete from products where id = ?"
	_, err := db.Exec(query, productId)

	if err != nil {
		return err
	}

	return nil
}
