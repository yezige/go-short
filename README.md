# go-short

go 短链接服务，类似于 git.io

## 运行方式

### 修改配置文件

```bash
mv config_production.yaml.example /data/config_production.yaml
```

### 运行

1. 编译并运行

```bash
docker build -t short .
docker run -d --expose 8088 -v "/data/log:/var/log/short.liu.app" -v "/data/config_production.yaml:/short.liu.app/config_production.yaml" --name short short
```

2. 或者直接运行我编译好的

```bash
docker run -d --expose 8088 -v "/data/log:/var/log/short.liu.app" -v "/data/config_production.yaml:/short.liu.app/config_production.yaml" --name short ghcr.io/yezige/short.liu.app:latest
```

### 增加 nginx

```nginx
location ^~ /s/ {
  proxy_pass http://short:8088/;
}
```

### redis 增加对应 key

set goshort:path:trace https://raw.githubusercontent.com/yezige/trace/main/run.sh

### 成功

则可使用短链接访问

```bash
curl liu.app/s/trace | bash
```
