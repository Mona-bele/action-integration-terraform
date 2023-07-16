package run_test

import (
	"context"
	"github.com/Mona-bele/action-integration-terraform/pkg/run"
	"github.com/Mona-bele/action-integration-terraform/pkg/utils"
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

func TestVariableRun_Read(t *testing.T) {

	config, err := utils.LoadEnv(".")
	if err != nil {
		t.Fatal(err)
	}

	setting := run.NewSetting(config)
	if err != nil {
		t.Fatal(err)
	}

	client := run.NewNetHttpAdapter(config)

	workspace := run.NewWorkspaceRun(client, setting)

	vR := run.NewVariableRun(client, config, workspace)

	ctx := context.Background()

	_, err = vR.Read(ctx)
	if err != nil {
		t.Errorf("vR.Read() error = %v", err)
	}
}

func TestVariableRun_Create(t *testing.T) {
	config, err := utils.LoadEnv(".")
	if err != nil {
		t.Fatal(err)
	}

	setting := run.NewSetting(config)
	if err != nil {
		t.Fatal(err)
	}

	client := run.NewNetHttpAdapter(config)

	workspace := run.NewWorkspaceRun(client, setting)

	vR := run.NewVariableRun(client, config, workspace)

	ctx := context.Background()

	err = vR.Create(ctx)
	if err != nil {
		t.Errorf("vR.Create() error = %v", err)
	}
}

func TestVariableRun_Update(t *testing.T) {
	config, err := utils.LoadEnv(".")
	if err != nil {
		t.Fatal(err)
	}

	setting := run.NewSetting(config)
	if err != nil {
		t.Fatal(err)
	}

	client := run.NewNetHttpAdapter(config)

	workspace := run.NewWorkspaceRun(client, setting)

	vR := run.NewVariableRun(client, config, workspace)

	ctx := context.Background()

	err = vR.Update(ctx)
	if err != nil {
		t.Errorf("vR.Update() error = %v", err)
	}
}

func TestVariableRun_Delete(t *testing.T) {

	config, err := utils.LoadEnv(".")
	if err != nil {
		t.Fatal(err)
	}

	setting := run.NewSetting(config)
	if err != nil {
		t.Fatal(err)
	}

	client := run.NewNetHttpAdapter(config)

	workspace := run.NewWorkspaceRun(client, setting)

	vR := run.NewVariableRun(client, config, workspace)

	ctx := context.Background()

	err = vR.Delete(ctx)
	if err != nil {
		t.Errorf("error = %v", err)
	}
}
