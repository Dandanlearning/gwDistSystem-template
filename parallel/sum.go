/**
* @Author: huadong.hu@outlook.com
* @Date: 8/10/20 10:20
* @Desc:
 */
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"runtime"
)

// Read a list of integers from `fileName`
// and launch `goRoutineNums` go routines to do calculation
// return sum of these Integers
// You Must start exact `goRoutineNums` go routines or you lose points here
func Sum(goRoutineNums int, fileName string) int {
	//TODO Add your code here
	var res, error = readInts(fileName)
	if error != nil{
		return 0
	}
	runtime.GOMAXPROCS(goRoutineNums)
	sum := 0
	for _, value := range res{
		sum += value
	}

	return sum
}

//Read integers from file
//Do not modify this function
func readInts(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	cin := bufio.NewScanner(file)
	cin.Split(bufio.ScanWords)
	var res []int
	for cin.Scan() {
		val, err := strconv.Atoi(cin.Text())
		if err != nil {
			return res, err
		}
		res = append(res, val)
	}
	return res, nil
}

