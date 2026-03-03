# 晨夕 (DawnDusk) - Flutter 移动端使用指南

## 🚀 快速开始

由于你已经安装了 Flutter，现在可以直接运行项目了！

### 1. 进入项目目录

```bash
cd D:\workspace\DawnDusk\mobile
```

### 2. 安装依赖

```bash
flutter pub get
```

### 3. 检查设备

```bash
flutter devices
```

### 4. 运行应用

```bash
# 在 iOS 模拟器运行
flutter run -d ios

# 在 Android 模拟器运行
flutter run -d android

# 或者直接运行（会自动选择设备）
flutter run
```

## ⚙️ 配置后端地址

在运行应用之前，需要配置后端 API 地址。

打开 `lib/core/constants/api_constants.dart`，修改 `baseUrl`:

```dart
// 如果使用 iOS 模拟器
static const String baseUrl = 'http://localhost:8080/api/v1';

// 如果使用 Android 模拟器
static const String baseUrl = 'http://10.0.2.2:8080/api/v1';

// 如果使用真机（替换为你的电脑 IP）
static const String baseUrl = 'http://192.168.1.100:8080/api/v1';
```

### 如何查找电脑 IP 地址

**Windows**:
```bash
ipconfig
# 查找 "IPv4 地址"
```

**macOS/Linux**:
```bash
ifconfig
# 或
ip addr show
```

## 📱 已实现的功能

### ✅ 认证系统
- 用户注册页面
- 用户登录页面
- JWT token 管理
- 表单验证
- 错误提示

### ✅ UI/UX
- Material Design 3
- 亮色/暗色主题
- 响应式布局
- 流畅的页面转场

## 🎯 测试流程

### 1. 启动后端服务

```bash
cd D:\workspace\DawnDusk\backend
go run cmd/server/main.go
```

确保后端运行在 `http://localhost:8080`

### 2. 启动 Flutter 应用

```bash
cd D:\workspace\DawnDusk\mobile
flutter run
```

### 3. 测试注册功能

1. 点击"立即注册"
2. 填写用户名、邮箱、密码
3. 点击"注册"按钮
4. 成功后会自动跳转到首页

### 4. 测试登录功能

1. 输入邮箱和密码
2. 点击"登录"按钮
3. 成功后会跳转到首页

## 🛠️ 开发工具

### VS Code 插件（推荐）
- Flutter
- Dart
- Flutter Riverpod Snippets

### Android Studio 插件
- Flutter
- Dart

### 常用命令

```bash
# 热重载（在运行时按 r）
r

# 热重启（在运行时按 R）
R

# 查看日志
flutter logs

# 清理项目
flutter clean

# 重新获取依赖
flutter pub get

# 代码格式化
flutter format lib/

# 分析代码
flutter analyze
```

## 📂 项目结构说明

```
mobile/
├── lib/
│   ├── main.dart                    # 应用入口
│   ├── app.dart                     # App 根组件
│   ├── core/
│   │   ├── router/
│   │   │   └── app_router.dart     # 路由配置
│   │   ├── theme/
│   │   │   ├── app_theme.dart      # 主题配置
│   │   │   └── colors.dart         # 颜色定义
│   │   └── constants/
│   │       └── api_constants.dart  # API 常量
│   └── features/
│       ├── auth/                    # 认证模块
│       │   ├── data/
│       │   │   ├── models/         # 数据模型
│       │   │   └── services/       # API 服务
│       │   └── presentation/
│       │       ├── providers/      # 状态管理
│       │       └── screens/        # 页面
│       └── home/                    # 首页模块
│           └── presentation/
│               └── screens/
```

## 🐛 常见问题

### 1. 无法连接后端

**问题**: 登录/注册时提示网络错误

**解决方案**:
- 确保后端服务已启动
- 检查 `api_constants.dart` 中的 `baseUrl` 配置
- Android 模拟器使用 `10.0.2.2` 而不是 `localhost`
- 真机需要使用电脑的局域网 IP

### 2. 依赖安装失败

**解决方案**:
```bash
flutter clean
flutter pub get
```

### 3. 构建失败

**解决方案**:
```bash
flutter clean
rm -rf pubspec.lock
flutter pub get
flutter run
```

### 4. iOS 模拟器无法连接网络

**解决方案**:
在 `ios/Runner/Info.plist` 中添加:
```xml
<key>NSAppTransportSecurity</key>
<dict>
    <key>NSAllowsArbitraryLoads</key>
    <true/>
</dict>
```

### 5. Android 模拟器无法连接网络

**解决方案**:
在 `android/app/src/main/AndroidManifest.xml` 中添加:
```xml
<uses-permission android:name="android.permission.INTERNET"/>
```

## 📝 下一步开发

### 待实现功能

1. **打卡功能**
   - 早晚打卡 UI
   - 打卡历史
   - 打卡统计

2. **宠物系统**
   - 宠物展示
   - 宠物升级
   - 装饰品系统

3. **社交功能**
   - 小组列表
   - 好友系统
   - 实时聊天

4. **其他功能**
   - 树洞
   - 排行榜
   - 每日语录

## 💡 开发提示

1. **状态管理**: 使用 Riverpod，所有状态都在 `providers` 目录
2. **网络请求**: 使用 Dio，所有 API 调用在 `services` 目录
3. **路由**: 使用 GoRouter，配置在 `app_router.dart`
4. **主题**: 在 `app_theme.dart` 中统一管理
5. **常量**: API 地址等常量在 `constants` 目录

## 🎨 UI 设计规范

- **主色**: Indigo (#6366F1)
- **次色**: Amber (#F59E0B)
- **圆角**: 12px
- **间距**: 8, 16, 24, 32, 48
- **字体大小**: 12, 14, 16, 18, 24, 28, 32

## 📚 参考资源

- [Flutter 官方文档](https://docs.flutter.dev/)
- [Riverpod 文档](https://riverpod.dev/)
- [GoRouter 文档](https://pub.dev/packages/go_router)
- [Dio 文档](https://pub.dev/packages/dio)
- [Material Design 3](https://m3.material.io/)

---

如有问题，请查看项目根目录的 `PROJECT_SUMMARY.md` 或后端的 `README.md`。
