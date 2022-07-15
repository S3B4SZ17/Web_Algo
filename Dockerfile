# syntax=docker/dockerfile:1

# Alpine image is chosen for its small footprint
FROM golang:1.18-alpine as builder

WORKDIR /web_algo

# Copy all the source code
COPY . .

RUN go mod tidy

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/web_algo

FROM scratch

COPY --from=builder /bin/web_algo /usr/bin/
# Run the binary
CMD [ "web_algo" ]