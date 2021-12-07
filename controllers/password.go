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

	if err != nil || identity.Role != consts.Admin {
		Err(c, helpers.Error(helpers.ErrForbid, err))
		return
	}

	s := &service.PasswordReset{
		// ID: helpers.Int64(c.Param("id")),
	}
	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))
		return
	}

	if s.ID == consts.Admin {
		Err(c, helpers.Error(helpers.ErrForbid), "禁止重置")
		return
	}

	fmt.Println("Id: ", s.ID)

	if err := s.Do(); err != nil {
		Err(c, err)

		return
	}

	OK(c)
}
