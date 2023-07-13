package main

import (
	"context"
	"github.com/Mona-bele/action-integration-terraform/internal/utils"
	"github.com/Mona-bele/action-integration-terraform/pkg/run"
	"os"
)

func main() {
	tfRunType := os.Getenv("tf_run_type")
	builder, err := run.GetBuilder(tfRunType)
	if err != nil {
		utils.SetGithubEnvOutput("error", err.Error())
	}
	ctx := context.Background()
	builder.GetRunBuilder(ctx)
}
