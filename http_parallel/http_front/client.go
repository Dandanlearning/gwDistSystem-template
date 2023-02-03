/**
* @Author: huadong.hu@outlook.com
* @Date: 7/7/20 9:40 PM
* @Desc:
 */
package main

import (
	"net/rpc"
)

type SumServiceReq struct {
	FileName string
	GoRoutineNums 	int
}

type SumServiceResp struct {
	Sum	int
}

type SumServiceClient struct {
	Client *rpc.Client
}

func (s *SumServiceClient) Sum(fileName string, goRoutineNums int) (int, error) {
	//TODO Add your code here
	//Hint: Here is RPC Client request
	return 0, nil
}


