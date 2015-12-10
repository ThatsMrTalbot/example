package server

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

// Server implements a server instance
type Server struct {
	Address  string `yaml:"address"`
	SSL      bool   `yaml:"ssl"`
	KeyFile  string `yaml:"keyfile"`
	CertFile string `yaml:"certfile"`
}

// Start server
func (server *Server) Start(handler http.Handler) error {
	s := &http.Server{
		Addr:    server.Address,
		Handler: handler,
	}

	if server.SSL {
		logrus.Infof("Serving HTTPS on %s", server.Address)
		return s.ListenAndServeTLS(server.CertFile, server.KeyFile)
	}

	logrus.Infof("Serving HTTP on %s", server.Address)
	return s.ListenAndServe()
}
