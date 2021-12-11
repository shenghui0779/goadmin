package service

import (
	"errors"
	"goadmin/pkg/consts"
	"goadmin/pkg/ent"
	"goadmin/pkg/ent/user"
	"goadmin/pkg/logger"
	"goadmin/pkg/result"
	"goadmin/pkg/service/lib"
	"goadmin/pkg/session"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

type Auth interface {
	Index(c *gin.Context)
	Captcha(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

func NewAuth() Auth {
	return new(auth)
}

type auth struct{}

type ParamsLogin struct {
	ID       string `json:"id" valid:"required"`
	Account  string `json:"account" valid:"required"`
	Password string `json:"password" valid:"required"`
	Captcha  string `json:"captcha" valid:"required"`
}

func (a *auth) Index(c *gin.Context) {
	if identity := session.GetIdentity(c); identity.ID != 0 {
		result.Redirect(c, "/")

		return
	}

	result.Render(c, "login", gin.H{
		"title": "GoAdmin | 登录",
	})
}

type Captcha struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func (a *auth) Captcha(c *gin.Context) {
	captcha := base64Captcha.NewCaptcha(lib.CaptchaDriver, base64Captcha.DefaultMemStore)

	var err error

	resp := new(Captcha)

	resp.ID, resp.Content, err = captcha.Generate()

	if err != nil {
		logger.Err(c.Request.Context(), "err generate captcha", zap.Error(err))
	}

	result.OK(result.Data(resp)).JSON(c)
}

func (a *auth) Login(c *gin.Context) {
	params := new(ParamsLogin)

	if err := c.ShouldBindJSON(params); err != nil {
		result.ErrParams(result.Err(err)).JSON(c)

		return
	}

	ctx := c.Request.Context()

	// 验证码验证
	if v := strings.ToLower(base64Captcha.DefaultMemStore.Get(params.ID, true)); v != strings.ToLower(params.Captcha) {
		logger.Err(ctx, "err captcha verify", zap.String("correct", v))
		result.ErrAuth(result.Err(errors.New("验证码错误"))).JSON(c)

		return
	}

	record, err := ent.DB.User.Query().Unique(false).Where(user.Name(params.Account), user.DeletedAt(0)).First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			result.ErrAuth(result.Err(errors.New("用户不存在"))).JSON(c)

			return
		}

		logger.Err(ctx, "err query user", zap.Error(err))

		result.ErrSystem(result.Err(err)).JSON(c)

		return
	}

	if yiigo.MD5(params.Password+record.Salt) != record.Password {
		result.ErrAuth(result.Err(errors.New("密码错误"))).JSON(c)
	}

	_, err = ent.DB.User.Update().Where(user.ID(record.ID)).SetLastLoginAt(time.Now().Unix()).Save(ctx)

	if err != nil {
		logger.Err(ctx, "err update user", zap.Error(err))
	}

	err = session.SetIdentity(c, &session.Identity{
		ID:   record.ID,
		Name: record.Name,
		Role: consts.Role(record.Role),
	}, time.Hour*12)

	if err != nil {
		logger.Err(ctx, "err session set identity", zap.Error(err))
		result.ErrAuth(result.Err(err)).JSON(c)

		return
	}

	result.OK().JSON(c)
}

func (a *auth) Logout(c *gin.Context) {
	if err := session.Destroy(c); err != nil {
		logger.Err(c.Request.Context(), "err destroy session", zap.Error(err))
	}

	result.Redirect(c, "/login")
}
