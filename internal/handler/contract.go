package handler

import (
	"gotu/bookstore/internal/types"
	"gotu/bookstore/internal/util"
)

func toBookListContract(books []*types.Book) []map[string]interface{} {
	var res []map[string]interface{}

	for _, b := range books {
		res = append(res, map[string]interface{}{
			"id":          b.ID,
			"title":       b.Title,
			"author":      util.GetNullableString(b.Author),
			"description": util.GetNullableString(b.Description),
			"created_at":  b.CreatedAt,
			"updated_at":  b.UpdatedAt,
		})
	}

	return res
}
