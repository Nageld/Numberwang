FROM golang:1.10.4 AS build-env
WORKDIR /go/src/Numberwang
COPY . /go/src/Numberwang
RUN cd /go/src/Numberwang && go get -d -v ./...
RUN go build -o goapp

# Hot-Reloading in local
FROM build-env AS local
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
ENTRYPOINT CompileDaemon -log-prefix=false -build="go install ." -command="Numberwang"

# final stage
# FROM alpine
FROM ubuntu:18.04
WORKDIR /temp
ENV PATH=$PATH:/temp
COPY --from=build-env /go/src/Numberwang/goapp /temp
CMD ["goapp"]