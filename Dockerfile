FROM golang:1.21.1-alpine3.17
RUN apk update && apk add git curl alpine-sdk
RUN mkdir /go/src/samplepb
WORKDIR /go/src/samplepb
COPY ./* /go/src/samplepb
# protocのインストール
RUN apk add --no-cache protobuf-dev
# protocプラグインのインストール
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# protobufモジュールのインストール
RUN go get -u google.golang.org/grpc@latest
