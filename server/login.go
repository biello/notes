package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/biello/notes/db"
	"github.com/biello/notes/web"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var loginFirst = `# Please login first
<a href="/login">click to login</a>`

func (s *Server) loginPage(ctx *gin.Context) {
	target := ctx.Query("target")
	logrus.Infof("loginPage target: %s", target)

	web.Login.Execute(ctx.Writer, data{"Target": target})
}

func (s *Server) login(ctx *gin.Context) {
	target := ctx.PostForm("target")
	sessionID, err := ctx.Request.Cookie("sessionID")
	logrus.Infof("target: %s", target)
	if err != nil {
		userName := ctx.PostForm("user")
		password := ctx.PostForm("password")
		if len(userName) == 0 || len(password) == 0 {
			ctx.HTML(http.StatusOK, "invalid username or password, %s", loginFirst)
			return
		}

		logrus.Infof("user: %s, password: %s", userName, password)
		err := s.db.View(func(tx *db.Tx) error {
			u, err := tx.User([]byte(userName))
			if err != nil || len(u.Password) == 0 {
				return db.ErrUserNotFound
			}

			if password == string(u.Password) {
				session := strconv.FormatInt(time.Now().Unix(), 10)
				s.sessions.Store(session, userName)
			} else {
				ctx.String(http.StatusUnauthorized, "wrong password")
				return db.ErrWrongPassword
			}

			return nil
		})
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		http.Redirect(ctx.Writer, ctx.Request, target, 302)
	} else {
		if _, ok := s.sessions.Load(sessionID); ok { // already logined
			http.Redirect(ctx.Writer, ctx.Request, target, 302)
		} else { // sessionID not exist
			http.Redirect(ctx.Writer, ctx.Request, "/login?target="+target, 302)
		}
	}

}

func (s *Server) logout(ctx *gin.Context) {

}
