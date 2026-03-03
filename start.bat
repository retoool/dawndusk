@echo off
REM 晨夕项目启动脚本 (Windows)

echo ======================================
echo   晨夕 (DawnDusk) 项目启动脚本
echo ======================================
echo.

REM 检查 Go 是否安装
where go >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo ❌ Go 未安装，请先安装 Go 1.22+
    exit /b 1
)
echo ✅ Go 已安装

REM 检查 Flutter 是否安装
where flutter >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo ❌ Flutter 未安装，请先安装 Flutter 3.0+
    exit /b 1
)
echo ✅ Flutter 已安装

REM 检查 Docker 是否安装
where docker >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo ⚠️  Docker 未安装，将无法启动数据库
    echo    请手动安装 PostgreSQL 和 Redis
) else (
    echo ✅ Docker 已安装
)

echo.
echo ======================================
echo   1. 启动数据库服务
echo ======================================

where docker >nul 2>nul
if %ERRORLEVEL% EQU 0 (
    echo 启动 PostgreSQL 和 Redis...
    docker-compose up -d
    if %ERRORLEVEL% EQU 0 (
        echo ✅ 数据库服务已启动
    ) else (
        echo ❌ 数据库服务启动失败
        exit /b 1
    )
) else (
    echo ⚠️  跳过数据库启动 (Docker 未安装)
)

echo.
echo ======================================
echo   2. 配置后端
echo ======================================

cd backend

REM 检查 .env 文件
if not exist .env (
    echo 创建 .env 文件...
    copy .env.example .env
    echo ✅ .env 文件已创建，请根据需要修改配置
) else (
    echo ✅ .env 文件已存在
)

REM 安装依赖
echo 安装 Go 依赖...
go mod download
echo ✅ Go 依赖已安装

echo.
echo ======================================
echo   3. 配置前端
echo ======================================

cd ..\mobile

REM 安装依赖
echo 安装 Flutter 依赖...
flutter pub get
echo ✅ Flutter 依赖已安装

cd ..

echo.
echo ======================================
echo   启动完成！
echo ======================================
echo.
echo 📝 下一步操作：
echo.
echo 1. 启动后端服务器：
echo    cd backend ^&^& go run cmd/server/main.go
echo.
echo 2. 启动 Flutter 应用：
echo    cd mobile ^&^& flutter run
echo.
echo 3. 测试 API：
echo    curl http://localhost:8080/health
echo.
echo 📚 查看文档：
echo    - 项目总结: FINAL_SUMMARY.md
echo    - Flutter 指南: FLUTTER_GUIDE.md
echo    - API 文档: shared/docs/API.md
echo.
echo 🎉 祝开发顺利！
echo.
pause
