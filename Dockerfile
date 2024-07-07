FROM golang:1.22.4
WORKDIR /app
COPY . .
RUN go install github.com/air-verse/air@latest
EXPOSE 8080
CMD ["air"]