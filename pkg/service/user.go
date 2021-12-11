package service

import (
	"errors"
	"goadmin/pkg/consts"
	"goadmin/pkg/ent"
	"goadmin/pkg/ent/predicate"
	"goadmin/pkg/ent/user"
	"goadmin/pkg/helpers"
	"goadmin/pkg/logger"
	"goadmin/pkg/result"
	"goadmin/pkg/service/lib"
	"goadmin/pkg/session"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

type User interface {
	Index(c *gin.Context)
	List(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewUser() User {
	return new(users)
}

type users struct{}

func (u *users) Index(c *gin.Context) {
	result.Render(c, "user", gin.H{"menu": "9"})
}

type ParamsUserList struct {
	Page int    `json:"page" valid:"required"`
	Size int    `json:"size" valid:"required"`
	Name string `json:"name"`
	Role uint8  `json:"role"`
}

type UserListData struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	EMail       string `json:"email"`
	Role        uint8  `json:"role"`
	RoleName    string `json:"role_name"`
	RegistedAt  string `json:"registed_at"`
	LastLoginAt string `json:"last_login_at"`
}

type RespUserList struct {
	Total int64           `json:"total"`
	List  []*UserListData `json:"list"`
}

func (u *users) List(c *gin.Context) {
	params := new(ParamsUserList)

	if err := c.ShouldBindJSON(params); err != nil {
		result.ErrParams(result.Err(err)).JSON(c)

		return
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Size == 0 {
		params.Size = 10
	}

	where := make([]predicate.User, 0)

	if len(params.Name) != 0 {
		where = append(where, user.Name(params.Name))
	}

	if params.Role != 0 {
		where = append(where, user.Role(params.Role))
	}

	where = append(where, user.DeletedAt(0))

	ctx := c.Request.Context()

	builder := ent.DB.User.Query().Unique(false).Where(where...)

	resp := new(RespUserList)

	if params.Page == 1 {
		total, err := builder.Clone().Count(ctx)

		if err != nil {
			logger.Err(ctx, "err count user", zap.Error(err))
			result.ErrSystem(result.Err(err)).JSON(c)

			return
		}

		resp.Total = int64(total)

		if total == 0 {
			resp.List = make([]*UserListData, 0)
			result.OK(result.Data(resp)).JSON(c)

			return
		}
	}

	records, err := builder.Offset((params.Page - 1) * params.Size).Limit(params.Size).All(ctx)

	if err != nil {
		logger.Err(ctx, "err query user", zap.Error(err))
		result.ErrSystem(result.Err(err)).JSON(c)

		return
	}

	resp.List = make([]*UserListData, 0, len(records))

	for _, v := range records {
		item := &UserListData{
			ID:          v.ID,
			Name:        v.Name,
			EMail:       v.Email,
			Role:        v.Role,
			RoleName:    "-",
			RegistedAt:  yiigo.Date(v.RegistedAt),
			LastLoginAt: "-",
		}

		if v.LastLoginAt != 0 {
			item.LastLoginAt = yiigo.Date(v.LastLoginAt)
		}

		switch consts.Role(v.Role) {
		case consts.SuperManager:
			item.RoleName = "超级管理员"
		case consts.SeniorManager:
			item.RoleName = "高级管理员"
		case consts.NormalManger:
			item.RoleName = "普通管理员"
		}

		resp.List = append(resp.List, item)
	}

	result.OK(result.Data(resp)).JSON(c)
}

type ParamsUserCreate struct {
	Name  string `json:"name" valid:"required"`
	EMail string `json:"email" valid:"required"`
	Role  uint8  `json:"role" valid:"required"`
}

func (u *users) Create(c *gin.Context) {
	params := new(ParamsUserCreate)

	if err := c.ShouldBindJSON(params); err != nil {
		result.ErrParams(result.Err(err)).JSON(c)

		return
	}

	identity := session.GetIdentity(c)

	if identity.Role == 0 || consts.Role(identity.Role) != consts.NormalManger {
		result.ErrPerm(result.Err(errors.New("权限不足"))).JSON(c)

		return
	}

	ctx := c.Request.Context()

	records, err := ent.DB.User.Query().Unique(false).Where(user.Name(params.Name), user.DeletedAt(0)).Limit(1).All(ctx)

	if err != nil {
		logger.Err(ctx, "err query user", zap.Error(err))
		result.ErrSystem(result.Err(err)).JSON(c)

		return
	}

	if len(records) != 0 {
		result.ErrParams(result.Err(errors.New("用户名已被使用")))

		return
	}

	salt := lib.GenSalt()

	_, err = ent.DB.User.Create().
		SetName(params.Name).
		SetEmail(params.EMail).
		SetRole(params.Role).
		SetPassword(yiigo.MD5("123456" + salt)).
		SetSalt(salt).
		Save(ctx)

	if err != nil {
		logger.Err(ctx, "err create user", zap.Error(err))
		result.ErrSystem(result.Err(err)).JSON(c)

		return
	}

	result.OK().JSON(c)
}

type ParamsUserUpdate struct {
	EMail string `json:"email" valid:"required"`
	Role  uint8  `json:"role" valid:"required"`
}

func (u *users) Update(c *gin.Context) {
	params := new(ParamsUserUpdate)

	if err := c.ShouldBindJSON(params); err != nil {
		result.ErrParams(result.Err(err)).JSON(c)

		return
	}

	identity := session.GetIdentity(c)

	if identity.Role == 0 || consts.Role(identity.Role) != consts.NormalManger {
		result.ErrPerm(result.Err(errors.New("权限不足"))).JSON(c)

		return
	}

	ctx := c.Request.Context()
	uid := helpers.URLParamInt(c, "uid")

	_, err := ent.DB.User.Query().Unique(false).Where(user.ID(uid), user.DeletedAt(0)).First(ctx)

	if err != nil {
		logger.Err(ctx, "err query user", zap.Error(err))
		result.ErrSystem(result.Err(err)).JSON(c)

		return
	}

	_, err = ent.DB.User.Update().Where(user.ID(uid)).
		SetEmail(params.EMail).
		SetRole(params.Role).
		Save(ctx)

	if err != nil {
		logger.Err(ctx, "err update user", zap.Error(err))
		result.ErrSystem(result.Err(err)).JSON(c)

		return
	}

	result.OK().JSON(c)
}

func (u *users) Delete(c *gin.Context) {
	identity := session.GetIdentity(c)

	if consts.Role(identity.Role) != consts.SuperManager {
		result.ErrPerm(result.Err(errors.New("权限不足"))).JSON(c)

		return
	}

	ctx := c.Request.Context()
	uid := helpers.URLParamInt(c, "uid")

	_, err := ent.DB.User.Query().Unique(false).Where(user.ID(uid), user.DeletedAt(0)).First(ctx)

	if err != nil {
		logger.Err(ctx, "err query user", zap.Error(err))
		result.ErrSystem(result.Err(err)).JSON(c)

		return
	}

	_, err = ent.DB.User.Update().Where(user.ID(uid)).
		SetDeletedAt(time.Now().Unix()).
		Save(ctx)

	if err != nil {
		logger.Err(ctx, "err update user", zap.Error(err))
		result.ErrSystem(result.Err(err)).JSON(c)

		return
	}

	result.OK().JSON(c)
}
