FROM golang:1.12-alpine

# Install packages
RUN echo 'http://dl-cdn.alpinelinux.org/alpine/edge/testing' >> /etc/apk/repositories && \
  apk update && \
  apk upgrade && \
  apk --no-cache add \
    supervisor \
    git

WORKDIR $GOPATH/src/app

COPY . .
COPY docker/start.sh /bin/original_start.sh
COPY docker/conf/supervisord.conf /etc/supervisor.d/supervisord.ini

# Turn on Go 1.11 Modules and build
ENV GO111MODULE=on
RUN go build -o /bin/app

# Set up start script
RUN tr -d '\r' < /bin/original_start.sh > /bin/start.sh && \
    chmod -R 700 /bin/start.sh

EXPOSE 8080

ENTRYPOINT ["/bin/start.sh"]