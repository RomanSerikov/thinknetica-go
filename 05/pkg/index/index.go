package index

import (
	"fmt"
	"strings"
)

// Service - stores search index, documents and documentID counter
type Service struct {
	Index     Index
	Documents *Tree
	LastID    uint
}

// Index - Inverted index for fast document search
// Store unique document ids for each word
type Index map[string]map[uint]struct{}

// New - creates new index Service
func New() *Service {
	return &Service{
		Index:     make(Index),
		Documents: new(Tree),
	}
}

// BuildIndex - takes raw documents from spider package and build index on it.
// Stores Index and Documents in Service object
func (s *Service) BuildIndex(documents map[string]string) {
	if len(documents) == 0 {
		fmt.Println("no documents for indexing")
		return
	}

	for url, title := range documents {
		id := s.Documents.Insert(&Document{
			URL:   url,
			Title: title,
		})

		words := strings.Split(title, " ")

		for _, w := range words {
			word := strings.ToLower(w)

			if _, ok := s.Index[word]; !ok {
				s.Index[word] = make(map[uint]struct{})
			}

			s.Index[word][id] = struct{}{}
		}
	}
}
