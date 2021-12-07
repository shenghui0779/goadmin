package controllers

import "github.com/gin-gonic/gin"

func OnePicture(c *gin.Context) {
	Render(c, "one", gin.H{"menu": "9"})
}
