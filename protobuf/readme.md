#

```
go install github.com/franklihub/nrpc/protoc-gen-nrpc@latest
```

~~protoc -I ./mpcCommon/protobuf --go_out=./api --nrpc_out=./api ./mpcCommon/protobuf/tfa/tfa.prot~~

```
protoc  -I ./mpcCommon/protobuf --go_out=./api --nrpc_out=./api ./mpcCommon/protobuf/riskengine/riskengine.proto
```

```
protoc  -I ./protobuf --go_out=./riskCtrlService  --nrpc_out=./riskCtrlService ./protobuf/riskctrl.proto
```
```
protoc  -I ./protobuf --go_out=./authService  --nrpc_out=./authService ./protobuf/auth.proto
```
