package storage

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
