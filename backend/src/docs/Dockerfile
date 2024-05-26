FROM golang:1.21.4 AS build

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/src/docs src/docs
COPY backend/pkg pkg

RUN go build -o docs ./src/docs

FROM debian:12.5

WORKDIR /app

COPY --from=build /app/docs .
COPY backend/src/docs/openapi.json .

EXPOSE 8080

RUN addgroup --system nonroot \
    && adduser --system nonroot \
    && adduser nonroot nonroot

USER nonroot

CMD ["./docs"]
