# BETTER-EMBY-SERVER

获取最好的公益服服务器地址

## 使用

首先感谢 [emby 公益服](https://rartv.gitbook.io/emby-public/)，感谢水管工大佬。

### WEB 版本

使用纯浏览器实现，所有操作都在本地实现。

另外，由于浏览器安全策略原因，还分如下两个版本

* [https 版本](https://kejunmao.github.io/better-emby-server/index.html)

  仅能测试 https 链接的 emby 服务器
* [http 版本](http://bes.kejun.me)

  可以测试 http 和 https 的 emby 服务器

### 命令行
#### 基础

使用 `-f` 指定服务器列表，只要把通过 `/create` 命令，厂妹发给你的信息另存为一个文件即可

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

#### 使用代理

使用 `-p` 指定代理，仅支持 `http` `https` `socks5`

```bash
❯ bes -f servers.txt -p socks5://192.168.1.1:1080
http://*****.emby.****:8096 233ms not emby
http://*****.emby.****:80 303ms not emby
http://*****.emby.****:80 403ms not emby
```
#### 帮助

```bash
❯ bes -h
bes - better emby server

并发 http 请求获取公益服地址请求耗时
  -f string
        厂妹发给你的服务器消息另存为文件的路径
  -p string
        使用指定的代理测试
```

## 常见问题

* 全是 0ms error

  检查网络或代理是否有效
* not emby 是什么

  厂妹返回的服务器地址，但是首页是 nginx 哎
