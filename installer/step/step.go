package step

import (
	"fmt"

	survey "github.com/AlecAivazis/survey"
)

type AskOption int

const (
	RetryOption AskOption = iota
	SkipOption
	ExitOption
)


type InstallStep int

const (
	PrepareSystem InstallStep = iota
	DownloadOSInstallMedia
	InstallConfigFiles
	SetupNetworking
	BuildRouterVMs
)

type step func() error

func AskRetry(s step) error {
	if err := s(); err != nil {
		var q = &survey.Select{
			Message: fmt.Sprintf("Step failed due to error: %v\nRetry?", err),
			Options: []string{"retry", "skip", "exit"},
			Default: "retry",
		}
		var response string
		survey.AskOne(q, &resp)
		switch response {
		case "skip":
			return nil
		case "exit":
			fmt.Println("[ERROR] Unable to proceed:", err)
			os.Exit(1)
		default: // retry
			return AskRetry(s)
		}
	}
	return nil
}
