FROM golang:1.19.5
ARG ARCH=amd64
WORKDIR /server
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/app ./server.go
ENTRYPOINT ["app"]
