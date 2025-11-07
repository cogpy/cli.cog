package gitcog

import (
	"testing"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/stretchr/testify/assert"
)

func TestNewCmdGitCog(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := NewCmdGitCog(f)
	assert.NotNil(t, cmd)
	assert.Equal(t, "gitcog <command>", cmd.Use)
	assert.Contains(t, cmd.Short, "autonomous multi-agent orchestration")
	
	// Verify subcommands exist
	assert.True(t, cmd.HasSubCommands())
	
	// Test that help works
	cmd.SetOut(stdout)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
}

func TestGitCogStart(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdStart(f)
	assert.NotNil(t, cmd)
	
	cmd.SetOut(stdout)
	err := cmd.Execute()
	assert.NoError(t, err)
	
	output := stdout.String()
	assert.Contains(t, output, "GitCog Orchestration System Starting")
	assert.Contains(t, output, "Meta-cognitive layer initialized")
}

func TestGitCogStatus(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdStatus(f)
	assert.NotNil(t, cmd)
	
	cmd.SetOut(stdout)
	err := cmd.Execute()
	assert.NoError(t, err)
	
	output := stdout.String()
	assert.Contains(t, output, "GitCog Orchestration System Status")
	assert.Contains(t, output, "System State:")
	assert.Contains(t, output, "Active")
}

func TestGitCogAgents(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdAgents(f)
	assert.NotNil(t, cmd)
	
	cmd.SetOut(stdout)
	err := cmd.Execute()
	assert.NoError(t, err)
	
	output := stdout.String()
	assert.Contains(t, output, "Active Cognitive Agents")
	assert.Contains(t, output, "No agents currently active")
}

func TestGitCogOrchestrate(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdOrchestrate(f)
	assert.NotNil(t, cmd)
	
	cmd.SetOut(stdout)
	cmd.SetArgs([]string{"Optimize workflow"})
	err := cmd.Execute()
	assert.NoError(t, err)
	
	output := stdout.String()
	assert.Contains(t, output, "Meta-Cognitive Analysis")
	assert.Contains(t, output, "Optimize workflow")
	assert.Contains(t, output, "Abstract Reasoning")
	assert.Contains(t, output, "Orchestration complete")
}

func TestGitCogOrchestrateNoArgs(t *testing.T) {
	ios, _, _, stderr := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdOrchestrate(f)
	cmd.SetErr(stderr)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	
	// Should fail with no arguments
	assert.Error(t, err)
}

func TestGitCogDaemon(t *testing.T) {
	ios, _, _, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdDaemon(f)
	assert.NotNil(t, cmd)
	assert.Equal(t, "daemon <command>", cmd.Use)
	assert.Contains(t, cmd.Short, "daemon")
	
	// Verify subcommands exist
	assert.True(t, cmd.HasSubCommands())
}

func TestGitCogDaemonStart(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdDaemonStart(f)
	assert.NotNil(t, cmd)
	
	cmd.SetOut(stdout)
	cmd.SetArgs([]string{"--name", "test-daemon", "--type", "code-review"})
	err := cmd.Execute()
	assert.NoError(t, err)
	
	output := stdout.String()
	assert.Contains(t, output, "Starting GitCog daemon")
	assert.Contains(t, output, "test-daemon")
	assert.Contains(t, output, "code-review")
	assert.Contains(t, output, "Daemon process initialized")
}

func TestGitCogDaemonStartAutoName(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdDaemonStart(f)
	cmd.SetOut(stdout)
	err := cmd.Execute()
	assert.NoError(t, err)
	
	output := stdout.String()
	assert.Contains(t, output, "daemon-")
}

func TestGitCogDaemonStop(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdDaemonStop(f)
	assert.NotNil(t, cmd)
	
	cmd.SetOut(stdout)
	cmd.SetArgs([]string{"test-daemon"})
	err := cmd.Execute()
	assert.NoError(t, err)
	
	output := stdout.String()
	assert.Contains(t, output, "Stopping daemon")
	assert.Contains(t, output, "test-daemon")
	assert.Contains(t, output, "Graceful shutdown initiated")
}

func TestGitCogDaemonStopNoArgs(t *testing.T) {
	ios, _, _, stderr := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdDaemonStop(f)
	cmd.SetErr(stderr)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	
	// Should fail with no arguments
	assert.Error(t, err)
}

func TestGitCogDaemonList(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdDaemonList(f)
	assert.NotNil(t, cmd)
	
	cmd.SetOut(stdout)
	err := cmd.Execute()
	assert.NoError(t, err)
	
	output := stdout.String()
	assert.Contains(t, output, "Active GitCog Daemons")
	assert.Contains(t, output, "No daemons currently running")
}

func TestGitCogDaemonLogs(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdDaemonLogs(f)
	assert.NotNil(t, cmd)
	
	cmd.SetOut(stdout)
	cmd.SetArgs([]string{"test-daemon"})
	err := cmd.Execute()
	assert.NoError(t, err)
	
	output := stdout.String()
	assert.Contains(t, output, "Logs for daemon")
	assert.Contains(t, output, "test-daemon")
	assert.Contains(t, output, "Daemon initialized")
}

func TestGitCogDaemonLogsFollow(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdDaemonLogs(f)
	cmd.SetOut(stdout)
	cmd.SetArgs([]string{"test-daemon", "--follow"})
	err := cmd.Execute()
	assert.NoError(t, err)
	
	output := stdout.String()
	assert.Contains(t, output, "Following logs")
}

func TestGitCogConstellation(t *testing.T) {
	ios, _, _, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdConstellation(f)
	assert.NotNil(t, cmd)
	assert.Equal(t, "constellation <command>", cmd.Use)
	assert.Contains(t, cmd.Short, "constellation")
	
	// Verify subcommands exist
	assert.True(t, cmd.HasSubCommands())
}

func TestGitCogConstellationCreate(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdConstellationCreate(f)
	assert.NotNil(t, cmd)
	
	cmd.SetOut(stdout)
	cmd.SetArgs([]string{"--name", "ci-pipeline", "--purpose", "CI/CD automation"})
	err := cmd.Execute()
	assert.NoError(t, err)
	
	output := stdout.String()
	assert.Contains(t, output, "Creating constellation")
	assert.Contains(t, output, "ci-pipeline")
	assert.Contains(t, output, "CI/CD automation")
	assert.Contains(t, output, "Constellation framework initialized")
}

func TestGitCogConstellationCreateNoName(t *testing.T) {
	ios, _, _, stderr := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdConstellationCreate(f)
	cmd.SetErr(stderr)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	
	// Should fail without --name flag
	assert.Error(t, err)
}

func TestGitCogConstellationList(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdConstellationList(f)
	assert.NotNil(t, cmd)
	
	cmd.SetOut(stdout)
	err := cmd.Execute()
	assert.NoError(t, err)
	
	output := stdout.String()
	assert.Contains(t, output, "Active Daemon Constellations")
	assert.Contains(t, output, "No constellations currently active")
}

func TestGitCogConstellationAdd(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdConstellationAdd(f)
	assert.NotNil(t, cmd)
	
	cmd.SetOut(stdout)
	cmd.SetArgs([]string{"ci-pipeline", "test-daemon"})
	err := cmd.Execute()
	assert.NoError(t, err)
	
	output := stdout.String()
	assert.Contains(t, output, "Adding daemon")
	assert.Contains(t, output, "test-daemon")
	assert.Contains(t, output, "ci-pipeline")
	assert.Contains(t, output, "Coordination protocols established")
}

func TestGitCogConstellationAddNoArgs(t *testing.T) {
	ios, _, _, stderr := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdConstellationAdd(f)
	cmd.SetErr(stderr)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	
	// Should fail with no arguments
	assert.Error(t, err)
}

func TestGitCogConstellationStatus(t *testing.T) {
	ios, _, stdout, _ := iostreams.Test()
	f := &cmdutil.Factory{
		IOStreams: ios,
	}

	cmd := newCmdConstellationStatus(f)
	assert.NotNil(t, cmd)
	
	cmd.SetOut(stdout)
	cmd.SetArgs([]string{"ci-pipeline"})
	err := cmd.Execute()
	assert.NoError(t, err)
	
	output := stdout.String()
	assert.Contains(t, output, "Constellation Status")
	assert.Contains(t, output, "ci-pipeline")
	assert.Contains(t, output, "Member Daemons")
	assert.Contains(t, output, "Active Workflows")
}
