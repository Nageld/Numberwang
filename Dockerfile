FROM golang:1.10.4

WORKDIR /go/src/Numberwang

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

# EXPOSE 8000
# ENV PORT=8000

ENTRYPOINT CompileDaemon -log-prefix=false -build="go install ." -command="Numberwang"

# CMD ["app"]

