package install

import (
	machine "./machine"
	terminal "./terminal"
)

func (self *Context) InstallAndUninstallPackagesFor(m machine.Type) error {
	terminal.Output("Updating package lists...")

	//step.AskRetry(pm.Update)
	terminal.Output("Upgrading packages...")
	//step.AskRetry(pm.Upgrade)
	terminal.Output("Installing packages...")

	//if err := pm.RemovePackages(self.Packages.Remove); err != nil {
	//	panic(fmt.Errorf("can't remove packages: %v\n", err))
	//}
	//if err := pm.InstallPackages(self.Packages.Install); err != nil {
	//	panic(fmt.Errorf("can't install packages: %v\n", err))
	//}
	return nil
}
