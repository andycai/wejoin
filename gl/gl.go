package gl

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var App *fiber.App
var Log *zap.Logger

func Logger() {
	// Define log level
	level := zap.NewAtomicLevel()
	level.SetLevel(zap.DebugLevel)

	// Create new zap logger config
	encoderCfg := zap.NewProductionEncoderConfig()

	// Create new zap logger
	Log = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		level,
	))
	defer Log.Sync()
}
