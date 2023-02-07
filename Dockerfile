FROM golang:1.20-alpine AS build
ENV CGO_ENABLED=0
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o /bin/protohackers main.go

FROM alpine:latest
COPY --from=build /bin/protohackers /bin/protohackers

ENTRYPOINT ["/bin/protohackers"]
