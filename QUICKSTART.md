# 晨夕 (DawnDusk) - 快速开始指南

## 🚀 5 分钟快速启动

### 方式一：使用启动脚本（推荐）

**Windows**:
```bash
start.bat
```

**macOS/Linux**:
```bash
./start.sh
```

脚本会自动：
- ✅ 检查环境（Go, Flutter, Docker）
- ✅ 启动数据库（PostgreSQL + Redis）
- ✅ 配置环境变量
- ✅ 安装所有依赖

### 方式二：手动启动

#### 1. 启动数据库

```bash
docker-compose up -d
```

#### 2. 启动后端

```bash
cd backend
go run cmd/server/main.go
```

后端运行在: `http://localhost:8080`

#### 3. 启动前端

```bash
cd mobile
flutter run
```

## 📱 测试应用

### 1. 注册新用户

打开应用 → 点击"立即注册" → 填写信息 → 注册

### 2. 登录

输入邮箱和密码 → 点击"登录"

### 3. 查看首页

登录成功后会看到欢迎页面

## 🔧 配置说明

### 后端配置

编辑 `backend/.env`:

```bash
# 数据库
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=dawndusk

# JWT 密钥（生产环境请修改）
JWT_SECRET=your-secret-key-change-in-production
```

### 前端配置

编辑 `mobile/lib/core/constants/api_constants.dart`:

```dart
// iOS 模拟器
static const String baseUrl = 'http://localhost:8080/api/v1';

// Android 模拟器
static const String baseUrl = 'http://10.0.2.2:8080/api/v1';

// 真机（替换为你的 IP）
static const String baseUrl = 'http://192.168.1.100:8080/api/v1';
```

## 🧪 测试 API

### 使用 curl

```bash
# 健康检查
curl http://localhost:8080/health

# 注册用户
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'

# 登录
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

### 使用 Postman

1. 导入 API 文档: `shared/docs/API.md`
2. 设置 Base URL: `http://localhost:8080/api/v1`
3. 测试注册和登录接口

## 📚 文档导航

| 文档 | 说明 |
|------|------|
| [FINAL_SUMMARY.md](FINAL_SUMMARY.md) | 📊 完整项目总结 |
| [FLUTTER_GUIDE.md](FLUTTER_GUIDE.md) | 📱 Flutter 使用指南 |
| [shared/docs/API.md](shared/docs/API.md) | 🔌 API 接口文档 |
| [backend/README.md](backend/README.md) | 🔧 后端开发文档 |
| [mobile/README.md](mobile/README.md) | 📱 前端开发文档 |

## ❓ 常见问题

### 1. 无法连接数据库

**问题**: 后端启动失败，提示数据库连接错误

**解决**:
```bash
# 检查 Docker 容器状态
docker ps

# 重启数据库
docker-compose restart

# 查看日志
docker-compose logs
```

### 2. Flutter 无法连接后端

**问题**: 登录/注册时提示网络错误

**解决**:
- 确保后端已启动: `curl http://localhost:8080/health`
- Android 模拟器使用 `10.0.2.2` 而不是 `localhost`
- 真机使用电脑的局域网 IP

### 3. 依赖安装失败

**后端**:
```bash
cd backend
go clean -modcache
go mod download
```

**前端**:
```bash
cd mobile
flutter clean
flutter pub get
```

### 4. 数据库迁移失败

**安装 migrate 工具**:
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

**运行迁移**:
```bash
cd backend
make migrate-up
```

## 🎯 开发流程

### 后端开发

```bash
cd backend

# 运行服务器
go run cmd/server/main.go

# 运行测试
go test -v ./...

# 构建
go build -o bin/server cmd/server/main.go
```

### 前端开发

```bash
cd mobile

# 运行应用
flutter run

# 热重载（运行时按 r）
r

# 运行测试
flutter test

# 代码格式化
flutter format lib/
```

## 📊 项目状态

- ✅ 后端认证系统（100%）
- ✅ 前端认证 UI（100%）
- ✅ 数据库设计（100%）
- ⏳ 打卡功能（0%）
- ⏳ 宠物系统（0%）
- ⏳ 社交功能（0%）

**总体进度**: 第一阶段 80% 完成

## 🎉 下一步

1. **完善打卡功能**: 实现早晚打卡 API 和 UI
2. **实现宠物系统**: 宠物展示和经验值系统
3. **添加单元测试**: 提高代码质量
4. **实现 WebSocket**: 准备实时聊天功能

## 💬 获取帮助

- 📖 查看 [FINAL_SUMMARY.md](FINAL_SUMMARY.md) 了解完整项目信息
- 📖 查看 [FLUTTER_GUIDE.md](FLUTTER_GUIDE.md) 了解 Flutter 开发
- 📖 查看 [shared/docs/API.md](shared/docs/API.md) 了解 API 接口

---

**祝开发顺利！** 🚀
