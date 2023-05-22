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
		if readBytes, err := reader.ReadBytes('\n'); err == nil {
			fmt.Println("读取内容：", string(readBytes))
		}
		// 读取并返回一个字节
		if readByte, err := reader.ReadByte(); err == nil {
			fmt.Println("读取内容：", string(readByte))
		}
		// 读取直到遇到delim界定符的位置，返回一个已读取的包含delim字节的byte切片
		if readLine, err := reader.ReadSlice('\n'); err == nil {
			fmt.Println("读取内容：", string(readLine))
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
		// 实现 io.WriterTo 接口，将缓冲区的数据写入到指定的变量中
		bf := bytes.NewBuffer(make([]byte, 4096))
		if n, err := reader.WriteTo(bf); err == nil {
			fmt.Println("写入大小：", n, " 写入内容：", bf)
		}
		// 丢弃缓冲中的数据，清除任何错误，将reader重设为其下层从新设置的缓冲对象读取数据
		bf2 := bytes.NewBuffer(make([]byte, 4096))
		reader.Reset(bf2)
		// 创建一个具有默认大小缓冲、写入w的*Writer
		writer := bufio.NewWriter(file)
		// 创建一个具有最少有size尺寸的缓冲、写入w的*Writer
		writer = bufio.NewWriterSize(file, 4096)

		p2 := make([]byte, 1024)
		// 将p的内容写入缓冲
		if n, err := writer.Write(p2); err == nil {
			fmt.Println("写入长度：", n, " 写入内容：", string(p[:n]))
		}
		// 写入一个字符串
		if writeString, err := writer.WriteString(""); err == nil {
			fmt.Println("写入长度：", writeString)
		}
		// 写入单个字节
		if err := writer.WriteByte('\n'); err != nil {
			log.Fatalln(err)
		}
		// 写入一个unicode码值（的utf-8编码）
		if writeRune, err := writer.WriteRune('测'); err == nil {
			fmt.Println("写入长度：", writeRune)
		}
		// 返回基础缓冲区的大小
		fmt.Println("基础缓冲区的大小：", writer.Size())
		// 返回缓冲中已使用的字节数
		fmt.Println("缓冲中已使用的字节数", writer.Buffered())
		// 返回缓冲中还有多少字节未使用
		fmt.Println("未使用的字节数：", writer.Available())
		// 返回一个容量为b.Available（未使用字节大小）的空缓冲区
		fmt.Println("空缓冲区：", writer.AvailableBuffer())
		// 将缓冲中的数据写入下层的io.Writer接口
		if err := writer.Flush(); err != nil {
			log.Fatalln(err)
		}
		// Reset丢弃缓冲中的数据，清除任何错误，将b重设为将其输出写入w
		bf3 := bytes.NewBuffer(make([]byte, 4096))
		writer.Reset(bf3)
		// 将任何缓冲的数据写入底层的 io.Writer，将缓冲区的数据写入到指定的变量中
		if n, err := writer.ReadFrom(file); err == nil {
			fmt.Println("读取长度：", n)
		}

		// 申请创建一个新的、将读写操作分派给r和w 的ReadWriter
		fmt.Println("读写对象：", bufio.NewReadWriter(reader, writer))

		// 创建并返回一个从r读取数据的Scanner，默认的分割函数是ScanLines
		scanner := bufio.NewScanner(reader)
		// 设置扫描时使用的初始缓冲区，以及扫描期间可能分配的最大缓冲区大小
		bf4 := make([]byte, 4096)
		scanner.Buffer(bf4, 128)
		// 设置该Scanner的分割函数，本方法必须在Scan之前调用
		/*
		   ScanBytes：会将每个字节作为一个token返回
		   ScanRunes：将每个utf-8编码的unicode码值作为一个token返回
		   ScanWords：会将空白（参见unicode.IsSpace）分隔的片段（去掉前后空白后）作为一个token返回
		   ScanLines：将每一行文本去掉末尾的换行标记作为一个token返回
		*/
		scanner.Split(bufio.ScanWords)
		// 扫描将scanner推进到下一个标记，然后通过字节或文本方法使用该标记，当扫描因为抵达输入流结尾或者遇到错误而停止时，返回false
		fmt.Println("推进是否成功：", scanner.Scan())
		// 返回最近一次Scan调用生成的token，会申请创建一个字符串保存token并返回该字符串
		fmt.Println("token：", scanner.Text())
		// 返回最近一次Scan调用生成的token
		fmt.Println("token：", string(scanner.Bytes()))
		// 返回Scanner遇到的第一个非EOF的错误
		if err := scanner.Err(); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}

}
