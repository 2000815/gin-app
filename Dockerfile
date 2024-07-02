FROM golang:1.22.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./

# Install air
RUN go install github.com/air-verse/air

COPY .air.toml ./

EXPOSE 8080

CMD ["air", "-c"]
