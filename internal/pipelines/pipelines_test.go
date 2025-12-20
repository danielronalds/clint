package pipelines

import "testing"

func TestRun(t *testing.T) {
	t.Run("Full Success", func(t *testing.T) {
		pipeline := &Pipeline{
			Name: "Happy Path",
			Steps: []Step{
				{Name: "Step 1", Cmd: "true"},
				{Name: "Step 2", Cmd: "echo hello"},
			},
		}

		if got := Run(pipeline); got != true {
			t.Errorf("Run() = %v; want true", got)
		}
	})

	t.Run("Stop on Failure", func(t *testing.T) {
		pipeline := &Pipeline{
			Name: "Sad Path",
			Steps: []Step{
				{Name: "Fail Fast", Cmd: "false"},
				{Name: "Never Runs", Cmd: "echo should_not_see_this"},
			},
		}

		if got := Run(pipeline); got != false {
			t.Errorf("Run() = %v; want false", got)
		}
	})
}

func TestRunStep(t *testing.T) {
	tests := []struct {
		name     string
		cmd      string
		expected bool
	}{
		{"valid echo", "echo hi", true},
		{"true", "true", true},
		{"false", "false", false},
		{"invalid command", "fakedoer", false},
		{"empty command", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			step := Step{Cmd: tt.cmd}
			if _, got := runStep(step); got != tt.expected {
				t.Errorf("runStep(%q) = %v; want %v", tt.cmd, got, tt.expected)
			}
		})
	}
}
