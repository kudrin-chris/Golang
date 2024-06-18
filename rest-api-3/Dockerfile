FROM golang:1.22.4-alpine as build
WORKDIR /build
COPY go.mod . 
RUN go mod download
COPY . .
RUN go build -o build/main main.go storage.go handler.go
CMD ["build/main"]