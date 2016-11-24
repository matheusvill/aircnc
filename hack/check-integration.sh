#!/bin/bash

go run /go/src/github.com/matheusvill/aircnc/hack/create-db.go
go test -tags integration -p 1 -v -covermode=atomic -timeout 100s $(glide novendor) | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
