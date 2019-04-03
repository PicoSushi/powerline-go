package main

import (
	"fmt"
	"os"
	"strings"
)

// getShell returns current shell name.
func getShellName() string {
	env_shell, _ := os.LookupEnv("SHELL")
	path := strings.Split(env_shell, "/")

	return path[len(path)-1]
}

func segmentExectionTime(p *powerline) {
	shell := getShellName()
	if shell != "fish" {
		// Only support fish, currently
		return
	}

	te, _ := os.LookupEnv("CMD_DURATION")
	fmt.Println("TE:", te)
	p.appendSegment("execution-time", segment{
		content:    te,
		foreground: p.theme.ExecutionTimeFg,
		background: p.theme.ExecutionTimeBg,
	})
}
