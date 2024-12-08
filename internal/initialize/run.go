package initialize

func Run() {
	LoadConfig()
	InitLogger()
	// global.Logger.Info("CONFIG_LOG_SUCCESS", zap.String("OK", "SUCCESS"))
	InitMySQL()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}