package install

type Context struct {
	GitRepository string
	



type Installation struct {
	MacheType machine.Type
	Name string
	Description string

	Steps []*InstallStep
}

type InstallStep struct {
	Name string
	Description string
	Action func (c Context, error) error

	PreviousStep *InstallStep
	NextStep *InstallStep
}

func InstallSteps() []string {
	return []string{
		"Install & Remove Packages",
		"Prepare Filesystem",
		"Install Configuration Files",
		"Configure Kernel & Other System Settings",
		"Setup Networking",
		"Setup Default Controller",
	}
}
