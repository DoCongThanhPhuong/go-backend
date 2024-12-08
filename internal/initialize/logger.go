package initialize

import (
	"github.com/DoCongThanhPhuong/go-backend/global"
	"github.com/DoCongThanhPhuong/go-backend/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}