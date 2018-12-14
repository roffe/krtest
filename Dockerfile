FROM golang:1.11-alpine3.8 AS build
RUN apk add --no-cache git
WORKDIR /go/src/github.com/roffe/krtest
COPY . .
RUN go get github.com/gorilla/websocket
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o server cmd/server/server.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o client cmd/client/client.go

FROM scratch
COPY --from=build /go/src/github.com/roffe/krtest/server /go/bin/server
COPY --from=build /go/src/github.com/roffe/krtest/client /go/bin/client
ENTRYPOINT [ "/go/bin/server" ]

