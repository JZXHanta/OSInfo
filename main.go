package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func getPlatform() string {
	return runtime.GOOS
}

func windowsOS() {
	cmd := "(Get-WmiObject -class Win32_OperatingSystem).Caption"
	// cmd := "echo"
	// arg := "[System.Environment]::OSVersion.Version.Build"
	c := exec.Command("powershell", "-NoProfile", cmd)

	stdout, err := c.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
}

func moreInfo() {
	switch runtime.GOOS {
	case "linux":
		fmt.Println("LINUX FUNCTION HERE")
	case "windows":
		windowsOS()
	}
}

func main() {
	fmt.Println(getPlatform())
	moreInfo()
}
