package main

import (
	"fmt"
	"github.com/pkg/errors"
)

// ErrorString 自定义错误类型
type ErrorString struct {
	s string
}

// 实现错误方法
func (e *ErrorString) Error() string {
	return e.s
}

func main() {
	//TODO 使用pkg/errors操作
	// go get -u github.com/pkg/errors

	// 新生成一个错误，带堆栈信息
	err := errors.New("new error")

	// 只附加新的信息 字符串
	err2 := errors.WithMessage(err, "with message 111")

	// 只附加新的信息 格式说明符
	err3 := errors.WithMessagef(err2, "with mesaage %s", "111")

	// 只附加调用堆栈信息
	err4 := errors.WithStack(err3)

	// 同时附加堆栈和信息 字符串
	err5 := errors.Wrap(err4, "wrap 111")

	// 同时附加堆栈和信息 格式说明符
	err6 := errors.Wrapf(err5, "wrap %s", "111")

	// 返回最底层的原始error对象（该错误链中最顶级的错误）
	err7 := errors.Cause(err6)

	// 打印出堆栈信息
	/*
	   %s,%v 	//功能一样，输出错误信息，不包含堆栈
	   %q 		//输出的错误信息带引号，不包含堆栈
	   %+v 		//输出错误信息和堆栈
	*/
	err8 := errors.Errorf("%+v", err7)
	fmt.Println("错误信息：", err8.Error())

	// 判断错误链中是否有错误和 target 匹配 （如果能匹配上则返回true，否则返回false）
	fmt.Println("是否能匹配错误：", errors.Is(err6, err5))

	// 寻找错误链中第一个和 target 匹配的error（如果能找到就将err的值设置到target并返回true，否则返回false）
	var targetErr *ErrorString
	err10 := fmt.Errorf("new error:[%w]", &ErrorString{s: "target err"})
	fmt.Println("是否能匹配错误", errors.As(err10, &targetErr))

	// 获取指定错误上一级的error对象（一般通过递归调用，如果为最顶级的错误则返回nil）
	err5 = errors.Unwrap(err6)
}
