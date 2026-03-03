# Flutter 移动端开发指南

由于当前环境未安装 Flutter，请按照以下步骤手动初始化 Flutter 项目。

## 安装 Flutter

### Windows
1. 下载 Flutter SDK: https://docs.flutter.dev/get-started/install/windows
2. 解压到 `C:\src\flutter`
3. 添加到环境变量 PATH: `C:\src\flutter\bin`
4. 运行 `flutter doctor` 检查环境

### macOS
```bash
brew install flutter
flutter doctor
```

### Linux
```bash
snap install flutter --classic
flutter doctor
```

## 初始化项目

在 `DawnDusk` 目录下运行：

```bash
cd /d/workspace/DawnDusk
flutter create mobile --org com.dawndusk --platforms ios,android
cd mobile
```

## 配置 pubspec.yaml

将以下依赖添加到 `mobile/pubspec.yaml`:

```yaml
name: dawndusk
description: 晨夕 - 健康生活习惯养成应用
version: 1.0.0+1

environment:
  sdk: '>=3.0.0 <4.0.0'

dependencies:
  flutter:
    sdk: flutter

  # 状态管理
  flutter_riverpod: ^2.5.0
  riverpod_annotation: ^2.3.0

  # 路由
  go_router: ^14.0.0

  # 网络请求
  dio: ^5.4.0
  retrofit: ^4.1.0
  json_annotation: ^4.8.0

  # WebSocket
  web_socket_channel: ^2.4.0

  # 本地存储
  flutter_secure_storage: ^9.0.0
  hive: ^2.2.3
  hive_flutter: ^1.1.0

  # 推送通知
  firebase_core: ^2.24.0
  firebase_messaging: ^14.7.0
  flutter_local_notifications: ^16.3.0

  # UI组件
  flutter_svg: ^2.0.9
  cached_network_image: ^3.3.0
  shimmer: ^3.0.0

  # 工具库
  intl: ^0.19.0
  equatable: ^2.0.5
  dartz: ^0.10.1
  logger: ^2.0.2

dev_dependencies:
  flutter_test:
    sdk: flutter

  # 代码生成
  build_runner: ^2.4.7
  riverpod_generator: ^2.3.0
  retrofit_generator: ^8.0.0
  json_serializable: ^6.7.0
  hive_generator: ^2.0.1

  # 测试
  mockito: ^5.4.4
  integration_test:
    sdk: flutter

  # 代码质量
  flutter_lints: ^3.0.1

flutter:
  uses-material-design: true
  assets:
    - assets/images/
    - assets/icons/
```

## 创建项目结构

```bash
cd mobile
mkdir -p lib/core/{router,theme,constants,utils}
mkdir -p lib/features/{auth,checkin,pet,group,friends,chat}/{data,domain,presentation}
mkdir -p lib/shared/{widgets,services,models}
mkdir -p lib/l10n
mkdir -p assets/{images,icons,fonts}
mkdir -p test/{unit,widget,integration}
```

## 配置 API 端点

创建 `lib/core/constants/api_constants.dart`:

```dart
class ApiConstants {
  static const String baseUrl = 'http://localhost:8080/api/v1';

  // Auth endpoints
  static const String register = '/auth/register';
  static const String login = '/auth/login';
  static const String refresh = '/auth/refresh';
  static const String logout = '/auth/logout';

  // User endpoints
  static const String userMe = '/users/me';

  // Check-in endpoints
  static const String checkIns = '/check-ins';
  static const String checkInsToday = '/check-ins/today';
  static const String checkInsStats = '/check-ins/stats';

  // Pet endpoints
  static const String pet = '/pet';
  static const String petDecorations = '/pet/decorations';
}
```

## 安装依赖

```bash
flutter pub get
```

## 运行应用

### iOS
```bash
flutter run -d ios
```

### Android
```bash
flutter run -d android
```

### 模拟器
```bash
# 列出可用设备
flutter devices

# 在特定设备上运行
flutter run -d <device-id>
```

## 开发工具

### VS Code 插件
- Flutter
- Dart
- Flutter Riverpod Snippets

### Android Studio 插件
- Flutter
- Dart

## 下一步

1. 实现认证 UI (登录/注册页面)
2. 配置 Riverpod 状态管理
3. 实现 API 服务层
4. 创建打卡功能 UI
5. 实现宠物展示页面

## 常见问题

### Flutter Doctor 报错
运行 `flutter doctor -v` 查看详细信息，根据提示安装缺失的依赖。

### iOS 开发
需要 Xcode 和 CocoaPods:
```bash
sudo gem install cocoapods
cd ios && pod install
```

### Android 开发
需要 Android Studio 和 Android SDK。

## 参考资源

- [Flutter 官方文档](https://docs.flutter.dev/)
- [Riverpod 文档](https://riverpod.dev/)
- [Go Router 文档](https://pub.dev/packages/go_router)
- [Dio 文档](https://pub.dev/packages/dio)
