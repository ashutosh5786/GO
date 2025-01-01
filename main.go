package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func findPackagemanger() string {
	packageManagers := []string{"apt", "yum", "dnf", "zypper", "pacman", "emerge"}
	for _, pm := range packageManagers {
		if _, err := os.Stat("/usr/bin/" + pm); err == nil {
			return pm
		}
	}
	return ""
}

func main() {
	if runtime.GOOS != "linux" && runtime.GOARCH != "amd64" {
		fmt.Println("This program only runs on Linux 64-bit.")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <package1> <package2> ...")
		os.Exit(1)
	}

	packages := os.Args[1:]
	fmt.Println("Packages to install :", strings.Join(packages, ", "))


	PackageManger := findPackagemanger() // Checking if package Manger Exist
	if PackageManager == "" {
		fmt.Println("No package manager found.")
		os.Exit(1)
	}
}
