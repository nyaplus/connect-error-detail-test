# connect-error-detail-test
connect-error-detail-test

```bash
grpcurl -plaintext -d '{"name": ""}' localhost:8082 greet.v1.GreetService.Greet
```


## go.mod
```
// add
replace github.com/bufbuild/connect-grpcreflect-go => github.com/2yanpath/connect-grpcreflect-go fix/error-detail
```