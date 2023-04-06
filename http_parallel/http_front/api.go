/**
* @Author: huadong.hu@outlook.com
* @Date: 7/7/20 9:44 PM
* @Desc:
 */


package main

import (
	"strconv"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"time"
)


var sumCli = SumServiceClient{}
func SumHandler(wr http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wr.Write([]byte(err.Error()))
		return
	}
	var fileName string
	var goRoutineNums int
	var totalSum int
	//TODO You should add your code here
	//HINT: Receive fileName(f) and goRoutineNums(g) from URL
	//HINT: Call sumCli.Sum(fileName, goRoutineNums) to retrieve sum result
	vars := r.URL.Query()
	// nums := vars.Get("g")
	// fileName = r.URL.Query().Get("f")
	goRoutineNumsStr := r.URL.Query().Get("g")
	goRoutineNums, err = strconv.Atoi(goRoutineNumsStr)
	fileName = vars.Get("f")
	totalSum, err = sumCli.Sum(fileName, goRoutineNums)
	// log.Println("err:", err)

	//DO NOT MODIFY OUTPUT FORMAT!
	type OutPut struct {
		FileName string `json:"file_name"`
		GoRoutineNums int `json:"go_routine_nums"`
		TotalSum int `json:"sum"`
	}

	out := OutPut{
		FileName: fileName,
		GoRoutineNums: goRoutineNums,
		TotalSum: totalSum,
	}
	data, err := json.Marshal(out)
	// log.Println("err:", err)
	if err != nil {
		panic(err)
	}
	wr.Write(data)
	writeLen, err := wr.Write(msg)
	if err != nil || writeLen != len(msg) {
		log.Println(err, "write len:", writeLen)
	}
}

func main() {
	time.Sleep(100 * time.Millisecond)
	go regRpcService()
	handleHttp()
}

func regRpcService() {
	client, err := rpc.Dial("tcp",  ":8083")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	sumCli.Client = client
}


func handleHttp() {
	http.HandleFunc("/", SumHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}



