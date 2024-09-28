#stage 1: build stage
FROM golang:1.22.0-alpine AS build

#install necessary packages
RUN apk add --no-cache git make protobuf

#untuk set destination for copy
WORKDIR /app

#copy go modul
COPY go.mod go.sum ./
RUN go mod download

#copy resource code
COPY . .

#install golangci-lint
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

#run linter with increase timeout
RUN golangci-lint run --timeout 5m

#build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./go-crud-user .

#stage 2: run stage
FROM alpine:edge

#set working directory
WORKDIR /app

#copy binary from build stage
COPY --from=build /app/go-crud-user .

#install CA certificate
RUN apk --no-cache add ca-certificates

#run
CMD [ "./go-crud-user" ]