package main

import (
	"testing"
)

//单元测试
func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		//输出错误日志
		t.Fatalf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
	}

	//输出成功日志
	t.Logf("AddUpper(10) 执行成功")
}
