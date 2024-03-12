FROM golang:1.21.4 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src/the-wall src/the-wall
COPY pkg pkg

RUN ls
RUN go build -o the-wall ./src/the-wall/cmd/the-wall

EXPOSE 8080

CMD ["./the-wall"]
