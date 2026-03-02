package main

import (
	"fmt"
    "bufio"
	"os"

	"streaming/services"
	"streaming/storage"
)

var reader = bufio.NewReader(os.Stdin)

func pause() {
	fmt.Print("\nPresiona ENTER para volver al menú...")
	reader.ReadString('\n')
}

func main() {

	// 1️⃣ Base de datos en memoria
	db := storage.NewMemoryStorage()

	// 2️⃣ Servicios (lógica del sistema)
	userService := services.NewUserService(db)
	catalogService := services.NewCatalogService(db)

	// 3️⃣ Ejecutar menú
	menu(userService, catalogService)
}

func menu(userService *services.UserService, catalogService *services.CatalogService) {

	var opcion int

	for {
		fmt.Println("\nSISTEMA DE GESTION DE STREAMING")
		fmt.Println("================================")
		fmt.Println("1. Registrar usuario")
		fmt.Println("2. Ver usuarios")
		fmt.Println("3. Registrar contenido")
		fmt.Println("4. Ver catalogo")
		fmt.Println("5. Salir")

		fmt.Print("Elige una opción: ")
		fmt.Scan(&opcion)
		reader.ReadString('\n') // limpia buffer

		switch opcion {

		case 1:
			var id int
			var nombre, email string

			fmt.Print("ID: ")
			fmt.Scanln(&id)

			fmt.Print("Nombre: ")
			fmt.Scanln(&nombre)

			fmt.Print("Email: ")
			fmt.Scanln(&email)

			userService.AddUser(id, nombre, email)

		case 2:
			userService.ShowUsers ()
        pause ()

		case 3:
			var id int
			var titulo, genero string

			fmt.Print("ID contenido: ")
			fmt.Scanln(&id)

			fmt.Print("Titulo: ")
			fmt.Scanln(&titulo)

			fmt.Print("Genero: ")
			fmt.Scanln(&genero)

			catalogService.AddContent(id, titulo, genero)

        case 4:
        catalogService.ShowCatalog()
        pause()

		case 5:
			fmt.Println("Saliendo del sistema...")
			return

		default:
			fmt.Println("Opción inválida")
		}
	}
}
