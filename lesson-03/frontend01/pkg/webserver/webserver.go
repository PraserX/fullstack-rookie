package webserver

import (
	"fmt"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
)

// Server ...
type Server struct {
	content string
	router  *gin.Engine
}

// New ...
func New(opts ...Option) *Server {
	var options = &Options{}
	options.Host = "127.0.0.1"
	options.Port = "28080"
	options.Content = "./public"
	options.DevMode = false

	for _, opt := range opts {
		opt(options)
	}

	if !options.DevMode {
		gin.SetMode(gin.ReleaseMode)
	}

	webserver := Server{
		content: options.Content,
		router:  gin.New(),
	}

	return &webserver
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

	s.router.Run(fmt.Sprintf("%s:%s", "127.0.0.1", "28080"))
}

// DefaultSecurityConfig ...
func (s *Server) DefaultSecurityConfig() {
	cspOpts := map[string]string{
		"default-src": "'self'",
		"img-src":     "'self' data:",
	}

	s.router.Use(gin.Recovery())

	s.router.Use(helmet.DNSPrefetchControl())
	s.router.Use(helmet.FrameGuard())
	s.router.Use(helmet.XSSFilter())
	s.router.Use(helmet.NoCache())
	s.router.Use(helmet.ContentSecurityPolicy(cspOpts, true))
}

// serveStaticContent serves static pages content through the router engine.
func serveStaticContent(router *gin.Engine, webRoot string) {
	router.Static("/images", webRoot+"/images")
	router.StaticFile("/favicon.png", webRoot+"/favicon.png")
	router.StaticFile("/css/styles.css", webRoot+"/css/styles.css")
	router.StaticFile("/js/app.js", webRoot+"/js/app.js")
	router.StaticFile("/index.html", webRoot+"/index.html")
	router.StaticFile("/", webRoot+"/index.html")
}
