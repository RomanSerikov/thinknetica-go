package oldest

import (
	"reflect"
	"testing"
)

func TestOldest(t *testing.T) {
	tests := []struct {
		name   string
		people []Person
		want   Person
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
			want: Employee{
				name: "Bob",
				age:  32,
			},
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
			want: Customer{
				age:     60,
				premium: true,
			},
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
			want: Customer{
				age:     60,
				premium: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Oldest(tt.people...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
