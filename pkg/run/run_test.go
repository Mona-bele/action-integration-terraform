package run_test

import (
	"context"
	"github.com/Mona-bele/action-integration-terraform/pkg/run"
	"os"
	"testing"
)

func init() {
	os.Setenv("INPUT_VARIABLE_TYPE", "vars")
	os.Setenv("INPUT_TF_VARIABLE_KEY", "some_key")
	os.Setenv("INPUT_TF_VARIABLE_VALUE", "some_value")
	os.Setenv("INPUT_TF_VARIABLE_DESCRIPTION", "some description")
	os.Setenv("INPUT_TF_VARIABLE_CATEGORY", "terraform")
	os.Setenv("INPUT_TF_VARIABLE_HCL", "false")
	os.Setenv("INPUT_TF_VARIABLE_SENSITIVE", "false")
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
