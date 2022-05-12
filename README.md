# douyin-12306

## 项目结构

> [项目结构 · gin实践系列 · 看云 (kancloud.cn)](https://static.kancloud.cn/lhj0702/sockstack_gin/1805357)

```bash
├─config		# 配置模块目录
├─controller	# 控制器目录
├─logger		# 日志模块目录
├─models		# 模型目录，负责项目的数据存储部分，例如各个模块的Mysql表的读写模型。
├─pkg			# 自定义的工具类等
│  ├─e			# 项目统一的响应定义，如错误码，通用的错误信息，响应的结构体
│  └─util		# 工具类目录
├─public		# 静态资源目录
├─repository	# 数据操作层，定义各种数据操作。
├─requests		# 定义入参及入参校验规则
├─responses		# 定义响应的数据
├─router		# 路由目录
├─service		# 服务定义目录
├─temp			# 临时文件目录，包含日志等信息。
├─config.yml	# 项目配置文件
└─main.go		# 项目入口，负责Gin框架的初始化
```
