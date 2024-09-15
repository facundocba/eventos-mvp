# Usa la imagen oficial de Golang como base
FROM golang:1.21-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los manifiestos de módulos Go
COPY go.mod go.sum ./

# Descarga y cachea las dependencias
RUN go mod download

# Copia el código fuente al contenedor
COPY . .

# Construye la aplicación Go
RUN go build -o main .

# Expone el puerto en el que la aplicación escuchará
EXPOSE 8080

# Ejecuta el binario
CMD ["./main"]
