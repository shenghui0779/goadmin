// Package session ...
package session

import (
	"encoding/gob"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/shenghui0779/yiigo"
	"github.com/pkg/errors"

	"github.com/iiinsomnia/goadmin/models"
)

const gosessid = "GOSESSID"

var store *sessions.CookieStore

// Start start session
func Start() {
	store = sessions.NewCookieStore([]byte(yiigo.Env("session.secret").String("N0awmAuS2OziVFu^9!*0LY7MeCRgQ&z0")))

	// gob register
	gob.Register(new(models.Identity))
}

// Get get session key - value
func Get(c *gin.Context, key string, defaultValule ...interface{}) (interface{}, error) {
	session, err := store.Get(c.Request, gosessid)

	if err != nil {
		return nil, errors.Wrap(err, "get session error")
	}

	// Get some session values.
	v, ok := session.Values[key]

	if !ok {
		if len(defaultValule) > 0 {
			return defaultValule[0], nil
		}

		return nil, nil
	}

	return v, nil
}

// Set set session key - value, duration: seconds
func Set(c *gin.Context, key string, data interface{}, duration ...int) error {
	session, err := store.Get(c.Request, gosessid)

	if err != nil {
		return errors.Wrap(err, "set session error")
	}

	if len(duration) > 0 {
		session.Options = &sessions.Options{
			Path:   "/",
			MaxAge: duration[0],
		}
	}

	// Set some session values.
	session.Values[key] = data
	// Save it before we write to the response/return from the handler.
	if err = session.Save(c.Request, c.Writer); err != nil {
		return errors.Wrap(err, "get session error")
	}

	return nil
}

// Delete delete session key
func Delete(c *gin.Context, key string) error {
	session, err := store.Get(c.Request, gosessid)

	if err != nil {
		return errors.Wrap(err, "delete session error")
	}

	delete(session.Values, key)

	if err = session.Save(c.Request, c.Writer); err != nil {
		return errors.Wrap(err, "delete session error")
	}

	return nil
}

// Destroy destroy session
func Destroy(c *gin.Context) error {
	session, err := store.Get(c.Request, gosessid)

	if err != nil {
		return errors.Wrap(err, "destroy session error")
	}

	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: -1,
	}

	if err = session.Save(c.Request, c.Writer); err != nil {
		return errors.Wrap(err, "destory session error")
	}

	return nil
}
