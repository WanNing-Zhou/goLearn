# golang debug
1. 本地调试
2. 附加到进程
3. 远程调试


# golang 的编译于与运 

## go build

1. 编译当前目录

```
go build 
或
go build .
```

2. 编译执行文件或目录

``` 

# 编译main文件. 将会生成可执行文件
go build ./main.go
# 编译执行文件目录p1. 不会产生编译文件, 仅进行编译检查
go build ./p1
```


3. main包由多个文件的情况, 执行文件编译包

```
# 编译时由多个文件的情况,指定文件编译main包
go build ./main.go ./hello.go
```

4. 编译时指定编译输出结果

``` 
go build -o ./out/app ./main.gon ./hello.go
```

## 运行 go run 

1. 直接通过 go run 运行main包,即可运行应用程序

``` 
# 运行应用程序时, 指定main包所有文件
go run ./main.go ./hello.go
# 运行应用程序时, 指定main包所在目录
```

## 交叉编译

交叉编译需要修改`GOOS,GOARCH,CGO_ENABLED`三个环境变量  
GOOS:目标平台的操作系统(darwin,freebsd,linux,windows)  
GOARCH: 目标平台的体系架构32位还是64位(()386,adm64,arm)  
CGO_ENABLED: 是否启用CGO,交叉编译不支持CGO所以要禁用它

windows编译Linux与Mac可执行程序

```  
# 设置环境变量
$Env:CGO_ENABKED=0;$Env:GOARCH="amd64";$Env:GOOS="linux"
#编译,并输出到app文件
go build -o ./out/app .

# 设置环境变量
$Env: CGO_ENABLED=0;$Env:GOARCH="amd64";$Env:GOOS="darwin"

# 编译,并输出到app文件
go build -o ./out/app .
```

![img1](./markimgs/img.png)


# go mod命令

```
# 将模块下载到本地缓存,需要执行模块路径及版本阿红
go mod download
# 例如
go mod download github.com/gin-gonic/gin@v1.9.0

# 初始化一个新的模块到当前目录
go mod init
# 例如
go mod init gomodcase

```









