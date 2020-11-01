package index

import (
	"fmt"
	"sort"
	"strings"
)

// Service - stores search index, documents and documentID counter
type Service struct {
	Index     Index
	Documents []Document
	LastID    uint
}

// Index - Inverted index for fast document search
// Store unique document ids for each word
type Index map[string]map[uint]struct{}

// Document - stores information about web page
type Document struct {
	ID    uint
	URL   string
	Title string
}

// New - creates new index Service
func New() *Service {
	return &Service{
		Index: make(Index),
	}
}

// getID - returns id for new document and increment LastID counter
func (s *Service) getID() uint {
	id := s.LastID
	s.LastID++
	return id
}

// BuildIndex - takes raw documents from spider package and build index on it.
// Stores Index and Documents in Service object
func (s *Service) BuildIndex(documents map[string]string) {
	if len(documents) == 0 {
		fmt.Println("no documents for indexing")
		return
	}

	for url, title := range documents {
		id := s.getID()
		words := strings.Split(title, " ")

		for _, w := range words {
			word := strings.ToLower(w)

			if _, ok := s.Index[word]; !ok {
				s.Index[word] = make(map[uint]struct{})
			}

			s.Index[word][id] = struct{}{}
		}

		s.Documents = append(s.Documents, Document{
			ID:    id,
			URL:   url,
			Title: title,
		})
	}

	sort.Slice(s.Documents, func(i, j int) bool {
		return s.Documents[i].ID < s.Documents[j].ID
	})
}
