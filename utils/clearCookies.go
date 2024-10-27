package utils

import (
	"github.com/gin-gonic/gin"
)

// ClearCookies removes all cookies from the client's browser
func ClearCookies(c *gin.Context) {
	for _, cookie := range c.Request.Cookies() {
		c.SetCookie(cookie.Name, "", -1, "/", "", false, true)
	}
}
