# BETTER-EMBY-SERVER

获取最好的公益服服务器地址

## 使用

```bash
❯ bes -f ./servers.txt
http://*****.emby.****:8096 1390ms ok
http://*****.emby.****:80 9999ms error
http://*****.emby.****:80 2957ms not emby
...
```

每一列分别对应：

- 服务器地址
- 延时
- 消息
  - `error` 请求失败
  - `ok` 请求成功
  - `not emby` 请求成功但不是 emby 服务器

## 帮助

```bash
❯ bes -h
bes - better emby server

利用 go 并发 http 请求获取最快的公益服地址

  -f string
        厂妹发给你的服务器消息另存为文件的路径
```
