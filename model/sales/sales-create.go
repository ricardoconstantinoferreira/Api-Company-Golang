package sales

import (
	companyDB "company/db"
	productsModel "company/model/products"
	"company/structs"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

func CreateSalesHandler(w http.ResponseWriter, r *http.Request) {
	db := companyDB.Validate(w)
	defer db.Close()

	var sales structs.Sales

	render.DecodeJSON(r.Body, &sales)

	grandTotal := GrandTotalSum(sales.SalesItems...)
	sales.GrandTotal = grandTotal

	lastInsertId, err := createSales(db, sales)

	if err != nil {
		http.Error(w, "Failed create sales", http.StatusInternalServerError)
		return
	}

	validErrors := validQtdeItems(db, sales.SalesItems...)

	if validErrors > 0 {
		http.Error(w, "There is item or more that dont have quantities", http.StatusInternalServerError)
		return
	}

	errorItems := createSalesItems(db, lastInsertId, sales.SalesItems...)

	if errorItems > 0 {
		http.Error(w, "Failed create sales items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Sales created successfully!")

}

func createSalesItems(db *sql.DB, salesId int64, salesItems ...structs.SalesItems) int {

	contError := 0

	for i := 0; i < len(salesItems); i++ {
		var totalItem int = int(salesItems[i].PriceItem * float64(salesItems[i].Qtde))

		query := "insert into sales_items (id, product_id, sales_id, price_item, qtde, total_item) values (?, ?, ?, ?, ?, ?)"
		_, err := db.Exec(query, 0, salesItems[i].ProductId, salesId, salesItems[i].PriceItem, salesItems[i].Qtde, totalItem)

		if err != nil {
			contError++
		} else {
			products, _ := productsModel.GetProductById(db, salesItems[i].ProductId)
			qtde := products.Qtde - salesItems[i].Qtde
			productsModel.UpdateQtdeProductsById(db, salesItems[i].ProductId, qtde)
		}
	}

	return contError
}

func createSales(db *sql.DB, sales structs.Sales) (int64, error) {
	query := "insert into sales (id, employee_id, grand_total) values (?, ?, ?)"
	res, err := db.Exec(query, 0, sales.Employee, sales.GrandTotal)
	id, _ := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func validQtdeItems(db *sql.DB, salesItems ...structs.SalesItems) int {

	contError := 0

	for i := 0; i < len(salesItems); i++ {
		products, _ := productsModel.GetProductById(db, salesItems[i].ProductId)

		if products.Qtde < salesItems[i].Qtde {
			contError++
		}
	}

	return contError
}
