FROM golang:1.21.4 AS build

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./

RUN go build -o the-wall ./src/the-wall/cmd/the-wall

FROM debian:12.5

WORKDIR /app

COPY --from=build /app/the-wall .

EXPOSE 8080

RUN addgroup --system nonroot \
    && adduser --system nonroot \
    && adduser nonroot nonroot

USER nonroot

CMD ["./the-wall"]
