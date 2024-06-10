package handler

import (
	"log/slog"
	"net/http"

	"gotu/bookstore/internal/book"
	"gotu/bookstore/internal/util"
)

type ApiHandler struct {
	service *book.BookService
}

func NewApiHandler(sv *book.BookService) *ApiHandler {
	return &ApiHandler{
		service: sv,
	}
}

func (h *ApiHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get books
	books, err := h.service.List(ctx)
	if err != nil {
		slog.Error("ListBooks", slog.Any("error", err))
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteJson(w, http.StatusOK, books)
}
