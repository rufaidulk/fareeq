FROM golang:1.22-alpine

WORKDIR /go 

ADD . .

RUN chmod +x start.sh

ENTRYPOINT [  "./app.sh" ]