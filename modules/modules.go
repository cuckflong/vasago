package modules

import (
	"context"
	"reflect"

	"github.com/cuckflong/vasago/agent"
	"github.com/cuckflong/vasago/job"
	"github.com/sirupsen/logrus"
)

// events we want:
// on join
// on disconnect
// on reconnect
// on leave

type LeaderCore interface {
	GetDefaultProvider(typ ProviderType) interface{}
	GetProviderConcrete(provider interface{}) bool
}

type LeaderImpl struct {
}

func (l *LeaderImpl) GetProviderConcrete(provider interface{}) {
	want := reflect.TypeOf(provider)
	for _, loadedModule := range registeredLeaders {
		if reflect.TypeOf(loadedModule) == want {
			reflect.ValueOf(provider).Elem().Set(reflect.ValueOf(loadedModule))
			return
		}
	}

	return
}

func (l *LeaderImpl) GetDefaultProvider(want ProviderType) interface{} {
	for _, loadedModule := range registeredLeaders {
		if GetProviderType(&loadedModule) == want {
			return loadedModule
		}
	}

	return nil
}

type System interface {
	CreateJob(ctx context.Context) job.LeaderAgentJob
	ID() string
}

type Leader interface {
}

type Configuration interface {
	ConfigDesc(typ ProviderType) interface{} // TODO: support concrete type and structure (maybe only concrete type and use a
	// helper function to generate from structure)
	SetConfig(typ ProviderType, config interface{}) error
}

type CommandResponse interface {
	Logger() logrus.FieldLogger
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

type LeaderInit func(core LeaderCore) (Leader, error)
type AgentInit func(core agent.Core) (Agent, error)

var IsLeader = true

var registeredLeaders = make(map[string]Leader)
var registeredAgents = make(map[string]Agent)

func Register(moduleName string, ldr LeaderInit, agt AgentInit) {
	var leaderCore LeaderCore
	// var agentCore agent.Core

	if IsLeader {
		var err error
		registeredLeaders[moduleName], err = ldr(leaderCore)
		if err != nil {
			panic(err)
		}
	}
	// TODO: etc...
}
