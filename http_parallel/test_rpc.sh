#!/bin/sh
#########################################################
#DO NOT MAKE ANY CHANGES TO THIS FILE
#OR YOU MAY LOSE YOUR POINTS!
#########################################################
set -x
rm -rf tmp_res.txt

#export HOME="$PWD";
kill -9 $(lsof -i:8080|tail -1|awk '"$1"!=""{print $2}')

# shellcheck disable=SC1009
kill -9 $(lsof -i:8083|tail -1|awk '"$1"!=""{print $2}')

cd rpc_back
go build -o server
./server &
sleep 1s

perpid=`lsof -i :8083|grep -v "PID" | awk '{print $2}'`
if [ $perpid ];then
        echo "RPC Status OK"
        rm -rf ./server
        kill -9 $(lsof -i:8083|tail -1|awk '"$1"!=""{print $2}')
else
        echo "RPC Status ERR"
        rm -rf ./server
        exit 1
fi


printf "\n"
printf "#####################  TEST PASSES  ####################################"
printf "\n"


