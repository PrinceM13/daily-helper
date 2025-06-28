# Use official Go image as base
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all source code
COPY . .

# Build or run directly (depends on your approach)
CMD ["go", "run", "main.go"]

# Or for production, build binary and run that
