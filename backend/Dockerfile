# Use official Go image as base
FROM golang:1.24

# Set working directory
WORKDIR /app

# Copy go.mod only (since go.sum doesn't exist yet)
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy all source code
COPY . .

# Build or run directly (depends on your approach)
CMD ["go", "run", "/app/cmd/daily-helper/main.go"]

# Or for production, build binary and run that
