package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func RegisterHandlers() *chi.Mux {
	chi := chi.NewRouter()

	chi.Get("/status", StatusHandler)

	return chi
}
