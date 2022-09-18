FROM golang:1.19
RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y dumb-init
RUN mkdir /home/gameshelf
RUN useradd -u 8877 gameshelf
WORKDIR /app/
COPY ./ ./
RUN go build  --tags="fts5"  .
EXPOSE 80
USER gameshelf
ENTRYPOINT  ["/usr/bin/dumb-init", "--"]
CMD ./gameshelf
