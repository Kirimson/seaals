FROM golang:1.25-alpine3.23 AS build

RUN apk add build-base vips-dev jpeg-dev tiff-dev libexif-dev libpng-dev libwebp-dev pango-dev librsvg

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go install github.com/cshum/vipsgen/cmd/vipsgen@latest && vipsgen -out ./vips

ENV CGO_ENABLED=1
ENV CGO_CFLAGS_ALLOW=-Xpreprocessor
RUN go build -v -o /usr/local/bin/seaals .

FROM alpine:3.23

RUN apk add vips-dev jpeg-dev tiff-dev libexif-dev libpng-dev libwebp-dev pango-dev librsvg font-noto

RUN adduser seaals -D -H
USER seaals

WORKDIR /opt/seaals
COPY --from=build /usr/local/bin/seaals /usr/local/bin/seaals
COPY --chown=seaals:seaals ./public /opt/seaals/public
RUN mkdir /opt/seaals/data

ENTRYPOINT [ "seaals" ]
CMD ["serve"]
