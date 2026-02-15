package models

// User representa un usuario del sistema
type User struct {
	id    int
	name  string
	email string
}

// Constructor
func NewUser(id int, name string, email string) User {
	return User{id: id, name: name, email: email}
}

// Getters (encapsulación)
func (u User) GetID() int {
	return u.id
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetEmail() string {
	return u.email
}
