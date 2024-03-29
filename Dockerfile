FROM golang:1.22-alpine
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
COPY go.mod go.sum ./
RUN go mod download
WORKDIR /app/cmd/server
CMD ["air", "main.go", "wire_gen.go"]
