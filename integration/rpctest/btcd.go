// Copyright (c) 2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rpctest

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	// compileMtx guards access to the executable path so that the project is
	// only compiled once.
	compileMtx sync.Mutex

	// executablePath is the path to the compiled executable. This is the empty
	// string until vipsd is compiled. This should not be accessed directly;
	// instead use the function vipsdExecutablePath().
	executablePath string
)

// vipsdExecutablePath returns a path to the vipsd executable to be used by
// rpctests. To ensure the code tests against the most up-to-date version of
// vipsd, this method compiles vipsd the first time it is called. After that, the
// generated binary is used for subsequent test harnesses. The executable file
// is not cleaned up, but since it lives at a static path in a temp directory,
// it is not a big deal.
func vipsdExecutablePath() (string, error) {
	compileMtx.Lock()
	defer compileMtx.Unlock()

	// If vipsd has already been compiled, just use that.
	if len(executablePath) != 0 {
		return executablePath, nil
	}

	testDir, err := baseDir()
	if err != nil {
		return "", err
	}

	// Build vipsd and output an executable in a static temp path.
	outputPath := filepath.Join(testDir, "vipsd")
	if runtime.GOOS == "windows" {
		outputPath += ".exe"
	}
	cmd := exec.Command(
		"go", "build", "-o", outputPath, "github.com/vipstar-dev/vipsd",
	)
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Failed to build vipsd: %v", err)
	}

	// Save executable path so future calls do not recompile.
	executablePath = outputPath
	return executablePath, nil
}
