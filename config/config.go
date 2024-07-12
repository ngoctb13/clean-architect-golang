package config

import (
	"io/ioutil"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type MySQLConfig struct {
	DriverName        string `yaml:"driver_name"`
	DataSource        string `yaml:"data_source"`
	MaxOpenConns      int    `yaml:"max_open_conns"`
	MaxIdleConns      int    `yaml:"max_idle_conns"`
	ConnMaxLifeTimeMs int    `yaml:"conn_max_life_time_ms"`
	MigrationConnUrl  string `yaml:"migration_conn_url"`
	IsDevMode         bool   `yaml:"is_dev_mode"`
}

type AppConfig struct {
	DB *MySQLConfig `yaml:"db"`
}

func Load(filePath string) (*AppConfig, error) {
	if len(filePath) == 0 {
		filePath = os.Getenv("CONFIG_FILE")
	}

	// Tạo một danh sách các trường để logging
	fields := []interface{}{
		"func",
		"config.readFromFile",
		"filePath",
		filePath,
	}

	// Tạo logger với các trường đã định nghĩa
	sugar := zap.S().With(fields...)
	sugar.Debug("Load config...")
	// Log đường dẫn tệp cấu hình
	zap.S().Debugf("CONFIG_FILE: %v", filePath)

	// Đọc têp cấu hình
	configBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		sugar.Error("Failed to load config file")
		return nil, err
	}

	// Thay thế các biến môi trường trong nội dung tệp cấu hình
	configBytes = []byte(os.ExpandEnv(string(configBytes)))
	cfg := &AppConfig{}

	// Giải mã nội dung YAML thành cấu trúc AppConfig
	err = yaml.Unmarshal(configBytes, cfg)
	if err != nil {
		sugar.Error("Failed to parse config file")
		return nil, err
	}

	// Debug log cấu hình đã nạp
	zap.S().Debugf("config: %+v", cfg)
	zap.S().Debug("======================================")
	zap.S().Debugf("database config: %+v", cfg.DB)

	return cfg, nil
}
