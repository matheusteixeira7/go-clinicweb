# Etapa de construção
FROM golang:1.22-alpine as base

# Define o diretório de trabalho
WORKDIR /app

# Copia todos os arquivos do diretório atual para o diretório de trabalho no container
COPY . .

# Instala o Air para hot reload
RUN go install github.com/cosmtrek/air@latest
RUN go mod download

# Define o diretório de trabalho para /usr/src/cmd/server
WORKDIR /app/cmd/server

# Define o comando para executar o aplicativo
CMD ["air", "main.go", "wire_gen.go"]
