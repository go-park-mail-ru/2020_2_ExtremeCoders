#!/bin/bash
#chmod ugo+x runMainApplicationTest.sh
pwd
cd MainApplication
pwd
go test -v ./test/test/...
