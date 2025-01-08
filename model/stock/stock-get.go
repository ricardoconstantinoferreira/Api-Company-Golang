package stock

import (
	"company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetListStockHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	stocks, err := getListStock(db)

	if err != nil {
		http.Error(w, "No one stock found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stocks)
}

func GetListStockByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	stockId, err := strconv.Atoi(id)
	stock, err := getStockById(db, stockId)

	if err != nil {
		http.Error(w, "No one stock found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}

func getStockById(db *sql.DB, id int) (*structs.Stock, error) {
	query := "select id, name, address from stock where id = ?"
	result := db.QueryRow(query, id)

	stock := &structs.Stock{}

	err := result.Scan(&stock.Id, &stock.Name, &stock.Address)

	if err != nil {
		return nil, err
	}

	return stock, nil
}

func getListStock(db *sql.DB) (*map[int]structs.Stock, error) {
	results, err := db.Query("select id, name, address from stock")

	if err != nil {
		panic(err.Error())
	}

	stocks := map[int]structs.Stock{}
	stock := structs.Stock{}
	cont := 0

	for results.Next() {
		err = results.Scan(&stock.Id, &stock.Name, &stock.Address)

		if err != nil {
			panic(err.Error())
		}

		stocks[cont] = stock
		cont += 1
	}

	return &stocks, nil
}
