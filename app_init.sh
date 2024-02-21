#!/usr/bin/env sh

sleep 3

go get -u github.com/jteeuwen/go-bindata/...

sleep 3

sh ./build/make_graphql.sh 1

sleep 5

go build -o graphql_app ./cmd/main.go

sleep 3

./graphql_app