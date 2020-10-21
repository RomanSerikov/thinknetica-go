package mock

// Scanner -
type Scanner struct{}

// Scan -
func (s *Scanner) Scan(url string, depth int) (map[string]string, error) {
	return map[string]string{
		"https://go.dev/about":     "About - go.dev",
		"https://go.dev/learn":     "Learn - go.dev",
		"https://go.dev/solutions": "Why Go - go.dev",
	}, nil
}
