@echo off
chcp 65001 > nul

echo [1/4] 清理旧产物...
if exist dist rmdir /s /q dist
mkdir dist

echo [2/4] 编译 Windows x64...
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -o dist/heritage-go_windows_amd64.exe main.go

echo [3/4] 编译 Linux x64...
set GOOS=linux
set GOARCH=amd64
go build -ldflags="-s -w" -o dist/heritage-go_linux_amd64 main.go

echo [4/4] 编译 Linux ARM64...
set GOOS=linux
set GOARCH=arm64
go build -ldflags="-s -w" -o dist/heritage-go_linux_arm64 main.go

echo.
echo ====== 编译成功！产物已存入 dist/ 目录 ======
pause