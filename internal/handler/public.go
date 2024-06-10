package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"gotu/bookstore/internal/request"
	"gotu/bookstore/internal/service"
	"gotu/bookstore/internal/util"
)

type ApiHandler struct {
	bookService *service.BookService
	userService *service.UserService
}

func NewApiHandler(sv1 *service.BookService, sv2 *service.UserService) *ApiHandler {
	return &ApiHandler{
		bookService: sv1,
		userService: sv2,
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
	payload := &request.SignUpRequest{}
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		util.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	// validate
	if errors := util.ValidateRequest(payload); len(errors) > 0 {
		slog.Error("SignUp", slog.String("error", errors[0]))
		util.WriteError(w, http.StatusBadRequest, errors[0])
		return
	}

	// create
	user, err := h.userService.Create(ctx, payload)
	if err != nil {
		slog.Error("SignUp", slog.Any("error", err))
		util.WriteErrorf(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteJson(w, http.StatusOK, user)
}
