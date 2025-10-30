package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)

	return &Logger{
		Debug:   log.New(writer, prefix+"[DEBUG] ", log.Ldate|log.Ltime),
		Info:    log.New(writer, prefix+"[INFO] ", log.Ldate|log.Ltime),
		Warning: log.New(writer, prefix+"[WARN] ", log.Ldate|log.Ltime),
		Error:   log.New(writer, prefix+"[ERROR] ", log.Ldate|log.Ltime),
		writer:  writer,
	}
}

// Métodos personalizados para facilitar o uso

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Info.Printf(format, v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.Warning.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Error.Printf(format, v...)
}
