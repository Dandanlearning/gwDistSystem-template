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
	// "runtime"
)



// Read a list of integers from `fileName`
// and launch `goRoutineNums` go routines to do sum up
// return sum of these Integers
func Sum(goRoutineNums int, fileName string) int {
	var res, err = readInts(fileName)
	if err != nil{
		log.Println("Error reading file:", err)
		return 0
	}
	chunkSize := (len(res) + goRoutineNums - 1) / goRoutineNums
	sums := make(chan int, goRoutineNums)

	// Spawn a goroutine for each chunk of integers
	for i := 0; i < goRoutineNums; i++ {
		start := i * chunkSize
		end := (i+1)*chunkSize
		if end > len(res) {
			end = len(res)
		}
		go func(res []int) {
			sum := 0
			for _, n := range res {
				sum += n
			}
			sums <- sum
		}(res[start:end])
	}

	// Collect the computed sums from each goroutine and add them up
	totalSum := 0
	for i := 0; i < goRoutineNums; i++ {
		sum := <-sums
		totalSum += sum
	}

	// Return the computed sum
	return totalSum

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

