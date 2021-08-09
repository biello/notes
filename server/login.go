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
	logrus.Infof("login target: %s", target)

	userName := ctx.PostForm("user")
	password := ctx.PostForm("password")
	if len(userName) == 0 || len(password) == 0 {
		ctx.String(http.StatusOK, "invalid username or password")
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
			ctx.JSON(http.StatusOK, gin.H{"ErrNo": "0", "SessionId": session})
			return nil
		} else {
			return db.ErrWrongPassword
		}

	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ErrNo": "1", "ErrMsg": err.Error()})
		return
	}
}

func (s *Server) logout(ctx *gin.Context) {

}
