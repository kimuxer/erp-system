# 创建后端目录结构
$backendDirs = @(
    "backend/cmd",
    "backend/config",
    "backend/internal/controller",
    "backend/internal/service",
    "backend/internal/repository",
    "backend/internal/model",
    "backend/internal/middleware",
    "backend/internal/router",
    "backend/pkg/auth",
    "backend/pkg/database",
    "backend/pkg/util",
    "backend/migrations",
    "backend/docs",
    "backend/logs"
)

# 创建前端目录结构
$frontendDirs = @(
    "frontend/app/(dashboard)",
    "frontend/app/(auth)/login",
    "frontend/app/(auth)/register",
    "frontend/app/api",
    "frontend/components/core",
    "frontend/components/ui",
    "frontend/lib/redux",
    "frontend/lib/api",
    "frontend/styles/theme",
    "frontend/public"
)

# 其他基础设施目录
$infraDirs = @(
    "deployments/docker",
    "deployments/vercel",
    "deployments/nginx",
    "scripts/db",
    "scripts/ci",
    "tests/__tests__",
    "tests/integration_tests",
    ".github/workflows"
)

# 合并所有目录路径
$allDirs = $backendDirs + $frontendDirs + $infraDirs

# 创建目录
foreach ($dir in $allDirs) {
    New-Item -Path "D:\ERP_SYSTEM\$dir" -ItemType Directory -Force
}

Write-Host "目录结构创建完成！使用 tree /F 命令查看完整结构" 