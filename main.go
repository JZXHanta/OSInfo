package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func getPlatform() string {
	return runtime.GOOS
}

func windowsOS() string {
	cmd := "(Get-WmiObject -class Win32_OperatingSystem).Caption"
	c := exec.Command("powershell", "-NoProfile", cmd)

	stdout, err := c.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(stdout)
}

func windowsKernel() string {
	cmd := "[System.Environment]::OSVersion.Version.Build"
	c := exec.Command("powershell", "-NoProfile", cmd)

	stdout, err := c.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(stdout)
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
	cleanest := strings.TrimSpace(cleaner)

	return cleanest
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
	cleanest := strings.TrimSpace(clean)

	return cleanest
}

func moreInfo() {
	switch runtime.GOOS {
	case "linux":
		fmt.Println("LINUX FUNCTION HERE")
		fmt.Print("OS      : ", linuxDistro())
		fmt.Print("Version : ", linuxVersion())
	case "windows":
		fmt.Print("OS      : ", windowsOS())
		fmt.Print("Version : ", windowsKernel())
	}
}

func main() {
	fmt.Println(getPlatform())
	moreInfo()
}
