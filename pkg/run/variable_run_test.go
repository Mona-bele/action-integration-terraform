package run_test

import (
	"context"
	"github.com/Mona-bele/action-integration-terraform/pkg/run"
	"github.com/Mona-bele/action-integration-terraform/pkg/utils"
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
