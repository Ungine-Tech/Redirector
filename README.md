# Redirector

跳转用Web程序

## 参数

* `-D` 开启Gin调试模式
* `-t <url>` 指定要跳转的模板URL，默认为 `http://localhost`
* `-p <port>` 跳转端口，默认为 `:80`
* `-s` 是否启用HTTPS
* `-sp <port>` HTTPS端口，默认为 `:443`
* `-pem <path>` 证书路径，默认为 `./cert.pem`
* `-key <path>` 私钥路径，默认为 `./cert.key`
