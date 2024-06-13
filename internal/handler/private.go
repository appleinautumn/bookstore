package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"gotu/bookstore/internal/request"
	"gotu/bookstore/internal/util"
)

func (h *ApiHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIDHeader := r.Header.Get("user_id")

	// convert user id to int64
	userID, err := strconv.ParseInt(userIDHeader, 10, 64)
	if err != nil {
		slog.Error("CreateOrder", slog.Any("error", err))
		util.WriteErrorf(w, http.StatusBadRequest, err)
		return
	}

	// get payload
	var payload request.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		slog.Error("CreateOrder", slog.Any("error", err))
		util.WriteErrorf(w, http.StatusBadRequest, err)
		return
	}

	payload.UserID = userID

	// validate
	if errors := util.ValidateRequest(payload); len(errors) > 0 {
		slog.Error("SignUp", slog.String("validate error", errors[0]))
		util.WriteError(w, http.StatusBadRequest, errors[0])
		return
	}

	// create
	order, err := h.orderService.CreateOrder(ctx, &payload)
	if err != nil {
		slog.Error("SignUp", slog.Any("error", err))
		util.WriteErrorf(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteJson(w, http.StatusOK, order)
}

func (h *ApiHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIDHeader := r.Header.Get("user_id")

	// convert user id to int64
	userID, err := strconv.ParseInt(userIDHeader, 10, 64)
	if err != nil {
		slog.Error("ListOrders", slog.Any("convert error", err))
		util.WriteErrorf(w, http.StatusBadRequest, err)
		return
	}

	// get orders
	orders, err := h.orderService.ListOrdersByUserId(ctx, userID)
	if err != nil {
		slog.Error("ListOrders", slog.Any("list error", err))
		util.WriteErrorf(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteJson(w, http.StatusOK, orders)
}
