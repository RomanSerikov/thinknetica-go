package spider

import (
	"testing"
)

func TestScan(t *testing.T) {
	if !testing.Short() {
		t.Skip("use -short flag")
	}

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
