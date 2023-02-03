#!/bin/sh
#########################################################
#DO NOT MAKE ANY CHANGES TO THIS FILE
#OR YOU MAY LOSE YOUR POINTS!
#########################################################
set -x
# shellcheck disable=SC2129
echo "#DO NOT MODIFY THIS FILE" >> tmp_res.txt ;
go run main.go sum.go -f test1.txt -g 1 >> tmp_res.txt ;
go run main.go sum.go -f test2.txt -g 5 >> tmp_res.txt ;
go run main.go sum.go -f test3.txt -g 10 >> tmp_res.txt ;
if cmp test_exp.txt tmp_res.txt
then
  echo '---' parallel sum basic test: PASS
  rm -rf tmp_res.txt
else
  echo '---' parallel sum basic test: FAIL
  rm -rf tmp_res.txt
  exit 1
fi

printf "\n"
printf "#####################  TEST PASSES  ####################################"
printf "\n"
