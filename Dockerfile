FROM golang:1.19
RUN mkdir /home/gameshelf
RUN useradd -u 8877 gameshelf
WORKDIR /app/
COPY ./ ./
RUN go build  --tags="fts5"  .
EXPOSE 80
USER gameshelf
CMD ./gameshelf
