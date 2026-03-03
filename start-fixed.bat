@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion

echo ========================================
echo   DawnDusk Project Startup
echo ========================================
echo.

REM Check Docker
echo [1/5] Checking Docker...
docker --version >nul 2>&1
if errorlevel 1 (
    echo [ERROR] Docker not found. Please install Docker Desktop.
    echo Download: https://www.docker.com/products/docker-desktop
    pause
    exit /b 1
)
echo [OK] Docker is installed

REM Check Go
echo [2/5] Checking Go...
go version >nul 2>&1
if errorlevel 1 (
    echo [ERROR] Go not found. Please install Go 1.22+
    echo Download: https://golang.org/dl/
    pause
    exit /b 1
)
echo [OK] Go is installed

REM Check Flutter
echo [3/5] Checking Flutter...
flutter --version >nul 2>&1
if errorlevel 1 (
    echo [WARNING] Flutter not found. Frontend will not start.
    echo Download: https://flutter.dev/docs/get-started/install
    set FLUTTER_AVAILABLE=0
) else (
    echo [OK] Flutter is installed
    set FLUTTER_AVAILABLE=1
)

REM Start Database
echo [4/5] Starting database...
cd /d "%~dp0"
docker-compose up -d
if errorlevel 1 (
    echo [ERROR] Failed to start database
    pause
    exit /b 1
)
echo [OK] Database started

REM Wait for database
echo Waiting for database to be ready...
timeout /t 10 /nobreak >nul

REM Setup backend
echo [5/5] Setting up backend...
cd backend
if not exist .env (
    echo Creating .env file...
    copy .env.example .env >nul
)

echo.
echo ========================================
echo   Startup Complete!
echo ========================================
echo.
echo Next steps:
echo.
echo 1. Start backend (in new terminal):
echo    cd backend
echo    go run cmd/server/main.go
echo.
echo 2. Start frontend (in new terminal):
echo    cd mobile
echo    flutter run
echo.
echo 3. Test API:
echo    curl http://localhost:8080/health
echo.
echo Documentation:
echo - Quick Start: QUICKSTART.md
echo - Troubleshooting: TROUBLESHOOTING.md
echo - API Docs: shared/docs/API.md
echo.
pause
