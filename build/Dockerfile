FROM golang:1.16.3 AS base

ENV GOPRIVATE=github.com/blinkops
WORKDIR /go/src/github.com/blinkops/blink-plugin-base

COPY go.mod go.sum ./
RUN go mod download
COPY . .

##################
FROM base AS builder

# Install the package
RUN go install ./cmd/intpack

##################
FROM base AS test

RUN mkdir -p /reports && \
    go get -u github.com/jstemmer/go-junit-report && \
    go get github.com/axw/gocov/... && \
    go get github.com/AlekSi/gocov-xml && \
    go mod tidy

#  Run Tests
RUN go test -v -cover -covermode=count -coverprofile=profile.cov ... | \
    go-junit-report > /reports/test_report.xml && \
    go tool cover -func profile.cov && \
    gocov convert profile.cov | gocov-xml > /reports/controller_coverage.xml

##################
FROM scratch AS reports

COPY --from=test /reports/controller_coverage.xml .
COPY --from=test /reports/test_report.xml .

##################
FROM ubuntu AS plugin

WORKDIR /blink-plugin-base

# Copy only the required artifacts.
COPY --from=builder /go/bin/intpack .
COPY plugin/config.yaml config.yaml

# Expose the gRPC port.
EXPOSE 1337

# Set the default entrypoint (Should not be changed in inheriting layers).
ENTRYPOINT ./intpack
