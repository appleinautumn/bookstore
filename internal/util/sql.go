package util

import (
	"gotu/bookstore/internal/types"
)

func GetNullableString(ns types.NullString) interface{} {
	if ns.Valid {
		return ns.String
	}

	return nil
}
