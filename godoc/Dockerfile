FROM golang:1.12.0-alpine3.9

WORKDIR /app

COPY . /app

RUN go mod init godoc

CMD [ "go","run","." ]

