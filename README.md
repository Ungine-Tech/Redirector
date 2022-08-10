# Redirector

跳转用Web程序

## 参数

* `-D` 开启Gin调试模式
* `-t <url>` 指定要跳转的模板URL，默认为 `http://localhost/{{path}}`
* `-p <port>` 跳转端口，默认为 `:80`
* `-s` 是否启用HTTPS
* `-sp <port>` HTTPS端口，默认为 `:443`
* `-pem <path>` 证书路径，默认为 `./cert.pem`
* `-key <path>` 私钥路径，默认为 `./cert.key`

## 编译

编译需要 Golang 环境

```bash
go build -o redirector
chmod +x redirector
```

## 运行

```bash
# 请自行修改参数
./redirector -D -t http://localhost -p :80 -s -sp :443 -pem ./cert.pem -key ./cert.key
```

或使用 Docker

```bash
docker build . -t redirector
docker run -p 80:8080 -p 443:8443 -d -t redirector
```

请自行修改端口和调整相关参数
