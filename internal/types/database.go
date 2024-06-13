package types

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	// Check for JSON null value
	if string(data) == "null" {
		ns.String, ns.Valid = "", false
		return nil
	}

	// Unmarshal the string value
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	ns.String, ns.Valid = s, true
	return nil
}
