package packagemanager

import (
	"strings"

	machine "../"
)


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
func (self PackageManager) EnvironmentalVaraibles() string {
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

///////////////////////////////////////////////////////////////////////////////
func (self PackageManager) Install() string {
	switch self.OS.Name {
	case "alpine":
		return "add"
	case "debian": // Apt
		return "install"
	default:
		return ""
	}
}


///////////////////////////////////////////////////////////////////////////////
type Action int 

const (
	Install Action = iota
	Remove
	Update
	Upgrade
	Maintainance // Autoremove, cleanup, etc
)


func (self PackageManager) Action(action Action) error {

	return terminal.Run(fmt.Sprintf("%s %s %s", self.EnvironmentalVariables(), self.Name, self.Action)
}

///////////////////////////////////////////////////////////////////////////////
func (self PackageManager) InstallPackage(pkg string) error {
}

func (self PackageManager) InstallPackages(pkgs []string) error {
	return terminal.Run(self.Install() + ` ` + strings.Join(pkgs, " "))
}

///////////////////////////////////////////////////////////////////////////////
func (self PackageManager) Remove() string {
	switch self.OS.Name {
	case "alpine":
		return "rm"
	case "debian":
		return "remove"
	default:
		return ""
	}
}

func (self PackageManager) RemovePackage(pkg string) error {
	return Terminal(self.Remove() + ` ` + pkg)
}

func (self PackageManager) RemovePackages(pkgs []string) error {
	return Terminal(self.Remove() + ` ` + strings.Join(pkgs, " "))
}

func (self PackageManager) Maintainance() error {
	switch self.OS.Name {
	case "alpine":
		return ""
	case "debian":
		return "autoremove"
	default:
		return ""
	}
}

///////////////////////////////////////////////////////////////////////////////
func (self PackageManager) Update() error {
	switch self.OS.Name {
	case "alpine":
		return "update"
	case "debian":
		return "update"
	default:
		return ""
	}
}

func (self PackageManager) Upgrade() error {
	switch self.OS.Name {
	case "alpine":
		return "upgrade"
	case "debian":
		return "upgrade"
	default:
		return ""
	}
}

func (self PackageManager) DistUpgrade() error {
	switch self.OS.Name {
	case "alpine":
		return ""
	case "debian":
		return "dist-upgrade"
	default:
		return ""
	}
}




