package main

import (
	"fmt"
	"net/http"

	"streaming/api"
	"streaming/services"
	"streaming/storage"
)

func main() {
	// Base de datos en memoria
	db := storage.NewMemoryStorage()

	// Servicios (lógica del sistema)
	userService := services.NewUserService(db)
	catalogService := services.NewCatalogService(db)

	// Router (API REST)
	router := api.NewRouter(userService, catalogService)

	fmt.Println("✅ API corriendo en: http://localhost:8080")
	fmt.Println("✅ Health check:  http://localhost:8080/health")

	// Iniciar servidor
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("❌ Error al iniciar servidor:", err)
	}
=======
go
package main
func main () {
     // Punto de entrada del sistema
}