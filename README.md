# TODOLIST
![build](https://travis-ci.org/kangbb/todolist.svg?branch=master)
<br/>一个简单的网站demo。
## 主要特点
1. 完全自动化部署：Travis CLI持续集成+阿里云docker容器服务+阿里云ECS主机
2. 使用docker及docker-compose快速编排服务
3. golang作为后台语言
4. mysql实现主从服务配置
5. ngix实现反向代理
## 项目内容说明
- `core/`：后端服务核心，controller处理前端请求，service提供数据库服务，models定义数据模型
- `front_end/`：前端文件，基于Vue
- `mysql/`：主从配置相关文件。包括mysql的构建文件及docker-compose.yml文件
- `nginx/`：nginx相关配置文件
- `static/`：Vue编译生成的真正的端文件
- `main.go`： 后端程序入口文件
- `docker-compose.yml`：部署应用及nginx
- `Dockerfile`：生成镜像
- `.travis.yml`：自动集成及部署文件

## 其他
如何使用travis cli免密登录服务器？
>[一点都不高大上，手把手教你使用Travis CI实现持续部署](https://zhuanlan.zhihu.com/p/25066056)

如何加密密钥？通过travis-cli运行如下命令：
```
$ travis encrypt-file ~/.ssh/id_rsa --add
```