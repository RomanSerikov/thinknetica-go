package index

import (
	"strings"

	"github.com/romanserikov/thinknetica-go/08/pkg/storage"
)

// Service - stores search index
type Service struct {
	index map[string]map[uint]struct{}
}

// New - creates new index Service
func New() *Service {
	return &Service{
		index: make(map[string]map[uint]struct{}),
	}
}

// Add - document to search index
func (s *Service) Add(doc storage.Document) {
	tokens := strings.Split(doc.Title, " ")

	for _, t := range tokens {
		token := strings.ToLower(t)

		if _, ok := s.index[token]; !ok {
			s.index[token] = make(map[uint]struct{})
		}

		s.index[token][doc.ID] = struct{}{}
	}
}

// Get - returns document ids by token
func (s *Service) Get(token string) map[uint]struct{} {
	ids, ok := s.index[strings.ToLower(token)]
	if !ok {
		return make(map[uint]struct{})
	}

	return ids
}
