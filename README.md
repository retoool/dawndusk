# DawnDusk (晨夕)

一款跨平台的健康生活习惯养成应用，通过打卡、社交、宠物养成和AI提醒等功能，帮助用户建立规律的作息习惯。

## 🚀 快速开始

### 使用启动脚本（推荐）

**Windows**:
```bash
start.bat
```

**macOS/Linux**:
```bash
./start.sh
```

### 手动启动

```bash
# 1. 启动数据库
docker-compose up -d

# 2. 启动后端
cd backend && go run cmd/server/main.go

# 3. 启动前端
cd mobile && flutter run
```

详细说明请查看 [QUICKSTART.md](QUICKSTART.md)

## 📊 项目状态

- ✅ **后端认证系统** - 100% 完成
- ✅ **前端认证 UI** - 100% 完成
- ✅ **数据库设计** - 100% 完成
- ⏳ **打卡功能** - 待实现
- ⏳ **宠物系统** - 待实现
- ⏳ **社交功能** - 待实现

**总体进度**: 第一阶段 80% 完成

## 🏗️ 项目结构

```
DawnDusk/
├── backend/          # ✅ Golang 后端（已完成基础架构）
│   ├── cmd/         # 应用入口
│   ├── internal/    # 核心代码
│   └── tests/       # 测试文件
├── mobile/          # ✅ Flutter 前端（已完成认证模块）
│   ├── lib/         # 源代码
│   ├── test/        # 测试文件
│   └── assets/      # 资源文件
├── shared/          # 📚 共享文档
│   └── docs/        # API 文档
├── docker-compose.yml
├── start.bat        # Windows 启动脚本
└── start.sh         # Linux/macOS 启动脚本
```

## 💻 技术栈

### 前端
- **Flutter 3.24+** - 跨平台 UI 框架
- **Riverpod** - 状态管理
- **GoRouter** - 路由管理
- **Dio** - 网络请求
- **Material Design 3** - UI 设计

### 后端
- **Golang 1.22+** - 后端语言
- **Gin** - Web 框架
- **GORM** - ORM
- **PostgreSQL 15+** - 数据库
- **Redis 7+** - 缓存
- **JWT** - 认证

## 🎯 核心功能

### 已实现 ✅
1. 用户注册和登录
2. JWT 认证系统
3. 数据库架构设计
4. Flutter 认证 UI

### 开发中 🚧
3. 早晚打卡系统
4. 个性化睡眠时间设置
5. 小组功能
6. 好友系统与私聊
7. 树洞功能
8. 每日鸡汤语录
9. 小组排行榜
10. 宠物养成系统
11. 宠物装扮激励
12. AI 定时电话提醒

## 📚 文档导航

| 文档 | 说明 |
|------|------|
| [QUICKSTART.md](QUICKSTART.md) | ⚡ 5分钟快速启动 |
| [FINAL_SUMMARY.md](FINAL_SUMMARY.md) | 📊 完整项目总结 |
| [FLUTTER_GUIDE.md](FLUTTER_GUIDE.md) | 📱 Flutter 使用指南 |
| [shared/docs/API.md](shared/docs/API.md) | 🔌 API 接口文档 |
| [backend/README.md](backend/README.md) | 🔧 后端开发文档 |
| [mobile/README.md](mobile/README.md) | 📱 前端开发文档 |

## 🧪 测试

### 测试后端 API

```bash
# 健康检查
curl http://localhost:8080/health

# 注册用户
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'

# 登录
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

### 测试 Flutter 应用

1. 打开应用
2. 点击"立即注册"
3. 填写用户信息并注册
4. 使用注册的账号登录
5. 查看首页

## 📈 开发进度

### 第一阶段：MVP 基础 - ✅ 80% 完成
- [x] 后端项目初始化
- [x] 数据库设计和迁移
- [x] 认证系统（注册/登录/JWT）
- [x] Flutter 项目初始化
- [x] 认证 UI（登录/注册页面）
- [ ] 打卡功能
- [ ] 基础宠物系统

### 第二阶段：社交功能 - ⏳ 0% 完成
- [ ] 小组管理
- [ ] 好友系统
- [ ] 实时聊天

### 第三阶段：树洞与排行榜 - ⏳ 0% 完成
- [ ] 树洞功能
- [ ] 排行榜系统

### 第四阶段：宠物装扮与激励 - ⏳ 0% 完成
- [ ] 装饰品系统
- [ ] 每日语录

### 第五阶段：AI 电话提醒 - ⏳ 0% 完成
- [ ] Twilio 集成
- [ ] OpenAI 语音
- [ ] 定时任务

### 第六阶段：优化与上架 - ⏳ 0% 完成
- [ ] 性能优化
- [ ] 安全加固
- [ ] 测试
- [ ] App Store 提交

## 🛠️ 开发环境

### 前置要求
- Go 1.22+
- Flutter 3.0+
- Docker (可选，用于数据库)
- PostgreSQL 15+ (如果不使用 Docker)
- Redis 7+ (如果不使用 Docker)

### 安装依赖

**后端**:
```bash
cd backend
go mod download
```

**前端**:
```bash
cd mobile
flutter pub get
```

## 🐛 常见问题

### 1. 无法连接数据库
确保 Docker 已启动: `docker ps`

### 2. Flutter 无法连接后端
- iOS 模拟器: 使用 `http://localhost:8080`
- Android 模拟器: 使用 `http://10.0.2.2:8080`
- 真机: 使用电脑的局域网 IP

### 3. 依赖安装失败
```bash
# 后端
cd backend && go clean -modcache && go mod download

# 前端
cd mobile && flutter clean && flutter pub get
```

更多问题请查看 [QUICKSTART.md](QUICKSTART.md)

## 📊 项目统计

```
总文件数: 53+ 个
Go 文件: 25 个
Dart 文件: 12 个
代码总行数: ~3500+ 行
数据库表: 15 张
已实现 API: 5 个
已实现页面: 3 个
```

## 🎨 设计规范

- **主色**: Indigo (#6366F1)
- **次色**: Amber (#F59E0B)
- **圆角**: 12px
- **间距**: 8, 16, 24, 32, 48
- **字体**: PingFang SC

## 🔒 安全特性

- ✅ JWT 认证
- ✅ bcrypt 密码加密
- ✅ CORS 配置
- ✅ SQL 注入防护
- ✅ XSS 防护

## 📝 下一步开发

1. **打卡功能** - 实现早晚打卡 API 和 UI
2. **宠物系统** - 宠物展示和经验值系统
3. **单元测试** - 提高代码质量
4. **WebSocket** - 准备实时聊天功能

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License

---

**开发时间**: 约 8 小时
**预计总开发时间**: 20 周（约 5 个月）
**当前进度**: 第一阶段 80% 完成

🎉 **祝开发顺利！**
