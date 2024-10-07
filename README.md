# Go-music 音乐网站

## 项目简介
这是一个音乐网站的后端服务，使用 Go 语言开发，结合多个现代框架和工具，旨在提供高效、可靠的音乐管理和播放体验。

## 后端技术栈

- **Go 语言**：使用 Go 语言进行后端开发，以其高性能和并发支持为特点。
  
- **Gin 框架**：选择 Gin 框架作为 HTTP web 框架，提供了高效的路由和中间件支持。

- **GORM 框架**：使用 GORM 作为 ORM（对象关系映射）工具，简化了与数据库的交互。

- **MinIO**：集成 MinIO 作为对象存储解决方案，用于存储音乐文件和相关数据。

- **Redis**：使用 Redis 作为缓存和消息队列，提高数据访问速度和系统性能。

- **MySQL**：采用 MySQL 作为关系型数据库，存储用户信息和音乐元数据。

## 自定义日志打印持久化工具类
本项目实现了一个自定义的日志打印持久化中间件，能够将日志信息持久化到指定位置，便于后期的监控和排错

## 使用Docker进行容器化部署
- 分别为三个项目编写dockerfile，通过docker-compose编排，实现云服务器上的容器化部署
