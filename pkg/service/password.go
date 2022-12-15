package service

import (
	"errors"
	"time"

	"goadmin/pkg/consts"
	"goadmin/pkg/ent"
	"goadmin/pkg/ent/user"
	"goadmin/pkg/lib"
	"goadmin/pkg/logger"
	"goadmin/pkg/result"
	"goadmin/pkg/session"

	"github.com/gin-gonic/gin"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

type Password interface {
	Index(c *gin.Context)
	Change(c *gin.Context)
	Reset(c *gin.Context)
}

func NewPassword() Password {
	return new(password)
}

type password struct{}

type ParamsPasswordChange struct {
	Password string `json:"password" valid:"required"`
	Confirm  string `json:"confirm" valid:"required"`
}

func (p *password) Index(c *gin.Context) {
	result.Render(c, "password", gin.H{
		"title": "修改密码",
	})
}

func (p *password) Change(c *gin.Context) {
	params := new(ParamsPasswordChange)

	if err := c.ShouldBindJSON(params); err != nil {
		result.ErrParams(result.Err(err)).JSON(c)

		return
	}

	if params.Confirm != params.Password {
		result.ErrParams(result.Err(errors.New("密码确认错误"))).JSON(c)

		return
	}

	ctx := c.Request.Context()

	identity := session.GetIdentity(c)
	salt := lib.Nonce()

	_, err := ent.DB.User.Update().Where(user.ID(identity.ID)).
		SetPassword(yiigo.MD5(params.Password + salt)).
		SetSalt(salt).
		SetUpdatedAt(time.Now().Unix()).
		Save(ctx)

	if err != nil {
		logger.Err(ctx, "err update user", zap.Error(err))
		result.ErrSystem(result.Err(err)).JSON(c)

		return
	}

	result.OK().JSON(c)
}

func (p *password) Reset(c *gin.Context) {
	identity := session.GetIdentity(c)

	if consts.Role(identity.Role) != consts.SuperManager {
		result.ErrPerm(result.Err(errors.New("权限不足"))).JSON(c)

		return
	}

	ctx := c.Request.Context()

	uid := lib.URLParamInt(c, "uid")
	salt := lib.Nonce()

	_, err := ent.DB.User.Update().Where(user.ID(uid)).
		SetPassword(yiigo.MD5("123456" + salt)).
		SetSalt(salt).
		SetUpdatedAt(time.Now().Unix()).
		Save(ctx)

	if err != nil {
		logger.Err(ctx, "err update user", zap.Error(err))
		result.ErrSystem(result.Err(err)).JSON(c)

		return
	}

	result.OK().JSON(c)
}
