#!/bin/bash
#chmod ugo+x runUserServiceTest.sh
pwd
cd UserService
pwd
go test -v ./test/test/...
