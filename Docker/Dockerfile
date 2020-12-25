
### BUILD IMAGE
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/locker/
COPY . .

RUN go get -d -v

# Statically compile Locker, to run on Scratch
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/locker


### BUILD APP IMAGE
FROM scratch

COPY --from=builder /go/bin/locker /go/bin/locker
## Copy config yml
#COPY cfg-server.yml /go/bin/cfg-server.yml

EXPOSE 27001

ENTRYPOINT [ "/go/bin/locker", "server" ]