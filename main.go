package main

import (
	"context"
	"github.com/Mona-bele/action-integration-terraform/pkg/run"
	"github.com/Mona-bele/action-integration-terraform/pkg/utils"
	"os"
)

func main() {
	utils.ArgsInput()
	tfRunType := os.Getenv("tf_run_type")
	builder, err := run.GetBuilder(tfRunType)
	if err != nil {
		utils.SetGithubEnvOutput("error", err.Error())
	}
	ctx := context.Background()
	builder.GetRunBuilder(ctx)
}
