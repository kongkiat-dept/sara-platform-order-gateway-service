FROM golang:1.16.3-alpine3.13 as builder

WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.* ./
RUN go mod tidy
# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -v -o service cmd/api/main.go

# multi-stage ...
FROM alpine:3.14.0 as release
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/service /app/cmd/
COPY --from=builder /app/configs /app/configs
WORKDIR /app

# Run the web server on container startup.
#CMD ./cmd/server
ENTRYPOINT ["./cmd/service", "--env=develop"]