
https://gangmax.me/blog/2019/09/10/fix-the-google-slash-protobuf-slash-descriptor-dot-proto-file-not-found-issue/

```
git clone https://github.com/protocolbuffers/protobuf.git

```

```powershell

protoc -I ../../protocolbuffers/protobuf/src -I ./ --go_out=./  --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative  helloworld/helloworld.proto 

```