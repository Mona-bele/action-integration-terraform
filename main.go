package main

import (
	"context"
	"github.com/Mona-bele/action-integration-terraform/pkg/run"
	"github.com/Mona-bele/action-integration-terraform/pkg/utils"
	"os"
)

func main() {
	tfRunType := os.Getenv("INPUT_TF_RUN_TYPE")
	builder, err := run.GetBuilder(tfRunType)
	if err != nil {
		utils.SetGithubEnvOutput("error", err.Error())
	}
	ctx := context.Background()
	builder.GetRunBuilder(ctx)
}
