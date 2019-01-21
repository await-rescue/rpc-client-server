FROM golang:latest 
ADD . /go/src/server
WORKDIR /go/src/server
RUN go get ./
RUN go install /go/src/server
ENTRYPOINT /go/bin/server

# docker build --tag=chuck .

# docker run -it --name=ch chuck

# docker run -it -d --name=ch chucky
# docker attach ch
# ctrl+p, ctrl+q to exit