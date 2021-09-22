# grpc_go
grpc 简单搭建demo

# 安装grpc
- 参考：https://segmentfault.com/a/1190000022808406
- 第一步：进入到$GOPATH/src目录，执行git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
- 第二步：进入到$GOPATH/src目录，下载对应包（发现会有很多错误，根据提示可以发现是由于缺少包的原因）
```
(1)所需包：
    git clone https://github.com/golang/text.git ./golang.org/x/text
    git clone https://github.com/golang/net.git ./golang.org/x/net
    git clone https://github.com/google/go-genproto.git ./google.golang.org/genproto
    git clone https://github.com/protocolbuffers/protobuf-go.git ./google.golang.org/protobuf
    git clone https://github.com/golang/protobuf.git ./github.com/golang/protobuf(看情况省略)

(2)问题及解决：Failed to connect to github.com port 443:connection timed out | OpenSSL SSL_read: Connection was reset, errno 10054
    取消验证
        git config --global http.sslVerify "false"
    取消代理（视情况省略）
        git config --global --unset http.proxy
        git config --global --unset https.proxy

```
- 第三步：go get google.golang.org/grpc（前面两个做好的前提下）

# 安装protoc
1. Protocol buffer 3 下载：https://grpc.io/docs/protoc-installation/
```
说明：解压后的protoc.exe，放入到$GOPATH/bin目录
```
2. Go plugins for the protocol compiler 下载：https://grpc.io/docs/languages/go/quickstart/
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26  --> helloworld.pb.go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1  --> helloworld_grpc.pb.go

说明：自动生成在$GOPATH/bin目录。
```
3. protoc执行命令：
```
protoc --go_out=. --go-grpc_out=.  *.proto
说明：--go_out=.生成helloworld.pb.go，--go-grpc_out=.生成helloworld_grpc.pb.go
```

4. option go_package = "./;helloworld";
```
option go_package = "google.golang.org/grpc/examples/helloworld/helloworld";
一般用法：声明生成的 go 文件所属的包"github.com/grpc_go/helloworld",方便导入pb

option go_package = "./;helloworld";
说明：./;默认路径，helloworld包名。
```

# 运行服务端、客户端
1. 报错不要紧，配置问题
2. 找不到包问题
```
greeter_client\main.go:9:2: cannot find package "grpc_go/helloerdan" in any of:
        D:\CodingSoft\go\go.16.4\src\grpc_go\helloerdan (from $GOROOT)
        D:\CodingSpace\go_workspace\src\grpc_go\helloerdan (from $GOPATH)
解决办法：把项目放到对应目录下
```

# demo参考
- https://github.com/grpc/grpc-go/tree/master/examples/helloworld


# grpc 接口测试工具
- 参考：https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-08-grpcurl.html
- 第一步：启动反射服务
```
    pb.RegisterYourOwnServer(s, &server{})

    // Register reflection service on gRPC server.
    reflection.Register(s)

    s.Serve(lis)
```
- 第二步：grpcurl手工安装
```
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
说明：
两个命令轮流来安装,遇到443和10054用下面命令解决
    取消验证
        git config --global http.sslVerify "false"
    取消代理（视情况省略）
        git config --global --unset http.proxy
        git config --global --unset https.proxy
```
- 第三步：使用命令
```
前提说明：
    1.在使用grpcurl时，需要通过-cert和-key参数设置公钥和私钥文件，链接启用了tls协议的服务。
    2.对于没有没用tls协议的grpc服务，通过-plaintext参数忽略tls证书的验证过程。

常用命令：
（1）list命令查看服务列表
grpcurl -plaintext localhost:1234 list
（2）list子命令查看服务的方法列表
grpcurl -plaintext localhost:1234 list HelloService.HelloService
（3）describe子命令查看更详细的描述信息
grpcurl -plaintext localhost:1234 describe HelloService.HelloService
（4）describe命令查看参数HelloService.String类型的信息
grpcurl -plaintext localhost:1234 describe HelloService.String
（5）-d 参数传入一个json字符串作为输入参数，调用服务方法
grpcurl -plaintext -d '{"value": "gopher"}' localhost:1234 HelloService.HelloService/Hello
（6）-d 参数是 @ 则表示从标准输入读取json输入参数，这一般用于比较输入复杂的json数据，也可以用于测试流方法
grpcurl -plaintext -d @ localhost:1234 HelloService.HelloService/Channel //回车
{"value": "gopher"} //输入参数
```