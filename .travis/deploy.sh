#!/usr/bin/env bash
#chmod ugo+x deploy.sh
echo DEPLLLOY
pwd
ssh -o StrictHostKeyChecking=no ubuntu@95.163.209.195
echo DEPLLLOYMachine
pwd
ls
echo SUCCCES > alahbabah.txt
#cd go
#rm -rf 2020_2_ExtremeCoders
#git clone https://github.com/sergii1/2020_2_ExtremeCoders
#cd 2020_2_ExtremeCoders
./cmd/runAllServices.sh