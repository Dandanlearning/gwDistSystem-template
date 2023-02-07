/**
* @Author: huadong.hu@outlook.com
* @Date: 8/17/20 02:09
* @Desc:
 */
package main

import (
	"fmt"
	"flag"
)


func main() {
	sum := 0
	//@TODO read file name and goroutine numbers from command line and output its sum
	//the argument should be `-f` `-g`
	var GoroutineNums int
	flag.IntVar(&GoroutineNums, "g", 0, "")
	var fileName string
	flag.StringVar(&fileName, "f", "", "")
	flag.Parse()
	sum = Sum(GoroutineNums, fileName)
	//DO NOT OUTPUT ANYTHING ABOVE THIS LINE
	//DO NOT MODIFY OUTPUT FORMAT!!
	fmt.Println(sum)
}
