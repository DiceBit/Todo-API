FROM golang:1.23

LABEL stage=gobuilder

ENV GOOS=linux
ENV TZ="Europe/Minsk"
ENV APP_ROOT="/app"

WORKDIR /app


ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

WORKDIR /app/cmd/app
RUN go build -o app

EXPOSE 8080

CMD ["./app"]


