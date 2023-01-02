# Golang 笔记

1. float32跟float64默认零值为0

2. 字符类型用byte表示，原样输出只能用格式化c%的形式

3. 布尔类型格式化输出t%，类型T%

4. 整数如果不指定类型为int，根据系统代表的是int32或者int64，rune类型代表的是int32

5. 字符串不可变代表的是字符不可变，不能用str[0]=e的方式重新赋值，可以通过强转转换成byte或rune数组在修改！

6. 强制类型转换T(v)

7. 基本数据类型转换为字符串：1.sprintf里边的是变量的类型2.strconv，q%代表加双引号输出变量，FormatFloat第二个参数代表的是格式，'f'代表十进制形式，第三个参数代表小数点后位数，第四个参数代表Float变量是32还是64位，ParseFloat类型的函数，返回值有两个，第二个代表转换产生的错误，可不接收，第二个参数代表的是转换后的位数32或者64，ParseInt第二个参数代表的是进制，第三个代表的是转换后的位数

8. 指针:一定是一个地址值，基本数据类型都有对应的指针类型，＆age可以拿到内存地址，*ptr可以拿到变量对应的值*

9. 算数运算符，10/3结果为整数，10.0/3结果为浮点数，取余a%b=a-a/b*b，在go中只有a++，代表的是自增一！

10. 键盘录入，使用的是fmt包下的Scanln或者Scanf，可以单个录入，或者同时录入，使用空格分隔即可

11. 流程控制:使用if i:=1;i >0{}，for {}可以执行死循环，for循环对字节就行遍历，for-range对字符进行遍历

12. break跟continue配合lable在双重循环中使用，注意区别

13. goto 配合lable标签使用，跳过某些逻辑，但不建议使用，容易造成混乱，return也可以用在流程控制中

14. 函数返回值只有一个时，返回值列表可以省略括号，值传递内存分析图，函数可变参数三个点在数据类型前边，函数支持返回值命名：在函数体中直接return，在返回值列表中(sub int,sum int)！

15. Go中基础类型跟复合类型(数组跟结构体)是值传递，引用类型是地址传递

16. init函数执行顺序:当导入一个包时，会执行这个包中的全局变量、函数定义、及init函数，然后执行本包中的全局变量、函数、及init函数，最后执行main函数

17. 闭包=返回的匿名函数+匿名函数以外的变量，这个变量会一直存在在内存中，供我们使用，但不能滥用闭包！

18. defer关键字:函数执行完毕之后释放资源，按照先进后出的顺序执行defer栈

19. 字符串相关函数：统计长度：内置函数len，按照字节进行统计，所有的内置函数都在builtin包中,遍历字符串:[]rune，字符串转整数strconv.Atoi,整数转字符串strconv.Itoa，字符串是否包含子串strings.Contains,子串第一次出现的位置strings.Index，忽略大小写比较strings.EqualFold，子串出现的次数strings.Count，分割字符串strings.Split，替换字符串strings.Replace最后一个参数是-1代表全部替换，strings.ToLower跟ToUpper，strings.Trim跟TrimSpace，TrimLeft，TrimRight，HasPrefix跟HasSuffix

20. 时间日期函数，time.Now()返回的是一个time.Time的结构体，通过这个变量可以获取年月日时分秒，另外可以通过Format跟Parse实现日期与字符串类型的互转

21. 内置函数:new(int)返回一个对应类型的指针，make ()返回的是引用类型本身

22. defer匿名函数+recover内建函数可以处理异常，保证程序正常执行下边的流程，自定义错误使用errors.New()抛出error，调用者接收之后，可以使用panic()终止程序执行，通过recover可以收集到panic的异常，保证程序正常执行！

23. 标准库中log包log.Panic跟log.Fatal的区别，标准库中没有Error方法
    ```azure
    Panic ： 即记录当前错误，记录为失败，但是继续执行，可以被defer函数捕获，类似于panic
    
    Fatal ： 即记录当前错误，记录为失败，不继续执行,defer函数不会执行
    ```
    
    Go内置的log库功能有限，例如无法满足记录不同级别日志的情况，我们在实际的项目中根据自己的需要选择使用第三方的日志库，如logrus、zap等。
    

24. 单元测试
    
规则如下：
    
```

    1、文件名必须以xx_test.go命名

    2、方法必须是Test[^a-z]开头

    3、方法参数必须 t *testing.T

    4、使用go test执行单元测试
    
```

执行命令：

```go

 go test  ./...
     
    ok      GoBasic/gotest  (cached)

 go test -v  ./...
     
    === RUN   TestAdd
    calc_test.go:12: 测试通过，期望值：3,实际值：3
    --- PASS: TestAdd (0.00s)
    === RUN   TestSub
    calc_test.go:21: 测试通过，期望值：7,实际值：7
    --- PASS: TestSub (0.00s)
    PASS


 go test -v -run TestAdd ./...
     
    === RUN   TestAdd
    calc_test.go:12: 测试通过，期望值：3,实际值：3
    --- PASS: TestAdd (0.00s)
    PASS

```


25. 关于日志跟错误的理解

   日志库：github.com/sirupsen/logrus
   错误处理库：github.com/pkg/errors
  
   源码解析：

   - bulitin包中定义error接口

```go
type error interface {
    Error() string
}
```
## 创建错误

- errors.New

  Go 语言中内置了一个处理错误的标准包，你不需要自己去实现error接口，New函数返回了一个实现了Error接口的结构体：

```go
func New(text string) error {
	return &errorString{text}
}


type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```

```go
import "errors"
var ErrNotFound = errors.New("not found")
```

- fmt.Errorf

fmt 标准包内也有一个创建错误的函数 Errorf ，该函数可以使用占位符设置错误信息，比 errors.New函数更灵活。
```go
func Errorf(format string, a ...interface{}) error
    
import "fmt"    
var ErrHuman = fmt.Errorf("%s不符合我们人类要求", "老苗")    
```

## 打印错误

在项目开发中，错误常常通过函数或方法的返回值携带，返回的位置也通常被放置在最后一位。

```go
// error/file.go
package main

import "io/ioutil"

func LoadConfig() (string, error) {
    filename := "./config.json"
    b, err := ioutil.ReadFile(filename)
    if err != nil {
        return "", err
    }

    content := string(b)
    if len(content) == 0 {
        return "", errors.New("内容为空")
    }
    return content, nil
}
```

 - ReadFile 函数读取 "config.json" 文件内容。
 - (string, error) 返回两个值，第一个为文件内容，第二个为错误。
 - err != nil 用于判断是否有错误，如果有 return 直接返回。
 - string(b) 变量 b 的类型为 []byte ，该操作是将 []byte 类型转为 string 。
 - 增加了一个“内容为空”的错误判断，该错误也可以直接保存到变量中返回。
  
```go
var ErrEmpty = errors.New("内容为空")

func LoadConfig() (string, error) {
// ...
    return "", ErrEmpty
// ...
}
```  

现在假设 "config.json" 文件不存在，调用 LoadConfig 函数看看结果。

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    content, err := LoadConfig()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("内容：", content)
}

// 输出
2021/09/23 16:57:25 open ./config.json: The system cannot find the file specified.
```

- 当 err 不等于 nil 时，打印错误，并退出程序。
- log 标准包包含打印日志的函数集。
- log.Fatal 函数打印日志，并终止程序向下执行。
- 输出的错误消息显示没有找到指定的文件。
- 打印错误时，也可以使用 fmt 包，例如：fmt.Println(err) ，只是输出信息没 log 包多。
     
### os.Exit

该函数通知程序退出，并且该函数后的逻辑将不会被执行。在调用时需要指定退出码，为 0 时，表示正常退出程序。

```go
os.Exit(0)
```

不主动调用该函数，即程序从 main 函数自然结束时，默认的退出码为 0。在使用编写工具时或许能看到成功的退出码信息，例如我使用的是 Goland，执行代码后输出的结果末尾会显示如下信息。

```go
Process finished with exit code 0
```

如果不正常退出，退出码则为非 0，通常使用 1 表示未知错误。

```go
os.Exit(1)
```

在使用 log.Fatal 函数时，内部就调用了 os.Exit(1) 。

## 错误判断

### errors.Is

源码位置：github.com/pkg/errors/go113.go

```go
func Is(err, target error) bool
```

- err 参数为要判断的错误。
- target 参数为要比对的错误对象。

```go
err := GetError()
if errors.Is(err, ErrNotFoundRequest) {
    // 错误 1,错误3
} else if errors.Is(err, ErrBadRequest) {
    // 错误 2
}
```

### errors.As

这个和 errors.Is 函数类似，不同点就是该函数只判断错误类型，而 errors.Is 函数不仅判断类型，也要判断值（错误消息）。

源码位置：github.com/pkg/errors/go113.go

```go
func As(err error, target interface{}) bool { return stderrors.As(err, target) }
```

```go
// error/as.go
package main

import (
    "errors"
    "fmt"
)

type ErrorString struct {
    s string
}

func (e *ErrorString) Error() string {
    return e.s
}

func main() {
    var targetErr *ErrorString
    err := fmt.Errorf("new error:[%w]", &ErrorString{s: "target err"})

    fmt.Println(errors.As(err, &targetErr))
}

// 输出
true
```

targetErr 变量有几点要求：

- 无需初始化。
- 必须是指针类型，并且实现了 error 接口。
- As 函数不接受nil ，因此不能直接使用targetErr 变量，要使用其引用&targetErr。

## 错误加工

### 错误拼接

在返回错误时，如果想携带附加的错误消息时，可以使用 fmt.Errorf ，现在修改 LoadConfig 函数。

```go
func LoadConfig() (string, error) {
    filename := "./config.json"
    b, err := ioutil.ReadFile(filename)
    if err != nil {
        return "", fmt.Errorf("读取文件出错:%v", err)
    }

    // ...
}
```

现在重新执行上面的 main 函数，还是假设 "config.json" 文件不存在。

```go
// ...
content, err := LoadConfig()
if err != nil {
    log.Fatal(err)
}
//...

// 输出
2021/09/24 11:37:33 读取文件出错:open ./config.json: The system cannot find the file specified.
```

LoadConfig 函数返回的 err 变量中携带了附加的错误消息，但这样有个问题，附加的错误和原始错误消息杂糅在一块导致不能分离。

### 错误嵌套和 errors.Unwrap

源码位置：github.com/pkg/errors/go113.go

```go
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
```

那怎么实现这种嵌套关系呢，还是使用 fmt.Errorf 函数，只是使用另外一个占位符%w ，w 的英文全名就是 wrap 。

```go
func LoadConfig() (string, error) {
    filename := "./config.json"
    b, err := ioutil.ReadFile(filename)
    if err != nil {
        return "", fmt.Errorf("读取文件出错:%w", err)
    }

    // ...
}
```

```go
package main

import (
    "errors"
    "fmt"
    "log"
)

func main() {
   content, err := LoadConfig()
   if err != nil {
      log.Fatal(errors.Unwrap(err))
   }

   fmt.Println("内容：", content)
}

// 输出
2021/09/24 18:11:09 open ./config.json: The system cannot find the file specified.
```

在打印错误时，增加了一个 errors.Unwrap 函数，该函数就是用来取出嵌套的错误，再看看输出的结果，附加的错误信息”读取文件出错:“已经没有了。

## logrus

logrus的使用非常简单，与标准库log类似。logrus支持更多的日志级别：

Panic：记录日志，然后panic。

Fatal：致命错误，出现错误时程序无法正常运转。输出日志后，程序退出；

Error：错误日志，需要查看原因；

Warn：警告信息，提醒程序员注意；

Info：关键操作，核心流程的日志；

Debug：一般程序中输出的调试信息；

Trace：很细粒度的信息，一般用不到；

### 添加字段

有时候需要在输出中添加一些字段，可以通过调用logrus.WithField和logrus.WithFields实现。logrus.WithFields接受一个logrus.Fields类型的参数

```go
package main

import (
  "github.com/sirupsen/logrus"
)

func main() {
  logrus.WithFields(logrus.Fields{
    "name": "dj",
    "age": 18,
  }).Info("info msg")
}
```