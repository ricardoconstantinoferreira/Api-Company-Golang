package sales

import (
	"company/structs"
)

func GrandTotalSum(total ...structs.SalesItems) float64 {

	grandTotal := 0.00

	for i := 0; i < len(total); i++ {
		grandTotal += total[i].PriceItem
	}

	return grandTotal
}
