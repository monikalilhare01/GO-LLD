FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /survey-backend ./

EXPOSE 8080
ENV APP_PORT=8080

ENTRYPOINT ["/survey-backend"]
