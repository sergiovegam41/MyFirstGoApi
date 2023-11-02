# Utiliza la imagen oficial de Go como base
FROM golang:1.17-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al directorio de trabajo
COPY . .

# Compila el proyecto
RUN go build -o main .

# Exponer el puerto en el que se ejecuta la aplicación
EXPOSE 8080

# Ejecutar la aplicación cuando se inicie el contenedor
CMD ["./main"]