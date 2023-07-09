FROM golang:alpine
COPY ./bin/getmark-home /release/
COPY ./cmd/api/config.env /release/
RUN mkdir /release/tags

WORKDIR "/release"
ENTRYPOINT ["./getmark-home", "-l", ":8080"]
