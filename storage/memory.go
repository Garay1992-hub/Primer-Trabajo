package storage

import "streaming/models"

type MemoryStorage struct {
	users    []models.User
	contents []models.Content
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users:    []models.User{},
		contents: []models.Content{},
	}
}

// ================= USERS =================

func (db *MemoryStorage) AddUser(user models.User) {
	db.users = append(db.users, user)
}

func (db *MemoryStorage) GetUsers() []models.User {
	return db.users
}

func (db *MemoryStorage) GetUserByID(id int) (models.User, bool) {
	for _, u := range db.users {
		if u.ID == id {
			return u, true
		}
	}
	return models.User{}, false
}

func (db *MemoryStorage) DeleteUser(id int) bool {
	for i, u := range db.users {
		if u.ID == id {
			db.users = append(db.users[:i], db.users[i+1:]...)
			return true
		}
	}
	return false
}

// ================= CONTENT =================

func (db *MemoryStorage) AddContent(content models.Content) {
	db.contents = append(db.contents, content)
}

func (db *MemoryStorage) GetContents() []models.Content {
	return db.contents
}

func (db *MemoryStorage) GetContentByID(id int) (models.Content, bool) {
	for _, c := range db.contents {
		if c.ID == id {
			return c, true
		}
	}
	return models.Content{}, false
}

func (db *MemoryStorage) DeleteContent(id int) bool {
	for i, c := range db.contents {
		if c.ID == id {
			db.contents = append(db.contents[:i], db.contents[i+1:]...)
			return true
		}
	}
	return false
}