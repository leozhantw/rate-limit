FROM golang:1.15.6-alpine3.12 AS build
WORKDIR /src
COPY . .
RUN go build -mod=vendor -o bin/server ./cmd/server

FROM alpine:3.12
COPY --from=build /src/bin/server .
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ./server
