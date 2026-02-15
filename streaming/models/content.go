package models

// Content representa una película o serie
type Content struct {
	id    int
	title string
	genre string
}

// Constructor
func NewContent(id int, title string, genre string) Content {
	return Content{id: id, title: title, genre: genre}
}

// Getters (encapsulación)
func (c Content) GetID() int {
	return c.id
}

func (c Content) GetTitle() string {
	return c.title
}

func (c Content) GetGenre() string {
	return c.genre
}
