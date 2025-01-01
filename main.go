package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func findPackagemanger() string {
	packageManagers := []string{"apt", "yum", "dnf", "zypper", "pacman", "emerge", "apk"}
	for _, pm := range packageManagers {
		if _, err := os.Stat("/usr/bin/" + pm); err == nil {
			return pm
		} else if _, err := os.Stat("sbin/" + pm); err == nil {
			return pm
		}
	}
	return ""
}

func InstallPackage(pm, pkg string) {
	fmt.Printf("Installing %s with %s\n", pkg, pm)
	var cmd *exec.Cmd
	switch pm {
	case "apt":
		fallthrough
	case "yum":
		fallthrough
	case "dnf":
		fallthrough
	case "zypper":
		fallthrough
	case "pacman":
		fallthrough
	case "emerge":
		fmt.Printf("Installing %s with %s\n", pkg, pm)

		cmdStr := fmt.Sprintf("%s install -y %s", pm, pkg)

		cmd = exec.Command("sh", "-c", cmdStr)

		// Capture the output and err
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error installing package: ", err)
			return
		}
		fmt.Println(string(out))
		fmt.Printf("Package %s installed successfully with %s\n", pkg, pm)
	case "apk":
		fmt.Printf("Installing %s with %s\n", pkg, pm)
		cmdStr := fmt.Sprintf("%s add %s", pm, pkg)
		cmd = exec.Command("sh", "-c", cmdStr)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error installing package: ", err)
			return
		}
		fmt.Println(string(out))
		fmt.Printf("Package %s installed successfully with %s\n", pkg, pm)

	default:
		fmt.Println("Unknown package manager.")

	}
}

func main() {
	if runtime.GOOS != "linux" && runtime.GOARCH != "amd64" {
		fmt.Println("This program only runs on Linux 64-bit.")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: installerx <package1> <package2> ...")
		os.Exit(1)
	}

	packages := os.Args[1:]
	fmt.Println("Packages to install :", strings.Join(packages, ", "))

	PackageManger := findPackagemanger() // Checking if package Manger Exist
	if PackageManger == "" {
		fmt.Println("No package manager found.")
		os.Exit(1)
	}

	// Installing the packages using the package manager
	for _, pkg := range packages {
		InstallPackage(PackageManger, pkg)
	}
}
