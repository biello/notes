package server

import (
	"github.com/biello/notes/db"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var server *Server

func Init(logger *logrus.Logger, db *db.DB) {
	server = &Server{logger: logger, db: db}
}

func HTTP(r *gin.Engine) {

	r.Static("/static/", "./web/assets")

	rootGroup := r.Group("/")
	{
		rootGroup.GET("", server.home)
		rootGroup.GET("favicon.ico", server.favicon)
		rootGroup.GET("login", server.loginPage)
		rootGroup.POST("login", server.login)
		rootGroup.POST("signCheck", server.SignCheck)
		rootGroup.GET("password", server.passwordPage)
		rootGroup.POST("password", server.password)
	}

	homeGroup := r.Group("/home")
	{
		homeGroup.GET("", server.show)
		// todo
	}

	adminGroup := r.Group("/admin", server.AdminCheckMiddleware)
	{
		adminGroup.GET("/register", server.registerPage)
		adminGroup.POST("/register", server.register)
	}

	userGroup := r.Group("/notes", server.SignCheckMiddleware)
	{
		userGroup.GET("", server.home)
		userGroup.GET(":user", server.notes)
		userGroup.GET(":user/:page", server.show)
		userGroup.GET(":user/:page/edit", server.edit)

		userGroup.POST(":user/:page", server.save)
	}

}
