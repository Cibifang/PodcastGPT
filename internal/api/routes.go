package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"your-project-name/internal/entities"
	"your-project-name/internal/services"
)

func SetupRoutes(router *mux.Router, podcastService services.PodcastService) {
	router.HandleFunc("/podcasts", getPodcastsHandler(podcastService)).Methods(http.MethodGet)
	router.HandleFunc("/podcasts", createPodcastHandler(podcastService)).Methods(http.MethodPost)
}

func getPodcastsHandler(podcastService services.PodcastService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		podcasts, err := podcastService.GetAllPodcasts()
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err, "Failed to get all podcasts")
			return
		}
		writeJSONResponse(w, http.StatusOK, podcasts)
	}
}

func createPodcastHandler(podcastService services.PodcastService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var podcast entities.Podcast
		if err := json.NewDecoder(r.Body).Decode(&podcast); err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err, "Invalid request payload")
			return
		}

		if err := podcastService.CreatePodcast(&podcast); err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err, "Failed to create podcast")
			return
		}

		writeJSONResponse(w, http.StatusCreated, &podcast)
	}
}