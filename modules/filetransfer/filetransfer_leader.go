package example

import (
	"context"
	"io"
	"os"

	"github.com/cuckflong/vasago/modules"
)

type Leader struct {
}

type CommandArgs struct {
	RemotePath string
	LocalPath  string
	Receiving  bool
}

func (l *Leader) Args() interface{} {
	return &CommandArgs{}
}

func (l *Leader) OnCommand(args interface{}, targets []modules.System, resp modules.CommandResponse) {
	parsedArgs := args.(*CommandArgs)
	for _, target := range targets {
		go func(target modules.System) {
			job := target.CreateJob(context.Background())
			if parsedArgs.Receiving {
				localPath := parsedArgs.LocalPath + "." + target.ID()
				job.Encode(JobReq{
					Receiving: true,
					Path:      parsedArgs.RemotePath,
				})
				job.Close()

				file, err := os.Create(localPath)
				if err != nil {
					job.Cancel(err)
					return
					// TODO: do something
				}

				if _, err := io.Copy(file, job); err != nil {
					job.Cancel(err)
				}
			} else {
				job.Encode(JobReq{
					Receiving: false,
					Path:      parsedArgs.RemotePath,
				})

				file, err := os.Open(parsedArgs.LocalPath)
				if err != nil {
					// TODO: do something
				}

				if _, err := io.Copy(job, file); err != nil {
					job.Cancel(err)
				}
			}
		}(target)
	}

	resp.UseAutomaticProgress()
}
