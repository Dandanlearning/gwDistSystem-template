/**
* @Author: huadong.hu@outlook.com
* @Date: 8/10/20 10:55
* @Desc:
 */
package main

import (
	"fmt"
	"testing"
)

func sequentialSumTest(t *testing.T, fileName string,  expected int) {
	res := Sum(fileName)
	if res != expected {
		t.Fatal(fmt.Sprintf(
			"Test %s Failed: expected %d, got %d\n", fileName,expected, res))
	}
}

func Test0_1(t *testing.T) {
	sequentialSumTest(t, "test1.txt", -83)
}

func Test0_2(t *testing.T) {
	sequentialSumTest(t, "test2.txt", 19605)
}

func Test0_3(t *testing.T) {
	sequentialSumTest(t, "test3.txt", 31418)
}
