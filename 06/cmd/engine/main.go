package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/romanserikov/thinknetica-go/06/pkg/engine"
	"github.com/romanserikov/thinknetica-go/06/pkg/spider"
)

func main() {
	engine := new(engine.Engine)
	engine.Crawler = spider.New()
	engine.CacheFile = "documents.json"

	sites := []string{"https://go.dev"}
	depth := 2

	if err := engine.Start(sites, depth); err != nil {
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
		response := engine.Search(words)

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
