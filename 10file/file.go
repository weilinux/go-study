package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "C:/Users/loril/Desktop/aaa.txt"

	//第一种打开文件方式 适用于大文件

	//打开文件
	//默认打开
	//open, err := os.Open(filePath)
	//通过权限打开
	open, err := os.OpenFile(filePath, os.O_RDONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	fmt.Printf("文件类型：%T\n", open)
	fmt.Println(open.Name())

	//关闭文件
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			return
		}
	}(open)

	//创建带缓冲的*Reader实例, 默认缓冲大小为4096
	reader := bufio.NewReader(open)
	//读取文件
	var fileInfo string
	for {
		readString, err := reader.ReadString('\n')
		fileInfo += readString
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("文件内容： %v", fileInfo)

	//创建带缓冲的*Writer实例，默认缓冲大小为4096
	writer := bufio.NewWriter(open)
	for i := 0; i < 5; i++ {
		_, err := writer.WriteString("追加写入```\n")
		if err != nil {
			return
		}
	}

	//写入数据需要通过Flush方法将缓存数据写入到指定的文件中，不写则无任何效果
	err = writer.Flush()
	if err != nil {
		return
	}

	//判断文件是否存在
	stat, err := os.Stat("C:/Users/loril/Desktop/aaa.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
		}
		return
	}
	fmt.Println(stat)

	/*//第二种打开文件方式 适用于小文件
	file, err := os.ReadFile(filePath)
	if err != nil {
		return
	}
	fmt.Printf("文件内容： %s", file)
	//file2 := string(file)
	//fmt.Println(file2)*/

}
