package mdb

type Product struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	Proteins      float32 `json:"proteins"`
	Fats          float32 `json:"fast"`
	Carbohydrates float32 `json:"carbohydrates"`
}
