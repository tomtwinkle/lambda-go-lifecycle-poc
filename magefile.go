// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path/filepath"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build
type BuildEnv struct {
	OS           string
	ARCH         string
	File         string
	Compress     string
	CompressFile string
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)

	var buildEnvList = []BuildEnv{
		{"linux", "amd64", "main", "zip", "main.zip"},
	}

	fmt.Println("Building ...")
	for _, e := range buildEnvList {
		filePath := filepath.Join("builds", e.File)
		cmd := exec.Command("go", "build", "-o", filePath, "-ldflags", "-s -w", ".")
		cmd.Env = append(os.Environ(),
			fmt.Sprintf("GOOS=%s", e.OS),
			fmt.Sprintf("GOARCH=%s", e.ARCH),
		)
		if err := cmd.Run(); err != nil {
			fmt.Printf("%v\n", err)
			return errors.WithStack(err)
		}
		compressFilePath := filepath.Join("builds", e.CompressFile)

		cmd = exec.Command("build-lambda-zip", "-o", compressFilePath, filePath)
		if err := cmd.Run(); err != nil {
			fmt.Printf("%v\n", err)
			return errors.WithStack(err)
		}

		os.Remove(filePath)
	}
	return nil
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	return nil
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "get", "-u", "github.com/aws/aws-lambda-go/cmd/build-lambda-zip")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("builds")
}
