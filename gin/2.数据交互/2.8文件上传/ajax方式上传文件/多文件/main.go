package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		fmt.Println(file.Filename)
		// 防止文件名冲突，使用时间戳命名
		timer := time.Now().Unix()               // 时间戳，int类型
		timerStr := strconv.FormatInt(timer, 10) // 讲int类型转为string类型，方便拼接，使用sprinf也可以

		// 要先创建好upload目录
		filePath := "upload/" + timerStr + file.Filename // 设置保存文件的路径，不要忘了后面的文件名
		err := ctx.SaveUploadedFile(file, filePath)      // 保存文件
		if err != nil {
			fmt.Println("save file failed, err: ", err.Error())
		}
	}
}

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index/index.html", nil)
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("template/**/*")
	engine.GET("/", Index)
	engine.POST("/upload", Upload)
	engine.Run(":9000")
}
