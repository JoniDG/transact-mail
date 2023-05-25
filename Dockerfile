# Imagen base que contiene el entorno de ejecución de Go
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos necesarios para compilar y ejecutar la aplicación
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copia el código fuente de tu aplicación en el contenedor
COPY . ./

# Compila la aplicación
RUN go build -o bin ./cmd/main.go

# Expone el puerto en el que tu aplicación escucha las conexiones
EXPOSE 8080

# Comando para ejecutar tu aplicación cuando se inicie el contenedor
CMD ["./bin"]
