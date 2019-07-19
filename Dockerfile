FROM golang:1.10.4 AS build-env

WORKDIR /go/src/Numberwang

COPY . /go/src/Numberwang

RUN cd /go/src/Numberwang && go get -d -v ./...

RUN go install -v ./...
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN go build -o goapp
ENTRYPOINT CompileDaemon -log-prefix=false -build="go install ." -command="Numberwang"



# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/Numberwang/goapp /Numberwang/
ENTRYPOINT ./goapp
