package engine

import (
	"testing"

	"github.com/romanserikov/thinknetica-go/09/pkg/index"
	"github.com/romanserikov/thinknetica-go/09/pkg/storage"
	"github.com/romanserikov/thinknetica-go/09/pkg/storage/bst"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name          string
		docs          []storage.Document
		searchRequest string
		want          map[string]string
		wantLen       int
	}{
		{
			name: "one doc, one response",
			docs: []storage.Document{
				{
					URL:   "http://go.dev",
					Title: "About go",
				},
			},
			searchRequest: "about",
			want: map[string]string{
				"http://go.dev": "About go",
			},
			wantLen: 1,
		},
		{
			name: "two docs, one response",
			docs: []storage.Document{
				{
					URL:   "http://go.dev",
					Title: "About go",
				},
				{
					URL:   "http://go.dev/started",
					Title: "getting started",
				},
			},
			searchRequest: "about",
			want: map[string]string{
				"http://go.dev": "About go",
			},
			wantLen: 1,
		},
		{
			name: "two docs, two responses",
			docs: []storage.Document{
				{
					URL:   "http://go.dev",
					Title: "About go",
				},
				{
					URL:   "http://go.dev/about",
					Title: "About me",
				},
			},
			searchRequest: "about",
			want: map[string]string{
				"http://go.dev":       "About go",
				"http://go.dev/about": "About me",
			},
			wantLen: 2,
		},
		{
			name: "two docs, zero responses",
			docs: []storage.Document{
				{
					URL:   "http://go.dev",
					Title: "About go",
				},
				{
					URL:   "http://go.dev/about",
					Title: "About me",
				},
			},
			searchRequest: "zero",
			want:          map[string]string{},
			wantLen:       0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ind := index.New()
			strg := bst.New()

			for i, doc := range tt.docs {
				tt.docs[i].ID = strg.Insert(doc)
				ind.Add(tt.docs[i])
			}

			service := New(ind, strg)
			response := service.Search(tt.searchRequest)

			if got := len(response); got != tt.wantLen {
				t.Errorf("got %v, want %v", got, tt.wantLen)
			}

			for url, wantTitle := range tt.want {
				if gotTitle := response[url]; gotTitle != wantTitle {
					t.Errorf("got %v, want %v", gotTitle, wantTitle)
				}
			}
		})
	}
}
