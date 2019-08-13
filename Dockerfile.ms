FROM golang:alpine

MAINTAINER alexdnn11

WORKDIR $GOPATH/src/go-grpc-k8s/server
COPY server .
COPY pb $GOPATH/src/go-grpc-k8s/server/vendor/go-grpc-k8s/pb

CMD [ "go", "run", "main.go" ]