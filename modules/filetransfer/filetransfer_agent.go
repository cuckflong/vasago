package example

import (
	"context"
	"encoding/gob"
	"io"
	"os"

	"github.com/cuckflong/vasago/agent"
	"github.com/cuckflong/vasago/job"
	"github.com/cuckflong/vasago/modules"
)

const ModuleName = "example"

type Agent struct {
	core *agent.Core
}

type JobReq struct {
	Receiving bool
	Path      string
}

func init() {
	gob.Register(JobReq{})
}

type readerFunc func(p []byte) (n int, err error)

func (rf readerFunc) Read(p []byte) (n int, err error) { return rf(p) }

func copyWithCtx(ctx context.Context, dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, readerFunc(func(p []byte) (int, error) {
		// golang non-blocking channel: https://gobyexample.com/non-blocking-channel-operations
		select {

		// if context has been canceled
		case <-ctx.Done():
			// stop process and propagate "context canceled" error
			return 0, ctx.Err()
		default:
			// otherwise just run default io.Reader implementation
			return src.Read(p)
		}
	}))
	return err
}

func (a *Agent) OnReceiveJob(jb job.AgentJob) error {
	var jobReq JobReq
	if err := jb.Decode(&jobReq); err != nil {
		return err
	}

	if jobReq.Receiving {
		file, err := os.Open(jobReq.Path)
		if err != nil {
			return err
		}

		return copyWithCtx(jb.Context(), jb, file)
	} else {
		file, err := os.Create(jobReq.Path)
		if err != nil {
			return err
		}

		return copyWithCtx(jb.Context(), file, jb)
	}
}

func init() {
	modules.Register(ModuleName,
		func(core *modules.LeaderCore) (modules.Leader, error) {
			return &Leader{}, nil
		},
		func(core *agent.Core) (modules.Agent, error) {
			return &Agent{
				core: core,
			}, nil
		},
	)
}
