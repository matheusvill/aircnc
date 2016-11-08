#!/bin/bash

go test -tags unit -v -race -covermode=atomic -timeout 30s $(glide novendor) | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
