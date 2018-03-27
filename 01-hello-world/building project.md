## Workspace
Define location of the workspace with GOPATH
```shell
export GOPATH=$(pwd)
```

Create pkg bin src directories


PATH should contain GOPATH/bin directory
```shell
export PATH=$PATH:$GOPATH/bin
```
## install protocol buffers 
Install the protocol compiler (prtotoc)
Install the go proto package:
```shell
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

## Download/build/install packages 
In workspace directory: download/build/install packages defined in src as dependencies
```shell
go get ./…
```



##Run the standalone executable 
In one step (compile + run):
```shell
go run main.go  #(or go run /path/to/go_file , example: go run src/grpc/01-hello-world/server/main.go)
```

In 2 steps:
```shell
go build    # make go source file into an executable binary file
./server    # run the executable binary file (the parent directory name)
```







