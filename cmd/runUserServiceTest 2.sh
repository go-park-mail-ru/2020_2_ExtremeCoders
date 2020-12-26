#!/bin/bash
#chmod ugo+x runUserServiceTest.sh
pwd
cd UserService
pwd
go generate ./...
go test -v ./test/test/...
