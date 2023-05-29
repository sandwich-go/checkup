package checkup

import (
	"log"
)

func LogError(err error) {
	GetOptions().GetLogErrorFunc()(err)
}

type Logger struct {
	*log.Logger
}

func NewLogger(l *log.Logger) *Logger {
	if l == nil {
		l = log.Default()
	}
	return &Logger{l}
}

// LogErrorAndEatError 输出 err
func (l *Logger) LogErrorAndEatError(err error) {
	if err != nil {
		l.Logger.Println(
			"[Error] ", err.Error(),
		)
	}
}
