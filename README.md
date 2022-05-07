# douyin-12306

## 项目结构

```
├─config        配置模块 => 框架 viper, 优先级: 环境变量 > config.yml > 默认配置 
├─controller    Management of the REST interface to the business logic
├─logger        日志模块 => 框架 logrus
├─repository	Storage of the entity beans in the system
│  └─mysqlDB        MySQL => 框架 gorm
├─router        路由配置
├─service       Business Logic implementations
├─temp          临时文件目录
├─config.yml    
├─docker-compose.yml
├─Dockerfile
└─main.go
```

![img](https://miro.medium.com/max/1400/1*neBcAZJyLGpE7KHc3sH8bw.png)

## 项目运行

> [Docker Compose配置文件详解（V3）](https://blog.51cto.com/u_15329153/3371134)

+ 启动

```bash
$ docker-compose up
```

+ 启动并重新构建所有镜像

```bash
$ docker-compose up --build
```

+ 关闭

```bash
$ docker-compose down
```

+ 关闭并删除所有本地镜像

```bash
$ docker-compose down --rmi local
```