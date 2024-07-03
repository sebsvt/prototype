FROM golang:1.22.4

WORKDIR /backend

RUN go install github.com/air-verse/air@latest


COPY . .
RUN go mod tidy
