FROM golang:1.22.5-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

#RUN go build -o main .

EXPOSE 8080

#CMD ["./main"]
#CMD ["go", "run", "main.go"]
CMD ["air", "-c", ".air.toml"]