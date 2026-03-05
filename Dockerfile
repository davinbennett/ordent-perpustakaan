# ========= BUILD STAGE =========
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN rm -f ordent-perpustakaan-backend

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app-bin ./main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app-bin /app/ordent-perpustakaan-backend
COPY config /app/config

RUN chmod +x /app/ordent-perpustakaan-backend

EXPOSE 8080

CMD ["/app/ordent-perpustakaan-backend"]