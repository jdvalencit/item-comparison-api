package dto

// ProductResponse represents the structure for product data sent in responses
type ProductResponse struct {
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	Description    string            `json:"description,omitempty"`
	Price          float32           `json:"price"`
	Brand          string            `json:"brand"`
	ImageUrl       string            `json:"image_url"`
	Rating         float32           `json:"rating"`
	Specifications map[string]string `json:"specifications"`
	SellerID       string            `json:"seller_id"`
}
