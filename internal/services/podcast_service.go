package services

import (
	"errors"

	"your-project-name/internal/entities"
	"your-project-name/internal/repositories"
)

type PodcastService interface {
	GetAllPodcasts() ([]*entities.Podcast, error)
	CreatePodcast(podcast *entities.Podcast) error
}

type podcastService struct {
	podcastRepo repositories.PodcastRepository
}

func NewPodcastService(podcastRepo repositories.PodcastRepository) PodcastService {
	return &podcastService{
		podcastRepo: podcastRepo,
	}
}

func (s *podcastService) GetAllPodcasts() ([]*entities.Podcast, error) {
	return s.podcastRepo.FindAll()
}

func (s *podcastService) CreatePodcast(podcast *entities.Podcast) error {
	if podcast.Title == "" || podcast.Description == "" || podcast.Uploader == "" || podcast.PublishedAt.IsZero() {
		return errors.New("Missing required fields")
	}
	return s.podcastRepo.Save(podcast)
}