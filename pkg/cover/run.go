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

	// gas lint warns about possible injection here, but we're fine
	cmd := exec.Command(goExe, "test", "-cover", "./...") // nolint: gas,gosec

	output, err := cmd.CombinedOutput()
	return output, err
}
