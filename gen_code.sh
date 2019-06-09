go build pbplugin
mv pbplugin /usr/bin/protoc-gen-go-netrpc
cd src/self_proto/comm
protoc --go-netrpc_out=plugins=netrpc:. hello.proto
chown ruima:ruima -Rf *
cd ../../../
cd src/gRPC/proto
protoc --go_out=plugins=grpc:. hello.proto
chown ruima:ruima -Rf *
