package util

import "testing"

func TestNullString(t *testing.T) {
	tests := []struct {
		text  string
		want  string
		valid bool
	}{
		{text: "爱情", want: "爱情", valid: true},
		{text: "", want: "", valid: false},
	}

	for i, tt := range tests {
		res := NullString(tt.text)

		if res.String != tt.want || res.Valid != tt.valid {
			t.Errorf(`test #%d: got: %s, want: %s`, i+1, res.String, tt.want)
		}
	}
}

func TestGetNullableString(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		originalString := "爱情"
		ns := NullString(originalString)
		res := GetNullableString(ns)

		if res != originalString {
			t.Errorf("got: %s want: %s", res, originalString)
		}
	})

	t.Run("Empty string", func(t *testing.T) {
		ns := NullString("")
		res := GetNullableString(ns)

		if res != nil {
			t.Errorf("res should be nil")
		}
	})
}
