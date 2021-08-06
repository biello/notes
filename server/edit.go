package server

import (
	"github.com/biello/notes/db"
	"github.com/biello/notes/web"
	"github.com/gin-gonic/gin"
)

func (s *Server) edit(ctx *gin.Context) {
	userName := ctx.Param("user")
	pageName := ctx.Param("page")
	s.db.View(func(tx *db.Tx) error {
		p, _ := tx.Page([]byte(userName), []byte(pageName))

		return web.Edit.Execute(ctx.Writer, data{
			"Title": pageName,
			"Path":  "/notes/" + userName + "/" + pageName,
			"Text":  string(p.Text),
		})
	})
}
