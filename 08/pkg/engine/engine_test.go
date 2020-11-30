package engine

import (
	"testing"

	"github.com/romanserikov/thinknetica-go/08/pkg/index"
	"github.com/romanserikov/thinknetica-go/08/pkg/storage"
	"github.com/romanserikov/thinknetica-go/08/pkg/storage/bst"
)

func TestService_Search(t *testing.T) {
	ind := index.New()
	strg := bst.New()

	doc := storage.Document{
		URL:   "http://go.dev",
		Title: "About go",
	}

	doc.ID = strg.Insert(doc)
	ind.Add(doc)

	service := New(ind, strg)
	resp := service.Search("about")

	if len(resp) == 0 {
		t.Errorf("got %d, want %d", len(resp), 1)
	}

	if got := resp["http://go.dev"]; got != "About go" {
		t.Errorf("got %v want %v", got, "About go")
	}
}
