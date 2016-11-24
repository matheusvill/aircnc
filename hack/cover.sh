#!/bin/bash

./hack/build.sh

go run /go/src/github.com/matheusvill/aircnc/hack/create-db.go
touch coverage.out
echo "mode: atomic" > full_coverage.out
for pkg in $(go list ./... | grep -v /vendor/); do
	go test -tags integration -coverprofile=coverage.out -race -covermode=atomic -timeout 60s $pkg
	grep -h -v "^mode: atomic" coverage.out >> full_coverage.out
done

go tool cover -html=full_coverage.out -o=coverage.html
rm -f coverage.out full_coverage.out
