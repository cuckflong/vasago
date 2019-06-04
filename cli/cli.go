package cli

import (
	"fmt"
	"io"
	"strings"

	"github.com/chzyer/readline"
)

var completer *readline.PrefixCompleter
var shellStage = "main"

func Shell() {
	prompt, err := readline.NewEx(&readline.Config{
		Prompt:              "\u001b[34m»>\u001b[0m ",
		HistoryFile:         "/tmp/readline.tmp",
		AutoComplete:        completer,
		InterruptPrompt:     "^C",
		EOFPrompt:           "exit",
		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		panic(err)
	}
	defer prompt.Close()

	for {
		line, err := prompt.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		cmd := strings.Fields(line)

		if len(cmd) > 0 {
			switch shellStage {
			case "main":
				switch cmd[0] {
				case "test":
					fmt.Println("Testing prompt")
				}
			}
		}
	}
}

func getCompleter(stage string) *readline.PrefixCompleter {
	return nil
}

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}
