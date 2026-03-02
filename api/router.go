package api

import (
	"net/http"
	"streaming/services"
)

func NewRouter(userService *services.UserService, catalogService *services.CatalogService) *http.ServeMux {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	userHandler := NewUserHandler(userService)
	contentHandler := NewContentHandler(catalogService)

	userHandler.RegisterRoutes(mux)
	contentHandler.RegisterRoutes(mux)

	return mux
}