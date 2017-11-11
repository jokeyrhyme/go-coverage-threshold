package cover

import (
	"os/exec"
)

// Run executes `go test -cover ./...` and returns the raw result.
func Run() ([]byte, error) {
	goExe, err := exec.LookPath("go")
	if err != nil {
		return nil, nil
	}

	cmd := exec.Command(goExe, "test", "-cover", "./...")

	// TODO: stream the output to stdout

	output, err := cmd.CombinedOutput()
	return output, err
}
