# go-short
go 短链接服务，类似于git.io

## 运行方式
### 修改配置文件
```bash
mv config_production.yaml.example config_production.yaml
```
### 运行
1. 编译并运行
```bash
docker build -t short .
docker run -dp 8088:8088 -v "/data/log:/var/log/short.liu.app" --name short short
```
2. 或者直接运行我编译好的
```bash
docker run -dp 8088:8088 -v "/data/log:/var/log/short.liu.app" ghcr.io/yezige/short.liu.app:latest --name short
```
### 增加nginx
```nginx
location ^~ /s/ {
  proxy_pass http://short:8088/;
}
```