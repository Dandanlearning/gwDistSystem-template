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

func test(t *testing.T, fileName string, num int, expected int) {
	res := Sum(num, fileName)
	if res != expected {
		t.Fatal(fmt.Sprintf(
			"Test %s Failed: expected %d, got %d\n", fileName,expected, res))
	}
}

//following tests are for parallel sum tests
func Test1_1(t *testing.T) {
	test(t, "test1.txt", 1, -83)
}

func Test1_2(t *testing.T) {
	test(t, "test1.txt", 10, -83)
}

func Test2_1(t *testing.T)  {
	test(t, "test2.txt", 1, 19605)

}

func Test2_2(t *testing.T)  {
	test(t, "test2.txt", 10, 19605)
}

func Test3_1(t *testing.T)  {
	test(t, "test3.txt", 5, 31418)
}

func Test3_2(t *testing.T)  {
	test(t, "test3.txt", 20, 31418)
}

