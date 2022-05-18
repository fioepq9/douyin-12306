# douyin-12306

## 项目结构

```bash
├─cmd
│  ├─api			# api 服务，负责监听 http 请求并转发到其他服务
│  │  ├─handlers	# 控制器目录
│  │  ├─middleware		# 中间件
│  │  ├─router			# 路由注册
│  │  └─rpc				# 封装了调用其他 rpc 服务的逻辑
│  └─user			# user 服务，提供 user 相关的接口，包括注册、登录、获取用户信息等
│      ├─dal			# DAL 层，负责存储
│      │  ├─db				# MySQL 相关代码
│      │  └─rds				# Redis 相关代码
│      ├─output			# 自动生成的代码
│      ├─pack			# 数据打包/处理
│      ├─script			# 运行脚本目录
│      ├─service		# 封装了业务逻辑
│      ├─build.sh		# 构建脚本
│      ├─handler.go		# rpc接口定义，使用kitex自动生成
│      └─main.go
├─config			# 配置模块
├─idl				# idl，proto接口定义文件
├─kitex_gen			# kitex 自动生成的 rpc 接口包			
├─logger			# 日志模块
├─pkg				
│  ├─errno			# 错误码
│  ├─middleware		# rpc 中间件
│  ├─tracer			# Jarger 初始化
│  └─util			# 工具模块
├─repo				# 负责数据存储层的初始化定义
├─config.yml		# 配置文件
├─build.sh			# 构建项目依赖的脚本
└─run.sh			# 运行项目的脚本
```

## 项目运行

> 环境要求：dokcer

### 构建(导入新的包时需要执行)

```bash
$ sh build.sh
```

### 运行

```bash
$ sh run.sh
```