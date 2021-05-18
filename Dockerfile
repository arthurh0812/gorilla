FROM golang:1.16.3-alpine

ENV PORT=8080
# Go Modules
ENV GO111MODULE=on
ENV GOMOD=/app/go.mod
ENV GOPATH=/go
ENV GOMODCACHE=/go/pkg/mod
# Go Env
ENV GOPRIVATE=github.com/arthurh0812

RUN mkdir /go/pkg/mod

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download && go mod vendor && go mod verify

COPY . .

RUN go build -o /cmd cmd

EXPOSE ${PORT}

CMD ["cmd"]