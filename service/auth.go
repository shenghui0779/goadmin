package service

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/shenghui0779/yiigo"

	"github.com/iiinsomnia/goadmin/consts"
	"github.com/iiinsomnia/goadmin/dao"
	"github.com/iiinsomnia/goadmin/helpers"
	"github.com/iiinsomnia/goadmin/models"
	"github.com/iiinsomnia/goadmin/session"
)

// Login 用户登录
type Login struct {
	ID       string `json:"id" valid:"required"`
	Account  string `json:"account" valid:"required"`
	Password string `json:"password" valid:"required"`
	Captcha  string `json:"captcha" valid:"required"`
}

func (l *Login) Do(c *gin.Context) error {
	// 验证码验证
	if v := base64Captcha.DefaultMemStore.Get(l.ID, true); strings.ToLower(v) != strings.ToLower(l.Captcha) {
		return helpers.Error(helpers.ErrCaptcha)
	}

	userDao := dao.NewUser()

	user, err := userDao.FindByName(l.Account)

	if err != nil {
		return helpers.Error(helpers.ErrLogin, err)
	}

	// 账号密码验证

	fmt.Println("l.Password: ", l.Password)
	fmt.Println("user.Password: ", user.Password)
	if user == nil || l.Password != user.Password {
		return helpers.Error(helpers.ErrAuth)
	}

	// 更新信息
	if err = userDao.UpdateByID(user.ID, yiigo.X{
		"last_login_ip":   c.ClientIP(),
		"last_login_time": time.Now().Unix(),
	}); err != nil {
		return helpers.Error(helpers.ErrLogin, err)
	}

	// 设置session
	if err = session.Set(c, consts.SessionID, &models.Identity{
		ID:   user.ID,
		Name: user.Name,
		Role: user.Role,
	}, consts.SessionDuration); err != nil {
		return helpers.Error(helpers.ErrLogin, err)
	}

	return nil
}

var Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// buildSalt 生成随机加密盐
func buildSalt() string {
	salt := make([]string, 0, 16)
	pattern := "abcdef!ghijklm@nopqrst#uvwxyz$12345%67890^ABCDEFGH&IJKLMNOP*QRSTUVWXYZ"

	l := len(pattern)

	for i := 0; i < 16; i++ {
		n := Rand.Intn(l)
		salt = append(salt, pattern[n:n+1])
	}

	return strings.Join(salt, "")
}
