package entities

import "time"

type Podcast struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Uploader    string     `json:"uploader"`
	PublishedAt time.Time  `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func NewPodcast(title, description, uploader string, publishedAt time.Time) *Podcast {
	now := time.Now()
	return &Podcast{
		Title:       title,
		Description: description,
		Uploader:    uploader,
		PublishedAt: publishedAt,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (p *Podcast) UpdateFields(title, description, uploader string, publishedAt time.Time) {
	p.Title = title
	p.Description = description
	p.Uploader = uploader
	p.PublishedAt = publishedAt
	p.UpdatedAt = time.Now()
}