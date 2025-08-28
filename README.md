# EmbarcAPI Backend - API de ejemplo 

Este proyecto es una API sencilla en Go usando Gin Gonic, con almacenamiento en memoria y estructura profesional.

## Estructura

- `cmd/main.go`: Punto de entrada de la aplicación
- `internal/user/`: Lógica y rutas del CRUD de usuarios
- `go.mod`, `go.sum`: Dependencias

## Endpoints

- `GET /ping` — Healthcheck
- `GET /users/` — Listar usuarios
- `POST /users/` — Crear usuario (JSON: `{ "name": "Nombre" }`)
- `GET /users/:id` — Obtener usuario por ID
- `PUT /users/:id` — Actualizar usuario
- `DELETE /users/:id` — Eliminar usuario

## Cómo correr

1. Instala dependencias:
   ```sh
   go mod tidy
   ```
2. Ejecuta la API:
   ```sh
   go run ./cmd/main.go
   ```

La API estará disponible en http://localhost:8080
