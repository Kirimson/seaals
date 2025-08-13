FROM golang:1.24-alpine AS build

RUN apk add --no-cache build-base imagemagick-dev imagemagick

COPY . /code
WORKDIR /code
ENV CGO_ENABLED=1
ENV CGO_CFLAGS_ALLOW=-Xpreprocessor
RUN go mod tidy && \
  go build -a -installsuffix cgo -o app .

FROM alpine:3.22
RUN apk add --no-cache imagemagick-dev imagemagick
COPY --from=build /code/app /app/seaals
COPY seal.jpeg /app/seal.jpeg

WORKDIR /app
ENTRYPOINT [ "./seaals" ]
