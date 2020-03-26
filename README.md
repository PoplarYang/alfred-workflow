# :rocket: alfred-workflow

##  Search Password in Yaml File
从yaml格式的文件中查找密码，并复制到剪切板。实现了某些应用或者网站密码的快速查找。

1. password file template
```yaml
default_user: echo@foxmail.com
default_password: 142857
group:
  - name: baidu
    tag: 
      - web
    user: xxxx@126.com
    pass: xxxx

  - name: zhihu
    tag: 
      - web
    user: default
    pass: default
```

2. workflow
* :rabbit: [workflow下载](workflows/Search-Password-in-Yaml-File.alfredworkflow)
* 修改 workflow 环境变量 `PASSWORD_PATH` 为 password file 的位置

## Github Url Convert to Raw

:rabbit: [workflow下载](workflows/Github-Url-Convert-to-Raw.alfredworkflow)

将纯文本格式的 github 文件转化为可以下载的 url。

如将 `https://github.com/aikuyun/iterm2-zmodem/blob/master/iterm2-send-zmodem.sh` 转化为 `https://raw.githubusercontent.com/aikuyun/iterm2-zmodem/master/iterm2-send-zmodem.sh`。

## [Upload clipboard PNG to Minio](img2url)

