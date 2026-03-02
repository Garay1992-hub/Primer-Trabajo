package services

import "streaming/models"

type CatalogService struct {
	repo ContentRepository
}

func NewCatalogService(repo ContentRepository) *CatalogService {
	return &CatalogService{repo: repo}
}

func (c *CatalogService) AddContent(id int, title string, genre string) {
	content := models.Content{ID: id, Title: title, Genre: genre}
	c.repo.AddContent(content)
}

func (c *CatalogService) GetContents() []models.Content {
	return c.repo.GetContents()
}

func (c *CatalogService) GetContentByID(id int) (models.Content, bool) {
	return c.repo.GetContentByID(id)
}

func (c *CatalogService) DeleteContent(id int) bool {
	return c.repo.DeleteContent(id)
}