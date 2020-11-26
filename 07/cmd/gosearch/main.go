package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/romanserikov/thinknetica-go/07/pkg/crawler/spider"
	"github.com/romanserikov/thinknetica-go/07/pkg/engine"
)

// Server - struct for gosearch service
type Server struct {
	engine *engine.Service
	sites  []string
	depth  int
}

func main() {
	gosearch := New()
	gosearch.Init()
	gosearch.Start()
}

// New - creates new gosearch service
func New() *Server {
	return &Server{
		engine: engine.New(spider.New(), "documents.json"),
		sites:  []string{"https://go.dev"},
		depth:  2,
	}
}

// Init - loads data from cache and sync in routine or start a new sync
func (s *Server) Init() {
	if !s.engine.HasCache() {
		s.engine.Sync(s.sites, s.depth)
		return
	}

	if err := s.engine.LoadCache(); err != nil {
		s.engine.Sync(s.sites, s.depth)
	} else {
		go s.engine.Sync(s.sites, s.depth)
	}
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

		words := strings.Split(input, " ")
		response := s.engine.Search(words)

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
