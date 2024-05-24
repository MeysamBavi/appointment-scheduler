FROM golang:1.21.4 AS build

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/src/business-manager src/business-manager
COPY backend/pkg pkg

RUN go build -o business-manager ./src/business-manager/cmd/business-manager

FROM debian:12.5

WORKDIR /app

COPY --from=build /app/business-manager .

EXPOSE 8080

CMD ["./business-manager"]
