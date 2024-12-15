# Proyect-Go-JWT-Gorm

**Luca Pierini**

---

## Descripción

**Proyect-Go-JWT-Gorm** es un sistema de autenticación desarrollado en Go que utiliza Gin como framework web y Gorm para la integración con PostgreSQL. Implementa autenticación mediante JSON Web Tokens (JWT) y está diseñado para ser fácilmente desplegable con Docker y Docker Compose, lo que permite su uso tanto en entornos locales como en producción.

## Funcionalidades Principales

- **Registro de Usuarios**: Permite registrar nuevos usuarios con contraseñas hasheadas para mayor seguridad.
- **Inicio de Sesión**: Valida las credenciales del usuario y genera un token de acceso JWT.
- **Protección de Endpoints**: Middleware para validar y proteger rutas privadas con JWT.
- **Persistencia de Datos**: Almacena información de usuarios en una base de datos PostgreSQL.

## Endpoints

### Usuarios

- **POST /signup**: Registro de nuevos usuarios.
- **POST /login**: Inicio de sesión y generación de tokens.
- **POST /validate**: Endpoint protegido para validar la sesión activa.

### Ping
- **GET /ping**: Endpoint de prueba para verificar la conectividad del servicio.

## Tecnologías Utilizadas

- **Backend**: Go (Gin Framework)
- **Base de Datos**: PostgreSQL con Gorm
- **Autenticación**: JWT
- **Contenedores**: Docker y Docker Compose

## Estructura del Proyecto

- **controllers**: Controladores para manejar la lógica de usuario.
- **initializers**: Configuración inicial de la base de datos y variables de entorno.
- **middleware**: Middleware para validación de autenticación.
- **models**: Definición de modelos para la base de datos.
- **main.go**: Punto de entrada de la aplicación.

## Requisitos Técnicos

1. **Docker**: Asegúrate de tener instalado Docker y Docker Compose.
2. **Variables de Entorno**:
   - **PORT**: Puerto donde correrá la API.
   - **SECRET**: Clave secreta para firmar tokens JWT.
   - **DB**: URL de conexión a PostgreSQL.

## Instrucciones de Instalación

1. Clona este repositorio:

   ```bash
   git clone https://github.com/lucapierini/proyect-go-jwt.git
   cd proyect-go-jwt
   ```

2. Configura las variables de entorno creando un archivo `.env` en la raíz del proyecto:

   ```env
   PORT=3000
   SECRET=secret_key
   DB=postgres://userTest:123456@localhost:5432/gorm-jwt-auth?sslmode=disable
   ```

3. Levanta los contenedores utilizando Docker Compose:

   ```bash
   docker-compose up --build
   ```

4. La API estará disponible en `http://localhost:3000`.

## Uso de la API

1. **Registro de Usuario**:
   - Endpoint: `POST /signup`
   - Body:

     ```json
     {
       "Email": "test@example.com",
       "Password": "securepassword123"
     }
     ```

2. **Inicio de Sesión**:
   - Endpoint: `POST /login`
   - Body:

     ```json
     {
       "Email": "test@example.com",
       "Password": "securepassword123"
     }
     ```

3. **Validación de Usuario**:
   - Endpoint: `POST /validate`
   - Header:
     - Cookie: Authorization=<JWT_TOKEN>

## Instrucciones de Desarrollo

1. Instala las dependencias locales:

   ```bash
   go mod tidy
   ```

2. Ejecuta el servidor:

   ```bash
   go run main.go
   ```

3. Para levantar una base de datos PostgreSQL de manera local, utiliza Docker:

   ```bash
   docker run -d --name postgres-db -e POSTGRES_USER=userTest -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=gorm-jwt-auth -p 5432:5432 postgres:15
   ```

