FROM golang:1.19.5 as builder
ARG ARCH=amd64
WORKDIR ./
COPY go.* .
RUN go mod download
COPY . .
RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates
ENTRYPOINT ["./server.go"]