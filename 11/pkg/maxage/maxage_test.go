package maxage

import (
	"testing"
)

func TestMaxAge(t *testing.T) {
	tests := []struct {
		name    string
		people  []Person
		wantAge int
	}{
		{
			name: "only employees",
			people: []Person{
				Employee{
					name: "Alice",
					age:  17,
				},
				Employee{
					name: "Bob",
					age:  32,
				},
			},
			wantAge: 32,
		},
		{
			name: "only customers",
			people: []Person{
				Customer{
					age:     42,
					premium: false,
				},
				Customer{
					age:     60,
					premium: true,
				},
			},
			wantAge: 60,
		},
		{
			name: "employees and customers",
			people: []Person{
				Employee{
					name: "Alice",
					age:  17,
				},
				Employee{
					name: "Bob",
					age:  32,
				},
				Customer{
					age:     42,
					premium: false,
				},
				Customer{
					age:     60,
					premium: true,
				},
			},
			wantAge: 60,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge(tt.people...); got != tt.wantAge {
				t.Errorf("got %v, want %v", got, tt.wantAge)
			}
		})
	}
}
