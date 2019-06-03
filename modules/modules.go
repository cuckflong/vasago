package modules

import (
	"context"
	"io"

	"github.com/cuckflong/vasago/agent"
	"github.com/cuckflong/vasago/job"
)

// events we want:
// on join
// on disconnect
// on reconnect
// on leave

type LeaderCore struct {
}

type System interface {
	CreateJob(ctx context.Context) job.LeaderAgentJob
	ID() string
}

type Leader interface {
}

type Commandable interface {
	OnCommand(args interface{}, targets []System, resp CommandResponse) error
	Args() interface{} // TODO: support concrete type and structure (maybe only concrete type and use a
	// helper function to generate from structure)
	Validate(args interface{}) error
}

type CommandResponse interface {
	Logger() io.Writer
	UseAutomaticProgress()
}

type Agent interface {
	OnReceiveJob(job job.AgentJob) error
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
type AgentInit func(core *agent.Core) (Agent, error)

func Register(moduleName string, leader LeaderInit, agent AgentInit) {

}
