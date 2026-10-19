# ==========================================
# STAGE 1: Build the Go 1.27.2 Binary
# ==========================================
FROM golang:1.27.2-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy dependency files first to leverage Docker caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build a statically linked binary for Linux
# CGO_ENABLED=0 ensures the binary doesn't depend on system C libraries
RUN CGO_ENABLED=0 GOOS=linux go build -o /my-go-app .

# ==========================================
# STAGE 2: Final Minimal Run Image
# ==========================================
FROM alpine:3.19

# Add a non-root user for security purposes
RUN adduser -D appuser
USER appuser

# Set working directory
WORKDIR /

# Copy only the compiled binary from the builder stage
COPY --from=builder /ChatAppGo /ChatAppGo

# Document the port your app listens on (e.g., 8080)
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/ChatAppGo"]
