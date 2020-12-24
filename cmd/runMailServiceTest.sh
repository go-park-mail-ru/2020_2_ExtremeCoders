#!/bin/bash
#chmod ugo+x runMailServiceTest.sh
pwd
cd MailService
pwd
go generate ./...
go test -v ./test/...
