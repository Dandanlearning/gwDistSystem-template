/**
* @Author: huadong.hu@outlook.com
* @Date: 7/7/20 8:51 PM
* @Desc:
 */

package main

//You Should not use time.sleep() to block go routines

import (
	"bufio"
	"log"
	"os"
	"strconv"
)



// Read a list of integers from `fileName`
// and launch `goRoutineNums` go routines to do sum up
// return sum of these Integers
func Sum(goRoutineNums int, fileName string) int {
	sum := 0
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

//Read integers from reader
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

