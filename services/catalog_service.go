package services

import (
	"fmt"
	"streaming/models"
)

// CatalogService (capa de negocio)
type CatalogService struct {
	db ContentRepository // 👈 interfaz
}

func NewCatalogService(db ContentRepository) *CatalogService {
	return &CatalogService{db: db}
}

func (s *CatalogService) AddContent(id int, title string, genre string) {
	content := models.NewContent(id, title, genre)
	s.db.AddContent(content)
	fmt.Println("✅ Contenido agregado correctamente.")
}

func (s *CatalogService) ShowCatalog() {
	catalog := s.db.GetContents()

	if len(catalog) == 0 {
		fmt.Println("⚠️ El catálogo está vacío.")
		return
	}

	fmt.Println("\n🎞️ CATÁLOGO DISPONIBLE")
	for _, c := range catalog {
		fmt.Printf("ID: %d | Título: %s | Género: %s\n", c.GetID(), c.GetTitle(), c.GetGenre())
	}
}
