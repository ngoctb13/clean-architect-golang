package setting

import (
	"clean-arch-repo/config"
	"clean-arch-repo/infra"
	"os"
	"os/signal"
	"syscall"

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
