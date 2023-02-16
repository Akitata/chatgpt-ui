package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
)

var Sessions = session.New()

func SessionHandler(c *fiber.Ctx) error {
	sess, err := Sessions.Get(c)
	if err != nil {
		panic("get session error: " + err.Error())
	}
	uid := sess.Get("uid")
	if uid != nil {
		return c.Next()
	} else {
		sess.Set("uid", uuid.NewString()[0:8])
		if err := sess.Save(); err != nil {
			panic(err)
		}
		return c.Next()
	}
}
