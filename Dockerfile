FROM golang:1.23-alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o api ./cmd/api

FROM alpine:latest
WORKDIR /app

COPY --from=build /app/api .
COPY .env .

EXPOSE 8080
CMD ["./api"]
