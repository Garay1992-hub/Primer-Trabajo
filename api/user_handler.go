package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"streaming/models"
	"streaming/services"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// RegisterRoutes registra las rutas del módulo User en el mux
func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {

	// /users -> GET (listar) y POST (crear)
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetUsers(w, r)
		case http.MethodPost:
			h.CreateUser(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// /users/{id} -> GET (ver por id) y DELETE (eliminar)
	mux.HandleFunc("/users/", h.UserByID)
}

// GET /users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.Service.GetUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// POST /users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	h.Service.AddUser(user.ID, user.Name, user.Email)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Usuario creado correctamente",
	})
}

// Handler para /users/{id}
func (h *UserHandler) UserByID(w http.ResponseWriter, r *http.Request) {

	// Extraer el id desde /users/{id}
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodGet:
		user, ok := h.Service.GetUserByID(id)
		if !ok {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

	case http.MethodDelete:
		deleted := h.Service.DeleteUser(id)
		if !deleted {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Usuario eliminado correctamente",
		})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}