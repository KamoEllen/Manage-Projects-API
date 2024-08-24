# Stage 1: Build
FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go build -o api .

# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/api .
EXPOSE 8000
CMD ["./api"]
