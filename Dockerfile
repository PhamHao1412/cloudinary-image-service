FROM golang:1.24-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /image-service ./cmd

# ---- Runtime stage ----
FROM gcr.io/distroless/base-debian12:nonroot
WORKDIR /app
COPY --from=build /image-service /app/image-service
COPY .env /app/.env
EXPOSE 8085
USER nonroot
ENTRYPOINT ["/app/image-service"]
