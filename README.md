# BilibiliMallDownload
 ~工房不让下载？我偏要下载！~
 
工坊可以下载了，是时候改成public仓库了

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

> 可~使用 [BilibiliLogin-Lite](https://github.com/FangCunWuChang/BilibiliLogin-Lite) 生成登录信息~（缺少一些必要项，会获取失败）, 或直接从浏览器开发者工具中复制cookie字符串
