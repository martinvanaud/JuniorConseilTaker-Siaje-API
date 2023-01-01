FROM golang:1.18-alpine

WORKDIR /siaje-api/app/golang/

COPY . .

RUN go mod tidy

EXPOSE 8018

CMD go run .