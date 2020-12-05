package mock

import "github.com/romanserikov/thinknetica-go/09/pkg/storage"

// Service - empty struct for mock cache service
type Service struct{}

// NewCache - create mock cache service
func NewCache() *Service {
	return new(Service)
}

// Set - empty mock method
func (s *Service) Set(key string, value []storage.Document) error {
	return nil
}

// Get - mocked return
func (s *Service) Get(key string) ([]storage.Document, error) {
	docs := []storage.Document{
		{
			URL:   "https://go.dev/about",
			Title: "About - go.dev",
		},
		{
			URL:   "https://go.dev/learn",
			Title: "Getting Started - go.dev",
		},
		{
			URL:   "https://go.dev/solutions",
			Title: "Why Go - go.dev",
		},
	}

	return docs, nil
}

// Exists - empty mocked method
func (s *Service) Exists(key string) bool {
	return true
}
