package install

import (
	"github.com/AlecAivazis/survey"
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
		var resp string
		survey.AskOne(q, &resp)
		if resp == "retry" {
			return AskRetry(s)
		} else if resp == "skip" {
			return nil
		} else if resp == "exit" {
			log.Println(Fail(err.Error()))
			os.Exit(1)
			return err
		}
	}
	return nil
}
