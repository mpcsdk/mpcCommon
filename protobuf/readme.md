#

```
go install github.com/nats-rpc/nrpc/protoc-gen-nrpc@latest
```

~~protoc -I ./mpcCommon/protobuf --go_out=./api --nrpc_out=./api ./mpcCommon/protobuf/tfa/tfa.prot~~

```
protoc  -I ./mpcCommon/protobuf --go_out=./api --nrpc_out=./api ./mpcCommon/protobuf/riskengine/riskengine.proto
```

```
protoc  -I ./mpcCommon/protobuf --go_out=./api --nrpc_out=./api ./mpcCommon/protobuf/riskctrl/riskctrl.proto
```
