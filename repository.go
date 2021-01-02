package ash

import "net/url"

type Repository struct {
	URL      *url.URL
	Location string
}

func NewRepository(URL url.URL, location string) (*Repository, error) {
	return &Repository{
		URL:      &URL,
		Location: location,
	}, nil
}

func (r *Repository) Clone() error { return nil }
