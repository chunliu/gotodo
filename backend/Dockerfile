# Build with golang image
FROM golang:latest as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gotodo

# Create image
FROM alpine:latest 

RUN apk --no-cache add ca-certificates
WORKDIR /gotodoapp
COPY --from=builder /app/gotodo ./
COPY --from=builder /app/pages ./pages
COPY --from=builder /app/static ./static
RUN chmod +x ./gotodo \
    && chmod +r ./pages/index.html \
    && chmod +r ./static/gotodo.js \
    && chmod +r ./static/style.css

EXPOSE 8080

CMD ["./gotodo"]