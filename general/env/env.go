package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload" //自动加载，引入此文件无需再使用Load、OverLoad进行手动加载
)

func main() {
	//TODO 使用joho/godotenv操作
	// go get -u github.com/joho/godotenv
	//默认数据格式为：S3_BUCKET=YOURS3BUCKET，还可以使用 YAML 格式：name: awesome web

	// 从指定的文件加载环境变量（不会覆盖已经存在的env变量）
	if err := godotenv.Load("./s/test1.env", "./s/test2.env"); err != nil {
		log.Fatalln(err)
	}

	// 从指定的文件加载环境变量（将覆盖已经存在的env变量）
	if err := godotenv.Overload("./s/test1.env", "./s/test2.env"); err != nil {
		log.Fatalln(err)
	}

	// 从指定的文件获取所有的环境变量数据，返回map[string]string
	envMap, err := godotenv.Read("./s/test1.env", "./s/test2.env")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("环境变量map：", envMap)

	// 将map环境变量写入到指定路径的文件中
	envMap["TEST"] = "val"
	if err = godotenv.Write(envMap, "./s/test3.env"); err != nil {
		log.Fatalln(err)
	}

	// 将map环境变量转成dotenv格式，每行的格式为：KEY=“VALUE” 返回string
	envString, err := godotenv.Marshal(envMap)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("转成string：", envString)

	// 将string转成map环境变量格式 返回map[string]string
	envMap, err = godotenv.Unmarshal(envString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("转成map：", envMap)

	// 打开文件（返回文件对象指针），以只读的方式
	reader, err := os.OpenFile("./s/test1.env", os.O_RDONLY, 0755)
	if err == nil {
		// 从io.Reader中读取环境变量 返回map[string]string
		if envMap, err = godotenv.Parse(reader); err != nil {
			log.Fatalln(err)
		} else {
			fmt.Println("环境变量map：", envMap)
		}
	}

	// 从指定的文件加载环境变量并执行指定的cmd命令
	filenames := []string{"./s/test1.env", "./s/test2.env"}
	if err = godotenv.Exec(filenames, "/bin/bash", []string{"-c", "sh /test.sh"}); err != nil {
		log.Fatalln(err)
	}

	// 加载环境变量后还可调用os.env包的方法来调用其环境变量 更多详情请查看os.env包
	fmt.Println("环境变量S3_BUCKET的值：", os.Getenv("S3_BUCKET"))
}
