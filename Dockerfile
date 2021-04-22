FROM golang AS builder

ENV GOPRIVATE=github.com/blinkops

WORKDIR /go/src/github.com/blinkops/shell-script-runner/plugin

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .


# Install the package
RUN go install -v ./...

FROM ubuntu AS plugin

WORKDIR /shell-script-runner

COPY --from=builder /go/bin/plugin .
COPY plugin.yaml plugin.yaml
COPY actions/ actions/

EXPOSE 1337

COPY plugin/config.yaml config.yaml

ENTRYPOINT ./plugin






