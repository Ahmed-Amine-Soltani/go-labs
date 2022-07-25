FROM golang:alpine AS builder

# Install git and ca-certificates (needed to be able to call HTTPS)
RUN apk --update add ca-certificates git

# Move to working directory /app
WORKDIR /app

# Copy the code into the container
COPY . .

# Download dependencies using go mod
RUN go mod download

# Build the application's binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# Build a smaller image that will only contain the application's binary
FROM scratch

# Move to working directory /app
WORKDIR /app

# Copy certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copy application's binary
COPY --from=builder /app .

# Command to run the application when starting the container
CMD ["./main"]
