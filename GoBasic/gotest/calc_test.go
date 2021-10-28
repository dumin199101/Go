package main

import (
	"testing"
)

// 测试单文件 go test -v 单文件 依赖文件
// 测试单个函数 go test -v -test.run 函数名
func TestAdd(t *testing.T) {
	res := add(1, 2)
	if res == 3 {
		t.Logf("测试通过，期望值：%d,实际值：%d", 3, res)
	} else {
		t.Fatalf("测试失败")
	}
}

func TestSub(t *testing.T) {
	res := sub(10, 3)
	if res == 7 {
		t.Logf("测试通过，期望值：%d,实际值：%d", 7, res)
	} else {
		t.Fatalf("测试失败")
	}
}
