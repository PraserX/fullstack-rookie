package webserver

import (
	"fmt"

	"github.com/PraserX/fullstack-rookie/pkg/database"
	v1Comm "github.com/PraserX/fullstack-rookie/pkg/webserver/api/v1/comments"
	v1User "github.com/PraserX/fullstack-rookie/pkg/webserver/api/v1/users"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
)

// Server ...
type Server struct {
	host     string
	port     string
	content  string
	router   *gin.Engine
	database *database.Database
}

// New ...
func New(opts ...Option) (*Server, error) {
	var err error

	var options = &Options{}
	options.Host = "127.0.0.1"
	options.Port = "28080"
	options.Content = "./public"
	options.DevMode = false
	options.Database = nil

	for _, opt := range opts {
		opt(options)
	}

	if !options.DevMode {
		gin.SetMode(gin.ReleaseMode)
	}

	if options.Database == nil {
		if options.Database, err = database.New(); err != nil {
			return nil, fmt.Errorf("cannot create database")
		}
	}

	webserver := Server{
		host:     options.Host,
		port:     options.Port,
		content:  options.Content,
		router:   gin.New(),
		database: options.Database,
	}

	return &webserver, nil
}

// Serve runs http Golang Gin web server.
func (s *Server) Serve() {
	// Setup basic security configuration
	s.DefaultSecurityConfig()

	// Serve static files
	serveStaticContent(s.router, s.content)

	s.router.NoRoute(func(c *gin.Context) {
		c.File(s.content + "/404.html")
	})

	// {base}/api/v1
	v1Router := s.router.Group("/api/v1")
	{
		v1Comm.LocationComments(v1Router, s.database)
		v1User.LocationUsers(v1Router, s.database)
	}

	s.router.Run(fmt.Sprintf("%s:%s", s.host, s.port))
}

// DefaultSecurityConfig ...
func (s *Server) DefaultSecurityConfig() {
	// cspOpts := map[string]string{
	// 	"default-src": "'self' cdn.jsdelivr.net",
	// 	"img-src":     "'self' data:",
	// }

	s.router.Use(gin.Recovery())

	s.router.Use(helmet.DNSPrefetchControl())
	s.router.Use(helmet.FrameGuard())
	s.router.Use(helmet.XSSFilter())
	s.router.Use(helmet.NoCache())
	// s.router.Use(helmet.ContentSecurityPolicy(cspOpts, true))
}

// serveStaticContent serves static pages content through the router engine.
func serveStaticContent(router *gin.Engine, webRoot string) {
	router.Static("/images", webRoot+"/images")
	router.StaticFile("/favicon.png", webRoot+"/favicon.png")
	router.StaticFile("/css/styles.css", webRoot+"/css/styles.css") // ./public/css/styles.css
	router.StaticFile("/js/app.js", webRoot+"/js/app.js")
	router.StaticFile("/index.html", webRoot+"/index.html")
	router.StaticFile("/", webRoot+"/index.html")
}
