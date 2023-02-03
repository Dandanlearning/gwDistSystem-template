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
echo "#DO NOT MODIFY THIS FILE" >> tmp_res.txt ;

cd rpc_back
go build -o server
./server &
sleep 1s
cd ../http_front
go build -o api
./api &
cd ../
sleep 2s
# shellcheck disable=SC2129
curl --request GET -sL \
     --url 'http://localhost:8080/?f=test1.txt&g=1' >> tmp_res.txt;
printf "\n" >> tmp_res.txt;

curl --request GET -sL \
     --url 'http://localhost:8080/?f=test2.txt&g=5' >> tmp_res.txt;
printf "\n" >> tmp_res.txt;

curl --request GET -sL \
     --url 'http://localhost:8080/?f=test3.txt&g=10' >> tmp_res.txt;
printf "\n" >> tmp_res.txt;
rm -rf http_front/api
rm -rf rpc_back/server
#export HOME="$PWD";
kill -9 $(lsof -i:8080|tail -1|awk '"$1"!=""{print $2}')

# shellcheck disable=SC1009
kill -9 $(lsof -i:8083|tail -1|awk '"$1"!=""{print $2}')
#echo "#DO NOT MODIFY THIS FILE" >> tmp_res.txt ;
if cmp test_exp.txt tmp_res.txt
then
  echo '---' parallel sum http test: PASS
  rm -rf tmp_res.txt
else
  echo '---' parallel sum http test: FAIL
  rm -rf tmp_res.txt
  exit 1
fi


printf "\n"
printf "#####################  TEST PASSES  ####################################"
printf "\n"


