package main

import (
	"archive/tar"
	"archive/zip"
	"compress/flate"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

func main() {
	// TODO archive包详解

	// zip 压缩文件
	/*
		压缩方法
		const (
		    Store   uint16 = 0 无压缩
		    Deflate uint16 = 8 压缩
		)*/
	// 为指定的方法ID注册自定义压缩器
	zip.RegisterCompressor(zip.Deflate, func(w io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(w, flate.BestCompression)
	})
	// 为指定的方法ID定制解压缩器
	zip.RegisterDecompressor(zip.Deflate, func(r io.Reader) io.ReadCloser {
		return flate.NewReader(r)
	})
	// 获取文件信息对象
	if stat, err := os.Stat("archive.go"); err == nil {
		// 通过文件信息，创建 zip.FileHeader 的文件信息
		if header, err := zip.FileInfoHeader(stat); err == nil {
			// 将tar.Header 转化为os.FileInfo
			fmt.Println("文件信息：", header.FileInfo())
			// 打开zip文件
			if stat, err := os.Open("archive.zip"); err == nil {
				// 创建一个新的zip.Writer
				writer := zip.NewWriter(stat)
				// 关闭打开的writer，同时写入数据。必须调用close方法，结束文件读写，close会向zip文件写入结束字段
				defer func(writer *zip.Writer) {
					err := writer.Close()
					if err != nil {
						log.Fatalln(err)
					}
				}(writer)
				// 设置中央目录注释字段的结尾
				err := writer.SetComment("Comments")
				if err != nil {
					log.Fatalln(err)
				}
				// 设置底层编写器中压缩数据开始的偏移量
				//writer.SetOffset(10)
				// 创建指定名称的文件添加到zip文件中
				if writer2, err := writer.Create("test"); err == nil {
					// 写入内容
					if n, err := writer2.Write([]byte("test")); err == nil {
						fmt.Println("写入长度：", n)
					}
				}
				// 使用提供的文件头将文件添加到zip存档中
				if writer3, err := writer.CreateRaw(header); err == nil {
					// 写入内容
					if n, err := writer3.Write([]byte("test")); err == nil {
						fmt.Println("写入长度：", n)
					}
				}
				// 将文件f（从读卡器获得）复制到w。它直接绕过解压缩、压缩和验证来复制原始表单
				//err := writer.Copy()
				// 将所有缓冲数据刷新到基础写入程序。通常不需要调用Flush；调用Close就足够了
				//err := writer.Flush()
				// 为特定的方法ID注册或覆盖自定义压缩器。如果找不到用于给定方法的压缩器，则Writer将默认在包级别查找该压缩器
				writer.RegisterCompressor(zip.Deflate, func(w io.Writer) (io.WriteCloser, error) {
					return flate.NewWriter(w, flate.BestCompression)
				})
				// 使用给出的*FileHeader来作为文件的元数据添加一个文件进zip文件
				if writer4, err := writer.CreateHeader(header); err == nil {
					// 写入内容
					if n, err := writer4.Write([]byte("test")); err == nil {
						fmt.Println("写入长度：", n)
					}
				}
				// 返回一个从r读取数据的*Reader，r被假设其大小为size字节
				if reader, err := zip.NewReader(stat, 1024); err == nil {
					// 返回一个io.ReadCloser接口，提供读取文件内容的方法。可以同时读取多个文件
					if file, err := reader.Open("zip.log"); err == nil {
						// 关闭文件
						defer func(file fs.File) {
							if err := file.Close(); err != nil {
								log.Fatalln(err)
							}
						}(file)
						// 读取文件内容
						bt := make([]byte, 1024)
						if n, err := file.Read(bt); err == nil {
							fmt.Println("读取长度：", n)
						}
					}
					// 注册或覆盖特定方法 ID 的自定义解压缩程序。如果找不到给定方法的解压缩程序，则 Reader 将默认在包级别查找解压缩程序
					reader.RegisterDecompressor(zip.Deflate, func(r io.Reader) io.ReadCloser {
						return flate.NewReader(r)
					})
				}
			}
		}
	}
	// 将打开由名称指定的Zip文件
	if reader, err := zip.OpenReader("archive.zip"); err == nil {
		// 使用io.fs的语义在ZIP存档中打开命名文件
		if file, err := reader.Open("zip.log"); err == nil {
			// 关闭文件
			defer func(file fs.File) {
				if err := file.Close(); err != nil {
					log.Fatalln(err)
				}
			}(file)
			// 读取文件内容
			bt := make([]byte, 1024)
			if n, err := file.Read(bt); err == nil {
				fmt.Println("读取长度：", n)
			}
		}
	}

	// tar 备份文件
	/*const (
	    // 类型
	    TypeReg           = '0'    // 普通文件
	    TypeRegA          = '\x00' // 普通文件
	    TypeLink          = '1'    // 硬链接
	    TypeSymlink       = '2'    // 符号链接
	    TypeChar          = '3'    // 字符设备节点
	    TypeBlock         = '4'    // 块设备节点
	    TypeDir           = '5'    // 目录
	    TypeFifo          = '6'    // 先进先出队列节点
	    TypeCont          = '7'    // 保留位
	    TypeXHeader       = 'x'    // 扩展头
	    TypeXGlobalHeader = 'g'    // 全局扩展头
	    TypeGNULongName   = 'L'    // 下一个文件记录有个长名字
	    TypeGNULongLink   = 'K'    // 下一个文件记录指向一个具有长名字的文件
	    TypeGNUSparse     = 'S'    // 稀疏文件
	)*/
	/*type Header struct {
	    Name       string    // 记录头域的文件名
	    Mode       int64     // 权限和模式位
	    Uid        int       // 所有者的用户ID
	    Gid        int       // 所有者的组ID
	    Size       int64     // 字节数（长度）
	    ModTime    time.Time // 修改时间
	    Typeflag   byte      // 记录头的类型
	    Linkname   string    // 链接的目标名
	    Uname      string    // 所有者的用户名
	    Gname      string    // 所有者的组名
	    Devmajor   int64     // 字符设备或块设备的major number
	    Devminor   int64     // 字符设备或块设备的minor number
	    AccessTime time.Time // 访问时间
	    ChangeTime time.Time- // 状态改变时间
	    Xattrs     map[string]string
	}*/
	// 获取文件信息对象
	if stat, err := os.Stat("archive.go"); err == nil {
		// 将 FileInfo 实例转化为tar.Header
		// 当文件为软连接文件的时候，link字段为源文件的全路径
		if header, err := tar.FileInfoHeader(stat, ""); err == nil {
			// 将tar.Header 转化为os.FileInfo
			fmt.Println("文件信息：", header.FileInfo())
			// 打开tar文件
			if stat, err := os.Open("archive.tar"); err == nil {
				// 创建一个新的tar.Writer
				writer := tar.NewWriter(stat)
				// 关闭打开的writer，同时写入数据。必须调用close方法，结束文件读写，close会向tar文件写入结束字段
				defer func(writer *tar.Writer) {
					err := writer.Close()
					if err != nil {
						log.Fatalln(err)
					}
				}(writer)
				// 将文件头部信息header写入当前的tar文件
				err = writer.WriteHeader(header)
				if err != nil {
					log.Fatalln(err)
				}
				// 将数据写入当前存档文件。如果在writeHeader之后写入超过Header.Size字节，则Write返回ErrWriteTooLong错误
				if n, err := writer.Write([]byte("test")); err == nil {
					fmt.Println("写入长度：", n)
				}
				// 将数据写入tar文件
				err = writer.Flush()
				if err != nil {
					log.Fatalln(err)
				}
				// Reader提供对tar文件的顺序访问
				reader := tar.NewReader(stat)
				// 从文件中读取数据，到达记录末端时返回(0, EOF)，直到调用Next方法转入下一记录
				bt := make([]byte, 1024)
				if n, err := reader.Read(bt); err == nil {
					fmt.Println("读取的字节数：", n, " 读取的内容：", string(bt))
				} else if err == io.EOF {
					fmt.Println("文件已读取完毕")
				}
				// 读取下一个文件Header
				if header, err := reader.Next(); err == nil {
					fmt.Println("下一个文件Header：", header)
				}
			}
		}
	}
}
