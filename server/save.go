package server

import (
	"strings"

	"github.com/biello/notes/db"
	"github.com/gin-gonic/gin"
)

func (s *Server) save(ctx *gin.Context) {
	userName := ctx.Param("user")
	pageName := ctx.Param("page")
	s.logger.Infof("save page: %s, %s", userName, pageName)
	s.db.Update(func(tx *db.Tx) error {
		ctx.Request.ParseForm()

		p := db.Page{
			Tx:   tx,
			User: []byte(userName),
			Name: []byte(pageName),
		}

		p.Text = []byte(strings.TrimSpace(ctx.Request.FormValue("text")))

		return p.Save()
	})

	s.redirect(ctx)
}
