# Go study

## 文件以及变量命名规范

- Go 的文件访问是范围是 package
- 同一个目录下面的文件的 package 必须一样，但不一定需要和 dir 的名字一样

### 公司要求

- 首字母缩略词
    - 缩略部分全部大写或者小写
    - 不能使用驼峰的方式
- 方法接收者命名
    - 一般为 1-2 个字符
    - 同一接收者命名一直
    - 指针或者值的传递，遵循最小原则
- 使用sync.WaitGroup，需要在main goroutine中Add，在goroutine中Done
- 建议使用提前定义的错误，通过使用==或者errors.Is()区分错误类型
- 在执行SQL语句时，对于参数查询的情况，必须使用预编译或使用GORM等ORM执行SQL操作方式进行，从而在根本上消除SQL注入风险。
- 命令执行检查
- 文件路径执行检查

### 包

- 命名
    - 包命名必须是小写字母
- 规范
    - 从上到下依次
        - 标准库
        - 第三方
        - 本地库
    - 分组
        - 相同组之间不能有空行
        - 不同组之间需要一行空行
        - 分组内的包按照分组里面的顺序依次升序排序

### 文件命名规范

- Go 的文件命名规范和 Dart 语言的是一样的
- 小写字母 + 下划线

### 变量命名规范

- 变量的命名方式和 Dart 语言的是一样的
- 使用驼峰的方式命名

## 关键字

- func
    - 用来定义函数
    - 函数有匿名函数和具名函数
        - 匿名函数的名字在当前的函数里面是 `func1, func2, ...`
        - 具名函数的名字就是函数定义的时候的名字

- new & make

    - new(malloc realloc calloc)
        - 最终返回的是空间的一个 *指针*
        - 不会进行初始化
        - 可以用于任何类型
    - make
        - 返回已经初始化了的实例
        - 返回的实例也就是具体的实例，而不是地址
        - 用于 slice map chan slices 等
        - make(type, size)
            - 返回的是一个 slice，chan，map
            - 返回的 slice 的 len 和 cap 都是 size
        - make(type, size, cap)
            - 返回的是一个 slice
            - 返回的 slice 的 len 是 size，cap 是 cap
- switch

    - switch 比较类似浅拷贝
    - 基本数值类型
    - 可以直接对类型进行判定 case: int; case: bool
    - 表达式
    - 特殊的 switch 断言代码

  ```
  var typeFindFun = func(v any){
    switch t := v.(type){
    case int:
      fmt.Println("int", t)
    case bool:
      fmt.Println("bool", t)
    default:
      fmt.Println("I don't know", t)
    }
  }
  ```

- `...`
    - Go 的 ... 的使用很具象
    - ... 可以用来解构
        - arr...
    - ... 可以用来结构化
        - ...arr

- interface

    - 接口其实就是特殊的指针，传递的时候需要注意
- struct
    - 命名嵌套
        - 命名嵌套的时候，我们必须通过名字访问
    - 匿名嵌套
        - 匿名嵌套，我们可以匿名访问
        - 也可以通过名字访问
    - 结构体定义方法
        - 需要注意是通过指针定义还是通过结构体本身定义
        - 针对一个结构体的方法的定义，需要满足统一、高性能两点
    - tag 的使用 (也就是反引号)

- defer
    - 用来处理收尾工作
    - defer 是按照栈的顺序执行
    - defer 让逻辑上面存在较远距离的代码在实际编写的时候写在了一起
    - defer 在 return 之前执行(也就是 return 永远是最后)
    - 需要注意 defer 的排队顺序
    - 同级遇到panic的时候，会按照FILO的顺序执行注册的defer里面的代码

- return
    - Go 返回 nil 和 没有返回值是不一样的
    - 在 Python 里面，没有返回值就是返回 None
- go
    - 用来启动 goroutine
    - 需要避免死锁
        - WaitGroup
            - Add()
            - Done() 相当于 -1 wg 根据 delta 的数量来判定是否需要 wait
            - Wait() 和其他语言的一样的
    - 主线程需要能够跑完，主线程跑完及时其他部分会死锁也没问题
- chan
    - chan 是一个特殊的队列
        - chan 可以用来实现线程同步
        - chan 作为函数参数的时候我们可以指定为 *仅仅能输入数据或者输出数据*
    - 默认没有缓冲区
    - <-
        - 每一次这个符号的执行都会卡在符号执行处
            - 入的时候没有多余空位
            - 出的时候没有其他东西
- range
    - range 可以用来遍历 chan
    - range 在 chan 是 close 并且为空的时候会结束 range
- select
    - select 是一个普通的顺序执行的语句
- panic
    - 是一个函数
    - 用来引出一个异常
    - 显式的异常处理机制
- recover
    - 类似其他语言的 catch
    - 用来捕获先有得 panic
- var
    - 用来申明变量
    - var a Person
        - 申明的同时申请了空间

## Go语言的数据类型

- 基本数据类型
- 指针类型
- 引用类型

## 断言

- val.(Type)
  1. 断言只能一层一层断言，不能一下子
  ```
  itemNew := itemOld.([]any)// itemOld 其实是一个 []map[string]any，但是断言的时候不能一下子就断言到 []map[string]any ，只能一层一层来
  ```

## 变量

- 值类型和引用类型
    - 基本数据类型都是值类型
    - 数据、结构体也是值类型
- 引用类型
    - 引用类型传递的是引用的副本
    - slice
        - slice 有些时候也需要传递指针
        - 需要重新切片或者重新分配切片的时候需要传递指针
    - map
    - chan

1. array

- array 申明的过程中可以使用 ... 而不用自己数有多少个
    - ... 只能用在申明语句的右边(var arr [5]int = [...]int{0,1,2, 6: 6,7})

## 函数(内置)

- clear
    - 将一个切片的元素全部置为默认值
- delete
- append
- len
- cap
- make
- new
- init
  - 无参、无返回值
  - 该函数针对 package 使用
  - 一个 package 里面的 go 文件里面可以有多个 init 函数(go 语言不允许函数同 package 重名)
    - init 函数的执行顺序按照在文件里面的出现顺序执行

## 内置常量

- iota(aɪˈəʊtə)
    - 极微小
    - const iota = 0
    - 在编译的时候会被处理，在一个定义域里面的使用会逐渐 ++，换一个定义域会置零从新开始
    - iota 生成的常量默认是无类型的常量 untyped constant，知道被使用的时候才指定敌营的类型

## package

### fmt
1. fmt.Fprintln()
   - F 表示的是 File 的意思

### time

- NewTimer()
    - 是一个定时器，用来实现定时任务
- NewTicker()
    - 提供了 Stop 方法，用来实现周期任务
- Tick()
    - 返回的是一个在程序整个执行过程中不需要停止的 Tick (没有提供 Stop 方法)

### context

context 在设计的时候就是分层次的

- 用来处理并发(如下是使用例子)
    - 取消信号
        - 对于取消函数 cancelFunc 按照逻辑是可以重复调用的(实际上也是)
    - 超时控制(WithTimeout)
    - 请求范围控制
        - 这个主要是在 http 请求里面使用
    - 传递元数据
    - 层级管理
    - 资源释放

- Context 接口

    ```
    type Context interface{
        Done()<-chan struct{} // 返回一个通道，通道在 context 截止时间到了、取消或者超时的时候会关闭
        Deadline() (deadline time.Time, ok bool) // 返回 context 被取消的时间
        Err() error           // 返回 context 被取消的原因
        Value(key any) any    // 返回 context 中与 key 关联的值
  }
    ```

- package 提供的几个重要函数
    - 基本的两个创建 空 Context(空 Context 也就是 struct{}) 的函数
        - context.Background()
            - 返回一个空的 Context，一般是不可操作的
        - context.TODO()
            - 返回一个空的 context，用在尚未确定需要使用什么 Context 的场景
    - 可以取消或者有截止时间的 Context
        - context.WithCancel(parent Context)
            - 返回一个可以取消的 Context
        - context.WithDeadline(parent Context,d time.Time)[使用的是绝对的截止时间]
            - 返回一个具有截止时间的 Context
        - context.WithTimeout(parent Context, timeout Time.Duration)[使用的是相对的超时时间]
            - 返回一个具有超时时间的 Context

### json

- json.Marshal/json.Unmarshal
  1. 根据字段的可访问性来进行序列化
- fmt.Println()
  1. 根据

### sqlx

[sqlx](http://jmoiron.github.io/sqlx/) 是基于 Go 的基本数据库连接内置包 database/sql 构建的第三方包

无论是 Go 的 sqlx 还是 Java 的 jdbc 等，最终都是通过 sql 来进行数据库访问的

X 表示的查询结果是 sqlx 定义的类型

- 主要有四种 handle types
    1. sqlx.DB `database`
    2. sqlx.Tx `transaction`
    3. sqlx.Stmt `statement`
    4. sqlx.NamedStmt `named statement`

#### 数据库的连接

- sqlite3
    - sqlite3 的每一个 *.db 文件对应一个数据库，区别于其他类型的关系型数据库

#### Querying 101(101 在大学表示的是初级课程)

1. sql 带有以及 sqlx 也带有的类型
    1. Exec
    2. Query() *sql.Rows
    3. QueryRow
2. sqlx 提供的方法
    1. MustExec() sql.Result
    2. Queryx
    3. QueryRowx
3. 拓展(将查询的结果以及目标对象都会存储在内存里面)
    1. Get
        - 作用于单行数据
    2. Select
        - 作用于多行数据

#### transactions

开启事务不能使用 sql 来执行因为这样的话可能是来自不同的数据库连接，需要使用 sqlx.DB 携带的 Beginx

- db.Beginx()
- tx.Exec("sql")
- tx.Commit()

#### prepared statements

用来预处理 sql 查询，减少编译(重复使用)，防止 sql 注入

- sql 注入通过参数化查询解决问题
    - 一般查询直接把参数拼接在了 sql 语句里面
    - 参数化查询先解析编译 sql 句型，在传入 参数

#### Query Helper

- in queries
- named queries
    1. 使用的是 `:name` 的语法

#### Rebind

- 用来处理不同数据库驱动的占位符问题
- 用来解决 sqlx.In 的 动态占位符问题

#### scan

- 针对一行数据，将数据结果按照查询顺序使用 scan 进行赋值

### goqu

goqu 通过函数式编程的方式将 sql 语句使用 Go 语言的形式表现出来  
是一个 SQL 生成库

### testify

- 用来测试的第三方包
- 在 Go 的测试里面
    - 用来测试的 Go 文件需要使用 _test.go 命名
    - 一般测试的函数需要使用 Test 来开头
    - 如果需要进行基准测试的话需要使用 `BenchmarkXXX(b *testing.B)` 这里的 `b.N` 就是一个循环次数

- assert
- suite
- mock
    - 用于单元测试中模拟接口的实现，以便于在测试中替代真实的依赖项。

### thrift

### validate

用来校验结构体字段是否符合 tag 部分的约束

### gin 框架

Golang 的 web 框架

- gin 的 HandlerFunc 怎么接受参数的，这个是闭包实现的吗?
- context 是怎么使用的？我怎么没有察觉到呢？还是说Context自己处理了呢？

### viper

- Golang 的配置管理库
  1. 本身就是用来对配置进行管理的
  2. 读取的配置将会保存在 viper 的实例里面

- viper 的基本操作
  1. AutomaticEnv
     - 用来自动读取环境变量，但是读取到的环境变量不会存储在 viper 实例中，而是每次调用 Get 方法的时候就会动态获取
  2. Get
     - 用来获取配置值(获取先后顺序)
       1. Set 现实设置的
       2. 环境变量的(唯不能 Marshal 的)
       3. 配置文件的
       4. SetDefault 默认值
```
func main() {
	var v = viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error read config file: %s \n", err))
	}
	_m := make(map[string]interface{})
	err = v.Unmarshal(&_m)
	if err != nil {
		panic(fmt.Errorf("Fatal error marshal config file: %s \n", err))
	}
	_jsonString, _ := json.Marshal(_m)
	fmt.Println(string(_jsonString))
}
```

### cobra

Golang 的命令行库  
CLI: command line interface

1. 定义 rootCmd 
   1. rootCmd 下面可以添加 subCmd
   2. 编译之后的生成的可执行文件的名字代表该命令的名字，因此 Use 这个字段在这里有点小问题
   3. 需要让 main 编译之后的可执行命令和 rootCmd 的 Use 名字一致
2. 定义 subCmd
   1. 定义的方法和 rootCmd 一样(cmd的定义的方法只有一个)
   2. 使用: `./rootCmd subCmd`
3. flag
   1. 就是用来添加 - 类型参数
      1. 参数类型有 bool string int 等
         1. BoolVarP(bool var p(p 表示短的意思 shorthand))
         2. bool 类型的只要传入的时候就是表示 true
            - `rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "启用详细输出")`
            - 上面的表示 verbose 这个命令默认为 false, 只要执行的时候 `rootCmd -v xxx` 里面有 `-v` 就是 true
4. 编写 main 函数
   1. main 函数输出

### zap

Golang 的高性能日志库
- 类似的高性能日志库还有 zerolog
- 提供了多个日志接口
- core
  - zapcore.Core
    - 决定哪些日志需要记录(LevelEnabler)
    - 决定编码的规则(Encoder)
      - console
        - 是一种通用的日志格式概念，不是 zap 特有的(一般包含了时间戳、日志级别、消息内容等)
        - 是一种方便在终端/控制台查看的格式
      - json
    - 决定日志写在哪(WriteSyncer)

### wire 
- Golang 的依赖注入库
- Java 对应的有 Spring 系列的 @Autowire
- Go 需要简洁
  - 使用 wire.Build 接收的时候，Provider 不能多，也不能少，但是如果是使用Set的话，Set里面元素的可以多余
- 两个重要的函数
  - NewSet -> ProviderSet
    - 对于经常被使用的 component，可以装在 Set 里面，以后要使用的时候，直接使用就可以了
  - Build
    - 这个是核心，传入 一系列的 component 就会被 build 解析，然后造出目标 component 
- 其他函数
  - wire.Bind
    - wire.Bind(interface, interfaceImplType)
    - 绑定现在需要的接口和实际的类型
  - wire.Struct
    - wire.Struct(new(_Struct,"field1", "field2"))
    - 用来构造需要的结构体，可以使用 "*" 代表所有的 fields
  - wire.Value
    - wire.Value(Struct_{V: literalVal})
    - 绑定的值只能使用字面量
  - wire.InterfaceValue
    - wire.InterfaceValue(new(interface), interfaceImplObj)
    - 第一个参数是接口，第二个参数是接口实现的结构体的instance
- 两个重要的概念
  - provider
    - 简单来说就是一个工厂函数
  - injector(initialize/inject)
    - 调用 wire.build() 的函数
- 注入的理解
  1. wire 什么时候函数的参数部分可以使用指针？什么时候返回值的那儿可以使用指针？
  2. 222

### casbin

Golang 的控制访问库

### reflect

提供了在运行时动态操作对象的能力(运行时对变量进行操作)

- DeepEqual(x,y any)
    - DeepEqual是递归进行判断，在Go里面<span style="color:red">类型</span>不一样就判定为 false ，只有类型一样才能进行对比
        - 使用了 type 定义之后，和原来的就不是同一种类型了，类型一样，这里的要求就是类型名字一模一样，是同一个，而不是本质上一样
        - 使用 == 来进行对比的时候，如果编译器检测出类型
### go-testy
？？？
### runtime

runtime 包提供了与 Go 运行时系统交互的功能，管控 Go 程序的运行是行为(协程干礼、内存分配、垃圾回收)

### os

- 临时文件以及文件夹
    - 很多时候我们的文件在在使用结束之后就不再需要了，因此需要临时文件/文件夹
    - 没有指定创建的位置的时候，就会创建在默认的临时文件目录下，这样我们可以不用显式删除
    - `os.CreateTemp(dir, pattern string)` pattern 也仅仅是前缀

## 代码生成和构建指令

- 语法
    - `//go:`
- 指令类型
    - `//go:generate <command>`
        - 在编译前生成代码
    - `//go:build <constraint>`
        - 用来指定文件的构建约束
    - `// +build`
        - 旧版本的，现在都是使用 `//go:build`
    - `//go:noinline`
        - 禁止编译器对函数进行内联优化
    - `//go:nosplit`
        - 禁止编译器在函数中插入栈分割检查代码
    - `//go:linkname myFunction runtime.myFunction`
        - 将一个函数或者变量的名称链接到另一个保重的非导出名称
