package repositories

import (
	"backend_project/models"

	"gorm.io/gorm"
)

type EpisodeRepository interface {
	FindEpisode() ([]models.Episode, error)
	GetEpisode(ID int) (models.Episode, error)
	CreateEpisode(episode models.Episode) (models.Episode, error)
	UpdateEpisode(episode models.Episode) (models.Episode, error)
	DeleteEpisode(episode models.Episode) (models.Episode, error)
}

func RepositoryEpisode(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindEpisode() ([]models.Episode, error) {
	var episodes []models.Episode
	err := r.db.Preload("Film").Find(&episodes).Error // Using Find method

	return episodes, err
}

func (r *repository) GetEpisode(ID int) (models.Episode, error) {
	var episode models.Episode
	err := r.db.Preload("Film", ID).First(&episode, ID).Error // Using First method

	return episode, err
}

func (r *repository) CreateEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Create(&episode).Error // Using Create method

	return episode, err
}

func (r *repository) UpdateEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Save(&episode).Error // Using Save method

	return episode, err
}

func (r *repository) DeleteEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Delete(&episode).Error // Using Delete method

	return episode, err
}
