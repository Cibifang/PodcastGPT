package parsers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"time"

	"your-project-name/internal/entities"
)

type RssParser struct{}

type RssItem struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	PubDate     string    `xml:"pubDate"`
	GUID        string    `xml:"guid"`
	Enclosure   Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	URL    string `xml:"url,attr"`
	Type   string `xml:"type,attr"`
	Length int64  `xml:"length,attr"`
}

type RssFeed struct {
	Title       string    `xml:"channel>title"`
	Description string    `xml:"channel>description"`
	Link        string    `xml:"channel>link"`
	PubDate     string    `xml:"channel>pubDate"`
	Items       []RssItem `xml:"channel>item"`
}

func NewRssParser() *RssParser {
	return &RssParser{}
}

func (p *RssParser) Parse(url string) ([]*entities.Podcast, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Unexpected status code: %d", resp.StatusCode))
	}

	var feed RssFeed
	if err := xml.NewDecoder(resp.Body).Decode(&feed); err != nil {
		return nil, err
	}

	podcasts := make([]*entities.Podcast, 0)
	for _, item := range feed.Items {
		publishedAt, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			return nil, err
		}
		podcast := &entities.Podcast{
			Title:       item.Title,
			Description: item.Description,
			Uploader:    feed.Title,
			PublishedAt: publishedAt,
		}
		if item.Enclosure.URL != "" {
			podcast.AudioURL = item.Enclosure.URL
		} else if item.Link != "" {
			podcast.AudioURL = item.Link
		} else {
			return nil, errors.New("Unable to find audio URL")
		}
		podcasts = append(podcasts, podcast)
	}

	return podcasts, nil
}
