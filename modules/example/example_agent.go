package example

import (
	"encoding/gob"
	"time"

	"github.com/cuckflong/vasago/agent"
	"github.com/cuckflong/vasago/job"
	"github.com/cuckflong/vasago/modules"
)

const ModuleName = "example"

type Agent struct {
	core *agent.Core
}

type JobReq struct {
	PingMessage string
}

type JobResp struct {
	Response string
}

func init() {
	gob.Register(JobReq{})
	gob.Register(JobResp{})
}

func (a *Agent) OnReceiveJob(jb job.AgentJob) error {
	var jobReq JobReq
	if err := jb.Decode(&jobReq); err != nil {
		return err
	}

	select {
	case <-time.After(10 * time.Second):
		break
	case <-jb.Context().Done():
		return jb.Context().Err()
	}

	return jb.Encode(JobResp{Response: "You said: " + jobReq.PingMessage})
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
