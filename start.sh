#!/bin/bash

# 晨夕项目启动脚本

echo "======================================"
echo "  晨夕 (DawnDusk) 项目启动脚本"
echo "======================================"
echo ""

# 检查 Go 是否安装
if ! command -v go &> /dev/null; then
    echo "❌ Go 未安装，请先安装 Go 1.22+"
    exit 1
fi
echo "✅ Go 已安装: $(go version)"

# 检查 Flutter 是否安装
if ! command -v flutter &> /dev/null; then
    echo "❌ Flutter 未安装，请先安装 Flutter 3.0+"
    exit 1
fi
echo "✅ Flutter 已安装: $(flutter --version | head -n 1)"

# 检查 Docker 是否安装
if ! command -v docker &> /dev/null; then
    echo "⚠️  Docker 未安装，将无法启动数据库"
    echo "   请手动安装 PostgreSQL 和 Redis"
else
    echo "✅ Docker 已安装"
fi

echo ""
echo "======================================"
echo "  1. 启动数据库服务"
echo "======================================"

if command -v docker &> /dev/null; then
    echo "启动 PostgreSQL 和 Redis..."
    docker-compose up -d

    if [ $? -eq 0 ]; then
        echo "✅ 数据库服务已启动"
    else
        echo "❌ 数据库服务启动失败"
        exit 1
    fi
else
    echo "⚠️  跳过数据库启动（Docker 未安装）"
fi

echo ""
echo "======================================"
echo "  2. 配置后端"
echo "======================================"

cd backend

# 检查 .env 文件
if [ ! -f .env ]; then
    echo "创建 .env 文件..."
    cp .env.example .env
    echo "✅ .env 文件已创建，请根据需要修改配置"
else
    echo "✅ .env 文件已存在"
fi

# 安装依赖
echo "安装 Go 依赖..."
go mod download
echo "✅ Go 依赖已安装"

# 运行数据库迁移
if command -v migrate &> /dev/null; then
    echo "运行数据库迁移..."
    make migrate-up
    echo "✅ 数据库迁移完成"
else
    echo "⚠️  migrate 工具未安装，跳过数据库迁移"
    echo "   请手动运行: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
fi

echo ""
echo "======================================"
echo "  3. 配置前端"
echo "======================================"

cd ../mobile

# 安装依赖
echo "安装 Flutter 依赖..."
flutter pub get
echo "✅ Flutter 依赖已安装"

echo ""
echo "======================================"
echo "  启动完成！"
echo "======================================"
echo ""
echo "📝 下一步操作："
echo ""
echo "1. 启动后端服务器："
echo "   cd backend && make run"
echo "   或: cd backend && go run cmd/server/main.go"
echo ""
echo "2. 启动 Flutter 应用："
echo "   cd mobile && flutter run"
echo ""
echo "3. 测试 API："
echo "   curl http://localhost:8080/health"
echo ""
echo "📚 查看文档："
echo "   - 项目总结: FINAL_SUMMARY.md"
echo "   - Flutter 指南: FLUTTER_GUIDE.md"
echo "   - API 文档: shared/docs/API.md"
echo ""
echo "🎉 祝开发顺利！"
