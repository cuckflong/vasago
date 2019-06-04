package example

import (
	"context"
	"errors"

	"github.com/cuckflong/vasago/modules"
	"github.com/cuckflong/vasago/modules/filetransfer"
)

type Leader struct {
	Args *CommandArgs
	core modules.LeaderCore
}

type CommandArgs struct {
	UserAge     int    `vasago:"Age of user"`
	PingMessage string `vasago:"Ping message"`
}

func (l *Leader) ConfigDesc(typ modules.ProviderType) interface{} {
	return &CommandArgs{}
}

func (l *Leader) SetConfig(args interface{}) error {
	parsedArgs := args.(*CommandArgs)
	if parsedArgs.UserAge < 18 {
		return errors.New("user must be over 18")
	}

	if parsedArgs.PingMessage == "" {
		return errors.New("ping message must have content")
	}

	l.Args = parsedArgs

	return nil
}

func (l *Leader) Command(targets []modules.System, resp modules.CommandResponse) error {
	dataProvider := l.core.GetDefaultProvider(modules.TypeDataProvider).(modules.DataProvider)
	dataProvider.Lock("test")

	providers := l.core.GetAllProviders()

	var fileTransferProvider *filetransfer.Leader
	ok := l.core.GetProviderConcrete(&fileTransferProvider)
	if !ok {
		panic("rip")
	}

	parsedArgs := l.Args
	for _, target := range targets {
		job := target.CreateJob(context.Background())
		job.Encode(JobReq{
			PingMessage: parsedArgs.PingMessage,
		})
		job.Close()
	}

	resp.UseAutomaticProgress()

	return nil
}

var _ modules.CommandProvider = (*Leader)(nil)
