# ========= BUILD STAGE =========
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ordent-backend ./main.go


FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/ordent-backend /app/ordent-backend

COPY config /app/config

RUN chmod +x /app/ordent-backend

EXPOSE 8081

CMD ["/app/ordent-backend"]
