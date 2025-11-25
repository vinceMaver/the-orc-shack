FROM golang:1.25 AS build

WORKDIR /app
COPY . .

ENV CGO_ENABLED=1
ENV GOOS=linux

# Install certificates FIRST so HTTPS apt can work
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

# Install build deps for CGO + SQLite
RUN apt-get update && apt-get install -y \
    gcc g++ make libc6-dev sqlite3 libsqlite3-dev && \
    rm -rf /var/lib/apt/lists/*

RUN go mod download
RUN go build -o api ./cmd/api



FROM debian:bookworm-slim
WORKDIR /app

# Install certificates FIRST
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

# Fix source list (HTTPS)
RUN sed -i 's|http://deb.debian.org|https://deb.debian.org|g' /etc/apt/sources.list.d/debian.sources

# Install SQLite runtime
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-0 && \
    rm -rf /var/lib/apt/lists/*

COPY --from=build /app/api .

EXPOSE 8080
CMD ["./api"]
