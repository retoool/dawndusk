# 🔧 晨夕项目 - 常见问题解决方案

## 📋 目录

1. [数据库连接问题](#数据库连接问题)
2. [Flutter 依赖下载慢](#flutter-依赖下载慢)
3. [启动脚本问题](#启动脚本问题)
4. [端口占用问题](#端口占用问题)
5. [其他常见问题](#其他常见问题)

---

## 1. 数据库连接问题

### 问题：密码认证失败

```
failed SASL auth: FATAL: password authentication failed for user "postgres"
```

### ✅ 解决方案

**已修复**：`.env` 文件中的密码已更新为 `postgres`

如果仍有问题，请检查：

```bash
# 1. 确认 Docker 容器正在运行
docker ps

# 应该看到：
# dawndusk-postgres
# dawndusk-redis

# 2. 如果没有运行，启动容器
docker-compose up -d

# 3. 检查容器日志
docker-compose logs postgres

# 4. 测试数据库连接
docker exec -it dawndusk-postgres psql -U postgres -d dawndusk
```

### 手动修复步骤

如果需要手动修复：

1. 编辑 `backend/.env` 文件
2. 将 `DB_PASSWORD=your_password` 改为 `DB_PASSWORD=postgres`
3. 保存文件
4. 重启后端服务

---

## 2. Flutter 依赖下载慢

### 问题：下载包需要很长时间

```
Downloading packages... (6:14.7s)
```

### ✅ 解决方案

#### 方案 A：使用国内镜像（推荐）

**临时设置**（在 PowerShell 中）：
```powershell
# 设置镜像
$env:PUB_HOSTED_URL="https://pub.flutter-io.cn"
$env:FLUTTER_STORAGE_BASE_URL="https://storage.flutter-io.cn"

# 重新下载
cd D:\workspace\DawnDusk\mobile
flutter pub get
```

**永久设置**：
1. 打开"系统属性" → "环境变量"
2. 新建用户变量：
   - `PUB_HOSTED_URL` = `https://pub.flutter-io.cn`
   - `FLUTTER_STORAGE_BASE_URL` = `https://storage.flutter-io.cn`
3. 重启 PowerShell

#### 方案 B：继续等待

- ✅ 首次下载需要 5-10 分钟（正常）
- ✅ 下载完成后会缓存
- ✅ 以后只需要几秒钟

#### 方案 C：使用阿里云镜像

```powershell
$env:PUB_HOSTED_URL="https://mirrors.aliyun.com/pub"
$env:FLUTTER_STORAGE_BASE_URL="https://mirrors.aliyun.com/flutter"
```

---

## 3. 启动脚本问题

### 问题：start.bat 出现乱码

```
'is' 不是内部或外部命令
'鍒涘缓' 不是内部或外部命令
```

### ✅ 解决方案

**原因**：脚本编码问题导致中文乱码

**临时解决**：手动启动各个服务

```powershell
# 1. 启动数据库
docker-compose up -d

# 2. 启动后端（新终端）
cd backend
go run cmd/server/main.go

# 3. 启动前端（新终端）
cd mobile
flutter run
```

**永久解决**：我会创建一个新的启动脚本

---

## 4. 端口占用问题

### 问题：端口 8080 被占用

```
bind: address already in use
```

### ✅ 解决方案

#### 方案 A：查找并关闭占用进程

```powershell
# 查找占用 8080 端口的进程
netstat -ano | findstr :8080

# 关闭进程（替换 <PID> 为实际进程 ID）
taskkill /PID <PID> /F
```

#### 方案 B：修改端口

编辑 `backend/.env`：
```
SERVER_PORT=8081
```

然后重启后端服务。

---

## 5. 其他常见问题

### 问题：Flutter 无法连接后端

**症状**：登录/注册时提示网络错误

**解决方案**：

1. **确认后端已启动**：
   ```bash
   curl http://localhost:8080/health
   ```

2. **检查 API 地址配置**：

   编辑 `mobile/lib/core/constants/api_constants.dart`：

   ```dart
   // iOS 模拟器
   static const String baseUrl = 'http://localhost:8080/api/v1';

   // Android 模拟器
   static const String baseUrl = 'http://10.0.2.2:8080/api/v1';

   // 真机（替换为你的电脑 IP）
   static const String baseUrl = 'http://192.168.1.100:8080/api/v1';
   ```

3. **查找电脑 IP**：
   ```powershell
   ipconfig
   # 查找 "IPv4 地址"
   ```

### 问题：Go 依赖下载失败

**解决方案**：

```bash
cd backend

# 清理缓存
go clean -modcache

# 删除 go.sum
rm go.sum

# 重新下载
go mod download
go mod tidy
```

### 问题：Docker 容器无法启动

**解决方案**：

```bash
# 停止所有容器
docker-compose down

# 删除卷（会清空数据）
docker-compose down -v

# 重新启动
docker-compose up -d

# 查看日志
docker-compose logs -f
```

---

## 🚀 推荐的启动流程

### 完整启动步骤

```powershell
# 1. 进入项目目录
cd D:\workspace\DawnDusk

# 2. 启动数据库（如果使用 Docker）
docker-compose up -d

# 3. 等待数据库启动（约 10 秒）
timeout /t 10

# 4. 启动后端（新终端）
cd backend
go run cmd/server/main.go

# 5. 等待后端启动（约 5 秒）
# 看到 "Server starting on port 8080" 即可

# 6. 启动前端（新终端）
cd mobile

# 设置镜像（可选，加速下载）
$env:PUB_HOSTED_URL="https://pub.flutter-io.cn"
$env:FLUTTER_STORAGE_BASE_URL="https://storage.flutter-io.cn"

# 运行应用
flutter run
```

### 快速测试后端

```bash
# 健康检查
curl http://localhost:8080/health

# 注册用户
curl -X POST http://localhost:8080/api/v1/auth/register `
  -H "Content-Type: application/json" `
  -d '{\"username\":\"test\",\"email\":\"test@example.com\",\"password\":\"123456\"}'

# 登录
curl -X POST http://localhost:8080/api/v1/auth/login `
  -H "Content-Type: application/json" `
  -d '{\"email\":\"test@example.com\",\"password\":\"123456\"}'
```

---

## 📞 获取帮助

如果以上方案都无法解决问题：

1. 查看详细文档：
   - [QUICKSTART.md](QUICKSTART.md)
   - [FINAL_SUMMARY.md](FINAL_SUMMARY.md)
   - [FLUTTER_GUIDE.md](FLUTTER_GUIDE.md)

2. 检查日志：
   - 后端日志：控制台输出
   - 数据库日志：`docker-compose logs postgres`
   - Flutter 日志：控制台输出

3. 常用调试命令：
   ```bash
   # 检查 Go 版本
   go version

   # 检查 Flutter 版本
   flutter --version

   # 检查 Docker 版本
   docker --version

   # 检查端口占用
   netstat -ano | findstr :8080
   netstat -ano | findstr :5432
   ```

---

## ✅ 当前状态

根据你的情况：

- ✅ **数据库密码问题** - 已修复
- ⏳ **Flutter 依赖下载** - 正在进行中（建议等待完成或使用镜像加速）
- ⚠️ **启动脚本** - 有编码问题，建议手动启动

**建议下一步**：

1. 等待 Flutter 依赖下载完成（或按 Ctrl+C 取消，使用镜像重新下载）
2. 下载完成后运行 `flutter run`
3. 测试登录/注册功能

---

**最后更新**: 2026-03-03
**项目状态**: 第一阶段 80% 完成
