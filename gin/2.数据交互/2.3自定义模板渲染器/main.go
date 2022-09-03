package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func main() {
	router := gin.Default()

	html := template.Must(template.ParseFiles("index.html"))
	router.SetHTMLTemplate(html)
	router.GET("/", Index)

	router.Run(":9000")
}
