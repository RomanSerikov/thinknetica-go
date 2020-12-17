package oldest

// Employee - employee user
type Employee struct {
	name string
	age  int
}

// Customer - customer user
type Customer struct {
	premium bool
	age     int
}

// Oldest - returns the oldest person
func Oldest(people ...interface{}) interface{} {
	var maxAge int
	var oldest interface{}

	for _, person := range people {
		if p, ok := person.(Employee); ok {
			if p.age > maxAge {
				oldest = p
				maxAge = p.age
			}
		}

		if p, ok := person.(Customer); ok {
			if p.age > maxAge {
				oldest = p
				maxAge = p.age
			}
		}
	}

	return oldest
}
