FROM golang:alpine

RUN apk update && apk add git

WORKDIR /app

RUN ls

COPY . .

RUN go mod tidy
RUN go build -o hello-world

ENTRYPOINT [ "./hello-world" ]