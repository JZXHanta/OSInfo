package osinfo

import (
	"fmt"
	"os/exec"

	// "runtime"
	"strings"
)

func WindowsOS() string {
	cmd := "(Get-WmiObject -class Win32_OperatingSystem).Caption"
	c := exec.Command("powershell", "-NoProfile", cmd)

	stdout, err := c.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return strings.TrimSpace(string(stdout))
}

func WindowsKernel() string {
	cmd := "[System.Environment]::OSVersion.Version.Build"
	c := exec.Command("powershell", "-NoProfile", cmd)

	stdout, err := c.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return strings.TrimSpace(string(stdout))
}

func LinuxDistro() string {
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

func LinuxVersion() string {
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

// func moreInfo() {
// 	switch runtime.GOOS {
// 	case "linux":
// 		fmt.Println("OS      : ", LinuxDistro())
// 		fmt.Println("Version : ", LinuxVersion())
// 	case "windows":
// 		fmt.Println("OS      : ", WindowsOS())
// 		fmt.Println("Version : ", WindowsKernel())
// 	}
// }

// func main() {
// 	moreInfo()
// }
