package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/romanserikov/thinknetica-go/02/pkg/spider"
)

const exit = "exit"

func main() {
	sites := []string{
		"https://habr.com",
	}

	store := make(map[string]string)
	for _, site := range sites {
		fmt.Printf("Scanning %s\n", site)

		data, err := spider.Scan(site, 2)
		if err != nil {
			log.Fatalf("error during scanning %s: %v\n", site, err)
		}

		for url, title := range data {
			store[url] = title
		}
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
