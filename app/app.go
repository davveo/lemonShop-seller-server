package app

import (
	"fmt"
	"github.com/davveo/lemonShop-framework/cache"
	"github.com/davveo/lemonShop-framework/db"
	"github.com/davveo/lemonShop-seller-server/app/middleware"
	"github.com/davveo/lemonShop-seller-server/app/router"
	"github.com/davveo/lemonShop-seller-server/conf"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"syscall"
)

type Server struct {
	db      db.Repo
	cache   cache.Repo
	lg      *zap.Logger
	appConf *conf.AppConf
}

func NewServer(appConf *conf.AppConf, lg *zap.Logger,
	dbRepo db.Repo, cacheRepo cache.Repo) Server {
	return Server{
		db:      dbRepo,
		lg:      lg,
		appConf: appConf,
		cache:   cacheRepo,
	}
}

func (s *Server) Init() {
	defer s.Clean() // 资源清理

	gin.SetMode(conf.Conf.Server.RunMode)

	engine := gin.New()
	middlewares := []gin.HandlerFunc{
		gin.Recovery(),
		gin.Logger(),
		middleware.Cors(),
		middleware.WrapperCtx(),
		middleware.RequestId(),
	}
	engine.Use(middlewares...)

	router.Init(engine)

	endless.DefaultMaxHeaderBytes = 1 << 20
	endless.DefaultReadTimeOut = conf.Conf.Server.ReadTimeout
	endless.DefaultWriteTimeOut = conf.Conf.Server.WriteTimeout
	endPoint := fmt.Sprintf(":%d", conf.Conf.Server.HttpPort)

	server := endless.NewServer(endPoint, engine)
	server.BeforeBegin = func(add string) {
		log.Printf("[info] pid is %d, start "+
			"http server listening %s", syscall.Getpid(), endPoint)
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
		return
	}
}

func (s *Server) Clean() {
	if s.db != nil {
		if err := s.db.DbWClose(); err != nil {
			s.lg.Info("dbw close err, ", zap.Error(err))
			return
		}
		if err := s.db.DbRClose(); err != nil {
			s.lg.Info("dbr close err, ", zap.Error(err))
			return
		}
	}

	if s.cache != nil {
		if err := s.cache.Close(); err != nil {
			s.lg.Info("cache close err, ", zap.Error(err))
			return
		}
	}

	if err := s.lg.Sync(); err != nil {
		s.lg.Info("log sync err, ", zap.Error(err))
		return
	}
}
