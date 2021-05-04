FROM golang AS builder

ENV GOPRIVATE=github.com/blinkops

WORKDIR /go/src/github.com/blinkops/shell-script-runner/plugin

COPY go.mod go.sum .
RUN go mod download
COPY . .

# Install the package
RUN go install -v ./...

FROM ubuntu AS plugin

WORKDIR /shell-script-runner

# Copy only the required artifacts.
COPY --from=builder /go/bin/plugin .
COPY plugin/config.yaml config.yaml

# Expose the gRPC port.
EXPOSE 1337

# Set the default entrypoint (Should not be changed in inheriting layers).
ENTRYPOINT ./plugin
