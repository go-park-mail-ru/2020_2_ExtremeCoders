#!/bin/bash
#chmod ugo+x runMailServiceTest.sh
pwd
cd MailService
pwd
go test -v ./test/...
