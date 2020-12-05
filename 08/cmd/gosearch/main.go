package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/romanserikov/thinknetica-go/08/pkg/cache"
	"github.com/romanserikov/thinknetica-go/08/pkg/cache/local"
	"github.com/romanserikov/thinknetica-go/08/pkg/crawler"
	"github.com/romanserikov/thinknetica-go/08/pkg/crawler/spider"
	"github.com/romanserikov/thinknetica-go/08/pkg/engine"
	"github.com/romanserikov/thinknetica-go/08/pkg/index"
	"github.com/romanserikov/thinknetica-go/08/pkg/storage"
	"github.com/romanserikov/thinknetica-go/08/pkg/storage/bst"
)

// Server - struct for gosearch service
type Server struct {
	crawler crawler.Scanner
	index   *index.Service
	storage storage.Service
	engine  *engine.Service
	cache   cache.Service

	sites []string
	depth int
}

func main() {
	gosearch := New()
	if err := gosearch.Init(); err != nil {
		log.Fatal(err)
	}

	gosearch.Start()
}

// New - creates new gosearch service
func New() *Server {
	s := new(Server)

	s.crawler = spider.New()
	s.index = index.New()
	s.storage = bst.New()
	s.cache = local.New()
	s.engine = engine.New(s.index, s.storage)

	s.sites = []string{"https://go.dev"}
	s.depth = 2

	return s
}

// Init - loads data from cache and sync in routine or start a new sync
func (s *Server) Init() error {
	if !s.cache.Exists("documents.json") {
		return s.sync()
	}

	docs, err := s.cache.Get("documents.json")
	if err != nil {
		return s.sync()
	}

	for _, doc := range docs {
		doc.ID = s.storage.Insert(doc)
		s.index.Add(doc)
	}

	go s.sync()

	return nil
}

func (s *Server) sync() error {
	var docs []storage.Document

	for _, site := range s.sites {
		data, err := s.crawler.Scan(site, s.depth)
		if err != nil {
			return err
		}

		for url, title := range data {
			doc := storage.Document{
				URL:   url,
				Title: title,
			}

			doc.ID = s.storage.Insert(doc)
			s.index.Add(doc)
			docs = append(docs, doc)
		}
	}

	return s.cache.Set("documents.json", docs)
}

// Start - starts gosearch service
func (s *Server) Start() {
	for {
		fmt.Println("Please, enter search request. Type exit to stop.")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("See you next time.")
			return
		}

		response := s.engine.Search(input)
		if len(response) == 0 {
			fmt.Println("Sorry, nothing found.")
			continue
		}

		fmt.Println("Results:")
		for url, title := range response {
			fmt.Printf("  * %s - %s\n", url, title)
		}
	}
}
