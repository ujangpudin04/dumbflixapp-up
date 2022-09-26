package episodedto

import "backend_project/models"

// for association relation with another table (user)
type EpisodeResponse struct {
	ID            int                 `json:"id"`
	Title         string              `json:"title"`
	Thumbnailfilm string              `json:"thumbnailfilm"`
	Linkfilm      string              `json:"linkfilm"`
	FilmID        string              `json:"film_id"`
	Film          models.FilmResponse `json:"film"`
}
