package incontainer

import (
	"bytes"
	"os"
)

func Detect() bool {
	return hasDockerEnv() || hasDockerCgroup() || pid1EnvHasContainer()
}

func hasDockerEnv() bool {
	_, err := os.Stat("/.dockerenv")
	return err == nil
}

func hasDockerCgroup() bool {
	cgroup, err := os.ReadFile("/proc/self/cgroup")
	if err != nil {
		return false
	}
	return bytes.Contains(cgroup, []byte("/docker/"))
}

func pid1EnvHasContainer() bool {
	bs, err := os.ReadFile("/proc/1/environ")
	if err != nil {
		return false
	}
	return bytes.Contains(bs, []byte("CONTAINER="))
}
