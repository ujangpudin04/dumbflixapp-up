package models

type Episode struct {
	ID            int          `json:"id" gorm:"primary_key:auto_increment"`
	Title         string       `json:"title" gorm:"type: varchar(255)"`
	Thumbnailfilm string       `json:"thumbnailfilm" gorm:"type: varchar(255)"`
	Linkfilm      string       `json:"linkfilm" gorm:"type: varchar(255)"`
	FilmID        int          `json:"film_id"`
	Film          FilmResponse `json:"film" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type EpisodeResponse struct {
	ID            int          `json:"id"`
	Title         string       `json:"title"`
	Thumbnailfilm string       `json:"thumbnailfilm"`
	Linkfilm      string       `json:"linkfilm"`
	FilmID        int          `json:"film_id"`
	Film          FilmResponse `json:"film" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (EpisodeResponse) TableName() string {
	return "episode"
}
