go build pbplugin
mv pbplugin /usr/bin/protoc-gen-go-netrpc
cd src/self_proto/comm
protoc --go-netrpc_out=plugins=netrpc:. hello.proto

chown ruima:ruima -Rf *


