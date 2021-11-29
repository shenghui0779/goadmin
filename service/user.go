package service

import (
	"fmt"

	"github.com/iiinsomnia/goadmin/consts"
	"github.com/iiinsomnia/goadmin/dao"
	"github.com/iiinsomnia/goadmin/helpers"
	"github.com/iiinsomnia/goadmin/reply"

	"time"

	"github.com/shenghui0779/yiigo"
)

type UserList struct {
	Page int    `json:"page" valid:"required"`
	Size int    `json:"size" valid:"required"`
	Name string `json:"name"`
	Role int    `json:"role"`
}

func (u *UserList) Do() (*reply.UserListReply, error) {
	query := &dao.UserQuery{
		Name:  u.Name,
		Role:  u.Role,
		Page:  u.Page,
		Limit: u.Size,
	}

	if u.Page == 0 {
		query.Page = 1
	}

	if u.Size == 0 {
		query.Limit = 10
	}

	userDao := dao.NewUser()

	var count int64

	if u.Page == 1 {
		v, err := userDao.CountList(query)

		if err != nil {
			return nil, helpers.Error(helpers.ErrSystem, err)
		}

		count = v
	}

	users, err := userDao.FindList(query)

	if err != nil {
		return nil, helpers.Error(helpers.ErrSystem, err)
	}

	l := len(users)

	resp := &reply.UserListReply{
		Count: count,
		List:  make([]*reply.UserListItem, 0, l),
	}

	if l == 0 {
		return resp, nil
	}

	for _, v := range users {
		item := &reply.UserListItem{
			ID:            v.ID,
			Name:          v.Name,
			EMail:         v.Email,
			Role:          v.Role,
			LastLoginIP:   v.LastLoginIP,
			LastLoginTime: "-",
			CreatedAt:     yiigo.Date(v.CreatedAt),
			UpdatedAt:     yiigo.Date(v.UpdatedAt),
		}

		if v.LastLoginTime != 0 {
			item.LastLoginTime = yiigo.Date(v.LastLoginTime)
		}

		switch v.Role {
		case consts.SuperManager:
			item.RoleName = "xxxxxxx"
		case consts.SeniorManager:
			item.RoleName = "高级管理员"
		case consts.GeneralManger:
			item.RoleName = "普通管理员"
		}

		resp.List = append(resp.List, item)
	}

	return resp, nil
}

type UserAdd struct {
	Name  string `json:"name" valid:"required"`
	EMail string `json:"email" valid:"required"`
	Role  int    `json:"role" valid:"required"`
}

func (u *UserAdd) Do() error {
	defaultPass := yiigo.Env("app.default_pass").String("123")

	salt := buildSalt()

	now := time.Now().Unix()

	data := &dao.UserAddData{
		Name:      u.Name,
		EMail:     u.EMail,
		Role:      u.Role,
		Password:  yiigo.MD5(defaultPass + salt),
		Salt:      salt,
		CreatedAt: now,
		UpdatedAt: now,
	}

	userDao := dao.NewUser()

	if err := userDao.Add(data); err != nil {
		return helpers.Error(helpers.ErrOpt, err)
	}

	return nil
}

type UserEdit struct {
	ID    int64  `json:"id" valid:"required"`
	Name  string `json:"name" valid:"required"`
	EMail string `json:"email" valid:"required"`
	Role  int    `json:"role" valid:"required"`
}

func (u *UserEdit) Do() error {
	data := yiigo.X{
		"name":       u.Name,
		"email":      u.EMail,
		"role":       u.Role,
		"updated_at": time.Now().Unix(),
	}

	fmt.Println(data)

	userDao := dao.NewUser()

	if err := userDao.UpdateByID(u.ID, data); err != nil {
		return helpers.Error(helpers.ErrSystem, err)
	}

	return nil
}

type UserDelete struct {
	ID int64 `json:"id"`
}

func (u *UserDelete) Do() error {
	userDao := dao.NewUser()

	if err := userDao.DeleteByID(u.ID); err != nil {
		return helpers.Error(helpers.ErrSystem, err)
	}

	return nil
}

func CheckUserUnique(name string, id ...int64) (bool, error) {
	userDao := dao.NewUser()

	user, err := userDao.FindForUniqueCheck(name, id...)

	if err != nil {
		return false, err
	}

	if user != nil {
		return false, nil
	}

	return true, nil
}
