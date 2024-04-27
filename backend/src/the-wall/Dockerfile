FROM golang:1.21.4 AS build

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/src/the-wall src/the-wall
COPY backend/pkg pkg

RUN go build -o the-wall ./src/the-wall/cmd/the-wall

FROM debian:12.5

WORKDIR /app

COPY --from=build /app/the-wall .

EXPOSE 8080

CMD ["./the-wall"]
