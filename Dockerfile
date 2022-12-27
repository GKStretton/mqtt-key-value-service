FROM golang:1.19-alpine

WORKDIR /src

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY . .
RUN go build

ENTRYPOINT ["./mqtt-key-value-service"]