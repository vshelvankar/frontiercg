FROM golang:1.12.0-alpine

WORKDIR /go-modules

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o frontiercg 

EXPOSE 8080
CMD ["./frontiercg"]
