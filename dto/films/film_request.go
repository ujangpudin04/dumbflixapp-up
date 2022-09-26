package filmsdto

type CreateFilmRequest struct {
	Title         string `json:"title" form:"title" validate:"required" `
	Thumbnailfilm string `json:"thumbnailfilm" form:"thumbnailfilm" validate:"required"`
	Year          string `json:"year" form:"year"`
	Linkfilm      string `json:"linkfilm" form:"linkfilm" validate:"required"`
	CategoryID    string `json:"category_id" form:"category_id"`
	Description   string `json:"description" form:"description"`
}

type UpdateFilmRequest struct {
	Title         string `json:"title" form:"title" `
	Thumbnailfilm string `json:"thumbnailfilm" form:"thumbnailfilm" `
	Year          string `json:"year" form:"year" `
	Linkfilm      string `json:"linkfilm" form:"linkfilm"`
	CategoryID    string `json:"category_id" form:"category_id" `
	Description   string `json:"description" form:"description" `
}
