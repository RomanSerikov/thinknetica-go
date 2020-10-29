package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/romanserikov/thinknetica-go/04/pkg/index"
	"github.com/romanserikov/thinknetica-go/04/pkg/spider"
)

const exit = "exit"

// Scanner - interface
type Scanner interface {
	Scan(url string, depth int) (data map[string]string, err error)
}

func main() {
	crawler := spider.New()
	sites := []string{"https://go.dev"}
	depth := 2

	searchIndex, documents, err := buildSearchIndex(crawler, sites, depth)
	if err != nil {
		log.Println("error occured while building search index", err)
		return
	}

	for {
		fmt.Printf("Please, enter search request. Type %s to stop.\n", exit)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == exit {
			fmt.Println("See you next time.")
			return
		}

		docIDs := make(map[uint]struct{})
		words := strings.Split(input, " ")
		for _, word := range words {
			documentIDs, ok := searchIndex[strings.ToLower(word)]
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
			idx := sort.Search(len(documents), func(j int) bool { return documents[j].ID >= id })
			doc := documents[idx]
			response[doc.URL] = doc.Title
		}

		fmt.Println("Results:")
		for url, title := range response {
			fmt.Printf("  * %s - %s\n", url, title)
		}
	}
}

// buildSearchIndex - builds search index for sites and returns this index and sorted documents with ids
func buildSearchIndex(scanner Scanner, sites []string, depth int) (index.Index, []index.Document, error) {
	i := index.New()

	for _, site := range sites {
		fmt.Printf("scanning %s...\n", site)
		docs, err := scanner.Scan(site, depth)
		if err != nil {
			return nil, nil, err
		}

		i.BuildIndex(docs)
	}

	return i.Index, i.Documents, nil
}
