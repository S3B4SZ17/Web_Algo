# syntax=docker/dockerfile:1

############################
# STEP 1 build executable binary
############################
# Alpine image is chosen for its small footprint
FROM golang:1.18-alpine as builder

# Install SSL ca certificates.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache ca-certificates tzdata && update-ca-certificates

# Create appuser.
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR $GOPATH/src/web_algo/app
# Copy all the source code
COPY . .

# Build the binary
# Removing debug informations and compile only for linux target and disabling cross compilation.
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/web_algo

############################
# STEP 2 build a small image
############################
FROM scratch

# Import the user and group files from the builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group


# Copy our static executable.
COPY --from=builder /go/bin/web_algo /go/bin/web_algo

# Use an unprivileged user.
USER appuser:appuser

# Run the binary
ENTRYPOINT [ "/go/bin/web_algo" ]