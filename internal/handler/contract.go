package handler

import (
	"time"

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

type OrderViewContract struct {
	ID        int64                    `json:"id"`
	UserID    int64                    `json:"user_id"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
	Items     []map[string]interface{} `json:"items"`
}

func toOrderViewsContract(ovs []*types.OrderView) []*OrderViewContract {
	var orders []*OrderViewContract

	// Create a map to group data by order ID
	groupedList := make(map[int64][]*types.OrderView)

	for _, ov := range ovs {
		groupedList[ov.ID] = append(groupedList[ov.ID], ov)
	}

	// Loop the second time to create the contract
	for orderID, ov := range groupedList {
		var ovc OrderViewContract
		ovc.ID = orderID
		ovc.UserID = ov[0].UserID
		ovc.CreatedAt = ov[0].CreatedAt
		ovc.UpdatedAt = ov[0].UpdatedAt

		for _, item := range ov {

			ovc.Items = append(ovc.Items, map[string]interface{}{
				"book_id":     item.ID,
				"book_title":  item.BookTitle,
				"book_author": util.GetNullableString(item.BookAuthor),
				"quantity":    item.Quantity,
			})
		}

		orders = append(orders, &ovc)
	}

	return orders
}
