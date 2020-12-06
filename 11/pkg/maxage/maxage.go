package maxage

// Person - interface for user with age
type Person interface {
	Age() int
}

// Employee - employee user
type Employee struct {
	name string
	age  int
}

// Age - getter for age field
func (e Employee) Age() int {
	return e.age
}

// Customer - customer user
type Customer struct {
	premium bool
	age     int
}

// Age - getter for age field
func (c Customer) Age() int {
	return c.age
}

// MaxAge - returns age of the oldest person
func MaxAge(people ...Person) int {
	var max int

	for _, person := range people {
		if max < person.Age() {
			max = person.Age()
		}
	}

	return max
}
