package dto

type ProductResponse struct {
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	Price          float32           `json:"price"`
	Brand          string            `json:"brand"`
	ImageUrl       string            `json:"image_url"`
	Rating         float32           `json:"rating"`
	Specifications map[string]string `json:"specifications"`
}
