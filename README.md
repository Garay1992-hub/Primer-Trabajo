Proyecto Final – Servicios Web en Go
Sistema de Gestión tipo Streaming
Estudiante: José Luis
Asignatura: Programación Estructurada y Funcional
Fecha: 03 de Marzo 2026
Repositorio: Primer-Trabajo

🎯 Objetivo del Proyecto
Desarrollar un sistema de gestión tipo plataforma de streaming utilizando el lenguaje Go, aplicando arquitectura por capas, principios de programación estructurada y funcional, y generación de servicios web REST con serialización en formato JSON.

🏗 Arquitectura Implementada
El proyecto está organizado en capas:
•	models/ → Estructuras de datos (User, Content)
•	services/ → Lógica de negocio
•	storage/ → Almacenamiento en memoria
•	api/ → Handlers y rutas HTTP
•	main.go → Punto de entrada del servidor
Esta estructura permite separación de responsabilidades y mayor mantenibilidad.

🚀 Funcionalidades Implementadas
El sistema incluye los siguientes servicios web:
Usuarios
•	Crear usuario (POST /users)
•	Listar usuarios (GET /users)
•	Obtener usuario por ID (GET /users/{id})
•	Eliminar usuario (DELETE /users/{id})
Contenidos
•	Crear contenido (POST /contents)
•	Listar contenidos (GET /contents)
•	Obtener contenido por ID (GET /contents/{id})
•	Eliminar contenido (DELETE /contents/{id})
Sistema
•	Verificación de estado (GET /health)
Se implementaron más de 8 endpoints distintos, cumpliendo el requisito de la asignatura.

🔄 Serialización
La comunicación entre cliente y servidor se realiza en formato JSON, utilizando el paquete estándar encoding/json del lenguaje Go.

📚 Tecnologías Utilizadas
•	Lenguaje: Go
•	Protocolo: HTTP
•	Arquitectura: REST
•	Almacenamiento: Memoria (in-memory storage)
•	Control de versiones: Git y GitHub

🧠 Conocimientos Integrados
Este proyecto integra:
•	Programación estructurada
•	Programación funcional
•	Encapsulación
•	Interfaces
•	Manejo de errores
•	Arquitectura por capas
•	Servicios Web REST
•	Serialización JSON
