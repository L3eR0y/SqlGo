package handlers

import (
	"encoding/json"
	"fmt"
	"internal/mdb"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Root Handler work!")
}

func GetProducts(w http.ResponseWriter, _ *http.Request) {
	items := []mdb.Product{
		{
			Id:            1,
			Name:          "q1",
			Proteins:      0.1,
			Fats:          0.1,
			Carbohydrates: 0.1,
		},
		{
			Id:            1,
			Name:          "q1",
			Proteins:      0.1,
			Fats:          0.1,
			Carbohydrates: 0.1,
		},
		{
			Id:            1,
			Name:          "q1",
			Proteins:      0.1,
			Fats:          0.1,
			Carbohydrates: 0.1,
		},
		{
			Id:            1,
			Name:          "q1",
			Proteins:      0.1,
			Fats:          0.1,
			Carbohydrates: 0.1,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// func getProductById() {}
// func createProducs()  {}
// func deleteProduct()  {}
// func updateItem()     {}
