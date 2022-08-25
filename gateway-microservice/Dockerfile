FROM golang:1.17-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY [^go.sum]* ./
COPY src src
COPY proto proto
COPY config.yml /app

WORKDIR /app/src
RUN go build -o /gateway-microservice
RUN chmod a+x /gateway-microservice

EXPOSE 3000
CMD ["/gateway-microservice"]

