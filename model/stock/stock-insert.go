package stock

import (
	"company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateStockHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	var stock structs.Stock
	json.NewDecoder(r.Body).Decode(&stock)

	err := createStock(db, stock)

	if err != nil {
		http.Error(w, "Failed create company", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Stock created successfully!")
}

func createStock(db *sql.DB, stock structs.Stock) error {
	query := "insert into stock (id, name, address) values (?, ?, ?)"
	_, err := db.Exec(query, 0, stock.Name, stock.Address)

	if err != nil {
		return err
	}

	return nil
}
