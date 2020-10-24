package index

import (
	"fmt"
	"sort"
	"strings"

	"github.com/romanserikov/thinknetica-go/04/pkg/spider"
)

// Indexer -
type Indexer struct {
	Scanner
	Index
	Documents []Document
}

// Scanner - interface
type Scanner interface {
	Scan(url string, depth int) (data map[string]string, err error)
}

// Index -
type Index map[string]map[uint]struct{}

// Document -
type Document struct {
	ID    uint
	URL   string
	Title string
}

// New -
func New() *Indexer {
	return &Indexer{
		Scanner: new(spider.Spider),
		Index:   make(Index),
	}
}

// Run -
func (i *Indexer) Run(sites []string, depth int) error {
	if err := i.buildDocuments(sites, depth); err != nil {
		return fmt.Errorf("error during building documents: %w", err)
	}

	return i.buildIndex()
}

func (i *Indexer) buildDocuments(sites []string, depth int) error {
	var documents []Document
	var id uint

	for _, site := range sites {
		fmt.Printf("Scanning %s\n", site)

		data, err := i.Scan(site, depth)
		if err != nil {
			return fmt.Errorf("error during scanning %s: %w", site, err)
		}

		for url, title := range data {
			documents = append(documents, Document{
				ID:    id,
				URL:   url,
				Title: title,
			})

			id++
		}
	}

	sort.Slice(documents, func(i, j int) bool {
		return documents[i].ID < documents[j].ID
	})

	i.Documents = documents

	return nil
}

func (i *Indexer) buildIndex() error {
	if len(i.Documents) == 0 {
		return fmt.Errorf("no documents for indexing")
	}

	for _, doc := range i.Documents {
		words := strings.Split(doc.Title, " ")

		for _, w := range words {
			word := strings.ToLower(w)

			if _, ok := i.Index[word]; !ok {
				i.Index[word] = make(map[uint]struct{})
			}

			i.Index[word][doc.ID] = struct{}{}
		}
	}

	return nil
}

// GetDocument -
func (i *Indexer) GetDocument(id uint) Document {
	idx := sort.Search(len(i.Documents), func(j int) bool { return i.Documents[j].ID >= id })
	return i.Documents[idx]
}
