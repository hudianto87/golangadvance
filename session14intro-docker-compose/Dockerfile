#stage 1: build stage
FROM golang:1.22.0-alpine AS build

#untuk set destination for copy
WORKDIR /app

#copy resource code
COPY . .

#build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./go-hello-world .

#stage 2: run stage
FROM alpine:edge

#set working directory
WORKDIR /app

#copy binary from build stage
COPY --from=build /app/go-hello-world .

#optional: expose port
EXPOSE 8080

#run
CMD [ "./go-hello-world" ]