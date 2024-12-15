# Usa la imagen oficial de Go
FROM golang:1.22

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo go.mod y go.sum, e instala las dependencias
COPY api/go.mod api/go.sum ./
# COPY go.mod go.sum ./
RUN go mod download

# Copia el código de la API al contenedor
COPY api .

# Compila la aplicación
RUN go build -o main .

# Expone el puerto de la aplicación
EXPOSE 3000

# Define el comando para ejecutar la aplicación
CMD ["./main"]
