# alfred-workflow

##  Search Password in Yaml File

1. password file
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
* [workflow下载](workflows/search-password)
* 修改环境变量 `PASSWORD_PATH` 为 password file 的位置
