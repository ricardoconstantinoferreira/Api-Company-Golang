package stock

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

func UpdateStockByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	stockId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Impossible change to int", http.StatusInternalServerError)
		return
	}

	var stock structs.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		http.Error(w, "Impossible decode", http.StatusInternalServerError)
		return
	}

	error := updateStockById(db, stockId, stock)
	if error != nil {
		http.Error(w, "Stock not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Stock updated successfully")
}

func updateStockById(db *sql.DB, stockId int, stock structs.Stock) error {
	query := "update stock set name = ?, address = ? where id = ?"
	_, err := db.Exec(query, stock.Name, stock.Address, stockId)

	if err != nil {
		return err
	}

	return nil
}
