package commands

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/ribeirosaimon/motion-go/cli/domain"
	"github.com/ribeirosaimon/motion-go/confighub/util"
	"github.com/spf13/cobra"
)

func UpAppCommand(use, help string) *cobra.Command {
	// motionFlag := domain.NewCommandFlagMotion("nome", "n", "Testar Comando", "World")
	ctx, cancel := context.WithCancel(context.Background())

	return domain.NewCommandMotion(use, help, func(cmd *cobra.Command, strings []string) {
		uppMotionApi(ctx, cancel)
	})
}

func uppMotionApi(ctx context.Context, cancel context.CancelFunc) {
	// dir, err := util.FindRootDir()
	dir, err := util.FindModuleDir("api")
	if err != nil {
		panic(err)
	}
	programPath := fmt.Sprintf("%s/cmd/main.go", dir)
	cmd := exec.CommandContext(ctx, "go", "run", programPath)

	// Set the environment variables if needed
	// cmd.Env = append(os.Environ(), "MY_ENV_VARIABLE=value")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		cancel()
		contextDone(ctx)
	}

}

func contextDone(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Context is canceled")
	case <-time.After(5 * time.Second):
		// Simulate some work
		fmt.Println("Work is done")
	}
}
