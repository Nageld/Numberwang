FROM golang:1.10.4

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...


EXPOSE 8000
ENV PORT=8000

CMD ["app"]