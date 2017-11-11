package cover

import (
	"os/exec"
)

// Run executes `go test -cover ./...` and returns the raw result.
func Run() ([]byte, error) {
	goExe, err := exec.LookPath("go")
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(goExe, "test", "-cover", "./...")

	output, err := cmd.CombinedOutput()
	return output, err
}
