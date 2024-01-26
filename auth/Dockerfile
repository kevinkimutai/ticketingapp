FROM golang:1.22rc2-alpine3.19  AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go
FROM scratch
ENV DB_URL=root:P@ssw0rd@tcp(127.0.0.1:3307)/users?parseTime=true \
    PORT=8080 \
    ENV=development \
    JWT_SECRET_KEY=bMcnkA3PxJh!
COPY --from=builder /app/main ./main
CMD ["./main"] 