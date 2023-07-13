package main

import (
	"context"
	"github.com/Mona-bele/action-integration-terraform/pkg/run"
	"log"
	"os"
)

func main() {
	tfRunType := os.Getenv("tf_run_type")
	builder, err := run.GetBuilder(tfRunType)
	if err != nil {
		log.Println(err)
	}
	ctx := context.Background()
	builder.GetRunBuilder(ctx)
}
