package gin

import (
	"github.com/A-SoulFan/acao-homework/internal/pkg/transports/http"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var ProviderSet = wire.NewSet(NewServer)

type Server struct {
	httpServer *http.Server
	logger     *zap.Logger
}

func NewServer(logger *zap.Logger, hs *http.Server) *Server {
	return &Server{logger: logger, httpServer: hs}
}

func (server *Server) Start() error {
	if server.httpServer != nil {
		if err := server.httpServer.Start(); err != nil {
			return errors.Wrap(err, "http server start error")
		}
	}

	return nil
}

func (server *Server) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		server.logger.Info("receive server signal", zap.String("signal", s.String()))
		if server.httpServer != nil {
			if err := server.httpServer.Stop(); err != nil {
				server.logger.Warn("stop http server error", zap.Error(err))
			}
		}
		os.Exit(0)
	}
}
