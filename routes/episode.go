package routes

import (
	"backend_project/handlers"
	"backend_project/pkg/mysql"
	"backend_project/repositories"

	"github.com/gorilla/mux"
)

func EpisodeRoutes(r *mux.Router) {
	EpisodeRepository := repositories.RepositoryEpisode(mysql.DB)
	h := handlers.HandlerEpisode(EpisodeRepository)

	r.HandleFunc("/episode", h.FindEpisode).Methods("GET")
	r.HandleFunc("/episode/{id}", h.GetEpisode).Methods("GET")
	r.HandleFunc("/episode", h.CreateEpisode).Methods("POST")
	r.HandleFunc("/episode/{id}", h.UpdateEpisode).Methods("PATCH")
	r.HandleFunc("/episode/{id}", h.DeleteEpisode).Methods("DELETE")
}
