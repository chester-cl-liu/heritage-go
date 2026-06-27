@echo off
chcp 65001 > nul

:: ==========================================
::             统一版本号配置中心
:: ==========================================
set VERSION=v1.0.1
set BIN_NAME=heritage-go

echo ===================================================
echo   HeritageGo 华夏谱系系统 - 自动化编译工具 (%VERSION%)
echo ===================================================
echo.

echo [1/4] 清理旧产物...
if exist dist rmdir /s /q dist
mkdir dist

:: 💡 极客提示：-X 可以将版本号变量动态注入到代码的 main.Version 中
set LDFLAGS="-s -w -X main.Version=%VERSION%"

echo [2/4] 编译 Windows x64...
set GOOS=windows
set GOARCH=amd64
go build -ldflags=%LDFLAGS% -o dist/%BIN_NAME%_%VERSION%_windows_amd64.exe main.go

echo [3/4] 编译 Linux x64...
set GOOS=linux
set GOARCH=amd64
go build -ldflags=%LDFLAGS% -o dist/%BIN_NAME%_%VERSION%_linux_amd64 main.go

echo [4/4] 编译 Linux ARM64...
set GOOS=linux
set GOARCH=arm64
go build -ldflags=%LDFLAGS% -o dist/%BIN_NAME%_%VERSION%_linux_arm64 main.go

echo.
echo ====== 编译成功！产物已带版本号存入 dist/ 目录 ======
pause