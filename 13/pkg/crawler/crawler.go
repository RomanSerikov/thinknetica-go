package crawler

// Scanner - interface
type Scanner interface {
	Scan(url string, depth int) (data map[string]string, err error)
	BulkScan(sites []string, depth int) (map[string]string, error)
}
