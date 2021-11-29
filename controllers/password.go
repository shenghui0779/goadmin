package controllers

import (
	"fmt"

	"github.com/iiinsomnia/goadmin/consts"
	"github.com/iiinsomnia/goadmin/helpers"
	"github.com/iiinsomnia/goadmin/service"

	"github.com/gin-gonic/gin"
)

func PasswordChange(c *gin.Context) {
	if c.Request.Method == "GET" {
		Render(c, "password", gin.H{
			"title": "修改密码",
		})

		return
	}

	s := new(service.PasswordChange)

	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))

		return
	}

	if s.Password != s.Confirm {
		Err(c, helpers.Error(helpers.ErrParams), "密码确认错误")

		return
	}

	identity, err := Identity(c)

	if err != nil {
		Err(c, helpers.Error(helpers.ErrForbid, err))

		return
	}

	s.AuthID = identity.ID

	if err := s.Do(); err != nil {
		Err(c, err)

		return
	}

	OK(c)
}

func PasswordReset(c *gin.Context) {
	identity, err := Identity(c)

	if err != nil || identity.Role != consts.SuperManager {
		Err(c, helpers.Error(helpers.ErrForbid, err))

		return
	}

	s := &service.PasswordReset{
		// ID: helpers.Int64(c.Param("id")),
	}
	c.ShouldBindJSON(s)

	fmt.Println("s:  --------- ", s)
	if err := s.Do(); err != nil {
		Err(c, err)

		return
	}

	OK(c)
}
