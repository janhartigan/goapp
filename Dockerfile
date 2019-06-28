FROM golang:1.12

WORKDIR $GOPATH/src/app

COPY . .
COPY docker/start.sh /bin/original_start.sh

ENV GO111MODULE=on

RUN go build -o /bin/app

# Set up start script
RUN tr -d '\r' < /bin/original_start.sh > /bin/start.sh && \
    chmod -R 700 /bin/start.sh

EXPOSE 8080

ENTRYPOINT ["/bin/start.sh"]