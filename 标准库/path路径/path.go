package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {
	//TODO path包详解
	// path包只能用于由正斜杠分隔的路径，例如 URL 中的路径。此软件包不处理带有驱动器号或反斜杠的 Windows 路径
	// filepath包兼容各操作系统的文件路径

	// 测试地址
	testPath := "D:\\GoProject\\goStudy\\s"

	// 返回路径的最后一个元素  如果路径为空字符串，返回. 如果路径只有斜线，返回/
	fmt.Println("路径的最后一个元素：", path.Base(testPath))
	fmt.Println("路径的最后一个元素：", filepath.Base(testPath))

	// 返回路径最后一个元素的目录 如果路径为空则返回.
	fmt.Println("路径的最后一个元素的目录：", path.Dir(testPath))
	fmt.Println("路径的最后一个元素的目录：", filepath.Dir(testPath))

	// 返回路径中的扩展名 如果没有点，返回空
	fmt.Println("路径的最后一个元素的目录：", path.Ext(testPath))
	fmt.Println("路径的最后一个元素的目录：", filepath.Ext(testPath))

	// 判断路径是否为绝对路径
	fmt.Println("路径是否为绝对路径：", path.IsAbs(testPath))
	fmt.Println("路径是否为绝对路径：", filepath.IsAbs(testPath))

	// 连接路径，返回已经clean处理后的路径
	fmt.Println("连接路径：", path.Join(testPath, "request.log"))
	fmt.Println("连接路径：", filepath.Join(testPath, "request.log"))

	// 分割路径中的目录与文件
	dir, file := path.Split(testPath)
	//dir, file := filepath.Split(testPath)
	fmt.Println("分割后的目录：", dir, " 文件：", file)

	// 根据pattern进行文件名匹配，完全匹配则返回true
	// 可用的匹配字符如下:
	// '*'                                  匹配0或多个非/的字符
	// '?'                                  匹配1个非/的字符
	// '[' [ '^' ] { character-range } ']'  字符组（必须非空）(支持三种格式[abc],[^abc],[a-c])
	// c                                    匹配字符c（c != '*', '?', '\\', '['）
	// '\\' c                               匹配字符c(可上面c的区别是 可以支持字符 * ? \\ [的匹配)
	matched, err := path.Match("*", "fileName")
	//matched, err := filepath.Match("*", "fileName")
	if err != nil {
		// 表示匹配模式格式不正确
		//filepath.ErrBadPattern
		if err == path.ErrBadPattern {
			log.Fatalln(err)
		}
		panic(err)
	}
	fmt.Println("是否匹配成功：", matched)

	// 返回与指定路径等效的最短路径（通过）
	// 此函数迭代地应用以下规则，直到无法进行进一步处理为止：
	// 将连续的多个斜杠替换为单个斜杠
	// 剔除每一个 . 路径名元素（代表当前目录）
	// 剔除每一个路径内的 .. 路径名元素（代表父目录）和它前面的非 .. 路径名元素
	// 剔除开始一个根路径的 .. 路径名元素，即将路径开始处的 /.. 替换为 /
	fmt.Println("最短路径：", path.Clean("..srv/./../local"))
	fmt.Println("最短路径：", filepath.Clean("..srv/./../local"))

	// 返回文件的实际路径
	symlinks, err := filepath.EvalSymlinks(testPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("实际路径：", symlinks)

	// 将路径中的/替换为路径分隔符（\）
	fmt.Println("替换后路径：", filepath.FromSlash(testPath))

	//将路径分隔符使用/替换
	fmt.Println("替换后路径：", filepath.ToSlash(testPath))

	//返回所有匹配的文件 返回[]string（所有匹配到的文件路径）
	match, err := filepath.Glob(filepath.Join(testPath, "*.go"))
	if err != nil {
		panic(err)
	}
	fmt.Println("匹配到的文件：", match)

	// 返回以basePath为基准的相对路径
	path2, err := filepath.Rel("C:/a/b", "C:/a/b/c/d/../e")
	if err != nil {
		panic(err)
	}
	fmt.Println("相对经：", path2)

	// 将路径使用路径列表分隔符分开，见os.PathListSeparator 返回[]string（所有分开到的文件路径）
	// linux下默认为:，windows下为;
	fmt.Println(filepath.SplitList("C:/windows;C:/windows/system"))

	// 返回分区名
	fmt.Println("分区名：", filepath.VolumeName(testPath))

	// 遍历指定目录下所有文件
	err = filepath.Walk(testPath, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	// 遍历指定目录下所有目录
	//filepath.WalkDir()
}
