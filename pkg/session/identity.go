package session

import (
	"goadmin/pkg/consts"
	"time"

	"github.com/gin-gonic/gin"
)

const IdentityKey = "goadmin_identity"

type Identity struct {
	ID   int64       `json:"id"`
	Name string      `json:"name"`
	Role consts.Role `json:"role"`
}

func SetIdentity(c *gin.Context, identity *Identity, duration time.Duration) error {
	return Set(c, IdentityKey, identity, duration)
}

func GetIdentity(c *gin.Context) *Identity {
	v, ok := Get(c, IdentityKey)

	if !ok {
		return new(Identity)
	}

	identity, ok := v.(*Identity)

	if !ok {
		return new(Identity)
	}

	return identity
}
