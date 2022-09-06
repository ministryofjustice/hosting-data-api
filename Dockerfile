FROM golang:1.19.0-alpine3.15

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

RUN mkdir -p /opt/your-folder && \
    adduser --disabled-password nonRootUser -u 1001 && \
    chown -R nonRootUser:nonRootUser /opt/your-folder

USER 1001

CMD ["app"]
