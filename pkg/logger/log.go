package logger

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/color"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type Logger struct {
	zerolog.Logger
}

var (
	logger *Logger
	mu     sync.RWMutex
)

const (
	KeyServiceName = "service_name"
	KeyLogId       = "log_id"
	KeyFileError   = "file_error"
)

func InitLog(serviceName string) {
	mu.Lock()
	defer mu.Unlock()
	if logger != nil {
		return
	}

	out := zerolog.ConsoleWriter{
		Out:        color.Output(),
		NoColor:    false,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			switch i {
			case "error":
				return color.Red("["+strings.ToUpper(i.(string))+"] ") + "|"
			case "info":
				return color.Green("["+strings.ToUpper(i.(string))+"] ") + "|"
			case "warn":
				return color.Yellow("["+strings.ToUpper(i.(string))+"] ") + "|"
			case "debug":
				return color.Blue("["+strings.ToUpper(i.(string))+"] ") + "|"
			default:
				return strings.ToUpper(fmt.Sprintf("[%s]", i)) + "|"
			}
		},
		FormatFieldName: func(i interface{}) string {
			return color.Cyan(i) + "="
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("| %s |", i)
		},
	}

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log := zerolog.New(out).With().Timestamp().Caller().Logger()
	confLog := log.With().Str(KeyServiceName, serviceName)
	lg := confLog.Logger()
	logger = &Logger{lg}
}

func GetLogger() *Logger {
	mu.RLock()
	defer mu.RUnlock()
	// handle generate log id default
	uid := uuid.NewString()
	lg := logger.With().Str(KeyLogId, uid).Logger()
	return &Logger{lg}
}

func (lg *Logger) StackTrace() *Logger {
	stack := getFullStack()
	newLg := lg.With().Str(KeyFileError, stack).Logger()
	return &Logger{newLg}
}

func getFullStack() string {
	buf := make([]byte, 1<<16)
	stackSize := runtime.Stack(buf, true)
	stack := fmt.Sprintf("%s", buf[0:stackSize])
	stackTemp := strings.Split(stack, "\n")
	stackFile := fmt.Sprintf("file: %s, func: %s", strings.TrimSpace(stackTemp[6]), strings.TrimSpace(stackTemp[5]))
	return stackFile
}

func (lg *Logger) GetFunctionName() string {
	pc, _, _, _ := runtime.Caller(2) // Adjust the call stack index as needed
	fullFnName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(fullFnName, ".")
	fnName := parts[len(parts)-1]
	return fnName
}
