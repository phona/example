FROM golang:latest AS builder
WORKDIR /app
COPY app.go ./
COPY go.mod ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM rockylinux:latest
WORKDIR /root/
COPY --from=builder /app/app ./
CMD ["./app"]
