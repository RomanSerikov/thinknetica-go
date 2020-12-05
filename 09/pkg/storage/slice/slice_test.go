package slice

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

// BenchmarkInsert/1_document-4         	 5503113	       197 ns/op	      80 B/op	       2 allocs/op
// BenchmarkInsert/10_documents-4       	  375548	      3209 ns/op	    2432 B/op	      33 allocs/op
// BenchmarkInsert/100_documents-4      	    5080	    248463 ns/op	   23296 B/op	     306 allocs/op
// PASS
// ok  	github.com/romanserikov/thinknetica-go/09/pkg/storage/slice	3.842s

func BenchmarkSearch(b *testing.B) {
	tests := []struct {
		name string
		docs []storage.Document
	}{
		{
			name: "bst 1 document",
			docs: storage.GenerateTestDocuments(1),
		},
		{
			name: "bst 10 documents",
			docs: storage.GenerateTestDocuments(10),
		},
		{
			name: "bst 100 documents",
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

// BenchmarkSearch/bst_1_document-4         	100000000	        11.7 ns/op	       0 B/op	       0 allocs/op
// BenchmarkSearch/bst_10_documents-4       	62425102	        20.2 ns/op	       0 B/op	       0 allocs/op
// BenchmarkSearch/bst_100_documents-4      	31717461	        37.2 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/romanserikov/thinknetica-go/09/pkg/storage/slice	3.689s
