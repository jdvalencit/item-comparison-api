package models

type Product struct {
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	Description    string            `json:"description,omitempty"`
	Price          float32           `json:"price"`
	Brand          string            `json:"brand"`
	SellerID       string            `json:"seller_id"`
	ImageUrl       string            `json:"image_url"`
	Rating         float32           `json:"rating"`
	CreatedAt      string            `json:"created_at"`
	Specifications map[string]string `json:"specifications"`
}
