package episodedto

type CreateEpisodeRequest struct {
	ID            int    `json:"id"`
	Title         string `json:"title" gorm:"type: varchar(255)" validate:"required"`
	Thumbnailfilm string `json:"thumbnailfilm" gorm:"type: varchar(255)" validate:"required"`
	Linkfilm      string `json:"linkfilm" gorm:"type: varchar(255)" validate:"required"`
	FilmID        int    `json:"film_id"`
}

// for association relation with another table (user)
type UpdateEpisodeRequest struct {
	Title         string `json:"title" form:"title" `
	Thumbnailfilm string `json:"thumbnailfilm" form:"thumbnailfilm"`
	Linkfilm      string `json:"linkfilm" form:"linkfilm"`
	FilmID        int    `json:"film_id"`
}
