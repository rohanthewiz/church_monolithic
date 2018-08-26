package app

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/rohanthewiz/church/resource/auth"
	"github.com/rohanthewiz/church/church_redis"
	"time"
	"github.com/rohanthewiz/church/web/flash"
	"github.com/rohanthewiz/serr"
	"github.com/rohanthewiz/church/web/context"
)

// Note on Redirect: the SeeOther code (303) is the preferred code when redirecting after a post
// so the browser knows to do a fresh get request

func Redirect(c echo.Context, url, fl_msg string) {
	if fl_msg != "" {
		fl := flash.NewFlash()
		fl.Info = fl_msg  // todo warn and error
		fl.Set(c)
	}
	c.Redirect(http.StatusSeeOther, url)
}

// Generate and persist form token
func GenerateFormToken() (token string, err error) {
	tokenLifetime := 3600 * time.Second
	token = auth.RandomKey()
	err = church_redis.Set(token, "true", tokenLifetime)
	if err != nil {
		return token, serr.Wrap(err)
	}
	return
}

func VerifyFormToken(token string) bool {
	str, err := church_redis.Get(token)
	if err != nil { return false }
	if str == "true" { return true }
	return false
}

func IsLoggedIn(c echo.Context) bool {
	cc, ok := c.(*context.CustomContext)
	return ok && cc.Admin
}