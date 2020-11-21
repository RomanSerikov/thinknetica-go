package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/romanserikov/thinknetica-go/07/pkg/crawler"
	"github.com/romanserikov/thinknetica-go/07/pkg/crawler/spider"
	"github.com/romanserikov/thinknetica-go/07/pkg/engine"
)

// Service - struct for gosearch service
type Service struct {
	engine  *engine.Service
	crawler crawler.Scanner
	sites   []string
	depth   int
}

func main() {
	gosearch := New()
	gosearch.Start()
}

// New - creates new gosearch service
func New() *Service {
	crawler := spider.New()

	return &Service{
		engine:  engine.New(crawler, "documents.json"),
		crawler: crawler,
		sites:   []string{"https://go.dev"},
		depth:   2,
	}
}

// Start - starts gosearch service
func (s *Service) Start() {
	if err := s.engine.Start(s.sites, s.depth); err != nil {
		log.Println(err)
		return
	}

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
