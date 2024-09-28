FROM golang:1.22.0-alpine

#untuk set destination for copy
WORKDIR /app

#download go modules
COPY go.mod ./
RUN go mod download

#copy resource code
COPY *.go ./

#build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./go-hello-world

#optional: expose port
EXPOSE 8080

#run
CMD [ "./go-hello-world" ]