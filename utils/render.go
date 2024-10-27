package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func RenderHtml(c *gin.Context, status int, html string, data interface{}) error {
	c.HTML(status, html, gin.H{
		"data": data,
	})
	return nil
}

func Redirect(c *gin.Context, url string, code int) {
	c.Render(-1, render.Redirect{
		Code:     code,
		Location: url,
		Request:  c.Request,
	})
}
