package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//TODO 文件拷贝

	//源文件地址
	srcPath := "D:/adaf2edda3cc7cd970a17b3d3b01213fb90e9142.jpg"

	//打开文件
	file, err := os.OpenFile(srcPath, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println("源文件打开失败")
		return
	}

	//关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("文件关闭失败")
		}
	}(file)

	//创建读缓存
	reader := bufio.NewReader(file)

	//目标地址
	dstPath := "E:/adaf2edda3cc7cd970a17b3d3b01213fb90e9142.jpg"

	//打开文件
	openFile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("目标地址打开失败")
		return
	}

	//关闭文件
	defer func(openFile *os.File) {
		err := openFile.Close()
		if err != nil {
			fmt.Println("文件关闭失败")
		}
	}(openFile)

	//创建写缓存
	writer := bufio.NewWriter(openFile)

	//copy文件到指定目录
	written, err := io.Copy(writer, reader)
	if err != nil {
		fmt.Println("文件copy失败")
		return
	}

	fmt.Printf("文件写入大小： %v\n", written)
}
