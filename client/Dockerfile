FROM golang:latest 
ADD . /go/src/client
WORKDIR /go/src/client
RUN go get ./
RUN go install /go/src/client
ENTRYPOINT /go/bin/client

# docker build --tag=chuck .

# docker run -it --name=ch chuck

# docker run -it -d --name=ch chucky
# docker attach ch
# ctrl+p, ctrl+q to exit