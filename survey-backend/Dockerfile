# Dockerfile
# ---- Build stage ----
# Use Go version matching your go.mod. If golang:1.24 doesn't exist on your machine, use the latest available (e.g. 1.23).
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Enable modules + faster downloads
ENV CGO_ENABLED=0 \
    GO111MODULE=on

# Cache deps first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# If your main.go lives at ./cmd/server/main.go, set that path.
# Otherwise, if main.go is at repo root, change to `./`.
RUN go build -trimpath -ldflags="-s -w" -o /out/survey-backend ./ || \
    (echo "Adjust build path in Dockerfile if your main package isn't in ./cmd/server" && false)

# ---- Runtime stage ----
# Distroless includes CA certs and runs as nonroot by default
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /
COPY --from=builder /out/survey-backend /survey-backend

# App listens on APP_PORT (8080 by default)
EXPOSE 8080

# Use environment variables for configuration
ENV APP_PORT=8080

# Non-root user from distroless image
USER nonroot:nonroot

ENTRYPOINT ["/survey-backend"]
