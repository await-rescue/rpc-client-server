FROM golang:latest 
ADD . /go/src/client
WORKDIR /go/src/client
RUN go get ./
RUN go install /go/src/client
ENTRYPOINT /go/bin/client