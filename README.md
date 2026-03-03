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
- ✅ **打卡功能** - 100% 完成
- ✅ **宠物系统** - 100% 完成
- ✅ **睡眠时间设置** - 100% 完成
- ✅ **小组功能** - 100% 完成
- ✅ **消息系统** - 100% 完成
- ✅ **装饰品系统** - 100% 完成
- ✅ **好友系统** - 100% 完成

**总体进度**: 第二阶段 100% 完成

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
1. **用户认证系统**
   - 用户注册和登录
   - JWT 认证（access + refresh token）
   - 用户信息管理

2. **打卡功能**
   - 早晚打卡记录
   - 打卡历史查看
   - 打卡统计（连续天数、准时率等）
   - 心情记录和备注

3. **宠物养成系统**
   - 宠物创建和展示
   - 经验值和等级系统
   - 健康值和快乐值
   - 宠物重命名

4. **装饰品系统**
   - 装饰品目录浏览
   - 稀有度分类（普通、稀有、史诗、传说）
   - 装饰品解锁和装备
   - 已拥有装饰品管理

5. **小组功能**
   - 创建和加入小组
   - 小组详情和成员管理
   - 邀请码分享
   - 退出和解散小组

6. **消息系统**
   - 私聊功能
   - 消息列表和会话管理
   - 未读消息提醒
   - 消息已读状态

7. **设置功能**
   - 睡眠时间设置
   - 用户信息修改
   - 时区设置

8. **好友系统**
   - 发送和接收好友请求
   - 接受或拒绝好友请求
   - 好友列表管理
   - 删除好友功能
   - 请求状态追踪

### 开发中 🚧
9. 树洞功能
10. 每日鸡汤语录
11. 小组排行榜
12. AI 定时电话提醒
13. WebSocket 实时通信

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

### 第一阶段：MVP 基础 - ✅ 100% 完成
- [x] 后端项目初始化
- [x] 数据库设计和迁移
- [x] 认证系统（注册/登录/JWT）
- [x] Flutter 项目初始化
- [x] 认证 UI（登录/注册页面）
- [x] 打卡功能（前后端完整实现）
- [x] 基础宠物系统（前后端完整实现）
- [x] 睡眠时间设置

### 第二阶段：社交功能 - ✅ 100% 完成
- [x] 小组管理（创建、加入、退出、解散）
- [x] 小组详情和成员展示
- [x] 消息系统（私聊、会话列表）
- [x] 消息已读状态
- [x] 好友系统（请求、接受、拒绝、删除）
- [ ] WebSocket 实时通信

### 第三阶段：装饰品与激励 - ✅ 80% 完成
- [x] 装饰品系统（浏览、解锁、装备）
- [x] 稀有度分类
- [ ] 树洞功能
- [ ] 排行榜系统
- [ ] 每日语录

### 第四阶段：AI 电话提醒 - ⏳ 0% 完成
- [ ] Twilio 集成
- [ ] OpenAI 语音
- [ ] 定时任务

### 第五阶段：优化与上架 - ⏳ 0% 完成
- [ ] 性能优化
- [ ] 安全加固
- [ ] 单元测试
- [ ] 集成测试
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
总文件数: 100+ 个
Go 文件: 45+ 个
Dart 文件: 40+ 个
代码总行数: ~11000+ 行
数据库表: 16 张
已实现 API: 50+ 个
已实现页面: 18+ 个
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

1. **WebSocket 实时通信** - 实时消息推送和在线状态
2. **树洞功能** - 匿名发帖、浏览、点赞
3. **排行榜系统** - 个人排行、小组排行
4. **每日语录** - 语录推送和分享
5. **AI 电话提醒** - Twilio + OpenAI 集成
6. **单元测试** - 提高代码质量和覆盖率

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License

---

**开发时间**: 约 8 小时
**预计总开发时间**: 20 周（约 5 个月）
**当前进度**: 第一阶段 80% 完成

🎉 **祝开发顺利！**
