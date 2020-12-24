#!/bin/bash
#chmod ugo+x runFileServiceTest.sh
pwd
cd FileService
pwd
go test -v ./test/tests/...
