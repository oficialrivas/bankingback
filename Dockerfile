# Usa la imagen oficial de Go como imagen base
FROM golang:1.22 as builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente de tu aplicación al contenedor
COPY . .

# Descarga las dependencias de tu aplicación
RUN go mod download

# Compila la aplicación para producción
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Usa la imagen de scratch para una imagen final pequeña
FROM scratch

# Copia el ejecutable a la imagen final
COPY --from=builder /app/main .
# Expone el puerto 8080
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
