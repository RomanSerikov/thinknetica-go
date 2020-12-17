package stringwriter

import (
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		want string
	}{
		{
			name: "one string",
			args: []interface{}{"hello"},
			want: "hello",
		},
		{
			name: "two strings",
			args: []interface{}{"hello", "world"},
			want: "helloworld",
		},
		{
			name: "string and numbers",
			args: []interface{}{"hello", 42, 3.14},
			want: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := new(strings.Builder)

			Print(w, tt.args...)

			if got := w.String(); got != tt.want {
				t.Errorf("StringWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
