package main

import (
	"flag"

	"github.com/ngoctb13/clean-architect-golang/config"
	"github.com/ngoctb13/clean-architect-golang/server"
	"github.com/ngoctb13/clean-architect-golang/setting"

	"go.uber.org/zap"
)

func main() {
	var configFile, port string
	flag.StringVar(&configFile, "config-file", "", "Specify config file path")
	flag.StringVar(&port, "port", "", "Specify port")
	flag.Parse()

	defer setting.WaitOSSignal()

	//load config
	cfg, err := config.Load(configFile)
	if err != nil {
		zap.S().Errorf("load config fail with err: %v", err)
		panic(err)
	}

	// migrate database
	go setting.MigrateDatabase(cfg.DB)

	//start new server
	s := server.NewServer(cfg)
	s.Init()

	if err := s.ListenHTTP(); err != nil {
		zap.S().Errorf("start server fail with err: %v", err)
		panic(err)
	}
}
