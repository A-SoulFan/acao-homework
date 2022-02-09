package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(NewServer, NewRouter, NewOptions)

type Options struct {
	Host               string
	Port               int
	Mode               string
	MaxMultipartMemory uint8
}

type Server struct {
	o          *Options
	host       string
	port       int
	logger     *zap.Logger
	router     *gin.Engine
	httpServer http.Server
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := &Options{}
	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal http option error")
	}

	logger.Info("load http server options success")

	return o, err
}

type InitRouters func(r *gin.Engine)

func NewRouter(o *Options, logger *zap.Logger, init InitRouters) *gin.Engine {
	gin.SetMode(o.Mode)
	// 初始化 gin
	r := gin.New()

	// panic之后自动恢复
	r.Use(gin.Recovery())
	r.Use(ginZap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginZap.RecoveryWithZap(logger, true))

	if o.MaxMultipartMemory != 0 {
		// 最大上传文件大小 mb
		r.MaxMultipartMemory = int64(o.MaxMultipartMemory) << 20
	}

	init(r)

	return r
}

func NewServer(o *Options, logger *zap.Logger, router *gin.Engine) (*Server, error) {
	var s = &Server{
		logger: logger.With(zap.String("type", "http.Server")),
		router: router,
		o:      o,
	}

	return s, nil
}

func (s *Server) Start() error {
	s.port = s.o.Port
	if s.port == 0 {
		s.port = 8080
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	s.httpServer = http.Server{Addr: addr, Handler: s.router}

	s.logger.Info("http server starting ...", zap.String("addr", addr))
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("start http server err", zap.Error(err))
			return
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	s.logger.Info("stopping http server")
	// 平滑关闭,等待5秒钟处理
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}

	return nil
}
