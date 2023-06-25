# BilibiliMallDownload
 工房不让下载？我偏要下载！

## 用法

### 下载可执行文件

```cmd
BilibiliMallDownload.exe your-config-file.json
```

### 手动构建

```cmd
go run main.go your-config-file.json
```

### 配置文件格式
```json
{
    "accessKey": "",    // 非必要
    "cookie": "",       // 登录信息
}
```

> 可使用 [BilibiliLogin-Lite](https://github.com/FangCunWuChang/BilibiliLogin-Lite) 生成登录信息, 或直接从浏览器开发者工具中复制cookie字符串