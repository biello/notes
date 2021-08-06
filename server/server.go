package server

import (
	"html/template"
	"net/http"
	"strings"
	"sync"

	"github.com/biello/notes/db"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
)

// Server is the notes server
type Server struct {
	logger   *logrus.Logger
	db       *db.DB
	sessions sync.Map // sessionID:userName
}

// New creates a new notes server
func New(logger *logrus.Logger, db *db.DB) *Server {
	return &Server{logger: logger, db: db, sessions: sync.Map{}}
}

func (s *Server) SignCheckMiddleware(ctx *gin.Context) {
	userName := ctx.Param("user")
	err := s.db.View(func(tx *db.Tx) error {
		u, err := tx.User([]byte(userName))
		if err != nil || len(u.Password) == 0 {
			return db.ErrUserNotFound
		}

		return nil
	})
	if err != nil {
		ctx.Redirect(http.StatusFound, "/login?target="+ctx.Request.URL.Path)
		ctx.Abort()
		return
	}
	ctx.Next()
}

func (s *Server) redirect(ctx *gin.Context) {
	if path := ctx.Request.URL.Path; len(path) > 1 {
		target := strings.TrimSuffix(path, "/")

		if target == "/home" {
			target = "/"
		}

		http.Redirect(ctx.Writer, ctx.Request, target, 302)
	}
}

type data map[string]interface{}

// bytesAsHTML returns the template bytes as HTML
func bytesAsHTML(b []byte) template.HTML {
	return template.HTML(string(b))
}

// parsedMarkdown returns provided bytes parsed as Markdown
func parsedMarkdown(b []byte) []byte {
	return blackfriday.MarkdownCommon(b)
}
