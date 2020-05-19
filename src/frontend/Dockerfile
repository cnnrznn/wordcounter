# Use base golang image from Docker Hub
FROM golang:1.14

WORKDIR /src/frontend

# Copy the go.mod and go.sum, download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the application source code
COPY . ./

# Compile the application to /app/frontend.
RUN go build -o /app/frontend .

ENTRYPOINT ["/app/frontend"]
