package run_test

import (
	"context"
	"github.com/Mona-bele/action-integration-terraform/pkg/run"
	"log"
	"os"
	"testing"
)

func init() {
	log.Println("TF_API_TOKEN-TO", os.Getenv("TF_API_TOKEN"))
	os.Setenv("tf_api_token", os.Getenv("TF_API_TOKEN"))
	os.Setenv("tf_organization", os.Getenv("TF_ORGANIZATION"))
	os.Setenv("tf_workspace", os.Getenv("TF_WORKSPACE"))
	os.Setenv("tf_run_type", "teste")

	os.Setenv("variable_type", "vars")
	os.Setenv("variable_key", "some_key")
	os.Setenv("variable_value", "some_value")
	os.Setenv("variable_description", "some description")
	os.Setenv("variable_category", "terraform")
	os.Setenv("variable_hcl", "false")
	os.Setenv("variable_sensitive", "false")
}

func TestVariableRunBuilder_GetRunBuilder_Create(t *testing.T) {
	builder, err := run.GetBuilder("variable_create")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	builder.GetRunBuilder(ctx)
}

func TestVariableRunBuilder_GetRunBuilder_Delete(t *testing.T) {
	builder, err := run.GetBuilder("variable_delete")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	builder.GetRunBuilder(ctx)
}
