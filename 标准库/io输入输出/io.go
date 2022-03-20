package main

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// TODO io包详解

	// 文件操作
	// 读取到结尾错误
	fmt.Println("读取到结尾错误：", io.EOF)

	// 管道关闭的错误
	fmt.Println("读取到结尾错误：", io.ErrClosedPipe)

	// 多次调用reader都没有返回数据，一般来说就是封装的reader有问题
	fmt.Println("读取到结尾错误：", io.ErrNoProgress)

	// 缓冲池不够大
	fmt.Println("读取到结尾错误：", io.ErrShortBuffer)

	// 写的数据比提供的少
	fmt.Println("读取到结尾错误：", io.ErrShortWrite)

	// 这个表示在没有读取对应的数据时候，没有读取完成返回的错误
	fmt.Println("读取到结尾错误：", io.ErrUnexpectedEOF)

	// io.Reader接口定义了 Read 方法，用于读取数据到字节数组中

	// io.Writer接口定义了 Write 方法，用于写数据到文件中

	// io.Closer接口定义了 Close 方法，该方法用于关闭连接

	// io.Seeker接口定义了 Seek 方法，该方法用于指定下次读取或者写入时的偏移量

	// ioutil  常用、方便的IO操作函数
	// Discard 是一个 io.Writer 接口，调用它的 Write 方法将不做任何事情
	//discard := ioutil.Discard
	if r, err := os.Open("io.log"); err != nil {
		// ReadAll从 r 读取，直到出错或 EOF 并返回它读取的数据
		if content, err := ioutil.ReadAll(r); err == nil {
			fmt.Println("读取内容：", string(content))
		}
		// NopCloser 将 r 包装为一个 ReadCloser 类型，但 Close 方法不做任何事情
		ioutil.NopCloser(r)
	}
	// ReadFile 读取文件中的所有数据，返回读取的数据和遇到的错误
	if content, err := ioutil.ReadFile(`D:\GoProject\goStudy\s\io.go`); err == nil {
		fmt.Println("读取内容：", string(content))
	}
	// ReadDir 读取指定目录中的所有目录和文件（不包括子目录） 返回fs.FileInfo切片
	if dir, err := ioutil.ReadDir(`D:\GoProject\goStudy\s`); err == nil {
		fmt.Println("读取目录和文件：", dir)
	}
	// WriteFile 向文件中写入数据，写入前会清空文件
	if err := ioutil.WriteFile(`D:\GoProject\goStudy\s\io.go`, []byte("test"), 0755); err != nil {
		log.Fatalln(err)
	}
	// TempDir在dir目录中创建一个新的临时目录
	if dirName, err := ioutil.TempDir(`D:\GoProject\goStudy\s`, "testTemp"); err == nil {
		fmt.Println("临时目录名：", dirName)
	}
	// TempFile在dir目录中创建一个新的临时文件，并将其以读写模式打开
	if file, err := ioutil.TempFile(`D:\GoProject\goStudy\s`, "testTemp.log"); err == nil {
		fmt.Println("临时文件对象：", file)
	}

	// fs 统一标准库文件io相关的访问方式
	// 返回一个文件系统（fs.fs），用于目录dir下的文件树
	dirFs := os.DirFS(`D:\GoProject\goStudy\s`)
	// 返回一个FS，该FS对应于根在fsys目录下的子树
	if subtreeFs, err := fs.Sub(dirFs, "testDir"); err != nil {
		fmt.Println("根sys下的子树sys：", subtreeFs)
	}
	// 获取文件信息对象，返回的FileInfo描述该符号链接指向的文件的信息，本函数会尝试跳转该链接
	if fileInfo1, err := fs.Stat(dirFs, "io.go"); err == nil {
		fmt.Println("文件对象信息：", fileInfo1)
		// y返回一个DirEntry，该DirEntry从fileInfo返回信息
		fmt.Println("该文件的目录对象：", fs.FileInfoToDirEntry(fileInfo1))
	}
	// 返回与pattern或nil匹配的所有文件的名称
	if fileNames, err := fs.Glob(dirFs, "log"); err != nil {
		fmt.Println("匹配到的所有文件名称：", fileNames)
	}
	// 读取命名目录并返回按文件名排序的目录项列表，返回[]DirEntry
	if dirs, err := fs.ReadDir(dirFs, "testDir"); err != nil {
		fmt.Println("目录对象：", dirs)
	}
	// 从文件系统fs读取命名文件并返回其内容
	if content, err := fs.ReadFile(dirFs, "testFile.log"); err == nil {
		fmt.Println("读取的内容：", string(content))
	}
	// 遍历根目录下的文件树（目录），为每个文件或文件夹调用fn方法
	if err := fs.WalkDir(dirFs, `D:\GoProject\goStudy\s`, func(path string, d fs.DirEntry, err error) error {
		if err == nil {
			fmt.Println("当前目录路径：", path)
			fmt.Println("当前目录对象：", d)
		}
		return nil
	}); err != nil {
		log.Fatalln(err)
	}

}
