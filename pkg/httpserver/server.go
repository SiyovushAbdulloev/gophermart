package httpserver

import "github.com/gin-gonic/gin"

type Server struct {
	App     *gin.Engine
	Address string
}

func New(opts ...Option) *Server {
	server := &Server{
		App:     nil,
		Address: "",
	}

	for _, opt := range opts {
		opt(server)
	}

	server.App = gin.New()

	return server
}

func (s *Server) Run() error {
	return s.App.Run(s.Address)
}
