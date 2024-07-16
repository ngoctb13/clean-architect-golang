package setting

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ngoctb13/clean-architect-golang/config"
	"github.com/ngoctb13/clean-architect-golang/infra"

	"go.uber.org/zap"
)

const (
	migrationFile = "file://./migrations/sql"
)

// Migration database ...
func MigrateDatabase(cfg *config.MySQLConfig) {
	infra.CreateDBAndMigrate(cfg, migrationFile)
}

// Wait OS Signal ...
func WaitOSSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	s := <-c
	zap.S().Infof("Recieve os.Signal: %s", s.String())
}
