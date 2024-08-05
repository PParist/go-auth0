// platform/router/router.go

package router

import (
	"encoding/gob"
	"net/http"

	"github.com/PParist/go-auth0/platform/authenticator"
	"github.com/PParist/go-auth0/web/app/callback"
	"github.com/PParist/go-auth0/web/app/login"
	"github.com/PParist/go-auth0/web/app/logout"
	"github.com/PParist/go-auth0/web/app/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", user.Handler)
	router.GET("/logout", logout.Handler)

	return router
}
