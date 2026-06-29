# HeritageGo (华夏谱系开源通用系统)

[![Go Version](https://img.shields.io/github/go-mod/go-version/chester-cl-liu/heritage-go)](https://github.com/chester-cl-liu/heritage-go)
[![Release](https://img.shields.io/github/v/release/chester-cl-liu/heritage-go?color=orange)](https://github.com/chester-cl-liu/heritage-go/releases)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

HeritageGo 是一款专为数字化安全传承华夏家族血脉记忆而设计的开源通用谱系管理系统。系统采用现代化的技术栈，兼顾了极客的本地私有化部署需求与普通长辈“开箱即用”的低门槛体验。

---

## ✨ 核心特性

* **🚀 零依赖单文件分发**：基于 Go 的 `embed` 技术，将前端静态资源、画布引擎全量内嵌至单个二进制文件中。用户无需配置任何 Node.js、Nginx 或 Web 服务器，双击即可运行。
* **💾 磐石本地存储**：后端基于 SQLite3 嵌入式数据库，所有家族谱系数据完全沉淀在本地，数据 100% 私有，彻底杜绝云端泄露风险，守护家族隐私。
* **📊 Excel 拖拽一键清洗**：内置强大的多叉树血脉数据清洗引擎。支持标准的 Excel 谱系表一键拖拽上传，后端自动自动校验、纠错并建立代际关联。
* **🎨 动态高清拓扑画布**：前端基于 AntV G6 高性能图可视化引擎打造。完美支持 600+ 成员的动态降采样，提供全景平滑缩放、拖拽及节点高亮交互，清晰呈现绵延不绝的血脉网络。

---

## 🛠️ 技术栈

* **后端 (Backend)**：Go + Gin Web Framework + GORM
* **数据库 (Database)**：SQLite3 (嵌入式单文件)
* **前端 (Frontend)**：HTML5 / JavaScript + AntV G6 (关系图谱引擎)
* **分发技术 (Embed)**：Go standard library `embed` / `io/fs`

---

## 📦 快速开始 (使用者指南)

### Windows 用户 (推荐)
1. 前往项目的 [Releases 页面](https://github.com/chester-cl-liu/heritage-go/releases) 下载最新版的 `heritage-go_v1.0.1_windows_amd64.zip`。
2. 解压压缩包到任意本地目录。
3. 双击运行 `heritage-go_windows_amd64.exe`。系统会自动在当前目录下初始化 `storage/` 数据库目录。
4. 打开浏览器，访问 `http://localhost:8080`，即可开启您的数字化寻根之旅！

### Linux / 树莓派用户
```bash
# 下载对应架构的二进制文件，以 Linux amd64 为例
wget [https://github.com/chester-cl-liu/heritage-go/releases/download/v1.0.1/heritage-go_linux_amd64](https://github.com/chester-cl-liu/heritage-go/releases/download/v1.0.1/heritage-go_linux_amd64)

# 赋予执行权限
chmod +x heritage-go_linux_amd64

# 后台运行
nohup ./heritage-go_linux_amd64 > stdout.log 2>&1 &

```

---

# 💻 开发者指南 (本地编译)
如果您希望基于本项目进行二次开发，请确保本地已安装 Go 1.25+ 环境。

### 1. 克隆仓库
```bash
git clone [https://github.com/chester-cl-liu/heritage-go.git](https://github.com/chester-cl-liu/heritage-go.git)
cd heritage-go

```
### 2. 整理依赖
```bash
go mod tidy

```
### 3. 一键跨平台编译
项目根目录下内置了专为 Windows 环境编写的自动化跨平台编译脚本。直接在 PowerShell 中运行：
```PowerShell
.\build.bat

```
编译完成后，适用于 Windows x64、Linux x64 以及 Linux ARM64 (树莓派/信创平台) 的瘦身版二进制文件将整齐集结在 dist/ 目录下。

📜 开源协议
本项目基于 MIT License 协议开源，允许自由修改和商业使用，但请保留原作者的版权声明。

👥 贡献与支持
欢迎通过提交 Issue 提报 Bug 或新功能建议，也欢迎直接提交 Pull Request 共同完善华夏谱系系统的建设！
