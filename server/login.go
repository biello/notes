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
	userName := ctx.PostForm("user")
	password := ctx.PostForm("password")

	logrus.Infof("login user: %s, target: %s", userName, target)
	if len(userName) == 0 || len(password) == 0 {
		ctx.String(http.StatusOK, "invalid username or password")
		return
	}

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

func (s *Server) passwordPage(ctx *gin.Context) {
	web.Password.Execute(ctx.Writer, data{})
}

type ChangePassword struct {
	User        string
	Password    string
	NewPassword string
}

func (s *Server) password(ctx *gin.Context) {
	var up ChangePassword
	if err := ctx.ShouldBind(&up); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ErrNo": "1", "ErrMsg": err.Error()})
		return
	}

	// check
	err := s.db.View(func(tx *db.Tx) error {
		u, err := tx.User([]byte(up.User))
		if err != nil || len(u.Password) == 0 {
			return db.ErrUserNotFound
		}

		if up.Password == string(u.Password) {
			return nil
		} else {
			return db.ErrWrongPassword
		}
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ErrNo": "2", "ErrMsg": err.Error()})
		return
	}

	// update password
	err = s.db.Update(func(tx *db.Tx) error {
		u, err := tx.User([]byte(up.User))
		if err != nil || len(u.Password) == 0 {
			return db.ErrUserNotFound
		}

		u.Password = []byte(up.NewPassword)
		return u.Save()
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ErrNo": "3", "ErrMsg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ErrNo": "0", "ErrMsg": "success"})
}

func (s *Server) registerPage(ctx *gin.Context) {
	web.Register.Execute(ctx.Writer, data{})
}

type UserPassword struct {
	User     string
	Password string
}

func (s *Server) register(ctx *gin.Context) {
	var up UserPassword
	if err := ctx.ShouldBind(&up); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ErrNo": "1", "ErrMsg": err.Error()})
		return
	}

	// check
	err := s.db.View(func(tx *db.Tx) error {
		u, err := tx.User([]byte(up.User))
		if err != nil || len(u.Password) == 0 {
			return nil
		} else {
			return db.ErrUserNameExists
		}

	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ErrNo": "2", "ErrMsg": err.Error()})
		return
	}

	// add user password
	err = s.db.Update(func(tx *db.Tx) error {
		u := db.User{
			Tx:       tx,
			Name:     []byte(up.User),
			Password: []byte(up.Password),
		}
		return u.Save()
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ErrNo": "3", "ErrMsg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ErrNo": "0", "ErrMsg": "success"})
}
