package index

import (
	"reflect"
	"testing"

	"github.com/romanserikov/thinknetica-go/09/pkg/storage"
)

func TestService_Add(t *testing.T) {
	tests := []struct {
		name    string
		doc     storage.Document
		token   string
		wantIDs []uint
	}{
		{
			name: "test 1",
			doc: storage.Document{
				ID:    1,
				URL:   "http://go.dev",
				Title: "About go",
			},
			token:   "about",
			wantIDs: []uint{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ind := New()
			ind.Add(tt.doc)

			var ids []uint
			for id := range ind.index[tt.token] {
				ids = append(ids, id)
			}

			if !reflect.DeepEqual(ids, tt.wantIDs) {
				t.Errorf("got %v, want %v", ids, tt.wantIDs)
			}
		})
	}
}

func TestService_Get(t *testing.T) {
	tests := []struct {
		name  string
		docs  []storage.Document
		token string
		ids   []uint
	}{
		{
			name: "test 1",
			docs: []storage.Document{
				{
					ID:    1,
					URL:   "http://go.dev",
					Title: "About go",
				},
			},
			token: "about",
			ids:   []uint{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ind := New()
			for _, doc := range tt.docs {
				ind.Add(doc)
			}
			var ids []uint
			for id := range ind.Get(tt.token) {
				ids = append(ids, id)
			}

			if !reflect.DeepEqual(ids, tt.ids) {
				t.Errorf("got %v, want %v", ids, tt.ids)
			}
		})
	}
}
