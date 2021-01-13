package packagemanager

import (
	"fmt"

	machine "../"
	terminal "../../terminal"
)

///////////////////////////////////////////////////////////////////////////////
// TODO: Need sources management. Ability to detect and manage the soruces file
//       and probably track at least for debian BULLSEYE vs BUSTER vs ...
///////////////////////////////////////////////////////////////////////////////
type Manager interface {
	Install() (bool, error)
	Uninstall() (bool, error)
	Package() string
	Dependencies() []string
	InstallDependencies() (bool, error)
	Configs() ([]string, error)
	InstallConfigs() (bool, error)
	PostInstallCommands() []string
	RunPostInstallCommands() (bool, error)
	Installed() bool
}

type Source struct {
	Type string
	Address string
	Options []string
}

///////////////////////////////////////////////////////////////////////////////
type PackageManager struct {
	OS machine.OperatingSystem
	Sources []Source
}

func (self PackageManager) SourcesFile() string {
	return `
deb http://deb.debian.org/debian/ `+self.OS.Version+` main
deb-src http://deb.debian.org/debian/ bullseye main

deb http://security.debian.org/debian-security buster/updates main
deb-src http://security.debian.org/debian-security buster/updates main

# buster-updates, previously known as 'volatile'
deb http://deb.debian.org/debian/ bullseye-updates main
deb-src http://deb.debian.org/debian/ bullseye-updates main
`
}

///////////////////////////////////////////////////////////////////////////////
func New(os machine.OperatingSystem) PackageManager {
	return PackageManager{
		OS: os,
	}
}

///////////////////////////////////////////////////////////////////////////////
func (self PackageManager) Name() string {
	switch self.OS.Name {
	case "alpine":
		return "apk"
	case "debian":
		return "apt"
	default:
		return ""
	}
}

func (self PackageManager) EnvVars() string {
	switch self.OS.Name {
	case "alpine":
		return ""
	case "debian":
		return "DEBIAN_FRONTEND=noninteractive"
	default:
		return ""
	}
}

func (self PackageManager) Flags() string {
	switch self.OS.Name {
	case "alpine":
		return ""
	case "debian":
		return " -o Dpkg::Options::=--force-confdef -o Dpkg::Options::=--force-confnew -o APT::Install-Recommends=0 -y"
	default:
		return ""
	}
}

///////////////////////////////////////////////////////////////////////////////
type ActionType int

const (
	Install ActionType = iota
	Remove
	Update
	Upgrade
	DistributionUpgrade
	AutoRemove
	Clean // Autoremove, cleanup, etc
)

func (self ActionType) String(osName string) string {
	switch osName {
	case "alpine":
		switch self {
		case Install:
			return "add"
		case Remove:
			return "rm"
		case Update:
			return "update"
		case Upgrade:
			return "upgrade"
		case DistributionUpgrade:
			return  ""
		case AutoRemove:
			return ""
		case Clean:
			return "clean"
		default:
			return ""
		}
	case "debian": // Apt
		switch self {
		case Install:
			return "install"
		case Remove:
			return "remove"
		case Update:
			return "update"
		case Upgrade:
			return "upgrade"
		case DistributionUpgrade:
			return  "dist-upgrade"
		case AutoRemove:
			return "auto-remove"
		case Clean:
			return "clean"
		default:
			return ""
		}
	default:
		return ""
	}
}

func (self PackageManager) Action(action ActionType) string {
	return fmt.Sprintf("%s %s %s", self.EnvVars(), self.Name, action.String(self.OS.Name), self.Flags())
}

///////////////////////////////////////////////////////////////////////////////
func (self PackageManager) InstallPackage(pkg string) error { return terminal.Run(self.Action(Install) + fmt.Sprintf(" %s", pkg)) }

func (self PackageManager) InstallPackages(pkgs []string) (err error) {
	for _, pkg := range pkgs {
		err = self.InstallPackage(pkg)
	}
	return err
}

///////////////////////////////////////////////////////////////////////////////
func (self PackageManager) RemovePackage(pkg string) error { return terminal.Run(self.Action(Remove) + fmt.Sprintf(" %s", pkg)) }

func (self PackageManager) RemovePackages(pkgs []string) (err error) {
	for _, pkg := range pkgs {
		err = self.RemovePackage(pkg)
	}
	return err
}
