FROM golang:1.25

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8083

CMD ["go", "run", "cmd/app/main.go"]