# use go new version
FROM golang:1.24.1

# setup folder work on contaier
WORKDIR /app

# Setup go.mod and go.sum to download dependencies
COPY  go.mod go.sum ./
RUN go mod download

# Copy source to contaier
COPY . .

# Compile the application into binary file
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/main ./