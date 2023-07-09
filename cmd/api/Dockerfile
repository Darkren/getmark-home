FROM golang:alpine
COPY ./bin/getmark-home /release/
COPY ./assets /release/assets
RUN mkdir /release/tags

WORKDIR "/release"
ENTRYPOINT ["./getmark-home", "-l", ":8080"]
