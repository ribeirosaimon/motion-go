package dto

type Order struct {
	Id         string  `json:"id"`
	Value      float64 `json:"value"`
	Quantity   float64 `json:"quantity"`
	IsNational bool    `json:"isNational"`
}
