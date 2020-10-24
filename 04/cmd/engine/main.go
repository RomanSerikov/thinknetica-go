package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/romanserikov/thinknetica-go/04/pkg/index"
)

const exit = "exit"

func main() {
	sites := []string{"https://go.dev"}
	depth := 2

	indexer := index.New()
	if err := indexer.Run(sites, depth); err != nil {
		log.Println("error occured while running indexer", err)
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
			ids, ok := indexer.Index[strings.ToLower(word)]
			if !ok {
				continue
			}

			for id := range ids {
				docIDs[id] = struct{}{}
			}
		}

		if len(docIDs) == 0 {
			fmt.Println("Sorry, nothing found.")
			continue
		}

		response := make(map[string]string)
		for id := range docIDs {
			doc := indexer.GetDocument(id)
			response[doc.URL] = doc.Title
		}

		fmt.Println("Results:")
		for url, title := range response {
			fmt.Printf("  * %s - %s\n", url, title)
		}
	}
}
