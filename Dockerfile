FROM golang:1.25-trixie AS build

RUN apt-get update && apt-get install -y libmagickwand-dev libmagickcore-dev \
  imagemagick libjpeg-dev libpng-dev libgif-dev fonts-cantarell

WORKDIR /code
ENV CGO_ENABLED=1
COPY . /code
RUN go mod tidy && \
  go build -a -installsuffix cgo -o seaals .

FROM debian:trixie-slim

RUN apt-get update && apt-get install -y libmagickwand-dev libmagickcore-dev \
  libjpeg-dev libpng-dev libgif-dev

RUN adduser seaals
USER seaals
WORKDIR /app
ENV PATH="$PATH:/app"
COPY --from=build --chown=seaals:seaals /code/seaals /app/seaals
COPY --chown=seaals:seaals ./public /app/public

ENTRYPOINT [ "/app/seaals" ]
CMD [ "serve" ]
