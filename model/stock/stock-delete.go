package stock

import (
	"company/db"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteStockByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	stockId, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Error StockId", http.StatusBadRequest)
		return
	}

	stock := deleteStockById(db, stockId)

	if stock != nil {
		http.Error(w, "Stock not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Stock deleted successfully")
}

func deleteStockById(db *sql.DB, id int) error {
	query := "delete from stock where id = ?"
	_, err := db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
