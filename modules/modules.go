package modules

import (
	"context"

	"github.com/cuckflong/vasago/job"
)

// events we want:
// on join
// on disconnect
// on reconnect
// on leave

type LeaderCore struct {
}

type AgentCore struct {
}

type System interface {
	CreateJob(ctx context.Context) *job.LeaderAgentJob
}

type Leader interface {
	OnCommand(args interface{}, targets []*System)
}

type Agent interface {
	OnReceiveJob(job *job.AgentJob)
}

type Selectable interface {
	Selection(filter Filter)
}

type ConnectionEvents interface {
	OnJoin(system *System)
	OnDisconnect(system *System)
	OnReconnect(system *System)
	OnLeave(system *System)
}

type LeaderInit func(core *LeaderCore) (Leader, error)
type AgentInit func(core *AgentCore) (Agent, error)

func Register(moduleName string, leader LeaderInit, agent AgentInit) {

}
