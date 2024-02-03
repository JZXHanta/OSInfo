package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func windowsOS() string {
	cmd := "(Get-WmiObject -class Win32_OperatingSystem).Caption"
	c := exec.Command("powershell", "-NoProfile", cmd)

	stdout, err := c.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return strings.TrimSpace(string(stdout))
}

func windowsKernel() string {
	cmd := "[System.Environment]::OSVersion.Version.Build"
	c := exec.Command("powershell", "-NoProfile", cmd)

	stdout, err := c.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return strings.TrimSpace(string(stdout))
}

func linuxDistro() string {
	cmd := "lsb_release"
	arg1 := "-i"

	c := exec.Command(cmd, arg1)

	stdout, err := c.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	out := string(stdout)
	clean := strings.SplitAfter(out, " ")[1]
	cleaner := strings.SplitAfter(clean, ":")[1]

	return strings.TrimSpace(cleaner)
}

func linuxVersion() string {
	cmd := "lsb_release"
	arg1 := "-r"

	c := exec.Command(cmd, arg1)

	stdout, err := c.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	out := string(stdout)

	clean := strings.SplitAfter(out, ":")[1]

	return strings.TrimSpace(clean)
}

func moreInfo() {
	switch runtime.GOOS {
	case "linux":
		fmt.Println("OS      : ", linuxDistro())
		fmt.Println("Version : ", linuxVersion())
	case "windows":
		fmt.Println("OS      : ", windowsOS())
		fmt.Println("Version : ", windowsKernel())
	}
}

func main() {
	moreInfo()
}
