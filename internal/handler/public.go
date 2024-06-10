package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"gotu/bookstore/internal/request"
	"gotu/bookstore/internal/service"
	"gotu/bookstore/internal/util"
)

type ApiHandler struct {
	bookService  *service.BookService
	userService  *service.UserService
	orderService *service.OrderService
}

func NewApiHandler(sv1 *service.BookService, sv2 *service.UserService, sv3 *service.OrderService) *ApiHandler {
	return &ApiHandler{
		bookService:  sv1,
		userService:  sv2,
		orderService: sv3,
	}
}

func (h *ApiHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get books
	books, err := h.bookService.List(ctx)
	if err != nil {
		slog.Error("ListBooks", slog.Any("error", err))
		util.WriteErrorf(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteJson(w, http.StatusOK, books)
}

func (h *ApiHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get payload
	var payload request.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		util.WriteErrorf(w, http.StatusBadRequest, err)
		return
	}

	// validate
	if errors := util.ValidateRequest(payload); len(errors) > 0 {
		slog.Error("SignUp", slog.String("error", errors[0]))
		util.WriteError(w, http.StatusBadRequest, errors[0])
		return
	}

	// create
	user, err := h.userService.CreateUser(ctx, &payload)
	if err != nil {
		slog.Error("SignUp", slog.Any("error", err))
		util.WriteErrorf(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteJson(w, http.StatusOK, user)
}

func (h *ApiHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIDHeader := r.Header.Get("user_id")

	// convert user id to int64
	userID, err := strconv.ParseInt(userIDHeader, 10, 64)
	if err != nil {
		slog.Error("CreateOrder", slog.Any("convert error", err))
		util.WriteErrorf(w, http.StatusBadRequest, err)
		return
	}

	// get payload
	var payload request.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		slog.Error("CreateOrder", slog.Any("decoding error", err))
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
		slog.Error("SignUp", slog.Any("create error", err))
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
