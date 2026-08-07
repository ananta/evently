FROM golang:1.26-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /api ./cmd/api

FROM alpine:3.21

COPY --from=build /api /usr/local/bin/api

EXPOSE 4000
ENTRYPOINT ["api"]
