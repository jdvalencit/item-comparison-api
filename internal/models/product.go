package models

type Product struct {
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	Description    string            `json:"description,omitempty"`
	Price          float64           `json:"price"`
	Brand          string            `json:"brand"`
	Image_url      string            `json:"image_url"`
	Rating         float32           `json:"rating"`
	Specifications map[string]string `json:"specifications"`
}
