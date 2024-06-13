package handler

import (
	"database/sql"
	"testing"
	"time"

	"gotu/bookstore/internal/types"
)

func TestToBookListContract(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name   string
		books  []*types.Book
		expect []map[string]interface{}
	}{
		{
			name: "all fields populated",
			books: []*types.Book{
				{
					ID:          1,
					Title:       "Title1",
					Author:      types.NullString{NullString: sql.NullString{String: "Author1", Valid: true}},
					Description: types.NullString{NullString: sql.NullString{String: "Description1", Valid: true}},
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			expect: []map[string]interface{}{
				{
					"id":          int64(1),
					"title":       "Title1",
					"author":      "Author1",
					"description": "Description1",
					"created_at":  now,
					"updated_at":  now,
				},
			},
		},
		{
			name: "null author and description",
			books: []*types.Book{
				{
					ID:          2,
					Title:       "Title2",
					Author:      types.NullString{NullString: sql.NullString{String: "", Valid: false}},
					Description: types.NullString{NullString: sql.NullString{String: "", Valid: false}},
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			expect: []map[string]interface{}{
				{
					"id":          int64(2),
					"title":       "Title2",
					"author":      nil,
					"description": nil,
					"created_at":  now,
					"updated_at":  now,
				},
			},
		},
		{
			name:   "empty list",
			books:  []*types.Book{},
			expect: []map[string]interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := toBookListContract(tt.books)
			if len(result) != len(tt.expect) {
				t.Errorf("expected length %d, got %d", len(tt.expect), len(result))
			}
			for i := range result {
				for key, expectedValue := range tt.expect[i] {
					if result[i][key] != expectedValue {
						t.Errorf("expected %v for key %q, got %v", expectedValue, key, result[i][key])
					}
				}
			}
		})
	}
}
