package util

import (
	"database/sql"
	"testing"

	"gotu/bookstore/internal/types"
)

func TestGetNullableString(t *testing.T) {
	tests := []struct {
		input    types.NullString
		expected interface{}
	}{
		{
			input:    types.NullString{NullString: sql.NullString{String: "test string", Valid: true}},
			expected: "test string",
		},
		{
			input:    types.NullString{NullString: sql.NullString{String: "", Valid: false}},
			expected: nil,
		},
	}

	for _, test := range tests {
		result := GetNullableString(test.input)
		if result != test.expected {
			t.Errorf("GetNullableString(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}
