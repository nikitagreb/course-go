FROM golang:1.17

WORKDIR /go
#WORKDIR /go/src/app

#RUN go get -d -v ./...
#RUN go install -v ./...
#
#CMD ["app"]


# Compile stage
#FROM golang:1.17 AS build-env
#
#ADD . /dockerdev
#WORKDIR /dockerdev
#
#RUN go build -o /server
#
## Final stage
#FROM debian:buster
#
#EXPOSE 8000
#
#WORKDIR /
#COPY --from=build-env /server /
#
#CMD ["/server"]