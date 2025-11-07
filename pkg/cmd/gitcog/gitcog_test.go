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
