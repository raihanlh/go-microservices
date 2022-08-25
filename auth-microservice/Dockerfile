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
RUN go build -o /auth-microservice
RUN chmod a+x /auth-microservice

EXPOSE 3001
CMD ["/auth-microservice"]

