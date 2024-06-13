package types

import (
	"database/sql"
	"encoding/json"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {

	t.Run("success and null", func(t *testing.T) {
		tests := []struct {
			input    string
			expected NullString
		}{
			{
				input:    `"test string"`,
				expected: NullString{sql.NullString{String: "test string", Valid: true}},
			},
			{
				input:    `null`,
				expected: NullString{sql.NullString{String: "", Valid: false}},
			},
		}

		for _, test := range tests {
			var ns NullString
			err := json.Unmarshal([]byte(test.input), &ns)
			if err != nil {
				t.Errorf("UnmarshalJSON(%s) returned error: %v", test.input, err)
			}

			if ns != test.expected {
				t.Errorf("UnmarshalJSON(%s) = %v, want %v", test.input, ns, test.expected)
			}
		}
	})

	t.Run("error - unmarshalJSON", func(t *testing.T) {
		var ns NullString
		if err := json.Unmarshal([]byte("{}"), &ns); err == nil {
			t.Errorf("expecting an error unmarshaling object into string")
		}
	})
}
