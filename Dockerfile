# Etapa 1: build
FROM golang:1.24.0-alpine AS builder

WORKDIR /app

# Copia os arquivos do projeto
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila o binário de forma estática
RUN go build -o main -ldflags="-s -w"

# Etapa 2: imagem final
FROM alpine:latest

WORKDIR /app

# Copia apenas o binário do build
COPY --from=builder /app/main .

# Expõe a porta da aplicação
EXPOSE 8080

# Executa o binário
CMD ["./main"]
