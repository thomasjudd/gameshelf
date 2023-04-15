FROM golang:1.19
RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y dumb-init
RUN mkdir /home/gameshelf
#RUN useradd -u 8877 gameshelf
WORKDIR /app/
COPY ./controller ./controller
COPY ./entity ./entity
COPY ./main.go ./main.go
COPY ./go.sum ./go.sum
COPY ./go.mod ./go.mod
RUN go build  .
COPY ./templates ./templates
COPY ./static ./static
#USER gameshelf
ENTRYPOINT  ["/usr/bin/dumb-init", "--"]
CMD ./gameshelf
