package migrate

import (
	"fmt"
	"log/slog"

	"github.com/pressly/goose"
)

type GooseSlogger struct {
	slogr *slog.Logger
}

var _ goose.Logger = &GooseSlogger{}

func NewGooseSlogger(s *slog.Logger) *GooseSlogger {
	return &GooseSlogger{
		slogr: s,
	}
}

// Fatal implements goose.Logger.
func (g *GooseSlogger) Fatal(v ...interface{}) {
	g.slogr.Error(fmt.Sprintf("%s", v))
	panic(v)
}

// Fatalf implements goose.Logger.
func (g *GooseSlogger) Fatalf(format string, v ...interface{}) {
	g.slogr.Error(fmt.Sprintf(format, v...))
	panic(v)
}

// Print implements goose.Logger.
func (g *GooseSlogger) Print(v ...interface{}) {
	g.slogr.Info(fmt.Sprintf("%s", v))
}

// Printf implements goose.Logger.
func (g *GooseSlogger) Printf(format string, v ...interface{}) {
	g.slogr.Info(fmt.Sprintf(format, v...))
}

// Println implements goose.Logger.
func (g *GooseSlogger) Println(v ...interface{}) {
	g.slogr.Info(fmt.Sprintf("%v\n", v))
}
