# DawnDusk 项目启动脚本
# PowerShell 版本

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  晨夕 (DawnDusk) 项目启动" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# 检查 Docker
Write-Host "[1/5] 检查 Docker..." -ForegroundColor Yellow
try {
    $dockerVersion = docker --version 2>&1
    Write-Host "[OK] Docker 已安装: $dockerVersion" -ForegroundColor Green
} catch {
    Write-Host "[错误] 未找到 Docker，请安装 Docker Desktop" -ForegroundColor Red
    Write-Host "下载地址: https://www.docker.com/products/docker-desktop" -ForegroundColor Yellow
    pause
    exit 1
}

# 检查 Go
Write-Host "[2/5] 检查 Go..." -ForegroundColor Yellow
try {
    $goVersion = go version 2>&1
    Write-Host "[OK] Go 已安装: $goVersion" -ForegroundColor Green
} catch {
    Write-Host "[错误] 未找到 Go，请安装 Go 1.22+" -ForegroundColor Red
    Write-Host "下载地址: https://golang.org/dl/" -ForegroundColor Yellow
    pause
    exit 1
}

# 检查 Flutter
Write-Host "[3/5] 检查 Flutter..." -ForegroundColor Yellow
try {
    $flutterVersion = flutter --version 2>&1 | Select-Object -First 1
    Write-Host "[OK] Flutter 已安装: $flutterVersion" -ForegroundColor Green
    $flutterAvailable = $true
} catch {
    Write-Host "[警告] 未找到 Flutter，前端将无法启动" -ForegroundColor Yellow
    Write-Host "下载地址: https://flutter.dev/docs/get-started/install" -ForegroundColor Yellow
    $flutterAvailable = $false
}

# 启动数据库
Write-Host "[4/5] 启动数据库..." -ForegroundColor Yellow
Set-Location $PSScriptRoot
docker-compose up -d
if ($LASTEXITCODE -ne 0) {
    Write-Host "[错误] 数据库启动失败" -ForegroundColor Red
    pause
    exit 1
}
Write-Host "[OK] 数据库已启动" -ForegroundColor Green

# 等待数据库就绪
Write-Host "等待数据库就绪..." -ForegroundColor Yellow
Start-Sleep -Seconds 10

# 设置后端
Write-Host "[5/5] 配置后端..." -ForegroundColor Yellow
Set-Location backend
if (-not (Test-Path .env)) {
    Write-Host "创建 .env 文件..." -ForegroundColor Yellow
    Copy-Item .env.example .env
}
Write-Host "[OK] 后端配置完成" -ForegroundColor Green

# 设置 Flutter 镜像（可选）
if ($flutterAvailable) {
    Write-Host ""
    Write-Host "提示: 如果 Flutter 依赖下载慢，可以设置国内镜像:" -ForegroundColor Cyan
    Write-Host '  $env:PUB_HOSTED_URL="https://pub.flutter-io.cn"' -ForegroundColor Gray
    Write-Host '  $env:FLUTTER_STORAGE_BASE_URL="https://storage.flutter-io.cn"' -ForegroundColor Gray
}

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  启动完成！" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "下一步操作:" -ForegroundColor Yellow
Write-Host ""
Write-Host "1. 启动后端 (新终端):" -ForegroundColor White
Write-Host "   cd backend" -ForegroundColor Gray
Write-Host "   go run cmd/server/main.go" -ForegroundColor Gray
Write-Host ""
Write-Host "2. 启动前端 (新终端):" -ForegroundColor White
Write-Host "   cd mobile" -ForegroundColor Gray
Write-Host "   flutter run" -ForegroundColor Gray
Write-Host ""
Write-Host "3. 测试 API:" -ForegroundColor White
Write-Host "   curl http://localhost:8080/health" -ForegroundColor Gray
Write-Host ""
Write-Host "文档:" -ForegroundColor Yellow
Write-Host "- 快速开始: QUICKSTART.md" -ForegroundColor Gray
Write-Host "- 故障排查: TROUBLESHOOTING.md" -ForegroundColor Gray
Write-Host "- API 文档: shared/docs/API.md" -ForegroundColor Gray
Write-Host ""

Set-Location $PSScriptRoot
pause
