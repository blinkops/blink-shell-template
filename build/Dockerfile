FROM golang AS builder

ENV GOPRIVATE=github.com/blinkops

WORKDIR /go/src/github.com/blinkops/blink-plugin-base

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Install the package
RUN go install ./cmd/intpack

FROM ubuntu AS plugin

WORKDIR /blink-plugin-base

# Copy only the required artifacts.
COPY --from=builder /go/bin/intpack .
COPY plugin/config.yaml config.yaml

# Expose the gRPC port.
EXPOSE 1337

# Set the default entrypoint (Should not be changed in inheriting layers).
ENTRYPOINT ./intpack
