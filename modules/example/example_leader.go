package example

import (
	"context"
	"errors"

	"github.com/cuckflong/vasago/modules"
)

type Leader struct {
}

type CommandArgs struct {
	UserAge     int    `vasago:"Age of user"`
	PingMessage string `vasago:"Ping message"`
}

func (l *Leader) Args() interface{} {
	return &CommandArgs{}
}

func (l *Leader) Validate(args interface{}) error {
	parsedArgs := args.(*CommandArgs)
	if parsedArgs.UserAge < 18 {
		return errors.New("user must be over 18")
	}

	if parsedArgs.PingMessage == "" {
		return errors.New("ping message must have content")
	}

	return nil
}

func (l *Leader) OnCommand(args interface{}, targets []modules.System, resp modules.CommandResponse) {
	parsedArgs := args.(*CommandArgs)
	for _, target := range targets {
		job := target.CreateJob(context.Background())
		job.Encode(JobReq{
			PingMessage: parsedArgs.PingMessage,
		})
		job.Close()
	}

	resp.UseAutomaticProgress()
}
