package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//源文件路径
	srcFile := "C:\\Users\\loril\\Pictures\\Saved Pictures\\adaf2edda3cc7cd970a17b3d3b01213fb90e9142.jpg"

	//目标文件路径
	destFile := srcFile[strings.LastIndex(srcFile, "\\")+1:]

	//临时文件路径
	tempFile := destFile + "temp.txt"

	//打开源文件 （只读）
	file1, err := os.Open(srcFile)
	handleErr(err)

	//打开目标文件 （不存在则创建、只写  linux操作系统下文件模式为：0777）
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	handleErr(err)

	//打开临时文件 （不存在则创建、读写  linux操作系统下文件模式为：0777）
	file3, err := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)

	//关闭文件
	defer func(file1, fiel2 *os.File) {
		err = file1.Close()
		handleErr(err)

		err = file2.Close()
		handleErr(err)
	}(file1, file2)

	//设置读取临时文件的偏移量，从起始位置开始读取
	_, err = file3.Seek(0, io.SeekStart)
	handleErr(err)

	//创建一个byte切片 （用于保存读取临时文件的内容）
	bs := make([]byte, 100)

	//读取临时文件中的偏移值
	n1, err := file3.Read(bs)
	if err != io.EOF {
		handleErr(err)
	}

	//将该byte切片取到临时文件读取的长度
	countStr := string(bs[:n1])

	//创建一个int64的变量 （用于保存临时文件的内容（转成int64类型））
	var count int64

	//如果不为空时（程序中断等原因）
	if countStr != "" {
		//将临时文件中读取到的长度按十进制转成int64类型
		count, err = strconv.ParseInt(countStr, 10, 64)
		handleErr(err)
	}

	//设置读取源文件的偏移量，从临时文件存储的位置开始读取 （默认是0，从首位开始）
	_, err = file1.Seek(count, io.SeekStart)
	handleErr(err)

	//设置读取目标文件的偏移量，从临时文件存储的位置开始读取 （默认是0，从首位开始）
	_, err = file2.Seek(count, io.SeekStart)
	handleErr(err)

	//创建一个byte切片 （用于保存读取源文件的内容）
	data := make([]byte, 1024)

	//创建一个int变量  （用于保存读取源文件的大小）
	n2 := -1

	//创建一个int变量  （用于保存写入目标文件的大小）
	n3 := -1

	//创建一个int变量  （用于保存源文件读取进度）
	total := int(count)

	for {
		//读取源文件
		n2, err = file1.Read(data)

		//如果文件读取结束或者源文件为空
		if err == io.EOF || n2 == 0 {
			fmt.Println("文件已读取完毕")

			//关闭临时文件
			err = file3.Close()
			handleErr(err)

			//删除临时文件
			err = os.Remove(tempFile)
			handleErr(err)

			//跳出该循环体
			break
		}

		//将源文件内容写入目标文件
		n3, err = file2.Write(data[:n2])

		//累加写入目标文件后的进度
		total += n3

		//设置临时文件的偏移量，从0开始 （覆盖，每次源文件和目标文件的读写进度更改都更新临时文件储存的进度值）
		_, err = file3.Seek(0, io.SeekStart)
		handleErr(err)

		//将读写进度
		_, err = file3.WriteString(strconv.Itoa(total))
		handleErr(err)

		//模拟程序终止
		/*if total > 8000 {
			panic("程序发生意外情况已终止运行")
		}*/
	}
}
