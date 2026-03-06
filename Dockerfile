# syntax=docker/dockerfile:1

FROM golang:1.25-alpine AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/app ./cmd/main.go
COPY readme.md /readme.md


FROM scratch
COPY --from=builder /out/app /app

EXPOSE 8080
ENTRYPOINT ["/app"]