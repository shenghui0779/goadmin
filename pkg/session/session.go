// Package session ...
package session

import (
	"encoding/gob"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
)

const gosessid = "GOSESSID"

var store sessions.Store

// Start start session
func Start() {
	store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))

	// gob register
	gob.Register(new(Identity))
}

// Get get session key - value
func Get(c *gin.Context, key string) (interface{}, bool) {
	session, _ := store.Get(c.Request, gosessid)

	v, ok := session.Values[key]

	return v, ok
}

// Set set session key - value, duration: seconds
func Set(c *gin.Context, key string, data interface{}, duration time.Duration) error {
	session, _ := store.Get(c.Request, gosessid)

	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: int(duration.Seconds()),
	}

	// Set some session values.
	session.Values[key] = data

	// Save it before we write to the response/return from the handler.
	if err := session.Save(c.Request, c.Writer); err != nil {
		return errors.Wrap(err, "err session save")
	}

	return nil
}

// Delete delete session key
func Delete(c *gin.Context, key string) error {
	session, _ := store.Get(c.Request, gosessid)

	delete(session.Values, key)

	if err := session.Save(c.Request, c.Writer); err != nil {
		return errors.Wrap(err, "err session delete")
	}

	return nil
}

// Destroy destroy session
func Destroy(c *gin.Context) error {
	session, _ := store.Get(c.Request, gosessid)

	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: -1,
	}

	if err := session.Save(c.Request, c.Writer); err != nil {
		return errors.Wrap(err, "err session destroy")
	}

	return nil
}
