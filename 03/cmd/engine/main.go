package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/romanserikov/thinknetica-go/03/pkg/spider"
)

const exit = "exit"

// Scanner - interface
type Scanner interface {
	Scan(url string, depth int) (data map[string]string, err error)
}

func main() {
	scanner := new(spider.Spider)
	sites := []string{"https://go.dev"}
	depth := 2

	store, err := buildStore(scanner, sites, depth)
	if err != nil {
		log.Fatal(err)
	}

	var input string

	for {
		fmt.Printf("Please, enter search request. Type %s to stop.\n", exit)
		fmt.Scanln(&input)

		if input == exit {
			fmt.Println("See you next time.")
			return
		}

		request := strings.ToLower(input)
		response := make(map[string]string)

		for url, title := range store {
			if strings.Contains(strings.ToLower(url), request) || strings.Contains(strings.ToLower(title), request) {
				response[url] = title
			}
		}

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

func buildStore(s Scanner, sites []string, depth int) (map[string]string, error) {
	store := make(map[string]string)

	for _, site := range sites {
		fmt.Printf("Scanning %s\n", site)

		data, err := s.Scan(site, 2)
		if err != nil {
			return nil, fmt.Errorf("error during scanning %s: %w", site, err)
		}

		for url, title := range data {
			store[url] = title
		}
	}

	return store, nil
}
