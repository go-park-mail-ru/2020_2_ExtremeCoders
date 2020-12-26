#!/bin/bash
#chmod ugo+x runFileServiceTest.sh
pwd
cd FileService
pwd
go generate ./...
go test -v ./test/tests/...
