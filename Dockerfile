FROM golang:1.18.0-alpine3.15 AS build
WORKDIR /todo-app
COPY . .
RUN go build -o todo-api
CMD [ "./todo-api" ]
EXPOSE 8080