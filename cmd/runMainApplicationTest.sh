#!/bin/bash
#chmod ugo+x runMainApplicationTest.sh
pwd
cd MainApplication
pwd
go generate ./...
go test -v ./test/test/...
