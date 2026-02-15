package storage

import "streaming/models"

// MemoryStorage simula una base de datos en memoria
type MemoryStorage struct {
	users    []models.User
	contents []models.Content
}

// Constructor
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users:    []models.User{},
		contents: []models.Content{},
	}
}

/* =========================
   USERS (cumple UserRepository)
   ========================= */

func (db *MemoryStorage) AddUser(user models.User) {
	db.users = append(db.users, user)
}

func (db *MemoryStorage) GetUsers() []models.User {
	return db.users
}

/* =========================
   CONTENT (cumple ContentRepository)
   ========================= */

func (db *MemoryStorage) AddContent(content models.Content) {
	db.contents = append(db.contents, content)
}

func (db *MemoryStorage) GetContents() []models.Content {
	return db.contents
}
