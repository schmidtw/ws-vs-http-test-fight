To build:

## Build
export GOPATH=`pwd`
go get github.com/gorilla/websocket
go build http-responder.go
go build http-sender.go
go build ws-responder.go
go build ws-sender.go

To run:

./http-responder (in one terminal)
./http-sender (in a different terminal)

or

./ws-responder.go (in one terminal)
./ws-sender (in a different terminal)
