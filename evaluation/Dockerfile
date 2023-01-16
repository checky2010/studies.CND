# Use golang image with version 1.19 as base
FROM golang:1.19-buster as builder

# Create a directory and enter it
WORKDIR /app/code

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build service
COPY ./ ./
RUN go build -o /app/build

# Use an extremely small image as base for deployment
FROM gcr.io/distroless/base-debian10 AS app

# Change directory
WORKDIR /

# Copy the executable from the builder
COPY --from=builder /app/build /app

# Change user to a non-root user
USER nonroot:nonroot

# Expose port 8080 and set entrypoint
EXPOSE 8080
ENTRYPOINT ["/app"]