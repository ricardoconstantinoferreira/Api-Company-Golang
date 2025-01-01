package products

import (
	"company/db"
	"company/structs"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetListProductsHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	products, err := getListProducts(db)

	if err != nil {
		http.Error(w, "No one products found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getListProducts(db *sql.DB) (*map[int]structs.Products, error) {
	query := "select id, description, sku, price from products"

	result, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	rowsProducts := map[int]structs.Products{}
	products := &structs.Products{}
	cont := 0

	for result.Next() {
		err := result.Scan(&products.Id, &products.Description, &products.Sku, &products.Price)

		if err != nil {
			panic(err.Error())
		}

		rowsProducts[cont] = *products
		cont += 1
	}

	return &rowsProducts, nil
}

func GetListProductsByIdHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Validate(w)
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	productId, err := strconv.Atoi(id)
	product, err := getProductById(db, productId)

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func getProductById(db *sql.DB, productId int) (*structs.Products, error) {
	query := "select id, description, sku, price from products where id = ?"
	result := db.QueryRow(query, productId)

	product := &structs.Products{}

	err := result.Scan(&product.Id, &product.Description, &product.Sku, &product.Price)

	if err != nil {
		return nil, err
	}

	return product, nil
}
