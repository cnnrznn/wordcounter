# Use base golang image from Docker Hub
FROM golang:1.14

WORKDIR /src/backend

# Copy the go.mod and go.sum, download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the application source code
COPY . ./

# Compile the application to /app/backend.
RUN go build -o /app/backend .

ENTRYPOINT ["/app/backend"]
