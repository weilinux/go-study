package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

type Config struct {
	AppName  string `ini:"app_name"`
	LogLevel string `ini:"log_level"`

	MySQL *MySQLConfig `ini:"mysql"`
	Redis *RedisConfig `ini:"redis"`
}

type MySQLConfig struct {
	IP       string `ini:"ip"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type RedisConfig struct {
	IP   string `ini:"ip"`
	Port int    `ini:"port"`
}

func main() {
	//TODO 使用go-ini/ini操作
	// go get -u gopkg.in/ini.v1

	// 自定义配置
	options := ini.LoadOptions{ // LoadOptions包含用于加载数据源的所有自定义选项
		Loose:                       false, // 是否应该忽略不存在的文件还是返回错误
		Insensitive:                 false, // 是否强制所有分区和键名称为小写
		InsensitiveSections:         false, // 是否强制所有分区为小写
		InsensitiveKeys:             false, // 是否强制所有键名为小写
		IgnoreContinuation:          false, // 在解析时是否忽略延续行 (换行符\)
		IgnoreInlineComment:         false, // 是否忽略值末尾的注释，并将其视为值的一部分
		SkipUnrecognizableLines:     false, // 是否跳过不符合键/值对的不可识别行
		ShortCircuit:                false, // 加载第一个可用配置源后是否忽略其他配置源
		AllowBooleanKeys:            false, // 是否允许bool类型的键（允许则值为true），或者是否将其视为缺少值
		AllowShadows:                true,  // 是否跟踪同一分区下具有相同名称的键
		AllowNestedValues:           true,  // 是否允许类似AWS的嵌套值
		AllowPythonMultilineValues:  false, // 是否允许类似Python的多行值
		SpaceBeforeInlineComment:    false, // 是否允许使用注释符号（\#和\）出现在值的内部
		UnescapeValueDoubleQuotes:   false, // 是否将值内的双引号取消为常规格式
		UnescapeValueCommentSymbols: false, // 使用unescape注释符号（\#和\）将内部值转换为常规格式
		UnparseableSections:         nil,   // 存储一个允许包含原始内容的块列表，否则不允许包含原始内容
		KeyValueDelimiters:          "",    // 用于分隔键和值的分隔符序列 默认为"=:"
		KeyValueDelimiterOnWrite:    "",    // 用于分隔键和值输出的分隔符 默认为"="
		ChildSectionDelimiter:       "",    // 用于分隔子节的分隔符 默认为"."
		PreserveSurroundedQuote:     false, // 是否保留包围的引号（单引号和双引号）
		DebugFunc:                   nil,   // 调用DebugFunc是为了收集调试信息（目前仅对调试解析Python样式的多行值有用）
		ReaderBufferSize:            0,     // 读卡器的缓冲区大小，以字节为单位
		AllowNonUniqueSections:      false, // 是否允许多次使用相同名称的分区
		AllowDuplicateShadowValues:  false, // 是否应删除阴影关键点的值
	}

	// TODO 加载配置
	// 从.ini数据源加载和解析
	//cfg, err := ini.Load("D:\\GoProject\\goStudy\\s\\my.ini")
	// 应用自定义选项从数据源加载
	cfg, err := ini.LoadSources(options, "D:\\GoProject\\goStudy\\s\\my.ini")
	// 忽略不存在的文件，而不是返回错误
	//cfg, err := ini.LooseLoad("D:\\GoProject\\goStudy\\s\\my.ini")
	// 强制所有节和键名称都使用小写
	//cfg, err := ini.InsensitiveLoad("D:\\GoProject\\goStudy\\s\\my.ini")
	// 允许有阴影键
	//cfg, err := ini.ShadowLoad("D:\\GoProject\\goStudy\\s\\my.ini")
	// 创建一个空文件对象（后续需使用append追加需要加入的数据源）
	//cfg := ini.Empty(options)
	// 将文件映射到给定的结构体中
	c := &Config{MySQL: &MySQLConfig{}, Redis: &RedisConfig{}}
	if err = ini.MapTo(c, "D:\\GoProject\\goStudy\\s\\my.ini"); err != nil {
		log.Fatalln(err)
	}
	// 使用名称的映射函数（func(string) string）将文件映射到给定的结构体中
	if err = ini.MapToWithMapper(c, cfg.NameMapper, "D:\\GoProject\\goStudy\\s\\my.ini"); err != nil {
		log.Fatalln(err)
	}
	// 以严格模式将文件映射到给定的结构体中
	if err = ini.StrictMapTo(c, "D:\\GoProject\\goStudy\\s\\my.ini"); err != nil {
		log.Fatalln(err)
	}
	// 使用名称的映射函数（func(string) string）在严格模式将文件映射到给定的结构体中
	if err = ini.StrictMapToWithMapper(c, cfg.NameMapper, "D:\\GoProject\\goStudy\\s\\my.ini"); err != nil {
		log.Fatalln(err)
	}
	// 追加一个结构体数据源（反射结构体数据将其保存进配置对象中）
	config := &Config{
		AppName:  "awesome web",
		LogLevel: "DEBUG",
		MySQL: &MySQLConfig{
			IP:       "127.0.0.1",
			Port:     3306,
			User:     "root",
			Password: "123456",
			Database: "awesome",
		},
		Redis: &RedisConfig{
			IP:   "127.0.0.1",
			Port: 6381,
		},
	}
	if err = ini.ReflectFrom(cfg, config); err != nil {
		log.Fatalln(err)
	}
	// 使用名称的映射函数（func(string) string）追加一个结构体数据源（反射结构体数据将其保存进配置对象中）
	if err = ini.ReflectFromWithMapper(cfg, config, cfg.NameMapper); err != nil {
		log.Fatalln(err)
	}
	// 设置为安全模式（读写锁）
	cfg.BlockMode = true
	// 名称的映射函数（操作名称会自动触发该函数）两种内置名称映射器函数
	// ini.SnackCase：转换为大小写格式
	// ini.TitleUnderscore：转换为标题下划线格式
	cfg.NameMapper = ini.SnackCase
	// 键值的映射函数（操作键值会自动触发该函数）无内置键值映射器函数，但可使用任意func(string) string类型的方法作为该函数的处理函数
	// os.ExpandEnv：替换字符串中的${var}或者$var
	cfg.ValueMapper = os.ExpandEnv
	// 重新加载并解析所有数据源
	if err = cfg.Reload(); err != nil {
		log.Fatalln(err)
	}
	// 追加一个或多个数据源并自动重新加载
	if err = cfg.Append("D:\\GoProject\\goStudy\\s\\app.ini"); err != nil {
		log.Fatalln(err)
	}
	// 追加一个结构体数据源（反射结构体数据将其保存进配置对象中）
	if err = cfg.ReflectFrom(config); err != nil {
		log.Fatalln(err)
	}
	// 将文件内容保存到指定的.ini文件中
	if err = cfg.SaveTo("./s/test1.ini"); err != nil {
		log.Fatalln(err)
	}
	// 以给定的缩进方式将文件内容保存到指定的.ini文件中
	if err = cfg.SaveToIndent("./s/test2.ini", "\t"); err != nil {
		log.Fatalln(err)
	}
	// 打开文件（返回文件对象指针），以只写的方式
	write1, err := os.OpenFile("./s/test1.txt", os.O_WRONLY, 0755)
	if err == nil {
		// 将文件内容写入到实现了io.writer方法的对象中
		if n, err := cfg.WriteTo(write1); err != nil {
			log.Fatalln(err)
		} else {
			fmt.Println("写入字符长度：", n)
		}
	}
	// 打开文件（返回文件对象指针），以只写的方式
	write2, err := os.OpenFile("./s/test2.txt", os.O_WRONLY, 0755)
	if err == nil {
		// 以给定的缩进方式将文件内容写入到实现了io.writer方法的对象中
		if n, err := cfg.WriteToIndent(write2, "\t"); err != nil {
			log.Fatalln(err)
		} else {
			fmt.Println("写入字符长度：", n)
		}
	}
	// 将文件映射到给定的结构体中
	if err = cfg.MapTo(c); err != nil {
		log.Fatalln(err)
	}
	// 以严格模式将文件映射到给定的结构体中
	if err = cfg.StrictMapTo(c); err != nil {
		log.Fatalln(err)
	}
	// 判断错误是否为ini.ErrDelimiterNotFound（指未找到应存在的分隔符的错误类型）
	fmt.Println("是否为指定错误类型：", ini.IsErrDelimiterNotFound(err))

	// TODO 操作分区
	// 创建一个新分区
	//section, err := cfg.NewSection("test")
	// 创建一个具有不可拆分实体的新分区
	//section, err := cfg.NewRawSection("test1", "body")
	// 创建一个区分列表
	//err := cfg.NewSections("test2", "test3")
	// 返回指定名称的区分，不存在则返回err
	//section, err := cfg.GetSection("")
	// 返回指定名称的区分，不存在则返回nil（抑制错误）
	section := cfg.Section("") // 默认分区的名字为""，也可以使用ini.DefaultSection
	// 返回当前实例中的分区列表 []*Section
	//sections := cfg.Sections()
	// 返回指定名称的所有分区列表
	//sections, err := cfg.SectionsByName("server")
	// 返回指定名称和索引位的分区，如果不存在则返回一个新分区
	//section := cfg.SectionWithIndex("server", 1)
	// 返回指定名称的子区分列表
	//sections := cfg.ChildSections("child1")
	// 获取所有分区名 默认分区名为：DEFAULT
	fmt.Println("所有分区名：", cfg.SectionStrings())
	// 判断文件是否包含指定的分区
	fmt.Println("指定分区是否存在：", cfg.HasSection("test"))
	// 删除指定名称的分区
	cfg.DeleteSection("test")
	// 删除指定名称和索引位的分区
	if err = cfg.DeleteSectionWithIndex("test", 0); err != nil {
		log.Fatalln(err)
	}

	// TODO 操作键
	// 获取分区下的键，不存在则返回err
	key, err := section.GetKey("app_mode")
	// 获取分区下的键，不存在则返回nil（抑制错误）
	//key := section.Key("app_mode")
	// 返回分区的键列表 类型：[]*Key
	fmt.Println("分区的键列表：", section.Keys())
	// 返回分区的map键值对列表 类型：map[string]string
	fmt.Println("分区的map键值对列表：", section.KeysHash())
	// 返回分区的键名称列表 类型：[]string
	fmt.Println("分区的键名称列表：", section.KeyStrings())
	// 从分区中删除一个指定名称的键
	section.DeleteKey("app_name")
	// 判断分区是否包含指定的键
	fmt.Println("指定键是否存在：", section.HasKey("app_mode"))
	// 为分区创建一个bool类型的新键
	if key2, err := section.NewBooleanKey("bool_key"); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("键指针地址：", key2)
	}
	// 为分区创建一个bool类型的新键
	if key3, err := section.NewKey("string_key", "string_val"); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("键指针地址：", key3)
	}
	// 返回父级分区的键列表 类型：[]*Key
	fmt.Println("父级分区的键列表：", section.ParentKeys())

	// TODO 操作值
	// 获取键的注释信息：
	fmt.Println("键的注释信息：", key.Comment)
	// 修改键的值
	key.SetValue("development test")
	// 获取键的名称
	fmt.Println("键的名称：", key.Name())
	// 获取键的值
	fmt.Println("键的值：", key.Value())
	// 向键添加嵌套值
	if err = key.AddNestedValue("test value"); err != nil {
		log.Fatalln(err)
	}
	// 获取键的嵌套值
	fmt.Println("键的嵌套值：", key.NestedValues())
	// 向自身添加新的阴影关键点
	if err = key.AddShadow("test shadow"); err != nil {
		log.Fatalln(err)
	}
	// 获取键的值及其阴影（如果有）的原始值
	fmt.Println("键的值及其阴影（如果有）的原始值：", key.ValueWithShadows())
	// 其他转换值类型、反射结构体、抑制错误的转换类型、检查值是否在给定的范围中、已分隔符分割值到切片中等等方法请查看：https://ini.unknwon.io/docs/howto/work_with_values
}
