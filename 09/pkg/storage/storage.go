package storage

import "fmt"

// Document - basic data structure
type Document struct {
	ID    uint
	URL   string
	Title string
}

// Service - storage interface
type Service interface {
	Insert(Document) uint
	Search(id uint) *Document
}

// GenerateTestDocuments - generates documents for benchmark tests
func GenerateTestDocuments(count int) []Document {
	var docs []Document

	for i := 0; i < count; i++ {
		docs = append(docs, Document{
			URL:   "http://go.dev",
			Title: fmt.Sprintf("About go %d", i),
		})
	}

	return docs
}
