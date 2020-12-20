// Package spider реализует сканер содержимого веб-сайтов.
// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.
package spider

import (
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

const maxWorkers = 10

// Service - служба поискового робота.
type Service struct {
	wg *sync.WaitGroup
}

// New - констрктор службы поискового робота.
func New() *Service {
	s := Service{
		wg: new(sync.WaitGroup),
	}
	return &s
}

// BulkScan - scans multiple sites simultaneously
func (s *Service) BulkScan(sites []string, depth int) (map[string]string, error) {
	jobs := make(chan string, len(sites))
	results := make(chan map[string]string, len(sites))
	errors := make(chan error)

	defer close(results)
	defer close(errors)

	for i := 0; i < maxWorkers; i++ {
		s.wg.Add(1)
		go s.worker(jobs, results, depth, errors)
	}

	for _, site := range sites {
		jobs <- site
	}
	close(jobs)

	s.wg.Wait()

	ret := make(map[string]string)
	for i := 0; i < len(sites); i++ {
		select {
		case err := <-errors:
			return nil, err
		case res := <-results:
			for url, title := range res {
				ret[url] = title
			}
		}
	}

	return ret, nil
}

// worker - gets jobs from channel and start Scan
func (s *Service) worker(jobs <-chan string, results chan<- map[string]string, depth int, errs chan<- error) {
	defer s.wg.Done()
	for url := range jobs {
		data, err := s.Scan(url, depth)
		if err != nil {
			errs <- err
		}

		results <- data
	}
}

// Scan осуществляет рекурсивный обход ссылок сайта, указанного в URL,
// с учётом глубины перехода по ссылкам, переданной в depth.
func (s *Service) Scan(url string, depth int) (map[string]string, error) {
	data := make(map[string]string)

	parse(url, url, depth, data)

	return data, nil
}

// parse рекурсивно обходит ссылки на странице, переданной в url.
// Глубина рекурсии задаётся в depth.
// Каждая найденная ссылка записывается в ассоциативный массив
// data вместе с названием страницы.
func parse(url, baseurl string, depth int, data map[string]string) error {
	if depth == 0 {
		return nil
	}

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	page, err := html.Parse(response.Body)
	if err != nil {
		return err
	}

	data[url] = pageTitle(page)

	links := pageLinks(nil, page)
	for _, link := range links {
		// ссылка уже отсканирована
		if data[link] != "" {
			continue
		}
		// ссылка содержит базовый url полностью
		if strings.HasPrefix(link, baseurl) {
			parse(link, baseurl, depth-1, data)
		}
		// относительная ссылка
		if strings.HasPrefix(link, "/") && len(link) > 1 {
			next := baseurl + link[1:]
			parse(next, baseurl, depth-1, data)
		}
	}

	return nil
}

// pageTitle осуществляет рекурсивный обход HTML-страницы и возвращает значение элемента <tittle>.
func pageTitle(n *html.Node) string {
	var title string
	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = pageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

// pageLinks рекурсивно сканирует узлы HTML-страницы и возвращает все найденные ссылки без дубликатов.
func pageLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if !sliceContains(links, a.Val) {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLinks(links, c)
	}
	return links
}

// sliceContains возвращает true если массив содержит переданное значение
func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
