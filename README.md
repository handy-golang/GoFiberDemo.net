# GoFiberDemo.net

基于 goFiber 实现的一整套 web 服务器样例

## go.work

```
go 1.18

use (
	./
)

replace (
	github.com/EasyGolang/goTools => /root/EasyGolang/goTools
)

```

## server_env.yaml

```yaml
# IP
LocalIP: "xx.xx.xx.xx"
# Mongodb
MongoAddress: "158.125.11.11:5689"
MongoUserName: "root"
MongoPassword: "123456"

# 运行模式
RunMod: 1
```
