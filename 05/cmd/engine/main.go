package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/romanserikov/thinknetica-go/05/pkg/index"
	"github.com/romanserikov/thinknetica-go/05/pkg/spider"
)

// Scanner - interface
type Scanner interface {
	Scan(url string, depth int) (data map[string]string, err error)
}

func main() {
	crawler := spider.New()
	sites := []string{"https://go.dev"}
	depth := 2

	scanned, err := search(crawler, sites, depth)
	if err != nil {
		log.Println("error occured while building search index", err)
		return
	}

	ind := index.New()
	ind.BuildIndex(scanned)

	for {
		fmt.Println("Please, enter search request. Type exit to stop.")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("See you next time.")
			return
		}

		docIDs := make(map[uint]struct{})
		words := strings.Split(input, " ")
		for _, word := range words {
			documentIDs, ok := ind.Index[strings.ToLower(word)]
			if !ok {
				continue
			}

			for id := range documentIDs {
				docIDs[id] = struct{}{}
			}
		}

		if len(docIDs) == 0 {
			fmt.Println("Sorry, nothing found.")
			continue
		}

		response := make(map[string]string)
		for id := range docIDs {
			if doc := ind.Documents.Search(id); doc != nil {
				response[doc.URL] = doc.Title
			}
		}

		fmt.Println("Results:")
		for url, title := range response {
			fmt.Printf("  * %s - %s\n", url, title)
		}
	}
}

// search - search for new data from sites
func search(scanner Scanner, sites []string, depth int) (map[string]string, error) {
	data := make(map[string]string)

	for _, site := range sites {
		fmt.Printf("scanning %s...\n", site)
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
