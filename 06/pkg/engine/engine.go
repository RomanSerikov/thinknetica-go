package engine

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/romanserikov/thinknetica-go/06/pkg/bst"
	"github.com/romanserikov/thinknetica-go/06/pkg/index"
)

// Scanner - interface
type Scanner interface {
	Scan(url string, depth int) (data map[string]string, err error)
}

// Engine - struct for search engine
type Engine struct {
	indexer   *index.Service
	Crawler   Scanner
	CacheFile string
}

// Start - starts the search engine
func (e *Engine) Start(sites []string, depth int) error {
	if fileExists(e.CacheFile) {
		if err := e.LoadCache(); err != nil {
			return err
		}

		go e.Sync(sites, depth)
		return nil
	}

	e.Sync(sites, depth)
	return nil
}

// LoadCache - loads cache from local file
func (e *Engine) LoadCache() error {
	file, err := ioutil.ReadFile(e.CacheFile)
	if err != nil {
		return err
	}

	var docs []bst.Document
	if err := json.Unmarshal(file, &docs); err != nil {
		return err
	}

	data := make(map[string]string)
	for _, doc := range docs {
		data[doc.URL] = doc.Title
	}

	ind := index.New()
	ind.BuildIndex(data)

	e.indexer = ind

	return nil
}

// Sync - get fresh data from crawler, save it to local file and update indexed documents
func (e *Engine) Sync(sites []string, depth int) {
	data, err := sync(e.Crawler, sites, depth)
	if err != nil {
		log.Println("error occured while getting new data", err)
		return
	}

	var docs []bst.Document
	for url, title := range data {
		docs = append(docs, bst.Document{
			URL:   url,
			Title: title,
		})
	}

	saveDocuments(e.CacheFile, docs)

	ind := index.New()
	ind.BuildIndex(data)

	e.indexer = ind
}

// Search - search for result
func (e *Engine) Search(words []string) map[string]string {
	docIDs := make(map[uint]struct{})
	for _, word := range words {
		documentIDs, ok := e.indexer.Index[strings.ToLower(word)]
		if !ok {
			continue
		}

		for id := range documentIDs {
			docIDs[id] = struct{}{}
		}
	}

	response := make(map[string]string)
	for id := range docIDs {
		if doc := e.indexer.Documents.Search(id); doc != nil {
			response[doc.URL] = doc.Title
		}
	}

	return response
}

// search - search for new data from sites
func sync(scanner Scanner, sites []string, depth int) (map[string]string, error) {
	data := make(map[string]string)

	for _, site := range sites {
		scanned, err := scanner.Scan(site, depth)
		if err != nil {
			return nil, err
		}

		for url, title := range scanned {
			data[url] = title
		}
	}

	return data, nil
}

// saveDocuments - saves data to local cache
func saveDocuments(filename string, docs []bst.Document) {
	f, err := os.Create(filename)
	if err != nil {
		log.Println("error while creating cache", err)
		return
	}
	defer f.Close()

	data, err := json.MarshalIndent(docs, "", "  ")
	if err != nil {
		log.Println("error while creating cache", err)
		return
	}

	if _, err := f.Write(data); err != nil {
		log.Println("error while creating cache", err)
		return
	}
}

// fileExists - checks if file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
