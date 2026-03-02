package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"streaming/models"
	"streaming/services"
)

type ContentHandler struct {
	Service *services.CatalogService
}

func NewContentHandler(service *services.CatalogService) *ContentHandler {
	return &ContentHandler{Service: service}
}

func (h *ContentHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/contents", h.contents)
	mux.HandleFunc("/contents/", h.contentByID)
}

func (h *ContentHandler) contents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		contents := h.Service.GetContents()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contents)

	case http.MethodPost:
		var content models.Content
		if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		h.Service.AddContent(content.ID, content.Title, content.Genre)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Contenido creado correctamente",
		})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *ContentHandler) contentByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/contents/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		content, ok := h.Service.GetContentByID(id)
		if !ok {
			http.Error(w, "Contenido no encontrado", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(content)

	case http.MethodDelete:
		deleted := h.Service.DeleteContent(id)
		if !deleted {
			http.Error(w, "Contenido no encontrado", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Contenido eliminado correctamente",
		})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}