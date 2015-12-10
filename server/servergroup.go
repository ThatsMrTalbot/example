package server

import (
	"net/http"
	"sync"

	"github.com/Sirupsen/logrus"
)

// Group defines a list of servers
type Group []*Server

// Start all servers in group
func (serverGroup Group) Start(handler http.Handler) {
	waitGroup := sync.WaitGroup{}

	logrus.Info("Starting server group")

	waitGroup.Add(len(serverGroup))
	for _, server := range serverGroup {
		go func(server *Server) {
			if err := server.Start(handler); err != nil {
				logrus.WithError(err).Error("Error starting server")
			}
			waitGroup.Done()
		}(server)
	}

	waitGroup.Wait()

	logrus.Info("All servers closed")
}
