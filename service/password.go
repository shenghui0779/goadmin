package service

import (
	"time"

	"github.com/iiinsomnia/goadmin/dao"
	"github.com/iiinsomnia/goadmin/helpers"

	"github.com/shenghui0779/yiigo"
)

type PasswordChange struct {
	AuthID   int64
	Password string `json:"password" valid:"required"`
	Confirm  string `json:"confirm" valid:"required"`
}

// ChangePassword ...
func (p *PasswordChange) Do() error {
	salt := buildSalt()

	userDao := dao.NewUser()

	if err := userDao.UpdateByID(p.AuthID, yiigo.X{
		"password":  yiigo.MD5(p.Password + salt),
		"salt":      salt,
		"update_at": time.Now().Unix(),
	}); err != nil {
		return helpers.Error(helpers.ErrSystem, err)
	}

	return nil
}

type PasswordReset struct {
	ID int64 `json:"id" valid:"required"`
}

func (p *PasswordReset) Do() error {
	defaultPass := yiigo.Env("app.default_pass").String("123")
	salt := buildSalt()

	userDao := dao.NewUser()

	if err := userDao.UpdateByID(p.ID, yiigo.X{
		"password":  yiigo.MD5(defaultPass + salt),
		"salt":      salt,
		"update_at": time.Now().Unix(),
	}); err != nil {
		return helpers.Error(helpers.ErrSystem, err)
	}

	return nil
}
