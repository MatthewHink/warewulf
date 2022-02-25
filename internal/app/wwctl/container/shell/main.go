//go:build linux
// +build linux

package shell

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	//wwapi "github.com/hpcng/warewulf/internal/pkg/api/container"
	//"github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"

	"github.com/hpcng/warewulf/internal/pkg/container"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
	"github.com/spf13/cobra"
)

func CobraRunE(cmd *cobra.Command, args []string) (err error) {

	/*
	csp := &wwapiv1.ContainerShellParameter{
		ContainerName: args[0],
		Binds: binds,
		Args: args[1:],
	}

	return wwapi.ContainerShell(csp)
	*/

	
	containerName := args[0]
	var allargs []string

	if !container.ValidSource(containerName) {
		wwlog.Printf(wwlog.ERROR, "Unknown Warewulf container: %s\n", containerName)
		os.Exit(1)
	}

	for _, b := range binds {
		allargs = append(allargs, "--bind", b)
	}
	allargs = append(allargs, args...)
	allargs = append(allargs, "/usr/bin/bash")

	fmt.Printf("allargs: %v\n", allargs)

	c := exec.Command("/proc/self/exe", append([]string{"container", "exec"}, allargs...)...)

	//c := exec.Command("/bin/sh")
	c.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	wwlog.Printf(wwlog.VERBOSE, "command: %s\n", c)

	if err := c.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil

}
