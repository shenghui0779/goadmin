package controllers

import (
	"fmt"

	"github.com/iiinsomnia/goadmin/consts"
	"github.com/iiinsomnia/goadmin/helpers"
	"github.com/iiinsomnia/goadmin/service"

	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	Render(c, "user", gin.H{"menu": "9"})
}

func UserQuery(c *gin.Context) {
	s := new(service.UserList)

	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))

		return
	}

	data, err := s.Do()

	if err != nil {
		Err(c, err)

		return
	}

	OK(c, data)
}

func UserAdd(c *gin.Context) {
	s := new(service.UserAdd)

	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))
		return
	}

	identity, err := Identity(c)

	if err != nil || identity.Role != consts.Admin {
		Err(c, helpers.Error(helpers.ErrForbid, err))
		return
	}

	unique, err := service.CheckUserUnique(s.Name)

	if err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))

		return
	}

	if !unique {
		Err(c, helpers.Error(helpers.ErrParams), "用户名已被使用")

		return
	}

	if err = s.Do(); err != nil {
		Err(c, err)

		return
	}

	OK(c)
}

func UserEdit(c *gin.Context) {
	s := new(service.UserEdit)

	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))

		return
	}

	identity, err := Identity(c)

	if err != nil || identity.Role != consts.Admin {
		Err(c, helpers.Error(helpers.ErrForbid, err))

		return
	}

	if s.ID != consts.Admin && s.ID != consts.User {
		Err(c, helpers.Error(helpers.ErrParams, err))
		return
	}

	unique, err := service.CheckUserUnique(s.Name, s.ID)

	if err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))

		return
	}

	if !unique {
		Err(c, helpers.Error(helpers.ErrParams), "用户名已被使用")

		return
	}

	if s.Name == "admin" {
		Err(c, helpers.Error(helpers.ErrForbid), "禁止修改")
		return
	}

	if err = s.Do(); err != nil {
		Err(c, err)

		return
	}

	OK(c)
}

func UserDelete(c *gin.Context) {
	identity, err := Identity(c)

	if err != nil || identity.Role != consts.Admin {
		Err(c, helpers.Error(helpers.ErrForbid, err))

		return
	}

	// s := &service.UserDelete{ID: helpers.Int64(c.Param("id"))}
	s := &service.UserDelete{}
	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))
		return
	}

	if s.ID == consts.Admin {
		Err(c, helpers.Error(helpers.ErrForbid), "禁止删除")
		return
	}

	fmt.Println("UserDelete: ", s)
	if err := s.Do(); err != nil {
		Err(c, err)

		return
	}

	OK(c)
}
