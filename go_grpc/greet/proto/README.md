to generate *_grpc.pd.go and *.pd.go

```bash
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/Ahmed-Amine-Soltani/go-labs/go_grpc --go-grpc_out=. --go-grpc_opt=module=github.com/Ahmed-Amine-Soltani/go-labs/go_grpc greet/proto/dummy.proto
```

we can use the Makefile from the cours

