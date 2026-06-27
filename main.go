package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	
	"heritage-go/model"
    "heritage-go/repository"
    "heritage-go/service"
)

//go:embed web/index.html static/*
var embeddedFiles embed.FS

func main() {
	// 1. 初始化本地 SQLite 数据库底座
	repository.InitDB()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 静态资源与主页托管
	staticFS, _ := fs.Sub(embeddedFiles, "static")
	r.StaticFS("/static", http.FS(staticFS))
	r.GET("/", func(c *gin.Context) {
		htmlContent, err := embeddedFiles.ReadFile("web/index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "模板丢失")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", htmlContent)
	})

	// 🚀 路由 API 1：获取真实的家族多叉拓扑树
	r.GET("/api/v1/tree", func(c *gin.Context) {
		var dbMembers []model.Member
		// 从真实的 SQLite 数据库中取出全量扁平数据
		if err := repository.DB.Order("generation asc, id asc").Find(&dbMembers).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "读取谱系失败"})
			return
		}

		// 转义为前端渲染树
		tree := service.BuildFamilyTree(dbMembers)
		c.JSON(http.StatusOK, tree)
	})

	// 🚀 路由 API 2：上传并批量导入 Excel 谱系表
	r.POST("/api/v1/import", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "未检测到上传的文件"})
			return
		}

		// 临时保存到本地准备解析
		tempPath := filepath.Join("storage", "temp_"+file.Filename)
		if err := c.SaveUploadedFile(file, tempPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "缓存文件失败"})
			return
		}
		// 解析完毕后销毁临时文件
		defer func() { _ = os.Remove(tempPath) }() 

		// 调用核心解析服务
		count, err := service.ImportFromExcel(tempPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("成功清洗并批量导入 %d 位家族成员！", count),
		})
	})

	log.Println("华夏谱系开源系统已在安全沙箱中启动：http://localhost:8080")
	r.Run(":8080")
}
