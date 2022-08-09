FROM golang:1.18-alpine AS build
WORKDIR /app/

COPY . .

RUN go build -o redirector

FROM alpine AS run
WORKDIR /app/

COPY --from=build /app/redirector .

ENTRYPOINT ["/app/redirector"]
