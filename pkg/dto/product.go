package dto

// Product -
type Product struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}