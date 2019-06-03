package example

import (
	"encoding/gob"
	"time"

	"github.com/cuckflong/vasago/job"
)

type Agent struct {
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
