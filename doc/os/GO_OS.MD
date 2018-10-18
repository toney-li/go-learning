## OS包操作相关

### 权限相关

[权限相关内容](./auth/GO_AUTH.MD)

### 文件操作

### 系统相关

``` go
func Chdir(dir string) error   //chdir将当前工作目录更改为dir目录．

func Getwd() (dir string, err error)    //获取当前目录，类似linux中的pwd
func Chmod(name string, mode FileMode) error     //更改文件的权限（读写执行，分为三类：all-group-owner）
func Chown(name string, uid, gid int) error  //更改文件拥有者owner
func Chtimes(name string, atime time.Time, mtime time.Time) error    //更改文件的访问时间和修改时间，atime表示访问时间，mtime表示更改时间
func Clearenv()    //清除所有环境变量（慎用）
func Environ() []string  //返回所有环境变量
func Exit(code int)     //系统退出，并返回code，其中０表示执行成功并退出，非０表示错误并退出，其中执行Exit后程序会直接退出，defer函数不会执行．

func Expand(s string, mapping func(string) string) string   //Expand用mapping 函数指定的规则替换字符串中的${var}或者$var（注：变量之前必须有$符号）。比如，os.ExpandEnv(s)等效于os.Expand(s, os.Getenv)。
func ExpandEnv(s string) string  //ExpandEnv根据当前环境变量的值来替换字符串中的${var}或者$var。如果引用变量没有定义，则用空字符串替换。
func Getgroups() ([]int, error)  //返回调用者属于的group，其和chown配合使用，改变文件属于的group．


func Getpagesize() int　　　//获取底层系统内存页的数量
func Getpid() int　　　　//获取进程id
func Getppid() int             //获取调用者进程父id

func Hostname() (name string, err error)    //获取主机名
func IsExist(err error) bool    　　　　　 //返回一个布尔值，它指明err错误是否报告了一个文件或者目录已经存在。它被ErrExist和其它系统调用满足。

func IsNotExist(err error) bool　　　　　//返回一个布尔值，它指明err错误是否报告了一个文件或者目录不存在。它被ErrNotExist 和其它系统调用满足。
func IsPathSeparator(c uint8) bool         //判断c是否是一个路径分割符号，是的话返回true,否则返回false
func IsPermission(err error) bool   //判定err错误是否是权限错误。它被ErrPermission 和其它系统调用满足。
func Lchown(name string, uid, gid int) error　　　//改变了文件的gid和uid。如果文件是一个符号链接，它改变的链接自己。如果出错，则会是*PathError类型。
func Link(oldname, newname string) error       //创建一个从oldname指向newname的硬连接，对一个进行操作，则另外一个也会被修改．
func Mkdir(name string, perm FileMode) error　//创建一个新目录，该目录具有FileMode权限，==当创建一个已经存在的目录时会报错==

func MkdirAll(path string, perm FileMode) error　//创建一个新目录，该目录是利用路径（包括绝对路径和相对路径）进行创建的，如果需要创建对应的父目录，也一起进行创建，如果已经有了该目录，则不进行新的创建，当创建一个已经存在的目录时，不会报错.
func NewSyscallError(syscall string, err error) error    //NewSyscallError返回一个SyscallError 错误，带有给出的系统调用名字和详细的错误信息。也就是说，如果err为空，则返回空

func Readlink(name string) (string, error)         //返回符号链接的目标。如果出错，将会是 *PathError类型。
func Remove(name string) error           //删除文件或者目录
func RemoveAll(path string) error　　//删除目录以及其子目录和文件，如果path不存在的话，返回nil

func Rename(oldpath, newpath string) error　　//重命名文件，如果oldpath不存在，则报错no such file or directory
func SameFile(fi1, fi2 FileInfo) bool　　　　　　//查看f1和f2这两个是否是同一个文件，如果再Unix系统，这意味着底层结构的device和inode完全一致，在其他系统上可能是基于文件绝对路径的．SameFile只适用于本文件包stat返回的状态，其他情况下都返回false
func Setenv(key, value string) error           //设定环境变量，经常与Getenv连用，用来设定环境变量的值

func Symlink(oldname, newname string) error　　　//创建一个newname作为oldname的符号连接，这是一个符号连接（Link是硬连接），与link的硬连接不同，利用Link创建的硬连接，则newname和oldname的file互不影响，一个文件删除，另外一个文件不受影响；但是利用SymLink创建的符号连接，其newname只是一个指向oldname文件的符号连接，当oldname　file删除之后，则newname的文件也就不能够继续使用．
func TempDir() string　　　　　　　　//创建临时目录用来存放临时文件，这个临时目录一般为/tmp
func Truncate(name string, size int64) error     //按照指定长度size将文件截断，如果这个文件是一个符号链接，则同时也改变其目标连接的长度，如果有错误，则返回一个错误．

func Create(name string) (file *File, err error)　　//创建一个文件，文件mode是0666(读写权限)，如果文件已经存在，则重新创建一个，原文件被覆盖，创建的新文件具有读写权限，能够备用与i/o操作．其相当于OpenFile的快捷操作，等同于OpenFile(name string, O_CREATE,0666)
func NewFile(fd uintptr, name string) *File　　　//根据文件描述符和名字创建一个新的文件
func OpenFile(name string, flag int, perm FileMode) (file *File, err error)　//指定文件权限和打开方式打开name文件或者create文件，其中flag标志如下:

打开标记：

O_RDONLY：只读模式(read-only)
O_WRONLY：只写模式(write-only)
O_RDWR：读写模式(read-write)
O_APPEND：追加模式(append)
O_CREATE：文件不存在就创建(create a new file if none exists.)
O_EXCL：与 O_CREATE 一起用，构成一个新建文件的功能，它要求文件必须不存在(used with O_CREATE, file must not exist)
O_SYNC：同步方式打开，即不使用缓存，直接写入硬盘
O_TRUNC：打开并清空文件
至于操作权限perm，除非创建文件时才需要指定，不需要创建新文件时可以将其设定为０.虽然go语言给perm权限设定了很多的常量，但是习惯上也可以直接使用数字，如0666(具体含义和Unix系统的一致).
func Pipe() (r *File, w *File, err error)        //返回一对连接的文件，从r中读取写入w中的数据，即首先向w中写入数据，此时从r中变能够读取到写入w中的数据，Pipe()函数返回文件和该过程中产生的错误．
func (f *File) Chdir() error　　　　　　　//改变工作目录到f，其中f必须为一个目录，否则便会报错
func (f *File) Chmod(mode FileMode) error　　　//更改文件权限，其等价与os.Chmod(name string,mode FileMode)error
func (f *File) Chown(uid, gid int) error                     //更改文件所有者，与os.Chown等价
func (f *File) Close() error　　　　　　　　　　//关闭文件，使其不能够再进行i/o操作，其经常和defer一起使用，用在创建或者打开某个文件之后，这样在程序退出前变能够自己关闭响应的已经打开的文件．
func (f *File) Fd() uintptr　　　//返回系统文件描述符，也叫做文件句柄．
func (f *File) Name() string　　//返回文件名字，与file.Stat().Name()等价
func (f *File) Read(b []byte) (n int, err error)　//将len(b)的数据从f中读取到b中，如果无错，则返回n和nil,否则返回读取的字节数n以及响应的错误
func (f *File) ReadAt(b []byte, off int64) (n int, err error)　//和Read类似，不过ReadAt指定开始读取的位置offset
func (f *File) Readdir(n int) (fi []FileInfo, err error)            
Readdir读取file指定的目录的内容，然后返回一个切片，它最多包含n个FileInfo值，这些值可能是按照目录顺序的Lstat返回的。接下来调用相同的文件会产生更多的FileInfos。

如果n>0，Readdir返回最多n个FileInfo结构。在这种情况下，如果Readdir返回一个空的切片，它将会返回一个非空的错误来解释原因。在目录的结尾，错误将会是io.EOF。

如果n<=0，Readdir返回目录的所有的FileInfo，用一个切片表示。在这种情况下，如果Readdir成功（读取直到目录的结尾），它会返回切片和一个空的错误。如果它在目录的结尾前遇到了一个错误，Readdir返回直到当前所读到的FIleInfo和一个非空的错误。


func (f *File) Readdirnames(n int) (names []string, err error)
Readdirnames读取并返回目录f里面的文件的名字切片。

如果n>0，Readdirnames返回最多n个名字。在这种情况下，如果Readdirnames返回一个空的切片，它会返回一个非空的错误来解释原因。在目录的结尾，错误为EOF。

如果n<0，Readdirnames返回目录下所有的文件的名字，用一个切片表示。在这种情况下，如果用一个切片表示成功（读取直到目录结尾），它返回切片和一个空的错误。如果在目录结尾之前遇到了一个错误，Readdirnames返回直到当前所读到的names和一个非空的错误。
func (f *File) Seek(offset int64, whence int) (ret int64, err error)　　　　//Seek设置下一次读或写操作的偏移量offset，根据whence来解析：0意味着相对于文件的原始位置，1意味着相对于当前偏移量，2意味着相对于文件结尾。它返回新的偏移量和错误（如果存在）。
func (f *File) Stat() (fi FileInfo, err error)　　//返回文件描述相关信息，包括大小，名字等．等价于os.Stat(filename string)
func (f *File) Sync() (err error)                        //同步操作，将当前存在内存中的文件内容写入硬盘．
func (f *File) Truncate(size int64) error        //类似  os.Truncate(name, size),，将文件进行截断
func (f *File) Write(b []byte) (n int, err error)　　//将b中的数据写入f文件
func (f *File) WriteAt(b []byte, off int64) (n int, err error)　//将b中数据写入f文件中，写入时从offset位置开始进行写入操作
func (f *File) WriteString(s string) (ret int, err error)　　　//将字符串s写入文件中




**FileInfo经常被Stat和Lstat返回来描述一个文件**

func Lstat(name string) (fi FileInfo, err error)      //返回描述文件的FileInfo信息。如果文件是符号链接，返回的FileInfo描述的符号链接。Lstat不会试着去追溯link。如果出错，将是 *PathError类型。
func Stat(name string) (fi FileInfo, err error)       //返回描述文件的FileInfo信息。如果出错，将是 *PathError类型。

FileMode代表文件的模式和权限标志位。标志位在所有的操作系统有相同的定义，因此文件的信息可以从一个操作系统移动到另外一个操作系统。不是所有的标志位是用所有的系统。唯一要求的标志位是目录的ModeDir。
FileModel的方法主要用来进行判断和输出权限
func (m FileMode) IsDir() bool              //判断m是否是目录，也就是检查文件是否有设置的ModeDir位
func (m FileMode) IsRegular() bool　　//判断m是否是普通文件，也就是说检查m中是否有设置mode type
func (m FileMode) Perm() FileMode　　//返回m的权限位
func (m FileMode) String() string　　　　//返回m的字符串表示

type LinkError

func (e *LinkError) Error() string　　　　//LinkError记录了一个在链接或者符号连接或者重命名的系统调用中发生的错误和引起错误的文件的路径。

type PathError
func (e *PathError) Error() string　　//返回一个有操作者，路径以及错误组成的字符串形式

type ProcAttr // 进程相关操作

func FindProcess(pid int) (p *Process, err error)　　　　//通过进程pid查找运行的进程，返回相关进程信息及在该过程中遇到的errorfunc StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)  //StartProcess启动一个新的进程，其传入的name、argv和addr指定了程序、参数和属性；StartProcess是一个低层次的接口。os/exec包提供了高层次的接口；如果出错，将会是*PathError错误。func (p *Process) Kill() error　　　　　　　　　　　//杀死进程并直接退出func (p *Process) Release() error　　　　　　　　//释放进程p的所有资源，之后进程p便不能够再被使用，只有Wati没有被调用时，才需要调用Release释放资源

func (p *Process) Signal(sig Signal) error　　　　//发送一个信号给进程p, 在windows中没有实现发送中断interrupt
func (p *Process) Wait() (*ProcessState, error)　　//Wait等待进程退出，并返回进程状态和错误。Wait释放进程相关的资源。在大多数的系统上，进程p必须是当前进程的子进程，否则会返回一个错误。
type ProcessState     //ProcessState存储了Wait函数报告的进程信息

func (p *ProcessState) Exited() bool　　　　　　// 判断程序是否已经退出
func (p *ProcessState) Pid() int                                //返回退出进程的pid
func (p *ProcessState) String() string                     //获取进程状态的字符串表示
func (p *ProcessState) Success() bool                    //判断程序是否成功退出，而Exited则仅仅只是判断其是否退出
func (p *ProcessState) Sys() interface{}                //返回有关进程的系统独立的退出信息。并将它转换为恰当的底层类型（比如Unix上的syscall.WaitStatus），主要用来获取进程退出相关资源。
func (p *ProcessState) SysUsage() interface{}         //SysUsage返回关于退出进程的系统独立的资源使用信息。主要用来获取进程使用系统资源
func (p *ProcessState) SystemTime() time.Duration      //获取退出进程的内核态cpu使用时间
func (p *ProcessState) UserTime() time.Duration     //返回退出进程和子进程的用户态CPU使用时间

type Signal
代表操作系统的信号．底层的实现是操作系统独立的：在Unix中是syscal.Signal．
var (
	Interrupt Signal = syscall.SIGINT
	Kill      Signal = syscall.SIGKILL
)
在所有系统中都能够使用的是interrupe,给进程发送一个信号，强制杀死该进程
type SyscallError　//SyscallError记录了一个特定系统调用的错误，主要用来返回SyscallError的字符串表示
func (e *SyscallError) Error() string　　　//返回SyscallError的字符串表示
```


### 附录
参考文章
[golang中os包用法](https://studygolang.com/articles/5024)
