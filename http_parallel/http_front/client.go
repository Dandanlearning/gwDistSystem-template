/**
* @Author: huadong.hu@outlook.com
* @Date: 7/7/20 9:40 PM
* @Desc:
 */
package main

import (
	"net/rpc"
	// "fmt"

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
	var args SumServiceReq 
	args.FileName = fileName
	args.GoRoutineNums = goRoutineNums
	var reply SumServiceResp
	err := sumCli.Client.Call("SumService.CalcSum", &args, &reply)
	if err != nil {
		// fmt.Println(err)
		return 0, err
	}
	return reply.Sum, err
}


