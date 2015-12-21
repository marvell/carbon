package loggers

import (
	"fmt"
	"log"
	"os"
)

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{log.New(os.Stdout, "", log.LstdFlags)}
}

type ConsoleLogger struct {
	l *log.Logger
}

func (l *ConsoleLogger) SetPrefix(s string) {
	l.l.SetPrefix(s)
}

func (l *ConsoleLogger) Debug(msg interface{})   { l.l.Printf("[DBG] %s", msg) }
func (l *ConsoleLogger) Info(msg interface{})    { l.l.Printf("[INF] %s", msg) }
func (l *ConsoleLogger) Warning(msg interface{}) { l.l.Printf("[WRN] %s", msg) }
func (l *ConsoleLogger) Error(msg interface{})   { l.l.Printf("[ERR] %s", msg) }
func (l *ConsoleLogger) Fatal(msg interface{})   { l.l.Printf("[FAT] %s", msg); os.Exit(1) }

func (l *ConsoleLogger) Debugf(f string, msg interface{})   { l.Debug(fmt.Sprintf(f, msg)) }
func (l *ConsoleLogger) Infof(f string, msg interface{})    { l.Info(fmt.Sprintf(f, msg)) }
func (l *ConsoleLogger) Warningf(f string, msg interface{}) { l.Warning(fmt.Sprintf(f, msg)) }
func (l *ConsoleLogger) Errorf(f string, msg interface{})   { l.Error(fmt.Sprintf(f, msg)) }
func (l *ConsoleLogger) Fatalf(f string, msg interface{})   { l.Fatal(fmt.Sprintf(f, msg)) }
