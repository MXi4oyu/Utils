package subprocess


import (
	"fmt"
	"context"
	"os/exec"
	"bytes"
)

func RunCommand(ctx context.Context, cmd string, args ...string) (string, error) {

	var c *exec.Cmd

	if ctx != nil {
		c = exec.CommandContext(ctx, cmd, args...)
	} else {
		c = exec.Command(cmd, args...)
	}


	var out bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &out
	c.Stderr = &stderr

	err:=c.Run()

	if err != nil {
		fmt.Println(fmt.Sprint("exec Command error") + ": " + stderr.String())
		return out.String(),err
	}

	// check for exec context timeout
	if ctx != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("command %s timed out", cmd)
		}
	}

	return out.String(), nil
}
