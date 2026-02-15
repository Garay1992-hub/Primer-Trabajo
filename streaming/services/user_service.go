package services

import (
	"fmt"
	"streaming/models"
)

// UserService (capa de negocio)
type UserService struct {
	db UserRepository // 👈 interfaz
}

func NewUserService(db UserRepository) *UserService {
	return &UserService{db: db}
}

func (s *UserService) AddUser(id int, name string, email string) {
	user := models.NewUser(id, name, email)
	s.db.AddUser(user)
	fmt.Println("✅ Usuario registrado correctamente.")
}

func (s *UserService) ShowUsers() {
	users := s.db.GetUsers()

	if len(users) == 0 {
		fmt.Println("⚠️ No hay usuarios registrados.")
		return
	}

	fmt.Println("\n👥 USUARIOS REGISTRADOS")
	for _, u := range users {
		fmt.Printf("ID: %d | Nombre: %s | Email: %s\n", u.GetID(), u.GetName(), u.GetEmail())
	}
}
