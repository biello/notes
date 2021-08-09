package server

import (
	"html/template"
	"net/http"
	"strings"
	"sync"

	"github.com/biello/notes/db"
	"github.com/biello/notes/web"
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
	cookiePair, err := ctx.Request.Cookie("SID")
	logrus.Infof("sign check middleware cookiePair: %s", cookiePair)
	if err != nil {
		logrus.Infof("cookie SID not found: %s", err.Error())
		ctx.Redirect(http.StatusFound, "/login?target="+ctx.Request.URL.Path)
		ctx.Abort()
		return
	}

	sessionID := cookiePair.Value
	user, ok := s.sessions.Load(sessionID)
	if !ok {
		logrus.Infof("SID: %d not in sessions", sessionID)
		ctx.Redirect(http.StatusFound, "/login?target="+ctx.Request.URL.Path)
		ctx.Abort()
		return
	}
	sessionUser := user.(string)
	if userName != sessionUser {
		logrus.Infof("SID: %d user not macth %s != %s", sessionID, sessionUser, user)
		ctx.String(http.StatusOK, string(web.UnauthorizedText))
		ctx.Abort()
		return
	}
	ctx.Next()
}

func (s *Server) SignCheck(ctx *gin.Context) {
	sessionID, err := ctx.Request.Cookie("SID")
	logrus.Infof("sign check SID: %s", sessionID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	_, ok := s.sessions.Load(sessionID)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.String(http.StatusOK, "success")
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
