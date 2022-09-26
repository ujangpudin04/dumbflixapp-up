package models

type Film struct {
	ID            int              `json:"id" gorm:"primary_key:auto_increment"`
	Title         string           `json:"title" gorm:"type: varchar(255)"`
	Thumbnailfilm string           `json:"thumbnailfilm" gorm:"type: varchar(255)"`
	Year          string           `json:"year" gorm:"type: varchar(255)"`
	Linkfilm      string           `json:"linkfilm" gorm:"type: varchar(255)"`
	CategoryID    string           `json:"category_id"`
	Category      CategoryResponse `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Description   string           `json:"description"`
}

// for association relation with another table (user)
type FilmResponse struct {
	ID            int              `json:"id"`
	Title         string           `json:"title"`
	Thumbnailfilm string           `json:"thumbnailfilm"`
	Year          string           `json:"year"`
	Linkfilm      string           `json:"linkfilm"`
	CategoryID    string           `json:"category_id"`
	Category      CategoryResponse `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Description   string           `json:"description"`
}

func (FilmResponse) TableName() string {
	return "films"
}
