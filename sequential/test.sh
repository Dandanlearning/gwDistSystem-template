#!/bin/sh
#########################################################
#DO NOT MAKE ANY CHANGES TO THIS FILE
#OR YOU MAY LOSE YOUR POINTS!
#########################################################
set -x
# shellcheck disable=SC2129
echo "#DO NOT MODIFY THIS FILE" >> tmp_res.txt ;
go build -o sequential
./sequential -f test1.txt >> tmp_res.txt ;
./sequential -f test2.txt >> tmp_res.txt ;
./sequential -f test3.txt >> tmp_res.txt ;
if cmp test_exp.txt tmp_res.txt
then
  echo '---' sequentail sum test: PASS
  rm -rf tmp_res.txt
else
  echo '---' sequential sum test: FAIL
  rm -rf tmp_res.txt
  rm -rf ./sequential
  exit 1
fi

rm -rf ./sequential

printf "\n"
printf "#####################  TEST PASSES  ####################################"
printf "\n"
