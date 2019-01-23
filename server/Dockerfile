FROM golang:latest 
ADD . /go/src/server
WORKDIR /go/src/server
RUN go get ./
RUN go install /go/src/server
ENTRYPOINT /go/bin/server
