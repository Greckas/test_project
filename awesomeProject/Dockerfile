FROM golang:1.12.6 AS builder

# Installing godep
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN mkdir -p /go/src/github.com/Greckas/test_project/awesomeProject
WORKDIR /go/src/github.com/Greckas/test_project/awesomeProject

# copies the Gopkg.toml and Gopkg.lock to WORKDIR
COPY ./ ./

# install the dependencies without checking for go code
RUN dep ensure -vendor-only

# Build my app
RUN go build -o /medievaltimes/server .
CMD ["/Greckas/server"]
