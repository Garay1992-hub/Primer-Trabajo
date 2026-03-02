package services

import "streaming/models"

type UserRepository interface {
	AddUser(user models.User)
	GetUsers() []models.User
	GetUserByID(id int) (models.User, bool)
	DeleteUser(id int) bool
}

type ContentRepository interface {
	AddContent(content models.Content)
	GetContents() []models.Content
	GetContentByID(id int) (models.Content, bool)
	DeleteContent(id int) bool
}