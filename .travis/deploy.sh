#!/usr/bin/env bash
#chmod ugo+x deploy.sh
echo DEPLLLOY
ssh ubuntu@95.163.209.195 << EOF
pwd
echo DEPLLLOYMachine > flsfadlf.txt
pwd
ls
cd go
ls
rm -rf 2020_2_ExtremeCoders
git clone https://github.com/go-park-mail-ru/2020_2_ExtremeCoders
cd 2020_2_ExtremeCoders
ls
git checkout webSocket
git pull
#cat ~/myconf.go > config/conf.go
#cat config/config.go
./cmd/runAllServices.sh
sleep 35
echo res
cat FileService.txt
cat MailService.txt
cat UserService.txt
cat MainApplication.txt
EOF