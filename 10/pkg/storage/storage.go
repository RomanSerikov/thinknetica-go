package storage

// Хранилище отсканированных документов.

import (
	"github.com/romanserikov/thinknetica-go/10/pkg/crawler"
)

// Interface определяет контракт хранилища данных.
type Interface interface {
	Docs([]int) []crawler.Document
	StoreDocs([]crawler.Document) error
}
