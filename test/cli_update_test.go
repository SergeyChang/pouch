package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"

	"github.com/alibaba/pouch/apis/types"
	"github.com/alibaba/pouch/test/command"
	"github.com/alibaba/pouch/test/environment"

	"github.com/go-check/check"
	"github.com/gotestyourself/gotestyourself/icmd"
)

// PouchUpdateSuite is the test suite for help CLI.
type PouchUpdateSuite struct{}

func init() {
	check.Suite(&PouchUpdateSuite{})
}

// SetUpSuite does common setup in the beginning of each test suite.
func (suite *PouchUpdateSuite) SetUpSuite(c *check.C) {
	SkipIfFalse(c, environment.IsLinux)

	environment.PruneAllContainers(apiClient)

	command.PouchRun("pull", busyboxImage).Assert(c, icmd.Success)
}

// TearDownTest does cleanup work in the end of each test.
func (suite *PouchUpdateSuite) TearDownTest(c *check.C) {
}

// TestUpdateMemory is to verify the correctness of updating container memory.
func (suite *PouchUpdateSuite) TestUpdateMemory(c *check.C) {
	//TODO
}

// TestUpdateRunningContainer is to verify the correctness of updating a running container.
func (suite *PouchUpdateSuite) TestUpdateRunningContainer(c *check.C) {
	name := "update-running-container"

	command.PouchRun("run", "-d", "-m", "300M", "--name", name, busyboxImage).Assert(c, icmd.Success)

	output := command.PouchRun("inspect", name).Stdout()
	result := &types.ContainerJSON{}
	if err := json.Unmarshal([]byte(output), result); err != nil {
		c.Errorf("failed to decode inspect output: %v", err)
	}
	containerID := result.ID

	file := "/sys/fs/cgroup/memory/default/" + containerID + "/memory.limit_in_bytes"
	if _, err := os.Stat(file); err != nil {
		c.Fatalf("container %s cgroup mountpoint not exists", containerID)
	}

	command.PouchRun("update", "-m", "500M", name).Assert(c, icmd.Success)

	out, err := exec.Command("cat", file).Output()
	if err != nil {
		c.Fatalf("execute cat command failed: %v", err)
	}

	if !strings.Contains(string(out), "524288000") {
		c.Fatalf("unexpected output %s expected %s\n", string(out), "524288000")
	}
}

// TestUpdateStoppedContainer is to verify the correctness of updating a stopped container.
func (suite *PouchUpdateSuite) TestUpdateStoppedContainer(c *check.C) {
	name := "update-stopped-container"

	command.PouchRun("create", "-m", "300M", "--name", name, busyboxImage).Assert(c, icmd.Success)

	output := command.PouchRun("inspect", name).Stdout()
	result := &types.ContainerJSON{}
	if err := json.Unmarshal([]byte(output), result); err != nil {
		c.Errorf("failed to decode inspect output: %v", err)
	}
	containerID := result.ID

	command.PouchRun("update", "-m", "500M", name).Assert(c, icmd.Success)

	command.PouchRun("start", name).Assert(c, icmd.Success)

	file := "/sys/fs/cgroup/memory/default/" + containerID + "/memory.limit_in_bytes"
	if _, err := os.Stat(file); err != nil {
		c.Fatalf("container %s cgroup mountpoint not exists", containerID)
	}

	out, err := exec.Command("cat", file).Output()
	if err != nil {
		c.Fatalf("execute cat command failed: %v", err)
	}

	if !strings.Contains(string(out), "524288000") {
		c.Fatalf("unexpected output %s expected %s\n", string(out), "524288000")
	}
}

// TestUpdateContainerInvalidValue is to verify the correctness of updating a container with invalid value.
func (suite *PouchUpdateSuite) TestUpdateContainerInvalidValue(c *check.C) {
	name := "update-container-with-invalid-value"

	command.PouchRun("run", "-d", "-m", "300M", "--name", name, busyboxImage).Assert(c, icmd.Success)

	res := command.PouchRun("update", "-m", "2M", name)
	c.Assert(res.Error, check.IsNil)

	// TODO add invalid value test.
}

// TestUpdateContainerWithoutFlag is to verify the correctness of updating a container without any flag.
func (suite *PouchUpdateSuite) TestUpdateContainerWithoutFlag(c *check.C) {
	name := "update-container-without-flag"

	command.PouchRun("run", "-d", "-m", "300M", "--name", name, busyboxImage).Assert(c, icmd.Success)

	command.PouchRun("update", name).Assert(c, icmd.Success)

}
