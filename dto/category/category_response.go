package categorydto

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"  form:"name" validate:"required"`
	// Subscribe bool   `json:"subscribe" form:"subscribe" `
}
