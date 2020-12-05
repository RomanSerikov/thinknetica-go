package cache

import "github.com/romanserikov/thinknetica-go/09/pkg/storage"

// Service - interface for cache service
type Service interface {
	Set(key string, value []storage.Document) error
	Get(key string) ([]storage.Document, error)
	Exists(key string) bool
}
