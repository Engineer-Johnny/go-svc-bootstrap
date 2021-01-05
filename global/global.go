package global

import (
	"go-svc-bootstrap/config_model"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

var (
	G_CONFIG config_model.Server
	G_VP     *viper.Viper
	G_LOG    *zap.Logger
	G_DB     *gorm.DB
)

const (
	ConfigFile = "config.yaml"
)
