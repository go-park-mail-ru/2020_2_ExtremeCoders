#!/bin/bash
#chmod ugo+x runAllServices.sh
#cd ..
#pwd
echo this is deploy in path
pwd
echo 'asdfasdf' > ~/asdfsdfs.txt
sudo kill 9 `sudo lsof -t -i:8080`
sudo kill 9 `sudo lsof -t -i:8081`
sudo kill 9 `sudo lsof -t -i:8082`
sudo kill 9 `sudo lsof -t -i:8083`
sudo kill 9 `sudo lsof -t -i:8084`
go mod vendor
go run ./FileService/cmd/main.go &>FileService.txt &
go run ./MailService/cmd/main.go &>MailService.txt &
go run ./UserService/cmd/main.go &>UserService.txt &
go run ./MainApplication/cmd/main.go &>MainApplication.txt &
