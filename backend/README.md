# DawnDusk Backend

晨夕应用后端服务 - 使用 Golang + Gin + PostgreSQL + Redis 构建

## 项目结构

```
backend/
├── cmd/server/          # 应用入口
├── internal/
│   ├── api/            # API层（handlers, middlewares, routes）
│   ├── domain/         # 领域层（entities, repositories, services）
│   ├── infrastructure/ # 基础设施层（database, cache, websocket）
│   ├── shared/         # 共享代码（config, errors, utils）
│   └── dto/            # 数据传输对象
├── tests/              # 测试文件
└── Makefile           # 构建脚本
```

## 快速开始

### 前置要求

- Go 1.22+
- PostgreSQL 15+
- Redis 7+

### 安装依赖

```bash
make install-deps
```

### 配置环境变量

复制 `.env.example` 到 `.env` 并修改配置：

```bash
cp .env.example .env
```

### 运行数据库迁移

```bash
make migrate-up
```

### 启动服务器

```bash
make run
```

服务器将在 `http://localhost:8080` 启动

## API 端点

### 认证

- `POST /api/v1/auth/register` - 用户注册
- `POST /api/v1/auth/login` - 用户登录
- `POST /api/v1/auth/refresh` - 刷新token
- `POST /api/v1/auth/logout` - 登出

### 用户

- `GET /api/v1/users/me` - 获取当前用户信息（需要认证）

### 打卡

- `POST /api/v1/check-ins` - 创建打卡记录（需要认证）
- `GET /api/v1/check-ins` - 获取打卡历史（需要认证）

### 宠物

- `GET /api/v1/pet` - 获取宠物信息（需要认证）
- `POST /api/v1/pet` - 创建宠物（需要认证）

## 开发

### 构建

```bash
make build
```

### 运行测试

```bash
make test
```

### 使用Docker

```bash
make docker-up    # 启动服务
make docker-down  # 停止服务
```

## 技术栈

- **Web框架**: Gin
- **ORM**: GORM
- **数据库**: PostgreSQL
- **缓存**: Redis
- **认证**: JWT
- **WebSocket**: Gorilla WebSocket
