package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/ViolaPeracia/Weather-cli/internal/config"
)

func TestTuiQuitCommand(t *testing.T) {
	input := "q\n"
	stdin := bytes.NewBufferString(input)
	var stdout bytes.Buffer

	cfg := config.Config{
		DefaultCity: "Hanoi",
		Unit:        "celsius",
	}

	runTUI(stdin, &stdout, cfg)

	output := stdout.String()
	if !strings.Contains(output, "Goodbye!") {
		t.Errorf("Expected output to contain 'Goodbye!', got:\n%s", output)
	}
}

func TestTuiHelpToggle(t *testing.T) {
	// Toggle help on ("h") and then quit ("q")
	input := "h\nq\n"
	stdin := bytes.NewBufferString(input)
	var stdout bytes.Buffer

	cfg := config.Config{
		DefaultCity: "Hanoi",
		Unit:        "celsius",
	}

	runTUI(stdin, &stdout, cfg)

	output := stdout.String()
	if !strings.Contains(output, "Help & Guidance") {
		t.Errorf("Expected output to contain Help page, got:\n%s", output)
	}
}
