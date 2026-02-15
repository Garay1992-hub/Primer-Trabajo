package services

import "streaming/models"

// UserRepository define lo que el sistema necesita para manejar usuarios
type UserRepository interface {
	AddUser(user models.User)
	GetUsers() []models.User
}

// ContentRepository define lo necesario para manejar el catálogo
type ContentRepository interface {
	AddContent(content models.Content)
	GetContents() []models.Content
}
