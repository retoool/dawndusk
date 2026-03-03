# 🚀 Flutter 依赖下载加速指南

## 问题说明

Flutter 首次下载依赖时会很慢（可能需要 5-10 分钟），这是因为：
1. 需要下载大量的 Dart 包
2. 默认使用国外镜像源
3. 某些包体积较大

## 解决方案

### 方案一：使用国内镜像（推荐）

在 PowerShell 中设置环境变量：

```powershell
# 临时设置（当前会话有效）
$env:PUB_HOSTED_URL="https://pub.flutter-io.cn"
$env:FLUTTER_STORAGE_BASE_URL="https://storage.flutter-io.cn"

# 然后运行
flutter pub get
```

### 方案二：永久配置镜像

**Windows 系统环境变量**：
1. 右键"此电脑" → "属性" → "高级系统设置" → "环境变量"
2. 在"用户变量"中新建：
   - 变量名：`PUB_HOSTED_URL`
   - 变量值：`https://pub.flutter-io.cn`
3. 再新建：
   - 变量名：`FLUTTER_STORAGE_BASE_URL`
   - 变量值：`https://storage.flutter-io.cn`
4. 重启 PowerShell

### 方案三：等待完成（如果已经在下载）

如果已经在下载中，建议：
- ✅ **继续等待**（第一次会慢，后续会快很多）
- ✅ 下载完成后会自动缓存
- ✅ 以后运行 `flutter pub get` 会很快（几秒钟）

## 当前进度说明

```
Resolving dependencies... (40.4s)     ✅ 已完成
Downloading packages... (6:14.7s)     ⏳ 正在下载（这是最慢的步骤）
```

**预计总时间**：
- 首次下载：5-10 分钟
- 使用镜像：2-3 分钟
- 后续更新：10-30 秒

## 快速测试（跳过 Chrome）

如果你想快速测试，可以先不在 Chrome 上运行：

```bash
# 取消当前下载（Ctrl+C）

# 使用镜像加速
$env:PUB_HOSTED_URL="https://pub.flutter-io.cn"
$env:FLUTTER_STORAGE_BASE_URL="https://storage.flutter-io.cn"

# 重新下载依赖
cd D:\workspace\DawnDusk\mobile
flutter pub get

# 检查可用设备
flutter devices

# 在 Android 模拟器或真机上运行（比 Chrome 快）
flutter run
```

## 其他加速技巧

### 1. 使用阿里云镜像（备选）

```powershell
$env:PUB_HOSTED_URL="https://mirrors.aliyun.com/pub"
$env:FLUTTER_STORAGE_BASE_URL="https://mirrors.aliyun.com/flutter"
```

### 2. 清理缓存重试

```bash
flutter clean
flutter pub cache repair
flutter pub get
```

### 3. 检查网络

```bash
# 测试网络连接
ping pub.dev
ping pub.flutter-io.cn
```

## 建议

**如果当前正在下载**：
- ✅ 建议继续等待完成（只需要等这一次）
- ✅ 下载完成后会自动缓存
- ✅ 以后就不会这么慢了

**如果想加速**：
1. 按 `Ctrl+C` 取消当前下载
2. 设置国内镜像
3. 重新运行 `flutter pub get`

## 预计时间线

```
当前进度：Downloading packages... (6:14.7s)
预计剩余：3-5 分钟
总计时间：约 10 分钟（首次）
```

**耐心等待，这是正常现象！** 🎉

下载完成后，后续的 `flutter pub get` 只需要几秒钟。
