package flags

import (
	"github.com/Lee-xy-z/recommend/pkg/healthcheck"
	"os"
	"os/signal"
	"syscall"
)

// Service represents an abstract Recommend backend compoent with some basic shared functionality.
type Service struct {
	signalsChannel  chan os.Signal
	hcStatusChannel chan healthcheck.Status
}

// NewService creates a new Service
func NewService() *Service {

	signalsChannel := make(chan os.Signal, 1)
	hcStatusChannel := make(chan healthcheck.Status)

	signal.Notify(signalsChannel, os.Interrupt, syscall.SIGTERM)
	return &Service{
		signalsChannel:  signalsChannel,
		hcStatusChannel: hcStatusChannel,
	}
}
