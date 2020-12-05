package slice

import (
	"sort"

	"github.com/romanserikov/thinknetica-go/09/pkg/storage"
)

// Slice - struct to store documents and current document id
type Slice struct {
	docs      []storage.Document
	currentID uint
}

// New - creates Slice object
func New() *Slice {
	return &Slice{
		docs: make([]storage.Document, 0),
	}
}

// Insert document, autoincrements currentID, returns id of inserted document
func (s *Slice) Insert(doc storage.Document) uint {
	doc.ID = s.currentID
	s.currentID++

	s.docs = append(s.docs, doc)

	sort.Slice(s.docs, func(i, j int) bool {
		return s.docs[i].ID < s.docs[j].ID
	})

	return doc.ID
}

// Search for document
func (s *Slice) Search(docID uint) *storage.Document {
	indx := sort.Search(len(s.docs), func(j int) bool { return s.docs[j].ID >= docID })
	if indx < len(s.docs) && s.docs[indx].ID == docID {
		return &s.docs[indx]
	}

	return nil
}
