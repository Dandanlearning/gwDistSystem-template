#!/bin/sh
#########################################################
#DO NOT MAKE ANY CHANGES TO THIS FILE
#OR YOU MAY LOSE YOUR POINTS!
#########################################################
set -x
rm -rf tmp_res.txt
rm -rf ./rpc_back/server
rm -rf ./http_front/api
#export HOME="$PWD";
kill -9 $(lsof -i:8080|tail -1|awk '"$1"!=""{print $2}')

# shellcheck disable=SC1009
kill -9 $(lsof -i:8083|tail -1|awk '"$1"!=""{print $2}')

cd rpc_back
go build -o server
./server &
sleep 2s
cd ../http_front
go build -o api
./api &
cd ../
sleep 3s
perpid=`lsof -i :8080|grep -v "PID" | awk '{print $2}'`
if [ $perpid ];then
        echo "HTTP Status OK"
        rm -rf ./rpc_back/server
        rm -rf ./http_front/api
        kill -9 $(lsof -i:8083|tail -1|awk '"$1"!=""{print $2}')
        kill -9 $(lsof -i:8080|tail -1|awk '"$1"!=""{print $2}')
else
        echo "HTTP Status ERR"
        rm -rf ./rpc_back/server
        rm -rf ./http_front/api
        kill -9 $(lsof -i:8083|tail -1|awk '"$1"!=""{print $2}')
        kill -9 $(lsof -i:8080|tail -1|awk '"$1"!=""{print $2}')
        exit 1
fi

printf "\n"
printf "#####################  TEST PASSES  ####################################"
printf "\n"

