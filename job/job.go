package job

import (
	"context"
	"io"

	"github.com/sirupsen/logrus"
)

type Status int

const (
	StatusPending Status = iota
	StatusExecuting
	StatusCancelled
	StatusFailed
	StatusStreaming
	StatusSuccess
)

type LeaderAgentJob interface {
	ID() string

	Logger() logrus.FieldLogger
	LogOutput() io.Reader

	Write(data []byte) (int, error)
	Encode(data interface{}) error
	Close() error

	Read(data []byte) (int, error)
	Decode(data interface{}) error

	Cancel(err error)

	Status() <-chan int
	Context() context.Context

	CurrentStatus() int
}

type AgentJob interface {
	ID() string

	Logger() logrus.FieldLogger
	Context() context.Context

	Read(data []byte) (int, error)
	Decode(data interface{}) error

	Write(data []byte) (int, error)
	Encode(data interface{}) error
	Close() error

	Failed(err error)
}
