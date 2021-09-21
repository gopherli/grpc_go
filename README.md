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

(2)问题及解决：Failed to connect to github.com port 443:connection timed out
    取消验证
        git config --global http.sslVerify "false"
    取消代理（视情况省略）
        git config --global --unset http.proxy
        git config --global --unset https.proxy

```

# 安装protoc
1. Protocol buffer 3 下载：https://grpc.io/docs/protoc-installation/
```
说明：解压后的protoc.exe，放入到$GOPATH/bin目录
```
2. Go plugins for the protocol compiler 下载：https://grpc.io/docs/languages/go/quickstart/
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

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
