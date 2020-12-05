package bst

import (
	"testing"

	"github.com/romanserikov/thinknetica-go/09/pkg/storage"
)

func BenchmarkInsert(b *testing.B) {
	tests := []struct {
		name string
		docs []storage.Document
	}{
		{
			name: "1 document",
			docs: storage.GenerateTestDocuments(1),
		},
		{
			name: "10 documents",
			docs: storage.GenerateTestDocuments(10),
		},
		{
			name: "100 documents",
			docs: storage.GenerateTestDocuments(100),
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {

			for i := 0; i < b.N; i++ {
				store := New()
				for _, doc := range tt.docs {
					id := store.Insert(doc)
					_ = id
				}
			}
		})
	}
}

// BenchmarkInsert/1_document-4         	10687707	       108 ns/op	      80 B/op	       2 allocs/op
// BenchmarkInsert/10_documents-4       	 1000000	      1264 ns/op	     800 B/op	      20 allocs/op
// BenchmarkInsert/100_documents-4      	   30252	     37911 ns/op	    8000 B/op	     200 allocs/op
// PASS
// ok  	github.com/romanserikov/thinknetica-go/09/pkg/storage/bst	4.106s

func BenchmarkSearch(b *testing.B) {
	tests := []struct {
		name string
		docs []storage.Document
	}{
		{
			name: "1 document",
			docs: storage.GenerateTestDocuments(1),
		},
		{
			name: "10 documents",
			docs: storage.GenerateTestDocuments(10),
		},
		{
			name: "100 documents",
			docs: storage.GenerateTestDocuments(100),
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			store := New()

			var id uint
			for _, doc := range tt.docs {
				id = store.Insert(doc)
			}

			for i := 0; i < b.N; i++ {
				doc := store.Search(id)
				_ = doc
			}
		})
	}
}

// BenchmarkSearch/1_document-4          337445816	       3.65 ns/op	       0 B/op	       0 allocs/op
// BenchmarkSearch/10_documents-4       	31778230	       35.8 ns/op	       0 B/op	       0 allocs/op
// BenchmarkSearch/100_documents-4      	 2265498	        546 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/romanserikov/thinknetica-go/09/pkg/storage/bst	4.552s
