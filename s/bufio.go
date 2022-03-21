package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	// TODO bufio包详解

	// 打开文件
	if file, err := os.Open("s/bufio.go"); err == nil {
		// 创建一个具有默认大小缓冲、从r读取的*Reader
		reader := bufio.NewReader(file)
		// 创建一个具有最少有size尺寸的缓冲、从r读取的*Reader
		reader = bufio.NewReaderSize(file, 4096)
		p := make([]byte, 1024)
		// 读取数据并写入到p中
		if n, err := reader.Read(p); err == nil {
			fmt.Println("读取长度：", n, " 读取内容：", string(p[:n]))
		}
		// 读取一行数据（ReadLine是一个低水平的行数据读取原语。大多数调用者应使用ReadBytes('\n')或ReadString('\n')代替，或者使用Scanner）
		if line, b, err := reader.ReadLine(); err == nil {
			fmt.Println("读取内容：", string(line), " 是否超过缓冲：", b)
		}
		// 读取直到遇到delim界定符的位置，返回一个已读取的包含delim字节的字符串
		if readString, err := reader.ReadString('\n'); err == nil {
			fmt.Println("读取内容：", readString)
		}
		// 读取直到遇到delim界定符的位置，返回一个已读取的包含delim字节的byte切片
		if bytes, err := reader.ReadBytes('\n'); err == nil {
			fmt.Println("读取内容：", string(bytes))
		}
		// 读取并返回一个字节
		if readByte, err := reader.ReadByte(); err == nil {
			fmt.Println("读取内容：", string(readByte))
		}
		// 读取直到遇到delim界定符的位置，返回一个已读取的包含delim字节的byte切片
		if line, err := reader.ReadSlice('\n'); err == nil {
			fmt.Println("读取内容：", string(line))
		}
		// 读取一个utf-8编码的unicode码值
		if readRune, i, err := reader.ReadRune(); err == nil {
			fmt.Println("unicode字节编码：", readRune, " 字节大小：", i)
		}
		// 回滚最近一次读取操作读取的最后一个字节
		if err := reader.UnreadByte(); err != nil {
			log.Fatalln(err)
		}
		// 回滚最近一次ReadRune调用读取的unicode码值
		/*if err := reader.UnreadRune(); err != nil {
			log.Fatalln(err)
		}*/
		// 返回缓冲中已使用的字节数
		fmt.Println("缓冲中已使用的字节数", reader.Buffered())
		// 返回接下来的n个字节，而不会推进阅读器。字节在下次读取呼叫时停止有效
		if peek, err := reader.Peek(10); err == nil {
			fmt.Println("接下来预读取的10个字节：", peek, " 读取内容：", string(peek))
		}
		// 跳过接下来的n个字节，返回被丢弃的字节数
		if discard, err := reader.Discard(10); err == nil {
			fmt.Println("被丢弃的字节数：", discard)
		}
		// 返回基础缓冲区的大小
		fmt.Println("基础缓冲区的大小：", reader.Size())
		// 实现 io.WriterTo 接口
		bf := bytes.NewBuffer(make([]byte, 1024))
		if n, err := reader.WriteTo(bf

		); err == nil {
			fmt.Println(n)
		}

		/*reader.Reset()
		reader.WriteTo()*/

	} else {
		fmt.Println(err)
	}

}
