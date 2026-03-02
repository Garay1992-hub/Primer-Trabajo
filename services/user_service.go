package services

import "streaming/models"

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Crea usuario
func (s *UserService) AddUser(id int, name string, email string) {
	user := models.User{ID: id, Name: name, Email: email}
	s.repo.AddUser(user)
}

// Lista usuarios
func (s *UserService) GetUsers() []models.User {
	return s.repo.GetUsers()
}

// Buscar usuario por ID
func (s *UserService) GetUserByID(id int) (models.User, bool) {
	return s.repo.GetUserByID(id)
}

// Eliminar usuario por ID
func (s *UserService) DeleteUser(id int) bool {
	return s.repo.DeleteUser(id)
}