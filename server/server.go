package server

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/ngoctb13/clean-architect-golang/config"
	"github.com/ngoctb13/clean-architect-golang/infra"
	"github.com/ngoctb13/clean-architect-golang/infra/repos"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	httpServer *http.Server
	router     *gin.Engine
	cfg        *config.AppConfig
}

func NewServer(cfg *config.AppConfig) *Server {
	router := gin.New()
	return &Server{
		router: router,
		cfg:    cfg,
	}
}

func (s *Server) Init() {
	db, err := infra.InitMySQL(s.cfg.DB)
	if err != nil {
		zap.S().Errorf("Init db error: %v", err)
		panic(err)
	}
	repo := repos.NewSQLRepo(db, s.cfg.DB)
	domains := s.initDomains(repo)
	s.initCORS()
	s.initRouter(domains, repo)
}

func (s *Server) ListenHTTP() error {
	listen, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		zap.S().Errorf("err %v", err)
		panic(err)
	}
	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Println(address)
	s.httpServer = &http.Server{
		Addr:    address,
		Handler: s.router,
	}

	zap.S().Infof("starting http server at port %v ...", os.Getenv("PORT"))

	return s.httpServer.Serve(listen)
}
