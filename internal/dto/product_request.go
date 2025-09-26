package dto

type ProductRequest struct {
	ID             int               `json:"id"`
	Name           string            `json:"name" validate:"required"`
	Description    string            `json:"description,omitempty"`
	Price          float32           `json:"price" validate:"required,gt=0"`
	Brand          string            `json:"brand"`
	ImageUrl       string            `json:"image_url"`
	Rating         float32           `json:"rating"`
	Specifications map[string]string `json:"specifications"`
}
