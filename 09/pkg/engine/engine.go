package engine

import (
	"strings"

	"github.com/romanserikov/thinknetica-go/09/pkg/index"
	"github.com/romanserikov/thinknetica-go/09/pkg/storage"
)

// Service - struct for search engine
type Service struct {
	indexer *index.Service
	storage storage.Service
}

// New - creates new engine service
func New(indexer *index.Service, storage storage.Service) *Service {
	return &Service{
		indexer: indexer,
		storage: storage,
	}
}

// Search - search for result
func (s *Service) Search(input string) map[string]string {
	tokens := strings.Split(input, " ")
	docIDs := make(map[uint]struct{})

	for _, token := range tokens {
		documentIDs := s.indexer.Get(token)
		if len(documentIDs) == 0 {
			continue
		}

		for id := range documentIDs {
			docIDs[id] = struct{}{}
		}
	}

	response := make(map[string]string)
	for id := range docIDs {
		if doc := s.storage.Search(id); doc != nil {
			response[doc.URL] = doc.Title
		}
	}

	return response
}
