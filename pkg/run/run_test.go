package run_test

import (
	"context"
	"github.com/Mona-bele/action-integration-terraform/pkg/run"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func init() {
	_ = godotenv.Load()
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
