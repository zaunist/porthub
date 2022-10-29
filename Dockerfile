# Build the manager binary
FROM golang:1.19 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source and configuration files
COPY main.go main.go
COPY cmd/ cmd/
COPY pkg/ pkg/

# Build
RUN CGO_ENABLED=0 go build -a -o porthub main.go

# Use alpine as base images
FROM alpine:3.16
WORKDIR /
COPY --from=builder /workspace/porthub .

EXPOSE 38210
EXPOSE 38211
EXPOSE 38212

ENTRYPOINT ["/porthub"]
