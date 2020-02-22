# alfred-workflow

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
* [workflow下载](workflows/Search-Password-in-Yaml-File.alfredworkflow)
* 修改 workflow 环境变量 `PASSWORD_PATH` 为 password file 的位置
