# Flutter Mobile App

晨夕移动端应用 - 使用 Flutter 构建

## 项目结构

```
mobile/
├── lib/
│   ├── main.dart           # 应用入口
│   ├── app.dart            # App根组件
│   ├── core/               # 核心功能
│   │   ├── router/         # 路由配置
│   │   ├── theme/          # 主题配置
│   │   └── constants/      # 常量定义
│   ├── features/           # 功能模块
│   │   ├── auth/          # 认证模块
│   │   └── home/          # 首页模块
│   └── shared/            # 共享组件
├── assets/                # 资源文件
├── test/                  # 测试文件
└── pubspec.yaml          # 依赖配置
```

## 快速开始

### 前置要求

- Flutter 3.0+
- Dart 3.0+

### 安装依赖

```bash
flutter pub get
```

### 运行应用

```bash
# iOS
flutter run -d ios

# Android
flutter run -d android

# 选择设备
flutter devices
flutter run -d <device-id>
```

### 配置后端地址

修改 `lib/core/constants/api_constants.dart` 中的 `baseUrl`:

```dart
static const String baseUrl = 'http://your-backend-url:8080/api/v1';
```

**注意**:
- iOS 模拟器使用 `http://localhost:8080/api/v1`
- Android 模拟器使用 `http://10.0.2.2:8080/api/v1`
- 真机使用你的电脑 IP 地址，如 `http://192.168.1.100:8080/api/v1`

## 已实现功能

### 认证模块 ✅
- 用户注册
- 用户登录
- 状态管理（Riverpod）
- 表单验证
- 错误处理

### UI 组件 ✅
- 主题配置（亮色/暗色）
- 自定义颜色系统
- 响应式布局

## 开发中功能

- 打卡功能
- 宠物系统
- 小组功能
- 好友系统
- 聊天功能

## 技术栈

- **状态管理**: Riverpod
- **路由**: GoRouter
- **网络请求**: Dio
- **本地存储**: Hive + SecureStorage
- **UI**: Material Design 3

## 开发命令

```bash
# 运行应用
flutter run

# 热重载
r (在运行时按 r)

# 热重启
R (在运行时按 R)

# 构建 APK
flutter build apk

# 构建 iOS
flutter build ios

# 运行测试
flutter test

# 代码生成
flutter pub run build_runner build

# 清理
flutter clean
```

## 目录说明

- `lib/core/`: 核心功能（路由、主题、常量）
- `lib/features/`: 按功能模块组织的代码
- `lib/shared/`: 共享组件和工具
- `assets/`: 图片、字体等资源文件
- `test/`: 测试文件

## 注意事项

1. 确保后端服务已启动
2. 检查网络权限配置
3. iOS 需要配置 Info.plist 的网络权限
4. Android 需要配置 AndroidManifest.xml 的网络权限

## 常见问题

### 无法连接后端

1. 检查后端是否运行在 `http://localhost:8080`
2. Android 模拟器使用 `10.0.2.2` 代替 `localhost`
3. 真机需要使用电脑的局域网 IP

### 依赖安装失败

```bash
flutter clean
flutter pub get
```

### 构建失败

```bash
flutter clean
flutter pub get
flutter run
```
