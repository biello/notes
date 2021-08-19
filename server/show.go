package server

import (
	"github.com/biello/notes/db"
	"github.com/biello/notes/web"
	"github.com/gin-gonic/gin"
)

func (s *Server) show(ctx *gin.Context) {
	userName := ctx.Param("user")
	pageName := ctx.Param("page")
	s.db.View(func(tx *db.Tx) error {
		p, err := tx.Page([]byte(userName), []byte(pageName))

		if err != nil || len(p.Text) == 0 {
			p.Text = web.EmptyPageText
		}

		return web.Show.Execute(ctx.Writer, data{
			"Title": pageName,
			"Path":  "/notes/" + userName + "/" + pageName + "/edit",
			"Text":  bytesAsHTML(parsedMarkdown(p.Text)),
		})
	})
}

func (s *Server) notes(ctx *gin.Context) {
	userName := ctx.Param("user")
	s.db.View(func(tx *db.Tx) error {
		notes, err := tx.Notes([]byte(userName))

		if err != nil || len(notes.Notes) == 0 {

		}

		return web.Notes.Execute(ctx.Writer, data{
			"User":  userName,
			"Notes": notes,
		})

	})
}

func (s *Server) home(ctx *gin.Context) {
	userName := "home"
	pageName := "home"
	s.db.View(func(tx *db.Tx) error {
		p, err := tx.Page([]byte(userName), []byte(pageName))

		if err != nil || len(p.Text) == 0 {
			p.Text = web.EmptyPageText
		}

		return web.Show.Execute(ctx.Writer, data{
			"Title": "home",
			"Path":  "/notes/" + userName + "/" + pageName + "/edit",
			"Text":  bytesAsHTML(parsedMarkdown(p.Text)),
		})
	})
}
