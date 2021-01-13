package packagemanager

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"unicode"
	"unicode/utf8"
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

func Source struct {
	Type string
	Address string
	Options []string
}


///////////////////////////////////////////////////////////////////////////////
type PackageManager struct {
	OS OperatingSystem
	
	Sources []Source
}

func (self PackageManager) SourcesFile() string {
	return `
deb http://deb.debian.org/debian/ `+self.OS.Version.Name+` main
deb-src http://deb.debian.org/debian/ bullseye main

deb http://security.debian.org/debian-security buster/updates main
deb-src http://security.debian.org/debian-security buster/updates main

# buster-updates, previously known as 'volatile'
deb http://deb.debian.org/debian/ bullseye-updates main
deb-src http://deb.debian.org/debian/ bullseye-updates main
`
}

///////////////////////////////////////////////////////////////////////////////
func New(os OperatingSystem) PackageManager {
	return PackageManager{
		OperatingSystem: os,
	}
}

///////////////////////////////////////////////////////////////////////////////
type OperatingSystem struct {
	Name OSName
	Version OSVersion

}

type OSName int

const (
	Debian OSName = iota
	Alpine
)

type OSVersion struct {
	Name string
	Number string
}


///////////////////////////////////////////////////////////////////////////////
func (self PackageManager) EnvironmentalVaraibles() string {
	return "DEBIAN_FRONTEND=noninteractive"
}

func (self PackageManager) Flags() string {
	switch self.OS {
	case Debian:
		return " -o Dpkg::Options::=--force-confdef -o Dpkg::Options::=--force-confnew -o APT::Install-Recommends=0 -y"
	case Apt:
		return ""
	default:
		return ""
	}
}

func (self PaackageManager) Name() string {
	switch self {
	case Debian:
		return "apt"
	case Alpine:
		return "apt"
	}
}

///////////////////////////////////////////////////////////////////////////////
func (self PackageManager) Install() string {
	switch self.Type {
	case Apk:
		return "apk add"
	default: // Apt
		return "DEBIAN_FRONTEND=noninteractive apt install -o Dpkg::Options::=--force-confdef -o Dpkg::Options::=--force-confnew -o APT::Install-Recommends=0 -y"
	}
}

func (self PackageManager) InstallPackage(pkg string) error {
	return Terminal(self.Install() + ` ` + pkg)
}

func (self PackageManager) InstallPackages(pkgs []string) error {
	return Terminal(self.Install() + ` ` + strings.Join(pkgs, " "))
}

///////////////////////////////////////////////////////////////////////////////
func (self PackageManager) Remove() string {
	switch self.Type {
	case Apk:
		return "apk rm"
	default: // Apt
		return "DEBIAN_FRONTEND=noninteractive apt remove -y"
	}
}

func (self PackageManager) RemovePackage(pkg string) error {
	return Terminal(self.Remove() + ` ` + pkg)
}

func (self PackageManager) RemovePackages(pkgs []string) error {
	return Terminal(self.Remove() + ` ` + strings.Join(pkgs, " "))
}

func (self PackageManager) Autoremove() error {
	switch self.Type {
	default: // Apt
		return terminal.Run("DEBIAN_FRONTEND=noninteractive apt autoremove -y")
	}
}

///////////////////////////////////////////////////////////////////////////////
func (self PackageManager) Update() error {
	switch self.Type {
	case Apk:
		return Terminal("apk update")
	default: // Apt
		return Terminal("DEBIAN_FRONTEND=noninteractive apt update -y")
	}
}

func (self PackageManager) Upgrade() error {
	switch self.Type {
	case Apk:
		return Terminal("apk upgrade")
	default: // Apt
		return Terminal("DEBIAN_FRONTEND=noninteractive apt upgrade -o Dpkg::Options::=--force-confdef -o Dpkg::Options::=--force-confnew -y")
	}
}

func (self PackageManager) DistUpgrade() error {
	switch self.Type {
	default: // Apt
		return Terminal("DEBIAN_FRONTEND=noninteractive apt dist-upgrade -o Dpkg::Options::=--force-confdef -o Dpkg::Options::=--force-confnew -y")
	}
}




