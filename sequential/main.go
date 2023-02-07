/**
* @Author: huadong.hu@outlook.com
* @Date: 8/17/20 02:10
* @Desc:
 */
package main

import (
	"flag"
	"fmt"
)

func main() {
	sum := 0
	//@TODO read file name from command line
	//the argument should be `-f`
	var fileName string
	flag.StringVar(&fileName, "f", "", "")
	flag.Parse()
	sum = Sum(fileName)
	//DO NOT OUTPUT ANYTHING ABOVE THIS LINE
	//DO NOT MODIFY OUTPUT FORMAT!!
	fmt.Println(sum)
}
