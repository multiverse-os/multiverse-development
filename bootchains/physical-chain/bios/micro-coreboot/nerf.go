// Copyright 2012-2018 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	flag "github.com/spf13/pflag"
)

const initramfs = "initramfs.linux_amd64.cpio"

var (
	configTxt = `loglevel=1
	init=/init
rootwait
`
	apt           = flag.Bool("apt", false, "apt-get all the things we need")
	fetch         = flag.Bool("fetch", false, "Fetch all the things we need")
	skipkern      = flag.Bool("skipkern", false, "Don't build the kernel")
	extra         = flag.String("extra", "", "Comma-separated list of extra packages to include")
	kernelVersion = "v4.12.7"
	workingDir    = ""
	linuxVersion  = "linux-stable"
	homeDir       = ""
	threads       = runtime.NumCPU() + 2 // Number of threads to use when calling make.
	packageList   = []string{
		"bc",
		"git",
		"golang",
		"build-essential",
		"git-core",
		"gitk",
		"git-gui",
		"iasl",
		"curl",
		"python2.7",
		"libyaml-dev",
		"liblzma-dev",
		"uuid-dev",
		"libssl-dev",
	}
)

func cp(inputLoc string, outputLoc string) error {
	// Don't check for an error, there are all kinds of
	// reasons a remove can fail even if the file is
	// writeable
	os.Remove(outputLoc)

	if _, err := os.Stat(inputLoc); err != nil {
		return err
	}
	fileContent, err := ioutil.ReadFile(inputLoc)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(outputLoc, fileContent, 0777)
}

func goGet() error {
	cmd := exec.Command("go", "get", "github.com/u-root/u-root")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	return cmd.Run()
}

func goBuildStatic() error {
	oFile := filepath.Join(workingDir, "linux-stable", initramfs)
	args := []string{"run", "github.com/u-root/u-root", "-o", oFile, "-build=bb"}
	cmd := exec.Command("go", append(args, staticCmdList...)...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	cmd = exec.Command("xz", "-f", "--check=crc32", "--lzma2=dict=512KiB", "linux-stable/initramfs.linux_amd64.cpio")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Printf("Created %v\n", oFile)
	return nil
}

func kernelGet() error {
	var args = []string{"clone", "--depth", "1", "-b", kernelVersion, "git://git.kernel.org/pub/scm/linux/kernel/git/stable/" + linuxVersion + ".git"}
	fmt.Printf("-------- Getting the kernel via git %v\n", args)
	cmd := exec.Command("git", args...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("didn't clone kernel %v", err)
		return err
	}
	return nil
}

func corebootGet() error {
	var args = []string{"https://coreboot.org/releases/coreboot-4.9.tar.xz"}
	fmt.Printf("-------- Getting coreboot via wget %v\n", "https://coreboot.org/releases/coreboot-4.9.tar.xz")
	cmd := exec.Command("wget", args...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("didn't wget coreboot %v", err)
		return err
	}
	cmd = exec.Command("tar", "xvf", "coreboot-4.9.tar.xz")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("untar failed %v", err)
		return err
	}
	cmd = exec.Command("make", "-j"+strconv.Itoa(threads), "crossgcc-i386", "iasl")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Dir = "coreboot-4.9"
	if err := cmd.Run(); err != nil {
		fmt.Printf("untar failed %v", err)
		return err
	}
	return nil
}

func buildKernel() error {
	if err := ioutil.WriteFile("linux-stable/.config", []byte(linuxconfig), 0666); err != nil {
		fmt.Printf("writing linux-stable/.config: %v", err)
		return err
	}

	cmd := exec.Command("make", "--directory", "linux-stable", "-j"+strconv.Itoa(threads))
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	// TODO: this is OK for now. Later we'll need to do something
	// with a map and GOARCH.
	cmd.Env = append(os.Environ(), "ARCH=x86_64")
	err := cmd.Run()
	if err != nil {
		return err
	}
	if _, err := os.Stat(filepath.Join("linux-stable", "/arch/x86/boot/bzImage")); err != nil {
		return err
	}
	fmt.Printf("bzImage created")
	return nil
}

func buildCoreboot() error {
	if err := ioutil.WriteFile("coreboot-4.9/.config", []byte(corebootconfig), 0666); err != nil {
		fmt.Printf("writing corebootconfig: %v", err)
		return err
	}
	if err := cp("linux-stable/arch/x86/boot/bzImage", "coreboot-4.9/bzImage"); err != nil {
		fmt.Printf("copying %v to linux-stable/.config: %v", err)
	}

	cmd := exec.Command("make", "-j"+strconv.Itoa(threads))
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Env = append(os.Environ(), "ARCH=x86_64")
	cmd.Dir = "coreboot-4.9"
	err := cmd.Run()
	if err != nil {
		return err
	}
	if _, err := os.Stat("coreboot-4.9/build/coreboot.rom"); err != nil {
		return err
	}
	fmt.Printf("bzImage created")
	return nil
}

func run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%v %v: %v", name, args, err)
	}
	return nil
}

func check() error {
	if os.Getenv("GOPATH") == "" {
		return fmt.Errorf("You have to set GOPATH.")
	}
	return nil
}

func cleanup() error {
	filesToRemove := [...]string{"linux-stable", "vboot_reference", "linux-firmware"}
	fmt.Printf("-------- Removing problematic files %v\n", filesToRemove)
	for _, file := range filesToRemove {
		if _, err := os.Stat(file); err != nil {
			if os.IsNotExist(err) {
				continue
			}
		}
		err := os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func aptget() error {
	missing := []string{}
	for _, packageName := range packageList {
		cmd := exec.Command("dpkg", "-s", packageName)
		if err := cmd.Run(); err != nil {
			missing = append(missing, packageName)
		}
	}

	if len(missing) == 0 {
		fmt.Println("No missing dependencies to install")
		return nil
	}

	fmt.Printf("Using apt-get to get %v\n", missing)
	get := []string{"apt-get", "-y", "install"}
	get = append(get, missing...)
	cmd := exec.Command("sudo", get...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	return cmd.Run()

}

func allFunc() error {
	var cmds = []struct {
		f      func() error
		skip   bool
		ignore bool
		n      string
	}{
		{f: check, skip: false, ignore: false, n: "check environment"},
		{f: cleanup, skip: *skipkern || !*fetch, ignore: false, n: "cleanup"},
		{f: goGet, skip: *skipkern || !*fetch, ignore: false, n: "Get u-root source"},
		{f: aptget, skip: !*apt, ignore: false, n: "apt get"},
		{f: kernelGet, skip: *skipkern || !*fetch, ignore: false, n: "Git clone the kernel"},
		{f: corebootGet, skip: *skipkern || !*fetch, ignore: false, n: "Git clone coreboot"},
		{f: goBuildStatic, skip: *skipkern, ignore: false, n: "Build static initramfs"},
		{f: buildKernel, skip: *skipkern, ignore: false, n: "build the kernel"},
		{f: buildCoreboot, skip: *skipkern, ignore: false, n: "build coreboot"},
	}

	for _, c := range cmds {
		log.Printf("-----> Step %v: ", c.n)
		if c.skip {
			log.Printf("-------> Skip")
			continue
		}
		log.Printf("----------> Start")
		err := c.f()
		if c.ignore {
			log.Printf("----------> Ignore result")
			continue
		}
		if err != nil {
			return fmt.Errorf("%v: %v", c.n, err)
		}
		log.Printf("----------> Finished %v\n", c.n)
	}
	return nil
}

func main() {
	flag.Parse()
	log.Printf("Using kernel %v\n", kernelVersion)
	if err := allFunc(); err != nil {
		log.Fatalf("fail error is : %v", err)
	}
	log.Printf("execution completed successfully\n")
}
