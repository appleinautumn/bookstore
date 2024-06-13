package util

import "database/sql"

func NullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func GetNullableString(ns sql.NullString) interface{} {
	if ns.Valid {
		return ns.String
	}

	return nil
}
