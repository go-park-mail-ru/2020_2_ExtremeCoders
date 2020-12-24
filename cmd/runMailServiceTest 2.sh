#!/bin/bash
#chmod ugo+x runMailServiceTest.sh
pwd
cd MailService
pwd
go generate ./internal/...
go test -v ./test/...
