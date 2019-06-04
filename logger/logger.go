package logger

import "io"

type Logger struct {
	Writer io.Writer
}
