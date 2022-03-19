package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"time"
)

func main() {
	// TODO os包详解

	// 环境变量
	// 获取所有环境变量, 返回变量列表
	envs := os.Environ()
	for _, env := range envs {
		fmt.Println(env)
		/*cache := strings.Split(env, "=")
		fmt.Printf("key: %s value: %s\n", cache[0], cache[1])*/
	}
	// 获取指定的环境变量
	fmt.Println("GOROOT的路径为：", os.Getenv("GOROOT"))
	// 检索指定的环境变量 如果环境中存在变量，则返回值（可能为空），布尔值为真。否则，返回值将为空，布尔值将为false
	if envValue, ok := os.LookupEnv("GOROOT"); ok {
		fmt.Println("GOROOT的路径为：", envValue)
	}
	// 设置指定的环境变量
	//err := os.Setenv("GOROOT", "D:\\GoProject\\goStudy")
	// 设置指定的环境变量
	//err := os.Unsetenv("GOROOT")
	// 清除所有的环境变量
	//os.Clearenv()
	// 将当前工作目录（GOPATH）更改为dir目录
	//err := os.Chdir("D:\\GoProject")

	// 操作目录/文件
	// 打开方式
	/*const (
		//只读模式
		O_RDONLY int = syscall.O_RDONLY // open the file read-only.
		//只写模式
		O_WRONLY int = syscall.O_WRONLY // open the file write-only.
		//可读可写
		O_RDWR int = syscall.O_RDWR // open the file read-write.
		//追加内容
		O_APPEND int = syscall.O_APPEND // append data to the file when writing.
		//创建文件,如果文件不存在
		O_CREATE int = syscall.O_CREAT // create a new file if none exists.
		//与创建文件一同使用,文件必须存在
		O_EXCL int = syscall.O_EXCL // used with O_CREATE, file must not exist
		//打开一个同步的文件流
		O_SYNC int = syscall.O_SYNC // open for synchronous I/O.
		//如果可能,打开时缩短文件
		O_TRUNC int = syscall.O_TRUNC // if possible, truncate file when opened.
	)*/
	// 打开模式
	/*const (
		// 单字符是被String方法用于格式化的属性缩写。
		ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
		ModeAppend                                     // a: 只能写入，且只能写入到末尾
		ModeExclusive                                  // l: 用于执行
		ModeTemporary                                  // T: 临时文件（非备份文件）
		ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
		ModeDevice                                     // D: 设备
		ModeNamedPipe                                  // p: 命名管道（FIFO）
		ModeSocket                                     // S: Unix域socket
		ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
		ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
		ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
		ModeSticky                                     // t: 只有root/创建者能删除/移动文件
		// 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
		ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
		ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
	)*/
	// 文件信息
	/*type FileInfo interface {
		Name() string       // 文件的名字（不含扩展名）
		Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
		Mode() FileMode     // 文件的模式位
		ModTime() time.Time // 文件的修改时间
		IsDir() bool        // 等价于Mode().IsDir()
		Sys() interface{}   // 底层数据来源（可以返回nil）
	}*/
	// 获取文件信息对象，返回的FileInfo描述该符号链接指向的文件的信息，本函数会尝试跳转该链接
	fileInfo1, err := os.Stat(`D:\GoProject\goStudy\s\os.go`)
	if err == nil {
		fmt.Println("文件对象信息：", fileInfo1)
	} else {
		// 根据错误，判断文件或目录是否存在
		fmt.Println("是否存在：", os.IsExist(err))
		// 根据错误，判断是否为权限错误
		fmt.Println("是否为权限错误：", os.IsPermission(err))
	}
	// 获取文件信息对象，返回的FileInfo描述该符号链接指向的文件的信息，本函数不会试图跳转该链接
	fileInfo2, err := os.Lstat(`D:\GoProject\goStudy\s\os.go`)
	// 比较两个文件信息对象，是否指向同一目录
	fmt.Println("是否在同一目录：", os.SameFile(fileInfo1, fileInfo2))
	// 获取当前工作目录
	if getwd, err := os.Getwd(); err == nil {
		fmt.Println("工作目录：", getwd)
	}
	// 修改文件的权限
	err = os.Chmod("os.go", 755)
	// 更改文件拥有者
	err = os.Chown("os.go", 1025, 1026)
	// 更改文件拥有者。如果文件是一个符号链接，它改变的链接自己
	err = os.Lchown("os.go", 1025, 1026)
	// 修改文件的访问时间和修改时间
	err = os.Chtimes("os.go", time.Now(), time.Now())
	// 创建一个从原地址指向新地址的硬连接，对一个进行操作，则另外一个也会被修改
	err = os.Link("old.txt", "new.txt")
	// 创建一个从原地址指向新地址的软链接（符号链接），对一个进行操作，则另外一个也会被修改
	err = os.Symlink("old.txt", "new.txt")
	// 返回符号链接的目标文件
	if readlink, err := os.Readlink("new.txt"); err == nil {
		fmt.Println("符号链接的目标文件：", readlink)
	}
	// 读取指定的目录 返回[]DirEntry
	if readDir, err := os.ReadDir("new.txt"); err == nil {
		fmt.Println("目录对象：", readDir)
	}
	// 读取命名文件并返回内容
	if readFile, err := os.ReadFile("new.txt"); err == nil {
		fmt.Println("文件内容：", string(readFile))
	}
	// 创建目录
	err = os.Mkdir("test", os.ModeDir)
	// 递归创建目录
	err = os.MkdirAll("test", os.ModeDir)
	// 创建临时目录，如果dir是一个空字符串，则使用默认临时目录，目录名通过pattern+随机字符串生成
	tempDir, err := os.MkdirTemp("", "test")
	fmt.Println("临时文件路径：", tempDir)
	// 删除文件或目录
	err = os.Remove("test")
	// 递归删除文件或目录
	err = os.RemoveAll("test")
	// 文件重命名或移动路径
	err = os.Rename("test.txt", "test2.txt")      //重命名
	err = os.Rename("test.txt", "test/test2.txt") //移动路径
	// 修改文件大小
	err = os.Truncate("test.txt", 1024)
	// 创建文件（返回文件对象指针）
	file, err := os.Create("test.txt")
	fmt.Println("文件信息：", file)
	// 创建临时文件（返回文件对象指针），如果dir是一个空字符串，则使用默认临时目录，文件名通过pattern+随机字符串生成
	tempFile, err := os.CreateTemp("", "test")
	fmt.Println("临时文件信息：", tempFile)
	// 打开文件（返回文件对象指针），以只读的方式
	file, err = os.Open("test.txt")
	// 打开文件（返回文件对象指针），以指定的方式、指定的权限
	file, err = os.OpenFile("test.txt", os.O_WRONLY|os.O_APPEND, 0755)
	// 返回一对连接的文件，从r文件中读取数据，从w文件中写入数据
	if r, w, err := os.Pipe(); err == nil {
		fmt.Println("读取数据文件对象：", r, " 写入数据文件对象：", w)
	}

	// 文件对象
	// 获取文件路径
	fmt.Println("文件路径：", file.Name())
	// 获取文件信息对象
	fileInfo3, err := file.Stat()
	fmt.Println("文件信息对象：", fileInfo3)
	// 将当前工作路径修改为文件对象目录，文件对象必须为目录，该接口不支持window
	err = file.Chdir()
	// 修改文件模式
	err = file.Chmod(0755)
	// 修改文件大小
	err = file.Truncate(1024)
	// 读取文件内容
	bt := make([]byte, 1024)
	n, err := file.Read(bt)
	if err == io.EOF {
		fmt.Println("文件已读取完毕")
	}
	fmt.Println("读取的字节数：", n, " 读取的内容：", string(bt))
	// 从指定的位置读取文件内容，off代表字节数位置
	n, err = file.ReadAt(bt, 10)
	// Readdir读取与文件关联的目录的内容，并按目录顺序返回最多n个FileInfo值的片段，就像Lstat返回的那样。对同一文件的后续调用将产生更多的FileInfos
	// n代表最多读取多个文件对象返回
	if readdir, err := file.Readdir(10); err == nil {
		fmt.Println("指定文件同一目录下的文件对象：", readdir)
	}
	// ReadDir读取与文件f关联的目录的内容，并按目录顺序返回一段DirEntry值。对同一文件的后续调用将在目录中生成更晚的DirEntry记录
	// n代表最多读取多个目录对象返回
	if readDir, err := file.ReadDir(10); err == nil {
		fmt.Println("指定文件同一目录下的目录对象：", readDir)
	}
	// 读取并返回目录f里面的文件(或文件夹)的名称列表
	// n代表最多读取多个目录名称返回
	if readdirnames, err := file.Readdirnames(10); err == nil {
		fmt.Println("指定文件同一目录下的目录名称：", readdirnames)
	}
	// 写入内容
	n, err = file.Write(bt)
	// 写入字符串
	n, err = file.WriteString("abc")
	// 从指定的位置写入内容，off代表字节数位置
	n, err = file.WriteAt(bt, 10)
	fmt.Println("写入的字节数：", n)
	// 设置下一次读取或写入文件的偏移量，返回此处的偏移量
	/*const (
		SeekStart   = 0 // 0为相对文件开头
		SeekCurrent = 1 // 1为相对当前位置
		SeekEnd     = 2 // 2为相对文件结尾
	)*/
	seek, err := file.Seek(10, io.SeekStart)
	fmt.Println("此处的偏移量：", seek)
	// 关闭文件
	err = file.Close()
	// 将当前文件稳定的存储到磁盘中
	err = file.Sync()
	// 返回引用打开文件的Windows句柄
	fmt.Println("Windows句柄：", file.Fd())

	// 系统
	// 返回内核报告的主机名
	if hostname, err := os.Hostname(); err == nil {
		fmt.Println("主机名：", hostname)
	}
	// 让当前程序以给出的状态码 code 退出，一般来说，状态码 0 表示成功，非 0 表示出错。程序会立刻终止，并且 defer 的函数不会被执行
	//os.Exit(0)

	// 用户
	// 获取当前用户信息
	if current, err := user.Current(); err == nil {
		// 返回登录名
		fmt.Println("用户名：", current.Username)
		// 返回用户的真实名称或显示名称
		fmt.Println("用户名：", current.Name)
		// 返回用户id
		fmt.Println("uid：", current.Uid)
		// 返回所属组id
		fmt.Println("gid：", current.Gid)
		// 返回用户主目录的路径
		fmt.Println("用户主目录的路径：", current.HomeDir)
		// 返回用户所属的组ID列表
		if ids, err := current.GroupIds(); err == nil {
			fmt.Println("用户所属的组ID列表：", ids)
		}
	}
	// 创建用户实例
	userInfo := &user.User{
		Uid:      "1001",
		Gid:      "g01",
		Username: "lori",
		Name:     "lxz",
		HomeDir:  "usr/lxz",
	}
	// 按用户名查找用户
	if u, err := user.Lookup(userInfo.Username); err == nil {
		fmt.Println("用户：", u)
	}
	// 按用户id查找用户
	if u, err := user.LookupId(userInfo.Uid); err == nil {
		fmt.Println("用户：", u)
	}
	// 按组名称查找组
	if group, err := user.LookupGroup(userInfo.Gid); err != nil {
		fmt.Println("组：", group)
	}
	// 按组id查找组
	if group, err := user.LookupGroupId(userInfo.Gid); err != nil {
		fmt.Println("组：", group)
	}
	// 返回调用者的数字用户id
	fmt.Println("数字用户id", os.Getuid())
	// 返回调用者的数字有效用户id
	fmt.Println("数字有效用户id", os.Geteuid())
	// 返回调用者的数字组id
	fmt.Println("数字组id", os.Getgid())
	// 返回调用者的数字有效组id
	fmt.Println("数字有效组id", os.Getegid())
	// 返回调用者所属组 返回[]int
	if groups, err := os.Getgroups(); err == nil {
		fmt.Println("的数字ID列表", groups)
	}
	// 返回用于用户特定缓存数据的默认根目录
	if dir, err := os.UserCacheDir(); err == nil {
		fmt.Println("特定缓存数据的默认根目录：", dir)
	}
	// 返回用于用户特定配置数据的默认根目录
	if dir, err := os.UserConfigDir(); err != nil {
		fmt.Println("特定配置数据的默认根目录：", dir)
	}
	// 返回当前用户的主目录
	if dir, err := os.UserHomeDir(); err == nil {
		fmt.Println("当前用户的主目录：", dir)
	}

	// 进程
	/*
		type Process struct {
		    Pid int             //当前进程id
		    handle uintptr 　　  //处理指针
		    isdone uint32       //如果进程正在等待则该值非0，没有等待该值为0
		    sigMu  sync.RWMutex //锁
		}
	*/
	// StartProcess启动一个新的进程，其传入的name、argv和addr指定了程序、参数和属性
	// StartProcess是一个低层次的接口。os/exec包提供了高层次的接口
	//os.StartProcess()
	// 通过进程pid查找运行的进程，返回相关进程信息及在该过程中遇到的error
	if process, err := os.FindProcess(0); err == nil {
		fmt.Println("运行的进程：", process)
		// 发送一个信号给进程p, 在windows中没有实现发送中断interrupt
		//err := process.Signal()
		if err != nil {
			log.Fatalln(err)
		}
		// Wait等待进程退出，并返回进程状态和错误。Wait释放进程相关的资源。在大多数的系统上，进程p必须是当前进程的子进程，否则会返回一个错误
		if processState, err := process.Wait(); err == nil {
			fmt.Println("进程状态：", processState)
			// 判断判断程序是否已经退出
			fmt.Println("程序是否已经退出：", processState.Exited())
			// 判断程序是否成功退出，而Exited则仅仅只是判断其是否退出
			fmt.Println("程序是否成功退出：", processState.Success())
			// 返回有关进程的系统独立的退出信息。并将它转换为恰当的底层类型（比如Unix上的syscall.WaitStatus），主要用来获取进程退出相关资源
			fmt.Println("进程的系统独立的退出信息：", processState.Sys())
			// SysUsage返回关于退出进程的系统独立的资源使用信息。主要用来获取进程使用系统资源
			fmt.Println("关于退出进程的系统独立的资源使用信息：", processState.SysUsage())
			// 返回退出进程的内核态cpu使用时间
			fmt.Println("退出进程的内核态cpu使用时间：", processState.SystemTime())
			// 返回退出进程和子进程的用户态CPU使用时间
			fmt.Println("退出进程和子进程的用户态CPU使用时间：", processState.UserTime())
			// 返回退出进程的pid
			fmt.Println("进程的pid：", processState.Pid())
			// 返回进程状态的字符串表示
			fmt.Println("进程状态（字符串）：", processState.String())
		}
		// 杀死进程并直接退出
		err = process.Kill()
		if err != nil {
			log.Fatalln(err)
		}
	}

	// 执行外部命令
	// 在环境变量path中查找可执行二进制文件，返回完整路径或者相对于当前目录的一个相对路径
	if path, err := exec.LookPath("go"); err == nil {
		fmt.Println("相对路径：", path)
	}
	/*
		type Cmd struct {
			Path         string　　　//运行命令的路径，绝对路径或者相对路径
			Args         []string　　 // 命令参数
			Env          []string         //进程环境，如果环境为空，则使用当前进程的环境
			Dir          string　　　//指定command的工作目录，如果dir为空，则comman在调用进程所在当前目录中运行
			Stdin        io.Reader　　//标准输入，如果stdin是nil的话，进程从null device中读取（os.DevNull），stdin也可以时一个文件，否则的话则在运行过程中再开一个goroutine去
		　　　　　　　　　　　　　//读取标准输入
			Stdout       io.Writer       //标准输出
			Stderr       io.Writer　　//错误输出，如果这两个（Stdout和Stderr）为空的话，则command运行时将响应的文件描述符连接到os.DevNull
			ExtraFiles   []*os.File
			SysProcAttr  *syscall.SysProcAttr
			Process      *os.Process    //Process是底层进程，只启动一次
			ProcessState *os.ProcessState　　//ProcessState包含一个退出进程的信息，当进程调用Wait或者Run时便会产生该信息．
		}
	*/
	// command返回cmd结构来执行带有相关参数的命令，它仅仅设定cmd结构中的Path和Args参数
	cmd := exec.Command("ls")
	// CommandContext与Command类似，但包含一个上下文
	//cmd := exec.CommandContext(context.TODO(), "ls")
	// 返回人类可读的c语言描述
	fmt.Println("c语言描述：", cmd.String())
	// 运行命令，并返回标准输出和标准错误
	if output, err := cmd.CombinedOutput(); err == nil {
		fmt.Println("标准输出：", string(output))
	}
	// 运行命令，并返回其标准输出
	if output, err := cmd.Output(); err == nil {
		fmt.Println("标准输出：", string(output))
	}
	// 返回一个pipe，这个管道连接到command的标准错误，当command命令退出时，Wait将关闭这些pipe
	if pipe, err := cmd.StderrPipe(); err == nil {
		// 读取pipe
		bt := make([]byte, 1024)
		if n, err := pipe.Read(bt); err == nil {
			fmt.Println("读取内容：", string(bt[:n]))
		}
		// 关闭pipe
		if err := pipe.Close(); err == nil {
			return
		}
	}
	// 返回一个连接到command标准输入的管道pipe
	if pipe, err := cmd.StdinPipe(); err == nil {
		// 写入pipe
		bt := make([]byte, 1024)
		if n, err := pipe.Write(bt); err == nil {
			fmt.Println("写入内容：", string(bt[:n]))
		}
		// 关闭pipe
		if err := pipe.Close(); err == nil {
			return
		}
	}
	// 开始指定命令并且等待他执行结束，如果命令能够成功执行完毕，则返回nil，否则的话边会产生错误
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	// 使某个命令开始执行，但是并不等到他执行结束，这点和Run命令有区别．然后使用Wait方法等待命令执行完毕并且释放响应的资源
	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}
	// Wait等待command退出，他必须和Start一起使用，如果命令能够顺利执行完并顺利退出则返回nil，否则的话便会返回error，其中Wait会是放掉所有与cmd命令相关的资源
	if err := cmd.Wait(); err != nil {
		log.Fatalln(err)
	}
}
