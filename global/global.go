package global

import (
	"github.com/DoCongThanhPhuong/go-backend/pkg/logger"
	"github.com/DoCongThanhPhuong/go-backend/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
