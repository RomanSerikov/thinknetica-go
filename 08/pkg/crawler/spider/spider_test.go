package spider

import (
	"testing"
)

func TestScanSite(t *testing.T) {
	if !testing.Short() {
		const url = "https://go.dev"
		const depth = 2
		data, err := New().Scan(url, depth)
		if err != nil {
			t.Fatal(err)
		}

		for k, v := range data {
			t.Logf("%s -> %s\n", k, v)
		}
	}
}
