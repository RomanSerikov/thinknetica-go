package local

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/romanserikov/thinknetica-go/08/pkg/storage"
)

// Service - empty struct for local cache service
type Service struct{}

// New - create local cache service
func New() *Service {
	return new(Service)
}

// Set - set cache for specified key
func (s *Service) Set(key string, value []storage.Document) error {
	f, err := os.Create(key)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}

	if _, err := f.Write(data); err != nil {
		return err
	}

	return nil
}

// Get - gets cache by key
func (s *Service) Get(key string) ([]storage.Document, error) {
	file, err := ioutil.ReadFile(key)
	if err != nil {
		return nil, err
	}

	var docs []storage.Document
	if err := json.Unmarshal(file, &docs); err != nil {
		return nil, err
	}

	return docs, nil
}

// Exists - checks if key exist in cache
func (s *Service) Exists(key string) bool {
	info, err := os.Stat(key)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
